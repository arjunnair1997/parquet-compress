package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	kzstd "github.com/klauspost/compress/zstd"
	"github.com/parquet-go/parquet-go"
	"github.com/parquet-go/parquet-go/compress"
	"github.com/parquet-go/parquet-go/compress/snappy"
	parquetzstd "github.com/parquet-go/parquet-go/compress/zstd"
	"github.com/parquet-go/parquet-go/encoding"
)

const (
	defaultInput              = "data/clickbench/hits.tsv.gz"
	defaultOutputDir          = "clickbench parquet files"
	defaultResultsDir         = "experiment results"
	defaultRows               = int64(1_000_000)
	defaultPageSize           = 256 * 1024
	defaultDictionaryPageSize = 1 * 1024 * 1024
)

type columnKind int

const (
	kindInt16 columnKind = iota
	kindInt32
	kindInt64
	kindDate
	kindTimestampMillis
	kindString
)

func (k columnKind) String() string {
	switch k {
	case kindInt16:
		return "int16"
	case kindInt32:
		return "int32"
	case kindInt64:
		return "int64"
	case kindDate:
		return "date"
	case kindTimestampMillis:
		return "timestamp_millis"
	case kindString:
		return "string"
	default:
		return "unknown"
	}
}

type columnSpec struct {
	Name string
	Kind columnKind
}

type config struct {
	Input             string
	OutputDir         string
	ResultsDir        string
	TSVDir            string
	Rows              int64
	MaxPageSize       int64
	MaxDictPageSize   int64
	MaxRowGroupRows   int64
	MaxRowGroupSize   int64
	MaxFileSize       int64
	Encoding          string
	IntEncoding       string
	StringEncoding    string
	DateEncoding      string
	TimestampEncoding string
	Compression       string
	ZstdLevel         int
	Verify            bool
	VerifyOnly        bool
	ComparePlainBase  bool
	BaselineOutputDir string
}

type runStats struct {
	InputBytes      int64
	Rows            int64
	Files           []fileStat
	Columns         []columnStat
	StartedAt       time.Time
	FinishedAt      time.Time
	ResultPath      string
	ColumnStatsPath string
	OutputDir       string
	SchemaColumns   int
	CompressionName string
	EncodingByGroup map[string]string
	Verification    *verifyStats
}

type verifyStats struct {
	Rows        int64
	Files       int
	Elapsed     time.Duration
	SourceBytes int64
}

type fileStat struct {
	Path               string
	Rows               int64
	Size               int64
	PhysicalSize       int64
	EncodedSize        int64
	CompressedDataSize int64
}

type columnStat struct {
	Name               string
	Kind               string
	ConfiguredEncoding string
	SourceBytes        int64
	PhysicalBytes      int64
	EncodedBytes       int64
	CompressedBytes    int64
	Values             int64
	MetadataEncodings  map[string]struct{}
	PageEncodingStats  map[string]int64
}

func main() {
	cfg, err := parseFlags(os.Args[1:])
	if err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return
		}
		exitf("%v", err)
	}
	if cfg.VerifyOnly {
		if err := runVerifyOnly(cfg); err != nil {
			exitf("%v", err)
		}
		return
	}
	if cfg.ComparePlainBase {
		if err := runWithPlainBaseline(cfg); err != nil {
			exitf("%v", err)
		}
		return
	}
	if _, err := run(cfg); err != nil {
		exitf("%v", err)
	}
}

func parseFlags(args []string) (config, error) {
	var cfg config
	fs := flag.NewFlagSet("clickbench-parquet", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	cfg.MaxPageSize = defaultPageSize
	cfg.MaxDictPageSize = defaultDictionaryPageSize
	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "Usage of %s:\n", fs.Name())
		fs.PrintDefaults()
		fmt.Fprintln(fs.Output())
		fmt.Fprintln(fs.Output(), "Encoding choices by type group:")
		fmt.Fprintln(fs.Output(), "  int/date/timestamp: plain, rle-dict")
		fmt.Fprintln(fs.Output(), "  string:             plain, rle-dict, delta-byte-array, delta-length-byte-array")
		fmt.Fprintln(fs.Output())
		fmt.Fprintln(fs.Output(), "Aliases accepted:")
		fmt.Fprintln(fs.Output(), "  rle-dictionary, dict -> rle-dict")
		fmt.Fprintln(fs.Output(), "  delta-bytearray      -> delta-byte-array")
		fmt.Fprintln(fs.Output(), "  delta-length-bytearray -> delta-length-byte-array")
	}

	fs.StringVar(&cfg.Input, "input", defaultInput, "path to hits.tsv or hits.tsv.gz")
	fs.StringVar(&cfg.OutputDir, "output-dir", defaultOutputDir, "directory for generated parquet part files")
	fs.StringVar(&cfg.ResultsDir, "results-dir", defaultResultsDir, "directory for markdown experiment result files")
	fs.StringVar(&cfg.TSVDir, "tsv-dir", "", "directory for TSV result files; defaults to a tsvs directory near --results-dir")
	fs.Int64Var(&cfg.Rows, "rows", defaultRows, "number of input rows to write")
	fs.Var(sizeFlag{&cfg.MaxPageSize}, "max-page-size", "target parquet page buffer size, e.g. 256KiB, 1MiB")
	fs.Var(sizeFlag{&cfg.MaxDictPageSize}, "max-dictionary-page-size", "maximum per-column dictionary bytes before falling back to plain encoding; 0 disables the cap")
	fs.Int64Var(&cfg.MaxRowGroupRows, "max-row-group-rows", 0, "approximate maximum rows per row group; 0 disables the row-count limit")
	fs.Var(sizeFlag{&cfg.MaxRowGroupSize}, "max-row-group-size", "approximate row group byte-size threshold; 0 disables the byte-size limit")
	fs.Var(sizeFlag{&cfg.MaxFileSize}, "max-file-size", "approximate parquet file byte-size threshold; 0 writes one file")
	fs.StringVar(&cfg.Encoding, "encoding", "plain", "default encoding for all type groups; must be valid for every group unless overridden")
	fs.StringVar(&cfg.IntEncoding, "int-encoding", "", "encoding for integer columns: plain, rle-dict; defaults to --encoding")
	fs.StringVar(&cfg.StringEncoding, "string-encoding", "", "encoding for string columns: plain, rle-dict, delta-byte-array, delta-length-byte-array; defaults to --encoding")
	fs.StringVar(&cfg.DateEncoding, "date-encoding", "", "encoding for date columns: plain, rle-dict; defaults to --encoding")
	fs.StringVar(&cfg.TimestampEncoding, "timestamp-encoding", "", "encoding for timestamp columns: plain, rle-dict; defaults to --encoding")
	fs.StringVar(&cfg.Compression, "compression", "none", "compression: none, snappy, zstd")
	fs.IntVar(&cfg.ZstdLevel, "zstd-level", 3, "zstd compression level when --compression=zstd")
	fs.BoolVar(&cfg.Verify, "verify", false, "read generated parquet files and compare them to parsed source rows after writing")
	fs.BoolVar(&cfg.VerifyOnly, "verify-only", false, "verify existing parquet files in --output-dir without writing new files")
	fs.BoolVar(&cfg.ComparePlainBase, "compare-plain-baseline", false, "also run a plain-encoding, uncompressed baseline with the same rows/page/row-group/file settings and write a comparison result")
	fs.StringVar(&cfg.BaselineOutputDir, "baseline-output-dir", "", "output directory for --compare-plain-baseline; defaults to --output-dir plus -plain-uncompressed-baseline")

	if err := fs.Parse(args); err != nil {
		return cfg, err
	}
	if cfg.Rows < 0 {
		return cfg, fmt.Errorf("--rows must be >= 0")
	}
	if cfg.MaxPageSize <= 0 {
		return cfg, fmt.Errorf("--max-page-size must be > 0")
	}
	if cfg.MaxDictPageSize < 0 {
		return cfg, fmt.Errorf("--max-dictionary-page-size must be >= 0")
	}
	if cfg.MaxRowGroupRows < 0 {
		return cfg, fmt.Errorf("--max-row-group-rows must be >= 0")
	}
	if cfg.MaxRowGroupSize < 0 {
		return cfg, fmt.Errorf("--max-row-group-size must be >= 0")
	}
	if cfg.MaxFileSize < 0 {
		return cfg, fmt.Errorf("--max-file-size must be >= 0")
	}
	cfg.Encoding = normalizeEncodingName(cfg.Encoding)
	if cfg.IntEncoding == "" {
		cfg.IntEncoding = cfg.Encoding
	}
	if cfg.StringEncoding == "" {
		cfg.StringEncoding = cfg.Encoding
	}
	if cfg.DateEncoding == "" {
		cfg.DateEncoding = cfg.Encoding
	}
	if cfg.TimestampEncoding == "" {
		cfg.TimestampEncoding = cfg.Encoding
	}
	cfg.IntEncoding = normalizeEncodingName(cfg.IntEncoding)
	cfg.StringEncoding = normalizeEncodingName(cfg.StringEncoding)
	cfg.DateEncoding = normalizeEncodingName(cfg.DateEncoding)
	cfg.TimestampEncoding = normalizeEncodingName(cfg.TimestampEncoding)
	cfg.Compression = normalizeCompressionName(cfg.Compression)
	if cfg.TSVDir == "" {
		cfg.TSVDir = defaultTSVDir(cfg.ResultsDir)
	}
	return cfg, validateEncodingGroups(cfg)
}

func defaultTSVDir(resultsDir string) string {
	if filepath.Base(resultsDir) == "results" {
		return filepath.Join(filepath.Dir(resultsDir), "tsvs")
	}
	return filepath.Join(resultsDir, "tsvs")
}

func runVerifyOnly(cfg config) error {
	start := time.Now()
	schema := clickBenchSchema()
	stats, err := verifyOutput(cfg, schema)
	if err != nil {
		return err
	}
	fmt.Printf("verified %d rows from %d parquet file(s)\n", stats.Rows, stats.Files)
	fmt.Printf("elapsed: %s\n", stats.Elapsed.Round(time.Millisecond))
	fmt.Printf("source TSV bytes checked: %d (%s)\n", stats.SourceBytes, formatBytes(stats.SourceBytes))
	fmt.Printf("output dir: %s\n", cfg.OutputDir)
	fmt.Printf("started: %s\n", start.Format(time.RFC3339))
	return nil
}

func run(cfg config) (runStats, error) {
	started := time.Now()
	schema := clickBenchSchema()
	columnIndexes, err := columnIndexesByTSVPosition(schema)
	if err != nil {
		return runStats{}, err
	}
	writerOptions, compressionName, encodingByGroup, err := writerOptions(cfg, schema)
	if err != nil {
		return runStats{}, err
	}

	if err := os.MkdirAll(cfg.OutputDir, 0o755); err != nil {
		return runStats{}, err
	}
	if err := clearPartFiles(cfg.OutputDir); err != nil {
		return runStats{}, err
	}
	if err := os.MkdirAll(cfg.ResultsDir, 0o755); err != nil {
		return runStats{}, err
	}
	if err := os.MkdirAll(cfg.TSVDir, 0o755); err != nil {
		return runStats{}, err
	}

	input, closeInput, err := openInput(cfg.Input)
	if err != nil {
		return runStats{}, err
	}
	defer closeInput()

	reader := bufio.NewReaderSize(input, 4<<20)
	fields := make([][]byte, 0, len(clickBenchColumns))
	row := make(parquet.Row, len(schema.Columns()))
	stats := runStats{
		StartedAt:       started,
		OutputDir:       cfg.OutputDir,
		SchemaColumns:   len(schema.Columns()),
		CompressionName: compressionName,
		EncodingByGroup: encodingByGroup,
		Columns:         newColumnStats(cfg),
	}

	var (
		currentFile              *os.File
		currentWriter            *parquet.Writer
		currentFileRows          int64
		currentFilePhysicalBytes int64
		rowGroupRows             int64
		rowGroupStart            int64
		partIndex                int
		longLineScratch          []byte
	)

	closePart := func() error {
		if currentWriter == nil {
			return nil
		}
		if err := currentWriter.Close(); err != nil {
			currentWriter = nil
			_ = currentFile.Close()
			currentFile = nil
			return err
		}
		if err := currentFile.Close(); err != nil {
			currentWriter = nil
			currentFile = nil
			return err
		}
		fileStats, err := collectParquetFileStats(currentFile.Name(), currentFileRows, stats.Columns)
		if err != nil {
			currentWriter = nil
			currentFile = nil
			return err
		}
		fileStats.PhysicalSize = currentFilePhysicalBytes
		stats.Files = append(stats.Files, fileStats)
		currentWriter = nil
		currentFile = nil
		currentFileRows = 0
		currentFilePhysicalBytes = 0
		rowGroupRows = 0
		rowGroupStart = 0
		return nil
	}
	defer closePart()

	openPart := func() error {
		if currentWriter != nil {
			return nil
		}
		path := filepath.Join(cfg.OutputDir, fmt.Sprintf("part-%05d.parquet", partIndex))
		partIndex++
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		currentFile = f
		currentWriter = parquet.NewWriter(f, writerOptions...)
		currentFileRows = 0
		currentFilePhysicalBytes = 0
		rowGroupRows = 0
		rowGroupStart = currentWriter.Size()
		return nil
	}

	for cfg.Rows == 0 || stats.Rows < cfg.Rows {
		line, err := readLine(reader, &longLineScratch)
		if err != nil && !(errors.Is(err, io.EOF) && len(line) > 0) {
			if errors.Is(err, io.EOF) {
				break
			}
			return runStats{}, err
		}
		stats.InputBytes += int64(len(line))
		line = trimLineEnding(line)
		fields = splitTabs(line, fields)
		if len(fields) != len(clickBenchColumns) {
			return runStats{}, fmt.Errorf("row %d has %d fields, expected %d", stats.Rows+1, len(fields), len(clickBenchColumns))
		}
		for i, field := range fields {
			stats.Columns[i].SourceBytes += int64(len(field))
		}
		physicalRowBytes, err := buildRow(row, fields, columnIndexes, stats.Columns)
		if err != nil {
			return runStats{}, fmt.Errorf("row %d: %w", stats.Rows+1, err)
		}
		if err := openPart(); err != nil {
			return runStats{}, err
		}
		if _, err := currentWriter.WriteRows([]parquet.Row{row}); err != nil {
			return runStats{}, fmt.Errorf("write row %d: %w", stats.Rows+1, err)
		}

		stats.Rows++
		currentFileRows++
		currentFilePhysicalBytes += physicalRowBytes
		rowGroupRows++

		if cfg.MaxRowGroupRows > 0 && rowGroupRows >= cfg.MaxRowGroupRows {
			if err := currentWriter.Flush(); err != nil {
				return runStats{}, err
			}
			rowGroupRows = 0
			rowGroupStart = currentWriter.Size()
		} else if cfg.MaxRowGroupSize > 0 && currentWriter.Size()-rowGroupStart >= cfg.MaxRowGroupSize {
			if err := currentWriter.Flush(); err != nil {
				return runStats{}, err
			}
			rowGroupRows = 0
			rowGroupStart = currentWriter.Size()
		}

		if cfg.MaxFileSize > 0 && currentWriter.Size() >= cfg.MaxFileSize {
			if err := closePart(); err != nil {
				return runStats{}, err
			}
		}
	}

	if err := closePart(); err != nil {
		return runStats{}, err
	}
	stats.FinishedAt = time.Now()
	if stats.Rows == 0 {
		return runStats{}, fmt.Errorf("no rows were written")
	}
	if cfg.Verify {
		verifyStats, err := verifyOutput(cfg, schema)
		if err != nil {
			return runStats{}, err
		}
		stats.Verification = verifyStats
	}

	if err := writeResultFile(cfg, &stats); err != nil {
		return runStats{}, err
	}
	printSummary(cfg, stats)
	return stats, nil
}

func runWithPlainBaseline(cfg config) error {
	candidateStats, err := run(cfg)
	if err != nil {
		return err
	}

	baselineCfg := plainBaselineConfig(cfg)
	baselineStats, err := run(baselineCfg)
	if err != nil {
		return err
	}

	comparisonPath, err := writeComparisonResultFile(cfg, candidateStats, baselineCfg, baselineStats)
	if err != nil {
		return err
	}
	printComparisonSummary(candidateStats, baselineStats, comparisonPath)
	return nil
}

func plainBaselineConfig(cfg config) config {
	baseline := cfg
	baseline.ComparePlainBase = false
	if baseline.BaselineOutputDir != "" {
		baseline.OutputDir = baseline.BaselineOutputDir
	} else {
		baseline.OutputDir = cfg.OutputDir + "-plain-uncompressed-baseline"
	}
	baseline.Encoding = "plain"
	baseline.IntEncoding = "plain"
	baseline.StringEncoding = "plain"
	baseline.DateEncoding = "plain"
	baseline.TimestampEncoding = "plain"
	baseline.Compression = "none"
	return baseline
}

func writerOptions(cfg config, schema *parquet.Schema) ([]parquet.WriterOption, string, map[string]string, error) {
	compressionCodec, compressionName, err := compressionCodec(cfg)
	if err != nil {
		return nil, "", nil, err
	}
	intEncoding, err := encodingByName(cfg.IntEncoding)
	if err != nil {
		return nil, "", nil, err
	}
	stringEncoding, err := encodingByName(cfg.StringEncoding)
	if err != nil {
		return nil, "", nil, err
	}
	dateEncoding, err := encodingByName(cfg.DateEncoding)
	if err != nil {
		return nil, "", nil, err
	}
	timestampEncoding, err := encodingByName(cfg.TimestampEncoding)
	if err != nil {
		return nil, "", nil, err
	}

	opts := []parquet.WriterOption{
		schema,
		parquet.PageBufferSize(int(cfg.MaxPageSize)),
		parquet.DictionaryMaxBytes(cfg.MaxDictPageSize),
		parquet.Compression(compressionCodec),
		parquet.DefaultEncodingFor(parquet.Boolean, &parquet.Plain),
		parquet.DefaultEncodingFor(parquet.Int32, intEncoding),
		parquet.DefaultEncodingFor(parquet.Int64, intEncoding),
		parquet.DefaultEncodingFor(parquet.Int96, &parquet.Plain),
		parquet.DefaultEncodingFor(parquet.Float, &parquet.Plain),
		parquet.DefaultEncodingFor(parquet.Double, &parquet.Plain),
		parquet.DefaultEncodingFor(parquet.ByteArray, stringEncoding),
		parquet.DefaultEncodingFor(parquet.FixedLenByteArray, &parquet.Plain),
	}

	// Dates are physically INT32 and timestamps are physically INT64. If either
	// group differs from normal integer encoding, pin those columns directly.
	if cfg.DateEncoding != cfg.IntEncoding || cfg.TimestampEncoding != cfg.IntEncoding {
		opts[0] = clickBenchSchemaWithSpecialEncodings(dateEncoding, timestampEncoding, cfg.DateEncoding != cfg.IntEncoding, cfg.TimestampEncoding != cfg.IntEncoding)
	}

	return opts, compressionName, map[string]string{
		"int":       cfg.IntEncoding,
		"string":    cfg.StringEncoding,
		"date":      cfg.DateEncoding,
		"timestamp": cfg.TimestampEncoding,
	}, nil
}

func clearPartFiles(dir string) error {
	paths, err := filepath.Glob(filepath.Join(dir, "part-*.parquet"))
	if err != nil {
		return err
	}
	for _, path := range paths {
		if err := os.Remove(path); err != nil {
			return err
		}
	}
	return nil
}

func newColumnStats(cfg config) []columnStat {
	columns := make([]columnStat, len(clickBenchColumns))
	for i, col := range clickBenchColumns {
		columns[i] = columnStat{
			Name:               col.Name,
			Kind:               col.Kind.String(),
			ConfiguredEncoding: configuredEncodingForKind(cfg, col.Kind),
			MetadataEncodings:  make(map[string]struct{}),
			PageEncodingStats:  make(map[string]int64),
		}
	}
	return columns
}

func configuredEncodingForKind(cfg config, kind columnKind) string {
	switch kind {
	case kindInt16, kindInt32, kindInt64:
		return cfg.IntEncoding
	case kindDate:
		return cfg.DateEncoding
	case kindTimestampMillis:
		return cfg.TimestampEncoding
	case kindString:
		return cfg.StringEncoding
	default:
		return cfg.Encoding
	}
}

func collectParquetFileStats(path string, fallbackRows int64, columns []columnStat) (fileStat, error) {
	info, err := os.Stat(path)
	if err != nil {
		return fileStat{}, err
	}
	f, err := os.Open(path)
	if err != nil {
		return fileStat{}, err
	}
	defer f.Close()

	pf, err := parquet.OpenFile(f, info.Size(), parquet.SkipPageIndex(true), parquet.SkipBloomFilters(true))
	if err != nil {
		return fileStat{}, err
	}

	stat := fileStat{
		Path: path,
		Rows: fallbackRows,
		Size: info.Size(),
	}
	var metadataRows int64
	for _, rowGroup := range pf.Metadata().RowGroups {
		metadataRows += rowGroup.NumRows
		for _, chunk := range rowGroup.Columns {
			md := chunk.MetaData
			stat.EncodedSize += md.TotalUncompressedSize
			stat.CompressedDataSize += md.TotalCompressedSize
			columnStatIndex := columnStatIndexByPath(columns, md.PathInSchema)
			if columnStatIndex < 0 {
				continue
			}
			column := &columns[columnStatIndex]
			column.EncodedBytes += md.TotalUncompressedSize
			column.CompressedBytes += md.TotalCompressedSize
			column.Values += md.NumValues
			for _, encoding := range md.Encoding {
				column.MetadataEncodings[encoding.String()] = struct{}{}
			}
			for _, pageStat := range md.EncodingStats {
				key := fmt.Sprintf("%s/%s", pageStat.PageType.String(), pageStat.Encoding.String())
				column.PageEncodingStats[key] += int64(pageStat.Count)
			}
		}
	}
	if metadataRows > 0 {
		stat.Rows = metadataRows
	}
	return stat, nil
}

func columnStatIndexByPath(columns []columnStat, path []string) int {
	name := strings.Join(path, ".")
	for i := range columns {
		if columns[i].Name == name {
			return i
		}
	}
	return -1
}

func verifyOutput(cfg config, schema *parquet.Schema) (*verifyStats, error) {
	started := time.Now()
	paths, err := parquetFiles(cfg.OutputDir)
	if err != nil {
		return nil, err
	}
	if len(paths) == 0 {
		return nil, fmt.Errorf("no parquet files found in %s", cfg.OutputDir)
	}

	columnIndexes, err := columnIndexesByTSVPosition(schema)
	if err != nil {
		return nil, err
	}
	columnNames := columnNamesByIndex(schema)
	expected, err := newExpectedRows(cfg.Input, schema, columnIndexes)
	if err != nil {
		return nil, err
	}
	defer expected.Close()

	rowBuf := make([]parquet.Row, 1024)
	var verifiedRows int64

	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		reader, err := newParquetReader(f)
		if err != nil {
			_ = f.Close()
			return nil, fmt.Errorf("%s: %w", path, err)
		}
		if !parquet.EqualNodes(schema, reader.Schema()) {
			_ = reader.Close()
			return nil, fmt.Errorf("%s: parquet schema differs from expected ClickBench schema", path)
		}
		for {
			n, readErr := reader.ReadRows(rowBuf)
			for i := 0; i < n; i++ {
				if cfg.Rows > 0 && verifiedRows >= cfg.Rows {
					_ = reader.Close()
					return nil, fmt.Errorf("parquet output has more than %d rows", cfg.Rows)
				}
				want, err := expected.Next()
				if err != nil {
					_ = reader.Close()
					return nil, fmt.Errorf("source ended before parquet output at row %d: %w", verifiedRows+1, err)
				}
				got := rowBuf[i]
				if !want.Equal(got) {
					_ = reader.Close()
					return nil, fmt.Errorf("row %d mismatch: %s", verifiedRows+1, rowDiff(want, got, columnNames))
				}
				verifiedRows++
			}
			if readErr != nil {
				if errors.Is(readErr, io.EOF) {
					break
				}
				_ = reader.Close()
				return nil, fmt.Errorf("%s: read rows: %w", path, readErr)
			}
		}
		if err := reader.Close(); err != nil {
			return nil, fmt.Errorf("%s: close reader: %w", path, err)
		}
	}

	if cfg.Rows > 0 && verifiedRows != cfg.Rows {
		return nil, fmt.Errorf("verified %d rows, expected %d", verifiedRows, cfg.Rows)
	}
	return &verifyStats{
		Rows:        verifiedRows,
		Files:       len(paths),
		Elapsed:     time.Since(started),
		SourceBytes: expected.InputBytes(),
	}, nil
}

func parquetFiles(dir string) ([]string, error) {
	paths, err := filepath.Glob(filepath.Join(dir, "*.parquet"))
	if err != nil {
		return nil, err
	}
	sort.Strings(paths)
	return paths, nil
}

func newParquetReader(f *os.File) (reader *parquet.Reader, err error) {
	defer func() {
		if v := recover(); v != nil {
			err = fmt.Errorf("open parquet reader: %v", v)
		}
	}()
	return parquet.NewReader(f), nil
}

type expectedRows struct {
	closer        func() error
	reader        *bufio.Reader
	scratch       []byte
	fields        [][]byte
	row           parquet.Row
	columnIndexes []int
	inputBytes    int64
	rowNumber     int64
}

func newExpectedRows(inputPath string, schema *parquet.Schema, columnIndexes []int) (*expectedRows, error) {
	input, closer, err := openInput(inputPath)
	if err != nil {
		return nil, err
	}
	return &expectedRows{
		closer:        closer,
		reader:        bufio.NewReaderSize(input, 4<<20),
		fields:        make([][]byte, 0, len(clickBenchColumns)),
		row:           make(parquet.Row, len(schema.Columns())),
		columnIndexes: columnIndexes,
	}, nil
}

func (e *expectedRows) Next() (parquet.Row, error) {
	line, err := readLine(e.reader, &e.scratch)
	if err != nil && !(errors.Is(err, io.EOF) && len(line) > 0) {
		return nil, err
	}
	e.inputBytes += int64(len(line))
	line = trimLineEnding(line)
	e.fields = splitTabs(line, e.fields)
	if len(e.fields) != len(clickBenchColumns) {
		return nil, fmt.Errorf("source row %d has %d fields, expected %d", e.rowNumber+1, len(e.fields), len(clickBenchColumns))
	}
	if _, err := buildRow(e.row, e.fields, e.columnIndexes, nil); err != nil {
		return nil, err
	}
	e.rowNumber++
	return e.row, nil
}

func (e *expectedRows) Close() error {
	return e.closer()
}

func (e *expectedRows) InputBytes() int64 {
	return e.inputBytes
}

func columnNamesByIndex(schema *parquet.Schema) []string {
	columns := schema.Columns()
	names := make([]string, len(columns))
	for i, path := range columns {
		names[i] = strings.Join(path, ".")
	}
	return names
}

func rowDiff(want, got parquet.Row, columnNames []string) string {
	if len(want) != len(got) {
		return fmt.Sprintf("row length differs: want %d values, got %d values", len(want), len(got))
	}
	for i := range want {
		if !parquet.Equal(want[i], got[i]) ||
			want[i].Column() != got[i].Column() ||
			want[i].RepetitionLevel() != got[i].RepetitionLevel() ||
			want[i].DefinitionLevel() != got[i].DefinitionLevel() {
			col := want[i].Column()
			name := fmt.Sprintf("column-%d", col)
			if col >= 0 && col < len(columnNames) {
				name = columnNames[col]
			}
			return fmt.Sprintf("%s: want %#v, got %#v", name, want[i], got[i])
		}
	}
	return "unknown difference"
}

func compressionCodec(cfg config) (compress.Codec, string, error) {
	switch cfg.Compression {
	case "none":
		return &parquet.Uncompressed, "none", nil
	case "snappy":
		return &snappy.Codec{}, "snappy", nil
	case "zstd":
		if cfg.ZstdLevel < -5 || cfg.ZstdLevel > 22 {
			return nil, "", fmt.Errorf("--zstd-level must be between -5 and 22")
		}
		return &parquetzstd.Codec{Level: kzstd.EncoderLevelFromZstd(cfg.ZstdLevel)}, fmt.Sprintf("zstd-%d", cfg.ZstdLevel), nil
	default:
		return nil, "", fmt.Errorf("unknown compression %q", cfg.Compression)
	}
}

func openInput(path string) (io.Reader, func() error, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	if strings.HasSuffix(path, ".gz") {
		gz, err := gzip.NewReader(f)
		if err != nil {
			_ = f.Close()
			return nil, nil, err
		}
		return gz, func() error {
			err1 := gz.Close()
			err2 := f.Close()
			if err1 != nil {
				return err1
			}
			return err2
		}, nil
	}
	return f, f.Close, nil
}

func readLine(r *bufio.Reader, scratch *[]byte) ([]byte, error) {
	line, err := r.ReadSlice('\n')
	if !errors.Is(err, bufio.ErrBufferFull) {
		return line, err
	}
	*scratch = append((*scratch)[:0], line...)
	for errors.Is(err, bufio.ErrBufferFull) {
		line, err = r.ReadSlice('\n')
		*scratch = append(*scratch, line...)
	}
	return *scratch, err
}

func trimLineEnding(line []byte) []byte {
	line = bytes.TrimSuffix(line, []byte{'\n'})
	line = bytes.TrimSuffix(line, []byte{'\r'})
	return line
}

func splitTabs(line []byte, fields [][]byte) [][]byte {
	fields = fields[:0]
	start := 0
	for i, b := range line {
		if b == '\t' {
			fields = append(fields, line[start:i])
			start = i + 1
		}
	}
	return append(fields, line[start:])
}

func buildRow(row parquet.Row, fields [][]byte, columnIndexes []int, columns []columnStat) (int64, error) {
	var rowPhysicalBytes int64
	for tsvIndex, field := range fields {
		spec := clickBenchColumns[tsvIndex]
		columnIndex := columnIndexes[tsvIndex]
		var value parquet.Value
		var physicalBytes int64
		switch spec.Kind {
		case kindInt16, kindInt32:
			n, err := parseInt64Bytes(field)
			if err != nil {
				return 0, fmt.Errorf("%s: %w", spec.Name, err)
			}
			value = parquet.Int32Value(int32(n))
			physicalBytes = 4
		case kindInt64:
			n, err := parseInt64Bytes(field)
			if err != nil {
				return 0, fmt.Errorf("%s: %w", spec.Name, err)
			}
			value = parquet.Int64Value(n)
			physicalBytes = 8
		case kindDate:
			days, err := parseDateDays(field)
			if err != nil {
				return 0, fmt.Errorf("%s: %w", spec.Name, err)
			}
			value = parquet.Int32Value(days)
			physicalBytes = 4
		case kindTimestampMillis:
			millis, err := parseTimestampMillis(field)
			if err != nil {
				return 0, fmt.Errorf("%s: %w", spec.Name, err)
			}
			value = parquet.Int64Value(millis)
			physicalBytes = 8
		case kindString:
			unescaped := unescapeClickHouseTSV(field)
			value = parquet.ByteArrayValue(unescaped)
			physicalBytes = int64(len(unescaped))
		default:
			return 0, fmt.Errorf("%s: unsupported column kind", spec.Name)
		}
		if columns != nil {
			columns[tsvIndex].PhysicalBytes += physicalBytes
		}
		rowPhysicalBytes += physicalBytes
		row[columnIndex] = value.Level(0, 0, columnIndex)
	}
	return rowPhysicalBytes, nil
}

func parseInt64Bytes(b []byte) (int64, error) {
	if len(b) == 0 {
		return 0, fmt.Errorf("empty integer")
	}
	neg := false
	i := 0
	if b[0] == '-' {
		neg = true
		i = 1
		if i == len(b) {
			return 0, fmt.Errorf("invalid integer %q", b)
		}
	}
	var n int64
	for ; i < len(b); i++ {
		c := b[i]
		if c < '0' || c > '9' {
			return 0, fmt.Errorf("invalid integer %q", b)
		}
		n = n*10 + int64(c-'0')
	}
	if neg {
		n = -n
	}
	return n, nil
}

func parseDateDays(b []byte) (int32, error) {
	if len(b) != len("2006-01-02") || b[4] != '-' || b[7] != '-' {
		return 0, fmt.Errorf("invalid date %q", b)
	}
	y := parseNDigits(b[0:4])
	m := parseNDigits(b[5:7])
	d := parseNDigits(b[8:10])
	if y < 0 || m < 1 || m > 12 || d < 1 || d > 31 {
		return 0, fmt.Errorf("invalid date %q", b)
	}
	return int32(daysFromCivil(y, m, d)), nil
}

func parseTimestampMillis(b []byte) (int64, error) {
	if len(b) != len("2006-01-02 15:04:05") || b[10] != ' ' || b[13] != ':' || b[16] != ':' {
		return 0, fmt.Errorf("invalid timestamp %q", b)
	}
	days, err := parseDateDays(b[:10])
	if err != nil {
		return 0, err
	}
	hh := parseNDigits(b[11:13])
	mm := parseNDigits(b[14:16])
	ss := parseNDigits(b[17:19])
	if hh < 0 || hh > 23 || mm < 0 || mm > 59 || ss < 0 || ss > 60 {
		return 0, fmt.Errorf("invalid timestamp %q", b)
	}
	return int64(days)*86_400_000 + int64(hh)*3_600_000 + int64(mm)*60_000 + int64(ss)*1_000, nil
}

func parseNDigits(b []byte) int {
	n := 0
	for _, c := range b {
		if c < '0' || c > '9' {
			return -1
		}
		n = n*10 + int(c-'0')
	}
	return n
}

func daysFromCivil(year, month, day int) int {
	if month <= 2 {
		year--
	}
	era := year / 400
	yoe := year - era*400
	mp := month
	if mp > 2 {
		mp -= 3
	} else {
		mp += 9
	}
	doy := (153*mp+2)/5 + day - 1
	doe := yoe*365 + yoe/4 - yoe/100 + doy
	return era*146097 + doe - 719468
}

func unescapeClickHouseTSV(field []byte) []byte {
	i := bytes.IndexByte(field, '\\')
	if i < 0 {
		return field
	}
	out := make([]byte, 0, len(field))
	out = append(out, field[:i]...)
	for i < len(field) {
		c := field[i]
		if c != '\\' || i+1 == len(field) {
			out = append(out, c)
			i++
			continue
		}
		i++
		switch field[i] {
		case '\\':
			out = append(out, '\\')
		case '0':
			out = append(out, 0)
		case 'a':
			out = append(out, '\a')
		case 'b':
			out = append(out, '\b')
		case 'f':
			out = append(out, '\f')
		case 'n':
			out = append(out, '\n')
		case 'r':
			out = append(out, '\r')
		case 't':
			out = append(out, '\t')
		case 'v':
			out = append(out, '\v')
		case 'x':
			if i+2 < len(field) {
				if v, ok := parseHexByte(field[i+1], field[i+2]); ok {
					out = append(out, v)
					i += 2
				} else {
					out = append(out, 'x')
				}
			} else {
				out = append(out, 'x')
			}
		default:
			out = append(out, field[i])
		}
		i++
	}
	return out
}

func parseHexByte(a, b byte) (byte, bool) {
	hi, ok := hexNibble(a)
	if !ok {
		return 0, false
	}
	lo, ok := hexNibble(b)
	if !ok {
		return 0, false
	}
	return hi<<4 | lo, true
}

func hexNibble(c byte) (byte, bool) {
	switch {
	case c >= '0' && c <= '9':
		return c - '0', true
	case c >= 'a' && c <= 'f':
		return c - 'a' + 10, true
	case c >= 'A' && c <= 'F':
		return c - 'A' + 10, true
	default:
		return 0, false
	}
}

func clickBenchSchema() *parquet.Schema {
	return clickBenchSchemaWithSpecialEncodings(nil, nil, false, false)
}

func clickBenchSchemaWithSpecialEncodings(dateEncoding, timestampEncoding encoding.Encoding, encodeDates, encodeTimestamps bool) *parquet.Schema {
	group := make(parquet.Group, len(clickBenchColumns))
	for _, col := range clickBenchColumns {
		var node parquet.Node
		switch col.Kind {
		case kindInt16:
			node = parquet.Int(16)
		case kindInt32:
			node = parquet.Int(32)
		case kindInt64:
			node = parquet.Int(64)
		case kindDate:
			node = parquet.Date()
			if encodeDates && dateEncoding != nil {
				node = parquet.Encoded(node, dateEncoding)
			}
		case kindTimestampMillis:
			node = parquet.Timestamp(parquet.Millisecond)
			if encodeTimestamps && timestampEncoding != nil {
				node = parquet.Encoded(node, timestampEncoding)
			}
		case kindString:
			node = parquet.String()
		default:
			panic("unsupported column kind")
		}
		group[col.Name] = parquet.Required(node)
	}
	return parquet.NewSchema("hits", group)
}

func columnIndexesByTSVPosition(schema *parquet.Schema) ([]int, error) {
	indexes := make([]int, len(clickBenchColumns))
	for i, col := range clickBenchColumns {
		leaf, ok := schema.Lookup(col.Name)
		if !ok {
			return nil, fmt.Errorf("schema missing column %s", col.Name)
		}
		indexes[i] = leaf.ColumnIndex
	}
	return indexes, nil
}

func validateEncodingGroups(cfg config) error {
	if !isIntEncoding(cfg.IntEncoding) {
		return fmt.Errorf("encoding %q is not allowed for int columns", cfg.IntEncoding)
	}
	if !isStringEncoding(cfg.StringEncoding) {
		return fmt.Errorf("encoding %q is not allowed for string columns", cfg.StringEncoding)
	}
	if !isIntEncoding(cfg.DateEncoding) {
		return fmt.Errorf("encoding %q is not allowed for date columns", cfg.DateEncoding)
	}
	if !isIntEncoding(cfg.TimestampEncoding) {
		return fmt.Errorf("encoding %q is not allowed for timestamp columns", cfg.TimestampEncoding)
	}
	return nil
}

func isIntEncoding(name string) bool {
	switch name {
	case "plain", "rle-dict":
		return true
	default:
		return false
	}
}

func isStringEncoding(name string) bool {
	switch name {
	case "plain", "rle-dict", "delta-byte-array", "delta-length-byte-array":
		return true
	default:
		return false
	}
}

func encodingByName(name string) (encoding.Encoding, error) {
	switch name {
	case "plain":
		return &parquet.Plain, nil
	case "rle-dict":
		return &parquet.RLEDictionary, nil
	case "delta-byte-array":
		return &parquet.DeltaByteArray, nil
	case "delta-length-byte-array":
		return &parquet.DeltaLengthByteArray, nil
	default:
		return nil, fmt.Errorf("unknown encoding %q", name)
	}
}

func normalizeEncodingName(name string) string {
	name = strings.ToLower(strings.TrimSpace(name))
	name = strings.ReplaceAll(name, "_", "-")
	name = strings.ReplaceAll(name, " ", "-")
	switch name {
	case "dict", "rle-dictionary", "rle-dict", "rledict":
		return "rle-dict"
	case "delta-bytearray", "delta-byte-array", "dba":
		return "delta-byte-array"
	case "delta-length-bytearray", "delta-length-byte-array", "dlba":
		return "delta-length-byte-array"
	default:
		return name
	}
}

func normalizeCompressionName(name string) string {
	name = strings.ToLower(strings.TrimSpace(name))
	switch name {
	case "uncompressed":
		return "none"
	default:
		return name
	}
}

type sizeFlag struct {
	dst *int64
}

func (f sizeFlag) Set(s string) error {
	n, err := parseSize(s)
	if err != nil {
		return err
	}
	*f.dst = n
	return nil
}

func (f sizeFlag) String() string {
	if f.dst == nil {
		return "0"
	}
	return formatBytes(*f.dst)
}

func parseSize(s string) (int64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("empty size")
	}
	i := 0
	for i < len(s) && ((s[i] >= '0' && s[i] <= '9') || s[i] == '.') {
		i++
	}
	numText := s[:i]
	unitText := strings.ToLower(strings.TrimSpace(s[i:]))
	if numText == "" {
		return 0, fmt.Errorf("invalid size %q", s)
	}
	n, err := strconv.ParseFloat(numText, 64)
	if err != nil {
		return 0, err
	}
	mult := float64(1)
	switch unitText {
	case "", "b":
	case "k", "kb", "kib":
		mult = 1024
	case "m", "mb", "mib":
		mult = 1024 * 1024
	case "g", "gb", "gib":
		mult = 1024 * 1024 * 1024
	default:
		return 0, fmt.Errorf("unknown size unit %q", unitText)
	}
	return int64(n * mult), nil
}

func formatBytes(n int64) string {
	const unit = 1024
	if n < unit {
		return fmt.Sprintf("%d B", n)
	}
	value := float64(n)
	for _, suffix := range []string{"KiB", "MiB", "GiB", "TiB"} {
		value /= unit
		if value < unit {
			return fmt.Sprintf("%.2f %s", value, suffix)
		}
	}
	return fmt.Sprintf("%.2f PiB", value/unit)
}

func writeResultFile(cfg config, stats *runStats) error {
	baseName := fmt.Sprintf("%s_%s", stats.StartedAt.Format("2006-01-02"), experimentDescription(cfg, *stats))
	resultPath := filepath.Join(cfg.ResultsDir, baseName+".md")
	columnStatsPath := filepath.Join(cfg.TSVDir, baseName+"_columns.tsv")

	stats.ResultPath = resultPath
	stats.ColumnStatsPath = columnStatsPath

	if err := writeColumnStatsTSV(columnStatsPath, stats.Columns); err != nil {
		return err
	}

	var b strings.Builder
	writeMarkdownSummary(&b, cfg, *stats)
	return os.WriteFile(resultPath, []byte(b.String()), 0o644)
}

func writeComparisonResultFile(candidateCfg config, candidate runStats, baselineCfg config, baseline runStats) (string, error) {
	baseName := fmt.Sprintf("%s_%s", candidate.StartedAt.Format("2006-01-02"), comparisonDescription(candidateCfg, candidate))
	path := filepath.Join(candidateCfg.ResultsDir, baseName+".md")
	columnComparisonPath := filepath.Join(candidateCfg.TSVDir, baseName+"_columns.tsv")
	if err := writeColumnComparisonTSV(columnComparisonPath, candidate.Columns, baseline.Columns); err != nil {
		return "", err
	}
	var b strings.Builder
	writeComparisonMarkdown(&b, candidateCfg, candidate, baselineCfg, baseline, path, columnComparisonPath)
	return path, os.WriteFile(path, []byte(b.String()), 0o644)
}

func experimentDescription(cfg config, stats runStats) string {
	parts := []string{
		fmt.Sprintf("rows-%d", stats.Rows),
		fmt.Sprintf("comp-%s", stats.CompressionName),
		fmt.Sprintf("int-%s", cfg.IntEncoding),
		fmt.Sprintf("str-%s", cfg.StringEncoding),
		fmt.Sprintf("date-%s", cfg.DateEncoding),
		fmt.Sprintf("ts-%s", cfg.TimestampEncoding),
	}
	return sanitizeFilename(strings.Join(parts, "_"))
}

func comparisonDescription(cfg config, stats runStats) string {
	parts := []string{"compare-plain-uncompressed", experimentDescription(cfg, stats)}
	return sanitizeFilename(strings.Join(parts, "_"))
}

func sanitizeFilename(s string) string {
	s = strings.ToLower(s)
	var b strings.Builder
	lastDash := false
	for _, r := range s {
		ok := (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9')
		if ok {
			b.WriteRune(r)
			lastDash = false
			continue
		}
		if !lastDash {
			b.WriteByte('-')
			lastDash = true
		}
	}
	return strings.Trim(b.String(), "-")
}

func writeMarkdownSummary(b *strings.Builder, cfg config, stats runStats) {
	parquetBytes := totalParquetBytes(stats.Files)
	physicalBytes := totalPhysicalBytes(stats.Files)
	encodedBytes := totalEncodedBytes(stats.Files)
	compressedDataBytes := totalCompressedDataBytes(stats.Files)
	elapsed := stats.FinishedAt.Sub(stats.StartedAt)
	fmt.Fprintf(b, "# ClickBench Parquet Experiment\n\n")
	fmt.Fprintf(b, "- Started: `%s`\n", stats.StartedAt.Format(time.RFC3339))
	fmt.Fprintf(b, "- Write elapsed: `%s`\n", elapsed.Round(time.Millisecond))
	fmt.Fprintf(b, "- Input: `%s`\n", cfg.Input)
	fmt.Fprintf(b, "- Output directory: `%s`\n", stats.OutputDir)
	fmt.Fprintf(b, "- Rows: `%d`\n", stats.Rows)
	fmt.Fprintf(b, "- Source TSV bytes for rows, reference only: `%d` (%s)\n", stats.InputBytes, formatBytes(stats.InputBytes))
	fmt.Fprintf(b, "- Parquet physical bytes before page encoding: `%d` (%s)\n", physicalBytes, formatBytes(physicalBytes))
	fmt.Fprintf(b, "- Encoded column bytes before codec compression: `%d` (%s)\n", encodedBytes, formatBytes(encodedBytes))
	fmt.Fprintf(b, "- Compressed column data bytes after codec compression: `%d` (%s)\n", compressedDataBytes, formatBytes(compressedDataBytes))
	fmt.Fprintf(b, "- Parquet file bytes: `%d` (%s)\n", parquetBytes, formatBytes(parquetBytes))
	fmt.Fprintf(b, "- Physical/encoded ratio: `%s`\n", formatMultiplier(physicalBytes, encodedBytes))
	fmt.Fprintf(b, "- Encoded/compressed-data ratio: `%s`\n", formatMultiplier(encodedBytes, compressedDataBytes))
	fmt.Fprintf(b, "- Physical/compressed-data ratio: `%s`\n", formatMultiplier(physicalBytes, compressedDataBytes))
	fmt.Fprintf(b, "- Physical/parquet-file ratio: `%s`\n", formatMultiplier(physicalBytes, parquetBytes))
	fmt.Fprintf(b, "- Files: `%d`\n\n", len(stats.Files))
	fmt.Fprintf(b, "## Settings\n\n")
	fmt.Fprintf(b, "- Compression: `%s`\n", stats.CompressionName)
	fmt.Fprintf(b, "- Int encoding: `%s`\n", cfg.IntEncoding)
	fmt.Fprintf(b, "- String encoding: `%s`\n", cfg.StringEncoding)
	fmt.Fprintf(b, "- Date encoding: `%s`\n", cfg.DateEncoding)
	fmt.Fprintf(b, "- Timestamp encoding: `%s`\n", cfg.TimestampEncoding)
	fmt.Fprintf(b, "- Max page size: `%s`\n", formatBytes(cfg.MaxPageSize))
	fmt.Fprintf(b, "- Max dictionary page size: `%s`\n", formatBytes(cfg.MaxDictPageSize))
	fmt.Fprintf(b, "- Max row group rows: `%d`\n", cfg.MaxRowGroupRows)
	fmt.Fprintf(b, "- Max row group size: `%s`\n", formatBytes(cfg.MaxRowGroupSize))
	fmt.Fprintf(b, "- Max file size: `%s`\n\n", formatBytes(cfg.MaxFileSize))
	writeSchemaMarkdown(b)
	if stats.Verification != nil {
		fmt.Fprintf(b, "## Verification\n\n")
		fmt.Fprintf(b, "- Status: `passed`\n")
		fmt.Fprintf(b, "- Rows read and compared: `%d`\n", stats.Verification.Rows)
		fmt.Fprintf(b, "- Files read: `%d`\n", stats.Verification.Files)
		fmt.Fprintf(b, "- Elapsed: `%s`\n", stats.Verification.Elapsed.Round(time.Millisecond))
		fmt.Fprintf(b, "- Source TSV bytes checked: `%d` (%s)\n\n", stats.Verification.SourceBytes, formatBytes(stats.Verification.SourceBytes))
	}
	writeColumnStatsMarkdown(b, stats.Columns, stats.ResultPath, stats.ColumnStatsPath)
	fmt.Fprintf(b, "## Files\n\n")
	for _, f := range stats.Files {
		fmt.Fprintf(b, "- `%s`: `%d` rows, `%d` file bytes (%s), `%d` physical bytes (%s), `%d` encoded bytes (%s), `%d` compressed data bytes (%s)\n",
			f.Path, f.Rows,
			f.Size, formatBytes(f.Size),
			f.PhysicalSize, formatBytes(f.PhysicalSize),
			f.EncodedSize, formatBytes(f.EncodedSize),
			f.CompressedDataSize, formatBytes(f.CompressedDataSize),
		)
	}
}

func printSummary(cfg config, stats runStats) {
	parquetBytes := totalParquetBytes(stats.Files)
	physicalBytes := totalPhysicalBytes(stats.Files)
	encodedBytes := totalEncodedBytes(stats.Files)
	compressedDataBytes := totalCompressedDataBytes(stats.Files)
	elapsed := stats.FinishedAt.Sub(stats.StartedAt)
	fmt.Printf("wrote %d rows into %d parquet file(s)\n", stats.Rows, len(stats.Files))
	fmt.Printf("write elapsed: %s\n", elapsed.Round(time.Millisecond))
	fmt.Printf("source TSV bytes for rows, reference only: %d (%s)\n", stats.InputBytes, formatBytes(stats.InputBytes))
	fmt.Printf("parquet physical bytes before page encoding: %d (%s)\n", physicalBytes, formatBytes(physicalBytes))
	fmt.Printf("encoded column bytes before codec compression: %d (%s)\n", encodedBytes, formatBytes(encodedBytes))
	fmt.Printf("compressed column data bytes after codec compression: %d (%s)\n", compressedDataBytes, formatBytes(compressedDataBytes))
	fmt.Printf("parquet file bytes: %d (%s)\n", parquetBytes, formatBytes(parquetBytes))
	fmt.Printf("physical/encoded ratio: %s\n", formatMultiplier(physicalBytes, encodedBytes))
	fmt.Printf("encoded/compressed-data ratio: %s\n", formatMultiplier(encodedBytes, compressedDataBytes))
	fmt.Printf("physical/compressed-data ratio: %s\n", formatMultiplier(physicalBytes, compressedDataBytes))
	fmt.Printf("physical/parquet-file ratio: %s\n", formatMultiplier(physicalBytes, parquetBytes))
	fmt.Printf("max dictionary page size: %s\n", formatBytes(cfg.MaxDictPageSize))
	fmt.Printf("output dir: %s\n", cfg.OutputDir)
	fmt.Printf("result file: %s\n", stats.ResultPath)
	if stats.Verification != nil {
		fmt.Printf("verification: passed, %d rows from %d file(s), %s\n", stats.Verification.Rows, stats.Verification.Files, stats.Verification.Elapsed.Round(time.Millisecond))
	}
	fmt.Println("columns:")
	for _, c := range stats.Columns {
		fmt.Printf("  %s: type=%s configured=%s metadata=%s physical=%s encoded=%s compressed=%s physical/encoded=%s encoded/compressed=%s physical/compressed=%s source-reference=%s\n",
			c.Name,
			c.Kind,
			c.ConfiguredEncoding,
			stringSet(c.MetadataEncodings),
			formatBytes(c.PhysicalBytes),
			formatBytes(c.EncodedBytes),
			formatBytes(c.CompressedBytes),
			formatMultiplier(c.PhysicalBytes, c.EncodedBytes),
			formatMultiplier(c.EncodedBytes, c.CompressedBytes),
			formatMultiplier(c.PhysicalBytes, c.CompressedBytes),
			formatBytes(c.SourceBytes),
		)
	}
	for _, f := range stats.Files {
		fmt.Printf("  %s: %d rows, file=%s, physical=%s, encoded=%s, compressed-data=%s\n",
			f.Path, f.Rows, formatBytes(f.Size), formatBytes(f.PhysicalSize), formatBytes(f.EncodedSize), formatBytes(f.CompressedDataSize))
	}
}

func printComparisonSummary(candidate, baseline runStats, path string) {
	baselineEncoded := totalEncodedBytes(baseline.Files)
	candidateEncoded := totalEncodedBytes(candidate.Files)
	baselineCompressed := totalCompressedDataBytes(baseline.Files)
	candidateCompressed := totalCompressedDataBytes(candidate.Files)
	baselineParquet := totalParquetBytes(baseline.Files)
	candidateParquet := totalParquetBytes(candidate.Files)

	fmt.Printf("plain baseline comparison result file: %s\n", path)
	fmt.Printf("encoded bytes baseline/candidate: %s (%s -> %s)\n",
		formatMultiplier(baselineEncoded, candidateEncoded),
		formatBytes(baselineEncoded),
		formatBytes(candidateEncoded),
	)
	fmt.Printf("compressed-data bytes baseline/candidate: %s (%s -> %s)\n",
		formatMultiplier(baselineCompressed, candidateCompressed),
		formatBytes(baselineCompressed),
		formatBytes(candidateCompressed),
	)
	fmt.Printf("parquet file bytes baseline/candidate: %s (%s -> %s)\n",
		formatMultiplier(baselineParquet, candidateParquet),
		formatBytes(baselineParquet),
		formatBytes(candidateParquet),
	)
}

func writeComparisonMarkdown(b *strings.Builder, candidateCfg config, candidate runStats, baselineCfg config, baseline runStats, comparisonPath, columnComparisonPath string) {
	baselinePhysical := totalPhysicalBytes(baseline.Files)
	candidatePhysical := totalPhysicalBytes(candidate.Files)
	baselineEncoded := totalEncodedBytes(baseline.Files)
	candidateEncoded := totalEncodedBytes(candidate.Files)
	baselineCompressed := totalCompressedDataBytes(baseline.Files)
	candidateCompressed := totalCompressedDataBytes(candidate.Files)
	baselineParquet := totalParquetBytes(baseline.Files)
	candidateParquet := totalParquetBytes(candidate.Files)

	fmt.Fprintf(b, "# ClickBench Plain-Uncompressed Baseline Comparison\n\n")
	fmt.Fprintf(b, "- Candidate result: `%s`\n", candidate.ResultPath)
	fmt.Fprintf(b, "- Plain baseline result: `%s`\n", baseline.ResultPath)
	fmt.Fprintf(b, "- Candidate output directory: `%s`\n", candidate.OutputDir)
	fmt.Fprintf(b, "- Plain baseline output directory: `%s`\n", baseline.OutputDir)
	fmt.Fprintf(b, "- Rows: candidate `%d`, baseline `%d`\n", candidate.Rows, baseline.Rows)
	fmt.Fprintf(b, "- Candidate write elapsed: `%s`\n", candidate.FinishedAt.Sub(candidate.StartedAt).Round(time.Millisecond))
	fmt.Fprintf(b, "- Plain baseline write elapsed: `%s`\n", baseline.FinishedAt.Sub(baseline.StartedAt).Round(time.Millisecond))
	fmt.Fprintf(b, "- Plain baseline definition: same rows/page/row-group/file settings, all encodings `plain`, compression `none`\n\n")
	writeSchemaMarkdown(b)

	fmt.Fprintf(b, "## Settings\n\n")
	fmt.Fprintf(b, "| Setting | Candidate | Plain baseline |\n")
	fmt.Fprintf(b, "| --- | --- | --- |\n")
	fmt.Fprintf(b, "| Compression | `%s` | `%s` |\n", candidate.CompressionName, baseline.CompressionName)
	fmt.Fprintf(b, "| Int encoding | `%s` | `%s` |\n", candidateCfg.IntEncoding, baselineCfg.IntEncoding)
	fmt.Fprintf(b, "| String encoding | `%s` | `%s` |\n", candidateCfg.StringEncoding, baselineCfg.StringEncoding)
	fmt.Fprintf(b, "| Date encoding | `%s` | `%s` |\n", candidateCfg.DateEncoding, baselineCfg.DateEncoding)
	fmt.Fprintf(b, "| Timestamp encoding | `%s` | `%s` |\n", candidateCfg.TimestampEncoding, baselineCfg.TimestampEncoding)
	fmt.Fprintf(b, "| Max page size | `%s` | `%s` |\n", formatBytes(candidateCfg.MaxPageSize), formatBytes(baselineCfg.MaxPageSize))
	fmt.Fprintf(b, "| Max dictionary page size | `%s` | `%s` |\n", formatBytes(candidateCfg.MaxDictPageSize), formatBytes(baselineCfg.MaxDictPageSize))
	fmt.Fprintf(b, "| Max row group rows | `%d` | `%d` |\n", candidateCfg.MaxRowGroupRows, baselineCfg.MaxRowGroupRows)
	fmt.Fprintf(b, "| Max row group size | `%s` | `%s` |\n", formatBytes(candidateCfg.MaxRowGroupSize), formatBytes(baselineCfg.MaxRowGroupSize))
	fmt.Fprintf(b, "| Max file size | `%s` | `%s` |\n\n", formatBytes(candidateCfg.MaxFileSize), formatBytes(baselineCfg.MaxFileSize))

	fmt.Fprintf(b, "## Totals\n\n")
	fmt.Fprintf(b, "| Metric | Plain baseline | Candidate | Baseline/candidate |\n")
	fmt.Fprintf(b, "| --- | ---: | ---: | ---: |\n")
	fmt.Fprintf(b, "| Physical bytes | `%d` (%s) | `%d` (%s) | `%s` |\n",
		baselinePhysical, formatBytes(baselinePhysical),
		candidatePhysical, formatBytes(candidatePhysical),
		formatMultiplier(baselinePhysical, candidatePhysical),
	)
	fmt.Fprintf(b, "| Encoded column bytes | `%d` (%s) | `%d` (%s) | `%s` |\n",
		baselineEncoded, formatBytes(baselineEncoded),
		candidateEncoded, formatBytes(candidateEncoded),
		formatMultiplier(baselineEncoded, candidateEncoded),
	)
	fmt.Fprintf(b, "| Compressed column data bytes | `%d` (%s) | `%d` (%s) | `%s` |\n",
		baselineCompressed, formatBytes(baselineCompressed),
		candidateCompressed, formatBytes(candidateCompressed),
		formatMultiplier(baselineCompressed, candidateCompressed),
	)
	fmt.Fprintf(b, "| Parquet file bytes | `%d` (%s) | `%d` (%s) | `%s` |\n\n",
		baselineParquet, formatBytes(baselineParquet),
		candidateParquet, formatBytes(candidateParquet),
		formatMultiplier(baselineParquet, candidateParquet),
	)

	fmt.Fprintf(b, "## Columns\n\n")
	fmt.Fprintf(b, "Encoding ratio compares plain-uncompressed baseline encoded bytes to candidate encoded bytes. Total ratio compares plain-uncompressed baseline encoded bytes to candidate compressed bytes.\n\n")
	if columnComparisonPath != "" {
		name := filepath.Base(columnComparisonPath)
		fmt.Fprintf(b, "Column comparison TSV: [%s](%s)\n\n", name, markdownLinkTarget(filepath.Dir(comparisonPath), columnComparisonPath))
	}
	fmt.Fprintf(b, "| Column | Type | Candidate encoding | Candidate metadata | Baseline encoded bytes | Candidate encoded bytes | Candidate compressed bytes | Encoding ratio | Candidate codec ratio | Total ratio |\n")
	fmt.Fprintf(b, "| --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: |\n")
	baselineColumns := columnStatsByName(baseline.Columns)
	for _, candidateColumn := range candidate.Columns {
		baselineColumn, ok := baselineColumns[candidateColumn.Name]
		if !ok {
			continue
		}
		fmt.Fprintf(b, "| `%s` | `%s` | `%s` | `%s` | `%d` (%s) | `%d` (%s) | `%d` (%s) | `%s` | `%s` | `%s` |\n",
			candidateColumn.Name,
			candidateColumn.Kind,
			candidateColumn.ConfiguredEncoding,
			stringSet(candidateColumn.MetadataEncodings),
			baselineColumn.EncodedBytes, formatBytes(baselineColumn.EncodedBytes),
			candidateColumn.EncodedBytes, formatBytes(candidateColumn.EncodedBytes),
			candidateColumn.CompressedBytes, formatBytes(candidateColumn.CompressedBytes),
			formatMultiplier(baselineColumn.EncodedBytes, candidateColumn.EncodedBytes),
			formatMultiplier(candidateColumn.EncodedBytes, candidateColumn.CompressedBytes),
			formatMultiplier(baselineColumn.EncodedBytes, candidateColumn.CompressedBytes),
		)
	}
	fmt.Fprintf(b, "\n")
}

func columnStatsByName(columns []columnStat) map[string]columnStat {
	byName := make(map[string]columnStat, len(columns))
	for _, column := range columns {
		byName[column.Name] = column
	}
	return byName
}

func writeSchemaMarkdown(b *strings.Builder) {
	fmt.Fprintf(b, "## Schema\n\n")
	fmt.Fprintf(b, "- Columns: `%d`, generated from the built-in ClickBench `hits` column list in source TSV field order.\n", len(clickBenchColumns))
	fmt.Fprintf(b, "- Mapping: each input row is split on tabs, and field `N` is written to ClickBench column `N` with the same name.\n")
	fmt.Fprintf(b, "- All Parquet columns are required.\n")
	fmt.Fprintf(b, "- String fields are ClickHouse TSV-unescaped before writing.\n\n")
	fmt.Fprintf(b, "| ClickBench kind | Parquet column type | Physical value written |\n")
	fmt.Fprintf(b, "| --- | --- | --- |\n")
	fmt.Fprintf(b, "| `int16` | `parquet.Int(16)` | `INT32`, signed 16-bit logical type |\n")
	fmt.Fprintf(b, "| `int32` | `parquet.Int(32)` | `INT32`, signed 32-bit logical type |\n")
	fmt.Fprintf(b, "| `int64` | `parquet.Int(64)` | `INT64`, signed 64-bit logical type |\n")
	fmt.Fprintf(b, "| `date` | `parquet.Date()` | `INT32` days since Unix epoch |\n")
	fmt.Fprintf(b, "| `timestamp_millis` | `parquet.Timestamp(parquet.Millisecond)` | `INT64` milliseconds since Unix epoch |\n")
	fmt.Fprintf(b, "| `string` | `parquet.String()` | `BYTE_ARRAY` UTF-8 string |\n\n")
}

func writeColumnStatsTSV(path string, columns []columnStat) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Comma = '\t'
	if err := w.Write([]string{
		"column",
		"type",
		"config_encoding",
		"metadata_encodings",
		"page_encodings",
		"values",
		"physical_bytes",
		"encoded_bytes",
		"compressed_bytes",
		"physical_encoded_ratio",
		"encoded_compressed_ratio",
		"physical_compressed_ratio",
		"source_field_bytes",
	}); err != nil {
		return err
	}
	for _, c := range columns {
		if err := w.Write([]string{
			c.Name,
			c.Kind,
			c.ConfiguredEncoding,
			stringSet(c.MetadataEncodings),
			int64MapString(c.PageEncodingStats),
			strconv.FormatInt(c.Values, 10),
			strconv.FormatInt(c.PhysicalBytes, 10),
			strconv.FormatInt(c.EncodedBytes, 10),
			strconv.FormatInt(c.CompressedBytes, 10),
			formatRatioDecimal(c.PhysicalBytes, c.EncodedBytes),
			formatRatioDecimal(c.EncodedBytes, c.CompressedBytes),
			formatRatioDecimal(c.PhysicalBytes, c.CompressedBytes),
			strconv.FormatInt(c.SourceBytes, 10),
		}); err != nil {
			return err
		}
	}
	w.Flush()
	return w.Error()
}

func writeColumnComparisonTSV(path string, candidateColumns, baselineColumns []columnStat) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Comma = '\t'
	if err := w.Write([]string{
		"column",
		"type",
		"candidate_encoding",
		"candidate_metadata_encodings",
		"candidate_page_encodings",
		"values",
		"baseline_encoded_bytes",
		"candidate_encoded_bytes",
		"candidate_compressed_bytes",
		"encoding_ratio",
		"candidate_codec_ratio",
		"total_ratio",
	}); err != nil {
		return err
	}

	baselineByName := columnStatsByName(baselineColumns)
	for _, candidate := range candidateColumns {
		baseline, ok := baselineByName[candidate.Name]
		if !ok {
			continue
		}
		if err := w.Write([]string{
			candidate.Name,
			candidate.Kind,
			candidate.ConfiguredEncoding,
			stringSet(candidate.MetadataEncodings),
			int64MapString(candidate.PageEncodingStats),
			strconv.FormatInt(candidate.Values, 10),
			strconv.FormatInt(baseline.EncodedBytes, 10),
			strconv.FormatInt(candidate.EncodedBytes, 10),
			strconv.FormatInt(candidate.CompressedBytes, 10),
			formatRatioDecimal(baseline.EncodedBytes, candidate.EncodedBytes),
			formatRatioDecimal(candidate.EncodedBytes, candidate.CompressedBytes),
			formatRatioDecimal(baseline.EncodedBytes, candidate.CompressedBytes),
		}); err != nil {
			return err
		}
	}
	w.Flush()
	return w.Error()
}

func writeColumnStatsMarkdown(b *strings.Builder, columns []columnStat, resultPath, columnStatsPath string) {
	fmt.Fprintf(b, "## Columns\n\n")
	fmt.Fprintf(b, "Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.\n\n")
	if columnStatsPath != "" {
		name := filepath.Base(columnStatsPath)
		fmt.Fprintf(b, "Column stats TSV: [%s](%s)\n\n", name, markdownLinkTarget(filepath.Dir(resultPath), columnStatsPath))
	}
	fmt.Fprintf(b, "| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |\n")
	fmt.Fprintf(b, "| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |\n")
	for _, c := range columns {
		fmt.Fprintf(b, "| `%s` | `%s` | `%s` | `%s` | `%s` | `%d` | `%d` (%s) | `%d` (%s) | `%d` (%s) | `%s` | `%s` | `%s` | `%d` (%s) |\n",
			c.Name,
			c.Kind,
			c.ConfiguredEncoding,
			stringSet(c.MetadataEncodings),
			int64MapString(c.PageEncodingStats),
			c.Values,
			c.PhysicalBytes, formatBytes(c.PhysicalBytes),
			c.EncodedBytes, formatBytes(c.EncodedBytes),
			c.CompressedBytes, formatBytes(c.CompressedBytes),
			formatMultiplier(c.PhysicalBytes, c.EncodedBytes),
			formatMultiplier(c.EncodedBytes, c.CompressedBytes),
			formatMultiplier(c.PhysicalBytes, c.CompressedBytes),
			c.SourceBytes, formatBytes(c.SourceBytes),
		)
	}
	fmt.Fprintf(b, "\n")
}

func markdownLinkTarget(fromDir, targetPath string) string {
	rel, err := filepath.Rel(fromDir, targetPath)
	if err != nil {
		return filepath.ToSlash(targetPath)
	}
	return filepath.ToSlash(rel)
}

func totalParquetBytes(files []fileStat) int64 {
	var total int64
	for _, f := range files {
		total += f.Size
	}
	return total
}

func totalPhysicalBytes(files []fileStat) int64 {
	var total int64
	for _, f := range files {
		total += f.PhysicalSize
	}
	return total
}

func totalEncodedBytes(files []fileStat) int64 {
	var total int64
	for _, f := range files {
		total += f.EncodedSize
	}
	return total
}

func totalCompressedDataBytes(files []fileStat) int64 {
	var total int64
	for _, f := range files {
		total += f.CompressedDataSize
	}
	return total
}

func formatMultiplier(numerator, denominator int64) string {
	if denominator == 0 {
		return "n/a"
	}
	return fmt.Sprintf("%.3fx", float64(numerator)/float64(denominator))
}

func formatRatioDecimal(numerator, denominator int64) string {
	if denominator == 0 {
		return ""
	}
	return fmt.Sprintf("%.6f", float64(numerator)/float64(denominator))
}

func stringSet(values map[string]struct{}) string {
	if len(values) == 0 {
		return "n/a"
	}
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return strings.Join(keys, ", ")
}

func int64MapString(values map[string]int64) string {
	if len(values) == 0 {
		return "n/a"
	}
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	parts := make([]string, len(keys))
	for i, key := range keys {
		parts[i] = fmt.Sprintf("%s:%d", key, values[key])
	}
	return strings.Join(parts, ", ")
}

func exitf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

var clickBenchColumns = []columnSpec{
	{"WatchID", kindInt64},
	{"JavaEnable", kindInt16},
	{"Title", kindString},
	{"GoodEvent", kindInt16},
	{"EventTime", kindTimestampMillis},
	{"EventDate", kindDate},
	{"CounterID", kindInt32},
	{"ClientIP", kindInt32},
	{"RegionID", kindInt32},
	{"UserID", kindInt64},
	{"CounterClass", kindInt16},
	{"OS", kindInt16},
	{"UserAgent", kindInt16},
	{"URL", kindString},
	{"Referer", kindString},
	{"IsRefresh", kindInt16},
	{"RefererCategoryID", kindInt16},
	{"RefererRegionID", kindInt32},
	{"URLCategoryID", kindInt16},
	{"URLRegionID", kindInt32},
	{"ResolutionWidth", kindInt16},
	{"ResolutionHeight", kindInt16},
	{"ResolutionDepth", kindInt16},
	{"FlashMajor", kindInt16},
	{"FlashMinor", kindInt16},
	{"FlashMinor2", kindString},
	{"NetMajor", kindInt16},
	{"NetMinor", kindInt16},
	{"UserAgentMajor", kindInt16},
	{"UserAgentMinor", kindString},
	{"CookieEnable", kindInt16},
	{"JavascriptEnable", kindInt16},
	{"IsMobile", kindInt16},
	{"MobilePhone", kindInt16},
	{"MobilePhoneModel", kindString},
	{"Params", kindString},
	{"IPNetworkID", kindInt32},
	{"TraficSourceID", kindInt16},
	{"SearchEngineID", kindInt16},
	{"SearchPhrase", kindString},
	{"AdvEngineID", kindInt16},
	{"IsArtifical", kindInt16},
	{"WindowClientWidth", kindInt16},
	{"WindowClientHeight", kindInt16},
	{"ClientTimeZone", kindInt16},
	{"ClientEventTime", kindTimestampMillis},
	{"SilverlightVersion1", kindInt16},
	{"SilverlightVersion2", kindInt16},
	{"SilverlightVersion3", kindInt32},
	{"SilverlightVersion4", kindInt16},
	{"PageCharset", kindString},
	{"CodeVersion", kindInt32},
	{"IsLink", kindInt16},
	{"IsDownload", kindInt16},
	{"IsNotBounce", kindInt16},
	{"FUniqID", kindInt64},
	{"OriginalURL", kindString},
	{"HID", kindInt32},
	{"IsOldCounter", kindInt16},
	{"IsEvent", kindInt16},
	{"IsParameter", kindInt16},
	{"DontCountHits", kindInt16},
	{"WithHash", kindInt16},
	{"HitColor", kindString},
	{"LocalEventTime", kindTimestampMillis},
	{"Age", kindInt16},
	{"Sex", kindInt16},
	{"Income", kindInt16},
	{"Interests", kindInt16},
	{"Robotness", kindInt16},
	{"RemoteIP", kindInt32},
	{"WindowName", kindInt32},
	{"OpenerName", kindInt32},
	{"HistoryLength", kindInt16},
	{"BrowserLanguage", kindString},
	{"BrowserCountry", kindString},
	{"SocialNetwork", kindString},
	{"SocialAction", kindString},
	{"HTTPError", kindInt16},
	{"SendTiming", kindInt32},
	{"DNSTiming", kindInt32},
	{"ConnectTiming", kindInt32},
	{"ResponseStartTiming", kindInt32},
	{"ResponseEndTiming", kindInt32},
	{"FetchTiming", kindInt32},
	{"SocialSourceNetworkID", kindInt16},
	{"SocialSourcePage", kindString},
	{"ParamPrice", kindInt64},
	{"ParamOrderID", kindString},
	{"ParamCurrency", kindString},
	{"ParamCurrencyID", kindInt16},
	{"OpenstatServiceName", kindString},
	{"OpenstatCampaignID", kindString},
	{"OpenstatAdID", kindString},
	{"OpenstatSourceID", kindString},
	{"UTMSource", kindString},
	{"UTMMedium", kindString},
	{"UTMCampaign", kindString},
	{"UTMContent", kindString},
	{"UTMTerm", kindString},
	{"FromTag", kindString},
	{"HasGCLID", kindInt16},
	{"RefererHash", kindInt64},
	{"URLHash", kindInt64},
	{"CLID", kindInt32},
}
