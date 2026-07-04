# page-256kib-rgsize-10mib-file-10mib-dictpage-1mib

Fixed settings represented by this folder:

- Max page size: `256KiB`
- Max dictionary page size: `1MiB`
- Max row group size: `10MiB`
- Max file size: `10MiB`

The dictionary page setting is passed to parquet-go as `DictionaryMaxBytes`.
When a column dictionary exceeds `1MiB`, later pages in that row group fall
back to plain encoding while earlier dictionary-encoded pages remain as written.

This folder contains one subdirectory per row count, for example `10000_rows/`
and `1000000_rows/`. Each row-count directory contains `results/` for markdown,
`results/configs/` for per-configuration markdown, `tsvs/` for tabular outputs,
and `parquet/` as the temporary parquet output root for that row count.

Run the full encoding/compression matrix:

```bash
go run ./cmd/encoding-matrix --rows 10000
```

The matrix runner covers 96 combinations: three compression settings
(`none`, `snappy`, `zstd`) times integer/date/timestamp encodings (`plain`,
`rle-dict`) times string encodings (`plain`, `rle-dict`, `delta-byte-array`,
`delta-length-byte-array`). It also writes aggregate TSV files ranking
experiments and selecting per-column winners.

Per-column aggregate TSVs include `post_compression_no_encoding_bytes`, copied
from the same column in the all-plain run for that compression setting, and
`post_compression_no_encoding_ratio`, computed against the plain/uncompressed
baseline bytes for that column.

The aggregate column TSVs and rendered Column Winners table also include
`physical_bytes` and `baseline_encoded_bytes`. Column ratios use
`baseline_encoded_bytes` as the denominator; `physical_bytes` is shown for
pre-page-encoding reference and is not the denominator for those ratios.

The Top Encoding Settings table includes `zstd_compression_no_encoding`, copied
from the all-plain ZSTD run for the whole experiment.

By default it runs one worker process. Use `--parallel 4` to run up to four
writer processes at once.

By default generated parquet output directories are deleted after each
successful experiment and a final cleanup pass runs when the matrix finishes.
This keeps disk usage bounded by the active workers.

Use `--keep-output` with `--parallel 1` when you need to retain generated
parquet files for inspection. Parallel runs always delete generated parquet
output, and the runner rejects `--keep-output` when `--parallel` is greater
than `1`.
