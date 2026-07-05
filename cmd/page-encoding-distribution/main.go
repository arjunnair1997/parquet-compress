package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"html"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	defaultInput      = "data/clickbench/hits.tsv.gz"
	defaultRows       = int64(1_000_000)
	defaultPageSize   = "256KiB"
	defaultDictSize   = "256KiB"
	defaultRowGroup   = "10MiB"
	defaultFileSize   = "10MiB"
	defaultZstdLevel  = 3
	plainLabel        = "plain+zstd"
	rleDictLabel      = "rle-dict+zstd"
	distributionDir   = "page_encoding_distribution"
	pageWinnerPlain   = "plain"
	pageWinnerRLEDict = "rle-dict"
	pageWinnerTie     = "tie"
)

type config struct {
	Rows             int64
	Input            string
	ZstdLevel        int
	MaxPageSize      string
	MaxDictSize      string
	MaxRowGroupRows  int64
	MaxRowGroupSize  string
	MaxFileSize      string
	ExperimentDir    string
	ResultDir        string
	ReportDir        string
	ConfigDir        string
	TSVDir           string
	OutputRoot       string
	PlainStatsJSON   string
	RLEDictStatsJSON string
	Verify           bool
	KeepOutput       bool
	GeneratePDF      bool
}

type writerRun struct {
	Name      string
	Encoding  string
	StatsPath string
	OutputDir string
	LogPath   string
	Elapsed   time.Duration
}

type statsSnapshot struct {
	Columns []columnStats `json:"columns"`
	Errors  []string      `json:"errors,omitempty"`
}

type columnStats struct {
	ColumnIndex       int             `json:"column_index"`
	Path              []string        `json:"path"`
	Name              string          `json:"name"`
	Type              string          `json:"type"`
	PhysicalType      string          `json:"physical_type"`
	SortedAscending   bool            `json:"sorted_ascending"`
	SortedDescending  bool            `json:"sorted_descending"`
	MinValueLength    int             `json:"min_value_length"`
	MedianValueLength float64         `json:"median_value_length"`
	MaxValueLength    int             `json:"max_value_length"`
	RowGroups         []rowGroupStats `json:"row_groups"`
	Pages             []pageStats     `json:"pages"`
}

type rowGroupStats struct {
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

type pageStats struct {
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

type pageCost struct {
	Column        string
	Type          string
	RowGroupIndex int64
	PageIndex     int
	Start         int64
	End           int64
	Rows          int64
	Cost          float64
	Encoding      string
}

type columnDistribution struct {
	Column                     string
	Type                       string
	PlainPages                 int
	RLEDictPages               int
	ComparisonWindows          int
	RowsCompared               int64
	PlainWindowWins            int
	RLEDictWindowWins          int
	TieWindowWins              int
	PlainRowsWon               int64
	RLEDictRowsWon             int64
	TieRowsWon                 int64
	PlainAllocatedBytes        float64
	RLEDictAllocatedBytes      float64
	PlainTotalPageBytes        float64
	RLEDictTotalPageBytes      float64
	ExactMatchedPages          int
	ExactPlainWins             int
	ExactRLEDictWins           int
	ExactTies                  int
	UnmatchedPlainPages        int
	UnmatchedRLEDictPages      int
	WinnerByAllocatedBytes     string
	PlainWindowWinPct          float64
	RLEDictWindowWinPct        float64
	TieWindowWinPct            float64
	PlainRowWinPct             float64
	RLEDictRowWinPct           float64
	TieRowWinPct               float64
	RLEDictAllocatedVsPlainPct float64
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
		Rows:            defaultRows,
		Input:           defaultInput,
		ZstdLevel:       defaultZstdLevel,
		MaxPageSize:     defaultPageSize,
		MaxDictSize:     defaultDictSize,
		MaxRowGroupSize: defaultRowGroup,
		MaxFileSize:     defaultFileSize,
		Verify:          true,
	}
	fs := flag.NewFlagSet("page-encoding-distribution", flag.ContinueOnError)
	fs.Int64Var(&cfg.Rows, "rows", cfg.Rows, "rows to write for each comparison run")
	fs.StringVar(&cfg.Input, "input", cfg.Input, "path to hits.tsv or hits.tsv.gz")
	fs.IntVar(&cfg.ZstdLevel, "zstd-level", cfg.ZstdLevel, "zstd compression level")
	fs.StringVar(&cfg.MaxPageSize, "max-page-size", cfg.MaxPageSize, "target parquet page buffer size")
	fs.StringVar(&cfg.MaxDictSize, "max-dictionary-page-size", cfg.MaxDictSize, "maximum dictionary bytes before fallback")
	fs.Int64Var(&cfg.MaxRowGroupRows, "max-row-group-rows", cfg.MaxRowGroupRows, "approximate maximum rows per row group; 0 disables")
	fs.StringVar(&cfg.MaxRowGroupSize, "max-row-group-size", cfg.MaxRowGroupSize, "approximate row group byte-size threshold; 0 disables")
	fs.StringVar(&cfg.MaxFileSize, "max-file-size", cfg.MaxFileSize, "approximate parquet file byte-size threshold; 0 writes one file")
	fs.StringVar(&cfg.ExperimentDir, "experiment-dir", "", "fixed-settings experiment directory; defaults from page/row-group/file/dictionary settings")
	fs.StringVar(&cfg.OutputRoot, "output-root", "", "root directory for generated parquet files; defaults under the row result directory")
	fs.StringVar(&cfg.PlainStatsJSON, "plain-stats-json", "", "existing writer stats JSON for plain+zstd; when both stats paths are set, writer runs are skipped")
	fs.StringVar(&cfg.RLEDictStatsJSON, "rle-dict-stats-json", "", "existing writer stats JSON for rle-dict+zstd; when both stats paths are set, writer runs are skipped")
	fs.BoolVar(&cfg.Verify, "verify", cfg.Verify, "verify generated parquet output")
	fs.BoolVar(&cfg.KeepOutput, "keep-output", cfg.KeepOutput, "keep generated parquet output directories")
	fs.BoolVar(&cfg.GeneratePDF, "generate-pdf", cfg.GeneratePDF, "write PDFs for the two underlying writer run markdown files")
	if err := fs.Parse(args); err != nil {
		return cfg, err
	}
	if cfg.Rows <= 0 {
		return cfg, fmt.Errorf("--rows must be > 0")
	}
	if cfg.ZstdLevel <= 0 {
		return cfg, fmt.Errorf("--zstd-level must be > 0")
	}
	for name, value := range map[string]string{
		"--max-page-size":            cfg.MaxPageSize,
		"--max-dictionary-page-size": cfg.MaxDictSize,
		"--max-row-group-size":       cfg.MaxRowGroupSize,
		"--max-file-size":            cfg.MaxFileSize,
	} {
		if _, err := parseSize(value); err != nil {
			return cfg, fmt.Errorf("%s: %w", name, err)
		}
	}
	if cfg.MaxRowGroupRows < 0 {
		return cfg, fmt.Errorf("--max-row-group-rows must be >= 0")
	}
	if cfg.ExperimentDir == "" {
		cfg.ExperimentDir = defaultExperimentDir(cfg)
	}
	cfg.ResultDir = filepath.Join(cfg.ExperimentDir, rowsDirName(cfg.Rows))
	cfg.ReportDir = filepath.Join(cfg.ResultDir, "results", distributionDir)
	cfg.ConfigDir = filepath.Join(cfg.ReportDir, "configs")
	cfg.TSVDir = filepath.Join(cfg.ResultDir, "tsvs", distributionDir)
	if cfg.OutputRoot == "" {
		cfg.OutputRoot = filepath.Join(cfg.ResultDir, "parquet", distributionDir)
	}
	if !cfg.KeepOutput {
		if err := requireChildPath(cfg.ResultDir, cfg.OutputRoot, "--output-root"); err != nil {
			return cfg, err
		}
	}
	if (cfg.PlainStatsJSON == "") != (cfg.RLEDictStatsJSON == "") {
		return cfg, fmt.Errorf("--plain-stats-json and --rle-dict-stats-json must be provided together")
	}
	return cfg, nil
}

func run(cfg config) error {
	started := time.Now()
	for _, dir := range []string{cfg.ExperimentDir, cfg.ResultDir, cfg.ReportDir, cfg.ConfigDir, cfg.TSVDir, cfg.OutputRoot, filepath.Join(cfg.ReportDir, "stats"), filepath.Join(cfg.ReportDir, "images"), filepath.Join(cfg.ReportDir, "logs")} {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}

	plainRun := writerRun{
		Name:      plainLabel,
		Encoding:  "plain",
		StatsPath: cfg.PlainStatsJSON,
		OutputDir: filepath.Join(cfg.OutputRoot, "plain-zstd"),
		LogPath:   filepath.Join(cfg.ReportDir, "logs", "plain-zstd.log"),
	}
	rleRun := writerRun{
		Name:      rleDictLabel,
		Encoding:  "rle-dict",
		StatsPath: cfg.RLEDictStatsJSON,
		OutputDir: filepath.Join(cfg.OutputRoot, "rle-dict-zstd"),
		LogPath:   filepath.Join(cfg.ReportDir, "logs", "rle-dict-zstd.log"),
	}
	if plainRun.StatsPath == "" {
		plainRun.StatsPath = filepath.Join(cfg.ReportDir, "stats", "plain-zstd_writer-stats.json")
		rleRun.StatsPath = filepath.Join(cfg.ReportDir, "stats", "rle-dict-zstd_writer-stats.json")
		toolPath, err := buildWriterTool()
		if err != nil {
			return err
		}
		defer os.Remove(toolPath)

		fmt.Printf("running %s writer stats run\n", plainRun.Name)
		if err := runWriter(toolPath, cfg, &plainRun); err != nil {
			return err
		}
		fmt.Printf("running %s writer stats run\n", rleRun.Name)
		if err := runWriter(toolPath, cfg, &rleRun); err != nil {
			return err
		}
		if !cfg.KeepOutput {
			if err := cleanOutputRoot(cfg.OutputRoot); err != nil {
				return err
			}
		}
	} else {
		fmt.Printf("using existing writer stats JSON files\n")
	}

	plainStats, err := loadStats(plainRun.StatsPath)
	if err != nil {
		return err
	}
	rleStats, err := loadStats(rleRun.StatsPath)
	if err != nil {
		return err
	}
	distributions := compareSnapshots(plainStats, rleStats)
	sort.Slice(distributions, func(i, j int) bool {
		if distributions[i].RLEDictWindowWinPct != distributions[j].RLEDictWindowWinPct {
			return distributions[i].RLEDictWindowWinPct > distributions[j].RLEDictWindowWinPct
		}
		return distributions[i].Column < distributions[j].Column
	})

	date := started.Format("2006-01-02")
	base := fmt.Sprintf("%s_rows-%d_plain-zstd_vs_rle-dict-zstd_page-distribution", date, cfg.Rows)
	tsvPath := filepath.Join(cfg.TSVDir, base+".tsv")
	mdPath := filepath.Join(cfg.ReportDir, base+".md")
	svgPath := filepath.Join(cfg.ReportDir, "images", base+".svg")
	if err := writeDistributionTSV(tsvPath, distributions); err != nil {
		return err
	}
	if err := writeDistributionSVG(svgPath, distributions); err != nil {
		return err
	}
	if err := writeMarkdown(mdPath, cfg, plainRun, rleRun, distributions, tsvPath, svgPath, started, time.Now()); err != nil {
		return err
	}
	fmt.Printf("wrote page distribution markdown: %s\n", mdPath)
	fmt.Printf("wrote page distribution TSV: %s\n", tsvPath)
	fmt.Printf("wrote page distribution SVG: %s\n", svgPath)
	return nil
}

func buildWriterTool() (string, error) {
	root, err := os.Getwd()
	if err != nil {
		return "", err
	}
	toolPath := filepath.Join(os.TempDir(), fmt.Sprintf("clickbench-parquet-writer-page-distribution-%d", os.Getpid()))
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

func runWriter(toolPath string, cfg config, run *writerRun) error {
	if !cfg.KeepOutput {
		if err := removeOutputDir(cfg.OutputRoot, run.OutputDir); err != nil {
			return err
		}
	}
	if err := os.MkdirAll(filepath.Dir(run.LogPath), 0o755); err != nil {
		return err
	}
	logFile, err := os.Create(run.LogPath)
	if err != nil {
		return err
	}
	defer logFile.Close()

	args := []string{
		"--input", cfg.Input,
		"--rows", strconv.FormatInt(cfg.Rows, 10),
		"--output-dir", run.OutputDir,
		"--results-dir", cfg.ConfigDir,
		"--tsv-dir", cfg.TSVDir,
		"--max-page-size", cfg.MaxPageSize,
		"--max-dictionary-page-size", cfg.MaxDictSize,
		"--max-row-group-rows", strconv.FormatInt(cfg.MaxRowGroupRows, 10),
		"--max-row-group-size", cfg.MaxRowGroupSize,
		"--max-file-size", cfg.MaxFileSize,
		"--compression", "zstd",
		"--zstd-level", strconv.Itoa(cfg.ZstdLevel),
		"--int-encoding", run.Encoding,
		"--date-encoding", run.Encoding,
		"--timestamp-encoding", run.Encoding,
		"--string-encoding", run.Encoding,
		"--writer-stats-json", run.StatsPath,
	}
	if cfg.Verify {
		args = append(args, "--verify")
	}
	if !cfg.GeneratePDF {
		args = append(args, "--generate-pdf=false")
	}
	started := time.Now()
	cmd := exec.Command(toolPath, args...)
	cmd.Stdout = logFile
	cmd.Stderr = logFile
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("run %s: %w (log: %s)", run.Name, err, run.LogPath)
	}
	run.Elapsed = time.Since(started)
	return nil
}

func loadStats(path string) (statsSnapshot, error) {
	f, err := os.Open(path)
	if err != nil {
		return statsSnapshot{}, err
	}
	defer f.Close()
	var snapshot statsSnapshot
	if err := json.NewDecoder(f).Decode(&snapshot); err != nil {
		return statsSnapshot{}, err
	}
	return snapshot, nil
}

func compareSnapshots(plain, rle statsSnapshot) []columnDistribution {
	plainByColumn := pagesByColumn(plain, false)
	rleByColumn := pagesByColumn(rle, true)
	types := columnTypes(plain, rle)

	columnSet := make(map[string]struct{})
	for column := range plainByColumn {
		columnSet[column] = struct{}{}
	}
	for column := range rleByColumn {
		columnSet[column] = struct{}{}
	}
	columns := make([]string, 0, len(columnSet))
	for column := range columnSet {
		columns = append(columns, column)
	}
	sort.Strings(columns)

	distributions := make([]columnDistribution, 0, len(columns))
	for _, column := range columns {
		dist := compareColumn(column, types[column], plainByColumn[column], rleByColumn[column])
		distributions = append(distributions, dist)
	}
	return distributions
}

func pagesByColumn(snapshot statsSnapshot, includeAmortizedDictionary bool) map[string][]pageCost {
	byColumn := make(map[string][]pageCost, len(snapshot.Columns))
	for _, col := range snapshot.Columns {
		pages := make([]pageCost, 0, len(col.Pages))
		for _, page := range col.Pages {
			if page.NumRows <= 0 {
				continue
			}
			start := page.AbsoluteFirstRowIndex
			if start == 0 && page.RowGroupFirstRowIndex != 0 {
				start = page.RowGroupFirstRowIndex + page.FirstRowIndex
			}
			cost := float64(page.CompressedPageBytesAfterCodec)
			if includeAmortizedDictionary {
				cost = page.CompressedPageBytesWithAmortizedDictionary
				if cost == 0 {
					cost = float64(page.CompressedPageBytesAfterCodec) + page.AmortizedDictionaryCompressedBytes
				}
			}
			pages = append(pages, pageCost{
				Column:        col.Name,
				Type:          col.Type,
				RowGroupIndex: page.RowGroupIndex,
				PageIndex:     page.PageIndex,
				Start:         start,
				End:           start + page.NumRows,
				Rows:          page.NumRows,
				Cost:          cost,
				Encoding:      page.Encoding,
			})
		}
		sort.Slice(pages, func(i, j int) bool {
			if pages[i].Start != pages[j].Start {
				return pages[i].Start < pages[j].Start
			}
			if pages[i].End != pages[j].End {
				return pages[i].End < pages[j].End
			}
			return pages[i].PageIndex < pages[j].PageIndex
		})
		byColumn[col.Name] = pages
	}
	return byColumn
}

func columnTypes(snapshots ...statsSnapshot) map[string]string {
	types := make(map[string]string)
	for _, snapshot := range snapshots {
		for _, col := range snapshot.Columns {
			if _, ok := types[col.Name]; !ok && col.Type != "" {
				types[col.Name] = col.Type
			}
		}
	}
	return types
}

func compareColumn(column, typ string, plainPages, rlePages []pageCost) columnDistribution {
	dist := columnDistribution{
		Column:       column,
		Type:         typ,
		PlainPages:   len(plainPages),
		RLEDictPages: len(rlePages),
	}
	for _, page := range plainPages {
		dist.PlainTotalPageBytes += page.Cost
	}
	for _, page := range rlePages {
		dist.RLEDictTotalPageBytes += page.Cost
	}

	dist.ExactMatchedPages, dist.ExactPlainWins, dist.ExactRLEDictWins, dist.ExactTies, dist.UnmatchedPlainPages, dist.UnmatchedRLEDictPages = exactMatchCounts(plainPages, rlePages)

	i, j := 0, 0
	for i < len(plainPages) && j < len(rlePages) {
		p := plainPages[i]
		r := rlePages[j]
		if p.End <= r.Start {
			i++
			continue
		}
		if r.End <= p.Start {
			j++
			continue
		}

		start := maxInt64(p.Start, r.Start)
		end := minInt64(p.End, r.End)
		if end > start {
			rows := end - start
			plainCost := p.Cost * float64(rows) / float64(p.Rows)
			rleCost := r.Cost * float64(rows) / float64(r.Rows)
			dist.ComparisonWindows++
			dist.RowsCompared += rows
			dist.PlainAllocatedBytes += plainCost
			dist.RLEDictAllocatedBytes += rleCost
			switch winner(plainCost, rleCost) {
			case pageWinnerPlain:
				dist.PlainWindowWins++
				dist.PlainRowsWon += rows
			case pageWinnerRLEDict:
				dist.RLEDictWindowWins++
				dist.RLEDictRowsWon += rows
			default:
				dist.TieWindowWins++
				dist.TieRowsWon += rows
			}
		}

		if p.End <= r.End {
			i++
		}
		if r.End <= p.End {
			j++
		}
	}

	if dist.ComparisonWindows > 0 {
		dist.PlainWindowWinPct = float64(dist.PlainWindowWins) / float64(dist.ComparisonWindows) * 100
		dist.RLEDictWindowWinPct = float64(dist.RLEDictWindowWins) / float64(dist.ComparisonWindows) * 100
		dist.TieWindowWinPct = float64(dist.TieWindowWins) / float64(dist.ComparisonWindows) * 100
	}
	if dist.RowsCompared > 0 {
		dist.PlainRowWinPct = float64(dist.PlainRowsWon) / float64(dist.RowsCompared) * 100
		dist.RLEDictRowWinPct = float64(dist.RLEDictRowsWon) / float64(dist.RowsCompared) * 100
		dist.TieRowWinPct = float64(dist.TieRowsWon) / float64(dist.RowsCompared) * 100
	}
	dist.WinnerByAllocatedBytes = "n/a"
	if dist.PlainAllocatedBytes > 0 || dist.RLEDictAllocatedBytes > 0 {
		dist.WinnerByAllocatedBytes = winner(dist.PlainAllocatedBytes, dist.RLEDictAllocatedBytes)
	}
	if dist.PlainAllocatedBytes > 0 {
		dist.RLEDictAllocatedVsPlainPct = ((dist.RLEDictAllocatedBytes / dist.PlainAllocatedBytes) - 1) * 100
	}
	return dist
}

func exactMatchCounts(plainPages, rlePages []pageCost) (matched, plainWins, rleWins, ties, unmatchedPlain, unmatchedRLE int) {
	rleByRange := make(map[string][]pageCost, len(rlePages))
	for _, page := range rlePages {
		key := rangeKey(page)
		rleByRange[key] = append(rleByRange[key], page)
	}
	usedRLE := make(map[string]int, len(rleByRange))
	for _, plain := range plainPages {
		key := rangeKey(plain)
		candidates := rleByRange[key]
		used := usedRLE[key]
		if used >= len(candidates) {
			unmatchedPlain++
			continue
		}
		rle := candidates[used]
		usedRLE[key]++
		matched++
		switch winner(plain.Cost, rle.Cost) {
		case pageWinnerPlain:
			plainWins++
		case pageWinnerRLEDict:
			rleWins++
		default:
			ties++
		}
	}
	for key, candidates := range rleByRange {
		unmatchedRLE += len(candidates) - usedRLE[key]
	}
	return matched, plainWins, rleWins, ties, unmatchedPlain, unmatchedRLE
}

func rangeKey(page pageCost) string {
	return strconv.FormatInt(page.Start, 10) + ":" + strconv.FormatInt(page.End, 10)
}

func winner(plainCost, rleCost float64) string {
	const epsilon = 1e-9
	if math.Abs(plainCost-rleCost) <= epsilon {
		return pageWinnerTie
	}
	if plainCost < rleCost {
		return pageWinnerPlain
	}
	return pageWinnerRLEDict
}

func writeDistributionTSV(path string, rows []columnDistribution) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	w.Comma = '\t'
	header := []string{
		"column",
		"type",
		"plain_pages",
		"rle_dict_pages",
		"comparison_windows",
		"rows_compared",
		"plain_window_wins",
		"rle_dict_window_wins",
		"tie_window_wins",
		"plain_window_win_pct",
		"rle_dict_window_win_pct",
		"tie_window_win_pct",
		"plain_rows_won",
		"rle_dict_rows_won",
		"tie_rows_won",
		"plain_row_win_pct",
		"rle_dict_row_win_pct",
		"tie_row_win_pct",
		"plain_allocated_compressed_bytes",
		"rle_dict_allocated_compressed_bytes_with_amortized_dictionary",
		"rle_dict_allocated_vs_plain_pct",
		"winner_by_allocated_bytes",
		"plain_total_page_compressed_bytes",
		"rle_dict_total_page_compressed_bytes_with_amortized_dictionary",
		"exact_matched_pages",
		"exact_plain_wins",
		"exact_rle_dict_wins",
		"exact_ties",
		"unmatched_plain_pages",
		"unmatched_rle_dict_pages",
	}
	if err := w.Write(header); err != nil {
		return err
	}
	for _, row := range rows {
		record := []string{
			row.Column,
			row.Type,
			strconv.Itoa(row.PlainPages),
			strconv.Itoa(row.RLEDictPages),
			strconv.Itoa(row.ComparisonWindows),
			strconv.FormatInt(row.RowsCompared, 10),
			strconv.Itoa(row.PlainWindowWins),
			strconv.Itoa(row.RLEDictWindowWins),
			strconv.Itoa(row.TieWindowWins),
			formatFloat(row.PlainWindowWinPct),
			formatFloat(row.RLEDictWindowWinPct),
			formatFloat(row.TieWindowWinPct),
			strconv.FormatInt(row.PlainRowsWon, 10),
			strconv.FormatInt(row.RLEDictRowsWon, 10),
			strconv.FormatInt(row.TieRowsWon, 10),
			formatFloat(row.PlainRowWinPct),
			formatFloat(row.RLEDictRowWinPct),
			formatFloat(row.TieRowWinPct),
			formatFloat(row.PlainAllocatedBytes),
			formatFloat(row.RLEDictAllocatedBytes),
			formatFloat(row.RLEDictAllocatedVsPlainPct),
			row.WinnerByAllocatedBytes,
			formatFloat(row.PlainTotalPageBytes),
			formatFloat(row.RLEDictTotalPageBytes),
			strconv.Itoa(row.ExactMatchedPages),
			strconv.Itoa(row.ExactPlainWins),
			strconv.Itoa(row.ExactRLEDictWins),
			strconv.Itoa(row.ExactTies),
			strconv.Itoa(row.UnmatchedPlainPages),
			strconv.Itoa(row.UnmatchedRLEDictPages),
		}
		if err := w.Write(record); err != nil {
			return err
		}
	}
	w.Flush()
	return w.Error()
}

func writeMarkdown(path string, cfg config, plainRun, rleRun writerRun, rows []columnDistribution, tsvPath, svgPath string, started, finished time.Time) error {
	var b strings.Builder
	reportDir := filepath.Dir(path)
	fmt.Fprintf(&b, "# Page-Level Encoding Distribution\n\n")
	fmt.Fprintf(&b, "- Started: `%s`\n", started.Format(time.RFC3339))
	fmt.Fprintf(&b, "- Elapsed: `%s`\n", finished.Sub(started).Round(time.Millisecond))
	fmt.Fprintf(&b, "- Rows: `%d`\n", cfg.Rows)
	fmt.Fprintf(&b, "- Compared configs: `%s` vs `%s`\n", plainLabel, rleDictLabel)
	fmt.Fprintf(&b, "- ZSTD level: `%d`\n", cfg.ZstdLevel)
	fmt.Fprintf(&b, "- Max page size: `%s`\n", cfg.MaxPageSize)
	fmt.Fprintf(&b, "- Max dictionary page size: `%s`\n", cfg.MaxDictSize)
	fmt.Fprintf(&b, "- Max row group rows: `%d`\n", cfg.MaxRowGroupRows)
	fmt.Fprintf(&b, "- Max row group size: `%s`\n", cfg.MaxRowGroupSize)
	fmt.Fprintf(&b, "- Max file size: `%s`\n", cfg.MaxFileSize)
	fmt.Fprintf(&b, "- TSV: [%s](%s)\n", filepath.Base(tsvPath), markdownLinkTarget(reportDir, tsvPath))
	fmt.Fprintf(&b, "- Plain stats JSON: [%s](%s)\n", filepath.Base(plainRun.StatsPath), markdownLinkTarget(reportDir, plainRun.StatsPath))
	fmt.Fprintf(&b, "- RLE dict stats JSON: [%s](%s)\n", filepath.Base(rleRun.StatsPath), markdownLinkTarget(reportDir, rleRun.StatsPath))
	if plainRun.Elapsed > 0 {
		fmt.Fprintf(&b, "- Plain writer elapsed: `%s`\n", plainRun.Elapsed.Round(time.Millisecond))
	}
	if rleRun.Elapsed > 0 {
		fmt.Fprintf(&b, "- RLE dict writer elapsed: `%s`\n", rleRun.Elapsed.Round(time.Millisecond))
	}
	fmt.Fprintf(&b, "\n")

	fmt.Fprintf(&b, "## Method\n\n")
	fmt.Fprintf(&b, "The primary distribution uses overlap windows from the union of page row ranges for each column. For each overlapping row span, the page compressed byte cost is allocated in proportion to row overlap. The RLE dictionary cost uses `compressed_page_bytes_with_amortized_dictionary`, meaning the compressed dictionary page bytes for a column chunk are divided evenly across that chunk's data pages before comparison.\n\n")
	fmt.Fprintf(&b, "`exact_matched_pages` counts only pages where both runs produced the same absolute row range. Exact matches are useful as a sanity check, but the overlap-window distribution is the full comparison when page boundaries differ.\n\n")

	fmt.Fprintf(&b, "## Distribution Chart\n\n")
	writeImage(&b, "Page-window winner distribution by column", svgPath, reportDir)

	fmt.Fprintf(&b, "## Column Distribution\n\n")
	fmt.Fprintf(&b, "| Column | Type | Windows | Plain wins | RLE dict wins | Ties | Rows compared | Row-weighted plain | Row-weighted RLE dict | Allocated plain bytes | Allocated RLE dict bytes | RLE dict vs plain | Exact matches | Unmatched plain | Unmatched RLE dict |\n")
	fmt.Fprintf(&b, "| --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |\n")
	for _, row := range rows {
		fmt.Fprintf(&b, "| `%s` | `%s` | `%d` | `%d` (%s) | `%d` (%s) | `%d` (%s) | `%d` | `%s` | `%s` | `%s` | `%s` | `%s` | `%d` | `%d` | `%d` |\n",
			row.Column,
			row.Type,
			row.ComparisonWindows,
			row.PlainWindowWins,
			formatPercent(row.PlainWindowWinPct),
			row.RLEDictWindowWins,
			formatPercent(row.RLEDictWindowWinPct),
			row.TieWindowWins,
			formatPercent(row.TieWindowWinPct),
			row.RowsCompared,
			formatPercent(row.PlainRowWinPct),
			formatPercent(row.RLEDictRowWinPct),
			formatBytesFloat(row.PlainAllocatedBytes),
			formatBytesFloat(row.RLEDictAllocatedBytes),
			formatPercent(row.RLEDictAllocatedVsPlainPct),
			row.ExactMatchedPages,
			row.UnmatchedPlainPages,
			row.UnmatchedRLEDictPages,
		)
	}
	data := []byte(b.String())
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func writeDistributionSVG(path string, rows []columnDistribution) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	const (
		width       = 1100
		leftMargin  = 230
		rightMargin = 40
		topMargin   = 58
		rowHeight   = 22
		bottom      = 54
	)
	height := topMargin + bottom + rowHeight*len(rows)
	if len(rows) == 0 {
		height = 180
	}
	plotWidth := float64(width - leftMargin - rightMargin)
	var b strings.Builder
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d">`+"\n", width, height, width, height)
	fmt.Fprintf(&b, `<rect width="100%%" height="100%%" fill="#ffffff"/>`+"\n")
	fmt.Fprintf(&b, `<text x="%d" y="28" font-family="Arial, sans-serif" font-size="18" font-weight="700" fill="#111827">%s</text>`+"\n", leftMargin, html.EscapeString("Page-window winner distribution by column"))
	fmt.Fprintf(&b, `<rect x="%d" y="40" width="14" height="10" fill="#2563eb"/><text x="%d" y="49" font-family="Arial, sans-serif" font-size="12" fill="#374151">plain + zstd</text>`+"\n", leftMargin, leftMargin+20)
	fmt.Fprintf(&b, `<rect x="%d" y="40" width="14" height="10" fill="#16a34a"/><text x="%d" y="49" font-family="Arial, sans-serif" font-size="12" fill="#374151">rle-dict + zstd</text>`+"\n", leftMargin+130, leftMargin+150)
	fmt.Fprintf(&b, `<rect x="%d" y="40" width="14" height="10" fill="#9ca3af"/><text x="%d" y="49" font-family="Arial, sans-serif" font-size="12" fill="#374151">tie</text>`+"\n", leftMargin+285, leftMargin+305)
	for i, row := range rows {
		y := topMargin + i*rowHeight
		label := row.Column
		if len(label) > 30 {
			label = label[:27] + "..."
		}
		fmt.Fprintf(&b, `<text x="%d" y="%d" font-family="Arial, sans-serif" font-size="11" text-anchor="end" fill="#374151">%s</text>`+"\n", leftMargin-10, y+13, html.EscapeString(label))
		fmt.Fprintf(&b, `<rect x="%d" y="%d" width="%.2f" height="14" fill="#f3f4f6"/>`+"\n", leftMargin, y+2, plotWidth)
		if row.ComparisonWindows == 0 {
			continue
		}
		x := float64(leftMargin)
		plainWidth := plotWidth * float64(row.PlainWindowWins) / float64(row.ComparisonWindows)
		rleWidth := plotWidth * float64(row.RLEDictWindowWins) / float64(row.ComparisonWindows)
		tieWidth := plotWidth * float64(row.TieWindowWins) / float64(row.ComparisonWindows)
		fmt.Fprintf(&b, `<rect x="%.2f" y="%d" width="%.2f" height="14" fill="#2563eb"/>`+"\n", x, y+2, plainWidth)
		x += plainWidth
		fmt.Fprintf(&b, `<rect x="%.2f" y="%d" width="%.2f" height="14" fill="#16a34a"/>`+"\n", x, y+2, rleWidth)
		x += rleWidth
		fmt.Fprintf(&b, `<rect x="%.2f" y="%d" width="%.2f" height="14" fill="#9ca3af"/>`+"\n", x, y+2, tieWidth)
		if row.ComparisonWindows > 0 {
			fmt.Fprintf(&b, `<text x="%d" y="%d" font-family="Arial, sans-serif" font-size="10" fill="#6b7280">%d</text>`+"\n", width-rightMargin+4, y+13, row.ComparisonWindows)
		}
	}
	fmt.Fprintf(&b, `<text x="%d" y="%d" font-family="Arial, sans-serif" font-size="11" fill="#6b7280">bar width = share of overlap comparison windows won</text>`+"\n", leftMargin, height-20)
	fmt.Fprintf(&b, `</svg>`+"\n")
	return os.WriteFile(path, []byte(b.String()), 0o644)
}

func writeImage(b *strings.Builder, alt, path, reportDir string) {
	fmt.Fprintf(b, "![%s](%s)\n\n", alt, markdownLinkTarget(reportDir, path))
}

func markdownLinkTarget(fromDir, path string) string {
	rel, err := filepath.Rel(fromDir, path)
	if err != nil {
		return filepath.ToSlash(path)
	}
	return strings.ReplaceAll(filepath.ToSlash(rel), " ", "%20")
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

func defaultExperimentDir(cfg config) string {
	return filepath.Join(
		"encoding_experiment",
		fmt.Sprintf("page-%s-rgsize-%s-file-%s-dictpage-%s",
			sanitizeFilename(cfg.MaxPageSize),
			sanitizeFilename(cfg.MaxRowGroupSize),
			sanitizeFilename(cfg.MaxFileSize),
			sanitizeFilename(cfg.MaxDictSize),
		),
	)
}

func rowsDirName(rows int64) string {
	return fmt.Sprintf("%d_rows", rows)
}

func sanitizeFilename(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
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

func formatFloat(v float64) string {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return ""
	}
	return strconv.FormatFloat(v, 'f', 6, 64)
}

func formatPercent(v float64) string {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return "n/a"
	}
	return fmt.Sprintf("%.2f%%", v)
}

func formatBytesFloat(v float64) string {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return "n/a"
	}
	return formatBytes(int64(math.Round(v)))
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

func minInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func exitf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
