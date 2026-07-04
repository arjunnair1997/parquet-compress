# ClickBench Parquet Experiment

- Started: `2026-07-03T23:35:31-04:00`
- Write elapsed: `11.703s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `811670273` (774.07 MiB)
- Compressed column data bytes after codec compression: `87947012` (83.87 MiB)
- Parquet file bytes: `88846206` (84.73 MiB)
- Physical/encoded ratio: `0.878x`
- Encoded/compressed-data ratio: `9.229x`
- Physical/compressed-data ratio: `8.100x`
- Physical/parquet-file ratio: `8.018x`
- Files: `29`

## Settings

- Compression: `zstd-3`
- Int encoding: `plain`
- String encoding: `plain`
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
- Elapsed: `7.48s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004555` (7.63 MiB) | `8005360` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `63975` (62.48 KiB) | `0.999x` | `62.580x` | `62.524x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:526` | `1000000` | `138409995` (132.00 MiB) | `142872027` (136.25 MiB) | `13947498` (13.30 MiB) | `0.969x` | `10.244x` | `9.924x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003530` (3.82 MiB) | `4921` (4.81 KiB) | `0.999x` | `813.560x` | `812.843x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `3367691` (3.21 MiB) | `2886917` (2.75 MiB) | `2.376x` | `1.167x` | `2.771x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003527` (3.82 MiB) | `4918` (4.80 KiB) | `0.999x` | `814.056x` | `813.339x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003527` (3.82 MiB) | `4935` (4.82 KiB) | `0.999x` | `811.252x` | `810.537x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003596` (3.82 MiB) | `408157` (398.59 KiB) | `0.999x` | `9.809x` | `9.800x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `190938` (186.46 KiB) | `0.999x` | `20.968x` | `20.949x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004551` (7.63 MiB) | `618182` (603.69 KiB) | `0.999x` | `12.949x` | `12.941x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003526` (3.82 MiB) | `4233` (4.13 KiB) | `0.999x` | `945.789x` | `944.956x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `106378` (103.88 KiB) | `0.999x` | `37.635x` | `37.602x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `134024` (130.88 KiB) | `0.999x` | `29.872x` | `29.845x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:352` | `1000000` | `88562192` (84.46 MiB) | `92652025` (88.36 MiB) | `15302640` (14.59 MiB) | `0.956x` | `6.055x` | `5.787x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:326` | `1000000` | `79583339` (75.90 MiB) | `83646709` (79.77 MiB) | `14214564` (13.56 MiB) | `0.951x` | `5.885x` | `5.599x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `178704` (174.52 KiB) | `0.999x` | `22.403x` | `22.383x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003582` (3.82 MiB) | `275092` (268.64 KiB) | `0.999x` | `14.554x` | `14.541x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003582` (3.82 MiB) | `231281` (225.86 KiB) | `0.999x` | `17.310x` | `17.295x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `87576` (85.52 KiB) | `0.999x` | `45.716x` | `45.675x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `48118` (46.99 KiB) | `0.999x` | `83.203x` | `83.129x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003582` (3.82 MiB) | `187144` (182.76 KiB) | `0.999x` | `21.393x` | `21.374x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `186061` (181.70 KiB) | `0.999x` | `21.518x` | `21.498x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `82489` (80.56 KiB) | `0.999x` | `48.535x` | `48.491x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `53323` (52.07 KiB) | `0.999x` | `75.082x` | `75.015x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `128354` (125.35 KiB) | `0.999x` | `31.192x` | `31.164x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3354477` (3.20 MiB) | `7357825` (7.02 MiB) | `246450` (240.67 KiB) | `0.456x` | `29.855x` | `13.611x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `26179` (25.57 KiB) | `0.999x` | `152.931x` | `152.794x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `24691` (24.11 KiB) | `0.999x` | `162.148x` | `162.002x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `154775` (151.15 KiB) | `0.999x` | `25.867x` | `25.844x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3767530` (3.59 MiB) | `7770885` (7.41 MiB) | `137091` (133.88 KiB) | `0.485x` | `56.684x` | `27.482x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003534` (3.82 MiB) | `6098` (5.96 KiB) | `0.999x` | `656.532x` | `655.953x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003542` (3.82 MiB) | `6547` (6.39 KiB) | `0.999x` | `611.508x` | `610.967x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `28826` (28.15 KiB) | `0.999x` | `138.888x` | `138.764x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `22463` (21.94 KiB) | `0.999x` | `178.230x` | `178.071x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `81583` (79.67 KiB) | `4084876` (3.90 MiB) | `22555` (22.03 KiB) | `0.020x` | `181.107x` | `3.617x` | `81583` (79.67 KiB) |
| `Params` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002158` (3.82 MiB) | `2865` (2.80 KiB) | `0.000x` | `1396.914x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003592` (3.82 MiB) | `323846` (316.26 KiB) | `0.999x` | `12.363x` | `12.352x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `288219` (281.46 KiB) | `0.999x` | `13.891x` | `13.878x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `101620` (99.24 KiB) | `0.999x` | `39.398x` | `39.362x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3528017` (3.36 MiB) | `7534787` (7.19 MiB) | `720328` (703.45 KiB) | `0.468x` | `10.460x` | `4.898x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003578` (3.82 MiB) | `30693` (29.97 KiB) | `0.999x` | `130.439x` | `130.323x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `164881` (161.02 KiB) | `0.999x` | `24.282x` | `24.260x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `305793` (298.63 KiB) | `0.999x` | `13.092x` | `13.081x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `319860` (312.36 KiB) | `0.999x` | `12.517x` | `12.505x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `99875` (97.53 KiB) | `0.999x` | `40.086x` | `40.050x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `3489020` (3.33 MiB) | `2899138` (2.76 MiB) | `2.293x` | `1.203x` | `2.759x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `90764` (88.64 KiB) | `0.999x` | `44.110x` | `44.070x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `74105` (72.37 KiB) | `0.999x` | `54.026x` | `53.977x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003581` (3.82 MiB) | `123778` (120.88 KiB) | `0.999x` | `32.345x` | `32.316x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4365` (4.26 KiB) | `0.999x` | `917.189x` | `916.380x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:85` | `1000000` | `13587860` (12.96 MiB) | `17594801` (16.78 MiB) | `14581` (14.24 KiB) | `0.772x` | `1206.694x` | `931.888x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003546` (3.82 MiB) | `7092` (6.93 KiB) | `0.999x` | `564.516x` | `564.016x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003554` (3.82 MiB) | `56774` (55.44 KiB) | `0.999x` | `70.517x` | `70.455x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003551` (3.82 MiB) | `8030` (7.84 KiB) | `0.999x` | `498.574x` | `498.132x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003534` (3.82 MiB) | `25188` (24.60 KiB) | `0.999x` | `158.946x` | `158.806x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004549` (7.63 MiB) | `693987` (677.72 KiB) | `0.999x` | `11.534x` | `11.528x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:137` | `1000000` | `27797671` (26.51 MiB) | `31859426` (30.38 MiB) | `5324529` (5.08 MiB) | `0.873x` | `5.984x` | `5.221x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003642` (3.82 MiB) | `3815331` (3.64 MiB) | `0.999x` | `1.049x` | `1.048x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003526` (3.82 MiB) | `4233` (4.13 KiB) | `0.999x` | `945.789x` | `944.956x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003526` (3.82 MiB) | `4233` (4.13 KiB) | `0.999x` | `945.789x` | `944.956x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003526` (3.82 MiB) | `4233` (4.13 KiB) | `0.999x` | `945.789x` | `944.956x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `82633` (80.70 KiB) | `0.999x` | `48.450x` | `48.407x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003526` (3.82 MiB) | `4233` (4.13 KiB) | `0.999x` | `945.789x` | `944.956x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `1000000` (976.56 KiB) | `5002901` (4.77 MiB) | `29299` (28.61 KiB) | `0.200x` | `170.753x` | `34.131x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `3381185` (3.22 MiB) | `2893866` (2.76 MiB) | `2.366x` | `1.168x` | `2.764x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `142996` (139.64 KiB) | `0.999x` | `27.998x` | `27.973x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `107738` (105.21 KiB) | `0.999x` | `37.160x` | `37.127x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `122536` (119.66 KiB) | `0.999x` | `32.673x` | `32.643x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `193528` (188.99 KiB) | `0.999x` | `20.687x` | `20.669x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `173744` (169.67 KiB) | `0.999x` | `23.043x` | `23.022x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003605` (3.82 MiB) | `426670` (416.67 KiB) | `0.999x` | `9.383x` | `9.375x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003579` (3.82 MiB) | `70834` (69.17 KiB) | `0.999x` | `56.521x` | `56.470x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003526` (3.82 MiB) | `4233` (4.13 KiB) | `0.999x` | `945.789x` | `944.956x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003571` (3.82 MiB) | `55124` (53.83 KiB) | `0.999x` | `72.628x` | `72.564x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `2001192` (1.91 MiB) | `6004433` (5.73 MiB) | `32241` (31.49 KiB) | `0.333x` | `186.236x` | `62.070x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3325142` (3.17 MiB) | `7328499` (6.99 MiB) | `122258` (119.39 KiB) | `0.454x` | `59.943x` | `27.198x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002158` (3.82 MiB) | `2865` (2.80 KiB) | `0.000x` | `1396.914x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002158` (3.82 MiB) | `2865` (2.80 KiB) | `0.000x` | `1396.914x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003526` (3.82 MiB) | `4233` (4.13 KiB) | `0.999x` | `945.789x` | `944.956x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003542` (3.82 MiB) | `61475` (60.03 KiB) | `0.999x` | `65.125x` | `65.067x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `135500` (132.32 KiB) | `0.999x` | `29.547x` | `29.520x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003594` (3.82 MiB) | `333308` (325.50 KiB) | `0.999x` | `12.012x` | `12.001x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003641` (3.82 MiB) | `1245872` (1.19 MiB) | `0.999x` | `3.214x` | `3.211x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003640` (3.82 MiB) | `937902` (915.92 KiB) | `0.999x` | `4.269x` | `4.265x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003622` (3.82 MiB) | `549952` (537.06 KiB) | `0.999x` | `7.280x` | `7.273x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003526` (3.82 MiB) | `5325` (5.20 KiB) | `0.999x` | `751.836x` | `751.174x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `1024` (1.00 KiB) | `4004022` (3.82 MiB) | `4888` (4.77 KiB) | `0.000x` | `819.153x` | `0.209x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004437` (7.63 MiB) | `5709` (5.58 KiB) | `0.999x` | `1402.073x` | `1401.296x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002158` (3.82 MiB) | `2865` (2.80 KiB) | `0.000x` | `1396.914x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3000000` (2.86 MiB) | `7003304` (6.68 MiB) | `5262` (5.14 KiB) | `0.428x` | `1330.921x` | `570.125x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003526` (3.82 MiB) | `4233` (4.13 KiB) | `0.999x` | `945.789x` | `944.956x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `58030` (56.67 KiB) | `4062525` (3.87 MiB) | `19169` (18.72 KiB) | `0.014x` | `211.932x` | `3.027x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `22051` (21.53 KiB) | `4025607` (3.84 MiB) | `16530` (16.14 KiB) | `0.005x` | `243.533x` | `1.334x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `25445` (24.85 KiB) | `4029080` (3.84 MiB) | `18948` (18.50 KiB) | `0.006x` | `212.639x` | `1.343x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `48191` (47.06 KiB) | `4051763` (3.86 MiB) | `13311` (13.00 KiB) | `0.012x` | `304.392x` | `3.620x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `49433` (48.27 KiB) | `4053231` (3.87 MiB) | `21089` (20.59 KiB) | `0.012x` | `192.196x` | `2.344x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `16873` (16.48 KiB) | `4019935` (3.83 MiB) | `17081` (16.68 KiB) | `0.004x` | `235.345x` | `0.988x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `91870` (89.72 KiB) | `4097410` (3.91 MiB) | `29800` (29.10 KiB) | `0.022x` | `137.497x` | `3.083x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `13001` (12.70 KiB) | `4016537` (3.83 MiB) | `13959` (13.63 KiB) | `0.003x` | `287.738x` | `0.931x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `28101` (27.44 KiB) | `4032755` (3.85 MiB) | `15692` (15.32 KiB) | `0.007x` | `256.994x` | `1.791x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `45607` (44.54 KiB) | `4048288` (3.86 MiB) | `28366` (27.70 KiB) | `0.011x` | `142.716x` | `1.608x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003557` (3.82 MiB) | `21190` (20.69 KiB) | `0.999x` | `188.936x` | `188.768x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004554` (7.63 MiB) | `2841910` (2.71 MiB) | `0.999x` | `2.817x` | `2.815x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004556` (7.63 MiB) | `3580352` (3.41 MiB) | `0.999x` | `2.236x` | `2.234x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003544` (3.82 MiB) | `5630` (5.50 KiB) | `0.999x` | `711.109x` | `710.480x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00000.parquet`: `35854` rows, `2901156` file bytes (2.77 MiB), `27844924` physical bytes (26.55 MiB), `31421923` encoded bytes (29.97 MiB), `2869108` compressed data bytes (2.74 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00001.parquet`: `35947` rows, `2855833` file bytes (2.72 MiB), `27614679` physical bytes (26.34 MiB), `31196994` encoded bytes (29.75 MiB), `2823243` compressed data bytes (2.69 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00002.parquet`: `35804` rows, `2914388` file bytes (2.78 MiB), `28045024` physical bytes (26.75 MiB), `31617960` encoded bytes (30.15 MiB), `2881109` compressed data bytes (2.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00003.parquet`: `36223` rows, `2871756` file bytes (2.74 MiB), `28041589` physical bytes (26.74 MiB), `31648140` encoded bytes (30.18 MiB), `2839321` compressed data bytes (2.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00004.parquet`: `35583` rows, `2825996` file bytes (2.70 MiB), `27676556` physical bytes (26.39 MiB), `31223662` encoded bytes (29.78 MiB), `2792824` compressed data bytes (2.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00005.parquet`: `35769` rows, `2818984` file bytes (2.69 MiB), `27421909` physical bytes (26.15 MiB), `30979751` encoded bytes (29.54 MiB), `2786570` compressed data bytes (2.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00006.parquet`: `36202` rows, `2876519` file bytes (2.74 MiB), `28032888` physical bytes (26.73 MiB), `31643203` encoded bytes (30.18 MiB), `2844060` compressed data bytes (2.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00007.parquet`: `36255` rows, `2862697` file bytes (2.73 MiB), `28048091` physical bytes (26.75 MiB), `31660225` encoded bytes (30.19 MiB), `2829839` compressed data bytes (2.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00008.parquet`: `36216` rows, `2841142` file bytes (2.71 MiB), `28069629` physical bytes (26.77 MiB), `31675459` encoded bytes (30.21 MiB), `2808443` compressed data bytes (2.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00009.parquet`: `35848` rows, `2846661` file bytes (2.71 MiB), `27649376` physical bytes (26.37 MiB), `31218576` encoded bytes (29.77 MiB), `2814519` compressed data bytes (2.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00010.parquet`: `36464` rows, `2842418` file bytes (2.71 MiB), `28033620` physical bytes (26.73 MiB), `31660990` encoded bytes (30.19 MiB), `2809830` compressed data bytes (2.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00011.parquet`: `36387` rows, `2828217` file bytes (2.70 MiB), `28070531` physical bytes (26.77 MiB), `31695344` encoded bytes (30.23 MiB), `2795731` compressed data bytes (2.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00012.parquet`: `36319` rows, `2824593` file bytes (2.69 MiB), `27840758` physical bytes (26.55 MiB), `31455601` encoded bytes (30.00 MiB), `2792472` compressed data bytes (2.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00013.parquet`: `36203` rows, `2868587` file bytes (2.74 MiB), `28028370` physical bytes (26.73 MiB), `31633886` encoded bytes (30.17 MiB), `2835897` compressed data bytes (2.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00014.parquet`: `35357` rows, `3280150` file bytes (3.13 MiB), `24474789` physical bytes (23.34 MiB), `27988536` encoded bytes (26.69 MiB), `3249325` compressed data bytes (3.10 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00015.parquet`: `36052` rows, `3438356` file bytes (3.28 MiB), `25185428` physical bytes (24.02 MiB), `28764377` encoded bytes (27.43 MiB), `3407947` compressed data bytes (3.25 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00016.parquet`: `35301` rows, `3446698` file bytes (3.29 MiB), `23226963` physical bytes (22.15 MiB), `26724072` encoded bytes (25.49 MiB), `3416438` compressed data bytes (3.26 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00017.parquet`: `34854` rows, `3464849` file bytes (3.30 MiB), `22487729` physical bytes (21.45 MiB), `25933457` encoded bytes (24.73 MiB), `3434120` compressed data bytes (3.28 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00018.parquet`: `34367` rows, `3552107` file bytes (3.39 MiB), `22419389` physical bytes (21.38 MiB), `25812050` encoded bytes (24.62 MiB), `3521776` compressed data bytes (3.36 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00019.parquet`: `34957` rows, `3443350` file bytes (3.28 MiB), `22245340` physical bytes (21.21 MiB), `25700149` encoded bytes (24.51 MiB), `3412612` compressed data bytes (3.25 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00020.parquet`: `34543` rows, `3507312` file bytes (3.34 MiB), `22211997` physical bytes (21.18 MiB), `25627451` encoded bytes (24.44 MiB), `3476413` compressed data bytes (3.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00021.parquet`: `35095` rows, `3458535` file bytes (3.30 MiB), `22433880` physical bytes (21.39 MiB), `25901848` encoded bytes (24.70 MiB), `3427907` compressed data bytes (3.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00022.parquet`: `34960` rows, `3464153` file bytes (3.30 MiB), `22428319` physical bytes (21.39 MiB), `25880834` encoded bytes (24.68 MiB), `3433868` compressed data bytes (3.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00023.parquet`: `34544` rows, `3527809` file bytes (3.36 MiB), `22413373` physical bytes (21.38 MiB), `25825041` encoded bytes (24.63 MiB), `3497054` compressed data bytes (3.34 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00024.parquet`: `35252` rows, `3450514` file bytes (3.29 MiB), `22398670` physical bytes (21.36 MiB), `25884421` encoded bytes (24.69 MiB), `3420372` compressed data bytes (3.26 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00025.parquet`: `34865` rows, `3463117` file bytes (3.30 MiB), `22218879` physical bytes (21.19 MiB), `25667032` encoded bytes (24.48 MiB), `3432429` compressed data bytes (3.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00026.parquet`: `35281` rows, `3442722` file bytes (3.28 MiB), `22437208` physical bytes (21.40 MiB), `25924566` encoded bytes (24.72 MiB), `3412239` compressed data bytes (3.25 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00027.parquet`: `34993` rows, `3475319` file bytes (3.31 MiB), `22450104` physical bytes (21.41 MiB), `25909381` encoded bytes (24.71 MiB), `3444882` compressed data bytes (3.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed/part-00028.parquet`: `4505` rows, `452268` file bytes (441.67 KiB), `2948612` physical bytes (2.81 MiB), `3395344` encoded bytes (3.24 MiB), `436664` compressed data bytes (426.43 KiB)
