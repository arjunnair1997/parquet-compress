# ClickBench Parquet Experiment

- Started: `2026-07-03T19:43:59-04:00`
- Write elapsed: `11.897s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `473346190` (451.42 MiB)
- Compressed column data bytes after codec compression: `79529029` (75.84 MiB)
- Parquet file bytes: `80431194` (76.71 MiB)
- Physical/encoded ratio: `1.505x`
- Encoded/compressed-data ratio: `5.952x`
- Physical/compressed-data ratio: `8.958x`
- Physical/parquet-file ratio: `8.857x`
- Files: `30`

## Settings

- Compression: `zstd-3`
- Int encoding: `plain`
- String encoding: `rle-dict`
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
- Elapsed: `7.109s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `8004792` (7.63 MiB) | `8005639` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003776` (3.82 MiB) | `64680` (63.16 KiB) | `0.999x` | `61.901x` | `61.843x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `138409995` (132.00 MiB) | `28831150` (27.50 MiB) | `7976017` (7.61 MiB) | `4.801x` | `3.615x` | `17.353x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `5178` (5.06 KiB) | `0.999x` | `773.216x` | `772.499x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `7359529` (7.02 MiB) | `4021744` (3.84 MiB) | `1.087x` | `1.830x` | `1.989x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `5182` (5.06 KiB) | `0.999x` | `772.620x` | `771.903x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `5197` (5.08 KiB) | `0.999x` | `770.389x` | `769.675x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003780` (3.82 MiB) | `408551` (398.98 KiB) | `0.999x` | `9.800x` | `9.791x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003772` (3.82 MiB) | `190959` (186.48 KiB) | `0.999x` | `20.967x` | `20.947x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `8004790` (7.63 MiB) | `618431` (603.94 KiB) | `0.999x` | `12.944x` | `12.936x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `4459` (4.35 KiB) | `0.999x` | `897.895x` | `897.062x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `107764` (105.24 KiB) | `0.999x` | `37.153x` | `37.118x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003778` (3.82 MiB) | `134905` (131.74 KiB) | `0.999x` | `29.678x` | `29.650x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `88562192` (84.46 MiB) | `44085355` (42.04 MiB) | `12716670` (12.13 MiB) | `2.009x` | `3.467x` | `6.964x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `79583339` (75.90 MiB) | `34396350` (32.80 MiB) | `11737340` (11.19 MiB) | `2.314x` | `2.931x` | `6.780x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003772` (3.82 MiB) | `179238` (175.04 KiB) | `0.999x` | `22.338x` | `22.317x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003771` (3.82 MiB) | `275831` (269.37 KiB) | `0.999x` | `14.515x` | `14.502x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `231744` (226.31 KiB) | `0.999x` | `17.277x` | `17.260x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `88077` (86.01 KiB) | `0.999x` | `45.458x` | `45.415x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003771` (3.82 MiB) | `48542` (47.40 KiB) | `0.999x` | `82.481x` | `82.403x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003776` (3.82 MiB) | `187865` (183.46 KiB) | `0.999x` | `21.312x` | `21.292x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003772` (3.82 MiB) | `186758` (182.38 KiB) | `0.999x` | `21.438x` | `21.418x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `83118` (81.17 KiB) | `0.999x` | `48.170x` | `48.124x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `53647` (52.39 KiB) | `0.999x` | `74.632x` | `74.561x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `129321` (126.29 KiB) | `0.999x` | `30.960x` | `30.931x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `3354477` (3.20 MiB) | `266150` (259.91 KiB) | `145709` (142.29 KiB) | `12.604x` | `1.827x` | `23.022x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003776` (3.82 MiB) | `26715` (26.09 KiB) | `0.999x` | `149.870x` | `149.729x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `25408` (24.81 KiB) | `0.999x` | `157.579x` | `157.431x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `156522` (152.85 KiB) | `0.999x` | `25.580x` | `25.556x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `3767530` (3.59 MiB) | `150462` (146.94 KiB) | `83866` (81.90 KiB) | `25.040x` | `1.794x` | `44.923x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003721` (3.82 MiB) | `6353` (6.20 KiB) | `0.999x` | `630.210x` | `629.624x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003725` (3.82 MiB) | `6801` (6.64 KiB) | `0.999x` | `588.697x` | `588.149x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003773` (3.82 MiB) | `29441` (28.75 KiB) | `0.999x` | `135.993x` | `135.865x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `22842` (22.31 KiB) | `0.999x` | `175.281x` | `175.116x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `81583` (79.67 KiB) | `25139` (24.55 KiB) | `20411` (19.93 KiB) | `3.245x` | `1.232x` | `3.997x` | `81583` (79.67 KiB) |
| `Params` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `0` (0 B) | `4249` (4.15 KiB) | `5329` (5.20 KiB) | `0.000x` | `0.797x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003772` (3.82 MiB) | `324260` (316.66 KiB) | `0.999x` | `12.347x` | `12.336x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003773` (3.82 MiB) | `288810` (282.04 KiB) | `0.999x` | `13.863x` | `13.850x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `102149` (99.75 KiB) | `0.999x` | `39.195x` | `39.158x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `3528017` (3.36 MiB) | `1642536` (1.57 MiB) | `636510` (621.59 KiB) | `2.148x` | `2.581x` | `5.543x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003767` (3.82 MiB) | `31005` (30.28 KiB) | `0.999x` | `129.133x` | `129.011x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003773` (3.82 MiB) | `165383` (161.51 KiB) | `0.999x` | `24.209x` | `24.186x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003777` (3.82 MiB) | `306341` (299.16 KiB) | `0.999x` | `13.070x` | `13.057x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003776` (3.82 MiB) | `319485` (312.00 KiB) | `0.999x` | `12.532x` | `12.520x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `100400` (98.05 KiB) | `0.999x` | `39.878x` | `39.841x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `7210536` (6.88 MiB) | `3958282` (3.77 MiB) | `1.109x` | `1.822x` | `2.021x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003776` (3.82 MiB) | `90905` (88.77 KiB) | `0.999x` | `44.044x` | `44.002x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003773` (3.82 MiB) | `74516` (72.77 KiB) | `0.999x` | `53.730x` | `53.680x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003776` (3.82 MiB) | `124398` (121.48 KiB) | `0.999x` | `32.185x` | `32.155x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `4590` (4.48 KiB) | `0.999x` | `872.269x` | `871.460x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `13587860` (12.96 MiB) | `8917` (8.71 KiB) | `9999` (9.76 KiB) | `1523.815x` | `0.892x` | `1358.922x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003735` (3.82 MiB) | `7337` (7.17 KiB) | `0.999x` | `545.691x` | `545.182x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003745` (3.82 MiB) | `57126` (55.79 KiB) | `0.999x` | `70.086x` | `70.021x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003741` (3.82 MiB) | `8325` (8.13 KiB) | `0.999x` | `480.930x` | `480.480x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003725` (3.82 MiB) | `25530` (24.93 KiB) | `0.999x` | `156.824x` | `156.678x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `8004790` (7.63 MiB) | `694331` (678.06 KiB) | `0.999x` | `11.529x` | `11.522x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `27797671` (26.51 MiB) | `21268274` (20.28 MiB) | `4889469` (4.66 MiB) | `1.307x` | `4.350x` | `5.685x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `3813061` (3.64 MiB) | `0.999x` | `1.050x` | `1.049x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `4459` (4.35 KiB) | `0.999x` | `897.895x` | `897.062x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `4459` (4.35 KiB) | `0.999x` | `897.895x` | `897.062x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `4459` (4.35 KiB) | `0.999x` | `897.895x` | `897.062x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `82924` (80.98 KiB) | `0.999x` | `48.282x` | `48.237x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `4459` (4.35 KiB) | `0.999x` | `897.895x` | `897.062x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `1000000` (976.56 KiB) | `33679` (32.89 KiB) | `29652` (28.96 KiB) | `29.692x` | `1.136x` | `33.725x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `7355383` (7.01 MiB) | `4023316` (3.84 MiB) | `1.088x` | `1.828x` | `1.988x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003773` (3.82 MiB) | `143233` (139.88 KiB) | `0.999x` | `27.953x` | `27.927x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003773` (3.82 MiB) | `107930` (105.40 KiB) | `0.999x` | `37.096x` | `37.061x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `122863` (119.98 KiB) | `0.999x` | `32.587x` | `32.557x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003771` (3.82 MiB) | `193825` (189.28 KiB) | `0.999x` | `20.657x` | `20.637x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `173402` (169.34 KiB) | `0.999x` | `23.090x` | `23.068x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003783` (3.82 MiB) | `426804` (416.80 KiB) | `0.999x` | `9.381x` | `9.372x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003764` (3.82 MiB) | `71135` (69.47 KiB) | `0.999x` | `56.284x` | `56.231x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `4462` (4.36 KiB) | `0.999x` | `897.292x` | `896.459x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003756` (3.82 MiB) | `55580` (54.28 KiB) | `0.999x` | `72.036x` | `71.968x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `2001192` (1.91 MiB) | `42619` (41.62 KiB) | `27625` (26.98 KiB) | `46.955x` | `1.543x` | `72.441x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `3325142` (3.17 MiB) | `136011` (132.82 KiB) | `69385` (67.76 KiB) | `24.448x` | `1.960x` | `47.923x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `0` (0 B) | `4249` (4.15 KiB) | `5329` (5.20 KiB) | `0.000x` | `0.797x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `0` (0 B) | `4249` (4.15 KiB) | `5329` (5.20 KiB) | `0.000x` | `0.797x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `4459` (4.35 KiB) | `0.999x` | `897.895x` | `897.062x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003728` (3.82 MiB) | `62047` (60.59 KiB) | `0.999x` | `64.527x` | `64.467x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `135712` (132.53 KiB) | `0.999x` | `29.502x` | `29.474x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003779` (3.82 MiB) | `334105` (326.27 KiB) | `0.999x` | `11.984x` | `11.972x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `1246851` (1.19 MiB) | `0.999x` | `3.211x` | `3.208x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `938224` (916.23 KiB) | `0.999x` | `4.267x` | `4.263x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003815` (3.82 MiB) | `550507` (537.60 KiB) | `0.999x` | `7.273x` | `7.266x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `5585` (5.45 KiB) | `0.999x` | `716.869x` | `716.204x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `1024` (1.00 KiB) | `5900` (5.76 KiB) | `6980` (6.82 KiB) | `0.174x` | `0.845x` | `0.147x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `8004674` (7.63 MiB) | `6011` (5.87 KiB) | `0.999x` | `1331.671x` | `1330.893x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `0` (0 B) | `4249` (4.15 KiB) | `5329` (5.20 KiB) | `0.000x` | `0.797x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `3000000` (2.86 MiB) | `5149` (5.03 KiB) | `6229` (6.08 KiB) | `582.637x` | `0.827x` | `481.618x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `4459` (4.35 KiB) | `0.999x` | `897.895x` | `897.062x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `58030` (56.67 KiB) | `18741` (18.30 KiB) | `18037` (17.61 KiB) | `3.096x` | `1.039x` | `3.217x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `22051` (21.53 KiB) | `16323` (15.94 KiB) | `15676` (15.31 KiB) | `1.351x` | `1.041x` | `1.407x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `25445` (24.85 KiB) | `22780` (22.25 KiB) | `18487` (18.05 KiB) | `1.117x` | `1.232x` | `1.376x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `48191` (47.06 KiB) | `14826` (14.48 KiB) | `13216` (12.91 KiB) | `3.250x` | `1.122x` | `3.646x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `49433` (48.27 KiB) | `22714` (22.18 KiB) | `18496` (18.06 KiB) | `2.176x` | `1.228x` | `2.673x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `16873` (16.48 KiB) | `16586` (16.20 KiB) | `14766` (14.42 KiB) | `1.017x` | `1.123x` | `1.143x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `91870` (89.72 KiB) | `37140` (36.27 KiB) | `27472` (26.83 KiB) | `2.474x` | `1.352x` | `3.344x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `13001` (12.70 KiB) | `14833` (14.49 KiB) | `15145` (14.79 KiB) | `0.876x` | `0.979x` | `0.858x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `28101` (27.44 KiB) | `16699` (16.31 KiB) | `16242` (15.86 KiB) | `1.683x` | `1.028x` | `1.730x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:60, DICTIONARY_PAGE/PLAIN:60` | `1000000` | `45607` (44.54 KiB) | `41092` (40.13 KiB) | `22794` (22.26 KiB) | `1.110x` | `1.803x` | `2.001x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003743` (3.82 MiB) | `21399` (20.90 KiB) | `0.999x` | `187.100x` | `186.925x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `8004795` (7.63 MiB) | `2843626` (2.71 MiB) | `0.999x` | `2.815x` | `2.813x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `8004796` (7.63 MiB) | `3582216` (3.42 MiB) | `0.999x` | `2.235x` | `2.233x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003727` (3.82 MiB) | `5863` (5.73 KiB) | `0.999x` | `682.880x` | `682.245x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00000.parquet`: `35632` rows, `2413079` file bytes (2.30 MiB), `27674993` physical bytes (26.39 MiB), `15808851` encoded bytes (15.08 MiB), `2382589` compressed data bytes (2.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00001.parquet`: `35762` rows, `2427081` file bytes (2.31 MiB), `27464721` physical bytes (26.19 MiB), `15769871` encoded bytes (15.04 MiB), `2396433` compressed data bytes (2.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00002.parquet`: `35552` rows, `2451933` file bytes (2.34 MiB), `27861711` physical bytes (26.57 MiB), `15793626` encoded bytes (15.06 MiB), `2420633` compressed data bytes (2.31 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00003.parquet`: `35735` rows, `2417389` file bytes (2.31 MiB), `27664178` physical bytes (26.38 MiB), `15792749` encoded bytes (15.06 MiB), `2387047` compressed data bytes (2.28 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00004.parquet`: `35810` rows, `2405694` file bytes (2.29 MiB), `27850610` physical bytes (26.56 MiB), `15758149` encoded bytes (15.03 MiB), `2374555` compressed data bytes (2.26 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00005.parquet`: `35972` rows, `2410649` file bytes (2.30 MiB), `27597125` physical bytes (26.32 MiB), `15758516` encoded bytes (15.03 MiB), `2380042` compressed data bytes (2.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00006.parquet`: `35862` rows, `2387708` file bytes (2.28 MiB), `27718074` physical bytes (26.43 MiB), `15790994` encoded bytes (15.06 MiB), `2357300` compressed data bytes (2.25 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00007.parquet`: `35744` rows, `2425250` file bytes (2.31 MiB), `27679850` physical bytes (26.40 MiB), `15771861` encoded bytes (15.04 MiB), `2394372` compressed data bytes (2.28 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00008.parquet`: `35826` rows, `2399016` file bytes (2.29 MiB), `27796333` physical bytes (26.51 MiB), `15762659` encoded bytes (15.03 MiB), `2368251` compressed data bytes (2.26 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00009.parquet`: `35849` rows, `2409332` file bytes (2.30 MiB), `27610093` physical bytes (26.33 MiB), `15785934` encoded bytes (15.05 MiB), `2378879` compressed data bytes (2.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00010.parquet`: `35873` rows, `2388480` file bytes (2.28 MiB), `27590087` physical bytes (26.31 MiB), `15779766` encoded bytes (15.05 MiB), `2357954` compressed data bytes (2.25 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00011.parquet`: `36053` rows, `2372888` file bytes (2.26 MiB), `27801432` physical bytes (26.51 MiB), `15767794` encoded bytes (15.04 MiB), `2342385` compressed data bytes (2.23 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00012.parquet`: `35933` rows, `2402125` file bytes (2.29 MiB), `27588538` physical bytes (26.31 MiB), `15789390` encoded bytes (15.06 MiB), `2371824` compressed data bytes (2.26 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00013.parquet`: `35787` rows, `2395913` file bytes (2.28 MiB), `27702159` physical bytes (26.42 MiB), `15787670` encoded bytes (15.06 MiB), `2365169` compressed data bytes (2.26 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00014.parquet`: `32008` rows, `2755875` file bytes (2.63 MiB), `22484106` physical bytes (21.44 MiB), `16030839` encoded bytes (15.29 MiB), `2725973` compressed data bytes (2.60 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00015.parquet`: `31229` rows, `2787759` file bytes (2.66 MiB), `21524857` physical bytes (20.53 MiB), `16114826` encoded bytes (15.37 MiB), `2758563` compressed data bytes (2.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00016.parquet`: `30575` rows, `2878112` file bytes (2.74 MiB), `21185841` physical bytes (20.20 MiB), `16169126` encoded bytes (15.42 MiB), `2848953` compressed data bytes (2.72 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00017.parquet`: `31984` rows, `2974514` file bytes (2.84 MiB), `20318663` physical bytes (19.38 MiB), `15942402` encoded bytes (15.20 MiB), `2945055` compressed data bytes (2.81 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00018.parquet`: `31276` rows, `3001535` file bytes (2.86 MiB), `20377693` physical bytes (19.43 MiB), `16007819` encoded bytes (15.27 MiB), `2971730` compressed data bytes (2.83 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00019.parquet`: `31361` rows, `3031777` file bytes (2.89 MiB), `20192414` physical bytes (19.26 MiB), `15962849` encoded bytes (15.22 MiB), `3002326` compressed data bytes (2.86 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00020.parquet`: `31501` rows, `3003202` file bytes (2.86 MiB), `20214997` physical bytes (19.28 MiB), `15994378` encoded bytes (15.25 MiB), `2973238` compressed data bytes (2.84 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00021.parquet`: `31532` rows, `3012950` file bytes (2.87 MiB), `20255665` physical bytes (19.32 MiB), `15966872` encoded bytes (15.23 MiB), `2983314` compressed data bytes (2.85 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00022.parquet`: `31754` rows, `2990279` file bytes (2.85 MiB), `20313035` physical bytes (19.37 MiB), `15958968` encoded bytes (15.22 MiB), `2960568` compressed data bytes (2.82 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00023.parquet`: `31653` rows, `2980029` file bytes (2.84 MiB), `20278410` physical bytes (19.34 MiB), `15964160` encoded bytes (15.22 MiB), `2950542` compressed data bytes (2.81 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00024.parquet`: `31280` rows, `3017192` file bytes (2.88 MiB), `20331437` physical bytes (19.39 MiB), `16000897` encoded bytes (15.26 MiB), `2987369` compressed data bytes (2.85 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00025.parquet`: `31643` rows, `3000697` file bytes (2.86 MiB), `20262677` physical bytes (19.32 MiB), `15939334` encoded bytes (15.20 MiB), `2971313` compressed data bytes (2.83 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00026.parquet`: `31949` rows, `2996220` file bytes (2.86 MiB), `20349498` physical bytes (19.41 MiB), `15960012` encoded bytes (15.22 MiB), `2966441` compressed data bytes (2.83 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00027.parquet`: `32047` rows, `2965673` file bytes (2.83 MiB), `20472012` physical bytes (19.52 MiB), `15931301` encoded bytes (15.19 MiB), `2936001` compressed data bytes (2.80 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00028.parquet`: `32032` rows, `2972134` file bytes (2.83 MiB), `20272522` physical bytes (19.33 MiB), `15939762` encoded bytes (15.20 MiB), `2942767` compressed data bytes (2.81 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict/part-00029.parquet`: `24786` rows, `2356709` file bytes (2.25 MiB), `15964893` physical bytes (15.23 MiB), `12544815` encoded bytes (11.96 MiB), `2327443` compressed data bytes (2.22 MiB)
