# ClickBench Parquet Experiment

- Started: `2026-07-03T14:57:54-04:00`
- Write elapsed: `11.271s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `721224961` (687.81 MiB)
- Compressed column data bytes after codec compression: `131218811` (125.14 MiB)
- Parquet file bytes: `132178929` (126.06 MiB)
- Physical/encoded ratio: `0.988x`
- Encoded/compressed-data ratio: `5.496x`
- Physical/compressed-data ratio: `5.429x`
- Physical/parquet-file ratio: `5.390x`
- Files: `31`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `delta-length-byte-array`
- Date encoding: `plain`
- Timestamp encoding: `plain`
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
- Rows read and compared: `1000000`
- Files read: `31`
- Elapsed: `6.967s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004878` (7.63 MiB) | `8005305` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `262100` (255.96 KiB) | `0.999x` | `15.276x` | `15.261x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:505` | `1000000` | `138409995` (132.00 MiB) | `140035329` (133.55 MiB) | `21330632` (20.34 MiB) | `0.988x` | `6.565x` | `6.489x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204380` (199.59 KiB) | `0.999x` | `19.590x` | `19.571x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004877` (7.63 MiB) | `4282936` (4.08 MiB) | `0.999x` | `1.869x` | `1.868x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204381` (199.59 KiB) | `0.999x` | `19.590x` | `19.571x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `204401` (199.61 KiB) | `0.999x` | `19.588x` | `19.569x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `719461` (702.60 KiB) | `0.999x` | `5.565x` | `5.560x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003845` (3.82 MiB) | `396522` (387.23 KiB) | `0.999x` | `10.097x` | `10.088x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004876` (7.63 MiB) | `1084964` (1.03 MiB) | `0.999x` | `7.378x` | `7.374x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `204209` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `318052` (310.60 KiB) | `0.999x` | `12.589x` | `12.577x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `342280` (334.26 KiB) | `0.999x` | `11.698x` | `11.686x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:342` | `1000000` | `88562192` (84.46 MiB) | `89784410` (85.63 MiB) | `20793159` (19.83 MiB) | `0.986x` | `4.318x` | `4.259x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:298` | `1000000` | `79583339` (75.90 MiB) | `80832605` (77.09 MiB) | `19459948` (18.56 MiB) | `0.985x` | `4.154x` | `4.090x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003872` (3.82 MiB) | `492025` (480.49 KiB) | `0.999x` | `8.138x` | `8.130x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003870` (3.82 MiB) | `509051` (497.12 KiB) | `0.999x` | `7.865x` | `7.858x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003857` (3.82 MiB) | `458145` (447.41 KiB) | `0.999x` | `8.739x` | `8.731x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `272025` (265.65 KiB) | `0.999x` | `14.719x` | `14.705x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `232260` (226.82 KiB) | `0.999x` | `17.239x` | `17.222x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `373575` (364.82 KiB) | `0.999x` | `10.718x` | `10.707x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `373255` (364.51 KiB) | `0.999x` | `10.727x` | `10.717x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `283049` (276.42 KiB) | `0.999x` | `14.145x` | `14.132x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `245380` (239.63 KiB) | `0.999x` | `16.317x` | `16.301x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `328477` (320.78 KiB) | `0.999x` | `12.189x` | `12.177x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3354477` (3.20 MiB) | `3707607` (3.54 MiB) | `433313` (423.16 KiB) | `0.905x` | `8.556x` | `7.741x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `219196` (214.06 KiB) | `0.999x` | `18.266x` | `18.249x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `217360` (212.27 KiB) | `0.999x` | `18.420x` | `18.403x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `365135` (356.58 KiB) | `0.999x` | `10.965x` | `10.955x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3767530` (3.59 MiB) | `4017088` (3.83 MiB) | `327938` (320.25 KiB) | `0.938x` | `12.250x` | `11.489x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `205010` (200.21 KiB) | `0.999x` | `19.530x` | `19.511x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204998` (200.19 KiB) | `0.999x` | `19.531x` | `19.512x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `218511` (213.39 KiB) | `0.999x` | `18.323x` | `18.306x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `215462` (210.41 KiB) | `0.999x` | `18.583x` | `18.565x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `81583` (79.67 KiB) | `232073` (226.63 KiB) | `42859` (41.85 KiB) | `0.352x` | `5.415x` | `1.904x` | `81583` (79.67 KiB) |
| `Params` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `0` (0 B) | `41860` (40.88 KiB) | `4896` (4.78 KiB) | `0.000x` | `8.550x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003876` (3.82 MiB) | `560892` (547.75 KiB) | `0.999x` | `7.138x` | `7.131x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003870` (3.82 MiB) | `526838` (514.49 KiB) | `0.999x` | `7.600x` | `7.592x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `293173` (286.30 KiB) | `0.999x` | `13.657x` | `13.644x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3528017` (3.36 MiB) | `4225613` (4.03 MiB) | `1015186` (991.39 KiB) | `0.835x` | `4.162x` | `3.475x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `220370` (215.21 KiB) | `0.999x` | `18.169x` | `18.151x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003876` (3.82 MiB) | `505852` (494.00 KiB) | `0.999x` | `7.915x` | `7.907x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003868` (3.82 MiB) | `516785` (504.67 KiB) | `0.999x` | `7.748x` | `7.740x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003882` (3.82 MiB) | `551793` (538.86 KiB) | `0.999x` | `7.256x` | `7.249x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `285699` (279.00 KiB) | `0.999x` | `14.014x` | `14.001x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004877` (7.63 MiB) | `4229262` (4.03 MiB) | `0.999x` | `1.893x` | `1.892x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `287274` (280.54 KiB) | `0.999x` | `13.937x` | `13.924x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `269759` (263.44 KiB) | `0.999x` | `14.842x` | `14.828x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `304934` (297.79 KiB) | `0.999x` | `13.130x` | `13.118x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `204246` (199.46 KiB) | `0.999x` | `19.603x` | `19.584x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:91` | `1000000` | `13587860` (12.96 MiB) | `13657791` (13.03 MiB) | `699176` (682.79 KiB) | `0.995x` | `19.534x` | `19.434x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `205234` (200.42 KiB) | `0.999x` | `19.509x` | `19.490x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `256514` (250.50 KiB) | `0.999x` | `15.609x` | `15.594x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `206696` (201.85 KiB) | `0.999x` | `19.371x` | `19.352x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003844` (3.82 MiB) | `235455` (229.94 KiB) | `0.999x` | `17.005x` | `16.988x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004874` (7.63 MiB) | `1151665` (1.10 MiB) | `0.999x` | `6.951x` | `6.946x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:142` | `1000000` | `27797671` (26.51 MiB) | `28799318` (27.47 MiB) | `7056438` (6.73 MiB) | `0.965x` | `4.081x` | `3.939x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `3688243` (3.52 MiB) | `0.999x` | `1.086x` | `1.085x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `204209` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `204209` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `204209` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `295696` (288.77 KiB) | `0.999x` | `13.540x` | `13.527x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `204209` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `1000000` (976.56 KiB) | `1043305` (1018.85 KiB) | `77129` (75.32 KiB) | `0.958x` | `13.527x` | `12.965x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004878` (7.63 MiB) | `4284692` (4.09 MiB) | `0.999x` | `1.868x` | `1.867x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `333985` (326.16 KiB) | `0.999x` | `11.988x` | `11.977x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `302637` (295.54 KiB) | `0.999x` | `13.230x` | `13.217x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `318285` (310.83 KiB) | `0.999x` | `12.579x` | `12.567x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `385763` (376.72 KiB) | `0.999x` | `10.379x` | `10.369x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `381797` (372.85 KiB) | `0.999x` | `10.487x` | `10.477x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `707203` (690.63 KiB) | `0.999x` | `5.662x` | `5.656x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003845` (3.82 MiB) | `295376` (288.45 KiB) | `0.999x` | `13.555x` | `13.542x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204208` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `258423` (252.37 KiB) | `0.999x` | `15.493x` | `15.478x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `2001192` (1.91 MiB) | `2051444` (1.96 MiB) | `125417` (122.48 KiB) | `0.976x` | `16.357x` | `15.956x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3325142` (3.17 MiB) | `3638240` (3.47 MiB) | `352592` (344.33 KiB) | `0.914x` | `10.319x` | `9.431x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `0` (0 B) | `41860` (40.88 KiB) | `4896` (4.78 KiB) | `0.000x` | `8.550x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `0` (0 B) | `41860` (40.88 KiB) | `4896` (4.78 KiB) | `0.000x` | `8.550x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `204209` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003844` (3.82 MiB) | `269176` (262.87 KiB) | `0.999x` | `14.874x` | `14.860x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `327371` (319.70 KiB) | `0.999x` | `12.230x` | `12.219x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003877` (3.82 MiB) | `564879` (551.64 KiB) | `0.999x` | `7.088x` | `7.081x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `1726025` (1.65 MiB) | `0.999x` | `2.320x` | `2.317x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `1283389` (1.22 MiB) | `0.999x` | `3.120x` | `3.117x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003897` (3.82 MiB) | `809784` (790.80 KiB) | `0.999x` | `4.944x` | `4.940x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `204600` (199.80 KiB) | `0.999x` | `19.569x` | `19.550x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `1024` (1.00 KiB) | `46327` (45.24 KiB) | `7285` (7.11 KiB) | `0.022x` | `6.359x` | `0.141x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004816` (7.63 MiB) | `405185` (395.69 KiB) | `0.999x` | `19.756x` | `19.744x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `0` (0 B) | `41860` (40.88 KiB) | `4896` (4.78 KiB) | `0.000x` | `8.550x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3000000` (2.86 MiB) | `3044375` (2.90 MiB) | `157361` (153.67 KiB) | `0.985x` | `19.346x` | `19.064x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `204209` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `58030` (56.67 KiB) | `216392` (211.32 KiB) | `37030` (36.16 KiB) | `0.268x` | `5.844x` | `1.567x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `22051` (21.53 KiB) | `122805` (119.93 KiB) | `28740` (28.07 KiB) | `0.180x` | `4.273x` | `0.767x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `25445` (24.85 KiB) | `130318` (127.26 KiB) | `34755` (33.94 KiB) | `0.195x` | `3.750x` | `0.732x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `48191` (47.06 KiB) | `155561` (151.92 KiB) | `24598` (24.02 KiB) | `0.310x` | `6.324x` | `1.959x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `49433` (48.27 KiB) | `188446` (184.03 KiB) | `41830` (40.85 KiB) | `0.262x` | `4.505x` | `1.182x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `16873` (16.48 KiB) | `133753` (130.62 KiB) | `28972` (28.29 KiB) | `0.126x` | `4.617x` | `0.582x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `91870` (89.72 KiB) | `253882` (247.93 KiB) | `57847` (56.49 KiB) | `0.362x` | `4.389x` | `1.588x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `13001` (12.70 KiB) | `94606` (92.39 KiB) | `22135` (21.62 KiB) | `0.137x` | `4.274x` | `0.587x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `28101` (27.44 KiB) | `129337` (126.31 KiB) | `25453` (24.86 KiB) | `0.217x` | `5.081x` | `1.104x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `45607` (44.54 KiB) | `211288` (206.34 KiB) | `46276` (45.19 KiB) | `0.216x` | `4.566x` | `0.986x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `213622` (208.62 KiB) | `0.999x` | `18.743x` | `18.725x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004874` (7.63 MiB) | `3641371` (3.47 MiB) | `0.999x` | `2.198x` | `2.197x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004877` (7.63 MiB) | `4387197` (4.18 MiB) | `0.999x` | `1.825x` | `1.823x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204651` (199.85 KiB) | `0.999x` | `19.564x` | `19.545x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00000.parquet`: `33119` rows, `4203971` file bytes (4.01 MiB), `25720845` physical bytes (24.53 MiB), `26036573` encoded bytes (24.83 MiB), `4172025` compressed data bytes (3.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00001.parquet`: `33167` rows, `4152569` file bytes (3.96 MiB), `25506057` physical bytes (24.32 MiB), `25815027` encoded bytes (24.62 MiB), `4120243` compressed data bytes (3.93 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00002.parquet`: `32697` rows, `4143882` file bytes (3.95 MiB), `25595573` physical bytes (24.41 MiB), `25906190` encoded bytes (24.71 MiB), `4111060` compressed data bytes (3.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00003.parquet`: `32955` rows, `4141848` file bytes (3.95 MiB), `25547953` physical bytes (24.36 MiB), `25858680` encoded bytes (24.66 MiB), `4109470` compressed data bytes (3.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00004.parquet`: `33163` rows, `4158205` file bytes (3.97 MiB), `25752495` physical bytes (24.56 MiB), `26068667` encoded bytes (24.86 MiB), `4125061` compressed data bytes (3.93 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00005.parquet`: `33393` rows, `4117258` file bytes (3.93 MiB), `25773610` physical bytes (24.58 MiB), `26082765` encoded bytes (24.87 MiB), `4084851` compressed data bytes (3.90 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00006.parquet`: `32963` rows, `4111557` file bytes (3.92 MiB), `25358705` physical bytes (24.18 MiB), `25669725` encoded bytes (24.48 MiB), `4079408` compressed data bytes (3.89 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00007.parquet`: `33213` rows, `4150256` file bytes (3.96 MiB), `25765774` physical bytes (24.57 MiB), `26077039` encoded bytes (24.87 MiB), `4117742` compressed data bytes (3.93 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00008.parquet`: `33043` rows, `4114479` file bytes (3.92 MiB), `25568606` physical bytes (24.38 MiB), `25881865` encoded bytes (24.68 MiB), `4081604` compressed data bytes (3.89 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00009.parquet`: `33158` rows, `4104586` file bytes (3.91 MiB), `25580346` physical bytes (24.40 MiB), `25893490` encoded bytes (24.69 MiB), `4072210` compressed data bytes (3.88 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00010.parquet`: `33058` rows, `4133245` file bytes (3.94 MiB), `25555443` physical bytes (24.37 MiB), `25865339` encoded bytes (24.67 MiB), `4101234` compressed data bytes (3.91 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00011.parquet`: `33063` rows, `4053638` file bytes (3.87 MiB), `25420048` physical bytes (24.24 MiB), `25729294` encoded bytes (24.54 MiB), `4021439` compressed data bytes (3.84 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00012.parquet`: `33230` rows, `4081328` file bytes (3.89 MiB), `25600474` physical bytes (24.41 MiB), `25913507` encoded bytes (24.71 MiB), `4049069` compressed data bytes (3.86 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00013.parquet`: `33277` rows, `4127157` file bytes (3.94 MiB), `25569189` physical bytes (24.38 MiB), `25880355` encoded bytes (24.68 MiB), `4095181` compressed data bytes (3.91 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00014.parquet`: `32747` rows, `4049148` file bytes (3.86 MiB), `25240106` physical bytes (24.07 MiB), `25546261` encoded bytes (24.36 MiB), `4016960` compressed data bytes (3.83 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00015.parquet`: `33435` rows, `4430356` file bytes (4.23 MiB), `24022868` physical bytes (22.91 MiB), `24303103` encoded bytes (23.18 MiB), `4399151` compressed data bytes (4.20 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00016.parquet`: `33317` rows, `4540893` file bytes (4.33 MiB), `22917517` physical bytes (21.86 MiB), `23173296` encoded bytes (22.10 MiB), `4510633` compressed data bytes (4.30 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00017.parquet`: `32435` rows, `4565717` file bytes (4.35 MiB), `22441811` physical bytes (21.40 MiB), `22694753` encoded bytes (21.64 MiB), `4535494` compressed data bytes (4.33 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00018.parquet`: `33009` rows, `4591820` file bytes (4.38 MiB), `21048664` physical bytes (20.07 MiB), `21328773` encoded bytes (20.34 MiB), `4561508` compressed data bytes (4.35 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00019.parquet`: `32738` rows, `4643665` file bytes (4.43 MiB), `21213946` physical bytes (20.23 MiB), `21481159` encoded bytes (20.49 MiB), `4613132` compressed data bytes (4.40 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00020.parquet`: `32258` rows, `4605255` file bytes (4.39 MiB), `20763341` physical bytes (19.80 MiB), `21035246` encoded bytes (20.06 MiB), `4574549` compressed data bytes (4.36 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00021.parquet`: `33113` rows, `4644510` file bytes (4.43 MiB), `21158805` physical bytes (20.18 MiB), `21433059` encoded bytes (20.44 MiB), `4613814` compressed data bytes (4.40 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00022.parquet`: `32752` rows, `4657447` file bytes (4.44 MiB), `21137313` physical bytes (20.16 MiB), `21411828` encoded bytes (20.42 MiB), `4626887` compressed data bytes (4.41 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00023.parquet`: `32939` rows, `4616241` file bytes (4.40 MiB), `21203293` physical bytes (20.22 MiB), `21476889` encoded bytes (20.48 MiB), `4585368` compressed data bytes (4.37 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00024.parquet`: `32750` rows, `4614434` file bytes (4.40 MiB), `21040284` physical bytes (20.07 MiB), `21312128` encoded bytes (20.32 MiB), `4584279` compressed data bytes (4.37 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00025.parquet`: `32803` rows, `4656840` file bytes (4.44 MiB), `21153884` physical bytes (20.17 MiB), `21426324` encoded bytes (20.43 MiB), `4626045` compressed data bytes (4.41 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00026.parquet`: `33278` rows, `4600811` file bytes (4.39 MiB), `21142972` physical bytes (20.16 MiB), `21417895` encoded bytes (20.43 MiB), `4570688` compressed data bytes (4.36 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00027.parquet`: `32921` rows, `4603290` file bytes (4.39 MiB), `20990285` physical bytes (20.02 MiB), `21261384` encoded bytes (20.28 MiB), `4572647` compressed data bytes (4.36 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00028.parquet`: `33003` rows, `4558247` file bytes (4.35 MiB), `21031024` physical bytes (20.06 MiB), `21305947` encoded bytes (20.32 MiB), `4527695` compressed data bytes (4.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00029.parquet`: `33217` rows, `4610709` file bytes (4.40 MiB), `21190180` physical bytes (20.21 MiB), `21468858` encoded bytes (20.47 MiB), `4580219` compressed data bytes (4.37 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain/part-00030.parquet`: `9786` rows, `1395567` file bytes (1.33 MiB), `6387213` physical bytes (6.09 MiB), `6469542` encoded bytes (6.17 MiB), `1379145` compressed data bytes (1.32 MiB)
