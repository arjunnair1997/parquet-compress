package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"html"
	"io"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"parquet_compress/internal/reportpdf"
)

const (
	defaultInput       = "data/clickbench/hits.tsv.gz"
	fixedMaxPageSize   = "256KiB"
	fixedMaxRowGroup   = "10MiB"
	fixedMaxFileSize   = "10MiB"
	defaultDictMaxSize = "1MiB"
)

type config struct {
	Rows           int64
	Parallel       int
	ZstdLevel      int
	Input          string
	MaxDictSize    string
	ExperimentDir  string
	ResultDir      string
	MarkdownDir    string
	ConfigDir      string
	TSVDir         string
	OutputRoot     string
	ShapeStatsJSON string
	Verify         bool
	SkipExisting   bool
	KeepOutput     bool
	GeneratePDF    bool
}

type combo struct {
	Index             int
	Compression       string
	CompressionName   string
	IntEncoding       string
	DateEncoding      string
	TimestampEncoding string
	StringEncoding    string
	Slug              string
}

type experimentResult struct {
	Combo           combo
	ResultPath      string
	ColumnTSVPath   string
	OutputDir       string
	LogPath         string
	Elapsed         time.Duration
	Columns         []columnResult
	PhysicalBytes   int64
	EncodedBytes    int64
	CompressedBytes int64
	Err             error
}

type columnResult struct {
	Column            string
	Type              string
	ConfigEncoding    string
	MetadataEncodings string
	PageEncodings     string
	Values            int64
	PhysicalBytes     int64
	EncodedBytes      int64
	CompressedBytes   int64
	SourceFieldBytes  int64
}

type experimentRanking struct {
	Result                  experimentResult
	BaselineEncodedBytes    int64
	PostEncodingRatio       float64
	PostCompressionRatio    float64
	CodecRatio              float64
	HasPostCompressionRatio bool
}

type columnObservation struct {
	Experiment                            experimentResult
	Column                                columnResult
	BaselineEncodedBytes                  int64
	PlainCompressedBytes                  int64
	PlainCompressionRatio                 float64
	PostEncodingRatio                     float64
	PostCompressionRatio                  float64
	CompressionRatioImprovementPct        float64
	CodecRatio                            float64
	TargetBytes                           int64
	TargetMetric                          string
	HasPostCompressionRatio               bool
	HasCompressionRatioImprovementPercent bool
	HasPlainCompressedBytes               bool
	HasPlainCompressionRatio              bool
}

type columnWinner struct {
	Scope       string
	Observation columnObservation
}

type settingSummary struct {
	IntEncoding       string
	DateEncoding      string
	TimestampEncoding string
	StringEncoding    string
	None              *experimentRanking
	Snappy            *experimentRanking
	Zstd              *experimentRanking
	ZstdNoEncoding    *experimentRanking
}

type columnShapeStatsSnapshot struct {
	Columns []columnShapeStats `json:"columns"`
	Errors  []string           `json:"errors,omitempty"`
}

type columnShapeStats struct {
	ColumnIndex      int                  `json:"column_index"`
	Path             []string             `json:"path"`
	Name             string               `json:"name"`
	Type             string               `json:"type"`
	PhysicalType     string               `json:"physical_type"`
	SortedAscending  bool                 `json:"sorted_ascending"`
	SortedDescending bool                 `json:"sorted_descending"`
	RowGroups        []shapeRowGroupStats `json:"row_groups"`
	Pages            []shapePageStats     `json:"pages"`
}

type shapeRowGroupStats struct {
	RowGroupIndex      int64 `json:"row_group_index"`
	NumRows            int64 `json:"num_rows"`
	Cardinality        int64 `json:"cardinality"`
	PageCount          int   `json:"page_count"`
	PageCardinalityMin int64 `json:"page_cardinality_min"`
	PageCardinalityMax int64 `json:"page_cardinality_max"`
	MinValueLength     int   `json:"min_value_length"`
	MaxValueLength     int   `json:"max_value_length"`
}

type shapePageStats struct {
	RowGroupIndex int64   `json:"row_group_index"`
	PageIndex     int     `json:"page_index"`
	FirstRowIndex int64   `json:"first_row_index"`
	NumRows       int64   `json:"num_rows"`
	NumValues     int64   `json:"num_values"`
	Cardinality   int64   `json:"cardinality"`
	HasBounds     bool    `json:"has_bounds"`
	MinValue      string  `json:"min_value,omitempty"`
	MaxValue      string  `json:"max_value,omitempty"`
	MinValueBytes string  `json:"min_value_bytes,omitempty"`
	MaxValueBytes string  `json:"max_value_bytes,omitempty"`
	HasNumeric    bool    `json:"has_numeric"`
	MinNumeric    float64 `json:"min_numeric,omitempty"`
	MaxNumeric    float64 `json:"max_numeric,omitempty"`
	MinLength     int     `json:"min_length"`
	MaxLength     int     `json:"max_length"`
}

type columnShapePlots struct {
	RowGroupCardinality string
	PageCardinality     string
	PageBounds          string
	ValueLength         string
}

type plotSeries struct {
	Name   string
	Color  string
	Values []float64
}

type winnerDistributionRow struct {
	Compression string
	Encoding    string
	Wins        int
}

type zstdPlainRLEDictComparison struct {
	Buckets            []zstdPlainRLEDictBucket
	ComparedColumns    int
	PlainBetterCount   int
	RLEDictBetterCount int
	TieCount           int
	MissingCount       int
}

type zstdPlainRLEDictBucket struct {
	Label         string
	PlainBetter   int
	RLEDictBetter int
}

type deltaBinaryPackedWinnerComparison struct {
	Buckets                []encodingImprovementBucket
	WinnerCount            int
	MissingSecondBestCount int
}

type snappyPlainRLEDictComparison struct {
	RLEDictBetterBuckets []encodingImprovementBucket
	PlainBetterBuckets   []encodingImprovementBucket
	ComparedColumns      int
	RLEDictBetterCount   int
	PlainBetterCount     int
	TieCount             int
	MissingCount         int
}

type encodingImprovementBucket struct {
	Label string
	Count int
}

type barChartBar struct {
	Label string
	Value int
	Color string
}

func main() {
	cfg, err := parseFlags(os.Args[1:])
	if err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return
		}
		exitf("%v", err)
	}
	if err := run(cfg); err != nil {
		exitf("%v", err)
	}
}

func parseFlags(args []string) (config, error) {
	cfg := config{
		Rows:        10_000,
		Parallel:    1,
		ZstdLevel:   3,
		Input:       defaultInput,
		MaxDictSize: defaultDictMaxSize,
		Verify:      true,
	}
	fs := flag.NewFlagSet("encoding-matrix", flag.ContinueOnError)
	fs.Int64Var(&cfg.Rows, "rows", cfg.Rows, "rows per experiment")
	fs.IntVar(&cfg.Parallel, "parallel", cfg.Parallel, "number of writer processes to run at once")
	fs.IntVar(&cfg.ZstdLevel, "zstd-level", cfg.ZstdLevel, "zstd level for the zstd compression matrix entries")
	fs.StringVar(&cfg.Input, "input", cfg.Input, "path to hits.tsv or hits.tsv.gz")
	fs.StringVar(&cfg.MaxDictSize, "max-dictionary-page-size", cfg.MaxDictSize, "maximum per-column dictionary bytes before falling back to plain encoding")
	fs.StringVar(&cfg.ExperimentDir, "experiment-dir", "", "fixed-settings experiment directory for result markdown/TSV files; defaults from fixed settings")
	fs.StringVar(&cfg.OutputRoot, "output-root", "", "root directory for generated parquet output directories; defaults under --experiment-dir")
	fs.StringVar(&cfg.ShapeStatsJSON, "column-shape-stats-json", "", "optional writer stats JSON used to enrich col_top_5.md; defaults under the row results directory")
	fs.BoolVar(&cfg.Verify, "verify", cfg.Verify, "verify every generated parquet output")
	fs.BoolVar(&cfg.SkipExisting, "skip-existing", cfg.SkipExisting, "reuse an existing result markdown/column TSV when present")
	fs.BoolVar(&cfg.KeepOutput, "keep-output", cfg.KeepOutput, "keep generated parquet output directories after the experiment; only valid with --parallel 1")
	fs.BoolVar(&cfg.GeneratePDF, "generate-pdf", cfg.GeneratePDF, "write sibling PDFs for generated markdown results; disabled by default")
	if err := fs.Parse(args); err != nil {
		return cfg, err
	}
	if cfg.Rows <= 0 {
		return cfg, fmt.Errorf("--rows must be > 0")
	}
	if cfg.Parallel <= 0 {
		return cfg, fmt.Errorf("--parallel must be > 0")
	}
	if cfg.KeepOutput && cfg.Parallel > 1 {
		return cfg, fmt.Errorf("--keep-output is only allowed with --parallel 1; parallel runs always delete generated parquet output")
	}
	if cfg.MaxDictSize == "" {
		return cfg, fmt.Errorf("--max-dictionary-page-size must not be empty")
	}
	if cfg.ExperimentDir == "" {
		cfg.ExperimentDir = defaultExperimentDir(cfg.MaxDictSize)
	}
	cfg.ResultDir = filepath.Join(cfg.ExperimentDir, rowsDirName(cfg.Rows))
	cfg.MarkdownDir = filepath.Join(cfg.ResultDir, "results")
	cfg.ConfigDir = filepath.Join(cfg.MarkdownDir, "configs")
	cfg.TSVDir = filepath.Join(cfg.ResultDir, "tsvs")
	if cfg.OutputRoot == "" {
		cfg.OutputRoot = filepath.Join(cfg.ResultDir, "parquet")
	}
	if cfg.ShapeStatsJSON == "" {
		cfg.ShapeStatsJSON = filepath.Join(cfg.MarkdownDir, "column_shape_stats", "column_shape_stats.json")
	}
	if !cfg.KeepOutput {
		if err := requireChildPath(cfg.ResultDir, cfg.OutputRoot, "--output-root"); err != nil {
			return cfg, err
		}
	}
	return cfg, nil
}

func requireChildPath(root, path, flagName string) error {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return err
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	rel, err := filepath.Rel(absRoot, absPath)
	if err != nil {
		return err
	}
	if rel == "." || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return fmt.Errorf("%s must be a child of the row-count result directory when generated parquet cleanup is enabled: %s", flagName, root)
	}
	return nil
}

func rowsDirName(rows int64) string {
	return fmt.Sprintf("%d_rows", rows)
}

func defaultExperimentDir(dictSize string) string {
	return filepath.Join(
		"encoding_experiment",
		fmt.Sprintf("page-256kib-rgsize-10mib-file-10mib-dictpage-%s", sanitizeFilename(dictSize)),
	)
}

func run(cfg config) error {
	started := time.Now()
	if err := os.MkdirAll(cfg.ExperimentDir, 0o755); err != nil {
		return err
	}
	if err := os.MkdirAll(cfg.ResultDir, 0o755); err != nil {
		return err
	}
	if err := os.MkdirAll(cfg.MarkdownDir, 0o755); err != nil {
		return err
	}
	if err := os.MkdirAll(cfg.ConfigDir, 0o755); err != nil {
		return err
	}
	if err := os.MkdirAll(cfg.TSVDir, 0o755); err != nil {
		return err
	}
	if err := os.MkdirAll(cfg.OutputRoot, 0o755); err != nil {
		return err
	}

	toolPath, err := buildWriterTool()
	if err != nil {
		return err
	}
	defer os.Remove(toolPath)

	combos := matrixCombos(cfg.Rows, cfg.ZstdLevel)
	fmt.Printf("running %d encoding/compression experiments with %d worker(s)\n", len(combos), cfg.Parallel)
	fmt.Printf("fixed settings: page=%s dictionary-page=%s row-group=%s file=%s rows=%d\n", fixedMaxPageSize, cfg.MaxDictSize, fixedMaxRowGroup, fixedMaxFileSize, cfg.Rows)
	fmt.Printf("result directory: %s\n", cfg.MarkdownDir)
	fmt.Printf("config result directory: %s\n", cfg.ConfigDir)
	fmt.Printf("tsv directory: %s\n", cfg.TSVDir)
	if cfg.KeepOutput {
		fmt.Printf("generated parquet output: kept under %s\n", cfg.OutputRoot)
	} else {
		fmt.Printf("generated parquet output: deleted after each successful experiment and again when the matrix finishes\n")
	}

	results := runExperiments(cfg, toolPath, combos)
	if !cfg.KeepOutput {
		if err := cleanOutputRoot(cfg.OutputRoot); err != nil {
			return err
		}
	}
	failures := 0
	successes := make([]experimentResult, 0, len(results))
	for _, result := range results {
		if result.Err != nil {
			failures++
			continue
		}
		successes = append(successes, result)
	}
	if failures > 0 {
		return fmt.Errorf("%d/%d experiments failed; first failure: %w", failures, len(results), firstFailure(results))
	}

	summary, err := writeAggregateFiles(cfg, successes, started, time.Now())
	if err != nil {
		return err
	}
	fmt.Printf("wrote matrix summary: %s\n", summary)
	if cfg.GeneratePDF {
		fmt.Printf("wrote matrix summary PDF: %s\n", reportpdf.PathForMarkdown(summary))
	}
	return nil
}

func buildWriterTool() (string, error) {
	root, err := os.Getwd()
	if err != nil {
		return "", err
	}
	toolPath := filepath.Join(os.TempDir(), fmt.Sprintf("clickbench-parquet-writer-%d", os.Getpid()))
	cmd := exec.Command("go", "build", "-o", toolPath, ".")
	cmd.Dir = root
	cmd.Env = buildEnv(root)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("build parquet writer: %w", err)
	}
	return toolPath, nil
}

func buildEnv(root string) []string {
	env := os.Environ()
	if os.Getenv("GOCACHE") == "" {
		env = append(env, "GOCACHE="+filepath.Join(root, ".gocache"))
	}
	return env
}

func matrixCombos(rows int64, zstdLevel int) []combo {
	compressions := []string{"none", "snappy", "zstd"}
	intEncodings := []string{"plain", "rle-dict", "delta-binary-packed"}
	dateEncodings := []string{"plain", "rle-dict", "delta-binary-packed"}
	timestampEncodings := []string{"plain", "rle-dict", "delta-binary-packed"}
	stringEncodings := []string{"plain", "rle-dict", "delta-byte-array", "delta-length-byte-array"}

	var combos []combo
	for _, compression := range compressions {
		compressionName := compression
		if compression == "zstd" {
			compressionName = fmt.Sprintf("zstd-%d", zstdLevel)
		}
		for _, intEncoding := range intEncodings {
			for _, dateEncoding := range dateEncodings {
				for _, timestampEncoding := range timestampEncodings {
					for _, stringEncoding := range stringEncodings {
						c := combo{
							Index:             len(combos) + 1,
							Compression:       compression,
							CompressionName:   compressionName,
							IntEncoding:       intEncoding,
							DateEncoding:      dateEncoding,
							TimestampEncoding: timestampEncoding,
							StringEncoding:    stringEncoding,
						}
						c.Slug = sanitizeFilename(fmt.Sprintf("rows-%d_comp-%s_int-%s_str-%s_date-%s_ts-%s",
							rows,
							compressionName,
							intEncoding,
							stringEncoding,
							dateEncoding,
							timestampEncoding,
						))
						combos = append(combos, c)
					}
				}
			}
		}
	}
	return combos
}

func runExperiments(cfg config, toolPath string, combos []combo) []experimentResult {
	workerCount := cfg.Parallel
	if workerCount > len(combos) {
		workerCount = len(combos)
	}

	jobs := make(chan combo)
	results := make(chan experimentResult)
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := range jobs {
				results <- runExperiment(cfg, toolPath, c)
			}
		}()
	}
	go func() {
		for _, c := range combos {
			jobs <- c
		}
		close(jobs)
		wg.Wait()
		close(results)
	}()

	ordered := make([]experimentResult, 0, len(combos))
	completed := 0
	for result := range results {
		completed++
		if result.Err != nil {
			fmt.Printf("[%d/%d] failed %s: %v (log: %s)\n", completed, len(combos), result.Combo.Slug, result.Err, result.LogPath)
		} else {
			fmt.Printf("[%d/%d] ok %s elapsed=%s encoded=%s compressed=%s\n",
				completed,
				len(combos),
				result.Combo.Slug,
				result.Elapsed.Round(time.Millisecond),
				formatBytes(result.EncodedBytes),
				formatBytes(result.CompressedBytes),
			)
		}
		ordered = append(ordered, result)
	}
	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].Combo.Index < ordered[j].Combo.Index
	})
	return ordered
}

func runExperiment(cfg config, toolPath string, c combo) experimentResult {
	result := experimentResult{
		Combo:     c,
		OutputDir: filepath.Join(cfg.OutputRoot, c.Slug),
		LogPath:   filepath.Join(os.TempDir(), fmt.Sprintf("encoding-matrix-logs-rows-%d", cfg.Rows), c.Slug+".log"),
	}
	started := time.Now()

	existingResult, existingColumns := findExistingResultFiles(cfg.ConfigDir, cfg.TSVDir, c.Slug)
	if cfg.SkipExisting && existingResult != "" && existingColumns != "" {
		if cfg.GeneratePDF {
			if err := ensurePDFForMarkdown(existingResult); err != nil {
				result.Elapsed = time.Since(started)
				result.Err = err
				return result
			}
		}
		columns, err := parseColumnStatsTSV(existingColumns)
		if elapsed, ok := parseWriteElapsed(existingResult); ok {
			result.Elapsed = elapsed
		} else {
			result.Elapsed = time.Since(started)
		}
		result.ResultPath = existingResult
		result.ColumnTSVPath = existingColumns
		result.Columns = columns
		result.Err = err
		result.sumColumns()
		return result
	}

	if !cfg.KeepOutput {
		if err := removeOutputDir(cfg.OutputRoot, result.OutputDir); err != nil {
			result.Elapsed = time.Since(started)
			result.Err = err
			return result
		}
	}

	if err := os.MkdirAll(filepath.Dir(result.LogPath), 0o755); err != nil {
		result.Err = err
		return result
	}
	logFile, err := os.Create(result.LogPath)
	if err != nil {
		result.Err = err
		return result
	}
	defer logFile.Close()

	args := []string{
		"--input", cfg.Input,
		"--rows", strconv.FormatInt(cfg.Rows, 10),
		"--output-dir", result.OutputDir,
		"--results-dir", cfg.ConfigDir,
		"--tsv-dir", cfg.TSVDir,
		"--max-page-size", fixedMaxPageSize,
		"--max-dictionary-page-size", cfg.MaxDictSize,
		"--max-row-group-size", fixedMaxRowGroup,
		"--max-file-size", fixedMaxFileSize,
		"--compression", c.Compression,
		"--int-encoding", c.IntEncoding,
		"--date-encoding", c.DateEncoding,
		"--timestamp-encoding", c.TimestampEncoding,
		"--string-encoding", c.StringEncoding,
	}
	if c.Compression == "zstd" {
		args = append(args, "--zstd-level", strconv.Itoa(cfg.ZstdLevel))
	}
	if cfg.Verify {
		args = append(args, "--verify")
	}
	if !cfg.GeneratePDF {
		args = append(args, "--generate-pdf=false")
	}

	cmd := exec.Command(toolPath, args...)
	cmd.Stdout = logFile
	cmd.Stderr = logFile
	if err := cmd.Run(); err != nil {
		result.Elapsed = time.Since(started)
		result.Err = fmt.Errorf("run writer: %w", err)
		return result
	}
	result.Elapsed = time.Since(started)

	result.ResultPath, result.ColumnTSVPath = findExistingResultFiles(cfg.ConfigDir, cfg.TSVDir, c.Slug)
	if result.ResultPath == "" || result.ColumnTSVPath == "" {
		result.Err = fmt.Errorf("result files not found for %s", c.Slug)
		return result
	}
	result.Columns, result.Err = parseColumnStatsTSV(result.ColumnTSVPath)
	result.sumColumns()
	if result.Err == nil && !cfg.KeepOutput {
		result.Err = removeOutputDir(cfg.OutputRoot, result.OutputDir)
	}
	return result
}

func parseWriteElapsed(markdownPath string) (time.Duration, bool) {
	data, err := os.ReadFile(markdownPath)
	if err != nil {
		return 0, false
	}
	const prefix = "- Write elapsed: `"
	for _, line := range strings.Split(string(data), "\n") {
		if !strings.HasPrefix(line, prefix) {
			continue
		}
		value := strings.TrimSuffix(strings.TrimPrefix(line, prefix), "`")
		elapsed, err := time.ParseDuration(value)
		if err != nil {
			return 0, false
		}
		return elapsed, true
	}
	return 0, false
}

func ensurePDFForMarkdown(markdownPath string) error {
	if _, err := os.Stat(reportpdf.PathForMarkdown(markdownPath)); err == nil {
		return nil
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}
	_, err := reportpdf.Generate(markdownPath)
	return err
}

func removeOutputDir(outputRoot, path string) error {
	if outputRoot == "" || path == "" || path == "." || path == string(filepath.Separator) {
		return fmt.Errorf("refusing to remove unsafe output directory %q under root %q", path, outputRoot)
	}
	absRoot, err := filepath.Abs(outputRoot)
	if err != nil {
		return err
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	rel, err := filepath.Rel(absRoot, absPath)
	if err != nil {
		return err
	}
	if rel == "." || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return fmt.Errorf("refusing to remove output directory %q outside root %q", path, outputRoot)
	}
	return os.RemoveAll(path)
}

func cleanOutputRoot(outputRoot string) error {
	if outputRoot == "" || outputRoot == "." || outputRoot == string(filepath.Separator) {
		return fmt.Errorf("refusing to clean unsafe output root %q", outputRoot)
	}
	entries, err := os.ReadDir(outputRoot)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if err := removeOutputDir(outputRoot, filepath.Join(outputRoot, entry.Name())); err != nil {
			return err
		}
	}
	return nil
}

func findExistingResultFiles(markdownDir, tsvDir, slug string) (string, string) {
	result := newestMatch(filepath.Join(markdownDir, "*_"+slug+".md"))
	columns := newestMatch(filepath.Join(tsvDir, "*_"+slug+"_columns.tsv"))
	return result, columns
}

func newestMatch(pattern string) string {
	matches, err := filepath.Glob(pattern)
	if err != nil || len(matches) == 0 {
		return ""
	}
	sort.Slice(matches, func(i, j int) bool {
		left, leftErr := os.Stat(matches[i])
		right, rightErr := os.Stat(matches[j])
		if leftErr != nil || rightErr != nil {
			return matches[i] > matches[j]
		}
		return left.ModTime().After(right.ModTime())
	})
	return matches[0]
}

func parseColumnStatsTSV(path string) ([]columnResult, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	header, err := reader.Read()
	if err != nil {
		return nil, err
	}
	index := make(map[string]int, len(header))
	for i, name := range header {
		index[name] = i
	}

	var columns []columnResult
	for {
		record, err := reader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		col := columnResult{
			Column:            field(record, index, "column"),
			Type:              field(record, index, "type"),
			ConfigEncoding:    field(record, index, "config_encoding"),
			MetadataEncodings: field(record, index, "metadata_encodings"),
			PageEncodings:     field(record, index, "page_encodings"),
			Values:            parseInt(field(record, index, "values")),
			PhysicalBytes:     parseInt(field(record, index, "physical_bytes")),
			EncodedBytes:      parseInt(field(record, index, "encoded_bytes")),
			CompressedBytes:   parseInt(field(record, index, "compressed_bytes")),
			SourceFieldBytes:  parseInt(field(record, index, "source_field_bytes")),
		}
		columns = append(columns, col)
	}
	return columns, nil
}

func field(record []string, index map[string]int, name string) string {
	i, ok := index[name]
	if !ok || i >= len(record) {
		return ""
	}
	return record[i]
}

func parseInt(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

func (r *experimentResult) sumColumns() {
	r.PhysicalBytes = 0
	r.EncodedBytes = 0
	r.CompressedBytes = 0
	for _, col := range r.Columns {
		r.PhysicalBytes += col.PhysicalBytes
		r.EncodedBytes += col.EncodedBytes
		r.CompressedBytes += col.CompressedBytes
	}
}

func writeAggregateFiles(cfg config, results []experimentResult, started, finished time.Time) (string, error) {
	baseline, err := findBaseline(results)
	if err != nil {
		return "", err
	}
	rankings := buildExperimentRankings(results, baseline.EncodedBytes)
	columnObservations, err := buildColumnObservations(results, baseline)
	if err != nil {
		return "", err
	}
	winners := pickColumnWinners(columnObservations)

	date := started.Format("2006-01-02")
	baseName := fmt.Sprintf("%s_rows-%d_encoding-matrix", date, cfg.Rows)
	tsvBase := filepath.Join(cfg.TSVDir, baseName)
	allExperimentsPath := tsvBase + "_experiments.tsv"
	settingsPath := tsvBase + "_settings.tsv"
	columnResultsPath := tsvBase + "_column-results.tsv"
	columnWinnersPath := tsvBase + "_column-winners.tsv"
	bestColumnEncodingsPath := tsvBase + "_best-column-encodings.tsv"
	summaryPath := filepath.Join(cfg.MarkdownDir, baseName+"_summary.md")
	columnTop5Path := filepath.Join(cfg.MarkdownDir, "col_top_5.md")

	if err := writeExperimentRankingsTSV(allExperimentsPath, rankings, "all"); err != nil {
		return "", err
	}
	settingSummaries, err := buildSettingSummaries(rankings)
	if err != nil {
		return "", err
	}
	if err := writeSettingSummariesTSV(settingsPath, settingSummaries); err != nil {
		return "", err
	}
	if err := writeColumnResultsTSV(columnResultsPath, columnObservations); err != nil {
		return "", err
	}
	if err := writeColumnWinnersTSV(columnWinnersPath, winners); err != nil {
		return "", err
	}
	bestColumns := bestColumnEncodings(winners)
	if err := writeBestColumnEncodingsTSV(bestColumnEncodingsPath, bestColumns); err != nil {
		return "", err
	}
	shapeStats, err := loadColumnShapeStats(cfg.ShapeStatsJSON)
	if err != nil {
		return "", err
	}
	shapePlots := map[string]columnShapePlots{}
	if shapeStats != nil {
		shapePlots, err = writeColumnShapePlots(*shapeStats, filepath.Join(filepath.Dir(cfg.ShapeStatsJSON), "images"))
		if err != nil {
			return "", err
		}
	}
	if err := writeColumnTop5Markdown(columnTop5Path, cfg, columnObservations, columnResultsPath, shapeStats, shapePlots); err != nil {
		return "", err
	}
	if err := writeSummaryMarkdown(summaryPath, cfg, rankings, settingSummaries, winners, bestColumns, baseline, started, finished, allExperimentsPath, settingsPath, columnResultsPath, columnWinnersPath, bestColumnEncodingsPath, columnTop5Path); err != nil {
		return "", err
	}
	return summaryPath, nil
}

func findBaseline(results []experimentResult) (experimentResult, error) {
	for _, result := range results {
		c := result.Combo
		if c.Compression == "none" && isPlainEncodingCombo(c) {
			return result, nil
		}
	}
	return experimentResult{}, fmt.Errorf("plain/uncompressed baseline not found")
}

func isPlainEncodingCombo(c combo) bool {
	return c.IntEncoding == "plain" &&
		c.DateEncoding == "plain" &&
		c.TimestampEncoding == "plain" &&
		c.StringEncoding == "plain"
}

func buildExperimentRankings(results []experimentResult, baselineEncodedBytes int64) []experimentRanking {
	rankings := make([]experimentRanking, 0, len(results))
	for _, result := range results {
		ranking := experimentRanking{
			Result:               result,
			BaselineEncodedBytes: baselineEncodedBytes,
			PostEncodingRatio:    ratio(baselineEncodedBytes, result.EncodedBytes),
			CodecRatio:           ratio(result.EncodedBytes, result.CompressedBytes),
			PostCompressionRatio: ratio(baselineEncodedBytes, result.CompressedBytes),
		}
		if result.Combo.Compression != "none" {
			ranking.HasPostCompressionRatio = true
		}
		rankings = append(rankings, ranking)
	}
	return rankings
}

func buildSettingSummaries(rankings []experimentRanking) ([]settingSummary, error) {
	zstdNoEncoding := findPlainCompressionRanking(rankings, "zstd")
	if zstdNoEncoding == nil {
		return nil, fmt.Errorf("zstd/plain baseline not found")
	}

	byKey := make(map[string]*settingSummary)
	for i := range rankings {
		ranking := &rankings[i]
		c := ranking.Result.Combo
		key := settingKey(c)
		summary := byKey[key]
		if summary == nil {
			summary = &settingSummary{
				IntEncoding:       c.IntEncoding,
				DateEncoding:      c.DateEncoding,
				TimestampEncoding: c.TimestampEncoding,
				StringEncoding:    c.StringEncoding,
			}
			byKey[key] = summary
		}
		switch c.Compression {
		case "none":
			summary.None = ranking
		case "snappy":
			summary.Snappy = ranking
		case "zstd":
			summary.Zstd = ranking
		default:
			return nil, fmt.Errorf("unknown compression %q in ranking", c.Compression)
		}
	}

	settings := make([]settingSummary, 0, len(byKey))
	for _, summary := range byKey {
		if summary.None == nil || summary.Snappy == nil || summary.Zstd == nil {
			return nil, fmt.Errorf("incomplete compression set for int=%s date=%s timestamp=%s string=%s", summary.IntEncoding, summary.DateEncoding, summary.TimestampEncoding, summary.StringEncoding)
		}
		summary.ZstdNoEncoding = zstdNoEncoding
		settings = append(settings, *summary)
	}
	sort.Slice(settings, func(i, j int) bool {
		left := settings[i]
		right := settings[j]
		if left.Zstd.PostCompressionRatio != right.Zstd.PostCompressionRatio {
			return left.Zstd.PostCompressionRatio > right.Zstd.PostCompressionRatio
		}
		if left.None.PostEncodingRatio != right.None.PostEncodingRatio {
			return left.None.PostEncodingRatio > right.None.PostEncodingRatio
		}
		return settingKey(left.None.Result.Combo) < settingKey(right.None.Result.Combo)
	})
	return settings, nil
}

func findPlainCompressionRanking(rankings []experimentRanking, compression string) *experimentRanking {
	for i := range rankings {
		ranking := &rankings[i]
		if ranking.Result.Combo.Compression == compression && isPlainEncodingCombo(ranking.Result.Combo) {
			return ranking
		}
	}
	return nil
}

func settingKey(c combo) string {
	return strings.Join([]string{c.IntEncoding, c.DateEncoding, c.TimestampEncoding, c.StringEncoding}, "\x00")
}

func buildColumnObservations(results []experimentResult, baseline experimentResult) ([]columnObservation, error) {
	baselineColumns := make(map[string]columnResult, len(baseline.Columns))
	for _, col := range baseline.Columns {
		baselineColumns[col.Column] = col
	}
	plainColumnsByCompression, err := buildPlainColumnsByCompression(results)
	if err != nil {
		return nil, err
	}

	var observations []columnObservation
	for _, result := range results {
		plainColumns, ok := plainColumnsByCompression[result.Combo.CompressionName]
		if !ok {
			return nil, fmt.Errorf("plain baseline missing for compression %s", result.Combo.CompressionName)
		}
		for _, col := range result.Columns {
			baselineCol, ok := baselineColumns[col.Column]
			if !ok {
				return nil, fmt.Errorf("baseline missing column %s", col.Column)
			}
			plainCol, ok := plainColumns[col.Column]
			if !ok {
				return nil, fmt.Errorf("plain baseline for compression %s missing column %s", result.Combo.CompressionName, col.Column)
			}
			obs := columnObservation{
				Experiment:               result,
				Column:                   col,
				BaselineEncodedBytes:     baselineCol.EncodedBytes,
				PlainCompressedBytes:     plainCol.CompressedBytes,
				PlainCompressionRatio:    ratio(baselineCol.EncodedBytes, plainCol.CompressedBytes),
				PostEncodingRatio:        ratio(baselineCol.EncodedBytes, col.EncodedBytes),
				CodecRatio:               ratio(col.EncodedBytes, col.CompressedBytes),
				HasPlainCompressedBytes:  true,
				HasPlainCompressionRatio: true,
			}
			if result.Combo.Compression == "none" {
				obs.TargetBytes = col.EncodedBytes
				obs.TargetMetric = "encoded_bytes"
			} else {
				obs.TargetBytes = col.CompressedBytes
				obs.TargetMetric = "compressed_bytes"
				obs.PostCompressionRatio = ratio(baselineCol.EncodedBytes, col.CompressedBytes)
				obs.HasPostCompressionRatio = true
				if obs.PlainCompressionRatio != 0 {
					obs.CompressionRatioImprovementPct = ((obs.PostCompressionRatio / obs.PlainCompressionRatio) - 1) * 100
					obs.HasCompressionRatioImprovementPercent = true
				}
			}
			observations = append(observations, obs)
		}
	}
	sort.Slice(observations, func(i, j int) bool {
		if observations[i].Column.Column != observations[j].Column.Column {
			return observations[i].Column.Column < observations[j].Column.Column
		}
		return observations[i].Experiment.Combo.Slug < observations[j].Experiment.Combo.Slug
	})
	return observations, nil
}

func buildPlainColumnsByCompression(results []experimentResult) (map[string]map[string]columnResult, error) {
	byCompression := make(map[string]map[string]columnResult)
	for _, result := range results {
		if !isPlainEncodingCombo(result.Combo) {
			continue
		}
		columns := make(map[string]columnResult, len(result.Columns))
		for _, col := range result.Columns {
			columns[col.Column] = col
		}
		byCompression[result.Combo.CompressionName] = columns
	}

	for _, compression := range compressionNames(results) {
		if _, ok := byCompression[compression]; !ok {
			return nil, fmt.Errorf("plain baseline missing for compression %s", compression)
		}
	}
	return byCompression, nil
}

func compressionNames(results []experimentResult) []string {
	seen := make(map[string]struct{})
	for _, result := range results {
		seen[result.Combo.CompressionName] = struct{}{}
	}
	names := make([]string, 0, len(seen))
	for name := range seen {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func pickColumnWinners(observations []columnObservation) []columnWinner {
	scopes := compressionScopes(observations)
	scopes = append(scopes, "overall")
	byColumn := make(map[string][]columnObservation)
	for _, obs := range observations {
		byColumn[obs.Column.Column] = append(byColumn[obs.Column.Column], obs)
	}

	columns := make([]string, 0, len(byColumn))
	for col := range byColumn {
		columns = append(columns, col)
	}
	sort.Strings(columns)

	var winners []columnWinner
	for _, column := range columns {
		for _, scope := range scopes {
			var best *columnObservation
			for i := range byColumn[column] {
				obs := byColumn[column][i]
				if scope != "overall" && obs.Experiment.Combo.CompressionName != scope {
					continue
				}
				if best == nil || betterColumnObservation(obs, *best) {
					copy := obs
					best = &copy
				}
			}
			if best != nil {
				winners = append(winners, columnWinner{
					Scope:       scope,
					Observation: *best,
				})
			}
		}
	}
	return winners
}

func compressionScopes(observations []columnObservation) []string {
	seen := make(map[string]struct{})
	for _, obs := range observations {
		seen[obs.Experiment.Combo.CompressionName] = struct{}{}
	}
	preferred := []string{"none", "snappy"}
	var scopes []string
	for _, scope := range preferred {
		if _, ok := seen[scope]; ok {
			scopes = append(scopes, scope)
			delete(seen, scope)
		}
	}
	remaining := make([]string, 0, len(seen))
	for scope := range seen {
		remaining = append(remaining, scope)
	}
	sort.Strings(remaining)
	return append(scopes, remaining...)
}

func betterColumnObservation(left, right columnObservation) bool {
	if left.TargetBytes != right.TargetBytes {
		return left.TargetBytes < right.TargetBytes
	}
	if left.Column.EncodedBytes != right.Column.EncodedBytes {
		return left.Column.EncodedBytes < right.Column.EncodedBytes
	}
	return left.Experiment.Combo.Slug < right.Experiment.Combo.Slug
}

func writeExperimentRankingsTSV(path string, rankings []experimentRanking, mode string) error {
	rows := make([]experimentRanking, 0, len(rankings))
	for _, ranking := range rankings {
		if mode == "compression" && !ranking.HasPostCompressionRatio {
			continue
		}
		if mode == "encoding" && ranking.Result.Combo.Compression != "none" {
			continue
		}
		rows = append(rows, ranking)
	}
	sort.Slice(rows, func(i, j int) bool {
		switch mode {
		case "compression":
			return rows[i].PostCompressionRatio > rows[j].PostCompressionRatio
		case "encoding":
			return rows[i].PostEncodingRatio > rows[j].PostEncodingRatio
		default:
			return rows[i].Result.Combo.Index < rows[j].Result.Combo.Index
		}
	})

	return writeTSV(path, []string{
		"rank",
		"experiment",
		"compression",
		"int_encoding",
		"date_encoding",
		"timestamp_encoding",
		"string_encoding",
		"physical_bytes",
		"baseline_encoded_bytes",
		"encoded_bytes",
		"compressed_bytes",
		"post_encoding_ratio",
		"post_compression_ratio",
		"codec_ratio",
		"elapsed_ms",
		"result_file",
		"column_stats_tsv",
		"output_dir",
	}, func(w *csv.Writer) error {
		for i, ranking := range rows {
			result := ranking.Result
			if err := w.Write([]string{
				strconv.Itoa(i + 1),
				result.Combo.Slug,
				result.Combo.CompressionName,
				result.Combo.IntEncoding,
				result.Combo.DateEncoding,
				result.Combo.TimestampEncoding,
				result.Combo.StringEncoding,
				strconv.FormatInt(result.PhysicalBytes, 10),
				strconv.FormatInt(ranking.BaselineEncodedBytes, 10),
				strconv.FormatInt(result.EncodedBytes, 10),
				strconv.FormatInt(result.CompressedBytes, 10),
				formatRatio(ranking.PostEncodingRatio),
				optionalRatio(ranking.PostCompressionRatio, ranking.HasPostCompressionRatio),
				formatRatio(ranking.CodecRatio),
				strconv.FormatInt(result.Elapsed.Milliseconds(), 10),
				result.ResultPath,
				result.ColumnTSVPath,
				result.OutputDir,
			}); err != nil {
				return err
			}
		}
		return nil
	})
}

func writeSettingSummariesTSV(path string, settings []settingSummary) error {
	return writeTSV(path, []string{
		"rank",
		"int_encoding",
		"date_encoding",
		"timestamp_encoding",
		"string_encoding",
		"baseline_encoded_bytes",
		"pre_compression",
		"snappy_compression",
		"zstd_compression",
		"zstd_compression_no_encoding",
		"pre_compression_result_file",
		"snappy_result_file",
		"zstd_result_file",
		"zstd_no_encoding_result_file",
		"pre_compression_column_stats_tsv",
		"snappy_column_stats_tsv",
		"zstd_column_stats_tsv",
		"zstd_no_encoding_column_stats_tsv",
	}, func(w *csv.Writer) error {
		for i, setting := range settings {
			if err := w.Write([]string{
				strconv.Itoa(i + 1),
				setting.IntEncoding,
				setting.DateEncoding,
				setting.TimestampEncoding,
				setting.StringEncoding,
				strconv.FormatInt(setting.None.BaselineEncodedBytes, 10),
				settingCell(*setting.None, false),
				settingCell(*setting.Snappy, true),
				settingCell(*setting.Zstd, true),
				settingCell(*setting.ZstdNoEncoding, true),
				setting.None.Result.ResultPath,
				setting.Snappy.Result.ResultPath,
				setting.Zstd.Result.ResultPath,
				setting.ZstdNoEncoding.Result.ResultPath,
				setting.None.Result.ColumnTSVPath,
				setting.Snappy.Result.ColumnTSVPath,
				setting.Zstd.Result.ColumnTSVPath,
				setting.ZstdNoEncoding.Result.ColumnTSVPath,
			}); err != nil {
				return err
			}
		}
		return nil
	})
}

func writeColumnResultsTSV(path string, observations []columnObservation) error {
	return writeTSV(path, []string{
		"column",
		"type",
		"compression",
		"column_config_encoding",
		"int_encoding",
		"date_encoding",
		"timestamp_encoding",
		"string_encoding",
		"physical_bytes",
		"baseline_encoded_bytes",
		"encoded_bytes",
		"compressed_bytes",
		"post_compression_no_encoding_bytes",
		"post_compression_no_encoding_ratio",
		"target_metric",
		"target_bytes",
		"post_encoding_ratio",
		"post_compression_ratio",
		"encoding_compression_ratio_improvement_pct",
		"codec_ratio",
		"experiment",
		"result_file",
		"column_stats_tsv",
	}, func(w *csv.Writer) error {
		for _, obs := range observations {
			c := obs.Experiment.Combo
			if err := w.Write([]string{
				obs.Column.Column,
				obs.Column.Type,
				c.CompressionName,
				obs.Column.ConfigEncoding,
				c.IntEncoding,
				c.DateEncoding,
				c.TimestampEncoding,
				c.StringEncoding,
				strconv.FormatInt(obs.Column.PhysicalBytes, 10),
				strconv.FormatInt(obs.BaselineEncodedBytes, 10),
				strconv.FormatInt(obs.Column.EncodedBytes, 10),
				strconv.FormatInt(obs.Column.CompressedBytes, 10),
				optionalInt(obs.PlainCompressedBytes, obs.HasPlainCompressedBytes),
				optionalRatio(obs.PlainCompressionRatio, obs.HasPlainCompressionRatio),
				obs.TargetMetric,
				strconv.FormatInt(obs.TargetBytes, 10),
				formatRatio(obs.PostEncodingRatio),
				optionalRatio(obs.PostCompressionRatio, obs.HasPostCompressionRatio),
				optionalPercent(obs.CompressionRatioImprovementPct, obs.HasCompressionRatioImprovementPercent),
				formatRatio(obs.CodecRatio),
				c.Slug,
				obs.Experiment.ResultPath,
				obs.Experiment.ColumnTSVPath,
			}); err != nil {
				return err
			}
		}
		return nil
	})
}

func writeColumnWinnersTSV(path string, winners []columnWinner) error {
	return writeTSV(path, []string{
		"column",
		"type",
		"scope",
		"winner_metric",
		"compression",
		"column_config_encoding",
		"int_encoding",
		"date_encoding",
		"timestamp_encoding",
		"string_encoding",
		"physical_bytes",
		"baseline_encoded_bytes",
		"encoded_bytes",
		"compressed_bytes",
		"post_compression_no_encoding_bytes",
		"post_compression_no_encoding_ratio",
		"target_bytes",
		"post_encoding_ratio",
		"post_compression_ratio",
		"encoding_compression_ratio_improvement_pct",
		"codec_ratio",
		"experiment",
		"result_file",
		"column_stats_tsv",
	}, func(w *csv.Writer) error {
		for _, winner := range winners {
			obs := winner.Observation
			c := obs.Experiment.Combo
			if err := w.Write([]string{
				obs.Column.Column,
				obs.Column.Type,
				winner.Scope,
				obs.TargetMetric,
				c.CompressionName,
				obs.Column.ConfigEncoding,
				c.IntEncoding,
				c.DateEncoding,
				c.TimestampEncoding,
				c.StringEncoding,
				strconv.FormatInt(obs.Column.PhysicalBytes, 10),
				strconv.FormatInt(obs.BaselineEncodedBytes, 10),
				strconv.FormatInt(obs.Column.EncodedBytes, 10),
				strconv.FormatInt(obs.Column.CompressedBytes, 10),
				optionalInt(obs.PlainCompressedBytes, obs.HasPlainCompressedBytes),
				optionalRatio(obs.PlainCompressionRatio, obs.HasPlainCompressionRatio),
				strconv.FormatInt(obs.TargetBytes, 10),
				formatRatio(obs.PostEncodingRatio),
				optionalRatio(obs.PostCompressionRatio, obs.HasPostCompressionRatio),
				optionalPercent(obs.CompressionRatioImprovementPct, obs.HasCompressionRatioImprovementPercent),
				formatRatio(obs.CodecRatio),
				c.Slug,
				obs.Experiment.ResultPath,
				obs.Experiment.ColumnTSVPath,
			}); err != nil {
				return err
			}
		}
		return nil
	})
}

func bestColumnEncodings(winners []columnWinner) []columnWinner {
	best := make([]columnWinner, 0, len(winners)/4)
	for _, winner := range winners {
		if winner.Scope == "overall" {
			best = append(best, winner)
		}
	}
	sort.Slice(best, func(i, j int) bool {
		return best[i].Observation.Column.Column < best[j].Observation.Column.Column
	})
	return best
}

func writeBestColumnEncodingsTSV(path string, winners []columnWinner) error {
	return writeTSV(path, []string{
		"column",
		"type",
		"best_encoding",
		"compression",
		"winner_metric",
		"physical_bytes",
		"baseline_encoded_bytes",
		"target_bytes",
		"encoded_bytes",
		"compressed_bytes",
		"post_compression_no_encoding_bytes",
		"post_compression_no_encoding_ratio",
		"post_encoding_ratio",
		"post_compression_ratio",
		"encoding_compression_ratio_improvement_pct",
		"codec_ratio",
		"int_encoding",
		"date_encoding",
		"timestamp_encoding",
		"string_encoding",
		"experiment",
		"result_file",
		"column_stats_tsv",
	}, func(w *csv.Writer) error {
		for _, winner := range winners {
			obs := winner.Observation
			c := obs.Experiment.Combo
			if err := w.Write([]string{
				obs.Column.Column,
				obs.Column.Type,
				obs.Column.ConfigEncoding,
				c.CompressionName,
				obs.TargetMetric,
				strconv.FormatInt(obs.Column.PhysicalBytes, 10),
				strconv.FormatInt(obs.BaselineEncodedBytes, 10),
				strconv.FormatInt(obs.TargetBytes, 10),
				strconv.FormatInt(obs.Column.EncodedBytes, 10),
				strconv.FormatInt(obs.Column.CompressedBytes, 10),
				optionalInt(obs.PlainCompressedBytes, obs.HasPlainCompressedBytes),
				optionalRatio(obs.PlainCompressionRatio, obs.HasPlainCompressionRatio),
				formatRatio(obs.PostEncodingRatio),
				optionalRatio(obs.PostCompressionRatio, obs.HasPostCompressionRatio),
				optionalPercent(obs.CompressionRatioImprovementPct, obs.HasCompressionRatioImprovementPercent),
				formatRatio(obs.CodecRatio),
				c.IntEncoding,
				c.DateEncoding,
				c.TimestampEncoding,
				c.StringEncoding,
				c.Slug,
				obs.Experiment.ResultPath,
				obs.Experiment.ColumnTSVPath,
			}); err != nil {
				return err
			}
		}
		return nil
	})
}

func writeSummaryMarkdown(path string, cfg config, rankings []experimentRanking, settings []settingSummary, winners []columnWinner, bestColumnWinners []columnWinner, baseline experimentResult, started, finished time.Time, allExperimentsPath, settingsPath, columnResultsPath, columnWinnersPath, bestColumnEncodingsPath, columnTop5Path string) error {
	var b strings.Builder
	fmt.Fprintf(&b, "# Encoding Matrix Summary\n\n")
	fmt.Fprintf(&b, "- Started: `%s`\n", started.Format(time.RFC3339))
	fmt.Fprintf(&b, "- Finished: `%s`\n", finished.Format(time.RFC3339))
	fmt.Fprintf(&b, "- Elapsed: `%s`\n", finished.Sub(started).Round(time.Millisecond))
	fmt.Fprintf(&b, "- Rows per experiment: `%d`\n", cfg.Rows)
	fmt.Fprintf(&b, "- Experiments: `%d`\n", len(rankings))
	fmt.Fprintf(&b, "- Parallelism: `%d`\n", cfg.Parallel)
	fmt.Fprintf(&b, "- Fixed max page size: `%s`\n", fixedMaxPageSize)
	fmt.Fprintf(&b, "- Fixed max dictionary page size: `%s`\n", cfg.MaxDictSize)
	fmt.Fprintf(&b, "- Fixed max row group size: `%s`\n", fixedMaxRowGroup)
	fmt.Fprintf(&b, "- Fixed max file size: `%s`\n", fixedMaxFileSize)
	summaryDir := filepath.Dir(path)
	fmt.Fprintf(&b, "- Plain/uncompressed baseline: [`%s`](%s)\n\n", filepath.Base(baseline.ResultPath), markdownLinkTarget(summaryDir, baseline.ResultPath))

	fmt.Fprintf(&b, "## Outputs\n\n")
	fmt.Fprintf(&b, "- All experiments: [%s](%s)\n", filepath.Base(allExperimentsPath), markdownLinkTarget(summaryDir, allExperimentsPath))
	fmt.Fprintf(&b, "- Settings with pre/post compression side by side: [%s](%s)\n", filepath.Base(settingsPath), markdownLinkTarget(summaryDir, settingsPath))
	fmt.Fprintf(&b, "- All per-column observations: [%s](%s)\n", filepath.Base(columnResultsPath), markdownLinkTarget(summaryDir, columnResultsPath))
	fmt.Fprintf(&b, "- Per-column winners by scope: [%s](%s)\n", filepath.Base(columnWinnersPath), markdownLinkTarget(summaryDir, columnWinnersPath))
	fmt.Fprintf(&b, "- Best encoding per column: [%s](%s)\n", filepath.Base(bestColumnEncodingsPath), markdownLinkTarget(summaryDir, bestColumnEncodingsPath))
	fmt.Fprintf(&b, "- Column top 5 rankings with shape stats: [%s](%s)\n\n", filepath.Base(columnTop5Path), markdownLinkTarget(summaryDir, columnTop5Path))

	fmt.Fprintf(&b, "## Ranking Definitions\n\n")
	fmt.Fprintf(&b, "- Pre-compression uses the `none` run for the same encoding setting: plain/uncompressed baseline encoded bytes divided by that setting's encoded bytes.\n")
	fmt.Fprintf(&b, "- Snappy and ZSTD compression use the compressed bytes for the same encoding setting: plain/uncompressed baseline encoded bytes divided by compressed bytes.\n")
	fmt.Fprintf(&b, "- Column ratios use `baseline_encoded_bytes` as their denominator: the same column's encoded bytes from the all-plain/no-compression run. `physical_bytes` is shown separately and is not used as the ratio denominator.\n")
	fmt.Fprintf(&b, "- ZSTD compression without encoding is the all-plain ZSTD run, repeated in each Top Encoding Settings row as the zstd-only baseline.\n")
	fmt.Fprintf(&b, "- `post_compression_no_encoding_bytes` is the same column's compressed bytes from the all-plain run with the same compression setting; `post_compression_no_encoding_ratio` is plain/uncompressed baseline encoded bytes divided by those bytes.\n")
	fmt.Fprintf(&b, "- `encoding_compression_ratio_improvement_pct` is `((post_compression_ratio / post_compression_no_encoding_ratio) - 1) * 100`; positive means encoding improved the compression ratio versus compressing plain pages.\n")
	fmt.Fprintf(&b, "- Codec ratio: candidate encoded bytes divided by candidate compressed bytes.\n\n")

	fmt.Fprintf(&b, "## Top Encoding Settings\n\n")
	writeSettingTable(&b, settings, 20, summaryDir)

	fmt.Fprintf(&b, "## Column Winners\n\n")
	fmt.Fprintf(&b, "Best means smallest target bytes across all 96 runs for that column. For `none`, target bytes are encoded bytes; for Snappy/ZSTD, target bytes are compressed bytes.\n\n")
	writeBestColumnEncodingsMarkdown(&b, bestColumnWinners, summaryDir)
	if cfg.GeneratePDF {
		return reportpdf.WriteMarkdownAndPDF(path, []byte(b.String()))
	}
	return os.WriteFile(path, []byte(b.String()), 0o644)
}

func writeSettingTable(b *strings.Builder, settings []settingSummary, limit int, summaryDir string) {
	if len(settings) < limit {
		limit = len(settings)
	}
	fmt.Fprintf(b, "| Rank | Int | Date | Timestamp | String | Pre-compression | Snappy compression | ZSTD compression | ZSTD compression (no encoding) | Result |\n")
	fmt.Fprintf(b, "| ---: | --- | --- | --- | --- | --- | --- | --- | --- | --- |\n")
	for i := 0; i < limit; i++ {
		setting := settings[i]
		fmt.Fprintf(b, "| %d | `%s` | `%s` | `%s` | `%s` | %s | %s | %s | %s | [%s](%s) |\n",
			i+1,
			setting.IntEncoding,
			setting.DateEncoding,
			setting.TimestampEncoding,
			setting.StringEncoding,
			settingMarkdownCell(*setting.None, false),
			settingMarkdownCell(*setting.Snappy, true),
			settingMarkdownCell(*setting.Zstd, true),
			settingMarkdownCell(*setting.ZstdNoEncoding, true),
			filepath.Base(setting.Zstd.Result.ResultPath),
			markdownLinkTarget(summaryDir, setting.Zstd.Result.ResultPath),
		)
	}
	fmt.Fprintf(b, "\n")
}

func writeBestColumnEncodingsMarkdown(b *strings.Builder, winners []columnWinner, summaryDir string) {
	fmt.Fprintf(b, "| Column | Type | Best encoding | Compression | Physical bytes | Baseline encoded bytes | Target bytes | Encoded bytes | Compressed bytes | Post-compression no encoding bytes | Post-compression no encoding | Post-encoding | Post-compression | Encoding compression-ratio improvement | Result |\n")
	fmt.Fprintf(b, "| --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | --- |\n")
	for _, winner := range winners {
		obs := winner.Observation
		c := obs.Experiment.Combo
		fmt.Fprintf(b, "| `%s` | `%s` | `%s` | `%s` | `%d` | `%d` | `%d` | `%d` | `%d` | `%s` | `%s` | `%s` | `%s` | `%s` | [%s](%s) |\n",
			obs.Column.Column,
			obs.Column.Type,
			obs.Column.ConfigEncoding,
			c.CompressionName,
			obs.Column.PhysicalBytes,
			obs.BaselineEncodedBytes,
			obs.TargetBytes,
			obs.Column.EncodedBytes,
			obs.Column.CompressedBytes,
			optionalInt(obs.PlainCompressedBytes, obs.HasPlainCompressedBytes),
			optionalRatio(obs.PlainCompressionRatio, obs.HasPlainCompressionRatio),
			formatRatio(obs.PostEncodingRatio),
			optionalRatio(obs.PostCompressionRatio, obs.HasPostCompressionRatio),
			optionalPercentMarkdown(obs.CompressionRatioImprovementPct, obs.HasCompressionRatioImprovementPercent),
			filepath.Base(obs.Experiment.ResultPath),
			markdownLinkTarget(summaryDir, obs.Experiment.ResultPath),
		)
	}
	fmt.Fprintf(b, "\n")
}

func loadColumnShapeStats(path string) (*columnShapeStatsSnapshot, error) {
	f, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var snapshot columnShapeStatsSnapshot
	if err := json.NewDecoder(f).Decode(&snapshot); err != nil {
		return nil, err
	}
	return &snapshot, nil
}

func writeColumnTop5Markdown(path string, cfg config, observations []columnObservation, columnResultsPath string, shapeStats *columnShapeStatsSnapshot, shapePlots map[string]columnShapePlots) error {
	reportDir := filepath.Dir(path)
	byColumn := make(map[string][]columnObservation)
	for _, obs := range observations {
		byColumn[obs.Column.Column] = append(byColumn[obs.Column.Column], obs)
	}
	shapeByColumn := make(map[string]columnShapeStats)
	if shapeStats != nil {
		for _, col := range shapeStats.Columns {
			shapeByColumn[col.Name] = col
		}
	}

	columns := make([]string, 0, len(byColumn))
	for column := range byColumn {
		columns = append(columns, column)
	}
	sort.Strings(columns)

	var b strings.Builder
	fmt.Fprintf(&b, "# Column Top 5 Encoding Rankings\n\n")
	fmt.Fprintf(&b, "- Experiment: `%s/%s`\n", filepath.Base(cfg.ExperimentDir), filepath.Base(cfg.ResultDir))
	fmt.Fprintf(&b, "- Source data: [%s](%s)\n", filepath.Base(columnResultsPath), markdownLinkTarget(reportDir, columnResultsPath))
	fmt.Fprintf(&b, "- Rows: `%s`\n", formatCount(cfg.Rows))
	fmt.Fprintf(&b, "- Ranking metric: per-column `compressed_bytes`, after Parquet page encoding and Snappy/ZSTD compression.\n")
	fmt.Fprintf(&b, "- Each numbered item starts with the achieved compressed size for that encoding/compression choice.\n")
	fmt.Fprintf(&b, "- Duplicate matrix rows with the same effective column encoding are collapsed to the best observed row before ranking.\n")
	fmt.Fprintf(&b, "- Encodings in this matrix: `plain`, `rle-dict`, `delta-binary-packed`, `delta-byte-array`, `delta-length-byte-array`.\n")
	if shapeStats != nil {
		fmt.Fprintf(&b, "- Column shape stats: [%s](%s)\n", filepath.Base(cfg.ShapeStatsJSON), markdownLinkTarget(reportDir, cfg.ShapeStatsJSON))
		if len(shapeStats.Errors) > 0 {
			fmt.Fprintf(&b, "- Writer stats collection errors: `%d`\n", len(shapeStats.Errors))
		}
	}
	fmt.Fprintf(&b, "\n")

	winnerDistribution := buildCompressedWinnerDistribution(byColumn)
	winnerDistributionPath := filepath.Join(reportDir, "images", "column_winner_distribution.svg")
	if err := writeWinnerDistributionSVG(winnerDistributionPath, winnerDistribution); err != nil {
		return err
	}
	fmt.Fprintf(&b, "## Winner Distribution\n\n")
	fmt.Fprintf(&b, "Counts are based on each column's first `Compressed overall` ranking below: one winner per column, grouped by compression algorithm and configured column encoding.\n\n")
	writeShapeImage(&b, "Column winner distribution", winnerDistributionPath, reportDir)
	writeWinnerDistributionTable(&b, winnerDistribution)

	zstdComparison := buildZstdPlainRLEDictComparison(byColumn)
	zstdComparisonPath := filepath.Join(reportDir, "images", "zstd_plain_vs_rle_dict_improvement.svg")
	if err := writeZstdPlainRLEDictComparisonSVG(zstdComparisonPath, zstdComparison); err != nil {
		return err
	}
	fmt.Fprintf(&b, "## ZSTD Plain vs RLE Dict Improvement Distribution\n\n")
	fmt.Fprintf(&b, "For each column, this compares the best observed `zstd + plain` compressed byte count with the best observed `zstd + rle-dict` compressed byte count. Improvement is `(larger compressed bytes - smaller compressed bytes) / larger compressed bytes * 100`.\n\n")
	writeShapeImage(&b, "ZSTD plain versus RLE dictionary improvement distribution", zstdComparisonPath, reportDir)
	writeZstdPlainRLEDictComparisonTable(&b, zstdComparison)

	deltaBinaryPackedComparison := buildDeltaBinaryPackedWinnerComparison(byColumn)
	deltaBinaryPackedComparisonPath := filepath.Join(reportDir, "images", "delta_binary_packed_winner_vs_second_best_improvement.svg")
	if err := writeDeltaBinaryPackedWinnerComparisonSVG(deltaBinaryPackedComparisonPath, deltaBinaryPackedComparison); err != nil {
		return err
	}
	fmt.Fprintf(&b, "## Delta Binary Packed Winner vs Second Best Improvement Distribution\n\n")
	fmt.Fprintf(&b, "For each column, this looks at the `Compressed overall` ranking below. Only columns where `delta-binary-packed` is the best observed compressed result are bucketed. Improvement is `(second-best compressed bytes - delta-binary-packed compressed bytes) / second-best compressed bytes * 100`.\n\n")
	writeShapeImage(&b, "Delta binary packed winner improvement over second best", deltaBinaryPackedComparisonPath, reportDir)
	writeDeltaBinaryPackedWinnerComparisonTable(&b, deltaBinaryPackedComparison)

	snappyComparison := buildSnappyPlainRLEDictComparison(byColumn)
	snappyRLEDictBetterPath := filepath.Join(reportDir, "images", "snappy_rle_dict_better_than_plain_improvement.svg")
	if err := writeSnappyRLEDictBetterComparisonSVG(snappyRLEDictBetterPath, snappyComparison); err != nil {
		return err
	}
	snappyPlainBetterPath := filepath.Join(reportDir, "images", "snappy_plain_better_than_rle_dict_improvement.svg")
	if err := writeSnappyPlainBetterComparisonSVG(snappyPlainBetterPath, snappyComparison); err != nil {
		return err
	}
	fmt.Fprintf(&b, "## Snappy Plain vs RLE Dict Improvement Distribution\n\n")
	fmt.Fprintf(&b, "For each column, this compares the best observed `snappy + plain` compressed byte count with the best observed `snappy + rle-dict` compressed byte count. Improvement is `(larger compressed bytes - smaller compressed bytes) / larger compressed bytes * 100`.\n\n")
	fmt.Fprintf(&b, "- Compared columns: `%d`\n", snappyComparison.ComparedColumns)
	fmt.Fprintf(&b, "- `snappy + rle-dict` smaller: `%d`; `snappy + plain` smaller: `%d`; ties: `%d`; missing comparisons: `%d`\n\n",
		snappyComparison.RLEDictBetterCount,
		snappyComparison.PlainBetterCount,
		snappyComparison.TieCount,
		snappyComparison.MissingCount,
	)
	writeShapeImage(&b, "Snappy RLE dictionary improvement over plain", snappyRLEDictBetterPath, reportDir)
	writeSnappyRLEDictBetterComparisonTable(&b, snappyComparison)
	writeShapeImage(&b, "Snappy plain improvement over RLE dictionary", snappyPlainBetterPath, reportDir)
	writeSnappyPlainBetterComparisonTable(&b, snappyComparison)

	for _, column := range columns {
		observations := byColumn[column]
		columnType := observations[0].Column.Type
		fmt.Fprintf(&b, "## %s (%s)\n\n", column, columnType)
		if shape, ok := shapeByColumn[column]; ok {
			writeColumnShapeStatsMarkdown(&b, shape, shapePlots[column], reportDir)
		}
		writeRankingList(&b, "Compressed overall", topColumnObservations(observations, "overall", 5), true)
		writeRankingList(&b, "ZSTD", topColumnObservations(observations, "zstd", 5), false)
		writeRankingList(&b, "Snappy", topColumnObservations(observations, "snappy", 5), false)
	}

	if cfg.GeneratePDF {
		return reportpdf.WriteMarkdownAndPDF(path, []byte(b.String()))
	}
	return os.WriteFile(path, []byte(b.String()), 0o644)
}

func writeColumnShapeStatsMarkdown(b *strings.Builder, shape columnShapeStats, plots columnShapePlots, reportDir string) {
	rowGroupCardinalities := make([]int64, 0, len(shape.RowGroups))
	pageCardinalityMins := make([]int64, 0, len(shape.RowGroups))
	pageCardinalityMaxes := make([]int64, 0, len(shape.RowGroups))
	minLengths := make([]int, 0, len(shape.RowGroups))
	maxLengths := make([]int, 0, len(shape.RowGroups))
	for _, rg := range shape.RowGroups {
		rowGroupCardinalities = append(rowGroupCardinalities, rg.Cardinality)
		pageCardinalityMins = append(pageCardinalityMins, rg.PageCardinalityMin)
		pageCardinalityMaxes = append(pageCardinalityMaxes, rg.PageCardinalityMax)
		minLengths = append(minLengths, rg.MinValueLength)
		maxLengths = append(maxLengths, rg.MaxValueLength)
	}

	fmt.Fprintf(b, "Column shape stats:\n")
	fmt.Fprintf(b, "- Parquet type: `%s`; physical type: `%s`\n", shape.Type, shape.PhysicalType)
	fmt.Fprintf(b, "- Sorted ascending across written rows: `%t`; sorted descending: `%t`\n", shape.SortedAscending, shape.SortedDescending)
	fmt.Fprintf(b, "- Row groups: `%d`; pages: `%d`\n", len(shape.RowGroups), len(shape.Pages))
	fmt.Fprintf(b, "- Row-group cardinality min/median/max: `%s`\n", summarizeInt64(rowGroupCardinalities))
	fmt.Fprintf(b, "- Page cardinality per row group min/median/max of mins: `%s`; of maxes: `%s`\n", summarizeInt64(pageCardinalityMins), summarizeInt64(pageCardinalityMaxes))
	fmt.Fprintf(b, "- Value length per row group min/median/max of mins: `%s`; of maxes: `%s`\n\n", summarizeInt(minLengths), summarizeInt(maxLengths))

	writeShapeImage(b, "Row-group cardinality", plots.RowGroupCardinality, reportDir)
	writeShapeImage(b, "Page cardinality min/max per row group", plots.PageCardinality, reportDir)
	writeShapeImage(b, "Page min/max distribution", plots.PageBounds, reportDir)
	writeShapeImage(b, "Value length min/max per row group", plots.ValueLength, reportDir)
	fmt.Fprintf(b, "\n")
}

func writeShapeImage(b *strings.Builder, alt, path, reportDir string) {
	if path == "" {
		return
	}
	fmt.Fprintf(b, "![%s](%s)\n\n", alt, markdownImageTarget(reportDir, path))
}

func buildCompressedWinnerDistribution(byColumn map[string][]columnObservation) []winnerDistributionRow {
	counts := make(map[string]*winnerDistributionRow)
	for _, observations := range byColumn {
		top := topColumnObservations(observations, "overall", 1)
		if len(top) == 0 {
			continue
		}
		obs := top[0]
		c := obs.Experiment.Combo
		key := c.CompressionName + "\x00" + obs.Column.ConfigEncoding
		row := counts[key]
		if row == nil {
			row = &winnerDistributionRow{
				Compression: c.CompressionName,
				Encoding:    obs.Column.ConfigEncoding,
			}
			counts[key] = row
		}
		row.Wins++
	}

	rows := make([]winnerDistributionRow, 0, len(counts))
	for _, row := range counts {
		rows = append(rows, *row)
	}
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].Wins != rows[j].Wins {
			return rows[i].Wins > rows[j].Wins
		}
		if rows[i].Compression != rows[j].Compression {
			return rows[i].Compression < rows[j].Compression
		}
		return rows[i].Encoding < rows[j].Encoding
	})
	return rows
}

func writeWinnerDistributionTable(b *strings.Builder, rows []winnerDistributionRow) {
	if len(rows) == 0 {
		fmt.Fprintf(b, "No compressed winner data was available.\n\n")
		return
	}
	fmt.Fprintf(b, "| Compression | Encoding | Column wins |\n")
	fmt.Fprintf(b, "| --- | --- | ---: |\n")
	for _, row := range rows {
		fmt.Fprintf(b, "| `%s` | `%s` | %d |\n", row.Compression, row.Encoding, row.Wins)
	}
	fmt.Fprintf(b, "\n")
}

func writeWinnerDistributionSVG(path string, rows []winnerDistributionRow) error {
	bars := make([]barChartBar, 0, len(rows))
	for _, row := range rows {
		bars = append(bars, barChartBar{
			Label: fmt.Sprintf("%s + %s", row.Compression, row.Encoding),
			Value: row.Wins,
			Color: compressionColor(row.Compression),
		})
	}
	return writeHorizontalBarChartSVG(path, "Column wins by compression + encoding", "columns won", bars)
}

func buildZstdPlainRLEDictComparison(byColumn map[string][]columnObservation) zstdPlainRLEDictComparison {
	type comparisonValue struct {
		plainBetter bool
		pct         float64
	}

	var values []comparisonValue
	summary := zstdPlainRLEDictComparison{}
	for _, observations := range byColumn {
		plain, plainOK := bestCompressedObservationFor(observations, "zstd", "plain")
		rleDict, rleDictOK := bestCompressedObservationFor(observations, "zstd", "rle-dict")
		if !plainOK || !rleDictOK {
			summary.MissingCount++
			continue
		}
		summary.ComparedColumns++
		plainBytes := plain.Column.CompressedBytes
		rleDictBytes := rleDict.Column.CompressedBytes
		switch {
		case plainBytes < rleDictBytes:
			summary.PlainBetterCount++
			values = append(values, comparisonValue{
				plainBetter: true,
				pct:         percentSmaller(plainBytes, rleDictBytes),
			})
		case rleDictBytes < plainBytes:
			summary.RLEDictBetterCount++
			values = append(values, comparisonValue{
				pct: percentSmaller(rleDictBytes, plainBytes),
			})
		default:
			summary.TieCount++
		}
	}

	maxBucket := 1
	for _, value := range values {
		index := improvementBucketIndex(value.pct)
		if index > maxBucket {
			maxBucket = index
		}
	}
	if maxBucket > 10 {
		maxBucket = 10
	}
	summary.Buckets = make([]zstdPlainRLEDictBucket, maxBucket+1)
	for i := range summary.Buckets {
		summary.Buckets[i].Label = improvementBucketLabel(i)
	}
	for _, value := range values {
		index := improvementBucketIndex(value.pct)
		if index >= len(summary.Buckets) {
			index = len(summary.Buckets) - 1
		}
		if value.plainBetter {
			summary.Buckets[index].PlainBetter++
		} else {
			summary.Buckets[index].RLEDictBetter++
		}
	}
	return summary
}

func bestCompressedObservationFor(observations []columnObservation, compressionPrefix, encoding string) (columnObservation, bool) {
	var best columnObservation
	found := false
	for _, obs := range observations {
		if !strings.HasPrefix(obs.Experiment.Combo.CompressionName, compressionPrefix) {
			continue
		}
		if obs.Column.ConfigEncoding != encoding {
			continue
		}
		if !found || betterCompressedObservation(obs, best) {
			best = obs
			found = true
		}
	}
	return best, found
}

func percentSmaller(smaller, larger int64) float64 {
	if larger <= 0 {
		return 0
	}
	return (float64(larger-smaller) / float64(larger)) * 100
}

func improvementBucketIndex(pct float64) int {
	if pct < 0 || math.IsNaN(pct) || math.IsInf(pct, 0) {
		return 0
	}
	index := int(pct / 10)
	if index > 10 {
		return 10
	}
	return index
}

func improvementBucketLabel(index int) string {
	if index >= 10 {
		return "100%+"
	}
	return fmt.Sprintf("%d-%d%%", index*10, (index+1)*10)
}

func writeZstdPlainRLEDictComparisonTable(b *strings.Builder, summary zstdPlainRLEDictComparison) {
	fmt.Fprintf(b, "- Compared columns: `%d`\n", summary.ComparedColumns)
	fmt.Fprintf(b, "- `zstd + plain` smaller: `%d`; `zstd + rle-dict` smaller: `%d`; ties: `%d`; missing comparisons: `%d`\n\n",
		summary.PlainBetterCount,
		summary.RLEDictBetterCount,
		summary.TieCount,
		summary.MissingCount,
	)
	fmt.Fprintf(b, "| Improvement bucket | `zstd + plain` better | `zstd + rle-dict` better |\n")
	fmt.Fprintf(b, "| --- | ---: | ---: |\n")
	for _, bucket := range summary.Buckets {
		fmt.Fprintf(b, "| `%s` | %d | %d |\n", bucket.Label, bucket.PlainBetter, bucket.RLEDictBetter)
	}
	fmt.Fprintf(b, "\n")
}

func writeZstdPlainRLEDictComparisonSVG(path string, summary zstdPlainRLEDictComparison) error {
	return writeGroupedHistogramSVG(
		path,
		"ZSTD plain vs RLE dictionary improvement",
		"columns",
		summary.Buckets,
	)
}

func buildDeltaBinaryPackedWinnerComparison(byColumn map[string][]columnObservation) deltaBinaryPackedWinnerComparison {
	var improvements []float64
	summary := deltaBinaryPackedWinnerComparison{}
	for _, observations := range byColumn {
		top := topColumnObservations(observations, "overall", 2)
		if len(top) == 0 || top[0].Column.ConfigEncoding != "delta-binary-packed" {
			continue
		}
		summary.WinnerCount++
		if len(top) < 2 {
			summary.MissingSecondBestCount++
			continue
		}
		improvements = append(improvements, percentSmaller(top[0].Column.CompressedBytes, top[1].Column.CompressedBytes))
	}
	summary.Buckets = buildImprovementBuckets(improvements)
	return summary
}

func buildImprovementBuckets(values []float64) []encodingImprovementBucket {
	maxBucket := 1
	for _, value := range values {
		index := improvementBucketIndex(value)
		if index > maxBucket {
			maxBucket = index
		}
	}
	if maxBucket > 10 {
		maxBucket = 10
	}
	buckets := make([]encodingImprovementBucket, maxBucket+1)
	for i := range buckets {
		buckets[i].Label = improvementBucketLabel(i)
	}
	for _, value := range values {
		index := improvementBucketIndex(value)
		if index >= len(buckets) {
			index = len(buckets) - 1
		}
		buckets[index].Count++
	}
	return buckets
}

func writeDeltaBinaryPackedWinnerComparisonTable(b *strings.Builder, summary deltaBinaryPackedWinnerComparison) {
	fmt.Fprintf(b, "- Delta-binary-packed winner columns: `%d`\n", summary.WinnerCount)
	fmt.Fprintf(b, "- Missing second-best rows: `%d`\n\n", summary.MissingSecondBestCount)
	writeImprovementBucketTable(b, "`delta-binary-packed` better than second best", summary.Buckets)
}

func writeDeltaBinaryPackedWinnerComparisonSVG(path string, summary deltaBinaryPackedWinnerComparison) error {
	return writeImprovementBucketChartSVG(
		path,
		"Delta-binary-packed winner improvement over second best",
		"columns where delta-binary-packed is best",
		summary.Buckets,
		"#7c3aed",
	)
}

func buildSnappyPlainRLEDictComparison(byColumn map[string][]columnObservation) snappyPlainRLEDictComparison {
	var rleDictImprovements []float64
	var plainImprovements []float64
	summary := snappyPlainRLEDictComparison{}
	for _, observations := range byColumn {
		plain, plainOK := bestCompressedObservationFor(observations, "snappy", "plain")
		rleDict, rleDictOK := bestCompressedObservationFor(observations, "snappy", "rle-dict")
		if !plainOK || !rleDictOK {
			summary.MissingCount++
			continue
		}
		summary.ComparedColumns++
		plainBytes := plain.Column.CompressedBytes
		rleDictBytes := rleDict.Column.CompressedBytes
		switch {
		case rleDictBytes < plainBytes:
			summary.RLEDictBetterCount++
			rleDictImprovements = append(rleDictImprovements, percentSmaller(rleDictBytes, plainBytes))
		case plainBytes < rleDictBytes:
			summary.PlainBetterCount++
			plainImprovements = append(plainImprovements, percentSmaller(plainBytes, rleDictBytes))
		default:
			summary.TieCount++
		}
	}
	summary.RLEDictBetterBuckets = buildImprovementBuckets(rleDictImprovements)
	summary.PlainBetterBuckets = buildImprovementBuckets(plainImprovements)
	return summary
}

func writeSnappyRLEDictBetterComparisonTable(b *strings.Builder, summary snappyPlainRLEDictComparison) {
	fmt.Fprintf(b, "`snappy + rle-dict` better buckets:\n\n")
	writeImprovementBucketTable(b, "`snappy + rle-dict` better", summary.RLEDictBetterBuckets)
}

func writeSnappyPlainBetterComparisonTable(b *strings.Builder, summary snappyPlainRLEDictComparison) {
	fmt.Fprintf(b, "`snappy + plain` better buckets:\n\n")
	writeImprovementBucketTable(b, "`snappy + plain` better", summary.PlainBetterBuckets)
}

func writeImprovementBucketTable(b *strings.Builder, valueHeader string, buckets []encodingImprovementBucket) {
	fmt.Fprintf(b, "| Improvement bucket | %s |\n", valueHeader)
	fmt.Fprintf(b, "| --- | ---: |\n")
	for _, bucket := range buckets {
		fmt.Fprintf(b, "| `%s` | %d |\n", bucket.Label, bucket.Count)
	}
	fmt.Fprintf(b, "\n")
}

func writeSnappyRLEDictBetterComparisonSVG(path string, summary snappyPlainRLEDictComparison) error {
	return writeImprovementBucketChartSVG(
		path,
		"Snappy RLE dictionary improvement over plain",
		"columns where rle-dict is smaller",
		summary.RLEDictBetterBuckets,
		"#0f766e",
	)
}

func writeSnappyPlainBetterComparisonSVG(path string, summary snappyPlainRLEDictComparison) error {
	return writeImprovementBucketChartSVG(
		path,
		"Snappy plain improvement over RLE dictionary",
		"columns where plain is smaller",
		summary.PlainBetterBuckets,
		"#2563eb",
	)
}

func writeImprovementBucketChartSVG(path, title, xLabel string, buckets []encodingImprovementBucket, color string) error {
	bars := make([]barChartBar, 0, len(buckets))
	for _, bucket := range buckets {
		bars = append(bars, barChartBar{
			Label: bucket.Label,
			Value: bucket.Count,
			Color: color,
		})
	}
	return writeHorizontalBarChartSVG(path, title, xLabel, bars)
}

func compressionColor(compression string) string {
	switch {
	case strings.HasPrefix(compression, "zstd"):
		return "#2563eb"
	case compression == "snappy":
		return "#0f766e"
	default:
		return "#6b7280"
	}
}

func markdownImageTarget(fromDir, path string) string {
	rel, err := filepath.Rel(fromDir, path)
	if err != nil {
		return filepath.ToSlash(path)
	}
	return strings.ReplaceAll(filepath.ToSlash(rel), " ", "%20")
}

func writeRankingList(b *strings.Builder, title string, observations []columnObservation, includeCompression bool) {
	fmt.Fprintf(b, "%s:\n", title)
	if len(observations) == 0 {
		fmt.Fprintf(b, "- No compressed observations.\n\n")
		return
	}
	for i, obs := range observations {
		c := obs.Experiment.Combo
		prefix := fmt.Sprintf("`%s` + `%s`", c.CompressionName, obs.Column.ConfigEncoding)
		improvement := ""
		if obs.HasCompressionRatioImprovementPercent {
			improvement = fmt.Sprintf("; %s vs plain + %s", formatPercent(obs.CompressionRatioImprovementPct), c.CompressionName)
		}
		fmt.Fprintf(b, "%d. %s compressed - %s; %s encoded; %sx post-compression ratio%s; experiment `%s`\n",
			i+1,
			formatByteCount(obs.Column.CompressedBytes),
			prefix,
			formatByteCount(obs.Column.EncodedBytes),
			formatRatio(obs.PostCompressionRatio),
			improvement,
			c.Slug,
		)
	}
	fmt.Fprintf(b, "\n")
}

func topColumnObservations(observations []columnObservation, scope string, limit int) []columnObservation {
	bestByKey := make(map[string]columnObservation)
	for _, obs := range observations {
		c := obs.Experiment.Combo
		if c.Compression == "none" {
			continue
		}
		switch scope {
		case "zstd":
			if !strings.HasPrefix(c.CompressionName, "zstd") {
				continue
			}
		case "snappy":
			if c.CompressionName != "snappy" {
				continue
			}
		}
		key := c.CompressionName + "\x00" + obs.Column.ConfigEncoding
		if scope != "overall" {
			key = obs.Column.ConfigEncoding
		}
		if existing, ok := bestByKey[key]; !ok || betterCompressedObservation(obs, existing) {
			bestByKey[key] = obs
		}
	}

	rows := make([]columnObservation, 0, len(bestByKey))
	for _, obs := range bestByKey {
		rows = append(rows, obs)
	}
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].Column.CompressedBytes != rows[j].Column.CompressedBytes {
			return rows[i].Column.CompressedBytes < rows[j].Column.CompressedBytes
		}
		if rows[i].Column.EncodedBytes != rows[j].Column.EncodedBytes {
			return rows[i].Column.EncodedBytes < rows[j].Column.EncodedBytes
		}
		return rows[i].Experiment.Combo.Slug < rows[j].Experiment.Combo.Slug
	})
	if len(rows) > limit {
		rows = rows[:limit]
	}
	return rows
}

func betterCompressedObservation(left, right columnObservation) bool {
	if left.Column.CompressedBytes != right.Column.CompressedBytes {
		return left.Column.CompressedBytes < right.Column.CompressedBytes
	}
	if left.Column.EncodedBytes != right.Column.EncodedBytes {
		return left.Column.EncodedBytes < right.Column.EncodedBytes
	}
	return left.Experiment.Combo.Slug < right.Experiment.Combo.Slug
}

func writeColumnShapePlots(snapshot columnShapeStatsSnapshot, imagesDir string) (map[string]columnShapePlots, error) {
	if err := os.MkdirAll(imagesDir, 0o755); err != nil {
		return nil, err
	}
	plots := make(map[string]columnShapePlots, len(snapshot.Columns))
	for _, col := range snapshot.Columns {
		slug := sanitizeFilename(col.Name)
		if slug == "" {
			slug = fmt.Sprintf("column-%d", col.ColumnIndex)
		}
		plot := columnShapePlots{
			RowGroupCardinality: filepath.Join(imagesDir, slug+"_row_group_cardinality.svg"),
			PageCardinality:     filepath.Join(imagesDir, slug+"_page_cardinality.svg"),
			PageBounds:          filepath.Join(imagesDir, slug+"_page_bounds.svg"),
			ValueLength:         filepath.Join(imagesDir, slug+"_value_length.svg"),
		}
		if err := writeLinePlotSVG(plot.RowGroupCardinality, col.Name+" row-group cardinality", "distinct values", []plotSeries{
			{Name: "row group cardinality", Color: "#2563eb", Values: rowGroupCardinalityValues(col.RowGroups)},
		}); err != nil {
			return nil, err
		}
		if err := writeLinePlotSVG(plot.PageCardinality, col.Name+" page cardinality per row group", "distinct values", []plotSeries{
			{Name: "min page cardinality", Color: "#0f766e", Values: rowGroupPageCardinalityValues(col.RowGroups, false)},
			{Name: "max page cardinality", Color: "#b91c1c", Values: rowGroupPageCardinalityValues(col.RowGroups, true)},
		}); err != nil {
			return nil, err
		}
		title, yLabel, boundSeries := pageBoundsPlotSeries(col)
		if err := writeLinePlotSVG(plot.PageBounds, title, yLabel, boundSeries); err != nil {
			return nil, err
		}
		if err := writeLinePlotSVG(plot.ValueLength, col.Name+" value length per row group", "bytes", []plotSeries{
			{Name: "min length", Color: "#7c3aed", Values: rowGroupLengthValues(col.RowGroups, false)},
			{Name: "max length", Color: "#ea580c", Values: rowGroupLengthValues(col.RowGroups, true)},
		}); err != nil {
			return nil, err
		}
		plots[col.Name] = plot
	}
	return plots, nil
}

func rowGroupCardinalityValues(rowGroups []shapeRowGroupStats) []float64 {
	values := make([]float64, len(rowGroups))
	for i, rg := range rowGroups {
		values[i] = float64(rg.Cardinality)
	}
	return values
}

func rowGroupPageCardinalityValues(rowGroups []shapeRowGroupStats, max bool) []float64 {
	values := make([]float64, len(rowGroups))
	for i, rg := range rowGroups {
		if max {
			values[i] = float64(rg.PageCardinalityMax)
		} else {
			values[i] = float64(rg.PageCardinalityMin)
		}
	}
	return values
}

func rowGroupLengthValues(rowGroups []shapeRowGroupStats, max bool) []float64 {
	values := make([]float64, len(rowGroups))
	for i, rg := range rowGroups {
		if max {
			values[i] = float64(rg.MaxValueLength)
		} else {
			values[i] = float64(rg.MinValueLength)
		}
	}
	return values
}

func pageBoundsPlotSeries(col columnShapeStats) (string, string, []plotSeries) {
	hasNumeric := false
	for _, page := range col.Pages {
		if page.HasBounds && page.HasNumeric {
			hasNumeric = true
			break
		}
	}
	if hasNumeric {
		mins := make([]float64, 0, len(col.Pages))
		maxes := make([]float64, 0, len(col.Pages))
		for _, page := range col.Pages {
			if !page.HasBounds || !page.HasNumeric {
				continue
			}
			mins = append(mins, page.MinNumeric)
			maxes = append(maxes, page.MaxNumeric)
		}
		return col.Name + " page min/max", "value", []plotSeries{
			{Name: "page min", Color: "#2563eb", Values: mins},
			{Name: "page max", Color: "#b91c1c", Values: maxes},
		}
	}

	ranks := lexicalBoundRanks(col.Pages)
	mins := make([]float64, 0, len(col.Pages))
	maxes := make([]float64, 0, len(col.Pages))
	for _, page := range col.Pages {
		if !page.HasBounds {
			continue
		}
		mins = append(mins, float64(ranks[page.MinValueBytes]))
		maxes = append(maxes, float64(ranks[page.MaxValueBytes]))
	}
	return col.Name + " page min/max lexical rank", "lexical rank", []plotSeries{
		{Name: "page min rank", Color: "#2563eb", Values: mins},
		{Name: "page max rank", Color: "#b91c1c", Values: maxes},
	}
}

func lexicalBoundRanks(pages []shapePageStats) map[string]int {
	seen := make(map[string]struct{})
	for _, page := range pages {
		if !page.HasBounds {
			continue
		}
		seen[page.MinValueBytes] = struct{}{}
		seen[page.MaxValueBytes] = struct{}{}
	}
	values := make([]string, 0, len(seen))
	for value := range seen {
		values = append(values, value)
	}
	sort.Strings(values)
	ranks := make(map[string]int, len(values))
	for i, value := range values {
		ranks[value] = i + 1
	}
	return ranks
}

func writeLinePlotSVG(path, title, yLabel string, series []plotSeries) error {
	const (
		width        = 920
		height       = 300
		leftMargin   = 72
		rightMargin  = 28
		topMargin    = 42
		bottomMargin = 54
	)

	maxPoints := 0
	minY := math.Inf(1)
	maxY := math.Inf(-1)
	for _, s := range series {
		if len(s.Values) > maxPoints {
			maxPoints = len(s.Values)
		}
		for _, value := range s.Values {
			if math.IsNaN(value) || math.IsInf(value, 0) {
				continue
			}
			if value < minY {
				minY = value
			}
			if value > maxY {
				maxY = value
			}
		}
	}
	if maxPoints == 0 || math.IsInf(minY, 0) || math.IsInf(maxY, 0) {
		return writeEmptySVG(path, title, "no data")
	}
	if minY == maxY {
		pad := math.Max(1, math.Abs(minY)*0.05)
		minY -= pad
		maxY += pad
	}

	plotWidth := float64(width - leftMargin - rightMargin)
	plotHeight := float64(height - topMargin - bottomMargin)
	x := func(i int) float64 {
		if maxPoints == 1 {
			return float64(leftMargin) + plotWidth/2
		}
		return float64(leftMargin) + (float64(i) * plotWidth / float64(maxPoints-1))
	}
	y := func(value float64) float64 {
		return float64(topMargin) + ((maxY - value) * plotHeight / (maxY - minY))
	}

	var b strings.Builder
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d">`+"\n", width, height, width, height)
	fmt.Fprintf(&b, `<rect width="100%%" height="100%%" fill="#ffffff"/>`+"\n")
	fmt.Fprintf(&b, `<text x="%d" y="24" font-family="Arial, sans-serif" font-size="16" font-weight="700" fill="#111827">%s</text>`+"\n", leftMargin, html.EscapeString(title))
	fmt.Fprintf(&b, `<text x="%d" y="%d" font-family="Arial, sans-serif" font-size="11" fill="#6b7280">%s</text>`+"\n", leftMargin, height-12, html.EscapeString("page/row-group index"))
	fmt.Fprintf(&b, `<text x="18" y="%d" transform="rotate(-90 18 %d)" font-family="Arial, sans-serif" font-size="11" fill="#6b7280">%s</text>`+"\n", topMargin+int(plotHeight/2), topMargin+int(plotHeight/2), html.EscapeString(yLabel))

	for i := 0; i <= 4; i++ {
		value := minY + (float64(i) * (maxY - minY) / 4)
		yy := y(value)
		fmt.Fprintf(&b, `<line x1="%d" y1="%.2f" x2="%d" y2="%.2f" stroke="#e5e7eb" stroke-width="1"/>`+"\n", leftMargin, yy, width-rightMargin, yy)
		fmt.Fprintf(&b, `<text x="%d" y="%.2f" font-family="Arial, sans-serif" font-size="10" text-anchor="end" fill="#6b7280">%s</text>`+"\n", leftMargin-8, yy+3, formatPlotNumber(value))
	}
	fmt.Fprintf(&b, `<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="#9ca3af" stroke-width="1"/>`+"\n", leftMargin, height-bottomMargin, width-rightMargin, height-bottomMargin)
	fmt.Fprintf(&b, `<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="#9ca3af" stroke-width="1"/>`+"\n", leftMargin, topMargin, leftMargin, height-bottomMargin)

	legendX := leftMargin
	for _, s := range series {
		fmt.Fprintf(&b, `<circle cx="%d" cy="%d" r="4" fill="%s"/>`+"\n", legendX, height-34, s.Color)
		fmt.Fprintf(&b, `<text x="%d" y="%d" font-family="Arial, sans-serif" font-size="11" fill="#374151">%s</text>`+"\n", legendX+8, height-30, html.EscapeString(s.Name))
		legendX += 180
	}

	for _, s := range series {
		if len(s.Values) == 0 {
			continue
		}
		var points strings.Builder
		for i, value := range s.Values {
			if math.IsNaN(value) || math.IsInf(value, 0) {
				continue
			}
			fmt.Fprintf(&points, "%.2f,%.2f ", x(i), y(value))
		}
		fmt.Fprintf(&b, `<polyline fill="none" stroke="%s" stroke-width="1.8" points="%s"/>`+"\n", s.Color, strings.TrimSpace(points.String()))
		if len(s.Values) == 1 {
			fmt.Fprintf(&b, `<circle cx="%.2f" cy="%.2f" r="3" fill="%s"/>`+"\n", x(0), y(s.Values[0]), s.Color)
		}
	}
	fmt.Fprintf(&b, "</svg>\n")
	return os.WriteFile(path, []byte(b.String()), 0o644)
}

func writeHorizontalBarChartSVG(path, title, xLabel string, bars []barChartBar) error {
	if len(bars) == 0 {
		return writeEmptySVG(path, title, "no data")
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	const (
		width        = 920
		leftMargin   = 260
		rightMargin  = 92
		topMargin    = 48
		bottomMargin = 52
		rowHeight    = 34
	)
	height := topMargin + bottomMargin + rowHeight*len(bars)
	maxValue := 0
	for _, bar := range bars {
		if bar.Value > maxValue {
			maxValue = bar.Value
		}
	}
	if maxValue == 0 {
		return writeEmptySVG(path, title, "no data")
	}

	plotWidth := float64(width - leftMargin - rightMargin)
	scale := func(value int) float64 {
		return float64(value) * plotWidth / float64(maxValue)
	}

	var b strings.Builder
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d">`+"\n", width, height, width, height)
	fmt.Fprintf(&b, `<rect width="100%%" height="100%%" fill="#ffffff"/>`+"\n")
	fmt.Fprintf(&b, `<text x="%d" y="26" font-family="Arial, sans-serif" font-size="16" font-weight="700" fill="#111827">%s</text>`+"\n", leftMargin, html.EscapeString(title))
	fmt.Fprintf(&b, `<text x="%d" y="%d" font-family="Arial, sans-serif" font-size="11" fill="#6b7280">%s</text>`+"\n", leftMargin, height-14, html.EscapeString(xLabel))

	plotTop := topMargin - 6
	plotBottom := height - bottomMargin + 4
	for i := 0; i <= 4; i++ {
		value := float64(i) * float64(maxValue) / 4
		x := float64(leftMargin) + value*plotWidth/float64(maxValue)
		fmt.Fprintf(&b, `<line x1="%.2f" y1="%d" x2="%.2f" y2="%d" stroke="#eef2f7" stroke-width="1"/>`+"\n", x, plotTop, x, plotBottom)
		fmt.Fprintf(&b, `<text x="%.2f" y="%d" font-family="Arial, sans-serif" font-size="10" text-anchor="middle" fill="#6b7280">%s</text>`+"\n", x, height-32, formatPlotNumber(value))
	}
	fmt.Fprintf(&b, `<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="#9ca3af" stroke-width="1"/>`+"\n", leftMargin, plotBottom, width-rightMargin, plotBottom)

	for i, bar := range bars {
		y := topMargin + i*rowHeight
		barWidth := scale(bar.Value)
		fmt.Fprintf(&b, `<text x="%d" y="%d" font-family="Arial, sans-serif" font-size="12" text-anchor="end" fill="#374151">%s</text>`+"\n", leftMargin-12, y+19, html.EscapeString(bar.Label))
		fmt.Fprintf(&b, `<rect x="%d" y="%d" width="%.2f" height="18" rx="3" fill="%s"/>`+"\n", leftMargin, y+4, barWidth, bar.Color)
		fmt.Fprintf(&b, `<text x="%.2f" y="%d" font-family="Arial, sans-serif" font-size="12" font-weight="700" fill="#111827">%d</text>`+"\n", float64(leftMargin)+barWidth+8, y+18, bar.Value)
	}
	fmt.Fprintf(&b, "</svg>\n")
	return os.WriteFile(path, []byte(b.String()), 0o644)
}

func writeGroupedHistogramSVG(path, title, yLabel string, buckets []zstdPlainRLEDictBucket) error {
	if len(buckets) == 0 {
		return writeEmptySVG(path, title, "no data")
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	const (
		width        = 920
		height       = 360
		leftMargin   = 64
		rightMargin  = 28
		topMargin    = 62
		bottomMargin = 72
	)

	maxValue := 0
	for _, bucket := range buckets {
		if bucket.PlainBetter > maxValue {
			maxValue = bucket.PlainBetter
		}
		if bucket.RLEDictBetter > maxValue {
			maxValue = bucket.RLEDictBetter
		}
	}
	if maxValue == 0 {
		return writeEmptySVG(path, title, "no directional wins")
	}

	plotWidth := float64(width - leftMargin - rightMargin)
	plotHeight := float64(height - topMargin - bottomMargin)
	groupWidth := plotWidth / float64(len(buckets))
	barGap := 4.0
	barWidth := math.Min(28, (groupWidth-16-barGap)/2)
	if barWidth < 2 {
		barWidth = 2
	}
	y := func(value int) float64 {
		return float64(topMargin) + ((float64(maxValue-value) * plotHeight) / float64(maxValue))
	}
	barHeight := func(value int) float64 {
		return (float64(value) * plotHeight) / float64(maxValue)
	}

	var b strings.Builder
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d">`+"\n", width, height, width, height)
	fmt.Fprintf(&b, `<rect width="100%%" height="100%%" fill="#ffffff"/>`+"\n")
	fmt.Fprintf(&b, `<text x="%d" y="28" font-family="Arial, sans-serif" font-size="16" font-weight="700" fill="#111827">%s</text>`+"\n", leftMargin, html.EscapeString(title))
	fmt.Fprintf(&b, `<rect x="%d" y="42" width="10" height="10" fill="#2563eb"/>`+"\n", leftMargin)
	fmt.Fprintf(&b, `<text x="%d" y="52" font-family="Arial, sans-serif" font-size="11" fill="#374151">zstd + plain better</text>`+"\n", leftMargin+16)
	fmt.Fprintf(&b, `<rect x="%d" y="42" width="10" height="10" fill="#0f766e"/>`+"\n", leftMargin+160)
	fmt.Fprintf(&b, `<text x="%d" y="52" font-family="Arial, sans-serif" font-size="11" fill="#374151">zstd + rle-dict better</text>`+"\n", leftMargin+176)
	fmt.Fprintf(&b, `<text x="18" y="%d" transform="rotate(-90 18 %d)" font-family="Arial, sans-serif" font-size="11" fill="#6b7280">%s</text>`+"\n", topMargin+int(plotHeight/2), topMargin+int(plotHeight/2), html.EscapeString(yLabel))

	for i := 0; i <= 4; i++ {
		value := int(math.Round(float64(i) * float64(maxValue) / 4))
		yy := y(value)
		fmt.Fprintf(&b, `<line x1="%d" y1="%.2f" x2="%d" y2="%.2f" stroke="#e5e7eb" stroke-width="1"/>`+"\n", leftMargin, yy, width-rightMargin, yy)
		fmt.Fprintf(&b, `<text x="%d" y="%.2f" font-family="Arial, sans-serif" font-size="10" text-anchor="end" fill="#6b7280">%d</text>`+"\n", leftMargin-8, yy+3, value)
	}
	fmt.Fprintf(&b, `<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="#9ca3af" stroke-width="1"/>`+"\n", leftMargin, height-bottomMargin, width-rightMargin, height-bottomMargin)
	fmt.Fprintf(&b, `<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="#9ca3af" stroke-width="1"/>`+"\n", leftMargin, topMargin, leftMargin, height-bottomMargin)

	for i, bucket := range buckets {
		center := float64(leftMargin) + float64(i)*groupWidth + groupWidth/2
		plainX := center - barGap/2 - barWidth
		rleDictX := center + barGap/2
		plainHeight := barHeight(bucket.PlainBetter)
		rleDictHeight := barHeight(bucket.RLEDictBetter)
		plainY := float64(height-bottomMargin) - plainHeight
		rleDictY := float64(height-bottomMargin) - rleDictHeight

		fmt.Fprintf(&b, `<rect x="%.2f" y="%.2f" width="%.2f" height="%.2f" rx="2" fill="#2563eb"/>`+"\n", plainX, plainY, barWidth, plainHeight)
		fmt.Fprintf(&b, `<rect x="%.2f" y="%.2f" width="%.2f" height="%.2f" rx="2" fill="#0f766e"/>`+"\n", rleDictX, rleDictY, barWidth, rleDictHeight)
		if bucket.PlainBetter > 0 {
			fmt.Fprintf(&b, `<text x="%.2f" y="%.2f" font-family="Arial, sans-serif" font-size="10" text-anchor="middle" fill="#111827">%d</text>`+"\n", plainX+barWidth/2, plainY-4, bucket.PlainBetter)
		}
		if bucket.RLEDictBetter > 0 {
			fmt.Fprintf(&b, `<text x="%.2f" y="%.2f" font-family="Arial, sans-serif" font-size="10" text-anchor="middle" fill="#111827">%d</text>`+"\n", rleDictX+barWidth/2, rleDictY-4, bucket.RLEDictBetter)
		}
		fmt.Fprintf(&b, `<text x="%.2f" y="%d" font-family="Arial, sans-serif" font-size="10" text-anchor="middle" fill="#374151">%s</text>`+"\n", center, height-bottomMargin+20, html.EscapeString(bucket.Label))
	}
	fmt.Fprintf(&b, `<text x="%d" y="%d" font-family="Arial, sans-serif" font-size="11" fill="#6b7280">compressed-byte improvement bucket</text>`+"\n", leftMargin, height-14)
	fmt.Fprintf(&b, "</svg>\n")
	return os.WriteFile(path, []byte(b.String()), 0o644)
}

func writeEmptySVG(path, title, message string) error {
	const width, height = 920, 180
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	var b strings.Builder
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d">`+"\n", width, height, width, height)
	fmt.Fprintf(&b, `<rect width="100%%" height="100%%" fill="#ffffff"/>`+"\n")
	fmt.Fprintf(&b, `<text x="36" y="42" font-family="Arial, sans-serif" font-size="16" font-weight="700" fill="#111827">%s</text>`+"\n", html.EscapeString(title))
	fmt.Fprintf(&b, `<text x="36" y="88" font-family="Arial, sans-serif" font-size="13" fill="#6b7280">%s</text>`+"\n", html.EscapeString(message))
	fmt.Fprintf(&b, "</svg>\n")
	return os.WriteFile(path, []byte(b.String()), 0o644)
}

func summarizeInt64(values []int64) string {
	if len(values) == 0 {
		return ""
	}
	values = append([]int64(nil), values...)
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	return fmt.Sprintf("%s / %s / %s", formatCount(values[0]), formatCount(values[len(values)/2]), formatCount(values[len(values)-1]))
}

func summarizeInt(values []int) string {
	if len(values) == 0 {
		return ""
	}
	values = append([]int(nil), values...)
	sort.Ints(values)
	return fmt.Sprintf("%s / %s / %s", formatCount(int64(values[0])), formatCount(int64(values[len(values)/2])), formatCount(int64(values[len(values)-1])))
}

func formatByteCount(n int64) string {
	return fmt.Sprintf("%s B (%s)", formatCount(n), formatBytes(n))
}

func formatCount(n int64) string {
	s := strconv.FormatInt(n, 10)
	if len(s) <= 3 {
		return s
	}
	var b strings.Builder
	prefix := len(s) % 3
	if prefix == 0 {
		prefix = 3
	}
	b.WriteString(s[:prefix])
	for i := prefix; i < len(s); i += 3 {
		b.WriteByte(',')
		b.WriteString(s[i : i+3])
	}
	return b.String()
}

func formatPercent(value float64) string {
	return fmt.Sprintf("%.6f%%", value)
}

func formatPlotNumber(value float64) string {
	abs := math.Abs(value)
	if abs >= 1_000_000 || (abs > 0 && abs < 0.01) {
		return fmt.Sprintf("%.3g", value)
	}
	if math.Abs(value-math.Round(value)) < 0.001 {
		return fmt.Sprintf("%.0f", value)
	}
	return fmt.Sprintf("%.2f", value)
}

func markdownLinkTarget(fromDir, targetPath string) string {
	rel, err := filepath.Rel(fromDir, targetPath)
	if err != nil {
		return filepath.ToSlash(targetPath)
	}
	return filepath.ToSlash(rel)
}

func settingMarkdownCell(ranking experimentRanking, compressed bool) string {
	if compressed {
		return fmt.Sprintf("`%d` bytes, `%sx`", ranking.Result.CompressedBytes, formatRatio(ranking.PostCompressionRatio))
	}
	return fmt.Sprintf("`%d` bytes, `%sx`", ranking.Result.EncodedBytes, formatRatio(ranking.PostEncodingRatio))
}

func settingCell(ranking experimentRanking, compressed bool) string {
	if compressed {
		return fmt.Sprintf("compressed_bytes=%d post_compression_ratio=%s codec_ratio=%s", ranking.Result.CompressedBytes, formatRatio(ranking.PostCompressionRatio), formatRatio(ranking.CodecRatio))
	}
	return fmt.Sprintf("encoded_bytes=%d post_encoding_ratio=%s", ranking.Result.EncodedBytes, formatRatio(ranking.PostEncodingRatio))
}

func writeTSV(path string, header []string, writeRows func(*csv.Writer) error) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	w.Comma = '\t'
	if err := w.Write(header); err != nil {
		return err
	}
	if err := writeRows(w); err != nil {
		return err
	}
	w.Flush()
	return w.Error()
}

func firstFailure(results []experimentResult) error {
	for _, result := range results {
		if result.Err != nil {
			return result.Err
		}
	}
	return nil
}

func ratio(numerator, denominator int64) float64 {
	if denominator == 0 {
		return 0
	}
	return float64(numerator) / float64(denominator)
}

func formatRatio(value float64) string {
	if value == 0 {
		return ""
	}
	return fmt.Sprintf("%.6f", value)
}

func optionalRatio(value float64, ok bool) string {
	if !ok {
		return ""
	}
	return formatRatio(value)
}

func optionalPercent(value float64, ok bool) string {
	if !ok {
		return ""
	}
	return fmt.Sprintf("%.6f", value)
}

func optionalPercentMarkdown(value float64, ok bool) string {
	if !ok {
		return ""
	}
	return fmt.Sprintf("%.6f%%", value)
}

func optionalInt(value int64, ok bool) string {
	if !ok {
		return ""
	}
	return strconv.FormatInt(value, 10)
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

func exitf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
