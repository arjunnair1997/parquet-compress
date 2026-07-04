# ClickBench Parquet Experiment

- Started: `2026-07-03T14:56:06-04:00`
- Write elapsed: `10.84s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `470646188` (448.84 MiB)
- Compressed column data bytes after codec compression: `470646188` (448.84 MiB)
- Parquet file bytes: `472065040` (450.20 MiB)
- Physical/encoded ratio: `1.514x`
- Encoded/compressed-data ratio: `1.000x`
- Physical/compressed-data ratio: `1.514x`
- Physical/parquet-file ratio: `1.509x`
- Files: `47`

## Settings

- Compression: `none`
- Int encoding: `plain`
- String encoding: `rle-dict`
- Date encoding: `rle-dict`
- Timestamp encoding: `rle-dict`
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
- Rows read and compared: `1000000`
- Files read: `47`
- Elapsed: `6.772s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `8000000` (7.63 MiB) | `8007344` (7.64 MiB) | `8007344` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `138409995` (132.00 MiB) | `29434285` (28.07 MiB) | `29434285` (28.07 MiB) | `4.702x` | `1.000x` | `4.702x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005852` (3.82 MiB) | `4005852` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `8000000` (7.63 MiB) | `7349361` (7.01 MiB) | `7349361` (7.01 MiB) | `1.089x` | `1.000x` | `1.089x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `7947` (7.76 KiB) | `7947` (7.76 KiB) | `503.335x` | `1.000x` | `503.335x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005853` (3.82 MiB) | `4005853` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005855` (3.82 MiB) | `4005855` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `8000000` (7.63 MiB) | `8007342` (7.64 MiB) | `8007342` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005850` (3.82 MiB) | `4005850` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005855` (3.82 MiB) | `4005855` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `88562192` (84.46 MiB) | `44347200` (42.29 MiB) | `44347200` (42.29 MiB) | `1.997x` | `1.000x` | `1.997x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `79583339` (75.90 MiB) | `34568866` (32.97 MiB) | `34568866` (32.97 MiB) | `2.302x` | `1.000x` | `2.302x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005853` (3.82 MiB) | `4005853` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005849` (3.82 MiB) | `4005849` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005852` (3.82 MiB) | `4005852` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005853` (3.82 MiB) | `4005853` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005855` (3.82 MiB) | `4005855` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005853` (3.82 MiB) | `4005853` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005855` (3.82 MiB) | `4005855` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `3354477` (3.20 MiB) | `270712` (264.37 KiB) | `270712` (264.37 KiB) | `12.391x` | `1.000x` | `12.391x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005853` (3.82 MiB) | `4005853` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `3767530` (3.59 MiB) | `153226` (149.63 KiB) | `153226` (149.63 KiB) | `24.588x` | `1.000x` | `24.588x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005855` (3.82 MiB) | `4005855` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005851` (3.82 MiB) | `4005851` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `81583` (79.67 KiB) | `27804` (27.15 KiB) | `27804` (27.15 KiB) | `2.934x` | `1.000x` | `2.934x` | `81583` (79.67 KiB) |
| `Params` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `0` (0 B) | `6459` (6.31 KiB) | `6459` (6.31 KiB) | `0.000x` | `1.000x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `3528017` (3.36 MiB) | `1648351` (1.57 MiB) | `1648351` (1.57 MiB) | `2.140x` | `1.000x` | `2.140x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005857` (3.82 MiB) | `4005857` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005848` (3.82 MiB) | `4005848` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `8000000` (7.63 MiB) | `7222097` (6.89 MiB) | `7222097` (6.89 MiB) | `1.108x` | `1.000x` | `1.108x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005857` (3.82 MiB) | `4005857` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005852` (3.82 MiB) | `4005852` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `13587860` (12.96 MiB) | `13155` (12.85 KiB) | `13155` (12.85 KiB) | `1032.905x` | `1.000x` | `1032.905x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005852` (3.82 MiB) | `4005852` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005852` (3.82 MiB) | `4005852` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005850` (3.82 MiB) | `4005850` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `8000000` (7.63 MiB) | `8007345` (7.64 MiB) | `8007345` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `27797671` (26.51 MiB) | `21303356` (20.32 MiB) | `21303356` (20.32 MiB) | `1.305x` | `1.000x` | `1.305x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005850` (3.82 MiB) | `4005850` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005850` (3.82 MiB) | `4005850` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005850` (3.82 MiB) | `4005850` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005853` (3.82 MiB) | `4005853` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005850` (3.82 MiB) | `4005850` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `1000000` (976.56 KiB) | `36204` (35.36 KiB) | `36204` (35.36 KiB) | `27.621x` | `1.000x` | `27.621x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `8000000` (7.63 MiB) | `7347509` (7.01 MiB) | `7347509` (7.01 MiB) | `1.089x` | `1.000x` | `1.089x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005851` (3.82 MiB) | `4005851` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005855` (3.82 MiB) | `4005855` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005855` (3.82 MiB) | `4005855` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005852` (3.82 MiB) | `4005852` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005851` (3.82 MiB) | `4005851` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005858` (3.82 MiB) | `4005858` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005850` (3.82 MiB) | `4005850` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `2001192` (1.91 MiB) | `44847` (43.80 KiB) | `44847` (43.80 KiB) | `44.623x` | `1.000x` | `44.623x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `3325142` (3.17 MiB) | `138792` (135.54 KiB) | `138792` (135.54 KiB) | `23.958x` | `1.000x` | `23.958x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `0` (0 B) | `6459` (6.31 KiB) | `6459` (6.31 KiB) | `0.000x` | `1.000x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `0` (0 B) | `6459` (6.31 KiB) | `6459` (6.31 KiB) | `0.000x` | `1.000x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005850` (3.82 MiB) | `4005850` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005850` (3.82 MiB) | `4005850` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005856` (3.82 MiB) | `4005856` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005853` (3.82 MiB) | `4005853` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005859` (3.82 MiB) | `4005859` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005852` (3.82 MiB) | `4005852` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005855` (3.82 MiB) | `4005855` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `1024` (1.00 KiB) | `8138` (7.95 KiB) | `8138` (7.95 KiB) | `0.126x` | `1.000x` | `0.126x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `8000000` (7.63 MiB) | `8007343` (7.64 MiB) | `8007343` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `0` (0 B) | `6459` (6.31 KiB) | `6459` (6.31 KiB) | `0.000x` | `1.000x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `3000000` (2.86 MiB) | `7761` (7.58 KiB) | `7761` (7.58 KiB) | `386.548x` | `1.000x` | `386.548x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005850` (3.82 MiB) | `4005850` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `58030` (56.67 KiB) | `22693` (22.16 KiB) | `22693` (22.16 KiB) | `2.557x` | `1.000x` | `2.557x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `22051` (21.53 KiB) | `19355` (18.90 KiB) | `19355` (18.90 KiB) | `1.139x` | `1.000x` | `1.139x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `25445` (24.85 KiB) | `25557` (24.96 KiB) | `25557` (24.96 KiB) | `0.996x` | `1.000x` | `0.996x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `48191` (47.06 KiB) | `17920` (17.50 KiB) | `17920` (17.50 KiB) | `2.689x` | `1.000x` | `2.689x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `49433` (48.27 KiB) | `26261` (25.65 KiB) | `26261` (25.65 KiB) | `1.882x` | `1.000x` | `1.882x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `16873` (16.48 KiB) | `19443` (18.99 KiB) | `19443` (18.99 KiB) | `0.868x` | `1.000x` | `0.868x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `91870` (89.72 KiB) | `42397` (41.40 KiB) | `42397` (41.40 KiB) | `2.167x` | `1.000x` | `2.167x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `13001` (12.70 KiB) | `17725` (17.31 KiB) | `17725` (17.31 KiB) | `0.733x` | `1.000x` | `0.733x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `28101` (27.44 KiB) | `19566` (19.11 KiB) | `19566` (19.11 KiB) | `1.436x` | `1.000x` | `1.436x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:93, DICTIONARY_PAGE/PLAIN:93` | `1000000` | `45607` (44.54 KiB) | `43588` (42.57 KiB) | `43588` (42.57 KiB) | `1.046x` | `1.000x` | `1.046x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005854` (3.82 MiB) | `4005854` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `8000000` (7.63 MiB) | `8007342` (7.64 MiB) | `8007342` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `8000000` (7.63 MiB) | `8007341` (7.64 MiB) | `8007341` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:93` | `1000000` | `4000000` (3.81 MiB) | `4005853` (3.82 MiB) | `4005853` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00000.parquet`: `22813` rows, `10094224` file bytes (9.63 MiB), `17715500` physical bytes (16.89 MiB), `10063106` encoded bytes (9.60 MiB), `10063106` compressed data bytes (9.60 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00001.parquet`: `22825` rows, `10082540` file bytes (9.62 MiB), `17678669` physical bytes (16.86 MiB), `10051562` encoded bytes (9.59 MiB), `10051562` compressed data bytes (9.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00002.parquet`: `22846` rows, `10083443` file bytes (9.62 MiB), `17552742` physical bytes (16.74 MiB), `10052062` encoded bytes (9.59 MiB), `10052062` compressed data bytes (9.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00003.parquet`: `22749` rows, `10088348` file bytes (9.62 MiB), `17769088` physical bytes (16.95 MiB), `10056691` encoded bytes (9.59 MiB), `10056691` compressed data bytes (9.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00004.parquet`: `22838` rows, `10084496` file bytes (9.62 MiB), `17845177` physical bytes (17.02 MiB), `10053485` encoded bytes (9.59 MiB), `10053485` compressed data bytes (9.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00005.parquet`: `22846` rows, `10079235` file bytes (9.61 MiB), `17626585` physical bytes (16.81 MiB), `10047909` encoded bytes (9.58 MiB), `10047909` compressed data bytes (9.58 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00006.parquet`: `22835` rows, `10086201` file bytes (9.62 MiB), `17770088` physical bytes (16.95 MiB), `10054312` encoded bytes (9.59 MiB), `10054312` compressed data bytes (9.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00007.parquet`: `22950` rows, `10083830` file bytes (9.62 MiB), `17821175` physical bytes (17.00 MiB), `10052500` encoded bytes (9.59 MiB), `10052500` compressed data bytes (9.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00008.parquet`: `22997` rows, `10087218` file bytes (9.62 MiB), `17697562` physical bytes (16.88 MiB), `10056164` encoded bytes (9.59 MiB), `10056164` compressed data bytes (9.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00009.parquet`: `23015` rows, `10079863` file bytes (9.61 MiB), `17621968` physical bytes (16.81 MiB), `10048717` encoded bytes (9.58 MiB), `10048717` compressed data bytes (9.58 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00010.parquet`: `22924` rows, `10083199` file bytes (9.62 MiB), `17767302` physical bytes (16.94 MiB), `10052022` encoded bytes (9.59 MiB), `10052022` compressed data bytes (9.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00011.parquet`: `22807` rows, `10086343` file bytes (9.62 MiB), `17740084` physical bytes (16.92 MiB), `10055169` encoded bytes (9.59 MiB), `10055169` compressed data bytes (9.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00012.parquet`: `22959` rows, `10081825` file bytes (9.61 MiB), `17762465` physical bytes (16.94 MiB), `10050492` encoded bytes (9.58 MiB), `10050492` compressed data bytes (9.58 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00013.parquet`: `22983` rows, `10076996` file bytes (9.61 MiB), `17808920` physical bytes (16.98 MiB), `10046045` encoded bytes (9.58 MiB), `10046045` compressed data bytes (9.58 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00014.parquet`: `23043` rows, `10081305` file bytes (9.61 MiB), `17613932` physical bytes (16.80 MiB), `10050381` encoded bytes (9.58 MiB), `10050381` compressed data bytes (9.58 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00015.parquet`: `22686` rows, `10092180` file bytes (9.62 MiB), `17657245` physical bytes (16.84 MiB), `10061829` encoded bytes (9.60 MiB), `10061829` compressed data bytes (9.60 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00016.parquet`: `23033` rows, `10076502` file bytes (9.61 MiB), `17634320` physical bytes (16.82 MiB), `10045535` encoded bytes (9.58 MiB), `10045535` compressed data bytes (9.58 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00017.parquet`: `23130` rows, `10075387` file bytes (9.61 MiB), `17880752` physical bytes (17.05 MiB), `10043903` encoded bytes (9.58 MiB), `10043903` compressed data bytes (9.58 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00018.parquet`: `23082` rows, `10075345` file bytes (9.61 MiB), `17789202` physical bytes (16.97 MiB), `10043961` encoded bytes (9.58 MiB), `10043961` compressed data bytes (9.58 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00019.parquet`: `22985` rows, `10086057` file bytes (9.62 MiB), `17556876` physical bytes (16.74 MiB), `10055467` encoded bytes (9.59 MiB), `10055467` compressed data bytes (9.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00020.parquet`: `22925` rows, `10084983` file bytes (9.62 MiB), `17589182` physical bytes (16.77 MiB), `10054164` encoded bytes (9.59 MiB), `10054164` compressed data bytes (9.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00021.parquet`: `22891` rows, `10083228` file bytes (9.62 MiB), `17796025` physical bytes (16.97 MiB), `10052027` encoded bytes (9.59 MiB), `10052027` compressed data bytes (9.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00022.parquet`: `20204` rows, `10168211` file bytes (9.70 MiB), `14205511` physical bytes (13.55 MiB), `10138118` encoded bytes (9.67 MiB), `10138118` compressed data bytes (9.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00023.parquet`: `20109` rows, `10166493` file bytes (9.70 MiB), `13747584` physical bytes (13.11 MiB), `10136665` encoded bytes (9.67 MiB), `10136665` compressed data bytes (9.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00024.parquet`: `19610` rows, `10184033` file bytes (9.71 MiB), `13599679` physical bytes (12.97 MiB), `10154421` encoded bytes (9.68 MiB), `10154421` compressed data bytes (9.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00025.parquet`: `18917` rows, `10211279` file bytes (9.74 MiB), `13496858` physical bytes (12.87 MiB), `10181499` encoded bytes (9.71 MiB), `10181499` compressed data bytes (9.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00026.parquet`: `20428` rows, `10163120` file bytes (9.69 MiB), `13192996` physical bytes (12.58 MiB), `10133418` encoded bytes (9.66 MiB), `10133418` compressed data bytes (9.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00027.parquet`: `20368` rows, `10162743` file bytes (9.69 MiB), `12910850` physical bytes (12.31 MiB), `10132822` encoded bytes (9.66 MiB), `10132822` compressed data bytes (9.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00028.parquet`: `20012` rows, `10181648` file bytes (9.71 MiB), `13112795` physical bytes (12.51 MiB), `10151585` encoded bytes (9.68 MiB), `10151585` compressed data bytes (9.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00029.parquet`: `20052` rows, `10177799` file bytes (9.71 MiB), `12979017` physical bytes (12.38 MiB), `10148101` encoded bytes (9.68 MiB), `10148101` compressed data bytes (9.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00030.parquet`: `19634` rows, `10177888` file bytes (9.71 MiB), `12843053` physical bytes (12.25 MiB), `10147998` encoded bytes (9.68 MiB), `10147998` compressed data bytes (9.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00031.parquet`: `20209` rows, `10192470` file bytes (9.72 MiB), `12876901` physical bytes (12.28 MiB), `10162342` encoded bytes (9.69 MiB), `10162342` compressed data bytes (9.69 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00032.parquet`: `20473` rows, `10164314` file bytes (9.69 MiB), `13009073` physical bytes (12.41 MiB), `10134158` encoded bytes (9.66 MiB), `10134158` compressed data bytes (9.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00033.parquet`: `19938` rows, `10168449` file bytes (9.70 MiB), `12946456` physical bytes (12.35 MiB), `10138732` encoded bytes (9.67 MiB), `10138732` compressed data bytes (9.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00034.parquet`: `20359` rows, `10171298` file bytes (9.70 MiB), `12955522` physical bytes (12.36 MiB), `10141330` encoded bytes (9.67 MiB), `10141330` compressed data bytes (9.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00035.parquet`: `20316` rows, `10171753` file bytes (9.70 MiB), `13017351` physical bytes (12.41 MiB), `10141440` encoded bytes (9.67 MiB), `10141440` compressed data bytes (9.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00036.parquet`: `20147` rows, `10167909` file bytes (9.70 MiB), `12991343` physical bytes (12.39 MiB), `10138225` encoded bytes (9.67 MiB), `10138225` compressed data bytes (9.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00037.parquet`: `20175` rows, `10182526` file bytes (9.71 MiB), `13001452` physical bytes (12.40 MiB), `10152756` encoded bytes (9.68 MiB), `10152756` compressed data bytes (9.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00038.parquet`: `20016` rows, `10178824` file bytes (9.71 MiB), `12942606` physical bytes (12.34 MiB), `10148654` encoded bytes (9.68 MiB), `10148654` compressed data bytes (9.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00039.parquet`: `19953` rows, `10184149` file bytes (9.71 MiB), `12900762` physical bytes (12.30 MiB), `10154551` encoded bytes (9.68 MiB), `10154551` compressed data bytes (9.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00040.parquet`: `20813` rows, `10151491` file bytes (9.68 MiB), `13085650` physical bytes (12.48 MiB), `10121843` encoded bytes (9.65 MiB), `10121843` compressed data bytes (9.65 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00041.parquet`: `20375` rows, `10175340` file bytes (9.70 MiB), `13013515` physical bytes (12.41 MiB), `10145467` encoded bytes (9.68 MiB), `10145467` compressed data bytes (9.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00042.parquet`: `20492` rows, `10163961` file bytes (9.69 MiB), `13018450` physical bytes (12.42 MiB), `10134000` encoded bytes (9.66 MiB), `10134000` compressed data bytes (9.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00043.parquet`: `20454` rows, `10157192` file bytes (9.69 MiB), `13154122` physical bytes (12.54 MiB), `10127196` encoded bytes (9.66 MiB), `10127196` compressed data bytes (9.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00044.parquet`: `20729` rows, `10153341` file bytes (9.68 MiB), `12940335` physical bytes (12.34 MiB), `10123736` encoded bytes (9.65 MiB), `10123736` compressed data bytes (9.65 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00045.parquet`: `20121` rows, `10183273` file bytes (9.71 MiB), `13006657` physical bytes (12.40 MiB), `10153456` encoded bytes (9.68 MiB), `10153456` compressed data bytes (9.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00046.parquet`: `11934` rows, `6072788` file bytes (5.79 MiB), `7755227` physical bytes (7.40 MiB), `6056172` encoded bytes (5.78 MiB), `6056172` compressed data bytes (5.78 MiB)
