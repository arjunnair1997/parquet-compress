# ClickBench Parquet Experiment

- Started: `2026-07-03T23:36:36-04:00`
- Write elapsed: `11.723s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `715139897` (682.01 MiB)
- Compressed column data bytes after codec compression: `93993752` (89.64 MiB)
- Parquet file bytes: `94906889` (90.51 MiB)
- Physical/encoded ratio: `0.996x`
- Encoded/compressed-data ratio: `7.608x`
- Physical/compressed-data ratio: `7.579x`
- Physical/parquet-file ratio: `7.506x`
- Files: `29`

## Settings

- Compression: `zstd-3`
- Int encoding: `plain`
- String encoding: `delta-length-byte-array`
- Date encoding: `delta-binary-packed`
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
- Elapsed: `7.29s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `8004636` (7.63 MiB) | `8005441` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003648` (3.82 MiB) | `64832` (63.31 KiB) | `0.999x` | `61.754x` | `61.698x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:520` | `1000000` | `138409995` (132.00 MiB) | `140027539` (133.54 MiB) | `14467840` (13.80 MiB) | `0.988x` | `9.679x` | `9.567x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `4990` (4.87 KiB) | `0.999x` | `802.322x` | `801.603x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:58, DICTIONARY_PAGE/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `7355779` (7.02 MiB) | `4037844` (3.85 MiB) | `1.088x` | `1.822x` | `1.981x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:58` | `1000000` | `4000000` (3.81 MiB) | `50869` (49.68 KiB) | `6420` (6.27 KiB) | `78.633x` | `7.924x` | `623.053x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `5008` (4.89 KiB) | `0.999x` | `799.439x` | `798.722x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003659` (3.82 MiB) | `408423` (398.85 KiB) | `0.999x` | `9.803x` | `9.794x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003649` (3.82 MiB) | `190999` (186.52 KiB) | `0.999x` | `20.962x` | `20.943x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `8004630` (7.63 MiB) | `618192` (603.70 KiB) | `0.999x` | `12.948x` | `12.941x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `4294` (4.19 KiB) | `0.999x` | `932.368x` | `931.532x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003650` (3.82 MiB) | `107017` (104.51 KiB) | `0.999x` | `37.411x` | `37.377x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003642` (3.82 MiB) | `135386` (132.21 KiB) | `0.999x` | `29.572x` | `29.545x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:346` | `1000000` | `88562192` (84.46 MiB) | `89785373` (85.63 MiB) | `16034342` (15.29 MiB) | `0.986x` | `5.600x` | `5.523x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:326` | `1000000` | `79583339` (75.90 MiB) | `80834242` (77.09 MiB) | `15009209` (14.31 MiB) | `0.985x` | `5.386x` | `5.302x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003649` (3.82 MiB) | `178809` (174.62 KiB) | `0.999x` | `22.391x` | `22.370x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003650` (3.82 MiB) | `275459` (269.00 KiB) | `0.999x` | `14.534x` | `14.521x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003648` (3.82 MiB) | `231379` (225.96 KiB) | `0.999x` | `17.303x` | `17.288x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003649` (3.82 MiB) | `87717` (85.66 KiB) | `0.999x` | `45.643x` | `45.601x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003648` (3.82 MiB) | `48190` (47.06 KiB) | `0.999x` | `83.080x` | `83.005x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003647` (3.82 MiB) | `187365` (182.97 KiB) | `0.999x` | `21.368x` | `21.349x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003650` (3.82 MiB) | `186320` (181.95 KiB) | `0.999x` | `21.488x` | `21.468x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003645` (3.82 MiB) | `82361` (80.43 KiB) | `0.999x` | `48.611x` | `48.567x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003646` (3.82 MiB) | `53289` (52.04 KiB) | `0.999x` | `75.131x` | `75.062x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003650` (3.82 MiB) | `129482` (126.45 KiB) | `0.999x` | `30.921x` | `30.892x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `3354477` (3.20 MiB) | `3707337` (3.54 MiB) | `267059` (260.80 KiB) | `0.905x` | `13.882x` | `12.561x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003647` (3.82 MiB) | `26688` (26.06 KiB) | `0.999x` | `150.017x` | `149.880x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003648` (3.82 MiB) | `25276` (24.68 KiB) | `0.999x` | `158.397x` | `158.253x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003645` (3.82 MiB) | `155893` (152.24 KiB) | `0.999x` | `25.682x` | `25.659x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `3767530` (3.59 MiB) | `4017116` (3.83 MiB) | `151166` (147.62 KiB) | `0.938x` | `26.574x` | `24.923x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003597` (3.82 MiB) | `6167` (6.02 KiB) | `0.999x` | `649.197x` | `648.614x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003601` (3.82 MiB) | `6650` (6.49 KiB) | `0.999x` | `602.045x` | `601.504x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003644` (3.82 MiB) | `29115` (28.43 KiB) | `0.999x` | `137.511x` | `137.386x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003647` (3.82 MiB) | `22808` (22.27 KiB) | `0.999x` | `175.537x` | `175.377x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `81583` (79.67 KiB) | `231047` (225.63 KiB) | `30833` (30.11 KiB) | `0.353x` | `7.493x` | `2.646x` | `81583` (79.67 KiB) |
| `Params` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `0` (0 B) | `41739` (40.76 KiB) | `3623` (3.54 KiB) | `0.000x` | `11.521x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003649` (3.82 MiB) | `323808` (316.22 KiB) | `0.999x` | `12.364x` | `12.353x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003651` (3.82 MiB) | `288480` (281.72 KiB) | `0.999x` | `13.878x` | `13.866x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003645` (3.82 MiB) | `101857` (99.47 KiB) | `0.999x` | `39.307x` | `39.271x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `3528017` (3.36 MiB) | `4222019` (4.03 MiB) | `812367` (793.33 KiB) | `0.836x` | `5.197x` | `4.343x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003643` (3.82 MiB) | `30881` (30.16 KiB) | `0.999x` | `129.647x` | `129.529x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003648` (3.82 MiB) | `165163` (161.29 KiB) | `0.999x` | `24.241x` | `24.218x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003647` (3.82 MiB) | `305932` (298.76 KiB) | `0.999x` | `13.087x` | `13.075x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003650` (3.82 MiB) | `319342` (311.86 KiB) | `0.999x` | `12.537x` | `12.526x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003646` (3.82 MiB) | `100104` (97.76 KiB) | `0.999x` | `39.995x` | `39.958x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:58, DICTIONARY_PAGE/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `7209702` (6.88 MiB) | `3977999` (3.79 MiB) | `1.110x` | `1.812x` | `2.011x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003648` (3.82 MiB) | `89177` (87.09 KiB) | `0.999x` | `44.896x` | `44.855x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003646` (3.82 MiB) | `73776` (72.05 KiB) | `0.999x` | `54.268x` | `54.218x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003646` (3.82 MiB) | `123971` (121.07 KiB) | `0.999x` | `32.295x` | `32.266x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `4424` (4.32 KiB) | `0.999x` | `904.970x` | `904.159x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:86` | `1000000` | `13587860` (12.96 MiB) | `13657397` (13.02 MiB) | `18949` (18.50 KiB) | `0.995x` | `720.745x` | `717.075x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003609` (3.82 MiB) | `7159` (6.99 KiB) | `0.999x` | `559.241x` | `558.737x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003619` (3.82 MiB) | `56763` (55.43 KiB) | `0.999x` | `70.532x` | `70.468x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003609` (3.82 MiB) | `8158` (7.97 KiB) | `0.999x` | `490.759x` | `490.316x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003597` (3.82 MiB) | `25336` (24.74 KiB) | `0.999x` | `158.020x` | `157.878x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `8004635` (7.63 MiB) | `694095` (677.83 KiB) | `0.999x` | `11.532x` | `11.526x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:136` | `1000000` | `27797671` (26.51 MiB) | `28791599` (27.46 MiB) | `5612911` (5.35 MiB) | `0.965x` | `5.130x` | `4.952x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003704` (3.82 MiB) | `3848754` (3.67 MiB) | `0.999x` | `1.040x` | `1.039x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `4294` (4.19 KiB) | `0.999x` | `932.368x` | `931.532x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `4294` (4.19 KiB) | `0.999x` | `932.368x` | `931.532x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `4294` (4.19 KiB) | `0.999x` | `932.368x` | `931.532x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003650` (3.82 MiB) | `82515` (80.58 KiB) | `0.999x` | `48.520x` | `48.476x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `4294` (4.19 KiB) | `0.999x` | `932.368x` | `931.532x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `1000000` (976.56 KiB) | `1043084` (1018.64 KiB) | `28214` (27.55 KiB) | `0.959x` | `36.970x` | `35.443x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:58, DICTIONARY_PAGE/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `7351272` (7.01 MiB) | `4039334` (3.85 MiB) | `1.088x` | `1.820x` | `1.981x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003650` (3.82 MiB) | `142674` (139.33 KiB) | `0.999x` | `28.062x` | `28.036x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003645` (3.82 MiB) | `107618` (105.10 KiB) | `0.999x` | `37.202x` | `37.169x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003647` (3.82 MiB) | `122893` (120.01 KiB) | `0.999x` | `32.578x` | `32.549x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003640` (3.82 MiB) | `193643` (189.10 KiB) | `0.999x` | `20.675x` | `20.657x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003647` (3.82 MiB) | `174140` (170.06 KiB) | `0.999x` | `22.991x` | `22.970x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003663` (3.82 MiB) | `426492` (416.50 KiB) | `0.999x` | `9.387x` | `9.379x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003638` (3.82 MiB) | `70885` (69.22 KiB) | `0.999x` | `56.481x` | `56.429x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `4296` (4.20 KiB) | `0.999x` | `931.934x` | `931.099x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003632` (3.82 MiB) | `55412` (54.11 KiB) | `0.999x` | `72.252x` | `72.187x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `2001192` (1.91 MiB) | `2051147` (1.96 MiB) | `32230` (31.47 KiB) | `0.976x` | `63.641x` | `62.091x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `3325142` (3.17 MiB) | `3637930` (3.47 MiB) | `185208` (180.87 KiB) | `0.914x` | `19.642x` | `17.954x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `0` (0 B) | `41739` (40.76 KiB) | `3623` (3.54 KiB) | `0.000x` | `11.521x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `0` (0 B) | `41739` (40.76 KiB) | `3623` (3.54 KiB) | `0.000x` | `11.521x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `4294` (4.19 KiB) | `0.999x` | `932.368x` | `931.532x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003601` (3.82 MiB) | `61653` (60.21 KiB) | `0.999x` | `64.938x` | `64.879x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003645` (3.82 MiB) | `135244` (132.07 KiB) | `0.999x` | `29.603x` | `29.576x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003660` (3.82 MiB) | `333011` (325.21 KiB) | `0.999x` | `12.023x` | `12.012x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003706` (3.82 MiB) | `1246621` (1.19 MiB) | `0.999x` | `3.212x` | `3.209x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003708` (3.82 MiB) | `938094` (916.11 KiB) | `0.999x` | `4.268x` | `4.264x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003686` (3.82 MiB) | `550311` (537.41 KiB) | `0.999x` | `7.275x` | `7.269x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `5395` (5.27 KiB) | `0.999x` | `742.092x` | `741.427x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `1024` (1.00 KiB) | `46358` (45.27 KiB) | `6145` (6.00 KiB) | `0.022x` | `7.544x` | `0.167x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `8004518` (7.63 MiB) | `5750` (5.62 KiB) | `0.999x` | `1392.090x` | `1391.304x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `0` (0 B) | `41739` (40.76 KiB) | `3623` (3.54 KiB) | `0.000x` | `11.521x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `3000000` (2.86 MiB) | `3044033` (2.90 MiB) | `5878` (5.74 KiB) | `0.986x` | `517.869x` | `510.378x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `4294` (4.19 KiB) | `0.999x` | `932.368x` | `931.532x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `58030` (56.67 KiB) | `217671` (212.57 KiB) | `27497` (26.85 KiB) | `0.267x` | `7.916x` | `2.110x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `22051` (21.53 KiB) | `123110` (120.22 KiB) | `23279` (22.73 KiB) | `0.179x` | `5.288x` | `0.947x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `25445` (24.85 KiB) | `130898` (127.83 KiB) | `28592` (27.92 KiB) | `0.194x` | `4.578x` | `0.890x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `48191` (47.06 KiB) | `156136` (152.48 KiB) | `18327` (17.90 KiB) | `0.309x` | `8.519x` | `2.630x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `49433` (48.27 KiB) | `188306` (183.89 KiB) | `33994` (33.20 KiB) | `0.263x` | `5.539x` | `1.454x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `16873` (16.48 KiB) | `133591` (130.46 KiB) | `21981` (21.47 KiB) | `0.126x` | `6.078x` | `0.768x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `91870` (89.72 KiB) | `253413` (247.47 KiB) | `47710` (46.59 KiB) | `0.363x` | `5.312x` | `1.926x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `13001` (12.70 KiB) | `94580` (92.36 KiB) | `18217` (17.79 KiB) | `0.137x` | `5.192x` | `0.714x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `28101` (27.44 KiB) | `129074` (126.05 KiB) | `19968` (19.50 KiB) | `0.218x` | `6.464x` | `1.407x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `45607` (44.54 KiB) | `211114` (206.17 KiB) | `36431` (35.58 KiB) | `0.216x` | `5.795x` | `1.252x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003617` (3.82 MiB) | `21279` (20.78 KiB) | `0.999x` | `188.149x` | `187.979x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `8004634` (7.63 MiB) | `2843274` (2.71 MiB) | `0.999x` | `2.815x` | `2.814x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `8004634` (7.63 MiB) | `3581945` (3.42 MiB) | `0.999x` | `2.235x` | `2.233x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003605` (3.82 MiB) | `5677` (5.54 KiB) | `0.999x` | `705.233x` | `704.597x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00000.parquet`: `35117` rows, `3065590` file bytes (2.92 MiB), `27262167` physical bytes (26.00 MiB), `27219215` encoded bytes (25.96 MiB), `3033565` compressed data bytes (2.89 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00001.parquet`: `35138` rows, `3016108` file bytes (2.88 MiB), `27043411` physical bytes (25.79 MiB), `26996495` encoded bytes (25.75 MiB), `2983727` compressed data bytes (2.85 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00002.parquet`: `35625` rows, `3118530` file bytes (2.97 MiB), `27849002` physical bytes (26.56 MiB), `27803439` encoded bytes (26.52 MiB), `3085323` compressed data bytes (2.94 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00003.parquet`: `35468` rows, `3032834` file bytes (2.89 MiB), `27454380` physical bytes (26.18 MiB), `27401502` encoded bytes (26.13 MiB), `3000501` compressed data bytes (2.86 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00004.parquet`: `35336` rows, `3036143` file bytes (2.90 MiB), `27507139` physical bytes (26.23 MiB), `27454665` encoded bytes (26.18 MiB), `3003421` compressed data bytes (2.86 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00005.parquet`: `35456` rows, `3013752` file bytes (2.87 MiB), `27256358` physical bytes (25.99 MiB), `27208179` encoded bytes (25.95 MiB), `2981294` compressed data bytes (2.84 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00006.parquet`: `35614` rows, `3039890` file bytes (2.90 MiB), `27458009` physical bytes (26.19 MiB), `27411057` encoded bytes (26.14 MiB), `3007331` compressed data bytes (2.87 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00007.parquet`: `35674` rows, `3078557` file bytes (2.94 MiB), `27658707` physical bytes (26.38 MiB), `27615440` encoded bytes (26.34 MiB), `3045824` compressed data bytes (2.90 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00008.parquet`: `35939` rows, `3049591` file bytes (2.91 MiB), `27884057` physical bytes (26.59 MiB), `27831403` encoded bytes (26.54 MiB), `3016909` compressed data bytes (2.88 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00009.parquet`: `35766` rows, `3038956` file bytes (2.90 MiB), `27447046` physical bytes (26.18 MiB), `27395002` encoded bytes (26.13 MiB), `3006733` compressed data bytes (2.87 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00010.parquet`: `36082` rows, `3062151` file bytes (2.92 MiB), `27838417` physical bytes (26.55 MiB), `27789682` encoded bytes (26.50 MiB), `3029565` compressed data bytes (2.89 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00011.parquet`: `36143` rows, `3040555` file bytes (2.90 MiB), `27893777` physical bytes (26.60 MiB), `27842869` encoded bytes (26.55 MiB), `3007877` compressed data bytes (2.87 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00012.parquet`: `35742` rows, `3048423` file bytes (2.91 MiB), `27470090` physical bytes (26.20 MiB), `27422534` encoded bytes (26.15 MiB), `3016259` compressed data bytes (2.88 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00013.parquet`: `35632` rows, `3007589` file bytes (2.87 MiB), `27493230` physical bytes (26.22 MiB), `27440954` encoded bytes (26.17 MiB), `2974959` compressed data bytes (2.84 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00014.parquet`: `35263` rows, `3377723` file bytes (3.22 MiB), `24957101` physical bytes (23.80 MiB), `25060229` encoded bytes (23.90 MiB), `3346505` compressed data bytes (3.19 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00015.parquet`: `35006` rows, `3494275` file bytes (3.33 MiB), `24262326` physical bytes (23.14 MiB), `24411077` encoded bytes (23.28 MiB), `3463966` compressed data bytes (3.30 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00016.parquet`: `34197` rows, `3494469` file bytes (3.33 MiB), `23062256` physical bytes (21.99 MiB), `23273626` encoded bytes (22.20 MiB), `3464101` compressed data bytes (3.30 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00017.parquet`: `34273` rows, `3625148` file bytes (3.46 MiB), `22048374` physical bytes (21.03 MiB), `22305303` encoded bytes (21.27 MiB), `3594555` compressed data bytes (3.43 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00018.parquet`: `34137` rows, `3705475` file bytes (3.53 MiB), `22257118` physical bytes (21.23 MiB), `22508154` encoded bytes (21.47 MiB), `3674787` compressed data bytes (3.50 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00019.parquet`: `34187` rows, `3618396` file bytes (3.45 MiB), `21826156` physical bytes (20.82 MiB), `22081094` encoded bytes (21.06 MiB), `3587744` compressed data bytes (3.42 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00020.parquet`: `34214` rows, `3654922` file bytes (3.49 MiB), `22021784` physical bytes (21.00 MiB), `22282317` encoded bytes (21.25 MiB), `3624301` compressed data bytes (3.46 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00021.parquet`: `34411` rows, `3593736` file bytes (3.43 MiB), `22035243` physical bytes (21.01 MiB), `22294618` encoded bytes (21.26 MiB), `3562874` compressed data bytes (3.40 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00022.parquet`: `34338` rows, `3633495` file bytes (3.47 MiB), `22016616` physical bytes (21.00 MiB), `22269437` encoded bytes (21.24 MiB), `3603035` compressed data bytes (3.44 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00023.parquet`: `34100` rows, `3626751` file bytes (3.46 MiB), `22062086` physical bytes (21.04 MiB), `22315558` encoded bytes (21.28 MiB), `3596013` compressed data bytes (3.43 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00024.parquet`: `34030` rows, `3621385` file bytes (3.45 MiB), `21830113` physical bytes (20.82 MiB), `22082231` encoded bytes (21.06 MiB), `3591153` compressed data bytes (3.42 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00025.parquet`: `34401` rows, `3580635` file bytes (3.41 MiB), `21885263` physical bytes (20.87 MiB), `22143218` encoded bytes (21.12 MiB), `3549900` compressed data bytes (3.39 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00026.parquet`: `34318` rows, `3540867` file bytes (3.38 MiB), `21893251` physical bytes (20.88 MiB), `22145026` encoded bytes (21.12 MiB), `3510266` compressed data bytes (3.35 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00027.parquet`: `34290` rows, `3547799` file bytes (3.38 MiB), `21704100` physical bytes (20.70 MiB), `21959448` encoded bytes (20.94 MiB), `3517560` compressed data bytes (3.35 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict/part-00028.parquet`: `20103` rows, `2143144` file bytes (2.04 MiB), `13021047` physical bytes (12.42 MiB), `13176125` encoded bytes (12.57 MiB), `2113704` compressed data bytes (2.02 MiB)
