# ClickBench Parquet Experiment

- Started: `2026-07-03T23:35:38-04:00`
- Write elapsed: `12.068s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `505127589` (481.73 MiB)
- Compressed column data bytes after codec compression: `88897061` (84.78 MiB)
- Parquet file bytes: `89795970` (85.64 MiB)
- Physical/encoded ratio: `1.410x`
- Encoded/compressed-data ratio: `5.682x`
- Physical/compressed-data ratio: `8.014x`
- Physical/parquet-file ratio: `7.934x`
- Files: `29`

## Settings

- Compression: `zstd-3`
- Int encoding: `plain`
- String encoding: `delta-byte-array`
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
- Elapsed: `7.485s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004554` (7.63 MiB) | `8005359` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `64098` (62.60 KiB) | `0.999x` | `62.460x` | `62.404x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:526` | `1000000` | `138409995` (132.00 MiB) | `64472113` (61.49 MiB) | `13975352` (13.33 MiB) | `2.147x` | `4.613x` | `9.904x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `4922` (4.81 KiB) | `0.999x` | `813.395x` | `812.678x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `3368014` (3.21 MiB) | `2885749` (2.75 MiB) | `2.375x` | `1.167x` | `2.772x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003527` (3.82 MiB) | `4918` (4.80 KiB) | `0.999x` | `814.056x` | `813.339x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4936` (4.82 KiB) | `0.999x` | `811.088x` | `810.373x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003598` (3.82 MiB) | `408104` (398.54 KiB) | `0.999x` | `9.810x` | `9.801x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `190981` (186.50 KiB) | `0.999x` | `20.963x` | `20.944x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004552` (7.63 MiB) | `618265` (603.77 KiB) | `0.999x` | `12.947x` | `12.939x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4235` (4.14 KiB) | `0.999x` | `945.343x` | `944.510x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `106183` (103.69 KiB) | `0.999x` | `37.705x` | `37.671x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003580` (3.82 MiB) | `135014` (131.85 KiB) | `0.999x` | `29.653x` | `29.627x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:352` | `1000000` | `88562192` (84.46 MiB) | `40464364` (38.59 MiB) | `15089787` (14.39 MiB) | `2.189x` | `2.682x` | `5.869x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:326` | `1000000` | `79583339` (75.90 MiB) | `38990609` (37.18 MiB) | `14550283` (13.88 MiB) | `2.041x` | `2.680x` | `5.470x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `178675` (174.49 KiB) | `0.999x` | `22.407x` | `22.387x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `275096` (268.65 KiB) | `0.999x` | `14.553x` | `14.540x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `231223` (225.80 KiB) | `0.999x` | `17.315x` | `17.299x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `87590` (85.54 KiB) | `0.999x` | `45.708x` | `45.667x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `48114` (46.99 KiB) | `0.999x` | `83.210x` | `83.136x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `187216` (182.83 KiB) | `0.999x` | `21.385x` | `21.366x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `186064` (181.70 KiB) | `0.999x` | `21.517x` | `21.498x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `82613` (80.68 KiB) | `0.999x` | `48.462x` | `48.419x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `53457` (52.20 KiB) | `0.999x` | `74.894x` | `74.826x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `128125` (125.12 KiB) | `0.999x` | `31.247x` | `31.220x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `3354477` (3.20 MiB) | `1042809` (1018.37 KiB) | `326271` (318.62 KiB) | `3.217x` | `3.196x` | `10.281x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `26167` (25.55 KiB) | `0.999x` | `153.001x` | `152.864x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `24703` (24.12 KiB) | `0.999x` | `162.069x` | `161.924x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `155422` (151.78 KiB) | `0.999x` | `25.759x` | `25.736x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `3767530` (3.59 MiB) | `861852` (841.65 KiB) | `179849` (175.63 KiB) | `4.371x` | `4.792x` | `20.948x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003537` (3.82 MiB) | `6082` (5.94 KiB) | `0.999x` | `658.260x` | `657.678x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003541` (3.82 MiB) | `6530` (6.38 KiB) | `0.999x` | `613.100x` | `612.557x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003579` (3.82 MiB) | `28850` (28.17 KiB) | `0.999x` | `138.772x` | `138.648x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003579` (3.82 MiB) | `22576` (22.05 KiB) | `0.999x` | `177.338x` | `177.179x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `81583` (79.67 KiB) | `282592` (275.97 KiB) | `40717` (39.76 KiB) | `0.289x` | `6.940x` | `2.004x` | `81583` (79.67 KiB) |
| `Params` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `81281` (79.38 KiB) | `3685` (3.60 KiB) | `0.000x` | `22.057x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003592` (3.82 MiB) | `323862` (316.27 KiB) | `0.999x` | `12.362x` | `12.351x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `288306` (281.55 KiB) | `0.999x` | `13.887x` | `13.874x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `101595` (99.21 KiB) | `0.999x` | `39.407x` | `39.372x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `3528017` (3.36 MiB) | `2960127` (2.82 MiB) | `826447` (807.08 KiB) | `1.192x` | `3.582x` | `4.269x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003577` (3.82 MiB) | `30711` (29.99 KiB) | `0.999x` | `130.363x` | `130.246x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `164885` (161.02 KiB) | `0.999x` | `24.281x` | `24.259x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `305781` (298.61 KiB) | `0.999x` | `13.093x` | `13.081x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `319842` (312.35 KiB) | `0.999x` | `12.517x` | `12.506x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `99928` (97.59 KiB) | `0.999x` | `40.065x` | `40.029x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `3488733` (3.33 MiB) | `2897914` (2.76 MiB) | `2.293x` | `1.204x` | `2.761x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `90420` (88.30 KiB) | `0.999x` | `44.278x` | `44.238x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `74086` (72.35 KiB) | `0.999x` | `54.040x` | `53.991x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `123785` (120.88 KiB) | `0.999x` | `32.343x` | `32.314x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4365` (4.26 KiB) | `0.999x` | `917.189x` | `916.380x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:85` | `1000000` | `13587860` (12.96 MiB) | `141140` (137.83 KiB) | `21463` (20.96 KiB) | `96.272x` | `6.576x` | `633.083x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003550` (3.82 MiB) | `7086` (6.92 KiB) | `0.999x` | `564.994x` | `564.493x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003555` (3.82 MiB) | `56707` (55.38 KiB) | `0.999x` | `70.601x` | `70.538x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003552` (3.82 MiB) | `8054` (7.87 KiB) | `0.999x` | `497.089x` | `496.648x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003535` (3.82 MiB) | `25218` (24.63 KiB) | `0.999x` | `158.757x` | `158.617x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004554` (7.63 MiB) | `693939` (677.67 KiB) | `0.999x` | `11.535x` | `11.528x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:137` | `1000000` | `27797671` (26.51 MiB) | `21052840` (20.08 MiB) | `5597770` (5.34 MiB) | `1.320x` | `3.761x` | `4.966x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003641` (3.82 MiB) | `3824533` (3.65 MiB) | `0.999x` | `1.047x` | `1.046x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4235` (4.14 KiB) | `0.999x` | `945.343x` | `944.510x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4235` (4.14 KiB) | `0.999x` | `945.343x` | `944.510x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4235` (4.14 KiB) | `0.999x` | `945.343x` | `944.510x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003582` (3.82 MiB) | `82520` (80.59 KiB) | `0.999x` | `48.517x` | `48.473x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4235` (4.14 KiB) | `0.999x` | `945.343x` | `944.510x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `1000000` (976.56 KiB) | `206687` (201.84 KiB) | `50911` (49.72 KiB) | `4.838x` | `4.060x` | `19.642x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `3381623` (3.22 MiB) | `2894162` (2.76 MiB) | `2.366x` | `1.168x` | `2.764x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `142981` (139.63 KiB) | `0.999x` | `28.001x` | `27.976x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `107962` (105.43 KiB) | `0.999x` | `37.083x` | `37.050x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `122561` (119.69 KiB) | `0.999x` | `32.666x` | `32.637x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `193426` (188.89 KiB) | `0.999x` | `20.698x` | `20.680x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `173604` (169.54 KiB) | `0.999x` | `23.062x` | `23.041x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003601` (3.82 MiB) | `426579` (416.58 KiB) | `0.999x` | `9.385x` | `9.377x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003578` (3.82 MiB) | `70834` (69.17 KiB) | `0.999x` | `56.521x` | `56.470x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4235` (4.14 KiB) | `0.999x` | `945.343x` | `944.510x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003572` (3.82 MiB) | `55199` (53.91 KiB) | `0.999x` | `72.530x` | `72.465x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `2001192` (1.91 MiB) | `333815` (325.99 KiB) | `64686` (63.17 KiB) | `5.995x` | `5.161x` | `30.937x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `3325142` (3.17 MiB) | `960655` (938.14 KiB) | `195260` (190.68 KiB) | `3.461x` | `4.920x` | `17.029x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `81281` (79.38 KiB) | `3685` (3.60 KiB) | `0.000x` | `22.057x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `81281` (79.38 KiB) | `3685` (3.60 KiB) | `0.000x` | `22.057x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4235` (4.14 KiB) | `0.999x` | `945.343x` | `944.510x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003543` (3.82 MiB) | `61520` (60.08 KiB) | `0.999x` | `65.077x` | `65.020x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `135323` (132.15 KiB) | `0.999x` | `29.585x` | `29.559x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003599` (3.82 MiB) | `333632` (325.81 KiB) | `0.999x` | `12.000x` | `11.989x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003641` (3.82 MiB) | `1245927` (1.19 MiB) | `0.999x` | `3.213x` | `3.210x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003644` (3.82 MiB) | `938049` (916.06 KiB) | `0.999x` | `4.268x` | `4.264x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003625` (3.82 MiB) | `550008` (537.12 KiB) | `0.999x` | `7.279x` | `7.273x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003530` (3.82 MiB) | `5347` (5.22 KiB) | `0.999x` | `748.743x` | `748.083x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `1024` (1.00 KiB) | `86573` (84.54 KiB) | `6191` (6.05 KiB) | `0.012x` | `13.984x` | `0.165x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004442` (7.63 MiB) | `5714` (5.58 KiB) | `0.999x` | `1400.847x` | `1400.070x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `81281` (79.38 KiB) | `3685` (3.60 KiB) | `0.000x` | `22.057x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `3000000` (2.86 MiB) | `85969` (83.95 KiB) | `6881` (6.72 KiB) | `34.896x` | `12.494x` | `435.983x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4235` (4.14 KiB) | `0.999x` | `945.343x` | `944.510x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `58030` (56.67 KiB) | `302723` (295.63 KiB) | `33915` (33.12 KiB) | `0.192x` | `8.926x` | `1.711x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `22051` (21.53 KiB) | `198619` (193.96 KiB) | `28555` (27.89 KiB) | `0.111x` | `6.956x` | `0.772x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `25445` (24.85 KiB) | `210341` (205.41 KiB) | `33669` (32.88 KiB) | `0.121x` | `6.247x` | `0.756x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `48191` (47.06 KiB) | `233349` (227.88 KiB) | `23375` (22.83 KiB) | `0.207x` | `9.983x` | `2.062x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `49433` (48.27 KiB) | `266935` (260.68 KiB) | `40312` (39.37 KiB) | `0.185x` | `6.622x` | `1.226x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `16873` (16.48 KiB) | `212683` (207.70 KiB) | `28148` (27.49 KiB) | `0.079x` | `7.556x` | `0.599x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `91870` (89.72 KiB) | `326459` (318.81 KiB) | `53944` (52.68 KiB) | `0.281x` | `6.052x` | `1.703x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `13001` (12.70 KiB) | `149171` (145.67 KiB) | `20761` (20.27 KiB) | `0.087x` | `7.185x` | `0.626x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `28101` (27.44 KiB) | `189308` (184.87 KiB) | `23140` (22.60 KiB) | `0.148x` | `8.181x` | `1.214x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `45607` (44.54 KiB) | `262231` (256.08 KiB) | `38567` (37.66 KiB) | `0.174x` | `6.799x` | `1.183x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003558` (3.82 MiB) | `21182` (20.69 KiB) | `0.999x` | `189.008x` | `188.840x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004553` (7.63 MiB) | `2841905` (2.71 MiB) | `0.999x` | `2.817x` | `2.815x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004555` (7.63 MiB) | `3580060` (3.41 MiB) | `0.999x` | `2.236x` | `2.235x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003545` (3.82 MiB) | `5643` (5.51 KiB) | `0.999x` | `709.471x` | `708.843x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00000.parquet`: `35854` rows, `2918370` file bytes (2.78 MiB), `27844924` physical bytes (26.55 MiB), `17530990` encoded bytes (16.72 MiB), `2886337` compressed data bytes (2.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00001.parquet`: `35947` rows, `2870918` file bytes (2.74 MiB), `27614679` physical bytes (26.34 MiB), `17320307` encoded bytes (16.52 MiB), `2838342` compressed data bytes (2.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00002.parquet`: `35808` rows, `2928977` file bytes (2.79 MiB), `28048338` physical bytes (26.75 MiB), `17555621` encoded bytes (16.74 MiB), `2895719` compressed data bytes (2.76 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00003.parquet`: `36222` rows, `2888639` file bytes (2.75 MiB), `28040591` physical bytes (26.74 MiB), `17501187` encoded bytes (16.69 MiB), `2856237` compressed data bytes (2.72 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00004.parquet`: `35574` rows, `2847286` file bytes (2.72 MiB), `27669292` physical bytes (26.39 MiB), `17234519` encoded bytes (16.44 MiB), `2814130` compressed data bytes (2.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00005.parquet`: `35764` rows, `2836744` file bytes (2.71 MiB), `27419611` physical bytes (26.15 MiB), `17264722` encoded bytes (16.46 MiB), `2804349` compressed data bytes (2.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00006.parquet`: `36205` rows, `2893461` file bytes (2.76 MiB), `28033241` physical bytes (26.73 MiB), `17465214` encoded bytes (16.66 MiB), `2861019` compressed data bytes (2.73 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00007.parquet`: `36253` rows, `2881706` file bytes (2.75 MiB), `28046264` physical bytes (26.75 MiB), `17602913` encoded bytes (16.79 MiB), `2848870` compressed data bytes (2.72 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00008.parquet`: `36213` rows, `2856099` file bytes (2.72 MiB), `28068902` physical bytes (26.77 MiB), `17518765` encoded bytes (16.71 MiB), `2823418` compressed data bytes (2.69 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00009.parquet`: `35844` rows, `2865483` file bytes (2.73 MiB), `27646513` physical bytes (26.37 MiB), `17360295` encoded bytes (16.56 MiB), `2833360` compressed data bytes (2.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00010.parquet`: `36475` rows, `2855261` file bytes (2.72 MiB), `28040066` physical bytes (26.74 MiB), `17541718` encoded bytes (16.73 MiB), `2822691` compressed data bytes (2.69 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00011.parquet`: `36384` rows, `2846724` file bytes (2.71 MiB), `28069038` physical bytes (26.77 MiB), `17542322` encoded bytes (16.73 MiB), `2814252` compressed data bytes (2.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00012.parquet`: `36309` rows, `2851802` file bytes (2.72 MiB), `27832129` physical bytes (26.54 MiB), `17499929` encoded bytes (16.69 MiB), `2819696` compressed data bytes (2.69 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00013.parquet`: `36189` rows, `2886980` file bytes (2.75 MiB), `28024166` physical bytes (26.73 MiB), `17509543` encoded bytes (16.70 MiB), `2854307` compressed data bytes (2.72 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00014.parquet`: `35369` rows, `3289294` file bytes (3.14 MiB), `24480490` physical bytes (23.35 MiB), `17401332` encoded bytes (16.60 MiB), `3258475` compressed data bytes (3.11 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00015.parquet`: `36072` rows, `3434987` file bytes (3.28 MiB), `25199026` physical bytes (24.03 MiB), `17836673` encoded bytes (17.01 MiB), `3404563` compressed data bytes (3.25 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00016.parquet`: `35261` rows, `3481760` file bytes (3.32 MiB), `23205241` physical bytes (22.13 MiB), `18308724` encoded bytes (17.46 MiB), `3451513` compressed data bytes (3.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00017.parquet`: `34788` rows, `3510051` file bytes (3.35 MiB), `22437829` physical bytes (21.40 MiB), `18644850` encoded bytes (17.78 MiB), `3479335` compressed data bytes (3.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00018.parquet`: `34274` rows, `3598379` file bytes (3.43 MiB), `22365122` physical bytes (21.33 MiB), `18565269` encoded bytes (17.71 MiB), `3568065` compressed data bytes (3.40 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00019.parquet`: `34893` rows, `3493744` file bytes (3.33 MiB), `22177451` physical bytes (21.15 MiB), `18538102` encoded bytes (17.68 MiB), `3463016` compressed data bytes (3.30 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00020.parquet`: `34654` rows, `3581346` file bytes (3.42 MiB), `22311138` physical bytes (21.28 MiB), `18561663` encoded bytes (17.70 MiB), `3550422` compressed data bytes (3.39 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00021.parquet`: `34994` rows, `3507477` file bytes (3.34 MiB), `22375286` physical bytes (21.34 MiB), `18616069` encoded bytes (17.75 MiB), `3476856` compressed data bytes (3.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00022.parquet`: `34961` rows, `3505603` file bytes (3.34 MiB), `22365615` physical bytes (21.33 MiB), `18552143` encoded bytes (17.69 MiB), `3475325` compressed data bytes (3.31 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00023.parquet`: `34399` rows, `3578055` file bytes (3.41 MiB), `22362121` physical bytes (21.33 MiB), `18696680` encoded bytes (17.83 MiB), `3547297` compressed data bytes (3.38 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00024.parquet`: `35152` rows, `3511280` file bytes (3.35 MiB), `22335285` physical bytes (21.30 MiB), `18531047` encoded bytes (17.67 MiB), `3481145` compressed data bytes (3.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00025.parquet`: `34734` rows, `3519185` file bytes (3.36 MiB), `22161071` physical bytes (21.13 MiB), `18458262` encoded bytes (17.60 MiB), `3488507` compressed data bytes (3.33 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00026.parquet`: `35152` rows, `3496500` file bytes (3.33 MiB), `22375859` physical bytes (21.34 MiB), `18505517` encoded bytes (17.65 MiB), `3466002` compressed data bytes (3.31 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00027.parquet`: `34944` rows, `3522020` file bytes (3.36 MiB), `22387057` physical bytes (21.35 MiB), `18618771` encoded bytes (17.76 MiB), `3491562` compressed data bytes (3.33 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00028.parquet`: `5312` rows, `537839` file bytes (525.23 KiB), `3462279` physical bytes (3.30 MiB), `2844442` encoded bytes (2.71 MiB), `522251` compressed data bytes (510.01 KiB)
