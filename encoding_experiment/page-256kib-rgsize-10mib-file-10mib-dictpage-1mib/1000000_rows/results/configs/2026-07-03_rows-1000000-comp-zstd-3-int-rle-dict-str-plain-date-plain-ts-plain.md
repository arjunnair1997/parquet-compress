# ClickBench Parquet Experiment

- Started: `2026-07-03T15:32:40-04:00`
- Write elapsed: `15.902s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `550724373` (525.21 MiB)
- Compressed column data bytes after codec compression: `93577314` (89.24 MiB)
- Parquet file bytes: `94525517` (90.15 MiB)
- Physical/encoded ratio: `1.294x`
- Encoded/compressed-data ratio: `5.885x`
- Physical/compressed-data ratio: `7.613x`
- Physical/parquet-file ratio: `7.537x`
- Files: `29`

## Settings

- Compression: `zstd-3`
- Int encoding: `rle-dict`
- String encoding: `plain`
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
- Files read: `29`
- Elapsed: `7.2s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `9851200` (9.39 MiB) | `9852575` (9.40 MiB) | `0.812x` | `1.000x` | `0.812x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `79757` (77.89 KiB) | `50727` (49.54 KiB) | `50.152x` | `1.572x` | `78.853x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:525` | `1000000` | `138409995` (132.00 MiB) | `142862851` (136.24 MiB) | `13926821` (13.28 MiB) | `0.969x` | `10.258x` | `9.938x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4949` (4.83 KiB) | `5975` (5.83 KiB) | `808.244x` | `0.828x` | `669.456x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004556` (7.63 MiB) | `2519179` (2.40 MiB) | `0.999x` | `3.177x` | `3.176x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4920` (4.80 KiB) | `0.999x` | `813.725x` | `813.008x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4994` (4.88 KiB) | `6020` (5.88 KiB) | `800.961x` | `0.830x` | `664.452x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `941748` (919.68 KiB) | `806924` (788.01 KiB) | `4.247x` | `1.167x` | `4.957x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `435425` (425.22 KiB) | `245586` (239.83 KiB) | `9.186x` | `1.773x` | `16.288x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `1228281` (1.17 MiB) | `1102165` (1.05 MiB) | `6.513x` | `1.114x` | `7.258x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4949` (4.83 KiB) | `5975` (5.83 KiB) | `808.244x` | `0.828x` | `669.456x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `229346` (223.97 KiB) | `119893` (117.08 KiB) | `17.441x` | `1.913x` | `33.363x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `226591` (221.28 KiB) | `125153` (122.22 KiB) | `17.653x` | `1.811x` | `31.961x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:352` | `1000000` | `88562192` (84.46 MiB) | `92653141` (88.36 MiB) | `15306710` (14.60 MiB) | `0.956x` | `6.053x` | `5.786x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:326` | `1000000` | `79583339` (75.90 MiB) | `83646114` (79.77 MiB) | `14212298` (13.55 MiB) | `0.951x` | `5.885x` | `5.600x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `92576` (90.41 KiB) | `83712` (81.75 KiB) | `43.208x` | `1.106x` | `47.783x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `515672` (503.59 KiB) | `216181` (211.11 KiB) | `7.757x` | `2.385x` | `18.503x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `379518` (370.62 KiB) | `164769` (160.91 KiB) | `10.540x` | `2.303x` | `24.276x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `161519` (157.73 KiB) | `71400` (69.73 KiB) | `24.765x` | `2.262x` | `56.022x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `74399` (72.66 KiB) | `40756` (39.80 KiB) | `53.764x` | `1.825x` | `98.145x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `368799` (360.16 KiB) | `203528` (198.76 KiB) | `10.846x` | `1.812x` | `19.653x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `371693` (362.98 KiB) | `205409` (200.59 KiB) | `10.762x` | `1.810x` | `19.473x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `111133` (108.53 KiB) | `61927` (60.48 KiB) | `35.993x` | `1.795x` | `64.592x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `88463` (86.39 KiB) | `49129` (47.98 KiB) | `45.217x` | `1.801x` | `81.418x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `214374` (209.35 KiB) | `113252` (110.60 KiB) | `18.659x` | `1.893x` | `35.319x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3354477` (3.20 MiB) | `7357824` (7.02 MiB) | `246366` (240.59 KiB) | `0.456x` | `29.865x` | `13.616x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `36548` (35.69 KiB) | `25367` (24.77 KiB) | `109.445x` | `1.441x` | `157.685x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `33050` (32.28 KiB) | `23182` (22.64 KiB) | `121.029x` | `1.426x` | `172.548x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `275570` (269.11 KiB) | `177456` (173.30 KiB) | `14.515x` | `1.553x` | `22.541x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3767530` (3.59 MiB) | `7770883` (7.41 MiB) | `137153` (133.94 KiB) | `0.485x` | `56.658x` | `27.470x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `6055` (5.91 KiB) | `7083` (6.92 KiB) | `660.611x` | `0.855x` | `564.732x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `6371` (6.22 KiB) | `7397` (7.22 KiB) | `627.845x` | `0.861x` | `540.760x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `29875` (29.17 KiB) | `24150` (23.58 KiB) | `133.891x` | `1.237x` | `165.631x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `34956` (34.14 KiB) | `23778` (23.22 KiB) | `114.430x` | `1.470x` | `168.223x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `81583` (79.67 KiB) | `4084889` (3.90 MiB) | `22587` (22.06 KiB) | `0.020x` | `180.851x` | `3.612x` | `81583` (79.67 KiB) |
| `Params` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002160` (3.82 MiB) | `2869` (2.80 KiB) | `0.000x` | `1394.967x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `747959` (730.43 KiB) | `619973` (605.44 KiB) | `5.348x` | `1.206x` | `6.452x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `289305` (282.52 KiB) | `180684` (176.45 KiB) | `13.826x` | `1.601x` | `22.138x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `167349` (163.43 KiB) | `76166` (74.38 KiB) | `23.902x` | `2.197x` | `52.517x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3528017` (3.36 MiB) | `7535016` (7.19 MiB) | `719279` (702.42 KiB) | `0.468x` | `10.476x` | `4.905x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `34609` (33.80 KiB) | `22787` (22.25 KiB) | `115.577x` | `1.519x` | `175.539x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `80984` (79.09 KiB) | `82044` (80.12 KiB) | `49.392x` | `0.987x` | `48.754x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `694445` (678.17 KiB) | `471396` (460.35 KiB) | `5.760x` | `1.473x` | `8.485x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `750345` (732.76 KiB) | `583325` (569.65 KiB) | `5.331x` | `1.286x` | `6.857x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `164240` (160.39 KiB) | `84507` (82.53 KiB) | `24.355x` | `1.944x` | `47.333x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004551` (7.63 MiB) | `2475975` (2.36 MiB) | `0.999x` | `3.233x` | `3.231x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `155901` (152.25 KiB) | `83442` (81.49 KiB) | `25.657x` | `1.868x` | `47.937x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `89019` (86.93 KiB) | `56082` (54.77 KiB) | `44.934x` | `1.587x` | `71.324x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `230469` (225.07 KiB) | `123923` (121.02 KiB) | `17.356x` | `1.860x` | `32.278x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `5062` (4.94 KiB) | `6088` (5.95 KiB) | `790.202x` | `0.831x` | `657.030x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:85` | `1000000` | `13587860` (12.96 MiB) | `17594822` (16.78 MiB) | `14724` (14.38 KiB) | `0.772x` | `1194.976x` | `922.838x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `6974` (6.81 KiB) | `8000` (7.81 KiB) | `573.559x` | `0.872x` | `500.000x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `54176` (52.91 KiB) | `36721` (35.86 KiB) | `73.833x` | `1.475x` | `108.929x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `8324` (8.13 KiB) | `9258` (9.04 KiB) | `480.538x` | `0.899x` | `432.059x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `16299` (15.92 KiB) | `17331` (16.92 KiB) | `245.414x` | `0.940x` | `230.800x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `1186777` (1.13 MiB) | `1060473` (1.01 MiB) | `6.741x` | `1.119x` | `7.544x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:137` | `1000000` | `27797671` (26.51 MiB) | `31857505` (30.38 MiB) | `5325631` (5.08 MiB) | `0.873x` | `5.982x` | `5.220x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4515884` (4.31 MiB) | `4517046` (4.31 MiB) | `0.886x` | `1.000x` | `0.886x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4949` (4.83 KiB) | `5975` (5.83 KiB) | `808.244x` | `0.828x` | `669.456x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4949` (4.83 KiB) | `5975` (5.83 KiB) | `808.244x` | `0.828x` | `669.456x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4949` (4.83 KiB) | `5975` (5.83 KiB) | `808.244x` | `0.828x` | `669.456x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `65818` (64.28 KiB) | `50316` (49.14 KiB) | `60.774x` | `1.308x` | `79.498x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4949` (4.83 KiB) | `5975` (5.83 KiB) | `808.244x` | `0.828x` | `669.456x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `1000000` (976.56 KiB) | `5002901` (4.77 MiB) | `28126` (27.47 KiB) | `0.200x` | `177.875x` | `35.554x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004554` (7.63 MiB) | `2518850` (2.40 MiB) | `0.999x` | `3.178x` | `3.176x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `185832` (181.48 KiB) | `118226` (115.46 KiB) | `21.525x` | `1.572x` | `33.834x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `133318` (130.19 KiB) | `80090` (78.21 KiB) | `30.003x` | `1.665x` | `49.944x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `142367` (139.03 KiB) | `90277` (88.16 KiB) | `28.096x` | `1.577x` | `44.308x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `493958` (482.38 KiB) | `318561` (311.09 KiB) | `8.098x` | `1.551x` | `12.556x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `412858` (403.18 KiB) | `242016` (236.34 KiB) | `9.689x` | `1.706x` | `16.528x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `931028` (909.21 KiB) | `703363` (686.88 KiB) | `4.296x` | `1.324x` | `5.687x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `174383` (170.30 KiB) | `154094` (150.48 KiB) | `22.938x` | `1.132x` | `25.958x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4949` (4.83 KiB) | `5975` (5.83 KiB) | `808.244x` | `0.828x` | `669.456x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `69739` (68.10 KiB) | `51489` (50.28 KiB) | `57.357x` | `1.354x` | `77.686x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `2001192` (1.91 MiB) | `6004436` (5.73 MiB) | `32311` (31.55 KiB) | `0.333x` | `185.833x` | `61.935x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3325142` (3.17 MiB) | `7328496` (6.99 MiB) | `122286` (119.42 KiB) | `0.454x` | `59.929x` | `27.192x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002160` (3.82 MiB) | `2869` (2.80 KiB) | `0.000x` | `1394.967x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002160` (3.82 MiB) | `2869` (2.80 KiB) | `0.000x` | `1394.967x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4949` (4.83 KiB) | `5975` (5.83 KiB) | `808.244x` | `0.828x` | `669.456x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `122652` (119.78 KiB) | `79121` (77.27 KiB) | `32.613x` | `1.550x` | `50.555x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `343868` (335.81 KiB) | `156600` (152.93 KiB) | `11.632x` | `2.196x` | `25.543x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `704511` (688.00 KiB) | `378863` (369.98 KiB) | `5.678x` | `1.860x` | `10.558x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `1804514` (1.72 MiB) | `1574230` (1.50 MiB) | `2.217x` | `1.146x` | `2.541x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `1357248` (1.29 MiB) | `1046981` (1022.44 KiB) | `2.947x` | `1.296x` | `3.821x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `1134386` (1.08 MiB) | `689064` (672.91 KiB) | `3.526x` | `1.646x` | `5.805x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `5895` (5.76 KiB) | `6921` (6.76 KiB) | `678.541x` | `0.852x` | `577.951x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `1024` (1.00 KiB) | `4003995` (3.82 MiB) | `4864` (4.75 KiB) | `0.000x` | `823.190x` | `0.211x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `6089` (5.95 KiB) | `7115` (6.95 KiB) | `1313.845x` | `0.856x` | `1124.385x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002160` (3.82 MiB) | `2869` (2.80 KiB) | `0.000x` | `1394.967x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3000000` (2.86 MiB) | `7003298` (6.68 MiB) | `5256` (5.13 KiB) | `0.428x` | `1332.439x` | `570.776x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4949` (4.83 KiB) | `5975` (5.83 KiB) | `808.244x` | `0.828x` | `669.456x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `58030` (56.67 KiB) | `4062487` (3.87 MiB) | `19196` (18.75 KiB) | `0.014x` | `211.632x` | `3.023x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `22051` (21.53 KiB) | `4025596` (3.84 MiB) | `16539` (16.15 KiB) | `0.005x` | `243.400x` | `1.333x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `25445` (24.85 KiB) | `4029016` (3.84 MiB) | `18939` (18.50 KiB) | `0.006x` | `212.736x` | `1.344x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `48191` (47.06 KiB) | `4051713` (3.86 MiB) | `13233` (12.92 KiB) | `0.012x` | `306.182x` | `3.642x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `49433` (48.27 KiB) | `4053175` (3.87 MiB) | `21004` (20.51 KiB) | `0.012x` | `192.972x` | `2.354x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `16873` (16.48 KiB) | `4019915` (3.83 MiB) | `16944` (16.55 KiB) | `0.004x` | `237.247x` | `0.996x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `91870` (89.72 KiB) | `4097338` (3.91 MiB) | `29663` (28.97 KiB) | `0.022x` | `138.130x` | `3.097x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `13001` (12.70 KiB) | `4016617` (3.83 MiB) | `14020` (13.69 KiB) | `0.003x` | `286.492x` | `0.927x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `28101` (27.44 KiB) | `4032690` (3.85 MiB) | `15649` (15.28 KiB) | `0.007x` | `257.696x` | `1.796x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `45607` (44.54 KiB) | `4048307` (3.86 MiB) | `28448` (27.78 KiB) | `0.011x` | `142.306x` | `1.603x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `21296` (20.80 KiB) | `16292` (15.91 KiB) | `187.829x` | `1.307x` | `245.519x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `3779873` (3.60 MiB) | `3531855` (3.37 MiB) | `2.116x` | `1.070x` | `2.265x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `4619429` (4.41 MiB) | `4545945` (4.34 MiB) | `1.732x` | `1.016x` | `1.760x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `5882` (5.74 KiB) | `6908` (6.75 KiB) | `680.041x` | `0.851x` | `579.039x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00000.parquet`: `36038` rows, `3079288` file bytes (2.94 MiB), `27980233` physical bytes (26.68 MiB), `22012339` encoded bytes (20.99 MiB), `3045510` compressed data bytes (2.90 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00001.parquet`: `36343` rows, `3047298` file bytes (2.91 MiB), `27937012` physical bytes (26.64 MiB), `21911342` encoded bytes (20.90 MiB), `3012957` compressed data bytes (2.87 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00002.parquet`: `35992` rows, `3069490` file bytes (2.93 MiB), `28192611` physical bytes (26.89 MiB), `22233859` encoded bytes (21.20 MiB), `3034474` compressed data bytes (2.89 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00003.parquet`: `36456` rows, `3033255` file bytes (2.89 MiB), `28183912` physical bytes (26.88 MiB), `22123850` encoded bytes (21.10 MiB), `2999061` compressed data bytes (2.86 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00004.parquet`: `35701` rows, `2988737` file bytes (2.85 MiB), `27820403` physical bytes (26.53 MiB), `21887878` encoded bytes (20.87 MiB), `2953970` compressed data bytes (2.82 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00005.parquet`: `36254` rows, `2985468` file bytes (2.85 MiB), `27758230` physical bytes (26.47 MiB), `21723831` encoded bytes (20.72 MiB), `2951320` compressed data bytes (2.81 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00006.parquet`: `36430` rows, `3049238` file bytes (2.91 MiB), `28174079` physical bytes (26.87 MiB), `22135482` encoded bytes (21.11 MiB), `3014620` compressed data bytes (2.87 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00007.parquet`: `36298` rows, `3037464` file bytes (2.90 MiB), `28201272` physical bytes (26.89 MiB), `22176757` encoded bytes (21.15 MiB), `3002958` compressed data bytes (2.86 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00008.parquet`: `36497` rows, `3003863` file bytes (2.86 MiB), `28222097` physical bytes (26.91 MiB), `22137351` encoded bytes (21.11 MiB), `2969538` compressed data bytes (2.83 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00009.parquet`: `36278` rows, `3031195` file bytes (2.89 MiB), `27970915` physical bytes (26.68 MiB), `21940357` encoded bytes (20.92 MiB), `2997266` compressed data bytes (2.86 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00010.parquet`: `36597` rows, `2999071` file bytes (2.86 MiB), `28188489` physical bytes (26.88 MiB), `22075179` encoded bytes (21.05 MiB), `2964929` compressed data bytes (2.83 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00011.parquet`: `36609` rows, `2989791` file bytes (2.85 MiB), `28222372` physical bytes (26.91 MiB), `22121236` encoded bytes (21.10 MiB), `2955672` compressed data bytes (2.82 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00012.parquet`: `36522` rows, `3035883` file bytes (2.90 MiB), `27950756` physical bytes (26.66 MiB), `21887848` encoded bytes (20.87 MiB), `3001857` compressed data bytes (2.86 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00013.parquet`: `36461` rows, `3062198` file bytes (2.92 MiB), `28165984` physical bytes (26.86 MiB), `22103332` encoded bytes (21.08 MiB), `3027851` compressed data bytes (2.89 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00014.parquet`: `35034` rows, `3477970` file bytes (3.32 MiB), `24035095` physical bytes (22.92 MiB), `18659788` encoded bytes (17.80 MiB), `3445958` compressed data bytes (3.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00015.parquet`: `35897` rows, `3691531` file bytes (3.52 MiB), `25161029` physical bytes (24.00 MiB), `19637633` encoded bytes (18.73 MiB), `3659479` compressed data bytes (3.49 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00016.parquet`: `35126` rows, `3695299` file bytes (3.52 MiB), `22871619` physical bytes (21.81 MiB), `17353117` encoded bytes (16.55 MiB), `3663134` compressed data bytes (3.49 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00017.parquet`: `34969` rows, `3733459` file bytes (3.56 MiB), `22564534` physical bytes (21.52 MiB), `17047322` encoded bytes (16.26 MiB), `3700903` compressed data bytes (3.53 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00018.parquet`: `34184` rows, `3788921` file bytes (3.61 MiB), `22301900` physical bytes (21.27 MiB), `16903599` encoded bytes (16.12 MiB), `3756853` compressed data bytes (3.58 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00019.parquet`: `34927` rows, `3715575` file bytes (3.54 MiB), `22289836` physical bytes (21.26 MiB), `16777605` encoded bytes (16.00 MiB), `3682940` compressed data bytes (3.51 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00020.parquet`: `34638` rows, `3757172` file bytes (3.58 MiB), `22269071` physical bytes (21.24 MiB), `16822270` encoded bytes (16.04 MiB), `3724803` compressed data bytes (3.55 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00021.parquet`: `34951` rows, `3686889` file bytes (3.52 MiB), `22339672` physical bytes (21.30 MiB), `16826925` encoded bytes (16.05 MiB), `3654588` compressed data bytes (3.49 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00022.parquet`: `34734` rows, `3712688` file bytes (3.54 MiB), `22327715` physical bytes (21.29 MiB), `16844797` encoded bytes (16.06 MiB), `3680755` compressed data bytes (3.51 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00023.parquet`: `34504` rows, `3750353` file bytes (3.58 MiB), `22297276` physical bytes (21.26 MiB), `16856807` encoded bytes (16.08 MiB), `3717802` compressed data bytes (3.55 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00024.parquet`: `35098` rows, `3690388` file bytes (3.52 MiB), `22293000` physical bytes (21.26 MiB), `16770926` encoded bytes (15.99 MiB), `3658500` compressed data bytes (3.49 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00025.parquet`: `34670` rows, `3706901` file bytes (3.54 MiB), `22101304` physical bytes (21.08 MiB), `16652012` encoded bytes (15.88 MiB), `3674478` compressed data bytes (3.50 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00026.parquet`: `35176` rows, `3677652` file bytes (3.51 MiB), `22343425` physical bytes (21.31 MiB), `16806657` encoded bytes (16.03 MiB), `3645381` compressed data bytes (3.48 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00027.parquet`: `34773` rows, `3712879` file bytes (3.54 MiB), `22348264` physical bytes (21.31 MiB), `16866585` encoded bytes (16.09 MiB), `3680641` compressed data bytes (3.51 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain/part-00028.parquet`: `2843` rows, `315601` file bytes (308.20 KiB), `1886519` physical bytes (1.80 MiB), `1427689` encoded bytes (1.36 MiB), `299116` compressed data bytes (292.11 KiB)
