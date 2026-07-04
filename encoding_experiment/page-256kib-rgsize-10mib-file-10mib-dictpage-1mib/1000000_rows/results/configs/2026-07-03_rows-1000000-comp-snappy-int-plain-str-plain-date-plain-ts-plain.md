# ClickBench Parquet Experiment

- Started: `2026-07-03T15:28:39-04:00`
- Write elapsed: `11.358s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `825475951` (787.24 MiB)
- Compressed column data bytes after codec compression: `134316393` (128.09 MiB)
- Parquet file bytes: `135277538` (129.01 MiB)
- Physical/encoded ratio: `0.863x`
- Encoded/compressed-data ratio: `6.146x`
- Physical/compressed-data ratio: `5.304x`
- Physical/parquet-file ratio: `5.266x`
- Files: `31`

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
- Rows read and compared: `1000000`
- Files read: `31`
- Elapsed: `6.93s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004879` (7.63 MiB) | `8005306` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `262070` (255.93 KiB) | `0.999x` | `15.278x` | `15.263x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:505` | `1000000` | `138409995` (132.00 MiB) | `142874702` (136.26 MiB) | `20906222` (19.94 MiB) | `0.969x` | `6.834x` | `6.621x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `204379` (199.59 KiB) | `0.999x` | `19.590x` | `19.571x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004875` (7.63 MiB) | `4282586` (4.08 MiB) | `0.999x` | `1.869x` | `1.868x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204376` (199.59 KiB) | `0.999x` | `19.591x` | `19.572x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204399` (199.61 KiB) | `0.999x` | `19.588x` | `19.570x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `719481` (702.62 KiB) | `0.999x` | `5.565x` | `5.560x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `396778` (387.48 KiB) | `0.999x` | `10.091x` | `10.081x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004874` (7.63 MiB) | `1084918` (1.03 MiB) | `0.999x` | `7.378x` | `7.374x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `204215` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `318110` (310.65 KiB) | `0.999x` | `12.586x` | `12.574x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `342202` (334.18 KiB) | `0.999x` | `11.700x` | `11.689x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:343` | `1000000` | `88562192` (84.46 MiB) | `92652558` (88.36 MiB) | `20480534` (19.53 MiB) | `0.956x` | `4.524x` | `4.324x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:299` | `1000000` | `79583339` (75.90 MiB) | `83645311` (79.77 MiB) | `19049622` (18.17 MiB) | `0.951x` | `4.391x` | `4.178x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003873` (3.82 MiB) | `492036` (480.50 KiB) | `0.999x` | `8.137x` | `8.129x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003868` (3.82 MiB) | `509008` (497.08 KiB) | `0.999x` | `7.866x` | `7.858x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003856` (3.82 MiB) | `458123` (447.39 KiB) | `0.999x` | `8.740x` | `8.731x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `272037` (265.66 KiB) | `0.999x` | `14.718x` | `14.704x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `232236` (226.79 KiB) | `0.999x` | `17.240x` | `17.224x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `373488` (364.73 KiB) | `0.999x` | `10.720x` | `10.710x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `373249` (364.50 KiB) | `0.999x` | `10.727x` | `10.717x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `283052` (276.42 KiB) | `0.999x` | `14.145x` | `14.132x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `245331` (239.58 KiB) | `0.999x` | `16.320x` | `16.305x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `328455` (320.76 KiB) | `0.999x` | `12.190x` | `12.178x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `3354477` (3.20 MiB) | `7358098` (7.02 MiB) | `536824` (524.24 KiB) | `0.456x` | `13.707x` | `6.249x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `219217` (214.08 KiB) | `0.999x` | `18.264x` | `18.247x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `217372` (212.28 KiB) | `0.999x` | `18.419x` | `18.402x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003844` (3.82 MiB) | `365174` (356.62 KiB) | `0.999x` | `10.964x` | `10.954x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `3767530` (3.59 MiB) | `7771150` (7.41 MiB) | `467164` (456.21 KiB) | `0.485x` | `16.635x` | `8.065x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204886` (200.08 KiB) | `0.999x` | `19.542x` | `19.523x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `204972` (200.17 KiB) | `0.999x` | `19.534x` | `19.515x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `218511` (213.39 KiB) | `0.999x` | `18.323x` | `18.306x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `215432` (210.38 KiB) | `0.999x` | `18.585x` | `18.567x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `81583` (79.67 KiB) | `4085137` (3.90 MiB) | `220038` (214.88 KiB) | `0.020x` | `18.566x` | `0.371x` | `81583` (79.67 KiB) |
| `Params` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `0` (0 B) | `4002379` (3.82 MiB) | `202751` (198.00 KiB) | `0.000x` | `19.740x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003873` (3.82 MiB) | `561124` (547.97 KiB) | `0.999x` | `7.135x` | `7.129x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003874` (3.82 MiB) | `526869` (514.52 KiB) | `0.999x` | `7.599x` | `7.592x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `293177` (286.31 KiB) | `0.999x` | `13.657x` | `13.644x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `3528017` (3.36 MiB) | `7535913` (7.19 MiB) | `1098051` (1.05 MiB) | `0.468x` | `6.863x` | `3.213x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `220366` (215.20 KiB) | `0.999x` | `18.169x` | `18.152x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003876` (3.82 MiB) | `505868` (494.01 KiB) | `0.999x` | `7.915x` | `7.907x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003870` (3.82 MiB) | `516714` (504.60 KiB) | `0.999x` | `7.749x` | `7.741x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003882` (3.82 MiB) | `551892` (538.96 KiB) | `0.999x` | `7.255x` | `7.248x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `285815` (279.12 KiB) | `0.999x` | `14.009x` | `13.995x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004879` (7.63 MiB) | `4229160` (4.03 MiB) | `0.999x` | `1.893x` | `1.892x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `287294` (280.56 KiB) | `0.999x` | `13.936x` | `13.923x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `269816` (263.49 KiB) | `0.999x` | `14.839x` | `14.825x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `304934` (297.79 KiB) | `0.999x` | `13.130x` | `13.118x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `204253` (199.47 KiB) | `0.999x` | `19.602x` | `19.584x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:91` | `1000000` | `13587860` (12.96 MiB) | `17595482` (16.78 MiB) | `924713` (903.04 KiB) | `0.772x` | `19.028x` | `14.694x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `205215` (200.41 KiB) | `0.999x` | `19.510x` | `19.492x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `256552` (250.54 KiB) | `0.999x` | `15.606x` | `15.591x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `206697` (201.85 KiB) | `0.999x` | `19.371x` | `19.352x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003845` (3.82 MiB) | `235444` (229.93 KiB) | `0.999x` | `17.006x` | `16.989x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004875` (7.63 MiB) | `1151604` (1.10 MiB) | `0.999x` | `6.951x` | `6.947x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:140` | `1000000` | `27797671` (26.51 MiB) | `31858218` (30.38 MiB) | `7037379` (6.71 MiB) | `0.873x` | `4.527x` | `3.950x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003897` (3.82 MiB) | `3688228` (3.52 MiB) | `0.999x` | `1.086x` | `1.085x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `204215` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `204215` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `204215` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `295708` (288.78 KiB) | `0.999x` | `13.540x` | `13.527x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `204215` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `1000000` (976.56 KiB) | `5003112` (4.77 MiB) | `294413` (287.51 KiB) | `0.200x` | `16.994x` | `3.397x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004879` (7.63 MiB) | `4284958` (4.09 MiB) | `0.999x` | `1.868x` | `1.867x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `333988` (326.16 KiB) | `0.999x` | `11.988x` | `11.976x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `302670` (295.58 KiB) | `0.999x` | `13.228x` | `13.216x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `318283` (310.82 KiB) | `0.999x` | `12.579x` | `12.567x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `385674` (376.63 KiB) | `0.999x` | `10.381x` | `10.371x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003844` (3.82 MiB) | `381887` (372.94 KiB) | `0.999x` | `10.484x` | `10.474x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `707621` (691.04 KiB) | `0.999x` | `5.658x` | `5.653x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `295315` (288.39 KiB) | `0.999x` | `13.558x` | `13.545x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204210` (199.42 KiB) | `0.999x` | `19.606x` | `19.588x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `258329` (252.27 KiB) | `0.999x` | `15.499x` | `15.484x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `2001192` (1.91 MiB) | `6004663` (5.73 MiB) | `321313` (313.78 KiB) | `0.333x` | `18.688x` | `6.228x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `3325142` (3.17 MiB) | `7328758` (6.99 MiB) | `457183` (446.47 KiB) | `0.454x` | `16.030x` | `7.273x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `0` (0 B) | `4002379` (3.82 MiB) | `202751` (198.00 KiB) | `0.000x` | `19.740x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `0` (0 B) | `4002379` (3.82 MiB) | `202751` (198.00 KiB) | `0.000x` | `19.740x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `204215` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003848` (3.82 MiB) | `269326` (263.01 KiB) | `0.999x` | `14.866x` | `14.852x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `327222` (319.55 KiB) | `0.999x` | `12.236x` | `12.224x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003876` (3.82 MiB) | `564916` (551.68 KiB) | `0.999x` | `7.088x` | `7.081x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `1726790` (1.65 MiB) | `0.999x` | `2.319x` | `2.316x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `1282976` (1.22 MiB) | `0.999x` | `3.121x` | `3.118x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003897` (3.82 MiB) | `809765` (790.79 KiB) | `0.999x` | `4.945x` | `4.940x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204595` (199.80 KiB) | `0.999x` | `19.570x` | `19.551x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `1024` (1.00 KiB) | `4004128` (3.82 MiB) | `204280` (199.49 KiB) | `0.000x` | `19.601x` | `0.005x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004814` (7.63 MiB) | `405190` (395.69 KiB) | `0.999x` | `19.756x` | `19.744x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `0` (0 B) | `4002379` (3.82 MiB) | `202751` (198.00 KiB) | `0.000x` | `19.740x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `3000000` (2.86 MiB) | `7003596` (6.68 MiB) | `354326` (346.02 KiB) | `0.428x` | `19.766x` | `8.467x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `204215` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `58030` (56.67 KiB) | `4062845` (3.87 MiB) | `218529` (213.41 KiB) | `0.014x` | `18.592x` | `0.266x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `22051` (21.53 KiB) | `4025886` (3.84 MiB) | `214750` (209.72 KiB) | `0.005x` | `18.747x` | `0.103x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `25445` (24.85 KiB) | `4029295` (3.84 MiB) | `219153` (214.02 KiB) | `0.006x` | `18.386x` | `0.116x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `48191` (47.06 KiB) | `4052055` (3.86 MiB) | `214370` (209.35 KiB) | `0.012x` | `18.902x` | `0.225x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `49433` (48.27 KiB) | `4053511` (3.87 MiB) | `219144` (214.01 KiB) | `0.012x` | `18.497x` | `0.226x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `16873` (16.48 KiB) | `4020163` (3.83 MiB) | `214394` (209.37 KiB) | `0.004x` | `18.751x` | `0.079x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `91870` (89.72 KiB) | `4097905` (3.91 MiB) | `230416` (225.02 KiB) | `0.022x` | `17.785x` | `0.399x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `13001` (12.70 KiB) | `4016831` (3.83 MiB) | `212235` (207.26 KiB) | `0.003x` | `18.926x` | `0.061x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `28101` (27.44 KiB) | `4032973` (3.85 MiB) | `214577` (209.55 KiB) | `0.007x` | `18.795x` | `0.131x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `45607` (44.54 KiB) | `4048505` (3.86 MiB) | `221659` (216.46 KiB) | `0.011x` | `18.265x` | `0.206x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `213611` (208.60 KiB) | `0.999x` | `18.744x` | `18.726x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004877` (7.63 MiB) | `3641169` (3.47 MiB) | `0.999x` | `2.198x` | `2.197x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004879` (7.63 MiB) | `4387520` (4.18 MiB) | `0.999x` | `1.824x` | `1.823x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `204666` (199.87 KiB) | `0.999x` | `19.563x` | `19.544x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00000.parquet`: `33077` rows, `4308208` file bytes (4.11 MiB), `25685097` physical bytes (24.50 MiB), `29434067` encoded bytes (28.07 MiB), `4276221` compressed data bytes (4.08 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00001.parquet`: `33115` rows, `4258978` file bytes (4.06 MiB), `25469156` physical bytes (24.29 MiB), `29229509` encoded bytes (27.88 MiB), `4226609` compressed data bytes (4.03 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00002.parquet`: `32648` rows, `4248295` file bytes (4.05 MiB), `25564850` physical bytes (24.38 MiB), `29274733` encoded bytes (27.92 MiB), `4215432` compressed data bytes (4.02 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00003.parquet`: `32902` rows, `4246065` file bytes (4.05 MiB), `25508845` physical bytes (24.33 MiB), `29243535` encoded bytes (27.89 MiB), `4213642` compressed data bytes (4.02 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00004.parquet`: `33129` rows, `4265819` file bytes (4.07 MiB), `25718062` physical bytes (24.53 MiB), `29480248` encoded bytes (28.11 MiB), `4232616` compressed data bytes (4.04 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00005.parquet`: `33357` rows, `4224755` file bytes (4.03 MiB), `25736140` physical bytes (24.54 MiB), `29521479` encoded bytes (28.15 MiB), `4192292` compressed data bytes (4.00 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00006.parquet`: `32904` rows, `4218657` file bytes (4.02 MiB), `25323061` physical bytes (24.15 MiB), `29056660` encoded bytes (27.71 MiB), `4186466` compressed data bytes (3.99 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00007.parquet`: `33157` rows, `4253255` file bytes (4.06 MiB), `25729177` physical bytes (24.54 MiB), `29495810` encoded bytes (28.13 MiB), `4220733` compressed data bytes (4.03 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00008.parquet`: `32977` rows, `4218378` file bytes (4.02 MiB), `25530433` physical bytes (24.35 MiB), `29270264` encoded bytes (27.91 MiB), `4185478` compressed data bytes (3.99 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00009.parquet`: `33113` rows, `4206712` file bytes (4.01 MiB), `25541272` physical bytes (24.36 MiB), `29297535` encoded bytes (27.94 MiB), `4174388` compressed data bytes (3.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00010.parquet`: `33045` rows, `4235028` file bytes (4.04 MiB), `25515470` physical bytes (24.33 MiB), `29262139` encoded bytes (27.91 MiB), `4202971` compressed data bytes (4.01 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00011.parquet`: `33006` rows, `4164459` file bytes (3.97 MiB), `25373153` physical bytes (24.20 MiB), `29118902` encoded bytes (27.77 MiB), `4132211` compressed data bytes (3.94 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00012.parquet`: `33174` rows, `4186369` file bytes (3.99 MiB), `25567713` physical bytes (24.38 MiB), `29333906` encoded bytes (27.97 MiB), `4154054` compressed data bytes (3.96 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00013.parquet`: `33200` rows, `4233677` file bytes (4.04 MiB), `25533356` physical bytes (24.35 MiB), `29298650` encoded bytes (27.94 MiB), `4201637` compressed data bytes (4.01 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00014.parquet`: `32780` rows, `4155307` file bytes (3.96 MiB), `25197272` physical bytes (24.03 MiB), `28912778` encoded bytes (27.57 MiB), `4123042` compressed data bytes (3.93 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00015.parquet`: `33304` rows, `4522873` file bytes (4.31 MiB), `24025511` physical bytes (22.91 MiB), `27791494` encoded bytes (26.50 MiB), `4491613` compressed data bytes (4.28 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00016.parquet`: `33251` rows, `4662214` file bytes (4.45 MiB), `22875596` physical bytes (21.82 MiB), `26626149` encoded bytes (25.39 MiB), `4631917` compressed data bytes (4.42 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00017.parquet`: `32291` rows, `4678039` file bytes (4.46 MiB), `22411302` physical bytes (21.37 MiB), `26053111` encoded bytes (24.85 MiB), `4647765` compressed data bytes (4.43 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00018.parquet`: `33151` rows, `4650750` file bytes (4.44 MiB), `21049185` physical bytes (20.07 MiB), `24781918` encoded bytes (23.63 MiB), `4620348` compressed data bytes (4.41 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00019.parquet`: `32633` rows, `4733090` file bytes (4.51 MiB), `21200374` physical bytes (20.22 MiB), `24876184` encoded bytes (23.72 MiB), `4702537` compressed data bytes (4.48 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00020.parquet`: `32184` rows, `4686865` file bytes (4.47 MiB), `20746091` physical bytes (19.79 MiB), `24371384` encoded bytes (23.24 MiB), `4656463` compressed data bytes (4.44 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00021.parquet`: `33056` rows, `4736189` file bytes (4.52 MiB), `21149614` physical bytes (20.17 MiB), `24874042` encoded bytes (23.72 MiB), `4705212` compressed data bytes (4.49 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00022.parquet`: `32763` rows, `4748175` file bytes (4.53 MiB), `21125942` physical bytes (20.15 MiB), `24816087` encoded bytes (23.67 MiB), `4717661` compressed data bytes (4.50 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00023.parquet`: `32960` rows, `4703659` file bytes (4.49 MiB), `21218106` physical bytes (20.24 MiB), `24931128` encoded bytes (23.78 MiB), `4672747` compressed data bytes (4.46 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00024.parquet`: `32788` rows, `4693020` file bytes (4.48 MiB), `21038142` physical bytes (20.06 MiB), `24730371` encoded bytes (23.58 MiB), `4662809` compressed data bytes (4.45 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00025.parquet`: `32791` rows, `4751169` file bytes (4.53 MiB), `21137224` physical bytes (20.16 MiB), `24830842` encoded bytes (23.68 MiB), `4720363` compressed data bytes (4.50 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00026.parquet`: `32995` rows, `4653591` file bytes (4.44 MiB), `20973194` physical bytes (20.00 MiB), `24688004` encoded bytes (23.54 MiB), `4623483` compressed data bytes (4.41 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00027.parquet`: `32922` rows, `4685074` file bytes (4.47 MiB), `21003996` physical bytes (20.03 MiB), `24713105` encoded bytes (23.57 MiB), `4654228` compressed data bytes (4.44 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00028.parquet`: `32920` rows, `4653549` file bytes (4.44 MiB), `21022446` physical bytes (20.05 MiB), `24730334` encoded bytes (23.58 MiB), `4622990` compressed data bytes (4.41 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00029.parquet`: `32881` rows, `4640959` file bytes (4.43 MiB), `20954457` physical bytes (19.98 MiB), `24657342` encoded bytes (23.52 MiB), `4610642` compressed data bytes (4.40 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-plain/part-00030.parquet`: `11526` rows, `1654360` file bytes (1.58 MiB), `7474387` physical bytes (7.13 MiB), `8774241` encoded bytes (8.37 MiB), `1637823` compressed data bytes (1.56 MiB)
