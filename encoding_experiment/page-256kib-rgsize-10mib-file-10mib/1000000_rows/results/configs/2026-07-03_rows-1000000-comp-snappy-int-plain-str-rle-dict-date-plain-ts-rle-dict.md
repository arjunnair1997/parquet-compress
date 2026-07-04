# ClickBench Parquet Experiment

- Started: `2026-07-03T14:58:10-04:00`
- Write elapsed: `11.655s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `473401958` (451.47 MiB)
- Compressed column data bytes after codec compression: `111242239` (106.09 MiB)
- Parquet file bytes: `112176317` (106.98 MiB)
- Physical/encoded ratio: `1.505x`
- Encoded/compressed-data ratio: `4.256x`
- Physical/compressed-data ratio: `6.404x`
- Physical/parquet-file ratio: `6.351x`
- Files: `31`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `rle-dict`
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
- Rows read and compared: `1000000`
- Files read: `31`
- Elapsed: `6.884s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004956` (7.63 MiB) | `8005389` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `262195` (256.05 KiB) | `0.999x` | `15.271x` | `15.256x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `138409995` (132.00 MiB) | `28853429` (27.52 MiB) | `10016710` (9.55 MiB) | `4.797x` | `2.881x` | `13.818x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `204459` (199.67 KiB) | `0.999x` | `19.583x` | `19.564x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `7344855` (7.00 MiB) | `5516743` (5.26 MiB) | `1.089x` | `1.331x` | `1.450x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `204460` (199.67 KiB) | `0.999x` | `19.583x` | `19.564x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204476` (199.68 KiB) | `0.999x` | `19.581x` | `19.562x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003961` (3.82 MiB) | `719536` (702.67 KiB) | `0.999x` | `5.565x` | `5.559x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003904` (3.82 MiB) | `396935` (387.63 KiB) | `0.999x` | `10.087x` | `10.077x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004950` (7.63 MiB) | `1085032` (1.03 MiB) | `0.999x` | `7.378x` | `7.373x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204276` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `318195` (310.74 KiB) | `0.999x` | `12.583x` | `12.571x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `342430` (334.40 KiB) | `0.999x` | `11.693x` | `11.681x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `88562192` (84.46 MiB) | `44097177` (42.05 MiB) | `16099145` (15.35 MiB) | `2.008x` | `2.739x` | `5.501x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `79583339` (75.90 MiB) | `34410100` (32.82 MiB) | `14809433` (14.12 MiB) | `2.313x` | `2.324x` | `5.374x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003927` (3.82 MiB) | `492045` (480.51 KiB) | `0.999x` | `8.137x` | `8.129x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003933` (3.82 MiB) | `509115` (497.18 KiB) | `0.999x` | `7.864x` | `7.857x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003920` (3.82 MiB) | `458270` (447.53 KiB) | `0.999x` | `8.737x` | `8.728x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `272053` (265.68 KiB) | `0.999x` | `14.717x` | `14.703x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `232330` (226.88 KiB) | `0.999x` | `17.234x` | `17.217x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `373613` (364.86 KiB) | `0.999x` | `10.717x` | `10.706x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `373411` (364.66 KiB) | `0.999x` | `10.723x` | `10.712x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `283101` (276.47 KiB) | `0.999x` | `14.143x` | `14.129x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `245379` (239.63 KiB) | `0.999x` | `16.317x` | `16.301x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `328589` (320.89 KiB) | `0.999x` | `12.185x` | `12.173x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `3354477` (3.20 MiB) | `267617` (261.34 KiB) | `183232` (178.94 KiB) | `12.535x` | `1.461x` | `18.307x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `219252` (214.11 KiB) | `0.999x` | `18.262x` | `18.244x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `217398` (212.30 KiB) | `0.999x` | `18.417x` | `18.399x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `365425` (356.86 KiB) | `0.999x` | `10.957x` | `10.946x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `3767530` (3.59 MiB) | `151415` (147.87 KiB) | `104702` (102.25 KiB) | `24.882x` | `1.446x` | `35.983x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204943` (200.14 KiB) | `0.999x` | `19.537x` | `19.518x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `205095` (200.29 KiB) | `0.999x` | `19.522x` | `19.503x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `218556` (213.43 KiB) | `0.999x` | `18.320x` | `18.302x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `215491` (210.44 KiB) | `0.999x` | `18.580x` | `18.562x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `81583` (79.67 KiB) | `25449` (24.85 KiB) | `23259` (22.71 KiB) | `3.206x` | `1.094x` | `3.508x` | `81583` (79.67 KiB) |
| `Params` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `0` (0 B) | `4394` (4.29 KiB) | `4642` (4.53 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003938` (3.82 MiB) | `561203` (548.05 KiB) | `0.999x` | `7.135x` | `7.128x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003934` (3.82 MiB) | `526873` (514.52 KiB) | `0.999x` | `7.599x` | `7.592x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `293280` (286.41 KiB) | `0.999x` | `13.652x` | `13.639x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `3528017` (3.36 MiB) | `1643075` (1.57 MiB) | `786209` (767.78 KiB) | `2.147x` | `2.090x` | `4.487x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `220409` (215.24 KiB) | `0.999x` | `18.166x` | `18.148x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003936` (3.82 MiB) | `505907` (494.05 KiB) | `0.999x` | `7.914x` | `7.907x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003933` (3.82 MiB) | `517031` (504.91 KiB) | `0.999x` | `7.744x` | `7.736x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003940` (3.82 MiB) | `552097` (539.16 KiB) | `0.999x` | `7.252x` | `7.245x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `285889` (279.19 KiB) | `0.999x` | `14.005x` | `13.991x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `7220966` (6.89 MiB) | `5443327` (5.19 MiB) | `1.108x` | `1.327x` | `1.470x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `287400` (280.66 KiB) | `0.999x` | `13.931x` | `13.918x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `269864` (263.54 KiB) | `0.999x` | `14.837x` | `14.822x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `305051` (297.90 KiB) | `0.999x` | `13.125x` | `13.113x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204310` (199.52 KiB) | `0.999x` | `19.597x` | `19.578x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `13587860` (12.96 MiB) | `9178` (8.96 KiB) | `9362` (9.14 KiB) | `1480.482x` | `0.980x` | `1451.384x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `205313` (200.50 KiB) | `0.999x` | `19.501x` | `19.482x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `256611` (250.60 KiB) | `0.999x` | `15.603x` | `15.588x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `206752` (201.91 KiB) | `0.999x` | `19.366x` | `19.347x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003904` (3.82 MiB) | `235481` (229.96 KiB) | `0.999x` | `17.003x` | `16.987x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004950` (7.63 MiB) | `1151747` (1.10 MiB) | `0.999x` | `6.950x` | `6.946x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `27797671` (26.51 MiB) | `21275560` (20.29 MiB) | `6166518` (5.88 MiB) | `1.307x` | `3.450x` | `4.508x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003958` (3.82 MiB) | `3688289` (3.52 MiB) | `0.999x` | `1.086x` | `1.085x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204276` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204276` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204276` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003908` (3.82 MiB) | `295776` (288.84 KiB) | `0.999x` | `13.537x` | `13.524x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204276` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `1000000` (976.56 KiB) | `33925` (33.13 KiB) | `32173` (31.42 KiB) | `29.477x` | `1.054x` | `31.082x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `7339797` (7.00 MiB) | `5514694` (5.26 MiB) | `1.090x` | `1.331x` | `1.451x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `334086` (326.26 KiB) | `0.999x` | `11.985x` | `11.973x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `302675` (295.58 KiB) | `0.999x` | `13.228x` | `13.215x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `318422` (310.96 KiB) | `0.999x` | `12.574x` | `12.562x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `385782` (376.74 KiB) | `0.999x` | `10.379x` | `10.369x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `382046` (373.09 KiB) | `0.999x` | `10.480x` | `10.470x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003961` (3.82 MiB) | `707578` (690.99 KiB) | `0.999x` | `5.659x` | `5.653x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003907` (3.82 MiB) | `295424` (288.50 KiB) | `0.999x` | `13.553x` | `13.540x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204276` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `258766` (252.70 KiB) | `0.999x` | `15.473x` | `15.458x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `2001192` (1.91 MiB) | `42807` (41.80 KiB) | `34726` (33.91 KiB) | `46.749x` | `1.233x` | `57.628x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `3325142` (3.17 MiB) | `138141` (134.90 KiB) | `95183` (92.95 KiB) | `24.071x` | `1.451x` | `34.934x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `0` (0 B) | `4394` (4.29 KiB) | `4642` (4.53 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `0` (0 B) | `4394` (4.29 KiB) | `4642` (4.53 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204276` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003908` (3.82 MiB) | `269242` (262.93 KiB) | `0.999x` | `14.871x` | `14.857x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `327581` (319.90 KiB) | `0.999x` | `12.223x` | `12.211x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003940` (3.82 MiB) | `565856` (552.59 KiB) | `0.999x` | `7.076x` | `7.069x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003964` (3.82 MiB) | `1726536` (1.65 MiB) | `0.999x` | `2.319x` | `2.317x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003961` (3.82 MiB) | `1284449` (1.22 MiB) | `0.999x` | `3.117x` | `3.114x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003955` (3.82 MiB) | `810747` (791.75 KiB) | `0.999x` | `4.939x` | `4.934x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `204770` (199.97 KiB) | `0.999x` | `19.553x` | `19.534x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `1024` (1.00 KiB) | `6010` (5.87 KiB) | `6260` (6.11 KiB) | `0.170x` | `0.960x` | `0.164x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004894` (7.63 MiB) | `405263` (395.76 KiB) | `0.999x` | `19.752x` | `19.740x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `0` (0 B) | `4394` (4.29 KiB) | `4642` (4.53 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `3000000` (2.86 MiB) | `5324` (5.20 KiB) | `5572` (5.44 KiB) | `563.486x` | `0.955x` | `538.406x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204276` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `58030` (56.67 KiB) | `18860` (18.42 KiB) | `18097` (17.67 KiB) | `3.077x` | `1.042x` | `3.207x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `22051` (21.53 KiB) | `16623` (16.23 KiB) | `16140` (15.76 KiB) | `1.327x` | `1.030x` | `1.366x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `25445` (24.85 KiB) | `23013` (22.47 KiB) | `21141` (20.65 KiB) | `1.106x` | `1.089x` | `1.204x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `48191` (47.06 KiB) | `14853` (14.50 KiB) | `13641` (13.32 KiB) | `3.245x` | `1.089x` | `3.533x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `49433` (48.27 KiB) | `22825` (22.29 KiB) | `20451` (19.97 KiB) | `2.166x` | `1.116x` | `2.417x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `16873` (16.48 KiB) | `16768` (16.38 KiB) | `16045` (15.67 KiB) | `1.006x` | `1.045x` | `1.052x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `91870` (89.72 KiB) | `37769` (36.88 KiB) | `32160` (31.41 KiB) | `2.432x` | `1.174x` | `2.857x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `13001` (12.70 KiB) | `14927` (14.58 KiB) | `14743` (14.40 KiB) | `0.871x` | `1.012x` | `0.882x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `28101` (27.44 KiB) | `16766` (16.37 KiB) | `15890` (15.52 KiB) | `1.676x` | `1.055x` | `1.768x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `45607` (44.54 KiB) | `41575` (40.60 KiB) | `28863` (28.19 KiB) | `1.097x` | `1.440x` | `1.580x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `213675` (208.67 KiB) | `0.999x` | `18.738x` | `18.720x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004953` (7.63 MiB) | `3638366` (3.47 MiB) | `0.999x` | `2.200x` | `2.199x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004952` (7.63 MiB) | `4382589` (4.18 MiB) | `0.999x` | `1.827x` | `1.825x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `204735` (199.94 KiB) | `0.999x` | `19.557x` | `19.537x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00000.parquet`: `34630` rows, `3359643` file bytes (3.20 MiB), `26895469` physical bytes (25.65 MiB), `15371784` encoded bytes (14.66 MiB), `3329104` compressed data bytes (3.17 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00001.parquet`: `34768` rows, `3334877` file bytes (3.18 MiB), `26768433` physical bytes (25.53 MiB), `15351223` encoded bytes (14.64 MiB), `3304338` compressed data bytes (3.15 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00002.parquet`: `34615` rows, `3354996` file bytes (3.20 MiB), `27036660` physical bytes (25.78 MiB), `15361809` encoded bytes (14.65 MiB), `3323730` compressed data bytes (3.17 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00003.parquet`: `34754` rows, `3321996` file bytes (3.17 MiB), `26917026` physical bytes (25.67 MiB), `15345812` encoded bytes (14.63 MiB), `3291636` compressed data bytes (3.14 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00004.parquet`: `34778` rows, `3330922` file bytes (3.18 MiB), `27080550` physical bytes (25.83 MiB), `15342449` encoded bytes (14.63 MiB), `3299754` compressed data bytes (3.15 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00005.parquet`: `34982` rows, `3315650` file bytes (3.16 MiB), `26908059` physical bytes (25.66 MiB), `15346771` encoded bytes (14.64 MiB), `3285059` compressed data bytes (3.13 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00006.parquet`: `34871` rows, `3323878` file bytes (3.17 MiB), `26922404` physical bytes (25.68 MiB), `15345674` encoded bytes (14.63 MiB), `3293337` compressed data bytes (3.14 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00007.parquet`: `34774` rows, `3332011` file bytes (3.18 MiB), `26971176` physical bytes (25.72 MiB), `15353461` encoded bytes (14.64 MiB), `3301132` compressed data bytes (3.15 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00008.parquet`: `34863` rows, `3315213` file bytes (3.16 MiB), `27009357` physical bytes (25.76 MiB), `15347171` encoded bytes (14.64 MiB), `3284277` compressed data bytes (3.13 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00009.parquet`: `34979` rows, `3306098` file bytes (3.15 MiB), `26838309` physical bytes (25.60 MiB), `15347002` encoded bytes (14.64 MiB), `3275640` compressed data bytes (3.12 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00010.parquet`: `34719` rows, `3313127` file bytes (3.16 MiB), `26829065` physical bytes (25.59 MiB), `15357382` encoded bytes (14.65 MiB), `3282612` compressed data bytes (3.13 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00011.parquet`: `35099` rows, `3280082` file bytes (3.13 MiB), `27081961` physical bytes (25.83 MiB), `15345004` encoded bytes (14.63 MiB), `3249466` compressed data bytes (3.10 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00012.parquet`: `35091` rows, `3285489` file bytes (3.13 MiB), `26877575` physical bytes (25.63 MiB), `15345269` encoded bytes (14.63 MiB), `3255173` compressed data bytes (3.10 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00013.parquet`: `34783` rows, `3325401` file bytes (3.17 MiB), `26790375` physical bytes (25.55 MiB), `15356330` encoded bytes (14.64 MiB), `3294724` compressed data bytes (3.14 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00014.parquet`: `32952` rows, `3550431` file bytes (3.39 MiB), `24463432` physical bytes (23.33 MiB), `15579517` encoded bytes (14.86 MiB), `3519992` compressed data bytes (3.36 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00015.parquet`: `30685` rows, `3752429` file bytes (3.58 MiB), `20912180` physical bytes (19.94 MiB), `15579802` encoded bytes (14.86 MiB), `3722937` compressed data bytes (3.55 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00016.parquet`: `29355` rows, `3806543` file bytes (3.63 MiB), `20756430` physical bytes (19.79 MiB), `15698203` encoded bytes (14.97 MiB), `3777355` compressed data bytes (3.60 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00017.parquet`: `30637` rows, `3953260` file bytes (3.77 MiB), `19953486` physical bytes (19.03 MiB), `15518189` encoded bytes (14.80 MiB), `3923895` compressed data bytes (3.74 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00018.parquet`: `30749` rows, `3970525` file bytes (3.79 MiB), `19808459` physical bytes (18.89 MiB), `15464317` encoded bytes (14.75 MiB), `3940630` compressed data bytes (3.76 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00019.parquet`: `29931` rows, `4060040` file bytes (3.87 MiB), `19578950` physical bytes (18.67 MiB), `15535346` encoded bytes (14.82 MiB), `4030554` compressed data bytes (3.84 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00020.parquet`: `30686` rows, `4003656` file bytes (3.82 MiB), `19556634` physical bytes (18.65 MiB), `15481861` encoded bytes (14.76 MiB), `3973809` compressed data bytes (3.79 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00021.parquet`: `30719` rows, `4002456` file bytes (3.82 MiB), `19648144` physical bytes (18.74 MiB), `15469781` encoded bytes (14.75 MiB), `3972692` compressed data bytes (3.79 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00022.parquet`: `30502` rows, `3993142` file bytes (3.81 MiB), `19621393` physical bytes (18.71 MiB), `15474362` encoded bytes (14.76 MiB), `3963394` compressed data bytes (3.78 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00023.parquet`: `30682` rows, `3978236` file bytes (3.79 MiB), `19749652` physical bytes (18.83 MiB), `15476860` encoded bytes (14.76 MiB), `3948227` compressed data bytes (3.77 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00024.parquet`: `30611` rows, `3991477` file bytes (3.81 MiB), `19672012` physical bytes (18.76 MiB), `15483783` encoded bytes (14.77 MiB), `3962028` compressed data bytes (3.78 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00025.parquet`: `30370` rows, `4007849` file bytes (3.82 MiB), `19602797` physical bytes (18.69 MiB), `15500755` encoded bytes (14.78 MiB), `3977929` compressed data bytes (3.79 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00026.parquet`: `31006` rows, `3971271` file bytes (3.79 MiB), `19681974` physical bytes (18.77 MiB), `15409579` encoded bytes (14.70 MiB), `3941899` compressed data bytes (3.76 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00027.parquet`: `30695` rows, `4008671` file bytes (3.82 MiB), `19675076` physical bytes (18.76 MiB), `15467243` encoded bytes (14.75 MiB), `3978704` compressed data bytes (3.79 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00028.parquet`: `31155` rows, `3951610` file bytes (3.77 MiB), `19835136` physical bytes (18.92 MiB), `15414113` encoded bytes (14.70 MiB), `3921838` compressed data bytes (3.74 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00029.parquet`: `30924` rows, `3971525` file bytes (3.79 MiB), `19615489` physical bytes (18.71 MiB), `15470401` encoded bytes (14.75 MiB), `3942064` compressed data bytes (3.76 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00030.parquet`: `20635` rows, `2703813` file bytes (2.58 MiB), `13340961` physical bytes (12.72 MiB), `10460705` encoded bytes (9.98 MiB), `2674310` compressed data bytes (2.55 MiB)
