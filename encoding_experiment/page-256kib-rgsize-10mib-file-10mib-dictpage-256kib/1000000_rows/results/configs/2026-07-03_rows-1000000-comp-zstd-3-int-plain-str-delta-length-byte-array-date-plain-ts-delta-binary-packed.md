# ClickBench Parquet Experiment

- Started: `2026-07-03T23:35:39-04:00`
- Write elapsed: `11.752s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `707414847` (674.64 MiB)
- Compressed column data bytes after codec compression: `90576159` (86.38 MiB)
- Parquet file bytes: `91474601` (87.24 MiB)
- Physical/encoded ratio: `1.007x`
- Encoded/compressed-data ratio: `7.810x`
- Physical/compressed-data ratio: `7.865x`
- Physical/parquet-file ratio: `7.788x`
- Files: `29`

## Settings

- Compression: `zstd-3`
- Int encoding: `plain`
- String encoding: `delta-length-byte-array`
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
- Files read: `29`
- Elapsed: `7.303s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004555` (7.63 MiB) | `8005353` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `64701` (63.18 KiB) | `0.999x` | `61.878x` | `61.823x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:523` | `1000000` | `138409995` (132.00 MiB) | `140031813` (133.54 MiB) | `14458941` (13.79 MiB) | `0.988x` | `9.685x` | `9.573x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003530` (3.82 MiB) | `4915` (4.80 KiB) | `0.999x` | `814.553x` | `813.835x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `3365748` (3.21 MiB) | `2886254` (2.75 MiB) | `2.377x` | `1.166x` | `2.772x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003533` (3.82 MiB) | `4918` (4.80 KiB) | `0.999x` | `814.057x` | `813.339x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4931` (4.82 KiB) | `0.999x` | `811.910x` | `811.194x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003603` (3.82 MiB) | `408251` (398.68 KiB) | `0.999x` | `9.807x` | `9.798x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `191035` (186.56 KiB) | `0.999x` | `20.957x` | `20.939x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004550` (7.63 MiB) | `617840` (603.36 KiB) | `0.999x` | `12.956x` | `12.948x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4230` (4.13 KiB) | `0.999x` | `946.461x` | `945.626x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `107189` (104.68 KiB) | `0.999x` | `37.351x` | `37.317x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `134460` (131.31 KiB) | `0.999x` | `29.775x` | `29.749x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:355` | `1000000` | `88562192` (84.46 MiB) | `89787449` (85.63 MiB) | `16035097` (15.29 MiB) | `0.986x` | `5.599x` | `5.523x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:329` | `1000000` | `79583339` (75.90 MiB) | `80836227` (77.09 MiB) | `15001362` (14.31 MiB) | `0.985x` | `5.389x` | `5.305x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `178724` (174.54 KiB) | `0.999x` | `22.401x` | `22.381x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `275189` (268.74 KiB) | `0.999x` | `14.549x` | `14.535x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `231360` (225.94 KiB) | `0.999x` | `17.305x` | `17.289x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `87564` (85.51 KiB) | `0.999x` | `45.722x` | `45.681x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `48053` (46.93 KiB) | `0.999x` | `83.316x` | `83.241x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `187278` (182.89 KiB) | `0.999x` | `21.378x` | `21.359x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `186157` (181.79 KiB) | `0.999x` | `21.507x` | `21.487x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `81806` (79.89 KiB) | `0.999x` | `48.940x` | `48.896x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `53248` (52.00 KiB) | `0.999x` | `75.188x` | `75.120x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `129589` (126.55 KiB) | `0.999x` | `30.895x` | `30.867x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `3354477` (3.20 MiB) | `3707577` (3.54 MiB) | `266701` (260.45 KiB) | `0.905x` | `13.902x` | `12.578x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `26692` (26.07 KiB) | `0.999x` | `149.992x` | `149.858x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `25127` (24.54 KiB) | `0.999x` | `159.334x` | `159.191x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `155991` (152.33 KiB) | `0.999x` | `25.665x` | `25.643x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `3767530` (3.59 MiB) | `4015693` (3.83 MiB) | `151173` (147.63 KiB) | `0.938x` | `26.564x` | `24.922x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003538` (3.82 MiB) | `6109` (5.97 KiB) | `0.999x` | `655.351x` | `654.772x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003545` (3.82 MiB) | `6566` (6.41 KiB) | `0.999x` | `609.739x` | `609.199x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `29072` (28.39 KiB) | `0.999x` | `137.713x` | `137.589x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `22493` (21.97 KiB) | `0.999x` | `177.993x` | `177.833x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `81583` (79.67 KiB) | `230994` (225.58 KiB) | `30730` (30.01 KiB) | `0.353x` | `7.517x` | `2.655x` | `81583` (79.67 KiB) |
| `Params` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `41679` (40.70 KiB) | `3567` (3.48 KiB) | `0.000x` | `11.685x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `323726` (316.14 KiB) | `0.999x` | `12.367x` | `12.356x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `288532` (281.77 KiB) | `0.999x` | `13.876x` | `13.863x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `101873` (99.49 KiB) | `0.999x` | `39.300x` | `39.265x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `3528017` (3.36 MiB) | `4224078` (4.03 MiB) | `813041` (793.99 KiB) | `0.835x` | `5.195x` | `4.339x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003579` (3.82 MiB) | `30720` (30.00 KiB) | `0.999x` | `130.325x` | `130.208x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `165041` (161.17 KiB) | `0.999x` | `24.258x` | `24.236x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `305724` (298.56 KiB) | `0.999x` | `13.095x` | `13.084x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `319906` (312.41 KiB) | `0.999x` | `12.515x` | `12.504x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `99972` (97.63 KiB) | `0.999x` | `40.047x` | `40.011x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `3489542` (3.33 MiB) | `2901085` (2.77 MiB) | `2.293x` | `1.203x` | `2.758x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `88747` (86.67 KiB) | `0.999x` | `45.112x` | `45.072x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `74220` (72.48 KiB) | `0.999x` | `53.942x` | `53.894x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `123881` (120.98 KiB) | `0.999x` | `32.318x` | `32.289x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003530` (3.82 MiB) | `4360` (4.26 KiB) | `0.999x` | `918.241x` | `917.431x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:85` | `1000000` | `13587860` (12.96 MiB) | `13657002` (13.02 MiB) | `18725` (18.29 KiB) | `0.995x` | `729.346x` | `725.653x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003552` (3.82 MiB) | `7138` (6.97 KiB) | `0.999x` | `560.879x` | `560.381x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003559` (3.82 MiB) | `56770` (55.44 KiB) | `0.999x` | `70.522x` | `70.460x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003554` (3.82 MiB) | `8024` (7.84 KiB) | `0.999x` | `498.947x` | `498.504x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003540` (3.82 MiB) | `25184` (24.59 KiB) | `0.999x` | `158.972x` | `158.831x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004553` (7.63 MiB) | `693951` (677.69 KiB) | `0.999x` | `11.535x` | `11.528x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:136` | `1000000` | `27797671` (26.51 MiB) | `28790277` (27.46 MiB) | `5605309` (5.35 MiB) | `0.966x` | `5.136x` | `4.959x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003642` (3.82 MiB) | `3837303` (3.66 MiB) | `0.999x` | `1.043x` | `1.042x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4230` (4.13 KiB) | `0.999x` | `946.461x` | `945.626x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4230` (4.13 KiB) | `0.999x` | `946.461x` | `945.626x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4230` (4.13 KiB) | `0.999x` | `946.461x` | `945.626x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `82610` (80.67 KiB) | `0.999x` | `48.464x` | `48.420x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4230` (4.13 KiB) | `0.999x` | `946.461x` | `945.626x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `1000000` (976.56 KiB) | `1043049` (1018.60 KiB) | `28225` (27.56 KiB) | `0.959x` | `36.955x` | `35.430x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `3381225` (3.22 MiB) | `2894606` (2.76 MiB) | `2.366x` | `1.168x` | `2.764x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `142936` (139.59 KiB) | `0.999x` | `28.010x` | `27.985x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `107246` (104.73 KiB) | `0.999x` | `37.331x` | `37.297x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `122931` (120.05 KiB) | `0.999x` | `32.568x` | `32.539x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `193533` (189.00 KiB) | `0.999x` | `20.687x` | `20.668x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `173608` (169.54 KiB) | `0.999x` | `23.061x` | `23.040x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003603` (3.82 MiB) | `426549` (416.55 KiB) | `0.999x` | `9.386x` | `9.378x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `70910` (69.25 KiB) | `0.999x` | `56.460x` | `56.410x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `4232` (4.13 KiB) | `0.999x` | `946.014x` | `945.180x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003573` (3.82 MiB) | `55348` (54.05 KiB) | `0.999x` | `72.335x` | `72.270x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `2001192` (1.91 MiB) | `2051114` (1.96 MiB) | `32079` (31.33 KiB) | `0.976x` | `63.939x` | `62.383x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `3325142` (3.17 MiB) | `3639209` (3.47 MiB) | `185035` (180.70 KiB) | `0.914x` | `19.668x` | `17.970x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `41679` (40.70 KiB) | `3567` (3.48 KiB) | `0.000x` | `11.685x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `41679` (40.70 KiB) | `3567` (3.48 KiB) | `0.000x` | `11.685x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4230` (4.13 KiB) | `0.999x` | `946.461x` | `945.626x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003546` (3.82 MiB) | `61590` (60.15 KiB) | `0.999x` | `65.003x` | `64.946x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `135085` (131.92 KiB) | `0.999x` | `29.638x` | `29.611x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003600` (3.82 MiB) | `333175` (325.37 KiB) | `0.999x` | `12.017x` | `12.006x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003646` (3.82 MiB) | `1245841` (1.19 MiB) | `0.999x` | `3.214x` | `3.211x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003639` (3.82 MiB) | `937854` (915.87 KiB) | `0.999x` | `4.269x` | `4.265x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003627` (3.82 MiB) | `549872` (536.98 KiB) | `0.999x` | `7.281x` | `7.274x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `5326` (5.20 KiB) | `0.999x` | `751.696x` | `751.033x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `1024` (1.00 KiB) | `46330` (45.24 KiB) | `6134` (5.99 KiB) | `0.022x` | `7.553x` | `0.167x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004444` (7.63 MiB) | `5678` (5.54 KiB) | `0.999x` | `1409.729x` | `1408.947x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `41679` (40.70 KiB) | `3567` (3.48 KiB) | `0.000x` | `11.685x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `3000000` (2.86 MiB) | `3044023` (2.90 MiB) | `5773` (5.64 KiB) | `0.986x` | `527.286x` | `519.660x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4230` (4.13 KiB) | `0.999x` | `946.461x` | `945.626x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `58030` (56.67 KiB) | `216077` (211.01 KiB) | `27602` (26.96 KiB) | `0.269x` | `7.828x` | `2.102x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `22051` (21.53 KiB) | `122517` (119.65 KiB) | `23382` (22.83 KiB) | `0.180x` | `5.240x` | `0.943x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `25445` (24.85 KiB) | `130304` (127.25 KiB) | `28706` (28.03 KiB) | `0.195x` | `4.539x` | `0.886x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `48191` (47.06 KiB) | `155573` (151.93 KiB) | `18375` (17.94 KiB) | `0.310x` | `8.467x` | `2.623x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `49433` (48.27 KiB) | `188268` (183.86 KiB) | `33902` (33.11 KiB) | `0.263x` | `5.553x` | `1.458x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `16873` (16.48 KiB) | `133597` (130.47 KiB) | `22096` (21.58 KiB) | `0.126x` | `6.046x` | `0.764x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `91870` (89.72 KiB) | `253767` (247.82 KiB) | `47776` (46.66 KiB) | `0.362x` | `5.312x` | `1.923x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `13001` (12.70 KiB) | `94779` (92.56 KiB) | `18314` (17.88 KiB) | `0.137x` | `5.175x` | `0.710x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `28101` (27.44 KiB) | `129437` (126.40 KiB) | `19911` (19.44 KiB) | `0.217x` | `6.501x` | `1.411x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `45607` (44.54 KiB) | `212178` (207.21 KiB) | `36426` (35.57 KiB) | `0.215x` | `5.825x` | `1.252x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003559` (3.82 MiB) | `21185` (20.69 KiB) | `0.999x` | `188.981x` | `188.813x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004557` (7.63 MiB) | `2843318` (2.71 MiB) | `0.999x` | `2.815x` | `2.814x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004558` (7.63 MiB) | `3581370` (3.42 MiB) | `0.999x` | `2.235x` | `2.234x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003541` (3.82 MiB) | `5641` (5.51 KiB) | `0.999x` | `709.722x` | `709.094x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00000.parquet`: `35430` rows, `2983389` file bytes (2.85 MiB), `27510352` physical bytes (26.24 MiB), `27364500` encoded bytes (26.10 MiB), `2951421` compressed data bytes (2.81 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00001.parquet`: `35216` rows, `2928999` file bytes (2.79 MiB), `27076840` physical bytes (25.82 MiB), `26917368` encoded bytes (25.67 MiB), `2896601` compressed data bytes (2.76 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00002.parquet`: `35687` rows, `2995614` file bytes (2.86 MiB), `27924974` physical bytes (26.63 MiB), `27769647` encoded bytes (26.48 MiB), `2962468` compressed data bytes (2.83 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00003.parquet`: `35521` rows, `2947846` file bytes (2.81 MiB), `27502590` physical bytes (26.23 MiB), `27344707` encoded bytes (26.08 MiB), `2915589` compressed data bytes (2.78 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00004.parquet`: `35422` rows, `2932233` file bytes (2.80 MiB), `27560533` physical bytes (26.28 MiB), `27396966` encoded bytes (26.13 MiB), `2899524` compressed data bytes (2.77 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00005.parquet`: `36061` rows, `2948235` file bytes (2.81 MiB), `27704662` physical bytes (26.42 MiB), `27542927` encoded bytes (26.27 MiB), `2915733` compressed data bytes (2.78 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00006.parquet`: `35645` rows, `2940517` file bytes (2.80 MiB), `27521918` physical bytes (26.25 MiB), `27363450` encoded bytes (26.10 MiB), `2908045` compressed data bytes (2.77 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00007.parquet`: `35774` rows, `2975587` file bytes (2.84 MiB), `27716115` physical bytes (26.43 MiB), `27560805` encoded bytes (26.28 MiB), `2942793` compressed data bytes (2.81 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00008.parquet`: `35981` rows, `2949615` file bytes (2.81 MiB), `27931964` physical bytes (26.64 MiB), `27764998` encoded bytes (26.48 MiB), `2917002` compressed data bytes (2.78 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00009.parquet`: `35767` rows, `2938066` file bytes (2.80 MiB), `27504777` physical bytes (26.23 MiB), `27344864` encoded bytes (26.08 MiB), `2905925` compressed data bytes (2.77 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00010.parquet`: `35759` rows, `2921382` file bytes (2.79 MiB), `27515244` physical bytes (26.24 MiB), `27353019` encoded bytes (26.09 MiB), `2889042` compressed data bytes (2.76 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00011.parquet`: `36235` rows, `2936588` file bytes (2.80 MiB), `27938515` physical bytes (26.64 MiB), `27774665` encoded bytes (26.49 MiB), `2904124` compressed data bytes (2.77 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00012.parquet`: `36055` rows, `2967941` file bytes (2.83 MiB), `27704832` physical bytes (26.42 MiB), `27547393` encoded bytes (26.27 MiB), `2935812` compressed data bytes (2.80 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00013.parquet`: `35621` rows, `2915444` file bytes (2.78 MiB), `27552942` physical bytes (26.28 MiB), `27394495` encoded bytes (26.13 MiB), `2882910` compressed data bytes (2.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00014.parquet`: `35460` rows, `3294815` file bytes (3.14 MiB), `24924737` physical bytes (23.77 MiB), `24731823` encoded bytes (23.59 MiB), `3263753` compressed data bytes (3.11 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00015.parquet`: `35323` rows, `3413266` file bytes (3.26 MiB), `24538743` physical bytes (23.40 MiB), `24325166` encoded bytes (23.20 MiB), `3383086` compressed data bytes (3.23 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00016.parquet`: `35078` rows, `3471086` file bytes (3.31 MiB), `23478931` physical bytes (22.39 MiB), `23286720` encoded bytes (22.21 MiB), `3440837` compressed data bytes (3.28 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00017.parquet`: `34776` rows, `3528413` file bytes (3.36 MiB), `22385042` physical bytes (21.35 MiB), `22197291` encoded bytes (21.17 MiB), `3497724` compressed data bytes (3.34 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00018.parquet`: `34588` rows, `3620481` file bytes (3.45 MiB), `22524702` physical bytes (21.48 MiB), `22329113` encoded bytes (21.29 MiB), `3590107` compressed data bytes (3.42 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00019.parquet`: `34990` rows, `3518613` file bytes (3.36 MiB), `22312801` physical bytes (21.28 MiB), `22123462` encoded bytes (21.10 MiB), `3487825` compressed data bytes (3.33 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00020.parquet`: `34292` rows, `3577085` file bytes (3.41 MiB), `22145645` physical bytes (21.12 MiB), `21964132` encoded bytes (20.95 MiB), `3546494` compressed data bytes (3.38 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00021.parquet`: `34914` rows, `3532740` file bytes (3.37 MiB), `22360995` physical bytes (21.33 MiB), `22170716` encoded bytes (21.14 MiB), `3501913` compressed data bytes (3.34 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00022.parquet`: `35010` rows, `3509316` file bytes (3.35 MiB), `22353735` physical bytes (21.32 MiB), `22152295` encoded bytes (21.13 MiB), `3478926` compressed data bytes (3.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00023.parquet`: `34405` rows, `3575104` file bytes (3.41 MiB), `22374713` physical bytes (21.34 MiB), `22188082` encoded bytes (21.16 MiB), `3544323` compressed data bytes (3.38 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00024.parquet`: `34849` rows, `3484058` file bytes (3.32 MiB), `22161012` physical bytes (21.13 MiB), `21967976` encoded bytes (20.95 MiB), `3453929` compressed data bytes (3.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00025.parquet`: `34780` rows, `3509792` file bytes (3.35 MiB), `22174573` physical bytes (21.15 MiB), `21984354` encoded bytes (20.97 MiB), `3479269` compressed data bytes (3.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00026.parquet`: `34811` rows, `3472873` file bytes (3.31 MiB), `22204380` physical bytes (21.18 MiB), `22014502` encoded bytes (20.99 MiB), `3442180` compressed data bytes (3.28 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00027.parquet`: `35125` rows, `3527536` file bytes (3.36 MiB), `22386073` physical bytes (21.35 MiB), `22197456` encoded bytes (21.17 MiB), `3497291` compressed data bytes (3.34 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00028.parquet`: `11425` rows, `1157967` file bytes (1.10 MiB), `7406284` physical bytes (7.06 MiB), `7341955` encoded bytes (7.00 MiB), `1141513` compressed data bytes (1.09 MiB)
