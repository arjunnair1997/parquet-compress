# ClickBench Parquet Experiment

- Started: `2026-07-03T23:35:50-04:00`
- Write elapsed: `11.633s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `807674172` (770.26 MiB)
- Compressed column data bytes after codec compression: `87959994` (83.89 MiB)
- Parquet file bytes: `88859844` (84.74 MiB)
- Physical/encoded ratio: `0.882x`
- Encoded/compressed-data ratio: `9.182x`
- Physical/compressed-data ratio: `8.099x`
- Physical/parquet-file ratio: `8.017x`
- Files: `29`

## Settings

- Compression: `zstd-3`
- Int encoding: `plain`
- String encoding: `plain`
- Date encoding: `rle-dict`
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
- Elapsed: `7.46s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004555` (7.63 MiB) | `8005360` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `64068` (62.57 KiB) | `0.999x` | `62.490x` | `62.434x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:526` | `1000000` | `138409995` (132.00 MiB) | `142871743` (136.25 MiB) | `13947575` (13.30 MiB) | `0.969x` | `10.243x` | `9.924x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003530` (3.82 MiB) | `4921` (4.81 KiB) | `0.999x` | `813.560x` | `812.843x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `3368666` (3.21 MiB) | `2886539` (2.75 MiB) | `2.375x` | `1.167x` | `2.771x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4953` (4.84 KiB) | `5979` (5.84 KiB) | `807.591x` | `0.828x` | `669.008x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003525` (3.82 MiB) | `4933` (4.82 KiB) | `0.999x` | `811.580x` | `810.866x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003597` (3.82 MiB) | `408157` (398.59 KiB) | `0.999x` | `9.809x` | `9.800x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `190977` (186.50 KiB) | `0.999x` | `20.964x` | `20.945x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004552` (7.63 MiB) | `618183` (603.69 KiB) | `0.999x` | `12.949x` | `12.941x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003527` (3.82 MiB) | `4234` (4.13 KiB) | `0.999x` | `945.566x` | `944.733x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `106214` (103.72 KiB) | `0.999x` | `37.694x` | `37.660x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `134027` (130.89 KiB) | `0.999x` | `29.871x` | `29.845x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:352` | `1000000` | `88562192` (84.46 MiB) | `92652142` (88.36 MiB) | `15302624` (14.59 MiB) | `0.956x` | `6.055x` | `5.787x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:326` | `1000000` | `79583339` (75.90 MiB) | `83646861` (79.77 MiB) | `14215728` (13.56 MiB) | `0.951x` | `5.884x` | `5.598x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `178727` (174.54 KiB) | `0.999x` | `22.401x` | `22.381x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003581` (3.82 MiB) | `275159` (268.71 KiB) | `0.999x` | `14.550x` | `14.537x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `231286` (225.87 KiB) | `0.999x` | `17.310x` | `17.295x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `87566` (85.51 KiB) | `0.999x` | `45.721x` | `45.680x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `48097` (46.97 KiB) | `0.999x` | `83.240x` | `83.165x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `187140` (182.75 KiB) | `0.999x` | `21.394x` | `21.374x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003581` (3.82 MiB) | `186031` (181.67 KiB) | `0.999x` | `21.521x` | `21.502x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `82416` (80.48 KiB) | `0.999x` | `48.578x` | `48.534x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `53319` (52.07 KiB) | `0.999x` | `75.087x` | `75.020x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `128036` (125.04 KiB) | `0.999x` | `31.269x` | `31.241x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3354477` (3.20 MiB) | `7357825` (7.02 MiB) | `246338` (240.56 KiB) | `0.456x` | `29.869x` | `13.617x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `26174` (25.56 KiB) | `0.999x` | `152.960x` | `152.823x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `24688` (24.11 KiB) | `0.999x` | `162.167x` | `162.022x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `154759` (151.13 KiB) | `0.999x` | `25.870x` | `25.847x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3767530` (3.59 MiB) | `7770886` (7.41 MiB) | `137047` (133.83 KiB) | `0.485x` | `56.702x` | `27.491x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003535` (3.82 MiB) | `6100` (5.96 KiB) | `0.999x` | `656.317x` | `655.738x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003541` (3.82 MiB) | `6546` (6.39 KiB) | `0.999x` | `611.601x` | `611.060x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `28834` (28.16 KiB) | `0.999x` | `138.850x` | `138.725x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `22468` (21.94 KiB) | `0.999x` | `178.191x` | `178.031x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `81583` (79.67 KiB) | `4084877` (3.90 MiB) | `22583` (22.05 KiB) | `0.020x` | `180.883x` | `3.613x` | `81583` (79.67 KiB) |
| `Params` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002159` (3.82 MiB) | `2866` (2.80 KiB) | `0.000x` | `1396.427x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003592` (3.82 MiB) | `323851` (316.26 KiB) | `0.999x` | `12.362x` | `12.351x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `288261` (281.50 KiB) | `0.999x` | `13.889x` | `13.876x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `101565` (99.18 KiB) | `0.999x` | `39.419x` | `39.384x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3528017` (3.36 MiB) | `7534788` (7.19 MiB) | `720379` (703.50 KiB) | `0.468x` | `10.459x` | `4.897x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003578` (3.82 MiB) | `30684` (29.96 KiB) | `0.999x` | `130.478x` | `130.361x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `164871` (161.01 KiB) | `0.999x` | `24.283x` | `24.261x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `305824` (298.66 KiB) | `0.999x` | `13.091x` | `13.079x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003591` (3.82 MiB) | `319926` (312.43 KiB) | `0.999x` | `12.514x` | `12.503x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `99888` (97.55 KiB) | `0.999x` | `40.081x` | `40.045x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `3489240` (3.33 MiB) | `2896992` (2.76 MiB) | `2.293x` | `1.204x` | `2.761x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `90799` (88.67 KiB) | `0.999x` | `44.093x` | `44.053x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `74098` (72.36 KiB) | `0.999x` | `54.031x` | `53.983x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `123783` (120.88 KiB) | `0.999x` | `32.344x` | `32.315x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4365` (4.26 KiB) | `0.999x` | `917.189x` | `916.380x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:85` | `1000000` | `13587860` (12.96 MiB) | `17594795` (16.78 MiB) | `14575` (14.23 KiB) | `0.772x` | `1207.190x` | `932.272x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003548` (3.82 MiB) | `7094` (6.93 KiB) | `0.999x` | `564.357x` | `563.857x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003556` (3.82 MiB) | `56826` (55.49 KiB) | `0.999x` | `70.453x` | `70.390x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003551` (3.82 MiB) | `8064` (7.88 KiB) | `0.999x` | `496.472x` | `496.032x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003535` (3.82 MiB) | `25207` (24.62 KiB) | `0.999x` | `158.826x` | `158.686x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004552` (7.63 MiB) | `693958` (677.69 KiB) | `0.999x` | `11.535x` | `11.528x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:137` | `1000000` | `27797671` (26.51 MiB) | `31859424` (30.38 MiB) | `5325095` (5.08 MiB) | `0.873x` | `5.983x` | `5.220x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003641` (3.82 MiB) | `3827433` (3.65 MiB) | `0.999x` | `1.046x` | `1.045x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003527` (3.82 MiB) | `4234` (4.13 KiB) | `0.999x` | `945.566x` | `944.733x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003527` (3.82 MiB) | `4234` (4.13 KiB) | `0.999x` | `945.566x` | `944.733x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003527` (3.82 MiB) | `4234` (4.13 KiB) | `0.999x` | `945.566x` | `944.733x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `82697` (80.76 KiB) | `0.999x` | `48.413x` | `48.369x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003527` (3.82 MiB) | `4234` (4.13 KiB) | `0.999x` | `945.566x` | `944.733x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `1000000` (976.56 KiB) | `5002903` (4.77 MiB) | `29070` (28.39 KiB) | `0.200x` | `172.098x` | `34.400x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:57` | `1000000` | `8000000` (7.63 MiB) | `3382465` (3.23 MiB) | `2894643` (2.76 MiB) | `2.365x` | `1.169x` | `2.764x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `142995` (139.64 KiB) | `0.999x` | `27.998x` | `27.973x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `107830` (105.30 KiB) | `0.999x` | `37.129x` | `37.095x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `122651` (119.78 KiB) | `0.999x` | `32.642x` | `32.613x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `193538` (189.00 KiB) | `0.999x` | `20.686x` | `20.668x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `173740` (169.67 KiB) | `0.999x` | `23.044x` | `23.023x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003603` (3.82 MiB) | `426634` (416.63 KiB) | `0.999x` | `9.384x` | `9.376x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003578` (3.82 MiB) | `70828` (69.17 KiB) | `0.999x` | `56.525x` | `56.475x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4236` (4.14 KiB) | `0.999x` | `945.120x` | `944.287x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003570` (3.82 MiB) | `55128` (53.84 KiB) | `0.999x` | `72.623x` | `72.558x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `2001192` (1.91 MiB) | `6004430` (5.73 MiB) | `32240` (31.48 KiB) | `0.333x` | `186.242x` | `62.072x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3325142` (3.17 MiB) | `7328500` (6.99 MiB) | `122230` (119.37 KiB) | `0.454x` | `59.957x` | `27.204x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002159` (3.82 MiB) | `2866` (2.80 KiB) | `0.000x` | `1396.427x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002159` (3.82 MiB) | `2866` (2.80 KiB) | `0.000x` | `1396.427x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003527` (3.82 MiB) | `4234` (4.13 KiB) | `0.999x` | `945.566x` | `944.733x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003541` (3.82 MiB) | `61437` (60.00 KiB) | `0.999x` | `65.165x` | `65.107x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `135476` (132.30 KiB) | `0.999x` | `29.552x` | `29.526x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003594` (3.82 MiB) | `333302` (325.49 KiB) | `0.999x` | `12.012x` | `12.001x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003643` (3.82 MiB) | `1245942` (1.19 MiB) | `0.999x` | `3.213x` | `3.210x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003641` (3.82 MiB) | `937820` (915.84 KiB) | `0.999x` | `4.269x` | `4.265x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003622` (3.82 MiB) | `549988` (537.10 KiB) | `0.999x` | `7.279x` | `7.273x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003524` (3.82 MiB) | `5323` (5.20 KiB) | `0.999x` | `752.118x` | `751.456x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `1024` (1.00 KiB) | `4004021` (3.82 MiB) | `4887` (4.77 KiB) | `0.000x` | `819.321x` | `0.210x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004439` (7.63 MiB) | `5711` (5.58 KiB) | `0.999x` | `1401.583x` | `1400.805x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002159` (3.82 MiB) | `2866` (2.80 KiB) | `0.000x` | `1396.427x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3000000` (2.86 MiB) | `7003303` (6.68 MiB) | `5261` (5.14 KiB) | `0.428x` | `1331.173x` | `570.234x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003527` (3.82 MiB) | `4234` (4.13 KiB) | `0.999x` | `945.566x` | `944.733x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `58030` (56.67 KiB) | `4062525` (3.87 MiB) | `19170` (18.72 KiB) | `0.014x` | `211.921x` | `3.027x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `22051` (21.53 KiB) | `4025605` (3.84 MiB) | `16530` (16.14 KiB) | `0.005x` | `243.533x` | `1.334x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `25445` (24.85 KiB) | `4029082` (3.84 MiB) | `18952` (18.51 KiB) | `0.006x` | `212.594x` | `1.343x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `48191` (47.06 KiB) | `4051765` (3.86 MiB) | `13313` (13.00 KiB) | `0.012x` | `304.347x` | `3.620x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `49433` (48.27 KiB) | `4053229` (3.87 MiB) | `21089` (20.59 KiB) | `0.012x` | `192.196x` | `2.344x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `16873` (16.48 KiB) | `4019935` (3.83 MiB) | `17079` (16.68 KiB) | `0.004x` | `235.373x` | `0.988x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `91870` (89.72 KiB) | `4097408` (3.91 MiB) | `29800` (29.10 KiB) | `0.022x` | `137.497x` | `3.083x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `13001` (12.70 KiB) | `4016539` (3.83 MiB) | `13961` (13.63 KiB) | `0.003x` | `287.697x` | `0.931x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `28101` (27.44 KiB) | `4032756` (3.85 MiB) | `15694` (15.33 KiB) | `0.007x` | `256.962x` | `1.791x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `45607` (44.54 KiB) | `4048288` (3.86 MiB) | `28367` (27.70 KiB) | `0.011x` | `142.711x` | `1.608x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003557` (3.82 MiB) | `21190` (20.69 KiB) | `0.999x` | `188.936x` | `188.768x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004553` (7.63 MiB) | `2841973` (2.71 MiB) | `0.999x` | `2.817x` | `2.815x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004553` (7.63 MiB) | `3580403` (3.41 MiB) | `0.999x` | `2.236x` | `2.234x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003542` (3.82 MiB) | `5627` (5.50 KiB) | `0.999x` | `711.488x` | `710.858x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00000.parquet`: `35854` rows, `2901215` file bytes (2.77 MiB), `27844924` physical bytes (26.55 MiB), `31278557` encoded bytes (29.83 MiB), `2869144` compressed data bytes (2.74 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00001.parquet`: `35947` rows, `2855891` file bytes (2.72 MiB), `27614679` physical bytes (26.34 MiB), `31053255` encoded bytes (29.61 MiB), `2823278` compressed data bytes (2.69 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00002.parquet`: `35804` rows, `2914449` file bytes (2.78 MiB), `28045024` physical bytes (26.75 MiB), `31474794` encoded bytes (30.02 MiB), `2881147` compressed data bytes (2.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00003.parquet`: `36223` rows, `2871815` file bytes (2.74 MiB), `28041589` physical bytes (26.74 MiB), `31503298` encoded bytes (30.04 MiB), `2839357` compressed data bytes (2.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00004.parquet`: `35583` rows, `2826056` file bytes (2.70 MiB), `27676556` physical bytes (26.39 MiB), `31081381` encoded bytes (29.64 MiB), `2792861` compressed data bytes (2.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00005.parquet`: `35769` rows, `2819045` file bytes (2.69 MiB), `27421909` physical bytes (26.15 MiB), `30836725` encoded bytes (29.41 MiB), `2786608` compressed data bytes (2.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00006.parquet`: `36202` rows, `2876578` file bytes (2.74 MiB), `28032888` physical bytes (26.73 MiB), `31498445` encoded bytes (30.04 MiB), `2844096` compressed data bytes (2.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00007.parquet`: `36255` rows, `2862756` file bytes (2.73 MiB), `28048091` physical bytes (26.75 MiB), `31515255` encoded bytes (30.06 MiB), `2829875` compressed data bytes (2.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00008.parquet`: `36216` rows, `2841201` file bytes (2.71 MiB), `28069629` physical bytes (26.77 MiB), `31530645` encoded bytes (30.07 MiB), `2808479` compressed data bytes (2.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00009.parquet`: `35848` rows, `2846720` file bytes (2.71 MiB), `27649376` physical bytes (26.37 MiB), `31075234` encoded bytes (29.64 MiB), `2814555` compressed data bytes (2.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00010.parquet`: `36464` rows, `2842477` file bytes (2.71 MiB), `28033620` physical bytes (26.73 MiB), `31515184` encoded bytes (30.06 MiB), `2809866` compressed data bytes (2.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00011.parquet`: `36386` rows, `2828218` file bytes (2.70 MiB), `28069654` physical bytes (26.77 MiB), `31548883` encoded bytes (30.09 MiB), `2795709` compressed data bytes (2.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00012.parquet`: `36306` rows, `2835080` file bytes (2.70 MiB), `27829887` physical bytes (26.54 MiB), `31298522` encoded bytes (29.85 MiB), `2802936` compressed data bytes (2.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00013.parquet`: `36202` rows, `2867182` file bytes (2.73 MiB), `28030660` physical bytes (26.73 MiB), `31491495` encoded bytes (30.03 MiB), `2834469` compressed data bytes (2.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00014.parquet`: `35358` rows, `3278785` file bytes (3.13 MiB), `24474438` physical bytes (23.34 MiB), `27846716` encoded bytes (26.56 MiB), `3247937` compressed data bytes (3.10 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00015.parquet`: `36048` rows, `3439103` file bytes (3.28 MiB), `25184942` physical bytes (24.02 MiB), `28619288` encoded bytes (27.29 MiB), `3408670` compressed data bytes (3.25 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00016.parquet`: `35303` rows, `3446549` file bytes (3.29 MiB), `23228130` physical bytes (22.15 MiB), `26584728` encoded bytes (25.35 MiB), `3416266` compressed data bytes (3.26 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00017.parquet`: `34855` rows, `3464501` file bytes (3.30 MiB), `22488678` physical bytes (21.45 MiB), `25795509` encoded bytes (24.60 MiB), `3433749` compressed data bytes (3.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00018.parquet`: `34370` rows, `3551797` file bytes (3.39 MiB), `22419656` physical bytes (21.38 MiB), `25674965` encoded bytes (24.49 MiB), `3521443` compressed data bytes (3.36 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00019.parquet`: `34956` rows, `3444115` file bytes (3.28 MiB), `22245512` physical bytes (21.21 MiB), `25560962` encoded bytes (24.38 MiB), `3413354` compressed data bytes (3.26 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00020.parquet`: `34543` rows, `3508504` file bytes (3.35 MiB), `22211905` physical bytes (21.18 MiB), `25489492` encoded bytes (24.31 MiB), `3477582` compressed data bytes (3.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00021.parquet`: `35097` rows, `3458045` file bytes (3.30 MiB), `22434470` physical bytes (21.40 MiB), `25761882` encoded bytes (24.57 MiB), `3427394` compressed data bytes (3.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00022.parquet`: `34964` rows, `3464621` file bytes (3.30 MiB), `22426238` physical bytes (21.39 MiB), `25739925` encoded bytes (24.55 MiB), `3434313` compressed data bytes (3.28 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00023.parquet`: `34536` rows, `3527741` file bytes (3.36 MiB), `22414326` physical bytes (21.38 MiB), `25686389` encoded bytes (24.50 MiB), `3496963` compressed data bytes (3.33 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00024.parquet`: `35254` rows, `3450767` file bytes (3.29 MiB), `22398930` physical bytes (21.36 MiB), `25744144` encoded bytes (24.55 MiB), `3420602` compressed data bytes (3.26 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00025.parquet`: `34864` rows, `3462876` file bytes (3.30 MiB), `22219181` physical bytes (21.19 MiB), `25527567` encoded bytes (24.34 MiB), `3432165` compressed data bytes (3.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00026.parquet`: `35280` rows, `3443712` file bytes (3.28 MiB), `22435753` physical bytes (21.40 MiB), `25782214` encoded bytes (24.59 MiB), `3413206` compressed data bytes (3.26 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00027.parquet`: `34989` rows, `3475521` file bytes (3.31 MiB), `22448627` physical bytes (21.41 MiB), `25768091` encoded bytes (24.57 MiB), `3445061` compressed data bytes (3.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed/part-00028.parquet`: `4524` rows, `454524` file bytes (443.87 KiB), `2959352` physical bytes (2.82 MiB), `3390627` encoded bytes (3.23 MiB), `438909` compressed data bytes (428.62 KiB)
