package parquet

import (
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"slices"
	"sort"
	"strconv"
	"sync"

	"parquet_compress/parquet-go/deprecated"
)

// WriterStats collects in-memory statistics about pages written by a Writer.
//
// The collector is intentionally out-of-band: it is not serialized into Parquet
// files and has no effect unless attached with CaptureWriterStats.
type WriterStats struct {
	mu           sync.Mutex
	nextRowGroup int64
	nextRowIndex int64
	columns      map[int]*writerStatsColumn
	errors       []string
}

// WriterStatsSnapshot is a stable, JSON-friendly copy of WriterStats.
type WriterStatsSnapshot struct {
	Columns []WriterColumnStats `json:"columns"`
	Errors  []string            `json:"errors,omitempty"`
}

// WriterColumnStats contains collected stats for a single leaf column.
type WriterColumnStats struct {
	ColumnIndex       int                   `json:"column_index"`
	Path              []string              `json:"path"`
	Name              string                `json:"name"`
	Type              string                `json:"type"`
	PhysicalType      string                `json:"physical_type"`
	SortedAscending   bool                  `json:"sorted_ascending"`
	SortedDescending  bool                  `json:"sorted_descending"`
	MinValueLength    int                   `json:"min_value_length"`
	MedianValueLength float64               `json:"median_value_length"`
	MaxValueLength    int                   `json:"max_value_length"`
	RowGroups         []WriterRowGroupStats `json:"row_groups"`
	Pages             []WriterPageStats     `json:"pages"`
}

// WriterRowGroupStats contains one row of stats for a column in one row group.
type WriterRowGroupStats struct {
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

// WriterPageStats contains one row of stats for an encoded data page.
type WriterPageStats struct {
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

type writerStatsColumn struct {
	stats            WriterColumnStats
	hasLength        bool
	lengthCounts     map[int]int64
	lengthValueCount int64
	hasLast          bool
	last             Value
}

type writerStatsColumnAccumulator struct {
	numRows                             int64
	unique                              map[string]struct{}
	pages                               []WriterPageStats
	pageCardinalityMin                  int64
	pageCardinalityMax                  int64
	minLength                           int
	maxLength                           int
	lengthCounts                        map[int]int64
	lengthValueCount                    int64
	hasLength                           bool
	hasFirst                            bool
	first                               Value
	hasLast                             bool
	last                                Value
	sortedAscending                     bool
	sortedDescending                    bool
	encodedDataPageBytesBeforeCodec     int64
	compressedDataPageBytesAfterCodec   int64
	dictionaryPageCount                 int
	dictionaryEncodedBytesBeforeCodec   int64
	dictionaryCompressedBytesAfterCodec int64
	errors                              []string
}

type writerStatsPageLayout struct {
	pageType                      string
	encoding                      string
	encodingID                    int32
	headerBytes                   int64
	encodedBodyBytesBeforeCodec   int64
	compressedBodyBytesAfterCodec int64
	encodedPageBytesBeforeCodec   int64
	compressedPageBytesAfterCodec int64
}

// NewWriterStats constructs a writer stats collector.
func NewWriterStats() *WriterStats {
	return &WriterStats{
		columns: make(map[int]*writerStatsColumn),
	}
}

// CaptureWriterStats attaches an in-memory writer stats collector to a writer.
func CaptureWriterStats(stats *WriterStats) WriterOption {
	return writerOption(func(config *WriterConfig) { config.WriterStats = stats })
}

// Snapshot returns a stable copy of collected writer stats.
func (s *WriterStats) Snapshot() WriterStatsSnapshot {
	if s == nil {
		return WriterStatsSnapshot{}
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	columns := make([]WriterColumnStats, 0, len(s.columns))
	for _, col := range s.columns {
		copyCol := col.stats
		copyCol.MedianValueLength = medianLengthFromCounts(col.lengthCounts, col.lengthValueCount)
		copyCol.Path = slices.Clone(copyCol.Path)
		copyCol.RowGroups = slices.Clone(copyCol.RowGroups)
		copyCol.Pages = slices.Clone(copyCol.Pages)
		columns = append(columns, copyCol)
	}
	sort.Slice(columns, func(i, j int) bool {
		return columns[i].ColumnIndex < columns[j].ColumnIndex
	})
	return WriterStatsSnapshot{
		Columns: columns,
		Errors:  slices.Clone(s.errors),
	}
}

func (s *WriterStats) reserveRowGroup(numRows int64) (int64, int64) {
	if s == nil {
		return 0, 0
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	index := s.nextRowGroup
	firstRowIndex := s.nextRowIndex
	s.nextRowGroup++
	s.nextRowIndex += numRows
	return index, firstRowIndex
}

func (s *WriterStats) finishColumnRowGroup(columnIndex int, path columnPath, typ Type, rowGroupIndex, rowGroupFirstRowIndex int64, acc *writerStatsColumnAccumulator) {
	if s == nil || acc == nil || len(acc.pages) == 0 {
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	col := s.columns[columnIndex]
	if col == nil {
		pathCopy := make([]string, len(path))
		copy(pathCopy, path)
		col = &writerStatsColumn{
			stats: WriterColumnStats{
				ColumnIndex:      columnIndex,
				Path:             pathCopy,
				Name:             columnPathString(path),
				Type:             typ.String(),
				PhysicalType:     typ.Kind().String(),
				SortedAscending:  true,
				SortedDescending: true,
			},
		}
		s.columns[columnIndex] = col
	}

	amortizedDictionaryEncodedBytes := float64(0)
	amortizedDictionaryCompressedBytes := float64(0)
	if len(acc.pages) > 0 {
		amortizedDictionaryEncodedBytes = float64(acc.dictionaryEncodedBytesBeforeCodec) / float64(len(acc.pages))
		amortizedDictionaryCompressedBytes = float64(acc.dictionaryCompressedBytesAfterCodec) / float64(len(acc.pages))
	}
	for i := range acc.pages {
		acc.pages[i].RowGroupIndex = rowGroupIndex
		acc.pages[i].RowGroupFirstRowIndex = rowGroupFirstRowIndex
		acc.pages[i].AbsoluteFirstRowIndex = rowGroupFirstRowIndex + acc.pages[i].FirstRowIndex
		acc.pages[i].DataPageCountInColumnChunk = len(acc.pages)
		acc.pages[i].DictionaryPageCount = acc.dictionaryPageCount
		acc.pages[i].DictionaryEncodedBytesBeforeCodec = acc.dictionaryEncodedBytesBeforeCodec
		acc.pages[i].DictionaryCompressedBytesAfterCodec = acc.dictionaryCompressedBytesAfterCodec
		acc.pages[i].AmortizedDictionaryEncodedBytes = amortizedDictionaryEncodedBytes
		acc.pages[i].AmortizedDictionaryCompressedBytes = amortizedDictionaryCompressedBytes
		acc.pages[i].EncodedPageBytesWithAmortizedDictionary = float64(acc.pages[i].EncodedPageBytesBeforeCodec) + amortizedDictionaryEncodedBytes
		acc.pages[i].CompressedPageBytesWithAmortizedDictionary = float64(acc.pages[i].CompressedPageBytesAfterCodec) + amortizedDictionaryCompressedBytes
	}

	if col.hasLast && acc.hasFirst {
		if typ.Compare(col.last, acc.first) > 0 {
			col.stats.SortedAscending = false
		}
		if typ.Compare(col.last, acc.first) < 0 {
			col.stats.SortedDescending = false
		}
	}
	if !acc.sortedAscending {
		col.stats.SortedAscending = false
	}
	if !acc.sortedDescending {
		col.stats.SortedDescending = false
	}
	if acc.hasLast {
		col.hasLast = true
		col.last = acc.last
	}
	if acc.hasLength {
		if !col.hasLength || acc.minLength < col.stats.MinValueLength {
			col.stats.MinValueLength = acc.minLength
		}
		if !col.hasLength || acc.maxLength > col.stats.MaxValueLength {
			col.stats.MaxValueLength = acc.maxLength
		}
		col.hasLength = true
		if col.lengthCounts == nil {
			col.lengthCounts = make(map[int]int64, len(acc.lengthCounts))
		}
		for length, count := range acc.lengthCounts {
			col.lengthCounts[length] += count
			col.lengthValueCount += count
		}
	}

	col.stats.RowGroups = append(col.stats.RowGroups, WriterRowGroupStats{
		RowGroupIndex:                       rowGroupIndex,
		FirstRowIndex:                       rowGroupFirstRowIndex,
		NumRows:                             acc.numRows,
		Cardinality:                         int64(len(acc.unique)),
		PageCount:                           len(acc.pages),
		PageCardinalityMin:                  acc.pageCardinalityMin,
		PageCardinalityMax:                  acc.pageCardinalityMax,
		MinValueLength:                      acc.minLength,
		MedianValueLength:                   medianLengthFromCounts(acc.lengthCounts, acc.lengthValueCount),
		MaxValueLength:                      acc.maxLength,
		EncodedDataPageBytesBeforeCodec:     acc.encodedDataPageBytesBeforeCodec,
		CompressedDataPageBytesAfterCodec:   acc.compressedDataPageBytesAfterCodec,
		DictionaryPageCount:                 acc.dictionaryPageCount,
		DictionaryEncodedBytesBeforeCodec:   acc.dictionaryEncodedBytesBeforeCodec,
		DictionaryCompressedBytesAfterCodec: acc.dictionaryCompressedBytesAfterCodec,
		AmortizedDictionaryEncodedBytes:     amortizedDictionaryEncodedBytes,
		AmortizedDictionaryCompressedBytes:  amortizedDictionaryCompressedBytes,
		EncodedBytesWithDictionary:          acc.encodedDataPageBytesBeforeCodec + acc.dictionaryEncodedBytesBeforeCodec,
		CompressedBytesWithDictionary:       acc.compressedDataPageBytesAfterCodec + acc.dictionaryCompressedBytesAfterCodec,
	})
	col.stats.Pages = append(col.stats.Pages, acc.pages...)
	s.errors = append(s.errors, acc.errors...)
}

func (s *WriterStats) recordError(format string, args ...any) {
	if s == nil {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.errors = append(s.errors, fmt.Sprintf(format, args...))
}

func newWriterStatsColumnAccumulator() *writerStatsColumnAccumulator {
	return &writerStatsColumnAccumulator{
		unique:           make(map[string]struct{}),
		lengthCounts:     make(map[int]int64),
		sortedAscending:  true,
		sortedDescending: true,
	}
}

func (a *writerStatsColumnAccumulator) recordDataPage(columnType Type, firstRowIndex int64, page Page, layout writerStatsPageLayout) {
	pageStat := WriterPageStats{
		PageIndex:                     len(a.pages),
		FirstRowIndex:                 firstRowIndex,
		NumRows:                       page.NumRows(),
		NumValues:                     page.NumValues(),
		PageType:                      layout.pageType,
		Encoding:                      layout.encoding,
		EncodingID:                    layout.encodingID,
		HeaderBytes:                   layout.headerBytes,
		EncodedBodyBytesBeforeCodec:   layout.encodedBodyBytesBeforeCodec,
		CompressedBodyBytesAfterCodec: layout.compressedBodyBytesAfterCodec,
		EncodedPageBytesBeforeCodec:   layout.encodedPageBytesBeforeCodec,
		CompressedPageBytesAfterCodec: layout.compressedPageBytesAfterCodec,
	}
	if minValue, maxValue, ok := page.Bounds(); ok {
		pageStat.HasBounds = true
		pageStat.MinValue = minValue.String()
		pageStat.MaxValue = maxValue.String()
		pageStat.MinValueBytes = hex.EncodeToString(minValue.Bytes())
		pageStat.MaxValueBytes = hex.EncodeToString(maxValue.Bytes())
		if n, ok := valueNumeric(minValue); ok {
			pageStat.HasNumeric = true
			pageStat.MinNumeric = n
		}
		if n, ok := valueNumeric(maxValue); ok {
			pageStat.HasNumeric = true
			pageStat.MaxNumeric = n
		}
	}

	pageUnique := make(map[string]struct{})
	values := page.Values()
	buf := make([]Value, defaultValueBufferSize)
	pageHasLength := false
	for {
		n, err := values.ReadValues(buf)
		for i := 0; i < n; i++ {
			v := buf[i]
			key := valueKey(v)
			a.unique[key] = struct{}{}
			pageUnique[key] = struct{}{}

			length := valueLength(v)
			a.lengthCounts[length]++
			a.lengthValueCount++
			if !a.hasLength || length < a.minLength {
				a.minLength = length
			}
			if !a.hasLength || length > a.maxLength {
				a.maxLength = length
			}
			if !pageHasLength || length < pageStat.MinLength {
				pageStat.MinLength = length
			}
			if !pageHasLength || length > pageStat.MaxLength {
				pageStat.MaxLength = length
			}
			a.hasLength = true
			pageHasLength = true

			if v.IsNull() {
				continue
			}
			stable := cloneValue(v)
			if !a.hasFirst {
				a.first = stable
				a.hasFirst = true
			}
			if a.hasLast {
				if columnType.Compare(a.last, stable) > 0 {
					a.sortedAscending = false
				}
				if columnType.Compare(a.last, stable) < 0 {
					a.sortedDescending = false
				}
			}
			a.last = stable
			a.hasLast = true
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			a.errors = append(a.errors, err.Error())
			break
		}
		if n == 0 {
			a.errors = append(a.errors, io.ErrNoProgress.Error())
			break
		}
	}

	pageStat.Cardinality = int64(len(pageUnique))
	if len(a.pages) == 0 || pageStat.Cardinality < a.pageCardinalityMin {
		a.pageCardinalityMin = pageStat.Cardinality
	}
	if len(a.pages) == 0 || pageStat.Cardinality > a.pageCardinalityMax {
		a.pageCardinalityMax = pageStat.Cardinality
	}
	a.numRows += page.NumRows()
	a.encodedDataPageBytesBeforeCodec += layout.encodedPageBytesBeforeCodec
	a.compressedDataPageBytesAfterCodec += layout.compressedPageBytesAfterCodec
	a.pages = append(a.pages, pageStat)
}

func (a *writerStatsColumnAccumulator) recordDictionaryPage(layout writerStatsPageLayout) {
	a.dictionaryPageCount++
	a.dictionaryEncodedBytesBeforeCodec += layout.encodedPageBytesBeforeCodec
	a.dictionaryCompressedBytesAfterCodec += layout.compressedPageBytesAfterCodec
}

func cloneValue(v Value) Value {
	switch v.Kind() {
	case Boolean:
		return BooleanValue(v.Boolean())
	case Int32:
		return Int32Value(v.Int32())
	case Int64:
		return Int64Value(v.Int64())
	case Int96:
		return Int96Value(v.Int96())
	case Float:
		return FloatValue(v.Float())
	case Double:
		return DoubleValue(v.Double())
	case ByteArray:
		return ByteArrayValue(slices.Clone(v.ByteArray()))
	case FixedLenByteArray:
		return FixedLenByteArrayValue(slices.Clone(v.ByteArray()))
	default:
		return NullValue()
	}
}

func valueKey(v Value) string {
	if v.IsNull() {
		return "null"
	}
	switch v.Kind() {
	case Boolean:
		return "bool:" + strconv.FormatBool(v.Boolean())
	case Int32:
		return "i32:" + strconv.FormatInt(int64(v.Int32()), 10)
	case Int64:
		return "i64:" + strconv.FormatInt(v.Int64(), 10)
	case Int96:
		return "i96:" + deprecated.Int96(v.Int96()).String()
	case Float:
		return "f32:" + strconv.FormatFloat(float64(v.Float()), 'g', -1, 32)
	case Double:
		return "f64:" + strconv.FormatFloat(v.Double(), 'g', -1, 64)
	case ByteArray:
		return "ba:" + string(v.ByteArray())
	case FixedLenByteArray:
		return "flba:" + string(v.ByteArray())
	default:
		return v.Kind().String() + ":" + string(v.AppendBytes(nil))
	}
}

func valueLength(v Value) int {
	if v.IsNull() {
		return 0
	}
	switch v.Kind() {
	case Boolean:
		return 1
	case Int32, Float:
		return 4
	case Int64, Double:
		return 8
	case Int96:
		return 12
	case ByteArray, FixedLenByteArray:
		return len(v.ByteArray())
	default:
		return len(v.Bytes())
	}
}

func medianLengthFromCounts(counts map[int]int64, total int64) float64 {
	if total <= 0 || len(counts) == 0 {
		return 0
	}
	leftRank := (total + 1) / 2
	rightRank := (total + 2) / 2
	var left, right int
	leftSet := false
	var seen int64
	lengths := make([]int, 0, len(counts))
	for length := range counts {
		lengths = append(lengths, length)
	}
	sort.Ints(lengths)
	for _, length := range lengths {
		seen += counts[length]
		if seen >= leftRank && !leftSet {
			left = length
			leftSet = true
		}
		if seen >= rightRank {
			right = length
			break
		}
	}
	return float64(left+right) / 2
}

func valueNumeric(v Value) (float64, bool) {
	if v.IsNull() {
		return 0, false
	}
	switch v.Kind() {
	case Boolean:
		if v.Boolean() {
			return 1, true
		}
		return 0, true
	case Int32:
		return float64(v.Int32()), true
	case Int64:
		return float64(v.Int64()), true
	case Float:
		n := float64(v.Float())
		return n, !math.IsNaN(n)
	case Double:
		n := v.Double()
		return n, !math.IsNaN(n)
	default:
		return 0, false
	}
}
