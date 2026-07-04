# ClickBench Parquet Experiment

- Started: `2026-07-03T19:41:15-04:00`
- Write elapsed: `11.902s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `516844504` (492.90 MiB)
- Compressed column data bytes after codec compression: `124558596` (118.79 MiB)
- Parquet file bytes: `125503377` (119.69 MiB)
- Physical/encoded ratio: `1.378x`
- Encoded/compressed-data ratio: `4.149x`
- Physical/compressed-data ratio: `5.719x`
- Physical/parquet-file ratio: `5.676x`
- Files: `30`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `delta-byte-array`
- Date encoding: `plain`
- Timestamp encoding: `rle-dict`
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
- Elapsed: `7.155s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `8004798` (7.63 MiB) | `8005218` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003779` (3.82 MiB) | `262074` (255.93 KiB) | `0.999x` | `15.277x` | `15.263x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:516` | `1000000` | `138409995` (132.00 MiB) | `64474980` (61.49 MiB) | `17113334` (16.32 MiB) | `2.147x` | `3.768x` | `8.088x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003777` (3.82 MiB) | `204313` (199.52 KiB) | `0.999x` | `19.596x` | `19.578x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `7351408` (7.01 MiB) | `5533118` (5.28 MiB) | `1.088x` | `1.329x` | `1.446x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `204310` (199.52 KiB) | `0.999x` | `19.597x` | `19.578x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003780` (3.82 MiB) | `204341` (199.55 KiB) | `0.999x` | `19.594x` | `19.575x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003832` (3.82 MiB) | `719335` (702.48 KiB) | `0.999x` | `5.566x` | `5.561x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003781` (3.82 MiB) | `396592` (387.30 KiB) | `0.999x` | `10.095x` | `10.086x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `8004797` (7.63 MiB) | `1084853` (1.03 MiB) | `0.999x` | `7.379x` | `7.374x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `204145` (199.36 KiB) | `0.999x` | `19.612x` | `19.594x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `318015` (310.56 KiB) | `0.999x` | `12.590x` | `12.578x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003779` (3.82 MiB) | `342190` (334.17 KiB) | `0.999x` | `11.700x` | `11.689x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:340` | `1000000` | `88562192` (84.46 MiB) | `40459972` (38.59 MiB) | `18056749` (17.22 MiB) | `2.189x` | `2.241x` | `4.905x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:308` | `1000000` | `79583339` (75.90 MiB) | `38993379` (37.19 MiB) | `17334141` (16.53 MiB) | `2.041x` | `2.250x` | `4.591x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003809` (3.82 MiB) | `491896` (480.37 KiB) | `0.999x` | `8.140x` | `8.132x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003807` (3.82 MiB) | `509011` (497.08 KiB) | `0.999x` | `7.866x` | `7.858x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003798` (3.82 MiB) | `458114` (447.38 KiB) | `0.999x` | `8.740x` | `8.731x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `271974` (265.60 KiB) | `0.999x` | `14.721x` | `14.707x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003773` (3.82 MiB) | `232266` (226.82 KiB) | `0.999x` | `17.238x` | `17.222x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003781` (3.82 MiB) | `373383` (364.63 KiB) | `0.999x` | `10.723x` | `10.713x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003776` (3.82 MiB) | `373210` (364.46 KiB) | `0.999x` | `10.728x` | `10.718x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003776` (3.82 MiB) | `282953` (276.32 KiB) | `0.999x` | `14.150x` | `14.137x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003777` (3.82 MiB) | `245310` (239.56 KiB) | `0.999x` | `16.321x` | `16.306x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003778` (3.82 MiB) | `328373` (320.68 KiB) | `0.999x` | `12.193x` | `12.181x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `3354477` (3.20 MiB) | `1041625` (1017.21 KiB) | `395517` (386.25 KiB) | `3.220x` | `2.634x` | `8.481x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003778` (3.82 MiB) | `219168` (214.03 KiB) | `0.999x` | `18.268x` | `18.251x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003776` (3.82 MiB) | `217309` (212.22 KiB) | `0.999x` | `18.424x` | `18.407x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `365318` (356.76 KiB) | `0.999x` | `10.960x` | `10.949x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `3767530` (3.59 MiB) | `862795` (842.57 KiB) | `237078` (231.52 KiB) | `4.367x` | `3.639x` | `15.892x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003777` (3.82 MiB) | `204770` (199.97 KiB) | `0.999x` | `19.553x` | `19.534x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `204916` (200.11 KiB) | `0.999x` | `19.539x` | `19.520x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003777` (3.82 MiB) | `218464` (213.34 KiB) | `0.999x` | `18.327x` | `18.310x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003777` (3.82 MiB) | `215394` (210.35 KiB) | `0.999x` | `18.588x` | `18.571x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `81583` (79.67 KiB) | `282207` (275.59 KiB) | `57114` (55.78 KiB) | `0.289x` | `4.941x` | `1.428x` | `81583` (79.67 KiB) |
| `Params` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `0` (0 B) | `81485` (79.58 KiB) | `6943` (6.78 KiB) | `0.000x` | `11.736x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003813` (3.82 MiB) | `560874` (547.73 KiB) | `0.999x` | `7.139x` | `7.132x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003811` (3.82 MiB) | `526656` (514.31 KiB) | `0.999x` | `7.602x` | `7.595x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003776` (3.82 MiB) | `293127` (286.26 KiB) | `0.999x` | `13.659x` | `13.646x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `3528017` (3.36 MiB) | `2960390` (2.82 MiB) | `994289` (970.99 KiB) | `1.192x` | `2.977x` | `3.548x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003778` (3.82 MiB) | `220304` (215.14 KiB) | `0.999x` | `18.174x` | `18.157x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003812` (3.82 MiB) | `505772` (493.92 KiB) | `0.999x` | `7.916x` | `7.909x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003817` (3.82 MiB) | `516638` (504.53 KiB) | `0.999x` | `7.750x` | `7.742x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003818` (3.82 MiB) | `551461` (538.54 KiB) | `0.999x` | `7.260x` | `7.253x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003777` (3.82 MiB) | `285591` (278.90 KiB) | `0.999x` | `14.019x` | `14.006x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `7232264` (6.90 MiB) | `5462513` (5.21 MiB) | `1.106x` | `1.324x` | `1.465x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003776` (3.82 MiB) | `287205` (280.47 KiB) | `0.999x` | `13.940x` | `13.927x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003777` (3.82 MiB) | `269706` (263.38 KiB) | `0.999x` | `14.845x` | `14.831x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003780` (3.82 MiB) | `304875` (297.73 KiB) | `0.999x` | `13.133x` | `13.120x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `204183` (199.40 KiB) | `0.999x` | `19.609x` | `19.590x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:90` | `1000000` | `13587860` (12.96 MiB) | `142484` (139.14 KiB) | `27091` (26.46 KiB) | `95.364x` | `5.259x` | `501.564x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003776` (3.82 MiB) | `205144` (200.34 KiB) | `0.999x` | `19.517x` | `19.498x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `256485` (250.47 KiB) | `0.999x` | `15.610x` | `15.595x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `206624` (201.78 KiB) | `0.999x` | `19.377x` | `19.359x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003778` (3.82 MiB) | `235354` (229.84 KiB) | `0.999x` | `17.012x` | `16.996x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `8004797` (7.63 MiB) | `1151517` (1.10 MiB) | `0.999x` | `6.952x` | `6.947x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:140` | `1000000` | `27797671` (26.51 MiB) | `21056093` (20.08 MiB) | `6749265` (6.44 MiB) | `1.320x` | `3.120x` | `4.119x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `3688163` (3.52 MiB) | `0.999x` | `1.086x` | `1.085x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `204145` (199.36 KiB) | `0.999x` | `19.612x` | `19.594x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `204145` (199.36 KiB) | `0.999x` | `19.612x` | `19.594x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `204145` (199.36 KiB) | `0.999x` | `19.612x` | `19.594x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003780` (3.82 MiB) | `295641` (288.71 KiB) | `0.999x` | `13.543x` | `13.530x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `204145` (199.36 KiB) | `0.999x` | `19.612x` | `19.594x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `1000000` (976.56 KiB) | `207131` (202.28 KiB) | `74553` (72.81 KiB) | `4.828x` | `2.778x` | `13.413x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `7351801` (7.01 MiB) | `5534296` (5.28 MiB) | `1.088x` | `1.328x` | `1.446x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `333944` (326.12 KiB) | `0.999x` | `11.989x` | `11.978x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003777` (3.82 MiB) | `302534` (295.44 KiB) | `0.999x` | `13.234x` | `13.222x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003778` (3.82 MiB) | `318260` (310.80 KiB) | `0.999x` | `12.580x` | `12.568x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003780` (3.82 MiB) | `385521` (376.49 KiB) | `0.999x` | `10.385x` | `10.376x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003781` (3.82 MiB) | `381888` (372.94 KiB) | `0.999x` | `10.484x` | `10.474x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003830` (3.82 MiB) | `707048` (690.48 KiB) | `0.999x` | `5.663x` | `5.657x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003780` (3.82 MiB) | `295207` (288.29 KiB) | `0.999x` | `13.563x` | `13.550x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `204144` (199.36 KiB) | `0.999x` | `19.612x` | `19.594x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003782` (3.82 MiB) | `258301` (252.25 KiB) | `0.999x` | `15.500x` | `15.486x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `2001192` (1.91 MiB) | `333083` (325.28 KiB) | `89771` (87.67 KiB) | `6.008x` | `3.710x` | `22.292x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `3325142` (3.17 MiB) | `960812` (938.29 KiB) | `254160` (248.20 KiB) | `3.461x` | `3.780x` | `13.083x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `0` (0 B) | `81485` (79.58 KiB) | `6943` (6.78 KiB) | `0.000x` | `11.736x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `0` (0 B) | `81485` (79.58 KiB) | `6943` (6.78 KiB) | `0.000x` | `11.736x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `204145` (199.36 KiB) | `0.999x` | `19.612x` | `19.594x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `268732` (262.43 KiB) | `0.999x` | `14.899x` | `14.885x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003779` (3.82 MiB) | `327233` (319.56 KiB) | `0.999x` | `12.235x` | `12.224x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003814` (3.82 MiB) | `564579` (551.35 KiB) | `0.999x` | `7.092x` | `7.085x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `1724288` (1.64 MiB) | `0.999x` | `2.322x` | `2.320x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `1282306` (1.22 MiB) | `0.999x` | `3.122x` | `3.119x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `808929` (789.97 KiB) | `0.999x` | `4.950x` | `4.945x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003778` (3.82 MiB) | `204528` (199.73 KiB) | `0.999x` | `19.576x` | `19.557x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `1024` (1.00 KiB) | `86946` (84.91 KiB) | `9659` (9.43 KiB) | `0.012x` | `9.002x` | `0.106x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `8004738` (7.63 MiB) | `405111` (395.62 KiB) | `0.999x` | `19.759x` | `19.748x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `0` (0 B) | `81485` (79.58 KiB) | `6943` (6.78 KiB) | `0.000x` | `11.736x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `3000000` (2.86 MiB) | `86429` (84.40 KiB) | `10387` (10.14 KiB) | `34.711x` | `8.321x` | `288.823x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `204145` (199.36 KiB) | `0.999x` | `19.612x` | `19.594x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `58030` (56.67 KiB) | `301030` (293.97 KiB) | `48423` (47.29 KiB) | `0.193x` | `6.217x` | `1.198x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `22051` (21.53 KiB) | `197399` (192.77 KiB) | `38937` (38.02 KiB) | `0.112x` | `5.070x` | `0.566x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `25445` (24.85 KiB) | `208729` (203.84 KiB) | `44826` (43.78 KiB) | `0.122x` | `4.656x` | `0.568x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `48191` (47.06 KiB) | `231548` (226.12 KiB) | `34263` (33.46 KiB) | `0.208x` | `6.758x` | `1.407x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `49433` (48.27 KiB) | `267526` (261.26 KiB) | `52333` (51.11 KiB) | `0.185x` | `5.112x` | `0.945x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `16873` (16.48 KiB) | `213145` (208.15 KiB) | `39347` (38.42 KiB) | `0.079x` | `5.417x` | `0.429x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `91870` (89.72 KiB) | `327556` (319.88 KiB) | `68515` (66.91 KiB) | `0.280x` | `4.781x` | `1.341x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `13001` (12.70 KiB) | `149903` (146.39 KiB) | `27678` (27.03 KiB) | `0.087x` | `5.416x` | `0.470x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `28101` (27.44 KiB) | `190581` (186.11 KiB) | `31227` (30.50 KiB) | `0.147x` | `6.103x` | `0.900x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `45607` (44.54 KiB) | `261147` (255.03 KiB) | `50531` (49.35 KiB) | `0.175x` | `5.168x` | `0.903x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003778` (3.82 MiB) | `213557` (208.55 KiB) | `0.999x` | `18.748x` | `18.730x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `8004797` (7.63 MiB) | `3639089` (3.47 MiB) | `0.999x` | `2.200x` | `2.198x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `8004798` (7.63 MiB) | `4384522` (4.18 MiB) | `0.999x` | `1.826x` | `1.825x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003773` (3.82 MiB) | `204583` (199.79 KiB) | `0.999x` | `19.570x` | `19.552x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00000.parquet`: `34088` rows, `3899982` file bytes (3.72 MiB), `26480301` physical bytes (25.25 MiB), `16898839` encoded bytes (16.12 MiB), `3867966` compressed data bytes (3.69 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00001.parquet`: `33910` rows, `3805585` file bytes (3.63 MiB), `26076767` physical bytes (24.87 MiB), `16577328` encoded bytes (15.81 MiB), `3773243` compressed data bytes (3.60 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00002.parquet`: `33876` rows, `3876919` file bytes (3.70 MiB), `26524633` physical bytes (25.30 MiB), `16837007` encoded bytes (16.06 MiB), `3843735` compressed data bytes (3.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00003.parquet`: `34251` rows, `3841133` file bytes (3.66 MiB), `26505091` physical bytes (25.28 MiB), `16803622` encoded bytes (16.03 MiB), `3809152` compressed data bytes (3.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00004.parquet`: `34084` rows, `3861899` file bytes (3.68 MiB), `26504101` physical bytes (25.28 MiB), `16812057` encoded bytes (16.03 MiB), `3828676` compressed data bytes (3.65 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00005.parquet`: `34663` rows, `3848426` file bytes (3.67 MiB), `26729300` physical bytes (25.49 MiB), `16960498` encoded bytes (16.17 MiB), `3816019` compressed data bytes (3.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00006.parquet`: `34411` rows, `3861630` file bytes (3.68 MiB), `26477505` physical bytes (25.25 MiB), `16825278` encoded bytes (16.05 MiB), `3829046` compressed data bytes (3.65 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00007.parquet`: `34172` rows, `3843737` file bytes (3.67 MiB), `26534642` physical bytes (25.31 MiB), `16793923` encoded bytes (16.02 MiB), `3811039` compressed data bytes (3.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00008.parquet`: `34209` rows, `3841619` file bytes (3.66 MiB), `26514414` physical bytes (25.29 MiB), `16853828` encoded bytes (16.07 MiB), `3809002` compressed data bytes (3.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00009.parquet`: `34083` rows, `3769458` file bytes (3.59 MiB), `26137086` physical bytes (24.93 MiB), `16657776` encoded bytes (15.89 MiB), `3737019` compressed data bytes (3.56 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00010.parquet`: `33960` rows, `3815656` file bytes (3.64 MiB), `26318412` physical bytes (25.10 MiB), `16695408` encoded bytes (15.92 MiB), `3783593` compressed data bytes (3.61 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00011.parquet`: `34226` rows, `3751944` file bytes (3.58 MiB), `26360307` physical bytes (25.14 MiB), `16681487` encoded bytes (15.91 MiB), `3719345` compressed data bytes (3.55 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00012.parquet`: `34273` rows, `3799728` file bytes (3.62 MiB), `26311704` physical bytes (25.09 MiB), `16756965` encoded bytes (15.98 MiB), `3767556` compressed data bytes (3.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00013.parquet`: `34015` rows, `3797644` file bytes (3.62 MiB), `26067204` physical bytes (24.86 MiB), `16646939` encoded bytes (15.88 MiB), `3765336` compressed data bytes (3.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00014.parquet`: `34967` rows, `4052154` file bytes (3.86 MiB), `26739038` physical bytes (25.50 MiB), `17263559` encoded bytes (16.46 MiB), `4019628` compressed data bytes (3.83 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00015.parquet`: `33631` rows, `4350877` file bytes (4.15 MiB), `23004255` physical bytes (21.94 MiB), `17049548` encoded bytes (16.26 MiB), `4320488` compressed data bytes (4.12 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00016.parquet`: `32852` rows, `4391373` file bytes (4.19 MiB), `23065602` physical bytes (22.00 MiB), `16731351` encoded bytes (15.96 MiB), `4361160` compressed data bytes (4.16 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00017.parquet`: `32706` rows, `4496745` file bytes (4.29 MiB), `21308263` physical bytes (20.32 MiB), `17582667` encoded bytes (16.77 MiB), `4466483` compressed data bytes (4.26 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00018.parquet`: `32386` rows, `4573559` file bytes (4.36 MiB), `20935944` physical bytes (19.97 MiB), `17952242` encoded bytes (17.12 MiB), `4542802` compressed data bytes (4.33 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00019.parquet`: `32290` rows, `4738457` file bytes (4.52 MiB), `21184275` physical bytes (20.20 MiB), `18134630` encoded bytes (17.29 MiB), `4708029` compressed data bytes (4.49 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00020.parquet`: `32614` rows, `4489666` file bytes (4.28 MiB), `20607598` physical bytes (19.65 MiB), `17771131` encoded bytes (16.95 MiB), `4458966` compressed data bytes (4.25 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00021.parquet`: `32056` rows, `4614514` file bytes (4.40 MiB), `20710462` physical bytes (19.75 MiB), `17778997` encoded bytes (16.96 MiB), `4583941` compressed data bytes (4.37 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00022.parquet`: `32487` rows, `4536362` file bytes (4.33 MiB), `20759757` physical bytes (19.80 MiB), `17880816` encoded bytes (17.05 MiB), `4505638` compressed data bytes (4.30 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00023.parquet`: `32777` rows, `4615527` file bytes (4.40 MiB), `21027112` physical bytes (20.05 MiB), `17986922` encoded bytes (17.15 MiB), `4585019` compressed data bytes (4.37 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00024.parquet`: `32580` rows, `4613976` file bytes (4.40 MiB), `21045178` physical bytes (20.07 MiB), `18091780` encoded bytes (17.25 MiB), `4583260` compressed data bytes (4.37 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00025.parquet`: `32221` rows, `4600451` file bytes (4.39 MiB), `20698756` physical bytes (19.74 MiB), `17827069` encoded bytes (17.00 MiB), `4570136` compressed data bytes (4.36 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00026.parquet`: `32797` rows, `4511969` file bytes (4.30 MiB), `20789845` physical bytes (19.83 MiB), `17782254` encoded bytes (16.96 MiB), `4481483` compressed data bytes (4.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00027.parquet`: `32579` rows, `4583467` file bytes (4.37 MiB), `20856497` physical bytes (19.89 MiB), `17865904` encoded bytes (17.04 MiB), `4552641` compressed data bytes (4.34 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00028.parquet`: `32480` rows, `4407610` file bytes (4.20 MiB), `20456897` physical bytes (19.51 MiB), `17511520` encoded bytes (16.70 MiB), `4377224` compressed data bytes (4.17 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00029.parquet`: `30356` rows, `4311310` file bytes (4.11 MiB), `19667678` physical bytes (18.76 MiB), `16833159` encoded bytes (16.05 MiB), `4280971` compressed data bytes (4.08 MiB)
