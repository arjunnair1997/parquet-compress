# ClickBench Parquet Experiment

- Started: `2026-07-03T23:40:21-04:00`
- Write elapsed: `12.444s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `586845134` (559.66 MiB)
- Compressed column data bytes after codec compression: `101538503` (96.83 MiB)
- Parquet file bytes: `102451396` (97.71 MiB)
- Physical/encoded ratio: `1.214x`
- Encoded/compressed-data ratio: `5.780x`
- Physical/compressed-data ratio: `7.016x`
- Physical/parquet-file ratio: `6.954x`
- Files: `29`

## Settings

- Compression: `zstd-3`
- Int encoding: `delta-binary-packed`
- String encoding: `plain`
- Date encoding: `rle-dict`
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
- Files read: `29`
- Elapsed: `7.469s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `8000000` (7.63 MiB) | `7989725` (7.62 MiB) | `7990506` (7.62 MiB) | `1.001x` | `1.000x` | `1.001x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `197896` (193.26 KiB) | `68025` (66.43 KiB) | `20.213x` | `2.909x` | `58.802x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:512` | `1000000` | `138409995` (132.00 MiB) | `142873218` (136.25 MiB) | `13946791` (13.30 MiB) | `0.969x` | `10.244x` | `9.924x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `43674` (42.65 KiB) | `5710` (5.58 KiB) | `91.588x` | `7.649x` | `700.525x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:58, DICTIONARY_PAGE/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `7352575` (7.01 MiB) | `4030847` (3.84 MiB) | `1.088x` | `1.824x` | `1.985x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:58, DICTIONARY_PAGE/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `5040` (4.92 KiB) | `6084` (5.94 KiB) | `793.651x` | `0.828x` | `657.462x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `46468` (45.38 KiB) | `5963` (5.82 KiB) | `86.081x` | `7.793x` | `670.803x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `3747719` (3.57 MiB) | `859710` (839.56 KiB) | `1.067x` | `4.359x` | `4.653x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1484741` (1.42 MiB) | `464447` (453.56 KiB) | `2.694x` | `3.197x` | `8.612x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `8000000` (7.63 MiB) | `4390888` (4.19 MiB) | `760460` (742.64 KiB) | `1.822x` | `5.774x` | `10.520x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `43122` (42.11 KiB) | `5009` (4.89 KiB) | `92.760x` | `8.609x` | `798.563x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `843341` (823.58 KiB) | `216248` (211.18 KiB) | `4.743x` | `3.900x` | `18.497x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `687026` (670.92 KiB) | `236074` (230.54 KiB) | `5.822x` | `2.910x` | `16.944x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:336` | `1000000` | `88562192` (84.46 MiB) | `92648954` (88.36 MiB) | `15301281` (14.59 MiB) | `0.956x` | `6.055x` | `5.788x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:324` | `1000000` | `79583339` (75.90 MiB) | `83647148` (79.77 MiB) | `14213083` (13.55 MiB) | `0.951x` | `5.885x` | `5.599x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `233999` (228.51 KiB) | `98917` (96.60 KiB) | `17.094x` | `2.366x` | `40.438x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1863398` (1.78 MiB) | `522232` (509.99 KiB) | `2.147x` | `3.568x` | `7.659x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1366749` (1.30 MiB) | `336739` (328.85 KiB) | `2.927x` | `4.059x` | `11.879x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1081286` (1.03 MiB) | `203270` (198.51 KiB) | `3.699x` | `5.319x` | `19.678x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `721391` (704.48 KiB) | `99981` (97.64 KiB) | `5.545x` | `7.215x` | `40.008x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1282199` (1.22 MiB) | `410283` (400.67 KiB) | `3.120x` | `3.125x` | `9.749x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1341007` (1.28 MiB) | `358942` (350.53 KiB) | `2.983x` | `3.736x` | `11.144x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `552682` (539.73 KiB) | `136543` (133.34 KiB) | `7.237x` | `4.048x` | `29.295x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `398798` (389.45 KiB) | `93241` (91.06 KiB) | `10.030x` | `4.277x` | `42.900x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `523711` (511.44 KiB) | `199670` (194.99 KiB) | `7.638x` | `2.623x` | `20.033x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `3354477` (3.20 MiB) | `7357886` (7.02 MiB) | `247683` (241.88 KiB) | `0.456x` | `29.707x` | `13.543x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `159772` (156.03 KiB) | `42520` (41.52 KiB) | `25.036x` | `3.758x` | `94.073x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `197352` (192.73 KiB) | `36867` (36.00 KiB) | `20.268x` | `5.353x` | `108.498x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `661531` (646.03 KiB) | `259692` (253.61 KiB) | `6.047x` | `2.547x` | `15.403x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `3767530` (3.59 MiB) | `7770933` (7.41 MiB) | `137235` (134.02 KiB) | `0.485x` | `56.625x` | `27.453x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `46567` (45.48 KiB) | `7544` (7.37 KiB) | `85.898x` | `6.173x` | `530.223x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `47605` (46.49 KiB) | `8211` (8.02 KiB) | `84.025x` | `5.798x` | `487.151x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `114801` (112.11 KiB) | `33143` (32.37 KiB) | `34.843x` | `3.464x` | `120.689x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `144977` (141.58 KiB) | `39830` (38.90 KiB) | `27.591x` | `3.640x` | `100.427x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `81583` (79.67 KiB) | `4084960` (3.90 MiB) | `22719` (22.19 KiB) | `0.020x` | `179.804x` | `3.591x` | `81583` (79.67 KiB) |
| `Params` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `0` (0 B) | `4002197` (3.82 MiB) | `2893` (2.83 KiB) | `0.000x` | `1383.407x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `2652086` (2.53 MiB) | `782155` (763.82 KiB) | `1.508x` | `3.391x` | `5.114x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `496085` (484.46 KiB) | `238115` (232.53 KiB) | `8.063x` | `2.083x` | `16.799x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `377677` (368.83 KiB) | `150787` (147.25 KiB) | `10.591x` | `2.505x` | `26.527x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `3528017` (3.36 MiB) | `7535081` (7.19 MiB) | `720823` (703.93 KiB) | `0.468x` | `10.453x` | `4.894x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `202417` (197.67 KiB) | `45312` (44.25 KiB) | `19.761x` | `4.467x` | `88.277x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `196781` (192.17 KiB) | `84568` (82.59 KiB) | `20.327x` | `2.327x` | `47.299x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1358854` (1.30 MiB) | `575203` (561.72 KiB) | `2.944x` | `2.362x` | `6.954x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1249910` (1.19 MiB) | `552858` (539.90 KiB) | `3.200x` | `2.261x` | `7.235x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `975678` (952.81 KiB) | `212407` (207.43 KiB) | `4.100x` | `4.593x` | `18.832x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:58, DICTIONARY_PAGE/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `7216092` (6.88 MiB) | `3979746` (3.80 MiB) | `1.109x` | `1.813x` | `2.010x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `436556` (426.32 KiB) | `119334` (116.54 KiB) | `9.163x` | `3.658x` | `33.519x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `207482` (202.62 KiB) | `73471` (71.75 KiB) | `19.279x` | `2.824x` | `54.443x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1857414` (1.77 MiB) | `317700` (310.25 KiB) | `2.154x` | `5.846x` | `12.590x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `43609` (42.59 KiB) | `5252` (5.13 KiB) | `91.724x` | `8.303x` | `761.615x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:86` | `1000000` | `13587860` (12.96 MiB) | `17594920` (16.78 MiB) | `14731` (14.39 KiB) | `0.772x` | `1194.415x` | `922.399x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `95854` (93.61 KiB) | `12696` (12.40 KiB) | `41.730x` | `7.550x` | `315.060x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `140688` (137.39 KiB) | `44570` (43.53 KiB) | `28.432x` | `3.157x` | `89.746x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `53695` (52.44 KiB) | `10132` (9.89 KiB) | `74.495x` | `5.300x` | `394.789x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `63990` (62.49 KiB) | `16551` (16.16 KiB) | `62.510x` | `3.866x` | `241.677x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `8000000` (7.63 MiB) | `7451875` (7.11 MiB) | `1417172` (1.35 MiB) | `1.074x` | `5.258x` | `5.645x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:138` | `1000000` | `27797671` (26.51 MiB) | `31856006` (30.38 MiB) | `5328684` (5.08 MiB) | `0.873x` | `5.978x` | `5.217x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `3953077` (3.77 MiB) | `3952453` (3.77 MiB) | `1.012x` | `1.000x` | `1.012x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `43122` (42.11 KiB) | `5009` (4.89 KiB) | `92.760x` | `8.609x` | `798.563x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `43122` (42.11 KiB) | `5009` (4.89 KiB) | `92.760x` | `8.609x` | `798.563x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `43122` (42.11 KiB) | `5009` (4.89 KiB) | `92.760x` | `8.609x` | `798.563x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `168482` (164.53 KiB) | `60659` (59.24 KiB) | `23.741x` | `2.778x` | `65.942x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `43122` (42.11 KiB) | `5009` (4.89 KiB) | `92.760x` | `8.609x` | `798.563x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `1000000` (976.56 KiB) | `5002950` (4.77 MiB) | `28474` (27.81 KiB) | `0.200x` | `175.702x` | `35.120x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:58, DICTIONARY_PAGE/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `7349435` (7.01 MiB) | `4033241` (3.85 MiB) | `1.089x` | `1.822x` | `1.984x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `743117` (725.70 KiB) | `247180` (241.39 KiB) | `5.383x` | `3.006x` | `16.183x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `253582` (247.64 KiB) | `109113` (106.56 KiB) | `15.774x` | `2.324x` | `36.659x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `328430` (320.73 KiB) | `140318` (137.03 KiB) | `12.179x` | `2.341x` | `28.507x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1590350` (1.52 MiB) | `492249` (480.71 KiB) | `2.515x` | `3.231x` | `8.126x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `841533` (821.81 KiB) | `337731` (329.82 KiB) | `4.753x` | `2.492x` | `11.844x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `3797728` (3.62 MiB) | `927682` (905.94 KiB) | `1.053x` | `4.094x` | `4.312x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `274616` (268.18 KiB) | `121698` (118.85 KiB) | `14.566x` | `2.257x` | `32.868x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `43124` (42.11 KiB) | `5069` (4.95 KiB) | `92.756x` | `8.507x` | `789.110x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `116172` (113.45 KiB) | `67291` (65.71 KiB) | `34.432x` | `1.726x` | `59.443x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `2001192` (1.91 MiB) | `6004491` (5.73 MiB) | `32518` (31.76 KiB) | `0.333x` | `184.651x` | `61.541x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `3325142` (3.17 MiB) | `7328559` (6.99 MiB) | `122319` (119.45 KiB) | `0.454x` | `59.913x` | `27.184x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `0` (0 B) | `4002197` (3.82 MiB) | `2893` (2.83 KiB) | `0.000x` | `1383.407x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `0` (0 B) | `4002197` (3.82 MiB) | `2893` (2.83 KiB) | `0.000x` | `1383.407x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `43122` (42.11 KiB) | `5009` (4.89 KiB) | `92.760x` | `8.609x` | `798.563x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `173668` (169.60 KiB) | `108284` (105.75 KiB) | `23.032x` | `1.604x` | `36.940x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `808462` (789.51 KiB) | `302357` (295.27 KiB) | `4.948x` | `2.674x` | `13.229x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1092343` (1.04 MiB) | `581754` (568.12 KiB) | `3.662x` | `1.878x` | `6.876x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1703507` (1.62 MiB) | `1514914` (1.44 MiB) | `2.348x` | `1.124x` | `2.640x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1502343` (1.43 MiB) | `1303985` (1.24 MiB) | `2.663x` | `1.152x` | `3.068x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `1547985` (1.48 MiB) | `964776` (942.16 KiB) | `2.584x` | `1.605x` | `4.146x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `47369` (46.26 KiB) | `7090` (6.92 KiB) | `84.443x` | `6.681x` | `564.175x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `1024` (1.00 KiB) | `4003989` (3.82 MiB) | `4854` (4.74 KiB) | `0.000x` | `824.884x` | `0.211x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `8000000` (7.63 MiB) | `44050` (43.02 KiB) | `5937` (5.80 KiB) | `181.612x` | `7.420x` | `1347.482x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `0` (0 B) | `4002197` (3.82 MiB) | `2893` (2.83 KiB) | `0.000x` | `1383.407x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `3000000` (2.86 MiB) | `7003356` (6.68 MiB) | `5355` (5.23 KiB) | `0.428x` | `1307.816x` | `560.224x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `43122` (42.11 KiB) | `5009` (4.89 KiB) | `92.760x` | `8.609x` | `798.563x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `58030` (56.67 KiB) | `4062605` (3.87 MiB) | `19455` (19.00 KiB) | `0.014x` | `208.821x` | `2.983x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `22051` (21.53 KiB) | `4025686` (3.84 MiB) | `16700` (16.31 KiB) | `0.005x` | `241.059x` | `1.320x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `25445` (24.85 KiB) | `4029095` (3.84 MiB) | `19052` (18.61 KiB) | `0.006x` | `211.479x` | `1.336x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `48191` (47.06 KiB) | `4051797` (3.86 MiB) | `13389` (13.08 KiB) | `0.012x` | `302.621x` | `3.599x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `49433` (48.27 KiB) | `4053268` (3.87 MiB) | `21151` (20.66 KiB) | `0.012x` | `191.635x` | `2.337x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `16873` (16.48 KiB) | `4019990` (3.83 MiB) | `17211` (16.81 KiB) | `0.004x` | `233.571x` | `0.980x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `91870` (89.72 KiB) | `4097617` (3.91 MiB) | `29886` (29.19 KiB) | `0.022x` | `137.108x` | `3.074x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `13001` (12.70 KiB) | `4016652` (3.83 MiB) | `14070` (13.74 KiB) | `0.003x` | `285.476x` | `0.924x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `28101` (27.44 KiB) | `4032972` (3.85 MiB) | `16017` (15.64 KiB) | `0.007x` | `251.793x` | `1.754x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `45607` (44.54 KiB) | `4048342` (3.86 MiB) | `28413` (27.75 KiB) | `0.011x` | `142.482x` | `1.605x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `92489` (90.32 KiB) | `22552` (22.02 KiB) | `43.248x` | `4.101x` | `177.368x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `8000000` (7.63 MiB) | `8118534` (7.74 MiB) | `4467090` (4.26 MiB) | `0.985x` | `1.817x` | `1.791x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `8000000` (7.63 MiB) | `8121057` (7.74 MiB) | `5218206` (4.98 MiB) | `0.985x` | `1.556x` | `1.533x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `68903` (67.29 KiB) | `8617` (8.42 KiB) | `58.053x` | `7.996x` | `464.199x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00000.parquet`: `35062` rows, `3297460` file bytes (3.14 MiB), `27216012` physical bytes (25.96 MiB), `22673201` encoded bytes (21.62 MiB), `3265419` compressed data bytes (3.11 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00001.parquet`: `35082` rows, `3237759` file bytes (3.09 MiB), `27005931` physical bytes (25.75 MiB), `22378525` encoded bytes (21.34 MiB), `3205366` compressed data bytes (3.06 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00002.parquet`: `35588` rows, `3336785` file bytes (3.18 MiB), `27825856` physical bytes (26.54 MiB), `23167217` encoded bytes (22.09 MiB), `3303547` compressed data bytes (3.15 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00003.parquet`: `35454` rows, `3249601` file bytes (3.10 MiB), `27422407` physical bytes (26.15 MiB), `22709844` encoded bytes (21.66 MiB), `3217259` compressed data bytes (3.07 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00004.parquet`: `35286` rows, `3246503` file bytes (3.10 MiB), `27482828` physical bytes (26.21 MiB), `22826898` encoded bytes (21.77 MiB), `3213770` compressed data bytes (3.06 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00005.parquet`: `35411` rows, `3235088` file bytes (3.09 MiB), `27227967` physical bytes (25.97 MiB), `22557598` encoded bytes (21.51 MiB), `3202600` compressed data bytes (3.05 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00006.parquet`: `35595` rows, `3256545` file bytes (3.11 MiB), `27439095` physical bytes (26.17 MiB), `22745284` encoded bytes (21.69 MiB), `3223993` compressed data bytes (3.07 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00007.parquet`: `35888` rows, `3310856` file bytes (3.16 MiB), `27824263` physical bytes (26.54 MiB), `23108878` encoded bytes (22.04 MiB), `3278042` compressed data bytes (3.13 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00008.parquet`: `35399` rows, `3219222` file bytes (3.07 MiB), `27457404` physical bytes (26.19 MiB), `22773019` encoded bytes (21.72 MiB), `3186636` compressed data bytes (3.04 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00009.parquet`: `35231` rows, `3208515` file bytes (3.06 MiB), `27024403` physical bytes (25.77 MiB), `22365857` encoded bytes (21.33 MiB), `3176371` compressed data bytes (3.03 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00010.parquet`: `35503` rows, `3227172` file bytes (3.08 MiB), `27414960` physical bytes (26.14 MiB), `22710706` encoded bytes (21.66 MiB), `3194838` compressed data bytes (3.05 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00011.parquet`: `35657` rows, `3198857` file bytes (3.05 MiB), `27474418` physical bytes (26.20 MiB), `22726623` encoded bytes (21.67 MiB), `3166183` compressed data bytes (3.02 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00012.parquet`: `35614` rows, `3250995` file bytes (3.10 MiB), `27415161` physical bytes (26.15 MiB), `22745508` encoded bytes (21.69 MiB), `3218974` compressed data bytes (3.07 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00013.parquet`: `35633` rows, `3238659` file bytes (3.09 MiB), `27443332` physical bytes (26.17 MiB), `22726762` encoded bytes (21.67 MiB), `3206222` compressed data bytes (3.06 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00014.parquet`: `34695` rows, `3584227` file bytes (3.42 MiB), `24860310` physical bytes (23.71 MiB), `20743674` encoded bytes (19.78 MiB), `3553064` compressed data bytes (3.39 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00015.parquet`: `35010` rows, `3830240` file bytes (3.65 MiB), `24118888` physical bytes (23.00 MiB), `20171553` encoded bytes (19.24 MiB), `3799869` compressed data bytes (3.62 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00016.parquet`: `33923` rows, `3799121` file bytes (3.62 MiB), `23169493` physical bytes (22.10 MiB), `19273844` encoded bytes (18.38 MiB), `3768859` compressed data bytes (3.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00017.parquet`: `34088` rows, `3853339` file bytes (3.67 MiB), `21750699` physical bytes (20.74 MiB), `17631589` encoded bytes (16.81 MiB), `3822803` compressed data bytes (3.65 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00018.parquet`: `33672` rows, `3925932` file bytes (3.74 MiB), `21959689` physical bytes (20.94 MiB), `17873824` encoded bytes (17.05 MiB), `3895348` compressed data bytes (3.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00019.parquet`: `33992` rows, `3941564` file bytes (3.76 MiB), `21853515` physical bytes (20.84 MiB), `17758529` encoded bytes (16.94 MiB), `3910742` compressed data bytes (3.73 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00020.parquet`: `34143` rows, `3935396` file bytes (3.75 MiB), `21904056` physical bytes (20.89 MiB), `17790853` encoded bytes (16.97 MiB), `3904761` compressed data bytes (3.72 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00021.parquet`: `33910` rows, `3878332` file bytes (3.70 MiB), `21742694` physical bytes (20.74 MiB), `17695371` encoded bytes (16.88 MiB), `3847845` compressed data bytes (3.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00022.parquet`: `34203` rows, `3895741` file bytes (3.72 MiB), `21952057` physical bytes (20.94 MiB), `17817119` encoded bytes (16.99 MiB), `3865144` compressed data bytes (3.69 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00023.parquet`: `34160` rows, `3894166` file bytes (3.71 MiB), `21942379` physical bytes (20.93 MiB), `17829367` encoded bytes (17.00 MiB), `3863519` compressed data bytes (3.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00024.parquet`: `33867` rows, `3959747` file bytes (3.78 MiB), `21879006` physical bytes (20.87 MiB), `17806767` encoded bytes (16.98 MiB), `3929301` compressed data bytes (3.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00025.parquet`: `34323` rows, `3849655` file bytes (3.67 MiB), `21793021` physical bytes (20.78 MiB), `17668934` encoded bytes (16.85 MiB), `3819163` compressed data bytes (3.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00026.parquet`: `33926` rows, `3882613` file bytes (3.70 MiB), `21783545` physical bytes (20.77 MiB), `17754356` encoded bytes (16.93 MiB), `3851914` compressed data bytes (3.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00027.parquet`: `34049` rows, `3776306` file bytes (3.60 MiB), `21467513` physical bytes (20.47 MiB), `17397330` encoded bytes (16.59 MiB), `3745825` compressed data bytes (3.57 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict/part-00028.parquet`: `25636` rows, `2931000` file bytes (2.80 MiB), `16547722` physical bytes (15.78 MiB), `13416104` encoded bytes (12.79 MiB), `2901126` compressed data bytes (2.77 MiB)
