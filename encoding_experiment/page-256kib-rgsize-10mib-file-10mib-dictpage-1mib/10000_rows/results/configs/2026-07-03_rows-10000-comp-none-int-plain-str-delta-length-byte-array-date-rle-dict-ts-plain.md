# ClickBench Parquet Experiment

- Started: `2026-07-03T15:25:51-04:00`
- Write elapsed: `113ms`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/10000_rows/parquet/rows-10000-comp-none-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain`
- Rows: `10000`
- Source TSV bytes for rows, reference only: `8365486` (7.98 MiB)
- Parquet physical bytes before page encoding: `7711890` (7.35 MiB)
- Encoded column bytes before codec compression: `7773379` (7.41 MiB)
- Compressed column data bytes after codec compression: `7773379` (7.41 MiB)
- Parquet file bytes: `7790865` (7.43 MiB)
- Physical/encoded ratio: `0.992x`
- Encoded/compressed-data ratio: `1.000x`
- Physical/compressed-data ratio: `0.992x`
- Physical/parquet-file ratio: `0.990x`
- Files: `1`

## Settings

- Compression: `none`
- Int encoding: `plain`
- String encoding: `delta-length-byte-array`
- Date encoding: `rle-dict`
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
- Rows read and compared: `10000`
- Files read: `1`
- Elapsed: `72ms`
- Source TSV bytes checked: `8365486` (7.98 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-10000-comp-none-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-10000-comp-none-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `80080` (78.20 KiB) | `0.999x` | `1.000x` | `0.999x` | `190000` (185.55 KiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `Title` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:9` | `10000` | `2200605` (2.10 MiB) | `2220013` (2.12 MiB) | `2220013` (2.12 MiB) | `0.991x` | `1.000x` | `0.991x` | `2200673` (2.10 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `40063` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `80080` (78.20 KiB) | `0.999x` | `1.000x` | `0.999x` | `190000` (185.55 KiB) |
| `EventDate` | `date` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `87` (87 B) | `459.770x` | `1.000x` | `459.770x` | `100000` (97.66 KiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `20000` (19.53 KiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `100283` (97.93 KiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `24459` (23.89 KiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `80080` (78.20 KiB) | `0.999x` | `1.000x` | `0.999x` | `199987` (195.30 KiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `16130` (15.75 KiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `11313` (11.05 KiB) |
| `URL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:4` | `10000` | `760012` (742.20 KiB) | `771042` (752.97 KiB) | `771042` (752.97 KiB) | `0.986x` | `1.000x` | `0.986x` | `760022` (742.21 KiB) |
| `Referer` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:4` | `10000` | `797798` (779.10 KiB) | `810095` (791.11 KiB) | `810095` (791.11 KiB) | `0.985x` | `1.000x` | `0.985x` | `797822` (779.12 KiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `47318` (46.21 KiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `28586` (27.92 KiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `49424` (48.27 KiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `29545` (28.85 KiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `39093` (38.18 KiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `34257` (33.45 KiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `19813` (19.35 KiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `18909` (18.47 KiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `11787` (11.51 KiB) |
| `FlashMinor2` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `37170` (36.30 KiB) | `41363` (40.39 KiB) | `41363` (40.39 KiB) | `0.899x` | `1.000x` | `0.899x` | `37170` (36.30 KiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `18125` (17.70 KiB) |
| `UserAgentMinor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `36930` (36.06 KiB) | `39807` (38.87 KiB) | `39807` (38.87 KiB) | `0.928x` | `1.000x` | `0.928x` | `37108` (36.24 KiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10011` (9.78 KiB) |
| `MobilePhoneModel` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `1140` (1.11 KiB) | `2455` (2.40 KiB) | `2455` (2.40 KiB) | `0.464x` | `1.000x` | `0.464x` | `1140` (1.11 KiB) |
| `Params` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `0` (0 B) | `439` (439 B) | `439` (439 B) | `0.000x` | `1.000x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `68529` (66.92 KiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `17645` (17.23 KiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10166` (9.93 KiB) |
| `SearchPhrase` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `65246` (63.72 KiB) | `74809` (73.06 KiB) | `74809` (73.06 KiB) | `0.872x` | `1.000x` | `0.872x` | `65250` (63.72 KiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10142` (9.90 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `37695` (36.81 KiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `30023` (29.32 KiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `29805` (29.11 KiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `80080` (78.20 KiB) | `0.999x` | `1.000x` | `0.999x` | `190000` (185.55 KiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10017` (9.78 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `26504` (25.88 KiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `PageCharset` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `196540` (191.93 KiB) | `198785` (194.13 KiB) | `198785` (194.13 KiB) | `0.989x` | `1.000x` | `0.989x` | `196540` (191.93 KiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `39532` (38.61 KiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `80080` (78.20 KiB) | `0.999x` | `1.000x` | `0.999x` | `176860` (172.71 KiB) |
| `OriginalURL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `74750` (73.00 KiB) | `82311` (80.38 KiB) | `82311` (80.38 KiB) | `0.908x` | `1.000x` | `0.908x` | `74750` (73.00 KiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `88404` (86.33 KiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `HitColor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `10000` (9.77 KiB) | `10457` (10.21 KiB) | `10457` (10.21 KiB) | `0.956x` | `1.000x` | `0.956x` | `10000` (9.77 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `80080` (78.20 KiB) | `0.999x` | `1.000x` | `0.999x` | `190000` (185.55 KiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `17412` (17.00 KiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `22570` (22.04 KiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `13859` (13.53 KiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `100205` (97.86 KiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `20010` (19.54 KiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `20000` (19.53 KiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `19996` (19.53 KiB) |
| `BrowserLanguage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `20156` (19.68 KiB) | `21071` (20.58 KiB) | `21071` (20.58 KiB) | `0.957x` | `1.000x` | `0.957x` | `20156` (19.68 KiB) |
| `BrowserCountry` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `32430` (31.67 KiB) | `35647` (34.81 KiB) | `35647` (34.81 KiB) | `0.910x` | `1.000x` | `0.910x` | `38645` (37.74 KiB) |
| `SocialNetwork` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `0` (0 B) | `439` (439 B) | `439` (439 B) | `0.000x` | `1.000x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `0` (0 B) | `439` (439 B) | `439` (439 B) | `0.000x` | `1.000x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `40063` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10125` (9.89 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10493` (10.25 KiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `16080` (15.70 KiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `13892` (13.57 KiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `11851` (11.57 KiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `SocialSourcePage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `42` (42 B) | `554` (554 B) | `554` (554 B) | `0.076x` | `1.000x` | `0.076x` | `42` (42 B) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `80080` (78.20 KiB) | `0.999x` | `1.000x` | `0.999x` | `10000` (9.77 KiB) |
| `ParamOrderID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `0` (0 B) | `439` (439 B) | `439` (439 B) | `0.000x` | `1.000x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `30000` (29.30 KiB) | `30469` (29.75 KiB) | `30469` (29.75 KiB) | `0.985x` | `1.000x` | `0.985x` | `30000` (29.30 KiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `OpenstatServiceName` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `560` (560 B) | `2299` (2.25 KiB) | `2299` (2.25 KiB) | `0.244x` | `1.000x` | `0.244x` | `560` (560 B) |
| `OpenstatCampaignID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `272` (272 B) | `1583` (1.55 KiB) | `1583` (1.55 KiB) | `0.172x` | `1.000x` | `0.172x` | `272` (272 B) |
| `OpenstatAdID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `358` (358 B) | `1828` (1.79 KiB) | `1828` (1.79 KiB) | `0.196x` | `1.000x` | `0.196x` | `358` (358 B) |
| `OpenstatSourceID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `692` (692 B) | `2435` (2.38 KiB) | `2435` (2.38 KiB) | `0.284x` | `1.000x` | `0.284x` | `692` (692 B) |
| `UTMSource` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `1621` (1.58 KiB) | `4324` (4.22 KiB) | `4324` (4.22 KiB) | `0.375x` | `1.000x` | `0.375x` | `1621` (1.58 KiB) |
| `UTMMedium` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `502` (502 B) | `2717` (2.65 KiB) | `2717` (2.65 KiB) | `0.185x` | `1.000x` | `0.185x` | `502` (502 B) |
| `UTMCampaign` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `3304` (3.23 KiB) | `6651` (6.50 KiB) | `6651` (6.50 KiB) | `0.497x` | `1.000x` | `0.497x` | `3304` (3.23 KiB) |
| `UTMContent` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `474` (474 B) | `1933` (1.89 KiB) | `1933` (1.89 KiB) | `0.245x` | `1.000x` | `0.245x` | `474` (474 B) |
| `UTMTerm` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `1288` (1.26 KiB) | `3443` (3.36 KiB) | `3443` (3.36 KiB) | `0.374x` | `1.000x` | `0.374x` | `1288` (1.26 KiB) |
| `FromTag` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:1` | `10000` | `0` (0 B) | `439` (439 B) | `439` (439 B) | `0.000x` | `1.000x` | `0.000x` | `0` (0 B) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `80080` (78.20 KiB) | `0.999x` | `1.000x` | `0.999x` | `193030` (188.51 KiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `80080` (78.20 KiB) | `0.999x` | `1.000x` | `0.999x` | `193212` (188.68 KiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `40064` (39.12 KiB) | `0.998x` | `1.000x` | `0.998x` | `10000` (9.77 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/10000_rows/parquet/rows-10000-comp-none-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00000.parquet`: `10000` rows, `7790865` file bytes (7.43 MiB), `7711890` physical bytes (7.35 MiB), `7773379` encoded bytes (7.41 MiB), `7773379` compressed data bytes (7.41 MiB)
