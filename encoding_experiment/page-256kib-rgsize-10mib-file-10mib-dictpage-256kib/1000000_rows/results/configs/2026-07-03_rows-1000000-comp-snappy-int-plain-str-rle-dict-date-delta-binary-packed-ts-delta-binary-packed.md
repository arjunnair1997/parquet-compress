# ClickBench Parquet Experiment

- Started: `2026-07-03T23:30:34-04:00`
- Write elapsed: `11.024s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `457737148` (436.53 MiB)
- Compressed column data bytes after codec compression: `103660946` (98.86 MiB)
- Parquet file bytes: `104578216` (99.73 MiB)
- Physical/encoded ratio: `1.556x`
- Encoded/compressed-data ratio: `4.416x`
- Physical/compressed-data ratio: `6.872x`
- Physical/parquet-file ratio: `6.812x`
- Files: `31`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `rle-dict`
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
- Elapsed: `6.753s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004875` (7.63 MiB) | `8005301` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `262070` (255.93 KiB) | `0.999x` | `15.278x` | `15.263x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `138409995` (132.00 MiB) | `28821487` (27.49 MiB) | `10001864` (9.54 MiB) | `4.802x` | `2.882x` | `13.838x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `204374` (199.58 KiB) | `0.999x` | `19.591x` | `19.572x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `3369665` (3.21 MiB) | `3034244` (2.89 MiB) | `2.374x` | `1.111x` | `2.637x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `51688` (50.48 KiB) | `8096` (7.91 KiB) | `77.387x` | `6.384x` | `494.071x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `204393` (199.60 KiB) | `0.999x` | `19.589x` | `19.570x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003896` (3.82 MiB) | `719432` (702.57 KiB) | `0.999x` | `5.565x` | `5.560x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `396857` (387.56 KiB) | `0.999x` | `10.089x` | `10.079x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004875` (7.63 MiB) | `1084943` (1.03 MiB) | `0.999x` | `7.378x` | `7.374x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204199` (199.41 KiB) | `0.999x` | `19.608x` | `19.589x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `318189` (310.73 KiB) | `0.999x` | `12.583x` | `12.571x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `342256` (334.23 KiB) | `0.999x` | `11.698x` | `11.687x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `88562192` (84.46 MiB) | `44096189` (42.05 MiB) | `16090293` (15.34 MiB) | `2.008x` | `2.741x` | `5.504x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `79583339` (75.90 MiB) | `34404316` (32.81 MiB) | `14799906` (14.11 MiB) | `2.313x` | `2.325x` | `5.377x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003866` (3.82 MiB) | `492015` (480.48 KiB) | `0.999x` | `8.138x` | `8.130x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003868` (3.82 MiB) | `509246` (497.31 KiB) | `0.999x` | `7.862x` | `7.855x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003854` (3.82 MiB) | `458156` (447.42 KiB) | `0.999x` | `8.739x` | `8.731x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `272038` (265.66 KiB) | `0.999x` | `14.718x` | `14.704x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `232304` (226.86 KiB) | `0.999x` | `17.235x` | `17.219x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `373729` (364.97 KiB) | `0.999x` | `10.713x` | `10.703x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `373469` (364.72 KiB) | `0.999x` | `10.721x` | `10.710x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `283028` (276.39 KiB) | `0.999x` | `14.146x` | `14.133x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003833` (3.82 MiB) | `245353` (239.60 KiB) | `0.999x` | `16.319x` | `16.303x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `328412` (320.71 KiB) | `0.999x` | `12.192x` | `12.180x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `3354477` (3.20 MiB) | `267342` (261.08 KiB) | `182847` (178.56 KiB) | `12.548x` | `1.462x` | `18.346x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `219211` (214.07 KiB) | `0.999x` | `18.265x` | `18.247x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `217364` (212.27 KiB) | `0.999x` | `18.420x` | `18.402x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `365480` (356.91 KiB) | `0.999x` | `10.955x` | `10.945x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `3767530` (3.59 MiB) | `150824` (147.29 KiB) | `104664` (102.21 KiB) | `24.980x` | `1.441x` | `35.996x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204897` (200.09 KiB) | `0.999x` | `19.541x` | `19.522x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `205002` (200.20 KiB) | `0.999x` | `19.531x` | `19.512x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `218547` (213.42 KiB) | `0.999x` | `18.320x` | `18.303x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `215430` (210.38 KiB) | `0.999x` | `18.585x` | `18.568x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `81583` (79.67 KiB) | `25410` (24.81 KiB) | `23198` (22.65 KiB) | `3.211x` | `1.095x` | `3.517x` | `81583` (79.67 KiB) |
| `Params` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `0` (0 B) | `4328` (4.23 KiB) | `4572` (4.46 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003880` (3.82 MiB) | `561062` (547.91 KiB) | `0.999x` | `7.136x` | `7.129x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003874` (3.82 MiB) | `526760` (514.41 KiB) | `0.999x` | `7.601x` | `7.594x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `293174` (286.30 KiB) | `0.999x` | `13.657x` | `13.644x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `3528017` (3.36 MiB) | `1642695` (1.57 MiB) | `785115` (766.71 KiB) | `2.148x` | `2.092x` | `4.494x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `220370` (215.21 KiB) | `0.999x` | `18.169x` | `18.151x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003872` (3.82 MiB) | `505824` (493.97 KiB) | `0.999x` | `7.916x` | `7.908x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003878` (3.82 MiB) | `516709` (504.60 KiB) | `0.999x` | `7.749x` | `7.741x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003881` (3.82 MiB) | `552016` (539.08 KiB) | `0.999x` | `7.253x` | `7.246x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `285851` (279.15 KiB) | `0.999x` | `14.007x` | `13.993x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `3491586` (3.33 MiB) | `3057783` (2.92 MiB) | `2.291x` | `1.142x` | `2.616x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `287233` (280.50 KiB) | `0.999x` | `13.939x` | `13.926x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `269780` (263.46 KiB) | `0.999x` | `14.841x` | `14.827x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `304903` (297.76 KiB) | `0.999x` | `13.132x` | `13.119x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204235` (199.45 KiB) | `0.999x` | `19.604x` | `19.585x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `13587860` (12.96 MiB) | `9059` (8.85 KiB) | `9246` (9.03 KiB) | `1499.929x` | `0.980x` | `1469.593x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `205274` (200.46 KiB) | `0.999x` | `19.505x` | `19.486x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `256535` (250.52 KiB) | `0.999x` | `15.607x` | `15.592x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `206688` (201.84 KiB) | `0.999x` | `19.371x` | `19.353x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `235402` (229.88 KiB) | `0.999x` | `17.009x` | `16.992x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004876` (7.63 MiB) | `1151651` (1.10 MiB) | `0.999x` | `6.951x` | `6.947x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `27797671` (26.51 MiB) | `21272405` (20.29 MiB) | `6159030` (5.87 MiB) | `1.307x` | `3.454x` | `4.513x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003897` (3.82 MiB) | `3689300` (3.52 MiB) | `0.999x` | `1.085x` | `1.084x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204199` (199.41 KiB) | `0.999x` | `19.608x` | `19.589x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204199` (199.41 KiB) | `0.999x` | `19.608x` | `19.589x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204199` (199.41 KiB) | `0.999x` | `19.608x` | `19.589x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003845` (3.82 MiB) | `295697` (288.77 KiB) | `0.999x` | `13.540x` | `13.527x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204199` (199.41 KiB) | `0.999x` | `19.608x` | `19.589x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `1000000` (976.56 KiB) | `33725` (32.93 KiB) | `32067` (31.32 KiB) | `29.652x` | `1.052x` | `31.185x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `3383782` (3.23 MiB) | `3048110` (2.91 MiB) | `2.364x` | `1.110x` | `2.625x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `334123` (326.29 KiB) | `0.999x` | `11.983x` | `11.972x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `302611` (295.52 KiB) | `0.999x` | `13.231x` | `13.218x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `318355` (310.89 KiB) | `0.999x` | `12.577x` | `12.565x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `385668` (376.63 KiB) | `0.999x` | `10.382x` | `10.372x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `381906` (372.96 KiB) | `0.999x` | `10.484x` | `10.474x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `707536` (690.95 KiB) | `0.999x` | `5.659x` | `5.653x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `295456` (288.53 KiB) | `0.999x` | `13.551x` | `13.538x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204198` (199.41 KiB) | `0.999x` | `19.608x` | `19.589x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `258772` (252.71 KiB) | `0.999x` | `15.472x` | `15.458x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `2001192` (1.91 MiB) | `42793` (41.79 KiB) | `34681` (33.87 KiB) | `46.764x` | `1.234x` | `57.703x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `3325142` (3.17 MiB) | `136422` (133.22 KiB) | `94628` (92.41 KiB) | `24.374x` | `1.442x` | `35.139x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `0` (0 B) | `4328` (4.23 KiB) | `4572` (4.46 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `0` (0 B) | `4328` (4.23 KiB) | `4572` (4.46 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204199` (199.41 KiB) | `0.999x` | `19.608x` | `19.589x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003844` (3.82 MiB) | `269232` (262.92 KiB) | `0.999x` | `14.871x` | `14.857x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `327591` (319.91 KiB) | `0.999x` | `12.222x` | `12.210x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003874` (3.82 MiB) | `565337` (552.09 KiB) | `0.999x` | `7.082x` | `7.075x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003897` (3.82 MiB) | `1726518` (1.65 MiB) | `0.999x` | `2.319x` | `2.317x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `1283635` (1.22 MiB) | `0.999x` | `3.119x` | `3.116x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003894` (3.82 MiB) | `810147` (791.16 KiB) | `0.999x` | `4.942x` | `4.937x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `204576` (199.78 KiB) | `0.999x` | `19.571x` | `19.553x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `1024` (1.00 KiB) | `5901` (5.76 KiB) | `6146` (6.00 KiB) | `0.174x` | `0.960x` | `0.167x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004814` (7.63 MiB) | `405186` (395.69 KiB) | `0.999x` | `19.756x` | `19.744x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `0` (0 B) | `4328` (4.23 KiB) | `4572` (4.46 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `3000000` (2.86 MiB) | `5243` (5.12 KiB) | `5487` (5.36 KiB) | `572.191x` | `0.956x` | `546.747x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204199` (199.41 KiB) | `0.999x` | `19.608x` | `19.589x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `58030` (56.67 KiB) | `18611` (18.17 KiB) | `18033` (17.61 KiB) | `3.118x` | `1.032x` | `3.218x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `22051` (21.53 KiB) | `16264` (15.88 KiB) | `15863` (15.49 KiB) | `1.356x` | `1.025x` | `1.390x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `25445` (24.85 KiB) | `22600` (22.07 KiB) | `20852` (20.36 KiB) | `1.126x` | `1.084x` | `1.220x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `48191` (47.06 KiB) | `14695` (14.35 KiB) | `13506` (13.19 KiB) | `3.279x` | `1.088x` | `3.568x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `49433` (48.27 KiB) | `22747` (22.21 KiB) | `20402` (19.92 KiB) | `2.173x` | `1.115x` | `2.423x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `16873` (16.48 KiB) | `16706` (16.31 KiB) | `16000` (15.62 KiB) | `1.010x` | `1.044x` | `1.055x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `91870` (89.72 KiB) | `37855` (36.97 KiB) | `32052` (31.30 KiB) | `2.427x` | `1.181x` | `2.866x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `13001` (12.70 KiB) | `15066` (14.71 KiB) | `14774` (14.43 KiB) | `0.863x` | `1.020x` | `0.880x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `28101` (27.44 KiB) | `16728` (16.34 KiB) | `15862` (15.49 KiB) | `1.680x` | `1.055x` | `1.772x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `45607` (44.54 KiB) | `41029` (40.07 KiB) | `28764` (28.09 KiB) | `1.112x` | `1.426x` | `1.586x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `213595` (208.59 KiB) | `0.999x` | `18.745x` | `18.727x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004875` (7.63 MiB) | `3638101` (3.47 MiB) | `0.999x` | `2.200x` | `2.199x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004876` (7.63 MiB) | `4382621` (4.18 MiB) | `0.999x` | `1.827x` | `1.825x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204628` (199.83 KiB) | `0.999x` | `19.566x` | `19.548x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00000.parquet`: `34816` rows, `3198760` file bytes (3.05 MiB), `27043590` physical bytes (25.79 MiB), `15083576` encoded bytes (14.38 MiB), `3168303` compressed data bytes (3.02 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00001.parquet`: `34959` rows, `3166426` file bytes (3.02 MiB), `26902796` physical bytes (25.66 MiB), `15049221` encoded bytes (14.35 MiB), `3135933` compressed data bytes (2.99 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00002.parquet`: `34794` rows, `3193087` file bytes (3.05 MiB), `27189409` physical bytes (25.93 MiB), `15063949` encoded bytes (14.37 MiB), `3161923` compressed data bytes (3.02 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00003.parquet`: `34955` rows, `3158259` file bytes (3.01 MiB), `27054112` physical bytes (25.80 MiB), `15050053` encoded bytes (14.35 MiB), `3127979` compressed data bytes (2.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00004.parquet`: `34987` rows, `3162371` file bytes (3.02 MiB), `27218093` physical bytes (25.96 MiB), `15040070` encoded bytes (14.34 MiB), `3131285` compressed data bytes (2.99 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00005.parquet`: `35219` rows, `3146483` file bytes (3.00 MiB), `27094642` physical bytes (25.84 MiB), `15046057` encoded bytes (14.35 MiB), `3115559` compressed data bytes (2.97 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00006.parquet`: `35032` rows, `3157599` file bytes (3.01 MiB), `27028229` physical bytes (25.78 MiB), `15041364` encoded bytes (14.34 MiB), `3127184` compressed data bytes (2.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00007.parquet`: `34921` rows, `3169902` file bytes (3.02 MiB), `27168661` physical bytes (25.91 MiB), `15052981` encoded bytes (14.36 MiB), `3139172` compressed data bytes (2.99 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00008.parquet`: `35072` rows, `3148459` file bytes (3.00 MiB), `27135440` physical bytes (25.88 MiB), `15044208` encoded bytes (14.35 MiB), `3117673` compressed data bytes (2.97 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00009.parquet`: `35152` rows, `3133806` file bytes (2.99 MiB), `26961410` physical bytes (25.71 MiB), `15037571` encoded bytes (14.34 MiB), `3103477` compressed data bytes (2.96 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00010.parquet`: `34916` rows, `3151276` file bytes (3.01 MiB), `26959121` physical bytes (25.71 MiB), `15054780` encoded bytes (14.36 MiB), `3120825` compressed data bytes (2.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00011.parquet`: `35301` rows, `3113291` file bytes (2.97 MiB), `27265402` physical bytes (26.00 MiB), `15048260` encoded bytes (14.35 MiB), `3082753` compressed data bytes (2.94 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00012.parquet`: `35253` rows, `3119825` file bytes (2.98 MiB), `27001012` physical bytes (25.75 MiB), `15045689` encoded bytes (14.35 MiB), `3089619` compressed data bytes (2.95 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00013.parquet`: `34951` rows, `3156442` file bytes (3.01 MiB), `26896354` physical bytes (25.65 MiB), `15063661` encoded bytes (14.37 MiB), `3125926` compressed data bytes (2.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00014.parquet`: `33072` rows, `3383211` file bytes (3.23 MiB), `24319744` physical bytes (23.19 MiB), `15283144` encoded bytes (14.58 MiB), `3353218` compressed data bytes (3.20 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00015.parquet`: `31108` rows, `3556786` file bytes (3.39 MiB), `21251360` physical bytes (20.27 MiB), `15321246` encoded bytes (14.61 MiB), `3527351` compressed data bytes (3.36 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00016.parquet`: `29804` rows, `3600710` file bytes (3.43 MiB), `21209531` physical bytes (20.23 MiB), `15455323` encoded bytes (14.74 MiB), `3571720` compressed data bytes (3.41 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00017.parquet`: `31673` rows, `3748841` file bytes (3.58 MiB), `20270139` physical bytes (19.33 MiB), `15193902` encoded bytes (14.49 MiB), `3719362` compressed data bytes (3.55 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00018.parquet`: `31466` rows, `3764718` file bytes (3.59 MiB), `20331140` physical bytes (19.39 MiB), `15216278` encoded bytes (14.51 MiB), `3734919` compressed data bytes (3.56 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00019.parquet`: `30523` rows, `3847489` file bytes (3.67 MiB), `20041499` physical bytes (19.11 MiB), `15287820` encoded bytes (14.58 MiB), `3818080` compressed data bytes (3.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00020.parquet`: `31825` rows, `3734349` file bytes (3.56 MiB), `20128966` physical bytes (19.20 MiB), `15159473` encoded bytes (14.46 MiB), `3704511` compressed data bytes (3.53 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00021.parquet`: `31033` rows, `3827325` file bytes (3.65 MiB), `20026803` physical bytes (19.10 MiB), `15219448` encoded bytes (14.51 MiB), `3797624` compressed data bytes (3.62 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00022.parquet`: `31472` rows, `3752783` file bytes (3.58 MiB), `20183221` physical bytes (19.25 MiB), `15192301` encoded bytes (14.49 MiB), `3723305` compressed data bytes (3.55 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00023.parquet`: `31528` rows, `3762798` file bytes (3.59 MiB), `20150864` physical bytes (19.22 MiB), `15187657` encoded bytes (14.48 MiB), `3733201` compressed data bytes (3.56 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00024.parquet`: `31380` rows, `3745271` file bytes (3.57 MiB), `20173514` physical bytes (19.24 MiB), `15196544` encoded bytes (14.49 MiB), `3715669` compressed data bytes (3.54 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00025.parquet`: `30822` rows, `3830034` file bytes (3.65 MiB), `20010167` physical bytes (19.08 MiB), `15226892` encoded bytes (14.52 MiB), `3800484` compressed data bytes (3.62 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00026.parquet`: `32036` rows, `3732275` file bytes (3.56 MiB), `20273764` physical bytes (19.33 MiB), `15150859` encoded bytes (14.45 MiB), `3702810` compressed data bytes (3.53 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00027.parquet`: `31527` rows, `3776670` file bytes (3.60 MiB), `20160076` physical bytes (19.23 MiB), `15153831` encoded bytes (14.45 MiB), `3747024` compressed data bytes (3.57 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00028.parquet`: `31973` rows, `3723832` file bytes (3.55 MiB), `20218073` physical bytes (19.28 MiB), `15153622` encoded bytes (14.45 MiB), `3694319` compressed data bytes (3.52 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00029.parquet`: `31501` rows, `3756897` file bytes (3.58 MiB), `20220733` physical bytes (19.28 MiB), `15190081` encoded bytes (14.49 MiB), `3727344` compressed data bytes (3.55 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00030.parquet`: `6929` rows, `858241` file bytes (838.13 KiB), `4510759` physical bytes (4.30 MiB), `3427287` encoded bytes (3.27 MiB), `842394` compressed data bytes (822.65 KiB)
