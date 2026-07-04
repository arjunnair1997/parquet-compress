# ClickBench Parquet Experiment

- Started: `2026-07-03T23:30:02-04:00`
- Write elapsed: `11.329s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `514971711` (491.12 MiB)
- Compressed column data bytes after codec compression: `120633530` (115.05 MiB)
- Parquet file bytes: `121563138` (115.93 MiB)
- Physical/encoded ratio: `1.383x`
- Encoded/compressed-data ratio: `4.269x`
- Physical/compressed-data ratio: `5.905x`
- Physical/parquet-file ratio: `5.860x`
- Files: `30`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `delta-byte-array`
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
- Files read: `30`
- Elapsed: `7.067s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004718` (7.63 MiB) | `8005131` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `261957` (255.82 KiB) | `0.999x` | `15.284x` | `15.270x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:513` | `1000000` | `138409995` (132.00 MiB) | `64480578` (61.49 MiB) | `17110242` (16.32 MiB) | `2.147x` | `3.769x` | `8.089x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `204229` (199.44 KiB) | `0.999x` | `19.604x` | `19.586x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004714` (7.63 MiB) | `4282417` (4.08 MiB) | `0.999x` | `1.869x` | `1.868x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:59` | `1000000` | `4000000` (3.81 MiB) | `51375` (50.17 KiB) | `7913` (7.73 KiB) | `77.859x` | `6.492x` | `505.497x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003717` (3.82 MiB) | `204259` (199.47 KiB) | `0.999x` | `19.601x` | `19.583x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003770` (3.82 MiB) | `719078` (702.22 KiB) | `0.999x` | `5.568x` | `5.563x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003720` (3.82 MiB) | `396370` (387.08 KiB) | `0.999x` | `10.101x` | `10.092x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004718` (7.63 MiB) | `1084731` (1.03 MiB) | `0.999x` | `7.379x` | `7.375x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `204062` (199.28 KiB) | `0.999x` | `19.620x` | `19.602x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003710` (3.82 MiB) | `317973` (310.52 KiB) | `0.999x` | `12.591x` | `12.580x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `342049` (334.03 KiB) | `0.999x` | `11.705x` | `11.694x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:340` | `1000000` | `88562192` (84.46 MiB) | `40460999` (38.59 MiB) | `18084728` (17.25 MiB) | `2.189x` | `2.237x` | `4.897x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:309` | `1000000` | `79583339` (75.90 MiB) | `38989764` (37.18 MiB) | `17325529` (16.52 MiB) | `2.041x` | `2.250x` | `4.593x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003744` (3.82 MiB) | `491836` (480.31 KiB) | `0.999x` | `8.140x` | `8.133x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003743` (3.82 MiB) | `508954` (497.03 KiB) | `0.999x` | `7.867x` | `7.859x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003736` (3.82 MiB) | `457949` (447.22 KiB) | `0.999x` | `8.743x` | `8.735x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `271864` (265.49 KiB) | `0.999x` | `14.727x` | `14.713x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003711` (3.82 MiB) | `232134` (226.69 KiB) | `0.999x` | `17.247x` | `17.231x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `373371` (364.62 KiB) | `0.999x` | `10.723x` | `10.713x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003717` (3.82 MiB) | `373104` (364.36 KiB) | `0.999x` | `10.731x` | `10.721x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `282833` (276.20 KiB) | `0.999x` | `14.156x` | `14.143x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `245222` (239.47 KiB) | `0.999x` | `16.327x` | `16.312x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `328328` (320.63 KiB) | `0.999x` | `12.194x` | `12.183x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `3354477` (3.20 MiB) | `1041082` (1016.68 KiB) | `394774` (385.52 KiB) | `3.222x` | `2.637x` | `8.497x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `219079` (213.94 KiB) | `0.999x` | `18.275x` | `18.258x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `217243` (212.15 KiB) | `0.999x` | `18.430x` | `18.413x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003717` (3.82 MiB) | `364844` (356.29 KiB) | `0.999x` | `10.974x` | `10.964x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `3767530` (3.59 MiB) | `860316` (840.15 KiB) | `236318` (230.78 KiB) | `4.379x` | `3.641x` | `15.943x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `204957` (200.15 KiB) | `0.999x` | `19.534x` | `19.516x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `204943` (200.14 KiB) | `0.999x` | `19.536x` | `19.518x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `218352` (213.23 KiB) | `0.999x` | `18.336x` | `18.319x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `215287` (210.24 KiB) | `0.999x` | `18.597x` | `18.580x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `81583` (79.67 KiB) | `281681` (275.08 KiB) | `56765` (55.43 KiB) | `0.290x` | `4.962x` | `1.437x` | `81583` (79.67 KiB) |
| `Params` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `0` (0 B) | `81427` (79.52 KiB) | `6905` (6.74 KiB) | `0.000x` | `11.792x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003755` (3.82 MiB) | `560148` (547.02 KiB) | `0.999x` | `7.148x` | `7.141x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003754` (3.82 MiB) | `526613` (514.27 KiB) | `0.999x` | `7.603x` | `7.596x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `293060` (286.19 KiB) | `0.999x` | `13.662x` | `13.649x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `3528017` (3.36 MiB) | `2961687` (2.82 MiB) | `994375` (971.07 KiB) | `1.191x` | `2.978x` | `3.548x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `220197` (215.04 KiB) | `0.999x` | `18.182x` | `18.166x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003747` (3.82 MiB) | `505684` (493.83 KiB) | `0.999x` | `7.917x` | `7.910x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003752` (3.82 MiB) | `516182` (504.08 KiB) | `0.999x` | `7.756x` | `7.749x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003758` (3.82 MiB) | `551349` (538.43 KiB) | `0.999x` | `7.262x` | `7.255x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `285549` (278.86 KiB) | `0.999x` | `14.021x` | `14.008x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004718` (7.63 MiB) | `4229255` (4.03 MiB) | `0.999x` | `1.893x` | `1.892x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `287153` (280.42 KiB) | `0.999x` | `13.943x` | `13.930x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `269661` (263.34 KiB) | `0.999x` | `14.847x` | `14.833x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `304662` (297.52 KiB) | `0.999x` | `13.142x` | `13.129x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `204100` (199.32 KiB) | `0.999x` | `19.616x` | `19.598x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:89` | `1000000` | `13587860` (12.96 MiB) | `142663` (139.32 KiB) | `26939` (26.31 KiB) | `95.244x` | `5.296x` | `504.394x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `205305` (200.49 KiB) | `0.999x` | `19.501x` | `19.483x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `256373` (250.36 KiB) | `0.999x` | `15.617x` | `15.602x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `206526` (201.69 KiB) | `0.999x` | `19.386x` | `19.368x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003720` (3.82 MiB) | `235269` (229.75 KiB) | `0.999x` | `17.018x` | `17.002x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004714` (7.63 MiB) | `1151408` (1.10 MiB) | `0.999x` | `6.952x` | `6.948x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:148` | `1000000` | `27797671` (26.51 MiB) | `21056193` (20.08 MiB) | `6754150` (6.44 MiB) | `1.320x` | `3.118x` | `4.116x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `3689970` (3.52 MiB) | `0.999x` | `1.085x` | `1.084x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `204062` (199.28 KiB) | `0.999x` | `19.620x` | `19.602x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `204062` (199.28 KiB) | `0.999x` | `19.620x` | `19.602x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `204062` (199.28 KiB) | `0.999x` | `19.620x` | `19.602x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003719` (3.82 MiB) | `295526` (288.60 KiB) | `0.999x` | `13.548x` | `13.535x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `204062` (199.28 KiB) | `0.999x` | `19.620x` | `19.602x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `1000000` (976.56 KiB) | `207173` (202.32 KiB) | `74334` (72.59 KiB) | `4.827x` | `2.787x` | `13.453x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004716` (7.63 MiB) | `4283734` (4.09 MiB) | `0.999x` | `1.869x` | `1.868x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `333934` (326.11 KiB) | `0.999x` | `11.990x` | `11.978x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `302532` (295.44 KiB) | `0.999x` | `13.234x` | `13.222x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003711` (3.82 MiB) | `318225` (310.77 KiB) | `0.999x` | `12.581x` | `12.570x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `385336` (376.30 KiB) | `0.999x` | `10.390x` | `10.381x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003720` (3.82 MiB) | `381604` (372.66 KiB) | `0.999x` | `10.492x` | `10.482x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `706757` (690.19 KiB) | `0.999x` | `5.665x` | `5.660x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `295234` (288.31 KiB) | `0.999x` | `13.561x` | `13.549x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `204061` (199.28 KiB) | `0.999x` | `19.620x` | `19.602x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003719` (3.82 MiB) | `258298` (252.24 KiB) | `0.999x` | `15.500x` | `15.486x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `2001192` (1.91 MiB) | `333442` (325.63 KiB) | `89945` (87.84 KiB) | `6.002x` | `3.707x` | `22.249x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `3325142` (3.17 MiB) | `960601` (938.09 KiB) | `253937` (247.99 KiB) | `3.462x` | `3.783x` | `13.094x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `0` (0 B) | `81427` (79.52 KiB) | `6905` (6.74 KiB) | `0.000x` | `11.792x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `0` (0 B) | `81427` (79.52 KiB) | `6905` (6.74 KiB) | `0.000x` | `11.792x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `204062` (199.28 KiB) | `0.999x` | `19.620x` | `19.602x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003720` (3.82 MiB) | `268901` (262.60 KiB) | `0.999x` | `14.889x` | `14.875x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003717` (3.82 MiB) | `326996` (319.33 KiB) | `0.999x` | `12.244x` | `12.233x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003751` (3.82 MiB) | `564355` (551.13 KiB) | `0.999x` | `7.094x` | `7.088x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003772` (3.82 MiB) | `1721105` (1.64 MiB) | `0.999x` | `2.326x` | `2.324x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003771` (3.82 MiB) | `1280655` (1.22 MiB) | `0.999x` | `3.126x` | `3.123x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003770` (3.82 MiB) | `808626` (789.67 KiB) | `0.999x` | `4.951x` | `4.947x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `204454` (199.66 KiB) | `0.999x` | `19.582x` | `19.564x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `1024` (1.00 KiB) | `86814` (84.78 KiB) | `9646` (9.42 KiB) | `0.012x` | `9.000x` | `0.106x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004654` (7.63 MiB) | `405011` (395.52 KiB) | `0.999x` | `19.764x` | `19.753x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `0` (0 B) | `81427` (79.52 KiB) | `6905` (6.74 KiB) | `0.000x` | `11.792x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `3000000` (2.86 MiB) | `86274` (84.25 KiB) | `10304` (10.06 KiB) | `34.773x` | `8.373x` | `291.149x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `204062` (199.28 KiB) | `0.999x` | `19.620x` | `19.602x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `58030` (56.67 KiB) | `301718` (294.65 KiB) | `48221` (47.09 KiB) | `0.192x` | `6.257x` | `1.203x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `22051` (21.53 KiB) | `198433` (193.78 KiB) | `39049` (38.13 KiB) | `0.111x` | `5.082x` | `0.565x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `25445` (24.85 KiB) | `210085` (205.16 KiB) | `44966` (43.91 KiB) | `0.121x` | `4.672x` | `0.566x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `48191` (47.06 KiB) | `232757` (227.30 KiB) | `34215` (33.41 KiB) | `0.207x` | `6.803x` | `1.408x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `49433` (48.27 KiB) | `267386` (261.12 KiB) | `51796` (50.58 KiB) | `0.185x` | `5.162x` | `0.954x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `16873` (16.48 KiB) | `213108` (208.11 KiB) | `38896` (37.98 KiB) | `0.079x` | `5.479x` | `0.434x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `91870` (89.72 KiB) | `327495` (319.82 KiB) | `68186` (66.59 KiB) | `0.281x` | `4.803x` | `1.347x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `13001` (12.70 KiB) | `150188` (146.67 KiB) | `27791` (27.14 KiB) | `0.087x` | `5.404x` | `0.468x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `28101` (27.44 KiB) | `190573` (186.11 KiB) | `31132` (30.40 KiB) | `0.147x` | `6.121x` | `0.903x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `45607` (44.54 KiB) | `261720` (255.59 KiB) | `50279` (49.10 KiB) | `0.174x` | `5.205x` | `0.907x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `213439` (208.44 KiB) | `0.999x` | `18.758x` | `18.741x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004717` (7.63 MiB) | `3638801` (3.47 MiB) | `0.999x` | `2.200x` | `2.199x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004717` (7.63 MiB) | `4384004` (4.18 MiB) | `0.999x` | `1.826x` | `1.825x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `204500` (199.71 KiB) | `0.999x` | `19.578x` | `19.560x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00000.parquet`: `34160` rows, `3802452` file bytes (3.63 MiB), `26536032` physical bytes (25.31 MiB), `17031906` encoded bytes (16.24 MiB), `3770540` compressed data bytes (3.60 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00001.parquet`: `34470` rows, `3770074` file bytes (3.60 MiB), `26521657` physical bytes (25.29 MiB), `16965370` encoded bytes (16.18 MiB), `3737768` compressed data bytes (3.56 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00002.parquet`: `33971` rows, `3782788` file bytes (3.61 MiB), `26580917` physical bytes (25.35 MiB), `16970387` encoded bytes (16.18 MiB), `3749687` compressed data bytes (3.58 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00003.parquet`: `34083` rows, `3710084` file bytes (3.54 MiB), `26377676` physical bytes (25.16 MiB), `16820187` encoded bytes (16.04 MiB), `3678253` compressed data bytes (3.51 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00004.parquet`: `34140` rows, `3757259` file bytes (3.58 MiB), `26560850` physical bytes (25.33 MiB), `16932507` encoded bytes (16.15 MiB), `3724162` compressed data bytes (3.55 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00005.parquet`: `34704` rows, `3759914` file bytes (3.59 MiB), `26772540` physical bytes (25.53 MiB), `17088009` encoded bytes (16.30 MiB), `3727611` compressed data bytes (3.55 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00006.parquet`: `34512` rows, `3765884` file bytes (3.59 MiB), `26527433` physical bytes (25.30 MiB), `16981950` encoded bytes (16.20 MiB), `3733370` compressed data bytes (3.56 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00007.parquet`: `34254` rows, `3748621` file bytes (3.57 MiB), `26592785` physical bytes (25.36 MiB), `16925020` encoded bytes (16.14 MiB), `3716044` compressed data bytes (3.54 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00008.parquet`: `34261` rows, `3731475` file bytes (3.56 MiB), `26582018` physical bytes (25.35 MiB), `16982300` encoded bytes (16.20 MiB), `3698980` compressed data bytes (3.53 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00009.parquet`: `34144` rows, `3683823` file bytes (3.51 MiB), `26179770` physical bytes (24.97 MiB), `16797902` encoded bytes (16.02 MiB), `3651463` compressed data bytes (3.48 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00010.parquet`: `34483` rows, `3766897` file bytes (3.59 MiB), `26720087` physical bytes (25.48 MiB), `17051401` encoded bytes (16.26 MiB), `3734807` compressed data bytes (3.56 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00011.parquet`: `34797` rows, `3712350` file bytes (3.54 MiB), `26784057` physical bytes (25.54 MiB), `17062698` encoded bytes (16.27 MiB), `3679825` compressed data bytes (3.51 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00012.parquet`: `34367` rows, `3691236` file bytes (3.52 MiB), `26383630` physical bytes (25.16 MiB), `16881841` encoded bytes (16.10 MiB), `3659202` compressed data bytes (3.49 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00013.parquet`: `34051` rows, `3711622` file bytes (3.54 MiB), `26117600` physical bytes (24.91 MiB), `16788614` encoded bytes (16.01 MiB), `3679494` compressed data bytes (3.51 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00014.parquet`: `34779` rows, `3946193` file bytes (3.76 MiB), `26398240` physical bytes (25.18 MiB), `17234022` encoded bytes (16.44 MiB), `3913924` compressed data bytes (3.73 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00015.parquet`: `34302` rows, `4295958` file bytes (4.10 MiB), `23438440` physical bytes (22.35 MiB), `17228479` encoded bytes (16.43 MiB), `4265594` compressed data bytes (4.07 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00016.parquet`: `33584` rows, `4352483` file bytes (4.15 MiB), `23717846` physical bytes (22.62 MiB), `16946178` encoded bytes (16.16 MiB), `4322217` compressed data bytes (4.12 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00017.parquet`: `33697` rows, `4496795` file bytes (4.29 MiB), `21665068` physical bytes (20.66 MiB), `18017599` encoded bytes (17.18 MiB), `4466386` compressed data bytes (4.26 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00018.parquet`: `33374` rows, `4536660` file bytes (4.33 MiB), `21556089` physical bytes (20.56 MiB), `18248721` encoded bytes (17.40 MiB), `4505890` compressed data bytes (4.30 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00019.parquet`: `32669` rows, `4592107` file bytes (4.38 MiB), `21309939` physical bytes (20.32 MiB), `18022328` encoded bytes (17.19 MiB), `4561809` compressed data bytes (4.35 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00020.parquet`: `33244` rows, `4475492` file bytes (4.27 MiB), `21183496` physical bytes (20.20 MiB), `18021926` encoded bytes (17.19 MiB), `4444751` compressed data bytes (4.24 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00021.parquet`: `32874` rows, `4530654` file bytes (4.32 MiB), `21139989` physical bytes (20.16 MiB), `17910229` encoded bytes (17.08 MiB), `4499852` compressed data bytes (4.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00022.parquet`: `33332` rows, `4505118` file bytes (4.30 MiB), `21337174` physical bytes (20.35 MiB), `18095081` encoded bytes (17.26 MiB), `4474329` compressed data bytes (4.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00023.parquet`: `33302` rows, `4512505` file bytes (4.30 MiB), `21318897` physical bytes (20.33 MiB), `18032762` encoded bytes (17.20 MiB), `4482146` compressed data bytes (4.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00024.parquet`: `32868` rows, `4557630` file bytes (4.35 MiB), `21349955` physical bytes (20.36 MiB), `18165803` encoded bytes (17.32 MiB), `4526830` compressed data bytes (4.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00025.parquet`: `33110` rows, `4481994` file bytes (4.27 MiB), `21148404` physical bytes (20.17 MiB), `17880851` encoded bytes (17.05 MiB), `4451822` compressed data bytes (4.25 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00026.parquet`: `33221` rows, `4486118` file bytes (4.28 MiB), `21200449` physical bytes (20.22 MiB), `17946938` encoded bytes (17.12 MiB), `4455524` compressed data bytes (4.25 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00027.parquet`: `33350` rows, `4428713` file bytes (4.22 MiB), `21235318` physical bytes (20.25 MiB), `17900193` encoded bytes (17.07 MiB), `4398168` compressed data bytes (4.19 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00028.parquet`: `33282` rows, `4438444` file bytes (4.23 MiB), `21063039` physical bytes (20.09 MiB), `17850859` encoded bytes (17.02 MiB), `4408287` compressed data bytes (4.20 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain/part-00029.parquet`: `18615` rows, `2531795` file bytes (2.41 MiB), `12099229` physical bytes (11.54 MiB), `10189650` encoded bytes (9.72 MiB), `2514795` compressed data bytes (2.40 MiB)
