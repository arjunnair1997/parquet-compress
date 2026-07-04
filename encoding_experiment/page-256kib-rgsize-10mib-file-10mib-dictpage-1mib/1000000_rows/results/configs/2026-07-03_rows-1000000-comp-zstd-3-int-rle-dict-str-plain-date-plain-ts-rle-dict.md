# ClickBench Parquet Experiment

- Started: `2026-07-03T15:33:03-04:00`
- Write elapsed: `16.292s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `548641980` (523.23 MiB)
- Compressed column data bytes after codec compression: `98099739` (93.56 MiB)
- Parquet file bytes: `99052011` (94.46 MiB)
- Physical/encoded ratio: `1.298x`
- Encoded/compressed-data ratio: `5.593x`
- Physical/compressed-data ratio: `7.262x`
- Physical/parquet-file ratio: `7.192x`
- Files: `29`

## Settings

- Compression: `zstd-3`
- Int encoding: `rle-dict`
- String encoding: `plain`
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
- Elapsed: `7.253s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `9844913` (9.39 MiB) | `9846281` (9.39 MiB) | `0.813x` | `1.000x` | `0.812x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `79769` (77.90 KiB) | `50814` (49.62 KiB) | `50.145x` | `1.570x` | `78.718x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:524` | `1000000` | `138409995` (132.00 MiB) | `142872536` (136.25 MiB) | `13933257` (13.29 MiB) | `0.969x` | `10.254x` | `9.934x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4953` (4.84 KiB) | `5979` (5.84 KiB) | `807.591x` | `0.828x` | `669.008x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `7358049` (7.02 MiB) | `4038947` (3.85 MiB) | `1.087x` | `1.822x` | `1.981x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003526` (3.82 MiB) | `4913` (4.80 KiB) | `0.999x` | `814.884x` | `814.166x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `5005` (4.89 KiB) | `6031` (5.89 KiB) | `799.201x` | `0.830x` | `663.240x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `943758` (921.64 KiB) | `808765` (789.81 KiB) | `4.238x` | `1.167x` | `4.946x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `435261` (425.06 KiB) | `244673` (238.94 KiB) | `9.190x` | `1.779x` | `16.348x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `1230940` (1.17 MiB) | `1101698` (1.05 MiB) | `6.499x` | `1.117x` | `7.262x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4953` (4.84 KiB) | `5979` (5.84 KiB) | `807.591x` | `0.828x` | `669.008x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `230073` (224.68 KiB) | `120553` (117.73 KiB) | `17.386x` | `1.908x` | `33.180x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `228598` (223.24 KiB) | `127212` (124.23 KiB) | `17.498x` | `1.797x` | `31.444x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:350` | `1000000` | `88562192` (84.46 MiB) | `92649549` (88.36 MiB) | `15301109` (14.59 MiB) | `0.956x` | `6.055x` | `5.788x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:326` | `1000000` | `79583339` (75.90 MiB) | `83647391` (79.77 MiB) | `14212850` (13.55 MiB) | `0.951x` | `5.885x` | `5.599x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `92548` (90.38 KiB) | `83596` (81.64 KiB) | `43.221x` | `1.107x` | `47.849x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `519299` (507.13 KiB) | `218942` (213.81 KiB) | `7.703x` | `2.372x` | `18.270x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `381080` (372.15 KiB) | `164015` (160.17 KiB) | `10.496x` | `2.323x` | `24.388x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `161667` (157.88 KiB) | `71391` (69.72 KiB) | `24.742x` | `2.265x` | `56.029x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `74764` (73.01 KiB) | `40209` (39.27 KiB) | `53.502x` | `1.859x` | `99.480x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `368617` (359.98 KiB) | `204892` (200.09 KiB) | `10.851x` | `1.799x` | `19.522x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `372543` (363.81 KiB) | `206453` (201.61 KiB) | `10.737x` | `1.804x` | `19.375x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `111260` (108.65 KiB) | `62094` (60.64 KiB) | `35.952x` | `1.792x` | `64.418x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `88009` (85.95 KiB) | `49214` (48.06 KiB) | `45.450x` | `1.788x` | `81.278x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `214805` (209.77 KiB) | `113651` (110.99 KiB) | `18.622x` | `1.890x` | `35.195x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3354477` (3.20 MiB) | `7357830` (7.02 MiB) | `246055` (240.29 KiB) | `0.456x` | `29.903x` | `13.633x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `36608` (35.75 KiB) | `25685` (25.08 KiB) | `109.266x` | `1.425x` | `155.733x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `33244` (32.46 KiB) | `23402` (22.85 KiB) | `120.322x` | `1.421x` | `170.926x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `276271` (269.80 KiB) | `178949` (174.75 KiB) | `14.479x` | `1.544x` | `22.353x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3767530` (3.59 MiB) | `7770891` (7.41 MiB) | `136957` (133.75 KiB) | `0.485x` | `56.740x` | `27.509x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `6058` (5.92 KiB) | `7086` (6.92 KiB) | `660.284x` | `0.855x` | `564.493x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `6398` (6.25 KiB) | `7424` (7.25 KiB) | `625.195x` | `0.862x` | `538.793x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `29843` (29.14 KiB) | `24526` (23.95 KiB) | `134.035x` | `1.217x` | `163.092x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `34411` (33.60 KiB) | `23615` (23.06 KiB) | `116.242x` | `1.457x` | `169.384x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `81583` (79.67 KiB) | `4084892` (3.90 MiB) | `22655` (22.12 KiB) | `0.020x` | `180.309x` | `3.601x` | `81583` (79.67 KiB) |
| `Params` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002166` (3.82 MiB) | `2869` (2.80 KiB) | `0.000x` | `1394.969x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `749925` (732.35 KiB) | `622101` (607.52 KiB) | `5.334x` | `1.205x` | `6.430x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `289466` (282.68 KiB) | `178689` (174.50 KiB) | `13.819x` | `1.620x` | `22.385x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `167711` (163.78 KiB) | `76289` (74.50 KiB) | `23.851x` | `2.198x` | `52.432x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3528017` (3.36 MiB) | `7535085` (7.19 MiB) | `720780` (703.89 KiB) | `0.468x` | `10.454x` | `4.895x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `34920` (34.10 KiB) | `22970` (22.43 KiB) | `114.548x` | `1.520x` | `174.140x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `80996` (79.10 KiB) | `82055` (80.13 KiB) | `49.385x` | `0.987x` | `48.748x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `695835` (679.53 KiB) | `474188` (463.07 KiB) | `5.748x` | `1.467x` | `8.435x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `749506` (731.94 KiB) | `585576` (571.85 KiB) | `5.337x` | `1.280x` | `6.831x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `164845` (160.98 KiB) | `85345` (83.34 KiB) | `24.265x` | `1.932x` | `46.869x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `7208499` (6.87 MiB) | `3976302` (3.79 MiB) | `1.110x` | `1.813x` | `2.012x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `155946` (152.29 KiB) | `83233` (81.28 KiB) | `25.650x` | `1.874x` | `48.058x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `89021` (86.93 KiB) | `56113` (54.80 KiB) | `44.933x` | `1.586x` | `71.285x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `231541` (226.11 KiB) | `124239` (121.33 KiB) | `17.276x` | `1.864x` | `32.196x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `5061` (4.94 KiB) | `6087` (5.94 KiB) | `790.358x` | `0.831x` | `657.138x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:85` | `1000000` | `13587860` (12.96 MiB) | `17594776` (16.78 MiB) | `14603` (14.26 KiB) | `0.772x` | `1204.874x` | `930.484x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `6944` (6.78 KiB) | `7970` (7.78 KiB) | `576.037x` | `0.871x` | `501.882x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `54174` (52.90 KiB) | `36527` (35.67 KiB) | `73.836x` | `1.483x` | `109.508x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `8306` (8.11 KiB) | `9274` (9.06 KiB) | `481.580x` | `0.896x` | `431.313x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `16291` (15.91 KiB) | `17322` (16.92 KiB) | `245.534x` | `0.940x` | `230.920x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `1186873` (1.13 MiB) | `1059313` (1.01 MiB) | `6.740x` | `1.120x` | `7.552x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:137` | `1000000` | `27797671` (26.51 MiB) | `31858095` (30.38 MiB) | `5317392` (5.07 MiB) | `0.873x` | `5.991x` | `5.228x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4509921` (4.30 MiB) | `4511083` (4.30 MiB) | `0.887x` | `1.000x` | `0.887x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4953` (4.84 KiB) | `5979` (5.84 KiB) | `807.591x` | `0.828x` | `669.008x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4953` (4.84 KiB) | `5979` (5.84 KiB) | `807.591x` | `0.828x` | `669.008x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4953` (4.84 KiB) | `5979` (5.84 KiB) | `807.591x` | `0.828x` | `669.008x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `65853` (64.31 KiB) | `50036` (48.86 KiB) | `60.741x` | `1.316x` | `79.942x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4953` (4.84 KiB) | `5979` (5.84 KiB) | `807.591x` | `0.828x` | `669.008x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `1000000` (976.56 KiB) | `5002903` (4.77 MiB) | `29498` (28.81 KiB) | `0.200x` | `169.601x` | `33.901x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `7353460` (7.01 MiB) | `4039977` (3.85 MiB) | `1.088x` | `1.820x` | `1.980x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `186099` (181.74 KiB) | `118015` (115.25 KiB) | `21.494x` | `1.577x` | `33.894x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `133484` (130.36 KiB) | `80425` (78.54 KiB) | `29.966x` | `1.660x` | `49.736x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `142392` (139.05 KiB) | `88791` (86.71 KiB) | `28.091x` | `1.604x` | `45.050x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `494029` (482.45 KiB) | `316317` (308.90 KiB) | `8.097x` | `1.562x` | `12.646x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `415620` (405.88 KiB) | `241022` (235.37 KiB) | `9.624x` | `1.724x` | `16.596x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `931810` (909.97 KiB) | `701088` (684.66 KiB) | `4.293x` | `1.329x` | `5.705x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `173332` (169.27 KiB) | `151404` (147.86 KiB) | `23.077x` | `1.145x` | `26.419x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4953` (4.84 KiB) | `5979` (5.84 KiB) | `807.591x` | `0.828x` | `669.008x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `71570` (69.89 KiB) | `52162` (50.94 KiB) | `55.889x` | `1.372x` | `76.684x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `2001192` (1.91 MiB) | `6004435` (5.73 MiB) | `32302` (31.54 KiB) | `0.333x` | `185.884x` | `61.953x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3325142` (3.17 MiB) | `7328503` (6.99 MiB) | `122281` (119.42 KiB) | `0.454x` | `59.932x` | `27.193x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002166` (3.82 MiB) | `2869` (2.80 KiB) | `0.000x` | `1394.969x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002166` (3.82 MiB) | `2869` (2.80 KiB) | `0.000x` | `1394.969x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4953` (4.84 KiB) | `5979` (5.84 KiB) | `807.591x` | `0.828x` | `669.008x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `122082` (119.22 KiB) | `76429` (74.64 KiB) | `32.765x` | `1.597x` | `52.336x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `343596` (335.54 KiB) | `156920` (153.24 KiB) | `11.642x` | `2.190x` | `25.491x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `704591` (688.08 KiB) | `378341` (369.47 KiB) | `5.677x` | `1.862x` | `10.572x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `1796013` (1.71 MiB) | `1571380` (1.50 MiB) | `2.227x` | `1.143x` | `2.546x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `1359205` (1.30 MiB) | `1053248` (1.00 MiB) | `2.943x` | `1.290x` | `3.798x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `1134136` (1.08 MiB) | `692628` (676.39 KiB) | `3.527x` | `1.637x` | `5.775x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `5902` (5.76 KiB) | `6928` (6.77 KiB) | `677.736x` | `0.852x` | `577.367x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `1024` (1.00 KiB) | `4004029` (3.82 MiB) | `4890` (4.78 KiB) | `0.000x` | `818.820x` | `0.209x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `6093` (5.95 KiB) | `7119` (6.95 KiB) | `1312.982x` | `0.856x` | `1123.753x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002166` (3.82 MiB) | `2869` (2.80 KiB) | `0.000x` | `1394.969x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3000000` (2.86 MiB) | `7003299` (6.68 MiB) | `5259` (5.14 KiB) | `0.428x` | `1331.679x` | `570.451x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4953` (4.84 KiB) | `5979` (5.84 KiB) | `807.591x` | `0.828x` | `669.008x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `58030` (56.67 KiB) | `4062529` (3.87 MiB) | `19298` (18.85 KiB) | `0.014x` | `210.516x` | `3.007x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `22051` (21.53 KiB) | `4025591` (3.84 MiB) | `16558` (16.17 KiB) | `0.005x` | `243.121x` | `1.332x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `25445` (24.85 KiB) | `4029080` (3.84 MiB) | `19054` (18.61 KiB) | `0.006x` | `211.456x` | `1.335x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `48191` (47.06 KiB) | `4051764` (3.86 MiB) | `13315` (13.00 KiB) | `0.012x` | `304.301x` | `3.619x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `49433` (48.27 KiB) | `4053243` (3.87 MiB) | `21125` (20.63 KiB) | `0.012x` | `191.869x` | `2.340x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `16873` (16.48 KiB) | `4019956` (3.83 MiB) | `17078` (16.68 KiB) | `0.004x` | `235.388x` | `0.988x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `91870` (89.72 KiB) | `4097549` (3.91 MiB) | `29924` (29.22 KiB) | `0.022x` | `136.932x` | `3.070x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `13001` (12.70 KiB) | `4016725` (3.83 MiB) | `14194` (13.86 KiB) | `0.003x` | `282.988x` | `0.916x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `28101` (27.44 KiB) | `4032786` (3.85 MiB) | `15753` (15.38 KiB) | `0.007x` | `256.001x` | `1.784x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `45607` (44.54 KiB) | `4048299` (3.86 MiB) | `28452` (27.79 KiB) | `0.011x` | `142.285x` | `1.603x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `21366` (20.87 KiB) | `16226` (15.85 KiB) | `187.213x` | `1.317x` | `246.518x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `3778928` (3.60 MiB) | `3524106` (3.36 MiB) | `2.117x` | `1.072x` | `2.270x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `4618457` (4.40 MiB) | `4540631` (4.33 MiB) | `1.732x` | `1.017x` | `1.762x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `5887` (5.75 KiB) | `6913` (6.75 KiB) | `679.463x` | `0.852x` | `578.620x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00000.parquet`: `35922` rows, `3239536` file bytes (3.09 MiB), `27897645` physical bytes (26.61 MiB), `21706807` encoded bytes (20.70 MiB), `3205679` compressed data bytes (3.06 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00001.parquet`: `35729` rows, `3154870` file bytes (3.01 MiB), `27446647` physical bytes (26.18 MiB), `21274472` encoded bytes (20.29 MiB), `3120520` compressed data bytes (2.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00002.parquet`: `35874` rows, `3237524` file bytes (3.09 MiB), `28097951` physical bytes (26.80 MiB), `21918275` encoded bytes (20.90 MiB), `3202428` compressed data bytes (3.05 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00003.parquet`: `36315` rows, `3196545` file bytes (3.05 MiB), `28113695` physical bytes (26.81 MiB), `21830822` encoded bytes (20.82 MiB), `3162321` compressed data bytes (3.02 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00004.parquet`: `35677` rows, `3147023` file bytes (3.00 MiB), `27744225` physical bytes (26.46 MiB), `21581340` encoded bytes (20.58 MiB), `3112037` compressed data bytes (2.97 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00005.parquet`: `35870` rows, `3127963` file bytes (2.98 MiB), `27493329` physical bytes (26.22 MiB), `21284887` encoded bytes (20.30 MiB), `3093736` compressed data bytes (2.95 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00006.parquet`: `36284` rows, `3210553` file bytes (3.06 MiB), `28097364` physical bytes (26.80 MiB), `21842578` encoded bytes (20.83 MiB), `3176037` compressed data bytes (3.03 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00007.parquet`: `35778` rows, `3157970` file bytes (3.01 MiB), `27689781` physical bytes (26.41 MiB), `21506075` encoded bytes (20.51 MiB), `3123381` compressed data bytes (2.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00008.parquet`: `36295` rows, `3172881` file bytes (3.03 MiB), `28133231` physical bytes (26.83 MiB), `21840335` encoded bytes (20.83 MiB), `3138366` compressed data bytes (2.99 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00009.parquet`: `35931` rows, `3175287` file bytes (3.03 MiB), `27712192` physical bytes (26.43 MiB), `21501408` encoded bytes (20.51 MiB), `3141272` compressed data bytes (3.00 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00010.parquet`: `36558` rows, `3169542` file bytes (3.02 MiB), `28112230` physical bytes (26.81 MiB), `21757061` encoded bytes (20.75 MiB), `3135146` compressed data bytes (2.99 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00011.parquet`: `36499` rows, `3157788` file bytes (3.01 MiB), `28149775` physical bytes (26.85 MiB), `21811954` encoded bytes (20.80 MiB), `3123491` compressed data bytes (2.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00012.parquet`: `36659` rows, `3206698` file bytes (3.06 MiB), `28083875` physical bytes (26.78 MiB), `21740443` encoded bytes (20.73 MiB), `3172709` compressed data bytes (3.03 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00013.parquet`: `36291` rows, `3203114` file bytes (3.05 MiB), `28095695` physical bytes (26.79 MiB), `21810224` encoded bytes (20.80 MiB), `3168654` compressed data bytes (3.02 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00014.parquet`: `34762` rows, `3597440` file bytes (3.43 MiB), `24034368` physical bytes (22.92 MiB), `18687603` encoded bytes (17.82 MiB), `3564787` compressed data bytes (3.40 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00015.parquet`: `35114` rows, `3739162` file bytes (3.57 MiB), `24500421` physical bytes (23.37 MiB), `19133347` encoded bytes (18.25 MiB), `3707017` compressed data bytes (3.54 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00016.parquet`: `34300` rows, `3742814` file bytes (3.57 MiB), `22686098` physical bytes (21.64 MiB), `17391301` encoded bytes (16.59 MiB), `3710696` compressed data bytes (3.54 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00017.parquet`: `34224` rows, `3800843` file bytes (3.62 MiB), `22058432` physical bytes (21.04 MiB), `16762858` encoded bytes (15.99 MiB), `3768266` compressed data bytes (3.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00018.parquet`: `33581` rows, `3892645` file bytes (3.71 MiB), `22007311` physical bytes (20.99 MiB), `16806516` encoded bytes (16.03 MiB), `3860444` compressed data bytes (3.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00019.parquet`: `34080` rows, `3690775` file bytes (3.52 MiB), `21501821` physical bytes (20.51 MiB), `16225927` encoded bytes (15.47 MiB), `3658247` compressed data bytes (3.49 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00020.parquet`: `33661` rows, `3846480` file bytes (3.67 MiB), `21776475` physical bytes (20.77 MiB), `16589882` encoded bytes (15.82 MiB), `3814014` compressed data bytes (3.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00021.parquet`: `34324` rows, `3810668` file bytes (3.63 MiB), `21987988` physical bytes (20.97 MiB), `16688967` encoded bytes (15.92 MiB), `3777959` compressed data bytes (3.60 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00022.parquet`: `33909` rows, `3744623` file bytes (3.57 MiB), `21677899` physical bytes (20.67 MiB), `16431142` encoded bytes (15.67 MiB), `3712406` compressed data bytes (3.54 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00023.parquet`: `33965` rows, `3837144` file bytes (3.66 MiB), `22025827` physical bytes (21.01 MiB), `16771769` encoded bytes (15.99 MiB), `3804579` compressed data bytes (3.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00024.parquet`: `34109` rows, `3789316` file bytes (3.61 MiB), `21782930` physical bytes (20.77 MiB), `16523724` encoded bytes (15.76 MiB), `3757218` compressed data bytes (3.58 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00025.parquet`: `34186` rows, `3792735` file bytes (3.62 MiB), `21828274` physical bytes (20.82 MiB), `16560157` encoded bytes (15.79 MiB), `3760341` compressed data bytes (3.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00026.parquet`: `34418` rows, `3729380` file bytes (3.56 MiB), `21857802` physical bytes (20.85 MiB), `16540163` encoded bytes (15.77 MiB), `3696977` compressed data bytes (3.53 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00027.parquet`: `34149` rows, `3743346` file bytes (3.57 MiB), `21686280` physical bytes (20.68 MiB), `16426819` encoded bytes (15.67 MiB), `3711289` compressed data bytes (3.54 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict/part-00028.parquet`: `15536` rows, `1737346` file bytes (1.66 MiB), `10119063` physical bytes (9.65 MiB), `7695124` encoded bytes (7.34 MiB), `1719722` compressed data bytes (1.64 MiB)
