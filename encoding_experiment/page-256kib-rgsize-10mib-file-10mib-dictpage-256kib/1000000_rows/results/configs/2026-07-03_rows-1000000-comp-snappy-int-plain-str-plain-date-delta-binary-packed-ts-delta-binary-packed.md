# ClickBench Parquet Experiment

- Started: `2026-07-03T23:30:34-04:00`
- Write elapsed: `11.19s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `807757787` (770.34 MiB)
- Compressed column data bytes after codec compression: `130468209` (124.42 MiB)
- Parquet file bytes: `131428587` (125.34 MiB)
- Physical/encoded ratio: `0.882x`
- Encoded/compressed-data ratio: `6.191x`
- Physical/compressed-data ratio: `5.460x`
- Physical/parquet-file ratio: `5.420x`
- Files: `31`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `plain`
- Date encoding: `delta-binary-packed`
- Timestamp encoding: `delta-binary-packed`
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
- Elapsed: `6.982s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004876` (7.63 MiB) | `8005302` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `262066` (255.92 KiB) | `0.999x` | `15.278x` | `15.263x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:504` | `1000000` | `138409995` (132.00 MiB) | `142876656` (136.26 MiB) | `20900308` (19.93 MiB) | `0.969x` | `6.836x` | `6.622x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003833` (3.82 MiB) | `204387` (199.60 KiB) | `0.999x` | `19.589x` | `19.571x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `3370446` (3.21 MiB) | `3037022` (2.90 MiB) | `2.374x` | `1.110x` | `2.634x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `51917` (50.70 KiB) | `8174` (7.98 KiB) | `77.046x` | `6.351x` | `489.356x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204410` (199.62 KiB) | `0.999x` | `19.587x` | `19.569x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `719412` (702.55 KiB) | `0.999x` | `5.566x` | `5.560x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `396659` (387.36 KiB) | `0.999x` | `10.094x` | `10.084x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004874` (7.63 MiB) | `1084918` (1.03 MiB) | `0.999x` | `7.378x` | `7.374x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204216` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `318155` (310.70 KiB) | `0.999x` | `12.585x` | `12.572x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `342328` (334.30 KiB) | `0.999x` | `11.696x` | `11.685x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:343` | `1000000` | `88562192` (84.46 MiB) | `92650307` (88.36 MiB) | `20478593` (19.53 MiB) | `0.956x` | `4.524x` | `4.325x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:299` | `1000000` | `79583339` (75.90 MiB) | `83645300` (79.77 MiB) | `19054369` (18.17 MiB) | `0.951x` | `4.390x` | `4.177x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003868` (3.82 MiB) | `492004` (480.47 KiB) | `0.999x` | `8.138x` | `8.130x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003870` (3.82 MiB) | `508995` (497.07 KiB) | `0.999x` | `7.866x` | `7.859x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003854` (3.82 MiB) | `458120` (447.38 KiB) | `0.999x` | `8.740x` | `8.731x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `272016` (265.64 KiB) | `0.999x` | `14.719x` | `14.705x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `232242` (226.80 KiB) | `0.999x` | `17.240x` | `17.223x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `373580` (364.82 KiB) | `0.999x` | `10.717x` | `10.707x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `373352` (364.60 KiB) | `0.999x` | `10.724x` | `10.714x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `283029` (276.40 KiB) | `0.999x` | `14.146x` | `14.133x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003833` (3.82 MiB) | `245412` (239.66 KiB) | `0.999x` | `16.315x` | `16.299x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `328536` (320.84 KiB) | `0.999x` | `12.187x` | `12.175x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `3354477` (3.20 MiB) | `7358093` (7.02 MiB) | `536971` (524.39 KiB) | `0.456x` | `13.703x` | `6.247x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `219199` (214.06 KiB) | `0.999x` | `18.266x` | `18.248x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `217372` (212.28 KiB) | `0.999x` | `18.419x` | `18.402x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `365237` (356.68 KiB) | `0.999x` | `10.962x` | `10.952x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `3767530` (3.59 MiB) | `7771141` (7.41 MiB) | `467409` (456.45 KiB) | `0.485x` | `16.626x` | `8.060x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003833` (3.82 MiB) | `204868` (200.07 KiB) | `0.999x` | `19.543x` | `19.525x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003832` (3.82 MiB) | `205009` (200.20 KiB) | `0.999x` | `19.530x` | `19.511x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `218519` (213.40 KiB) | `0.999x` | `18.323x` | `18.305x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `215504` (210.45 KiB) | `0.999x` | `18.579x` | `18.561x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `81583` (79.67 KiB) | `4085146` (3.90 MiB) | `220080` (214.92 KiB) | `0.020x` | `18.562x` | `0.371x` | `81583` (79.67 KiB) |
| `Params` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `0` (0 B) | `4002375` (3.82 MiB) | `202752` (198.00 KiB) | `0.000x` | `19.740x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003876` (3.82 MiB) | `560904` (547.76 KiB) | `0.999x` | `7.138x` | `7.131x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003871` (3.82 MiB) | `526856` (514.51 KiB) | `0.999x` | `7.600x` | `7.592x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `293200` (286.33 KiB) | `0.999x` | `13.656x` | `13.643x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `3528017` (3.36 MiB) | `7535557` (7.19 MiB) | `1097472` (1.05 MiB) | `0.468x` | `6.866x` | `3.215x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `220385` (215.22 KiB) | `0.999x` | `18.167x` | `18.150x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003871` (3.82 MiB) | `505917` (494.06 KiB) | `0.999x` | `7.914x` | `7.906x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003869` (3.82 MiB) | `516774` (504.66 KiB) | `0.999x` | `7.748x` | `7.740x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003882` (3.82 MiB) | `551915` (538.98 KiB) | `0.999x` | `7.255x` | `7.247x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `285721` (279.02 KiB) | `0.999x` | `14.013x` | `14.000x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `3491595` (3.33 MiB) | `3062866` (2.92 MiB) | `2.291x` | `1.140x` | `2.612x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `287306` (280.57 KiB) | `0.999x` | `13.936x` | `13.922x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `269748` (263.43 KiB) | `0.999x` | `14.843x` | `14.829x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `304979` (297.83 KiB) | `0.999x` | `13.128x` | `13.116x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204251` (199.46 KiB) | `0.999x` | `19.603x` | `19.584x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:91` | `1000000` | `13587860` (12.96 MiB) | `17595522` (16.78 MiB) | `919603` (898.05 KiB) | `0.772x` | `19.134x` | `14.776x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `205228` (200.42 KiB) | `0.999x` | `19.509x` | `19.491x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `256541` (250.53 KiB) | `0.999x` | `15.607x` | `15.592x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `206700` (201.86 KiB) | `0.999x` | `19.370x` | `19.352x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `235470` (229.95 KiB) | `0.999x` | `17.004x` | `16.987x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004871` (7.63 MiB) | `1151601` (1.10 MiB) | `0.999x` | `6.951x` | `6.947x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:144` | `1000000` | `27797671` (26.51 MiB) | `31860445` (30.38 MiB) | `7045811` (6.72 MiB) | `0.872x` | `4.522x` | `3.945x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003897` (3.82 MiB) | `3688235` (3.52 MiB) | `0.999x` | `1.086x` | `1.085x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204216` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204216` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204216` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `295687` (288.76 KiB) | `0.999x` | `13.541x` | `13.528x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204216` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `1000000` (976.56 KiB) | `5003108` (4.77 MiB) | `294410` (287.51 KiB) | `0.200x` | `16.994x` | `3.397x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `3385011` (3.23 MiB) | `3046889` (2.91 MiB) | `2.363x` | `1.111x` | `2.626x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `334034` (326.21 KiB) | `0.999x` | `11.986x` | `11.975x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `302582` (295.49 KiB) | `0.999x` | `13.232x` | `13.220x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `318296` (310.84 KiB) | `0.999x` | `12.579x` | `12.567x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `385808` (376.77 KiB) | `0.999x` | `10.378x` | `10.368x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `381822` (372.87 KiB) | `0.999x` | `10.486x` | `10.476x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003895` (3.82 MiB) | `707213` (690.64 KiB) | `0.999x` | `5.662x` | `5.656x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003846` (3.82 MiB) | `295444` (288.52 KiB) | `0.999x` | `13.552x` | `13.539x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204213` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003846` (3.82 MiB) | `258498` (252.44 KiB) | `0.999x` | `15.489x` | `15.474x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `2001192` (1.91 MiB) | `6004664` (5.73 MiB) | `321316` (313.79 KiB) | `0.333x` | `18.688x` | `6.228x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `3325142` (3.17 MiB) | `7328757` (6.99 MiB) | `457190` (446.47 KiB) | `0.454x` | `16.030x` | `7.273x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `0` (0 B) | `4002375` (3.82 MiB) | `202752` (198.00 KiB) | `0.000x` | `19.740x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `0` (0 B) | `4002375` (3.82 MiB) | `202752` (198.00 KiB) | `0.000x` | `19.740x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204216` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003845` (3.82 MiB) | `269159` (262.85 KiB) | `0.999x` | `14.875x` | `14.861x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `327395` (319.72 KiB) | `0.999x` | `12.229x` | `12.218x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003872` (3.82 MiB) | `564780` (551.54 KiB) | `0.999x` | `7.089x` | `7.082x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `1725463` (1.65 MiB) | `0.999x` | `2.320x` | `2.318x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003897` (3.82 MiB) | `1283028` (1.22 MiB) | `0.999x` | `3.121x` | `3.118x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003893` (3.82 MiB) | `810083` (791.10 KiB) | `0.999x` | `4.943x` | `4.938x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204596` (199.80 KiB) | `0.999x` | `19.569x` | `19.551x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `1024` (1.00 KiB) | `4004109` (3.82 MiB) | `204255` (199.47 KiB) | `0.000x` | `19.603x` | `0.005x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004813` (7.63 MiB) | `405188` (395.69 KiB) | `0.999x` | `19.756x` | `19.744x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `0` (0 B) | `4002375` (3.82 MiB) | `202752` (198.00 KiB) | `0.000x` | `19.740x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `3000000` (2.86 MiB) | `7003593` (6.68 MiB) | `354323` (346.02 KiB) | `0.428x` | `19.766x` | `8.467x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204216` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `58030` (56.67 KiB) | `4062840` (3.87 MiB) | `218520` (213.40 KiB) | `0.014x` | `18.593x` | `0.266x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `22051` (21.53 KiB) | `4025892` (3.84 MiB) | `214764` (209.73 KiB) | `0.005x` | `18.746x` | `0.103x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `25445` (24.85 KiB) | `4029293` (3.84 MiB) | `219139` (214.00 KiB) | `0.006x` | `18.387x` | `0.116x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `48191` (47.06 KiB) | `4052044` (3.86 MiB) | `214379` (209.35 KiB) | `0.012x` | `18.901x` | `0.225x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `49433` (48.27 KiB) | `4053494` (3.87 MiB) | `219161` (214.02 KiB) | `0.012x` | `18.496x` | `0.226x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `16873` (16.48 KiB) | `4020157` (3.83 MiB) | `214404` (209.38 KiB) | `0.004x` | `18.750x` | `0.079x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `91870` (89.72 KiB) | `4097848` (3.91 MiB) | `230233` (224.84 KiB) | `0.022x` | `17.799x` | `0.399x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `13001` (12.70 KiB) | `4016750` (3.83 MiB) | `212197` (207.22 KiB) | `0.003x` | `18.929x` | `0.061x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `28101` (27.44 KiB) | `4033131` (3.85 MiB) | `214701` (209.67 KiB) | `0.007x` | `18.785x` | `0.131x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `45607` (44.54 KiB) | `4048498` (3.86 MiB) | `221708` (216.51 KiB) | `0.011x` | `18.260x` | `0.206x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `213638` (208.63 KiB) | `0.999x` | `18.741x` | `18.723x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004874` (7.63 MiB) | `3640793` (3.47 MiB) | `0.999x` | `2.199x` | `2.197x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004875` (7.63 MiB) | `4386689` (4.18 MiB) | `0.999x` | `1.825x` | `1.824x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204651` (199.85 KiB) | `0.999x` | `19.564x` | `19.545x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00000.parquet`: `33119` rows, `4240839` file bytes (4.04 MiB), `25720845` physical bytes (24.53 MiB), `28897361` encoded bytes (27.56 MiB), `4208865` compressed data bytes (4.01 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00001.parquet`: `33169` rows, `4185593` file bytes (3.99 MiB), `25507934` physical bytes (24.33 MiB), `28682694` encoded bytes (27.35 MiB), `4153237` compressed data bytes (3.96 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00002.parquet`: `32700` rows, `4179037` file bytes (3.99 MiB), `25596334` physical bytes (24.41 MiB), `28731084` encoded bytes (27.40 MiB), `4146187` compressed data bytes (3.95 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00003.parquet`: `32955` rows, `4174150` file bytes (3.98 MiB), `25549834` physical bytes (24.37 MiB), `28702720` encoded bytes (27.37 MiB), `4141742` compressed data bytes (3.95 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00004.parquet`: `33168` rows, `4191585` file bytes (4.00 MiB), `25758139` physical bytes (24.56 MiB), `28936239` encoded bytes (27.60 MiB), `4158411` compressed data bytes (3.97 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00005.parquet`: `33398` rows, `4149139` file bytes (3.96 MiB), `25775747` physical bytes (24.58 MiB), `28967137` encoded bytes (27.63 MiB), `4116705` compressed data bytes (3.93 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00006.parquet`: `32965` rows, `4146423` file bytes (3.95 MiB), `25359682` physical bytes (24.18 MiB), `28517958` encoded bytes (27.20 MiB), `4114245` compressed data bytes (3.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00007.parquet`: `33213` rows, `4183932` file bytes (3.99 MiB), `25766003` physical bytes (24.57 MiB), `28945522` encoded bytes (27.60 MiB), `4151388` compressed data bytes (3.96 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00008.parquet`: `33050` rows, `4146675` file bytes (3.95 MiB), `25572436` physical bytes (24.39 MiB), `28733766` encoded bytes (27.40 MiB), `4113772` compressed data bytes (3.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00009.parquet`: `33164` rows, `4137383` file bytes (3.95 MiB), `25584530` physical bytes (24.40 MiB), `28758985` encoded bytes (27.43 MiB), `4104975` compressed data bytes (3.91 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00010.parquet`: `33058` rows, `4164128` file bytes (3.97 MiB), `25560477` physical bytes (24.38 MiB), `28721709` encoded bytes (27.39 MiB), `4132086` compressed data bytes (3.94 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00011.parquet`: `33070` rows, `4086508` file bytes (3.90 MiB), `25421040` physical bytes (24.24 MiB), `28584088` encoded bytes (27.26 MiB), `4054279` compressed data bytes (3.87 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00012.parquet`: `33230` rows, `4116172` file bytes (3.93 MiB), `25600485` physical bytes (24.41 MiB), `28780438` encoded bytes (27.45 MiB), `4083884` compressed data bytes (3.89 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00013.parquet`: `33279` rows, `4161693` file bytes (3.97 MiB), `25571528` physical bytes (24.39 MiB), `28753561` encoded bytes (27.42 MiB), `4129688` compressed data bytes (3.94 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00014.parquet`: `32749` rows, `4080768` file bytes (3.89 MiB), `25246115` physical bytes (24.08 MiB), `28377745` encoded bytes (27.06 MiB), `4048550` compressed data bytes (3.86 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00015.parquet`: `33232` rows, `4391245` file bytes (4.19 MiB), `23865122` physical bytes (22.76 MiB), `27040219` encoded bytes (25.79 MiB), `4360042` compressed data bytes (4.16 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00016.parquet`: `33349` rows, `4513504` file bytes (4.30 MiB), `22949355` physical bytes (21.89 MiB), `26130227` encoded bytes (24.92 MiB), `4483236` compressed data bytes (4.28 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00017.parquet`: `32473` rows, `4524287` file bytes (4.31 MiB), `22473276` physical bytes (21.43 MiB), `25568029` encoded bytes (24.38 MiB), `4494032` compressed data bytes (4.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00018.parquet`: `33155` rows, `4508619` file bytes (4.30 MiB), `21140297` physical bytes (20.16 MiB), `24288219` encoded bytes (23.16 MiB), `4478276` compressed data bytes (4.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00019.parquet`: `32871` rows, `4559448` file bytes (4.35 MiB), `21296723` physical bytes (20.31 MiB), `24414972` encoded bytes (23.28 MiB), `4528875` compressed data bytes (4.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00020.parquet`: `32858` rows, `4600153` file bytes (4.39 MiB), `21170469` physical bytes (20.19 MiB), `24287475` encoded bytes (23.16 MiB), `4569339` compressed data bytes (4.36 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00021.parquet`: `33246` rows, `4571557` file bytes (4.36 MiB), `21240144` physical bytes (20.26 MiB), `24395068` encoded bytes (23.26 MiB), `4540825` compressed data bytes (4.33 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00022.parquet`: `32928` rows, `4579505` file bytes (4.37 MiB), `21224492` physical bytes (20.24 MiB), `24349248` encoded bytes (23.22 MiB), `4548906` compressed data bytes (4.34 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00023.parquet`: `33035` rows, `4549034` file bytes (4.34 MiB), `21295226` physical bytes (20.31 MiB), `24430709` encoded bytes (23.30 MiB), `4518211` compressed data bytes (4.31 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00024.parquet`: `33180` rows, `4567690` file bytes (4.36 MiB), `21275192` physical bytes (20.29 MiB), `24422101` encoded bytes (23.29 MiB), `4537026` compressed data bytes (4.33 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00025.parquet`: `32879` rows, `4594583` file bytes (4.38 MiB), `21226277` physical bytes (20.24 MiB), `24345558` encoded bytes (23.22 MiB), `4564232` compressed data bytes (4.35 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00026.parquet`: `33452` rows, `4518689` file bytes (4.31 MiB), `21247079` physical bytes (20.26 MiB), `24423570` encoded bytes (23.29 MiB), `4488548` compressed data bytes (4.28 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00027.parquet`: `33075` rows, `4534257` file bytes (4.32 MiB), `21060943` physical bytes (20.09 MiB), `24202253` encoded bytes (23.08 MiB), `4503553` compressed data bytes (4.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00028.parquet`: `33143` rows, `4484162` file bytes (4.28 MiB), `21136681` physical bytes (20.16 MiB), `24283218` encoded bytes (23.16 MiB), `4453652` compressed data bytes (4.25 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00029.parquet`: `33289` rows, `4530701` file bytes (4.32 MiB), `21282561` physical bytes (20.30 MiB), `24444406` encoded bytes (23.31 MiB), `4500165` compressed data bytes (4.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00030.parquet`: `7548` rows, `1057128` file bytes (1.01 MiB), `4923658` physical bytes (4.70 MiB), `5641508` encoded bytes (5.38 MiB), `1041277` compressed data bytes (1016.87 KiB)
