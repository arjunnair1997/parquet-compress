# ClickBench Parquet Experiment

- Started: `2026-07-03T15:25:52-04:00`
- Write elapsed: `124ms`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/10000_rows/parquet/rows-10000-comp-snappy-int-plain-str-plain-date-plain-ts-plain`
- Rows: `10000`
- Source TSV bytes for rows, reference only: `8365486` (7.98 MiB)
- Parquet physical bytes before page encoding: `7711890` (7.35 MiB)
- Encoded column bytes before codec compression: `8846976` (8.44 MiB)
- Compressed column data bytes after codec compression: `1274794` (1.22 MiB)
- Parquet file bytes: `1292001` (1.23 MiB)
- Physical/encoded ratio: `0.872x`
- Encoded/compressed-data ratio: `6.940x`
- Physical/compressed-data ratio: `6.050x`
- Physical/parquet-file ratio: `5.969x`
- Files: `1`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `plain`
- Date encoding: `plain`
- Timestamp encoding: `plain`
- Max page size: `256.00 KiB`
- Max dictionary page size: `1.00 MiB`
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
- Elapsed: `74ms`
- Source TSV bytes checked: `8365486` (7.98 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-10000-comp-snappy-int-plain-str-plain-date-plain-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-10000-comp-snappy-int-plain-str-plain-date-plain-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `80087` (78.21 KiB) | `0.999x` | `1.000x` | `0.999x` | `190000` (185.55 KiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2641` (2.58 KiB) | `0.998x` | `15.170x` | `15.146x` | `10000` (9.77 KiB) |
| `Title` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:9` | `10000` | `2200605` (2.10 MiB) | `2247281` (2.14 MiB) | `352517` (344.25 KiB) | `0.979x` | `6.375x` | `6.243x` | `2200673` (2.10 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2072` (2.02 KiB) | `0.998x` | `19.335x` | `19.305x` | `10000` (9.77 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `35183` (34.36 KiB) | `0.999x` | `2.276x` | `2.274x` | `190000` (185.55 KiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2072` (2.02 KiB) | `0.998x` | `19.335x` | `19.305x` | `100000` (97.66 KiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2079` (2.03 KiB) | `0.998x` | `19.270x` | `19.240x` | `20000` (19.53 KiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40062` (39.12 KiB) | `6760` (6.60 KiB) | `0.998x` | `5.926x` | `5.917x` | `100283` (97.93 KiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3690` (3.60 KiB) | `0.998x` | `10.857x` | `10.840x` | `24459` (23.89 KiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `10611` (10.36 KiB) | `0.999x` | `7.547x` | `7.539x` | `199987` (195.30 KiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3137` (3.06 KiB) | `0.998x` | `12.771x` | `12.751x` | `16130` (15.75 KiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3477` (3.40 KiB) | `0.998x` | `11.522x` | `11.504x` | `11313` (11.05 KiB) |
| `URL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:4` | `10000` | `760012` (742.20 KiB) | `800511` (781.75 KiB) | `129117` (126.09 KiB) | `0.949x` | `6.200x` | `5.886x` | `760022` (742.21 KiB) |
| `Referer` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:4` | `10000` | `797798` (779.10 KiB) | `838448` (818.80 KiB) | `145930` (142.51 KiB) | `0.952x` | `5.746x` | `5.467x` | `797822` (779.12 KiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `7387` (7.21 KiB) | `0.998x` | `5.423x` | `5.415x` | `10000` (9.77 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `4754` (4.64 KiB) | `0.998x` | `8.427x` | `8.414x` | `47318` (46.21 KiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `4762` (4.65 KiB) | `0.998x` | `8.413x` | `8.400x` | `28586` (27.92 KiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2267` (2.21 KiB) | `0.998x` | `17.672x` | `17.644x` | `49424` (48.27 KiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2300` (2.25 KiB) | `0.998x` | `17.419x` | `17.391x` | `29545` (28.85 KiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3638` (3.55 KiB) | `0.998x` | `11.012x` | `10.995x` | `39093` (38.18 KiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3619` (3.53 KiB) | `0.998x` | `11.070x` | `11.053x` | `34257` (33.45 KiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2730` (2.67 KiB) | `0.998x` | `14.675x` | `14.652x` | `19813` (19.35 KiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2622` (2.56 KiB) | `0.998x` | `15.280x` | `15.256x` | `18909` (18.47 KiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3219` (3.14 KiB) | `0.998x` | `12.446x` | `12.426x` | `11787` (11.51 KiB) |
| `FlashMinor2` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `37170` (36.30 KiB) | `77229` (75.42 KiB) | `5704` (5.57 KiB) | `0.481x` | `13.539x` | `6.516x` | `37170` (36.30 KiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2270` (2.22 KiB) | `0.998x` | `17.649x` | `17.621x` | `10000` (9.77 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2241` (2.19 KiB) | `0.998x` | `17.877x` | `17.849x` | `10000` (9.77 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3641` (3.56 KiB) | `0.998x` | `11.003x` | `10.986x` | `18125` (17.70 KiB) |
| `UserAgentMinor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `36930` (36.06 KiB) | `76989` (75.18 KiB) | `4953` (4.84 KiB) | `0.480x` | `15.544x` | `7.456x` | `37108` (36.24 KiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2157` (2.11 KiB) | `0.998x` | `18.573x` | `18.544x` | `10000` (9.77 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2132` (2.08 KiB) | `0.998x` | `18.791x` | `18.762x` | `10000` (9.77 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2201` (2.15 KiB) | `0.998x` | `18.202x` | `18.174x` | `10000` (9.77 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2188` (2.14 KiB) | `0.998x` | `18.310x` | `18.282x` | `10011` (9.78 KiB) |
| `MobilePhoneModel` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `1140` (1.11 KiB) | `41195` (40.23 KiB) | `2220` (2.17 KiB) | `0.028x` | `18.556x` | `0.514x` | `1140` (1.11 KiB) |
| `Params` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `0` (0 B) | `40039` (39.10 KiB) | `2045` (2.00 KiB) | `0.000x` | `19.579x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `5433` (5.31 KiB) | `0.998x` | `7.374x` | `7.362x` | `68529` (66.92 KiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `4728` (4.62 KiB) | `0.998x` | `8.474x` | `8.460x` | `17645` (17.23 KiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3534` (3.45 KiB) | `0.998x` | `11.336x` | `11.319x` | `10166` (9.93 KiB) |
| `SearchPhrase` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `65246` (63.72 KiB) | `105322` (102.85 KiB) | `18055` (17.63 KiB) | `0.619x` | `5.833x` | `3.614x` | `65250` (63.72 KiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2434` (2.38 KiB) | `0.998x` | `16.460x` | `16.434x` | `10142` (9.90 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `7248` (7.08 KiB) | `0.998x` | `5.527x` | `5.519x` | `10000` (9.77 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `5249` (5.13 KiB) | `0.998x` | `7.633x` | `7.620x` | `37695` (36.81 KiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `4961` (4.84 KiB) | `0.998x` | `8.076x` | `8.063x` | `30023` (29.32 KiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40062` (39.12 KiB) | `2656` (2.59 KiB) | `0.998x` | `15.084x` | `15.060x` | `29805` (29.11 KiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80079` (78.20 KiB) | `33528` (32.74 KiB) | `0.999x` | `2.388x` | `2.386x` | `190000` (185.55 KiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2756` (2.69 KiB) | `0.998x` | `14.537x` | `14.514x` | `10017` (9.78 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2575` (2.51 KiB) | `0.998x` | `15.558x` | `15.534x` | `10000` (9.77 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2908` (2.84 KiB) | `0.998x` | `13.777x` | `13.755x` | `26504` (25.88 KiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `PageCharset` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `196540` (191.93 KiB) | `236628` (231.08 KiB) | `12041` (11.76 KiB) | `0.831x` | `19.652x` | `16.323x` | `196540` (191.93 KiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40062` (39.12 KiB) | `2132` (2.08 KiB) | `0.998x` | `18.791x` | `18.762x` | `39532` (38.61 KiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `10655` (10.41 KiB) | `0.999x` | `7.516x` | `7.508x` | `176860` (172.71 KiB) |
| `OriginalURL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `74750` (73.00 KiB) | `115520` (112.81 KiB) | `15139` (14.78 KiB) | `0.647x` | `7.631x` | `4.938x` | `74750` (73.00 KiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `34415` (33.61 KiB) | `0.998x` | `1.164x` | `1.162x` | `88404` (86.33 KiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2109` (2.06 KiB) | `0.998x` | `18.996x` | `18.966x` | `10000` (9.77 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `HitColor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `10000` (9.77 KiB) | `50051` (48.88 KiB) | `2602` (2.54 KiB) | `0.200x` | `19.236x` | `3.843x` | `10000` (9.77 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `35180` (34.36 KiB) | `0.999x` | `2.276x` | `2.274x` | `190000` (185.55 KiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3434` (3.35 KiB) | `0.998x` | `11.667x` | `11.648x` | `17412` (17.00 KiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3116` (3.04 KiB) | `0.998x` | `12.857x` | `12.837x` | `10000` (9.77 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3172` (3.10 KiB) | `0.998x` | `12.630x` | `12.610x` | `10000` (9.77 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3809` (3.72 KiB) | `0.998x` | `10.518x` | `10.501x` | `22570` (22.04 KiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3720` (3.63 KiB) | `0.998x` | `10.770x` | `10.753x` | `13859` (13.53 KiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `6807` (6.65 KiB) | `0.998x` | `5.886x` | `5.876x` | `100205` (97.86 KiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2086` (2.04 KiB) | `0.998x` | `19.206x` | `19.175x` | `20010` (19.54 KiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `20000` (19.53 KiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2080` (2.03 KiB) | `0.998x` | `19.261x` | `19.231x` | `19996` (19.53 KiB) |
| `BrowserLanguage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `20156` (19.68 KiB) | `60213` (58.80 KiB) | `3262` (3.19 KiB) | `0.335x` | `18.459x` | `6.179x` | `20156` (19.68 KiB) |
| `BrowserCountry` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `32430` (31.67 KiB) | `72489` (70.79 KiB) | `4473` (4.37 KiB) | `0.447x` | `16.206x` | `7.250x` | `38645` (37.74 KiB) |
| `SocialNetwork` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `0` (0 B) | `40039` (39.10 KiB) | `2045` (2.00 KiB) | `0.000x` | `19.579x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `0` (0 B) | `40039` (39.10 KiB) | `2045` (2.00 KiB) | `0.000x` | `19.579x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2693` (2.63 KiB) | `0.998x` | `14.877x` | `14.853x` | `10125` (9.89 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40062` (39.12 KiB) | `3882` (3.79 KiB) | `0.998x` | `10.320x` | `10.304x` | `10493` (10.25 KiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `11343` (11.08 KiB) | `0.998x` | `3.532x` | `3.526x` | `16080` (15.70 KiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `9934` (9.70 KiB) | `0.998x` | `4.033x` | `4.027x` | `13892` (13.57 KiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `7294` (7.12 KiB) | `0.998x` | `5.493x` | `5.484x` | `11851` (11.57 KiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2079` (2.03 KiB) | `0.998x` | `19.270x` | `19.240x` | `10000` (9.77 KiB) |
| `SocialSourcePage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `42` (42 B) | `40103` (39.16 KiB) | `2085` (2.04 KiB) | `0.001x` | `19.234x` | `0.020x` | `42` (42 B) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80079` (78.20 KiB) | `4086` (3.99 KiB) | `0.999x` | `19.598x` | `19.579x` | `10000` (9.77 KiB) |
| `ParamOrderID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `0` (0 B) | `40039` (39.10 KiB) | `2045` (2.00 KiB) | `0.000x` | `19.579x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `30000` (29.30 KiB) | `70059` (68.42 KiB) | `3571` (3.49 KiB) | `0.428x` | `19.619x` | `8.401x` | `30000` (29.30 KiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `OpenstatServiceName` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `560` (560 B) | `40639` (39.69 KiB) | `2217` (2.17 KiB) | `0.014x` | `18.331x` | `0.253x` | `560` (560 B) |
| `OpenstatCampaignID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `272` (272 B) | `40339` (39.39 KiB) | `2178` (2.13 KiB) | `0.007x` | `18.521x` | `0.125x` | `272` (272 B) |
| `OpenstatAdID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `358` (358 B) | `40425` (39.48 KiB) | `2276` (2.22 KiB) | `0.009x` | `17.761x` | `0.157x` | `358` (358 B) |
| `OpenstatSourceID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `692` (692 B) | `40775` (39.82 KiB) | `2240` (2.19 KiB) | `0.017x` | `18.203x` | `0.309x` | `692` (692 B) |
| `UTMSource` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `1621` (1.58 KiB) | `41688` (40.71 KiB) | `2496` (2.44 KiB) | `0.039x` | `16.702x` | `0.649x` | `1621` (1.58 KiB) |
| `UTMMedium` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `502` (502 B) | `40557` (39.61 KiB) | `2380` (2.32 KiB) | `0.012x` | `17.041x` | `0.211x` | `502` (502 B) |
| `UTMCampaign` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `3304` (3.23 KiB) | `43419` (42.40 KiB) | `2652` (2.59 KiB) | `0.076x` | `16.372x` | `1.246x` | `3304` (3.23 KiB) |
| `UTMContent` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `474` (474 B) | `40533` (39.58 KiB) | `2253` (2.20 KiB) | `0.012x` | `17.991x` | `0.210x` | `474` (474 B) |
| `UTMTerm` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `1288` (1.26 KiB) | `41367` (40.40 KiB) | `2332` (2.28 KiB) | `0.031x` | `17.739x` | `0.552x` | `1288` (1.26 KiB) |
| `FromTag` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `0` (0 B) | `40039` (39.10 KiB) | `2045` (2.00 KiB) | `0.000x` | `19.579x` | `0.000x` | `0` (0 B) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2273` (2.22 KiB) | `0.998x` | `17.626x` | `17.598x` | `10000` (9.77 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `34180` (33.38 KiB) | `0.999x` | `2.343x` | `2.341x` | `193030` (188.51 KiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `38484` (37.58 KiB) | `0.999x` | `2.081x` | `2.079x` | `193212` (188.68 KiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/10000_rows/parquet/rows-10000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00000.parquet`: `10000` rows, `1292001` file bytes (1.23 MiB), `7711890` physical bytes (7.35 MiB), `8846976` encoded bytes (8.44 MiB), `1274794` compressed data bytes (1.22 MiB)
