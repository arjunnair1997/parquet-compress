# page-256kib-rgsize-10mib-file-10mib-dictpage-256kib

Fixed settings represented by this folder:

- Max page size: `256KiB`
- Max dictionary page size: `256KiB`
- Max row group size: `10MiB`
- Max file size: `10MiB`

The dictionary page setting is passed to parquet-go as `DictionaryMaxBytes`.
When a column dictionary exceeds `256KiB`, later pages in that row group fall
back to plain encoding while earlier dictionary-encoded pages remain as written.

This folder contains one subdirectory per row count, for example `1000000_rows/`.
Each row-count directory contains `results/` for overall markdown,
`results/configs/` for per-configuration markdown, `tsvs/` for tabular outputs,
and `parquet/` as the temporary parquet output root for that row count.

Run the full encoding/compression matrix:

```bash
go run ./cmd/encoding-matrix --rows 1000000 --parallel 4 --max-dictionary-page-size 256KiB
```

The matrix runner covers 96 combinations: three compression settings (`none`,
`snappy`, `zstd`) times integer/date/timestamp encodings (`plain`, `rle-dict`)
times string encodings (`plain`, `rle-dict`, `delta-byte-array`,
`delta-length-byte-array`). It also writes aggregate TSV files ranking
experiments and selecting per-column winners.

By default generated parquet output directories are deleted after each
successful experiment and a final cleanup pass runs when the matrix finishes.
This keeps disk usage bounded by the active workers.
