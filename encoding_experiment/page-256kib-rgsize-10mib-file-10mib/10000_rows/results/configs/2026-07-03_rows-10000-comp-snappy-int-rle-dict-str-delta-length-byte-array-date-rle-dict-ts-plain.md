# ClickBench Parquet Experiment

- Started: `2026-07-03T14:42:14-04:00`
- Write elapsed: `174ms`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/10000_rows/parquet/rows-10000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
- Rows: `10000`
- Source TSV bytes for rows, reference only: `8365486` (7.98 MiB)
- Parquet physical bytes before page encoding: `7711890` (7.35 MiB)
- Encoded column bytes before codec compression: `4976255` (4.75 MiB)
- Compressed column data bytes after codec compression: `1139944` (1.09 MiB)
- Parquet file bytes: `1157925` (1.10 MiB)
- Physical/encoded ratio: `1.550x`
- Encoded/compressed-data ratio: `4.365x`
- Physical/compressed-data ratio: `6.765x`
- Physical/parquet-file ratio: `6.660x`
- Files: `1`

## Settings

- Compression: `snappy`
- Int encoding: `rle-dict`
- String encoding: `delta-length-byte-array`
- Date encoding: `rle-dict`
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
- Rows read and compared: `10000`
- Files read: `1`
- Elapsed: `72ms`
- Source TSV bytes checked: `8365486` (7.98 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-10000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-10000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `97608` (95.32 KiB) | `97621` (95.33 KiB) | `0.820x` | `1.000x` | `0.819x` | `190000` (185.55 KiB) |
| `JavaEnable` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `850` (850 B) | `789` (789 B) | `47.059x` | `1.077x` | `50.697x` | `10000` (9.77 KiB) |
| `Title` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:9` | `10000` | `2200605` (2.10 MiB) | `2220012` (2.12 MiB) | `356497` (348.14 KiB) | `0.991x` | `6.227x` | `6.173x` | `2200673` (2.10 MiB) |
| `GoodEvent` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `91` (91 B) | `459.770x` | `0.956x` | `439.560x` | `10000` (9.77 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `35183` (34.36 KiB) | `0.999x` | `2.276x` | `2.274x` | `190000` (185.55 KiB) |
| `EventDate` | `date` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `91` (91 B) | `459.770x` | `0.956x` | `439.560x` | `100000` (97.66 KiB) |
| `CounterID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `101` (101 B) | `105` (105 B) | `396.040x` | `0.962x` | `380.952x` | `20000` (19.53 KiB) |
| `ClientIP` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `8212` (8.02 KiB) | `7351` (7.18 KiB) | `4.871x` | `1.117x` | `5.441x` | `100283` (97.93 KiB) |
| `RegionID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `3687` (3.60 KiB) | `2779` (2.71 KiB) | `10.849x` | `1.327x` | `14.394x` | `24459` (23.89 KiB) |
| `UserID` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `10952` (10.70 KiB) | `10170` (9.93 KiB) | `7.305x` | `1.077x` | `7.866x` | `199987` (195.30 KiB) |
| `CounterClass` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `91` (91 B) | `459.770x` | `0.956x` | `439.560x` | `10000` (9.77 KiB) |
| `OS` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `2327` (2.27 KiB) | `1592` (1.55 KiB) | `17.190x` | `1.462x` | `25.126x` | `16130` (15.75 KiB) |
| `UserAgent` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `2298` (2.24 KiB) | `1873` (1.83 KiB) | `17.406x` | `1.227x` | `21.356x` | `11313` (11.05 KiB) |
| `URL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:4` | `10000` | `760012` (742.20 KiB) | `771041` (752.97 KiB) | `130733` (127.67 KiB) | `0.986x` | `5.898x` | `5.813x` | `760022` (742.21 KiB) |
| `Referer` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:4` | `10000` | `797798` (779.10 KiB) | `810095` (791.11 KiB) | `148995` (145.50 KiB) | `0.985x` | `5.437x` | `5.355x` | `797822` (779.12 KiB) |
| `IsRefresh` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `1352` (1.32 KiB) | `1359` (1.33 KiB) | `29.586x` | `0.995x` | `29.433x` | `10000` (9.77 KiB) |
| `RefererCategoryID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `4330` (4.23 KiB) | `2798` (2.73 KiB) | `9.238x` | `1.548x` | `14.296x` | `47318` (46.21 KiB) |
| `RefererRegionID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `3615` (3.53 KiB) | `2399` (2.34 KiB) | `11.065x` | `1.507x` | `16.674x` | `28586` (27.92 KiB) |
| `URLCategoryID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `546` (546 B) | `480` (480 B) | `73.260x` | `1.137x` | `83.333x` | `49424` (48.27 KiB) |
| `URLRegionID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `557` (557 B) | `505` (505 B) | `71.813x` | `1.103x` | `79.208x` | `29545` (28.85 KiB) |
| `ResolutionWidth` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `3278` (3.20 KiB) | `2495` (2.44 KiB) | `12.203x` | `1.314x` | `16.032x` | `39093` (38.18 KiB) |
| `ResolutionHeight` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `3235` (3.16 KiB) | `2421` (2.36 KiB) | `12.365x` | `1.336x` | `16.522x` | `34257` (33.45 KiB) |
| `ResolutionDepth` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `1094` (1.07 KiB) | `950` (950 B) | `36.563x` | `1.152x` | `42.105x` | `19813` (19.35 KiB) |
| `FlashMajor` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `1185` (1.16 KiB) | `964` (964 B) | `33.755x` | `1.229x` | `41.494x` | `18909` (18.47 KiB) |
| `FlashMinor` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `2093` (2.04 KiB) | `1625` (1.59 KiB) | `19.111x` | `1.288x` | `24.615x` | `11787` (11.51 KiB) |
| `FlashMinor2` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `37170` (36.30 KiB) | `41362` (40.39 KiB) | `4979` (4.86 KiB) | `0.899x` | `8.307x` | `7.465x` | `37170` (36.30 KiB) |
| `NetMajor` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `483` (483 B) | `444` (444 B) | `82.816x` | `1.088x` | `90.090x` | `10000` (9.77 KiB) |
| `NetMinor` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `438` (438 B) | `404` (404 B) | `91.324x` | `1.084x` | `99.010x` | `10000` (9.77 KiB) |
| `UserAgentMajor` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `2804` (2.74 KiB) | `2378` (2.32 KiB) | `14.265x` | `1.179x` | `16.821x` | `18125` (17.70 KiB) |
| `UserAgentMinor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `36930` (36.06 KiB) | `39806` (38.87 KiB) | `3635` (3.55 KiB) | `0.928x` | `10.951x` | `10.160x` | `37108` (36.24 KiB) |
| `CookieEnable` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `231` (231 B) | `237` (237 B) | `173.160x` | `0.975x` | `168.776x` | `10000` (9.77 KiB) |
| `JavascriptEnable` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `230` (230 B) | `236` (236 B) | `173.913x` | `0.975x` | `169.492x` | `10000` (9.77 KiB) |
| `IsMobile` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `304` (304 B) | `297` (297 B) | `131.579x` | `1.024x` | `134.680x` | `10000` (9.77 KiB) |
| `MobilePhone` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `355` (355 B) | `335` (335 B) | `112.676x` | `1.060x` | `119.403x` | `10011` (9.78 KiB) |
| `MobilePhoneModel` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `1140` (1.11 KiB) | `2455` (2.40 KiB) | `455` (455 B) | `0.464x` | `5.396x` | `2.505x` | `1140` (1.11 KiB) |
| `Params` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `0` (0 B) | `438` (438 B) | `67` (67 B) | `0.000x` | `6.537x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `6819` (6.66 KiB) | `6071` (5.93 KiB) | `5.866x` | `1.123x` | `6.589x` | `68529` (66.92 KiB) |
| `TraficSourceID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `2897` (2.83 KiB) | `2399` (2.34 KiB) | `13.807x` | `1.208x` | `16.674x` | `17645` (17.23 KiB) |
| `SearchEngineID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `2763` (2.70 KiB) | `1675` (1.64 KiB) | `14.477x` | `1.650x` | `23.881x` | `10166` (9.93 KiB) |
| `SearchPhrase` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `65246` (63.72 KiB) | `74810` (73.06 KiB) | `17809` (17.39 KiB) | `0.872x` | `4.201x` | `3.664x` | `65250` (63.72 KiB) |
| `AdvEngineID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `680` (680 B) | `600` (600 B) | `58.824x` | `1.133x` | `66.667x` | `10142` (9.90 KiB) |
| `IsArtifical` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `1351` (1.32 KiB) | `1358` (1.33 KiB) | `29.608x` | `0.995x` | `29.455x` | `10000` (9.77 KiB) |
| `WindowClientWidth` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `6544` (6.39 KiB) | `4654` (4.54 KiB) | `6.112x` | `1.406x` | `8.595x` | `37695` (36.81 KiB) |
| `WindowClientHeight` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `6872` (6.71 KiB) | `5980` (5.84 KiB) | `5.821x` | `1.149x` | `6.689x` | `30023` (29.32 KiB) |
| `ClientTimeZone` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `1228` (1.20 KiB) | `947` (947 B) | `32.573x` | `1.297x` | `42.239x` | `29805` (29.11 KiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80079` (78.20 KiB) | `33528` (32.74 KiB) | `0.999x` | `2.388x` | `2.386x` | `190000` (185.55 KiB) |
| `SilverlightVersion1` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `1479` (1.44 KiB) | `1132` (1.11 KiB) | `27.045x` | `1.307x` | `35.336x` | `10017` (9.78 KiB) |
| `SilverlightVersion2` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `832` (832 B) | `747` (747 B) | `48.077x` | `1.114x` | `53.548x` | `10000` (9.77 KiB) |
| `SilverlightVersion3` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `2142` (2.09 KiB) | `1527` (1.49 KiB) | `18.674x` | `1.403x` | `26.195x` | `26504` (25.88 KiB) |
| `SilverlightVersion4` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `91` (91 B) | `459.770x` | `0.956x` | `439.560x` | `10000` (9.77 KiB) |
| `PageCharset` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `196540` (191.93 KiB) | `198785` (194.13 KiB) | `10370` (10.13 KiB) | `0.989x` | `19.169x` | `18.953x` | `196540` (191.93 KiB) |
| `CodeVersion` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `231` (231 B) | `237` (237 B) | `173.160x` | `0.975x` | `168.776x` | `39532` (38.61 KiB) |
| `IsLink` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `91` (91 B) | `459.770x` | `0.956x` | `439.560x` | `10000` (9.77 KiB) |
| `IsDownload` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `91` (91 B) | `459.770x` | `0.956x` | `439.560x` | `10000` (9.77 KiB) |
| `IsNotBounce` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `91` (91 B) | `459.770x` | `0.956x` | `439.560x` | `10000` (9.77 KiB) |
| `FUniqID` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `10320` (10.08 KiB) | `9470` (9.25 KiB) | `7.752x` | `1.090x` | `8.448x` | `176860` (172.71 KiB) |
| `OriginalURL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `74750` (73.00 KiB) | `82311` (80.38 KiB) | `14544` (14.20 KiB) | `0.908x` | `5.659x` | `5.140x` | `74750` (73.00 KiB) |
| `HID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `35968` (35.12 KiB) | `35979` (35.14 KiB) | `1.112x` | `1.000x` | `1.112x` | `88404` (86.33 KiB) |
| `IsOldCounter` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `91` (91 B) | `459.770x` | `0.956x` | `439.560x` | `10000` (9.77 KiB) |
| `IsEvent` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `91` (91 B) | `459.770x` | `0.956x` | `439.560x` | `10000` (9.77 KiB) |
| `IsParameter` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `91` (91 B) | `459.770x` | `0.956x` | `439.560x` | `10000` (9.77 KiB) |
| `DontCountHits` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `194` (194 B) | `199` (199 B) | `206.186x` | `0.975x` | `201.005x` | `10000` (9.77 KiB) |
| `WithHash` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `91` (91 B) | `459.770x` | `0.956x` | `439.560x` | `10000` (9.77 KiB) |
| `HitColor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `10000` (9.77 KiB) | `10456` (10.21 KiB) | `630` (630 B) | `0.956x` | `16.597x` | `15.873x` | `10000` (9.77 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `35180` (34.36 KiB) | `0.999x` | `2.276x` | `2.274x` | `190000` (185.55 KiB) |
| `Age` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `1981` (1.93 KiB) | `1754` (1.71 KiB) | `20.192x` | `1.129x` | `22.805x` | `17412` (17.00 KiB) |
| `Sex` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `1433` (1.40 KiB) | `1252` (1.22 KiB) | `27.913x` | `1.145x` | `31.949x` | `10000` (9.77 KiB) |
| `Income` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `1474` (1.44 KiB) | `1313` (1.28 KiB) | `27.137x` | `1.123x` | `30.465x` | `10000` (9.77 KiB) |
| `Interests` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `4556` (4.45 KiB) | `3320` (3.24 KiB) | `8.780x` | `1.372x` | `12.048x` | `22570` (22.04 KiB) |
| `Robotness` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `3879` (3.79 KiB) | `3054` (2.98 KiB) | `10.312x` | `1.270x` | `13.098x` | `13859` (13.53 KiB) |
| `RemoteIP` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `8272` (8.08 KiB) | `7411` (7.24 KiB) | `4.836x` | `1.116x` | `5.397x` | `100205` (97.86 KiB) |
| `WindowName` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `106` (106 B) | `110` (110 B) | `377.358x` | `0.964x` | `363.636x` | `20010` (19.54 KiB) |
| `OpenerName` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `91` (91 B) | `459.770x` | `0.956x` | `439.560x` | `20000` (19.53 KiB) |
| `HistoryLength` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `99` (99 B) | `103` (103 B) | `404.040x` | `0.961x` | `388.350x` | `19996` (19.53 KiB) |
| `BrowserLanguage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `20156` (19.68 KiB) | `21070` (20.58 KiB) | `1516` (1.48 KiB) | `0.957x` | `13.898x` | `13.296x` | `20156` (19.68 KiB) |
| `BrowserCountry` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `32430` (31.67 KiB) | `35646` (34.81 KiB) | `3447` (3.37 KiB) | `0.910x` | `10.341x` | `9.408x` | `38645` (37.74 KiB) |
| `SocialNetwork` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `0` (0 B) | `438` (438 B) | `67` (67 B) | `0.000x` | `6.537x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `0` (0 B) | `438` (438 B) | `67` (67 B) | `0.000x` | `6.537x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `91` (91 B) | `459.770x` | `0.956x` | `439.560x` | `10000` (9.77 KiB) |
| `SendTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `91` (91 B) | `459.770x` | `0.956x` | `439.560x` | `10000` (9.77 KiB) |
| `DNSTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `2153` (2.10 KiB) | `1373` (1.34 KiB) | `18.579x` | `1.568x` | `29.133x` | `10125` (9.89 KiB) |
| `ConnectTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `3895` (3.80 KiB) | `2502` (2.44 KiB) | `10.270x` | `1.557x` | `15.987x` | `10493` (10.25 KiB) |
| `ResponseStartTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `13243` (12.93 KiB) | `11918` (11.64 KiB) | `3.020x` | `1.111x` | `3.356x` | `16080` (15.70 KiB) |
| `ResponseEndTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `11871` (11.59 KiB) | `10117` (9.88 KiB) | `3.370x` | `1.173x` | `3.954x` | `13892` (13.57 KiB) |
| `FetchTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `10559` (10.31 KiB) | `7448` (7.27 KiB) | `3.788x` | `1.418x` | `5.371x` | `11851` (11.57 KiB) |
| `SocialSourceNetworkID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `101` (101 B) | `105` (105 B) | `396.040x` | `0.962x` | `380.952x` | `10000` (9.77 KiB) |
| `SocialSourcePage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `42` (42 B) | `555` (555 B) | `125` (125 B) | `0.076x` | `4.440x` | `0.336x` | `42` (42 B) |
| `ParamPrice` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `107` (107 B) | `111` (111 B) | `747.664x` | `0.964x` | `720.721x` | `10000` (9.77 KiB) |
| `ParamOrderID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `0` (0 B) | `438` (438 B) | `67` (67 B) | `0.000x` | `6.537x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `30000` (29.30 KiB) | `30468` (29.75 KiB) | `1607` (1.57 KiB) | `0.985x` | `18.960x` | `18.668x` | `30000` (29.30 KiB) |
| `ParamCurrencyID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `91` (91 B) | `459.770x` | `0.956x` | `439.560x` | `10000` (9.77 KiB) |
| `OpenstatServiceName` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `560` (560 B) | `2299` (2.25 KiB) | `381` (381 B) | `0.244x` | `6.034x` | `1.470x` | `560` (560 B) |
| `OpenstatCampaignID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `272` (272 B) | `1583` (1.55 KiB) | `354` (354 B) | `0.172x` | `4.472x` | `0.768x` | `272` (272 B) |
| `OpenstatAdID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `358` (358 B) | `1829` (1.79 KiB) | `536` (536 B) | `0.196x` | `3.412x` | `0.668x` | `358` (358 B) |
| `OpenstatSourceID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `692` (692 B) | `2435` (2.38 KiB) | `419` (419 B) | `0.284x` | `5.811x` | `1.652x` | `692` (692 B) |
| `UTMSource` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `1621` (1.58 KiB) | `4324` (4.22 KiB) | `967` (967 B) | `0.375x` | `4.472x` | `1.676x` | `1621` (1.58 KiB) |
| `UTMMedium` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `502` (502 B) | `2717` (2.65 KiB) | `653` (653 B) | `0.185x` | `4.161x` | `0.769x` | `502` (502 B) |
| `UTMCampaign` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `3304` (3.23 KiB) | `6651` (6.50 KiB) | `1238` (1.21 KiB) | `0.497x` | `5.372x` | `2.669x` | `3304` (3.23 KiB) |
| `UTMContent` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `474` (474 B) | `1933` (1.89 KiB) | `399` (399 B) | `0.245x` | `4.845x` | `1.188x` | `474` (474 B) |
| `UTMTerm` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `1288` (1.26 KiB) | `3442` (3.36 KiB) | `580` (580 B) | `0.374x` | `5.934x` | `2.221x` | `1288` (1.26 KiB) |
| `FromTag` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `0` (0 B) | `438` (438 B) | `67` (67 B) | `0.000x` | `6.537x` | `0.000x` | `0` (0 B) |
| `HasGCLID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `397` (397 B) | `384` (384 B) | `100.756x` | `1.034x` | `104.167x` | `10000` (9.77 KiB) |
| `RefererHash` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `33011` (32.24 KiB) | `30243` (29.53 KiB) | `2.423x` | `1.092x` | `2.645x` | `193030` (188.51 KiB) |
| `URLHash` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `36727` (35.87 KiB) | `35293` (34.47 KiB) | `2.178x` | `1.041x` | `2.267x` | `193212` (188.68 KiB) |
| `CLID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `91` (91 B) | `459.770x` | `0.956x` | `439.560x` | `10000` (9.77 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/10000_rows/parquet/rows-10000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00000.parquet`: `10000` rows, `1157925` file bytes (1.10 MiB), `7711890` physical bytes (7.35 MiB), `4976255` encoded bytes (4.75 MiB), `1139944` compressed data bytes (1.09 MiB)
