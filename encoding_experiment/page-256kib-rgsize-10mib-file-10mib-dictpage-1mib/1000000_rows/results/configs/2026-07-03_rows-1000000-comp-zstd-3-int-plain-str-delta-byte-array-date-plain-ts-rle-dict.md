# ClickBench Parquet Experiment

- Started: `2026-07-03T15:31:45-04:00`
- Write elapsed: `12.245s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `516817500` (492.88 MiB)
- Compressed column data bytes after codec compression: `92266855` (87.99 MiB)
- Parquet file bytes: `93168787` (88.85 MiB)
- Physical/encoded ratio: `1.378x`
- Encoded/compressed-data ratio: `5.601x`
- Physical/compressed-data ratio: `7.721x`
- Physical/parquet-file ratio: `7.646x`
- Files: `29`

## Settings

- Compression: `zstd-3`
- Int encoding: `plain`
- String encoding: `delta-byte-array`
- Date encoding: `plain`
- Timestamp encoding: `rle-dict`
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
- Files read: `29`
- Elapsed: `7.59s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004558` (7.63 MiB) | `8005362` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `64123` (62.62 KiB) | `0.999x` | `62.436x` | `62.380x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:528` | `1000000` | `138409995` (132.00 MiB) | `64475339` (61.49 MiB) | `13981029` (13.33 MiB) | `2.147x` | `4.612x` | `9.900x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003530` (3.82 MiB) | `4921` (4.81 KiB) | `0.999x` | `813.560x` | `812.843x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `7359703` (7.02 MiB) | `4041475` (3.85 MiB) | `1.087x` | `1.821x` | `1.979x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4919` (4.80 KiB) | `0.999x` | `813.891x` | `813.173x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003527` (3.82 MiB) | `4936` (4.82 KiB) | `0.999x` | `811.087x` | `810.373x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003600` (3.82 MiB) | `408182` (398.62 KiB) | `0.999x` | `9.808x` | `9.800x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `190725` (186.25 KiB) | `0.999x` | `20.991x` | `20.973x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004556` (7.63 MiB) | `617987` (603.50 KiB) | `0.999x` | `12.953x` | `12.945x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4236` (4.14 KiB) | `0.999x` | `945.120x` | `944.287x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003582` (3.82 MiB) | `106910` (104.40 KiB) | `0.999x` | `37.448x` | `37.415x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `133855` (130.72 KiB) | `0.999x` | `29.910x` | `29.883x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:356` | `1000000` | `88562192` (84.46 MiB) | `40465757` (38.59 MiB) | `15085689` (14.39 MiB) | `2.189x` | `2.682x` | `5.871x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:328` | `1000000` | `79583339` (75.90 MiB) | `38991609` (37.19 MiB) | `14559562` (13.89 MiB) | `2.041x` | `2.678x` | `5.466x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `178944` (174.75 KiB) | `0.999x` | `22.373x` | `22.353x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003591` (3.82 MiB) | `275231` (268.78 KiB) | `0.999x` | `14.546x` | `14.533x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `231312` (225.89 KiB) | `0.999x` | `17.308x` | `17.293x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `87526` (85.47 KiB) | `0.999x` | `45.742x` | `45.701x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `48056` (46.93 KiB) | `0.999x` | `83.311x` | `83.236x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `187256` (182.87 KiB) | `0.999x` | `21.380x` | `21.361x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `186166` (181.80 KiB) | `0.999x` | `21.505x` | `21.486x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `82209` (80.28 KiB) | `0.999x` | `48.700x` | `48.656x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `53280` (52.03 KiB) | `0.999x` | `75.142x` | `75.075x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `128717` (125.70 KiB) | `0.999x` | `31.104x` | `31.076x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `3354477` (3.20 MiB) | `1044596` (1020.11 KiB) | `326497` (318.84 KiB) | `3.211x` | `3.199x` | `10.274x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `26697` (26.07 KiB) | `0.999x` | `149.964x` | `149.830x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `25196` (24.61 KiB) | `0.999x` | `158.898x` | `158.755x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `155696` (152.05 KiB) | `0.999x` | `25.714x` | `25.691x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `3767530` (3.59 MiB) | `861087` (840.91 KiB) | `179333` (175.13 KiB) | `4.375x` | `4.802x` | `21.009x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003541` (3.82 MiB) | `6063` (5.92 KiB) | `0.999x` | `660.323x` | `659.739x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003546` (3.82 MiB) | `6568` (6.41 KiB) | `0.999x` | `609.553x` | `609.013x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003582` (3.82 MiB) | `28938` (28.26 KiB) | `0.999x` | `138.350x` | `138.227x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `22963` (22.42 KiB) | `0.999x` | `174.350x` | `174.193x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `81583` (79.67 KiB) | `282106` (275.49 KiB) | `40971` (40.01 KiB) | `0.289x` | `6.886x` | `1.991x` | `81583` (79.67 KiB) |
| `Params` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `81269` (79.36 KiB) | `3684` (3.60 KiB) | `0.000x` | `22.060x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003591` (3.82 MiB) | `323725` (316.14 KiB) | `0.999x` | `12.367x` | `12.356x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `288216` (281.46 KiB) | `0.999x` | `13.891x` | `13.878x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003582` (3.82 MiB) | `101752` (99.37 KiB) | `0.999x` | `39.346x` | `39.311x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `3528017` (3.36 MiB) | `2962047` (2.82 MiB) | `826795` (807.42 KiB) | `1.191x` | `3.583x` | `4.267x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `30708` (29.99 KiB) | `0.999x` | `130.376x` | `130.259x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `164905` (161.04 KiB) | `0.999x` | `24.278x` | `24.256x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `305937` (298.77 KiB) | `0.999x` | `13.086x` | `13.075x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003592` (3.82 MiB) | `320051` (312.55 KiB) | `0.999x` | `12.509x` | `12.498x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `99919` (97.58 KiB) | `0.999x` | `40.068x` | `40.032x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `7210438` (6.88 MiB) | `3978352` (3.79 MiB) | `1.110x` | `1.812x` | `2.011x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `90795` (88.67 KiB) | `0.999x` | `44.095x` | `44.055x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `74555` (72.81 KiB) | `0.999x` | `53.700x` | `53.652x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `123971` (121.07 KiB) | `0.999x` | `32.295x` | `32.266x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4362` (4.26 KiB) | `0.999x` | `917.820x` | `917.011x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:85` | `1000000` | `13587860` (12.96 MiB) | `141442` (138.13 KiB) | `21595` (21.09 KiB) | `96.067x` | `6.550x` | `629.213x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003552` (3.82 MiB) | `7113` (6.95 KiB) | `0.999x` | `562.850x` | `562.351x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003558` (3.82 MiB) | `56764` (55.43 KiB) | `0.999x` | `70.530x` | `70.467x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003556` (3.82 MiB) | `8054` (7.87 KiB) | `0.999x` | `497.089x` | `496.648x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003539` (3.82 MiB) | `25192` (24.60 KiB) | `0.999x` | `158.921x` | `158.781x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004554` (7.63 MiB) | `694044` (677.78 KiB) | `0.999x` | `11.533x` | `11.527x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:137` | `1000000` | `27797671` (26.51 MiB) | `21050084` (20.07 MiB) | `5588654` (5.33 MiB) | `1.321x` | `3.767x` | `4.974x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003646` (3.82 MiB) | `3804388` (3.63 MiB) | `0.999x` | `1.052x` | `1.051x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4236` (4.14 KiB) | `0.999x` | `945.120x` | `944.287x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4236` (4.14 KiB) | `0.999x` | `945.120x` | `944.287x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4236` (4.14 KiB) | `0.999x` | `945.120x` | `944.287x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `82616` (80.68 KiB) | `0.999x` | `48.460x` | `48.417x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4236` (4.14 KiB) | `0.999x` | `945.120x` | `944.287x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `1000000` (976.56 KiB) | `206697` (201.85 KiB) | `50880` (49.69 KiB) | `4.838x` | `4.062x` | `19.654x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `7355946` (7.02 MiB) | `4042063` (3.85 MiB) | `1.088x` | `1.820x` | `1.979x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `142529` (139.19 KiB) | `0.999x` | `28.090x` | `28.064x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `107837` (105.31 KiB) | `0.999x` | `37.126x` | `37.093x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `123045` (120.16 KiB) | `0.999x` | `32.538x` | `32.508x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `193369` (188.84 KiB) | `0.999x` | `20.704x` | `20.686x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `174024` (169.95 KiB) | `0.999x` | `23.006x` | `22.985x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003606` (3.82 MiB) | `426433` (416.44 KiB) | `0.999x` | `9.389x` | `9.380x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003578` (3.82 MiB) | `70807` (69.15 KiB) | `0.999x` | `56.542x` | `56.492x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003532` (3.82 MiB) | `4239` (4.14 KiB) | `0.999x` | `944.452x` | `943.619x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003575` (3.82 MiB) | `55254` (53.96 KiB) | `0.999x` | `72.458x` | `72.393x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `2001192` (1.91 MiB) | `332831` (325.03 KiB) | `65298` (63.77 KiB) | `6.013x` | `5.097x` | `30.647x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `3325142` (3.17 MiB) | `962259` (939.71 KiB) | `195796` (191.21 KiB) | `3.456x` | `4.915x` | `16.983x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `81269` (79.36 KiB) | `3684` (3.60 KiB) | `0.000x` | `22.060x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `81269` (79.36 KiB) | `3684` (3.60 KiB) | `0.000x` | `22.060x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4236` (4.14 KiB) | `0.999x` | `945.120x` | `944.287x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003543` (3.82 MiB) | `61498` (60.06 KiB) | `0.999x` | `65.100x` | `65.043x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `135486` (132.31 KiB) | `0.999x` | `29.550x` | `29.523x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003600` (3.82 MiB) | `333280` (325.47 KiB) | `0.999x` | `12.013x` | `12.002x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003646` (3.82 MiB) | `1245880` (1.19 MiB) | `0.999x` | `3.214x` | `3.211x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003643` (3.82 MiB) | `937831` (915.85 KiB) | `0.999x` | `4.269x` | `4.265x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003619` (3.82 MiB) | `550034` (537.14 KiB) | `0.999x` | `7.279x` | `7.272x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `5331` (5.21 KiB) | `0.999x` | `750.991x` | `750.328x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `1024` (1.00 KiB) | `86602` (84.57 KiB) | `6236` (6.09 KiB) | `0.012x` | `13.887x` | `0.164x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004443` (7.63 MiB) | `5704` (5.57 KiB) | `0.999x` | `1403.303x` | `1402.525x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `81269` (79.36 KiB) | `3684` (3.60 KiB) | `0.000x` | `22.060x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `3000000` (2.86 MiB) | `86024` (84.01 KiB) | `6924` (6.76 KiB) | `34.874x` | `12.424x` | `433.276x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4236` (4.14 KiB) | `0.999x` | `945.120x` | `944.287x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `58030` (56.67 KiB) | `301258` (294.20 KiB) | `33746` (32.96 KiB) | `0.193x` | `8.927x` | `1.720x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `22051` (21.53 KiB) | `197351` (192.73 KiB) | `28153` (27.49 KiB) | `0.112x` | `7.010x` | `0.783x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `25445` (24.85 KiB) | `208868` (203.97 KiB) | `33554` (32.77 KiB) | `0.122x` | `6.225x` | `0.758x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `48191` (47.06 KiB) | `231838` (226.40 KiB) | `23281` (22.74 KiB) | `0.208x` | `9.958x` | `2.070x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `49433` (48.27 KiB) | `267615` (261.34 KiB) | `40094` (39.15 KiB) | `0.185x` | `6.675x` | `1.233x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `16873` (16.48 KiB) | `213106` (208.11 KiB) | `28030` (27.37 KiB) | `0.079x` | `7.603x` | `0.602x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `91870` (89.72 KiB) | `327554` (319.88 KiB) | `53982` (52.72 KiB) | `0.280x` | `6.068x` | `1.702x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `13001` (12.70 KiB) | `149829` (146.32 KiB) | `20974` (20.48 KiB) | `0.087x` | `7.144x` | `0.620x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `28101` (27.44 KiB) | `189788` (185.34 KiB) | `23133` (22.59 KiB) | `0.148x` | `8.204x` | `1.215x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `45607` (44.54 KiB) | `260361` (254.26 KiB) | `38224` (37.33 KiB) | `0.175x` | `6.811x` | `1.193x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003559` (3.82 MiB) | `21162` (20.67 KiB) | `0.999x` | `189.186x` | `189.018x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004556` (7.63 MiB) | `2842469` (2.71 MiB) | `0.999x` | `2.816x` | `2.814x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004556` (7.63 MiB) | `3580471` (3.41 MiB) | `0.999x` | `2.236x` | `2.234x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003546` (3.82 MiB) | `5659` (5.53 KiB) | `0.999x` | `707.465x` | `706.839x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00000.parquet`: `35785` rows, `3017646` file bytes (2.88 MiB), `27793026` physical bytes (26.51 MiB), `17739184` encoded bytes (16.92 MiB), `2985531` compressed data bytes (2.85 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00001.parquet`: `35363` rows, `2925544` file bytes (2.79 MiB), `27158090` physical bytes (25.90 MiB), `17285758` encoded bytes (16.48 MiB), `2892998` compressed data bytes (2.76 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00002.parquet`: `35731` rows, `3034848` file bytes (2.89 MiB), `27987501` physical bytes (26.69 MiB), `17776245` encoded bytes (16.95 MiB), `3001493` compressed data bytes (2.86 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00003.parquet`: `36145` rows, `2996306` file bytes (2.86 MiB), `27990190` physical bytes (26.69 MiB), `17739166` encoded bytes (16.92 MiB), `2963821` compressed data bytes (2.83 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00004.parquet`: `35510` rows, `2945339` file bytes (2.81 MiB), `27624614` physical bytes (26.34 MiB), `17457507` encoded bytes (16.65 MiB), `2912096` compressed data bytes (2.78 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00005.parquet`: `36194` rows, `2982906` file bytes (2.84 MiB), `27771666` physical bytes (26.49 MiB), `17725198` encoded bytes (16.90 MiB), `2950324` compressed data bytes (2.81 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00006.parquet`: `36172` rows, `2999083` file bytes (2.86 MiB), `27960059` physical bytes (26.66 MiB), `17691409` encoded bytes (16.87 MiB), `2966622` compressed data bytes (2.83 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00007.parquet`: `36147` rows, `3008414` file bytes (2.87 MiB), `27970789` physical bytes (26.68 MiB), `17798873` encoded bytes (16.97 MiB), `2975486` compressed data bytes (2.84 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00008.parquet`: `36155` rows, `2951683` file bytes (2.81 MiB), `28024913` physical bytes (26.73 MiB), `17742307` encoded bytes (16.92 MiB), `2918905` compressed data bytes (2.78 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00009.parquet`: `35746` rows, `2935014` file bytes (2.80 MiB), `27595449` physical bytes (26.32 MiB), `17567273` encoded bytes (16.75 MiB), `2902776` compressed data bytes (2.77 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00010.parquet`: `36388` rows, `2953201` file bytes (2.82 MiB), `27974353` physical bytes (26.68 MiB), `17764621` encoded bytes (16.94 MiB), `2920645` compressed data bytes (2.79 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00011.parquet`: `36320` rows, `2954881` file bytes (2.82 MiB), `28014999` physical bytes (26.72 MiB), `17778243` encoded bytes (16.95 MiB), `2922305` compressed data bytes (2.79 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00012.parquet`: `36213` rows, `2957972` file bytes (2.82 MiB), `27788519` physical bytes (26.50 MiB), `17716317` encoded bytes (16.90 MiB), `2925769` compressed data bytes (2.79 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00013.parquet`: `35753` rows, `2944313` file bytes (2.81 MiB), `27611358` physical bytes (26.33 MiB), `17518073` encoded bytes (16.71 MiB), `2911674` compressed data bytes (2.78 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00014.parquet`: `35431` rows, `3397277` file bytes (3.24 MiB), `24657401` physical bytes (23.52 MiB), `17873387` encoded bytes (17.05 MiB), `3366279` compressed data bytes (3.21 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00015.parquet`: `34981` rows, `3432781` file bytes (3.27 MiB), `24372438` physical bytes (23.24 MiB), `17796124` encoded bytes (16.97 MiB), `3402466` compressed data bytes (3.24 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00016.parquet`: `34717` rows, `3564398` file bytes (3.40 MiB), `23107155` physical bytes (22.04 MiB), `18533039` encoded bytes (17.67 MiB), `3534030` compressed data bytes (3.37 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00017.parquet`: `34347` rows, `3601036` file bytes (3.43 MiB), `22106380` physical bytes (21.08 MiB), `18968405` encoded bytes (18.09 MiB), `3570222` compressed data bytes (3.40 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00018.parquet`: `33890` rows, `3725445` file bytes (3.55 MiB), `22233443` physical bytes (21.20 MiB), `19028548` encoded bytes (18.15 MiB), `3695000` compressed data bytes (3.52 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00019.parquet`: `34459` rows, `3514678` file bytes (3.35 MiB), `21720929` physical bytes (20.71 MiB), `18733418` encoded bytes (17.87 MiB), `3483877` compressed data bytes (3.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00020.parquet`: `33754` rows, `3650566` file bytes (3.48 MiB), `21839022` physical bytes (20.83 MiB), `18728155` encoded bytes (17.86 MiB), `3619869` compressed data bytes (3.45 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00021.parquet`: `34401` rows, `3614030` file bytes (3.45 MiB), `22039516` physical bytes (21.02 MiB), `18978502` encoded bytes (18.10 MiB), `3583102` compressed data bytes (3.42 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00022.parquet`: `33998` rows, `3545844` file bytes (3.38 MiB), `21732565` physical bytes (20.73 MiB), `18577877` encoded bytes (17.72 MiB), `3515450` compressed data bytes (3.35 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00023.parquet`: `34020` rows, `3642947` file bytes (3.47 MiB), `22062062` physical bytes (21.04 MiB), `19006846` encoded bytes (18.13 MiB), `3612143` compressed data bytes (3.44 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00024.parquet`: `34199` rows, `3586573` file bytes (3.42 MiB), `21838812` physical bytes (20.83 MiB), `18715999` encoded bytes (17.85 MiB), `3556253` compressed data bytes (3.39 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00025.parquet`: `34286` rows, `3591303` file bytes (3.42 MiB), `21879786` physical bytes (20.87 MiB), `18769142` encoded bytes (17.90 MiB), `3560672` compressed data bytes (3.40 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00026.parquet`: `34457` rows, `3536875` file bytes (3.37 MiB), `21893956` physical bytes (20.88 MiB), `18682819` encoded bytes (17.82 MiB), `3506237` compressed data bytes (3.34 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00027.parquet`: `34235` rows, `3543692` file bytes (3.38 MiB), `21738211` physical bytes (20.73 MiB), `18679992` encoded bytes (17.81 MiB), `3513399` compressed data bytes (3.35 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict/part-00028.parquet`: `15203` rows, `1614192` file bytes (1.54 MiB), `9911422` physical bytes (9.45 MiB), `8445073` encoded bytes (8.05 MiB), `1597411` compressed data bytes (1.52 MiB)
