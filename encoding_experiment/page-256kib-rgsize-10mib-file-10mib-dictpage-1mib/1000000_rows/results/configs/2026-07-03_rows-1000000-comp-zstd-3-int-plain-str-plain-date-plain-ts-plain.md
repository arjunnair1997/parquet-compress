# ClickBench Parquet Experiment

- Started: `2026-07-03T15:31:23-04:00`
- Write elapsed: `11.697s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `825442338` (787.20 MiB)
- Compressed column data bytes after codec compression: `86812277` (82.79 MiB)
- Parquet file bytes: `87711872` (83.65 MiB)
- Physical/encoded ratio: `0.863x`
- Encoded/compressed-data ratio: `9.508x`
- Physical/compressed-data ratio: `8.206x`
- Physical/parquet-file ratio: `8.122x`
- Files: `29`

## Settings

- Compression: `zstd-3`
- Int encoding: `plain`
- String encoding: `plain`
- Date encoding: `plain`
- Timestamp encoding: `plain`
- Max page size: `256.00 KiB`
- Max dictionary page size: `1.00 MiB`
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
- Elapsed: `7.412s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004557` (7.63 MiB) | `8005365` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `63644` (62.15 KiB) | `0.999x` | `62.906x` | `62.850x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:528` | `1000000` | `138409995` (132.00 MiB) | `142869584` (136.25 MiB) | `13939135` (13.29 MiB) | `0.969x` | `10.250x` | `9.930x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `4925` (4.81 KiB) | `0.999x` | `812.900x` | `812.183x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004556` (7.63 MiB) | `2517896` (2.40 MiB) | `0.999x` | `3.179x` | `3.177x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4923` (4.81 KiB) | `0.999x` | `813.230x` | `812.513x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003530` (3.82 MiB) | `4941` (4.83 KiB) | `0.999x` | `810.267x` | `809.553x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003599` (3.82 MiB) | `408069` (398.50 KiB) | `0.999x` | `9.811x` | `9.802x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `191174` (186.69 KiB) | `0.999x` | `20.942x` | `20.923x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004550` (7.63 MiB) | `618298` (603.81 KiB) | `0.999x` | `12.946x` | `12.939x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4240` (4.14 KiB) | `0.999x` | `944.229x` | `943.396x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `105905` (103.42 KiB) | `0.999x` | `37.804x` | `37.770x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `135169` (132.00 KiB) | `0.999x` | `29.619x` | `29.593x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:356` | `1000000` | `88562192` (84.46 MiB) | `92652677` (88.36 MiB) | `15306955` (14.60 MiB) | `0.956x` | `6.053x` | `5.786x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:326` | `1000000` | `79583339` (75.90 MiB) | `83647728` (79.77 MiB) | `14218039` (13.56 MiB) | `0.951x` | `5.883x` | `5.597x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `178941` (174.75 KiB) | `0.999x` | `22.374x` | `22.354x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `275135` (268.69 KiB) | `0.999x` | `14.551x` | `14.538x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `231356` (225.93 KiB) | `0.999x` | `17.305x` | `17.289x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `87570` (85.52 KiB) | `0.999x` | `45.719x` | `45.678x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `48124` (47.00 KiB) | `0.999x` | `83.193x` | `83.119x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `187141` (182.75 KiB) | `0.999x` | `21.393x` | `21.374x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `186024` (181.66 KiB) | `0.999x` | `21.522x` | `21.503x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `82090` (80.17 KiB) | `0.999x` | `48.771x` | `48.727x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `53604` (52.35 KiB) | `0.999x` | `74.688x` | `74.621x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `126766` (123.79 KiB) | `0.999x` | `31.582x` | `31.554x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3354477` (3.20 MiB) | `7357828` (7.02 MiB) | `246539` (240.76 KiB) | `0.456x` | `29.844x` | `13.606x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `26145` (25.53 KiB) | `0.999x` | `153.130x` | `152.993x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `24917` (24.33 KiB) | `0.999x` | `160.677x` | `160.533x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `154189` (150.58 KiB) | `0.999x` | `25.965x` | `25.942x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3767530` (3.59 MiB) | `7770883` (7.41 MiB) | `137050` (133.84 KiB) | `0.485x` | `56.701x` | `27.490x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003537` (3.82 MiB) | `6094` (5.95 KiB) | `0.999x` | `656.964x` | `656.383x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003542` (3.82 MiB) | `6510` (6.36 KiB) | `0.999x` | `614.983x` | `614.439x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `28686` (28.01 KiB) | `0.999x` | `139.566x` | `139.441x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `22518` (21.99 KiB) | `0.999x` | `177.795x` | `177.636x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `81583` (79.67 KiB) | `4084882` (3.90 MiB) | `22607` (22.08 KiB) | `0.020x` | `180.691x` | `3.609x` | `81583` (79.67 KiB) |
| `Params` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002161` (3.82 MiB) | `2872` (2.80 KiB) | `0.000x` | `1393.510x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003592` (3.82 MiB) | `323783` (316.19 KiB) | `0.999x` | `12.365x` | `12.354x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `288326` (281.57 KiB) | `0.999x` | `13.886x` | `13.873x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `101613` (99.23 KiB) | `0.999x` | `39.400x` | `39.365x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3528017` (3.36 MiB) | `7534820` (7.19 MiB) | `719930` (703.06 KiB) | `0.468x` | `10.466x` | `4.901x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003573` (3.82 MiB) | `30696` (29.98 KiB) | `0.999x` | `130.427x` | `130.310x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `164989` (161.12 KiB) | `0.999x` | `24.266x` | `24.244x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `305751` (298.58 KiB) | `0.999x` | `13.094x` | `13.083x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `319807` (312.31 KiB) | `0.999x` | `12.519x` | `12.508x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003582` (3.82 MiB) | `99924` (97.58 KiB) | `0.999x` | `40.066x` | `40.030x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004556` (7.63 MiB) | `2475752` (2.36 MiB) | `0.999x` | `3.233x` | `3.231x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `90600` (88.48 KiB) | `0.999x` | `44.190x` | `44.150x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `74028` (72.29 KiB) | `0.999x` | `54.082x` | `54.034x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `123785` (120.88 KiB) | `0.999x` | `32.343x` | `32.314x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4368` (4.27 KiB) | `0.999x` | `916.559x` | `915.751x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:85` | `1000000` | `13587860` (12.96 MiB) | `17594801` (16.78 MiB) | `14612` (14.27 KiB) | `0.772x` | `1204.134x` | `929.911x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003550` (3.82 MiB) | `7064` (6.90 KiB) | `0.999x` | `566.754x` | `566.251x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003557` (3.82 MiB) | `56781` (55.45 KiB) | `0.999x` | `70.509x` | `70.446x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003551` (3.82 MiB) | `8096` (7.91 KiB) | `0.999x` | `494.510x` | `494.071x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003535` (3.82 MiB) | `25193` (24.60 KiB) | `0.999x` | `158.915x` | `158.774x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004554` (7.63 MiB) | `694033` (677.77 KiB) | `0.999x` | `11.533x` | `11.527x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:137` | `1000000` | `27797671` (26.51 MiB) | `31856344` (30.38 MiB) | `5323588` (5.08 MiB) | `0.873x` | `5.984x` | `5.222x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003642` (3.82 MiB) | `3849795` (3.67 MiB) | `0.999x` | `1.040x` | `1.039x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4240` (4.14 KiB) | `0.999x` | `944.229x` | `943.396x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4240` (4.14 KiB) | `0.999x` | `944.229x` | `943.396x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4240` (4.14 KiB) | `0.999x` | `944.229x` | `943.396x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `82558` (80.62 KiB) | `0.999x` | `48.494x` | `48.451x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4240` (4.14 KiB) | `0.999x` | `944.229x` | `943.396x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `1000000` (976.56 KiB) | `5002901` (4.77 MiB) | `29173` (28.49 KiB) | `0.200x` | `171.491x` | `34.278x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004555` (7.63 MiB) | `2520192` (2.40 MiB) | `0.999x` | `3.176x` | `3.174x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `142624` (139.28 KiB) | `0.999x` | `28.071x` | `28.046x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `107976` (105.45 KiB) | `0.999x` | `37.078x` | `37.045x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `122679` (119.80 KiB) | `0.999x` | `32.635x` | `32.605x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `193520` (188.98 KiB) | `0.999x` | `20.688x` | `20.670x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `173697` (169.63 KiB) | `0.999x` | `23.049x` | `23.029x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003603` (3.82 MiB) | `426607` (416.61 KiB) | `0.999x` | `9.385x` | `9.376x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003578` (3.82 MiB) | `70817` (69.16 KiB) | `0.999x` | `56.534x` | `56.484x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4239` (4.14 KiB) | `0.999x` | `944.451x` | `943.619x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003572` (3.82 MiB) | `55137` (53.84 KiB) | `0.999x` | `72.611x` | `72.547x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `2001192` (1.91 MiB) | `6004438` (5.73 MiB) | `32361` (31.60 KiB) | `0.333x` | `185.546x` | `61.840x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3325142` (3.17 MiB) | `7328499` (6.99 MiB) | `122222` (119.36 KiB) | `0.454x` | `59.961x` | `27.206x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002161` (3.82 MiB) | `2872` (2.80 KiB) | `0.000x` | `1393.510x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002161` (3.82 MiB) | `2872` (2.80 KiB) | `0.000x` | `1393.510x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4240` (4.14 KiB) | `0.999x` | `944.229x` | `943.396x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003537` (3.82 MiB) | `61444` (60.00 KiB) | `0.999x` | `65.157x` | `65.100x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `135564` (132.39 KiB) | `0.999x` | `29.533x` | `29.506x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003600` (3.82 MiB) | `333314` (325.50 KiB) | `0.999x` | `12.011x` | `12.001x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003643` (3.82 MiB) | `1245891` (1.19 MiB) | `0.999x` | `3.213x` | `3.211x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003641` (3.82 MiB) | `937782` (915.80 KiB) | `0.999x` | `4.269x` | `4.265x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003625` (3.82 MiB) | `549819` (536.93 KiB) | `0.999x` | `7.282x` | `7.275x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003530` (3.82 MiB) | `5318` (5.19 KiB) | `0.999x` | `752.826x` | `752.162x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `1024` (1.00 KiB) | `4003989` (3.82 MiB) | `4856` (4.74 KiB) | `0.000x` | `824.545x` | `0.211x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004440` (7.63 MiB) | `5725` (5.59 KiB) | `0.999x` | `1398.155x` | `1397.380x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002161` (3.82 MiB) | `2872` (2.80 KiB) | `0.000x` | `1393.510x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3000000` (2.86 MiB) | `7003303` (6.68 MiB) | `5261` (5.14 KiB) | `0.428x` | `1331.173x` | `570.234x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4240` (4.14 KiB) | `0.999x` | `944.229x` | `943.396x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `58030` (56.67 KiB) | `4062525` (3.87 MiB) | `19232` (18.78 KiB) | `0.014x` | `211.238x` | `3.017x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `22051` (21.53 KiB) | `4025605` (3.84 MiB) | `16542` (16.15 KiB) | `0.005x` | `243.357x` | `1.333x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `25445` (24.85 KiB) | `4029111` (3.84 MiB) | `19026` (18.58 KiB) | `0.006x` | `211.769x` | `1.337x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `48191` (47.06 KiB) | `4051759` (3.86 MiB) | `13322` (13.01 KiB) | `0.012x` | `304.140x` | `3.617x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `49433` (48.27 KiB) | `4053208` (3.87 MiB) | `21033` (20.54 KiB) | `0.012x` | `192.707x` | `2.350x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `16873` (16.48 KiB) | `4019934` (3.83 MiB) | `16986` (16.59 KiB) | `0.004x` | `236.662x` | `0.993x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `91870` (89.72 KiB) | `4097442` (3.91 MiB) | `29858` (29.16 KiB) | `0.022x` | `137.231x` | `3.077x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `13001` (12.70 KiB) | `4016539` (3.83 MiB) | `13966` (13.64 KiB) | `0.003x` | `287.594x` | `0.931x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `28101` (27.44 KiB) | `4032776` (3.85 MiB) | `15728` (15.36 KiB) | `0.007x` | `256.407x` | `1.787x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `45607` (44.54 KiB) | `4048288` (3.86 MiB) | `28372` (27.71 KiB) | `0.011x` | `142.686x` | `1.607x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003558` (3.82 MiB) | `21186` (20.69 KiB) | `0.999x` | `188.972x` | `188.804x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004554` (7.63 MiB) | `2842003` (2.71 MiB) | `0.999x` | `2.817x` | `2.815x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004551` (7.63 MiB) | `3580480` (3.41 MiB) | `0.999x` | `2.236x` | `2.234x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003542` (3.82 MiB) | `5641` (5.51 KiB) | `0.999x` | `709.722x` | `709.094x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00000.parquet`: `35889` rows, `2840342` file bytes (2.71 MiB), `27875599` physical bytes (26.58 MiB), `31941149` encoded bytes (30.46 MiB), `2808290` compressed data bytes (2.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00001.parquet`: `35984` rows, `2790939` file bytes (2.66 MiB), `27645074` physical bytes (26.36 MiB), `31729973` encoded bytes (30.26 MiB), `2758345` compressed data bytes (2.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00002.parquet`: `35875` rows, `2831903` file bytes (2.70 MiB), `28091748` physical bytes (26.79 MiB), `32165856` encoded bytes (30.68 MiB), `2798621` compressed data bytes (2.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00003.parquet`: `36277` rows, `2803634` file bytes (2.67 MiB), `28077471` physical bytes (26.78 MiB), `32193828` encoded bytes (30.70 MiB), `2771219` compressed data bytes (2.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00004.parquet`: `35612` rows, `2768246` file bytes (2.64 MiB), `27705258` physical bytes (26.42 MiB), `31748552` encoded bytes (30.28 MiB), `2735053` compressed data bytes (2.61 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00005.parquet`: `36089` rows, `2778100` file bytes (2.65 MiB), `27666502` physical bytes (26.38 MiB), `31757787` encoded bytes (30.29 MiB), `2745631` compressed data bytes (2.62 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00006.parquet`: `36248` rows, `2820105` file bytes (2.69 MiB), `28058832` physical bytes (26.76 MiB), `32175263` encoded bytes (30.68 MiB), `2787658` compressed data bytes (2.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00007.parquet`: `36246` rows, `2828681` file bytes (2.70 MiB), `28062130` physical bytes (26.76 MiB), `32174229` encoded bytes (30.68 MiB), `2795775` compressed data bytes (2.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00008.parquet`: `36295` rows, `2773385` file bytes (2.64 MiB), `28105131` physical bytes (26.80 MiB), `32218564` encoded bytes (30.73 MiB), `2740644` compressed data bytes (2.61 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00009.parquet`: `35881` rows, `2788876` file bytes (2.66 MiB), `27676562` physical bytes (26.39 MiB), `31744081` encoded bytes (30.27 MiB), `2756708` compressed data bytes (2.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00010.parquet`: `36424` rows, `2783348` file bytes (2.65 MiB), `28067097` physical bytes (26.77 MiB), `32201379` encoded bytes (30.71 MiB), `2750742` compressed data bytes (2.62 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00011.parquet`: `36019` rows, `2728605` file bytes (2.60 MiB), `27731317` physical bytes (26.45 MiB), `31819087` encoded bytes (30.35 MiB), `2696233` compressed data bytes (2.57 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00012.parquet`: `36350` rows, `2791522` file bytes (2.66 MiB), `27863603` physical bytes (26.57 MiB), `31981942` encoded bytes (30.50 MiB), `2759400` compressed data bytes (2.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00013.parquet`: `36240` rows, `2813976` file bytes (2.68 MiB), `28059163` physical bytes (26.76 MiB), `32169728` encoded bytes (30.68 MiB), `2781309` compressed data bytes (2.65 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00014.parquet`: `35149` rows, `3218893` file bytes (3.07 MiB), `24299472` physical bytes (23.17 MiB), `28266012` encoded bytes (26.96 MiB), `3188088` compressed data bytes (3.04 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00015.parquet`: `36069` rows, `3405874` file bytes (3.25 MiB), `25205299` physical bytes (24.04 MiB), `29271924` encoded bytes (27.92 MiB), `3375459` compressed data bytes (3.22 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00016.parquet`: `35325` rows, `3428728` file bytes (3.27 MiB), `23230806` physical bytes (22.15 MiB), `27211574` encoded bytes (25.95 MiB), `3398477` compressed data bytes (3.24 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00017.parquet`: `34862` rows, `3455043` file bytes (3.29 MiB), `22496277` physical bytes (21.45 MiB), `26422031` encoded bytes (25.20 MiB), `3424268` compressed data bytes (3.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00018.parquet`: `34371` rows, `3542179` file bytes (3.38 MiB), `22421574` physical bytes (21.38 MiB), `26292319` encoded bytes (25.07 MiB), `3511845` compressed data bytes (3.35 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00019.parquet`: `34982` rows, `3423648` file bytes (3.27 MiB), `22257208` physical bytes (21.23 MiB), `26196636` encoded bytes (24.98 MiB), `3392907` compressed data bytes (3.24 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00020.parquet`: `34541` rows, `3494938` file bytes (3.33 MiB), `22210127` physical bytes (21.18 MiB), `26099689` encoded bytes (24.89 MiB), `3464023` compressed data bytes (3.30 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00021.parquet`: `35110` rows, `3441728` file bytes (3.28 MiB), `22437782` physical bytes (21.40 MiB), `26390842` encoded bytes (25.17 MiB), `3411096` compressed data bytes (3.25 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00022.parquet`: `34847` rows, `3455835` file bytes (3.30 MiB), `22447978` physical bytes (21.41 MiB), `26370673` encoded bytes (25.15 MiB), `3425555` compressed data bytes (3.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00023.parquet`: `34665` rows, `3507488` file bytes (3.35 MiB), `22408521` physical bytes (21.37 MiB), `26311396` encoded bytes (25.09 MiB), `3476729` compressed data bytes (3.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00024.parquet`: `35257` rows, `3438208` file bytes (3.28 MiB), `22397011` physical bytes (21.36 MiB), `26365160` encoded bytes (25.14 MiB), `3408067` compressed data bytes (3.25 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00025.parquet`: `34867` rows, `3446844` file bytes (3.29 MiB), `22225491` physical bytes (21.20 MiB), `26151283` encoded bytes (24.94 MiB), `3416161` compressed data bytes (3.26 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00026.parquet`: `35306` rows, `3426468` file bytes (3.27 MiB), `22447576` physical bytes (21.41 MiB), `26422434` encoded bytes (25.20 MiB), `3395953` compressed data bytes (3.24 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00027.parquet`: `35237` rows, `3482290` file bytes (3.32 MiB), `22630942` physical bytes (21.58 MiB), `26598234` encoded bytes (25.37 MiB), `3451731` compressed data bytes (3.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain/part-00028.parquet`: `3983` rows, `402046` file bytes (392.62 KiB), `2597073` physical bytes (2.48 MiB), `3050713` encoded bytes (2.91 MiB), `386290` compressed data bytes (377.24 KiB)
