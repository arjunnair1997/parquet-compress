# Column Top 5 Encoding Rankings

- Experiment: `page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows`
- Source data: [2026-07-03_rows-1000000_encoding-matrix_column-results.tsv](../tsvs/2026-07-03_rows-1000000_encoding-matrix_column-results.tsv)
- Rows: `1,000,000`
- Ranking metric: per-column `compressed_bytes`, after Parquet page encoding and Snappy/ZSTD compression.
- Each numbered item starts with the achieved compressed size for that encoding/compression choice.
- Duplicate matrix rows with the same effective column encoding are collapsed to the best observed row before ranking.
- Encodings in this matrix: `plain`, `rle-dict`, `delta-byte-array`, `delta-length-byte-array`. `delta-binary-packed` was not included.
- Column shape stats: [column_shape_stats.json](column_shape_stats/column_shape_stats.json)

## AdvEngineID (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 3 / 4`
- Page cardinality per row group min/median/max of mins: `1 / 3 / 4`; of maxes: `1 / 3 / 4`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/advengineid_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/advengineid_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/advengineid_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/advengineid_value_length.svg)


Compressed overall:
1. 22,787 B (22.25 KiB) compressed - `zstd-3` + `rle-dict`; 34,609 B (33.80 KiB) encoded; 175.760434x post-compression ratio; 34.708386% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 29,888 B (29.19 KiB) compressed - `snappy` + `rle-dict`; 34,914 B (34.10 KiB) encoded; 134.002041x post-compression ratio; 637.305942% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
3. 30,642 B (29.92 KiB) compressed - `zstd-3` + `plain`; 4,003,579 B (3.82 MiB) encoded; 130.704686x post-compression ratio; 0.176229% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
4. 220,197 B (215.04 KiB) compressed - `snappy` + `plain`; 4,003,714 B (3.82 MiB) encoded; 18.188499x post-compression ratio; 0.076749% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 22,787 B (22.25 KiB) compressed - `rle-dict`; 34,609 B (33.80 KiB) encoded; 175.760434x post-compression ratio; 34.708386% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 30,642 B (29.92 KiB) compressed - `plain`; 4,003,579 B (3.82 MiB) encoded; 130.704686x post-compression ratio; 0.176229% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`

Snappy:
1. 29,888 B (29.19 KiB) compressed - `rle-dict`; 34,914 B (34.10 KiB) encoded; 134.002041x post-compression ratio; 637.305942% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 220,197 B (215.04 KiB) compressed - `plain`; 4,003,714 B (3.82 MiB) encoded; 18.188499x post-compression ratio; 0.076749% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## Age (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `6 / 6 / 6`
- Page cardinality per row group min/median/max of mins: `6 / 6 / 6`; of maxes: `6 / 6 / 6`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/age_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/age_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/age_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/age_value_length.svg)


Compressed overall:
1. 117,935 B (115.17 KiB) compressed - `zstd-3` + `rle-dict`; 186,008 B (181.65 KiB) encoded; 33.959825x post-compression ratio; 20.934413% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`
2. 142,526 B (139.19 KiB) compressed - `zstd-3` + `plain`; 4,003,588 B (3.82 MiB) encoded; 28.100501x post-compression ratio; 0.068759% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`
3. 155,906 B (152.25 KiB) compressed - `snappy` + `rle-dict`; 186,344 B (181.98 KiB) encoded; 25.688889x post-compression ratio; 114.223955% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
4. 333,926 B (326.10 KiB) compressed - `snappy` + `plain`; 4,003,714 B (3.82 MiB) encoded; 11.993831x post-compression ratio; 0.018567% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 117,935 B (115.17 KiB) compressed - `rle-dict`; 186,008 B (181.65 KiB) encoded; 33.959825x post-compression ratio; 20.934413% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`
2. 142,526 B (139.19 KiB) compressed - `plain`; 4,003,588 B (3.82 MiB) encoded; 28.100501x post-compression ratio; 0.068759% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`

Snappy:
1. 155,906 B (152.25 KiB) compressed - `rle-dict`; 186,344 B (181.98 KiB) encoded; 25.688889x post-compression ratio; 114.223955% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
2. 333,926 B (326.10 KiB) compressed - `plain`; 4,003,714 B (3.82 MiB) encoded; 11.993831x post-compression ratio; 0.018567% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## BrowserCountry (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `3 / 5 / 11`
- Page cardinality per row group min/median/max of mins: `3 / 5 / 11`; of maxes: `3 / 5 / 11`
- Value length per row group min/median/max of mins: `2 / 2 / 2`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/browsercountry_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/browsercountry_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/browsercountry_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/browsercountry_value_length.svg)


Compressed overall:
1. 69,316 B (67.69 KiB) compressed - `zstd-3` + `rle-dict`; 136,716 B (133.51 KiB) encoded; 105.745831x post-compression ratio; 76.325812% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 94,809 B (92.59 KiB) compressed - `snappy` + `rle-dict`; 137,260 B (134.04 KiB) encoded; 77.312048x post-compression ratio; 382.214769% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
3. 122,185 B (119.32 KiB) compressed - `zstd-3` + `plain`; 7,328,504 B (6.99 MiB) encoded; 59.989999x post-compression ratio; 0.030282% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-rle-dict`
4. 185,144 B (180.80 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 3,638,794 B (3.47 MiB) encoded; 39.590146x post-compression ratio; -33.985438% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
5. 194,675 B (190.11 KiB) compressed - `zstd-3` + `delta-byte-array`; 960,047 B (937.55 KiB) encoded; 37.651871x post-compression ratio; -37.217414% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`

ZSTD:
1. 69,316 B (67.69 KiB) compressed - `rle-dict`; 136,716 B (133.51 KiB) encoded; 105.745831x post-compression ratio; 76.325812% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 122,185 B (119.32 KiB) compressed - `plain`; 7,328,504 B (6.99 MiB) encoded; 59.989999x post-compression ratio; 0.030282% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-rle-dict`
3. 185,144 B (180.80 KiB) compressed - `delta-length-byte-array`; 3,638,794 B (3.47 MiB) encoded; 39.590146x post-compression ratio; -33.985438% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
4. 194,675 B (190.11 KiB) compressed - `delta-byte-array`; 960,047 B (937.55 KiB) encoded; 37.651871x post-compression ratio; -37.217414% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`

Snappy:
1. 94,809 B (92.59 KiB) compressed - `rle-dict`; 137,260 B (134.04 KiB) encoded; 77.312048x post-compression ratio; 382.214769% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 253,897 B (247.95 KiB) compressed - `delta-byte-array`; 963,796 B (941.21 KiB) encoded; 28.869494x post-compression ratio; 80.066326% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
3. 352,235 B (343.98 KiB) compressed - `delta-length-byte-array`; 3,638,826 B (3.47 MiB) encoded; 20.809624x post-compression ratio; 29.794881% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
4. 457,092 B (446.38 KiB) compressed - `plain`; 7,328,708 B (6.99 MiB) encoded; 16.035892x post-compression ratio; 0.019908% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`

## BrowserLanguage (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `3 / 7 / 14`
- Page cardinality per row group min/median/max of mins: `3 / 7 / 14`; of maxes: `3 / 7 / 14`
- Value length per row group min/median/max of mins: `2 / 2 / 2`; of maxes: `2 / 3 / 3`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/browserlanguage_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/browserlanguage_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/browserlanguage_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/browserlanguage_value_length.svg)


Compressed overall:
1. 27,489 B (26.84 KiB) compressed - `zstd-3` + `rle-dict`; 42,266 B (41.28 KiB) encoded; 218.478701x post-compression ratio; 17.723453% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 31,938 B (31.19 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 2,051,163 B (1.96 MiB) encoded; 188.044367x post-compression ratio; 1.324441% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
3. 32,234 B (31.48 KiB) compressed - `zstd-3` + `plain`; 6,004,437 B (5.73 MiB) encoded; 186.317584x post-compression ratio; 0.393994% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-rle-dict`
4. 34,489 B (33.68 KiB) compressed - `snappy` + `rle-dict`; 42,617 B (41.62 KiB) encoded; 174.135550x post-compression ratio; 831.639073% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
5. 65,004 B (63.48 KiB) compressed - `zstd-3` + `delta-byte-array`; 331,509 B (323.74 KiB) encoded; 92.390637x post-compression ratio; -50.216910% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 27,489 B (26.84 KiB) compressed - `rle-dict`; 42,266 B (41.28 KiB) encoded; 218.478701x post-compression ratio; 17.723453% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 31,938 B (31.19 KiB) compressed - `delta-length-byte-array`; 2,051,163 B (1.96 MiB) encoded; 188.044367x post-compression ratio; 1.324441% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
3. 32,234 B (31.48 KiB) compressed - `plain`; 6,004,437 B (5.73 MiB) encoded; 186.317584x post-compression ratio; 0.393994% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-rle-dict`
4. 65,004 B (63.48 KiB) compressed - `delta-byte-array`; 331,509 B (323.74 KiB) encoded; 92.390637x post-compression ratio; -50.216910% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`

Snappy:
1. 34,489 B (33.68 KiB) compressed - `rle-dict`; 42,617 B (41.62 KiB) encoded; 174.135550x post-compression ratio; 831.639073% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 89,626 B (87.53 KiB) compressed - `delta-byte-array`; 333,501 B (325.68 KiB) encoded; 67.009138x post-compression ratio; 258.504229% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
3. 125,361 B (122.42 KiB) compressed - `delta-length-byte-array`; 2,051,346 B (1.96 MiB) encoded; 47.907730x post-compression ratio; 156.310176% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-rle-dict`
4. 321,265 B (313.74 KiB) compressed - `plain`; 6,004,609 B (5.73 MiB) encoded; 18.694103x post-compression ratio; 0.014941% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`

## CLID (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 2`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 2`; of maxes: `1 / 1 / 2`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/clid_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/clid_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/clid_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/clid_value_length.svg)


Compressed overall:
1. 5,641 B (5.51 KiB) compressed - `zstd-3` + `plain`; 4,003,542 B (3.82 MiB) encoded; 709.990073x post-compression ratio; 0.000000% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`
2. 6,322 B (6.17 KiB) compressed - `snappy` + `rle-dict`; 6,085 B (5.94 KiB) encoded; 633.510598x post-compression ratio; 3137.361594% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
3. 6,908 B (6.75 KiB) compressed - `zstd-3` + `rle-dict`; 5,882 B (5.74 KiB) encoded; 579.770411x post-compression ratio; -18.341054% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
4. 204,500 B (199.71 KiB) compressed - `snappy` + `plain`; 4,003,715 B (3.82 MiB) encoded; 19.584616x post-compression ratio; 0.081174% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 5,641 B (5.51 KiB) compressed - `plain`; 4,003,542 B (3.82 MiB) encoded; 709.990073x post-compression ratio; 0.000000% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`
2. 6,908 B (6.75 KiB) compressed - `rle-dict`; 5,882 B (5.74 KiB) encoded; 579.770411x post-compression ratio; -18.341054% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 6,322 B (6.17 KiB) compressed - `rle-dict`; 6,085 B (5.94 KiB) encoded; 633.510598x post-compression ratio; 3137.361594% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 204,500 B (199.71 KiB) compressed - `plain`; 4,003,715 B (3.82 MiB) encoded; 19.584616x post-compression ratio; 0.081174% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## ClientEventTime (timestamp_millis)

Column shape stats:
- Parquet type: `TIMESTAMP(isAdjustedToUTC=true,unit=MILLIS)`; physical type: `INT64`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `5,666 / 5,882 / 13,078`
- Page cardinality per row group min/median/max of mins: `5,666 / 5,882 / 13,078`; of maxes: `5,666 / 5,882 / 13,078`
- Value length per row group min/median/max of mins: `8 / 8 / 8`; of maxes: `8 / 8 / 8`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/clienteventtime_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/clienteventtime_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/clienteventtime_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/clienteventtime_value_length.svg)


Compressed overall:
1. 2,474,364 B (2.36 MiB) compressed - `zstd-3` + `plain`; 8,004,715 B (7.63 MiB) encoded; 3.235706x post-compression ratio; 0.056095% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain`
2. 3,958,277 B (3.77 MiB) compressed - `zstd-3` + `rle-dict`; 7,210,522 B (6.88 MiB) encoded; 2.022677x post-compression ratio; -37.453796% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict`
3. 4,228,484 B (4.03 MiB) compressed - `snappy` + `plain`; 8,004,794 B (7.63 MiB) encoded; 1.893424x post-compression ratio; 0.015987% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-rle-dict-ts-plain`
4. 5,441,369 B (5.19 MiB) compressed - `snappy` + `rle-dict`; 7,218,475 B (6.88 MiB) encoded; 1.471379x post-compression ratio; -22.277647% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-rle-dict`

ZSTD:
1. 2,474,364 B (2.36 MiB) compressed - `plain`; 8,004,715 B (7.63 MiB) encoded; 3.235706x post-compression ratio; 0.056095% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain`
2. 3,958,277 B (3.77 MiB) compressed - `rle-dict`; 7,210,522 B (6.88 MiB) encoded; 2.022677x post-compression ratio; -37.453796% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict`

Snappy:
1. 4,228,484 B (4.03 MiB) compressed - `plain`; 8,004,794 B (7.63 MiB) encoded; 1.893424x post-compression ratio; 0.015987% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-rle-dict-ts-plain`
2. 5,441,369 B (5.19 MiB) compressed - `rle-dict`; 7,218,475 B (6.88 MiB) encoded; 1.471379x post-compression ratio; -22.277647% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-rle-dict`

## ClientIP (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `744 / 924 / 1,957`
- Page cardinality per row group min/median/max of mins: `744 / 924 / 1,957`; of maxes: `744 / 924 / 1,957`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/clientip_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/clientip_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/clientip_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/clientip_value_length.svg)


Compressed overall:
1. 408,058 B (398.49 KiB) compressed - `zstd-3` + `plain`; 4,003,592 B (3.82 MiB) encoded; 9.814904x post-compression ratio; 0.002696% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
2. 719,081 B (702.23 KiB) compressed - `snappy` + `plain`; 4,003,770 B (3.82 MiB) encoded; 5.569679x post-compression ratio; 0.055627% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
3. 806,924 B (788.01 KiB) compressed - `zstd-3` + `rle-dict`; 941,748 B (919.68 KiB) encoded; 4.963355x post-compression ratio; -49.429066% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
4. 828,313 B (808.90 KiB) compressed - `snappy` + `rle-dict`; 938,265 B (916.27 KiB) encoded; 4.835189x post-compression ratio; -13.138995% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`

ZSTD:
1. 408,058 B (398.49 KiB) compressed - `plain`; 4,003,592 B (3.82 MiB) encoded; 9.814904x post-compression ratio; 0.002696% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
2. 806,924 B (788.01 KiB) compressed - `rle-dict`; 941,748 B (919.68 KiB) encoded; 4.963355x post-compression ratio; -49.429066% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 719,081 B (702.23 KiB) compressed - `plain`; 4,003,770 B (3.82 MiB) encoded; 5.569679x post-compression ratio; 0.055627% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
2. 828,313 B (808.90 KiB) compressed - `rle-dict`; 938,265 B (916.27 KiB) encoded; 4.835189x post-compression ratio; -13.138995% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`

## ClientTimeZone (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `10 / 15 / 20`
- Page cardinality per row group min/median/max of mins: `10 / 15 / 20`; of maxes: `10 / 15 / 20`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/clienttimezone_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/clienttimezone_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/clienttimezone_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/clienttimezone_value_length.svg)


Compressed overall:
1. 84,413 B (82.43 KiB) compressed - `zstd-3` + `rle-dict`; 164,301 B (160.45 KiB) encoded; 47.445903x post-compression ratio; 18.375132% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
2. 99,843 B (97.50 KiB) compressed - `zstd-3` + `plain`; 4,003,590 B (3.82 MiB) encoded; 40.113488x post-compression ratio; 0.081127% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-rle-dict`
3. 110,114 B (107.53 KiB) compressed - `snappy` + `rle-dict`; 162,472 B (158.66 KiB) encoded; 36.371860x post-compression ratio; 159.562817% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
4. 285,543 B (278.85 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 14.026087x post-compression ratio; 0.095257% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 84,413 B (82.43 KiB) compressed - `rle-dict`; 164,301 B (160.45 KiB) encoded; 47.445903x post-compression ratio; 18.375132% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
2. 99,843 B (97.50 KiB) compressed - `plain`; 4,003,590 B (3.82 MiB) encoded; 40.113488x post-compression ratio; 0.081127% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-rle-dict`

Snappy:
1. 110,114 B (107.53 KiB) compressed - `rle-dict`; 162,472 B (158.66 KiB) encoded; 36.371860x post-compression ratio; 159.562817% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 285,543 B (278.85 KiB) compressed - `plain`; 4,003,713 B (3.82 MiB) encoded; 14.026087x post-compression ratio; 0.095257% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## CodeVersion (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 2 / 3`
- Page cardinality per row group min/median/max of mins: `1 / 2 / 3`; of maxes: `1 / 2 / 3`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/codeversion_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/codeversion_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/codeversion_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/codeversion_value_length.svg)


Compressed overall:
1. 7,031 B (6.87 KiB) compressed - `zstd-3` + `plain`; 4,003,547 B (3.82 MiB) encoded; 569.627791x post-compression ratio; 0.469350% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
2. 7,360 B (7.19 KiB) compressed - `snappy` + `rle-dict`; 7,136 B (6.97 KiB) encoded; 544.164810x post-compression ratio; 2688.247283% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-rle-dict`
3. 7,970 B (7.78 KiB) compressed - `zstd-3` + `rle-dict`; 6,944 B (6.78 KiB) encoded; 502.516060x post-compression ratio; -11.367629% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict`
4. 205,144 B (200.34 KiB) compressed - `snappy` + `plain`; 4,003,776 B (3.82 MiB) encoded; 19.523130x post-compression ratio; 0.034610% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`

ZSTD:
1. 7,031 B (6.87 KiB) compressed - `plain`; 4,003,547 B (3.82 MiB) encoded; 569.627791x post-compression ratio; 0.469350% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
2. 7,970 B (7.78 KiB) compressed - `rle-dict`; 6,944 B (6.78 KiB) encoded; 502.516060x post-compression ratio; -11.367629% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict`

Snappy:
1. 7,360 B (7.19 KiB) compressed - `rle-dict`; 7,136 B (6.97 KiB) encoded; 544.164810x post-compression ratio; 2688.247283% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-rle-dict`
2. 205,144 B (200.34 KiB) compressed - `plain`; 4,003,776 B (3.82 MiB) encoded; 19.523130x post-compression ratio; 0.034610% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`

## ConnectTiming (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `117 / 222 / 628`
- Page cardinality per row group min/median/max of mins: `117 / 222 / 628`; of maxes: `117 / 222 / 628`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/connecttiming_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/connecttiming_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/connecttiming_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/connecttiming_value_length.svg)


Compressed overall:
1. 333,011 B (325.21 KiB) compressed - `zstd-3` + `plain`; 4,003,660 B (3.82 MiB) encoded; 12.026783x post-compression ratio; 0.090988% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict`
2. 377,469 B (368.62 KiB) compressed - `zstd-3` + `rle-dict`; 704,170 B (687.67 KiB) encoded; 10.610278x post-compression ratio; -11.697649% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
3. 462,272 B (451.44 KiB) compressed - `snappy` + `rle-dict`; 700,120 B (683.71 KiB) encoded; 8.663841x post-compression ratio; 22.204243% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 564,302 B (551.08 KiB) compressed - `snappy` + `plain`; 4,003,750 B (3.82 MiB) encoded; 7.097354x post-compression ratio; 0.108807% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 333,011 B (325.21 KiB) compressed - `plain`; 4,003,660 B (3.82 MiB) encoded; 12.026783x post-compression ratio; 0.090988% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict`
2. 377,469 B (368.62 KiB) compressed - `rle-dict`; 704,170 B (687.67 KiB) encoded; 10.610278x post-compression ratio; -11.697649% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`

Snappy:
1. 462,272 B (451.44 KiB) compressed - `rle-dict`; 700,120 B (683.71 KiB) encoded; 8.663841x post-compression ratio; 22.204243% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
2. 564,302 B (551.08 KiB) compressed - `plain`; 4,003,750 B (3.82 MiB) encoded; 7.097354x post-compression ratio; 0.108807% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## CookieEnable (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 2 / 2`
- Page cardinality per row group min/median/max of mins: `1 / 2 / 2`; of maxes: `1 / 2 / 2`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/cookieenable_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/cookieenable_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/cookieenable_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/cookieenable_value_length.svg)


Compressed overall:
1. 6,053 B (5.91 KiB) compressed - `zstd-3` + `plain`; 4,003,539 B (3.82 MiB) encoded; 661.663968x post-compression ratio; 0.677350% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-rle-dict`
2. 6,490 B (6.34 KiB) compressed - `snappy` + `rle-dict`; 6,251 B (6.10 KiB) encoded; 617.111248x post-compression ratio; 3056.949153% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-rle-dict`
3. 7,077 B (6.91 KiB) compressed - `zstd-3` + `rle-dict`; 6,049 B (5.91 KiB) encoded; 565.925110x post-compression ratio; -13.890066% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
4. 204,770 B (199.97 KiB) compressed - `snappy` + `plain`; 4,003,777 B (3.82 MiB) encoded; 19.558783x post-compression ratio; 0.056649% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`

ZSTD:
1. 6,053 B (5.91 KiB) compressed - `plain`; 4,003,539 B (3.82 MiB) encoded; 661.663968x post-compression ratio; 0.677350% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-rle-dict`
2. 7,077 B (6.91 KiB) compressed - `rle-dict`; 6,049 B (5.91 KiB) encoded; 565.925110x post-compression ratio; -13.890066% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`

Snappy:
1. 6,490 B (6.34 KiB) compressed - `rle-dict`; 6,251 B (6.10 KiB) encoded; 617.111248x post-compression ratio; 3056.949153% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-rle-dict`
2. 204,770 B (199.97 KiB) compressed - `plain`; 4,003,777 B (3.82 MiB) encoded; 19.558783x post-compression ratio; 0.056649% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`

## CounterClass (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/counterclass_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/counterclass_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/counterclass_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/counterclass_value_length.svg)


Compressed overall:
1. 4,228 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; 0.283822% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 5,360 B (5.23 KiB) compressed - `snappy` + `rle-dict`; 5,124 B (5.00 KiB) encoded; 747.211567x post-compression ratio; 3709.981343% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
3. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; -29.037657% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
4. 204,063 B (199.28 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 19.626557x post-compression ratio; 0.074487% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 4,228 B (4.13 KiB) compressed - `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; 0.283822% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 5,975 B (5.83 KiB) compressed - `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; -29.037657% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 5,360 B (5.23 KiB) compressed - `rle-dict`; 5,124 B (5.00 KiB) encoded; 747.211567x post-compression ratio; 3709.981343% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 204,063 B (199.28 KiB) compressed - `plain`; 4,003,713 B (3.82 MiB) encoded; 19.626557x post-compression ratio; 0.074487% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## CounterID (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `true`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 4`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 4`; of maxes: `1 / 1 / 4`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/counterid_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/counterid_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/counterid_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/counterid_value_length.svg)


Compressed overall:
1. 4,931 B (4.82 KiB) compressed - `zstd-3` + `plain`; 4,003,528 B (3.82 MiB) encoded; 812.219023x post-compression ratio; 0.202799% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 5,383 B (5.26 KiB) compressed - `snappy` + `rle-dict`; 5,147 B (5.03 KiB) encoded; 744.018577x post-compression ratio; 3697.120565% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
3. 6,020 B (5.88 KiB) compressed - `zstd-3` + `rle-dict`; 4,994 B (4.88 KiB) encoded; 665.291030x post-compression ratio; -17.923588% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
4. 204,260 B (199.47 KiB) compressed - `snappy` + `plain`; 4,003,717 B (3.82 MiB) encoded; 19.607618x post-compression ratio; 0.068051% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 4,931 B (4.82 KiB) compressed - `plain`; 4,003,528 B (3.82 MiB) encoded; 812.219023x post-compression ratio; 0.202799% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 6,020 B (5.88 KiB) compressed - `rle-dict`; 4,994 B (4.88 KiB) encoded; 665.291030x post-compression ratio; -17.923588% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 5,383 B (5.26 KiB) compressed - `rle-dict`; 5,147 B (5.03 KiB) encoded; 744.018577x post-compression ratio; 3697.120565% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 204,260 B (199.47 KiB) compressed - `plain`; 4,003,717 B (3.82 MiB) encoded; 19.607618x post-compression ratio; 0.068051% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## DNSTiming (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `68 / 98 / 352`
- Page cardinality per row group min/median/max of mins: `68 / 98 / 352`; of maxes: `68 / 98 / 352`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/dnstiming_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/dnstiming_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/dnstiming_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/dnstiming_value_length.svg)


Compressed overall:
1. 135,244 B (132.07 KiB) compressed - `zstd-3` + `plain`; 4,003,645 B (3.82 MiB) encoded; 29.613521x post-compression ratio; 0.236609% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict`
2. 156,435 B (152.77 KiB) compressed - `zstd-3` + `rle-dict`; 343,545 B (335.49 KiB) encoded; 25.602014x post-compression ratio; -13.341643% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
3. 193,951 B (189.41 KiB) compressed - `snappy` + `rle-dict`; 344,174 B (336.11 KiB) encoded; 20.649808x post-compression ratio; 68.713747% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 326,994 B (319.33 KiB) compressed - `snappy` + `plain`; 4,003,716 B (3.82 MiB) encoded; 12.248087x post-compression ratio; 0.069726% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 135,244 B (132.07 KiB) compressed - `plain`; 4,003,645 B (3.82 MiB) encoded; 29.613521x post-compression ratio; 0.236609% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict`
2. 156,435 B (152.77 KiB) compressed - `rle-dict`; 343,545 B (335.49 KiB) encoded; 25.602014x post-compression ratio; -13.341643% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`

Snappy:
1. 193,951 B (189.41 KiB) compressed - `rle-dict`; 344,174 B (336.11 KiB) encoded; 20.649808x post-compression ratio; 68.713747% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
2. 326,994 B (319.33 KiB) compressed - `plain`; 4,003,716 B (3.82 MiB) encoded; 12.248087x post-compression ratio; 0.069726% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## DontCountHits (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `2 / 2 / 2`
- Page cardinality per row group min/median/max of mins: `2 / 2 / 2`; of maxes: `2 / 2 / 2`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/dontcounthits_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/dontcounthits_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/dontcounthits_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/dontcounthits_value_length.svg)


Compressed overall:
1. 49,438 B (48.28 KiB) compressed - `zstd-3` + `rle-dict`; 65,952 B (64.41 KiB) encoded; 81.011550x post-compression ratio; 66.993001% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
2. 59,571 B (58.17 KiB) compressed - `snappy` + `rle-dict`; 66,011 B (64.46 KiB) encoded; 67.231522x post-compression ratio; 396.395897% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
3. 82,484 B (80.55 KiB) compressed - `zstd-3` + `plain`; 4,003,585 B (3.82 MiB) encoded; 48.555465x post-compression ratio; 0.089714% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
4. 295,527 B (288.60 KiB) compressed - `snappy` + `plain`; 4,003,719 B (3.82 MiB) encoded; 13.552227x post-compression ratio; 0.061247% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 49,438 B (48.28 KiB) compressed - `rle-dict`; 65,952 B (64.41 KiB) encoded; 81.011550x post-compression ratio; 66.993001% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
2. 82,484 B (80.55 KiB) compressed - `plain`; 4,003,585 B (3.82 MiB) encoded; 48.555465x post-compression ratio; 0.089714% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`

Snappy:
1. 59,571 B (58.17 KiB) compressed - `rle-dict`; 66,011 B (64.46 KiB) encoded; 67.231522x post-compression ratio; 396.395897% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
2. 295,527 B (288.60 KiB) compressed - `plain`; 4,003,719 B (3.82 MiB) encoded; 13.552227x post-compression ratio; 0.061247% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## EventDate (date)

Column shape stats:
- Parquet type: `DATE`; physical type: `INT32`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/eventdate_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/eventdate_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/eventdate_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/eventdate_value_length.svg)


Compressed overall:
1. 4,913 B (4.80 KiB) compressed - `zstd-3` + `plain`; 4,003,526 B (3.82 MiB) encoded; 815.193365x post-compression ratio; 0.203542% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict`
2. 5,360 B (5.23 KiB) compressed - `snappy` + `rle-dict`; 5,124 B (5.00 KiB) encoded; 747.209888x post-compression ratio; 3712.985075% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
3. 5,976 B (5.84 KiB) compressed - `zstd-3` + `rle-dict`; 4,950 B (4.83 KiB) encoded; 670.188253x post-compression ratio; -17.620482% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
4. 204,233 B (199.45 KiB) compressed - `snappy` + `plain`; 4,003,709 B (3.82 MiB) encoded; 19.610176x post-compression ratio; 0.070018% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 4,913 B (4.80 KiB) compressed - `plain`; 4,003,526 B (3.82 MiB) encoded; 815.193365x post-compression ratio; 0.203542% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict`
2. 5,976 B (5.84 KiB) compressed - `rle-dict`; 4,950 B (4.83 KiB) encoded; 670.188253x post-compression ratio; -17.620482% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

Snappy:
1. 5,360 B (5.23 KiB) compressed - `rle-dict`; 5,124 B (5.00 KiB) encoded; 747.209888x post-compression ratio; 3712.985075% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 204,233 B (199.45 KiB) compressed - `plain`; 4,003,709 B (3.82 MiB) encoded; 19.610176x post-compression ratio; 0.070018% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`

## EventTime (timestamp_millis)

Column shape stats:
- Parquet type: `TIMESTAMP(isAdjustedToUTC=true,unit=MILLIS)`; physical type: `INT64`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `5,997 / 6,237 / 13,042`
- Page cardinality per row group min/median/max of mins: `5,997 / 6,237 / 13,042`; of maxes: `5,997 / 6,237 / 13,042`
- Value length per row group min/median/max of mins: `8 / 8 / 8`; of maxes: `8 / 8 / 8`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/eventtime_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/eventtime_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/eventtime_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/eventtime_value_length.svg)


Compressed overall:
1. 2,516,679 B (2.40 MiB) compressed - `zstd-3` + `plain`; 8,004,558 B (7.63 MiB) encoded; 3.181302x post-compression ratio; 0.048357% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 4,021,616 B (3.84 MiB) compressed - `zstd-3` + `rle-dict`; 7,360,398 B (7.02 MiB) encoded; 1.990820x post-compression ratio; -37.390939% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-rle-dict`
3. 4,282,415 B (4.08 MiB) compressed - `snappy` + `plain`; 8,004,714 B (7.63 MiB) encoded; 1.869579x post-compression ratio; 0.003993% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
4. 5,516,487 B (5.26 MiB) compressed - `snappy` + `rle-dict`; 7,344,624 B (7.00 MiB) encoded; 1.451343x post-compression ratio; -22.367514% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict`

ZSTD:
1. 2,516,679 B (2.40 MiB) compressed - `plain`; 8,004,558 B (7.63 MiB) encoded; 3.181302x post-compression ratio; 0.048357% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 4,021,616 B (3.84 MiB) compressed - `rle-dict`; 7,360,398 B (7.02 MiB) encoded; 1.990820x post-compression ratio; -37.390939% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-rle-dict`

Snappy:
1. 4,282,415 B (4.08 MiB) compressed - `plain`; 8,004,714 B (7.63 MiB) encoded; 1.869579x post-compression ratio; 0.003993% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
2. 5,516,487 B (5.26 MiB) compressed - `rle-dict`; 7,344,624 B (7.00 MiB) encoded; 1.451343x post-compression ratio; -22.367514% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict`

## FUniqID (int64)

Column shape stats:
- Parquet type: `INT(64,true)`; physical type: `INT64`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `628 / 812 / 1,703`
- Page cardinality per row group min/median/max of mins: `628 / 812 / 1,703`; of maxes: `628 / 812 / 1,703`
- Value length per row group min/median/max of mins: `8 / 8 / 8`; of maxes: `8 / 8 / 8`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/funiqid_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/funiqid_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/funiqid_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/funiqid_value_length.svg)


Compressed overall:
1. 693,965 B (677.70 KiB) compressed - `zstd-3` + `plain`; 8,004,558 B (7.63 MiB) encoded; 11.537062x post-compression ratio; 0.009799% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-rle-dict`
2. 1,058,950 B (1.01 MiB) compressed - `zstd-3` + `rle-dict`; 1,187,745 B (1.13 MiB) encoded; 7.560619x post-compression ratio; -34.460267% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 1,071,313 B (1.02 MiB) compressed - `snappy` + `rle-dict`; 1,179,670 B (1.13 MiB) encoded; 7.473369x post-compression ratio; 7.494635% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`
4. 1,151,403 B (1.10 MiB) compressed - `snappy` + `plain`; 8,004,714 B (7.63 MiB) encoded; 6.953531x post-compression ratio; 0.017457% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 693,965 B (677.70 KiB) compressed - `plain`; 8,004,558 B (7.63 MiB) encoded; 11.537062x post-compression ratio; 0.009799% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-rle-dict`
2. 1,058,950 B (1.01 MiB) compressed - `rle-dict`; 1,187,745 B (1.13 MiB) encoded; 7.560619x post-compression ratio; -34.460267% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`

Snappy:
1. 1,071,313 B (1.02 MiB) compressed - `rle-dict`; 1,179,670 B (1.13 MiB) encoded; 7.473369x post-compression ratio; 7.494635% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`
2. 1,151,403 B (1.10 MiB) compressed - `plain`; 8,004,714 B (7.63 MiB) encoded; 6.953531x post-compression ratio; 0.017457% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## FetchTiming (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `329 / 664 / 1,264`
- Page cardinality per row group min/median/max of mins: `329 / 664 / 1,264`; of maxes: `329 / 664 / 1,264`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/fetchtiming_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/fetchtiming_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/fetchtiming_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/fetchtiming_value_length.svg)


Compressed overall:
1. 549,819 B (536.93 KiB) compressed - `zstd-3` + `plain`; 4,003,625 B (3.82 MiB) encoded; 7.284312x post-compression ratio; 0.000000% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`
2. 688,917 B (672.77 KiB) compressed - `zstd-3` + `rle-dict`; 1,133,389 B (1.08 MiB) encoded; 5.813549x post-compression ratio; -20.190821% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
3. 808,624 B (789.67 KiB) compressed - `snappy` + `plain`; 4,003,769 B (3.82 MiB) encoded; 4.952924x post-compression ratio; 0.141104% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
4. 811,674 B (792.65 KiB) compressed - `snappy` + `rle-dict`; 1,127,050 B (1.07 MiB) encoded; 4.934312x post-compression ratio; -0.235193% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-rle-dict`

ZSTD:
1. 549,819 B (536.93 KiB) compressed - `plain`; 4,003,625 B (3.82 MiB) encoded; 7.284312x post-compression ratio; 0.000000% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`
2. 688,917 B (672.77 KiB) compressed - `rle-dict`; 1,133,389 B (1.08 MiB) encoded; 5.813549x post-compression ratio; -20.190821% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`

Snappy:
1. 808,624 B (789.67 KiB) compressed - `plain`; 4,003,769 B (3.82 MiB) encoded; 4.952924x post-compression ratio; 0.141104% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
2. 811,674 B (792.65 KiB) compressed - `rle-dict`; 1,127,050 B (1.07 MiB) encoded; 4.934312x post-compression ratio; -0.235193% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-rle-dict`

## FlashMajor (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `4 / 5 / 7`
- Page cardinality per row group min/median/max of mins: `4 / 5 / 7`; of maxes: `4 / 5 / 7`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/flashmajor_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/flashmajor_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/flashmajor_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/flashmajor_value_length.svg)


Compressed overall:
1. 49,128 B (47.98 KiB) compressed - `zstd-3` + `rle-dict`; 87,797 B (85.74 KiB) encoded; 81.522818x post-compression ratio; 9.110894% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`
2. 53,261 B (52.01 KiB) compressed - `zstd-3` + `plain`; 4,003,585 B (3.82 MiB) encoded; 75.196729x post-compression ratio; 0.643998% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`
3. 68,118 B (66.52 KiB) compressed - `snappy` + `rle-dict`; 88,473 B (86.40 KiB) encoded; 58.795810x post-compression ratio; 260.155906% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 245,221 B (239.47 KiB) compressed - `snappy` + `plain`; 4,003,714 B (3.82 MiB) encoded; 16.332423x post-compression ratio; 0.044857% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 49,128 B (47.98 KiB) compressed - `rle-dict`; 87,797 B (85.74 KiB) encoded; 81.522818x post-compression ratio; 9.110894% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`
2. 53,261 B (52.01 KiB) compressed - `plain`; 4,003,585 B (3.82 MiB) encoded; 75.196729x post-compression ratio; 0.643998% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`

Snappy:
1. 68,118 B (66.52 KiB) compressed - `rle-dict`; 88,473 B (86.40 KiB) encoded; 58.795810x post-compression ratio; 260.155906% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
2. 245,221 B (239.47 KiB) compressed - `plain`; 4,003,714 B (3.82 MiB) encoded; 16.332423x post-compression ratio; 0.044857% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## FlashMinor (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `9 / 9 / 9`
- Page cardinality per row group min/median/max of mins: `9 / 9 / 9`; of maxes: `9 / 9 / 9`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/flashminor_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/flashminor_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/flashminor_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/flashminor_value_length.svg)


Compressed overall:
1. 113,166 B (110.51 KiB) compressed - `zstd-3` + `rle-dict`; 214,760 B (209.73 KiB) encoded; 35.390912x post-compression ratio; 12.017744% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
2. 126,766 B (123.79 KiB) compressed - `zstd-3` + `plain`; 4,003,586 B (3.82 MiB) encoded; 31.594024x post-compression ratio; 0.000000% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`
3. 153,210 B (149.62 KiB) compressed - `snappy` + `rle-dict`; 215,224 B (210.18 KiB) encoded; 26.140905x post-compression ratio; 114.382220% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
4. 328,332 B (320.64 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 12.198165x post-compression ratio; 0.037462% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 113,166 B (110.51 KiB) compressed - `rle-dict`; 214,760 B (209.73 KiB) encoded; 35.390912x post-compression ratio; 12.017744% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
2. 126,766 B (123.79 KiB) compressed - `plain`; 4,003,586 B (3.82 MiB) encoded; 31.594024x post-compression ratio; 0.000000% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`

Snappy:
1. 153,210 B (149.62 KiB) compressed - `rle-dict`; 215,224 B (210.18 KiB) encoded; 26.140905x post-compression ratio; 114.382220% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
2. 328,332 B (320.64 KiB) compressed - `plain`; 4,003,713 B (3.82 MiB) encoded; 12.198165x post-compression ratio; 0.037462% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## FlashMinor2 (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `18 / 24 / 29`
- Page cardinality per row group min/median/max of mins: `18 / 24 / 29`; of maxes: `18 / 24 / 29`
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `8 / 8 / 8`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/flashminor2_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/flashminor2_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/flashminor2_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/flashminor2_value_length.svg)


Compressed overall:
1. 145,400 B (141.99 KiB) compressed - `zstd-3` + `rle-dict`; 266,705 B (260.45 KiB) encoded; 50.613508x post-compression ratio; 69.559147% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 182,446 B (178.17 KiB) compressed - `snappy` + `rle-dict`; 266,476 B (260.23 KiB) encoded; 40.336341x post-compression ratio; 194.237199% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
3. 246,055 B (240.29 KiB) compressed - `zstd-3` + `plain`; 7,357,830 B (7.02 MiB) encoded; 29.908776x post-compression ratio; 0.196704% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict`
4. 266,557 B (260.31 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 3,706,682 B (3.53 MiB) encoded; 27.608369x post-compression ratio; -7.509838% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
5. 325,949 B (318.31 KiB) compressed - `zstd-3` + `delta-byte-array`; 1,042,299 B (1017.87 KiB) encoded; 22.577778x post-compression ratio; -24.362707% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 145,400 B (141.99 KiB) compressed - `rle-dict`; 266,705 B (260.45 KiB) encoded; 50.613508x post-compression ratio; 69.559147% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 246,055 B (240.29 KiB) compressed - `plain`; 7,357,830 B (7.02 MiB) encoded; 29.908776x post-compression ratio; 0.196704% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict`
3. 266,557 B (260.31 KiB) compressed - `delta-length-byte-array`; 3,706,682 B (3.53 MiB) encoded; 27.608369x post-compression ratio; -7.509838% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
4. 325,949 B (318.31 KiB) compressed - `delta-byte-array`; 1,042,299 B (1017.87 KiB) encoded; 22.577778x post-compression ratio; -24.362707% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

Snappy:
1. 182,446 B (178.17 KiB) compressed - `rle-dict`; 266,476 B (260.23 KiB) encoded; 40.336341x post-compression ratio; 194.237199% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 394,599 B (385.35 KiB) compressed - `delta-byte-array`; 1,043,347 B (1018.89 KiB) encoded; 18.649829x post-compression ratio; 36.042920% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
3. 432,351 B (422.22 KiB) compressed - `delta-length-byte-array`; 3,707,874 B (3.54 MiB) encoded; 17.021365x post-compression ratio; 24.163932% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
4. 536,601 B (524.02 KiB) compressed - `plain`; 7,358,039 B (7.02 MiB) encoded; 13.714481x post-compression ratio; 0.041558% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`

## FromTag (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 15`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 15`; of maxes: `1 / 1 / 15`
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 0 / 12`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/fromtag_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/fromtag_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/fromtag_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/fromtag_value_length.svg)


Compressed overall:
1. 22,585 B (22.06 KiB) compressed - `zstd-3` + `rle-dict`; 41,244 B (40.28 KiB) encoded; 179.296967x post-compression ratio; 25.623201% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain`
2. 28,371 B (27.71 KiB) compressed - `zstd-3` + `plain`; 4,048,287 B (3.86 MiB) encoded; 142.731028x post-compression ratio; 0.003525% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain`
3. 28,608 B (27.94 KiB) compressed - `snappy` + `rle-dict`; 41,121 B (40.16 KiB) encoded; 141.548588x post-compression ratio; 674.814737% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
4. 36,191 B (35.34 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 211,430 B (206.47 KiB) encoded; 111.890304x post-compression ratio; -21.604819% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
5. 38,213 B (37.32 KiB) compressed - `zstd-3` + `delta-byte-array`; 260,375 B (254.27 KiB) encoded; 105.969749x post-compression ratio; -25.753016% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`

ZSTD:
1. 22,585 B (22.06 KiB) compressed - `rle-dict`; 41,244 B (40.28 KiB) encoded; 179.296967x post-compression ratio; 25.623201% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain`
2. 28,371 B (27.71 KiB) compressed - `plain`; 4,048,287 B (3.86 MiB) encoded; 142.731028x post-compression ratio; 0.003525% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain`
3. 36,191 B (35.34 KiB) compressed - `delta-length-byte-array`; 211,430 B (206.47 KiB) encoded; 111.890304x post-compression ratio; -21.604819% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
4. 38,213 B (37.32 KiB) compressed - `delta-byte-array`; 260,375 B (254.27 KiB) encoded; 105.969749x post-compression ratio; -25.753016% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`

Snappy:
1. 28,608 B (27.94 KiB) compressed - `rle-dict`; 41,121 B (40.16 KiB) encoded; 141.548588x post-compression ratio; 674.814737% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 46,104 B (45.02 KiB) compressed - `delta-length-byte-array`; 210,743 B (205.80 KiB) encoded; 87.832336x post-compression ratio; 380.780410% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 50,268 B (49.09 KiB) compressed - `delta-byte-array`; 261,732 B (255.60 KiB) encoded; 80.556656x post-compression ratio; 340.954484% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
4. 221,650 B (216.46 KiB) compressed - `plain`; 4,048,461 B (3.86 MiB) encoded; 18.269443x post-compression ratio; 0.004060% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`

## GoodEvent (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/goodevent_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/goodevent_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/goodevent_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/goodevent_value_length.svg)


Compressed overall:
1. 4,916 B (4.80 KiB) compressed - `zstd-3` + `plain`; 4,003,531 B (3.82 MiB) encoded; 814.697518x post-compression ratio; 0.183076% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 5,360 B (5.23 KiB) compressed - `snappy` + `rle-dict`; 5,124 B (5.00 KiB) encoded; 747.211381x post-compression ratio; 3713.041045% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
3. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301757x post-compression ratio; -17.573222% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
4. 204,229 B (199.44 KiB) compressed - `snappy` + `plain`; 4,003,712 B (3.82 MiB) encoded; 19.610599x post-compression ratio; 0.073447% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 4,916 B (4.80 KiB) compressed - `plain`; 4,003,531 B (3.82 MiB) encoded; 814.697518x post-compression ratio; 0.183076% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 5,975 B (5.83 KiB) compressed - `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301757x post-compression ratio; -17.573222% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 5,360 B (5.23 KiB) compressed - `rle-dict`; 5,124 B (5.00 KiB) encoded; 747.211381x post-compression ratio; 3713.041045% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 204,229 B (199.44 KiB) compressed - `plain`; 4,003,712 B (3.82 MiB) encoded; 19.610599x post-compression ratio; 0.073447% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## HID (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `5,818 / 5,965 / 13,281`
- Page cardinality per row group min/median/max of mins: `5,818 / 5,965 / 13,281`; of maxes: `5,818 / 5,965 / 13,281`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/hid_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/hid_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/hid_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/hid_value_length.svg)


Compressed overall:
1. 3,688,155 B (3.52 MiB) compressed - `snappy` + `plain`; 4,003,834 B (3.82 MiB) encoded; 1.085923x post-compression ratio; 0.001979% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`
2. 3,804,388 B (3.63 MiB) compressed - `zstd-3` + `plain`; 4,003,646 B (3.82 MiB) encoded; 1.052745x post-compression ratio; 1.193543% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`
3. 4,468,486 B (4.26 MiB) compressed - `snappy` + `rle-dict`; 4,467,753 B (4.26 MiB) encoded; 0.896288x post-compression ratio; -17.461350% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-rle-dict`
4. 4,492,493 B (4.28 MiB) compressed - `zstd-3` + `rle-dict`; 4,491,293 B (4.28 MiB) encoded; 0.891499x post-compression ratio; -14.306043% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-rle-dict`

ZSTD:
1. 3,804,388 B (3.63 MiB) compressed - `plain`; 4,003,646 B (3.82 MiB) encoded; 1.052745x post-compression ratio; 1.193543% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`
2. 4,492,493 B (4.28 MiB) compressed - `rle-dict`; 4,491,293 B (4.28 MiB) encoded; 0.891499x post-compression ratio; -14.306043% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-rle-dict`

Snappy:
1. 3,688,155 B (3.52 MiB) compressed - `plain`; 4,003,834 B (3.82 MiB) encoded; 1.085923x post-compression ratio; 0.001979% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`
2. 4,468,486 B (4.26 MiB) compressed - `rle-dict`; 4,467,753 B (4.26 MiB) encoded; 0.896288x post-compression ratio; -17.461350% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-rle-dict`

## HTTPError (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/httperror_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/httperror_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/httperror_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/httperror_value_length.svg)


Compressed overall:
1. 4,228 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; 0.283822% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 5,360 B (5.23 KiB) compressed - `snappy` + `rle-dict`; 5,124 B (5.00 KiB) encoded; 747.211567x post-compression ratio; 3709.981343% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
3. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; -29.037657% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
4. 204,063 B (199.28 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 19.626557x post-compression ratio; 0.074487% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 4,228 B (4.13 KiB) compressed - `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; 0.283822% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 5,975 B (5.83 KiB) compressed - `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; -29.037657% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 5,360 B (5.23 KiB) compressed - `rle-dict`; 5,124 B (5.00 KiB) encoded; 747.211567x post-compression ratio; 3709.981343% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 204,063 B (199.28 KiB) compressed - `plain`; 4,003,713 B (3.82 MiB) encoded; 19.626557x post-compression ratio; 0.074487% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## HasGCLID (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 2 / 2`
- Page cardinality per row group min/median/max of mins: `1 / 2 / 2`; of maxes: `1 / 2 / 2`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/hasgclid_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/hasgclid_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/hasgclid_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/hasgclid_value_length.svg)


Compressed overall:
1. 16,205 B (15.83 KiB) compressed - `zstd-3` + `rle-dict`; 21,326 B (20.83 KiB) encoded; 247.149151x post-compression ratio; 30.737427% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
2. 19,819 B (19.35 KiB) compressed - `snappy` + `rle-dict`; 21,557 B (21.05 KiB) encoded; 202.081437x post-compression ratio; 977.809173% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-rle-dict`
3. 21,114 B (20.62 KiB) compressed - `zstd-3` + `plain`; 4,003,558 B (3.82 MiB) encoded; 189.687032x post-compression ratio; 0.341006% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
4. 213,441 B (208.44 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 18.764211x post-compression ratio; 0.079647% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 16,205 B (15.83 KiB) compressed - `rle-dict`; 21,326 B (20.83 KiB) encoded; 247.149151x post-compression ratio; 30.737427% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
2. 21,114 B (20.62 KiB) compressed - `plain`; 4,003,558 B (3.82 MiB) encoded; 189.687032x post-compression ratio; 0.341006% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`

Snappy:
1. 19,819 B (19.35 KiB) compressed - `rle-dict`; 21,557 B (21.05 KiB) encoded; 202.081437x post-compression ratio; 977.809173% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-rle-dict`
2. 213,441 B (208.44 KiB) compressed - `plain`; 4,003,713 B (3.82 MiB) encoded; 18.764211x post-compression ratio; 0.079647% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## HistoryLength (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 4 / 87`
- Page cardinality per row group min/median/max of mins: `1 / 4 / 87`; of maxes: `1 / 4 / 87`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/historylength_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/historylength_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/historylength_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/historylength_value_length.svg)


Compressed overall:
1. 51,071 B (49.87 KiB) compressed - `zstd-3` + `rle-dict`; 68,732 B (67.12 KiB) encoded; 78.421296x post-compression ratio; 7.961465% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-rle-dict`
2. 55,094 B (53.80 KiB) compressed - `zstd-3` + `plain`; 4,003,572 B (3.82 MiB) encoded; 72.694921x post-compression ratio; 0.078048% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
3. 58,253 B (56.89 KiB) compressed - `snappy` + `rle-dict`; 69,021 B (67.40 KiB) encoded; 68.752751x post-compression ratio; 343.460423% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
4. 258,297 B (252.24 KiB) compressed - `snappy` + `plain`; 4,003,719 B (3.82 MiB) encoded; 15.505616x post-compression ratio; 0.012389% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 51,071 B (49.87 KiB) compressed - `rle-dict`; 68,732 B (67.12 KiB) encoded; 78.421296x post-compression ratio; 7.961465% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-rle-dict`
2. 55,094 B (53.80 KiB) compressed - `plain`; 4,003,572 B (3.82 MiB) encoded; 72.694921x post-compression ratio; 0.078048% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

Snappy:
1. 58,253 B (56.89 KiB) compressed - `rle-dict`; 69,021 B (67.40 KiB) encoded; 68.752751x post-compression ratio; 343.460423% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 258,297 B (252.24 KiB) compressed - `plain`; 4,003,719 B (3.82 MiB) encoded; 15.505616x post-compression ratio; 0.012389% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## HitColor (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `2 / 3 / 3`
- Page cardinality per row group min/median/max of mins: `2 / 3 / 3`; of maxes: `2 / 3 / 3`
- Value length per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/hitcolor_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/hitcolor_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/hitcolor_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/hitcolor_value_length.svg)


Compressed overall:
1. 24,265 B (23.70 KiB) compressed - `zstd-3` + `rle-dict`; 33,649 B (32.86 KiB) encoded; 206.227323x post-compression ratio; 20.226664% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain`
2. 28,126 B (27.47 KiB) compressed - `zstd-3` + `plain`; 5,002,901 B (4.77 MiB) encoded; 177.917443x post-compression ratio; 3.722534% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 28,214 B (27.55 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 1,043,084 B (1018.64 KiB) encoded; 177.362515x post-compression ratio; 3.399022% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict`
4. 31,891 B (31.14 KiB) compressed - `snappy` + `rle-dict`; 33,597 B (32.81 KiB) encoded; 156.912797x post-compression ratio; 823.185225% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
5. 50,769 B (49.58 KiB) compressed - `zstd-3` + `delta-byte-array`; 207,139 B (202.28 KiB) encoded; 98.566172x post-compression ratio; -42.537769% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 24,265 B (23.70 KiB) compressed - `rle-dict`; 33,649 B (32.86 KiB) encoded; 206.227323x post-compression ratio; 20.226664% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain`
2. 28,126 B (27.47 KiB) compressed - `plain`; 5,002,901 B (4.77 MiB) encoded; 177.917443x post-compression ratio; 3.722534% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 28,214 B (27.55 KiB) compressed - `delta-length-byte-array`; 1,043,084 B (1018.64 KiB) encoded; 177.362515x post-compression ratio; 3.399022% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict`
4. 50,769 B (49.58 KiB) compressed - `delta-byte-array`; 207,139 B (202.28 KiB) encoded; 98.566172x post-compression ratio; -42.537769% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`

Snappy:
1. 31,891 B (31.14 KiB) compressed - `rle-dict`; 33,597 B (32.81 KiB) encoded; 156.912797x post-compression ratio; 823.185225% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 74,324 B (72.58 KiB) compressed - `delta-byte-array`; 207,013 B (202.16 KiB) encoded; 67.328265x post-compression ratio; 296.121038% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`
3. 76,943 B (75.14 KiB) compressed - `delta-length-byte-array`; 1,043,282 B (1018.83 KiB) encoded; 65.036534x post-compression ratio; 282.637797% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
4. 294,288 B (287.39 KiB) compressed - `plain`; 5,003,058 B (4.77 MiB) encoded; 17.004112x post-compression ratio; 0.042475% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`

## IPNetworkID (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `498 / 600 / 1,095`
- Page cardinality per row group min/median/max of mins: `498 / 600 / 1,095`; of maxes: `498 / 600 / 1,095`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/ipnetworkid_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/ipnetworkid_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/ipnetworkid_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/ipnetworkid_value_length.svg)


Compressed overall:
1. 323,719 B (316.13 KiB) compressed - `zstd-3` + `plain`; 4,003,592 B (3.82 MiB) encoded; 12.372002x post-compression ratio; 0.019770% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 560,148 B (547.02 KiB) compressed - `snappy` + `plain`; 4,003,756 B (3.82 MiB) encoded; 7.149989x post-compression ratio; 0.174240% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
3. 619,973 B (605.44 KiB) compressed - `zstd-3` + `rle-dict`; 747,959 B (730.43 KiB) encoded; 6.460043x post-compression ratio; -47.774661% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
4. 638,221 B (623.26 KiB) compressed - `snappy` + `rle-dict`; 749,603 B (732.03 KiB) encoded; 6.275337x post-compression ratio; -12.079985% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 323,719 B (316.13 KiB) compressed - `plain`; 4,003,592 B (3.82 MiB) encoded; 12.372002x post-compression ratio; 0.019770% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 619,973 B (605.44 KiB) compressed - `rle-dict`; 747,959 B (730.43 KiB) encoded; 6.460043x post-compression ratio; -47.774661% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 560,148 B (547.02 KiB) compressed - `plain`; 4,003,756 B (3.82 MiB) encoded; 7.149989x post-compression ratio; 0.174240% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
2. 638,221 B (623.26 KiB) compressed - `rle-dict`; 749,603 B (732.03 KiB) encoded; 6.275337x post-compression ratio; -12.079985% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`

## Income (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `3 / 4 / 4`
- Page cardinality per row group min/median/max of mins: `3 / 4 / 4`; of maxes: `3 / 4 / 4`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/income_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/income_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/income_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/income_value_length.svg)


Compressed overall:
1. 88,577 B (86.50 KiB) compressed - `zstd-3` + `rle-dict`; 142,532 B (139.19 KiB) encoded; 45.215507x post-compression ratio; 38.499836% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
2. 122,679 B (119.80 KiB) compressed - `zstd-3` + `plain`; 4,003,583 B (3.82 MiB) encoded; 32.646614x post-compression ratio; 0.000000% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`
3. 122,841 B (119.96 KiB) compressed - `snappy` + `rle-dict`; 142,510 B (139.17 KiB) encoded; 32.603561x post-compression ratio; 159.101603% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
4. 318,224 B (310.77 KiB) compressed - `snappy` + `plain`; 4,003,711 B (3.82 MiB) encoded; 12.585644x post-compression ratio; 0.018540% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 88,577 B (86.50 KiB) compressed - `rle-dict`; 142,532 B (139.19 KiB) encoded; 45.215507x post-compression ratio; 38.499836% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
2. 122,679 B (119.80 KiB) compressed - `plain`; 4,003,583 B (3.82 MiB) encoded; 32.646614x post-compression ratio; 0.000000% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`

Snappy:
1. 122,841 B (119.96 KiB) compressed - `rle-dict`; 142,510 B (139.17 KiB) encoded; 32.603561x post-compression ratio; 159.101603% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 318,224 B (310.77 KiB) compressed - `plain`; 4,003,711 B (3.82 MiB) encoded; 12.585644x post-compression ratio; 0.018540% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## Interests (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `151 / 217 / 395`
- Page cardinality per row group min/median/max of mins: `151 / 217 / 395`; of maxes: `151 / 217 / 395`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/interests_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/interests_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/interests_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/interests_value_length.svg)


Compressed overall:
1. 193,369 B (188.84 KiB) compressed - `zstd-3` + `plain`; 4,003,589 B (3.82 MiB) encoded; 20.711955x post-compression ratio; 0.078089% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`
2. 310,021 B (302.75 KiB) compressed - `zstd-3` + `rle-dict`; 487,582 B (476.15 KiB) encoded; 12.918641x post-compression ratio; -37.578422% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-rle-dict`
3. 343,969 B (335.91 KiB) compressed - `snappy` + `rle-dict`; 483,872 B (472.53 KiB) encoded; 11.643636x post-compression ratio; 12.124639% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-rle-dict`
4. 385,336 B (376.30 KiB) compressed - `snappy` + `plain`; 4,003,711 B (3.82 MiB) encoded; 10.393656x post-compression ratio; 0.087716% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 193,369 B (188.84 KiB) compressed - `plain`; 4,003,589 B (3.82 MiB) encoded; 20.711955x post-compression ratio; 0.078089% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`
2. 310,021 B (302.75 KiB) compressed - `rle-dict`; 487,582 B (476.15 KiB) encoded; 12.918641x post-compression ratio; -37.578422% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-rle-dict`

Snappy:
1. 343,969 B (335.91 KiB) compressed - `rle-dict`; 483,872 B (472.53 KiB) encoded; 11.643636x post-compression ratio; 12.124639% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-rle-dict`
2. 385,336 B (376.30 KiB) compressed - `plain`; 4,003,711 B (3.82 MiB) encoded; 10.393656x post-compression ratio; 0.087716% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## IsArtifical (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `2 / 2 / 2`
- Page cardinality per row group min/median/max of mins: `2 / 2 / 2`; of maxes: `2 / 2 / 2`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isartifical_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isartifical_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isartifical_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isartifical_value_length.svg)


Compressed overall:
1. 81,445 B (79.54 KiB) compressed - `snappy` + `rle-dict`; 81,087 B (79.19 KiB) encoded; 49.174879x post-compression ratio; 521.116091% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
2. 82,011 B (80.09 KiB) compressed - `zstd-3` + `rle-dict`; 80,952 B (79.05 KiB) encoded; 48.835498x post-compression ratio; 101.179110% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
3. 164,905 B (161.04 KiB) compressed - `zstd-3` + `plain`; 4,003,586 B (3.82 MiB) encoded; 24.287002x post-compression ratio; 0.050938% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`
4. 505,690 B (493.84 KiB) compressed - `snappy` + `plain`; 4,003,747 B (3.82 MiB) encoded; 7.919967x post-compression ratio; 0.035199% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 82,011 B (80.09 KiB) compressed - `rle-dict`; 80,952 B (79.05 KiB) encoded; 48.835498x post-compression ratio; 101.179110% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
2. 164,905 B (161.04 KiB) compressed - `plain`; 4,003,586 B (3.82 MiB) encoded; 24.287002x post-compression ratio; 0.050938% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`

Snappy:
1. 81,445 B (79.54 KiB) compressed - `rle-dict`; 81,087 B (79.19 KiB) encoded; 49.174879x post-compression ratio; 521.116091% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
2. 505,690 B (493.84 KiB) compressed - `plain`; 4,003,747 B (3.82 MiB) encoded; 7.919967x post-compression ratio; 0.035199% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## IsDownload (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 2`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 2`; of maxes: `1 / 1 / 2`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isdownload_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isdownload_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isdownload_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isdownload_value_length.svg)


Compressed overall:
1. 8,030 B (7.84 KiB) compressed - `zstd-3` + `plain`; 4,003,553 B (3.82 MiB) encoded; 498.761519x post-compression ratio; 0.821918% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
2. 8,709 B (8.50 KiB) compressed - `snappy` + `rle-dict`; 8,467 B (8.27 KiB) encoded; 459.875416x post-compression ratio; 2273.372373% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
3. 9,254 B (9.04 KiB) compressed - `zstd-3` + `rle-dict`; 8,328 B (8.13 KiB) encoded; 432.791766x post-compression ratio; -12.513508% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
4. 206,526 B (201.69 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 19.392498x post-compression ratio; 0.082798% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 8,030 B (7.84 KiB) compressed - `plain`; 4,003,553 B (3.82 MiB) encoded; 498.761519x post-compression ratio; 0.821918% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
2. 9,254 B (9.04 KiB) compressed - `rle-dict`; 8,328 B (8.13 KiB) encoded; 432.791766x post-compression ratio; -12.513508% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`

Snappy:
1. 8,709 B (8.50 KiB) compressed - `rle-dict`; 8,467 B (8.27 KiB) encoded; 459.875416x post-compression ratio; 2273.372373% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 206,526 B (201.69 KiB) compressed - `plain`; 4,003,713 B (3.82 MiB) encoded; 19.392498x post-compression ratio; 0.082798% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## IsEvent (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isevent_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isevent_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isevent_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isevent_value_length.svg)


Compressed overall:
1. 4,228 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; 0.283822% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 5,360 B (5.23 KiB) compressed - `snappy` + `rle-dict`; 5,124 B (5.00 KiB) encoded; 747.211567x post-compression ratio; 3709.981343% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
3. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; -29.037657% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
4. 204,063 B (199.28 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 19.626557x post-compression ratio; 0.074487% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 4,228 B (4.13 KiB) compressed - `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; 0.283822% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 5,975 B (5.83 KiB) compressed - `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; -29.037657% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 5,360 B (5.23 KiB) compressed - `rle-dict`; 5,124 B (5.00 KiB) encoded; 747.211567x post-compression ratio; 3709.981343% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 204,063 B (199.28 KiB) compressed - `plain`; 4,003,713 B (3.82 MiB) encoded; 19.626557x post-compression ratio; 0.074487% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## IsLink (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 2`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 2`; of maxes: `1 / 1 / 2`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/islink_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/islink_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/islink_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/islink_value_length.svg)


Compressed overall:
1. 36,415 B (35.56 KiB) compressed - `zstd-3` + `rle-dict`; 54,152 B (52.88 KiB) encoded; 109.983606x post-compression ratio; 55.927502% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
2. 47,279 B (46.17 KiB) compressed - `snappy` + `rle-dict`; 54,371 B (53.10 KiB) encoded; 84.711034x post-compression ratio; 442.634150% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
3. 56,671 B (55.34 KiB) compressed - `zstd-3` + `plain`; 4,003,556 B (3.82 MiB) encoded; 70.672002x post-compression ratio; 0.194103% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
4. 256,374 B (250.37 KiB) compressed - `snappy` + `plain`; 4,003,716 B (3.82 MiB) encoded; 15.621916x post-compression ratio; 0.069430% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 36,415 B (35.56 KiB) compressed - `rle-dict`; 54,152 B (52.88 KiB) encoded; 109.983606x post-compression ratio; 55.927502% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
2. 56,671 B (55.34 KiB) compressed - `plain`; 4,003,556 B (3.82 MiB) encoded; 70.672002x post-compression ratio; 0.194103% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`

Snappy:
1. 47,279 B (46.17 KiB) compressed - `rle-dict`; 54,371 B (53.10 KiB) encoded; 84.711034x post-compression ratio; 442.634150% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
2. 256,374 B (250.37 KiB) compressed - `plain`; 4,003,716 B (3.82 MiB) encoded; 15.621916x post-compression ratio; 0.069430% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## IsMobile (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `2 / 2 / 2`
- Page cardinality per row group min/median/max of mins: `2 / 2 / 2`; of maxes: `2 / 2 / 2`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/ismobile_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/ismobile_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/ismobile_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/ismobile_value_length.svg)


Compressed overall:
1. 24,058 B (23.49 KiB) compressed - `zstd-3` + `rle-dict`; 30,002 B (29.30 KiB) encoded; 166.474645x post-compression ratio; 19.236844% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`
2. 28,288 B (27.62 KiB) compressed - `snappy` + `rle-dict`; 29,981 B (29.28 KiB) encoded; 141.581130x post-compression ratio; 672.451216% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
3. 28,623 B (27.95 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 139.924082x post-compression ratio; 0.220103% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
4. 218,352 B (213.23 KiB) compressed - `snappy` + `plain`; 4,003,715 B (3.82 MiB) encoded; 18.342159x post-compression ratio; 0.072818% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 24,058 B (23.49 KiB) compressed - `rle-dict`; 30,002 B (29.30 KiB) encoded; 166.474645x post-compression ratio; 19.236844% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`
2. 28,623 B (27.95 KiB) compressed - `plain`; 4,003,587 B (3.82 MiB) encoded; 139.924082x post-compression ratio; 0.220103% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`

Snappy:
1. 28,288 B (27.62 KiB) compressed - `rle-dict`; 29,981 B (29.28 KiB) encoded; 141.581130x post-compression ratio; 672.451216% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
2. 218,352 B (213.23 KiB) compressed - `plain`; 4,003,715 B (3.82 MiB) encoded; 18.342159x post-compression ratio; 0.072818% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## IsNotBounce (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 2`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 2`; of maxes: `1 / 1 / 2`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isnotbounce_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isnotbounce_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isnotbounce_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isnotbounce_value_length.svg)


Compressed overall:
1. 16,669 B (16.28 KiB) compressed - `snappy` + `rle-dict`; 16,417 B (16.03 KiB) encoded; 240.269602x post-compression ratio; 1312.466255% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
2. 17,295 B (16.89 KiB) compressed - `zstd-3` + `rle-dict`; 16,263 B (15.88 KiB) encoded; 231.572940x post-compression ratio; 45.666378% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
3. 25,161 B (24.57 KiB) compressed - `zstd-3` + `plain`; 4,003,536 B (3.82 MiB) encoded; 159.177060x post-compression ratio; 0.127181% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
4. 235,271 B (229.76 KiB) compressed - `snappy` + `plain`; 4,003,720 B (3.82 MiB) encoded; 17.023152x post-compression ratio; 0.073532% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 17,295 B (16.89 KiB) compressed - `rle-dict`; 16,263 B (15.88 KiB) encoded; 231.572940x post-compression ratio; 45.666378% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 25,161 B (24.57 KiB) compressed - `plain`; 4,003,536 B (3.82 MiB) encoded; 159.177060x post-compression ratio; 0.127181% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`

Snappy:
1. 16,669 B (16.28 KiB) compressed - `rle-dict`; 16,417 B (16.03 KiB) encoded; 240.269602x post-compression ratio; 1312.466255% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
2. 235,271 B (229.76 KiB) compressed - `plain`; 4,003,720 B (3.82 MiB) encoded; 17.023152x post-compression ratio; 0.073532% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## IsOldCounter (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isoldcounter_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isoldcounter_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isoldcounter_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isoldcounter_value_length.svg)


Compressed overall:
1. 4,228 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; 0.283822% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 5,360 B (5.23 KiB) compressed - `snappy` + `rle-dict`; 5,124 B (5.00 KiB) encoded; 747.211567x post-compression ratio; 3709.981343% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
3. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; -29.037657% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
4. 204,063 B (199.28 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 19.626557x post-compression ratio; 0.074487% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 4,228 B (4.13 KiB) compressed - `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; 0.283822% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 5,975 B (5.83 KiB) compressed - `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; -29.037657% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 5,360 B (5.23 KiB) compressed - `rle-dict`; 5,124 B (5.00 KiB) encoded; 747.211567x post-compression ratio; 3709.981343% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 204,063 B (199.28 KiB) compressed - `plain`; 4,003,713 B (3.82 MiB) encoded; 19.626557x post-compression ratio; 0.074487% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## IsParameter (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isparameter_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isparameter_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isparameter_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isparameter_value_length.svg)


Compressed overall:
1. 4,228 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; 0.283822% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 5,360 B (5.23 KiB) compressed - `snappy` + `rle-dict`; 5,124 B (5.00 KiB) encoded; 747.211567x post-compression ratio; 3709.981343% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
3. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; -29.037657% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
4. 204,063 B (199.28 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 19.626557x post-compression ratio; 0.074487% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 4,228 B (4.13 KiB) compressed - `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; 0.283822% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 5,975 B (5.83 KiB) compressed - `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; -29.037657% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 5,360 B (5.23 KiB) compressed - `rle-dict`; 5,124 B (5.00 KiB) encoded; 747.211567x post-compression ratio; 3709.981343% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 204,063 B (199.28 KiB) compressed - `plain`; 4,003,713 B (3.82 MiB) encoded; 19.626557x post-compression ratio; 0.074487% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## IsRefresh (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `2 / 2 / 2`
- Page cardinality per row group min/median/max of mins: `2 / 2 / 2`; of maxes: `2 / 2 / 2`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isrefresh_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isrefresh_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isrefresh_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/isrefresh_value_length.svg)


Compressed overall:
1. 83,542 B (81.58 KiB) compressed - `zstd-3` + `rle-dict`; 92,593 B (90.42 KiB) encoded; 47.940569x post-compression ratio; 114.192861% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
2. 88,713 B (86.63 KiB) compressed - `snappy` + `rle-dict`; 92,710 B (90.54 KiB) encoded; 45.146157x post-compression ratio; 454.637990% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
3. 178,752 B (174.56 KiB) compressed - `zstd-3` + `plain`; 4,003,585 B (3.82 MiB) encoded; 22.405629x post-compression ratio; 0.105733% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
4. 491,835 B (480.31 KiB) compressed - `snappy` + `plain`; 4,003,744 B (3.82 MiB) encoded; 8.143078x post-compression ratio; 0.040867% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 83,542 B (81.58 KiB) compressed - `rle-dict`; 92,593 B (90.42 KiB) encoded; 47.940569x post-compression ratio; 114.192861% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
2. 178,752 B (174.56 KiB) compressed - `plain`; 4,003,585 B (3.82 MiB) encoded; 22.405629x post-compression ratio; 0.105733% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`

Snappy:
1. 88,713 B (86.63 KiB) compressed - `rle-dict`; 92,710 B (90.54 KiB) encoded; 45.146157x post-compression ratio; 454.637990% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
2. 491,835 B (480.31 KiB) compressed - `plain`; 4,003,744 B (3.82 MiB) encoded; 8.143078x post-compression ratio; 0.040867% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## JavaEnable (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `2 / 2 / 2`
- Page cardinality per row group min/median/max of mins: `2 / 2 / 2`; of maxes: `2 / 2 / 2`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/javaenable_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/javaenable_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/javaenable_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/javaenable_value_length.svg)


Compressed overall:
1. 50,727 B (49.54 KiB) compressed - `zstd-3` + `rle-dict`; 79,757 B (77.89 KiB) encoded; 78.953023x post-compression ratio; 25.463757% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 63,363 B (61.88 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 63.208024x post-compression ratio; 0.443476% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
3. 69,384 B (67.76 KiB) compressed - `snappy` + `rle-dict`; 80,011 B (78.14 KiB) encoded; 57.722962x post-compression ratio; 277.709558% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-rle-dict`
4. 261,967 B (255.83 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 15.288376x post-compression ratio; 0.039318% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 50,727 B (49.54 KiB) compressed - `rle-dict`; 79,757 B (77.89 KiB) encoded; 78.953023x post-compression ratio; 25.463757% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 63,363 B (61.88 KiB) compressed - `plain`; 4,003,587 B (3.82 MiB) encoded; 63.208024x post-compression ratio; 0.443476% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

Snappy:
1. 69,384 B (67.76 KiB) compressed - `rle-dict`; 80,011 B (78.14 KiB) encoded; 57.722962x post-compression ratio; 277.709558% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-rle-dict`
2. 261,967 B (255.83 KiB) compressed - `plain`; 4,003,713 B (3.82 MiB) encoded; 15.288376x post-compression ratio; 0.039318% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## JavascriptEnable (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 2 / 2`
- Page cardinality per row group min/median/max of mins: `1 / 2 / 2`; of maxes: `1 / 2 / 2`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/javascriptenable_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/javascriptenable_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/javascriptenable_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/javascriptenable_value_length.svg)


Compressed overall:
1. 6,485 B (6.33 KiB) compressed - `zstd-3` + `plain`; 4,003,543 B (3.82 MiB) encoded; 617.586893x post-compression ratio; 0.385505% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
2. 6,780 B (6.62 KiB) compressed - `snappy` + `rle-dict`; 6,558 B (6.40 KiB) encoded; 590.715487x post-compression ratio; 2923.185841% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
3. 7,387 B (7.21 KiB) compressed - `zstd-3` + `rle-dict`; 6,361 B (6.21 KiB) encoded; 542.175579x post-compression ratio; -11.872208% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
4. 204,912 B (200.11 KiB) compressed - `snappy` + `plain`; 4,003,774 B (3.82 MiB) encoded; 19.545224x post-compression ratio; 0.029281% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`

ZSTD:
1. 6,485 B (6.33 KiB) compressed - `plain`; 4,003,543 B (3.82 MiB) encoded; 617.586893x post-compression ratio; 0.385505% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
2. 7,387 B (7.21 KiB) compressed - `rle-dict`; 6,361 B (6.21 KiB) encoded; 542.175579x post-compression ratio; -11.872208% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`

Snappy:
1. 6,780 B (6.62 KiB) compressed - `rle-dict`; 6,558 B (6.40 KiB) encoded; 590.715487x post-compression ratio; 2923.185841% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
2. 204,912 B (200.11 KiB) compressed - `plain`; 4,003,774 B (3.82 MiB) encoded; 19.545224x post-compression ratio; 0.029281% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`

## LocalEventTime (timestamp_millis)

Column shape stats:
- Parquet type: `TIMESTAMP(isAdjustedToUTC=true,unit=MILLIS)`; physical type: `INT64`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `5,968 / 6,254 / 13,047`
- Page cardinality per row group min/median/max of mins: `5,968 / 6,254 / 13,047`; of maxes: `5,968 / 6,254 / 13,047`
- Value length per row group min/median/max of mins: `8 / 8 / 8`; of maxes: `8 / 8 / 8`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/localeventtime_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/localeventtime_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/localeventtime_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/localeventtime_value_length.svg)


Compressed overall:
1. 2,517,265 B (2.40 MiB) compressed - `zstd-3` + `plain`; 8,004,711 B (7.63 MiB) encoded; 3.180562x post-compression ratio; 0.116277% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 4,023,316 B (3.84 MiB) compressed - `zstd-3` + `rle-dict`; 7,355,383 B (7.01 MiB) encoded; 1.989980x post-compression ratio; -37.360327% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict`
3. 4,283,740 B (4.09 MiB) compressed - `snappy` + `plain`; 8,004,715 B (7.63 MiB) encoded; 1.869002x post-compression ratio; 0.028433% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
4. 5,514,562 B (5.26 MiB) compressed - `snappy` + `rle-dict`; 7,339,519 B (7.00 MiB) encoded; 1.451850x post-compression ratio; -22.297401% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict`

ZSTD:
1. 2,517,265 B (2.40 MiB) compressed - `plain`; 8,004,711 B (7.63 MiB) encoded; 3.180562x post-compression ratio; 0.116277% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 4,023,316 B (3.84 MiB) compressed - `rle-dict`; 7,355,383 B (7.01 MiB) encoded; 1.989980x post-compression ratio; -37.360327% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict`

Snappy:
1. 4,283,740 B (4.09 MiB) compressed - `plain`; 8,004,715 B (7.63 MiB) encoded; 1.869002x post-compression ratio; 0.028433% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
2. 5,514,562 B (5.26 MiB) compressed - `rle-dict`; 7,339,519 B (7.00 MiB) encoded; 1.451850x post-compression ratio; -22.297401% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict`

## MobilePhone (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `3 / 7 / 11`
- Page cardinality per row group min/median/max of mins: `3 / 7 / 11`; of maxes: `3 / 7 / 11`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/mobilephone_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/mobilephone_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/mobilephone_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/mobilephone_value_length.svg)


Compressed overall:
1. 22,518 B (21.99 KiB) compressed - `zstd-3` + `plain`; 4,003,586 B (3.82 MiB) encoded; 177.859934x post-compression ratio; 0.000000% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`
2. 23,572 B (23.02 KiB) compressed - `zstd-3` + `rle-dict`; 34,399 B (33.59 KiB) encoded; 169.907093x post-compression ratio; -4.471407% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 29,778 B (29.08 KiB) compressed - `snappy` + `rle-dict`; 34,573 B (33.76 KiB) encoded; 134.496944x post-compression ratio; 623.460273% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
4. 215,287 B (210.24 KiB) compressed - `snappy` + `plain`; 4,003,712 B (3.82 MiB) encoded; 18.603306x post-compression ratio; 0.067352% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 22,518 B (21.99 KiB) compressed - `plain`; 4,003,586 B (3.82 MiB) encoded; 177.859934x post-compression ratio; 0.000000% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`
2. 23,572 B (23.02 KiB) compressed - `rle-dict`; 34,399 B (33.59 KiB) encoded; 169.907093x post-compression ratio; -4.471407% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`

Snappy:
1. 29,778 B (29.08 KiB) compressed - `rle-dict`; 34,573 B (33.76 KiB) encoded; 134.496944x post-compression ratio; 623.460273% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
2. 215,287 B (210.24 KiB) compressed - `plain`; 4,003,712 B (3.82 MiB) encoded; 18.603306x post-compression ratio; 0.067352% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## MobilePhoneModel (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `2 / 4 / 6`
- Page cardinality per row group min/median/max of mins: `2 / 4 / 6`; of maxes: `2 / 4 / 6`
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `4 / 6 / 17`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/mobilephonemodel_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/mobilephonemodel_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/mobilephonemodel_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/mobilephonemodel_value_length.svg)


Compressed overall:
1. 20,159 B (19.69 KiB) compressed - `zstd-3` + `rle-dict`; 25,254 B (24.66 KiB) encoded; 202.699687x post-compression ratio; 12.143459% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 22,582 B (22.05 KiB) compressed - `zstd-3` + `plain`; 4,084,891 B (3.90 MiB) encoded; 180.950447x post-compression ratio; 0.110708% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
3. 23,006 B (22.47 KiB) compressed - `snappy` + `rle-dict`; 25,137 B (24.55 KiB) encoded; 177.615535x post-compression ratio; 856.437451% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
4. 30,635 B (29.92 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 231,125 B (225.71 KiB) encoded; 133.384136x post-compression ratio; -26.205321% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
5. 40,719 B (39.76 KiB) compressed - `zstd-3` + `delta-byte-array`; 283,596 B (276.95 KiB) encoded; 100.351752x post-compression ratio; -44.480464% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 20,159 B (19.69 KiB) compressed - `rle-dict`; 25,254 B (24.66 KiB) encoded; 202.699687x post-compression ratio; 12.143459% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 22,582 B (22.05 KiB) compressed - `plain`; 4,084,891 B (3.90 MiB) encoded; 180.950447x post-compression ratio; 0.110708% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
3. 30,635 B (29.92 KiB) compressed - `delta-length-byte-array`; 231,125 B (225.71 KiB) encoded; 133.384136x post-compression ratio; -26.205321% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
4. 40,719 B (39.76 KiB) compressed - `delta-byte-array`; 283,596 B (276.95 KiB) encoded; 100.351752x post-compression ratio; -44.480464% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

Snappy:
1. 23,006 B (22.47 KiB) compressed - `rle-dict`; 25,137 B (24.55 KiB) encoded; 177.615535x post-compression ratio; 856.437451% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 42,471 B (41.48 KiB) compressed - `delta-length-byte-array`; 230,819 B (225.41 KiB) encoded; 96.212074x post-compression ratio; 418.089991% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 56,742 B (55.41 KiB) compressed - `delta-byte-array`; 281,770 B (275.17 KiB) encoded; 72.014081x post-compression ratio; 287.786825% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
4. 219,895 B (214.74 KiB) compressed - `plain`; 4,085,072 B (3.90 MiB) encoded; 18.582610x post-compression ratio; 0.065031% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`

## NetMajor (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `3 / 4 / 4`
- Page cardinality per row group min/median/max of mins: `3 / 4 / 4`; of maxes: `3 / 4 / 4`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/netmajor_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/netmajor_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/netmajor_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/netmajor_value_length.svg)


Compressed overall:
1. 25,367 B (24.77 KiB) compressed - `zstd-3` + `rle-dict`; 36,548 B (35.69 KiB) encoded; 157.884338x post-compression ratio; 3.066977% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 26,145 B (25.53 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 153.186154x post-compression ratio; 0.000000% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`
3. 32,013 B (31.26 KiB) compressed - `snappy` + `rle-dict`; 36,846 B (35.98 KiB) encoded; 125.107050x post-compression ratio; 584.774935% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
4. 219,072 B (213.94 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 18.281898x post-compression ratio; 0.066188% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 25,367 B (24.77 KiB) compressed - `rle-dict`; 36,548 B (35.69 KiB) encoded; 157.884338x post-compression ratio; 3.066977% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 26,145 B (25.53 KiB) compressed - `plain`; 4,003,587 B (3.82 MiB) encoded; 153.186154x post-compression ratio; 0.000000% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`

Snappy:
1. 32,013 B (31.26 KiB) compressed - `rle-dict`; 36,846 B (35.98 KiB) encoded; 125.107050x post-compression ratio; 584.774935% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 219,072 B (213.94 KiB) compressed - `plain`; 4,003,713 B (3.82 MiB) encoded; 18.281898x post-compression ratio; 0.066188% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## NetMinor (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `2 / 3 / 3`
- Page cardinality per row group min/median/max of mins: `2 / 3 / 3`; of maxes: `2 / 3 / 3`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/netminor_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/netminor_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/netminor_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/netminor_value_length.svg)


Compressed overall:
1. 23,100 B (22.56 KiB) compressed - `zstd-3` + `rle-dict`; 33,117 B (32.34 KiB) encoded; 173.378615x post-compression ratio; 7.865801% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 24,917 B (24.33 KiB) compressed - `zstd-3` + `plain`; 4,003,586 B (3.82 MiB) encoded; 160.735482x post-compression ratio; 0.000000% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`
3. 28,832 B (28.16 KiB) compressed - `snappy` + `rle-dict`; 33,064 B (32.29 KiB) encoded; 138.909753x post-compression ratio; 653.926193% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 217,237 B (212.15 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 18.436298x post-compression ratio; 0.062144% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 23,100 B (22.56 KiB) compressed - `rle-dict`; 33,117 B (32.34 KiB) encoded; 173.378615x post-compression ratio; 7.865801% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 24,917 B (24.33 KiB) compressed - `plain`; 4,003,586 B (3.82 MiB) encoded; 160.735482x post-compression ratio; 0.000000% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`

Snappy:
1. 28,832 B (28.16 KiB) compressed - `rle-dict`; 33,064 B (32.29 KiB) encoded; 138.909753x post-compression ratio; 653.926193% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
2. 217,237 B (212.15 KiB) compressed - `plain`; 4,003,713 B (3.82 MiB) encoded; 18.436298x post-compression ratio; 0.062144% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## OS (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `16 / 23 / 31`
- Page cardinality per row group min/median/max of mins: `16 / 23 / 31`; of maxes: `16 / 23 / 31`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/os_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/os_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/os_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/os_value_length.svg)


Compressed overall:
1. 105,900 B (103.42 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 37.819160x post-compression ratio; 0.004721% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain`
2. 119,893 B (117.08 KiB) compressed - `zstd-3` + `rle-dict`; 229,346 B (223.97 KiB) encoded; 33.405195x post-compression ratio; -11.667070% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 148,976 B (145.48 KiB) compressed - `snappy` + `rle-dict`; 229,949 B (224.56 KiB) encoded; 26.883854x post-compression ratio; 113.531039% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`
4. 317,972 B (310.52 KiB) compressed - `snappy` + `plain`; 4,003,710 B (3.82 MiB) encoded; 12.595603x post-compression ratio; 0.043400% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 105,900 B (103.42 KiB) compressed - `plain`; 4,003,587 B (3.82 MiB) encoded; 37.819160x post-compression ratio; 0.004721% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain`
2. 119,893 B (117.08 KiB) compressed - `rle-dict`; 229,346 B (223.97 KiB) encoded; 33.405195x post-compression ratio; -11.667070% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 148,976 B (145.48 KiB) compressed - `rle-dict`; 229,949 B (224.56 KiB) encoded; 26.883854x post-compression ratio; 113.531039% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`
2. 317,972 B (310.52 KiB) compressed - `plain`; 4,003,710 B (3.82 MiB) encoded; 12.595603x post-compression ratio; 0.043400% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## OpenerName (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/openername_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/openername_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/openername_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/openername_value_length.svg)


Compressed overall:
1. 4,233 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,532 B (3.82 MiB) encoded; 946.148358x post-compression ratio; 0.141743% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 5,360 B (5.23 KiB) compressed - `snappy` + `rle-dict`; 5,124 B (5.00 KiB) encoded; 747.210075x post-compression ratio; 3709.888060% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
3. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.300586x post-compression ratio; -29.054393% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
4. 204,063 B (199.28 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 19.626517x post-compression ratio; 0.072037% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 4,233 B (4.13 KiB) compressed - `plain`; 4,003,532 B (3.82 MiB) encoded; 946.148358x post-compression ratio; 0.141743% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 5,975 B (5.83 KiB) compressed - `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.300586x post-compression ratio; -29.054393% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 5,360 B (5.23 KiB) compressed - `rle-dict`; 5,124 B (5.00 KiB) encoded; 747.210075x post-compression ratio; 3709.888060% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 204,063 B (199.28 KiB) compressed - `plain`; 4,003,713 B (3.82 MiB) encoded; 19.626517x post-compression ratio; 0.072037% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## OpenstatAdID (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 9 / 18`
- Page cardinality per row group min/median/max of mins: `1 / 9 / 18`; of maxes: `1 / 9 / 18`
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 13 / 22`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/openstatadid_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/openstatadid_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/openstatadid_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/openstatadid_value_length.svg)


Compressed overall:
1. 18,284 B (17.86 KiB) compressed - `zstd-3` + `rle-dict`; 22,605 B (22.08 KiB) encoded; 220.437322x post-compression ratio; 4.058193% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 18,939 B (18.50 KiB) compressed - `zstd-3` + `plain`; 4,029,016 B (3.84 MiB) encoded; 212.813559x post-compression ratio; 0.459370% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 20,848 B (20.36 KiB) compressed - `snappy` + `rle-dict`; 22,744 B (22.21 KiB) encoded; 193.326746x post-compression ratio; 951.194359% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
4. 28,592 B (27.92 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 130,898 B (127.83 KiB) encoded; 140.965165x post-compression ratio; -33.456911% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict`
5. 33,547 B (32.76 KiB) compressed - `zstd-3` + `delta-byte-array`; 208,868 B (203.97 KiB) encoded; 120.144156x post-compression ratio; -43.285540% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`

ZSTD:
1. 18,284 B (17.86 KiB) compressed - `rle-dict`; 22,605 B (22.08 KiB) encoded; 220.437322x post-compression ratio; 4.058193% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 18,939 B (18.50 KiB) compressed - `plain`; 4,029,016 B (3.84 MiB) encoded; 212.813559x post-compression ratio; 0.459370% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 28,592 B (27.92 KiB) compressed - `delta-length-byte-array`; 130,898 B (127.83 KiB) encoded; 140.965165x post-compression ratio; -33.456911% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict`
4. 33,547 B (32.76 KiB) compressed - `delta-byte-array`; 208,868 B (203.97 KiB) encoded; 120.144156x post-compression ratio; -43.285540% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`

Snappy:
1. 20,848 B (20.36 KiB) compressed - `rle-dict`; 22,744 B (22.21 KiB) encoded; 193.326746x post-compression ratio; 951.194359% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 34,347 B (33.54 KiB) compressed - `delta-length-byte-array`; 129,819 B (126.78 KiB) encoded; 117.345794x post-compression ratio; 538.055725% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 44,359 B (43.32 KiB) compressed - `delta-byte-array`; 210,052 B (205.13 KiB) encoded; 90.860389x post-compression ratio; 394.044050% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-rle-dict`
4. 219,056 B (213.92 KiB) compressed - `plain`; 4,029,259 B (3.84 MiB) encoded; 18.399295x post-compression ratio; 0.044281% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`

## OpenstatCampaignID (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 4 / 6`
- Page cardinality per row group min/median/max of mins: `1 / 4 / 6`; of maxes: `1 / 4 / 6`
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 10 / 12`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/openstatcampaignid_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/openstatcampaignid_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/openstatcampaignid_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/openstatcampaignid_value_length.svg)


Compressed overall:
1. 15,591 B (15.23 KiB) compressed - `zstd-3` + `rle-dict`; 16,315 B (15.93 KiB) encoded; 258.295042x post-compression ratio; 6.099673% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 15,872 B (15.50 KiB) compressed - `snappy` + `rle-dict`; 16,343 B (15.96 KiB) encoded; 253.722152x post-compression ratio; 1253.011593% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
3. 16,486 B (16.10 KiB) compressed - `zstd-3` + `plain`; 4,025,591 B (3.84 MiB) encoded; 244.272595x post-compression ratio; 0.339682% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-rle-dict`
4. 23,202 B (22.66 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 122,478 B (119.61 KiB) encoded; 173.565986x post-compression ratio; -28.704422% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`
5. 28,149 B (27.49 KiB) compressed - `zstd-3` + `delta-byte-array`; 197,351 B (192.73 KiB) encoded; 143.062915x post-compression ratio; -41.234147% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`

ZSTD:
1. 15,591 B (15.23 KiB) compressed - `rle-dict`; 16,315 B (15.93 KiB) encoded; 258.295042x post-compression ratio; 6.099673% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 16,486 B (16.10 KiB) compressed - `plain`; 4,025,591 B (3.84 MiB) encoded; 244.272595x post-compression ratio; 0.339682% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-rle-dict`
3. 23,202 B (22.66 KiB) compressed - `delta-length-byte-array`; 122,478 B (119.61 KiB) encoded; 173.565986x post-compression ratio; -28.704422% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`
4. 28,149 B (27.49 KiB) compressed - `delta-byte-array`; 197,351 B (192.73 KiB) encoded; 143.062915x post-compression ratio; -41.234147% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`

Snappy:
1. 15,872 B (15.50 KiB) compressed - `rle-dict`; 16,343 B (15.96 KiB) encoded; 253.722152x post-compression ratio; 1253.011593% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 28,470 B (27.80 KiB) compressed - `delta-length-byte-array`; 122,381 B (119.51 KiB) encoded; 141.449877x post-compression ratio; 654.302775% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 38,441 B (37.54 KiB) compressed - `delta-byte-array`; 198,099 B (193.46 KiB) encoded; 104.759970x post-compression ratio; 458.648318% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-rle-dict`
4. 214,628 B (209.60 KiB) compressed - `plain`; 4,025,849 B (3.84 MiB) encoded; 18.763060x post-compression ratio; 0.056843% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`

## OpenstatServiceName (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 2 / 6`
- Page cardinality per row group min/median/max of mins: `1 / 2 / 6`; of maxes: `1 / 2 / 6`
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 16 / 16`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/openstatservicename_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/openstatservicename_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/openstatservicename_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/openstatservicename_value_length.svg)


Compressed overall:
1. 17,778 B (17.36 KiB) compressed - `zstd-3` + `rle-dict`; 18,506 B (18.07 KiB) encoded; 228.614411x post-compression ratio; 8.178648% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 17,876 B (17.46 KiB) compressed - `snappy` + `rle-dict`; 18,558 B (18.12 KiB) encoded; 227.361099x post-compression ratio; 1122.471470% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
3. 19,194 B (18.74 KiB) compressed - `zstd-3` + `plain`; 4,062,487 B (3.87 MiB) encoded; 211.748828x post-compression ratio; 0.197979% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
4. 27,305 B (26.67 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 216,495 B (211.42 KiB) encoded; 148.848453x post-compression ratio; -29.566014% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`
5. 33,746 B (32.96 KiB) compressed - `zstd-3` + `delta-byte-array`; 301,258 B (294.20 KiB) encoded; 120.438185x post-compression ratio; -43.009542% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`

ZSTD:
1. 17,778 B (17.36 KiB) compressed - `rle-dict`; 18,506 B (18.07 KiB) encoded; 228.614411x post-compression ratio; 8.178648% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 19,194 B (18.74 KiB) compressed - `plain`; 4,062,487 B (3.87 MiB) encoded; 211.748828x post-compression ratio; 0.197979% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
3. 27,305 B (26.67 KiB) compressed - `delta-length-byte-array`; 216,495 B (211.42 KiB) encoded; 148.848453x post-compression ratio; -29.566014% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`
4. 33,746 B (32.96 KiB) compressed - `delta-byte-array`; 301,258 B (294.20 KiB) encoded; 120.438185x post-compression ratio; -43.009542% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`

Snappy:
1. 17,876 B (17.46 KiB) compressed - `rle-dict`; 18,558 B (18.12 KiB) encoded; 227.361099x post-compression ratio; 1122.471470% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 36,951 B (36.08 KiB) compressed - `delta-length-byte-array`; 216,700 B (211.62 KiB) encoded; 109.991800x post-compression ratio; 491.402127% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
3. 47,971 B (46.85 KiB) compressed - `delta-byte-array`; 300,177 B (293.14 KiB) encoded; 84.724250x post-compression ratio; 355.543974% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
4. 218,413 B (213.29 KiB) compressed - `plain`; 4,062,767 B (3.87 MiB) encoded; 18.608357x post-compression ratio; 0.053110% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`

## OpenstatSourceID (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 3 / 7`
- Page cardinality per row group min/median/max of mins: `1 / 3 / 7`; of maxes: `1 / 3 / 7`
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 20 / 31`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/openstatsourceid_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/openstatsourceid_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/openstatsourceid_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/openstatsourceid_value_length.svg)


Compressed overall:
1. 12,812 B (12.51 KiB) compressed - `zstd-3` + `rle-dict`; 14,560 B (14.22 KiB) encoded; 316.374337x post-compression ratio; 3.980643% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 13,231 B (12.92 KiB) compressed - `zstd-3` + `plain`; 4,051,710 B (3.86 MiB) encoded; 306.355378x post-compression ratio; 0.687779% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
3. 13,428 B (13.11 KiB) compressed - `snappy` + `rle-dict`; 14,784 B (14.44 KiB) encoded; 301.860888x post-compression ratio; 1496.440274% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
4. 18,296 B (17.87 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 155,189 B (151.55 KiB) encoded; 221.545037x post-compression ratio; -27.186270% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`
5. 23,226 B (22.68 KiB) compressed - `zstd-3` + `delta-byte-array`; 231,876 B (226.44 KiB) encoded; 174.519418x post-compression ratio; -42.641867% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 12,812 B (12.51 KiB) compressed - `rle-dict`; 14,560 B (14.22 KiB) encoded; 316.374337x post-compression ratio; 3.980643% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 13,231 B (12.92 KiB) compressed - `plain`; 4,051,710 B (3.86 MiB) encoded; 306.355378x post-compression ratio; 0.687779% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
3. 18,296 B (17.87 KiB) compressed - `delta-length-byte-array`; 155,189 B (151.55 KiB) encoded; 221.545037x post-compression ratio; -27.186270% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`
4. 23,226 B (22.68 KiB) compressed - `delta-byte-array`; 231,876 B (226.44 KiB) encoded; 174.519418x post-compression ratio; -42.641867% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`

Snappy:
1. 13,428 B (13.11 KiB) compressed - `rle-dict`; 14,784 B (14.44 KiB) encoded; 301.860888x post-compression ratio; 1496.440274% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 24,426 B (23.85 KiB) compressed - `delta-length-byte-array`; 154,936 B (151.30 KiB) encoded; 165.945632x post-compression ratio; 777.630394% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 33,748 B (32.96 KiB) compressed - `delta-byte-array`; 232,759 B (227.30 KiB) encoded; 120.107503x post-compression ratio; 535.208012% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-rle-dict`
4. 214,337 B (209.31 KiB) compressed - `plain`; 4,052,055 B (3.86 MiB) encoded; 18.911285x post-compression ratio; 0.015396% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-plain`

## OriginalURL (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `173`
- Row-group cardinality min/median/max: `95 / 227 / 6,169`
- Page cardinality per row group min/median/max of mins: `12 / 183 / 850`; of maxes: `95 / 227 / 1,065`
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `357 / 517 / 3,723`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/originalurl_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/originalurl_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/originalurl_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/originalurl_value_length.svg)


Compressed overall:
1. 4,878,331 B (4.65 MiB) compressed - `zstd-3` + `rle-dict`; 21,272,072 B (20.29 MiB) encoded; 6.532149x post-compression ratio; 9.127240% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 5,316,593 B (5.07 MiB) compressed - `zstd-3` + `plain`; 31,856,668 B (30.38 MiB) encoded; 5.993685x post-compression ratio; 0.131569% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-rle-dict`
3. 5,587,537 B (5.33 MiB) compressed - `zstd-3` + `delta-byte-array`; 21,049,324 B (20.07 MiB) encoded; 5.703046x post-compression ratio; -4.723888% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
4. 5,605,855 B (5.35 MiB) compressed - `zstd-3` + `delta-length-byte-array`; 28,789,340 B (27.46 MiB) encoded; 5.684411x post-compression ratio; -5.035218% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
5. 6,160,774 B (5.88 MiB) compressed - `snappy` + `rle-dict`; 21,273,479 B (20.29 MiB) encoded; 5.172399x post-compression ratio; 14.228813% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`

ZSTD:
1. 4,878,331 B (4.65 MiB) compressed - `rle-dict`; 21,272,072 B (20.29 MiB) encoded; 6.532149x post-compression ratio; 9.127240% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 5,316,593 B (5.07 MiB) compressed - `plain`; 31,856,668 B (30.38 MiB) encoded; 5.993685x post-compression ratio; 0.131569% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-rle-dict`
3. 5,587,537 B (5.33 MiB) compressed - `delta-byte-array`; 21,049,324 B (20.07 MiB) encoded; 5.703046x post-compression ratio; -4.723888% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
4. 5,605,855 B (5.35 MiB) compressed - `delta-length-byte-array`; 28,789,340 B (27.46 MiB) encoded; 5.684411x post-compression ratio; -5.035218% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`

Snappy:
1. 6,160,774 B (5.88 MiB) compressed - `rle-dict`; 21,273,479 B (20.29 MiB) encoded; 5.172399x post-compression ratio; 14.228813% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 6,746,652 B (6.43 MiB) compressed - `delta-byte-array`; 21,050,091 B (20.07 MiB) encoded; 4.723229x post-compression ratio; 4.309204% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`
3. 7,036,093 B (6.71 MiB) compressed - `plain`; 31,860,219 B (30.38 MiB) encoded; 4.528931x post-compression ratio; 0.018277% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`
4. 7,046,903 B (6.72 MiB) compressed - `delta-length-byte-array`; 28,796,803 B (27.46 MiB) encoded; 4.521984x post-compression ratio; -0.135152% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-rle-dict`

## PageCharset (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `122`
- Row-group cardinality min/median/max: `1 / 2 / 3`
- Page cardinality per row group min/median/max of mins: `1 / 2 / 3`; of maxes: `1 / 2 / 3`
- Value length per row group min/median/max of mins: `0 / 0 / 20`; of maxes: `7 / 20 / 20`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/pagecharset_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/pagecharset_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/pagecharset_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/pagecharset_value_length.svg)


Compressed overall:
1. 9,051 B (8.84 KiB) compressed - `snappy` + `rle-dict`; 8,870 B (8.66 KiB) encoded; 1944.421611x post-compression ratio; 10116.694288% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 9,926 B (9.69 KiB) compressed - `zstd-3` + `rle-dict`; 8,862 B (8.65 KiB) encoded; 1773.016321x post-compression ratio; 47.209349% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain`
3. 14,594 B (14.25 KiB) compressed - `zstd-3` + `plain`; 17,594,774 B (16.78 MiB) encoded; 1205.903796x post-compression ratio; 0.123338% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`
4. 18,662 B (18.22 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 13,656,839 B (13.02 MiB) encoded; 943.037188x post-compression ratio; -21.701854% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
5. 21,415 B (20.91 KiB) compressed - `zstd-3` + `delta-byte-array`; 140,380 B (137.09 KiB) encoded; 821.805277x post-compression ratio; -31.767453% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`

ZSTD:
1. 9,926 B (9.69 KiB) compressed - `rle-dict`; 8,862 B (8.65 KiB) encoded; 1773.016321x post-compression ratio; 47.209349% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain`
2. 14,594 B (14.25 KiB) compressed - `plain`; 17,594,774 B (16.78 MiB) encoded; 1205.903796x post-compression ratio; 0.123338% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`
3. 18,662 B (18.22 KiB) compressed - `delta-length-byte-array`; 13,656,839 B (13.02 MiB) encoded; 943.037188x post-compression ratio; -21.701854% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
4. 21,415 B (20.91 KiB) compressed - `delta-byte-array`; 140,380 B (137.09 KiB) encoded; 821.805277x post-compression ratio; -31.767453% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`

Snappy:
1. 9,051 B (8.84 KiB) compressed - `rle-dict`; 8,870 B (8.66 KiB) encoded; 1944.421611x post-compression ratio; 10116.694288% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 26,681 B (26.06 KiB) compressed - `delta-byte-array`; 141,925 B (138.60 KiB) encoded; 659.606462x post-compression ratio; 3365.810877% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
3. 699,152 B (682.77 KiB) compressed - `delta-length-byte-array`; 13,658,214 B (13.03 MiB) encoded; 25.171865x post-compression ratio; 32.262083% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain`
4. 919,930 B (898.37 KiB) compressed - `plain`; 17,595,645 B (16.78 MiB) encoded; 19.130760x post-compression ratio; 0.519931% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`

## ParamCurrency (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length per row group min/median/max of mins: `3 / 3 / 3`; of maxes: `3 / 3 / 3`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/paramcurrency_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/paramcurrency_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/paramcurrency_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/paramcurrency_value_length.svg)


Compressed overall:
1. 5,256 B (5.13 KiB) compressed - `zstd-3` + `plain`; 7,003,298 B (6.68 MiB) encoded; 1332.711758x post-compression ratio; 0.095129% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 5,394 B (5.27 KiB) compressed - `snappy` + `rle-dict`; 5,154 B (5.03 KiB) encoded; 1298.615684x post-compression ratio; 6468.891361% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
3. 5,795 B (5.66 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 3,043,978 B (2.90 MiB) encoded; 1208.754616x post-compression ratio; -9.214840% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
4. 6,132 B (5.99 KiB) compressed - `zstd-3` + `rle-dict`; 5,070 B (4.95 KiB) encoded; 1142.324364x post-compression ratio; -14.204175% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain`
5. 6,869 B (6.71 KiB) compressed - `zstd-3` + `delta-byte-array`; 86,017 B (84.00 KiB) encoded; 1019.760227x post-compression ratio; -23.409521% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`

ZSTD:
1. 5,256 B (5.13 KiB) compressed - `plain`; 7,003,298 B (6.68 MiB) encoded; 1332.711758x post-compression ratio; 0.095129% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 5,795 B (5.66 KiB) compressed - `delta-length-byte-array`; 3,043,978 B (2.90 MiB) encoded; 1208.754616x post-compression ratio; -9.214840% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
3. 6,132 B (5.99 KiB) compressed - `rle-dict`; 5,070 B (4.95 KiB) encoded; 1142.324364x post-compression ratio; -14.204175% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain`
4. 6,869 B (6.71 KiB) compressed - `delta-byte-array`; 86,017 B (84.00 KiB) encoded; 1019.760227x post-compression ratio; -23.409521% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`

Snappy:
1. 5,394 B (5.27 KiB) compressed - `rle-dict`; 5,154 B (5.03 KiB) encoded; 1298.615684x post-compression ratio; 6468.891361% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 10,304 B (10.06 KiB) compressed - `delta-byte-array`; 86,274 B (84.25 KiB) encoded; 679.807162x post-compression ratio; 3338.722826% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
3. 157,287 B (153.60 KiB) compressed - `delta-length-byte-array`; 3,044,410 B (2.90 MiB) encoded; 44.534723x post-compression ratio; 125.273545% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
4. 354,236 B (345.93 KiB) compressed - `plain`; 7,003,535 B (6.68 MiB) encoded; 19.774199x post-compression ratio; 0.025407% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-rle-dict-ts-plain`

## ParamCurrencyID (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/paramcurrencyid_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/paramcurrencyid_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/paramcurrencyid_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/paramcurrencyid_value_length.svg)


Compressed overall:
1. 4,228 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; 0.283822% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 5,360 B (5.23 KiB) compressed - `snappy` + `rle-dict`; 5,124 B (5.00 KiB) encoded; 747.211567x post-compression ratio; 3709.981343% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
3. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; -29.037657% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
4. 204,063 B (199.28 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 19.626557x post-compression ratio; 0.074487% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 4,228 B (4.13 KiB) compressed - `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; 0.283822% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 5,975 B (5.83 KiB) compressed - `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; -29.037657% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 5,360 B (5.23 KiB) compressed - `rle-dict`; 5,124 B (5.00 KiB) encoded; 747.211567x post-compression ratio; 3709.981343% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 204,063 B (199.28 KiB) compressed - `plain`; 4,003,713 B (3.82 MiB) encoded; 19.626557x post-compression ratio; 0.074487% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## ParamOrderID (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 0 / 0`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/paramorderid_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/paramorderid_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/paramorderid_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/paramorderid_value_length.svg)


Compressed overall:
1. 2,866 B (2.80 KiB) compressed - `zstd-3` + `plain`; 4,002,163 B (3.82 MiB) encoded; 1396.775297x post-compression ratio; 0.209351% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`
2. 3,569 B (3.49 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 41,691 B (40.71 KiB) encoded; 1121.646960x post-compression ratio; -19.529280% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
3. 3,681 B (3.59 KiB) compressed - `zstd-3` + `delta-byte-array`; 81,277 B (79.37 KiB) encoded; 1087.519152x post-compression ratio; -21.977723% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 4,494 B (4.39 KiB) compressed - `snappy` + `rle-dict`; 4,254 B (4.15 KiB) encoded; 890.778371x post-compression ratio; 4411.593235% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
5. 4,840 B (4.73 KiB) compressed - `snappy` + `delta-length-byte-array`; 41,796 B (40.82 KiB) encoded; 827.098760x post-compression ratio; 4089.070248% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`

ZSTD:
1. 2,866 B (2.80 KiB) compressed - `plain`; 4,002,163 B (3.82 MiB) encoded; 1396.775297x post-compression ratio; 0.209351% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`
2. 3,569 B (3.49 KiB) compressed - `delta-length-byte-array`; 41,691 B (40.71 KiB) encoded; 1121.646960x post-compression ratio; -19.529280% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
3. 3,681 B (3.59 KiB) compressed - `delta-byte-array`; 81,277 B (79.37 KiB) encoded; 1087.519152x post-compression ratio; -21.977723% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 5,247 B (5.12 KiB) compressed - `rle-dict`; 4,185 B (4.09 KiB) encoded; 762.942253x post-compression ratio; -45.263960% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain`

Snappy:
1. 4,494 B (4.39 KiB) compressed - `rle-dict`; 4,254 B (4.15 KiB) encoded; 890.778371x post-compression ratio; 4411.593235% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 4,840 B (4.73 KiB) compressed - `delta-length-byte-array`; 41,796 B (40.82 KiB) encoded; 827.098760x post-compression ratio; 4089.070248% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 6,892 B (6.73 KiB) compressed - `delta-byte-array`; 81,417 B (79.51 KiB) encoded; 580.841265x post-compression ratio; 2841.831109% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 202,705 B (197.95 KiB) compressed - `plain`; 4,002,337 B (3.82 MiB) encoded; 19.748689x post-compression ratio; 0.022693% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`

## ParamPrice (int64)

Column shape stats:
- Parquet type: `INT(64,true)`; physical type: `INT64`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length per row group min/median/max of mins: `8 / 8 / 8`; of maxes: `8 / 8 / 8`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/paramprice_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/paramprice_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/paramprice_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/paramprice_value_length.svg)


Compressed overall:
1. 5,675 B (5.54 KiB) compressed - `zstd-3` + `plain`; 8,004,440 B (7.63 MiB) encoded; 1410.804934x post-compression ratio; 0.881057% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 6,540 B (6.39 KiB) compressed - `snappy` + `rle-dict`; 6,304 B (6.16 KiB) encoded; 1224.207645x post-compression ratio; 6095.565749% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
3. 7,115 B (6.95 KiB) compressed - `zstd-3` + `rle-dict`; 6,089 B (5.95 KiB) encoded; 1125.273085x post-compression ratio; -19.536191% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
4. 405,010 B (395.52 KiB) compressed - `snappy` + `plain`; 8,004,654 B (7.63 MiB) encoded; 19.768198x post-compression ratio; 0.044443% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 5,675 B (5.54 KiB) compressed - `plain`; 8,004,440 B (7.63 MiB) encoded; 1410.804934x post-compression ratio; 0.881057% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 7,115 B (6.95 KiB) compressed - `rle-dict`; 6,089 B (5.95 KiB) encoded; 1125.273085x post-compression ratio; -19.536191% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 6,540 B (6.39 KiB) compressed - `rle-dict`; 6,304 B (6.16 KiB) encoded; 1224.207645x post-compression ratio; 6095.565749% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 405,010 B (395.52 KiB) compressed - `plain`; 8,004,654 B (7.63 MiB) encoded; 19.768198x post-compression ratio; 0.044443% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## Params (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 0 / 0`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/params_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/params_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/params_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/params_value_length.svg)


Compressed overall:
1. 2,866 B (2.80 KiB) compressed - `zstd-3` + `plain`; 4,002,163 B (3.82 MiB) encoded; 1396.775297x post-compression ratio; 0.209351% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`
2. 3,569 B (3.49 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 41,691 B (40.71 KiB) encoded; 1121.646960x post-compression ratio; -19.529280% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
3. 3,681 B (3.59 KiB) compressed - `zstd-3` + `delta-byte-array`; 81,277 B (79.37 KiB) encoded; 1087.519152x post-compression ratio; -21.977723% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 4,494 B (4.39 KiB) compressed - `snappy` + `rle-dict`; 4,254 B (4.15 KiB) encoded; 890.778371x post-compression ratio; 4411.593235% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
5. 4,840 B (4.73 KiB) compressed - `snappy` + `delta-length-byte-array`; 41,796 B (40.82 KiB) encoded; 827.098760x post-compression ratio; 4089.070248% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`

ZSTD:
1. 2,866 B (2.80 KiB) compressed - `plain`; 4,002,163 B (3.82 MiB) encoded; 1396.775297x post-compression ratio; 0.209351% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`
2. 3,569 B (3.49 KiB) compressed - `delta-length-byte-array`; 41,691 B (40.71 KiB) encoded; 1121.646960x post-compression ratio; -19.529280% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
3. 3,681 B (3.59 KiB) compressed - `delta-byte-array`; 81,277 B (79.37 KiB) encoded; 1087.519152x post-compression ratio; -21.977723% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 5,247 B (5.12 KiB) compressed - `rle-dict`; 4,185 B (4.09 KiB) encoded; 762.942253x post-compression ratio; -45.263960% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain`

Snappy:
1. 4,494 B (4.39 KiB) compressed - `rle-dict`; 4,254 B (4.15 KiB) encoded; 890.778371x post-compression ratio; 4411.593235% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 4,840 B (4.73 KiB) compressed - `delta-length-byte-array`; 41,796 B (40.82 KiB) encoded; 827.098760x post-compression ratio; 4089.070248% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 6,892 B (6.73 KiB) compressed - `delta-byte-array`; 81,417 B (79.51 KiB) encoded; 580.841265x post-compression ratio; 2841.831109% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 202,705 B (197.95 KiB) compressed - `plain`; 4,002,337 B (3.82 MiB) encoded; 19.748689x post-compression ratio; 0.022693% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`

## Referer (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `349`
- Row-group cardinality min/median/max: `361 / 2,628 / 5,960`
- Page cardinality per row group min/median/max of mins: `1 / 670 / 1,130`; of maxes: `223 / 802 / 1,644`
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `514 / 1,014 / 2,007`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/referer_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/referer_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/referer_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/referer_value_length.svg)


Compressed overall:
1. 11,720,762 B (11.18 MiB) compressed - `zstd-3` + `rle-dict`; 34,380,236 B (32.79 MiB) encoded; 7.136665x post-compression ratio; 21.306439% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 14,212,015 B (13.55 MiB) compressed - `zstd-3` + `plain`; 83,646,116 B (79.77 MiB) encoded; 5.885665x post-compression ratio; 0.042387% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
3. 14,535,820 B (13.86 MiB) compressed - `zstd-3` + `delta-byte-array`; 38,986,977 B (37.18 MiB) encoded; 5.754554x post-compression ratio; -2.186192% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 14,797,085 B (14.11 MiB) compressed - `snappy` + `rle-dict`; 34,402,795 B (32.81 MiB) encoded; 5.652948x post-compression ratio; 28.739019% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
5. 14,991,338 B (14.30 MiB) compressed - `zstd-3` + `delta-length-byte-array`; 80,834,647 B (77.09 MiB) encoded; 5.579699x post-compression ratio; -5.158305% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 11,720,762 B (11.18 MiB) compressed - `rle-dict`; 34,380,236 B (32.79 MiB) encoded; 7.136665x post-compression ratio; 21.306439% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 14,212,015 B (13.55 MiB) compressed - `plain`; 83,646,116 B (79.77 MiB) encoded; 5.885665x post-compression ratio; 0.042387% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
3. 14,535,820 B (13.86 MiB) compressed - `delta-byte-array`; 38,986,977 B (37.18 MiB) encoded; 5.754554x post-compression ratio; -2.186192% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 14,991,338 B (14.30 MiB) compressed - `delta-length-byte-array`; 80,834,647 B (77.09 MiB) encoded; 5.579699x post-compression ratio; -5.158305% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`

Snappy:
1. 14,797,085 B (14.11 MiB) compressed - `rle-dict`; 34,402,795 B (32.81 MiB) encoded; 5.652948x post-compression ratio; 28.739019% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 17,316,834 B (16.51 MiB) compressed - `delta-byte-array`; 38,989,593 B (37.18 MiB) encoded; 4.830395x post-compression ratio; 10.006379% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
3. 19,041,603 B (18.16 MiB) compressed - `plain`; 83,646,065 B (79.77 MiB) encoded; 4.392863x post-compression ratio; 0.042113% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`
4. 19,456,070 B (18.55 MiB) compressed - `delta-length-byte-array`; 80,833,108 B (77.09 MiB) encoded; 4.299283x post-compression ratio; -2.089055% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain`

## RefererCategoryID (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `12 / 35 / 79`
- Page cardinality per row group min/median/max of mins: `12 / 35 / 79`; of maxes: `12 / 35 / 79`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/referercategoryid_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/referercategoryid_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/referercategoryid_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/referercategoryid_value_length.svg)


Compressed overall:
1. 215,941 B (210.88 KiB) compressed - `zstd-3` + `rle-dict`; 516,015 B (503.92 KiB) encoded; 18.546978x post-compression ratio; 27.412117% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
2. 275,126 B (268.68 KiB) compressed - `zstd-3` + `plain`; 4,003,588 B (3.82 MiB) encoded; 14.557159x post-compression ratio; 0.003271% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-rle-dict`
3. 281,038 B (274.45 KiB) compressed - `snappy` + `rle-dict`; 514,238 B (502.19 KiB) encoded; 14.250930x post-compression ratio; 81.117144% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
4. 508,952 B (497.02 KiB) compressed - `snappy` + `plain`; 4,003,743 B (3.82 MiB) encoded; 7.869216x post-compression ratio; 0.011003% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 215,941 B (210.88 KiB) compressed - `rle-dict`; 516,015 B (503.92 KiB) encoded; 18.546978x post-compression ratio; 27.412117% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
2. 275,126 B (268.68 KiB) compressed - `plain`; 4,003,588 B (3.82 MiB) encoded; 14.557159x post-compression ratio; 0.003271% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-rle-dict`

Snappy:
1. 281,038 B (274.45 KiB) compressed - `rle-dict`; 514,238 B (502.19 KiB) encoded; 14.250930x post-compression ratio; 81.117144% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 508,952 B (497.02 KiB) compressed - `plain`; 4,003,743 B (3.82 MiB) encoded; 7.869216x post-compression ratio; 0.011003% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## RefererHash (int64)

Column shape stats:
- Parquet type: `INT(64,true)`; physical type: `INT64`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `378 / 2,729 / 6,051`
- Page cardinality per row group min/median/max of mins: `378 / 2,729 / 6,051`; of maxes: `378 / 2,729 / 6,051`
- Value length per row group min/median/max of mins: `8 / 8 / 8`; of maxes: `8 / 8 / 8`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/refererhash_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/refererhash_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/refererhash_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/refererhash_value_length.svg)


Compressed overall:
1. 2,841,886 B (2.71 MiB) compressed - `zstd-3` + `plain`; 8,004,557 B (7.63 MiB) encoded; 2.817254x post-compression ratio; 0.004117% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
2. 3,502,907 B (3.34 MiB) compressed - `zstd-3` + `rle-dict`; 3,770,600 B (3.60 MiB) encoded; 2.285620x post-compression ratio; -18.867301% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
3. 3,517,282 B (3.35 MiB) compressed - `snappy` + `rle-dict`; 3,759,926 B (3.59 MiB) encoded; 2.276279x post-compression ratio; 3.522237% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
4. 3,637,791 B (3.47 MiB) compressed - `snappy` + `plain`; 8,004,876 B (7.63 MiB) encoded; 2.200873x post-compression ratio; 0.092859% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain`

ZSTD:
1. 2,841,886 B (2.71 MiB) compressed - `plain`; 8,004,557 B (7.63 MiB) encoded; 2.817254x post-compression ratio; 0.004117% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
2. 3,502,907 B (3.34 MiB) compressed - `rle-dict`; 3,770,600 B (3.60 MiB) encoded; 2.285620x post-compression ratio; -18.867301% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`

Snappy:
1. 3,517,282 B (3.35 MiB) compressed - `rle-dict`; 3,759,926 B (3.59 MiB) encoded; 2.276279x post-compression ratio; 3.522237% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
2. 3,637,791 B (3.47 MiB) compressed - `plain`; 8,004,876 B (7.63 MiB) encoded; 2.200873x post-compression ratio; 0.092859% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain`

## RefererRegionID (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `7 / 17 / 33`
- Page cardinality per row group min/median/max of mins: `7 / 17 / 33`; of maxes: `7 / 17 / 33`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/refererregionid_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/refererregionid_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/refererregionid_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/refererregionid_value_length.svg)


Compressed overall:
1. 162,823 B (159.01 KiB) compressed - `zstd-3` + `rle-dict`; 378,845 B (369.97 KiB) encoded; 24.597569x post-compression ratio; 42.090491% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`
2. 218,742 B (213.62 KiB) compressed - `snappy` + `rle-dict`; 377,636 B (368.79 KiB) encoded; 18.309470x post-compression ratio; 109.435316% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
3. 231,308 B (225.89 KiB) compressed - `zstd-3` + `plain`; 4,003,588 B (3.82 MiB) encoded; 17.314792x post-compression ratio; 0.020752% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`
4. 457,942 B (447.21 KiB) compressed - `snappy` + `plain`; 4,003,736 B (3.82 MiB) encoded; 8.745758x post-compression ratio; 0.039525% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 162,823 B (159.01 KiB) compressed - `rle-dict`; 378,845 B (369.97 KiB) encoded; 24.597569x post-compression ratio; 42.090491% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`
2. 231,308 B (225.89 KiB) compressed - `plain`; 4,003,588 B (3.82 MiB) encoded; 17.314792x post-compression ratio; 0.020752% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`

Snappy:
1. 218,742 B (213.62 KiB) compressed - `rle-dict`; 377,636 B (368.79 KiB) encoded; 18.309470x post-compression ratio; 109.435316% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
2. 457,942 B (447.21 KiB) compressed - `plain`; 4,003,736 B (3.82 MiB) encoded; 8.745758x post-compression ratio; 0.039525% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## RegionID (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `112 / 149 / 275`
- Page cardinality per row group min/median/max of mins: `112 / 149 / 275`; of maxes: `112 / 149 / 275`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/regionid_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/regionid_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/regionid_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/regionid_value_length.svg)


Compressed overall:
1. 190,719 B (186.25 KiB) compressed - `zstd-3` + `plain`; 4,003,589 B (3.82 MiB) encoded; 20.999733x post-compression ratio; 0.238571% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`
2. 244,341 B (238.61 KiB) compressed - `zstd-3` + `rle-dict`; 435,033 B (424.84 KiB) encoded; 16.391224x post-compression ratio; -21.759345% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 287,605 B (280.86 KiB) compressed - `snappy` + `rle-dict`; 436,454 B (426.22 KiB) encoded; 13.925516x post-compression ratio; 37.959354% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 396,364 B (387.07 KiB) compressed - `snappy` + `plain`; 4,003,719 B (3.82 MiB) encoded; 10.104470x post-compression ratio; 0.104449% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 190,719 B (186.25 KiB) compressed - `plain`; 4,003,589 B (3.82 MiB) encoded; 20.999733x post-compression ratio; 0.238571% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`
2. 244,341 B (238.61 KiB) compressed - `rle-dict`; 435,033 B (424.84 KiB) encoded; 16.391224x post-compression ratio; -21.759345% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`

Snappy:
1. 287,605 B (280.86 KiB) compressed - `rle-dict`; 436,454 B (426.22 KiB) encoded; 13.925516x post-compression ratio; 37.959354% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
2. 396,364 B (387.07 KiB) compressed - `plain`; 4,003,719 B (3.82 MiB) encoded; 10.104470x post-compression ratio; 0.104449% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## RemoteIP (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `508 / 851 / 1,951`
- Page cardinality per row group min/median/max of mins: `508 / 851 / 1,951`; of maxes: `508 / 851 / 1,951`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/remoteip_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/remoteip_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/remoteip_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/remoteip_value_length.svg)


Compressed overall:
1. 426,425 B (416.43 KiB) compressed - `zstd-3` + `plain`; 4,003,606 B (3.82 MiB) encoded; 9.392149x post-compression ratio; 0.042680% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`
2. 698,454 B (682.08 KiB) compressed - `zstd-3` + `rle-dict`; 927,182 B (905.45 KiB) encoded; 5.734160x post-compression ratio; -38.921246% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
3. 706,731 B (690.17 KiB) compressed - `snappy` + `plain`; 4,003,834 B (3.82 MiB) encoded; 5.667003x post-compression ratio; 0.125932% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`
4. 748,986 B (731.43 KiB) compressed - `snappy` + `rle-dict`; 927,799 B (906.05 KiB) encoded; 5.347292x post-compression ratio; -5.522800% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`

ZSTD:
1. 426,425 B (416.43 KiB) compressed - `plain`; 4,003,606 B (3.82 MiB) encoded; 9.392149x post-compression ratio; 0.042680% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`
2. 698,454 B (682.08 KiB) compressed - `rle-dict`; 927,182 B (905.45 KiB) encoded; 5.734160x post-compression ratio; -38.921246% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`

Snappy:
1. 706,731 B (690.17 KiB) compressed - `plain`; 4,003,834 B (3.82 MiB) encoded; 5.667003x post-compression ratio; 0.125932% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`
2. 748,986 B (731.43 KiB) compressed - `rle-dict`; 927,799 B (906.05 KiB) encoded; 5.347292x post-compression ratio; -5.522800% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`

## ResolutionDepth (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `4 / 4 / 5`
- Page cardinality per row group min/median/max of mins: `4 / 4 / 5`; of maxes: `4 / 4 / 5`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/resolutiondepth_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/resolutiondepth_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/resolutiondepth_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/resolutiondepth_value_length.svg)


Compressed overall:
1. 61,927 B (60.48 KiB) compressed - `zstd-3` + `rle-dict`; 111,133 B (108.53 KiB) encoded; 64.673745x post-compression ratio; 32.559304% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 81,692 B (79.78 KiB) compressed - `zstd-3` + `plain`; 4,003,589 B (3.82 MiB) encoded; 49.026233x post-compression ratio; 0.487196% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
3. 87,347 B (85.30 KiB) compressed - `snappy` + `rle-dict`; 111,326 B (108.72 KiB) encoded; 45.852187x post-compression ratio; 224.054633% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 282,833 B (276.20 KiB) compressed - `snappy` + `plain`; 4,003,714 B (3.82 MiB) encoded; 14.160480x post-compression ratio; 0.077431% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 61,927 B (60.48 KiB) compressed - `rle-dict`; 111,133 B (108.53 KiB) encoded; 64.673745x post-compression ratio; 32.559304% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 81,692 B (79.78 KiB) compressed - `plain`; 4,003,589 B (3.82 MiB) encoded; 49.026233x post-compression ratio; 0.487196% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`

Snappy:
1. 87,347 B (85.30 KiB) compressed - `rle-dict`; 111,326 B (108.72 KiB) encoded; 45.852187x post-compression ratio; 224.054633% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
2. 282,833 B (276.20 KiB) compressed - `plain`; 4,003,714 B (3.82 MiB) encoded; 14.160480x post-compression ratio; 0.077431% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## ResolutionHeight (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `49 / 70 / 103`
- Page cardinality per row group min/median/max of mins: `49 / 70 / 103`; of maxes: `49 / 70 / 103`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/resolutionheight_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/resolutionheight_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/resolutionheight_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/resolutionheight_value_length.svg)


Compressed overall:
1. 186,024 B (181.66 KiB) compressed - `zstd-3` + `plain`; 4,003,586 B (3.82 MiB) encoded; 21.529749x post-compression ratio; 0.000000% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`
2. 205,409 B (200.59 KiB) compressed - `zstd-3` + `rle-dict`; 371,693 B (362.98 KiB) encoded; 19.497929x post-compression ratio; -9.437269% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 242,446 B (236.76 KiB) compressed - `snappy` + `rle-dict`; 373,970 B (365.21 KiB) encoded; 16.519349x post-compression ratio; 53.951395% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
4. 373,098 B (364.35 KiB) compressed - `snappy` + `plain`; 4,003,717 B (3.82 MiB) encoded; 10.734579x post-compression ratio; 0.040472% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 186,024 B (181.66 KiB) compressed - `plain`; 4,003,586 B (3.82 MiB) encoded; 21.529749x post-compression ratio; 0.000000% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`
2. 205,409 B (200.59 KiB) compressed - `rle-dict`; 371,693 B (362.98 KiB) encoded; 19.497929x post-compression ratio; -9.437269% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 242,446 B (236.76 KiB) compressed - `rle-dict`; 373,970 B (365.21 KiB) encoded; 16.519349x post-compression ratio; 53.951395% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
2. 373,098 B (364.35 KiB) compressed - `plain`; 4,003,717 B (3.82 MiB) encoded; 10.734579x post-compression ratio; 0.040472% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## ResolutionWidth (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `48 / 64 / 84`
- Page cardinality per row group min/median/max of mins: `48 / 64 / 84`; of maxes: `48 / 64 / 84`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/resolutionwidth_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/resolutionwidth_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/resolutionwidth_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/resolutionwidth_value_length.svg)


Compressed overall:
1. 187,141 B (182.75 KiB) compressed - `zstd-3` + `plain`; 4,003,584 B (3.82 MiB) encoded; 21.401259x post-compression ratio; 0.000000% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`
2. 203,528 B (198.76 KiB) compressed - `zstd-3` + `rle-dict`; 368,799 B (360.16 KiB) encoded; 19.678143x post-compression ratio; -8.051472% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 242,030 B (236.36 KiB) compressed - `snappy` + `rle-dict`; 366,735 B (358.14 KiB) encoded; 16.547754x post-compression ratio; 54.314754% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
4. 373,369 B (364.62 KiB) compressed - `snappy` + `plain`; 4,003,711 B (3.82 MiB) encoded; 10.726796x post-compression ratio; 0.031872% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 187,141 B (182.75 KiB) compressed - `plain`; 4,003,584 B (3.82 MiB) encoded; 21.401259x post-compression ratio; 0.000000% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`
2. 203,528 B (198.76 KiB) compressed - `rle-dict`; 368,799 B (360.16 KiB) encoded; 19.678143x post-compression ratio; -8.051472% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 242,030 B (236.36 KiB) compressed - `rle-dict`; 366,735 B (358.14 KiB) encoded; 16.547754x post-compression ratio; 54.314754% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 373,369 B (364.62 KiB) compressed - `plain`; 4,003,711 B (3.82 MiB) encoded; 10.726796x post-compression ratio; 0.031872% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## ResponseEndTiming (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `528 / 673 / 1,577`
- Page cardinality per row group min/median/max of mins: `528 / 673 / 1,577`; of maxes: `528 / 673 / 1,577`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/responseendtiming_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/responseendtiming_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/responseendtiming_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/responseendtiming_value_length.svg)


Compressed overall:
1. 937,781 B (915.80 KiB) compressed - `zstd-3` + `plain`; 4,003,640 B (3.82 MiB) encoded; 4.270774x post-compression ratio; 0.000107% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain`
2. 1,038,430 B (1014.09 KiB) compressed - `zstd-3` + `rle-dict`; 1,360,610 B (1.30 MiB) encoded; 3.856833x post-compression ratio; -9.692324% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`
3. 1,211,566 B (1.16 MiB) compressed - `snappy` + `rle-dict`; 1,361,456 B (1.30 MiB) encoded; 3.305681x post-compression ratio; 5.894025% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
4. 1,280,646 B (1.22 MiB) compressed - `snappy` + `plain`; 4,003,770 B (3.82 MiB) encoded; 3.127368x post-compression ratio; 0.181939% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 937,781 B (915.80 KiB) compressed - `plain`; 4,003,640 B (3.82 MiB) encoded; 4.270774x post-compression ratio; 0.000107% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain`
2. 1,038,430 B (1014.09 KiB) compressed - `rle-dict`; 1,360,610 B (1.30 MiB) encoded; 3.856833x post-compression ratio; -9.692324% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`

Snappy:
1. 1,211,566 B (1.16 MiB) compressed - `rle-dict`; 1,361,456 B (1.30 MiB) encoded; 3.305681x post-compression ratio; 5.894025% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 1,280,646 B (1.22 MiB) compressed - `plain`; 4,003,770 B (3.82 MiB) encoded; 3.127368x post-compression ratio; 0.181939% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## ResponseStartTiming (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `800 / 1,112 / 3,761`
- Page cardinality per row group min/median/max of mins: `800 / 1,112 / 3,761`; of maxes: `800 / 1,112 / 3,761`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/responsestarttiming_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/responsestarttiming_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/responsestarttiming_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/responsestarttiming_value_length.svg)


Compressed overall:
1. 1,245,827 B (1.19 MiB) compressed - `zstd-3` + `plain`; 4,003,646 B (3.82 MiB) encoded; 3.214772x post-compression ratio; 0.005137% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-rle-dict`
2. 1,556,751 B (1.48 MiB) compressed - `zstd-3` + `rle-dict`; 1,799,520 B (1.72 MiB) encoded; 2.572698x post-compression ratio; -19.968511% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`
3. 1,715,676 B (1.64 MiB) compressed - `snappy` + `rle-dict`; 1,810,414 B (1.73 MiB) encoded; 2.334386x post-compression ratio; 0.647791% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
4. 1,721,098 B (1.64 MiB) compressed - `snappy` + `plain`; 4,003,772 B (3.82 MiB) encoded; 2.327032x post-compression ratio; 0.330719% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 1,245,827 B (1.19 MiB) compressed - `plain`; 4,003,646 B (3.82 MiB) encoded; 3.214772x post-compression ratio; 0.005137% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-rle-dict`
2. 1,556,751 B (1.48 MiB) compressed - `rle-dict`; 1,799,520 B (1.72 MiB) encoded; 2.572698x post-compression ratio; -19.968511% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`

Snappy:
1. 1,715,676 B (1.64 MiB) compressed - `rle-dict`; 1,810,414 B (1.73 MiB) encoded; 2.334386x post-compression ratio; 0.647791% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
2. 1,721,098 B (1.64 MiB) compressed - `plain`; 4,003,772 B (3.82 MiB) encoded; 2.327032x post-compression ratio; 0.330719% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## Robotness (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `85 / 114 / 200`
- Page cardinality per row group min/median/max of mins: `85 / 114 / 200`; of maxes: `85 / 114 / 200`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/robotness_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/robotness_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/robotness_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/robotness_value_length.svg)


Compressed overall:
1. 173,390 B (169.33 KiB) compressed - `zstd-3` + `plain`; 4,003,775 B (3.82 MiB) encoded; 23.098529x post-compression ratio; 0.177058% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict`
2. 240,793 B (235.15 KiB) compressed - `zstd-3` + `rle-dict`; 415,545 B (405.81 KiB) encoded; 16.632768x post-compression ratio; -27.864597% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 281,953 B (275.34 KiB) compressed - `snappy` + `rle-dict`; 416,715 B (406.95 KiB) encoded; 14.204687x post-compression ratio; 35.443496% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-rle-dict`
4. 381,598 B (372.65 KiB) compressed - `snappy` + `plain`; 4,003,720 B (3.82 MiB) encoded; 10.495480x post-compression ratio; 0.075734% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 173,390 B (169.33 KiB) compressed - `plain`; 4,003,775 B (3.82 MiB) encoded; 23.098529x post-compression ratio; 0.177058% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict`
2. 240,793 B (235.15 KiB) compressed - `rle-dict`; 415,545 B (405.81 KiB) encoded; 16.632768x post-compression ratio; -27.864597% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`

Snappy:
1. 281,953 B (275.34 KiB) compressed - `rle-dict`; 416,715 B (406.95 KiB) encoded; 14.204687x post-compression ratio; 35.443496% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-rle-dict`
2. 381,598 B (372.65 KiB) compressed - `plain`; 4,003,720 B (3.82 MiB) encoded; 10.495480x post-compression ratio; 0.075734% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## SearchEngineID (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `3 / 11 / 15`
- Page cardinality per row group min/median/max of mins: `3 / 11 / 15`; of maxes: `3 / 11 / 15`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/searchengineid_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/searchengineid_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/searchengineid_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/searchengineid_value_length.svg)


Compressed overall:
1. 76,166 B (74.38 KiB) compressed - `zstd-3` + `rle-dict`; 167,349 B (163.43 KiB) encoded; 52.583174x post-compression ratio; 33.409920% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 101,598 B (99.22 KiB) compressed - `zstd-3` + `plain`; 4,003,584 B (3.82 MiB) encoded; 39.420559x post-compression ratio; 0.014764% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
3. 104,349 B (101.90 KiB) compressed - `snappy` + `rle-dict`; 167,279 B (163.36 KiB) encoded; 38.381297x post-compression ratio; 180.958131% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
4. 293,059 B (286.19 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 13.666361x post-compression ratio; 0.040265% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 76,166 B (74.38 KiB) compressed - `rle-dict`; 167,349 B (163.43 KiB) encoded; 52.583174x post-compression ratio; 33.409920% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 101,598 B (99.22 KiB) compressed - `plain`; 4,003,584 B (3.82 MiB) encoded; 39.420559x post-compression ratio; 0.014764% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`

Snappy:
1. 104,349 B (101.90 KiB) compressed - `rle-dict`; 167,279 B (163.36 KiB) encoded; 38.381297x post-compression ratio; 180.958131% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 293,059 B (286.19 KiB) compressed - `plain`; 4,003,713 B (3.82 MiB) encoded; 13.666361x post-compression ratio; 0.040265% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## SearchPhrase (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `7 / 421 / 558`
- Page cardinality per row group min/median/max of mins: `7 / 421 / 558`; of maxes: `7 / 421 / 558`
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `37 / 148 / 1,939`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/searchphrase_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/searchphrase_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/searchphrase_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/searchphrase_value_length.svg)


Compressed overall:
1. 635,371 B (620.48 KiB) compressed - `zstd-3` + `rle-dict`; 1,641,580 B (1.57 MiB) encoded; 11.863795x post-compression ratio; 13.308602% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 719,279 B (702.42 KiB) compressed - `zstd-3` + `plain`; 7,535,016 B (7.19 MiB) encoded; 10.479815x post-compression ratio; 0.090507% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 784,272 B (765.89 KiB) compressed - `snappy` + `rle-dict`; 1,643,009 B (1.57 MiB) encoded; 9.611348x post-compression ratio; 40.008951% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
4. 811,698 B (792.67 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 4,222,053 B (4.03 MiB) encoded; 9.286596x post-compression ratio; -11.305683% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
5. 825,797 B (806.44 KiB) compressed - `zstd-3` + `delta-byte-array`; 2,957,835 B (2.82 MiB) encoded; 9.128044x post-compression ratio; -12.819979% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 635,371 B (620.48 KiB) compressed - `rle-dict`; 1,641,580 B (1.57 MiB) encoded; 11.863795x post-compression ratio; 13.308602% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 719,279 B (702.42 KiB) compressed - `plain`; 7,535,016 B (7.19 MiB) encoded; 10.479815x post-compression ratio; 0.090507% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 811,698 B (792.67 KiB) compressed - `delta-length-byte-array`; 4,222,053 B (4.03 MiB) encoded; 9.286596x post-compression ratio; -11.305683% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
4. 825,797 B (806.44 KiB) compressed - `delta-byte-array`; 2,957,835 B (2.82 MiB) encoded; 9.128044x post-compression ratio; -12.819979% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`

Snappy:
1. 784,272 B (765.89 KiB) compressed - `rle-dict`; 1,643,009 B (1.57 MiB) encoded; 9.611348x post-compression ratio; 40.008951% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 992,328 B (969.07 KiB) compressed - `delta-byte-array`; 2,964,725 B (2.83 MiB) encoded; 7.596189x post-compression ratio; 10.654038% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
3. 1,012,875 B (989.14 KiB) compressed - `delta-length-byte-array`; 4,223,716 B (4.03 MiB) encoded; 7.442094x post-compression ratio; 8.409330% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
4. 1,096,892 B (1.05 MiB) compressed - `plain`; 7,535,584 B (7.19 MiB) encoded; 6.872063x post-compression ratio; 0.105662% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-rle-dict-ts-plain`

## SendTiming (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 989`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 989`; of maxes: `1 / 1 / 989`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/sendtiming_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/sendtiming_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/sendtiming_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/sendtiming_value_length.svg)


Compressed overall:
1. 61,436 B (60.00 KiB) compressed - `zstd-3` + `plain`; 4,003,541 B (3.82 MiB) encoded; 65.190670x post-compression ratio; 0.013022% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
2. 76,243 B (74.46 KiB) compressed - `zstd-3` + `rle-dict`; 120,257 B (117.44 KiB) encoded; 52.530121x post-compression ratio; -19.410307% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
3. 92,894 B (90.72 KiB) compressed - `snappy` + `rle-dict`; 119,257 B (116.46 KiB) encoded; 43.114238x post-compression ratio; 189.928305% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`
4. 268,732 B (262.43 KiB) compressed - `snappy` + `plain`; 4,003,775 B (3.82 MiB) encoded; 14.903525x post-compression ratio; 0.221038% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`

ZSTD:
1. 61,436 B (60.00 KiB) compressed - `plain`; 4,003,541 B (3.82 MiB) encoded; 65.190670x post-compression ratio; 0.013022% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
2. 76,243 B (74.46 KiB) compressed - `rle-dict`; 120,257 B (117.44 KiB) encoded; 52.530121x post-compression ratio; -19.410307% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`

Snappy:
1. 92,894 B (90.72 KiB) compressed - `rle-dict`; 119,257 B (116.46 KiB) encoded; 43.114238x post-compression ratio; 189.928305% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`
2. 268,732 B (262.43 KiB) compressed - `plain`; 4,003,775 B (3.82 MiB) encoded; 14.903525x post-compression ratio; 0.221038% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`

## Sex (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `3 / 3 / 3`
- Page cardinality per row group min/median/max of mins: `3 / 3 / 3`; of maxes: `3 / 3 / 3`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/sex_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/sex_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/sex_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/sex_value_length.svg)


Compressed overall:
1. 80,090 B (78.21 KiB) compressed - `zstd-3` + `rle-dict`; 133,318 B (130.19 KiB) encoded; 50.006892x post-compression ratio; 34.818329% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 107,558 B (105.04 KiB) compressed - `zstd-3` + `plain`; 4,003,585 B (3.82 MiB) encoded; 37.236207x post-compression ratio; 0.388628% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
3. 112,460 B (109.82 KiB) compressed - `snappy` + `rle-dict`; 133,793 B (130.66 KiB) encoded; 35.613125x post-compression ratio; 169.135693% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 302,530 B (295.44 KiB) compressed - `snappy` + `plain`; 4,003,711 B (3.82 MiB) encoded; 13.238528x post-compression ratio; 0.046276% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 80,090 B (78.21 KiB) compressed - `rle-dict`; 133,318 B (130.19 KiB) encoded; 50.006892x post-compression ratio; 34.818329% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 107,558 B (105.04 KiB) compressed - `plain`; 4,003,585 B (3.82 MiB) encoded; 37.236207x post-compression ratio; 0.388628% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`

Snappy:
1. 112,460 B (109.82 KiB) compressed - `rle-dict`; 133,793 B (130.66 KiB) encoded; 35.613125x post-compression ratio; 169.135693% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
2. 302,530 B (295.44 KiB) compressed - `plain`; 4,003,711 B (3.82 MiB) encoded; 13.238528x post-compression ratio; 0.046276% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## SilverlightVersion1 (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `4 / 5 / 6`
- Page cardinality per row group min/median/max of mins: `4 / 5 / 6`; of maxes: `4 / 5 / 6`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/silverlightversion1_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/silverlightversion1_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/silverlightversion1_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/silverlightversion1_value_length.svg)


Compressed overall:
1. 83,180 B (81.23 KiB) compressed - `zstd-3` + `rle-dict`; 155,641 B (151.99 KiB) encoded; 48.149195x post-compression ratio; 8.920414% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`
2. 89,171 B (87.08 KiB) compressed - `zstd-3` + `plain`; 4,003,648 B (3.82 MiB) encoded; 44.914266x post-compression ratio; 1.602539% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict`
3. 112,565 B (109.93 KiB) compressed - `snappy` + `rle-dict`; 156,160 B (152.50 KiB) encoded; 35.579887x post-compression ratio; 155.224981% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-rle-dict`
4. 287,146 B (280.42 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 13.947783x post-compression ratio; 0.051542% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 83,180 B (81.23 KiB) compressed - `rle-dict`; 155,641 B (151.99 KiB) encoded; 48.149195x post-compression ratio; 8.920414% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`
2. 89,171 B (87.08 KiB) compressed - `plain`; 4,003,648 B (3.82 MiB) encoded; 44.914266x post-compression ratio; 1.602539% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict`

Snappy:
1. 112,565 B (109.93 KiB) compressed - `rle-dict`; 156,160 B (152.50 KiB) encoded; 35.579887x post-compression ratio; 155.224981% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-rle-dict`
2. 287,146 B (280.42 KiB) compressed - `plain`; 4,003,713 B (3.82 MiB) encoded; 13.947783x post-compression ratio; 0.051542% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## SilverlightVersion2 (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `2 / 2 / 2`
- Page cardinality per row group min/median/max of mins: `2 / 2 / 2`; of maxes: `2 / 2 / 2`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/silverlightversion2_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/silverlightversion2_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/silverlightversion2_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/silverlightversion2_value_length.svg)


Compressed overall:
1. 56,055 B (54.74 KiB) compressed - `zstd-3` + `rle-dict`; 88,952 B (86.87 KiB) encoded; 71.448666x post-compression ratio; 32.063152% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
2. 73,703 B (71.98 KiB) compressed - `zstd-3` + `plain`; 4,003,589 B (3.82 MiB) encoded; 54.340461x post-compression ratio; 0.440959% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
3. 76,291 B (74.50 KiB) compressed - `snappy` + `rle-dict`; 89,062 B (86.97 KiB) encoded; 52.497084x post-compression ratio; 253.666881% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
4. 269,661 B (263.34 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 14.852185x post-compression ratio; 0.057480% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 56,055 B (54.74 KiB) compressed - `rle-dict`; 88,952 B (86.87 KiB) encoded; 71.448666x post-compression ratio; 32.063152% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
2. 73,703 B (71.98 KiB) compressed - `plain`; 4,003,589 B (3.82 MiB) encoded; 54.340461x post-compression ratio; 0.440959% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`

Snappy:
1. 76,291 B (74.50 KiB) compressed - `rle-dict`; 89,062 B (86.97 KiB) encoded; 52.497084x post-compression ratio; 253.666881% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 269,661 B (263.34 KiB) compressed - `plain`; 4,003,713 B (3.82 MiB) encoded; 14.852185x post-compression ratio; 0.057480% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## SilverlightVersion3 (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `17 / 22 / 28`
- Page cardinality per row group min/median/max of mins: `17 / 22 / 28`; of maxes: `17 / 22 / 28`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/silverlightversion3_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/silverlightversion3_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/silverlightversion3_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/silverlightversion3_value_length.svg)


Compressed overall:
1. 123,642 B (120.74 KiB) compressed - `zstd-3` + `rle-dict`; 231,069 B (225.65 KiB) encoded; 32.392318x post-compression ratio; 0.115656% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
2. 123,727 B (120.83 KiB) compressed - `zstd-3` + `plain`; 4,003,583 B (3.82 MiB) encoded; 32.370065x post-compression ratio; 0.046877% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
3. 154,287 B (150.67 KiB) compressed - `snappy` + `rle-dict`; 231,303 B (225.88 KiB) encoded; 25.958448x post-compression ratio; 97.640760% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
4. 304,661 B (297.52 KiB) compressed - `snappy` + `plain`; 4,003,715 B (3.82 MiB) encoded; 13.145926x post-compression ratio; 0.089608% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 123,642 B (120.74 KiB) compressed - `rle-dict`; 231,069 B (225.65 KiB) encoded; 32.392318x post-compression ratio; 0.115656% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
2. 123,727 B (120.83 KiB) compressed - `plain`; 4,003,583 B (3.82 MiB) encoded; 32.370065x post-compression ratio; 0.046877% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`

Snappy:
1. 154,287 B (150.67 KiB) compressed - `rle-dict`; 231,303 B (225.88 KiB) encoded; 25.958448x post-compression ratio; 97.640760% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 304,661 B (297.52 KiB) compressed - `plain`; 4,003,715 B (3.82 MiB) encoded; 13.145926x post-compression ratio; 0.089608% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## SilverlightVersion4 (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 3`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 3`; of maxes: `1 / 1 / 3`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/silverlightversion4_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/silverlightversion4_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/silverlightversion4_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/silverlightversion4_value_length.svg)


Compressed overall:
1. 4,357 B (4.25 KiB) compressed - `zstd-3` + `plain`; 4,003,528 B (3.82 MiB) encoded; 919.222860x post-compression ratio; 0.252467% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 5,475 B (5.35 KiB) compressed - `snappy` + `rle-dict`; 5,239 B (5.12 KiB) encoded; 731.516712x post-compression ratio; 3630.648402% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
3. 6,084 B (5.94 KiB) compressed - `zstd-3` + `rle-dict`; 5,058 B (4.94 KiB) encoded; 658.292899x post-compression ratio; -28.205128% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`
4. 204,101 B (199.32 KiB) compressed - `snappy` + `plain`; 4,003,714 B (3.82 MiB) encoded; 19.622902x post-compression ratio; 0.074473% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 4,357 B (4.25 KiB) compressed - `plain`; 4,003,528 B (3.82 MiB) encoded; 919.222860x post-compression ratio; 0.252467% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 6,084 B (5.94 KiB) compressed - `rle-dict`; 5,058 B (4.94 KiB) encoded; 658.292899x post-compression ratio; -28.205128% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`

Snappy:
1. 5,475 B (5.35 KiB) compressed - `rle-dict`; 5,239 B (5.12 KiB) encoded; 731.516712x post-compression ratio; 3630.648402% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
2. 204,101 B (199.32 KiB) compressed - `plain`; 4,003,714 B (3.82 MiB) encoded; 19.622902x post-compression ratio; 0.074473% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## SocialAction (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 0 / 0`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/socialaction_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/socialaction_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/socialaction_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/socialaction_value_length.svg)


Compressed overall:
1. 2,866 B (2.80 KiB) compressed - `zstd-3` + `plain`; 4,002,163 B (3.82 MiB) encoded; 1396.775297x post-compression ratio; 0.209351% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`
2. 3,569 B (3.49 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 41,691 B (40.71 KiB) encoded; 1121.646960x post-compression ratio; -19.529280% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
3. 3,681 B (3.59 KiB) compressed - `zstd-3` + `delta-byte-array`; 81,277 B (79.37 KiB) encoded; 1087.519152x post-compression ratio; -21.977723% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 4,494 B (4.39 KiB) compressed - `snappy` + `rle-dict`; 4,254 B (4.15 KiB) encoded; 890.778371x post-compression ratio; 4411.593235% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
5. 4,840 B (4.73 KiB) compressed - `snappy` + `delta-length-byte-array`; 41,796 B (40.82 KiB) encoded; 827.098760x post-compression ratio; 4089.070248% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`

ZSTD:
1. 2,866 B (2.80 KiB) compressed - `plain`; 4,002,163 B (3.82 MiB) encoded; 1396.775297x post-compression ratio; 0.209351% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`
2. 3,569 B (3.49 KiB) compressed - `delta-length-byte-array`; 41,691 B (40.71 KiB) encoded; 1121.646960x post-compression ratio; -19.529280% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
3. 3,681 B (3.59 KiB) compressed - `delta-byte-array`; 81,277 B (79.37 KiB) encoded; 1087.519152x post-compression ratio; -21.977723% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 5,247 B (5.12 KiB) compressed - `rle-dict`; 4,185 B (4.09 KiB) encoded; 762.942253x post-compression ratio; -45.263960% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain`

Snappy:
1. 4,494 B (4.39 KiB) compressed - `rle-dict`; 4,254 B (4.15 KiB) encoded; 890.778371x post-compression ratio; 4411.593235% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 4,840 B (4.73 KiB) compressed - `delta-length-byte-array`; 41,796 B (40.82 KiB) encoded; 827.098760x post-compression ratio; 4089.070248% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 6,892 B (6.73 KiB) compressed - `delta-byte-array`; 81,417 B (79.51 KiB) encoded; 580.841265x post-compression ratio; 2841.831109% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 202,705 B (197.95 KiB) compressed - `plain`; 4,002,337 B (3.82 MiB) encoded; 19.748689x post-compression ratio; 0.022693% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`

## SocialNetwork (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 0 / 0`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/socialnetwork_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/socialnetwork_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/socialnetwork_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/socialnetwork_value_length.svg)


Compressed overall:
1. 2,866 B (2.80 KiB) compressed - `zstd-3` + `plain`; 4,002,163 B (3.82 MiB) encoded; 1396.775297x post-compression ratio; 0.209351% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`
2. 3,569 B (3.49 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 41,691 B (40.71 KiB) encoded; 1121.646960x post-compression ratio; -19.529280% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
3. 3,681 B (3.59 KiB) compressed - `zstd-3` + `delta-byte-array`; 81,277 B (79.37 KiB) encoded; 1087.519152x post-compression ratio; -21.977723% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 4,494 B (4.39 KiB) compressed - `snappy` + `rle-dict`; 4,254 B (4.15 KiB) encoded; 890.778371x post-compression ratio; 4411.593235% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
5. 4,840 B (4.73 KiB) compressed - `snappy` + `delta-length-byte-array`; 41,796 B (40.82 KiB) encoded; 827.098760x post-compression ratio; 4089.070248% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`

ZSTD:
1. 2,866 B (2.80 KiB) compressed - `plain`; 4,002,163 B (3.82 MiB) encoded; 1396.775297x post-compression ratio; 0.209351% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`
2. 3,569 B (3.49 KiB) compressed - `delta-length-byte-array`; 41,691 B (40.71 KiB) encoded; 1121.646960x post-compression ratio; -19.529280% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
3. 3,681 B (3.59 KiB) compressed - `delta-byte-array`; 81,277 B (79.37 KiB) encoded; 1087.519152x post-compression ratio; -21.977723% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 5,247 B (5.12 KiB) compressed - `rle-dict`; 4,185 B (4.09 KiB) encoded; 762.942253x post-compression ratio; -45.263960% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain`

Snappy:
1. 4,494 B (4.39 KiB) compressed - `rle-dict`; 4,254 B (4.15 KiB) encoded; 890.778371x post-compression ratio; 4411.593235% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 4,840 B (4.73 KiB) compressed - `delta-length-byte-array`; 41,796 B (40.82 KiB) encoded; 827.098760x post-compression ratio; 4089.070248% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 6,892 B (6.73 KiB) compressed - `delta-byte-array`; 81,417 B (79.51 KiB) encoded; 580.841265x post-compression ratio; 2841.831109% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 202,705 B (197.95 KiB) compressed - `plain`; 4,002,337 B (3.82 MiB) encoded; 19.748689x post-compression ratio; 0.022693% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`

## SocialSourceNetworkID (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 2 / 4`
- Page cardinality per row group min/median/max of mins: `1 / 2 / 4`; of maxes: `1 / 2 / 4`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/socialsourcenetworkid_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/socialsourcenetworkid_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/socialsourcenetworkid_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/socialsourcenetworkid_value_length.svg)


Compressed overall:
1. 5,318 B (5.19 KiB) compressed - `zstd-3` + `plain`; 4,003,530 B (3.82 MiB) encoded; 753.112636x post-compression ratio; 0.000000% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`
2. 6,273 B (6.13 KiB) compressed - `snappy` + `rle-dict`; 6,037 B (5.90 KiB) encoded; 638.458951x post-compression ratio; 3161.517615% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
3. 6,912 B (6.75 KiB) compressed - `zstd-3` + `rle-dict`; 5,886 B (5.75 KiB) encoded; 579.434751x post-compression ratio; -23.061343% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
4. 204,456 B (199.66 KiB) compressed - `snappy` + `plain`; 4,003,716 B (3.82 MiB) encoded; 19.588826x post-compression ratio; 0.067985% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 5,318 B (5.19 KiB) compressed - `plain`; 4,003,530 B (3.82 MiB) encoded; 753.112636x post-compression ratio; 0.000000% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`
2. 6,912 B (6.75 KiB) compressed - `rle-dict`; 5,886 B (5.75 KiB) encoded; 579.434751x post-compression ratio; -23.061343% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`

Snappy:
1. 6,273 B (6.13 KiB) compressed - `rle-dict`; 6,037 B (5.90 KiB) encoded; 638.458951x post-compression ratio; 3161.517615% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
2. 204,456 B (199.66 KiB) compressed - `plain`; 4,003,716 B (3.82 MiB) encoded; 19.588826x post-compression ratio; 0.067985% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## SocialSourcePage (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 5`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 5`; of maxes: `1 / 1 / 5`
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 0 / 28`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/socialsourcepage_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/socialsourcepage_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/socialsourcepage_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/socialsourcepage_value_length.svg)


Compressed overall:
1. 4,856 B (4.74 KiB) compressed - `zstd-3` + `plain`; 4,003,989 B (3.82 MiB) encoded; 824.784390x post-compression ratio; 0.000000% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`
2. 5,977 B (5.84 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 46,327 B (45.24 KiB) encoded; 670.094194x post-compression ratio; -18.755228% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
3. 6,055 B (5.91 KiB) compressed - `snappy` + `rle-dict`; 5,822 B (5.69 KiB) encoded; 661.462097x post-compression ratio; 3273.740710% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
4. 6,144 B (6.00 KiB) compressed - `zstd-3` + `delta-byte-array`; 86,534 B (84.51 KiB) encoded; 651.880371x post-compression ratio; -20.963542% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
5. 6,826 B (6.67 KiB) compressed - `zstd-3` + `rle-dict`; 5,764 B (5.63 KiB) encoded; 586.749634x post-compression ratio; -28.860240% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`

ZSTD:
1. 4,856 B (4.74 KiB) compressed - `plain`; 4,003,989 B (3.82 MiB) encoded; 824.784390x post-compression ratio; 0.000000% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`
2. 5,977 B (5.84 KiB) compressed - `delta-length-byte-array`; 46,327 B (45.24 KiB) encoded; 670.094194x post-compression ratio; -18.755228% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
3. 6,144 B (6.00 KiB) compressed - `delta-byte-array`; 86,534 B (84.51 KiB) encoded; 651.880371x post-compression ratio; -20.963542% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
4. 6,826 B (6.67 KiB) compressed - `rle-dict`; 5,764 B (5.63 KiB) encoded; 586.749634x post-compression ratio; -28.860240% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`

Snappy:
1. 6,055 B (5.91 KiB) compressed - `rle-dict`; 5,822 B (5.69 KiB) encoded; 661.462097x post-compression ratio; 3273.740710% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 7,285 B (7.11 KiB) compressed - `delta-length-byte-array`; 46,327 B (45.24 KiB) encoded; 549.780782x post-compression ratio; 2704.118051% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
3. 9,385 B (9.17 KiB) compressed - `delta-byte-array`; 86,587 B (84.56 KiB) encoded; 426.761108x post-compression ratio; 2076.664891% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
4. 204,267 B (199.48 KiB) compressed - `plain`; 4,004,127 B (3.82 MiB) encoded; 19.607440x post-compression ratio; 0.006364% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-plain`

## Title (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `603`
- Row-group cardinality min/median/max: `108 / 2,277 / 2,527`
- Page cardinality per row group min/median/max of mins: `15 / 135 / 258`; of maxes: `84 / 334 / 758`
- Value length per row group min/median/max of mins: `0 / 0 / 33`; of maxes: `143 / 523 / 1,026`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/title_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/title_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/title_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/title_value_length.svg)


Compressed overall:
1. 7,955,232 B (7.59 MiB) compressed - `zstd-3` + `rle-dict`; 28,774,771 B (27.44 MiB) encoded; 17.961193x post-compression ratio; 75.219717% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 9,995,770 B (9.53 MiB) compressed - `snappy` + `rle-dict`; 28,820,237 B (27.49 MiB) encoded; 14.294592x post-compression ratio; 109.150691% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
3. 13,925,809 B (13.28 MiB) compressed - `zstd-3` + `plain`; 142,862,125 B (136.24 MiB) encoded; 10.260478x post-compression ratio; 0.095693% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
4. 13,958,100 B (13.31 MiB) compressed - `zstd-3` + `delta-byte-array`; 64,469,305 B (61.48 MiB) encoded; 10.236741x post-compression ratio; -0.135871% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
5. 14,459,433 B (13.79 MiB) compressed - `zstd-3` + `delta-length-byte-array`; 140,028,024 B (133.54 MiB) encoded; 9.881816x post-compression ratio; -3.598329% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 7,955,232 B (7.59 MiB) compressed - `rle-dict`; 28,774,771 B (27.44 MiB) encoded; 17.961193x post-compression ratio; 75.219717% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 13,925,809 B (13.28 MiB) compressed - `plain`; 142,862,125 B (136.24 MiB) encoded; 10.260478x post-compression ratio; 0.095693% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
3. 13,958,100 B (13.31 MiB) compressed - `delta-byte-array`; 64,469,305 B (61.48 MiB) encoded; 10.236741x post-compression ratio; -0.135871% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 14,459,433 B (13.79 MiB) compressed - `delta-length-byte-array`; 140,028,024 B (133.54 MiB) encoded; 9.881816x post-compression ratio; -3.598329% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`

Snappy:
1. 9,995,770 B (9.53 MiB) compressed - `rle-dict`; 28,820,237 B (27.49 MiB) encoded; 14.294592x post-compression ratio; 109.150691% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 17,101,500 B (16.31 MiB) compressed - `delta-byte-array`; 64,469,312 B (61.48 MiB) encoded; 8.355142x post-compression ratio; 22.247885% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
3. 20,894,961 B (19.93 MiB) compressed - `plain`; 142,870,491 B (136.25 MiB) encoded; 6.838273x post-compression ratio; 0.053893% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`
4. 21,327,428 B (20.34 MiB) compressed - `delta-length-byte-array`; 140,035,564 B (133.55 MiB) encoded; 6.699610x post-compression ratio; -1.974950% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`

## TraficSourceID (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `5 / 7 / 8`
- Page cardinality per row group min/median/max of mins: `5 / 7 / 8`; of maxes: `5 / 7 / 8`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/traficsourceid_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/traficsourceid_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/traficsourceid_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/traficsourceid_value_length.svg)


Compressed overall:
1. 178,371 B (174.19 KiB) compressed - `zstd-3` + `rle-dict`; 289,824 B (283.03 KiB) encoded; 22.453487x post-compression ratio; 61.643989% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-rle-dict`
2. 243,838 B (238.12 KiB) compressed - `snappy` + `rle-dict`; 289,838 B (283.04 KiB) encoded; 16.425049x post-compression ratio; 116.073377% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
3. 288,165 B (281.41 KiB) compressed - `zstd-3` + `plain`; 4,003,588 B (3.82 MiB) encoded; 13.898464x post-compression ratio; 0.055871% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-rle-dict`
4. 526,610 B (514.27 KiB) compressed - `snappy` + `plain`; 4,003,754 B (3.82 MiB) encoded; 7.605346x post-compression ratio; 0.049183% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 178,371 B (174.19 KiB) compressed - `rle-dict`; 289,824 B (283.03 KiB) encoded; 22.453487x post-compression ratio; 61.643989% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-rle-dict`
2. 288,165 B (281.41 KiB) compressed - `plain`; 4,003,588 B (3.82 MiB) encoded; 13.898464x post-compression ratio; 0.055871% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-rle-dict`

Snappy:
1. 243,838 B (238.12 KiB) compressed - `rle-dict`; 289,838 B (283.04 KiB) encoded; 16.425049x post-compression ratio; 116.073377% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
2. 526,610 B (514.27 KiB) compressed - `plain`; 4,003,754 B (3.82 MiB) encoded; 7.605346x post-compression ratio; 0.049183% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## URL (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `396`
- Row-group cardinality min/median/max: `2,860 / 3,100 / 6,974`
- Page cardinality per row group min/median/max of mins: `116 / 646 / 935`; of maxes: `847 / 1,003 / 1,731`
- Value length per row group min/median/max of mins: `0 / 17 / 19`; of maxes: `252 / 483 / 1,991`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/url_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/url_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/url_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/url_value_length.svg)


Compressed overall:
1. 12,678,493 B (12.09 MiB) compressed - `zstd-3` + `rle-dict`; 44,047,032 B (42.01 MiB) encoded; 7.307902x post-compression ratio; 20.731660% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 15,079,693 B (14.38 MiB) compressed - `zstd-3` + `delta-byte-array`; 40,462,865 B (38.59 MiB) encoded; 6.144235x post-compression ratio; 1.507073% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
3. 15,299,932 B (14.59 MiB) compressed - `zstd-3` + `plain`; 92,649,077 B (88.36 MiB) encoded; 6.055791x post-compression ratio; 0.045902% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`
4. 16,029,529 B (15.29 MiB) compressed - `zstd-3` + `delta-length-byte-array`; 89,786,522 B (85.63 MiB) encoded; 5.780156x post-compression ratio; -4.507768% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
5. 16,088,677 B (15.34 MiB) compressed - `snappy` + `rle-dict`; 44,089,715 B (42.05 MiB) encoded; 5.758906x post-compression ratio; 27.297813% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`

ZSTD:
1. 12,678,493 B (12.09 MiB) compressed - `rle-dict`; 44,047,032 B (42.01 MiB) encoded; 7.307902x post-compression ratio; 20.731660% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 15,079,693 B (14.38 MiB) compressed - `delta-byte-array`; 40,462,865 B (38.59 MiB) encoded; 6.144235x post-compression ratio; 1.507073% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
3. 15,299,932 B (14.59 MiB) compressed - `plain`; 92,649,077 B (88.36 MiB) encoded; 6.055791x post-compression ratio; 0.045902% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`
4. 16,029,529 B (15.29 MiB) compressed - `delta-length-byte-array`; 89,786,522 B (85.63 MiB) encoded; 5.780156x post-compression ratio; -4.507768% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`

Snappy:
1. 16,088,677 B (15.34 MiB) compressed - `rle-dict`; 44,089,715 B (42.05 MiB) encoded; 5.758906x post-compression ratio; 27.297813% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 18,051,200 B (17.21 MiB) compressed - `delta-byte-array`; 40,462,221 B (38.59 MiB) encoded; 5.132799x post-compression ratio; 13.458019% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-rle-dict`
3. 20,460,071 B (19.51 MiB) compressed - `plain`; 92,649,058 B (88.36 MiB) encoded; 4.528488x post-compression ratio; 0.100014% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict`
4. 20,780,387 B (19.82 MiB) compressed - `delta-length-byte-array`; 89,783,474 B (85.62 MiB) encoded; 4.458684x post-compression ratio; -1.442962% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`

## URLCategoryID (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `2 / 4 / 78`
- Page cardinality per row group min/median/max of mins: `2 / 4 / 78`; of maxes: `2 / 4 / 78`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/urlcategoryid_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/urlcategoryid_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/urlcategoryid_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/urlcategoryid_value_length.svg)


Compressed overall:
1. 71,056 B (69.39 KiB) compressed - `zstd-3` + `rle-dict`; 161,772 B (157.98 KiB) encoded; 56.364670x post-compression ratio; 23.240824% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
2. 87,467 B (85.42 KiB) compressed - `zstd-3` + `plain`; 4,003,589 B (3.82 MiB) encoded; 45.789246x post-compression ratio; 0.117759% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
3. 91,524 B (89.38 KiB) compressed - `snappy` + `rle-dict`; 161,494 B (157.71 KiB) encoded; 43.759538x post-compression ratio; 197.230235% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
4. 271,863 B (265.49 KiB) compressed - `snappy` + `plain`; 4,003,714 B (3.82 MiB) encoded; 14.731861x post-compression ratio; 0.064003% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 71,056 B (69.39 KiB) compressed - `rle-dict`; 161,772 B (157.98 KiB) encoded; 56.364670x post-compression ratio; 23.240824% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
2. 87,467 B (85.42 KiB) compressed - `plain`; 4,003,589 B (3.82 MiB) encoded; 45.789246x post-compression ratio; 0.117759% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`

Snappy:
1. 91,524 B (89.38 KiB) compressed - `rle-dict`; 161,494 B (157.71 KiB) encoded; 43.759538x post-compression ratio; 197.230235% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 271,863 B (265.49 KiB) compressed - `plain`; 4,003,714 B (3.82 MiB) encoded; 14.731861x post-compression ratio; 0.064003% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## URLHash (int64)

Column shape stats:
- Parquet type: `INT(64,true)`; physical type: `INT64`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `3,001 / 3,292 / 7,420`
- Page cardinality per row group min/median/max of mins: `3,001 / 3,292 / 7,420`; of maxes: `3,001 / 3,292 / 7,420`
- Value length per row group min/median/max of mins: `8 / 8 / 8`; of maxes: `8 / 8 / 8`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/urlhash_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/urlhash_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/urlhash_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/urlhash_value_length.svg)


Compressed overall:
1. 3,580,374 B (3.41 MiB) compressed - `zstd-3` + `plain`; 8,004,552 B (7.63 MiB) encoded; 2.236166x post-compression ratio; 0.002961% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
2. 4,382,430 B (4.18 MiB) compressed - `snappy` + `plain`; 8,004,876 B (7.63 MiB) encoded; 1.826911x post-compression ratio; 0.116146% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain`
3. 4,499,210 B (4.29 MiB) compressed - `snappy` + `rle-dict`; 4,603,439 B (4.39 MiB) encoded; 1.779492x post-compression ratio; -2.482436% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-rle-dict`
4. 4,529,372 B (4.32 MiB) compressed - `zstd-3` + `rle-dict`; 4,619,317 B (4.41 MiB) encoded; 1.767642x post-compression ratio; -20.949748% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 3,580,374 B (3.41 MiB) compressed - `plain`; 8,004,552 B (7.63 MiB) encoded; 2.236166x post-compression ratio; 0.002961% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
2. 4,529,372 B (4.32 MiB) compressed - `rle-dict`; 4,619,317 B (4.41 MiB) encoded; 1.767642x post-compression ratio; -20.949748% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`

Snappy:
1. 4,382,430 B (4.18 MiB) compressed - `plain`; 8,004,876 B (7.63 MiB) encoded; 1.826911x post-compression ratio; 0.116146% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain`
2. 4,499,210 B (4.29 MiB) compressed - `rle-dict`; 4,603,439 B (4.39 MiB) encoded; 1.779492x post-compression ratio; -2.482436% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-rle-dict`

## URLRegionID (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `3 / 5 / 35`
- Page cardinality per row group min/median/max of mins: `3 / 5 / 35`; of maxes: `3 / 5 / 35`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/urlregionid_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/urlregionid_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/urlregionid_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/urlregionid_value_length.svg)


Compressed overall:
1. 40,097 B (39.16 KiB) compressed - `zstd-3` + `rle-dict`; 74,481 B (72.74 KiB) encoded; 99.884006x post-compression ratio; 20.018954% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`
2. 48,056 B (46.93 KiB) compressed - `zstd-3` + `plain`; 4,003,586 B (3.82 MiB) encoded; 83.341289x post-compression ratio; 0.141502% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`
3. 52,198 B (50.97 KiB) compressed - `snappy` + `rle-dict`; 73,898 B (72.17 KiB) encoded; 76.728016x post-compression ratio; 344.913598% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
4. 232,132 B (226.69 KiB) compressed - `snappy` + `plain`; 4,003,712 B (3.82 MiB) encoded; 17.253326x post-compression ratio; 0.044802% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 40,097 B (39.16 KiB) compressed - `rle-dict`; 74,481 B (72.74 KiB) encoded; 99.884006x post-compression ratio; 20.018954% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`
2. 48,056 B (46.93 KiB) compressed - `plain`; 4,003,586 B (3.82 MiB) encoded; 83.341289x post-compression ratio; 0.141502% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`

Snappy:
1. 52,198 B (50.97 KiB) compressed - `rle-dict`; 73,898 B (72.17 KiB) encoded; 76.728016x post-compression ratio; 344.913598% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 232,132 B (226.69 KiB) compressed - `plain`; 4,003,712 B (3.82 MiB) encoded; 17.253326x post-compression ratio; 0.044802% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## UTMCampaign (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 10 / 18`
- Page cardinality per row group min/median/max of mins: `1 / 10 / 18`; of maxes: `1 / 10 / 18`
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 43 / 66`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/utmcampaign_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/utmcampaign_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/utmcampaign_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/utmcampaign_value_length.svg)


Compressed overall:
1. 27,122 B (26.49 KiB) compressed - `zstd-3` + `rle-dict`; 36,919 B (36.05 KiB) encoded; 151.161419x post-compression ratio; 10.087752% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 29,663 B (28.97 KiB) compressed - `zstd-3` + `plain`; 4,097,338 B (3.91 MiB) encoded; 138.212588x post-compression ratio; 0.657385% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 31,688 B (30.95 KiB) compressed - `snappy` + `rle-dict`; 37,205 B (36.33 KiB) encoded; 129.380207x post-compression ratio; 627.139611% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
4. 47,525 B (46.41 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 251,823 B (245.92 KiB) encoded; 86.266176x post-compression ratio; -37.174119% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`
5. 53,982 B (52.72 KiB) compressed - `zstd-3` + `delta-byte-array`; 327,554 B (319.88 KiB) encoded; 75.947538x post-compression ratio; -44.688970% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`

ZSTD:
1. 27,122 B (26.49 KiB) compressed - `rle-dict`; 36,919 B (36.05 KiB) encoded; 151.161419x post-compression ratio; 10.087752% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 29,663 B (28.97 KiB) compressed - `plain`; 4,097,338 B (3.91 MiB) encoded; 138.212588x post-compression ratio; 0.657385% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 47,525 B (46.41 KiB) compressed - `delta-length-byte-array`; 251,823 B (245.92 KiB) encoded; 86.266176x post-compression ratio; -37.174119% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`
4. 53,982 B (52.72 KiB) compressed - `delta-byte-array`; 327,554 B (319.88 KiB) encoded; 75.947538x post-compression ratio; -44.688970% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`

Snappy:
1. 31,688 B (30.95 KiB) compressed - `rle-dict`; 37,205 B (36.33 KiB) encoded; 129.380207x post-compression ratio; 627.139611% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 57,499 B (56.15 KiB) compressed - `delta-length-byte-array`; 253,184 B (247.25 KiB) encoded; 71.302110x post-compression ratio; 300.730447% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-rle-dict`
3. 68,016 B (66.42 KiB) compressed - `delta-byte-array`; 327,678 B (320.00 KiB) encoded; 60.276994x post-compression ratio; 238.767349% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 230,256 B (224.86 KiB) compressed - `plain`; 4,097,843 B (3.91 MiB) encoded; 17.805399x post-compression ratio; 0.069488% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`

## UTMContent (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 3 / 25`
- Page cardinality per row group min/median/max of mins: `1 / 3 / 25`; of maxes: `1 / 3 / 25`
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 7 / 62`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/utmcontent_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/utmcontent_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/utmcontent_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/utmcontent_value_length.svg)


Compressed overall:
1. 13,965 B (13.64 KiB) compressed - `zstd-3` + `plain`; 4,016,538 B (3.83 MiB) encoded; 287.728679x post-compression ratio; 0.007161% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain`
2. 14,626 B (14.28 KiB) compressed - `snappy` + `rle-dict`; 14,851 B (14.50 KiB) encoded; 274.725215x post-compression ratio; 1351.080268% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-rle-dict`
3. 14,962 B (14.61 KiB) compressed - `zstd-3` + `rle-dict`; 14,707 B (14.36 KiB) encoded; 268.555741x post-compression ratio; -6.656864% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
4. 18,085 B (17.66 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 93,895 B (91.69 KiB) encoded; 222.180315x post-compression ratio; -22.775781% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
5. 20,607 B (20.12 KiB) compressed - `zstd-3` + `delta-byte-array`; 149,448 B (145.95 KiB) encoded; 194.988645x post-compression ratio; -32.226913% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 13,965 B (13.64 KiB) compressed - `plain`; 4,016,538 B (3.83 MiB) encoded; 287.728679x post-compression ratio; 0.007161% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain`
2. 14,962 B (14.61 KiB) compressed - `rle-dict`; 14,707 B (14.36 KiB) encoded; 268.555741x post-compression ratio; -6.656864% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
3. 18,085 B (17.66 KiB) compressed - `delta-length-byte-array`; 93,895 B (91.69 KiB) encoded; 222.180315x post-compression ratio; -22.775781% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
4. 20,607 B (20.12 KiB) compressed - `delta-byte-array`; 149,448 B (145.95 KiB) encoded; 194.988645x post-compression ratio; -32.226913% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`

Snappy:
1. 14,626 B (14.28 KiB) compressed - `rle-dict`; 14,851 B (14.50 KiB) encoded; 274.725215x post-compression ratio; 1351.080268% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-rle-dict`
2. 22,079 B (21.56 KiB) compressed - `delta-length-byte-array`; 94,028 B (91.82 KiB) encoded; 181.988813x post-compression ratio; 861.252774% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain`
3. 27,482 B (26.84 KiB) compressed - `delta-byte-array`; 149,158 B (145.66 KiB) encoded; 146.209555x post-compression ratio; 672.269122% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 212,228 B (207.25 KiB) compressed - `plain`; 4,016,774 B (3.83 MiB) encoded; 18.933086x post-compression ratio; 0.003298% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`

## UTMMedium (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 3 / 5`
- Page cardinality per row group min/median/max of mins: `1 / 3 / 5`; of maxes: `1 / 3 / 5`
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 4 / 16`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/utmmedium_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/utmmedium_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/utmmedium_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/utmmedium_value_length.svg)


Compressed overall:
1. 14,496 B (14.16 KiB) compressed - `zstd-3` + `rle-dict`; 16,556 B (16.17 KiB) encoded; 277.401214x post-compression ratio; 17.177152% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-rle-dict`
2. 15,832 B (15.46 KiB) compressed - `snappy` + `rle-dict`; 16,565 B (16.18 KiB) encoded; 253.992420x post-compression ratio; 1254.181405% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
3. 16,944 B (16.55 KiB) compressed - `zstd-3` + `plain`; 4,019,915 B (3.83 MiB) encoded; 237.323418x post-compression ratio; 0.247875% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
4. 21,973 B (21.46 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 132,622 B (129.51 KiB) encoded; 183.006781x post-compression ratio; -22.696036% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`
5. 27,882 B (27.23 KiB) compressed - `zstd-3` + `delta-byte-array`; 213,000 B (208.01 KiB) encoded; 144.222366x post-compression ratio; -39.078976% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 14,496 B (14.16 KiB) compressed - `rle-dict`; 16,556 B (16.17 KiB) encoded; 277.401214x post-compression ratio; 17.177152% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-rle-dict`
2. 16,944 B (16.55 KiB) compressed - `plain`; 4,019,915 B (3.83 MiB) encoded; 237.323418x post-compression ratio; 0.247875% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 21,973 B (21.46 KiB) compressed - `delta-length-byte-array`; 132,622 B (129.51 KiB) encoded; 183.006781x post-compression ratio; -22.696036% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`
4. 27,882 B (27.23 KiB) compressed - `delta-byte-array`; 213,000 B (208.01 KiB) encoded; 144.222366x post-compression ratio; -39.078976% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`

Snappy:
1. 15,832 B (15.46 KiB) compressed - `rle-dict`; 16,565 B (16.18 KiB) encoded; 253.992420x post-compression ratio; 1254.181405% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 28,630 B (27.96 KiB) compressed - `delta-length-byte-array`; 133,259 B (130.14 KiB) encoded; 140.454349x post-compression ratio; 648.843870% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 38,746 B (37.84 KiB) compressed - `delta-byte-array`; 213,611 B (208.60 KiB) encoded; 103.783823x post-compression ratio; 453.331957% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-rle-dict`
4. 214,350 B (209.33 KiB) compressed - `plain`; 4,020,123 B (3.83 MiB) encoded; 18.760009x post-compression ratio; 0.020527% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`

## UTMSource (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 6 / 8`
- Page cardinality per row group min/median/max of mins: `1 / 6 / 8`; of maxes: `1 / 6 / 8`
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 16 / 19`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/utmsource_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/utmsource_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/utmsource_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/utmsource_value_length.svg)


Compressed overall:
1. 18,339 B (17.91 KiB) compressed - `zstd-3` + `rle-dict`; 22,575 B (22.05 KiB) encoded; 221.098261x post-compression ratio; 14.690005% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 20,260 B (19.79 KiB) compressed - `snappy` + `rle-dict`; 22,681 B (22.15 KiB) encoded; 200.134304x post-compression ratio; 981.658440% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
3. 21,004 B (20.51 KiB) compressed - `zstd-3` + `plain`; 4,053,175 B (3.87 MiB) encoded; 193.045182x post-compression ratio; 0.138069% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
4. 33,547 B (32.76 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 187,509 B (183.11 KiB) encoded; 120.866873x post-compression ratio; -37.302888% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
5. 40,012 B (39.07 KiB) compressed - `zstd-3` + `delta-byte-array`; 267,362 B (261.10 KiB) encoded; 101.337624x post-compression ratio; -47.433270% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 18,339 B (17.91 KiB) compressed - `rle-dict`; 22,575 B (22.05 KiB) encoded; 221.098261x post-compression ratio; 14.690005% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 21,004 B (20.51 KiB) compressed - `plain`; 4,053,175 B (3.87 MiB) encoded; 193.045182x post-compression ratio; 0.138069% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 33,547 B (32.76 KiB) compressed - `delta-length-byte-array`; 187,509 B (183.11 KiB) encoded; 120.866873x post-compression ratio; -37.302888% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
4. 40,012 B (39.07 KiB) compressed - `delta-byte-array`; 267,362 B (261.10 KiB) encoded; 101.337624x post-compression ratio; -47.433270% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

Snappy:
1. 20,260 B (19.79 KiB) compressed - `rle-dict`; 22,681 B (22.15 KiB) encoded; 200.134304x post-compression ratio; 981.658440% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 41,244 B (40.28 KiB) compressed - `delta-length-byte-array`; 187,810 B (183.41 KiB) encoded; 98.310566x post-compression ratio; 431.335467% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-rle-dict`
3. 51,686 B (50.47 KiB) compressed - `delta-byte-array`; 267,959 B (261.68 KiB) encoded; 78.449116x post-compression ratio; 323.991023% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 219,121 B (213.99 KiB) compressed - `plain`; 4,053,453 B (3.87 MiB) encoded; 18.504484x post-compression ratio; 0.010496% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`

## UTMTerm (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 2 / 75`
- Page cardinality per row group min/median/max of mins: `1 / 2 / 75`; of maxes: `1 / 2 / 75`
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 16 / 72`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/utmterm_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/utmterm_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/utmterm_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/utmterm_value_length.svg)


Compressed overall:
1. 15,648 B (15.28 KiB) compressed - `zstd-3` + `plain`; 4,032,689 B (3.85 MiB) encoded; 257.827454x post-compression ratio; 0.511247% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
2. 15,724 B (15.36 KiB) compressed - `snappy` + `rle-dict`; 16,633 B (16.24 KiB) encoded; 256.581277x post-compression ratio; 1264.646400% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain`
3. 15,920 B (15.55 KiB) compressed - `zstd-3` + `rle-dict`; 16,405 B (16.02 KiB) encoded; 253.422362x post-compression ratio; -1.206030% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
4. 19,861 B (19.40 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 128,536 B (125.52 KiB) encoded; 203.135995x post-compression ratio; -20.809627% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
5. 22,914 B (22.38 KiB) compressed - `zstd-3` + `delta-byte-array`; 189,668 B (185.22 KiB) encoded; 176.070699x post-compression ratio; -31.360740% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 15,648 B (15.28 KiB) compressed - `plain`; 4,032,689 B (3.85 MiB) encoded; 257.827454x post-compression ratio; 0.511247% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-plain`
2. 15,920 B (15.55 KiB) compressed - `rle-dict`; 16,405 B (16.02 KiB) encoded; 253.422362x post-compression ratio; -1.206030% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
3. 19,861 B (19.40 KiB) compressed - `delta-length-byte-array`; 128,536 B (125.52 KiB) encoded; 203.135995x post-compression ratio; -20.809627% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
4. 22,914 B (22.38 KiB) compressed - `delta-byte-array`; 189,668 B (185.22 KiB) encoded; 176.070699x post-compression ratio; -31.360740% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`

Snappy:
1. 15,724 B (15.36 KiB) compressed - `rle-dict`; 16,633 B (16.24 KiB) encoded; 256.581277x post-compression ratio; 1264.646400% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain`
2. 25,137 B (24.55 KiB) compressed - `delta-length-byte-array`; 129,144 B (126.12 KiB) encoded; 160.499821x post-compression ratio; 753.630107% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
3. 30,896 B (30.17 KiB) compressed - `delta-byte-array`; 189,246 B (184.81 KiB) encoded; 130.582729x post-compression ratio; 594.513853% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 214,577 B (209.55 KiB) compressed - `plain`; 4,032,973 B (3.85 MiB) encoded; 18.802034x post-compression ratio; 0.000000% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain`

## UserAgent (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `11 / 14 / 18`
- Page cardinality per row group min/median/max of mins: `11 / 14 / 18`; of maxes: `11 / 14 / 18`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/useragent_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/useragent_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/useragent_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/useragent_value_length.svg)


Compressed overall:
1. 125,153 B (122.22 KiB) compressed - `zstd-3` + `rle-dict`; 226,591 B (221.28 KiB) encoded; 32.001254x post-compression ratio; 8.003004% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 133,847 B (130.71 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 29.922621x post-compression ratio; 0.987695% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-rle-dict`
3. 165,734 B (161.85 KiB) compressed - `snappy` + `rle-dict`; 224,191 B (218.94 KiB) encoded; 24.165548x post-compression ratio; 106.476643% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-rle-dict`
4. 342,044 B (334.03 KiB) compressed - `snappy` + `plain`; 4,003,714 B (3.82 MiB) encoded; 11.709175x post-compression ratio; 0.046193% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 125,153 B (122.22 KiB) compressed - `rle-dict`; 226,591 B (221.28 KiB) encoded; 32.001254x post-compression ratio; 8.003004% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 133,847 B (130.71 KiB) compressed - `plain`; 4,003,587 B (3.82 MiB) encoded; 29.922621x post-compression ratio; 0.987695% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-rle-dict`

Snappy:
1. 165,734 B (161.85 KiB) compressed - `rle-dict`; 224,191 B (218.94 KiB) encoded; 24.165548x post-compression ratio; 106.476643% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-rle-dict`
2. 342,044 B (334.03 KiB) compressed - `plain`; 4,003,714 B (3.82 MiB) encoded; 11.709175x post-compression ratio; 0.046193% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## UserAgentMajor (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `26 / 29 / 31`
- Page cardinality per row group min/median/max of mins: `26 / 29 / 31`; of maxes: `26 / 29 / 31`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/useragentmajor_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/useragentmajor_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/useragentmajor_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/useragentmajor_value_length.svg)


Compressed overall:
1. 154,186 B (150.57 KiB) compressed - `zstd-3` + `plain`; 4,003,585 B (3.82 MiB) encoded; 25.975445x post-compression ratio; 0.001946% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain`
2. 177,456 B (173.30 KiB) compressed - `zstd-3` + `rle-dict`; 275,570 B (269.11 KiB) encoded; 22.569257x post-compression ratio; -13.111419% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 214,551 B (209.52 KiB) compressed - `snappy` + `rle-dict`; 276,481 B (270.00 KiB) encoded; 18.667123x post-compression ratio; 70.203821% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
4. 364,837 B (356.29 KiB) compressed - `snappy` + `plain`; 4,003,717 B (3.82 MiB) encoded; 10.977642x post-compression ratio; 0.092370% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 154,186 B (150.57 KiB) compressed - `plain`; 4,003,585 B (3.82 MiB) encoded; 25.975445x post-compression ratio; 0.001946% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain`
2. 177,456 B (173.30 KiB) compressed - `rle-dict`; 275,570 B (269.11 KiB) encoded; 22.569257x post-compression ratio; -13.111419% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 214,551 B (209.52 KiB) compressed - `rle-dict`; 276,481 B (270.00 KiB) encoded; 18.667123x post-compression ratio; 70.203821% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
2. 364,837 B (356.29 KiB) compressed - `plain`; 4,003,717 B (3.82 MiB) encoded; 10.977642x post-compression ratio; 0.092370% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## UserAgentMinor (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `19 / 22 / 28`
- Page cardinality per row group min/median/max of mins: `19 / 22 / 28`; of maxes: `19 / 22 / 28`
- Value length per row group min/median/max of mins: `2 / 2 / 2`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/useragentminor_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/useragentminor_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/useragentminor_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/useragentminor_value_length.svg)


Compressed overall:
1. 83,594 B (81.63 KiB) compressed - `zstd-3` + `rle-dict`; 150,305 B (146.78 KiB) encoded; 92.976183x post-compression ratio; 63.947173% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 104,280 B (101.84 KiB) compressed - `snappy` + `rle-dict`; 150,883 B (147.35 KiB) encoded; 74.532518x post-compression ratio; 347.990027% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
3. 136,892 B (133.68 KiB) compressed - `zstd-3` + `plain`; 7,770,889 B (7.41 MiB) encoded; 56.776517x post-compression ratio; 0.115419% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`
4. 150,595 B (147.07 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 4,016,573 B (3.83 MiB) encoded; 51.610286x post-compression ratio; -8.994323% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
5. 178,932 B (174.74 KiB) compressed - `zstd-3` + `delta-byte-array`; 858,898 B (838.77 KiB) encoded; 43.436898x post-compression ratio; -23.406657% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`

ZSTD:
1. 83,594 B (81.63 KiB) compressed - `rle-dict`; 150,305 B (146.78 KiB) encoded; 92.976183x post-compression ratio; 63.947173% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 136,892 B (133.68 KiB) compressed - `plain`; 7,770,889 B (7.41 MiB) encoded; 56.776517x post-compression ratio; 0.115419% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-rle-dict-ts-rle-dict`
3. 150,595 B (147.07 KiB) compressed - `delta-length-byte-array`; 4,016,573 B (3.83 MiB) encoded; 51.610286x post-compression ratio; -8.994323% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
4. 178,932 B (174.74 KiB) compressed - `delta-byte-array`; 858,898 B (838.77 KiB) encoded; 43.436898x post-compression ratio; -23.406657% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`

Snappy:
1. 104,280 B (101.84 KiB) compressed - `rle-dict`; 150,883 B (147.35 KiB) encoded; 74.532518x post-compression ratio; 347.990027% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 235,612 B (230.09 KiB) compressed - `delta-byte-array`; 861,489 B (841.30 KiB) encoded; 32.987501x post-compression ratio; 98.276828% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
3. 327,938 B (320.25 KiB) compressed - `delta-length-byte-array`; 4,017,088 B (3.83 MiB) encoded; 23.700367x post-compression ratio; 42.454976% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
4. 467,072 B (456.12 KiB) compressed - `plain`; 7,771,080 B (7.41 MiB) encoded; 16.640370x post-compression ratio; 0.019697% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`

## UserID (int64)

Column shape stats:
- Parquet type: `INT(64,true)`; physical type: `INT64`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `716 / 898 / 1,805`
- Page cardinality per row group min/median/max of mins: `716 / 898 / 1,805`; of maxes: `716 / 898 / 1,805`
- Value length per row group min/median/max of mins: `8 / 8 / 8`; of maxes: `8 / 8 / 8`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/userid_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/userid_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/userid_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/userid_value_length.svg)


Compressed overall:
1. 617,935 B (603.45 KiB) compressed - `zstd-3` + `plain`; 8,004,547 B (7.63 MiB) encoded; 12.956565x post-compression ratio; 0.058744% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 1,084,731 B (1.03 MiB) compressed - `snappy` + `plain`; 8,004,717 B (7.63 MiB) encoded; 7.380922x post-compression ratio; 0.017239% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
3. 1,101,698 B (1.05 MiB) compressed - `zstd-3` + `rle-dict`; 1,230,940 B (1.17 MiB) encoded; 7.267250x post-compression ratio; -43.877723% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict`
4. 1,120,481 B (1.07 MiB) compressed - `snappy` + `rle-dict`; 1,225,906 B (1.17 MiB) encoded; 7.145427x post-compression ratio; -3.173905% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-rle-dict`

ZSTD:
1. 617,935 B (603.45 KiB) compressed - `plain`; 8,004,547 B (7.63 MiB) encoded; 12.956565x post-compression ratio; 0.058744% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 1,101,698 B (1.05 MiB) compressed - `rle-dict`; 1,230,940 B (1.17 MiB) encoded; 7.267250x post-compression ratio; -43.877723% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict`

Snappy:
1. 1,084,731 B (1.03 MiB) compressed - `plain`; 8,004,717 B (7.63 MiB) encoded; 7.380922x post-compression ratio; 0.017239% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
2. 1,120,481 B (1.07 MiB) compressed - `rle-dict`; 1,225,906 B (1.17 MiB) encoded; 7.145427x post-compression ratio; -3.173905% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-rle-dict`

## WatchID (int64)

Column shape stats:
- Parquet type: `INT(64,true)`; physical type: `INT64`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `9,315 / 11,938 / 14,202`
- Page cardinality per row group min/median/max of mins: `9,315 / 11,938 / 14,202`; of maxes: `9,315 / 11,938 / 14,202`
- Value length per row group min/median/max of mins: `8 / 8 / 8`; of maxes: `8 / 8 / 8`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/watchid_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/watchid_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/watchid_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/watchid_value_length.svg)


Compressed overall:
1. 8,005,132 B (7.63 MiB) compressed - `snappy` + `plain`; 8,004,719 B (7.63 MiB) encoded; 1.000147x post-compression ratio; 0.002174% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
2. 8,005,354 B (7.63 MiB) compressed - `zstd-3` + `plain`; 8,004,556 B (7.63 MiB) encoded; 1.000120x post-compression ratio; 0.000137% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
3. 9,826,284 B (9.37 MiB) compressed - `snappy` + `rle-dict`; 9,825,517 B (9.37 MiB) encoded; 0.814785x post-compression ratio; -18.531705% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
4. 9,835,734 B (9.38 MiB) compressed - `zstd-3` + `rle-dict`; 9,834,361 B (9.38 MiB) encoded; 0.814002x post-compression ratio; -18.609379% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`

ZSTD:
1. 8,005,354 B (7.63 MiB) compressed - `plain`; 8,004,556 B (7.63 MiB) encoded; 1.000120x post-compression ratio; 0.000137% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 9,835,734 B (9.38 MiB) compressed - `rle-dict`; 9,834,361 B (9.38 MiB) encoded; 0.814002x post-compression ratio; -18.609379% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-rle-dict`

Snappy:
1. 8,005,132 B (7.63 MiB) compressed - `plain`; 8,004,719 B (7.63 MiB) encoded; 1.000147x post-compression ratio; 0.002174% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
2. 9,826,284 B (9.37 MiB) compressed - `rle-dict`; 9,825,517 B (9.37 MiB) encoded; 0.814785x post-compression ratio; -18.531705% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`

## WindowClientHeight (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `318 / 435 / 575`
- Page cardinality per row group min/median/max of mins: `318 / 435 / 575`; of maxes: `318 / 435 / 575`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/windowclientheight_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/windowclientheight_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/windowclientheight_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/windowclientheight_value_length.svg)


Compressed overall:
1. 319,342 B (311.86 KiB) compressed - `zstd-3` + `plain`; 4,003,650 B (3.82 MiB) encoded; 12.541573x post-compression ratio; 0.145612% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict`
2. 551,345 B (538.42 KiB) compressed - `snappy` + `plain`; 4,003,759 B (3.82 MiB) encoded; 7.264147x post-compression ratio; 0.099212% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
3. 583,325 B (569.65 KiB) compressed - `zstd-3` + `rle-dict`; 750,345 B (732.76 KiB) encoded; 6.865900x post-compression ratio; -45.175160% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
4. 602,315 B (588.20 KiB) compressed - `snappy` + `rle-dict`; 751,558 B (733.94 KiB) encoded; 6.649429x post-compression ratio; -8.371533% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 319,342 B (311.86 KiB) compressed - `plain`; 4,003,650 B (3.82 MiB) encoded; 12.541573x post-compression ratio; 0.145612% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict`
2. 583,325 B (569.65 KiB) compressed - `rle-dict`; 750,345 B (732.76 KiB) encoded; 6.865900x post-compression ratio; -45.175160% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 551,345 B (538.42 KiB) compressed - `plain`; 4,003,759 B (3.82 MiB) encoded; 7.264147x post-compression ratio; 0.099212% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
2. 602,315 B (588.20 KiB) compressed - `rle-dict`; 751,558 B (733.94 KiB) encoded; 6.649429x post-compression ratio; -8.371533% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`

## WindowClientWidth (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `241 / 306 / 374`
- Page cardinality per row group min/median/max of mins: `241 / 306 / 374`; of maxes: `241 / 306 / 374`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/windowclientwidth_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/windowclientwidth_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/windowclientwidth_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/windowclientwidth_value_length.svg)


Compressed overall:
1. 305,748 B (298.58 KiB) compressed - `zstd-3` + `plain`; 4,003,588 B (3.82 MiB) encoded; 13.099193x post-compression ratio; 0.000981% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain`
2. 471,075 B (460.03 KiB) compressed - `zstd-3` + `rle-dict`; 695,366 B (679.07 KiB) encoded; 8.501941x post-compression ratio; -35.095049% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 513,098 B (501.07 KiB) compressed - `snappy` + `rle-dict`; 696,642 B (680.31 KiB) encoded; 7.805628x post-compression ratio; 0.704739% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
4. 516,176 B (504.08 KiB) compressed - `snappy` + `plain`; 4,003,751 B (3.82 MiB) encoded; 7.759082x post-compression ratio; 0.104228% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 305,748 B (298.58 KiB) compressed - `plain`; 4,003,588 B (3.82 MiB) encoded; 13.099193x post-compression ratio; 0.000981% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain`
2. 471,075 B (460.03 KiB) compressed - `rle-dict`; 695,366 B (679.07 KiB) encoded; 8.501941x post-compression ratio; -35.095049% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`

Snappy:
1. 513,098 B (501.07 KiB) compressed - `rle-dict`; 696,642 B (680.31 KiB) encoded; 7.805628x post-compression ratio; 0.704739% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 516,176 B (504.08 KiB) compressed - `plain`; 4,003,751 B (3.82 MiB) encoded; 7.759082x post-compression ratio; 0.104228% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## WindowName (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 6 / 3,150`
- Page cardinality per row group min/median/max of mins: `1 / 6 / 3,150`; of maxes: `1 / 6 / 3,150`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/windowname_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/windowname_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/windowname_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/windowname_value_length.svg)


Compressed overall:
1. 70,807 B (69.15 KiB) compressed - `zstd-3` + `plain`; 4,003,578 B (3.82 MiB) encoded; 56.562939x post-compression ratio; 0.014123% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`
2. 148,786 B (145.30 KiB) compressed - `zstd-3` + `rle-dict`; 171,618 B (167.60 KiB) encoded; 26.918205x post-compression ratio; -52.403452% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
3. 150,879 B (147.34 KiB) compressed - `snappy` + `rle-dict`; 171,162 B (167.15 KiB) encoded; 26.544794x post-compression ratio; 95.729691% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
4. 295,204 B (288.29 KiB) compressed - `snappy` + `plain`; 4,003,778 B (3.82 MiB) encoded; 13.567065x post-compression ratio; 0.037601% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 70,807 B (69.15 KiB) compressed - `plain`; 4,003,578 B (3.82 MiB) encoded; 56.562939x post-compression ratio; 0.014123% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`
2. 148,786 B (145.30 KiB) compressed - `rle-dict`; 171,618 B (167.60 KiB) encoded; 26.918205x post-compression ratio; -52.403452% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`

Snappy:
1. 150,879 B (147.34 KiB) compressed - `rle-dict`; 171,162 B (167.15 KiB) encoded; 26.544794x post-compression ratio; 95.729691% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 295,204 B (288.29 KiB) compressed - `plain`; 4,003,778 B (3.82 MiB) encoded; 13.567065x post-compression ratio; 0.037601% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain`

## WithHash (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/withhash_row_group_cardinality.svg)

![Page cardinality min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/withhash_page_cardinality.svg)

![Page min/max distribution](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/withhash_page_bounds.svg)

![Value length min/max per row group](/Users/arjunnair/Desktop/parquet_compress/encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/results/column_shape_stats/images/withhash_value_length.svg)


Compressed overall:
1. 4,228 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; 0.283822% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 5,360 B (5.23 KiB) compressed - `snappy` + `rle-dict`; 5,124 B (5.00 KiB) encoded; 747.211567x post-compression ratio; 3709.981343% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
3. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; -29.037657% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
4. 204,063 B (199.28 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 19.626557x post-compression ratio; 0.074487% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 4,228 B (4.13 KiB) compressed - `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; 0.283822% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
2. 5,975 B (5.83 KiB) compressed - `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; -29.037657% vs plain + zstd-3; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 5,360 B (5.23 KiB) compressed - `rle-dict`; 5,124 B (5.00 KiB) encoded; 747.211567x post-compression ratio; 3709.981343% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 204,063 B (199.28 KiB) compressed - `plain`; 4,003,713 B (3.82 MiB) encoded; 19.626557x post-compression ratio; 0.074487% vs plain + snappy; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

