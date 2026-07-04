# ClickBench Parquet Experiment

- Started: `2026-07-03T14:58:29-04:00`
- Write elapsed: `11.287s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `471455845` (449.62 MiB)
- Compressed column data bytes after codec compression: `107323777` (102.35 MiB)
- Parquet file bytes: `108242842` (103.23 MiB)
- Physical/encoded ratio: `1.511x`
- Encoded/compressed-data ratio: `4.393x`
- Physical/compressed-data ratio: `6.638x`
- Physical/parquet-file ratio: `6.581x`
- Files: `31`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `rle-dict`
- Date encoding: `rle-dict`
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
- Elapsed: `6.82s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004878` (7.63 MiB) | `8005305` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `262103` (255.96 KiB) | `0.999x` | `15.276x` | `15.261x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `138409995` (132.00 MiB) | `28819920` (27.48 MiB) | `9998730` (9.54 MiB) | `4.803x` | `2.882x` | `13.843x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204382` (199.59 KiB) | `0.999x` | `19.590x` | `19.571x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004875` (7.63 MiB) | `4285059` (4.09 MiB) | `0.999x` | `1.868x` | `1.867x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `5302` (5.18 KiB) | `5546` (5.42 KiB) | `754.432x` | `0.956x` | `721.241x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204404` (199.61 KiB) | `0.999x` | `19.588x` | `19.569x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003897` (3.82 MiB) | `719470` (702.61 KiB) | `0.999x` | `5.565x` | `5.560x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `396700` (387.40 KiB) | `0.999x` | `10.093x` | `10.083x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004876` (7.63 MiB) | `1084927` (1.03 MiB) | `0.999x` | `7.378x` | `7.374x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204209` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `318198` (310.74 KiB) | `0.999x` | `12.583x` | `12.571x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `342286` (334.26 KiB) | `0.999x` | `11.697x` | `11.686x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `88562192` (84.46 MiB) | `44093069` (42.05 MiB) | `16092558` (15.35 MiB) | `2.009x` | `2.740x` | `5.503x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `79583339` (75.90 MiB) | `34404040` (32.81 MiB) | `14801993` (14.12 MiB) | `2.313x` | `2.324x` | `5.377x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003871` (3.82 MiB) | `491988` (480.46 KiB) | `0.999x` | `8.138x` | `8.130x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003871` (3.82 MiB) | `509104` (497.17 KiB) | `0.999x` | `7.865x` | `7.857x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003859` (3.82 MiB) | `458134` (447.40 KiB) | `0.999x` | `8.739x` | `8.731x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `272065` (265.69 KiB) | `0.999x` | `14.716x` | `14.702x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `232275` (226.83 KiB) | `0.999x` | `17.237x` | `17.221x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `373598` (364.84 KiB) | `0.999x` | `10.717x` | `10.707x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `373296` (364.55 KiB) | `0.999x` | `10.726x` | `10.715x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `283057` (276.42 KiB) | `0.999x` | `14.145x` | `14.131x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `245334` (239.58 KiB) | `0.999x` | `16.320x` | `16.304x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `328457` (320.76 KiB) | `0.999x` | `12.190x` | `12.178x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `3354477` (3.20 MiB) | `267050` (260.79 KiB) | `182811` (178.53 KiB) | `12.561x` | `1.461x` | `18.349x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `219183` (214.05 KiB) | `0.999x` | `18.267x` | `18.250x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `217337` (212.24 KiB) | `0.999x` | `18.422x` | `18.405x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `365407` (356.84 KiB) | `0.999x` | `10.957x` | `10.947x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `3767530` (3.59 MiB) | `150911` (147.37 KiB) | `104428` (101.98 KiB) | `24.965x` | `1.445x` | `36.078x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204862` (200.06 KiB) | `0.999x` | `19.544x` | `19.525x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `205000` (200.20 KiB) | `0.999x` | `19.531x` | `19.512x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `218533` (213.41 KiB) | `0.999x` | `18.321x` | `18.304x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `215449` (210.40 KiB) | `0.999x` | `18.584x` | `18.566x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `81583` (79.67 KiB) | `25340` (24.75 KiB) | `23201` (22.66 KiB) | `3.220x` | `1.092x` | `3.516x` | `81583` (79.67 KiB) |
| `Params` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `0` (0 B) | `4326` (4.22 KiB) | `4570` (4.46 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003877` (3.82 MiB) | `561044` (547.89 KiB) | `0.999x` | `7.136x` | `7.130x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003872` (3.82 MiB) | `526795` (514.45 KiB) | `0.999x` | `7.600x` | `7.593x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `293140` (286.27 KiB) | `0.999x` | `13.658x` | `13.645x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `3528017` (3.36 MiB) | `1643253` (1.57 MiB) | `785429` (767.02 KiB) | `2.147x` | `2.092x` | `4.492x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `220369` (215.20 KiB) | `0.999x` | `18.169x` | `18.151x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003874` (3.82 MiB) | `505871` (494.01 KiB) | `0.999x` | `7.915x` | `7.907x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003879` (3.82 MiB) | `516526` (504.42 KiB) | `0.999x` | `7.752x` | `7.744x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003880` (3.82 MiB) | `551786` (538.85 KiB) | `0.999x` | `7.256x` | `7.249x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `285692` (279.00 KiB) | `0.999x` | `14.015x` | `14.001x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004874` (7.63 MiB) | `4231867` (4.04 MiB) | `0.999x` | `1.892x` | `1.890x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `287241` (280.51 KiB) | `0.999x` | `13.939x` | `13.926x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `269804` (263.48 KiB) | `0.999x` | `14.840x` | `14.826x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `304979` (297.83 KiB) | `0.999x` | `13.128x` | `13.116x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204250` (199.46 KiB) | `0.999x` | `19.603x` | `19.584x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `13587860` (12.96 MiB) | `9033` (8.82 KiB) | `9219` (9.00 KiB) | `1504.247x` | `0.980x` | `1473.897x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `205250` (200.44 KiB) | `0.999x` | `19.507x` | `19.488x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `256538` (250.53 KiB) | `0.999x` | `15.607x` | `15.592x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `206730` (201.88 KiB) | `0.999x` | `19.367x` | `19.349x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `235438` (229.92 KiB) | `0.999x` | `17.006x` | `16.990x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004875` (7.63 MiB) | `1151637` (1.10 MiB) | `0.999x` | `6.951x` | `6.947x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `27797671` (26.51 MiB) | `21272971` (20.29 MiB) | `6162385` (5.88 MiB) | `1.307x` | `3.452x` | `4.511x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `3689975` (3.52 MiB) | `0.999x` | `1.085x` | `1.084x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204209` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204209` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204209` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `295723` (288.79 KiB) | `0.999x` | `13.539x` | `13.526x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204209` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `1000000` (976.56 KiB) | `33745` (32.95 KiB) | `32076` (31.32 KiB) | `29.634x` | `1.052x` | `31.176x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004879` (7.63 MiB) | `4286395` (4.09 MiB) | `0.999x` | `1.868x` | `1.866x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `334074` (326.24 KiB) | `0.999x` | `11.985x` | `11.973x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `302687` (295.59 KiB) | `0.999x` | `13.228x` | `13.215x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `318348` (310.89 KiB) | `0.999x` | `12.577x` | `12.565x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `385678` (376.64 KiB) | `0.999x` | `10.381x` | `10.371x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `382166` (373.21 KiB) | `0.999x` | `10.477x` | `10.467x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `707258` (690.68 KiB) | `0.999x` | `5.661x` | `5.656x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003845` (3.82 MiB) | `295489` (288.56 KiB) | `0.999x` | `13.550x` | `13.537x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204210` (199.42 KiB) | `0.999x` | `19.606x` | `19.588x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003844` (3.82 MiB) | `258724` (252.66 KiB) | `0.999x` | `15.475x` | `15.460x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `2001192` (1.91 MiB) | `42301` (41.31 KiB) | `34547` (33.74 KiB) | `47.308x` | `1.224x` | `57.927x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `3325142` (3.17 MiB) | `136131` (132.94 KiB) | `94999` (92.77 KiB) | `24.426x` | `1.433x` | `35.002x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `0` (0 B) | `4326` (4.22 KiB) | `4570` (4.46 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `0` (0 B) | `4326` (4.22 KiB) | `4570` (4.46 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204209` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003845` (3.82 MiB) | `269095` (262.79 KiB) | `0.999x` | `14.879x` | `14.865x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `327333` (319.66 KiB) | `0.999x` | `12.232x` | `12.220x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003878` (3.82 MiB) | `565067` (551.82 KiB) | `0.999x` | `7.086x` | `7.079x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003897` (3.82 MiB) | `1726263` (1.65 MiB) | `0.999x` | `2.319x` | `2.317x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `1283297` (1.22 MiB) | `0.999x` | `3.120x` | `3.117x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003896` (3.82 MiB) | `809454` (790.48 KiB) | `0.999x` | `4.946x` | `4.942x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204589` (199.79 KiB) | `0.999x` | `19.570x` | `19.551x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `1024` (1.00 KiB) | `5889` (5.75 KiB) | `6134` (5.99 KiB) | `0.174x` | `0.960x` | `0.167x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004815` (7.63 MiB) | `405187` (395.69 KiB) | `0.999x` | `19.756x` | `19.744x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `0` (0 B) | `4326` (4.22 KiB) | `4570` (4.46 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `3000000` (2.86 MiB) | `5241` (5.12 KiB) | `5485` (5.36 KiB) | `572.410x` | `0.956x` | `546.946x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204209` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `58030` (56.67 KiB) | `18547` (18.11 KiB) | `17933` (17.51 KiB) | `3.129x` | `1.034x` | `3.236x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `22051` (21.53 KiB) | `16396` (16.01 KiB) | `15966` (15.59 KiB) | `1.345x` | `1.027x` | `1.381x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `25445` (24.85 KiB) | `22720` (22.19 KiB) | `20907` (20.42 KiB) | `1.120x` | `1.087x` | `1.217x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `48191` (47.06 KiB) | `14826` (14.48 KiB) | `13544` (13.23 KiB) | `3.250x` | `1.095x` | `3.558x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `49433` (48.27 KiB) | `22767` (22.23 KiB) | `20433` (19.95 KiB) | `2.171x` | `1.114x` | `2.419x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `16873` (16.48 KiB) | `16681` (16.29 KiB) | `15932` (15.56 KiB) | `1.012x` | `1.047x` | `1.059x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `91870` (89.72 KiB) | `37763` (36.88 KiB) | `31916` (31.17 KiB) | `2.433x` | `1.183x` | `2.878x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `13001` (12.70 KiB) | `15059` (14.71 KiB) | `14753` (14.41 KiB) | `0.863x` | `1.021x` | `0.881x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `28101` (27.44 KiB) | `16626` (16.24 KiB) | `15731` (15.36 KiB) | `1.690x` | `1.057x` | `1.786x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `45607` (44.54 KiB) | `41229` (40.26 KiB) | `28833` (28.16 KiB) | `1.106x` | `1.430x` | `1.582x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `213595` (208.59 KiB) | `0.999x` | `18.745x` | `18.727x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004876` (7.63 MiB) | `3637791` (3.47 MiB) | `0.999x` | `2.200x` | `2.199x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004877` (7.63 MiB) | `4382625` (4.18 MiB) | `0.999x` | `1.827x` | `1.825x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204652` (199.86 KiB) | `0.999x` | `19.564x` | `19.545x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00000.parquet`: `34747` rows, `3263278` file bytes (3.11 MiB), `26988909` physical bytes (25.74 MiB), `15520261` encoded bytes (14.80 MiB), `3232791` compressed data bytes (3.08 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00001.parquet`: `34883` rows, `3237312` file bytes (3.09 MiB), `26850638` physical bytes (25.61 MiB), `15499623` encoded bytes (14.78 MiB), `3206789` compressed data bytes (3.06 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00002.parquet`: `34715` rows, `3261642` file bytes (3.11 MiB), `27108514` physical bytes (25.85 MiB), `15514078` encoded bytes (14.80 MiB), `3230448` compressed data bytes (3.08 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00003.parquet`: `34886` rows, `3223565` file bytes (3.07 MiB), `27027381` physical bytes (25.78 MiB), `15495354` encoded bytes (14.78 MiB), `3193255` compressed data bytes (3.05 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00004.parquet`: `34917` rows, `3229868` file bytes (3.08 MiB), `27161351` physical bytes (25.90 MiB), `15487724` encoded bytes (14.77 MiB), `3198756` compressed data bytes (3.05 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00005.parquet`: `35109` rows, `3219988` file bytes (3.07 MiB), `27025820` physical bytes (25.77 MiB), `15499504` encoded bytes (14.78 MiB), `3189090` compressed data bytes (3.04 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00006.parquet`: `34965` rows, `3228841` file bytes (3.08 MiB), `26996278` physical bytes (25.75 MiB), `15495519` encoded bytes (14.78 MiB), `3198396` compressed data bytes (3.05 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00007.parquet`: `34881` rows, `3235283` file bytes (3.09 MiB), `27068700` physical bytes (25.81 MiB), `15502250` encoded bytes (14.78 MiB), `3204434` compressed data bytes (3.06 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00008.parquet`: `34980` rows, `3218308` file bytes (3.07 MiB), `27087584` physical bytes (25.83 MiB), `15495338` encoded bytes (14.78 MiB), `3187464` compressed data bytes (3.04 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00009.parquet`: `35079` rows, `3203689` file bytes (3.06 MiB), `26895615` physical bytes (25.65 MiB), `15493991` encoded bytes (14.78 MiB), `3173285` compressed data bytes (3.03 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00010.parquet`: `34855` rows, `3216344` file bytes (3.07 MiB), `26918659` physical bytes (25.67 MiB), `15501388` encoded bytes (14.78 MiB), `3185879` compressed data bytes (3.04 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00011.parquet`: `35210` rows, `3183390` file bytes (3.04 MiB), `27202222` physical bytes (25.94 MiB), `15499800` encoded bytes (14.78 MiB), `3152824` compressed data bytes (3.01 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00012.parquet`: `35184` rows, `3187887` file bytes (3.04 MiB), `26943596` physical bytes (25.70 MiB), `15493530` encoded bytes (14.78 MiB), `3157653` compressed data bytes (3.01 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00013.parquet`: `34853` rows, `3226856` file bytes (3.08 MiB), `26843528` physical bytes (25.60 MiB), `15511513` encoded bytes (14.79 MiB), `3196292` compressed data bytes (3.05 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00014.parquet`: `33127` rows, `3473839` file bytes (3.31 MiB), `24456872` physical bytes (23.32 MiB), `15698909` encoded bytes (14.97 MiB), `3443831` compressed data bytes (3.28 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00015.parquet`: `31036` rows, `3678103` file bytes (3.51 MiB), `21176041` physical bytes (20.20 MiB), `15672857` encoded bytes (14.95 MiB), `3648623` compressed data bytes (3.48 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00016.parquet`: `29725` rows, `3728762` file bytes (3.56 MiB), `21110164` physical bytes (20.13 MiB), `15784672` encoded bytes (15.05 MiB), `3699681` compressed data bytes (3.53 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00017.parquet`: `31451` rows, `3889030` file bytes (3.71 MiB), `20217133` physical bytes (19.28 MiB), `15556347` encoded bytes (14.84 MiB), `3859513` compressed data bytes (3.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00018.parquet`: `31219` rows, `3914810` file bytes (3.73 MiB), `20188913` physical bytes (19.25 MiB), `15565973` encoded bytes (14.84 MiB), `3885008` compressed data bytes (3.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00019.parquet`: `30499` rows, `3989986` file bytes (3.81 MiB), `19957018` physical bytes (19.03 MiB), `15612935` encoded bytes (14.89 MiB), `3960533` compressed data bytes (3.78 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00020.parquet`: `31485` rows, `3899500` file bytes (3.72 MiB), `19981995` physical bytes (19.06 MiB), `15528247` encoded bytes (14.81 MiB), `3869522` compressed data bytes (3.69 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00021.parquet`: `31001` rows, `3956033` file bytes (3.77 MiB), `19947662` physical bytes (19.02 MiB), `15562399` encoded bytes (14.84 MiB), `3926435` compressed data bytes (3.74 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00022.parquet`: `31300` rows, `3911491` file bytes (3.73 MiB), `20060404` physical bytes (19.13 MiB), `15538861` encoded bytes (14.82 MiB), `3881930` compressed data bytes (3.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00023.parquet`: `31194` rows, `3915733` file bytes (3.73 MiB), `20069531` physical bytes (19.14 MiB), `15566508` encoded bytes (14.85 MiB), `3886088` compressed data bytes (3.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00024.parquet`: `31402` rows, `3905267` file bytes (3.72 MiB), `20093163` physical bytes (19.16 MiB), `15555011` encoded bytes (14.83 MiB), `3875551` compressed data bytes (3.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00025.parquet`: `30618` rows, `3967873` file bytes (3.78 MiB), `19946457` physical bytes (19.02 MiB), `15578651` encoded bytes (14.86 MiB), `3938198` compressed data bytes (3.76 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00026.parquet`: `31910` rows, `3871803` file bytes (3.69 MiB), `20148970` physical bytes (19.22 MiB), `15492911` encoded bytes (14.78 MiB), `3842351` compressed data bytes (3.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00027.parquet`: `31448` rows, `3914840` file bytes (3.73 MiB), `20020732` physical bytes (19.09 MiB), `15496879` encoded bytes (14.78 MiB), `3885178` compressed data bytes (3.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00028.parquet`: `31636` rows, `3885412` file bytes (3.71 MiB), `20160208` physical bytes (19.23 MiB), `15519595` encoded bytes (14.80 MiB), `3855783` compressed data bytes (3.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00029.parquet`: `31436` rows, `3892750` file bytes (3.71 MiB), `20082750` physical bytes (19.15 MiB), `15540192` encoded bytes (14.82 MiB), `3863131` compressed data bytes (3.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00030.parquet`: `10249` rows, `1311359` file bytes (1.25 MiB), `6661816` physical bytes (6.35 MiB), `5175025` encoded bytes (4.94 MiB), `1295065` compressed data bytes (1.24 MiB)
