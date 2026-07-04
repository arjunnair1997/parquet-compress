# ClickBench Parquet Experiment

- Started: `2026-07-03T23:36:47-04:00`
- Write elapsed: `11.739s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `807720241` (770.30 MiB)
- Compressed column data bytes after codec compression: `87960151` (83.89 MiB)
- Parquet file bytes: `88859290` (84.74 MiB)
- Physical/encoded ratio: `0.882x`
- Encoded/compressed-data ratio: `9.183x`
- Physical/compressed-data ratio: `8.099x`
- Physical/parquet-file ratio: `8.017x`
- Files: `29`

## Settings

- Compression: `zstd-3`
- Int encoding: `plain`
- String encoding: `plain`
- Date encoding: `delta-binary-packed`
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
- Elapsed: `7.413s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004555` (7.63 MiB) | `8005360` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `64046` (62.54 KiB) | `0.999x` | `62.511x` | `62.455x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:526` | `1000000` | `138409995` (132.00 MiB) | `142871744` (136.25 MiB) | `13947549` (13.30 MiB) | `0.969x` | `10.244x` | `9.924x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003530` (3.82 MiB) | `4921` (4.81 KiB) | `0.999x` | `813.560x` | `812.843x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `3368626` (3.21 MiB) | `2886499` (2.75 MiB) | `2.375x` | `1.167x` | `2.772x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `4000000` (3.81 MiB) | `50747` (49.56 KiB) | `6336` (6.19 KiB) | `78.822x` | `8.009x` | `631.313x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003526` (3.82 MiB) | `4934` (4.82 KiB) | `0.999x` | `811.416x` | `810.701x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003597` (3.82 MiB) | `408132` (398.57 KiB) | `0.999x` | `9.810x` | `9.801x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `190966` (186.49 KiB) | `0.999x` | `20.965x` | `20.946x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004553` (7.63 MiB) | `618142` (603.65 KiB) | `0.999x` | `12.949x` | `12.942x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4235` (4.14 KiB) | `0.999x` | `945.343x` | `944.510x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `106218` (103.73 KiB) | `0.999x` | `37.692x` | `37.658x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `134037` (130.90 KiB) | `0.999x` | `29.869x` | `29.843x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:352` | `1000000` | `88562192` (84.46 MiB) | `92652139` (88.36 MiB) | `15302500` (14.59 MiB) | `0.956x` | `6.055x` | `5.787x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:326` | `1000000` | `79583339` (75.90 MiB) | `83646858` (79.77 MiB) | `14215690` (13.56 MiB) | `0.951x` | `5.884x` | `5.598x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `178733` (174.54 KiB) | `0.999x` | `22.400x` | `22.380x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003581` (3.82 MiB) | `275155` (268.71 KiB) | `0.999x` | `14.550x` | `14.537x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `231260` (225.84 KiB) | `0.999x` | `17.312x` | `17.297x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `87558` (85.51 KiB) | `0.999x` | `45.725x` | `45.684x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `48095` (46.97 KiB) | `0.999x` | `83.243x` | `83.169x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003581` (3.82 MiB) | `187130` (182.74 KiB) | `0.999x` | `21.395x` | `21.376x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `186039` (181.68 KiB) | `0.999x` | `21.520x` | `21.501x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `82431` (80.50 KiB) | `0.999x` | `48.569x` | `48.525x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `53323` (52.07 KiB) | `0.999x` | `75.082x` | `75.015x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `128013` (125.01 KiB) | `0.999x` | `31.275x` | `31.247x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3354477` (3.20 MiB) | `7357825` (7.02 MiB) | `246378` (240.60 KiB) | `0.456x` | `29.864x` | `13.615x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `26171` (25.56 KiB) | `0.999x` | `152.978x` | `152.841x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `24688` (24.11 KiB) | `0.999x` | `162.167x` | `162.022x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003582` (3.82 MiB) | `154767` (151.14 KiB) | `0.999x` | `25.868x` | `25.845x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3767530` (3.59 MiB) | `7770884` (7.41 MiB) | `137025` (133.81 KiB) | `0.485x` | `56.711x` | `27.495x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003534` (3.82 MiB) | `6099` (5.96 KiB) | `0.999x` | `656.425x` | `655.845x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003541` (3.82 MiB) | `6545` (6.39 KiB) | `0.999x` | `611.695x` | `611.154x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `28835` (28.16 KiB) | `0.999x` | `138.845x` | `138.720x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `22468` (21.94 KiB) | `0.999x` | `178.191x` | `178.031x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `81583` (79.67 KiB) | `4084877` (3.90 MiB) | `22583` (22.05 KiB) | `0.020x` | `180.883x` | `3.613x` | `81583` (79.67 KiB) |
| `Params` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002160` (3.82 MiB) | `2867` (2.80 KiB) | `0.000x` | `1395.940x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003592` (3.82 MiB) | `323834` (316.24 KiB) | `0.999x` | `12.363x` | `12.352x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `288237` (281.48 KiB) | `0.999x` | `13.890x` | `13.877x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `101579` (99.20 KiB) | `0.999x` | `39.414x` | `39.378x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3528017` (3.36 MiB) | `7534788` (7.19 MiB) | `720381` (703.50 KiB) | `0.468x` | `10.459x` | `4.897x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003578` (3.82 MiB) | `30691` (29.97 KiB) | `0.999x` | `130.448x` | `130.331x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `164870` (161.01 KiB) | `0.999x` | `24.283x` | `24.262x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `305821` (298.65 KiB) | `0.999x` | `13.091x` | `13.080x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003591` (3.82 MiB) | `319906` (312.41 KiB) | `0.999x` | `12.515x` | `12.504x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `99890` (97.55 KiB) | `0.999x` | `40.080x` | `40.044x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `3489451` (3.33 MiB) | `2897203` (2.76 MiB) | `2.293x` | `1.204x` | `2.761x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `90792` (88.66 KiB) | `0.999x` | `44.096x` | `44.057x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `74075` (72.34 KiB) | `0.999x` | `54.048x` | `53.999x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `123784` (120.88 KiB) | `0.999x` | `32.343x` | `32.314x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003530` (3.82 MiB) | `4366` (4.26 KiB) | `0.999x` | `916.979x` | `916.170x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:85` | `1000000` | `13587860` (12.96 MiB) | `17594797` (16.78 MiB) | `14577` (14.24 KiB) | `0.772x` | `1207.025x` | `932.144x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003549` (3.82 MiB) | `7095` (6.93 KiB) | `0.999x` | `564.278x` | `563.777x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003555` (3.82 MiB) | `56797` (55.47 KiB) | `0.999x` | `70.489x` | `70.426x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003550` (3.82 MiB) | `8030` (7.84 KiB) | `0.999x` | `498.574x` | `498.132x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003536` (3.82 MiB) | `25208` (24.62 KiB) | `0.999x` | `158.820x` | `158.680x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004551` (7.63 MiB) | `693947` (677.68 KiB) | `0.999x` | `11.535x` | `11.528x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:137` | `1000000` | `27797671` (26.51 MiB) | `31859425` (30.38 MiB) | `5325158` (5.08 MiB) | `0.873x` | `5.983x` | `5.220x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003642` (3.82 MiB) | `3827434` (3.65 MiB) | `0.999x` | `1.046x` | `1.045x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4235` (4.14 KiB) | `0.999x` | `945.343x` | `944.510x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4235` (4.14 KiB) | `0.999x` | `945.343x` | `944.510x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4235` (4.14 KiB) | `0.999x` | `945.343x` | `944.510x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `82652` (80.71 KiB) | `0.999x` | `48.439x` | `48.396x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4235` (4.14 KiB) | `0.999x` | `945.343x` | `944.510x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `1000000` (976.56 KiB) | `5002904` (4.77 MiB) | `29086` (28.40 KiB) | `0.200x` | `172.004x` | `34.381x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `3382549` (3.23 MiB) | `2894727` (2.76 MiB) | `2.365x` | `1.169x` | `2.764x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `143019` (139.67 KiB) | `0.999x` | `27.993x` | `27.968x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `107841` (105.31 KiB) | `0.999x` | `37.125x` | `37.092x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `122656` (119.78 KiB) | `0.999x` | `32.641x` | `32.612x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `193541` (189.00 KiB) | `0.999x` | `20.686x` | `20.667x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `173730` (169.66 KiB) | `0.999x` | `23.045x` | `23.024x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003604` (3.82 MiB) | `426629` (416.63 KiB) | `0.999x` | `9.384x` | `9.376x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003576` (3.82 MiB) | `70826` (69.17 KiB) | `0.999x` | `56.527x` | `56.476x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003530` (3.82 MiB) | `4237` (4.14 KiB) | `0.999x` | `944.897x` | `944.064x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003572` (3.82 MiB) | `55130` (53.84 KiB) | `0.999x` | `72.621x` | `72.556x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `2001192` (1.91 MiB) | `6004431` (5.73 MiB) | `32231` (31.48 KiB) | `0.333x` | `186.294x` | `62.089x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3325142` (3.17 MiB) | `7328500` (6.99 MiB) | `122232` (119.37 KiB) | `0.454x` | `59.956x` | `27.204x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002160` (3.82 MiB) | `2867` (2.80 KiB) | `0.000x` | `1395.940x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002160` (3.82 MiB) | `2867` (2.80 KiB) | `0.000x` | `1395.940x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4235` (4.14 KiB) | `0.999x` | `945.343x` | `944.510x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003541` (3.82 MiB) | `61437` (60.00 KiB) | `0.999x` | `65.165x` | `65.107x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `135448` (132.27 KiB) | `0.999x` | `29.558x` | `29.532x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003594` (3.82 MiB) | `333245` (325.43 KiB) | `0.999x` | `12.014x` | `12.003x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003643` (3.82 MiB) | `1245931` (1.19 MiB) | `0.999x` | `3.213x` | `3.210x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003641` (3.82 MiB) | `937808` (915.83 KiB) | `0.999x` | `4.269x` | `4.265x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003621` (3.82 MiB) | `549955` (537.07 KiB) | `0.999x` | `7.280x` | `7.273x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `5327` (5.20 KiB) | `0.999x` | `751.554x` | `750.892x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `1024` (1.00 KiB) | `4004021` (3.82 MiB) | `4887` (4.77 KiB) | `0.000x` | `819.321x` | `0.210x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004438` (7.63 MiB) | `5710` (5.58 KiB) | `0.999x` | `1401.828x` | `1401.051x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002160` (3.82 MiB) | `2867` (2.80 KiB) | `0.000x` | `1395.940x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3000000` (2.86 MiB) | `7003302` (6.68 MiB) | `5260` (5.14 KiB) | `0.428x` | `1331.426x` | `570.342x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4235` (4.14 KiB) | `0.999x` | `945.343x` | `944.510x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `58030` (56.67 KiB) | `4062526` (3.87 MiB) | `19174` (18.72 KiB) | `0.014x` | `211.877x` | `3.026x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `22051` (21.53 KiB) | `4025606` (3.84 MiB) | `16531` (16.14 KiB) | `0.005x` | `243.519x` | `1.334x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `25445` (24.85 KiB) | `4029083` (3.84 MiB) | `18953` (18.51 KiB) | `0.006x` | `212.583x` | `1.343x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `48191` (47.06 KiB) | `4051766` (3.86 MiB) | `13314` (13.00 KiB) | `0.012x` | `304.324x` | `3.620x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `49433` (48.27 KiB) | `4053228` (3.87 MiB) | `21088` (20.59 KiB) | `0.012x` | `192.205x` | `2.344x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `16873` (16.48 KiB) | `4019934` (3.83 MiB) | `17078` (16.68 KiB) | `0.004x` | `235.387x` | `0.988x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `91870` (89.72 KiB) | `4097408` (3.91 MiB) | `29800` (29.10 KiB) | `0.022x` | `137.497x` | `3.083x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `13001` (12.70 KiB) | `4016539` (3.83 MiB) | `13961` (13.63 KiB) | `0.003x` | `287.697x` | `0.931x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `28101` (27.44 KiB) | `4032753` (3.85 MiB) | `15691` (15.32 KiB) | `0.007x` | `257.011x` | `1.791x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `45607` (44.54 KiB) | `4048288` (3.86 MiB) | `28367` (27.70 KiB) | `0.011x` | `142.711x` | `1.608x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003558` (3.82 MiB) | `21191` (20.69 KiB) | `0.999x` | `188.927x` | `188.759x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004551` (7.63 MiB) | `2842017` (2.71 MiB) | `0.999x` | `2.817x` | `2.815x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004555` (7.63 MiB) | `3580398` (3.41 MiB) | `0.999x` | `2.236x` | `2.234x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003544` (3.82 MiB) | `5629` (5.50 KiB) | `0.999x` | `711.235x` | `710.606x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00000.parquet`: `35854` rows, `2901203` file bytes (2.77 MiB), `27844924` physical bytes (26.55 MiB), `31280268` encoded bytes (29.83 MiB), `2869157` compressed data bytes (2.74 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00001.parquet`: `35947` rows, `2855878` file bytes (2.72 MiB), `27614679` physical bytes (26.34 MiB), `31054916` encoded bytes (29.62 MiB), `2823290` compressed data bytes (2.69 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00002.parquet`: `35804` rows, `2914435` file bytes (2.78 MiB), `28045024` physical bytes (26.75 MiB), `31476393` encoded bytes (30.02 MiB), `2881158` compressed data bytes (2.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00003.parquet`: `36223` rows, `2871807` file bytes (2.74 MiB), `28041589` physical bytes (26.74 MiB), `31504968` encoded bytes (30.05 MiB), `2839374` compressed data bytes (2.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00004.parquet`: `35583` rows, `2826051` file bytes (2.70 MiB), `27676556` physical bytes (26.39 MiB), `31083026` encoded bytes (29.64 MiB), `2792881` compressed data bytes (2.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00005.parquet`: `35769` rows, `2819031` file bytes (2.69 MiB), `27421909` physical bytes (26.15 MiB), `30838267` encoded bytes (29.41 MiB), `2786619` compressed data bytes (2.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00006.parquet`: `36202` rows, `2876565` file bytes (2.74 MiB), `28032888` physical bytes (26.73 MiB), `31500059` encoded bytes (30.04 MiB), `2844108` compressed data bytes (2.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00007.parquet`: `36255` rows, `2862743` file bytes (2.73 MiB), `28048091` physical bytes (26.75 MiB), `31516981` encoded bytes (30.06 MiB), `2829887` compressed data bytes (2.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00008.parquet`: `36216` rows, `2841189` file bytes (2.71 MiB), `28069629` physical bytes (26.77 MiB), `31532315` encoded bytes (30.07 MiB), `2808492` compressed data bytes (2.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00009.parquet`: `35848` rows, `2846710` file bytes (2.71 MiB), `27649376` physical bytes (26.37 MiB), `31076889` encoded bytes (29.64 MiB), `2814570` compressed data bytes (2.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00010.parquet`: `36464` rows, `2842463` file bytes (2.71 MiB), `28033620` physical bytes (26.73 MiB), `31516808` encoded bytes (30.06 MiB), `2809877` compressed data bytes (2.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00011.parquet`: `36386` rows, `2828206` file bytes (2.70 MiB), `28069654` physical bytes (26.77 MiB), `31550614` encoded bytes (30.09 MiB), `2795722` compressed data bytes (2.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00012.parquet`: `36306` rows, `2835067` file bytes (2.70 MiB), `27829887` physical bytes (26.54 MiB), `31300360` encoded bytes (29.85 MiB), `2802948` compressed data bytes (2.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00013.parquet`: `36202` rows, `2867171` file bytes (2.73 MiB), `28030660` physical bytes (26.73 MiB), `31493109` encoded bytes (30.03 MiB), `2834483` compressed data bytes (2.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00014.parquet`: `35358` rows, `3278773` file bytes (3.13 MiB), `24474438` physical bytes (23.34 MiB), `27848407` encoded bytes (26.56 MiB), `3247950` compressed data bytes (3.10 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00015.parquet`: `36048` rows, `3439080` file bytes (3.28 MiB), `25184942` physical bytes (24.02 MiB), `28620896` encoded bytes (27.30 MiB), `3408672` compressed data bytes (3.25 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00016.parquet`: `35303` rows, `3446540` file bytes (3.29 MiB), `23228130` physical bytes (22.15 MiB), `26586306` encoded bytes (25.35 MiB), `3416282` compressed data bytes (3.26 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00017.parquet`: `34855` rows, `3464486` file bytes (3.30 MiB), `22488678` physical bytes (21.45 MiB), `25796960` encoded bytes (24.60 MiB), `3433759` compressed data bytes (3.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00018.parquet`: `34370` rows, `3551783` file bytes (3.39 MiB), `22419656` physical bytes (21.38 MiB), `25676452` encoded bytes (24.49 MiB), `3521454` compressed data bytes (3.36 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00019.parquet`: `34956` rows, `3444102` file bytes (3.28 MiB), `22245512` physical bytes (21.21 MiB), `25562581` encoded bytes (24.38 MiB), `3413366` compressed data bytes (3.26 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00020.parquet`: `34542` rows, `3508475` file bytes (3.35 MiB), `22211348` physical bytes (21.18 MiB), `25490405` encoded bytes (24.31 MiB), `3477578` compressed data bytes (3.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00021.parquet`: `35096` rows, `3458135` file bytes (3.30 MiB), `22434026` physical bytes (21.39 MiB), `25763073` encoded bytes (24.57 MiB), `3427509` compressed data bytes (3.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00022.parquet`: `34965` rows, `3464760` file bytes (3.30 MiB), `22426064` physical bytes (21.39 MiB), `25741540` encoded bytes (24.55 MiB), `3434477` compressed data bytes (3.28 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00023.parquet`: `34535` rows, `3527753` file bytes (3.36 MiB), `22414335` physical bytes (21.38 MiB), `25687946` encoded bytes (24.50 MiB), `3497000` compressed data bytes (3.33 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00024.parquet`: `35253` rows, `3450744` file bytes (3.29 MiB), `22398299` physical bytes (21.36 MiB), `25744781` encoded bytes (24.55 MiB), `3420604` compressed data bytes (3.26 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00025.parquet`: `34864` rows, `3462298` file bytes (3.30 MiB), `22219270` physical bytes (21.19 MiB), `25529279` encoded bytes (24.35 MiB), `3431612` compressed data bytes (3.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00026.parquet`: `35281` rows, `3443830` file bytes (3.28 MiB), `22436449` physical bytes (21.40 MiB), `25784793` encoded bytes (24.59 MiB), `3413348` compressed data bytes (3.26 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00027.parquet`: `34991` rows, `3475495` file bytes (3.31 MiB), `22449639` physical bytes (21.41 MiB), `25770946` encoded bytes (24.58 MiB), `3445060` compressed data bytes (3.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed/part-00028.parquet`: `4524` rows, `454517` file bytes (443.86 KiB), `2959352` physical bytes (2.82 MiB), `3390903` encoded bytes (3.23 MiB), `438914` compressed data bytes (428.63 KiB)
