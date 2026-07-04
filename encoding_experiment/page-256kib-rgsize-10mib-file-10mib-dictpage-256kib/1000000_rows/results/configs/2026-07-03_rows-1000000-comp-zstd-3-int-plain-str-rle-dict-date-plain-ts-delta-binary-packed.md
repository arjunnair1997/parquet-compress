# ClickBench Parquet Experiment

- Started: `2026-07-03T23:35:31-04:00`
- Write elapsed: `11.336s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `461624481` (440.24 MiB)
- Compressed column data bytes after codec compression: `76137222` (72.61 MiB)
- Parquet file bytes: `77024147` (73.46 MiB)
- Physical/encoded ratio: `1.543x`
- Encoded/compressed-data ratio: `6.063x`
- Physical/compressed-data ratio: `9.357x`
- Physical/parquet-file ratio: `9.249x`
- Files: `30`

## Settings

- Compression: `zstd-3`
- Int encoding: `plain`
- String encoding: `rle-dict`
- Date encoding: `plain`
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
- Files read: `30`
- Elapsed: `7.061s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004717` (7.63 MiB) | `8005551` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003717` (3.82 MiB) | `64674` (63.16 KiB) | `0.999x` | `61.906x` | `61.849x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `138409995` (132.00 MiB) | `28808809` (27.47 MiB) | `7970118` (7.60 MiB) | `4.804x` | `3.615x` | `17.366x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003655` (3.82 MiB) | `5096` (4.98 KiB) | `0.999x` | `785.647x` | `784.929x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:59` | `1000000` | `8000000` (7.63 MiB) | `3369402` (3.21 MiB) | `2887044` (2.75 MiB) | `2.374x` | `1.167x` | `2.771x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003657` (3.82 MiB) | `5098` (4.98 KiB) | `0.999x` | `785.339x` | `784.621x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003656` (3.82 MiB) | `5114` (4.99 KiB) | `0.999x` | `782.882x` | `782.167x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003718` (3.82 MiB) | `408395` (398.82 KiB) | `0.999x` | `9.804x` | `9.794x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `191103` (186.62 KiB) | `0.999x` | `20.951x` | `20.931x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004716` (7.63 MiB) | `618399` (603.91 KiB) | `0.999x` | `12.944x` | `12.937x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003656` (3.82 MiB) | `4389` (4.29 KiB) | `0.999x` | `912.202x` | `911.369x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `107883` (105.35 KiB) | `0.999x` | `37.112x` | `37.077x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `134663` (131.51 KiB) | `0.999x` | `29.731x` | `29.704x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `88562192` (84.46 MiB) | `44074155` (42.03 MiB) | `12701645` (12.11 MiB) | `2.009x` | `3.470x` | `6.972x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `79583339` (75.90 MiB) | `34391703` (32.80 MiB) | `11728316` (11.18 MiB) | `2.314x` | `2.932x` | `6.786x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `179044` (174.85 KiB) | `0.999x` | `22.362x` | `22.341x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `275434` (268.98 KiB) | `0.999x` | `14.536x` | `14.523x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003705` (3.82 MiB) | `231561` (226.13 KiB) | `0.999x` | `17.290x` | `17.274x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `87808` (85.75 KiB) | `0.999x` | `45.596x` | `45.554x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `48416` (47.28 KiB) | `0.999x` | `82.694x` | `82.617x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003711` (3.82 MiB) | `187720` (183.32 KiB) | `0.999x` | `21.328x` | `21.308x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003711` (3.82 MiB) | `186613` (182.24 KiB) | `0.999x` | `21.455x` | `21.435x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `82153` (80.23 KiB) | `0.999x` | `48.735x` | `48.690x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `53814` (52.55 KiB) | `0.999x` | `74.399x` | `74.330x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `130084` (127.04 KiB) | `0.999x` | `30.778x` | `30.749x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `3354477` (3.20 MiB) | `266488` (260.24 KiB) | `145457` (142.05 KiB) | `12.588x` | `1.832x` | `23.062x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `26556` (25.93 KiB) | `0.999x` | `150.765x` | `150.625x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `25258` (24.67 KiB) | `0.999x` | `158.513x` | `158.366x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003710` (3.82 MiB) | `154871` (151.24 KiB) | `0.999x` | `25.852x` | `25.828x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `3767530` (3.59 MiB) | `150508` (146.98 KiB) | `84730` (82.74 KiB) | `25.032x` | `1.776x` | `44.465x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003662` (3.82 MiB) | `6257` (6.11 KiB) | `0.999x` | `639.869x` | `639.284x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003667` (3.82 MiB) | `6706` (6.55 KiB) | `0.999x` | `597.028x` | `596.481x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `29409` (28.72 KiB) | `0.999x` | `136.139x` | `136.013x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `22768` (22.23 KiB) | `0.999x` | `175.848x` | `175.685x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `81583` (79.67 KiB) | `25139` (24.55 KiB) | `20143` (19.67 KiB) | `3.245x` | `1.248x` | `4.050x` | `81583` (79.67 KiB) |
| `Params` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `0` (0 B) | `4189` (4.09 KiB) | `5251` (5.13 KiB) | `0.000x` | `0.798x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `323967` (316.37 KiB) | `0.999x` | `12.358x` | `12.347x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003711` (3.82 MiB) | `288561` (281.80 KiB) | `0.999x` | `13.875x` | `13.862x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003717` (3.82 MiB) | `101892` (99.50 KiB) | `0.999x` | `39.294x` | `39.257x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `3528017` (3.36 MiB) | `1642647` (1.57 MiB) | `636987` (622.06 KiB) | `2.148x` | `2.579x` | `5.539x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003703` (3.82 MiB) | `30937` (30.21 KiB) | `0.999x` | `129.415x` | `129.295x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003710` (3.82 MiB) | `165195` (161.32 KiB) | `0.999x` | `24.236x` | `24.214x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `306175` (299.00 KiB) | `0.999x` | `13.077x` | `13.064x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `319637` (312.15 KiB) | `0.999x` | `12.526x` | `12.514x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `100296` (97.95 KiB) | `0.999x` | `39.919x` | `39.882x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:59` | `1000000` | `8000000` (7.63 MiB) | `3492113` (3.33 MiB) | `2898410` (2.76 MiB) | `2.291x` | `1.205x` | `2.760x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `90290` (88.17 KiB) | `0.999x` | `44.343x` | `44.302x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `74270` (72.53 KiB) | `0.999x` | `53.908x` | `53.858x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `124245` (121.33 KiB) | `0.999x` | `32.224x` | `32.194x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003654` (3.82 MiB) | `4515` (4.41 KiB) | `0.999x` | `886.745x` | `885.936x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `13587860` (12.96 MiB) | `8838` (8.63 KiB) | `9902` (9.67 KiB) | `1537.436x` | `0.893x` | `1372.234x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003675` (3.82 MiB) | `7242` (7.07 KiB) | `0.999x` | `552.841x` | `552.334x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003687` (3.82 MiB) | `56987` (55.65 KiB) | `0.999x` | `70.256x` | `70.191x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003682` (3.82 MiB) | `8227` (8.03 KiB) | `0.999x` | `486.652x` | `486.204x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003659` (3.82 MiB) | `25359` (24.76 KiB) | `0.999x` | `157.879x` | `157.735x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004711` (7.63 MiB) | `694163` (677.89 KiB) | `0.999x` | `11.531x` | `11.525x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `27797671` (26.51 MiB) | `21269791` (20.28 MiB) | `4882629` (4.66 MiB) | `1.307x` | `4.356x` | `5.693x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `3792143` (3.62 MiB) | `0.999x` | `1.056x` | `1.055x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003656` (3.82 MiB) | `4389` (4.29 KiB) | `0.999x` | `912.202x` | `911.369x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003656` (3.82 MiB) | `4389` (4.29 KiB) | `0.999x` | `912.202x` | `911.369x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003656` (3.82 MiB) | `4389` (4.29 KiB) | `0.999x` | `912.202x` | `911.369x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `82776` (80.84 KiB) | `0.999x` | `48.368x` | `48.323x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003656` (3.82 MiB) | `4389` (4.29 KiB) | `0.999x` | `912.202x` | `911.369x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `1000000` (976.56 KiB) | `33611` (32.82 KiB) | `27691` (27.04 KiB) | `29.752x` | `1.214x` | `36.113x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:59` | `1000000` | `8000000` (7.63 MiB) | `3383208` (3.23 MiB) | `2897955` (2.76 MiB) | `2.365x` | `1.167x` | `2.761x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `142946` (139.60 KiB) | `0.999x` | `28.009x` | `27.983x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `107776` (105.25 KiB) | `0.999x` | `37.148x` | `37.114x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `122773` (119.90 KiB) | `0.999x` | `32.611x` | `32.580x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `193822` (189.28 KiB) | `0.999x` | `20.657x` | `20.637x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `174111` (170.03 KiB) | `0.999x` | `22.995x` | `22.974x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003726` (3.82 MiB) | `426803` (416.80 KiB) | `0.999x` | `9.381x` | `9.372x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003708` (3.82 MiB) | `71067` (69.40 KiB) | `0.999x` | `56.337x` | `56.285x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003656` (3.82 MiB) | `4389` (4.29 KiB) | `0.999x` | `912.202x` | `911.369x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003701` (3.82 MiB) | `55501` (54.20 KiB) | `0.999x` | `72.137x` | `72.071x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `2001192` (1.91 MiB) | `42600` (41.60 KiB) | `27522` (26.88 KiB) | `46.976x` | `1.548x` | `72.712x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `3325142` (3.17 MiB) | `136700` (133.50 KiB) | `69523` (67.89 KiB) | `24.324x` | `1.966x` | `47.828x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `0` (0 B) | `4189` (4.09 KiB) | `5251` (5.13 KiB) | `0.000x` | `0.798x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `0` (0 B) | `4189` (4.09 KiB) | `5251` (5.13 KiB) | `0.000x` | `0.798x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003656` (3.82 MiB) | `4389` (4.29 KiB) | `0.999x` | `912.202x` | `911.369x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003672` (3.82 MiB) | `61817` (60.37 KiB) | `0.999x` | `64.767x` | `64.707x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `135613` (132.43 KiB) | `0.999x` | `29.523x` | `29.496x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003724` (3.82 MiB) | `333520` (325.70 KiB) | `0.999x` | `12.004x` | `11.993x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `1246254` (1.19 MiB) | `0.999x` | `3.213x` | `3.210x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003770` (3.82 MiB) | `937965` (915.98 KiB) | `0.999x` | `4.269x` | `4.265x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003754` (3.82 MiB) | `550441` (537.54 KiB) | `0.999x` | `7.274x` | `7.267x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003657` (3.82 MiB) | `5488` (5.36 KiB) | `0.999x` | `729.529x` | `728.863x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `1024` (1.00 KiB) | `5860` (5.72 KiB) | `6922` (6.76 KiB) | `0.175x` | `0.847x` | `0.148x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004599` (7.63 MiB) | `5912` (5.77 KiB) | `0.999x` | `1353.958x` | `1353.180x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `0` (0 B) | `4189` (4.09 KiB) | `5251` (5.13 KiB) | `0.000x` | `0.798x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `3000000` (2.86 MiB) | `5074` (4.96 KiB) | `6136` (5.99 KiB) | `591.250x` | `0.827x` | `488.918x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003656` (3.82 MiB) | `4389` (4.29 KiB) | `0.999x` | `912.202x` | `911.369x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `58030` (56.67 KiB) | `18472` (18.04 KiB) | `17512` (17.10 KiB) | `3.142x` | `1.055x` | `3.314x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `22051` (21.53 KiB) | `16312` (15.93 KiB) | `15892` (15.52 KiB) | `1.352x` | `1.026x` | `1.388x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `25445` (24.85 KiB) | `22844` (22.31 KiB) | `18461` (18.03 KiB) | `1.114x` | `1.237x` | `1.378x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `48191` (47.06 KiB) | `14693` (14.35 KiB) | `13087` (12.78 KiB) | `3.280x` | `1.123x` | `3.682x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `49433` (48.27 KiB) | `22647` (22.12 KiB) | `18356` (17.93 KiB) | `2.183x` | `1.234x` | `2.693x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `16873` (16.48 KiB) | `16585` (16.20 KiB) | `14712` (14.37 KiB) | `1.017x` | `1.127x` | `1.147x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `91870` (89.72 KiB) | `37204` (36.33 KiB) | `27288` (26.65 KiB) | `2.469x` | `1.363x` | `3.367x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `13001` (12.70 KiB) | `14752` (14.41 KiB) | `14839` (14.49 KiB) | `0.881x` | `0.994x` | `0.876x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `28101` (27.44 KiB) | `16508` (16.12 KiB) | `16077` (15.70 KiB) | `1.702x` | `1.027x` | `1.748x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `45607` (44.54 KiB) | `41274` (40.31 KiB) | `22617` (22.09 KiB) | `1.105x` | `1.825x` | `2.016x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003685` (3.82 MiB) | `21321` (20.82 KiB) | `0.999x` | `187.781x` | `187.608x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004713` (7.63 MiB) | `2843246` (2.71 MiB) | `0.999x` | `2.815x` | `2.814x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004713` (7.63 MiB) | `3581622` (3.42 MiB) | `0.999x` | `2.235x` | `2.234x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003664` (3.82 MiB) | `5812` (5.68 KiB) | `0.999x` | `688.862x` | `688.231x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00000.parquet`: `35743` rows, `2314454` file bytes (2.21 MiB), `27762714` physical bytes (26.48 MiB), `15614681` encoded bytes (14.89 MiB), `2284073` compressed data bytes (2.18 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00001.parquet`: `35886` rows, `2324303` file bytes (2.22 MiB), `27565035` physical bytes (26.29 MiB), `15559767` encoded bytes (14.84 MiB), `2293758` compressed data bytes (2.19 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00002.parquet`: `35701` rows, `2332048` file bytes (2.22 MiB), `27970399` physical bytes (26.67 MiB), `15608579` encoded bytes (14.89 MiB), `2300861` compressed data bytes (2.19 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00003.parquet`: `35833` rows, `2315863` file bytes (2.21 MiB), `27757206` physical bytes (26.47 MiB), `15580807` encoded bytes (14.86 MiB), `2285555` compressed data bytes (2.18 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00004.parquet`: `35905` rows, `2308248` file bytes (2.20 MiB), `27900240` physical bytes (26.61 MiB), `15550955` encoded bytes (14.83 MiB), `2277074` compressed data bytes (2.17 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00005.parquet`: `36117` rows, `2295274` file bytes (2.19 MiB), `27692706` physical bytes (26.41 MiB), `15566628` encoded bytes (14.85 MiB), `2264774` compressed data bytes (2.16 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00006.parquet`: `35924` rows, `2317576` file bytes (2.21 MiB), `27808620` physical bytes (26.52 MiB), `15570218` encoded bytes (14.85 MiB), `2287245` compressed data bytes (2.18 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00007.parquet`: `35846` rows, `2325137` file bytes (2.22 MiB), `27730397` physical bytes (26.45 MiB), `15565620` encoded bytes (14.84 MiB), `2294366` compressed data bytes (2.19 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00008.parquet`: `36007` rows, `2288656` file bytes (2.18 MiB), `27905250` physical bytes (26.61 MiB), `15551327` encoded bytes (14.83 MiB), `2258003` compressed data bytes (2.15 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00009.parquet`: `35866` rows, `2323623` file bytes (2.22 MiB), `27684833` physical bytes (26.40 MiB), `15582774` encoded bytes (14.86 MiB), `2293395` compressed data bytes (2.19 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00010.parquet`: `36025` rows, `2279665` file bytes (2.17 MiB), `27704740` physical bytes (26.42 MiB), `15569965` encoded bytes (14.85 MiB), `2249209` compressed data bytes (2.15 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00011.parquet`: `36169` rows, `2272690` file bytes (2.17 MiB), `27897850` physical bytes (26.61 MiB), `15560695` encoded bytes (14.84 MiB), `2242296` compressed data bytes (2.14 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00012.parquet`: `36109` rows, `2283737` file bytes (2.18 MiB), `27670527` physical bytes (26.39 MiB), `15599466` encoded bytes (14.88 MiB), `2253480` compressed data bytes (2.15 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00013.parquet`: `35918` rows, `2293618` file bytes (2.19 MiB), `27789935` physical bytes (26.50 MiB), `15582407` encoded bytes (14.86 MiB), `2262983` compressed data bytes (2.16 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00014.parquet`: `32171` rows, `2702850` file bytes (2.58 MiB), `22475774` physical bytes (21.43 MiB), `15820753` encoded bytes (15.09 MiB), `2673163` compressed data bytes (2.55 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00015.parquet`: `31549` rows, `2727674` file bytes (2.60 MiB), `21821478` physical bytes (20.81 MiB), `15905157` encoded bytes (15.17 MiB), `2698555` compressed data bytes (2.57 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00016.parquet`: `31363` rows, `2816078` file bytes (2.69 MiB), `21491123` physical bytes (20.50 MiB), `15914591` encoded bytes (15.18 MiB), `2786984` compressed data bytes (2.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00017.parquet`: `32362` rows, `2921637` file bytes (2.79 MiB), `20662724` physical bytes (19.71 MiB), `15701474` encoded bytes (14.97 MiB), `2892273` compressed data bytes (2.76 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00018.parquet`: `31905` rows, `2932109` file bytes (2.80 MiB), `20740717` physical bytes (19.78 MiB), `15759249` encoded bytes (15.03 MiB), `2902571` compressed data bytes (2.77 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00019.parquet`: `32014` rows, `2958442` file bytes (2.82 MiB), `20561622` physical bytes (19.61 MiB), `15702737` encoded bytes (14.98 MiB), `2928663` compressed data bytes (2.79 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00020.parquet`: `32234` rows, `2937293` file bytes (2.80 MiB), `20586707` physical bytes (19.63 MiB), `15739761` encoded bytes (15.01 MiB), `2907604` compressed data bytes (2.77 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00021.parquet`: `31912` rows, `2953815` file bytes (2.82 MiB), `20622327` physical bytes (19.67 MiB), `15724555` encoded bytes (15.00 MiB), `2924282` compressed data bytes (2.79 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00022.parquet`: `32234` rows, `2912064` file bytes (2.78 MiB), `20752678` physical bytes (19.79 MiB), `15736543` encoded bytes (15.01 MiB), `2882221` compressed data bytes (2.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00023.parquet`: `32308` rows, `2916310` file bytes (2.78 MiB), `20668205` physical bytes (19.71 MiB), `15699867` encoded bytes (14.97 MiB), `2887176` compressed data bytes (2.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00024.parquet`: `31812` rows, `2952463` file bytes (2.82 MiB), `20603679` physical bytes (19.65 MiB), `15763453` encoded bytes (15.03 MiB), `2922684` compressed data bytes (2.79 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00025.parquet`: `32677` rows, `2903366` file bytes (2.77 MiB), `20717195` physical bytes (19.76 MiB), `15651332` encoded bytes (14.93 MiB), `2874163` compressed data bytes (2.74 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00026.parquet`: `32419` rows, `2931031` file bytes (2.80 MiB), `20690167` physical bytes (19.73 MiB), `15694900` encoded bytes (14.97 MiB), `2901119` compressed data bytes (2.77 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00027.parquet`: `32521` rows, `2910475` file bytes (2.78 MiB), `20794108` physical bytes (19.83 MiB), `15677604` encoded bytes (14.95 MiB), `2880709` compressed data bytes (2.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00028.parquet`: `32636` rows, `2908472` file bytes (2.77 MiB), `20692602` physical bytes (19.73 MiB), `15705453` encoded bytes (14.98 MiB), `2879175` compressed data bytes (2.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed/part-00029.parquet`: `14834` rows, `1365176` file bytes (1.30 MiB), `9677066` physical bytes (9.23 MiB), `7363163` encoded bytes (7.02 MiB), `1348808` compressed data bytes (1.29 MiB)
