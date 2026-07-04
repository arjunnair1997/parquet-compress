# ClickBench Parquet Experiment

- Started: `2026-07-03T23:29:22-04:00`
- Write elapsed: `11.153s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `811705594` (774.10 MiB)
- Compressed column data bytes after codec compression: `130653694` (124.60 MiB)
- Parquet file bytes: `131614372` (125.52 MiB)
- Physical/encoded ratio: `0.878x`
- Encoded/compressed-data ratio: `6.213x`
- Physical/compressed-data ratio: `5.453x`
- Physical/parquet-file ratio: `5.413x`
- Files: `31`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `plain`
- Date encoding: `plain`
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
- Elapsed: `6.979s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004876` (7.63 MiB) | `8005302` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `262092` (255.95 KiB) | `0.999x` | `15.276x` | `15.262x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:504` | `1000000` | `138409995` (132.00 MiB) | `142875515` (136.26 MiB) | `20896289` (19.93 MiB) | `0.969x` | `6.837x` | `6.624x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204393` (199.60 KiB) | `0.999x` | `19.589x` | `19.570x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `3369375` (3.21 MiB) | `3037056` (2.90 MiB) | `2.374x` | `1.109x` | `2.634x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `204392` (199.60 KiB) | `0.999x` | `19.589x` | `19.570x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204411` (199.62 KiB) | `0.999x` | `19.587x` | `19.568x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `719398` (702.54 KiB) | `0.999x` | `5.566x` | `5.560x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `396561` (387.27 KiB) | `0.999x` | `10.096x` | `10.087x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004869` (7.63 MiB) | `1084908` (1.03 MiB) | `0.999x` | `7.378x` | `7.374x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003833` (3.82 MiB) | `204213` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `318169` (310.71 KiB) | `0.999x` | `12.584x` | `12.572x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `342314` (334.29 KiB) | `0.999x` | `11.696x` | `11.685x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:343` | `1000000` | `88562192` (84.46 MiB) | `92649506` (88.36 MiB) | `20479893` (19.53 MiB) | `0.956x` | `4.524x` | `4.324x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:299` | `1000000` | `79583339` (75.90 MiB) | `83645314` (79.77 MiB) | `19052284` (18.17 MiB) | `0.951x` | `4.390x` | `4.177x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003870` (3.82 MiB) | `491986` (480.46 KiB) | `0.999x` | `8.138x` | `8.130x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003869` (3.82 MiB) | `508960` (497.03 KiB) | `0.999x` | `7.867x` | `7.859x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003853` (3.82 MiB) | `458142` (447.40 KiB) | `0.999x` | `8.739x` | `8.731x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `271964` (265.59 KiB) | `0.999x` | `14.722x` | `14.708x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `232275` (226.83 KiB) | `0.999x` | `17.237x` | `17.221x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `373595` (364.84 KiB) | `0.999x` | `10.717x` | `10.707x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `373249` (364.50 KiB) | `0.999x` | `10.727x` | `10.717x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `283028` (276.39 KiB) | `0.999x` | `14.146x` | `14.133x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `245406` (239.65 KiB) | `0.999x` | `16.315x` | `16.300x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `328472` (320.77 KiB) | `0.999x` | `12.189x` | `12.178x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `3354477` (3.20 MiB) | `7358092` (7.02 MiB) | `537028` (524.44 KiB) | `0.456x` | `13.702x` | `6.246x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `219187` (214.05 KiB) | `0.999x` | `18.267x` | `18.249x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `217364` (212.27 KiB) | `0.999x` | `18.420x` | `18.402x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `365184` (356.62 KiB) | `0.999x` | `10.964x` | `10.953x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `3767530` (3.59 MiB) | `7771139` (7.41 MiB) | `467322` (456.37 KiB) | `0.485x` | `16.629x` | `8.062x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `204872` (200.07 KiB) | `0.999x` | `19.543x` | `19.524x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204997` (200.19 KiB) | `0.999x` | `19.531x` | `19.512x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `218496` (213.38 KiB) | `0.999x` | `18.325x` | `18.307x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `215468` (210.42 KiB) | `0.999x` | `18.582x` | `18.564x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `81583` (79.67 KiB) | `4085144` (3.90 MiB) | `220101` (214.94 KiB) | `0.020x` | `18.560x` | `0.371x` | `81583` (79.67 KiB) |
| `Params` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `0` (0 B) | `4002369` (3.82 MiB) | `202749` (198.00 KiB) | `0.000x` | `19.741x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003876` (3.82 MiB) | `560873` (547.73 KiB) | `0.999x` | `7.139x` | `7.132x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003872` (3.82 MiB) | `526869` (514.52 KiB) | `0.999x` | `7.599x` | `7.592x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `293189` (286.32 KiB) | `0.999x` | `13.656x` | `13.643x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `3528017` (3.36 MiB) | `7535557` (7.19 MiB) | `1097549` (1.05 MiB) | `0.468x` | `6.866x` | `3.214x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `220361` (215.20 KiB) | `0.999x` | `18.169x` | `18.152x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003873` (3.82 MiB) | `505906` (494.05 KiB) | `0.999x` | `7.914x` | `7.907x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003865` (3.82 MiB) | `516799` (504.69 KiB) | `0.999x` | `7.747x` | `7.740x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003882` (3.82 MiB) | `551806` (538.87 KiB) | `0.999x` | `7.256x` | `7.249x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `285711` (279.01 KiB) | `0.999x` | `14.014x` | `14.000x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `3492175` (3.33 MiB) | `3062417` (2.92 MiB) | `2.291x` | `1.140x` | `2.612x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `287305` (280.57 KiB) | `0.999x` | `13.936x` | `13.922x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `269783` (263.46 KiB) | `0.999x` | `14.841x` | `14.827x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `304908` (297.76 KiB) | `0.999x` | `13.131x` | `13.119x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `204246` (199.46 KiB) | `0.999x` | `19.603x` | `19.584x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:91` | `1000000` | `13587860` (12.96 MiB) | `17595523` (16.78 MiB) | `916332` (894.86 KiB) | `0.772x` | `19.202x` | `14.829x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `205229` (200.42 KiB) | `0.999x` | `19.509x` | `19.490x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `256522` (250.51 KiB) | `0.999x` | `15.608x` | `15.593x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `206701` (201.86 KiB) | `0.999x` | `19.370x` | `19.352x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `235419` (229.90 KiB) | `0.999x` | `17.007x` | `16.991x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004871` (7.63 MiB) | `1151577` (1.10 MiB) | `0.999x` | `6.951x` | `6.947x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:142` | `1000000` | `27797671` (26.51 MiB) | `31859855` (30.38 MiB) | `7044778` (6.72 MiB) | `0.872x` | `4.522x` | `3.946x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `3688237` (3.52 MiB) | `0.999x` | `1.086x` | `1.085x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003833` (3.82 MiB) | `204213` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003833` (3.82 MiB) | `204213` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003833` (3.82 MiB) | `204213` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `295705` (288.77 KiB) | `0.999x` | `13.540x` | `13.527x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003833` (3.82 MiB) | `204213` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `1000000` (976.56 KiB) | `5003110` (4.77 MiB) | `294404` (287.50 KiB) | `0.200x` | `16.994x` | `3.397x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `3383959` (3.23 MiB) | `3045598` (2.90 MiB) | `2.364x` | `1.111x` | `2.627x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `334032` (326.20 KiB) | `0.999x` | `11.986x` | `11.975x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `302610` (295.52 KiB) | `0.999x` | `13.231x` | `13.218x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `318259` (310.80 KiB) | `0.999x` | `12.580x` | `12.568x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `385872` (376.83 KiB) | `0.999x` | `10.376x` | `10.366x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `381791` (372.84 KiB) | `0.999x` | `10.487x` | `10.477x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `707140` (690.57 KiB) | `0.999x` | `5.662x` | `5.657x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003844` (3.82 MiB) | `295401` (288.48 KiB) | `0.999x` | `13.554x` | `13.541x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `204217` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `258463` (252.41 KiB) | `0.999x` | `15.491x` | `15.476x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `2001192` (1.91 MiB) | `6004660` (5.73 MiB) | `321328` (313.80 KiB) | `0.333x` | `18.687x` | `6.228x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `3325142` (3.17 MiB) | `7328756` (6.99 MiB) | `457201` (446.49 KiB) | `0.454x` | `16.030x` | `7.273x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `0` (0 B) | `4002369` (3.82 MiB) | `202749` (198.00 KiB) | `0.000x` | `19.741x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `0` (0 B) | `4002369` (3.82 MiB) | `202749` (198.00 KiB) | `0.000x` | `19.741x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003833` (3.82 MiB) | `204213` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `269419` (263.10 KiB) | `0.999x` | `14.861x` | `14.847x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `327389` (319.72 KiB) | `0.999x` | `12.230x` | `12.218x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003877` (3.82 MiB) | `565081` (551.84 KiB) | `0.999x` | `7.085x` | `7.079x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `1725442` (1.65 MiB) | `0.999x` | `2.321x` | `2.318x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `1282944` (1.22 MiB) | `0.999x` | `3.121x` | `3.118x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003894` (3.82 MiB) | `810285` (791.29 KiB) | `0.999x` | `4.941x` | `4.937x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204606` (199.81 KiB) | `0.999x` | `19.569x` | `19.550x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `1024` (1.00 KiB) | `4004109` (3.82 MiB) | `204262` (199.47 KiB) | `0.000x` | `19.603x` | `0.005x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004816` (7.63 MiB) | `405189` (395.69 KiB) | `0.999x` | `19.756x` | `19.744x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `0` (0 B) | `4002369` (3.82 MiB) | `202749` (198.00 KiB) | `0.000x` | `19.741x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `3000000` (2.86 MiB) | `7003596` (6.68 MiB) | `354319` (346.01 KiB) | `0.428x` | `19.766x` | `8.467x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003833` (3.82 MiB) | `204213` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `58030` (56.67 KiB) | `4062844` (3.87 MiB) | `218536` (213.41 KiB) | `0.014x` | `18.591x` | `0.266x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `22051` (21.53 KiB) | `4025898` (3.84 MiB) | `214774` (209.74 KiB) | `0.005x` | `18.745x` | `0.103x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `25445` (24.85 KiB) | `4029295` (3.84 MiB) | `219135` (214.00 KiB) | `0.006x` | `18.387x` | `0.116x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `48191` (47.06 KiB) | `4052041` (3.86 MiB) | `214361` (209.34 KiB) | `0.012x` | `18.903x` | `0.225x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `49433` (48.27 KiB) | `4053494` (3.87 MiB) | `219126` (213.99 KiB) | `0.012x` | `18.498x` | `0.226x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `16873` (16.48 KiB) | `4020155` (3.83 MiB) | `214375` (209.35 KiB) | `0.004x` | `18.753x` | `0.079x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `91870` (89.72 KiB) | `4097846` (3.91 MiB) | `230210` (224.81 KiB) | `0.022x` | `17.800x` | `0.399x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `13001` (12.70 KiB) | `4016750` (3.83 MiB) | `212182` (207.21 KiB) | `0.003x` | `18.931x` | `0.061x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `28101` (27.44 KiB) | `4033132` (3.85 MiB) | `214687` (209.66 KiB) | `0.007x` | `18.786x` | `0.131x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `45607` (44.54 KiB) | `4048499` (3.86 MiB) | `221741` (216.54 KiB) | `0.011x` | `18.258x` | `0.206x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `213608` (208.60 KiB) | `0.999x` | `18.744x` | `18.726x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004877` (7.63 MiB) | `3640888` (3.47 MiB) | `0.999x` | `2.199x` | `2.197x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004869` (7.63 MiB) | `4386878` (4.18 MiB) | `0.999x` | `1.825x` | `1.824x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `204674` (199.88 KiB) | `0.999x` | `19.562x` | `19.543x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00000.parquet`: `33116` rows, `4246989` file bytes (4.05 MiB), `25718243` physical bytes (24.53 MiB), `29025423` encoded bytes (27.68 MiB), `4215013` compressed data bytes (4.02 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00001.parquet`: `33162` rows, `4192432` file bytes (4.00 MiB), `25502086` physical bytes (24.32 MiB), `28808063` encoded bytes (27.47 MiB), `4160074` compressed data bytes (3.97 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00002.parquet`: `32696` rows, `4184219` file bytes (3.99 MiB), `25595426` physical bytes (24.41 MiB), `28858301` encoded bytes (27.52 MiB), `4151367` compressed data bytes (3.96 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00003.parquet`: `32955` rows, `4180990` file bytes (3.99 MiB), `25547098` physical bytes (24.36 MiB), `28830871` encoded bytes (27.50 MiB), `4148580` compressed data bytes (3.96 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00004.parquet`: `33165` rows, `4196296` file bytes (4.00 MiB), `25754251` physical bytes (24.56 MiB), `29063201` encoded bytes (27.72 MiB), `4163120` compressed data bytes (3.97 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00005.parquet`: `33394` rows, `4155112` file bytes (3.96 MiB), `25773995` physical bytes (24.58 MiB), `29098804` encoded bytes (27.75 MiB), `4122676` compressed data bytes (3.93 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00006.parquet`: `32962` rows, `4152049` file bytes (3.96 MiB), `25358704` physical bytes (24.18 MiB), `28644975` encoded bytes (27.32 MiB), `4119869` compressed data bytes (3.93 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00007.parquet`: `33212` rows, `4187535` file bytes (3.99 MiB), `25764907` physical bytes (24.57 MiB), `29075151` encoded bytes (27.73 MiB), `4154989` compressed data bytes (3.96 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00008.parquet`: `33044` rows, `4152310` file bytes (3.96 MiB), `25569786` physical bytes (24.39 MiB), `28860825` encoded bytes (27.52 MiB), `4119405` compressed data bytes (3.93 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00009.parquet`: `33163` rows, `4141367` file bytes (3.95 MiB), `25583966` physical bytes (24.40 MiB), `28888586` encoded bytes (27.55 MiB), `4108957` compressed data bytes (3.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00010.parquet`: `33062` rows, `4168750` file bytes (3.98 MiB), `25559236` physical bytes (24.38 MiB), `28850747` encoded bytes (27.51 MiB), `4136706` compressed data bytes (3.95 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00011.parquet`: `33063` rows, `4091411` file bytes (3.90 MiB), `25419224` physical bytes (24.24 MiB), `28711332` encoded bytes (27.38 MiB), `4059180` compressed data bytes (3.87 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00012.parquet`: `33232` rows, `4117130` file bytes (3.93 MiB), `25602723` physical bytes (24.42 MiB), `28912669` encoded bytes (27.57 MiB), `4084840` compressed data bytes (3.90 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00013.parquet`: `33276` rows, `4165595` file bytes (3.97 MiB), `25568442` physical bytes (24.38 MiB), `28882367` encoded bytes (27.54 MiB), `4133586` compressed data bytes (3.94 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00014.parquet`: `32748` rows, `4087014` file bytes (3.90 MiB), `25240257` physical bytes (24.07 MiB), `28500965` encoded bytes (27.18 MiB), `4054792` compressed data bytes (3.87 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00015.parquet`: `33439` rows, `4429621` file bytes (4.22 MiB), `24028010` physical bytes (22.91 MiB), `27355075` encoded bytes (26.09 MiB), `4398381` compressed data bytes (4.19 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00016.parquet`: `33364` rows, `4521658` file bytes (4.31 MiB), `22945447` physical bytes (21.88 MiB), `26257902` encoded bytes (25.04 MiB), `4491360` compressed data bytes (4.28 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00017.parquet`: `32490` rows, `4523796` file bytes (4.31 MiB), `22466040` physical bytes (21.43 MiB), `25690467` encoded bytes (24.50 MiB), `4493539` compressed data bytes (4.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00018.parquet`: `33137` rows, `4520610` file bytes (4.31 MiB), `21130510` physical bytes (20.15 MiB), `24408678` encoded bytes (23.28 MiB), `4490264` compressed data bytes (4.28 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00019.parquet`: `32874` rows, `4564110` file bytes (4.35 MiB), `21299260` physical bytes (20.31 MiB), `24547853` encoded bytes (23.41 MiB), `4533535` compressed data bytes (4.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00020.parquet`: `32855` rows, `4605724` file bytes (4.39 MiB), `21164183` physical bytes (20.18 MiB), `24413243` encoded bytes (23.28 MiB), `4574909` compressed data bytes (4.36 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00021.parquet`: `33212` rows, `4580630` file bytes (4.37 MiB), `21233772` physical bytes (20.25 MiB), `24516126` encoded bytes (23.38 MiB), `4549895` compressed data bytes (4.34 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00022.parquet`: `32916` rows, `4585916` file bytes (4.37 MiB), `21223511` physical bytes (20.24 MiB), `24477202` encoded bytes (23.34 MiB), `4555315` compressed data bytes (4.34 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00023.parquet`: `33055` rows, `4557240` file bytes (4.35 MiB), `21293075` physical bytes (20.31 MiB), `24560485` encoded bytes (23.42 MiB), `4526415` compressed data bytes (4.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00024.parquet`: `33164` rows, `4576564` file bytes (4.36 MiB), `21273828` physical bytes (20.29 MiB), `24550771` encoded bytes (23.41 MiB), `4545899` compressed data bytes (4.34 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00025.parquet`: `32877` rows, `4600014` file bytes (4.39 MiB), `21224420` physical bytes (20.24 MiB), `24473051` encoded bytes (23.34 MiB), `4569666` compressed data bytes (4.36 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00026.parquet`: `33446` rows, `4529477` file bytes (4.32 MiB), `21246248` physical bytes (20.26 MiB), `24554119` encoded bytes (23.42 MiB), `4499343` compressed data bytes (4.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00027.parquet`: `32823` rows, `4501567` file bytes (4.29 MiB), `20899762` physical bytes (19.93 MiB), `24146351` encoded bytes (23.03 MiB), `4470895` compressed data bytes (4.26 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00028.parquet`: `33156` rows, `4486440` file bytes (4.28 MiB), `21130864` physical bytes (20.15 MiB), `24409660` encoded bytes (23.28 MiB), `4455928` compressed data bytes (4.25 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00029.parquet`: `33264` rows, `4538086` file bytes (4.33 MiB), `21275166` physical bytes (20.29 MiB), `24565717` encoded bytes (23.43 MiB), `4507547` compressed data bytes (4.30 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00030.parquet`: `7678` rows, `1073720` file bytes (1.02 MiB), `5006184` physical bytes (4.77 MiB), `5766614` encoded bytes (5.50 MiB), `1057649` compressed data bytes (1.01 MiB)
