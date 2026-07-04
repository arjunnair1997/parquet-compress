# ClickBench Parquet Experiment

- Started: `2026-07-03T19:44:03-04:00`
- Write elapsed: `11.918s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `719093006` (685.78 MiB)
- Compressed column data bytes after codec compression: `93993545` (89.64 MiB)
- Parquet file bytes: `94906884` (90.51 MiB)
- Physical/encoded ratio: `0.991x`
- Encoded/compressed-data ratio: `7.650x`
- Physical/compressed-data ratio: `7.579x`
- Physical/parquet-file ratio: `7.506x`
- Files: `29`

## Settings

- Compression: `zstd-3`
- Int encoding: `plain`
- String encoding: `delta-length-byte-array`
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
- Files read: `29`
- Elapsed: `7.334s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `8004636` (7.63 MiB) | `8005441` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003650` (3.82 MiB) | `64838` (63.32 KiB) | `0.999x` | `61.749x` | `61.692x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:520` | `1000000` | `138409995` (132.00 MiB) | `140027499` (133.54 MiB) | `14467865` (13.80 MiB) | `0.988x` | `9.679x` | `9.567x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `4990` (4.87 KiB) | `0.999x` | `802.322x` | `801.603x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:58, DICTIONARY_PAGE/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `7355778` (7.02 MiB) | `4037833` (3.85 MiB) | `1.088x` | `1.822x` | `1.981x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `4988` (4.87 KiB) | `0.999x` | `802.644x` | `801.925x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `5007` (4.89 KiB) | `0.999x` | `799.598x` | `798.882x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003658` (3.82 MiB) | `408408` (398.84 KiB) | `0.999x` | `9.803x` | `9.794x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003649` (3.82 MiB) | `191008` (186.53 KiB) | `0.999x` | `20.961x` | `20.942x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `8004631` (7.63 MiB) | `618179` (603.69 KiB) | `0.999x` | `12.949x` | `12.941x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `4294` (4.19 KiB) | `0.999x` | `932.368x` | `931.532x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003650` (3.82 MiB) | `107024` (104.52 KiB) | `0.999x` | `37.409x` | `37.375x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003642` (3.82 MiB) | `135376` (132.20 KiB) | `0.999x` | `29.574x` | `29.547x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:346` | `1000000` | `88562192` (84.46 MiB) | `89785334` (85.63 MiB) | `16034827` (15.29 MiB) | `0.986x` | `5.599x` | `5.523x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:326` | `1000000` | `79583339` (75.90 MiB) | `80834241` (77.09 MiB) | `15009858` (14.31 MiB) | `0.985x` | `5.385x` | `5.302x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003649` (3.82 MiB) | `178816` (174.62 KiB) | `0.999x` | `22.390x` | `22.369x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003650` (3.82 MiB) | `275460` (269.00 KiB) | `0.999x` | `14.534x` | `14.521x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003649` (3.82 MiB) | `231405` (225.98 KiB) | `0.999x` | `17.301x` | `17.286x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003646` (3.82 MiB) | `87721` (85.67 KiB) | `0.999x` | `45.641x` | `45.599x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003649` (3.82 MiB) | `48196` (47.07 KiB) | `0.999x` | `83.070x` | `82.994x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003648` (3.82 MiB) | `187365` (182.97 KiB) | `0.999x` | `21.368x` | `21.349x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003650` (3.82 MiB) | `186323` (181.96 KiB) | `0.999x` | `21.488x` | `21.468x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003644` (3.82 MiB) | `82140` (80.21 KiB) | `0.999x` | `48.742x` | `48.697x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003646` (3.82 MiB) | `53287` (52.04 KiB) | `0.999x` | `75.134x` | `75.065x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003650` (3.82 MiB) | `129482` (126.45 KiB) | `0.999x` | `30.921x` | `30.892x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `3354477` (3.20 MiB) | `3707410` (3.54 MiB) | `267127` (260.87 KiB) | `0.905x` | `13.879x` | `12.558x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003646` (3.82 MiB) | `26686` (26.06 KiB) | `0.999x` | `150.028x` | `149.891x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003648` (3.82 MiB) | `25271` (24.68 KiB) | `0.999x` | `158.429x` | `158.284x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003644` (3.82 MiB) | `155885` (152.23 KiB) | `0.999x` | `25.683x` | `25.660x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `3767530` (3.59 MiB) | `4017216` (3.83 MiB) | `151218` (147.67 KiB) | `0.938x` | `26.566x` | `24.915x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003597` (3.82 MiB) | `6167` (6.02 KiB) | `0.999x` | `649.197x` | `648.614x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003602` (3.82 MiB) | `6651` (6.50 KiB) | `0.999x` | `601.955x` | `601.413x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003644` (3.82 MiB) | `29112` (28.43 KiB) | `0.999x` | `137.526x` | `137.400x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003648` (3.82 MiB) | `22809` (22.27 KiB) | `0.999x` | `175.529x` | `175.369x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `81583` (79.67 KiB) | `231175` (225.76 KiB) | `30837` (30.11 KiB) | `0.353x` | `7.497x` | `2.646x` | `81583` (79.67 KiB) |
| `Params` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `0` (0 B) | `41737` (40.76 KiB) | `3621` (3.54 KiB) | `0.000x` | `11.526x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003649` (3.82 MiB) | `323804` (316.21 KiB) | `0.999x` | `12.364x` | `12.353x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003651` (3.82 MiB) | `288488` (281.73 KiB) | `0.999x` | `13.878x` | `13.865x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003645` (3.82 MiB) | `101859` (99.47 KiB) | `0.999x` | `39.306x` | `39.270x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `3528017` (3.36 MiB) | `4222103` (4.03 MiB) | `812274` (793.24 KiB) | `0.836x` | `5.198x` | `4.343x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003643` (3.82 MiB) | `30880` (30.16 KiB) | `0.999x` | `129.652x` | `129.534x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003648` (3.82 MiB) | `165149` (161.28 KiB) | `0.999x` | `24.243x` | `24.221x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003648` (3.82 MiB) | `305930` (298.76 KiB) | `0.999x` | `13.087x` | `13.075x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003650` (3.82 MiB) | `319349` (311.86 KiB) | `0.999x` | `12.537x` | `12.525x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003647` (3.82 MiB) | `100101` (97.75 KiB) | `0.999x` | `39.996x` | `39.960x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:58, DICTIONARY_PAGE/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `7209701` (6.88 MiB) | `3978022` (3.79 MiB) | `1.110x` | `1.812x` | `2.011x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003648` (3.82 MiB) | `89171` (87.08 KiB) | `0.999x` | `44.899x` | `44.858x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003646` (3.82 MiB) | `73776` (72.05 KiB) | `0.999x` | `54.268x` | `54.218x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003646` (3.82 MiB) | `123967` (121.06 KiB) | `0.999x` | `32.296x` | `32.267x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `4424` (4.32 KiB) | `0.999x` | `904.970x` | `904.159x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:86` | `1000000` | `13587860` (12.96 MiB) | `13657412` (13.02 MiB) | `18939` (18.50 KiB) | `0.995x` | `721.126x` | `717.454x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003610` (3.82 MiB) | `7160` (6.99 KiB) | `0.999x` | `559.163x` | `558.659x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003619` (3.82 MiB) | `56763` (55.43 KiB) | `0.999x` | `70.532x` | `70.468x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003609` (3.82 MiB) | `8158` (7.97 KiB) | `0.999x` | `490.759x` | `490.316x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003597` (3.82 MiB) | `25336` (24.74 KiB) | `0.999x` | `158.020x` | `157.878x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `8004636` (7.63 MiB) | `694089` (677.82 KiB) | `0.999x` | `11.533x` | `11.526x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:136` | `1000000` | `27797671` (26.51 MiB) | `28791855` (27.46 MiB) | `5613035` (5.35 MiB) | `0.965x` | `5.129x` | `4.952x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003704` (3.82 MiB) | `3848756` (3.67 MiB) | `0.999x` | `1.040x` | `1.039x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `4294` (4.19 KiB) | `0.999x` | `932.368x` | `931.532x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `4294` (4.19 KiB) | `0.999x` | `932.368x` | `931.532x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `4294` (4.19 KiB) | `0.999x` | `932.368x` | `931.532x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003650` (3.82 MiB) | `82515` (80.58 KiB) | `0.999x` | `48.520x` | `48.476x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `4294` (4.19 KiB) | `0.999x` | `932.368x` | `931.532x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `1000000` (976.56 KiB) | `1043084` (1018.64 KiB) | `28214` (27.55 KiB) | `0.959x` | `36.970x` | `35.443x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:58, DICTIONARY_PAGE/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `7351263` (7.01 MiB) | `4039306` (3.85 MiB) | `1.088x` | `1.820x` | `1.981x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003650` (3.82 MiB) | `142641` (139.30 KiB) | `0.999x` | `28.068x` | `28.042x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003643` (3.82 MiB) | `107621` (105.10 KiB) | `0.999x` | `37.201x` | `37.167x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003647` (3.82 MiB) | `122891` (120.01 KiB) | `0.999x` | `32.579x` | `32.549x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003640` (3.82 MiB) | `193631` (189.09 KiB) | `0.999x` | `20.677x` | `20.658x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003647` (3.82 MiB) | `174139` (170.06 KiB) | `0.999x` | `22.991x` | `22.970x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003662` (3.82 MiB) | `426479` (416.48 KiB) | `0.999x` | `9.388x` | `9.379x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003638` (3.82 MiB) | `70885` (69.22 KiB) | `0.999x` | `56.481x` | `56.429x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `4296` (4.20 KiB) | `0.999x` | `931.934x` | `931.099x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003630` (3.82 MiB) | `55410` (54.11 KiB) | `0.999x` | `72.255x` | `72.189x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `2001192` (1.91 MiB) | `2051148` (1.96 MiB) | `32236` (31.48 KiB) | `0.976x` | `63.629x` | `62.079x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `3325142` (3.17 MiB) | `3638003` (3.47 MiB) | `185216` (180.88 KiB) | `0.914x` | `19.642x` | `17.953x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `0` (0 B) | `41737` (40.76 KiB) | `3621` (3.54 KiB) | `0.000x` | `11.526x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `0` (0 B) | `41737` (40.76 KiB) | `3621` (3.54 KiB) | `0.000x` | `11.526x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `4294` (4.19 KiB) | `0.999x` | `932.368x` | `931.532x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003601` (3.82 MiB) | `61653` (60.21 KiB) | `0.999x` | `64.938x` | `64.879x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003644` (3.82 MiB) | `135290` (132.12 KiB) | `0.999x` | `29.593x` | `29.566x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003661` (3.82 MiB) | `333080` (325.27 KiB) | `0.999x` | `12.020x` | `12.009x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003705` (3.82 MiB) | `1246618` (1.19 MiB) | `0.999x` | `3.212x` | `3.209x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003708` (3.82 MiB) | `938120` (916.13 KiB) | `0.999x` | `4.268x` | `4.264x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003686` (3.82 MiB) | `550327` (537.43 KiB) | `0.999x` | `7.275x` | `7.268x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `5395` (5.27 KiB) | `0.999x` | `742.092x` | `741.427x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `1024` (1.00 KiB) | `46358` (45.27 KiB) | `6149` (6.00 KiB) | `0.022x` | `7.539x` | `0.167x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `8004519` (7.63 MiB) | `5751` (5.62 KiB) | `0.999x` | `1391.848x` | `1391.062x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `0` (0 B) | `41737` (40.76 KiB) | `3621` (3.54 KiB) | `0.000x` | `11.526x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `3000000` (2.86 MiB) | `3044033` (2.90 MiB) | `5877` (5.74 KiB) | `0.986x` | `517.957x` | `510.465x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `4294` (4.19 KiB) | `0.999x` | `932.368x` | `931.532x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `58030` (56.67 KiB) | `217595` (212.50 KiB) | `27479` (26.83 KiB) | `0.267x` | `7.919x` | `2.112x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `22051` (21.53 KiB) | `123065` (120.18 KiB) | `23324` (22.78 KiB) | `0.179x` | `5.276x` | `0.945x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `25445` (24.85 KiB) | `130838` (127.77 KiB) | `28599` (27.93 KiB) | `0.194x` | `4.575x` | `0.890x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `48191` (47.06 KiB) | `156060` (152.40 KiB) | `18348` (17.92 KiB) | `0.309x` | `8.506x` | `2.626x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `49433` (48.27 KiB) | `188339` (183.92 KiB) | `34032` (33.23 KiB) | `0.262x` | `5.534x` | `1.453x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `16873` (16.48 KiB) | `133627` (130.50 KiB) | `21975` (21.46 KiB) | `0.126x` | `6.081x` | `0.768x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `91870` (89.72 KiB) | `253429` (247.49 KiB) | `47714` (46.60 KiB) | `0.363x` | `5.311x` | `1.925x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `13001` (12.70 KiB) | `94556` (92.34 KiB) | `18205` (17.78 KiB) | `0.137x` | `5.194x` | `0.714x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `28101` (27.44 KiB) | `129034` (126.01 KiB) | `19946` (19.48 KiB) | `0.218x` | `6.469x` | `1.409x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:58` | `1000000` | `45607` (44.54 KiB) | `211112` (206.16 KiB) | `36429` (35.58 KiB) | `0.216x` | `5.795x` | `1.252x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003617` (3.82 MiB) | `21280` (20.78 KiB) | `0.999x` | `188.140x` | `187.970x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `8004633` (7.63 MiB) | `2843292` (2.71 MiB) | `0.999x` | `2.815x` | `2.814x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `8000000` (7.63 MiB) | `8004634` (7.63 MiB) | `3581933` (3.42 MiB) | `0.999x` | `2.235x` | `2.233x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:58` | `1000000` | `4000000` (3.81 MiB) | `4003605` (3.82 MiB) | `5677` (5.54 KiB) | `0.999x` | `705.233x` | `704.597x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00000.parquet`: `35117` rows, `3065546` file bytes (2.92 MiB), `27262167` physical bytes (26.00 MiB), `27358172` encoded bytes (26.09 MiB), `3033514` compressed data bytes (2.89 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00001.parquet`: `35138` rows, `3016065` file bytes (2.88 MiB), `27043411` physical bytes (25.79 MiB), `27135479` encoded bytes (25.88 MiB), `2983677` compressed data bytes (2.85 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00002.parquet`: `35625` rows, `3118491` file bytes (2.97 MiB), `27849002` physical bytes (26.56 MiB), `27944189` encoded bytes (26.65 MiB), `3085277` compressed data bytes (2.94 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00003.parquet`: `35468` rows, `3032791` file bytes (2.89 MiB), `27454380` physical bytes (26.18 MiB), `27541683` encoded bytes (26.27 MiB), `3000451` compressed data bytes (2.86 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00004.parquet`: `35336` rows, `3036100` file bytes (2.90 MiB), `27507139` physical bytes (26.23 MiB), `27594324` encoded bytes (26.32 MiB), `3003371` compressed data bytes (2.86 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00005.parquet`: `35456` rows, `3013706` file bytes (2.87 MiB), `27256358` physical bytes (25.99 MiB), `27348313` encoded bytes (26.08 MiB), `2981241` compressed data bytes (2.84 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00006.parquet`: `35615` rows, `3039935` file bytes (2.90 MiB), `27458756` physical bytes (26.19 MiB), `27552510` encoded bytes (26.28 MiB), `3007369` compressed data bytes (2.87 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00007.parquet`: `35674` rows, `3078931` file bytes (2.94 MiB), `27658684` physical bytes (26.38 MiB), `27756952` encoded bytes (26.47 MiB), `3046191` compressed data bytes (2.91 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00008.parquet`: `35940` rows, `3049887` file bytes (2.91 MiB), `27884421` physical bytes (26.59 MiB), `27973719` encoded bytes (26.68 MiB), `3017198` compressed data bytes (2.88 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00009.parquet`: `35764` rows, `3039296` file bytes (2.90 MiB), `27445958` physical bytes (26.17 MiB), `27535278` encoded bytes (26.26 MiB), `3007066` compressed data bytes (2.87 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00010.parquet`: `36082` rows, `3062110` file bytes (2.92 MiB), `27838417` physical bytes (26.55 MiB), `27932295` encoded bytes (26.64 MiB), `3029517` compressed data bytes (2.89 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00011.parquet`: `36143` rows, `3040514` file bytes (2.90 MiB), `27893777` physical bytes (26.60 MiB), `27985833` encoded bytes (26.69 MiB), `3007829` compressed data bytes (2.87 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00012.parquet`: `35742` rows, `3048382` file bytes (2.91 MiB), `27470090` physical bytes (26.20 MiB), `27563964` encoded bytes (26.29 MiB), `3016211` compressed data bytes (2.88 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00013.parquet`: `35632` rows, `3007541` file bytes (2.87 MiB), `27493230` physical bytes (26.22 MiB), `27581732` encoded bytes (26.30 MiB), `2974904` compressed data bytes (2.84 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00014.parquet`: `35263` rows, `3377683` file bytes (3.22 MiB), `24957101` physical bytes (23.80 MiB), `25199765` encoded bytes (24.03 MiB), `3346458` compressed data bytes (3.19 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00015.parquet`: `35006` rows, `3494233` file bytes (3.33 MiB), `24262326` physical bytes (23.14 MiB), `24549539` encoded bytes (23.41 MiB), `3463917` compressed data bytes (3.30 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00016.parquet`: `34197` rows, `3494426` file bytes (3.33 MiB), `23062256` physical bytes (21.99 MiB), `23408719` encoded bytes (22.32 MiB), `3464051` compressed data bytes (3.30 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00017.parquet`: `34273` rows, `3625112` file bytes (3.46 MiB), `22048374` physical bytes (21.03 MiB), `22440644` encoded bytes (21.40 MiB), `3594512` compressed data bytes (3.43 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00018.parquet`: `34137` rows, `3705433` file bytes (3.53 MiB), `22257118` physical bytes (21.23 MiB), `22643119` encoded bytes (21.59 MiB), `3674738` compressed data bytes (3.50 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00019.parquet`: `34187` rows, `3618355` file bytes (3.45 MiB), `21826156` physical bytes (20.82 MiB), `22216204` encoded bytes (21.19 MiB), `3587696` compressed data bytes (3.42 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00020.parquet`: `34214` rows, `3654874` file bytes (3.49 MiB), `22021784` physical bytes (21.00 MiB), `22417478` encoded bytes (21.38 MiB), `3624246` compressed data bytes (3.46 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00021.parquet`: `34411` rows, `3593695` file bytes (3.43 MiB), `22035243` physical bytes (21.01 MiB), `22430669` encoded bytes (21.39 MiB), `3562826` compressed data bytes (3.40 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00022.parquet`: `34338` rows, `3633453` file bytes (3.47 MiB), `22016616` physical bytes (21.00 MiB), `22405308` encoded bytes (21.37 MiB), `3602986` compressed data bytes (3.44 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00023.parquet`: `34100` rows, `3626704` file bytes (3.46 MiB), `22062086` physical bytes (21.04 MiB), `22450268` encoded bytes (21.41 MiB), `3595959` compressed data bytes (3.43 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00024.parquet`: `34030` rows, `3621340` file bytes (3.45 MiB), `21830113` physical bytes (20.82 MiB), `22216773` encoded bytes (21.19 MiB), `3591101` compressed data bytes (3.42 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00025.parquet`: `34401` rows, `3580592` file bytes (3.41 MiB), `21885263` physical bytes (20.87 MiB), `22279011` encoded bytes (21.25 MiB), `3549850` compressed data bytes (3.39 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00026.parquet`: `34318` rows, `3540827` file bytes (3.38 MiB), `21893251` physical bytes (20.88 MiB), `22280655` encoded bytes (21.25 MiB), `3510219` compressed data bytes (3.35 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00027.parquet`: `34290` rows, `3547757` file bytes (3.38 MiB), `21704100` physical bytes (20.70 MiB), `22094964` encoded bytes (21.07 MiB), `3517511` compressed data bytes (3.35 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00028.parquet`: `20103` rows, `2143105` file bytes (2.04 MiB), `13021047` physical bytes (12.42 MiB), `13255447` encoded bytes (12.64 MiB), `2113659` compressed data bytes (2.02 MiB)
