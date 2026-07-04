# ClickBench Parquet Experiment

- Started: `2026-07-03T23:32:45-04:00`
- Write elapsed: `11.501s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `593041822` (565.57 MiB)
- Compressed column data bytes after codec compression: `131362456` (125.28 MiB)
- Parquet file bytes: `132321005` (126.19 MiB)
- Physical/encoded ratio: `1.201x`
- Encoded/compressed-data ratio: `4.515x`
- Physical/compressed-data ratio: `5.423x`
- Physical/parquet-file ratio: `5.384x`
- Files: `31`

## Settings

- Compression: `snappy`
- Int encoding: `delta-binary-packed`
- String encoding: `plain`
- Date encoding: `plain`
- Timestamp encoding: `plain`
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
- Files read: `31`
- Elapsed: `7.064s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `7990565` (7.62 MiB) | `7990992` (7.62 MiB) | `1.001x` | `1.000x` | `1.001x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `198321` (193.67 KiB) | `91180` (89.04 KiB) | `20.169x` | `2.175x` | `43.869x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:505` | `1000000` | `138409995` (132.00 MiB) | `142874365` (136.26 MiB) | `20896027` (19.93 MiB) | `0.969x` | `6.837x` | `6.624x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `43905` (42.88 KiB) | `7102` (6.94 KiB) | `91.106x` | `6.182x` | `563.222x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004879` (7.63 MiB) | `4283366` (4.08 MiB) | `0.999x` | `1.869x` | `1.868x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204379` (199.59 KiB) | `0.999x` | `19.590x` | `19.571x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `46757` (45.66 KiB) | `7472` (7.30 KiB) | `85.549x` | `6.258x` | `535.332x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `3753464` (3.58 MiB) | `1054799` (1.01 MiB) | `1.066x` | `3.558x` | `3.792x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `1488256` (1.42 MiB) | `547680` (534.84 KiB) | `2.688x` | `2.717x` | `7.304x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `4399494` (4.20 MiB) | `945699` (923.53 KiB) | `1.818x` | `4.652x` | `8.459x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `43345` (42.33 KiB) | `6364` (6.21 KiB) | `92.283x` | `6.811x` | `628.536x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `847004` (827.15 KiB) | `266025` (259.79 KiB) | `4.723x` | `3.184x` | `15.036x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `686549` (670.46 KiB) | `282852` (276.22 KiB) | `5.826x` | `2.427x` | `14.142x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:342` | `1000000` | `88562192` (84.46 MiB) | `92648307` (88.36 MiB) | `20475103` (19.53 MiB) | `0.956x` | `4.525x` | `4.325x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:300` | `1000000` | `79583339` (75.90 MiB) | `83646070` (79.77 MiB) | `19048392` (18.17 MiB) | `0.951x` | `4.391x` | `4.178x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `234163` (228.67 KiB) | `138319` (135.08 KiB) | `17.082x` | `1.693x` | `28.919x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `1864248` (1.78 MiB) | `591972` (578.10 KiB) | `2.146x` | `3.149x` | `6.757x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `1368894` (1.31 MiB) | `393827` (384.60 KiB) | `2.922x` | `3.476x` | `10.157x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `1078898` (1.03 MiB) | `242209` (236.53 KiB) | `3.707x` | `4.454x` | `16.515x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `723909` (706.94 KiB) | `127469` (124.48 KiB) | `5.526x` | `5.679x` | `31.380x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `1284158` (1.22 MiB) | `478596` (467.38 KiB) | `3.115x` | `2.683x` | `8.358x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `1343351` (1.28 MiB) | `423148` (413.23 KiB) | `2.978x` | `3.175x` | `9.453x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `552631` (539.68 KiB) | `163848` (160.01 KiB) | `7.238x` | `3.373x` | `24.413x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `400863` (391.47 KiB) | `113743` (111.08 KiB) | `9.978x` | `3.524x` | `35.167x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `524560` (512.27 KiB) | `235195` (229.68 KiB) | `7.625x` | `2.230x` | `17.007x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `3354477` (3.20 MiB) | `7358096` (7.02 MiB) | `536849` (524.27 KiB) | `0.456x` | `13.706x` | `6.248x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `159940` (156.19 KiB) | `52029` (50.81 KiB) | `25.009x` | `3.074x` | `76.880x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `197681` (193.05 KiB) | `48142` (47.01 KiB) | `20.235x` | `4.106x` | `83.088x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `662797` (647.26 KiB) | `301924` (294.85 KiB) | `6.035x` | `2.195x` | `13.248x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `3767530` (3.59 MiB) | `7771138` (7.41 MiB) | `467235` (456.28 KiB) | `0.485x` | `16.632x` | `8.063x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `46885` (45.79 KiB) | `9263` (9.05 KiB) | `85.315x` | `5.062x` | `431.826x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `47865` (46.74 KiB) | `9956` (9.72 KiB) | `83.568x` | `4.808x` | `401.768x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `114927` (112.23 KiB) | `42771` (41.77 KiB) | `34.805x` | `2.687x` | `93.521x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `144775` (141.38 KiB) | `48323` (47.19 KiB) | `27.629x` | `2.996x` | `82.776x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `81583` (79.67 KiB) | `4085142` (3.90 MiB) | `220093` (214.93 KiB) | `0.020x` | `18.561x` | `0.371x` | `81583` (79.67 KiB) |
| `Params` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `0` (0 B) | `4002377` (3.82 MiB) | `202741` (197.99 KiB) | `0.000x` | `19.741x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `2658357` (2.54 MiB) | `929624` (907.84 KiB) | `1.505x` | `2.860x` | `4.303x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `496636` (485.00 KiB) | `292901` (286.04 KiB) | `8.054x` | `1.696x` | `13.656x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `379344` (370.45 KiB) | `169013` (165.05 KiB) | `10.545x` | `2.244x` | `23.667x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `3528017` (3.36 MiB) | `7535473` (7.19 MiB) | `1097546` (1.05 MiB) | `0.468x` | `6.866x` | `3.214x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `202872` (198.12 KiB) | `53700` (52.44 KiB) | `19.717x` | `3.778x` | `74.488x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `196941` (192.33 KiB) | `119273` (116.48 KiB) | `20.311x` | `1.651x` | `33.537x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `1358300` (1.30 MiB) | `647041` (631.88 KiB) | `2.945x` | `2.099x` | `6.182x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `1249633` (1.19 MiB) | `617733` (603.25 KiB) | `3.201x` | `2.023x` | `6.475x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `978467` (955.53 KiB) | `252815` (246.89 KiB) | `4.088x` | `3.870x` | `15.822x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004874` (7.63 MiB) | `4229550` (4.03 MiB) | `0.999x` | `1.893x` | `1.891x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `437563` (427.31 KiB) | `152008` (148.45 KiB) | `9.142x` | `2.879x` | `26.314x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `207724` (202.86 KiB) | `98686` (96.37 KiB) | `19.256x` | `2.105x` | `40.533x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `1864414` (1.78 MiB) | `393080` (383.87 KiB) | `2.145x` | `4.743x` | `10.176x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `43822` (42.79 KiB) | `6596` (6.44 KiB) | `91.278x` | `6.644x` | `606.428x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:91` | `1000000` | `13587860` (12.96 MiB) | `17595561` (16.78 MiB) | `917471` (895.97 KiB) | `0.772x` | `19.178x` | `14.810x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `96494` (94.23 KiB) | `15517` (15.15 KiB) | `41.453x` | `6.219x` | `257.782x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `140983` (137.68 KiB) | `59998` (58.59 KiB) | `28.372x` | `2.350x` | `66.669x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `53898` (52.63 KiB) | `12414` (12.12 KiB) | `74.214x` | `4.342x` | `322.217x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `64177` (62.67 KiB) | `23359` (22.81 KiB) | `62.328x` | `2.747x` | `171.240x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `7478240` (7.13 MiB) | `1783283` (1.70 MiB) | `1.070x` | `4.194x` | `4.486x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:140` | `1000000` | `27797671` (26.51 MiB) | `31857520` (30.38 MiB) | `7043102` (6.72 MiB) | `0.873x` | `4.523x` | `3.947x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `3953571` (3.77 MiB) | `3761024` (3.59 MiB) | `1.012x` | `1.051x` | `1.064x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `43345` (42.33 KiB) | `6364` (6.21 KiB) | `92.283x` | `6.811x` | `628.536x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `43345` (42.33 KiB) | `6364` (6.21 KiB) | `92.283x` | `6.811x` | `628.536x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `43345` (42.33 KiB) | `6364` (6.21 KiB) | `92.283x` | `6.811x` | `628.536x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `168631` (164.68 KiB) | `83289` (81.34 KiB) | `23.720x` | `2.025x` | `48.026x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `43345` (42.33 KiB) | `6364` (6.21 KiB) | `92.283x` | `6.811x` | `628.536x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `1000000` (976.56 KiB) | `5003113` (4.77 MiB) | `294445` (287.54 KiB) | `0.200x` | `16.992x` | `3.396x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004871` (7.63 MiB) | `4284344` (4.09 MiB) | `0.999x` | `1.868x` | `1.867x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `743924` (726.49 KiB) | `287674` (280.93 KiB) | `5.377x` | `2.586x` | `13.905x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `254323` (248.36 KiB) | `136190` (133.00 KiB) | `15.728x` | `1.867x` | `29.371x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `329739` (322.01 KiB) | `168630` (164.68 KiB) | `12.131x` | `1.955x` | `23.721x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `1588250` (1.51 MiB) | `584217` (570.52 KiB) | `2.518x` | `2.719x` | `6.847x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `842705` (822.95 KiB) | `386667` (377.60 KiB) | `4.747x` | `2.179x` | `10.345x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `3797920` (3.62 MiB) | `1111214` (1.06 MiB) | `1.053x` | `3.418x` | `3.600x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `274785` (268.34 KiB) | `130627` (127.57 KiB) | `14.557x` | `2.104x` | `30.622x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `43340` (42.32 KiB) | `6410` (6.26 KiB) | `92.293x` | `6.761x` | `624.025x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `116473` (113.74 KiB) | `68934` (67.32 KiB) | `34.343x` | `1.690x` | `58.027x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `2001192` (1.91 MiB) | `6004669` (5.73 MiB) | `321298` (313.77 KiB) | `0.333x` | `18.689x` | `6.228x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `3325142` (3.17 MiB) | `7328761` (6.99 MiB) | `457244` (446.53 KiB) | `0.454x` | `16.028x` | `7.272x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `0` (0 B) | `4002377` (3.82 MiB) | `202741` (197.99 KiB) | `0.000x` | `19.741x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `0` (0 B) | `4002377` (3.82 MiB) | `202741` (197.99 KiB) | `0.000x` | `19.741x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `43345` (42.33 KiB) | `6364` (6.21 KiB) | `92.283x` | `6.811x` | `628.536x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `173548` (169.48 KiB) | `111934` (109.31 KiB) | `23.048x` | `1.550x` | `35.735x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `808003` (789.07 KiB) | `342456` (334.43 KiB) | `4.950x` | `2.359x` | `11.680x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `1091931` (1.04 MiB) | `626409` (611.73 KiB) | `3.663x` | `1.743x` | `6.386x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `1705430` (1.63 MiB) | `1522423` (1.45 MiB) | `2.345x` | `1.120x` | `2.627x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `1503555` (1.43 MiB) | `1305051` (1.24 MiB) | `2.660x` | `1.152x` | `3.065x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `1549374` (1.48 MiB) | `1014134` (990.37 KiB) | `2.582x` | `1.528x` | `3.944x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `47649` (46.53 KiB) | `8423` (8.23 KiB) | `83.947x` | `5.657x` | `474.890x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `1024` (1.00 KiB) | `4004127` (3.82 MiB) | `204273` (199.49 KiB) | `0.000x` | `19.602x` | `0.005x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `44321` (43.28 KiB) | `7340` (7.17 KiB) | `180.501x` | `6.038x` | `1089.918x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `0` (0 B) | `4002377` (3.82 MiB) | `202741` (197.99 KiB) | `0.000x` | `19.741x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `3000000` (2.86 MiB) | `7003594` (6.68 MiB) | `354312` (346.01 KiB) | `0.428x` | `19.767x` | `8.467x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `43345` (42.33 KiB) | `6364` (6.21 KiB) | `92.283x` | `6.811x` | `628.536x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `58030` (56.67 KiB) | `4062845` (3.87 MiB) | `218519` (213.40 KiB) | `0.014x` | `18.593x` | `0.266x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `22051` (21.53 KiB) | `4025894` (3.84 MiB) | `214778` (209.74 KiB) | `0.005x` | `18.744x` | `0.103x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `25445` (24.85 KiB) | `4029299` (3.84 MiB) | `219169` (214.03 KiB) | `0.006x` | `18.384x` | `0.116x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `48191` (47.06 KiB) | `4052047` (3.86 MiB) | `214366` (209.34 KiB) | `0.012x` | `18.902x` | `0.225x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `49433` (48.27 KiB) | `4053496` (3.87 MiB) | `219162` (214.03 KiB) | `0.012x` | `18.495x` | `0.226x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `16873` (16.48 KiB) | `4020162` (3.83 MiB) | `214408` (209.38 KiB) | `0.004x` | `18.750x` | `0.079x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `91870` (89.72 KiB) | `4097838` (3.91 MiB) | `230232` (224.84 KiB) | `0.022x` | `17.799x` | `0.399x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `13001` (12.70 KiB) | `4016770` (3.83 MiB) | `212207` (207.23 KiB) | `0.003x` | `18.929x` | `0.061x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `28101` (27.44 KiB) | `4033136` (3.85 MiB) | `214703` (209.67 KiB) | `0.007x` | `18.785x` | `0.131x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `45607` (44.54 KiB) | `4048506` (3.86 MiB) | `221714` (216.52 KiB) | `0.011x` | `18.260x` | `0.206x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `92519` (90.35 KiB) | `28610` (27.94 KiB) | `43.234x` | `3.234x` | `139.811x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `8120117` (7.74 MiB) | `5099000` (4.86 MiB) | `0.985x` | `1.592x` | `1.569x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `8122623` (7.75 MiB) | `5949824` (5.67 MiB) | `0.985x` | `1.365x` | `1.345x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `69050` (67.43 KiB) | `10601` (10.35 KiB) | `57.929x` | `6.514x` | `377.323x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00000.parquet`: `33130` rows, `4218444` file bytes (4.02 MiB), `25731463` physical bytes (24.54 MiB), `21792975` encoded bytes (20.78 MiB), `4186544` compressed data bytes (3.99 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00001.parquet`: `33194` rows, `4152926` file bytes (3.96 MiB), `25525840` physical bytes (24.34 MiB), `21498803` encoded bytes (20.50 MiB), `4120650` compressed data bytes (3.93 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00002.parquet`: `32722` rows, `4152390` file bytes (3.96 MiB), `25609560` physical bytes (24.42 MiB), `21680579` encoded bytes (20.68 MiB), `4119620` compressed data bytes (3.93 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00003.parquet`: `32967` rows, `4142160` file bytes (3.95 MiB), `25562830` physical bytes (24.38 MiB), `21571739` encoded bytes (20.57 MiB), `4109834` compressed data bytes (3.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00004.parquet`: `33197` rows, `4155774` file bytes (3.96 MiB), `25781925` physical bytes (24.59 MiB), `21769915` encoded bytes (20.76 MiB), `4122684` compressed data bytes (3.93 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00005.parquet`: `33440` rows, `4111001` file bytes (3.92 MiB), `25803957` physical bytes (24.61 MiB), `21713425` encoded bytes (20.71 MiB), `4078660` compressed data bytes (3.89 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00006.parquet`: `32981` rows, `4120022` file bytes (3.93 MiB), `25374894` physical bytes (24.20 MiB), `21429696` encoded bytes (20.44 MiB), `4087924` compressed data bytes (3.90 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00007.parquet`: `33232` rows, `4152781` file bytes (3.96 MiB), `25783969` physical bytes (24.59 MiB), `21763380` encoded bytes (20.76 MiB), `4120321` compressed data bytes (3.93 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00008.parquet`: `33078` rows, `4115319` file bytes (3.92 MiB), `25585251` physical bytes (24.40 MiB), `21563069` encoded bytes (20.56 MiB), `4082486` compressed data bytes (3.89 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00009.parquet`: `33194` rows, `4103005` file bytes (3.91 MiB), `25609390` physical bytes (24.42 MiB), `21620564` encoded bytes (20.62 MiB), `4070682` compressed data bytes (3.88 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00010.parquet`: `33079` rows, `4127081` file bytes (3.94 MiB), `25578742` physical bytes (24.39 MiB), `21556018` encoded bytes (20.56 MiB), `4095153` compressed data bytes (3.91 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00011.parquet`: `33301` rows, `4073715` file bytes (3.88 MiB), `25594966` physical bytes (24.41 MiB), `21538066` encoded bytes (20.54 MiB), `4041531` compressed data bytes (3.85 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00012.parquet`: `33239` rows, `4079085` file bytes (3.89 MiB), `25629228` physical bytes (24.44 MiB), `21574380` encoded bytes (20.57 MiB), `4046860` compressed data bytes (3.86 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00013.parquet`: `33321` rows, `4128858` file bytes (3.94 MiB), `25588583` physical bytes (24.40 MiB), `21575288` encoded bytes (20.58 MiB), `4096912` compressed data bytes (3.91 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00014.parquet`: `32774` rows, `4043525` file bytes (3.86 MiB), `25264913` physical bytes (24.09 MiB), `21295006` encoded bytes (20.31 MiB), `4011406` compressed data bytes (3.83 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00015.parquet`: `33241` rows, `4443725` file bytes (4.24 MiB), `23839307` physical bytes (22.73 MiB), `20098517` encoded bytes (19.17 MiB), `4412591` compressed data bytes (4.21 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00016.parquet`: `33312` rows, `4586973` file bytes (4.37 MiB), `22906386` physical bytes (21.85 MiB), `19265976` encoded bytes (18.37 MiB), `4556729` compressed data bytes (4.35 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00017.parquet`: `32450` rows, `4595153` file bytes (4.38 MiB), `22428721` physical bytes (21.39 MiB), `18791525` encoded bytes (17.92 MiB), `4564980` compressed data bytes (4.35 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00018.parquet`: `33064` rows, `4607214` file bytes (4.39 MiB), `21083297` physical bytes (20.11 MiB), `17127854` encoded bytes (16.33 MiB), `4576942` compressed data bytes (4.36 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00019.parquet`: `32811` rows, `4647895` file bytes (4.43 MiB), `21251948` physical bytes (20.27 MiB), `17296785` encoded bytes (16.50 MiB), `4617396` compressed data bytes (4.40 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00020.parquet`: `32280` rows, `4620804` file bytes (4.41 MiB), `20806036` physical bytes (19.84 MiB), `16947685` encoded bytes (16.16 MiB), `4590145` compressed data bytes (4.38 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00021.parquet`: `33202` rows, `4650738` file bytes (4.44 MiB), `21202785` physical bytes (20.22 MiB), `17201560` encoded bytes (16.40 MiB), `4620075` compressed data bytes (4.41 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00022.parquet`: `32798` rows, `4672436` file bytes (4.46 MiB), `21163747` physical bytes (20.18 MiB), `17245565` encoded bytes (16.45 MiB), `4641912` compressed data bytes (4.43 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00023.parquet`: `32994` rows, `4631227` file bytes (4.42 MiB), `21250124` physical bytes (20.27 MiB), `17315587` encoded bytes (16.51 MiB), `4600479` compressed data bytes (4.39 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00024.parquet`: `33085` rows, `4662327` file bytes (4.45 MiB), `21225284` physical bytes (20.24 MiB), `17286450` encoded bytes (16.49 MiB), `4631752` compressed data bytes (4.42 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00025.parquet`: `32822` rows, `4681011` file bytes (4.46 MiB), `21193343` physical bytes (20.21 MiB), `17245544` encoded bytes (16.45 MiB), `4650682` compressed data bytes (4.44 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00026.parquet`: `33381` rows, `4608422` file bytes (4.39 MiB), `21184091` physical bytes (20.20 MiB), `17245189` encoded bytes (16.45 MiB), `4578349` compressed data bytes (4.37 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00027.parquet`: `32690` rows, `4594828` file bytes (4.38 MiB), `20856054` physical bytes (19.89 MiB), `16988034` encoded bytes (16.20 MiB), `4564232` compressed data bytes (4.35 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00028.parquet`: `33076` rows, `4575170` file bytes (4.36 MiB), `21079904` physical bytes (20.10 MiB), `17159189` encoded bytes (16.36 MiB), `4544721` compressed data bytes (4.33 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00029.parquet`: `33175` rows, `4627993` file bytes (4.41 MiB), `21229852` physical bytes (20.25 MiB), `17286589` encoded bytes (16.49 MiB), `4597504` compressed data bytes (4.38 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain/part-00030.parquet`: `8770` rows, `1239003` file bytes (1.18 MiB), `5672234` physical bytes (5.41 MiB), `4596870` encoded bytes (4.38 MiB), `1222700` compressed data bytes (1.17 MiB)
