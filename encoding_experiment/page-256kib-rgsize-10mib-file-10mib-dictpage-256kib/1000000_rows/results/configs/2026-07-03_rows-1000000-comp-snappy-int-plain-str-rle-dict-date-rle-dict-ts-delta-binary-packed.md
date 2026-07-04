# ClickBench Parquet Experiment

- Started: `2026-07-03T23:29:40-04:00`
- Write elapsed: `11.016s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `457692500` (436.49 MiB)
- Compressed column data bytes after codec compression: `103662460` (98.86 MiB)
- Parquet file bytes: `104580491` (99.74 MiB)
- Physical/encoded ratio: `1.557x`
- Encoded/compressed-data ratio: `4.415x`
- Physical/compressed-data ratio: `6.872x`
- Physical/parquet-file ratio: `6.812x`
- Files: `31`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `rle-dict`
- Date encoding: `rle-dict`
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
- Elapsed: `6.758s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004872` (7.63 MiB) | `8005298` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `262072` (255.93 KiB) | `0.999x` | `15.278x` | `15.263x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `138409995` (132.00 MiB) | `28821394` (27.49 MiB) | `10001711` (9.54 MiB) | `4.802x` | `2.882x` | `13.839x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204376` (199.59 KiB) | `0.999x` | `19.591x` | `19.572x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `3370220` (3.21 MiB) | `3035850` (2.90 MiB) | `2.374x` | `1.110x` | `2.635x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `5304` (5.18 KiB) | `5548` (5.42 KiB) | `754.148x` | `0.956x` | `720.981x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204392` (199.60 KiB) | `0.999x` | `19.589x` | `19.570x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003896` (3.82 MiB) | `719431` (702.57 KiB) | `0.999x` | `5.565x` | `5.560x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `396862` (387.56 KiB) | `0.999x` | `10.089x` | `10.079x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004875` (7.63 MiB) | `1084937` (1.03 MiB) | `0.999x` | `7.378x` | `7.374x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204202` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `318184` (310.73 KiB) | `0.999x` | `12.583x` | `12.571x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `342251` (334.23 KiB) | `0.999x` | `11.699x` | `11.687x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `88562192` (84.46 MiB) | `44095717` (42.05 MiB) | `16089644` (15.34 MiB) | `2.008x` | `2.741x` | `5.504x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `79583339` (75.90 MiB) | `34404083` (32.81 MiB) | `14799894` (14.11 MiB) | `2.313x` | `2.325x` | `5.377x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003867` (3.82 MiB) | `492013` (480.48 KiB) | `0.999x` | `8.138x` | `8.130x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003866` (3.82 MiB) | `509232` (497.30 KiB) | `0.999x` | `7.863x` | `7.855x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003856` (3.82 MiB) | `458157` (447.42 KiB) | `0.999x` | `8.739x` | `8.731x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `272030` (265.65 KiB) | `0.999x` | `14.718x` | `14.704x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `232300` (226.86 KiB) | `0.999x` | `17.236x` | `17.219x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `373729` (364.97 KiB) | `0.999x` | `10.713x` | `10.703x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `373444` (364.69 KiB) | `0.999x` | `10.721x` | `10.711x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `283042` (276.41 KiB) | `0.999x` | `14.146x` | `14.132x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `245353` (239.60 KiB) | `0.999x` | `16.319x` | `16.303x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `328405` (320.71 KiB) | `0.999x` | `12.192x` | `12.180x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `3354477` (3.20 MiB) | `267518` (261.25 KiB) | `182940` (178.65 KiB) | `12.539x` | `1.462x` | `18.336x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `219217` (214.08 KiB) | `0.999x` | `18.264x` | `18.247x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `217368` (212.27 KiB) | `0.999x` | `18.420x` | `18.402x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `365502` (356.94 KiB) | `0.999x` | `10.954x` | `10.944x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `3767530` (3.59 MiB) | `150729` (147.20 KiB) | `104580` (102.13 KiB) | `24.995x` | `1.441x` | `36.025x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204902` (200.10 KiB) | `0.999x` | `19.540x` | `19.522x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `205000` (200.20 KiB) | `0.999x` | `19.531x` | `19.512x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `218554` (213.43 KiB) | `0.999x` | `18.320x` | `18.302x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `215429` (210.38 KiB) | `0.999x` | `18.585x` | `18.568x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `81583` (79.67 KiB) | `25389` (24.79 KiB) | `23187` (22.64 KiB) | `3.213x` | `1.095x` | `3.518x` | `81583` (79.67 KiB) |
| `Params` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `0` (0 B) | `4328` (4.23 KiB) | `4572` (4.46 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003879` (3.82 MiB) | `561062` (547.91 KiB) | `0.999x` | `7.136x` | `7.129x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003874` (3.82 MiB) | `526762` (514.42 KiB) | `0.999x` | `7.601x` | `7.594x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `293169` (286.30 KiB) | `0.999x` | `13.657x` | `13.644x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `3528017` (3.36 MiB) | `1642849` (1.57 MiB) | `785059` (766.66 KiB) | `2.147x` | `2.093x` | `4.494x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `220369` (215.20 KiB) | `0.999x` | `18.169x` | `18.151x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003871` (3.82 MiB) | `505823` (493.97 KiB) | `0.999x` | `7.916x` | `7.908x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003878` (3.82 MiB) | `516700` (504.59 KiB) | `0.999x` | `7.749x` | `7.741x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003882` (3.82 MiB) | `552024` (539.09 KiB) | `0.999x` | `7.253x` | `7.246x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `285855` (279.16 KiB) | `0.999x` | `14.007x` | `13.993x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `3492150` (3.33 MiB) | `3057997` (2.92 MiB) | `2.291x` | `1.142x` | `2.616x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `287228` (280.50 KiB) | `0.999x` | `13.940x` | `13.926x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003833` (3.82 MiB) | `269778` (263.46 KiB) | `0.999x` | `14.841x` | `14.827x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `304894` (297.75 KiB) | `0.999x` | `13.132x` | `13.119x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `204241` (199.45 KiB) | `0.999x` | `19.604x` | `19.585x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `13587860` (12.96 MiB) | `9080` (8.87 KiB) | `9267` (9.05 KiB) | `1496.460x` | `0.980x` | `1466.263x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `205270` (200.46 KiB) | `0.999x` | `19.505x` | `19.487x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `256530` (250.52 KiB) | `0.999x` | `15.608x` | `15.593x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `206696` (201.85 KiB) | `0.999x` | `19.371x` | `19.352x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003844` (3.82 MiB) | `235410` (229.89 KiB) | `0.999x` | `17.008x` | `16.992x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004876` (7.63 MiB) | `1151638` (1.10 MiB) | `0.999x` | `6.951x` | `6.947x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `27797671` (26.51 MiB) | `21273003` (20.29 MiB) | `6158978` (5.87 MiB) | `1.307x` | `3.454x` | `4.513x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `3689302` (3.52 MiB) | `0.999x` | `1.085x` | `1.084x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204202` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204202` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204202` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003844` (3.82 MiB) | `295690` (288.76 KiB) | `0.999x` | `13.541x` | `13.528x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204202` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `1000000` (976.56 KiB) | `33737` (32.95 KiB) | `32031` (31.28 KiB) | `29.641x` | `1.053x` | `31.220x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `3384257` (3.23 MiB) | `3051369` (2.91 MiB) | `2.364x` | `1.109x` | `2.622x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `334114` (326.28 KiB) | `0.999x` | `11.983x` | `11.972x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `302596` (295.50 KiB) | `0.999x` | `13.232x` | `13.219x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `318351` (310.89 KiB) | `0.999x` | `12.577x` | `12.565x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `385675` (376.64 KiB) | `0.999x` | `10.381x` | `10.371x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `381919` (372.97 KiB) | `0.999x` | `10.483x` | `10.473x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `707695` (691.11 KiB) | `0.999x` | `5.658x` | `5.652x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `295456` (288.53 KiB) | `0.999x` | `13.551x` | `13.538x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204200` (199.41 KiB) | `0.999x` | `19.607x` | `19.589x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `258768` (252.70 KiB) | `0.999x` | `15.473x` | `15.458x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `2001192` (1.91 MiB) | `42787` (41.78 KiB) | `34725` (33.91 KiB) | `46.771x` | `1.232x` | `57.630x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `3325142` (3.17 MiB) | `136542` (133.34 KiB) | `94628` (92.41 KiB) | `24.353x` | `1.443x` | `35.139x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `0` (0 B) | `4328` (4.23 KiB) | `4572` (4.46 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `0` (0 B) | `4328` (4.23 KiB) | `4572` (4.46 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204202` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `269225` (262.92 KiB) | `0.999x` | `14.872x` | `14.857x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `327544` (319.87 KiB) | `0.999x` | `12.224x` | `12.212x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003874` (3.82 MiB) | `565311` (552.06 KiB) | `0.999x` | `7.083x` | `7.076x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003897` (3.82 MiB) | `1726545` (1.65 MiB) | `0.999x` | `2.319x` | `2.317x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `1283424` (1.22 MiB) | `0.999x` | `3.120x` | `3.117x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003894` (3.82 MiB) | `810219` (791.23 KiB) | `0.999x` | `4.942x` | `4.937x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204571` (199.78 KiB) | `0.999x` | `19.572x` | `19.553x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `1024` (1.00 KiB) | `5900` (5.76 KiB) | `6145` (6.00 KiB) | `0.174x` | `0.960x` | `0.167x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004815` (7.63 MiB) | `405191` (395.69 KiB) | `0.999x` | `19.756x` | `19.744x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `0` (0 B) | `4328` (4.23 KiB) | `4572` (4.46 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `3000000` (2.86 MiB) | `5243` (5.12 KiB) | `5487` (5.36 KiB) | `572.191x` | `0.956x` | `546.747x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204202` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `58030` (56.67 KiB) | `18614` (18.18 KiB) | `18001` (17.58 KiB) | `3.118x` | `1.034x` | `3.224x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `22051` (21.53 KiB) | `16269` (15.89 KiB) | `15880` (15.51 KiB) | `1.355x` | `1.024x` | `1.389x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `25445` (24.85 KiB) | `22611` (22.08 KiB) | `20874` (20.38 KiB) | `1.125x` | `1.083x` | `1.219x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `48191` (47.06 KiB) | `14695` (14.35 KiB) | `13506` (13.19 KiB) | `3.279x` | `1.088x` | `3.568x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `49433` (48.27 KiB) | `22743` (22.21 KiB) | `20396` (19.92 KiB) | `2.174x` | `1.115x` | `2.424x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `16873` (16.48 KiB) | `16701` (16.31 KiB) | `15995` (15.62 KiB) | `1.010x` | `1.044x` | `1.055x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `91870` (89.72 KiB) | `37836` (36.95 KiB) | `32061` (31.31 KiB) | `2.428x` | `1.180x` | `2.865x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `13001` (12.70 KiB) | `15062` (14.71 KiB) | `14781` (14.43 KiB) | `0.863x` | `1.019x` | `0.880x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `28101` (27.44 KiB) | `16742` (16.35 KiB) | `15876` (15.50 KiB) | `1.678x` | `1.055x` | `1.770x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `45607` (44.54 KiB) | `41001` (40.04 KiB) | `28745` (28.07 KiB) | `1.112x` | `1.426x` | `1.587x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `213598` (208.59 KiB) | `0.999x` | `18.745x` | `18.727x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004877` (7.63 MiB) | `3638048` (3.47 MiB) | `0.999x` | `2.200x` | `2.199x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004875` (7.63 MiB) | `4382600` (4.18 MiB) | `0.999x` | `1.827x` | `1.825x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204632` (199.84 KiB) | `0.999x` | `19.566x` | `19.547x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00000.parquet`: `34816` rows, `3198694` file bytes (3.05 MiB), `27043590` physical bytes (25.79 MiB), `15081960` encoded bytes (14.38 MiB), `3168212` compressed data bytes (3.02 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00001.parquet`: `34959` rows, `3166360` file bytes (3.02 MiB), `26902796` physical bytes (25.66 MiB), `15047544` encoded bytes (14.35 MiB), `3135842` compressed data bytes (2.99 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00002.parquet`: `34794` rows, `3193022` file bytes (3.05 MiB), `27189409` physical bytes (25.93 MiB), `15062389` encoded bytes (14.36 MiB), `3161833` compressed data bytes (3.02 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00003.parquet`: `34955` rows, `3158187` file bytes (3.01 MiB), `27054112` physical bytes (25.80 MiB), `15048432` encoded bytes (14.35 MiB), `3127882` compressed data bytes (2.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00004.parquet`: `34987` rows, `3162300` file bytes (3.02 MiB), `27218093` physical bytes (25.96 MiB), `15038393` encoded bytes (14.34 MiB), `3131189` compressed data bytes (2.99 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00005.parquet`: `35219` rows, `3146447` file bytes (3.00 MiB), `27094642` physical bytes (25.84 MiB), `15044653` encoded bytes (14.35 MiB), `3115498` compressed data bytes (2.97 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00006.parquet`: `35032` rows, `3157535` file bytes (3.01 MiB), `27028229` physical bytes (25.78 MiB), `15039851` encoded bytes (14.34 MiB), `3127095` compressed data bytes (2.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00007.parquet`: `34921` rows, `3169834` file bytes (3.02 MiB), `27168661` physical bytes (25.91 MiB), `15051416` encoded bytes (14.35 MiB), `3139079` compressed data bytes (2.99 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00008.parquet`: `35072` rows, `3148396` file bytes (3.00 MiB), `27135440` physical bytes (25.88 MiB), `15042582` encoded bytes (14.35 MiB), `3117585` compressed data bytes (2.97 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00009.parquet`: `35152` rows, `3133749` file bytes (2.99 MiB), `26961410` physical bytes (25.71 MiB), `15035997` encoded bytes (14.34 MiB), `3103395` compressed data bytes (2.96 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00010.parquet`: `34916` rows, `3151203` file bytes (3.01 MiB), `26959121` physical bytes (25.71 MiB), `15052996` encoded bytes (14.36 MiB), `3120727` compressed data bytes (2.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00011.parquet`: `35301` rows, `3113229` file bytes (2.97 MiB), `27265402` physical bytes (26.00 MiB), `15046680` encoded bytes (14.35 MiB), `3082666` compressed data bytes (2.94 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00012.parquet`: `35253` rows, `3119765` file bytes (2.98 MiB), `27001012` physical bytes (25.75 MiB), `15044165` encoded bytes (14.35 MiB), `3089534` compressed data bytes (2.95 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00013.parquet`: `34951` rows, `3156382` file bytes (3.01 MiB), `26896354` physical bytes (25.65 MiB), `15062041` encoded bytes (14.36 MiB), `3125841` compressed data bytes (2.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00014.parquet`: `33072` rows, `3383163` file bytes (3.23 MiB), `24319744` physical bytes (23.19 MiB), `15281706` encoded bytes (14.57 MiB), `3353145` compressed data bytes (3.20 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00015.parquet`: `31108` rows, `3556725` file bytes (3.39 MiB), `21251360` physical bytes (20.27 MiB), `15319777` encoded bytes (14.61 MiB), `3527265` compressed data bytes (3.36 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00016.parquet`: `29804` rows, `3600662` file bytes (3.43 MiB), `21209531` physical bytes (20.23 MiB), `15453903` encoded bytes (14.74 MiB), `3571647` compressed data bytes (3.41 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00017.parquet`: `31673` rows, `3749062` file bytes (3.58 MiB), `20270139` physical bytes (19.33 MiB), `15192420` encoded bytes (14.49 MiB), `3719558` compressed data bytes (3.55 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00018.parquet`: `31466` rows, `3764661` file bytes (3.59 MiB), `20331140` physical bytes (19.39 MiB), `15214848` encoded bytes (14.51 MiB), `3734837` compressed data bytes (3.56 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00019.parquet`: `30523` rows, `3847427` file bytes (3.67 MiB), `20041499` physical bytes (19.11 MiB), `15286262` encoded bytes (14.58 MiB), `3817993` compressed data bytes (3.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00020.parquet`: `31826` rows, `3734401` file bytes (3.56 MiB), `20129549` physical bytes (19.20 MiB), `15158497` encoded bytes (14.46 MiB), `3704538` compressed data bytes (3.53 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00021.parquet`: `31033` rows, `3830392` file bytes (3.65 MiB), `20027205` physical bytes (19.10 MiB), `15218830` encoded bytes (14.51 MiB), `3800666` compressed data bytes (3.62 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00022.parquet`: `31470` rows, `3755526` file bytes (3.58 MiB), `20181182` physical bytes (19.25 MiB), `15188540` encoded bytes (14.48 MiB), `3726023` compressed data bytes (3.55 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00023.parquet`: `31527` rows, `3762532` file bytes (3.59 MiB), `20150940` physical bytes (19.22 MiB), `15186295` encoded bytes (14.48 MiB), `3732910` compressed data bytes (3.56 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00024.parquet`: `31379` rows, `3749121` file bytes (3.58 MiB), `20172577` physical bytes (19.24 MiB), `15194366` encoded bytes (14.49 MiB), `3719494` compressed data bytes (3.55 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00025.parquet`: `30822` rows, `3831241` file bytes (3.65 MiB), `20010321` physical bytes (19.08 MiB), `15226268` encoded bytes (14.52 MiB), `3801666` compressed data bytes (3.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00026.parquet`: `32041` rows, `3728826` file bytes (3.56 MiB), `20277404` physical bytes (19.34 MiB), `15151749` encoded bytes (14.45 MiB), `3699336` compressed data bytes (3.53 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00027.parquet`: `31527` rows, `3772888` file bytes (3.60 MiB), `20159851` physical bytes (19.23 MiB), `15152342` encoded bytes (14.45 MiB), `3743217` compressed data bytes (3.57 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00028.parquet`: `31973` rows, `3723715` file bytes (3.55 MiB), `20218411` physical bytes (19.28 MiB), `15152924` encoded bytes (14.45 MiB), `3694177` compressed data bytes (3.52 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00029.parquet`: `31501` rows, `3757122` file bytes (3.58 MiB), `20220559` physical bytes (19.28 MiB), `15188609` encoded bytes (14.48 MiB), `3727544` compressed data bytes (3.55 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-delta-binary-packed/part-00030.parquet`: `6927` rows, `857924` file bytes (837.82 KiB), `4508941` physical bytes (4.30 MiB), `3426065` encoded bytes (3.27 MiB), `842066` compressed data bytes (822.33 KiB)
