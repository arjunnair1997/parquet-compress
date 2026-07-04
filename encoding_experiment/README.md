# Encoding Experiment

Use this space to track experiments whose fixed write settings stay constant while
compression and Parquet encodings vary.

Fixed-settings groups:

- [page-256kib-rgsize-10mib-file-10mib-dictpage-256kib](page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/) warpstream default with 256KiB dictionary page cap
- [page-256kib-rgsize-10mib-file-10mib-dictpage-1mib](page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/) warpstream default with 1MiB dictionary page cap
- [page-256kib-rgsize-10mib-file-10mib](page-256kib-rgsize-10mib-file-10mib/) earlier warpstream default with unlimited dictionary pages
1. The idea is to try all possible combinations of encodings. Observe the compression ratios post encoding
and post compression.
2. The second idea is to figure out the optimal encodings for each of the 105 columns.

Suggested result layout:

```bash
go run . \
	  --rows 10000000 \
	  --max-file-size 10MiB \
	  --max-row-group-size 10MiB \
	  --max-page-size 256KiB \
	  --max-dictionary-page-size 256KiB \
	  --verify
```

The folder name records the fixed data page, dictionary page, row-group, and
file settings. Within each row-count folder, overall markdown files live in
`results/`, per-configuration markdown files live in `results/configs/`,
tabular files live in `tsvs/`, and generated parquet files use `parquet/`.
Generated result filenames record only the varying dimensions: rows,
compression, and encoding choices.
