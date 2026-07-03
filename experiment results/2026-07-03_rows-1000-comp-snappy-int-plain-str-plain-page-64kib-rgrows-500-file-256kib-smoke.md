# ClickBench Parquet Experiment

- Started: `2026-07-03T00:35:38-04:00`
- Elapsed: `22ms`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `clickbench parquet files/smoke-1000-plain-snappy`
- Rows: `1000`
- Source TSV bytes for rows: `855850` (835.79 KiB)
- Parquet bytes: `239024` (233.42 KiB)
- Parquet/source ratio: `0.279283`
- Source bytes/row: `855.850`
- Parquet bytes/row: `239.024`
- Files: `4`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `plain`
- Date encoding: `plain`
- Timestamp encoding: `plain`
- Max page size: `64.00 KiB`
- Max row group rows: `500`
- Max row group size: `0 B`
- Max file size: `256.00 KiB`
- Data page version: `2`

## Files

- `clickbench parquet files/smoke-1000-plain-snappy/part-00000.parquet`: `278` rows, `64845` bytes (63.33 KiB)
- `clickbench parquet files/smoke-1000-plain-snappy/part-00001.parquet`: `349` rows, `76934` bytes (75.13 KiB)
- `clickbench parquet files/smoke-1000-plain-snappy/part-00002.parquet`: `360` rows, `72065` bytes (70.38 KiB)
- `clickbench parquet files/smoke-1000-plain-snappy/part-00003.parquet`: `13` rows, `25180` bytes (24.59 KiB)
