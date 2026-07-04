# ClickBench Parquet Experiment

- Started: `2026-07-03T15:26:37-04:00`
- Write elapsed: `10.461s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `472710213` (450.81 MiB)
- Compressed column data bytes after codec compression: `472710213` (450.81 MiB)
- Parquet file bytes: `474125963` (452.16 MiB)
- Physical/encoded ratio: `1.507x`
- Encoded/compressed-data ratio: `1.000x`
- Physical/compressed-data ratio: `1.507x`
- Physical/parquet-file ratio: `1.503x`
- Files: `47`

## Settings

- Compression: `none`
- Int encoding: `plain`
- String encoding: `rle-dict`
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
- Rows read and compared: `1000000`
- Files read: `47`
- Elapsed: `6.719s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `8000000` (7.63 MiB) | `8007336` (7.64 MiB) | `8007336` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005857` (3.82 MiB) | `4005857` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `138409995` (132.00 MiB) | `29405339` (28.04 MiB) | `29405339` (28.04 MiB) | `4.707x` | `1.000x` | `4.707x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005850` (3.82 MiB) | `4005850` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `8000000` (7.63 MiB) | `8007341` (7.64 MiB) | `8007341` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `7947` (7.76 KiB) | `7947` (7.76 KiB) | `503.335x` | `1.000x` | `503.335x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005851` (3.82 MiB) | `4005851` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `8000000` (7.63 MiB) | `8007341` (7.64 MiB) | `8007341` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005853` (3.82 MiB) | `4005853` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005851` (3.82 MiB) | `4005851` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `88562192` (84.46 MiB) | `44345810` (42.29 MiB) | `44345810` (42.29 MiB) | `1.997x` | `1.000x` | `1.997x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `79583339` (75.90 MiB) | `34553829` (32.95 MiB) | `34553829` (32.95 MiB) | `2.303x` | `1.000x` | `2.303x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005852` (3.82 MiB) | `4005852` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005853` (3.82 MiB) | `4005853` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005857` (3.82 MiB) | `4005857` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005855` (3.82 MiB) | `4005855` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005857` (3.82 MiB) | `4005857` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005850` (3.82 MiB) | `4005850` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `3354477` (3.20 MiB) | `270935` (264.58 KiB) | `270935` (264.58 KiB) | `12.381x` | `1.000x` | `12.381x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005857` (3.82 MiB) | `4005857` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005850` (3.82 MiB) | `4005850` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `3767530` (3.59 MiB) | `153802` (150.20 KiB) | `153802` (150.20 KiB) | `24.496x` | `1.000x` | `24.496x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005852` (3.82 MiB) | `4005852` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005857` (3.82 MiB) | `4005857` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `81583` (79.67 KiB) | `27780` (27.13 KiB) | `27780` (27.13 KiB) | `2.937x` | `1.000x` | `2.937x` | `81583` (79.67 KiB) |
| `Params` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `0` (0 B) | `6459` (6.31 KiB) | `6459` (6.31 KiB) | `0.000x` | `1.000x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005853` (3.82 MiB) | `4005853` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `3528017` (3.36 MiB) | `1648718` (1.57 MiB) | `1648718` (1.57 MiB) | `2.140x` | `1.000x` | `2.140x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005853` (3.82 MiB) | `4005853` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005852` (3.82 MiB) | `4005852` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005855` (3.82 MiB) | `4005855` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005852` (3.82 MiB) | `4005852` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `8000000` (7.63 MiB) | `8007344` (7.64 MiB) | `8007344` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005853` (3.82 MiB) | `4005853` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005852` (3.82 MiB) | `4005852` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `13587860` (12.96 MiB) | `13152` (12.84 KiB) | `13152` (12.84 KiB) | `1033.140x` | `1.000x` | `1033.140x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005855` (3.82 MiB) | `4005855` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005853` (3.82 MiB) | `4005853` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `8000000` (7.63 MiB) | `8007339` (7.64 MiB) | `8007339` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `27797671` (26.51 MiB) | `21308797` (20.32 MiB) | `21308797` (20.32 MiB) | `1.305x` | `1.000x` | `1.305x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005853` (3.82 MiB) | `4005853` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005853` (3.82 MiB) | `4005853` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005853` (3.82 MiB) | `4005853` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005853` (3.82 MiB) | `4005853` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `1000000` (976.56 KiB) | `36230` (35.38 KiB) | `36230` (35.38 KiB) | `27.601x` | `1.000x` | `27.601x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `8000000` (7.63 MiB) | `8007340` (7.64 MiB) | `8007340` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005857` (3.82 MiB) | `4005857` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005855` (3.82 MiB) | `4005855` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005849` (3.82 MiB) | `4005849` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005857` (3.82 MiB) | `4005857` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005858` (3.82 MiB) | `4005858` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005855` (3.82 MiB) | `4005855` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `2001192` (1.91 MiB) | `45113` (44.06 KiB) | `45113` (44.06 KiB) | `44.360x` | `1.000x` | `44.360x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `3325142` (3.17 MiB) | `138792` (135.54 KiB) | `138792` (135.54 KiB) | `23.958x` | `1.000x` | `23.958x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `0` (0 B) | `6459` (6.31 KiB) | `6459` (6.31 KiB) | `0.000x` | `1.000x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `0` (0 B) | `6459` (6.31 KiB) | `6459` (6.31 KiB) | `0.000x` | `1.000x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005853` (3.82 MiB) | `4005853` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005855` (3.82 MiB) | `4005855` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005852` (3.82 MiB) | `4005852` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005851` (3.82 MiB) | `4005851` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005858` (3.82 MiB) | `4005858` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005858` (3.82 MiB) | `4005858` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `1024` (1.00 KiB) | `8213` (8.02 KiB) | `8213` (8.02 KiB) | `0.125x` | `1.000x` | `0.125x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `8000000` (7.63 MiB) | `8007341` (7.64 MiB) | `8007341` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `0` (0 B) | `6459` (6.31 KiB) | `6459` (6.31 KiB) | `0.000x` | `1.000x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `3000000` (2.86 MiB) | `7761` (7.58 KiB) | `7761` (7.58 KiB) | `386.548x` | `1.000x` | `386.548x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005853` (3.82 MiB) | `4005853` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `58030` (56.67 KiB) | `22687` (22.16 KiB) | `22687` (22.16 KiB) | `2.558x` | `1.000x` | `2.558x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `22051` (21.53 KiB) | `19492` (19.04 KiB) | `19492` (19.04 KiB) | `1.131x` | `1.000x` | `1.131x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `25445` (24.85 KiB) | `25201` (24.61 KiB) | `25201` (24.61 KiB) | `1.010x` | `1.000x` | `1.010x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `48191` (47.06 KiB) | `17802` (17.38 KiB) | `17802` (17.38 KiB) | `2.707x` | `1.000x` | `2.707x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `49433` (48.27 KiB) | `26324` (25.71 KiB) | `26324` (25.71 KiB) | `1.878x` | `1.000x` | `1.878x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `16873` (16.48 KiB) | `19386` (18.93 KiB) | `19386` (18.93 KiB) | `0.870x` | `1.000x` | `0.870x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `91870` (89.72 KiB) | `42458` (41.46 KiB) | `42458` (41.46 KiB) | `2.164x` | `1.000x` | `2.164x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `13001` (12.70 KiB) | `17620` (17.21 KiB) | `17620` (17.21 KiB) | `0.738x` | `1.000x` | `0.738x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `28101` (27.44 KiB) | `19725` (19.26 KiB) | `19725` (19.26 KiB) | `1.425x` | `1.000x` | `1.425x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `45607` (44.54 KiB) | `43164` (42.15 KiB) | `43164` (42.15 KiB) | `1.057x` | `1.000x` | `1.057x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `8000000` (7.63 MiB) | `8007342` (7.64 MiB) | `8007342` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `8000000` (7.63 MiB) | `8007339` (7.64 MiB) | `8007339` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00000.parquet`: `22598` rows, `10146681` file bytes (9.68 MiB), `17542847` physical bytes (16.73 MiB), `10115633` encoded bytes (9.65 MiB), `10115633` compressed data bytes (9.65 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00001.parquet`: `22594` rows, `10133700` file bytes (9.66 MiB), `17519787` physical bytes (16.71 MiB), `10102825` encoded bytes (9.63 MiB), `10102825` compressed data bytes (9.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00002.parquet`: `22625` rows, `10134837` file bytes (9.67 MiB), `17351916` physical bytes (16.55 MiB), `10103800` encoded bytes (9.64 MiB), `10103800` compressed data bytes (9.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00003.parquet`: `22524` rows, `10140325` file bytes (9.67 MiB), `17582092` physical bytes (16.77 MiB), `10108740` encoded bytes (9.64 MiB), `10108740` compressed data bytes (9.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00004.parquet`: `22596` rows, `10142584` file bytes (9.67 MiB), `17716110` physical bytes (16.90 MiB), `10111356` encoded bytes (9.64 MiB), `10111356` compressed data bytes (9.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00005.parquet`: `22655` rows, `10132931` file bytes (9.66 MiB), `17472327` physical bytes (16.66 MiB), `10101670` encoded bytes (9.63 MiB), `10101670` compressed data bytes (9.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00006.parquet`: `22599` rows, `10134931` file bytes (9.67 MiB), `17605943` physical bytes (16.79 MiB), `10103675` encoded bytes (9.64 MiB), `10103675` compressed data bytes (9.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00007.parquet`: `22733` rows, `10135621` file bytes (9.67 MiB), `17634555` physical bytes (16.82 MiB), `10103811` encoded bytes (9.64 MiB), `10103811` compressed data bytes (9.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00008.parquet`: `22809` rows, `10133134` file bytes (9.66 MiB), `17527232` physical bytes (16.72 MiB), `10101833` encoded bytes (9.63 MiB), `10101833` compressed data bytes (9.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00009.parquet`: `22737` rows, `10139571` file bytes (9.67 MiB), `17404062` physical bytes (16.60 MiB), `10108127` encoded bytes (9.64 MiB), `10108127` compressed data bytes (9.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00010.parquet`: `22710` rows, `10141949` file bytes (9.67 MiB), `17598983` physical bytes (16.78 MiB), `10110986` encoded bytes (9.64 MiB), `10110986` compressed data bytes (9.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00011.parquet`: `22608` rows, `10149361` file bytes (9.68 MiB), `17597435` physical bytes (16.78 MiB), `10117825` encoded bytes (9.65 MiB), `10117825` compressed data bytes (9.65 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00012.parquet`: `22736` rows, `10136589` file bytes (9.67 MiB), `17542644` physical bytes (16.73 MiB), `10105923` encoded bytes (9.64 MiB), `10105923` compressed data bytes (9.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00013.parquet`: `22742` rows, `10134341` file bytes (9.66 MiB), `17660305` physical bytes (16.84 MiB), `10103236` encoded bytes (9.64 MiB), `10103236` compressed data bytes (9.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00014.parquet`: `22840` rows, `10125743` file bytes (9.66 MiB), `17487432` physical bytes (16.68 MiB), `10094981` encoded bytes (9.63 MiB), `10094981` compressed data bytes (9.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00015.parquet`: `22548` rows, `10143376` file bytes (9.67 MiB), `17473318` physical bytes (16.66 MiB), `10112466` encoded bytes (9.64 MiB), `10112466` compressed data bytes (9.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00016.parquet`: `22728` rows, `10132785` file bytes (9.66 MiB), `17437323` physical bytes (16.63 MiB), `10102037` encoded bytes (9.63 MiB), `10102037` compressed data bytes (9.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00017.parquet`: `22881` rows, `10130884` file bytes (9.66 MiB), `17729774` physical bytes (16.91 MiB), `10099981` encoded bytes (9.63 MiB), `10099981` compressed data bytes (9.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00018.parquet`: `22827` rows, `10134195` file bytes (9.66 MiB), `17551455` physical bytes (16.74 MiB), `10103055` encoded bytes (9.64 MiB), `10103055` compressed data bytes (9.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00019.parquet`: `22860` rows, `10137333` file bytes (9.67 MiB), `17477414` physical bytes (16.67 MiB), `10106388` encoded bytes (9.64 MiB), `10106388` compressed data bytes (9.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00020.parquet`: `22649` rows, `10145169` file bytes (9.68 MiB), `17419786` physical bytes (16.61 MiB), `10114292` encoded bytes (9.65 MiB), `10114292` compressed data bytes (9.65 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00021.parquet`: `22653` rows, `10134120` file bytes (9.66 MiB), `17582355` physical bytes (16.77 MiB), `10103000` encoded bytes (9.63 MiB), `10103000` compressed data bytes (9.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00022.parquet`: `20738` rows, `10202366` file bytes (9.73 MiB), `15022203` physical bytes (14.33 MiB), `10172346` encoded bytes (9.70 MiB), `10172346` compressed data bytes (9.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00023.parquet`: `20175` rows, `10210270` file bytes (9.74 MiB), `13772607` physical bytes (13.13 MiB), `10180321` encoded bytes (9.71 MiB), `10180321` compressed data bytes (9.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00024.parquet`: `19913` rows, `10219546` file bytes (9.75 MiB), `13750413` physical bytes (13.11 MiB), `10189981` encoded bytes (9.72 MiB), `10189981` compressed data bytes (9.72 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00025.parquet`: `19164` rows, `10230108` file bytes (9.76 MiB), `13580633` physical bytes (12.95 MiB), `10200612` encoded bytes (9.73 MiB), `10200612` compressed data bytes (9.73 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00026.parquet`: `20177` rows, `10201423` file bytes (9.73 MiB), `13401895` physical bytes (12.78 MiB), `10171885` encoded bytes (9.70 MiB), `10171885` compressed data bytes (9.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00027.parquet`: `20591` rows, `10190511` file bytes (9.72 MiB), `13035004` physical bytes (12.43 MiB), `10160634` encoded bytes (9.69 MiB), `10160634` compressed data bytes (9.69 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00028.parquet`: `20344` rows, `10197623` file bytes (9.73 MiB), `13213869` physical bytes (12.60 MiB), `10167653` encoded bytes (9.70 MiB), `10167653` compressed data bytes (9.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00029.parquet`: `20184` rows, `10202512` file bytes (9.73 MiB), `13077961` physical bytes (12.47 MiB), `10172872` encoded bytes (9.70 MiB), `10172872` compressed data bytes (9.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00030.parquet`: `19678` rows, `10203645` file bytes (9.73 MiB), `12942539` physical bytes (12.34 MiB), `10173705` encoded bytes (9.70 MiB), `10173705` compressed data bytes (9.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00031.parquet`: `20667` rows, `10206257` file bytes (9.73 MiB), `13080318` physical bytes (12.47 MiB), `10176363` encoded bytes (9.70 MiB), `10176363` compressed data bytes (9.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00032.parquet`: `20374` rows, `10208990` file bytes (9.74 MiB), `13060699` physical bytes (12.46 MiB), `10178987` encoded bytes (9.71 MiB), `10178987` compressed data bytes (9.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00033.parquet`: `20156` rows, `10220251` file bytes (9.75 MiB), `13053075` physical bytes (12.45 MiB), `10190594` encoded bytes (9.72 MiB), `10190594` compressed data bytes (9.72 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00034.parquet`: `20658` rows, `10199029` file bytes (9.73 MiB), `13104822` physical bytes (12.50 MiB), `10169336` encoded bytes (9.70 MiB), `10169336` compressed data bytes (9.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00035.parquet`: `20392` rows, `10200333` file bytes (9.73 MiB), `13126165` physical bytes (12.52 MiB), `10170422` encoded bytes (9.70 MiB), `10170422` compressed data bytes (9.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00036.parquet`: `20405` rows, `10199227` file bytes (9.73 MiB), `13106979` physical bytes (12.50 MiB), `10169355` encoded bytes (9.70 MiB), `10169355` compressed data bytes (9.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00037.parquet`: `20458` rows, `10222499` file bytes (9.75 MiB), `13150088` physical bytes (12.54 MiB), `10193035` encoded bytes (9.72 MiB), `10193035` compressed data bytes (9.72 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00038.parquet`: `20161` rows, `10206891` file bytes (9.73 MiB), `13056583` physical bytes (12.45 MiB), `10176873` encoded bytes (9.71 MiB), `10176873` compressed data bytes (9.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00039.parquet`: `20104` rows, `10223394` file bytes (9.75 MiB), `13039362` physical bytes (12.44 MiB), `10193667` encoded bytes (9.72 MiB), `10193667` compressed data bytes (9.72 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00040.parquet`: `21001` rows, `10182553` file bytes (9.71 MiB), `13205618` physical bytes (12.59 MiB), `10153159` encoded bytes (9.68 MiB), `10153159` compressed data bytes (9.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00041.parquet`: `20717` rows, `10193981` file bytes (9.72 MiB), `13154058` physical bytes (12.54 MiB), `10164206` encoded bytes (9.69 MiB), `10164206` compressed data bytes (9.69 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00042.parquet`: `20561` rows, `10197094` file bytes (9.72 MiB), `13119266` physical bytes (12.51 MiB), `10167178` encoded bytes (9.70 MiB), `10167178` compressed data bytes (9.70 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00043.parquet`: `20642` rows, `10187131` file bytes (9.72 MiB), `13279868` physical bytes (12.66 MiB), `10157181` encoded bytes (9.69 MiB), `10157181` compressed data bytes (9.69 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00044.parquet`: `20892` rows, `10190623` file bytes (9.72 MiB), `13058722` physical bytes (12.45 MiB), `10161059` encoded bytes (9.69 MiB), `10161059` compressed data bytes (9.69 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00045.parquet`: `20378` rows, `10211071` file bytes (9.74 MiB), `13134228` physical bytes (12.53 MiB), `10181253` encoded bytes (9.71 MiB), `10181253` compressed data bytes (9.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-plain/part-00046.parquet`: `12218` rows, `6198475` file bytes (5.91 MiB), `7956554` physical bytes (7.59 MiB), `6181896` encoded bytes (5.90 MiB), `6181896` compressed data bytes (5.90 MiB)
