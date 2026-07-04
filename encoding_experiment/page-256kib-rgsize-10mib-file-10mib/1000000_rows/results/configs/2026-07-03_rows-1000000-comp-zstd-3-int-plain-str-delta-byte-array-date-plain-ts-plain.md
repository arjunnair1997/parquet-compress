# ClickBench Parquet Experiment

- Started: `2026-07-03T15:00:41-04:00`
- Write elapsed: `12.06s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `518911401` (494.87 MiB)
- Compressed column data bytes after codec compression: `87749823` (83.68 MiB)
- Parquet file bytes: `88648669` (84.54 MiB)
- Physical/encoded ratio: `1.373x`
- Encoded/compressed-data ratio: `5.914x`
- Physical/compressed-data ratio: `8.119x`
- Physical/parquet-file ratio: `8.036x`
- Files: `29`

## Settings

- Compression: `zstd-3`
- Int encoding: `plain`
- String encoding: `delta-byte-array`
- Date encoding: `plain`
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
- Files read: `29`
- Elapsed: `7.523s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004555` (7.63 MiB) | `8005363` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `63369` (61.88 KiB) | `0.999x` | `63.179x` | `63.122x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:528` | `1000000` | `138409995` (132.00 MiB) | `64477766` (61.49 MiB) | `13967505` (13.32 MiB) | `2.147x` | `4.616x` | `9.909x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003525` (3.82 MiB) | `4920` (4.80 KiB) | `0.999x` | `813.725x` | `813.008x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004552` (7.63 MiB) | `2517963` (2.40 MiB) | `0.999x` | `3.179x` | `3.177x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4923` (4.81 KiB) | `0.999x` | `813.229x` | `812.513x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `4943` (4.83 KiB) | `0.999x` | `809.940x` | `809.225x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003592` (3.82 MiB) | `408058` (398.49 KiB) | `0.999x` | `9.811x` | `9.803x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `191205` (186.72 KiB) | `0.999x` | `20.939x` | `20.920x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004550` (7.63 MiB) | `618369` (603.88 KiB) | `0.999x` | `12.945x` | `12.937x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `4242` (4.14 KiB) | `0.999x` | `943.784x` | `942.951x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `106112` (103.62 KiB) | `0.999x` | `37.730x` | `37.696x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `134990` (131.83 KiB) | `0.999x` | `29.658x` | `29.632x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:354` | `1000000` | `88562192` (84.46 MiB) | `40465826` (38.59 MiB) | `15098027` (14.40 MiB) | `2.189x` | `2.680x` | `5.866x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:326` | `1000000` | `79583339` (75.90 MiB) | `38994337` (37.19 MiB) | `14561303` (13.89 MiB) | `2.041x` | `2.678x` | `5.465x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `178752` (174.56 KiB) | `0.999x` | `22.397x` | `22.377x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `275181` (268.73 KiB) | `0.999x` | `14.549x` | `14.536x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `231326` (225.90 KiB) | `0.999x` | `17.307x` | `17.292x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `87584` (85.53 KiB) | `0.999x` | `45.711x` | `45.670x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `48144` (47.02 KiB) | `0.999x` | `83.159x` | `83.084x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `187160` (182.77 KiB) | `0.999x` | `21.391x` | `21.372x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `186046` (181.69 KiB) | `0.999x` | `21.519x` | `21.500x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `81834` (79.92 KiB) | `0.999x` | `48.923x` | `48.879x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `53667` (52.41 KiB) | `0.999x` | `74.600x` | `74.534x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `128106` (125.10 KiB) | `0.999x` | `31.252x` | `31.224x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `3354477` (3.20 MiB) | `1041835` (1017.42 KiB) | `326069` (318.43 KiB) | `3.220x` | `3.195x` | `10.288x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `26180` (25.57 KiB) | `0.999x` | `152.925x` | `152.788x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `24938` (24.35 KiB) | `0.999x` | `160.542x` | `160.398x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `156248` (152.59 KiB) | `0.999x` | `25.623x` | `25.600x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `3767530` (3.59 MiB) | `862465` (842.25 KiB) | `179729` (175.52 KiB) | `4.368x` | `4.799x` | `20.962x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003530` (3.82 MiB) | `6087` (5.94 KiB) | `0.999x` | `657.718x` | `657.138x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003543` (3.82 MiB) | `6485` (6.33 KiB) | `0.999x` | `617.354x` | `616.808x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `28623` (27.95 KiB) | `0.999x` | `139.873x` | `139.748x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `22552` (22.02 KiB) | `0.999x` | `177.527x` | `177.368x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `81583` (79.67 KiB) | `282800` (276.17 KiB) | `41002` (40.04 KiB) | `0.288x` | `6.897x` | `1.990x` | `81583` (79.67 KiB) |
| `Params` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `81280` (79.38 KiB) | `3683` (3.60 KiB) | `0.000x` | `22.069x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003592` (3.82 MiB) | `323794` (316.21 KiB) | `0.999x` | `12.365x` | `12.354x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `288272` (281.52 KiB) | `0.999x` | `13.888x` | `13.876x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `101598` (99.22 KiB) | `0.999x` | `39.406x` | `39.371x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `3528017` (3.36 MiB) | `2957835` (2.82 MiB) | `825797` (806.44 KiB) | `1.193x` | `3.582x` | `4.272x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003579` (3.82 MiB) | `30677` (29.96 KiB) | `0.999x` | `130.508x` | `130.391x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `164975` (161.11 KiB) | `0.999x` | `24.268x` | `24.246x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `305757` (298.59 KiB) | `0.999x` | `13.094x` | `13.082x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003591` (3.82 MiB) | `319858` (312.36 KiB) | `0.999x` | `12.517x` | `12.506x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `99975` (97.63 KiB) | `0.999x` | `40.046x` | `40.010x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004554` (7.63 MiB) | `2476339` (2.36 MiB) | `0.999x` | `3.232x` | `3.231x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `90123` (88.01 KiB) | `0.999x` | `44.424x` | `44.384x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `73933` (72.20 KiB) | `0.999x` | `54.152x` | `54.103x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `123727` (120.83 KiB) | `0.999x` | `32.358x` | `32.329x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003530` (3.82 MiB) | `4371` (4.27 KiB) | `0.999x` | `915.930x` | `915.122x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:85` | `1000000` | `13587860` (12.96 MiB) | `140931` (137.63 KiB) | `21510` (21.01 KiB) | `96.415x` | `6.552x` | `631.700x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003548` (3.82 MiB) | `7032` (6.87 KiB) | `0.999x` | `569.333x` | `568.828x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003561` (3.82 MiB) | `56699` (55.37 KiB) | `0.999x` | `70.611x` | `70.548x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003554` (3.82 MiB) | `8034` (7.85 KiB) | `0.999x` | `498.326x` | `497.884x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003539` (3.82 MiB) | `25188` (24.60 KiB) | `0.999x` | `158.946x` | `158.806x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004550` (7.63 MiB) | `694093` (677.83 KiB) | `0.999x` | `11.532x` | `11.526x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:137` | `1000000` | `27797671` (26.51 MiB) | `21049867` (20.07 MiB) | `5596601` (5.34 MiB) | `1.321x` | `3.761x` | `4.967x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003642` (3.82 MiB) | `3831613` (3.65 MiB) | `0.999x` | `1.045x` | `1.044x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `4242` (4.14 KiB) | `0.999x` | `943.784x` | `942.951x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `4242` (4.14 KiB) | `0.999x` | `943.784x` | `942.951x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `4242` (4.14 KiB) | `0.999x` | `943.784x` | `942.951x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `82484` (80.55 KiB) | `0.999x` | `48.538x` | `48.494x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `4242` (4.14 KiB) | `0.999x` | `943.784x` | `942.951x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `1000000` (976.56 KiB) | `206678` (201.83 KiB) | `51359` (50.16 KiB) | `4.838x` | `4.024x` | `19.471x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004557` (7.63 MiB) | `2520034` (2.40 MiB) | `0.999x` | `3.176x` | `3.175x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `142783` (139.44 KiB) | `0.999x` | `28.040x` | `28.015x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `108045` (105.51 KiB) | `0.999x` | `37.055x` | `37.022x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `122692` (119.82 KiB) | `0.999x` | `32.631x` | `32.602x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `193529` (188.99 KiB) | `0.999x` | `20.687x` | `20.669x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003582` (3.82 MiB) | `173664` (169.59 KiB) | `0.999x` | `23.054x` | `23.033x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003599` (3.82 MiB) | `426572` (416.57 KiB) | `0.999x` | `9.386x` | `9.377x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003579` (3.82 MiB) | `70829` (69.17 KiB) | `0.999x` | `56.525x` | `56.474x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003524` (3.82 MiB) | `4235` (4.14 KiB) | `0.999x` | `945.342x` | `944.510x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003571` (3.82 MiB) | `55149` (53.86 KiB) | `0.999x` | `72.596x` | `72.531x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `2001192` (1.91 MiB) | `331509` (323.74 KiB) | `65004` (63.48 KiB) | `6.037x` | `5.100x` | `30.786x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `3325142` (3.17 MiB) | `963284` (940.71 KiB) | `195515` (190.93 KiB) | `3.452x` | `4.927x` | `17.007x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `81280` (79.38 KiB) | `3683` (3.60 KiB) | `0.000x` | `22.069x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `81280` (79.38 KiB) | `3683` (3.60 KiB) | `0.000x` | `22.069x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `4242` (4.14 KiB) | `0.999x` | `943.784x` | `942.951x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003541` (3.82 MiB) | `61436` (60.00 KiB) | `0.999x` | `65.166x` | `65.108x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `135362` (132.19 KiB) | `0.999x` | `29.577x` | `29.550x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003598` (3.82 MiB) | `333300` (325.49 KiB) | `0.999x` | `12.012x` | `12.001x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003643` (3.82 MiB) | `1245933` (1.19 MiB) | `0.999x` | `3.213x` | `3.210x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003640` (3.82 MiB) | `938050` (916.06 KiB) | `0.999x` | `4.268x` | `4.264x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003624` (3.82 MiB) | `549889` (537.00 KiB) | `0.999x` | `7.281x` | `7.274x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `5329` (5.20 KiB) | `0.999x` | `751.272x` | `750.610x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `1024` (1.00 KiB) | `86525` (84.50 KiB) | `6158` (6.01 KiB) | `0.012x` | `14.051x` | `0.166x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004441` (7.63 MiB) | `5726` (5.59 KiB) | `0.999x` | `1397.911x` | `1397.136x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `81280` (79.38 KiB) | `3683` (3.60 KiB) | `0.000x` | `22.069x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `3000000` (2.86 MiB) | `86044` (84.03 KiB) | `6923` (6.76 KiB) | `34.866x` | `12.429x` | `433.338x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `4242` (4.14 KiB) | `0.999x` | `943.784x` | `942.951x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `58030` (56.67 KiB) | `303581` (296.47 KiB) | `33979` (33.18 KiB) | `0.191x` | `8.934x` | `1.708x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `22051` (21.53 KiB) | `198947` (194.28 KiB) | `28206` (27.54 KiB) | `0.111x` | `7.053x` | `0.782x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `25445` (24.85 KiB) | `211156` (206.21 KiB) | `33635` (32.85 KiB) | `0.121x` | `6.278x` | `0.757x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `48191` (47.06 KiB) | `234357` (228.86 KiB) | `23440` (22.89 KiB) | `0.206x` | `9.998x` | `2.056x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `49433` (48.27 KiB) | `267116` (260.86 KiB) | `40088` (39.15 KiB) | `0.185x` | `6.663x` | `1.233x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `16873` (16.48 KiB) | `213000` (208.01 KiB) | `27882` (27.23 KiB) | `0.079x` | `7.639x` | `0.605x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `91870` (89.72 KiB) | `326877` (319.22 KiB) | `54137` (52.87 KiB) | `0.281x` | `6.038x` | `1.697x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `13001` (12.70 KiB) | `149448` (145.95 KiB) | `20607` (20.12 KiB) | `0.087x` | `7.252x` | `0.631x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `28101` (27.44 KiB) | `189947` (185.50 KiB) | `23144` (22.60 KiB) | `0.148x` | `8.207x` | `1.214x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `45607` (44.54 KiB) | `260324` (254.22 KiB) | `38421` (37.52 KiB) | `0.175x` | `6.776x` | `1.187x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003557` (3.82 MiB) | `21189` (20.69 KiB) | `0.999x` | `188.945x` | `188.777x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004557` (7.63 MiB) | `2841886` (2.71 MiB) | `0.999x` | `2.817x` | `2.815x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004552` (7.63 MiB) | `3580402` (3.41 MiB) | `0.999x` | `2.236x` | `2.234x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003547` (3.82 MiB) | `5652` (5.52 KiB) | `0.999x` | `708.342x` | `707.714x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00000.parquet`: `35889` rows, `2857148` file bytes (2.72 MiB), `27875599` physical bytes (26.58 MiB), `18031236` encoded bytes (17.20 MiB), `2825111` compressed data bytes (2.69 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00001.parquet`: `35984` rows, `2806554` file bytes (2.68 MiB), `27645074` physical bytes (26.36 MiB), `17840307` encoded bytes (17.01 MiB), `2773975` compressed data bytes (2.65 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00002.parquet`: `35879` rows, `2848237` file bytes (2.72 MiB), `28094680` physical bytes (26.79 MiB), `18081420` encoded bytes (17.24 MiB), `2814978` compressed data bytes (2.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00003.parquet`: `36267` rows, `2824478` file bytes (2.69 MiB), `28070313` physical bytes (26.77 MiB), `18025354` encoded bytes (17.19 MiB), `2792081` compressed data bytes (2.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00004.parquet`: `35612` rows, `2783808` file bytes (2.65 MiB), `27705357` physical bytes (26.42 MiB), `17758072` encoded bytes (16.94 MiB), `2750632` compressed data bytes (2.62 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00005.parquet`: `36078` rows, `2800424` file bytes (2.67 MiB), `27657416` physical bytes (26.38 MiB), `17917171` encoded bytes (17.09 MiB), `2767970` compressed data bytes (2.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00006.parquet`: `36220` rows, `2847400` file bytes (2.72 MiB), `28044091` physical bytes (26.74 MiB), `17976753` encoded bytes (17.14 MiB), `2814970` compressed data bytes (2.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00007.parquet`: `36251` rows, `2850000` file bytes (2.72 MiB), `28054804` physical bytes (26.76 MiB), `18108919` encoded bytes (17.27 MiB), `2817188` compressed data bytes (2.69 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00008.parquet`: `36287` rows, `2790346` file bytes (2.66 MiB), `28103637` physical bytes (26.80 MiB), `18031570` encoded bytes (17.20 MiB), `2757624` compressed data bytes (2.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00009.parquet`: `35856` rows, `2820234` file bytes (2.69 MiB), `27656610` physical bytes (26.38 MiB), `17862031` encoded bytes (17.03 MiB), `2788084` compressed data bytes (2.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00010.parquet`: `36425` rows, `2790321` file bytes (2.66 MiB), `28072449` physical bytes (26.77 MiB), `18049243` encoded bytes (17.21 MiB), `2757731` compressed data bytes (2.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00011.parquet`: `36013` rows, `2745298` file bytes (2.62 MiB), `27730609` physical bytes (26.45 MiB), `17859281` encoded bytes (17.03 MiB), `2712925` compressed data bytes (2.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00012.parquet`: `36361` rows, `2806100` file bytes (2.68 MiB), `27867104` physical bytes (26.58 MiB), `18025640` encoded bytes (17.19 MiB), `2773992` compressed data bytes (2.65 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00013.parquet`: `36253` rows, `2825180` file bytes (2.69 MiB), `28066642` physical bytes (26.77 MiB), `18041535` encoded bytes (17.21 MiB), `2792560` compressed data bytes (2.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00014.parquet`: `35150` rows, `3230584` file bytes (3.08 MiB), `24307305` physical bytes (23.18 MiB), `17769329` encoded bytes (16.95 MiB), `3199788` compressed data bytes (3.05 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00015.parquet`: `36094` rows, `3402682` file bytes (3.25 MiB), `25222378` physical bytes (24.05 MiB), `18338001` encoded bytes (17.49 MiB), `3372257` compressed data bytes (3.22 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00016.parquet`: `35290` rows, `3461083` file bytes (3.30 MiB), `23209979` physical bytes (22.13 MiB), `18804475` encoded bytes (17.93 MiB), `3430844` compressed data bytes (3.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00017.parquet`: `34778` rows, `3504468` file bytes (3.34 MiB), `22442271` physical bytes (21.40 MiB), `19130193` encoded bytes (18.24 MiB), `3473721` compressed data bytes (3.31 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00018.parquet`: `34311` rows, `3581837` file bytes (3.42 MiB), `22381802` physical bytes (21.34 MiB), `19060393` encoded bytes (18.18 MiB), `3551519` compressed data bytes (3.39 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00019.parquet`: `34888` rows, `3478783` file bytes (3.32 MiB), `22194807` physical bytes (21.17 MiB), `19025495` encoded bytes (18.14 MiB), `3448052` compressed data bytes (3.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00020.parquet`: `34465` rows, `3539254` file bytes (3.38 MiB), `22170418` physical bytes (21.14 MiB), `18919686` encoded bytes (18.04 MiB), `3508369` compressed data bytes (3.35 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00021.parquet`: `35009` rows, `3492378` file bytes (3.33 MiB), `22384308` physical bytes (21.35 MiB), `19106926` encoded bytes (18.22 MiB), `3461756` compressed data bytes (3.30 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00022.parquet`: `34954` rows, `3492690` file bytes (3.33 MiB), `22364900` physical bytes (21.33 MiB), `19029376` encoded bytes (18.15 MiB), `3462411` compressed data bytes (3.30 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00023.parquet`: `34413` rows, `3562876` file bytes (3.40 MiB), `22375152` physical bytes (21.34 MiB), `19186709` encoded bytes (18.30 MiB), `3532115` compressed data bytes (3.37 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00024.parquet`: `35170` rows, `3488126` file bytes (3.33 MiB), `22346573` physical bytes (21.31 MiB), `19019796` encoded bytes (18.14 MiB), `3457988` compressed data bytes (3.30 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00025.parquet`: `34778` rows, `3497423` file bytes (3.34 MiB), `22178523` physical bytes (21.15 MiB), `18948713` encoded bytes (18.07 MiB), `3466745` compressed data bytes (3.31 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00026.parquet`: `35162` rows, `3483252` file bytes (3.32 MiB), `22379656` physical bytes (21.34 MiB), `18988706` encoded bytes (18.11 MiB), `3452767` compressed data bytes (3.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00027.parquet`: `34944` rows, `3510927` file bytes (3.35 MiB), `22389426` physical bytes (21.35 MiB), `19103412` encoded bytes (18.22 MiB), `3480485` compressed data bytes (3.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00028.parquet`: `5219` rows, `526778` file bytes (514.43 KiB), `3406741` physical bytes (3.25 MiB), `2871659` encoded bytes (2.74 MiB), `511185` compressed data bytes (499.20 KiB)
