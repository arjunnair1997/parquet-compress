# ClickBench Parquet Experiment

- Started: `2026-07-03T14:42:14-04:00`
- Write elapsed: `123ms`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/10000_rows/parquet/rows-10000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`
- Rows: `10000`
- Source TSV bytes for rows, reference only: `8365486` (7.98 MiB)
- Parquet physical bytes before page encoding: `7711890` (7.35 MiB)
- Encoded column bytes before codec compression: `4946819` (4.72 MiB)
- Compressed column data bytes after codec compression: `813122` (794.06 KiB)
- Parquet file bytes: `830082` (810.63 KiB)
- Physical/encoded ratio: `1.559x`
- Encoded/compressed-data ratio: `6.084x`
- Physical/compressed-data ratio: `9.484x`
- Physical/parquet-file ratio: `9.291x`
- Files: `1`

## Settings

- Compression: `zstd-3`
- Int encoding: `plain`
- String encoding: `delta-byte-array`
- Date encoding: `plain`
- Timestamp encoding: `rle-dict`
- Max page size: `256.00 KiB`
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
- Rows read and compared: `10000`
- Files read: `1`
- Elapsed: `79ms`
- Source TSV bytes checked: `8365486` (7.98 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-10000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict_columns.tsv](../../tsvs/2026-07-03_rows-10000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `80092` (78.21 KiB) | `0.999x` | `1.000x` | `0.999x` | `190000` (185.55 KiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `705` (705 B) | `0.998x` | `56.827x` | `56.738x` | `10000` (9.77 KiB) |
| `Title` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:9` | `10000` | `2200605` (2.10 MiB) | `900557` (879.45 KiB) | `230664` (225.26 KiB) | `2.444x` | `3.904x` | `9.540x` | `2200673` (2.10 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40062` (39.12 KiB) | `85` (85 B) | `0.998x` | `471.318x` | `470.588x` | `10000` (9.77 KiB) |
| `EventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `58661` (57.29 KiB) | `33469` (32.68 KiB) | `1.364x` | `1.753x` | `2.390x` | `190000` (185.55 KiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40062` (39.12 KiB) | `85` (85 B) | `0.998x` | `471.318x` | `470.588x` | `100000` (97.66 KiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40062` (39.12 KiB) | `89` (89 B) | `0.998x` | `450.135x` | `449.438x` | `20000` (19.53 KiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3641` (3.56 KiB) | `0.998x` | `11.003x` | `10.986x` | `100283` (97.93 KiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `1647` (1.61 KiB) | `0.998x` | `24.325x` | `24.287x` | `24459` (23.89 KiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80079` (78.20 KiB) | `5858` (5.72 KiB) | `0.999x` | `13.670x` | `13.657x` | `199987` (195.30 KiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40061` (39.12 KiB) | `72` (72 B) | `0.998x` | `556.403x` | `555.556x` | `10000` (9.77 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `955` (955 B) | `0.998x` | `41.951x` | `41.885x` | `16130` (15.75 KiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `1421` (1.39 KiB) | `0.998x` | `28.194x` | `28.149x` | `11313` (11.05 KiB) |
| `URL` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:4` | `10000` | `760012` (742.20 KiB) | `211638` (206.68 KiB) | `87048` (85.01 KiB) | `3.591x` | `2.431x` | `8.731x` | `760022` (742.21 KiB) |
| `Referer` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:4` | `10000` | `797798` (779.10 KiB) | `277125` (270.63 KiB) | `104736` (102.28 KiB) | `2.879x` | `2.646x` | `7.617x` | `797822` (779.12 KiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3154` (3.08 KiB) | `0.998x` | `12.702x` | `12.682x` | `10000` (9.77 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40062` (39.12 KiB) | `2611` (2.55 KiB) | `0.998x` | `15.344x` | `15.320x` | `47318` (46.21 KiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2656` (2.59 KiB) | `0.998x` | `15.084x` | `15.060x` | `28586` (27.92 KiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `430` (430 B) | `0.998x` | `93.170x` | `93.023x` | `49424` (48.27 KiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `445` (445 B) | `0.998x` | `90.029x` | `89.888x` | `29545` (28.85 KiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `1819` (1.78 KiB) | `0.998x` | `22.025x` | `21.990x` | `39093` (38.18 KiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `1773` (1.73 KiB) | `0.998x` | `22.596x` | `22.561x` | `34257` (33.45 KiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `907` (907 B) | `0.998x` | `44.171x` | `44.101x` | `19813` (19.35 KiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `791` (791 B) | `0.998x` | `50.649x` | `50.569x` | `18909` (18.47 KiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `1238` (1.21 KiB) | `0.998x` | `32.361x` | `32.310x` | `11787` (11.51 KiB) |
| `FlashMinor2` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `37170` (36.30 KiB) | `11521` (11.25 KiB) | `3722` (3.63 KiB) | `3.226x` | `3.095x` | `9.987x` | `37170` (36.30 KiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `340` (340 B) | `0.998x` | `117.832x` | `117.647x` | `10000` (9.77 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `354` (354 B) | `0.998x` | `113.172x` | `112.994x` | `10000` (9.77 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `1526` (1.49 KiB) | `0.998x` | `26.254x` | `26.212x` | `18125` (17.70 KiB) |
| `UserAgentMinor` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `36930` (36.06 KiB) | `10213` (9.97 KiB) | `2326` (2.27 KiB) | `3.616x` | `4.391x` | `15.877x` | `37108` (36.24 KiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `235` (235 B) | `0.998x` | `170.481x` | `170.213x` | `10000` (9.77 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `202` (202 B) | `0.998x` | `198.332x` | `198.020x` | `10000` (9.77 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `296` (296 B) | `0.998x` | `135.348x` | `135.135x` | `10000` (9.77 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `278` (278 B) | `0.998x` | `144.112x` | `143.885x` | `10011` (9.78 KiB) |
| `MobilePhoneModel` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `1140` (1.11 KiB) | `2652` (2.59 KiB) | `428` (428 B) | `0.430x` | `6.196x` | `2.664x` | `1140` (1.11 KiB) |
| `Params` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `0` (0 B) | `839` (839 B) | `64` (64 B) | `0.000x` | `13.109x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2953` (2.88 KiB) | `0.998x` | `13.567x` | `13.546x` | `68529` (66.92 KiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2586` (2.53 KiB) | `0.998x` | `15.492x` | `15.468x` | `17645` (17.23 KiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `1654` (1.62 KiB) | `0.998x` | `24.222x` | `24.184x` | `10166` (9.93 KiB) |
| `SearchPhrase` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `65246` (63.72 KiB) | `50156` (48.98 KiB) | `14444` (14.11 KiB) | `1.301x` | `3.472x` | `4.517x` | `65250` (63.72 KiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40062` (39.12 KiB) | `650` (650 B) | `0.998x` | `61.634x` | `61.538x` | `10142` (9.90 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2790` (2.72 KiB) | `0.998x` | `14.359x` | `14.337x` | `10000` (9.77 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2986` (2.92 KiB) | `0.998x` | `13.417x` | `13.396x` | `37695` (36.81 KiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2425` (2.37 KiB) | `0.998x` | `16.521x` | `16.495x` | `30023` (29.32 KiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `809` (809 B) | `0.998x` | `49.522x` | `49.444x` | `29805` (29.11 KiB) |
| `ClientEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `55120` (53.83 KiB) | `32305` (31.55 KiB) | `1.451x` | `1.706x` | `2.476x` | `190000` (185.55 KiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `611` (611 B) | `0.998x` | `65.570x` | `65.466x` | `10017` (9.78 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `671` (671 B) | `0.998x` | `59.706x` | `59.613x` | `10000` (9.77 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `1121` (1.09 KiB) | `0.998x` | `35.739x` | `35.682x` | `26504` (25.88 KiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40061` (39.12 KiB) | `72` (72 B) | `0.998x` | `556.403x` | `555.556x` | `10000` (9.77 KiB) |
| `PageCharset` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `196540` (191.93 KiB) | `5072` (4.95 KiB) | `658` (658 B) | `38.750x` | `7.708x` | `298.693x` | `196540` (191.93 KiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `232` (232 B) | `0.998x` | `172.685x` | `172.414x` | `39532` (38.61 KiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40061` (39.12 KiB) | `72` (72 B) | `0.998x` | `556.403x` | `555.556x` | `10000` (9.77 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40061` (39.12 KiB) | `72` (72 B) | `0.998x` | `556.403x` | `555.556x` | `10000` (9.77 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40061` (39.12 KiB) | `72` (72 B) | `0.998x` | `556.403x` | `555.556x` | `10000` (9.77 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80079` (78.20 KiB) | `6080` (5.94 KiB) | `0.999x` | `13.171x` | `13.158x` | `176860` (172.71 KiB) |
| `OriginalURL` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `74750` (73.00 KiB) | `42916` (41.91 KiB) | `11271` (11.01 KiB) | `1.742x` | `3.808x` | `6.632x` | `74750` (73.00 KiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `24108` (23.54 KiB) | `0.998x` | `1.662x` | `1.659x` | `88404` (86.33 KiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40061` (39.12 KiB) | `72` (72 B) | `0.998x` | `556.403x` | `555.556x` | `10000` (9.77 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40061` (39.12 KiB) | `72` (72 B) | `0.998x` | `556.403x` | `555.556x` | `10000` (9.77 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40061` (39.12 KiB) | `72` (72 B) | `0.998x` | `556.403x` | `555.556x` | `10000` (9.77 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `176` (176 B) | `0.998x` | `227.631x` | `227.273x` | `10000` (9.77 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40061` (39.12 KiB) | `72` (72 B) | `0.998x` | `556.403x` | `555.556x` | `10000` (9.77 KiB) |
| `HitColor` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `10000` (9.77 KiB) | `1191` (1.16 KiB) | `236` (236 B) | `8.396x` | `5.047x` | `42.373x` | `10000` (9.77 KiB) |
| `LocalEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `58525` (57.15 KiB) | `33468` (32.68 KiB) | `1.367x` | `1.749x` | `2.390x` | `190000` (185.55 KiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `1493` (1.46 KiB) | `0.998x` | `26.834x` | `26.792x` | `17412` (17.00 KiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `1052` (1.03 KiB) | `0.998x` | `38.083x` | `38.023x` | `10000` (9.77 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `1273` (1.24 KiB) | `0.998x` | `31.471x` | `31.422x` | `10000` (9.77 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40062` (39.12 KiB) | `1861` (1.82 KiB) | `0.998x` | `21.527x` | `21.494x` | `22570` (22.04 KiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `1567` (1.53 KiB) | `0.998x` | `25.567x` | `25.526x` | `13859` (13.53 KiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3684` (3.60 KiB) | `0.998x` | `10.875x` | `10.858x` | `100205` (97.86 KiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40062` (39.12 KiB) | `103` (103 B) | `0.998x` | `388.951x` | `388.350x` | `20010` (19.54 KiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40062` (39.12 KiB) | `73` (73 B) | `0.998x` | `548.795x` | `547.945x` | `20000` (19.53 KiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40061` (39.12 KiB) | `95` (95 B) | `0.998x` | `421.695x` | `421.053x` | `19996` (19.53 KiB) |
| `BrowserLanguage` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `20156` (19.68 KiB) | `3313` (3.24 KiB) | `856` (856 B) | `6.084x` | `3.870x` | `23.547x` | `20156` (19.68 KiB) |
| `BrowserCountry` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `32430` (31.67 KiB) | `9795` (9.57 KiB) | `1959` (1.91 KiB) | `3.311x` | `5.000x` | `16.554x` | `38645` (37.74 KiB) |
| `SocialNetwork` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `0` (0 B) | `839` (839 B) | `64` (64 B) | `0.000x` | `13.109x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `0` (0 B) | `839` (839 B) | `64` (64 B) | `0.000x` | `13.109x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40061` (39.12 KiB) | `72` (72 B) | `0.998x` | `556.403x` | `555.556x` | `10000` (9.77 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40061` (39.12 KiB) | `72` (72 B) | `0.998x` | `556.403x` | `555.556x` | `10000` (9.77 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `839` (839 B) | `0.998x` | `47.751x` | `47.676x` | `10125` (9.89 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `1748` (1.71 KiB) | `0.998x` | `22.919x` | `22.883x` | `10493` (10.25 KiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `8157` (7.97 KiB) | `0.998x` | `4.911x` | `4.904x` | `16080` (15.70 KiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `6951` (6.79 KiB) | `0.998x` | `5.764x` | `5.755x` | `13892` (13.57 KiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `4732` (4.62 KiB) | `0.998x` | `8.466x` | `8.453x` | `11851` (11.57 KiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40062` (39.12 KiB) | `100` (100 B) | `0.998x` | `400.620x` | `400.000x` | `10000` (9.77 KiB) |
| `SocialSourcePage` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `42` (42 B) | `972` (972 B) | `122` (122 B) | `0.043x` | `7.967x` | `0.344x` | `42` (42 B) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80078` (78.20 KiB) | `91` (91 B) | `0.999x` | `879.978x` | `879.121x` | `10000` (9.77 KiB) |
| `ParamOrderID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `0` (0 B) | `839` (839 B) | `64` (64 B) | `0.000x` | `13.109x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `30000` (29.30 KiB) | `910` (910 B) | `123` (123 B) | `32.967x` | `7.398x` | `243.902x` | `30000` (29.30 KiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40061` (39.12 KiB) | `72` (72 B) | `0.998x` | `556.403x` | `555.556x` | `10000` (9.77 KiB) |
| `OpenstatServiceName` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `560` (560 B) | `3588` (3.50 KiB) | `384` (384 B) | `0.156x` | `9.344x` | `1.458x` | `560` (560 B) |
| `OpenstatCampaignID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `272` (272 B) | `2619` (2.56 KiB) | `376` (376 B) | `0.104x` | `6.965x` | `0.723x` | `272` (272 B) |
| `OpenstatAdID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `358` (358 B) | `2982` (2.91 KiB) | `551` (551 B) | `0.120x` | `5.412x` | `0.650x` | `358` (358 B) |
| `OpenstatSourceID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `692` (692 B) | `3658` (3.57 KiB) | `420` (420 B) | `0.189x` | `8.710x` | `1.648x` | `692` (692 B) |
| `UTMSource` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `1621` (1.58 KiB) | `5530` (5.40 KiB) | `1003` (1003 B) | `0.293x` | `5.513x` | `1.616x` | `1621` (1.58 KiB) |
| `UTMMedium` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `502` (502 B) | `4077` (3.98 KiB) | `669` (669 B) | `0.123x` | `6.094x` | `0.750x` | `502` (502 B) |
| `UTMCampaign` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `3304` (3.23 KiB) | `7589` (7.41 KiB) | `1147` (1.12 KiB) | `0.435x` | `6.616x` | `2.881x` | `3304` (3.23 KiB) |
| `UTMContent` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `474` (474 B) | `2912` (2.84 KiB) | `390` (390 B) | `0.163x` | `7.467x` | `1.215x` | `474` (474 B) |
| `UTMTerm` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `1288` (1.26 KiB) | `4608` (4.50 KiB) | `560` (560 B) | `0.280x` | `8.229x` | `2.300x` | `1288` (1.26 KiB) |
| `FromTag` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:1` | `10000` | `0` (0 B) | `839` (839 B) | `64` (64 B) | `0.000x` | `13.109x` | `0.000x` | `0` (0 B) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `455` (455 B) | `0.998x` | `88.051x` | `87.912x` | `10000` (9.77 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `23587` (23.03 KiB) | `0.999x` | `3.395x` | `3.392x` | `193030` (188.51 KiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `27185` (26.55 KiB) | `0.999x` | `2.946x` | `2.943x` | `193212` (188.68 KiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40061` (39.12 KiB) | `72` (72 B) | `0.998x` | `556.403x` | `555.556x` | `10000` (9.77 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/10000_rows/parquet/rows-10000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00000.parquet`: `10000` rows, `830082` file bytes (810.63 KiB), `7711890` physical bytes (7.35 MiB), `4946819` encoded bytes (4.72 MiB), `813122` compressed data bytes (794.06 KiB)
