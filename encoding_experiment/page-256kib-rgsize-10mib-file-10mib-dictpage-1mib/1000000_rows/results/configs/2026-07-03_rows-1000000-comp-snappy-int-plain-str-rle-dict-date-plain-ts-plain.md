# ClickBench Parquet Experiment

- Started: `2026-07-03T15:28:39-04:00`
- Write elapsed: `11.029s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `475452603` (453.43 MiB)
- Compressed column data bytes after codec compression: `107523527` (102.54 MiB)
- Parquet file bytes: `108442185` (103.42 MiB)
- Physical/encoded ratio: `1.498x`
- Encoded/compressed-data ratio: `4.422x`
- Physical/compressed-data ratio: `6.626x`
- Physical/parquet-file ratio: `6.569x`
- Files: `31`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `rle-dict`
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
- Elapsed: `6.794s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004873` (7.63 MiB) | `8005300` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `262095` (255.95 KiB) | `0.999x` | `15.276x` | `15.262x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `138409995` (132.00 MiB) | `28817537` (27.48 MiB) | `9996348` (9.53 MiB) | `4.803x` | `2.883x` | `13.846x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204381` (199.59 KiB) | `0.999x` | `19.590x` | `19.571x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004878` (7.63 MiB) | `4285039` (4.09 MiB) | `0.999x` | `1.868x` | `1.867x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204377` (199.59 KiB) | `0.999x` | `19.590x` | `19.572x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204398` (199.61 KiB) | `0.999x` | `19.588x` | `19.570x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `719490` (702.63 KiB) | `0.999x` | `5.565x` | `5.559x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `396753` (387.45 KiB) | `0.999x` | `10.092x` | `10.082x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004875` (7.63 MiB) | `1084941` (1.03 MiB) | `0.999x` | `7.378x` | `7.374x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204202` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `318155` (310.70 KiB) | `0.999x` | `12.585x` | `12.572x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `342338` (334.31 KiB) | `0.999x` | `11.696x` | `11.684x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `88562192` (84.46 MiB) | `44090872` (42.05 MiB) | `16090587` (15.35 MiB) | `2.009x` | `2.740x` | `5.504x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `79583339` (75.90 MiB) | `34405980` (32.81 MiB) | `14804645` (14.12 MiB) | `2.313x` | `2.324x` | `5.376x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003871` (3.82 MiB) | `491931` (480.40 KiB) | `0.999x` | `8.139x` | `8.131x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003870` (3.82 MiB) | `509032` (497.10 KiB) | `0.999x` | `7.866x` | `7.858x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003858` (3.82 MiB) | `458150` (447.41 KiB) | `0.999x` | `8.739x` | `8.731x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `272092` (265.71 KiB) | `0.999x` | `14.715x` | `14.701x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `232259` (226.82 KiB) | `0.999x` | `17.239x` | `17.222x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `373600` (364.84 KiB) | `0.999x` | `10.717x` | `10.707x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `373335` (364.58 KiB) | `0.999x` | `10.725x` | `10.714x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `283056` (276.42 KiB) | `0.999x` | `14.145x` | `14.131x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `245383` (239.63 KiB) | `0.999x` | `16.317x` | `16.301x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `328495` (320.80 KiB) | `0.999x` | `12.188x` | `12.177x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `3354477` (3.20 MiB) | `267511` (261.24 KiB) | `182776` (178.49 KiB) | `12.540x` | `1.464x` | `18.353x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `219211` (214.07 KiB) | `0.999x` | `18.265x` | `18.247x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `217340` (212.25 KiB) | `0.999x` | `18.422x` | `18.404x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `365365` (356.80 KiB) | `0.999x` | `10.958x` | `10.948x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `3767530` (3.59 MiB) | `150942` (147.40 KiB) | `104518` (102.07 KiB) | `24.960x` | `1.444x` | `36.047x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `204841` (200.04 KiB) | `0.999x` | `19.546x` | `19.527x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `205005` (200.20 KiB) | `0.999x` | `19.530x` | `19.512x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `218547` (213.42 KiB) | `0.999x` | `18.320x` | `18.303x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `215434` (210.38 KiB) | `0.999x` | `18.585x` | `18.567x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `81583` (79.67 KiB) | `25363` (24.77 KiB) | `23210` (22.67 KiB) | `3.217x` | `1.093x` | `3.515x` | `81583` (79.67 KiB) |
| `Params` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `0` (0 B) | `4330` (4.23 KiB) | `4574` (4.47 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003878` (3.82 MiB) | `561058` (547.91 KiB) | `0.999x` | `7.136x` | `7.129x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003872` (3.82 MiB) | `526727` (514.38 KiB) | `0.999x` | `7.601x` | `7.594x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `293186` (286.31 KiB) | `0.999x` | `13.656x` | `13.643x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `3528017` (3.36 MiB) | `1644313` (1.57 MiB) | `785682` (767.27 KiB) | `2.146x` | `2.093x` | `4.490x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `220337` (215.17 KiB) | `0.999x` | `18.171x` | `18.154x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003873` (3.82 MiB) | `505847` (493.99 KiB) | `0.999x` | `7.915x` | `7.908x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003875` (3.82 MiB) | `516534` (504.43 KiB) | `0.999x` | `7.751x` | `7.744x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003880` (3.82 MiB) | `551751` (538.82 KiB) | `0.999x` | `7.257x` | `7.250x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `285763` (279.07 KiB) | `0.999x` | `14.011x` | `13.998x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004879` (7.63 MiB) | `4232065` (4.04 MiB) | `0.999x` | `1.891x` | `1.890x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `287262` (280.53 KiB) | `0.999x` | `13.938x` | `13.925x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `269786` (263.46 KiB) | `0.999x` | `14.841x` | `14.827x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `305000` (297.85 KiB) | `0.999x` | `13.127x` | `13.115x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `204239` (199.45 KiB) | `0.999x` | `19.604x` | `19.585x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `13587860` (12.96 MiB) | `9022` (8.81 KiB) | `9208` (8.99 KiB) | `1506.081x` | `0.980x` | `1475.658x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `205262` (200.45 KiB) | `0.999x` | `19.506x` | `19.487x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `256548` (250.54 KiB) | `0.999x` | `15.607x` | `15.592x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `206724` (201.88 KiB) | `0.999x` | `19.368x` | `19.349x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `235421` (229.90 KiB) | `0.999x` | `17.007x` | `16.991x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004877` (7.63 MiB) | `1151630` (1.10 MiB) | `0.999x` | `6.951x` | `6.947x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `27797671` (26.51 MiB) | `21272240` (20.29 MiB) | `6163240` (5.88 MiB) | `1.307x` | `3.451x` | `4.510x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `3690038` (3.52 MiB) | `0.999x` | `1.085x` | `1.084x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204202` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204202` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204202` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `295697` (288.77 KiB) | `0.999x` | `13.540x` | `13.527x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204202` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `1000000` (976.56 KiB) | `33758` (32.97 KiB) | `32068` (31.32 KiB) | `29.623x` | `1.053x` | `31.184x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004879` (7.63 MiB) | `4286442` (4.09 MiB) | `0.999x` | `1.867x` | `1.866x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003833` (3.82 MiB) | `334063` (326.23 KiB) | `0.999x` | `11.985x` | `11.974x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `302621` (295.53 KiB) | `0.999x` | `13.231x` | `13.218x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `318324` (310.86 KiB) | `0.999x` | `12.578x` | `12.566x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `385735` (376.69 KiB) | `0.999x` | `10.380x` | `10.370x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003845` (3.82 MiB) | `382155` (373.20 KiB) | `0.999x` | `10.477x` | `10.467x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003897` (3.82 MiB) | `707215` (690.64 KiB) | `0.999x` | `5.661x` | `5.656x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003844` (3.82 MiB) | `295348` (288.43 KiB) | `0.999x` | `13.556x` | `13.543x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `204201` (199.42 KiB) | `0.999x` | `19.607x` | `19.589x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `258707` (252.64 KiB) | `0.999x` | `15.476x` | `15.462x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `2001192` (1.91 MiB) | `42256` (41.27 KiB) | `34573` (33.76 KiB) | `47.359x` | `1.222x` | `57.883x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `3325142` (3.17 MiB) | `135990` (132.80 KiB) | `94899` (92.67 KiB) | `24.451x` | `1.433x` | `35.039x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `0` (0 B) | `4330` (4.23 KiB) | `4574` (4.47 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `0` (0 B) | `4330` (4.23 KiB) | `4574` (4.47 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204202` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `269138` (262.83 KiB) | `0.999x` | `14.877x` | `14.862x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `327405` (319.73 KiB) | `0.999x` | `12.229x` | `12.217x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003874` (3.82 MiB) | `565498` (552.24 KiB) | `0.999x` | `7.080x` | `7.073x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `1726268` (1.65 MiB) | `0.999x` | `2.319x` | `2.317x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `1283302` (1.22 MiB) | `0.999x` | `3.120x` | `3.117x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003897` (3.82 MiB) | `810446` (791.45 KiB) | `0.999x` | `4.940x` | `4.936x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `204567` (199.77 KiB) | `0.999x` | `19.572x` | `19.553x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `1024` (1.00 KiB) | `5890` (5.75 KiB) | `6135` (5.99 KiB) | `0.174x` | `0.960x` | `0.167x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004817` (7.63 MiB) | `405189` (395.69 KiB) | `0.999x` | `19.756x` | `19.744x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `0` (0 B) | `4330` (4.23 KiB) | `4574` (4.47 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `3000000` (2.86 MiB) | `5245` (5.12 KiB) | `5489` (5.36 KiB) | `571.973x` | `0.956x` | `546.548x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204202` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `58030` (56.67 KiB) | `18579` (18.14 KiB) | `17936` (17.52 KiB) | `3.123x` | `1.036x` | `3.235x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `22051` (21.53 KiB) | `16420` (16.04 KiB) | `15954` (15.58 KiB) | `1.343x` | `1.029x` | `1.382x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `25445` (24.85 KiB) | `22781` (22.25 KiB) | `20898` (20.41 KiB) | `1.117x` | `1.090x` | `1.218x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `48191` (47.06 KiB) | `14850` (14.50 KiB) | `13537` (13.22 KiB) | `3.245x` | `1.097x` | `3.560x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `49433` (48.27 KiB) | `22801` (22.27 KiB) | `20486` (20.01 KiB) | `2.168x` | `1.113x` | `2.413x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `16873` (16.48 KiB) | `16715` (16.32 KiB) | `15974` (15.60 KiB) | `1.009x` | `1.046x` | `1.056x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `91870` (89.72 KiB) | `37833` (36.95 KiB) | `31912` (31.16 KiB) | `2.428x` | `1.186x` | `2.879x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `13001` (12.70 KiB) | `15151` (14.80 KiB) | `14862` (14.51 KiB) | `0.858x` | `1.019x` | `0.875x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `28101` (27.44 KiB) | `16633` (16.24 KiB) | `15724` (15.36 KiB) | `1.689x` | `1.058x` | `1.787x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `45607` (44.54 KiB) | `41063` (40.10 KiB) | `28747` (28.07 KiB) | `1.111x` | `1.428x` | `1.586x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `213592` (208.59 KiB) | `0.999x` | `18.745x` | `18.727x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004878` (7.63 MiB) | `3637899` (3.47 MiB) | `0.999x` | `2.200x` | `2.199x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004876` (7.63 MiB) | `4382430` (4.18 MiB) | `0.999x` | `1.827x` | `1.825x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204646` (199.85 KiB) | `0.999x` | `19.565x` | `19.546x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00000.parquet`: `34740` rows, `3269540` file bytes (3.12 MiB), `26984110` physical bytes (25.73 MiB), `15656227` encoded bytes (14.93 MiB), `3239076` compressed data bytes (3.09 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00001.parquet`: `34875` rows, `3243633` file bytes (3.09 MiB), `26844666` physical bytes (25.60 MiB), `15636349` encoded bytes (14.91 MiB), `3213133` compressed data bytes (3.06 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00002.parquet`: `34709` rows, `3266936` file bytes (3.12 MiB), `27105438` physical bytes (25.85 MiB), `15648810` encoded bytes (14.92 MiB), `3235765` compressed data bytes (3.09 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00003.parquet`: `34882` rows, `3230352` file bytes (3.08 MiB), `27023048` physical bytes (25.77 MiB), `15633486` encoded bytes (14.91 MiB), `3200065` compressed data bytes (3.05 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00004.parquet`: `34902` rows, `3238419` file bytes (3.09 MiB), `27154154` physical bytes (25.90 MiB), `15625427` encoded bytes (14.90 MiB), `3207326` compressed data bytes (3.06 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00005.parquet`: `35106` rows, `3226029` file bytes (3.08 MiB), `27023649` physical bytes (25.77 MiB), `15637266` encoded bytes (14.91 MiB), `3195154` compressed data bytes (3.05 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00006.parquet`: `34970` rows, `3233699` file bytes (3.08 MiB), `26992902` physical bytes (25.74 MiB), `15632207` encoded bytes (14.91 MiB), `3203277` compressed data bytes (3.05 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00007.parquet`: `34877` rows, `3240536` file bytes (3.09 MiB), `27070071` physical bytes (25.82 MiB), `15638228` encoded bytes (14.91 MiB), `3209710` compressed data bytes (3.06 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00008.parquet`: `34964` rows, `3225833` file bytes (3.08 MiB), `27076039` physical bytes (25.82 MiB), `15632013` encoded bytes (14.91 MiB), `3195012` compressed data bytes (3.05 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00009.parquet`: `35075` rows, `3209919` file bytes (3.06 MiB), `26891257` physical bytes (25.65 MiB), `15630502` encoded bytes (14.91 MiB), `3179538` compressed data bytes (3.03 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00010.parquet`: `34850` rows, `3222442` file bytes (3.07 MiB), `26902493` physical bytes (25.66 MiB), `15637521` encoded bytes (14.91 MiB), `3192000` compressed data bytes (3.04 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00011.parquet`: `35190` rows, `3190329` file bytes (3.04 MiB), `27198549` physical bytes (25.94 MiB), `15638145` encoded bytes (14.91 MiB), `3159786` compressed data bytes (3.01 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00012.parquet`: `35177` rows, `3194621` file bytes (3.05 MiB), `26940272` physical bytes (25.69 MiB), `15631072` encoded bytes (14.91 MiB), `3164410` compressed data bytes (3.02 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00013.parquet`: `34846` rows, `3234957` file bytes (3.09 MiB), `26838163` physical bytes (25.59 MiB), `15649390` encoded bytes (14.92 MiB), `3204416` compressed data bytes (3.06 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00014.parquet`: `33147` rows, `3477120` file bytes (3.32 MiB), `24474433` physical bytes (23.34 MiB), `15829845` encoded bytes (15.10 MiB), `3447135` compressed data bytes (3.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00015.parquet`: `31030` rows, `3683173` file bytes (3.51 MiB), `21169412` physical bytes (20.19 MiB), `15794619` encoded bytes (15.06 MiB), `3653716` compressed data bytes (3.48 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00016.parquet`: `29723` rows, `3734828` file bytes (3.56 MiB), `21103741` physical bytes (20.13 MiB), `15901122` encoded bytes (15.16 MiB), `3705770` compressed data bytes (3.53 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00017.parquet`: `31445` rows, `3894388` file bytes (3.71 MiB), `20221514` physical bytes (19.28 MiB), `15680676` encoded bytes (14.95 MiB), `3864894` compressed data bytes (3.69 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00018.parquet`: `31197` rows, `3919046` file bytes (3.74 MiB), `20176605` physical bytes (19.24 MiB), `15683540` encoded bytes (14.96 MiB), `3889267` compressed data bytes (3.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00019.parquet`: `30538` rows, `3993676` file bytes (3.81 MiB), `19955974` physical bytes (19.03 MiB), `15724440` encoded bytes (15.00 MiB), `3964246` compressed data bytes (3.78 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00020.parquet`: `31414` rows, `3909914` file bytes (3.73 MiB), `19970112` physical bytes (19.04 MiB), `15654229` encoded bytes (14.93 MiB), `3879971` compressed data bytes (3.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00021.parquet`: `31002` rows, `3961268` file bytes (3.78 MiB), `19946834` physical bytes (19.02 MiB), `15684371` encoded bytes (14.96 MiB), `3931693` compressed data bytes (3.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00022.parquet`: `31312` rows, `3914773` file bytes (3.73 MiB), `20061866` physical bytes (19.13 MiB), `15661251` encoded bytes (14.94 MiB), `3885238` compressed data bytes (3.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00023.parquet`: `31183` rows, `3918868` file bytes (3.74 MiB), `20071926` physical bytes (19.14 MiB), `15686904` encoded bytes (14.96 MiB), `3889246` compressed data bytes (3.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00024.parquet`: `31404` rows, `3911811` file bytes (3.73 MiB), `20090936` physical bytes (19.16 MiB), `15678958` encoded bytes (14.95 MiB), `3882118` compressed data bytes (3.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00025.parquet`: `30624` rows, `3971821` file bytes (3.79 MiB), `19946538` physical bytes (19.02 MiB), `15699036` encoded bytes (14.97 MiB), `3942170` compressed data bytes (3.76 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00026.parquet`: `31895` rows, `3877000` file bytes (3.70 MiB), `20135872` physical bytes (19.20 MiB), `15614441` encoded bytes (14.89 MiB), `3847525` compressed data bytes (3.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00027.parquet`: `31421` rows, `3923783` file bytes (3.74 MiB), `20014009` physical bytes (19.09 MiB), `15619654` encoded bytes (14.90 MiB), `3893884` compressed data bytes (3.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00028.parquet`: `31637` rows, `3889147` file bytes (3.71 MiB), `20156708` physical bytes (19.22 MiB), `15641533` encoded bytes (14.92 MiB), `3859541` compressed data bytes (3.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00029.parquet`: `31426` rows, `3900081` file bytes (3.72 MiB), `20078628` physical bytes (19.15 MiB), `15666024` encoded bytes (14.94 MiB), `3870485` compressed data bytes (3.69 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00030.parquet`: `10439` rows, `1334243` file bytes (1.27 MiB), `6778705` physical bytes (6.46 MiB), `5305317` encoded bytes (5.06 MiB), `1317960` compressed data bytes (1.26 MiB)
