# ClickBench Parquet Experiment

- Started: `2026-07-03T23:39:23-04:00`
- Write elapsed: `12.301s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `590842555` (563.47 MiB)
- Compressed column data bytes after codec compression: `101538366` (96.83 MiB)
- Parquet file bytes: `102450591` (97.70 MiB)
- Physical/encoded ratio: `1.206x`
- Encoded/compressed-data ratio: `5.819x`
- Physical/compressed-data ratio: `7.016x`
- Physical/parquet-file ratio: `6.954x`
- Files: `29`

## Settings

- Compression: `zstd-3`
- Int encoding: `delta-binary-packed`
- String encoding: `plain`
- Date encoding: `plain`
- Timestamp encoding: `rle-dict`
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
- Elapsed: `7.386s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `8000000` (7.63 MiB) | `7989713` (7.62 MiB) | `7990494` (7.62 MiB) | `1.001x` | `1.000x` | `1.001x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `197832` (193.20 KiB) | `68004` (66.41 KiB) | `20.219x` | `2.909x` | `58.820x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:512` | `1000000` | `138409995` (132.00 MiB) | `142873220` (136.25 MiB) | `13947224` (13.30 MiB) | `0.969x` | `10.244x` | `9.924x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `43674` (42.65 KiB) | `5710` (5.58 KiB) | `91.588x` | `7.649x` | `700.525x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:58, DICTIONARY_PAGE/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `7352553` (7.01 MiB) | `4030794` (3.84 MiB) | `1.088x` | `1.824x` | `1.985x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003591` (3.82 MiB) | `4983` (4.87 KiB) | `0.999x` | `803.450x` | `802.729x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `46469` (45.38 KiB) | `5964` (5.82 KiB) | `86.079x` | `7.792x` | `670.691x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `3746807` (3.57 MiB) | `859655` (839.51 KiB) | `1.068x` | `4.359x` | `4.653x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1484498` (1.42 MiB) | `464431` (453.55 KiB) | `2.695x` | `3.196x` | `8.613x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `8000000` (7.63 MiB) | `4392093` (4.19 MiB) | `760398` (742.58 KiB) | `1.821x` | `5.776x` | `10.521x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `43122` (42.11 KiB) | `5009` (4.89 KiB) | `92.760x` | `8.609x` | `798.563x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `843273` (823.51 KiB) | `216053` (210.99 KiB) | `4.743x` | `3.903x` | `18.514x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `687026` (670.92 KiB) | `236030` (230.50 KiB) | `5.822x` | `2.911x` | `16.947x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:336` | `1000000` | `88562192` (84.46 MiB) | `92648956` (88.36 MiB) | `15301341` (14.59 MiB) | `0.956x` | `6.055x` | `5.788x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:324` | `1000000` | `79583339` (75.90 MiB) | `83647149` (79.77 MiB) | `14213189` (13.55 MiB) | `0.951x` | `5.885x` | `5.599x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `233983` (228.50 KiB) | `98883` (96.57 KiB) | `17.095x` | `2.366x` | `40.452x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1863337` (1.78 MiB) | `521907` (509.67 KiB) | `2.147x` | `3.570x` | `7.664x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1366550` (1.30 MiB) | `336961` (329.06 KiB) | `2.927x` | `4.056x` | `11.871x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1081336` (1.03 MiB) | `203193` (198.43 KiB) | `3.699x` | `5.322x` | `19.686x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `721822` (704.90 KiB) | `99972` (97.63 KiB) | `5.542x` | `7.220x` | `40.011x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1282044` (1.22 MiB) | `410279` (400.66 KiB) | `3.120x` | `3.125x` | `9.749x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1340966` (1.28 MiB) | `358949` (350.54 KiB) | `2.983x` | `3.736x` | `11.144x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `552639` (539.69 KiB) | `136576` (133.38 KiB) | `7.238x` | `4.046x` | `29.288x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `398930` (389.58 KiB) | `93314` (91.13 KiB) | `10.027x` | `4.275x` | `42.866x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `523654` (511.38 KiB) | `199515` (194.84 KiB) | `7.639x` | `2.625x` | `20.049x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `3354477` (3.20 MiB) | `7357884` (7.02 MiB) | `247645` (241.84 KiB) | `0.456x` | `29.711x` | `13.546x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `159800` (156.05 KiB) | `42536` (41.54 KiB) | `25.031x` | `3.757x` | `94.038x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `197395` (192.77 KiB) | `36821` (35.96 KiB) | `20.264x` | `5.361x` | `108.634x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `661568` (646.06 KiB) | `259802` (253.71 KiB) | `6.046x` | `2.546x` | `15.396x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `3767530` (3.59 MiB) | `7770933` (7.41 MiB) | `137239` (134.02 KiB) | `0.485x` | `56.623x` | `27.452x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `46567` (45.48 KiB) | `7545` (7.37 KiB) | `85.898x` | `6.172x` | `530.152x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `47605` (46.49 KiB) | `8211` (8.02 KiB) | `84.025x` | `5.798x` | `487.151x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `114774` (112.08 KiB) | `33174` (32.40 KiB) | `34.851x` | `3.460x` | `120.576x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `144999` (141.60 KiB) | `39874` (38.94 KiB) | `27.586x` | `3.636x` | `100.316x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `81583` (79.67 KiB) | `4084957` (3.90 MiB) | `22716` (22.18 KiB) | `0.020x` | `179.827x` | `3.591x` | `81583` (79.67 KiB) |
| `Params` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `0` (0 B) | `4002197` (3.82 MiB) | `2893` (2.83 KiB) | `0.000x` | `1383.407x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `2651766` (2.53 MiB) | `782126` (763.79 KiB) | `1.508x` | `3.390x` | `5.114x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `496161` (484.53 KiB) | `238053` (232.47 KiB) | `8.062x` | `2.084x` | `16.803x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `377666` (368.81 KiB) | `150858` (147.32 KiB) | `10.591x` | `2.503x` | `26.515x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `3528017` (3.36 MiB) | `7535082` (7.19 MiB) | `720822` (703.93 KiB) | `0.468x` | `10.453x` | `4.894x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `202433` (197.69 KiB) | `45280` (44.22 KiB) | `19.760x` | `4.471x` | `88.339x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `196757` (192.15 KiB) | `84588` (82.61 KiB) | `20.330x` | `2.326x` | `47.288x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1358840` (1.30 MiB) | `575275` (561.79 KiB) | `2.944x` | `2.362x` | `6.953x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1249898` (1.19 MiB) | `552674` (539.72 KiB) | `3.200x` | `2.262x` | `7.238x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `975690` (952.82 KiB) | `212564` (207.58 KiB) | `4.100x` | `4.590x` | `18.818x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:58, DICTIONARY_PAGE/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `7216022` (6.88 MiB) | `3979676` (3.80 MiB) | `1.109x` | `1.813x` | `2.010x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `436640` (426.41 KiB) | `119359` (116.56 KiB) | `9.161x` | `3.658x` | `33.512x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `207538` (202.67 KiB) | `73468` (71.75 KiB) | `19.274x` | `2.825x` | `54.445x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1857791` (1.77 MiB) | `317694` (310.25 KiB) | `2.153x` | `5.848x` | `12.591x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `43609` (42.59 KiB) | `5252` (5.13 KiB) | `91.724x` | `8.303x` | `761.615x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:86` | `1000000` | `13587860` (12.96 MiB) | `17594922` (16.78 MiB) | `14733` (14.39 KiB) | `0.772x` | `1194.252x` | `922.274x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `95855` (93.61 KiB) | `12715` (12.42 KiB) | `41.730x` | `7.539x` | `314.589x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `140676` (137.38 KiB) | `44553` (43.51 KiB) | `28.434x` | `3.157x` | `89.781x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `53682` (52.42 KiB) | `10139` (9.90 KiB) | `74.513x` | `5.295x` | `394.516x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `63989` (62.49 KiB) | `16547` (16.16 KiB) | `62.511x` | `3.867x` | `241.736x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `8000000` (7.63 MiB) | `7450724` (7.11 MiB) | `1416643` (1.35 MiB) | `1.074x` | `5.259x` | `5.647x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:138` | `1000000` | `27797671` (26.51 MiB) | `31856003` (30.38 MiB) | `5328657` (5.08 MiB) | `0.873x` | `5.978x` | `5.217x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `3953077` (3.77 MiB) | `3952453` (3.77 MiB) | `1.012x` | `1.000x` | `1.012x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `43122` (42.11 KiB) | `5009` (4.89 KiB) | `92.760x` | `8.609x` | `798.563x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `43122` (42.11 KiB) | `5009` (4.89 KiB) | `92.760x` | `8.609x` | `798.563x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `43122` (42.11 KiB) | `5009` (4.89 KiB) | `92.760x` | `8.609x` | `798.563x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `168465` (164.52 KiB) | `60606` (59.19 KiB) | `23.744x` | `2.780x` | `66.000x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `43122` (42.11 KiB) | `5009` (4.89 KiB) | `92.760x` | `8.609x` | `798.563x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `1000000` (976.56 KiB) | `5002947` (4.77 MiB) | `28467` (27.80 KiB) | `0.200x` | `175.745x` | `35.128x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:58, DICTIONARY_PAGE/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `7349389` (7.01 MiB) | `4033202` (3.85 MiB) | `1.089x` | `1.822x` | `1.984x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `742997` (725.58 KiB) | `247373` (241.58 KiB) | `5.384x` | `3.004x` | `16.170x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `253598` (247.65 KiB) | `109079` (106.52 KiB) | `15.773x` | `2.325x` | `36.671x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `328415` (320.72 KiB) | `140317` (137.03 KiB) | `12.180x` | `2.341x` | `28.507x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1590142` (1.52 MiB) | `492240` (480.70 KiB) | `2.515x` | `3.230x` | `8.126x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `841174` (821.46 KiB) | `337536` (329.62 KiB) | `4.755x` | `2.492x` | `11.851x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `3798142` (3.62 MiB) | `928255` (906.50 KiB) | `1.053x` | `4.092x` | `4.309x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `274616` (268.18 KiB) | `121713` (118.86 KiB) | `14.566x` | `2.256x` | `32.864x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `43124` (42.11 KiB) | `5069` (4.95 KiB) | `92.756x` | `8.507x` | `789.110x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `116172` (113.45 KiB) | `67304` (65.73 KiB) | `34.432x` | `1.726x` | `59.432x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `2001192` (1.91 MiB) | `6004494` (5.73 MiB) | `32503` (31.74 KiB) | `0.333x` | `184.737x` | `61.569x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `3325142` (3.17 MiB) | `7328559` (6.99 MiB) | `122338` (119.47 KiB) | `0.454x` | `59.904x` | `27.180x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `0` (0 B) | `4002197` (3.82 MiB) | `2893` (2.83 KiB) | `0.000x` | `1383.407x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `0` (0 B) | `4002197` (3.82 MiB) | `2893` (2.83 KiB) | `0.000x` | `1383.407x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `43122` (42.11 KiB) | `5009` (4.89 KiB) | `92.760x` | `8.609x` | `798.563x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `173667` (169.60 KiB) | `108282` (105.74 KiB) | `23.033x` | `1.604x` | `36.941x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `808318` (789.37 KiB) | `302428` (295.34 KiB) | `4.949x` | `2.673x` | `13.226x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1092442` (1.04 MiB) | `581737` (568.10 KiB) | `3.662x` | `1.878x` | `6.876x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1703583` (1.62 MiB) | `1515539` (1.45 MiB) | `2.348x` | `1.124x` | `2.639x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1502420` (1.43 MiB) | `1305594` (1.25 MiB) | `2.662x` | `1.151x` | `3.064x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1548038` (1.48 MiB) | `964646` (942.04 KiB) | `2.584x` | `1.605x` | `4.147x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `47373` (46.26 KiB) | `7091` (6.92 KiB) | `84.436x` | `6.681x` | `564.095x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `1024` (1.00 KiB) | `4003988` (3.82 MiB) | `4853` (4.74 KiB) | `0.000x` | `825.054x` | `0.211x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `8000000` (7.63 MiB) | `44050` (43.02 KiB) | `5937` (5.80 KiB) | `181.612x` | `7.420x` | `1347.482x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `0` (0 B) | `4002197` (3.82 MiB) | `2893` (2.83 KiB) | `0.000x` | `1383.407x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `3000000` (2.86 MiB) | `7003354` (6.68 MiB) | `5353` (5.23 KiB) | `0.428x` | `1308.305x` | `560.433x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `43122` (42.11 KiB) | `5009` (4.89 KiB) | `92.760x` | `8.609x` | `798.563x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `58030` (56.67 KiB) | `4062605` (3.87 MiB) | `19455` (19.00 KiB) | `0.014x` | `208.821x` | `2.983x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `22051` (21.53 KiB) | `4025685` (3.84 MiB) | `16699` (16.31 KiB) | `0.005x` | `241.073x` | `1.320x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `25445` (24.85 KiB) | `4029094` (3.84 MiB) | `19051` (18.60 KiB) | `0.006x` | `211.490x` | `1.336x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `48191` (47.06 KiB) | `4051797` (3.86 MiB) | `13389` (13.08 KiB) | `0.012x` | `302.621x` | `3.599x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `49433` (48.27 KiB) | `4053267` (3.87 MiB) | `21150` (20.65 KiB) | `0.012x` | `191.644x` | `2.337x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `16873` (16.48 KiB) | `4019990` (3.83 MiB) | `17211` (16.81 KiB) | `0.004x` | `233.571x` | `0.980x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `91870` (89.72 KiB) | `4097616` (3.91 MiB) | `29885` (29.18 KiB) | `0.022x` | `137.113x` | `3.074x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `13001` (12.70 KiB) | `4016652` (3.83 MiB) | `14070` (13.74 KiB) | `0.003x` | `285.476x` | `0.924x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `28101` (27.44 KiB) | `4032974` (3.85 MiB) | `16019` (15.64 KiB) | `0.007x` | `251.762x` | `1.754x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `45607` (44.54 KiB) | `4048344` (3.86 MiB) | `28425` (27.76 KiB) | `0.011x` | `142.422x` | `1.604x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `92475` (90.31 KiB) | `22534` (22.01 KiB) | `43.255x` | `4.104x` | `177.510x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `8000000` (7.63 MiB) | `8118553` (7.74 MiB) | `4466607` (4.26 MiB) | `0.985x` | `1.818x` | `1.791x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `8000000` (7.63 MiB) | `8121071` (7.74 MiB) | `5217669` (4.98 MiB) | `0.985x` | `1.556x` | `1.533x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `68905` (67.29 KiB) | `8594` (8.39 KiB) | `58.051x` | `8.018x` | `465.441x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00000.parquet`: `35062` rows, `3297399` file bytes (3.14 MiB), `27216012` physical bytes (25.96 MiB), `22813399` encoded bytes (21.76 MiB), `3265381` compressed data bytes (3.11 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00001.parquet`: `35082` rows, `3237698` file bytes (3.09 MiB), `27005931` physical bytes (25.75 MiB), `22518803` encoded bytes (21.48 MiB), `3205328` compressed data bytes (3.06 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00002.parquet`: `35588` rows, `3336724` file bytes (3.18 MiB), `27825856` physical bytes (26.54 MiB), `23309519` encoded bytes (22.23 MiB), `3303509` compressed data bytes (3.15 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00003.parquet`: `35454` rows, `3249540` file bytes (3.10 MiB), `27422407` physical bytes (26.15 MiB), `22851610` encoded bytes (21.79 MiB), `3217221` compressed data bytes (3.07 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00004.parquet`: `35286` rows, `3246442` file bytes (3.10 MiB), `27482828` physical bytes (26.21 MiB), `22967992` encoded bytes (21.90 MiB), `3213732` compressed data bytes (3.06 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00005.parquet`: `35411` rows, `3235026` file bytes (3.09 MiB), `27227967` physical bytes (25.97 MiB), `22699191` encoded bytes (21.65 MiB), `3202561` compressed data bytes (3.05 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00006.parquet`: `35595` rows, `3256484` file bytes (3.11 MiB), `27439095` physical bytes (26.17 MiB), `22887614` encoded bytes (21.83 MiB), `3223955` compressed data bytes (3.07 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00007.parquet`: `35888` rows, `3310796` file bytes (3.16 MiB), `27824263` physical bytes (26.54 MiB), `23252381` encoded bytes (22.18 MiB), `3278005` compressed data bytes (3.13 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00008.parquet`: `35399` rows, `3219161` file bytes (3.07 MiB), `27457404` physical bytes (26.19 MiB), `22914565` encoded bytes (21.85 MiB), `3186598` compressed data bytes (3.04 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00009.parquet`: `35231` rows, `3208454` file bytes (3.06 MiB), `27024403` physical bytes (25.77 MiB), `22506731` encoded bytes (21.46 MiB), `3176333` compressed data bytes (3.03 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00010.parquet`: `35503` rows, `3227111` file bytes (3.08 MiB), `27414960` physical bytes (26.14 MiB), `22852668` encoded bytes (21.79 MiB), `3194800` compressed data bytes (3.05 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00011.parquet`: `35657` rows, `3198796` file bytes (3.05 MiB), `27474418` physical bytes (26.20 MiB), `22869201` encoded bytes (21.81 MiB), `3166145` compressed data bytes (3.02 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00012.parquet`: `35615` rows, `3250891` file bytes (3.10 MiB), `27415805` physical bytes (26.15 MiB), `22888326` encoded bytes (21.83 MiB), `3218893` compressed data bytes (3.07 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00013.parquet`: `35632` rows, `3239014` file bytes (3.09 MiB), `27442688` physical bytes (26.17 MiB), `22871003` encoded bytes (21.81 MiB), `3206600` compressed data bytes (3.06 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00014.parquet`: `34695` rows, `3584166` file bytes (3.42 MiB), `24860310` physical bytes (23.71 MiB), `20882404` encoded bytes (19.92 MiB), `3553026` compressed data bytes (3.39 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00015.parquet`: `35010` rows, `3830179` file bytes (3.65 MiB), `24118888` physical bytes (23.00 MiB), `20311543` encoded bytes (19.37 MiB), `3799831` compressed data bytes (3.62 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00016.parquet`: `33923` rows, `3799060` file bytes (3.62 MiB), `23169493` physical bytes (22.10 MiB), `19409486` encoded bytes (18.51 MiB), `3768821` compressed data bytes (3.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00017.parquet`: `34088` rows, `3853278` file bytes (3.67 MiB), `21750699` physical bytes (20.74 MiB), `17767891` encoded bytes (16.94 MiB), `3822765` compressed data bytes (3.65 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00018.parquet`: `33672` rows, `3925871` file bytes (3.74 MiB), `21959689` physical bytes (20.94 MiB), `18008462` encoded bytes (17.17 MiB), `3895310` compressed data bytes (3.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00019.parquet`: `33992` rows, `3941503` file bytes (3.76 MiB), `21853515` physical bytes (20.84 MiB), `17894447` encoded bytes (17.07 MiB), `3910704` compressed data bytes (3.73 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00020.parquet`: `34143` rows, `3935334` file bytes (3.75 MiB), `21904056` physical bytes (20.89 MiB), `17927374` encoded bytes (17.10 MiB), `3904722` compressed data bytes (3.72 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00021.parquet`: `33910` rows, `3878271` file bytes (3.70 MiB), `21742694` physical bytes (20.74 MiB), `17830961` encoded bytes (17.00 MiB), `3847807` compressed data bytes (3.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00022.parquet`: `34203` rows, `3895680` file bytes (3.72 MiB), `21952057` physical bytes (20.94 MiB), `17953881` encoded bytes (17.12 MiB), `3865106` compressed data bytes (3.69 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00023.parquet`: `34160` rows, `3894105` file bytes (3.71 MiB), `21942379` physical bytes (20.93 MiB), `17965957` encoded bytes (17.13 MiB), `3863481` compressed data bytes (3.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00024.parquet`: `33868` rows, `3959764` file bytes (3.78 MiB), `21879535` physical bytes (20.87 MiB), `17942509` encoded bytes (17.11 MiB), `3929341` compressed data bytes (3.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00025.parquet`: `34323` rows, `3850730` file bytes (3.67 MiB), `21793004` physical bytes (20.78 MiB), `17804179` encoded bytes (16.98 MiB), `3820261` compressed data bytes (3.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00026.parquet`: `33923` rows, `3883913` file bytes (3.70 MiB), `21781931` physical bytes (20.77 MiB), `17888438` encoded bytes (17.06 MiB), `3853237` compressed data bytes (3.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00027.parquet`: `34051` rows, `3774260` file bytes (3.60 MiB), `21468615` physical bytes (20.47 MiB), `17533420` encoded bytes (16.72 MiB), `3743803` compressed data bytes (3.57 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict/part-00028.parquet`: `25636` rows, `2930941` file bytes (2.80 MiB), `16547722` physical bytes (15.78 MiB), `13518600` encoded bytes (12.89 MiB), `2901090` compressed data bytes (2.77 MiB)
