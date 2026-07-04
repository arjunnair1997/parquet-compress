# ClickBench Parquet Experiment

- Started: `2026-07-03T23:39:04-04:00`
- Write elapsed: `12.227s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `592940094` (565.47 MiB)
- Compressed column data bytes after codec compression: `96973184` (92.48 MiB)
- Parquet file bytes: `97871134` (93.34 MiB)
- Physical/encoded ratio: `1.201x`
- Encoded/compressed-data ratio: `6.114x`
- Physical/compressed-data ratio: `7.346x`
- Physical/parquet-file ratio: `7.279x`
- Files: `29`

## Settings

- Compression: `zstd-3`
- Int encoding: `delta-binary-packed`
- String encoding: `plain`
- Date encoding: `plain`
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
- Elapsed: `7.462s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `7988385` (7.62 MiB) | `7988609` (7.62 MiB) | `1.001x` | `1.000x` | `1.001x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `198000` (193.36 KiB) | `67964` (66.37 KiB) | `20.202x` | `2.913x` | `58.855x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:522` | `1000000` | `138409995` (132.00 MiB) | `142873290` (136.25 MiB) | `13946896` (13.30 MiB) | `0.969x` | `10.244x` | `9.924x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43591` (42.57 KiB) | `5614` (5.48 KiB) | `91.762x` | `7.765x` | `712.504x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004558` (7.63 MiB) | `2516727` (2.40 MiB) | `0.999x` | `3.181x` | `3.179x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `4908` (4.79 KiB) | `0.999x` | `815.715x` | `814.996x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `46275` (45.19 KiB) | `5906` (5.77 KiB) | `86.440x` | `7.835x` | `677.277x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `3745529` (3.57 MiB) | `861430` (841.24 KiB) | `1.068x` | `4.348x` | `4.643x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1484903` (1.42 MiB) | `464184` (453.30 KiB) | `2.694x` | `3.199x` | `8.617x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `4385321` (4.18 MiB) | `760999` (743.16 KiB) | `1.824x` | `5.763x` | `10.512x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4932` (4.82 KiB) | `92.896x` | `8.731x` | `811.030x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `844623` (824.83 KiB) | `216524` (211.45 KiB) | `4.736x` | `3.901x` | `18.474x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `687238` (671.13 KiB) | `236287` (230.75 KiB) | `5.820x` | `2.908x` | `16.929x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:348` | `1000000` | `88562192` (84.46 MiB) | `92648698` (88.36 MiB) | `15299291` (14.59 MiB) | `0.956x` | `6.056x` | `5.789x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:328` | `1000000` | `79583339` (75.90 MiB) | `83646063` (79.77 MiB) | `14217423` (13.56 MiB) | `0.951x` | `5.883x` | `5.598x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `233899` (228.42 KiB) | `98712` (96.40 KiB) | `17.101x` | `2.370x` | `40.522x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1863908` (1.78 MiB) | `520429` (508.23 KiB) | `2.146x` | `3.581x` | `7.686x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1368599` (1.31 MiB) | `336006` (328.13 KiB) | `2.923x` | `4.073x` | `11.905x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1084378` (1.03 MiB) | `202577` (197.83 KiB) | `3.689x` | `5.353x` | `19.746x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `726601` (709.57 KiB) | `99555` (97.22 KiB) | `5.505x` | `7.298x` | `40.179x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1281576` (1.22 MiB) | `409502` (399.90 KiB) | `3.121x` | `3.130x` | `9.768x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1340323` (1.28 MiB) | `358361` (349.96 KiB) | `2.984x` | `3.740x` | `11.162x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `553541` (540.57 KiB) | `136559` (133.36 KiB) | `7.226x` | `4.053x` | `29.291x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `399302` (389.94 KiB) | `92840` (90.66 KiB) | `10.017x` | `4.301x` | `43.085x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `523391` (511.12 KiB) | `199400` (194.73 KiB) | `7.642x` | `2.625x` | `20.060x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3354477` (3.20 MiB) | `7357826` (7.02 MiB) | `246806` (241.02 KiB) | `0.456x` | `29.812x` | `13.592x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `159925` (156.18 KiB) | `42244` (41.25 KiB) | `25.012x` | `3.786x` | `94.688x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `197686` (193.05 KiB) | `36526` (35.67 KiB) | `20.234x` | `5.412x` | `109.511x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `661075` (645.58 KiB) | `259009` (252.94 KiB) | `6.051x` | `2.552x` | `15.443x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3767530` (3.59 MiB) | `7770879` (7.41 MiB) | `136888` (133.68 KiB) | `0.485x` | `56.768x` | `27.523x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `46499` (45.41 KiB) | `7487` (7.31 KiB) | `86.023x` | `6.211x` | `534.259x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `47533` (46.42 KiB) | `8212` (8.02 KiB) | `84.152x` | `5.788x` | `487.092x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `114387` (111.71 KiB) | `32991` (32.22 KiB) | `34.969x` | `3.467x` | `121.245x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `144029` (140.65 KiB) | `39508` (38.58 KiB) | `27.772x` | `3.646x` | `101.245x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `81583` (79.67 KiB) | `4084902` (3.90 MiB) | `22660` (22.13 KiB) | `0.020x` | `180.269x` | `3.600x` | `81583` (79.67 KiB) |
| `Params` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002166` (3.82 MiB) | `2859` (2.79 KiB) | `0.000x` | `1399.848x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `2653923` (2.53 MiB) | `783338` (764.98 KiB) | `1.507x` | `3.388x` | `5.106x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `497268` (485.61 KiB) | `238013` (232.43 KiB) | `8.044x` | `2.089x` | `16.806x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `379390` (370.50 KiB) | `150547` (147.02 KiB) | `10.543x` | `2.520x` | `26.570x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3528017` (3.36 MiB) | `7534858` (7.19 MiB) | `720325` (703.44 KiB) | `0.468x` | `10.460x` | `4.898x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `203106` (198.35 KiB) | `44905` (43.85 KiB) | `19.694x` | `4.523x` | `89.077x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `196714` (192.10 KiB) | `84461` (82.48 KiB) | `20.334x` | `2.329x` | `47.359x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1357309` (1.29 MiB) | `574771` (561.30 KiB) | `2.947x` | `2.361x` | `6.959x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1249617` (1.19 MiB) | `552087` (539.15 KiB) | `3.201x` | `2.263x` | `7.245x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `972668` (949.87 KiB) | `212430` (207.45 KiB) | `4.112x` | `4.579x` | `18.830x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004556` (7.63 MiB) | `2474643` (2.36 MiB) | `0.999x` | `3.235x` | `3.233x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `436861` (426.62 KiB) | `118817` (116.03 KiB) | `9.156x` | `3.677x` | `33.665x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `207538` (202.67 KiB) | `73199` (71.48 KiB) | `19.274x` | `2.835x` | `54.646x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1860368` (1.77 MiB) | `316802` (309.38 KiB) | `2.150x` | `5.872x` | `12.626x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43549` (42.53 KiB) | `5164` (5.04 KiB) | `91.851x` | `8.433x` | `774.593x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:85` | `1000000` | `13587860` (12.96 MiB) | `17594806` (16.78 MiB) | `14654` (14.31 KiB) | `0.772x` | `1200.683x` | `927.246x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `95513` (93.27 KiB) | `12547` (12.25 KiB) | `41.879x` | `7.612x` | `318.801x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `140756` (137.46 KiB) | `44450` (43.41 KiB) | `28.418x` | `3.167x` | `89.989x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `53421` (52.17 KiB) | `10028` (9.79 KiB) | `74.877x` | `5.327x` | `398.883x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `63954` (62.46 KiB) | `16438` (16.05 KiB) | `62.545x` | `3.891x` | `243.339x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `7453534` (7.11 MiB) | `1414138` (1.35 MiB) | `1.073x` | `5.271x` | `5.657x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:137` | `1000000` | `27797671` (26.51 MiB) | `31855021` (30.38 MiB) | `5313297` (5.07 MiB) | `0.873x` | `5.995x` | `5.232x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `3952427` (3.77 MiB) | `3949708` (3.77 MiB) | `1.012x` | `1.001x` | `1.013x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4932` (4.82 KiB) | `92.896x` | `8.731x` | `811.030x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4932` (4.82 KiB) | `92.896x` | `8.731x` | `811.030x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4932` (4.82 KiB) | `92.896x` | `8.731x` | `811.030x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `168464` (164.52 KiB) | `60580` (59.16 KiB) | `23.744x` | `2.781x` | `66.028x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4932` (4.82 KiB) | `92.896x` | `8.731x` | `811.030x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `1000000` (976.56 KiB) | `5002900` (4.77 MiB) | `28827` (28.15 KiB) | `0.200x` | `173.549x` | `34.690x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004556` (7.63 MiB) | `2518831` (2.40 MiB) | `0.999x` | `3.178x` | `3.176x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `742797` (725.39 KiB) | `245748` (239.99 KiB) | `5.385x` | `3.023x` | `16.277x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `253014` (247.08 KiB) | `108797` (106.25 KiB) | `15.809x` | `2.326x` | `36.766x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `328619` (320.92 KiB) | `140203` (136.92 KiB) | `12.172x` | `2.344x` | `28.530x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1588484` (1.51 MiB) | `492104` (480.57 KiB) | `2.518x` | `3.228x` | `8.128x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `841378` (821.66 KiB) | `336636` (328.75 KiB) | `4.754x` | `2.499x` | `11.882x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `3804761` (3.63 MiB) | `928488` (906.73 KiB) | `1.051x` | `4.098x` | `4.308x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `274063` (267.64 KiB) | `121181` (118.34 KiB) | `14.595x` | `2.262x` | `33.008x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4989` (4.87 KiB) | `92.896x` | `8.631x` | `801.764x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `116069` (113.35 KiB) | `67114` (65.54 KiB) | `34.462x` | `1.729x` | `59.600x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `2001192` (1.91 MiB) | `6004436` (5.73 MiB) | `32283` (31.53 KiB) | `0.333x` | `185.994x` | `61.989x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3325142` (3.17 MiB) | `7328501` (6.99 MiB) | `122261` (119.40 KiB) | `0.454x` | `59.941x` | `27.197x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002166` (3.82 MiB) | `2859` (2.79 KiB) | `0.000x` | `1399.848x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002166` (3.82 MiB) | `2859` (2.79 KiB) | `0.000x` | `1399.848x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4932` (4.82 KiB) | `92.896x` | `8.731x` | `811.030x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `173895` (169.82 KiB) | `108388` (105.85 KiB) | `23.002x` | `1.604x` | `36.904x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `805807` (786.92 KiB) | `302074` (294.99 KiB) | `4.964x` | `2.668x` | `13.242x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1091320` (1.04 MiB) | `580696` (567.09 KiB) | `3.665x` | `1.879x` | `6.888x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1704268` (1.63 MiB) | `1521672` (1.45 MiB) | `2.347x` | `1.120x` | `2.629x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1502209` (1.43 MiB) | `1304927` (1.24 MiB) | `2.663x` | `1.151x` | `3.065x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1549598` (1.48 MiB) | `964601` (941.99 KiB) | `2.581x` | `1.606x` | `4.147x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `47408` (46.30 KiB) | `7030` (6.87 KiB) | `84.374x` | `6.744x` | `568.990x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `1024` (1.00 KiB) | `4004076` (3.82 MiB) | `4970` (4.85 KiB) | `0.000x` | `805.649x` | `0.206x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `43971` (42.94 KiB) | `5844` (5.71 KiB) | `181.938x` | `7.524x` | `1368.925x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002166` (3.82 MiB) | `2859` (2.79 KiB) | `0.000x` | `1399.848x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3000000` (2.86 MiB) | `7003302` (6.68 MiB) | `5262` (5.14 KiB) | `0.428x` | `1330.920x` | `570.125x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4932` (4.82 KiB) | `92.896x` | `8.731x` | `811.030x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `58030` (56.67 KiB) | `4062526` (3.87 MiB) | `19262` (18.81 KiB) | `0.014x` | `210.909x` | `3.013x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `22051` (21.53 KiB) | `4025596` (3.84 MiB) | `16516` (16.13 KiB) | `0.005x` | `243.739x` | `1.335x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `25445` (24.85 KiB) | `4029064` (3.84 MiB) | `19023` (18.58 KiB) | `0.006x` | `211.800x` | `1.338x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `48191` (47.06 KiB) | `4051765` (3.86 MiB) | `13409` (13.09 KiB) | `0.012x` | `302.168x` | `3.594x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `49433` (48.27 KiB) | `4053221` (3.87 MiB) | `21084` (20.59 KiB) | `0.012x` | `192.242x` | `2.345x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `16873` (16.48 KiB) | `4019953` (3.83 MiB) | `17161` (16.76 KiB) | `0.004x` | `234.249x` | `0.983x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `91870` (89.72 KiB) | `4097520` (3.91 MiB) | `29862` (29.16 KiB) | `0.022x` | `137.215x` | `3.076x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `13001` (12.70 KiB) | `4016777` (3.83 MiB) | `14231` (13.90 KiB) | `0.003x` | `282.255x` | `0.914x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `28101` (27.44 KiB) | `4032862` (3.85 MiB) | `15940` (15.57 KiB) | `0.007x` | `253.003x` | `1.763x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `45607` (44.54 KiB) | `4048294` (3.86 MiB) | `28402` (27.74 KiB) | `0.011x` | `142.536x` | `1.606x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `92219` (90.06 KiB) | `22369` (21.84 KiB) | `43.375x` | `4.123x` | `178.819x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `8117218` (7.74 MiB) | `4465570` (4.26 MiB) | `0.986x` | `1.818x` | `1.791x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `8119701` (7.74 MiB) | `5218576` (4.98 MiB) | `0.985x` | `1.556x` | `1.533x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `68925` (67.31 KiB) | `8517` (8.32 KiB) | `58.034x` | `8.093x` | `469.649x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00000.parquet`: `35158` rows, `3139083` file bytes (2.99 MiB), `27299118` physical bytes (26.03 MiB), `23114597` encoded bytes (22.04 MiB), `3107144` compressed data bytes (2.96 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00001.parquet`: `35211` rows, `3077223` file bytes (2.93 MiB), `27086260` physical bytes (25.83 MiB), `22829291` encoded bytes (21.77 MiB), `3044933` compressed data bytes (2.90 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00002.parquet`: `35681` rows, `3180597` file bytes (3.03 MiB), `27907170` physical bytes (26.61 MiB), `23617417` encoded bytes (22.52 MiB), `3147461` compressed data bytes (3.00 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00003.parquet`: `35546` rows, `3098131` file bytes (2.95 MiB), `27506526` physical bytes (26.23 MiB), `23218972` encoded bytes (22.14 MiB), `3065886` compressed data bytes (2.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00004.parquet`: `35412` rows, `3086498` file bytes (2.94 MiB), `27562740` physical bytes (26.29 MiB), `23276221` encoded bytes (22.20 MiB), `3053850` compressed data bytes (2.91 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00005.parquet`: `35523` rows, `3068893` file bytes (2.93 MiB), `27315405` physical bytes (26.05 MiB), `23009402` encoded bytes (21.94 MiB), `3036524` compressed data bytes (2.90 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00006.parquet`: `35707` rows, `3093303` file bytes (2.95 MiB), `27522288` physical bytes (26.25 MiB), `23195001` encoded bytes (22.12 MiB), `3060844` compressed data bytes (2.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00007.parquet`: `35977` rows, `3154083` file bytes (3.01 MiB), `27905765` physical bytes (26.61 MiB), `23553464` encoded bytes (22.46 MiB), `3121300` compressed data bytes (2.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00008.parquet`: `36000` rows, `3112855` file bytes (2.97 MiB), `27935072` physical bytes (26.64 MiB), `23554101` encoded bytes (22.46 MiB), `3080282` compressed data bytes (2.94 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00009.parquet`: `35824` rows, `3098937` file bytes (2.96 MiB), `27509577` physical bytes (26.24 MiB), `23163198` encoded bytes (22.09 MiB), `3066811` compressed data bytes (2.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00010.parquet`: `36190` rows, `3111620` file bytes (2.97 MiB), `27896247` physical bytes (26.60 MiB), `23482209` encoded bytes (22.39 MiB), `3079196` compressed data bytes (2.94 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00011.parquet`: `36248` rows, `3091429` file bytes (2.95 MiB), `27948049` physical bytes (26.65 MiB), `23526610` encoded bytes (22.44 MiB), `3058984` compressed data bytes (2.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00012.parquet`: `36073` rows, `3119979` file bytes (2.98 MiB), `27716770` physical bytes (26.43 MiB), `23370207` encoded bytes (22.29 MiB), `3087859` compressed data bytes (2.94 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00013.parquet`: `35635` rows, `3070152` file bytes (2.93 MiB), `27563033` physical bytes (26.29 MiB), `23228654` encoded bytes (22.15 MiB), `3037641` compressed data bytes (2.90 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00014.parquet`: `35378` rows, `3563665` file bytes (3.40 MiB), `24871574` physical bytes (23.72 MiB), `20944122` encoded bytes (19.97 MiB), `3532599` compressed data bytes (3.37 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00015.parquet`: `35154` rows, `3720416` file bytes (3.55 MiB), `24434286` physical bytes (23.30 MiB), `20583361` encoded bytes (19.63 MiB), `3690177` compressed data bytes (3.52 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00016.parquet`: `34957` rows, `3736432` file bytes (3.56 MiB), `23388671` physical bytes (22.31 MiB), `19378752` encoded bytes (18.48 MiB), `3706172` compressed data bytes (3.53 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00017.parquet`: `34604` rows, `3800408` file bytes (3.62 MiB), `22277260` physical bytes (21.25 MiB), `18132792` encoded bytes (17.29 MiB), `3770060` compressed data bytes (3.60 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00018.parquet`: `34154` rows, `3853634` file bytes (3.68 MiB), `22257817` physical bytes (21.23 MiB), `18112038` encoded bytes (17.27 MiB), `3822911` compressed data bytes (3.65 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00019.parquet`: `34797` rows, `3797750` file bytes (3.62 MiB), `22206646` physical bytes (21.18 MiB), `18031457` encoded bytes (17.20 MiB), `3767028` compressed data bytes (3.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00020.parquet`: `34190` rows, `3842683` file bytes (3.66 MiB), `22056640` physical bytes (21.03 MiB), `18011073` encoded bytes (17.18 MiB), `3812142` compressed data bytes (3.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00021.parquet`: `34749` rows, `3797443` file bytes (3.62 MiB), `22265303` physical bytes (21.23 MiB), `18117111` encoded bytes (17.28 MiB), `3766643` compressed data bytes (3.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00022.parquet`: `34585` rows, `3767773` file bytes (3.59 MiB), `22074199` physical bytes (21.05 MiB), `17948966` encoded bytes (17.12 MiB), `3737532` compressed data bytes (3.56 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00023.parquet`: `34326` rows, `3833672` file bytes (3.66 MiB), `22276324` physical bytes (21.24 MiB), `18174406` encoded bytes (17.33 MiB), `3802975` compressed data bytes (3.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00024.parquet`: `34738` rows, `3817354` file bytes (3.64 MiB), `22191471` physical bytes (21.16 MiB), `18055280` encoded bytes (17.22 MiB), `3787108` compressed data bytes (3.61 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00025.parquet`: `34628` rows, `3785314` file bytes (3.61 MiB), `22097195` physical bytes (21.07 MiB), `17971855` encoded bytes (17.14 MiB), `3754754` compressed data bytes (3.58 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00026.parquet`: `34803` rows, `3735129` file bytes (3.56 MiB), `22113726` physical bytes (21.09 MiB), `17968213` encoded bytes (17.14 MiB), `3704583` compressed data bytes (3.53 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00027.parquet`: `34534` rows, `3752275` file bytes (3.58 MiB), `21959353` physical bytes (20.94 MiB), `17882673` encoded bytes (17.05 MiB), `3722095` compressed data bytes (3.55 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00028.parquet`: `14218` rows, `1564403` file bytes (1.49 MiB), `9254139` physical bytes (8.83 MiB), `7488651` encoded bytes (7.14 MiB), `1547690` compressed data bytes (1.48 MiB)
