# ClickBench Parquet Experiment

- Started: `2026-07-03T19:44:23-04:00`
- Write elapsed: `11.669s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `717186145` (683.96 MiB)
- Compressed column data bytes after codec compression: `89440333` (85.30 MiB)
- Parquet file bytes: `90339703` (86.15 MiB)
- Physical/encoded ratio: `0.993x`
- Encoded/compressed-data ratio: `8.019x`
- Physical/compressed-data ratio: `7.965x`
- Physical/parquet-file ratio: `7.886x`
- Files: `29`

## Settings

- Compression: `zstd-3`
- Int encoding: `plain`
- String encoding: `delta-length-byte-array`
- Date encoding: `rle-dict`
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
- Files read: `29`
- Elapsed: `7.344s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004556` (7.63 MiB) | `8005354` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `65024` (63.50 KiB) | `0.999x` | `61.571x` | `61.516x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:523` | `1000000` | `138409995` (132.00 MiB) | `140026955` (133.54 MiB) | `14461085` (13.79 MiB) | `0.988x` | `9.683x` | `9.571x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `4916` (4.80 KiB) | `0.999x` | `814.388x` | `813.670x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004558` (7.63 MiB) | `2516679` (2.40 MiB) | `0.999x` | `3.181x` | `3.179x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4955` (4.84 KiB) | `5981` (5.84 KiB) | `807.265x` | `0.828x` | `668.784x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4931` (4.82 KiB) | `0.999x` | `811.910x` | `811.194x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003597` (3.82 MiB) | `408244` (398.68 KiB) | `0.999x` | `9.807x` | `9.798x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `191146` (186.67 KiB) | `0.999x` | `20.945x` | `20.926x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004547` (7.63 MiB) | `617935` (603.45 KiB) | `0.999x` | `12.954x` | `12.946x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003527` (3.82 MiB) | `4228` (4.13 KiB) | `0.999x` | `946.908x` | `946.074x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003591` (3.82 MiB) | `106696` (104.20 KiB) | `0.999x` | `37.523x` | `37.490x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `134614` (131.46 KiB) | `0.999x` | `29.741x` | `29.715x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:352` | `1000000` | `88562192` (84.46 MiB) | `89786522` (85.63 MiB) | `16029529` (15.29 MiB) | `0.986x` | `5.601x` | `5.525x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:328` | `1000000` | `79583339` (75.90 MiB) | `80835923` (77.09 MiB) | `14993079` (14.30 MiB) | `0.985x` | `5.392x` | `5.308x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `178854` (174.66 KiB) | `0.999x` | `22.385x` | `22.365x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003591` (3.82 MiB) | `275180` (268.73 KiB) | `0.999x` | `14.549x` | `14.536x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `231366` (225.94 KiB) | `0.999x` | `17.304x` | `17.289x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `87467` (85.42 KiB) | `0.999x` | `45.773x` | `45.732x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `48115` (46.99 KiB) | `0.999x` | `83.209x` | `83.134x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `187233` (182.84 KiB) | `0.999x` | `21.383x` | `21.364x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `186146` (181.78 KiB) | `0.999x` | `21.508x` | `21.489x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `81692` (79.78 KiB) | `0.999x` | `49.008x` | `48.964x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `53267` (52.02 KiB) | `0.999x` | `75.161x` | `75.093x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `129182` (126.15 KiB) | `0.999x` | `30.992x` | `30.964x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `3354477` (3.20 MiB) | `3707703` (3.54 MiB) | `266770` (260.52 KiB) | `0.905x` | `13.899x` | `12.574x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `26147` (25.53 KiB) | `0.999x` | `153.118x` | `152.981x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `25027` (24.44 KiB) | `0.999x` | `159.971x` | `159.827x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `155778` (152.13 KiB) | `0.999x` | `25.701x` | `25.678x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `3767530` (3.59 MiB) | `4016346` (3.83 MiB) | `150995` (147.46 KiB) | `0.938x` | `26.599x` | `24.951x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003538` (3.82 MiB) | `6078` (5.94 KiB) | `0.999x` | `658.693x` | `658.111x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003544` (3.82 MiB) | `6521` (6.37 KiB) | `0.999x` | `613.946x` | `613.403x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `29123` (28.44 KiB) | `0.999x` | `137.472x` | `137.348x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `22585` (22.06 KiB) | `0.999x` | `177.268x` | `177.109x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `81583` (79.67 KiB) | `231000` (225.59 KiB) | `30843` (30.12 KiB) | `0.353x` | `7.490x` | `2.645x` | `81583` (79.67 KiB) |
| `Params` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `41692` (40.71 KiB) | `3570` (3.49 KiB) | `0.000x` | `11.678x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003592` (3.82 MiB) | `323719` (316.13 KiB) | `0.999x` | `12.367x` | `12.356x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `288294` (281.54 KiB) | `0.999x` | `13.887x` | `13.875x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `101608` (99.23 KiB) | `0.999x` | `39.402x` | `39.367x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `3528017` (3.36 MiB) | `4222508` (4.03 MiB) | `813244` (794.18 KiB) | `0.836x` | `5.192x` | `4.338x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003579` (3.82 MiB) | `30642` (29.92 KiB) | `0.999x` | `130.657x` | `130.540x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `165119` (161.25 KiB) | `0.999x` | `24.247x` | `24.225x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `305818` (298.65 KiB) | `0.999x` | `13.091x` | `13.080x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `320157` (312.65 KiB) | `0.999x` | `12.505x` | `12.494x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `99854` (97.51 KiB) | `0.999x` | `40.094x` | `40.058x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004551` (7.63 MiB) | `2475238` (2.36 MiB) | `0.999x` | `3.234x` | `3.232x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `90428` (88.31 KiB) | `0.999x` | `44.274x` | `44.234x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `73703` (71.98 KiB) | `0.999x` | `54.321x` | `54.272x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `123751` (120.85 KiB) | `0.999x` | `32.352x` | `32.323x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4357` (4.25 KiB) | `0.999x` | `918.873x` | `918.063x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:85` | `1000000` | `13587860` (12.96 MiB) | `13657101` (13.02 MiB) | `18822` (18.38 KiB) | `0.995x` | `725.592x` | `721.914x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003550` (3.82 MiB) | `7092` (6.93 KiB) | `0.999x` | `564.516x` | `564.016x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003556` (3.82 MiB) | `56671` (55.34 KiB) | `0.999x` | `70.646x` | `70.583x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003551` (3.82 MiB) | `8045` (7.86 KiB) | `0.999x` | `497.645x` | `497.203x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003536` (3.82 MiB) | `25161` (24.57 KiB) | `0.999x` | `159.117x` | `158.976x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004556` (7.63 MiB) | `694026` (677.76 KiB) | `0.999x` | `11.534x` | `11.527x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:138` | `1000000` | `27797671` (26.51 MiB) | `28789340` (27.46 MiB) | `5605855` (5.35 MiB) | `0.966x` | `5.136x` | `4.959x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003643` (3.82 MiB) | `3878952` (3.70 MiB) | `0.999x` | `1.032x` | `1.031x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003527` (3.82 MiB) | `4228` (4.13 KiB) | `0.999x` | `946.908x` | `946.074x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003527` (3.82 MiB) | `4228` (4.13 KiB) | `0.999x` | `946.908x` | `946.074x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003527` (3.82 MiB) | `4228` (4.13 KiB) | `0.999x` | `946.908x` | `946.074x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `82496` (80.56 KiB) | `0.999x` | `48.531x` | `48.487x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003527` (3.82 MiB) | `4228` (4.13 KiB) | `0.999x` | `946.908x` | `946.074x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `1000000` (976.56 KiB) | `1043050` (1018.60 KiB) | `28261` (27.60 KiB) | `0.959x` | `36.908x` | `35.384x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004557` (7.63 MiB) | `2518514` (2.40 MiB) | `0.999x` | `3.178x` | `3.176x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `143144` (139.79 KiB) | `0.999x` | `27.969x` | `27.944x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `107558` (105.04 KiB) | `0.999x` | `37.223x` | `37.189x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `123055` (120.17 KiB) | `0.999x` | `32.535x` | `32.506x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `193482` (188.95 KiB) | `0.999x` | `20.692x` | `20.674x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `174006` (169.93 KiB) | `0.999x` | `23.008x` | `22.988x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003604` (3.82 MiB) | `426616` (416.62 KiB) | `0.999x` | `9.385x` | `9.376x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003581` (3.82 MiB) | `70871` (69.21 KiB) | `0.999x` | `56.491x` | `56.441x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003532` (3.82 MiB) | `4233` (4.13 KiB) | `0.999x` | `945.791x` | `944.956x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003571` (3.82 MiB) | `55202` (53.91 KiB) | `0.999x` | `72.526x` | `72.461x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `2001192` (1.91 MiB) | `2051163` (1.96 MiB) | `31938` (31.19 KiB) | `0.976x` | `64.223x` | `62.659x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `3325142` (3.17 MiB) | `3638794` (3.47 MiB) | `185144` (180.80 KiB) | `0.914x` | `19.654x` | `17.960x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `41692` (40.71 KiB) | `3570` (3.49 KiB) | `0.000x` | `11.678x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `41692` (40.71 KiB) | `3570` (3.49 KiB) | `0.000x` | `11.678x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003527` (3.82 MiB) | `4228` (4.13 KiB) | `0.999x` | `946.908x` | `946.074x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003543` (3.82 MiB) | `61457` (60.02 KiB) | `0.999x` | `65.144x` | `65.086x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `135374` (132.20 KiB) | `0.999x` | `29.574x` | `29.548x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003598` (3.82 MiB) | `333279` (325.47 KiB) | `0.999x` | `12.013x` | `12.002x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003646` (3.82 MiB) | `1246074` (1.19 MiB) | `0.999x` | `3.213x` | `3.210x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003640` (3.82 MiB) | `938075` (916.09 KiB) | `0.999x` | `4.268x` | `4.264x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003626` (3.82 MiB) | `550438` (537.54 KiB) | `0.999x` | `7.274x` | `7.267x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003534` (3.82 MiB) | `5333` (5.21 KiB) | `0.999x` | `750.710x` | `750.047x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `1024` (1.00 KiB) | `46292` (45.21 KiB) | `6095` (5.95 KiB) | `0.022x` | `7.595x` | `0.168x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004440` (7.63 MiB) | `5675` (5.54 KiB) | `0.999x` | `1410.474x` | `1409.692x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `41692` (40.71 KiB) | `3570` (3.49 KiB) | `0.000x` | `11.678x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `3000000` (2.86 MiB) | `3044016` (2.90 MiB) | `5811` (5.67 KiB) | `0.986x` | `523.837x` | `516.262x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003527` (3.82 MiB) | `4228` (4.13 KiB) | `0.999x` | `946.908x` | `946.074x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `58030` (56.67 KiB) | `216868` (211.79 KiB) | `27594` (26.95 KiB) | `0.268x` | `7.859x` | `2.103x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `22051` (21.53 KiB) | `122877` (120.00 KiB) | `23447` (22.90 KiB) | `0.179x` | `5.241x` | `0.940x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `25445` (24.85 KiB) | `130573` (127.51 KiB) | `28779` (28.10 KiB) | `0.195x` | `4.537x` | `0.884x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `48191` (47.06 KiB) | `155830` (152.18 KiB) | `18399` (17.97 KiB) | `0.309x` | `8.469x` | `2.619x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `49433` (48.27 KiB) | `188704` (184.28 KiB) | `33980` (33.18 KiB) | `0.262x` | `5.553x` | `1.455x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `16873` (16.48 KiB) | `133974` (130.83 KiB) | `22040` (21.52 KiB) | `0.126x` | `6.079x` | `0.766x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `91870` (89.72 KiB) | `253983` (248.03 KiB) | `47595` (46.48 KiB) | `0.362x` | `5.336x` | `1.930x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `13001` (12.70 KiB) | `94259` (92.05 KiB) | `18310` (17.88 KiB) | `0.138x` | `5.148x` | `0.710x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `28101` (27.44 KiB) | `128648` (125.63 KiB) | `19982` (19.51 KiB) | `0.218x` | `6.438x` | `1.406x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:57` | `1000000` | `45607` (44.54 KiB) | `211605` (206.65 KiB) | `36271` (35.42 KiB) | `0.216x` | `5.834x` | `1.257x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003558` (3.82 MiB) | `21114` (20.62 KiB) | `0.999x` | `189.616x` | `189.448x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004558` (7.63 MiB) | `2844347` (2.71 MiB) | `0.999x` | `2.814x` | `2.813x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004559` (7.63 MiB) | `3582683` (3.42 MiB) | `0.999x` | `2.234x` | `2.233x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003543` (3.82 MiB) | `5647` (5.51 KiB) | `0.999x` | `708.968x` | `708.341x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00000.parquet`: `35713` rows, `2944756` file bytes (2.81 MiB), `27742766` physical bytes (26.46 MiB), `27936236` encoded bytes (26.64 MiB), `2912711` compressed data bytes (2.78 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00001.parquet`: `35282` rows, `2867484` file bytes (2.73 MiB), `27096702` physical bytes (25.84 MiB), `27285553` encoded bytes (26.02 MiB), `2835057` compressed data bytes (2.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00002.parquet`: `35675` rows, `2962508` file bytes (2.83 MiB), `27929448` physical bytes (26.64 MiB), `28123301` encoded bytes (26.82 MiB), `2929282` compressed data bytes (2.79 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00003.parquet`: `35556` rows, `2883654` file bytes (2.75 MiB), `27535306` physical bytes (26.26 MiB), `27724563` encoded bytes (26.44 MiB), `2851386` compressed data bytes (2.72 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00004.parquet`: `35460` rows, `2874820` file bytes (2.74 MiB), `27583902` physical bytes (26.31 MiB), `27773608` encoded bytes (26.49 MiB), `2841782` compressed data bytes (2.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00005.parquet`: `36101` rows, `2899447` file bytes (2.77 MiB), `27729275` physical bytes (26.44 MiB), `27921241` encoded bytes (26.63 MiB), `2866920` compressed data bytes (2.73 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00006.parquet`: `35675` rows, `2881459` file bytes (2.75 MiB), `27557994` physical bytes (26.28 MiB), `27750551` encoded bytes (26.46 MiB), `2848969` compressed data bytes (2.72 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00007.parquet`: `36078` rows, `2923760` file bytes (2.79 MiB), `27928105` physical bytes (26.63 MiB), `28125869` encoded bytes (26.82 MiB), `2890890` compressed data bytes (2.76 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00008.parquet`: `36033` rows, `2889891` file bytes (2.76 MiB), `27960026` physical bytes (26.66 MiB), `28148774` encoded bytes (26.84 MiB), `2857313` compressed data bytes (2.72 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00009.parquet`: `35750` rows, `2884033` file bytes (2.75 MiB), `27542652` physical bytes (26.27 MiB), `27732914` encoded bytes (26.45 MiB), `2851756` compressed data bytes (2.72 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00010.parquet`: `36301` rows, `2898498` file bytes (2.76 MiB), `27906543` physical bytes (26.61 MiB), `28100489` encoded bytes (26.80 MiB), `2866015` compressed data bytes (2.73 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00011.parquet`: `36274` rows, `2873167` file bytes (2.74 MiB), `27965870` physical bytes (26.67 MiB), `28159899` encoded bytes (26.86 MiB), `2840709` compressed data bytes (2.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00012.parquet`: `35835` rows, `2881693` file bytes (2.75 MiB), `27520737` physical bytes (26.25 MiB), `27716325` encoded bytes (26.43 MiB), `2849584` compressed data bytes (2.72 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00013.parquet`: `35642` rows, `2859985` file bytes (2.73 MiB), `27588128` physical bytes (26.31 MiB), `27779271` encoded bytes (26.49 MiB), `2827470` compressed data bytes (2.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00014.parquet`: `35792` rows, `3295132` file bytes (3.14 MiB), `25061609` physical bytes (23.90 MiB), `25201966` encoded bytes (24.03 MiB), `3264007` compressed data bytes (3.11 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00015.parquet`: `35786` rows, `3440871` file bytes (3.28 MiB), `24927069` physical bytes (23.77 MiB), `25054231` encoded bytes (23.89 MiB), `3410605` compressed data bytes (3.25 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00016.parquet`: `34976` rows, `3449859` file bytes (3.29 MiB), `23260144` physical bytes (22.18 MiB), `23405428` encoded bytes (22.32 MiB), `3419643` compressed data bytes (3.26 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00017.parquet`: `34861` rows, `3487954` file bytes (3.33 MiB), `22419439` physical bytes (21.38 MiB), `22567896` encoded bytes (21.52 MiB), `3457222` compressed data bytes (3.30 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00018.parquet`: `34278` rows, `3652546` file bytes (3.48 MiB), `22533142` physical bytes (21.49 MiB), `22677095` encoded bytes (21.63 MiB), `3622168` compressed data bytes (3.45 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00019.parquet`: `34925` rows, `3425096` file bytes (3.27 MiB), `22048462` physical bytes (21.03 MiB), `22201008` encoded bytes (21.17 MiB), `3394381` compressed data bytes (3.24 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00020.parquet`: `34529` rows, `3584130` file bytes (3.42 MiB), `22303595` physical bytes (21.27 MiB), `22456052` encoded bytes (21.42 MiB), `3553444` compressed data bytes (3.39 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00021.parquet`: `34987` rows, `3506423` file bytes (3.34 MiB), `22355419` physical bytes (21.32 MiB), `22506863` encoded bytes (21.46 MiB), `3475532` compressed data bytes (3.31 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00022.parquet`: `34942` rows, `3505582` file bytes (3.34 MiB), `22359911` physical bytes (21.32 MiB), `22507930` encoded bytes (21.47 MiB), `3475161` compressed data bytes (3.31 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00023.parquet`: `34381` rows, `3567109` file bytes (3.40 MiB), `22390132` physical bytes (21.35 MiB), `22536236` encoded bytes (21.49 MiB), `3536441` compressed data bytes (3.37 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00024.parquet`: `34943` rows, `3468600` file bytes (3.31 MiB), `22180473` physical bytes (21.15 MiB), `22327778` encoded bytes (21.29 MiB), `3438534` compressed data bytes (3.28 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00025.parquet`: `34829` rows, `3481215` file bytes (3.32 MiB), `22183788` physical bytes (21.16 MiB), `22331977` encoded bytes (21.30 MiB), `3450454` compressed data bytes (3.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00026.parquet`: `34971` rows, `3501065` file bytes (3.34 MiB), `22345110` physical bytes (21.31 MiB), `22496855` encoded bytes (21.45 MiB), `3470357` compressed data bytes (3.31 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00027.parquet`: `34883` rows, `3471602` file bytes (3.31 MiB), `22230487` physical bytes (21.20 MiB), `22381589` encoded bytes (21.34 MiB), `3441262` compressed data bytes (3.28 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00028.parquet`: `9542` rows, `977364` file bytes (954.46 KiB), `6212390` physical bytes (5.92 MiB), `6254647` encoded bytes (5.96 MiB), `961278` compressed data bytes (938.75 KiB)
