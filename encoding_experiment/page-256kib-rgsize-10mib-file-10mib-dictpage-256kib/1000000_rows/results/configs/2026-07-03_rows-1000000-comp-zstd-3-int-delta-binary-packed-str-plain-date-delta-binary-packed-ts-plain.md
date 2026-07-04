# ClickBench Parquet Experiment

- Started: `2026-07-03T23:41:00-04:00`
- Write elapsed: `12.013s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `588992764` (561.71 MiB)
- Compressed column data bytes after codec compression: `96979729` (92.49 MiB)
- Parquet file bytes: `97877625` (93.34 MiB)
- Physical/encoded ratio: `1.210x`
- Encoded/compressed-data ratio: `6.073x`
- Physical/compressed-data ratio: `7.346x`
- Physical/parquet-file ratio: `7.278x`
- Files: `29`

## Settings

- Compression: `zstd-3`
- Int encoding: `delta-binary-packed`
- String encoding: `plain`
- Date encoding: `delta-binary-packed`
- Timestamp encoding: `plain`
- Max page size: `256.00 KiB`
- Max dictionary page size: `256.00 KiB`
- Max row group rows: `0`
- Max row group size: `10.00 MiB`
- Max file size: `10.00 MiB`

## Schema

- Columns: `105`, generated from the built-in ClickBench `hits` column list in source TSV field order.
- Mapping: each input row is split on tabs, and field `N` is written to ClickBench column `N` with the same name.
- All Parquet columns are required.
- String fields are ClickHouse TSV-unescaped before writing.

| ClickBench kind | Parquet column type | Physical value written |
| --- | --- | --- |
| `int16` | `parquet.Int(16)` | `INT32`, signed 16-bit logical type |
| `int32` | `parquet.Int(32)` | `INT32`, signed 32-bit logical type |
| `int64` | `parquet.Int(64)` | `INT64`, signed 64-bit logical type |
| `date` | `parquet.Date()` | `INT32` days since Unix epoch |
| `timestamp_millis` | `parquet.Timestamp(parquet.Millisecond)` | `INT64` milliseconds since Unix epoch |
| `string` | `parquet.String()` | `BYTE_ARRAY` UTF-8 string |

## Verification

- Status: `passed`
- Rows read and compared: `1000000`
- Files read: `29`
- Elapsed: `7.36s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `7988376` (7.62 MiB) | `7988600` (7.62 MiB) | `1.001x` | `1.000x` | `1.001x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `198177` (193.53 KiB) | `67945` (66.35 KiB) | `20.184x` | `2.917x` | `58.871x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:522` | `1000000` | `138409995` (132.00 MiB) | `142873291` (136.25 MiB) | `13947088` (13.30 MiB) | `0.969x` | `10.244x` | `9.924x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43592` (42.57 KiB) | `5615` (5.48 KiB) | `91.760x` | `7.763x` | `712.378x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004556` (7.63 MiB) | `2516691` (2.40 MiB) | `0.999x` | `3.181x` | `3.179x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `50680` (49.49 KiB) | `6336` (6.19 KiB) | `78.927x` | `7.999x` | `631.313x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `46276` (45.19 KiB) | `5913` (5.77 KiB) | `86.438x` | `7.826x` | `676.476x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `3747205` (3.57 MiB) | `860619` (840.45 KiB) | `1.067x` | `4.354x` | `4.648x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1485534` (1.42 MiB) | `463905` (453.03 KiB) | `2.693x` | `3.202x` | `8.622x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `4386803` (4.18 MiB) | `760998` (743.16 KiB) | `1.824x` | `5.765x` | `10.513x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4932` (4.82 KiB) | `92.896x` | `8.731x` | `811.030x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `844517` (824.72 KiB) | `216199` (211.13 KiB) | `4.736x` | `3.906x` | `18.501x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `686754` (670.66 KiB) | `236260` (230.72 KiB) | `5.825x` | `2.907x` | `16.931x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:348` | `1000000` | `88562192` (84.46 MiB) | `92648694` (88.36 MiB) | `15299307` (14.59 MiB) | `0.956x` | `6.056x` | `5.789x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:328` | `1000000` | `79583339` (75.90 MiB) | `83646058` (79.77 MiB) | `14217136` (13.56 MiB) | `0.951x` | `5.883x` | `5.598x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `233763` (228.28 KiB) | `98681` (96.37 KiB) | `17.111x` | `2.369x` | `40.535x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1863974` (1.78 MiB) | `520779` (508.57 KiB) | `2.146x` | `3.579x` | `7.681x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1368810` (1.31 MiB) | `335977` (328.10 KiB) | `2.922x` | `4.074x` | `11.906x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1084022` (1.03 MiB) | `202922` (198.17 KiB) | `3.690x` | `5.342x` | `19.712x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `725962` (708.95 KiB) | `99654` (97.32 KiB) | `5.510x` | `7.285x` | `40.139x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1281954` (1.22 MiB) | `409960` (400.35 KiB) | `3.120x` | `3.127x` | `9.757x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1339957` (1.28 MiB) | `359004` (350.59 KiB) | `2.985x` | `3.732x` | `11.142x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `553732` (540.75 KiB) | `136322` (133.13 KiB) | `7.224x` | `4.062x` | `29.342x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `399413` (390.05 KiB) | `92739` (90.57 KiB) | `10.015x` | `4.307x` | `43.132x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `523347` (511.08 KiB) | `199439` (194.76 KiB) | `7.643x` | `2.624x` | `20.056x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3354477` (3.20 MiB) | `7357828` (7.02 MiB) | `246848` (241.06 KiB) | `0.456x` | `29.807x` | `13.589x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `159936` (156.19 KiB) | `42209` (41.22 KiB) | `25.010x` | `3.789x` | `94.767x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `197682` (193.05 KiB) | `36638` (35.78 KiB) | `20.235x` | `5.396x` | `109.176x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `661505` (646.00 KiB) | `259270` (253.19 KiB) | `6.047x` | `2.551x` | `15.428x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3767530` (3.59 MiB) | `7770877` (7.41 MiB) | `136886` (133.68 KiB) | `0.485x` | `56.769x` | `27.523x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `46499` (45.41 KiB) | `7502` (7.33 KiB) | `86.023x` | `6.198x` | `533.191x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `47535` (46.42 KiB) | `8216` (8.02 KiB) | `84.149x` | `5.786x` | `486.855x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `114583` (111.90 KiB) | `32936` (32.16 KiB) | `34.909x` | `3.479x` | `121.448x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `144083` (140.71 KiB) | `39519` (38.59 KiB) | `27.762x` | `3.646x` | `101.217x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `81583` (79.67 KiB) | `4084902` (3.90 MiB) | `22652` (22.12 KiB) | `0.020x` | `180.333x` | `3.602x` | `81583` (79.67 KiB) |
| `Params` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002161` (3.82 MiB) | `2854` (2.79 KiB) | `0.000x` | `1402.299x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `2653337` (2.53 MiB) | `783477` (765.11 KiB) | `1.508x` | `3.387x` | `5.105x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `497063` (485.41 KiB) | `238143` (232.56 KiB) | `8.047x` | `2.087x` | `16.797x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `379857` (370.95 KiB) | `150786` (147.25 KiB) | `10.530x` | `2.519x` | `26.528x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3528017` (3.36 MiB) | `7534860` (7.19 MiB) | `720327` (703.44 KiB) | `0.468x` | `10.460x` | `4.898x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `203177` (198.42 KiB) | `44964` (43.91 KiB) | `19.687x` | `4.519x` | `88.960x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `196656` (192.05 KiB) | `84365` (82.39 KiB) | `20.340x` | `2.331x` | `47.413x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1357348` (1.29 MiB) | `574343` (560.88 KiB) | `2.947x` | `2.363x` | `6.964x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1250165` (1.19 MiB) | `551886` (538.95 KiB) | `3.200x` | `2.265x` | `7.248x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `973321` (950.51 KiB) | `212470` (207.49 KiB) | `4.110x` | `4.581x` | `18.826x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004555` (7.63 MiB) | `2474601` (2.36 MiB) | `0.999x` | `3.235x` | `3.233x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `437185` (426.94 KiB) | `118729` (115.95 KiB) | `9.149x` | `3.682x` | `33.690x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `207602` (202.74 KiB) | `73285` (71.57 KiB) | `19.268x` | `2.833x` | `54.581x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1861559` (1.78 MiB) | `317541` (310.10 KiB) | `2.149x` | `5.862x` | `12.597x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43548` (42.53 KiB) | `5165` (5.04 KiB) | `91.853x` | `8.431x` | `774.443x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:85` | `1000000` | `13587860` (12.96 MiB) | `17594807` (16.78 MiB) | `14656` (14.31 KiB) | `0.772x` | `1200.519x` | `927.119x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `95548` (93.31 KiB) | `12568` (12.27 KiB) | `41.864x` | `7.602x` | `318.269x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `140689` (137.39 KiB) | `44444` (43.40 KiB) | `28.432x` | `3.166x` | `90.001x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `53356` (52.11 KiB) | `10021` (9.79 KiB) | `74.968x` | `5.324x` | `399.162x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `63947` (62.45 KiB) | `16386` (16.00 KiB) | `62.552x` | `3.903x` | `244.111x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `7454669` (7.11 MiB) | `1414752` (1.35 MiB) | `1.073x` | `5.269x` | `5.655x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:137` | `1000000` | `27797671` (26.51 MiB) | `31855167` (30.38 MiB) | `5313330` (5.07 MiB) | `0.873x` | `5.995x` | `5.232x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `3952426` (3.77 MiB) | `3953052` (3.77 MiB) | `1.012x` | `1.000x` | `1.012x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4932` (4.82 KiB) | `92.896x` | `8.731x` | `811.030x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4932` (4.82 KiB) | `92.896x` | `8.731x` | `811.030x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4932` (4.82 KiB) | `92.896x` | `8.731x` | `811.030x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `168440` (164.49 KiB) | `60481` (59.06 KiB) | `23.747x` | `2.785x` | `66.136x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4932` (4.82 KiB) | `92.896x` | `8.731x` | `811.030x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `1000000` (976.56 KiB) | `5002900` (4.77 MiB) | `28569` (27.90 KiB) | `0.200x` | `175.116x` | `35.003x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004557` (7.63 MiB) | `2518765` (2.40 MiB) | `0.999x` | `3.178x` | `3.176x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `742581` (725.18 KiB) | `245941` (240.18 KiB) | `5.387x` | `3.019x` | `16.264x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `253027` (247.10 KiB) | `108767` (106.22 KiB) | `15.809x` | `2.326x` | `36.776x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `328564` (320.86 KiB) | `140096` (136.81 KiB) | `12.174x` | `2.345x` | `28.552x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1588333` (1.51 MiB) | `492727` (481.18 KiB) | `2.518x` | `3.224x` | `8.118x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `841432` (821.71 KiB) | `336809` (328.92 KiB) | `4.754x` | `2.498x` | `11.876x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `3803036` (3.63 MiB) | `928022` (906.27 KiB) | `1.052x` | `4.098x` | `4.310x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `274257` (267.83 KiB) | `121204` (118.36 KiB) | `14.585x` | `2.263x` | `33.002x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4989` (4.87 KiB) | `92.896x` | `8.631x` | `801.764x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `116036` (113.32 KiB) | `67065` (65.49 KiB) | `34.472x` | `1.730x` | `59.644x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `2001192` (1.91 MiB) | `6004436` (5.73 MiB) | `32294` (31.54 KiB) | `0.333x` | `185.930x` | `61.968x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3325142` (3.17 MiB) | `7328502` (6.99 MiB) | `122313` (119.45 KiB) | `0.454x` | `59.916x` | `27.186x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002161` (3.82 MiB) | `2854` (2.79 KiB) | `0.000x` | `1402.299x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002161` (3.82 MiB) | `2854` (2.79 KiB) | `0.000x` | `1402.299x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4932` (4.82 KiB) | `92.896x` | `8.731x` | `811.030x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `173800` (169.73 KiB) | `108404` (105.86 KiB) | `23.015x` | `1.603x` | `36.899x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `806831` (787.92 KiB) | `301666` (294.60 KiB) | `4.958x` | `2.675x` | `13.260x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1091438` (1.04 MiB) | `580870` (567.26 KiB) | `3.665x` | `1.879x` | `6.886x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1703906` (1.62 MiB) | `1519213` (1.45 MiB) | `2.348x` | `1.122x` | `2.633x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1501947` (1.43 MiB) | `1304658` (1.24 MiB) | `2.663x` | `1.151x` | `3.066x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1549550` (1.48 MiB) | `963977` (941.38 KiB) | `2.581x` | `1.607x` | `4.149x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `47410` (46.30 KiB) | `7063` (6.90 KiB) | `84.370x` | `6.712x` | `566.332x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `1024` (1.00 KiB) | `4004072` (3.82 MiB) | `4966` (4.85 KiB) | `0.000x` | `806.297x` | `0.206x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `43971` (42.94 KiB) | `5844` (5.71 KiB) | `181.938x` | `7.524x` | `1368.925x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002161` (3.82 MiB) | `2854` (2.79 KiB) | `0.000x` | `1402.299x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3000000` (2.86 MiB) | `7003302` (6.68 MiB) | `5262` (5.14 KiB) | `0.428x` | `1330.920x` | `570.125x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4932` (4.82 KiB) | `92.896x` | `8.731x` | `811.030x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `58030` (56.67 KiB) | `4062528` (3.87 MiB) | `19264` (18.81 KiB) | `0.014x` | `210.887x` | `3.012x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `22051` (21.53 KiB) | `4025596` (3.84 MiB) | `16532` (16.14 KiB) | `0.005x` | `243.503x` | `1.334x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `25445` (24.85 KiB) | `4029061` (3.84 MiB) | `19020` (18.57 KiB) | `0.006x` | `211.833x` | `1.338x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `48191` (47.06 KiB) | `4051760` (3.86 MiB) | `13404` (13.09 KiB) | `0.012x` | `302.280x` | `3.595x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `49433` (48.27 KiB) | `4053218` (3.87 MiB) | `21081` (20.59 KiB) | `0.012x` | `192.269x` | `2.345x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `16873` (16.48 KiB) | `4019951` (3.83 MiB) | `17160` (16.76 KiB) | `0.004x` | `234.263x` | `0.983x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `91870` (89.72 KiB) | `4097519` (3.91 MiB) | `29858` (29.16 KiB) | `0.022x` | `137.234x` | `3.077x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `13001` (12.70 KiB) | `4016774` (3.83 MiB) | `14228` (13.89 KiB) | `0.003x` | `282.315x` | `0.914x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `28101` (27.44 KiB) | `4032863` (3.85 MiB) | `15945` (15.57 KiB) | `0.007x` | `252.923x` | `1.762x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `45607` (44.54 KiB) | `4048295` (3.86 MiB) | `28407` (27.74 KiB) | `0.011x` | `142.510x` | `1.605x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `92215` (90.05 KiB) | `22369` (21.84 KiB) | `43.377x` | `4.122x` | `178.819x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `8117202` (7.74 MiB) | `4466644` (4.26 MiB) | `0.986x` | `1.817x` | `1.791x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `8119697` (7.74 MiB) | `5221293` (4.98 MiB) | `0.985x` | `1.555x` | `1.532x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `68922` (67.31 KiB) | `8563` (8.36 KiB) | `58.037x` | `8.049x` | `467.126x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00000.parquet`: `35158` rows, `3139131` file bytes (2.99 MiB), `27299118` physical bytes (26.03 MiB), `22975588` encoded bytes (21.91 MiB), `3107193` compressed data bytes (2.96 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00001.parquet`: `35211` rows, `3077271` file bytes (2.93 MiB), `27086260` physical bytes (25.83 MiB), `22690126` encoded bytes (21.64 MiB), `3044983` compressed data bytes (2.90 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00002.parquet`: `35681` rows, `3180644` file bytes (3.03 MiB), `27907170` physical bytes (26.61 MiB), `23476336` encoded bytes (22.39 MiB), `3147510` compressed data bytes (3.00 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00003.parquet`: `35546` rows, `3098178` file bytes (2.95 MiB), `27506526` physical bytes (26.23 MiB), `23078425` encoded bytes (22.01 MiB), `3065935` compressed data bytes (2.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00004.parquet`: `35412` rows, `3086547` file bytes (2.94 MiB), `27562740` physical bytes (26.29 MiB), `23136150` encoded bytes (22.06 MiB), `3053901` compressed data bytes (2.91 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00005.parquet`: `35523` rows, `3068940` file bytes (2.93 MiB), `27315405` physical bytes (26.05 MiB), `22868892` encoded bytes (21.81 MiB), `3036573` compressed data bytes (2.90 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00006.parquet`: `35707` rows, `3093351` file bytes (2.95 MiB), `27522288` physical bytes (26.25 MiB), `23053874` encoded bytes (21.99 MiB), `3060894` compressed data bytes (2.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00007.parquet`: `35977` rows, `3154136` file bytes (3.01 MiB), `27905765` physical bytes (26.61 MiB), `23411266` encoded bytes (22.33 MiB), `3121355` compressed data bytes (2.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00008.parquet`: `36000` rows, `3112904` file bytes (2.97 MiB), `27935072` physical bytes (26.64 MiB), `23411867` encoded bytes (22.33 MiB), `3080333` compressed data bytes (2.94 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00009.parquet`: `35824` rows, `3098989` file bytes (2.96 MiB), `27509577` physical bytes (26.24 MiB), `23021770` encoded bytes (21.96 MiB), `3066865` compressed data bytes (2.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00010.parquet`: `36190` rows, `3111667` file bytes (2.97 MiB), `27896247` physical bytes (26.60 MiB), `23339113` encoded bytes (22.26 MiB), `3079245` compressed data bytes (2.94 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00011.parquet`: `36248` rows, `3091475` file bytes (2.95 MiB), `27948049` physical bytes (26.65 MiB), `23383338` encoded bytes (22.30 MiB), `3059032` compressed data bytes (2.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00012.parquet`: `36073` rows, `3120024` file bytes (2.98 MiB), `27716770` physical bytes (26.43 MiB), `23227574` encoded bytes (22.15 MiB), `3087906` compressed data bytes (2.94 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00013.parquet`: `35634` rows, `3070241` file bytes (2.93 MiB), `27562367` physical bytes (26.29 MiB), `23087417` encoded bytes (22.02 MiB), `3037732` compressed data bytes (2.90 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00014.parquet`: `35378` rows, `3564290` file bytes (3.40 MiB), `24871448` physical bytes (23.72 MiB), `20804008` encoded bytes (19.84 MiB), `3533226` compressed data bytes (3.37 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00015.parquet`: `35151` rows, `3722889` file bytes (3.55 MiB), `24432344` physical bytes (23.30 MiB), `20443102` encoded bytes (19.50 MiB), `3692650` compressed data bytes (3.52 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00016.parquet`: `34955` rows, `3737768` file bytes (3.56 MiB), `23388253` physical bytes (22.30 MiB), `19238732` encoded bytes (18.35 MiB), `3707510` compressed data bytes (3.54 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00017.parquet`: `34607` rows, `3799489` file bytes (3.62 MiB), `22278719` physical bytes (21.25 MiB), `18002777` encoded bytes (17.17 MiB), `3769143` compressed data bytes (3.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00018.parquet`: `34154` rows, `3856405` file bytes (3.68 MiB), `22257675` physical bytes (21.23 MiB), `17979032` encoded bytes (17.15 MiB), `3825684` compressed data bytes (3.65 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00019.parquet`: `34796` rows, `3797356` file bytes (3.62 MiB), `22205747` physical bytes (21.18 MiB), `17889864` encoded bytes (17.06 MiB), `3766635` compressed data bytes (3.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00020.parquet`: `34190` rows, `3843144` file bytes (3.67 MiB), `22057091` physical bytes (21.04 MiB), `17870727` encoded bytes (17.04 MiB), `3812605` compressed data bytes (3.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00021.parquet`: `34750` rows, `3797144` file bytes (3.62 MiB), `22265152` physical bytes (21.23 MiB), `17974725` encoded bytes (17.14 MiB), `3766346` compressed data bytes (3.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00022.parquet`: `34586` rows, `3767221` file bytes (3.59 MiB), `22075440` physical bytes (21.05 MiB), `17816012` encoded bytes (16.99 MiB), `3736982` compressed data bytes (3.56 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00023.parquet`: `34326` rows, `3833949` file bytes (3.66 MiB), `22276260` physical bytes (21.24 MiB), `18042209` encoded bytes (17.21 MiB), `3803254` compressed data bytes (3.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00024.parquet`: `34737` rows, `3817481` file bytes (3.64 MiB), `22191004` physical bytes (21.16 MiB), `17918371` encoded bytes (17.09 MiB), `3787238` compressed data bytes (3.61 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00025.parquet`: `34627` rows, `3786989` file bytes (3.61 MiB), `22096029` physical bytes (21.07 MiB), `17836595` encoded bytes (17.01 MiB), `3756431` compressed data bytes (3.58 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00026.parquet`: `34804` rows, `3734591` file bytes (3.56 MiB), `22114849` physical bytes (21.09 MiB), `17833604` encoded bytes (17.01 MiB), `3704047` compressed data bytes (3.53 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00027.parquet`: `34535` rows, `3750584` file bytes (3.58 MiB), `21959831` physical bytes (20.94 MiB), `17743980` encoded bytes (16.92 MiB), `3720406` compressed data bytes (3.55 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain/part-00028.parquet`: `14220` rows, `1564827` file bytes (1.49 MiB), `9255428` physical bytes (8.83 MiB), `7437290` encoded bytes (7.09 MiB), `1548115` compressed data bytes (1.48 MiB)
