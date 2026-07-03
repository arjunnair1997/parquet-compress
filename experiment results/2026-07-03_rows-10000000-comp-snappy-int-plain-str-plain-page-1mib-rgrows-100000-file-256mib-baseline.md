# ClickBench Parquet Experiment

- Started: `2026-07-03T00:44:58-04:00`
- Write elapsed: `2m10.091s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `clickbench parquet files/10m-plain-snappy-page1m-rg100k-file256m`
- Rows: `10000000`
- Source TSV bytes for rows: `7493330544` (6.98 GiB)
- Parquet bytes: `1450801831` (1.35 GiB)
- Parquet/source ratio: `0.193612`
- Source bytes/row: `749.333`
- Parquet bytes/row: `145.080`
- Files: `7`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `plain`
- Date encoding: `plain`
- Timestamp encoding: `plain`
- Max page size: `1.00 MiB`
- Max row group rows: `100000`
- Max row group size: `0 B`
- Max file size: `256.00 MiB`
- Data page version: `2`

## Verification

- Status: `passed`
- Rows read and compared: `10000000`
- Files read: `7`
- Elapsed: `1m45.64s`
- Source TSV bytes checked: `7493330544` (6.98 GiB)

## Files

- `clickbench parquet files/10m-plain-snappy-page1m-rg100k-file256m/part-00000.parquet`: `1588377` rows, `233544573` bytes (222.73 MiB)
- `clickbench parquet files/10m-plain-snappy-page1m-rg100k-file256m/part-00001.parquet`: `1274846` rows, `238965239` bytes (227.90 MiB)
- `clickbench parquet files/10m-plain-snappy-page1m-rg100k-file256m/part-00002.parquet`: `1294160` rows, `233250752` bytes (222.45 MiB)
- `clickbench parquet files/10m-plain-snappy-page1m-rg100k-file256m/part-00003.parquet`: `1793015` rows, `228858075` bytes (218.26 MiB)
- `clickbench parquet files/10m-plain-snappy-page1m-rg100k-file256m/part-00004.parquet`: `1796845` rows, `227699770` bytes (217.15 MiB)
- `clickbench parquet files/10m-plain-snappy-page1m-rg100k-file256m/part-00005.parquet`: `1874871` rows, `235984310` bytes (225.05 MiB)
- `clickbench parquet files/10m-plain-snappy-page1m-rg100k-file256m/part-00006.parquet`: `377886` rows, `52499112` bytes (50.07 MiB)
