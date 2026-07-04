package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	defaultInput       = "data/clickbench/hits.tsv.gz"
	fixedMaxPageSize   = "256KiB"
	fixedMaxRowGroup   = "10MiB"
	fixedMaxFileSize   = "10MiB"
	defaultDictMaxSize = "1MiB"
)

type config struct {
	Rows          int64
	Parallel      int
	ZstdLevel     int
	Input         string
	MaxDictSize   string
	ExperimentDir string
	ResultDir     string
	MarkdownDir   string
	ConfigDir     string
	TSVDir        string
	OutputRoot    string
	Verify        bool
	SkipExisting  bool
	KeepOutput    bool
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
	fs.BoolVar(&cfg.Verify, "verify", cfg.Verify, "verify every generated parquet output")
	fs.BoolVar(&cfg.SkipExisting, "skip-existing", cfg.SkipExisting, "reuse an existing result markdown/column TSV when present")
	fs.BoolVar(&cfg.KeepOutput, "keep-output", cfg.KeepOutput, "keep generated parquet output directories after the experiment; only valid with --parallel 1")
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
	intEncodings := []string{"plain", "rle-dict"}
	dateEncodings := []string{"plain", "rle-dict"}
	timestampEncodings := []string{"plain", "rle-dict"}
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
		columns, err := parseColumnStatsTSV(existingColumns)
		result.Elapsed = time.Since(started)
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
	if err := writeSummaryMarkdown(summaryPath, cfg, rankings, settingSummaries, winners, bestColumns, baseline, started, finished, allExperimentsPath, settingsPath, columnResultsPath, columnWinnersPath, bestColumnEncodingsPath); err != nil {
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

func writeSummaryMarkdown(path string, cfg config, rankings []experimentRanking, settings []settingSummary, winners []columnWinner, bestColumnWinners []columnWinner, baseline experimentResult, started, finished time.Time, allExperimentsPath, settingsPath, columnResultsPath, columnWinnersPath, bestColumnEncodingsPath string) error {
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
	fmt.Fprintf(&b, "- Best encoding per column: [%s](%s)\n\n", filepath.Base(bestColumnEncodingsPath), markdownLinkTarget(summaryDir, bestColumnEncodingsPath))

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
