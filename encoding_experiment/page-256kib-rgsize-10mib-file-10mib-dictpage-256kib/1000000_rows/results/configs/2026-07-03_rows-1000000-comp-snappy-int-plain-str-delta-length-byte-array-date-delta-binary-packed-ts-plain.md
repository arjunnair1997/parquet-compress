# ClickBench Parquet Experiment

- Started: `2026-07-03T23:30:02-04:00`
- Write elapsed: `11.25s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `717271831` (684.04 MiB)
- Compressed column data bytes after codec compression: `131022775` (124.95 MiB)
- Parquet file bytes: `131982481` (125.87 MiB)
- Physical/encoded ratio: `0.993x`
- Encoded/compressed-data ratio: `5.474x`
- Physical/compressed-data ratio: `5.437x`
- Physical/parquet-file ratio: `5.398x`
- Files: `31`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `delta-length-byte-array`
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
- Files read: `31`
- Elapsed: `6.844s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004876` (7.63 MiB) | `8005303` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `262068` (255.93 KiB) | `0.999x` | `15.278x` | `15.263x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:505` | `1000000` | `138409995` (132.00 MiB) | `140039289` (133.55 MiB) | `21337537` (20.35 MiB) | `0.988x` | `6.563x` | `6.487x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `204389` (199.60 KiB) | `0.999x` | `19.589x` | `19.571x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004877` (7.63 MiB) | `4282790` (4.08 MiB) | `0.999x` | `1.869x` | `1.868x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `52125` (50.90 KiB) | `8141` (7.95 KiB) | `76.739x` | `6.403x` | `491.340x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `204405` (199.61 KiB) | `0.999x` | `19.588x` | `19.569x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `719392` (702.53 KiB) | `0.999x` | `5.566x` | `5.560x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `396584` (387.29 KiB) | `0.999x` | `10.096x` | `10.086x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004875` (7.63 MiB) | `1084901` (1.03 MiB) | `0.999x` | `7.378x` | `7.374x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204207` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `318082` (310.63 KiB) | `0.999x` | `12.587x` | `12.575x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `342291` (334.27 KiB) | `0.999x` | `11.697x` | `11.686x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:342` | `1000000` | `88562192` (84.46 MiB) | `89785621` (85.63 MiB) | `20795139` (19.83 MiB) | `0.986x` | `4.318x` | `4.259x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:298` | `1000000` | `79583339` (75.90 MiB) | `80832992` (77.09 MiB) | `19455042` (18.55 MiB) | `0.985x` | `4.155x` | `4.091x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003871` (3.82 MiB) | `492008` (480.48 KiB) | `0.999x` | `8.138x` | `8.130x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003870` (3.82 MiB) | `509023` (497.09 KiB) | `0.999x` | `7.866x` | `7.858x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003854` (3.82 MiB) | `458119` (447.38 KiB) | `0.999x` | `8.740x` | `8.731x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `272035` (265.66 KiB) | `0.999x` | `14.718x` | `14.704x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `232249` (226.81 KiB) | `0.999x` | `17.239x` | `17.223x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `373498` (364.74 KiB) | `0.999x` | `10.720x` | `10.710x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `373256` (364.51 KiB) | `0.999x` | `10.727x` | `10.717x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `282938` (276.31 KiB) | `0.999x` | `14.151x` | `14.137x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `245405` (239.65 KiB) | `0.999x` | `16.315x` | `16.300x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `328474` (320.78 KiB) | `0.999x` | `12.189x` | `12.178x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3354477` (3.20 MiB) | `3707308` (3.54 MiB) | `433479` (423.32 KiB) | `0.905x` | `8.552x` | `7.738x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `219194` (214.06 KiB) | `0.999x` | `18.266x` | `18.249x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `217370` (212.28 KiB) | `0.999x` | `18.419x` | `18.402x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `365189` (356.63 KiB) | `0.999x` | `10.964x` | `10.953x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3767530` (3.59 MiB) | `4017958` (3.83 MiB) | `328133` (320.44 KiB) | `0.938x` | `12.245x` | `11.482x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204986` (200.18 KiB) | `0.999x` | `19.532x` | `19.514x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204981` (200.18 KiB) | `0.999x` | `19.533x` | `19.514x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `218503` (213.38 KiB) | `0.999x` | `18.324x` | `18.306x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `215467` (210.42 KiB) | `0.999x` | `18.582x` | `18.564x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `81583` (79.67 KiB) | `231565` (226.14 KiB) | `42897` (41.89 KiB) | `0.352x` | `5.398x` | `1.902x` | `81583` (79.67 KiB) |
| `Params` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `0` (0 B) | `41853` (40.87 KiB) | `4896` (4.78 KiB) | `0.000x` | `8.548x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003879` (3.82 MiB) | `560906` (547.76 KiB) | `0.999x` | `7.138x` | `7.131x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003870` (3.82 MiB) | `526837` (514.49 KiB) | `0.999x` | `7.600x` | `7.592x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `293188` (286.32 KiB) | `0.999x` | `13.656x` | `13.643x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3528017` (3.36 MiB) | `4224051` (4.03 MiB) | `1015102` (991.31 KiB) | `0.835x` | `4.161x` | `3.476x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `220374` (215.21 KiB) | `0.999x` | `18.168x` | `18.151x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003876` (3.82 MiB) | `505897` (494.04 KiB) | `0.999x` | `7.914x` | `7.907x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003871` (3.82 MiB) | `516804` (504.69 KiB) | `0.999x` | `7.747x` | `7.740x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003882` (3.82 MiB) | `551959` (539.02 KiB) | `0.999x` | `7.254x` | `7.247x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `285764` (279.07 KiB) | `0.999x` | `14.011x` | `13.998x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004873` (7.63 MiB) | `4229176` (4.03 MiB) | `0.999x` | `1.893x` | `1.892x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `287293` (280.56 KiB) | `0.999x` | `13.936x` | `13.923x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `269749` (263.43 KiB) | `0.999x` | `14.843x` | `14.829x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `304921` (297.77 KiB) | `0.999x` | `13.131x` | `13.118x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204244` (199.46 KiB) | `0.999x` | `19.603x` | `19.584x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:91` | `1000000` | `13587860` (12.96 MiB) | `13658178` (13.03 MiB) | `699186` (682.80 KiB) | `0.995x` | `19.534x` | `19.434x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `205212` (200.40 KiB) | `0.999x` | `19.511x` | `19.492x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `256539` (250.53 KiB) | `0.999x` | `15.607x` | `15.592x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `206715` (201.87 KiB) | `0.999x` | `19.369x` | `19.350x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `235460` (229.94 KiB) | `0.999x` | `17.004x` | `16.988x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004873` (7.63 MiB) | `1151628` (1.10 MiB) | `0.999x` | `6.951x` | `6.947x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:140` | `1000000` | `27797671` (26.51 MiB) | `28795413` (27.46 MiB) | `7053886` (6.73 MiB) | `0.965x` | `4.082x` | `3.941x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `3688236` (3.52 MiB) | `0.999x` | `1.086x` | `1.085x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204207` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204207` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204207` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003844` (3.82 MiB) | `295713` (288.78 KiB) | `0.999x` | `13.540x` | `13.527x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204207` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `1000000` (976.56 KiB) | `1043317` (1018.86 KiB) | `77094` (75.29 KiB) | `0.958x` | `13.533x` | `12.971x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004877` (7.63 MiB) | `4284728` (4.09 MiB) | `0.999x` | `1.868x` | `1.867x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `334001` (326.17 KiB) | `0.999x` | `11.988x` | `11.976x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `302572` (295.48 KiB) | `0.999x` | `13.233x` | `13.220x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `318316` (310.86 KiB) | `0.999x` | `12.578x` | `12.566x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `385691` (376.65 KiB) | `0.999x` | `10.381x` | `10.371x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `381787` (372.84 KiB) | `0.999x` | `10.487x` | `10.477x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `706747` (690.18 KiB) | `0.999x` | `5.665x` | `5.660x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `295387` (288.46 KiB) | `0.999x` | `13.555x` | `13.542x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204209` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `258419` (252.36 KiB) | `0.999x` | `15.494x` | `15.479x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `2001192` (1.91 MiB) | `2051477` (1.96 MiB) | `125426` (122.49 KiB) | `0.975x` | `16.356x` | `15.955x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3325142` (3.17 MiB) | `3637745` (3.47 MiB) | `352615` (344.35 KiB) | `0.914x` | `10.316x` | `9.430x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `0` (0 B) | `41853` (40.87 KiB) | `4896` (4.78 KiB) | `0.000x` | `8.548x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `0` (0 B) | `41853` (40.87 KiB) | `4896` (4.78 KiB) | `0.000x` | `8.548x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204207` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `269184` (262.88 KiB) | `0.999x` | `14.874x` | `14.860x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `327374` (319.70 KiB) | `0.999x` | `12.230x` | `12.218x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003875` (3.82 MiB) | `564692` (551.46 KiB) | `0.999x` | `7.090x` | `7.084x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `1725906` (1.65 MiB) | `0.999x` | `2.320x` | `2.318x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003897` (3.82 MiB) | `1283257` (1.22 MiB) | `0.999x` | `3.120x` | `3.117x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `810373` (791.38 KiB) | `0.999x` | `4.941x` | `4.936x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204599` (199.80 KiB) | `0.999x` | `19.569x` | `19.550x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `1024` (1.00 KiB) | `46408` (45.32 KiB) | `7322` (7.15 KiB) | `0.022x` | `6.338x` | `0.140x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004817` (7.63 MiB) | `405187` (395.69 KiB) | `0.999x` | `19.756x` | `19.744x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `0` (0 B) | `41853` (40.87 KiB) | `4896` (4.78 KiB) | `0.000x` | `8.548x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3000000` (2.86 MiB) | `3044411` (2.90 MiB) | `157364` (153.68 KiB) | `0.985x` | `19.346x` | `19.064x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204207` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `58030` (56.67 KiB) | `216981` (211.90 KiB) | `37104` (36.23 KiB) | `0.267x` | `5.848x` | `1.564x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `22051` (21.53 KiB) | `122864` (119.98 KiB) | `28665` (27.99 KiB) | `0.179x` | `4.286x` | `0.769x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `25445` (24.85 KiB) | `130671` (127.61 KiB) | `34816` (34.00 KiB) | `0.195x` | `3.753x` | `0.731x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `48191` (47.06 KiB) | `155722` (152.07 KiB) | `24746` (24.17 KiB) | `0.309x` | `6.293x` | `1.947x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `49433` (48.27 KiB) | `187809` (183.41 KiB) | `41380` (40.41 KiB) | `0.263x` | `4.539x` | `1.195x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `16873` (16.48 KiB) | `133216` (130.09 KiB) | `28760` (28.09 KiB) | `0.127x` | `4.632x` | `0.587x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `91870` (89.72 KiB) | `252870` (246.94 KiB) | `57692` (56.34 KiB) | `0.363x` | `4.383x` | `1.592x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `13001` (12.70 KiB) | `94142` (91.94 KiB) | `22040` (21.52 KiB) | `0.138x` | `4.271x` | `0.590x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `28101` (27.44 KiB) | `128959` (125.94 KiB) | `25357` (24.76 KiB) | `0.218x` | `5.086x` | `1.108x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `45607` (44.54 KiB) | `211634` (206.67 KiB) | `46234` (45.15 KiB) | `0.215x` | `4.577x` | `0.986x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `213620` (208.61 KiB) | `0.999x` | `18.743x` | `18.725x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004876` (7.63 MiB) | `3641282` (3.47 MiB) | `0.999x` | `2.198x` | `2.197x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004875` (7.63 MiB) | `4387083` (4.18 MiB) | `0.999x` | `1.825x` | `1.824x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204656` (199.86 KiB) | `0.999x` | `19.564x` | `19.545x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00000.parquet`: `33122` rows, `4197753` file bytes (4.00 MiB), `25723561` physical bytes (24.53 MiB), `25908349` encoded bytes (24.71 MiB), `4165814` compressed data bytes (3.97 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00001.parquet`: `33171` rows, `4144701` file bytes (3.95 MiB), `25509930` physical bytes (24.33 MiB), `25687158` encoded bytes (24.50 MiB), `4112382` compressed data bytes (3.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00002.parquet`: `32706` rows, `4138469` file bytes (3.95 MiB), `25599070` physical bytes (24.41 MiB), `25780271` encoded bytes (24.59 MiB), `4105654` compressed data bytes (3.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00003.parquet`: `32955` rows, `4135775` file bytes (3.94 MiB), `25551069` physical bytes (24.37 MiB), `25731176` encoded bytes (24.54 MiB), `4103404` compressed data bytes (3.91 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00004.parquet`: `33164` rows, `4152904` file bytes (3.96 MiB), `25755586` physical bytes (24.56 MiB), `25939317` encoded bytes (24.74 MiB), `4119767` compressed data bytes (3.93 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00005.parquet`: `33399` rows, `4109957` file bytes (3.92 MiB), `25776558` physical bytes (24.58 MiB), `25953441` encoded bytes (24.75 MiB), `4077557` compressed data bytes (3.89 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00006.parquet`: `32963` rows, `4107262` file bytes (3.92 MiB), `25357286` physical bytes (24.18 MiB), `25540109` encoded bytes (24.36 MiB), `4075120` compressed data bytes (3.89 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00007.parquet`: `33215` rows, `4144818` file bytes (3.95 MiB), `25767417` physical bytes (24.57 MiB), `25947437` encoded bytes (24.75 MiB), `4112311` compressed data bytes (3.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00008.parquet`: `33047` rows, `4109535` file bytes (3.92 MiB), `25569453` physical bytes (24.38 MiB), `25751420` encoded bytes (24.56 MiB), `4076667` compressed data bytes (3.89 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00009.parquet`: `33161` rows, `4098457` file bytes (3.91 MiB), `25583001` physical bytes (24.40 MiB), `25764642` encoded bytes (24.57 MiB), `4066088` compressed data bytes (3.88 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00010.parquet`: `33054` rows, `4129674` file bytes (3.94 MiB), `25557558` physical bytes (24.37 MiB), `25736367` encoded bytes (24.54 MiB), `4097670` compressed data bytes (3.91 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00011.parquet`: `33068` rows, `4048444` file bytes (3.86 MiB), `25419204` physical bytes (24.24 MiB), `25597099` encoded bytes (24.41 MiB), `4016254` compressed data bytes (3.83 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00012.parquet`: `33230` rows, `4078430` file bytes (3.89 MiB), `25601404` physical bytes (24.42 MiB), `25783883` encoded bytes (24.59 MiB), `4046180` compressed data bytes (3.86 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00013.parquet`: `33281` rows, `4122162` file bytes (3.93 MiB), `25571581` physical bytes (24.39 MiB), `25752351` encoded bytes (24.56 MiB), `4090195` compressed data bytes (3.90 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00014.parquet`: `32751` rows, `4041338` file bytes (3.85 MiB), `25248451` physical bytes (24.08 MiB), `25425838` encoded bytes (24.25 MiB), `4009159` compressed data bytes (3.82 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00015.parquet`: `33225` rows, `4390759` file bytes (4.19 MiB), `23860162` physical bytes (22.75 MiB), `24006500` encoded bytes (22.89 MiB), `4359598` compressed data bytes (4.16 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00016.parquet`: `33304` rows, `4533711` file bytes (4.32 MiB), `22917642` physical bytes (21.86 MiB), `23042041` encoded bytes (21.97 MiB), `4503486` compressed data bytes (4.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00017.parquet`: `32421` rows, `4561227` file bytes (4.35 MiB), `22447324` physical bytes (21.41 MiB), `22572092` encoded bytes (21.53 MiB), `4530995` compressed data bytes (4.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00018.parquet`: `33024` rows, `4583505` file bytes (4.37 MiB), `21052416` physical bytes (20.08 MiB), `21201848` encoded bytes (20.22 MiB), `4553201` compressed data bytes (4.34 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00019.parquet`: `32752` rows, `4635059` file bytes (4.42 MiB), `21212519` physical bytes (20.23 MiB), `21351182` encoded bytes (20.36 MiB), `4604533` compressed data bytes (4.39 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00020.parquet`: `32246` rows, `4599708` file bytes (4.39 MiB), `20768577` physical bytes (19.81 MiB), `20912441` encoded bytes (19.94 MiB), `4569041` compressed data bytes (4.36 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00021.parquet`: `33129` rows, `4639220` file bytes (4.42 MiB), `21161128` physical bytes (20.18 MiB), `21305785` encoded bytes (20.32 MiB), `4608529` compressed data bytes (4.40 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00022.parquet`: `32739` rows, `4655625` file bytes (4.44 MiB), `21141221` physical bytes (20.16 MiB), `21286184` encoded bytes (20.30 MiB), `4625163` compressed data bytes (4.41 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00023.parquet`: `32955` rows, `4608270` file bytes (4.39 MiB), `21210480` physical bytes (20.23 MiB), `21353162` encoded bytes (20.36 MiB), `4577404` compressed data bytes (4.37 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00024.parquet`: `32766` rows, `4605906` file bytes (4.39 MiB), `21046959` physical bytes (20.07 MiB), `21187877` encoded bytes (20.21 MiB), `4575759` compressed data bytes (4.36 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00025.parquet`: `32797` rows, `4652708` file bytes (4.44 MiB), `21149039` physical bytes (20.17 MiB), `21291068` encoded bytes (20.30 MiB), `4621919` compressed data bytes (4.41 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00026.parquet`: `33021` rows, `4561330` file bytes (4.35 MiB), `20983488` physical bytes (20.01 MiB), `21127163` encoded bytes (20.15 MiB), `4531257` compressed data bytes (4.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00027.parquet`: `32932` rows, `4595375` file bytes (4.38 MiB), `20999318` physical bytes (20.03 MiB), `21140035` encoded bytes (20.16 MiB), `4564600` compressed data bytes (4.35 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00028.parquet`: `32994` rows, `4551411` file bytes (4.34 MiB), `21034486` physical bytes (20.06 MiB), `21179928` encoded bytes (20.20 MiB), `4520972` compressed data bytes (4.31 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00029.parquet`: `33211` rows, `4607406` file bytes (4.39 MiB), `21196925` physical bytes (20.21 MiB), `21344709` encoded bytes (20.36 MiB), `4576931` compressed data bytes (4.36 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain/part-00030.parquet`: `10197` rows, `1441582` file bytes (1.37 MiB), `6625811` physical bytes (6.32 MiB), `6670958` encoded bytes (6.36 MiB), `1425165` compressed data bytes (1.36 MiB)
