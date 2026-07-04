# ClickBench Parquet Experiment

- Started: `2026-07-03T23:40:02-04:00`
- Write elapsed: `11.809s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `588938191` (561.66 MiB)
- Compressed column data bytes after codec compression: `96980479` (92.49 MiB)
- Parquet file bytes: `97879085` (93.34 MiB)
- Physical/encoded ratio: `1.210x`
- Encoded/compressed-data ratio: `6.073x`
- Physical/compressed-data ratio: `7.346x`
- Physical/parquet-file ratio: `7.278x`
- Files: `29`

## Settings

- Compression: `zstd-3`
- Int encoding: `delta-binary-packed`
- String encoding: `plain`
- Date encoding: `rle-dict`
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
- Elapsed: `7.412s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `7988400` (7.62 MiB) | `7988624` (7.62 MiB) | `1.001x` | `1.000x` | `1.001x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `198162` (193.52 KiB) | `67960` (66.37 KiB) | `20.186x` | `2.916x` | `58.858x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:522` | `1000000` | `138409995` (132.00 MiB) | `142873292` (136.25 MiB) | `13946946` (13.30 MiB) | `0.969x` | `10.244x` | `9.924x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43591` (42.57 KiB) | `5614` (5.48 KiB) | `91.762x` | `7.765x` | `712.504x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004558` (7.63 MiB) | `2516678` (2.40 MiB) | `0.999x` | `3.181x` | `3.179x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4952` (4.84 KiB) | `5978` (5.84 KiB) | `807.754x` | `0.828x` | `669.120x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `46276` (45.19 KiB) | `5905` (5.77 KiB) | `86.438x` | `7.837x` | `677.392x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `3744971` (3.57 MiB) | `860643` (840.47 KiB) | `1.068x` | `4.351x` | `4.648x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1485401` (1.42 MiB) | `463745` (452.88 KiB) | `2.693x` | `3.203x` | `8.625x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `4384283` (4.18 MiB) | `761045` (743.21 KiB) | `1.825x` | `5.761x` | `10.512x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4932` (4.82 KiB) | `92.896x` | `8.731x` | `811.030x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `844787` (824.99 KiB) | `216080` (211.02 KiB) | `4.735x` | `3.910x` | `18.512x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `686522` (670.43 KiB) | `236309` (230.77 KiB) | `5.826x` | `2.905x` | `16.927x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:348` | `1000000` | `88562192` (84.46 MiB) | `92648688` (88.36 MiB) | `15298922` (14.59 MiB) | `0.956x` | `6.056x` | `5.789x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:328` | `1000000` | `79583339` (75.90 MiB) | `83646063` (79.77 MiB) | `14216997` (13.56 MiB) | `0.951x` | `5.884x` | `5.598x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `233754` (228.28 KiB) | `98684` (96.37 KiB) | `17.112x` | `2.369x` | `40.533x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1863915` (1.78 MiB) | `520794` (508.59 KiB) | `2.146x` | `3.579x` | `7.681x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1368441` (1.31 MiB) | `336033` (328.16 KiB) | `2.923x` | `4.072x` | `11.904x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1084217` (1.03 MiB) | `202792` (198.04 KiB) | `3.689x` | `5.346x` | `19.725x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `726618` (709.59 KiB) | `99549` (97.22 KiB) | `5.505x` | `7.299x` | `40.181x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1281619` (1.22 MiB) | `409754` (400.15 KiB) | `3.121x` | `3.128x` | `9.762x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1339704` (1.28 MiB) | `358630` (350.22 KiB) | `2.986x` | `3.736x` | `11.154x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `553593` (540.62 KiB) | `136246` (133.05 KiB) | `7.226x` | `4.063x` | `29.359x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `399366` (390.01 KiB) | `92819` (90.64 KiB) | `10.016x` | `4.303x` | `43.095x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `523229` (510.97 KiB) | `199300` (194.63 KiB) | `7.645x` | `2.625x` | `20.070x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3354477` (3.20 MiB) | `7357828` (7.02 MiB) | `246814` (241.03 KiB) | `0.456x` | `29.811x` | `13.591x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `159938` (156.19 KiB) | `42172` (41.18 KiB) | `25.010x` | `3.793x` | `94.850x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `197689` (193.06 KiB) | `36617` (35.76 KiB) | `20.234x` | `5.399x` | `109.239x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `661321` (645.82 KiB) | `259144` (253.07 KiB) | `6.048x` | `2.552x` | `15.435x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3767530` (3.59 MiB) | `7770876` (7.41 MiB) | `136912` (133.70 KiB) | `0.485x` | `56.758x` | `27.518x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `46499` (45.41 KiB) | `7495` (7.32 KiB) | `86.023x` | `6.204x` | `533.689x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `47534` (46.42 KiB) | `8220` (8.03 KiB) | `84.150x` | `5.783x` | `486.618x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `114588` (111.90 KiB) | `32896` (32.12 KiB) | `34.908x` | `3.483x` | `121.595x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `144134` (140.76 KiB) | `39562` (38.63 KiB) | `27.752x` | `3.643x` | `101.107x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `81583` (79.67 KiB) | `4084902` (3.90 MiB) | `22653` (22.12 KiB) | `0.020x` | `180.325x` | `3.601x` | `81583` (79.67 KiB) |
| `Params` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002162` (3.82 MiB) | `2855` (2.79 KiB) | `0.000x` | `1401.808x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `2652935` (2.53 MiB) | `783765` (765.40 KiB) | `1.508x` | `3.385x` | `5.104x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `497161` (485.51 KiB) | `238060` (232.48 KiB) | `8.046x` | `2.088x` | `16.802x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `380012` (371.11 KiB) | `150741` (147.21 KiB) | `10.526x` | `2.521x` | `26.536x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3528017` (3.36 MiB) | `7534860` (7.19 MiB) | `720330` (703.45 KiB) | `0.468x` | `10.460x` | `4.898x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `203178` (198.42 KiB) | `44948` (43.89 KiB) | `19.687x` | `4.520x` | `88.992x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `196606` (192.00 KiB) | `84329` (82.35 KiB) | `20.345x` | `2.331x` | `47.433x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1357303` (1.29 MiB) | `574264` (560.80 KiB) | `2.947x` | `2.364x` | `6.965x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1249475` (1.19 MiB) | `551652` (538.72 KiB) | `3.201x` | `2.265x` | `7.251x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `974183` (951.35 KiB) | `212392` (207.41 KiB) | `4.106x` | `4.587x` | `18.833x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004554` (7.63 MiB) | `2474569` (2.36 MiB) | `0.999x` | `3.235x` | `3.233x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `437184` (426.94 KiB) | `118746` (115.96 KiB) | `9.149x` | `3.682x` | `33.685x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `207589` (202.72 KiB) | `73271` (71.55 KiB) | `19.269x` | `2.833x` | `54.592x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1861430` (1.78 MiB) | `317292` (309.86 KiB) | `2.149x` | `5.867x` | `12.607x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43549` (42.53 KiB) | `5164` (5.04 KiB) | `91.851x` | `8.433x` | `774.593x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:85` | `1000000` | `13587860` (12.96 MiB) | `17594808` (16.78 MiB) | `14657` (14.31 KiB) | `0.772x` | `1200.437x` | `927.056x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `95539` (93.30 KiB) | `12549` (12.25 KiB) | `41.868x` | `7.613x` | `318.750x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `140713` (137.42 KiB) | `44401` (43.36 KiB) | `28.427x` | `3.169x` | `90.088x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `53341` (52.09 KiB) | `10029` (9.79 KiB) | `74.989x` | `5.319x` | `398.843x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `63944` (62.45 KiB) | `16388` (16.00 KiB) | `62.555x` | `3.902x` | `244.081x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `7454179` (7.11 MiB) | `1414357` (1.35 MiB) | `1.073x` | `5.270x` | `5.656x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:137` | `1000000` | `27797671` (26.51 MiB) | `31855165` (30.38 MiB) | `5313092` (5.07 MiB) | `0.873x` | `5.996x` | `5.232x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `3952431` (3.77 MiB) | `3953057` (3.77 MiB) | `1.012x` | `1.000x` | `1.012x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4932` (4.82 KiB) | `92.896x` | `8.731x` | `811.030x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4932` (4.82 KiB) | `92.896x` | `8.731x` | `811.030x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4932` (4.82 KiB) | `92.896x` | `8.731x` | `811.030x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `168447` (164.50 KiB) | `60479` (59.06 KiB) | `23.746x` | `2.785x` | `66.139x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4932` (4.82 KiB) | `92.896x` | `8.731x` | `811.030x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `1000000` (976.56 KiB) | `5002898` (4.77 MiB) | `28596` (27.93 KiB) | `0.200x` | `174.951x` | `34.970x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004557` (7.63 MiB) | `2518724` (2.40 MiB) | `0.999x` | `3.178x` | `3.176x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `742363` (724.96 KiB) | `246102` (240.33 KiB) | `5.388x` | `3.016x` | `16.253x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `252983` (247.05 KiB) | `108814` (106.26 KiB) | `15.811x` | `2.325x` | `36.760x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `328416` (320.72 KiB) | `140103` (136.82 KiB) | `12.180x` | `2.344x` | `28.550x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1587609` (1.51 MiB) | `492761` (481.21 KiB) | `2.520x` | `3.222x` | `8.118x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `841656` (821.93 KiB) | `336854` (328.96 KiB) | `4.753x` | `2.499x` | `11.875x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `3802216` (3.63 MiB) | `927885` (906.14 KiB) | `1.052x` | `4.098x` | `4.311x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `274105` (267.68 KiB) | `121210` (118.37 KiB) | `14.593x` | `2.261x` | `33.001x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43060` (42.05 KiB) | `4990` (4.87 KiB) | `92.894x` | `8.629x` | `801.603x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `116017` (113.30 KiB) | `67060` (65.49 KiB) | `34.478x` | `1.730x` | `59.648x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `2001192` (1.91 MiB) | `6004436` (5.73 MiB) | `32295` (31.54 KiB) | `0.333x` | `185.925x` | `61.966x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3325142` (3.17 MiB) | `7328501` (6.99 MiB) | `122303` (119.44 KiB) | `0.454x` | `59.921x` | `27.188x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002162` (3.82 MiB) | `2855` (2.79 KiB) | `0.000x` | `1401.808x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002162` (3.82 MiB) | `2855` (2.79 KiB) | `0.000x` | `1401.808x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4932` (4.82 KiB) | `92.896x` | `8.731x` | `811.030x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `173781` (169.71 KiB) | `108424` (105.88 KiB) | `23.017x` | `1.603x` | `36.892x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `806822` (787.91 KiB) | `301815` (294.74 KiB) | `4.958x` | `2.673x` | `13.253x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1091327` (1.04 MiB) | `580996` (567.38 KiB) | `3.665x` | `1.878x` | `6.885x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1703601` (1.62 MiB) | `1520463` (1.45 MiB) | `2.348x` | `1.120x` | `2.631x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1501583` (1.43 MiB) | `1306611` (1.25 MiB) | `2.664x` | `1.149x` | `3.061x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `1549539` (1.48 MiB) | `964176` (941.58 KiB) | `2.581x` | `1.607x` | `4.149x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `47411` (46.30 KiB) | `7063` (6.90 KiB) | `84.369x` | `6.713x` | `566.332x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `1024` (1.00 KiB) | `4004073` (3.82 MiB) | `4967` (4.85 KiB) | `0.000x` | `806.135x` | `0.206x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `43971` (42.94 KiB) | `5844` (5.71 KiB) | `181.938x` | `7.524x` | `1368.925x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002162` (3.82 MiB) | `2855` (2.79 KiB) | `0.000x` | `1401.808x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3000000` (2.86 MiB) | `7003302` (6.68 MiB) | `5262` (5.14 KiB) | `0.428x` | `1330.920x` | `570.125x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `43059` (42.05 KiB) | `4932` (4.82 KiB) | `92.896x` | `8.731x` | `811.030x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `58030` (56.67 KiB) | `4062527` (3.87 MiB) | `19264` (18.81 KiB) | `0.014x` | `210.887x` | `3.012x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `22051` (21.53 KiB) | `4025597` (3.84 MiB) | `16523` (16.14 KiB) | `0.005x` | `243.636x` | `1.335x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `25445` (24.85 KiB) | `4029062` (3.84 MiB) | `19020` (18.57 KiB) | `0.006x` | `211.833x` | `1.338x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `48191` (47.06 KiB) | `4051761` (3.86 MiB) | `13405` (13.09 KiB) | `0.012x` | `302.257x` | `3.595x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `49433` (48.27 KiB) | `4053217` (3.87 MiB) | `21081` (20.59 KiB) | `0.012x` | `192.269x` | `2.345x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `16873` (16.48 KiB) | `4019951` (3.83 MiB) | `17160` (16.76 KiB) | `0.004x` | `234.263x` | `0.983x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `91870` (89.72 KiB) | `4097517` (3.91 MiB) | `29858` (29.16 KiB) | `0.022x` | `137.233x` | `3.077x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `13001` (12.70 KiB) | `4016774` (3.83 MiB) | `14228` (13.89 KiB) | `0.003x` | `282.315x` | `0.914x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `28101` (27.44 KiB) | `4032862` (3.85 MiB) | `15947` (15.57 KiB) | `0.007x` | `252.892x` | `1.762x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `45607` (44.54 KiB) | `4048294` (3.86 MiB) | `28405` (27.74 KiB) | `0.011x` | `142.520x` | `1.606x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `92215` (90.05 KiB) | `22369` (21.84 KiB) | `43.377x` | `4.122x` | `178.819x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `8117209` (7.74 MiB) | `4465750` (4.26 MiB) | `0.986x` | `1.818x` | `1.791x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `8119694` (7.74 MiB) | `5222597` (4.98 MiB) | `0.985x` | `1.555x` | `1.532x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `68926` (67.31 KiB) | `8546` (8.35 KiB) | `58.033x` | `8.065x` | `468.055x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00000.parquet`: `35158` rows, `3139143` file bytes (2.99 MiB), `27299118` physical bytes (26.03 MiB), `22974014` encoded bytes (21.91 MiB), `3107181` compressed data bytes (2.96 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00001.parquet`: `35211` rows, `3077283` file bytes (2.93 MiB), `27086260` physical bytes (25.83 MiB), `22688496` encoded bytes (21.64 MiB), `3044970` compressed data bytes (2.90 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00002.parquet`: `35681` rows, `3180658` file bytes (3.03 MiB), `27907170` physical bytes (26.61 MiB), `23474743` encoded bytes (22.39 MiB), `3147499` compressed data bytes (3.00 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00003.parquet`: `35546` rows, `3098192` file bytes (2.95 MiB), `27506526` physical bytes (26.23 MiB), `23076838` encoded bytes (22.01 MiB), `3065924` compressed data bytes (2.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00004.parquet`: `35412` rows, `3086559` file bytes (2.94 MiB), `27562740` physical bytes (26.29 MiB), `23134623` encoded bytes (22.06 MiB), `3053888` compressed data bytes (2.91 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00005.parquet`: `35523` rows, `3068954` file bytes (2.93 MiB), `27315405` physical bytes (26.05 MiB), `22867360` encoded bytes (21.81 MiB), `3036562` compressed data bytes (2.90 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00006.parquet`: `35707` rows, `3093363` file bytes (2.95 MiB), `27522288` physical bytes (26.25 MiB), `23052224` encoded bytes (21.98 MiB), `3060881` compressed data bytes (2.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00007.parquet`: `35977` rows, `3154143` file bytes (3.01 MiB), `27905765` physical bytes (26.61 MiB), `23409605` encoded bytes (22.33 MiB), `3121337` compressed data bytes (2.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00008.parquet`: `36000` rows, `3112916` file bytes (2.97 MiB), `27935072` physical bytes (26.64 MiB), `23410151` encoded bytes (22.33 MiB), `3080320` compressed data bytes (2.94 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00009.parquet`: `35824` rows, `3098995` file bytes (2.96 MiB), `27509577` physical bytes (26.24 MiB), `23019951` encoded bytes (21.95 MiB), `3066846` compressed data bytes (2.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00010.parquet`: `36190` rows, `3111679` file bytes (2.97 MiB), `27896247` physical bytes (26.60 MiB), `23337499` encoded bytes (22.26 MiB), `3079232` compressed data bytes (2.94 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00011.parquet`: `36248` rows, `3091488` file bytes (2.95 MiB), `27948049` physical bytes (26.65 MiB), `23381668` encoded bytes (22.30 MiB), `3059020` compressed data bytes (2.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00012.parquet`: `36073` rows, `3120038` file bytes (2.98 MiB), `27716770` physical bytes (26.43 MiB), `23225965` encoded bytes (22.15 MiB), `3087895` compressed data bytes (2.94 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00013.parquet`: `35634` rows, `3070253` file bytes (2.93 MiB), `27562367` physical bytes (26.29 MiB), `23085718` encoded bytes (22.02 MiB), `3037719` compressed data bytes (2.90 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00014.parquet`: `35378` rows, `3564302` file bytes (3.40 MiB), `24871448` physical bytes (23.72 MiB), `20802537` encoded bytes (19.84 MiB), `3533213` compressed data bytes (3.37 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00015.parquet`: `35151` rows, `3722895` file bytes (3.55 MiB), `24432344` physical bytes (23.30 MiB), `20441366` encoded bytes (19.49 MiB), `3692631` compressed data bytes (3.52 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00016.parquet`: `34955` rows, `3737782` file bytes (3.56 MiB), `23388253` physical bytes (22.30 MiB), `19237113` encoded bytes (18.35 MiB), `3707499` compressed data bytes (3.54 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00017.parquet`: `34607` rows, `3799503` file bytes (3.62 MiB), `22278719` physical bytes (21.25 MiB), `18001117` encoded bytes (17.17 MiB), `3769132` compressed data bytes (3.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00018.parquet`: `34154` rows, `3856415` file bytes (3.68 MiB), `22257675` physical bytes (21.23 MiB), `17977497` encoded bytes (17.14 MiB), `3825669` compressed data bytes (3.65 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00019.parquet`: `34796` rows, `3797370` file bytes (3.62 MiB), `22205747` physical bytes (21.18 MiB), `17888251` encoded bytes (17.06 MiB), `3766624` compressed data bytes (3.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00020.parquet`: `34190` rows, `3843157` file bytes (3.67 MiB), `22057091` physical bytes (21.04 MiB), `17869138` encoded bytes (17.04 MiB), `3812593` compressed data bytes (3.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00021.parquet`: `34750` rows, `3797164` file bytes (3.62 MiB), `22265152` physical bytes (21.23 MiB), `17973224` encoded bytes (17.14 MiB), `3766341` compressed data bytes (3.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00022.parquet`: `34584` rows, `3767799` file bytes (3.59 MiB), `22074067` physical bytes (21.05 MiB), `17815816` encoded bytes (16.99 MiB), `3737535` compressed data bytes (3.56 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00023.parquet`: `34327` rows, `3833588` file bytes (3.66 MiB), `22276377` physical bytes (21.24 MiB), `18038730` encoded bytes (17.20 MiB), `3802868` compressed data bytes (3.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00024.parquet`: `34734` rows, `3818439` file bytes (3.64 MiB), `22189674` physical bytes (21.16 MiB), `17917197` encoded bytes (17.09 MiB), `3788171` compressed data bytes (3.61 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00025.parquet`: `34627` rows, `3789085` file bytes (3.61 MiB), `22096403` physical bytes (21.07 MiB), `17830793` encoded bytes (17.00 MiB), `3758502` compressed data bytes (3.58 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00026.parquet`: `34806` rows, `3732415` file bytes (3.56 MiB), `22115501` physical bytes (21.09 MiB), `17831902` encoded bytes (17.01 MiB), `3701848` compressed data bytes (3.53 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00027.parquet`: `34533` rows, `3752112` file bytes (3.58 MiB), `21959203` physical bytes (20.94 MiB), `17741457` encoded bytes (16.92 MiB), `3721908` compressed data bytes (3.55 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain/part-00028.parquet`: `14224` rows, `1563395` file bytes (1.49 MiB), `9257616` physical bytes (8.83 MiB), `7433198` encoded bytes (7.09 MiB), `1546671` compressed data bytes (1.48 MiB)
