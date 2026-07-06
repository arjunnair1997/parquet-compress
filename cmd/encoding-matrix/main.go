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
	Rows                          int64
	Parallel                      int
	ZstdLevel                     int
	Input                         string
	MaxDictSize                   string
	ExperimentDir                 string
	ResultDir                     string
	MarkdownDir                   string
	ConfigDir                     string
	TSVDir                        string
	OutputRoot                    string
	ShapeStatsJSON                string
	Verify                        bool
	SkipExisting                  bool
	RefreshMissingDictionaryStats bool
	KeepOutput                    bool
	GeneratePDF                   bool
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
	Column                                  string
	Type                                    string
	ConfigEncoding                          string
	MetadataEncodings                       string
	PageEncodings                           string
	Values                                  int64
	PhysicalBytes                           int64
	EncodedBytes                            int64
	CompressedBytes                         int64
	DictionaryPageCount                     int64
	DictionaryPageCompressedBytes           int64
	CompressedBytesWithoutDictionaryPages   int64
	HasCompressedBytesWithoutDictionaryPage bool
	SourceFieldBytes                        int64
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
	Experiment               experimentResult
	Column                   columnResult
	BaselineEncodedBytes     int64
	PlainCompressedBytes     int64
	PlainCompressionRatio    float64
	PostEncodingRatio        float64
	PostCompressionRatio     float64
	CodecRatio               float64
	TargetBytes              int64
	TargetMetric             string
	HasPostCompressionRatio  bool
	HasPlainCompressedBytes  bool
	HasPlainCompressionRatio bool
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
	ColumnIndex       int                  `json:"column_index"`
	Path              []string             `json:"path"`
	Name              string               `json:"name"`
	Type              string               `json:"type"`
	PhysicalType      string               `json:"physical_type"`
	SortedAscending   bool                 `json:"sorted_ascending"`
	SortedDescending  bool                 `json:"sorted_descending"`
	MinValueLength    int                  `json:"min_value_length"`
	MedianValueLength float64              `json:"median_value_length"`
	MaxValueLength    int                  `json:"max_value_length"`
	RowGroups         []shapeRowGroupStats `json:"row_groups"`
	Pages             []shapePageStats     `json:"pages"`
}

type shapeRowGroupStats struct {
	RowGroupIndex                       int64   `json:"row_group_index"`
	FirstRowIndex                       int64   `json:"first_row_index"`
	NumRows                             int64   `json:"num_rows"`
	Cardinality                         int64   `json:"cardinality"`
	PageCount                           int     `json:"page_count"`
	PageCardinalityMin                  int64   `json:"page_cardinality_min"`
	PageCardinalityMax                  int64   `json:"page_cardinality_max"`
	MinValueLength                      int     `json:"min_value_length"`
	MedianValueLength                   float64 `json:"median_value_length"`
	MaxValueLength                      int     `json:"max_value_length"`
	EncodedDataPageBytesBeforeCodec     int64   `json:"encoded_data_page_bytes_before_codec"`
	CompressedDataPageBytesAfterCodec   int64   `json:"compressed_data_page_bytes_after_codec"`
	DictionaryPageCount                 int     `json:"dictionary_page_count"`
	DictionaryEncodedBytesBeforeCodec   int64   `json:"dictionary_encoded_bytes_before_codec"`
	DictionaryCompressedBytesAfterCodec int64   `json:"dictionary_compressed_bytes_after_codec"`
	AmortizedDictionaryEncodedBytes     float64 `json:"amortized_dictionary_encoded_bytes_per_data_page"`
	AmortizedDictionaryCompressedBytes  float64 `json:"amortized_dictionary_compressed_bytes_per_data_page"`
	EncodedBytesWithDictionary          int64   `json:"encoded_bytes_with_dictionary"`
	CompressedBytesWithDictionary       int64   `json:"compressed_bytes_with_dictionary"`
}

type shapePageStats struct {
	RowGroupIndex                              int64   `json:"row_group_index"`
	RowGroupFirstRowIndex                      int64   `json:"row_group_first_row_index"`
	PageIndex                                  int     `json:"page_index"`
	FirstRowIndex                              int64   `json:"first_row_index"`
	AbsoluteFirstRowIndex                      int64   `json:"absolute_first_row_index"`
	NumRows                                    int64   `json:"num_rows"`
	NumValues                                  int64   `json:"num_values"`
	Cardinality                                int64   `json:"cardinality"`
	PageType                                   string  `json:"page_type"`
	Encoding                                   string  `json:"encoding"`
	EncodingID                                 int32   `json:"encoding_id"`
	HeaderBytes                                int64   `json:"header_bytes"`
	EncodedBodyBytesBeforeCodec                int64   `json:"encoded_body_bytes_before_codec"`
	CompressedBodyBytesAfterCodec              int64   `json:"compressed_body_bytes_after_codec"`
	EncodedPageBytesBeforeCodec                int64   `json:"encoded_page_bytes_before_codec"`
	CompressedPageBytesAfterCodec              int64   `json:"compressed_page_bytes_after_codec"`
	DataPageCountInColumnChunk                 int     `json:"data_page_count_in_column_chunk"`
	DictionaryPageCount                        int     `json:"dictionary_page_count"`
	DictionaryEncodedBytesBeforeCodec          int64   `json:"dictionary_encoded_bytes_before_codec"`
	DictionaryCompressedBytesAfterCodec        int64   `json:"dictionary_compressed_bytes_after_codec"`
	AmortizedDictionaryEncodedBytes            float64 `json:"amortized_dictionary_encoded_bytes"`
	AmortizedDictionaryCompressedBytes         float64 `json:"amortized_dictionary_compressed_bytes"`
	EncodedPageBytesWithAmortizedDictionary    float64 `json:"encoded_page_bytes_with_amortized_dictionary"`
	CompressedPageBytesWithAmortizedDictionary float64 `json:"compressed_page_bytes_with_amortized_dictionary"`
	HasBounds                                  bool    `json:"has_bounds"`
	MinValue                                   string  `json:"min_value,omitempty"`
	MaxValue                                   string  `json:"max_value,omitempty"`
	MinValueBytes                              string  `json:"min_value_bytes,omitempty"`
	MaxValueBytes                              string  `json:"max_value_bytes,omitempty"`
	HasNumeric                                 bool    `json:"has_numeric"`
	MinNumeric                                 float64 `json:"min_numeric,omitempty"`
	MaxNumeric                                 float64 `json:"max_numeric,omitempty"`
	MinLength                                  int     `json:"min_length"`
	MaxLength                                  int     `json:"max_length"`
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

type pageEncodingDistribution struct {
	Compression           string
	TSVPath               string
	ImagePath             string
	PlainWinsImagePath    string
	Columns               []pageEncodingDistributionColumn
	ComparedColumns       int
	MixedColumns          int
	PlainOnlyColumns      int
	RLEDictOnlyColumns    int
	TieOnlyColumns        int
	ComparisonWindows     int64
	RowsCompared          int64
	PlainWindowWins       int64
	RLEDictWindowWins     int64
	TieWindowWins         int64
	PlainRowsWon          int64
	RLEDictRowsWon        int64
	TieRowsWon            int64
	PlainAllocatedBytes   int64
	RLEDictAllocatedBytes int64
	RLEDictNoDictBytes    int64
	ExactMatchedPages     int64
	UnmatchedPlainPages   int64
	UnmatchedRLEDictPages int64
	DictOverheadFlipWins  int64
	DictOverheadFlipRows  int64
}

type pageEncodingDistributionColumn struct {
	Column                 string
	Type                   string
	PlainPages             int64
	RLEDictPages           int64
	ComparisonWindows      int64
	RowsCompared           int64
	PlainWindowWins        int64
	RLEDictWindowWins      int64
	TieWindowWins          int64
	PlainRowsWon           int64
	RLEDictRowsWon         int64
	TieRowsWon             int64
	PlainAllocatedBytes    int64
	RLEDictAllocatedBytes  int64
	RLEDictNoDictBytes     int64
	UncompressedBytes      int64
	PlainToUncompressed    float64
	RLEDictToUncompressed  float64
	RLEDictRatioAdvantage  float64
	AbsoluteRatioDiff      float64
	RLEDictToPlainRatio    float64
	RLEDictNoDictRatio     float64
	DictOverheadFlipWins   int64
	DictOverheadFlipRows   int64
	ExactMatchedPages      int64
	UnmatchedPlainPages    int64
	UnmatchedRLEDictPages  int64
	WinnerByAllocatedBytes string
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

type rleDictWorseCategoryComparison struct {
	Categories         []rleDictWorseCategory
	ComparedColumns    int
	RLEDictWorseCount  int
	RLEDictBetterCount int
	TieCount           int
	MissingCount       int
	MissingShapeStats  int
}

type rleDictWorseCategory struct {
	Name        string
	Slug        string
	Description string
	Rows        []rleDictWorseColumn
	Buckets     []encodingImprovementBucket
	MinPct      float64
	MedianPct   float64
	MaxPct      float64
}

type rleDictWorseColumn struct {
	Column                                  string
	Type                                    string
	Category                                string
	MeasuredFeature                         string
	MeasuredReason                          string
	RowGroupCardinality                     string
	MedianCardinalityRatio                  float64
	ValueLengthMin                          string
	ValueLengthMedian                       string
	ValueLengthMax                          string
	HasShapeStats                           bool
	PhysicalBytes                           int64
	BaselineEncodedBytes                    int64
	PlainBytes                              int64
	RLEDictBytes                            int64
	RLEDictBytesWithoutDictionaryPages      int64
	DictionaryPageCount                     int64
	HasCompressedBytesWithoutDictionaryPage bool
	WorseByPct                              float64
}

type snappyRLEDictWorseCategoryComparison struct {
	Categories         []snappyRLEDictWorseCategory
	ComparedColumns    int
	RLEDictWorseCount  int
	RLEDictBetterCount int
	TieCount           int
	MissingCount       int
	MissingShapeStats  int
}

type snappyRLEDictWorseCategory struct {
	Name        string
	Slug        string
	Description string
	Rows        []snappyRLEDictWorseColumn
	Buckets     []encodingImprovementBucket
	MinPct      float64
	MedianPct   float64
	MaxPct      float64
}

type snappyRLEDictWorseColumn struct {
	Column                                  string
	Type                                    string
	Category                                string
	MeasuredFeature                         string
	MeasuredReason                          string
	RowGroupCardinality                     string
	MedianCardinalityRatio                  float64
	ValueLengthMin                          string
	ValueLengthMedian                       string
	ValueLengthMax                          string
	HasShapeStats                           bool
	PhysicalBytes                           int64
	BaselineEncodedBytes                    int64
	RLEDictEncodedBytes                     int64
	PlainBytes                              int64
	RLEDictBytes                            int64
	RLEDictBytesWithoutDictionaryPages      int64
	DictionaryPageCount                     int64
	HasCompressedBytesWithoutDictionaryPage bool
	WorseByPct                              float64
}

type encodingImprovementBucket struct {
	Label string
	Count int
}

type overallAbsoluteDifferenceComparison struct {
	Compression        string
	RLEDictBetterRows  []overallAbsoluteDifferenceRow
	PlainBetterRows    []overallAbsoluteDifferenceRow
	ComparedColumns    int
	RLEDictBetterCount int
	PlainBetterCount   int
	TieCount           int
	MissingCount       int
}

type overallAbsoluteDifferenceRow struct {
	Column                 string
	Type                   string
	HasShapeStats          bool
	CardinalityRatioMin    float64
	CardinalityRatioMedian float64
	CardinalityRatioMax    float64
	UncompressedBytes      int64
	PlainBytes             int64
	RLEDictBytes           int64
	PlainRatio             float64
	RLEDictRatio           float64
	AbsoluteDifference     float64
}

type compressionRatioColorBucket struct {
	Label string
	Min   float64
	Max   float64
	Color string
}

type encodingRankDistribution struct {
	Rows    []encodingRankDistributionRow
	MaxRank int
}

type encodingRankDistributionRow struct {
	Compression string
	Encoding    string
	RankCounts  []int
	Total       int
}

type secondPlaceDistribution struct {
	Rows               []secondPlaceDistributionRow
	FirstPlaceEncoding string
	FirstPlaceCount    int
	MissingSecondPlace int
}

type secondPlaceDistributionRow struct {
	Encoding string
	Count    int
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
	fs.BoolVar(&cfg.RefreshMissingDictionaryStats, "refresh-missing-dictionary-stats", cfg.RefreshMissingDictionaryStats, "with --skip-existing, rerun compressed rle-dict configs whose column TSV lacks dictionary page byte stats")
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
		refreshMissingDictionaryStats, err := shouldRefreshMissingDictionaryStats(cfg, c, existingColumns)
		if err != nil {
			result.Elapsed = time.Since(started)
			result.Err = err
			return result
		}
		if refreshMissingDictionaryStats {
			fmt.Printf("refreshing %s because existing column TSV lacks dictionary page byte stats\n", c.Slug)
		} else {
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

func shouldRefreshMissingDictionaryStats(cfg config, c combo, columnTSVPath string) (bool, error) {
	if !cfg.RefreshMissingDictionaryStats || c.Compression == "none" || !comboUsesRLEDictionary(c) {
		return false, nil
	}
	hasStats, err := columnStatsTSVHasDictionaryPageBytes(columnTSVPath)
	if err != nil {
		return false, err
	}
	return !hasStats, nil
}

func comboUsesRLEDictionary(c combo) bool {
	return c.IntEncoding == "rle-dict" ||
		c.DateEncoding == "rle-dict" ||
		c.TimestampEncoding == "rle-dict" ||
		c.StringEncoding == "rle-dict"
}

func columnStatsTSVHasDictionaryPageBytes(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = '\t'
	header, err := reader.Read()
	if err != nil {
		return false, err
	}
	seen := make(map[string]bool, len(header))
	for _, name := range header {
		seen[name] = true
	}
	return seen["dictionary_page_compressed_bytes"] && seen["compressed_bytes_without_dictionary_pages"], nil
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
			Column:                        field(record, index, "column"),
			Type:                          field(record, index, "type"),
			ConfigEncoding:                field(record, index, "config_encoding"),
			MetadataEncodings:             field(record, index, "metadata_encodings"),
			PageEncodings:                 field(record, index, "page_encodings"),
			Values:                        parseInt(field(record, index, "values")),
			PhysicalBytes:                 parseInt(field(record, index, "physical_bytes")),
			EncodedBytes:                  parseInt(field(record, index, "encoded_bytes")),
			CompressedBytes:               parseInt(field(record, index, "compressed_bytes")),
			DictionaryPageCount:           parseInt(field(record, index, "dictionary_page_count")),
			DictionaryPageCompressedBytes: parseInt(field(record, index, "dictionary_page_compressed_bytes")),
			SourceFieldBytes:              parseInt(field(record, index, "source_field_bytes")),
		}
		if col.DictionaryPageCount == 0 {
			col.DictionaryPageCount = dictionaryPageCountFromPageEncodings(col.PageEncodings)
		}
		if value := field(record, index, "compressed_bytes_without_dictionary_pages"); value != "" {
			col.CompressedBytesWithoutDictionaryPages = parseInt(value)
			col.HasCompressedBytesWithoutDictionaryPage = true
		} else if col.DictionaryPageCompressedBytes != 0 {
			col.CompressedBytesWithoutDictionaryPages = col.CompressedBytes - col.DictionaryPageCompressedBytes
			col.HasCompressedBytesWithoutDictionaryPage = true
		} else {
			col.CompressedBytesWithoutDictionaryPages = col.CompressedBytes
		}
		columns = append(columns, col)
	}
	return columns, nil
}

func dictionaryPageCountFromPageEncodings(pageEncodings string) int64 {
	for _, part := range strings.Split(pageEncodings, ",") {
		part = strings.TrimSpace(part)
		if !strings.HasPrefix(part, "DICTIONARY_PAGE/") {
			continue
		}
		colon := strings.LastIndex(part, ":")
		if colon < 0 || colon == len(part)-1 {
			continue
		}
		return parseInt(part[colon+1:])
	}
	return 0
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

func parseFloat(s string) float64 {
	n, _ := strconv.ParseFloat(s, 64)
	return n
}

func parseRoundedFloat(s string) int64 {
	return int64(math.Round(parseFloat(s)))
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

func isRLEDictEncodingCombo(c combo) bool {
	return c.IntEncoding == "rle-dict" &&
		c.DateEncoding == "rle-dict" &&
		c.TimestampEncoding == "rle-dict" &&
		c.StringEncoding == "rle-dict"
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
		"dictionary_page_count",
		"dictionary_page_compressed_bytes",
		"compressed_bytes_without_dictionary_pages",
		"post_compression_no_encoding_bytes",
		"post_compression_no_encoding_ratio",
		"target_metric",
		"target_bytes",
		"post_encoding_ratio",
		"post_compression_ratio",
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
				strconv.FormatInt(obs.Column.DictionaryPageCount, 10),
				strconv.FormatInt(obs.Column.DictionaryPageCompressedBytes, 10),
				optionalInt(obs.Column.CompressedBytesWithoutDictionaryPages, obs.Column.HasCompressedBytesWithoutDictionaryPage),
				optionalInt(obs.PlainCompressedBytes, obs.HasPlainCompressedBytes),
				optionalRatio(obs.PlainCompressionRatio, obs.HasPlainCompressionRatio),
				obs.TargetMetric,
				strconv.FormatInt(obs.TargetBytes, 10),
				formatRatio(obs.PostEncodingRatio),
				optionalRatio(obs.PostCompressionRatio, obs.HasPostCompressionRatio),
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

type markdownDoc struct {
	b            *strings.Builder
	anchorCounts map[string]int
}

func newMarkdownDoc(b *strings.Builder) *markdownDoc {
	return &markdownDoc{
		b:            b,
		anchorCounts: make(map[string]int),
	}
}

func (d *markdownDoc) Heading(level int, title string) {
	if level < 1 {
		level = 1
	}
	if level > 6 {
		level = 6
	}
	id := d.nextAnchorID(title)
	fmt.Fprintf(d.b, "<a id=\"%s\"></a>\n%s %s [#](#%s)\n\n", id, strings.Repeat("#", level), title, id)
}

func (d *markdownDoc) nextAnchorID(title string) string {
	base := markdownAnchorSlug(title)
	d.anchorCounts[base]++
	if d.anchorCounts[base] == 1 {
		return base
	}
	return fmt.Sprintf("%s-%d", base, d.anchorCounts[base])
}

func markdownAnchorSlug(title string) string {
	title = strings.ToLower(strings.TrimSpace(title))
	var b strings.Builder
	lastDash := false
	for _, r := range title {
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
	slug := strings.Trim(b.String(), "-")
	if slug == "" {
		return "section"
	}
	return slug
}

func writeSummaryMarkdown(path string, cfg config, rankings []experimentRanking, settings []settingSummary, winners []columnWinner, bestColumnWinners []columnWinner, baseline experimentResult, started, finished time.Time, allExperimentsPath, settingsPath, columnResultsPath, columnWinnersPath, bestColumnEncodingsPath, columnTop5Path string) error {
	var b strings.Builder
	md := newMarkdownDoc(&b)
	md.Heading(1, "Encoding Matrix Summary")
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

	md.Heading(2, "Outputs")
	fmt.Fprintf(&b, "- All experiments: [%s](%s)\n", filepath.Base(allExperimentsPath), markdownLinkTarget(summaryDir, allExperimentsPath))
	fmt.Fprintf(&b, "- Settings with pre/post compression side by side: [%s](%s)\n", filepath.Base(settingsPath), markdownLinkTarget(summaryDir, settingsPath))
	fmt.Fprintf(&b, "- All per-column observations: [%s](%s)\n", filepath.Base(columnResultsPath), markdownLinkTarget(summaryDir, columnResultsPath))
	fmt.Fprintf(&b, "- Per-column winners by scope: [%s](%s)\n", filepath.Base(columnWinnersPath), markdownLinkTarget(summaryDir, columnWinnersPath))
	fmt.Fprintf(&b, "- Best encoding per column: [%s](%s)\n", filepath.Base(bestColumnEncodingsPath), markdownLinkTarget(summaryDir, bestColumnEncodingsPath))
	fmt.Fprintf(&b, "- Column top 5 rankings with shape stats: [%s](%s)\n\n", filepath.Base(columnTop5Path), markdownLinkTarget(summaryDir, columnTop5Path))

	md.Heading(2, "Ranking Definitions")
	fmt.Fprintf(&b, "- Pre-compression uses the `none` run for the same encoding setting: plain/uncompressed baseline encoded bytes divided by that setting's encoded bytes.\n")
	fmt.Fprintf(&b, "- Snappy and ZSTD compression use the compressed bytes for the same encoding setting: plain/uncompressed baseline encoded bytes divided by compressed bytes.\n")
	fmt.Fprintf(&b, "- Column ratios use `baseline_encoded_bytes` as their denominator: the same column's encoded bytes from the all-plain/no-compression run. `physical_bytes` is shown separately and is not used as the ratio denominator.\n")
	fmt.Fprintf(&b, "- ZSTD compression without encoding is the all-plain ZSTD run, repeated in each Top Encoding Settings row as the zstd-only baseline.\n")
	fmt.Fprintf(&b, "- `post_compression_no_encoding_bytes` is the same column's compressed bytes from the all-plain run with the same compression setting; `post_compression_no_encoding_ratio` is plain/uncompressed baseline encoded bytes divided by those bytes.\n")
	fmt.Fprintf(&b, "- Codec ratio: candidate encoded bytes divided by candidate compressed bytes.\n\n")

	md.Heading(2, "Top Encoding Settings")
	writeSettingTable(&b, settings, 20, summaryDir)

	md.Heading(2, "Column Winners")
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
	fmt.Fprintf(b, "| Column | Type | Best encoding | Compression | Physical bytes | Baseline encoded bytes | Target bytes | Encoded bytes | Compressed bytes | Post-compression no encoding bytes | Post-compression no encoding | Post-encoding | Post-compression | Result |\n")
	fmt.Fprintf(b, "| --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | --- |\n")
	for _, winner := range winners {
		obs := winner.Observation
		c := obs.Experiment.Combo
		fmt.Fprintf(b, "| `%s` | `%s` | `%s` | `%s` | `%d` | `%d` | `%d` | `%d` | `%d` | `%s` | `%s` | `%s` | `%s` | [%s](%s) |\n",
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
	md := newMarkdownDoc(&b)
	md.Heading(1, "Column Top 5 Encoding Rankings")
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
	md.Heading(2, "Winner Distribution")
	fmt.Fprintf(&b, "Counts are based on each column's first `Compressed overall` ranking below: one winner per column, grouped by compression algorithm and configured column encoding.\n\n")
	writeShapeImage(&b, "Column winner distribution", winnerDistributionPath, reportDir)
	writeWinnerDistributionTable(&b, winnerDistribution)

	rankDistribution := buildEncodingRankDistribution(byColumn)
	rankDistributionPath := filepath.Join(reportDir, "images", "encoding_rank_distribution.svg")
	if err := writeEncodingRankDistributionSVG(rankDistributionPath, rankDistribution); err != nil {
		return err
	}
	md.Heading(2, "Encoding Rank Distribution")
	fmt.Fprintf(&b, "For each column and compression codec, duplicate matrix rows with the same effective column encoding are collapsed to the smallest compressed byte count. The remaining encodings are sorted by compressed bytes; counts below show how often each compression + encoding landed at rank 1, rank 2, and so on. Encodings that are not valid for a column type are not counted for that column.\n\n")
	writeShapeImage(&b, "Encoding rank distribution by compression", rankDistributionPath, reportDir)
	writeEncodingRankDistributionTable(&b, rankDistribution)

	zstdPlainSecondPlace := buildZstdPlainWinnerSecondPlaceDistribution(byColumn)
	zstdPlainSecondPlacePath := filepath.Join(reportDir, "images", "zstd_plain_winner_second_place_distribution.svg")
	if err := writeZstdPlainWinnerSecondPlaceDistributionSVG(zstdPlainSecondPlacePath, zstdPlainSecondPlace); err != nil {
		return err
	}
	md.Heading(2, "ZSTD Plain Winner Second-Place Distribution")
	fmt.Fprintf(&b, "For columns where `zstd + plain` is rank 1 in the ZSTD-only compressed-byte ranking, this counts which encoding landed at rank 2 after collapsing duplicate matrix rows to each encoding's smallest compressed byte count.\n\n")
	writeShapeImage(&b, "ZSTD plain winner second-place distribution", zstdPlainSecondPlacePath, reportDir)
	writeSecondPlaceDistributionTable(&b, zstdPlainSecondPlace)

	zstdComparison := buildZstdPlainRLEDictComparison(byColumn)
	zstdComparisonPath := filepath.Join(reportDir, "images", "zstd_plain_vs_rle_dict_improvement.svg")
	if err := writeZstdPlainRLEDictComparisonSVG(zstdComparisonPath, zstdComparison); err != nil {
		return err
	}
	md.Heading(2, "ZSTD Plain vs RLE Dict Improvement Distribution")
	fmt.Fprintf(&b, "For each column, this compares the best observed `zstd + plain` compressed byte count with the best observed `zstd + rle-dict` compressed byte count. Improvement is `(larger compressed bytes - smaller compressed bytes) / larger compressed bytes * 100`.\n\n")
	writeShapeImage(&b, "ZSTD plain versus RLE dictionary improvement distribution", zstdComparisonPath, reportDir)
	writeZstdPlainRLEDictComparisonTable(&b, zstdComparison)
	writeOverallAbsoluteDifferenceComparison(md, buildOverallPlainRLEDictAbsoluteDifference(byColumn, shapeByColumn, "zstd"))
	pageDistribution, err := loadPageEncodingDistribution(cfg, "zstd")
	if err != nil {
		return err
	}
	if pageDistribution != nil {
		writePageEncodingDistributionMarkdown(md, *pageDistribution, reportDir)
	}

	rleDictWorseCategories := buildZstdRLEDictWorseCategoryComparison(byColumn, shapeByColumn)
	if err := writeRLEDictWorseCategoryImages(reportDir, &rleDictWorseCategories); err != nil {
		return err
	}
	md.Heading(2, "ZSTD RLE Dict Worse Distribution By Category")
	fmt.Fprintf(&b, "For columns where the best observed `zstd + plain` compressed byte count is smaller than the best observed `zstd + rle-dict` compressed byte count, each category image plots `plain + zstd` compressed bytes on the x-axis and `rle-dict + zstd` compressed bytes on the y-axis using the same log byte scale. Points above the diagonal are larger with RLE dictionary encoding. Point color is bucketed by `plain/no-compression encoded bytes / rle-dict + zstd compressed bytes`, so high-ratio colors identify columns where RLE dictionary lost the head-to-head but still compressed the baseline dramatically.\n\n")
	fmt.Fprintf(&b, "The bucket tables below each image still show how much worse RLE dictionary encoding was. Worse-by percentage is `(rle_dict_compressed_bytes / plain_compressed_bytes - 1) * 100`, so values can exceed 100%%.\n\n")
	fmt.Fprintf(&b, "The compressed bytes are Parquet column-chunk bytes, including dictionary pages and page headers.\n\n")
	fmt.Fprintf(&b, "`Plain encoded bytes before compression` is the same column's byte count from the all-plain/no-compression baseline run. The `/ plain encoded` percentage columns compare compressed column bytes against that baseline denominator.\n\n")
	fmt.Fprintf(&b, "Categorization uses only measured byte sizes, row-group cardinality, and column type: `True dictionary bloat` means RLE dictionary encoded bytes exceeded plain encoded bytes before ZSTD; `Tiny/constant plain stream` means median row-group cardinality is at most 2 or median cardinality/rows is at most 0.0006; `Structured medium/high-cardinality numeric streams` means a numeric or temporal column has median cardinality/rows at least 0.09; the remaining losing columns fall into `Small-domain fixed-width literals`. Sortedness, page min/max, and value-length distributions are shown elsewhere in this report but are not currently used for this category assignment.\n\n")
	writeRLEDictWorseCategoryComparison(md, rleDictWorseCategories, reportDir)

	snappyRLEDictWorseCategories := buildSnappyRLEDictWorseCategoryComparison(byColumn, shapeByColumn)
	if err := writeSnappyRLEDictWorseCategoryImages(reportDir, &snappyRLEDictWorseCategories); err != nil {
		return err
	}
	md.Heading(2, "Snappy RLE Dict Worse Distribution By Category")
	fmt.Fprintf(&b, "For columns where the best observed `snappy + plain` compressed byte count is smaller than the best observed `snappy + rle-dict` compressed byte count, each category image plots `plain + snappy` compressed bytes on the x-axis and `rle-dict + snappy` compressed bytes on the y-axis using the same log byte scale. Points above the diagonal are larger with RLE dictionary encoding. Point color is bucketed by `plain/no-compression encoded bytes / rle-dict + snappy compressed bytes`, so high-ratio colors identify columns where RLE dictionary lost the head-to-head but still compressed the baseline dramatically.\n\n")
	fmt.Fprintf(&b, "The bucket tables below each image show how much worse RLE dictionary encoding was. Worse-by percentage is `(rle_dict_snappy_compressed_bytes / plain_snappy_compressed_bytes - 1) * 100`, so values can exceed 100%%.\n\n")
	fmt.Fprintf(&b, "The compressed bytes are Parquet column-chunk bytes, including dictionary pages and page headers. Dictionary-page byte breakdown columns are left blank when the cached Snappy result TSV does not contain those byte counts.\n\n")
	fmt.Fprintf(&b, "`Plain encoded bytes before compression` is the same column's byte count from the all-plain/no-compression baseline run. The `/ plain encoded` percentage columns compare compressed column bytes against that baseline denominator.\n\n")
	fmt.Fprintf(&b, "Categorization uses measured row-group cardinality and column type: `Medium-cardinality fixed-width numeric streams` means a non-timestamp numeric column has median cardinality/rows below 9%%; `High-cardinality fixed-width IDs / hashes` means a non-timestamp numeric column has median cardinality/rows at least 9%%; `High-cardinality timestamp streams` covers timestamp columns. Value-length distributions are included in the table for context, but these Snappy categories are driven by fixed-width type plus cardinality.\n\n")
	writeSnappyRLEDictWorseCategoryComparison(md, snappyRLEDictWorseCategories, reportDir)

	deltaBinaryPackedComparison := buildDeltaBinaryPackedWinnerComparison(byColumn)
	deltaBinaryPackedComparisonPath := filepath.Join(reportDir, "images", "delta_binary_packed_winner_vs_second_best_improvement.svg")
	if err := writeDeltaBinaryPackedWinnerComparisonSVG(deltaBinaryPackedComparisonPath, deltaBinaryPackedComparison); err != nil {
		return err
	}
	md.Heading(2, "Delta Binary Packed Winner vs Second Best Improvement Distribution")
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
	md.Heading(2, "Snappy Plain vs RLE Dict Improvement Distribution")
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
	writeOverallAbsoluteDifferenceComparison(md, buildOverallPlainRLEDictAbsoluteDifference(byColumn, shapeByColumn, "snappy"))
	snappyPageDistribution, err := loadPageEncodingDistribution(cfg, "snappy")
	if err != nil {
		return err
	}
	if snappyPageDistribution != nil {
		writePageEncodingDistributionMarkdown(md, *snappyPageDistribution, reportDir)
	}

	for _, column := range columns {
		observations := byColumn[column]
		columnType := observations[0].Column.Type
		md.Heading(2, fmt.Sprintf("%s (%s)", column, columnType))
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
	valueLengthMin, valueLengthMedian, valueLengthMax := columnValueLengthSummary(shape, true)
	fmt.Fprintf(b, "- Value length min/median/max: `%s / %s / %s` bytes\n", valueLengthMin, valueLengthMedian, valueLengthMax)
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

func buildEncodingRankDistribution(byColumn map[string][]columnObservation) encodingRankDistribution {
	rowsByKey := make(map[string]*encodingRankDistributionRow)
	summary := encodingRankDistribution{}
	scopes := []string{"zstd", "snappy"}

	for _, observations := range byColumn {
		for _, scope := range scopes {
			ranked := topColumnObservations(observations, scope, 0)
			for i, obs := range ranked {
				rank := i + 1
				c := obs.Experiment.Combo
				key := c.CompressionName + "\x00" + obs.Column.ConfigEncoding
				row := rowsByKey[key]
				if row == nil {
					row = &encodingRankDistributionRow{
						Compression: c.CompressionName,
						Encoding:    obs.Column.ConfigEncoding,
					}
					rowsByKey[key] = row
				}
				for len(row.RankCounts) < rank {
					row.RankCounts = append(row.RankCounts, 0)
				}
				row.RankCounts[rank-1]++
				row.Total++
				if rank > summary.MaxRank {
					summary.MaxRank = rank
				}
			}
		}
	}

	summary.Rows = make([]encodingRankDistributionRow, 0, len(rowsByKey))
	for _, row := range rowsByKey {
		for len(row.RankCounts) < summary.MaxRank {
			row.RankCounts = append(row.RankCounts, 0)
		}
		summary.Rows = append(summary.Rows, *row)
	}
	sort.Slice(summary.Rows, func(i, j int) bool {
		left := summary.Rows[i]
		right := summary.Rows[j]
		if compressionRankOrder(left.Compression) != compressionRankOrder(right.Compression) {
			return compressionRankOrder(left.Compression) < compressionRankOrder(right.Compression)
		}
		if left.Compression != right.Compression {
			return left.Compression < right.Compression
		}
		if encodingOrder(left.Encoding) != encodingOrder(right.Encoding) {
			return encodingOrder(left.Encoding) < encodingOrder(right.Encoding)
		}
		return left.Encoding < right.Encoding
	})
	return summary
}

func writeEncodingRankDistributionTable(b *strings.Builder, summary encodingRankDistribution) {
	if len(summary.Rows) == 0 || summary.MaxRank == 0 {
		fmt.Fprintf(b, "No encoding rank distribution data was available.\n\n")
		return
	}
	fmt.Fprintf(b, "| Compression | Encoding | Ranked columns |")
	for rank := 1; rank <= summary.MaxRank; rank++ {
		fmt.Fprintf(b, " Rank %d |", rank)
	}
	fmt.Fprintf(b, "\n")
	fmt.Fprintf(b, "| --- | --- | ---: |")
	for rank := 1; rank <= summary.MaxRank; rank++ {
		fmt.Fprintf(b, " ---: |")
	}
	fmt.Fprintf(b, "\n")
	for _, row := range summary.Rows {
		fmt.Fprintf(b, "| `%s` | `%s` | %d |", row.Compression, row.Encoding, row.Total)
		for rank := 0; rank < summary.MaxRank; rank++ {
			count := 0
			if rank < len(row.RankCounts) {
				count = row.RankCounts[rank]
			}
			fmt.Fprintf(b, " %d |", count)
		}
		fmt.Fprintf(b, "\n")
	}
	fmt.Fprintf(b, "\n")
}

func buildZstdPlainWinnerSecondPlaceDistribution(byColumn map[string][]columnObservation) secondPlaceDistribution {
	counts := make(map[string]int)
	summary := secondPlaceDistribution{
		FirstPlaceEncoding: "plain",
	}
	for _, observations := range byColumn {
		ranked := topColumnObservations(observations, "zstd", 0)
		if len(ranked) == 0 || ranked[0].Column.ConfigEncoding != summary.FirstPlaceEncoding {
			continue
		}
		summary.FirstPlaceCount++
		if len(ranked) < 2 {
			summary.MissingSecondPlace++
			continue
		}
		counts[ranked[1].Column.ConfigEncoding]++
	}

	summary.Rows = make([]secondPlaceDistributionRow, 0, len(counts))
	for encoding, count := range counts {
		summary.Rows = append(summary.Rows, secondPlaceDistributionRow{
			Encoding: encoding,
			Count:    count,
		})
	}
	sort.Slice(summary.Rows, func(i, j int) bool {
		if summary.Rows[i].Count != summary.Rows[j].Count {
			return summary.Rows[i].Count > summary.Rows[j].Count
		}
		if encodingOrder(summary.Rows[i].Encoding) != encodingOrder(summary.Rows[j].Encoding) {
			return encodingOrder(summary.Rows[i].Encoding) < encodingOrder(summary.Rows[j].Encoding)
		}
		return summary.Rows[i].Encoding < summary.Rows[j].Encoding
	})
	return summary
}

func writeSecondPlaceDistributionTable(b *strings.Builder, summary secondPlaceDistribution) {
	fmt.Fprintf(b, "- Columns where `zstd + %s` ranked first: `%d`\n", summary.FirstPlaceEncoding, summary.FirstPlaceCount)
	fmt.Fprintf(b, "- Missing second-place rows: `%d`\n\n", summary.MissingSecondPlace)
	if len(summary.Rows) == 0 {
		fmt.Fprintf(b, "No second-place data was available.\n\n")
		return
	}
	fmt.Fprintf(b, "| Second-place encoding | Columns |\n")
	fmt.Fprintf(b, "| --- | ---: |\n")
	for _, row := range summary.Rows {
		fmt.Fprintf(b, "| `zstd + %s` | %d |\n", row.Encoding, row.Count)
	}
	fmt.Fprintf(b, "\n")
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

func loadPageEncodingDistribution(cfg config, compression string) (*pageEncodingDistribution, error) {
	tsvPattern := filepath.Join(
		cfg.TSVDir,
		"page_encoding_distribution",
		fmt.Sprintf("*_rows-%d_plain-%s_vs_rle-dict-%s_page-distribution.tsv", cfg.Rows, compression, compression),
	)
	tsvPath := newestMatch(tsvPattern)
	if tsvPath == "" {
		return nil, nil
	}
	distribution, err := parsePageEncodingDistributionTSV(tsvPath, compression)
	if err != nil {
		return nil, err
	}
	distribution.Compression = compression
	distribution.TSVPath = tsvPath
	distribution.ImagePath = newestMatch(filepath.Join(
		cfg.MarkdownDir,
		"page_encoding_distribution",
		"images",
		fmt.Sprintf("*_rows-%d_plain-%s_vs_rle-dict-%s_page-distribution.svg", cfg.Rows, compression, compression),
	))
	distribution.PlainWinsImagePath = newestMatch(filepath.Join(
		cfg.MarkdownDir,
		"page_encoding_distribution",
		"images",
		fmt.Sprintf("*_rows-%d_plain-%s_vs_rle-dict-%s_page-distribution-plain-%s-absolute-wins.svg", cfg.Rows, compression, compression, compression),
	))
	return &distribution, nil
}

func parsePageEncodingDistributionTSV(path, compression string) (pageEncodingDistribution, error) {
	f, err := os.Open(path)
	if err != nil {
		return pageEncodingDistribution{}, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	header, err := reader.Read()
	if err != nil {
		return pageEncodingDistribution{}, err
	}
	index := make(map[string]int, len(header))
	for i, name := range header {
		index[name] = i
	}

	var distribution pageEncodingDistribution
	for {
		record, err := reader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return pageEncodingDistribution{}, err
		}
		col := pageEncodingDistributionColumn{
			Column:                 field(record, index, "column"),
			Type:                   field(record, index, "type"),
			PlainPages:             parseInt(field(record, index, "plain_pages")),
			RLEDictPages:           parseInt(field(record, index, "rle_dict_pages")),
			ComparisonWindows:      parseInt(field(record, index, "comparison_windows")),
			RowsCompared:           parseInt(field(record, index, "rows_compared")),
			PlainWindowWins:        parseInt(field(record, index, "plain_window_wins")),
			RLEDictWindowWins:      parseInt(field(record, index, "rle_dict_window_wins")),
			TieWindowWins:          parseInt(field(record, index, "tie_window_wins")),
			PlainRowsWon:           parseInt(field(record, index, "plain_rows_won")),
			RLEDictRowsWon:         parseInt(field(record, index, "rle_dict_rows_won")),
			TieRowsWon:             parseInt(field(record, index, "tie_rows_won")),
			PlainAllocatedBytes:    parseRoundedFloat(field(record, index, "plain_allocated_compressed_bytes")),
			RLEDictAllocatedBytes:  parseRoundedFloat(field(record, index, "rle_dict_allocated_compressed_bytes_with_amortized_dictionary")),
			RLEDictNoDictBytes:     parseRoundedFloat(field(record, index, "rle_dict_allocated_compressed_bytes_without_dictionary_pages")),
			UncompressedBytes:      parseRoundedFloat(field(record, index, "uncompressed_bytes_for_ratio")),
			PlainToUncompressed:    parseFloat(field(record, index, "plain_"+compression+"_aggregate_to_uncompressed_ratio")),
			RLEDictToUncompressed:  parseFloat(field(record, index, "rle_dict_"+compression+"_aggregate_to_uncompressed_ratio")),
			RLEDictRatioAdvantage:  parseFloat(field(record, index, "rle_dict_ratio_advantage_vs_plain")),
			AbsoluteRatioDiff:      parseFloat(field(record, index, "absolute_plain_rle_ratio_difference")),
			RLEDictToPlainRatio:    parseFloat(field(record, index, "rle_dict_allocated_to_plain_ratio")),
			RLEDictNoDictRatio:     parseFloat(field(record, index, "rle_dict_without_dictionary_allocated_to_plain_ratio")),
			DictOverheadFlipWins:   parseInt(field(record, index, "rle_dict_without_dictionary_overhead_flip_window_wins")),
			DictOverheadFlipRows:   parseInt(field(record, index, "rle_dict_without_dictionary_overhead_flip_rows")),
			WinnerByAllocatedBytes: field(record, index, "winner_by_allocated_bytes"),
			ExactMatchedPages:      parseInt(field(record, index, "exact_matched_pages")),
			UnmatchedPlainPages:    parseInt(field(record, index, "unmatched_plain_pages")),
			UnmatchedRLEDictPages:  parseInt(field(record, index, "unmatched_rle_dict_pages")),
		}
		if col.RLEDictToPlainRatio == 0 {
			if pct := field(record, index, "rle_dict_allocated_vs_plain_pct"); pct != "" {
				col.RLEDictToPlainRatio = 1 + parseFloat(pct)/100
			} else if col.PlainAllocatedBytes > 0 {
				col.RLEDictToPlainRatio = ratio(col.RLEDictAllocatedBytes, col.PlainAllocatedBytes)
			}
		}
		if col.RLEDictNoDictRatio == 0 && col.PlainAllocatedBytes > 0 && col.RLEDictNoDictBytes > 0 {
			col.RLEDictNoDictRatio = ratio(col.RLEDictNoDictBytes, col.PlainAllocatedBytes)
		}
		if col.UncompressedBytes > 0 {
			if col.PlainToUncompressed == 0 {
				col.PlainToUncompressed = ratio(col.PlainAllocatedBytes, col.UncompressedBytes)
			}
			if col.RLEDictToUncompressed == 0 {
				col.RLEDictToUncompressed = ratio(col.RLEDictAllocatedBytes, col.UncompressedBytes)
			}
		}
		if col.RLEDictRatioAdvantage == 0 && (col.PlainToUncompressed != 0 || col.RLEDictToUncompressed != 0) {
			col.RLEDictRatioAdvantage = col.PlainToUncompressed - col.RLEDictToUncompressed
		}
		if col.AbsoluteRatioDiff == 0 && col.RLEDictRatioAdvantage != 0 {
			col.AbsoluteRatioDiff = math.Abs(col.RLEDictRatioAdvantage)
		}
		distribution.Columns = append(distribution.Columns, col)
		distribution.ComparedColumns++
		distribution.ComparisonWindows += col.ComparisonWindows
		distribution.RowsCompared += col.RowsCompared
		distribution.PlainWindowWins += col.PlainWindowWins
		distribution.RLEDictWindowWins += col.RLEDictWindowWins
		distribution.TieWindowWins += col.TieWindowWins
		distribution.PlainRowsWon += col.PlainRowsWon
		distribution.RLEDictRowsWon += col.RLEDictRowsWon
		distribution.TieRowsWon += col.TieRowsWon
		distribution.PlainAllocatedBytes += col.PlainAllocatedBytes
		distribution.RLEDictAllocatedBytes += col.RLEDictAllocatedBytes
		distribution.RLEDictNoDictBytes += col.RLEDictNoDictBytes
		distribution.ExactMatchedPages += col.ExactMatchedPages
		distribution.UnmatchedPlainPages += col.UnmatchedPlainPages
		distribution.UnmatchedRLEDictPages += col.UnmatchedRLEDictPages
		distribution.DictOverheadFlipWins += col.DictOverheadFlipWins
		distribution.DictOverheadFlipRows += col.DictOverheadFlipRows

		hasPlain := col.PlainWindowWins > 0
		hasRLEDict := col.RLEDictWindowWins > 0
		hasTie := col.TieWindowWins > 0
		winnerKinds := 0
		for _, ok := range []bool{hasPlain, hasRLEDict, hasTie} {
			if ok {
				winnerKinds++
			}
		}
		switch {
		case winnerKinds > 1:
			distribution.MixedColumns++
		case hasPlain:
			distribution.PlainOnlyColumns++
		case hasRLEDict:
			distribution.RLEDictOnlyColumns++
		case hasTie:
			distribution.TieOnlyColumns++
		}
	}
	sort.Slice(distribution.Columns, func(i, j int) bool {
		if distribution.Columns[i].RLEDictRatioAdvantage != distribution.Columns[j].RLEDictRatioAdvantage {
			return distribution.Columns[i].RLEDictRatioAdvantage > distribution.Columns[j].RLEDictRatioAdvantage
		}
		if distribution.Columns[i].AbsoluteRatioDiff != distribution.Columns[j].AbsoluteRatioDiff {
			return distribution.Columns[i].AbsoluteRatioDiff > distribution.Columns[j].AbsoluteRatioDiff
		}
		return distribution.Columns[i].Column < distribution.Columns[j].Column
	})
	return distribution, nil
}

func writePageEncodingDistributionMarkdown(md *markdownDoc, distribution pageEncodingDistribution, reportDir string) {
	b := md.b
	compression := distribution.Compression
	if compression == "" {
		compression = "zstd"
	}
	plainLabel := "plain + " + compression
	rleDictLabel := "rle-dict + " + compression
	md.Heading(3, strings.ToUpper(compression)+" Page-Level Winner Distribution")
	fmt.Fprintf(b, "This is the page-level version of the same `%s` vs `%s` comparison. Page ranges differ between the two runs, so the distribution is computed over overlap windows from the union of page row ranges. Red chart segments are windows where RLE would win if dictionary-page bytes were excluded, but does not win when amortized dictionary-page bytes are included. Rows are sorted by `%s aggregate ratio - %s aggregate ratio`, where each aggregate ratio is final encoded-and-compressed page bytes divided by the same plain uncompressed encoded bytes. Larger positive gaps are bigger absolute wins for RLE dictionary.\n\n", plainLabel, rleDictLabel, plainLabel, rleDictLabel)
	if distribution.TSVPath != "" {
		fmt.Fprintf(b, "- Source TSV: [%s](%s)\n", filepath.Base(distribution.TSVPath), markdownLinkTarget(reportDir, distribution.TSVPath))
	}
	fmt.Fprintf(b, "- Compared columns: `%s`; mixed page winners: `%s`; plain-only columns: `%s`; rle-dict-only columns: `%s`; tie-only columns: `%s`\n",
		formatCount(int64(distribution.ComparedColumns)),
		formatCount(int64(distribution.MixedColumns)),
		formatCount(int64(distribution.PlainOnlyColumns)),
		formatCount(int64(distribution.RLEDictOnlyColumns)),
		formatCount(int64(distribution.TieOnlyColumns)),
	)
	fmt.Fprintf(b, "- Overlap windows: `%s`; `%s` wins: `%s`; `%s` wins: `%s`; ties: `%s`\n",
		formatCount(distribution.ComparisonWindows),
		plainLabel,
		formatCount(distribution.PlainWindowWins),
		rleDictLabel,
		formatCount(distribution.RLEDictWindowWins),
		formatCount(distribution.TieWindowWins),
	)
	fmt.Fprintf(b, "- Row-weighted wins: `%s` %s; `%s` %s; ties %s\n",
		plainLabel,
		formatCountPercent(distribution.PlainRowsWon, distribution.RowsCompared),
		rleDictLabel,
		formatCountPercent(distribution.RLEDictRowsWon, distribution.RowsCompared),
		formatCountPercent(distribution.TieRowsWon, distribution.RowsCompared),
	)
	fmt.Fprintf(b, "- Red overhead-flip windows: %s; row-weighted: %s\n",
		formatCountPercent(distribution.DictOverheadFlipWins, distribution.ComparisonWindows),
		formatCountPercent(distribution.DictOverheadFlipRows, distribution.RowsCompared),
	)
	fmt.Fprintf(b, "- Allocated page bytes: `%s` `%s`; `%s` `%s`; `%s / %s` `%s`\n",
		plainLabel,
		formatByteCount(distribution.PlainAllocatedBytes),
		rleDictLabel,
		formatByteCount(distribution.RLEDictAllocatedBytes),
		rleDictLabel,
		plainLabel,
		formatRatio(ratio(distribution.RLEDictAllocatedBytes, distribution.PlainAllocatedBytes)),
	)
	fmt.Fprintf(b, "- Exact matched page ranges: `%s`; unmatched plain pages: `%s`; unmatched rle-dict pages: `%s`\n\n",
		formatCount(distribution.ExactMatchedPages),
		formatCount(distribution.UnmatchedPlainPages),
		formatCount(distribution.UnmatchedRLEDictPages),
	)
	writeShapeImage(b, strings.ToUpper(compression)+" page-window winner distribution by column", distribution.ImagePath, reportDir)
	writePageEncodingDistributionTable(b, distribution)
	plainRows := plainAbsoluteWinDistributionColumns(distribution.Columns)
	if len(plainRows) > 0 {
		md.Heading(3, strings.ToUpper(compression)+" Plain Absolute Wins")
		fmt.Fprintf(b, "These are the columns where `%s` has a lower aggregate final-bytes-to-uncompressed ratio than `%s`, sorted by the largest absolute plain advantage.\n\n", plainLabel, rleDictLabel)
		if distribution.PlainWinsImagePath != "" {
			writeShapeImage(b, strings.ToUpper(compression)+" plain absolute wins by column", distribution.PlainWinsImagePath, reportDir)
		}
		writePageEncodingDistributionRowsTable(b, plainRows, compression)
	}
}

func writePageEncodingDistributionTable(b *strings.Builder, distribution pageEncodingDistribution) {
	if len(distribution.Columns) == 0 {
		return
	}
	writePageEncodingDistributionRowsTable(b, distribution.Columns, distribution.Compression)
}

func writePageEncodingDistributionRowsTable(b *strings.Builder, rows []pageEncodingDistributionColumn, compression string) {
	if compression == "" {
		compression = "zstd"
	}
	fmt.Fprintf(b, "| Column | Type | Windows | Plain wins | RLE dict wins | Red overhead flips | Ties | Row-weighted plain | Row-weighted RLE dict | Plain allocated bytes | RLE dict allocated bytes | Plain ratio | RLE ratio | RLE+%s advantage | Abs ratio gap | RLE+%s / plain+%s | Exact matches | Unmatched pages |\n", compression, compression, compression)
	fmt.Fprintf(b, "| --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |\n")
	for _, row := range rows {
		fmt.Fprintf(
			b,
			"| `%s` | `%s` | `%s` | %s | %s | %s | %s | %s | %s | `%s` | `%s` | `%s` | `%s` | `%s` | `%s` | `%s` | `%s` | `%s / %s` |\n",
			row.Column,
			row.Type,
			formatCount(row.ComparisonWindows),
			formatCountPercent(row.PlainWindowWins, row.ComparisonWindows),
			formatCountPercent(row.RLEDictWindowWins, row.ComparisonWindows),
			formatCountPercent(row.DictOverheadFlipWins, row.ComparisonWindows),
			formatCountPercent(row.TieWindowWins, row.ComparisonWindows),
			formatCountPercent(row.PlainRowsWon, row.RowsCompared),
			formatCountPercent(row.RLEDictRowsWon, row.RowsCompared),
			formatByteCount(row.PlainAllocatedBytes),
			formatByteCount(row.RLEDictAllocatedBytes),
			formatCompactPercentRatio(row.PlainToUncompressed),
			formatCompactPercentRatio(row.RLEDictToUncompressed),
			formatSignedCompactPercentRatio(row.RLEDictRatioAdvantage),
			formatCompactPercentRatio(row.AbsoluteRatioDiff),
			formatRatio(row.RLEDictToPlainRatio),
			formatCount(row.ExactMatchedPages),
			formatCount(row.UnmatchedPlainPages),
			formatCount(row.UnmatchedRLEDictPages),
		)
	}
	fmt.Fprintf(b, "\n")
}

func plainAbsoluteWinDistributionColumns(rows []pageEncodingDistributionColumn) []pageEncodingDistributionColumn {
	plainRows := make([]pageEncodingDistributionColumn, 0, len(rows))
	for _, row := range rows {
		if row.RLEDictRatioAdvantage < 0 {
			plainRows = append(plainRows, row)
		}
	}
	sort.Slice(plainRows, func(i, j int) bool {
		left := -plainRows[i].RLEDictRatioAdvantage
		right := -plainRows[j].RLEDictRatioAdvantage
		if left != right {
			return left > right
		}
		if plainRows[i].AbsoluteRatioDiff != plainRows[j].AbsoluteRatioDiff {
			return plainRows[i].AbsoluteRatioDiff > plainRows[j].AbsoluteRatioDiff
		}
		return plainRows[i].Column < plainRows[j].Column
	})
	return plainRows
}

func buildZstdRLEDictWorseCategoryComparison(byColumn map[string][]columnObservation, shapeByColumn map[string]columnShapeStats) rleDictWorseCategoryComparison {
	categories := rleDictWorseCategoryDefinitions()
	byName := make(map[string]*rleDictWorseCategory, len(categories))
	for i := range categories {
		byName[categories[i].Name] = &categories[i]
	}

	summary := rleDictWorseCategoryComparison{
		Categories: categories,
	}
	for column, observations := range byColumn {
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
		case rleDictBytes > plainBytes:
			summary.RLEDictWorseCount++
		case rleDictBytes < plainBytes:
			summary.RLEDictBetterCount++
			continue
		default:
			summary.TieCount++
			continue
		}

		shape, ok := shapeByColumn[column]
		if !ok {
			summary.MissingShapeStats++
		}
		rowGroupCardinality, medianCardinalityRatio, hasShapeStats := rleDictWorseShapeSummary(shape, ok)
		valueLengthMin, valueLengthMedian, valueLengthMax := columnValueLengthSummary(shape, ok)
		categoryName := classifyRLEDictWorseColumn(plain, rleDict, shape, ok)
		measuredFeature, measuredReason := rleDictWorseMeasuredReason(categoryName, plain, rleDict, shape, ok)
		category := byName[categoryName]
		if category == nil {
			category = byName["Small-domain fixed-width literals"]
		}
		category.Rows = append(category.Rows, rleDictWorseColumn{
			Column:                                  column,
			Type:                                    plain.Column.Type,
			Category:                                categoryName,
			MeasuredFeature:                         measuredFeature,
			MeasuredReason:                          measuredReason,
			RowGroupCardinality:                     rowGroupCardinality,
			MedianCardinalityRatio:                  medianCardinalityRatio,
			ValueLengthMin:                          valueLengthMin,
			ValueLengthMedian:                       valueLengthMedian,
			ValueLengthMax:                          valueLengthMax,
			HasShapeStats:                           hasShapeStats,
			PhysicalBytes:                           plain.Column.PhysicalBytes,
			BaselineEncodedBytes:                    plain.BaselineEncodedBytes,
			PlainBytes:                              plainBytes,
			RLEDictBytes:                            rleDictBytes,
			RLEDictBytesWithoutDictionaryPages:      rleDict.Column.CompressedBytesWithoutDictionaryPages,
			DictionaryPageCount:                     rleDict.Column.DictionaryPageCount,
			HasCompressedBytesWithoutDictionaryPage: rleDict.Column.HasCompressedBytesWithoutDictionaryPage,
			WorseByPct:                              percentLarger(rleDictBytes, plainBytes),
		})
	}

	for i := range summary.Categories {
		category := &summary.Categories[i]
		sort.Slice(category.Rows, func(i, j int) bool {
			if category.Rows[i].WorseByPct != category.Rows[j].WorseByPct {
				return category.Rows[i].WorseByPct > category.Rows[j].WorseByPct
			}
			return category.Rows[i].Column < category.Rows[j].Column
		})
		values := make([]float64, 0, len(category.Rows))
		for _, row := range category.Rows {
			values = append(values, row.WorseByPct)
		}
		category.Buckets = buildWorseByBuckets(values)
		category.MinPct, category.MedianPct, category.MaxPct = summarizeFloat64(values)
	}
	return summary
}

func rleDictWorseCategoryDefinitions() []rleDictWorseCategory {
	return []rleDictWorseCategory{
		{
			Name:        "True dictionary bloat",
			Slug:        "true_dictionary_bloat",
			Description: "RLE dictionary encoding was already larger than plain before ZSTD, usually because the dictionary itself was too large for the column.",
		},
		{
			Name:        "Tiny/constant plain stream",
			Slug:        "tiny_constant_plain_stream",
			Description: "The column is tiny or nearly constant per row group; plain pages give ZSTD an extremely repetitive stream, while dictionary pages add overhead.",
		},
		{
			Name:        "Small-domain fixed-width literals",
			Slug:        "small_domain_fixed_width_literals",
			Description: "RLE dictionary shrank the encoded stream, but ZSTD compressed the repeated fixed-width plain literals better than dictionary IDs plus a dictionary page.",
		},
		{
			Name:        "Structured medium/high-cardinality numeric streams",
			Slug:        "structured_medium_high_cardinality_numeric_streams",
			Description: "The column has enough distinct numeric/timestamp values that the plain stream preserves structure ZSTD can exploit better than dictionary IDs.",
		},
	}
}

func classifyRLEDictWorseColumn(plain, rleDict columnObservation, shape columnShapeStats, hasShape bool) string {
	if rleDict.Column.EncodedBytes > plain.Column.EncodedBytes {
		return "True dictionary bloat"
	}
	if hasShape {
		medianCardinality, medianRows := medianRowGroupCardinalityAndRows(shape.RowGroups)
		cardinalityRatio := 0.0
		if medianRows > 0 {
			cardinalityRatio = medianCardinality / medianRows
		}
		if medianCardinality <= 2 || cardinalityRatio <= 0.0006 {
			return "Tiny/constant plain stream"
		}
		if isNumericOrTemporalColumnType(plain.Column.Type) && cardinalityRatio >= 0.09 {
			return "Structured medium/high-cardinality numeric streams"
		}
	}
	return "Small-domain fixed-width literals"
}

func rleDictWorseMeasuredReason(category string, plain, rleDict columnObservation, shape columnShapeStats, hasShape bool) (string, string) {
	cardinalityFeature := "shape stats unavailable"
	if hasShape {
		medianCardinality, medianRows := medianRowGroupCardinalityAndRows(shape.RowGroups)
		ratio := 0.0
		if medianRows > 0 {
			ratio = medianCardinality / medianRows
		}
		cardinalityFeature = fmt.Sprintf("median row-group cardinality %.0f; median cardinality/rows %.6f%%", medianCardinality, ratio*100)
	}

	encodedComparison := fmt.Sprintf("plain encoded %s; rle encoded %s", formatByteCount(plain.Column.EncodedBytes), formatByteCount(rleDict.Column.EncodedBytes))
	compressedComparison := fmt.Sprintf("plain+zstd %s; rle+zstd %s", formatByteCount(plain.Column.CompressedBytes), formatByteCount(rleDict.Column.CompressedBytes))

	switch category {
	case "True dictionary bloat":
		return encodedComparison, "RLE dictionary was larger than plain before ZSTD; the compressed result stayed larger."
	case "Tiny/constant plain stream":
		return cardinalityFeature, "Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead."
	case "Small-domain fixed-width literals":
		return cardinalityFeature + "; " + encodedComparison, "RLE dictionary reduced pre-codec bytes, but ZSTD compressed the plain fixed-width values to fewer bytes."
	case "Structured medium/high-cardinality numeric streams":
		return cardinalityFeature + "; " + compressedComparison, "The high-cardinality numeric/timestamp column produced a larger RLE-dict+ZSTD stream than plain+ZSTD."
	default:
		return cardinalityFeature, "RLE-dict+ZSTD was larger than plain+ZSTD for this measured column."
	}
}

func rleDictWorseShapeSummary(shape columnShapeStats, ok bool) (string, float64, bool) {
	if !ok || len(shape.RowGroups) == 0 {
		return "", 0, false
	}
	cardinalities := make([]int64, 0, len(shape.RowGroups))
	rows := make([]float64, 0, len(shape.RowGroups))
	for _, rowGroup := range shape.RowGroups {
		cardinalities = append(cardinalities, rowGroup.Cardinality)
		rows = append(rows, float64(rowGroup.NumRows))
	}
	medianCardinality := medianFloat64(int64sToFloat64s(cardinalities))
	medianRows := medianFloat64(rows)
	ratio := 0.0
	if medianRows > 0 {
		ratio = medianCardinality / medianRows
	}
	return summarizeInt64(cardinalities), ratio, true
}

func columnValueLengthSummary(shape columnShapeStats, ok bool) (string, string, string) {
	if !ok || len(shape.RowGroups) == 0 {
		return "", "", ""
	}
	haveLength := false
	minLength := 0
	maxLength := 0
	for _, rowGroup := range shape.RowGroups {
		if !haveLength || rowGroup.MinValueLength < minLength {
			minLength = rowGroup.MinValueLength
		}
		if !haveLength || rowGroup.MaxValueLength > maxLength {
			maxLength = rowGroup.MaxValueLength
		}
		haveLength = true
	}
	if !haveLength {
		return "", "", ""
	}
	median := shape.MedianValueLength
	if median == 0 && !(minLength == 0 && maxLength == 0) {
		medians := make([]float64, 0, len(shape.RowGroups))
		for _, rowGroup := range shape.RowGroups {
			if rowGroup.MedianValueLength != 0 || (rowGroup.MinValueLength == 0 && rowGroup.MaxValueLength == 0) {
				medians = append(medians, rowGroup.MedianValueLength)
			}
		}
		if len(medians) > 0 {
			median = medianFloat64(medians)
		}
	}
	medianText := ""
	if median != 0 || minLength == 0 {
		medianText = formatLengthFloat(median)
	}
	return formatCount(int64(minLength)), medianText, formatCount(int64(maxLength))
}

func int64sToFloat64s(values []int64) []float64 {
	out := make([]float64, 0, len(values))
	for _, value := range values {
		out = append(out, float64(value))
	}
	return out
}

func medianRowGroupCardinalityAndRows(rowGroups []shapeRowGroupStats) (float64, float64) {
	cardinalities := make([]float64, 0, len(rowGroups))
	rows := make([]float64, 0, len(rowGroups))
	for _, rowGroup := range rowGroups {
		cardinalities = append(cardinalities, float64(rowGroup.Cardinality))
		rows = append(rows, float64(rowGroup.NumRows))
	}
	return medianFloat64(cardinalities), medianFloat64(rows)
}

func isNumericOrTemporalColumnType(columnType string) bool {
	switch {
	case strings.HasPrefix(columnType, "int"):
		return true
	case strings.HasPrefix(columnType, "uint"):
		return true
	case strings.HasPrefix(columnType, "timestamp"):
		return true
	case columnType == "date":
		return true
	default:
		return false
	}
}

func isTimestampColumnType(columnType string) bool {
	return strings.HasPrefix(columnType, "timestamp")
}

func percentLarger(larger, smaller int64) float64 {
	if smaller <= 0 {
		return 0
	}
	return ((float64(larger) / float64(smaller)) - 1) * 100
}

func buildWorseByBuckets(values []float64) []encodingImprovementBucket {
	labels := []string{
		"0-10%",
		"10-20%",
		"20-30%",
		"30-40%",
		"40-50%",
		"50-60%",
		"60-70%",
		"70-80%",
		"80-90%",
		"90-100%",
		"100-200%",
		"200-500%",
		"500%+",
	}
	buckets := make([]encodingImprovementBucket, len(labels))
	for i, label := range labels {
		buckets[i].Label = label
	}
	for _, value := range values {
		index := worseByBucketIndex(value)
		buckets[index].Count++
	}
	return buckets
}

func worseByBucketIndex(value float64) int {
	if value < 0 || math.IsNaN(value) || math.IsInf(value, 0) {
		return 0
	}
	if value < 100 {
		return int(value / 10)
	}
	if value < 200 {
		return 10
	}
	if value < 500 {
		return 11
	}
	return 12
}

func summarizeFloat64(values []float64) (float64, float64, float64) {
	if len(values) == 0 {
		return 0, 0, 0
	}
	values = append([]float64(nil), values...)
	sort.Float64s(values)
	return values[0], medianFloat64(values), values[len(values)-1]
}

func medianFloat64(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	values = append([]float64(nil), values...)
	sort.Float64s(values)
	mid := len(values) / 2
	if len(values)%2 == 1 {
		return values[mid]
	}
	return (values[mid-1] + values[mid]) / 2
}

func writeRLEDictWorseCategoryImages(reportDir string, summary *rleDictWorseCategoryComparison) error {
	for i := range summary.Categories {
		category := &summary.Categories[i]
		path := filepath.Join(reportDir, "images", "zstd_rle_dict_worse_"+category.Slug+".svg")
		if err := writeRLEDictWorseCategoryScatterSVG(
			path,
			"RLE dictionary worse: "+category.Name,
			category.Rows,
		); err != nil {
			return err
		}
	}
	return nil
}

func writeRLEDictWorseCategoryScatterSVG(path, title string, rows []rleDictWorseColumn) error {
	if len(rows) == 0 {
		return writeEmptySVG(path, title, "no data")
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	const (
		width        = 980
		height       = 650
		leftMargin   = 104
		topMargin    = 72
		plotSize     = 440
		legendLeft   = 590
		bottomMargin = 96
	)

	minBytes := math.Inf(1)
	maxBytes := math.Inf(-1)
	for _, row := range rows {
		for _, value := range []int64{row.PlainBytes, row.RLEDictBytes} {
			if value <= 0 {
				continue
			}
			f := float64(value)
			if f < minBytes {
				minBytes = f
			}
			if f > maxBytes {
				maxBytes = f
			}
		}
	}
	if math.IsInf(minBytes, 0) || math.IsInf(maxBytes, 0) {
		return writeEmptySVG(path, title, "no positive byte counts")
	}

	axisMin := math.Pow(10, math.Floor(math.Log10(minBytes)))
	axisMax := math.Pow(10, math.Ceil(math.Log10(maxBytes)))
	if axisMin <= 0 {
		axisMin = 1
	}
	if axisMax <= axisMin {
		axisMax = axisMin * 10
	}
	logMin := math.Log10(axisMin)
	logMax := math.Log10(axisMax)
	scaleX := func(value int64) float64 {
		if value <= 0 {
			value = 1
		}
		return leftMargin + ((math.Log10(float64(value)) - logMin) / (logMax - logMin) * plotSize)
	}
	scaleY := func(value int64) float64 {
		if value <= 0 {
			value = 1
		}
		return topMargin + ((logMax - math.Log10(float64(value))) / (logMax - logMin) * plotSize)
	}

	var b strings.Builder
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d">`+"\n", width, height, width, height)
	fmt.Fprintf(&b, `<rect width="100%%" height="100%%" fill="#ffffff"/>`+"\n")
	fmt.Fprintf(&b, `<text x="%d" y="28" font-family="Arial, sans-serif" font-size="16" font-weight="700" fill="#111827">%s</text>`+"\n", leftMargin, html.EscapeString(title))
	fmt.Fprintf(&b, `<text x="%d" y="48" font-family="Arial, sans-serif" font-size="11" fill="#6b7280">%s</text>`+"\n", leftMargin, html.EscapeString("same log byte scale on both axes; points above the diagonal are larger with rle-dict + zstd"))

	plotRight := leftMargin + plotSize
	plotBottom := topMargin + plotSize
	for _, tick := range logByteTicks(axisMin, axisMax) {
		x := scaleX(int64(tick))
		y := scaleY(int64(tick))
		label := formatBytes(int64(math.Round(tick)))
		fmt.Fprintf(&b, `<line x1="%.2f" y1="%d" x2="%.2f" y2="%d" stroke="#eef2f7" stroke-width="1"/>`+"\n", x, topMargin, x, plotBottom)
		fmt.Fprintf(&b, `<line x1="%d" y1="%.2f" x2="%d" y2="%.2f" stroke="#eef2f7" stroke-width="1"/>`+"\n", leftMargin, y, plotRight, y)
		fmt.Fprintf(&b, `<text x="%.2f" y="%d" font-family="Arial, sans-serif" font-size="10" text-anchor="middle" fill="#6b7280">%s</text>`+"\n", x, plotBottom+22, html.EscapeString(label))
		fmt.Fprintf(&b, `<text x="%d" y="%.2f" font-family="Arial, sans-serif" font-size="10" text-anchor="end" fill="#6b7280">%s</text>`+"\n", leftMargin-8, y+4, html.EscapeString(label))
	}
	fmt.Fprintf(&b, `<rect x="%d" y="%d" width="%d" height="%d" fill="none" stroke="#9ca3af" stroke-width="1"/>`+"\n", leftMargin, topMargin, plotSize, plotSize)
	fmt.Fprintf(&b, `<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="#111827" stroke-width="1.4" stroke-dasharray="5 5"/>`+"\n", leftMargin, plotBottom, plotRight, topMargin)
	fmt.Fprintf(&b, `<text x="%d" y="%d" font-family="Arial, sans-serif" font-size="10" fill="#111827">equal size</text>`+"\n", leftMargin+12, plotBottom-12)

	for _, row := range rows {
		compressionRatio := ratio(row.BaselineEncodedBytes, row.RLEDictBytes)
		bucket := compressionRatioColor(compressionRatio)
		x := scaleX(row.PlainBytes)
		y := scaleY(row.RLEDictBytes)
		tooltip := fmt.Sprintf(
			"%s\nplain+zstd: %s\nrle-dict+zstd: %s\nworse by: %s\nplain baseline / rle-dict+zstd: %s",
			row.Column,
			formatByteCount(row.PlainBytes),
			formatByteCount(row.RLEDictBytes),
			formatPercent(row.WorseByPct),
			formatCompactRatio(compressionRatio),
		)
		fmt.Fprintf(&b, `<circle cx="%.2f" cy="%.2f" r="5.5" fill="%s" stroke="#ffffff" stroke-width="1.2" opacity="0.88"><title>%s</title></circle>`+"\n",
			x,
			y,
			bucket.Color,
			html.EscapeString(tooltip),
		)
	}

	fmt.Fprintf(&b, `<text x="%d" y="%d" font-family="Arial, sans-serif" font-size="11" text-anchor="middle" fill="#374151">%s</text>`+"\n", leftMargin+plotSize/2, height-34, html.EscapeString("plain + zstd compressed bytes"))
	fmt.Fprintf(&b, `<text x="20" y="%d" transform="rotate(-90 20 %d)" font-family="Arial, sans-serif" font-size="11" text-anchor="middle" fill="#374151">%s</text>`+"\n", topMargin+plotSize/2, topMargin+plotSize/2, html.EscapeString("rle-dict + zstd compressed bytes"))

	fmt.Fprintf(&b, `<text x="%d" y="%d" font-family="Arial, sans-serif" font-size="12" font-weight="700" fill="#111827">Color: plain baseline / rle-dict+zstd</text>`+"\n", legendLeft, topMargin+6)
	legendY := topMargin + 28
	for _, bucket := range compressionRatioColorBuckets() {
		fmt.Fprintf(&b, `<rect x="%d" y="%d" width="13" height="13" rx="2" fill="%s"/>`+"\n", legendLeft, legendY-10, bucket.Color)
		fmt.Fprintf(&b, `<text x="%d" y="%d" font-family="Arial, sans-serif" font-size="11" fill="#374151">%s</text>`+"\n", legendLeft+20, legendY+1, html.EscapeString(bucket.Label))
		legendY += 22
	}
	fmt.Fprintf(&b, `<text x="%d" y="%d" font-family="Arial, sans-serif" font-size="11" fill="#6b7280">%s</text>`+"\n", legendLeft, legendY+12, html.EscapeString(fmt.Sprintf("points: %d losing columns", len(rows))))
	fmt.Fprintf(&b, "</svg>\n")
	return os.WriteFile(path, []byte(b.String()), 0o644)
}

func logByteTicks(minValue, maxValue float64) []float64 {
	if minValue <= 0 || maxValue <= minValue {
		return nil
	}
	ticks := make([]float64, 0, 6)
	logMin := math.Log10(minValue)
	logMax := math.Log10(maxValue)
	for i := 0; i <= 4; i++ {
		ticks = append(ticks, math.Pow(10, logMin+(float64(i)*(logMax-logMin)/4)))
	}
	return ticks
}

func compressionRatioColor(value float64) compressionRatioColorBucket {
	buckets := compressionRatioColorBuckets()
	for _, bucket := range buckets {
		if value >= bucket.Min && value < bucket.Max {
			return bucket
		}
	}
	return buckets[len(buckets)-1]
}

func compressionRatioColorBuckets() []compressionRatioColorBucket {
	return []compressionRatioColorBucket{
		{Label: "<1x", Min: math.Inf(-1), Max: 1, Color: "#991b1b"},
		{Label: "1-2x", Min: 1, Max: 2, Color: "#dc2626"},
		{Label: "2-5x", Min: 2, Max: 5, Color: "#f97316"},
		{Label: "5-10x", Min: 5, Max: 10, Color: "#f59e0b"},
		{Label: "10-25x", Min: 10, Max: 25, Color: "#84cc16"},
		{Label: "25-50x", Min: 25, Max: 50, Color: "#22c55e"},
		{Label: "50-100x", Min: 50, Max: 100, Color: "#14b8a6"},
		{Label: "100x+", Min: 100, Max: math.Inf(1), Color: "#2563eb"},
	}
}

func formatCompactRatio(value float64) string {
	if value <= 0 || math.IsNaN(value) || math.IsInf(value, 0) {
		return "n/a"
	}
	if value >= 100 {
		return fmt.Sprintf("%.0fx", value)
	}
	if value >= 10 {
		return fmt.Sprintf("%.1fx", value)
	}
	return fmt.Sprintf("%.2fx", value)
}

func formatCompactPercentRatio(value float64) string {
	if value < 0 || math.IsNaN(value) || math.IsInf(value, 0) {
		return "n/a"
	}
	pct := value * 100
	if pct >= 100 {
		return fmt.Sprintf("%.0f%%", pct)
	}
	if pct >= 10 {
		return fmt.Sprintf("%.1f%%", pct)
	}
	return fmt.Sprintf("%.2f%%", pct)
}

func formatSignedCompactPercentRatio(value float64) string {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return "n/a"
	}
	if value == 0 {
		return "0.00%"
	}
	if value < 0 {
		return "-" + formatCompactPercentRatio(-value)
	}
	return "+" + formatCompactPercentRatio(value)
}

func writeRLEDictWorseCategoryComparison(md *markdownDoc, summary rleDictWorseCategoryComparison, reportDir string) {
	b := md.b
	fmt.Fprintf(b, "- Compared columns: `%d`\n", summary.ComparedColumns)
	fmt.Fprintf(b, "- `zstd + rle-dict` worse than `zstd + plain`: `%d`; better: `%d`; ties: `%d`; missing comparisons: `%d`\n", summary.RLEDictWorseCount, summary.RLEDictBetterCount, summary.TieCount, summary.MissingCount)
	fmt.Fprintf(b, "- Missing shape stats while categorizing: `%d`\n\n", summary.MissingShapeStats)

	fmt.Fprintf(b, "| Category | Columns | Worse by min/median/max |\n")
	fmt.Fprintf(b, "| --- | ---: | ---: |\n")
	for _, category := range summary.Categories {
		fmt.Fprintf(b, "| %s | %d | %s / %s / %s |\n",
			category.Name,
			len(category.Rows),
			formatPercent(category.MinPct),
			formatPercent(category.MedianPct),
			formatPercent(category.MaxPct),
		)
	}
	fmt.Fprintf(b, "\n")

	for _, category := range summary.Categories {
		md.Heading(3, category.Name)
		fmt.Fprintf(b, "%s\n\n", category.Description)
		imagePath := filepath.Join(reportDir, "images", "zstd_rle_dict_worse_"+category.Slug+".svg")
		writeShapeImage(b, "RLE dictionary worse: "+category.Name, imagePath, reportDir)
		writeImprovementBucketTable(b, "`zstd + rle-dict` worse by", category.Buckets)
		if len(category.Rows) == 0 {
			continue
		}
		fmt.Fprintf(b, "| Column | Type | Category | Measured feature | Measured reason | Row-group cardinality min/median/max | Median cardinality / rows | Min value length (B) | Median value length (B) | Max value length (B) | Physical bytes before encoding/compression | Plain encoded bytes before compression | Plain + ZSTD compressed bytes | Plain + ZSTD / physical | Plain + ZSTD / plain encoded | RLE dict + ZSTD compressed bytes | RLE dict + ZSTD / physical | RLE dict + ZSTD / plain encoded | RLE dict + ZSTD vs plain + ZSTD | RLE dict + ZSTD without dict pages | RLE dict without dict pages / physical | RLE dict without dict pages / plain encoded | RLE dict without dict pages vs plain + ZSTD | RLE + dict is better without including dict page | Dictionary pages |\n")
		fmt.Fprintf(b, "| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | --- | ---: |\n")
		for _, row := range category.Rows {
			fmt.Fprintf(b, "| `%s` | `%s` | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %d |\n",
				row.Column,
				row.Type,
				row.Category,
				row.MeasuredFeature,
				row.MeasuredReason,
				optionalString(row.RowGroupCardinality, row.HasShapeStats),
				optionalPercentMarkdown(row.MedianCardinalityRatio*100, row.HasShapeStats),
				optionalString(row.ValueLengthMin, row.HasShapeStats),
				optionalString(row.ValueLengthMedian, row.HasShapeStats),
				optionalString(row.ValueLengthMax, row.HasShapeStats),
				formatByteCount(row.PhysicalBytes),
				formatByteCount(row.BaselineEncodedBytes),
				formatByteCount(row.PlainBytes),
				formatBytesAsPercentOf(row.PlainBytes, row.PhysicalBytes),
				formatBytesAsPercentOf(row.PlainBytes, row.BaselineEncodedBytes),
				formatByteCount(row.RLEDictBytes),
				formatBytesAsPercentOf(row.RLEDictBytes, row.PhysicalBytes),
				formatBytesAsPercentOf(row.RLEDictBytes, row.BaselineEncodedBytes),
				formatPercent(row.WorseByPct),
				formatOptionalByteCount(row.RLEDictBytesWithoutDictionaryPages, row.HasCompressedBytesWithoutDictionaryPage),
				formatOptionalBytesAsPercentOf(row.RLEDictBytesWithoutDictionaryPages, row.PhysicalBytes, row.HasCompressedBytesWithoutDictionaryPage),
				formatOptionalBytesAsPercentOf(row.RLEDictBytesWithoutDictionaryPages, row.BaselineEncodedBytes, row.HasCompressedBytesWithoutDictionaryPage),
				formatOptionalPercentDifference(row.RLEDictBytesWithoutDictionaryPages, row.PlainBytes, row.HasCompressedBytesWithoutDictionaryPage),
				yesNo(row.HasCompressedBytesWithoutDictionaryPage && row.RLEDictBytesWithoutDictionaryPages < row.PlainBytes),
				row.DictionaryPageCount,
			)
		}
		fmt.Fprintf(b, "\n")
	}
}

func buildSnappyRLEDictWorseCategoryComparison(byColumn map[string][]columnObservation, shapeByColumn map[string]columnShapeStats) snappyRLEDictWorseCategoryComparison {
	categories := snappyRLEDictWorseCategoryDefinitions()
	byName := make(map[string]*snappyRLEDictWorseCategory, len(categories))
	for i := range categories {
		byName[categories[i].Name] = &categories[i]
	}

	summary := snappyRLEDictWorseCategoryComparison{
		Categories: categories,
	}
	for column, observations := range byColumn {
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
		case rleDictBytes > plainBytes:
			summary.RLEDictWorseCount++
		case rleDictBytes < plainBytes:
			summary.RLEDictBetterCount++
			continue
		default:
			summary.TieCount++
			continue
		}

		shape, ok := shapeByColumn[column]
		if !ok {
			summary.MissingShapeStats++
		}
		rowGroupCardinality, medianCardinalityRatio, hasShapeStats := rleDictWorseShapeSummary(shape, ok)
		valueLengthMin, valueLengthMedian, valueLengthMax := columnValueLengthSummary(shape, ok)
		categoryName := classifySnappyRLEDictWorseColumn(plain, shape, ok)
		measuredFeature, measuredReason := snappyRLEDictWorseMeasuredReason(categoryName, plain, rleDict, shape, ok)
		category := byName[categoryName]
		if category == nil {
			category = byName["Medium-cardinality fixed-width numeric streams"]
		}
		category.Rows = append(category.Rows, snappyRLEDictWorseColumn{
			Column:                                  column,
			Type:                                    plain.Column.Type,
			Category:                                categoryName,
			MeasuredFeature:                         measuredFeature,
			MeasuredReason:                          measuredReason,
			RowGroupCardinality:                     rowGroupCardinality,
			MedianCardinalityRatio:                  medianCardinalityRatio,
			ValueLengthMin:                          valueLengthMin,
			ValueLengthMedian:                       valueLengthMedian,
			ValueLengthMax:                          valueLengthMax,
			HasShapeStats:                           hasShapeStats,
			PhysicalBytes:                           plain.Column.PhysicalBytes,
			BaselineEncodedBytes:                    plain.BaselineEncodedBytes,
			RLEDictEncodedBytes:                     rleDict.Column.EncodedBytes,
			PlainBytes:                              plainBytes,
			RLEDictBytes:                            rleDictBytes,
			RLEDictBytesWithoutDictionaryPages:      rleDict.Column.CompressedBytesWithoutDictionaryPages,
			DictionaryPageCount:                     rleDict.Column.DictionaryPageCount,
			HasCompressedBytesWithoutDictionaryPage: rleDict.Column.HasCompressedBytesWithoutDictionaryPage,
			WorseByPct:                              percentLarger(rleDictBytes, plainBytes),
		})
	}

	for i := range summary.Categories {
		category := &summary.Categories[i]
		sort.Slice(category.Rows, func(i, j int) bool {
			if category.Rows[i].WorseByPct != category.Rows[j].WorseByPct {
				return category.Rows[i].WorseByPct > category.Rows[j].WorseByPct
			}
			return category.Rows[i].Column < category.Rows[j].Column
		})
		values := make([]float64, 0, len(category.Rows))
		for _, row := range category.Rows {
			values = append(values, row.WorseByPct)
		}
		category.Buckets = buildWorseByBuckets(values)
		category.MinPct, category.MedianPct, category.MaxPct = summarizeFloat64(values)
	}
	return summary
}

func snappyRLEDictWorseCategoryDefinitions() []snappyRLEDictWorseCategory {
	return []snappyRLEDictWorseCategory{
		{
			Name:        "Medium-cardinality fixed-width numeric streams",
			Slug:        "medium_cardinality_fixed_width_numeric_streams",
			Description: "Non-timestamp numeric columns with medium row-group cardinality; RLE dictionary reduced the pre-compression stream, but Snappy compressed the plain fixed-width stream to fewer bytes.",
		},
		{
			Name:        "High-cardinality fixed-width IDs / hashes",
			Slug:        "high_cardinality_fixed_width_ids_hashes",
			Description: "Non-timestamp numeric ID/hash-like columns with high row-group cardinality; dictionary IDs had too little repetition to beat Snappy over plain fixed-width values.",
		},
		{
			Name:        "High-cardinality timestamp streams",
			Slug:        "high_cardinality_timestamp_streams",
			Description: "Timestamp columns with high row-group cardinality; RLE dictionary barely reduced the encoded stream, and Snappy did better on the plain timestamp bytes.",
		},
	}
}

func classifySnappyRLEDictWorseColumn(plain columnObservation, shape columnShapeStats, hasShape bool) string {
	if isTimestampColumnType(plain.Column.Type) {
		return "High-cardinality timestamp streams"
	}
	if hasShape {
		medianCardinality, medianRows := medianRowGroupCardinalityAndRows(shape.RowGroups)
		cardinalityRatio := 0.0
		if medianRows > 0 {
			cardinalityRatio = medianCardinality / medianRows
		}
		if isNumericOrTemporalColumnType(plain.Column.Type) && cardinalityRatio >= 0.09 {
			return "High-cardinality fixed-width IDs / hashes"
		}
	}
	return "Medium-cardinality fixed-width numeric streams"
}

func snappyRLEDictWorseMeasuredReason(category string, plain, rleDict columnObservation, shape columnShapeStats, hasShape bool) (string, string) {
	cardinalityFeature := "shape stats unavailable"
	if hasShape {
		medianCardinality, medianRows := medianRowGroupCardinalityAndRows(shape.RowGroups)
		ratio := 0.0
		if medianRows > 0 {
			ratio = medianCardinality / medianRows
		}
		cardinalityFeature = fmt.Sprintf("median row-group cardinality %.0f; median cardinality/rows %.6f%%", medianCardinality, ratio*100)
	}
	encodedPct := ratio(rleDict.Column.EncodedBytes, plain.BaselineEncodedBytes) * 100
	encodedFeature := fmt.Sprintf("rle encoded %s of plain encoded", formatPercent(encodedPct))
	worseFeature := fmt.Sprintf("rle+snappy %s; plain+snappy %s; rle larger by %s", formatByteCount(rleDict.Column.CompressedBytes), formatByteCount(plain.Column.CompressedBytes), formatPercent(percentLarger(rleDict.Column.CompressedBytes, plain.Column.CompressedBytes)))

	switch category {
	case "High-cardinality timestamp streams":
		return cardinalityFeature + "; " + encodedFeature, "High-cardinality timestamp values left little dictionary repetition; Snappy compressed the plain timestamp stream to fewer bytes."
	case "High-cardinality fixed-width IDs / hashes":
		if rleDict.Column.EncodedBytes >= plain.BaselineEncodedBytes {
			return cardinalityFeature + "; " + encodedFeature, "RLE dictionary was already larger than plain before Snappy; the compressed result stayed larger."
		}
		return cardinalityFeature + "; " + encodedFeature, "High cardinality limited dictionary benefit; Snappy over plain fixed-width values stayed smaller."
	case "Medium-cardinality fixed-width numeric streams":
		return cardinalityFeature + "; " + encodedFeature, "RLE dictionary reduced pre-compression bytes, but Snappy compressed the plain fixed-width stream better for this column."
	default:
		return cardinalityFeature + "; " + worseFeature, "Snappy+RLE-dict was larger than Snappy+plain for this measured column."
	}
}

func writeSnappyRLEDictWorseCategoryImages(reportDir string, summary *snappyRLEDictWorseCategoryComparison) error {
	for i := range summary.Categories {
		category := &summary.Categories[i]
		if len(category.Rows) == 0 {
			continue
		}
		path := filepath.Join(reportDir, "images", "snappy_rle_dict_worse_"+category.Slug+".svg")
		if err := writeSnappyRLEDictWorseCategoryScatterSVG(
			path,
			"Snappy RLE dictionary worse: "+category.Name,
			category.Rows,
		); err != nil {
			return err
		}
	}
	return nil
}

func writeSnappyRLEDictWorseCategoryScatterSVG(path, title string, rows []snappyRLEDictWorseColumn) error {
	if len(rows) == 0 {
		return writeEmptySVG(path, title, "no data")
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	const (
		width        = 980
		height       = 650
		leftMargin   = 104
		topMargin    = 72
		plotSize     = 440
		legendLeft   = 590
		bottomMargin = 96
	)

	minBytes := math.Inf(1)
	maxBytes := math.Inf(-1)
	for _, row := range rows {
		for _, value := range []int64{row.PlainBytes, row.RLEDictBytes} {
			if value <= 0 {
				continue
			}
			f := float64(value)
			if f < minBytes {
				minBytes = f
			}
			if f > maxBytes {
				maxBytes = f
			}
		}
	}
	if math.IsInf(minBytes, 0) || math.IsInf(maxBytes, 0) {
		return writeEmptySVG(path, title, "no positive byte counts")
	}

	axisMin := math.Pow(10, math.Floor(math.Log10(minBytes)))
	axisMax := math.Pow(10, math.Ceil(math.Log10(maxBytes)))
	if axisMin <= 0 {
		axisMin = 1
	}
	if axisMax <= axisMin {
		axisMax = axisMin * 10
	}
	logMin := math.Log10(axisMin)
	logMax := math.Log10(axisMax)
	scaleX := func(value int64) float64 {
		if value <= 0 {
			value = 1
		}
		return leftMargin + ((math.Log10(float64(value)) - logMin) / (logMax - logMin) * plotSize)
	}
	scaleY := func(value int64) float64 {
		if value <= 0 {
			value = 1
		}
		return topMargin + ((logMax - math.Log10(float64(value))) / (logMax - logMin) * plotSize)
	}

	var b strings.Builder
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d">`+"\n", width, height, width, height)
	fmt.Fprintf(&b, `<rect width="100%%" height="100%%" fill="#ffffff"/>`+"\n")
	fmt.Fprintf(&b, `<text x="%d" y="28" font-family="Arial, sans-serif" font-size="16" font-weight="700" fill="#111827">%s</text>`+"\n", leftMargin, html.EscapeString(title))
	fmt.Fprintf(&b, `<text x="%d" y="48" font-family="Arial, sans-serif" font-size="11" fill="#6b7280">%s</text>`+"\n", leftMargin, html.EscapeString("same log byte scale on both axes; points above the diagonal are larger with rle-dict + snappy"))

	plotRight := leftMargin + plotSize
	plotBottom := topMargin + plotSize
	for _, tick := range logByteTicks(axisMin, axisMax) {
		x := scaleX(int64(tick))
		y := scaleY(int64(tick))
		label := formatBytes(int64(math.Round(tick)))
		fmt.Fprintf(&b, `<line x1="%.2f" y1="%d" x2="%.2f" y2="%d" stroke="#eef2f7" stroke-width="1"/>`+"\n", x, topMargin, x, plotBottom)
		fmt.Fprintf(&b, `<line x1="%d" y1="%.2f" x2="%d" y2="%.2f" stroke="#eef2f7" stroke-width="1"/>`+"\n", leftMargin, y, plotRight, y)
		fmt.Fprintf(&b, `<text x="%.2f" y="%d" font-family="Arial, sans-serif" font-size="10" text-anchor="middle" fill="#6b7280">%s</text>`+"\n", x, plotBottom+22, html.EscapeString(label))
		fmt.Fprintf(&b, `<text x="%d" y="%.2f" font-family="Arial, sans-serif" font-size="10" text-anchor="end" fill="#6b7280">%s</text>`+"\n", leftMargin-8, y+4, html.EscapeString(label))
	}
	fmt.Fprintf(&b, `<rect x="%d" y="%d" width="%d" height="%d" fill="none" stroke="#9ca3af" stroke-width="1"/>`+"\n", leftMargin, topMargin, plotSize, plotSize)
	fmt.Fprintf(&b, `<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="#111827" stroke-width="1.4" stroke-dasharray="5 5"/>`+"\n", leftMargin, plotBottom, plotRight, topMargin)
	fmt.Fprintf(&b, `<text x="%d" y="%d" font-family="Arial, sans-serif" font-size="10" fill="#111827">equal size</text>`+"\n", leftMargin+12, plotBottom-12)

	for _, row := range rows {
		compressionRatio := ratio(row.BaselineEncodedBytes, row.RLEDictBytes)
		bucket := compressionRatioColor(compressionRatio)
		x := scaleX(row.PlainBytes)
		y := scaleY(row.RLEDictBytes)
		tooltip := fmt.Sprintf(
			"%s\nplain+snappy: %s\nrle-dict+snappy: %s\nworse by: %s\nplain baseline / rle-dict+snappy: %s",
			row.Column,
			formatByteCount(row.PlainBytes),
			formatByteCount(row.RLEDictBytes),
			formatPercent(row.WorseByPct),
			formatCompactRatio(compressionRatio),
		)
		fmt.Fprintf(&b, `<circle cx="%.2f" cy="%.2f" r="5.5" fill="%s" stroke="#ffffff" stroke-width="1.2" opacity="0.88"><title>%s</title></circle>`+"\n",
			x,
			y,
			bucket.Color,
			html.EscapeString(tooltip),
		)
	}

	fmt.Fprintf(&b, `<text x="%d" y="%d" font-family="Arial, sans-serif" font-size="11" text-anchor="middle" fill="#374151">%s</text>`+"\n", leftMargin+plotSize/2, height-34, html.EscapeString("plain + snappy compressed bytes"))
	fmt.Fprintf(&b, `<text x="20" y="%d" transform="rotate(-90 20 %d)" font-family="Arial, sans-serif" font-size="11" text-anchor="middle" fill="#374151">%s</text>`+"\n", topMargin+plotSize/2, topMargin+plotSize/2, html.EscapeString("rle-dict + snappy compressed bytes"))

	fmt.Fprintf(&b, `<text x="%d" y="%d" font-family="Arial, sans-serif" font-size="12" font-weight="700" fill="#111827">Color: plain baseline / rle-dict+snappy</text>`+"\n", legendLeft, topMargin+6)
	legendY := topMargin + 28
	for _, bucket := range compressionRatioColorBuckets() {
		fmt.Fprintf(&b, `<rect x="%d" y="%d" width="13" height="13" rx="2" fill="%s"/>`+"\n", legendLeft, legendY-10, bucket.Color)
		fmt.Fprintf(&b, `<text x="%d" y="%d" font-family="Arial, sans-serif" font-size="11" fill="#374151">%s</text>`+"\n", legendLeft+20, legendY+1, html.EscapeString(bucket.Label))
		legendY += 22
	}
	fmt.Fprintf(&b, `<text x="%d" y="%d" font-family="Arial, sans-serif" font-size="11" fill="#6b7280">%s</text>`+"\n", legendLeft, legendY+12, html.EscapeString(fmt.Sprintf("points: %d losing columns", len(rows))))
	fmt.Fprintf(&b, "</svg>\n")
	return os.WriteFile(path, []byte(b.String()), 0o644)
}

func writeSnappyRLEDictWorseCategoryComparison(md *markdownDoc, summary snappyRLEDictWorseCategoryComparison, reportDir string) {
	b := md.b
	fmt.Fprintf(b, "- Compared columns: `%d`\n", summary.ComparedColumns)
	fmt.Fprintf(b, "- `snappy + rle-dict` worse than `snappy + plain`: `%d`; better: `%d`; ties: `%d`; missing comparisons: `%d`\n", summary.RLEDictWorseCount, summary.RLEDictBetterCount, summary.TieCount, summary.MissingCount)
	fmt.Fprintf(b, "- Missing shape stats while categorizing: `%d`\n\n", summary.MissingShapeStats)

	fmt.Fprintf(b, "| Category | Columns | Worse by min/median/max |\n")
	fmt.Fprintf(b, "| --- | ---: | ---: |\n")
	for _, category := range summary.Categories {
		if len(category.Rows) == 0 {
			continue
		}
		fmt.Fprintf(b, "| %s | %d | %s / %s / %s |\n",
			category.Name,
			len(category.Rows),
			formatPercent(category.MinPct),
			formatPercent(category.MedianPct),
			formatPercent(category.MaxPct),
		)
	}
	fmt.Fprintf(b, "\n")

	for _, category := range summary.Categories {
		if len(category.Rows) == 0 {
			continue
		}
		md.Heading(3, category.Name)
		fmt.Fprintf(b, "%s\n\n", category.Description)
		imagePath := filepath.Join(reportDir, "images", "snappy_rle_dict_worse_"+category.Slug+".svg")
		writeShapeImage(b, "Snappy RLE dictionary worse: "+category.Name, imagePath, reportDir)
		writeImprovementBucketTable(b, "`snappy + rle-dict` worse by", category.Buckets)
		fmt.Fprintf(b, "| Column | Type | Category | Measured feature | Measured reason | Row-group cardinality min/median/max | Median cardinality / rows | Min value length (B) | Median value length (B) | Max value length (B) | Physical bytes before encoding/compression | Plain encoded bytes before compression | RLE dict encoded bytes before compression | RLE dict encoded / plain encoded | Plain + Snappy compressed bytes | Plain + Snappy / physical | Plain + Snappy / plain encoded | RLE dict + Snappy compressed bytes | RLE dict + Snappy / physical | RLE dict + Snappy / plain encoded | RLE dict + Snappy vs plain + Snappy | RLE dict + Snappy without dict pages | RLE dict without dict pages / physical | RLE dict without dict pages / plain encoded | RLE dict without dict pages vs plain + Snappy | RLE + dict is better without including dict page | Dictionary pages |\n")
		fmt.Fprintf(b, "| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | --- | ---: |\n")
		for _, row := range category.Rows {
			withoutDictBetter := ""
			if row.HasCompressedBytesWithoutDictionaryPage {
				withoutDictBetter = yesNo(row.RLEDictBytesWithoutDictionaryPages < row.PlainBytes)
			}
			fmt.Fprintf(b, "| `%s` | `%s` | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %s | %d |\n",
				row.Column,
				row.Type,
				row.Category,
				row.MeasuredFeature,
				row.MeasuredReason,
				optionalString(row.RowGroupCardinality, row.HasShapeStats),
				optionalPercentMarkdown(row.MedianCardinalityRatio*100, row.HasShapeStats),
				optionalString(row.ValueLengthMin, row.HasShapeStats),
				optionalString(row.ValueLengthMedian, row.HasShapeStats),
				optionalString(row.ValueLengthMax, row.HasShapeStats),
				formatByteCount(row.PhysicalBytes),
				formatByteCount(row.BaselineEncodedBytes),
				formatByteCount(row.RLEDictEncodedBytes),
				formatBytesAsPercentOf(row.RLEDictEncodedBytes, row.BaselineEncodedBytes),
				formatByteCount(row.PlainBytes),
				formatBytesAsPercentOf(row.PlainBytes, row.PhysicalBytes),
				formatBytesAsPercentOf(row.PlainBytes, row.BaselineEncodedBytes),
				formatByteCount(row.RLEDictBytes),
				formatBytesAsPercentOf(row.RLEDictBytes, row.PhysicalBytes),
				formatBytesAsPercentOf(row.RLEDictBytes, row.BaselineEncodedBytes),
				formatPercent(row.WorseByPct),
				formatOptionalByteCount(row.RLEDictBytesWithoutDictionaryPages, row.HasCompressedBytesWithoutDictionaryPage),
				formatOptionalBytesAsPercentOf(row.RLEDictBytesWithoutDictionaryPages, row.PhysicalBytes, row.HasCompressedBytesWithoutDictionaryPage),
				formatOptionalBytesAsPercentOf(row.RLEDictBytesWithoutDictionaryPages, row.BaselineEncodedBytes, row.HasCompressedBytesWithoutDictionaryPage),
				formatOptionalPercentDifference(row.RLEDictBytesWithoutDictionaryPages, row.PlainBytes, row.HasCompressedBytesWithoutDictionaryPage),
				withoutDictBetter,
				row.DictionaryPageCount,
			)
		}
		fmt.Fprintf(b, "\n")
	}
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

func buildOverallPlainRLEDictAbsoluteDifference(byColumn map[string][]columnObservation, shapeByColumn map[string]columnShapeStats, compression string) overallAbsoluteDifferenceComparison {
	summary := overallAbsoluteDifferenceComparison{
		Compression: compression,
	}
	for _, observations := range byColumn {
		plain, plainOK := exactEncodingComboObservationFor(observations, compression, isPlainEncodingCombo)
		rleDict, rleDictOK := exactEncodingComboObservationFor(observations, compression, isRLEDictEncodingCombo)
		if !plainOK || !rleDictOK || plain.BaselineEncodedBytes <= 0 {
			summary.MissingCount++
			continue
		}
		summary.ComparedColumns++
		cardinalityRatioMin, cardinalityRatioMedian, cardinalityRatioMax, hasShapeStats := rowGroupCardinalityRatioSummary(shapeByColumn[plain.Column.Column])
		plainRatio := ratio(plain.Column.CompressedBytes, plain.BaselineEncodedBytes)
		rleDictRatio := ratio(rleDict.Column.CompressedBytes, plain.BaselineEncodedBytes)
		row := overallAbsoluteDifferenceRow{
			Column:                 plain.Column.Column,
			Type:                   plain.Column.Type,
			HasShapeStats:          hasShapeStats,
			CardinalityRatioMin:    cardinalityRatioMin,
			CardinalityRatioMedian: cardinalityRatioMedian,
			CardinalityRatioMax:    cardinalityRatioMax,
			UncompressedBytes:      plain.BaselineEncodedBytes,
			PlainBytes:             plain.Column.CompressedBytes,
			RLEDictBytes:           rleDict.Column.CompressedBytes,
			PlainRatio:             plainRatio,
			RLEDictRatio:           rleDictRatio,
			AbsoluteDifference:     math.Abs(plainRatio - rleDictRatio),
		}
		switch {
		case rleDictRatio < plainRatio:
			summary.RLEDictBetterRows = append(summary.RLEDictBetterRows, row)
			summary.RLEDictBetterCount++
		case plainRatio < rleDictRatio:
			summary.PlainBetterRows = append(summary.PlainBetterRows, row)
			summary.PlainBetterCount++
		default:
			summary.TieCount++
		}
	}
	sortOverallAbsoluteDifferenceRows(summary.RLEDictBetterRows)
	sortOverallAbsoluteDifferenceRows(summary.PlainBetterRows)
	return summary
}

func exactEncodingComboObservationFor(observations []columnObservation, compression string, match func(combo) bool) (columnObservation, bool) {
	for _, obs := range observations {
		if obs.Experiment.Combo.Compression != compression {
			continue
		}
		if match(obs.Experiment.Combo) {
			return obs, true
		}
	}
	return columnObservation{}, false
}

func rowGroupCardinalityRatioSummary(shape columnShapeStats) (float64, float64, float64, bool) {
	if len(shape.RowGroups) == 0 {
		return 0, 0, 0, false
	}
	ratios := make([]float64, 0, len(shape.RowGroups))
	for _, rowGroup := range shape.RowGroups {
		if rowGroup.NumRows <= 0 {
			continue
		}
		ratios = append(ratios, float64(rowGroup.Cardinality)/float64(rowGroup.NumRows))
	}
	if len(ratios) == 0 {
		return 0, 0, 0, false
	}
	minValue, medianValue, maxValue := summarizeFloat64(ratios)
	return minValue, medianValue, maxValue, true
}

func sortOverallAbsoluteDifferenceRows(rows []overallAbsoluteDifferenceRow) {
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].AbsoluteDifference != rows[j].AbsoluteDifference {
			return rows[i].AbsoluteDifference > rows[j].AbsoluteDifference
		}
		return rows[i].Column < rows[j].Column
	})
}

func writeOverallAbsoluteDifferenceComparison(md *markdownDoc, summary overallAbsoluteDifferenceComparison) {
	b := md.b
	display := strings.ToUpper(summary.Compression)
	plainLabel := summary.Compression + " + plain"
	rleDictLabel := summary.Compression + " + rle-dict"
	md.Heading(3, display+" Overall Absolute Difference")
	fmt.Fprintf(b, "These are overall column-level comparisons, not page-window comparisons. `%s` is the all-plain run for every type group, and `%s` is the all-rle-dict run for every type group. The ratio denominator is the same for both sides: the all-plain/no-compression Parquet encoded byte count for that column. Absolute difference is `abs((%s compressed bytes / plain uncompressed encoded bytes) - (%s compressed bytes / plain uncompressed encoded bytes))`.\n\n", plainLabel, rleDictLabel, plainLabel, rleDictLabel)
	fmt.Fprintf(b, "- Compared columns: `%d`; `%s` better: `%d`; `%s` better: `%d`; ties: `%d`; missing comparisons: `%d`\n\n",
		summary.ComparedColumns,
		rleDictLabel,
		summary.RLEDictBetterCount,
		plainLabel,
		summary.PlainBetterCount,
		summary.TieCount,
		summary.MissingCount,
	)
	md.Heading(4, display+" RLE Dict Better By Absolute Difference")
	writeOverallAbsoluteDifferenceRows(b, summary.RLEDictBetterRows, plainLabel, rleDictLabel)
	md.Heading(4, display+" Plain Better By Absolute Difference")
	writeOverallAbsoluteDifferenceRows(b, summary.PlainBetterRows, plainLabel, rleDictLabel)
}

func writeOverallAbsoluteDifferenceRows(b *strings.Builder, rows []overallAbsoluteDifferenceRow, plainLabel, rleDictLabel string) {
	if len(rows) == 0 {
		fmt.Fprintf(b, "No columns.\n\n")
		return
	}
	fmt.Fprintf(b, "| Column | Type | Row-group cardinality/rows min | Row-group cardinality/rows median | Row-group cardinality/rows max | Plain uncompressed encoded bytes | %s compressed bytes | %s compressed bytes | %s ratio | %s ratio | Absolute difference |\n", plainLabel, rleDictLabel, plainLabel, rleDictLabel)
	fmt.Fprintf(b, "| --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |\n")
	for _, row := range rows {
		fmt.Fprintf(b, "| `%s` | `%s` | %s | %s | %s | %s | %s | %s | `%s` | `%s` | `%s` |\n",
			row.Column,
			row.Type,
			optionalPercentMarkdown(row.CardinalityRatioMin*100, row.HasShapeStats),
			optionalPercentMarkdown(row.CardinalityRatioMedian*100, row.HasShapeStats),
			optionalPercentMarkdown(row.CardinalityRatioMax*100, row.HasShapeStats),
			formatByteCount(row.UncompressedBytes),
			formatByteCount(row.PlainBytes),
			formatByteCount(row.RLEDictBytes),
			formatRatio(row.PlainRatio),
			formatRatio(row.RLEDictRatio),
			formatRatio(row.AbsoluteDifference),
		)
	}
	fmt.Fprintf(b, "\n")
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

func encodingColor(encoding string) string {
	switch encoding {
	case "plain":
		return "#2563eb"
	case "rle-dict":
		return "#0f766e"
	case "delta-binary-packed":
		return "#7c3aed"
	case "delta-byte-array":
		return "#d97706"
	case "delta-length-byte-array":
		return "#be123c"
	default:
		return "#6b7280"
	}
}

func compressionRankOrder(compression string) int {
	switch {
	case strings.HasPrefix(compression, "zstd"):
		return 0
	case compression == "snappy":
		return 1
	default:
		return 2
	}
}

func encodingOrder(encoding string) int {
	switch encoding {
	case "plain":
		return 0
	case "rle-dict":
		return 1
	case "delta-binary-packed":
		return 2
	case "delta-byte-array":
		return 3
	case "delta-length-byte-array":
		return 4
	default:
		return 5
	}
}

func rankColor(rank int) string {
	colors := []string{
		"#16a34a",
		"#2563eb",
		"#f59e0b",
		"#dc2626",
		"#7c3aed",
		"#0891b2",
	}
	if rank <= 0 {
		return "#6b7280"
	}
	return colors[(rank-1)%len(colors)]
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
		fmt.Fprintf(b, "%d. %s compressed - %s; %s encoded; %sx post-compression ratio; experiment `%s`\n",
			i+1,
			formatByteCount(obs.Column.CompressedBytes),
			prefix,
			formatByteCount(obs.Column.EncodedBytes),
			formatRatio(obs.PostCompressionRatio),
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
	if limit > 0 && len(rows) > limit {
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

func writeEncodingRankDistributionSVG(path string, summary encodingRankDistribution) error {
	if len(summary.Rows) == 0 || summary.MaxRank == 0 {
		return writeEmptySVG(path, "Encoding rank distribution by compression", "no data")
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	const (
		width        = 1040
		leftMargin   = 260
		rightMargin  = 92
		topMargin    = 78
		bottomMargin = 54
		rowHeight    = 34
	)
	height := topMargin + bottomMargin + rowHeight*len(summary.Rows)
	maxTotal := 0
	for _, row := range summary.Rows {
		if row.Total > maxTotal {
			maxTotal = row.Total
		}
	}
	if maxTotal == 0 {
		return writeEmptySVG(path, "Encoding rank distribution by compression", "no ranked columns")
	}

	plotWidth := float64(width - leftMargin - rightMargin)
	scale := func(value int) float64 {
		return float64(value) * plotWidth / float64(maxTotal)
	}

	var b strings.Builder
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d">`+"\n", width, height, width, height)
	fmt.Fprintf(&b, `<rect width="100%%" height="100%%" fill="#ffffff"/>`+"\n")
	fmt.Fprintf(&b, `<text x="%d" y="28" font-family="Arial, sans-serif" font-size="16" font-weight="700" fill="#111827">Encoding rank distribution by compression</text>`+"\n", leftMargin)
	fmt.Fprintf(&b, `<text x="%d" y="%d" font-family="Arial, sans-serif" font-size="11" fill="#6b7280">columns ranked</text>`+"\n", leftMargin, height-14)

	legendX := leftMargin
	for rank := 1; rank <= summary.MaxRank; rank++ {
		fmt.Fprintf(&b, `<rect x="%d" y="44" width="10" height="10" rx="2" fill="%s"/>`+"\n", legendX, rankColor(rank))
		fmt.Fprintf(&b, `<text x="%d" y="54" font-family="Arial, sans-serif" font-size="11" fill="#374151">rank %d</text>`+"\n", legendX+16, rank)
		legendX += 78
	}

	plotTop := topMargin - 6
	plotBottom := height - bottomMargin + 4
	for i := 0; i <= 4; i++ {
		value := float64(i) * float64(maxTotal) / 4
		x := float64(leftMargin) + value*plotWidth/float64(maxTotal)
		fmt.Fprintf(&b, `<line x1="%.2f" y1="%d" x2="%.2f" y2="%d" stroke="#eef2f7" stroke-width="1"/>`+"\n", x, plotTop, x, plotBottom)
		fmt.Fprintf(&b, `<text x="%.2f" y="%d" font-family="Arial, sans-serif" font-size="10" text-anchor="middle" fill="#6b7280">%s</text>`+"\n", x, height-32, formatPlotNumber(value))
	}
	fmt.Fprintf(&b, `<line x1="%d" y1="%d" x2="%d" y2="%d" stroke="#9ca3af" stroke-width="1"/>`+"\n", leftMargin, plotBottom, width-rightMargin, plotBottom)

	for i, row := range summary.Rows {
		y := topMargin + i*rowHeight
		fmt.Fprintf(&b, `<text x="%d" y="%d" font-family="Arial, sans-serif" font-size="12" text-anchor="end" fill="#374151">%s</text>`+"\n", leftMargin-12, y+19, html.EscapeString(fmt.Sprintf("%s + %s", row.Compression, row.Encoding)))
		x := float64(leftMargin)
		for rank, count := range row.RankCounts {
			if count == 0 {
				continue
			}
			segmentWidth := scale(count)
			fmt.Fprintf(&b, `<rect x="%.2f" y="%d" width="%.2f" height="18" rx="2" fill="%s"/>`+"\n", x, y+4, segmentWidth, rankColor(rank+1))
			if segmentWidth >= 22 {
				fmt.Fprintf(&b, `<text x="%.2f" y="%d" font-family="Arial, sans-serif" font-size="10" font-weight="700" text-anchor="middle" fill="#ffffff">%d</text>`+"\n", x+segmentWidth/2, y+17, count)
			}
			x += segmentWidth
		}
		fmt.Fprintf(&b, `<text x="%.2f" y="%d" font-family="Arial, sans-serif" font-size="11" fill="#111827">%d</text>`+"\n", x+8, y+17, row.Total)
	}
	fmt.Fprintf(&b, "</svg>\n")
	return os.WriteFile(path, []byte(b.String()), 0o644)
}

func writeZstdPlainWinnerSecondPlaceDistributionSVG(path string, summary secondPlaceDistribution) error {
	bars := make([]barChartBar, 0, len(summary.Rows))
	for _, row := range summary.Rows {
		bars = append(bars, barChartBar{
			Label: fmt.Sprintf("zstd + %s", row.Encoding),
			Value: row.Count,
			Color: encodingColor(row.Encoding),
		})
	}
	return writeHorizontalBarChartSVG(
		path,
		"Second place when ZSTD plain ranks first",
		"columns where zstd + plain is rank 1",
		bars,
	)
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

func formatOptionalByteCount(n int64, ok bool) string {
	if !ok {
		return ""
	}
	return formatByteCount(n)
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

func formatLengthFloat(value float64) string {
	if math.Abs(value-math.Round(value)) < 0.000001 {
		return formatCount(int64(math.Round(value)))
	}
	return fmt.Sprintf("%.2f", value)
}

func formatPercent(value float64) string {
	return fmt.Sprintf("%.6f%%", value)
}

func formatCountPercent(count, total int64) string {
	if total <= 0 {
		return fmt.Sprintf("`%s` (`n/a`)", formatCount(count))
	}
	pct := (float64(count) / float64(total)) * 100
	return fmt.Sprintf("`%s` (`%.2f%%`)", formatCount(count), pct)
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

func yesNo(value bool) string {
	if value {
		return "yes"
	}
	return "no"
}

func formatBytesAsPercentOf(bytes, total int64) string {
	if total == 0 {
		return "n/a"
	}
	return formatPercent((float64(bytes) / float64(total)) * 100)
}

func formatOptionalBytesAsPercentOf(bytes, total int64, ok bool) string {
	if !ok {
		return ""
	}
	return formatBytesAsPercentOf(bytes, total)
}

func formatOptionalPercentDifference(candidate, baseline int64, ok bool) string {
	if !ok {
		return ""
	}
	return formatPercent(percentLarger(candidate, baseline))
}

func optionalInt(value int64, ok bool) string {
	if !ok {
		return ""
	}
	return strconv.FormatInt(value, 10)
}

func optionalString(value string, ok bool) string {
	if !ok {
		return ""
	}
	return value
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
