# ClickBench Parquet Experiment

- Started: `2026-07-03T14:55:14-04:00`
- Write elapsed: `10.284s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `476656260` (454.57 MiB)
- Compressed column data bytes after codec compression: `476656260` (454.57 MiB)
- Parquet file bytes: `478083868` (455.94 MiB)
- Physical/encoded ratio: `1.495x`
- Encoded/compressed-data ratio: `1.000x`
- Physical/compressed-data ratio: `1.495x`
- Physical/parquet-file ratio: `1.490x`
- Files: `47`

## Settings

- Compression: `none`
- Int encoding: `plain`
- String encoding: `rle-dict`
- Date encoding: `plain`
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
- Rows read and compared: `1000000`
- Files read: `47`
- Elapsed: `6.709s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `8000000` (7.63 MiB) | `8007421` (7.64 MiB) | `8007421` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005915` (3.82 MiB) | `4005915` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `138409995` (132.00 MiB) | `29384410` (28.02 MiB) | `29384410` (28.02 MiB) | `4.710x` | `1.000x` | `4.710x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005916` (3.82 MiB) | `4005916` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `8000000` (7.63 MiB) | `8007417` (7.64 MiB) | `8007417` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005914` (3.82 MiB) | `4005914` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005914` (3.82 MiB) | `4005914` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005912` (3.82 MiB) | `4005912` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005913` (3.82 MiB) | `4005913` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `8000000` (7.63 MiB) | `8007420` (7.64 MiB) | `8007420` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005912` (3.82 MiB) | `4005912` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005915` (3.82 MiB) | `4005915` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005913` (3.82 MiB) | `4005913` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `88562192` (84.46 MiB) | `44330230` (42.28 MiB) | `44330230` (42.28 MiB) | `1.998x` | `1.000x` | `1.998x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `79583339` (75.90 MiB) | `34540103` (32.94 MiB) | `34540103` (32.94 MiB) | `2.304x` | `1.000x` | `2.304x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005914` (3.82 MiB) | `4005914` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005911` (3.82 MiB) | `4005911` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005917` (3.82 MiB) | `4005917` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005916` (3.82 MiB) | `4005916` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005914` (3.82 MiB) | `4005914` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005917` (3.82 MiB) | `4005917` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005915` (3.82 MiB) | `4005915` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005917` (3.82 MiB) | `4005917` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005913` (3.82 MiB) | `4005913` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005915` (3.82 MiB) | `4005915` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `3354477` (3.20 MiB) | `271264` (264.91 KiB) | `271264` (264.91 KiB) | `12.366x` | `1.000x` | `12.366x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005911` (3.82 MiB) | `4005911` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005915` (3.82 MiB) | `4005915` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005918` (3.82 MiB) | `4005918` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `3767530` (3.59 MiB) | `153966` (150.36 KiB) | `153966` (150.36 KiB) | `24.470x` | `1.000x` | `24.470x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005916` (3.82 MiB) | `4005916` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005910` (3.82 MiB) | `4005910` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005913` (3.82 MiB) | `4005913` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005910` (3.82 MiB) | `4005910` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `81583` (79.67 KiB) | `27823` (27.17 KiB) | `27823` (27.17 KiB) | `2.932x` | `1.000x` | `2.932x` | `81583` (79.67 KiB) |
| `Params` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `0` (0 B) | `6525` (6.37 KiB) | `6525` (6.37 KiB) | `0.000x` | `1.000x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005918` (3.82 MiB) | `4005918` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005912` (3.82 MiB) | `4005912` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005913` (3.82 MiB) | `4005913` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `3528017` (3.36 MiB) | `1646953` (1.57 MiB) | `1646953` (1.57 MiB) | `2.142x` | `1.000x` | `2.142x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005913` (3.82 MiB) | `4005913` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005919` (3.82 MiB) | `4005919` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005914` (3.82 MiB) | `4005914` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005914` (3.82 MiB) | `4005914` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005914` (3.82 MiB) | `4005914` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `8000000` (7.63 MiB) | `8007417` (7.64 MiB) | `8007417` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005910` (3.82 MiB) | `4005910` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005915` (3.82 MiB) | `4005915` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005913` (3.82 MiB) | `4005913` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005912` (3.82 MiB) | `4005912` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `13587860` (12.96 MiB) | `13214` (12.90 KiB) | `13214` (12.90 KiB) | `1028.293x` | `1.000x` | `1028.293x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005913` (3.82 MiB) | `4005913` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005912` (3.82 MiB) | `4005912` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005912` (3.82 MiB) | `4005912` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005915` (3.82 MiB) | `4005915` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `8000000` (7.63 MiB) | `8007419` (7.64 MiB) | `8007419` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `27797671` (26.51 MiB) | `21302107` (20.32 MiB) | `21302107` (20.32 MiB) | `1.305x` | `1.000x` | `1.305x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005913` (3.82 MiB) | `4005913` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005912` (3.82 MiB) | `4005912` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005912` (3.82 MiB) | `4005912` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005912` (3.82 MiB) | `4005912` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005910` (3.82 MiB) | `4005910` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005912` (3.82 MiB) | `4005912` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `1000000` (976.56 KiB) | `36215` (35.37 KiB) | `36215` (35.37 KiB) | `27.613x` | `1.000x` | `27.613x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `8000000` (7.63 MiB) | `8007422` (7.64 MiB) | `8007422` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005914` (3.82 MiB) | `4005914` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005910` (3.82 MiB) | `4005910` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005912` (3.82 MiB) | `4005912` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005916` (3.82 MiB) | `4005916` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005913` (3.82 MiB) | `4005913` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005916` (3.82 MiB) | `4005916` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005913` (3.82 MiB) | `4005913` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005911` (3.82 MiB) | `4005911` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005914` (3.82 MiB) | `4005914` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `2001192` (1.91 MiB) | `45184` (44.12 KiB) | `45184` (44.12 KiB) | `44.290x` | `1.000x` | `44.290x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `3325142` (3.17 MiB) | `139624` (136.35 KiB) | `139624` (136.35 KiB) | `23.815x` | `1.000x` | `23.815x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `0` (0 B) | `6525` (6.37 KiB) | `6525` (6.37 KiB) | `0.000x` | `1.000x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `0` (0 B) | `6525` (6.37 KiB) | `6525` (6.37 KiB) | `0.000x` | `1.000x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005912` (3.82 MiB) | `4005912` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005912` (3.82 MiB) | `4005912` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005916` (3.82 MiB) | `4005916` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005918` (3.82 MiB) | `4005918` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005912` (3.82 MiB) | `4005912` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005913` (3.82 MiB) | `4005913` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005908` (3.82 MiB) | `4005908` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005917` (3.82 MiB) | `4005917` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `1024` (1.00 KiB) | `8166` (7.97 KiB) | `8166` (7.97 KiB) | `0.125x` | `1.000x` | `0.125x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `8000000` (7.63 MiB) | `8007417` (7.64 MiB) | `8007417` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `0` (0 B) | `6525` (6.37 KiB) | `6525` (6.37 KiB) | `0.000x` | `1.000x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `3000000` (2.86 MiB) | `7841` (7.66 KiB) | `7841` (7.66 KiB) | `382.604x` | `1.000x` | `382.604x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005912` (3.82 MiB) | `4005912` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `58030` (56.67 KiB) | `22533` (22.00 KiB) | `22533` (22.00 KiB) | `2.575x` | `1.000x` | `2.575x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `22051` (21.53 KiB) | `19386` (18.93 KiB) | `19386` (18.93 KiB) | `1.137x` | `1.000x` | `1.137x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `25445` (24.85 KiB) | `25390` (24.79 KiB) | `25390` (24.79 KiB) | `1.002x` | `1.000x` | `1.002x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `48191` (47.06 KiB) | `17852` (17.43 KiB) | `17852` (17.43 KiB) | `2.699x` | `1.000x` | `2.699x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `49433` (48.27 KiB) | `26185` (25.57 KiB) | `26185` (25.57 KiB) | `1.888x` | `1.000x` | `1.888x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `16873` (16.48 KiB) | `19403` (18.95 KiB) | `19403` (18.95 KiB) | `0.870x` | `1.000x` | `0.870x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `91870` (89.72 KiB) | `42313` (41.32 KiB) | `42313` (41.32 KiB) | `2.171x` | `1.000x` | `2.171x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `13001` (12.70 KiB) | `17641` (17.23 KiB) | `17641` (17.23 KiB) | `0.737x` | `1.000x` | `0.737x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `28101` (27.44 KiB) | `19813` (19.35 KiB) | `19813` (19.35 KiB) | `1.418x` | `1.000x` | `1.418x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:94, DICTIONARY_PAGE/PLAIN:94` | `1000000` | `45607` (44.54 KiB) | `43650` (42.63 KiB) | `43650` (42.63 KiB) | `1.045x` | `1.000x` | `1.045x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005914` (3.82 MiB) | `4005914` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `8000000` (7.63 MiB) | `8007419` (7.64 MiB) | `8007419` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `8000000` (7.63 MiB) | `8007420` (7.64 MiB) | `8007420` (7.64 MiB) | `0.999x` | `1.000x` | `0.999x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:94` | `1000000` | `4000000` (3.81 MiB) | `4005913` (3.82 MiB) | `4005913` (3.82 MiB) | `0.999x` | `1.000x` | `0.999x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00000.parquet`: `22458` rows, `10174837` file bytes (9.70 MiB), `17432285` physical bytes (16.62 MiB), `10143807` encoded bytes (9.67 MiB), `10143807` compressed data bytes (9.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00001.parquet`: `22463` rows, `10161464` file bytes (9.69 MiB), `17411291` physical bytes (16.60 MiB), `10130604` encoded bytes (9.66 MiB), `10130604` compressed data bytes (9.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00002.parquet`: `22499` rows, `10163099` file bytes (9.69 MiB), `17264396` physical bytes (16.46 MiB), `10131831` encoded bytes (9.66 MiB), `10131831` compressed data bytes (9.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00003.parquet`: `22365` rows, `10169079` file bytes (9.70 MiB), `17479114` physical bytes (16.67 MiB), `10138022` encoded bytes (9.67 MiB), `10138022` compressed data bytes (9.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00004.parquet`: `22498` rows, `10166232` file bytes (9.70 MiB), `17608364` physical bytes (16.79 MiB), `10135061` encoded bytes (9.67 MiB), `10135061` compressed data bytes (9.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00005.parquet`: `22483` rows, `10165729` file bytes (9.69 MiB), `17360144` physical bytes (16.56 MiB), `10134266` encoded bytes (9.66 MiB), `10134266` compressed data bytes (9.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00006.parquet`: `22489` rows, `10164635` file bytes (9.69 MiB), `17454020` physical bytes (16.65 MiB), `10133442` encoded bytes (9.66 MiB), `10133442` compressed data bytes (9.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00007.parquet`: `22580` rows, `10164392` file bytes (9.69 MiB), `17544248` physical bytes (16.73 MiB), `10132319` encoded bytes (9.66 MiB), `10132319` compressed data bytes (9.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00008.parquet`: `22655` rows, `10161714` file bytes (9.69 MiB), `17438478` physical bytes (16.63 MiB), `10130260` encoded bytes (9.66 MiB), `10130260` compressed data bytes (9.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00009.parquet`: `22608` rows, `10167855` file bytes (9.70 MiB), `17323492` physical bytes (16.52 MiB), `10136829` encoded bytes (9.67 MiB), `10136829` compressed data bytes (9.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00010.parquet`: `22594` rows, `10159186` file bytes (9.69 MiB), `17469853` physical bytes (16.66 MiB), `10127989` encoded bytes (9.66 MiB), `10127989` compressed data bytes (9.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00011.parquet`: `22507` rows, `10170968` file bytes (9.70 MiB), `17505099` physical bytes (16.69 MiB), `10139607` encoded bytes (9.67 MiB), `10139607` compressed data bytes (9.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00012.parquet`: `22574` rows, `10166692` file bytes (9.70 MiB), `17432317` physical bytes (16.62 MiB), `10135705` encoded bytes (9.67 MiB), `10135705` compressed data bytes (9.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00013.parquet`: `22547` rows, `10165165` file bytes (9.69 MiB), `17578156` physical bytes (16.76 MiB), `10134234` encoded bytes (9.66 MiB), `10134234` compressed data bytes (9.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00014.parquet`: `22689` rows, `10158914` file bytes (9.69 MiB), `17382112` physical bytes (16.58 MiB), `10128215` encoded bytes (9.66 MiB), `10128215` compressed data bytes (9.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00015.parquet`: `22470` rows, `10170920` file bytes (9.70 MiB), `17362649` physical bytes (16.56 MiB), `10140388` encoded bytes (9.67 MiB), `10140388` compressed data bytes (9.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00016.parquet`: `22572` rows, `10165484` file bytes (9.69 MiB), `17363877` physical bytes (16.56 MiB), `10134718` encoded bytes (9.67 MiB), `10134718` compressed data bytes (9.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00017.parquet`: `22754` rows, `10159859` file bytes (9.69 MiB), `17595183` physical bytes (16.78 MiB), `10129454` encoded bytes (9.66 MiB), `10129454` compressed data bytes (9.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00018.parquet`: `22718` rows, `10165349` file bytes (9.69 MiB), `17419964` physical bytes (16.61 MiB), `10134010` encoded bytes (9.66 MiB), `10134010` compressed data bytes (9.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00019.parquet`: `22727` rows, `10162952` file bytes (9.69 MiB), `17455389` physical bytes (16.65 MiB), `10132054` encoded bytes (9.66 MiB), `10132054` compressed data bytes (9.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00020.parquet`: `22538` rows, `10170107` file bytes (9.70 MiB), `17300528` physical bytes (16.50 MiB), `10138984` encoded bytes (9.67 MiB), `10138984` compressed data bytes (9.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00021.parquet`: `22481` rows, `10167595` file bytes (9.70 MiB), `17394532` physical bytes (16.59 MiB), `10136787` encoded bytes (9.67 MiB), `10136787` compressed data bytes (9.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00022.parquet`: `21074` rows, `10210508` file bytes (9.74 MiB), `15491273` physical bytes (14.77 MiB), `10180435` encoded bytes (9.71 MiB), `10180435` compressed data bytes (9.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00023.parquet`: `20016` rows, `10228324` file bytes (9.75 MiB), `13717867` physical bytes (13.08 MiB), `10199035` encoded bytes (9.73 MiB), `10199035` compressed data bytes (9.73 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00024.parquet`: `19871` rows, `10240285` file bytes (9.77 MiB), `13628159` physical bytes (13.00 MiB), `10210362` encoded bytes (9.74 MiB), `10210362` compressed data bytes (9.74 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00025.parquet`: `19064` rows, `10263938` file bytes (9.79 MiB), `13536334` physical bytes (12.91 MiB), `10234611` encoded bytes (9.76 MiB), `10234611` compressed data bytes (9.76 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00026.parquet`: `19799` rows, `10221358` file bytes (9.75 MiB), `13392321` physical bytes (12.77 MiB), `10191798` encoded bytes (9.72 MiB), `10191798` compressed data bytes (9.72 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00027.parquet`: `20457` rows, `10235490` file bytes (9.76 MiB), `12989678` physical bytes (12.39 MiB), `10205765` encoded bytes (9.73 MiB), `10205765` compressed data bytes (9.73 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00028.parquet`: `20361` rows, `10241341` file bytes (9.77 MiB), `13135024` physical bytes (12.53 MiB), `10211628` encoded bytes (9.74 MiB), `10211628` compressed data bytes (9.74 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00029.parquet`: `20346` rows, `10236833` file bytes (9.76 MiB), `13115885` physical bytes (12.51 MiB), `10206733` encoded bytes (9.73 MiB), `10206733` compressed data bytes (9.73 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00030.parquet`: `19312` rows, `10277660` file bytes (9.80 MiB), `12891503` physical bytes (12.29 MiB), `10248029` encoded bytes (9.77 MiB), `10248029` compressed data bytes (9.77 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00031.parquet`: `20814` rows, `10228634` file bytes (9.75 MiB), `13053162` physical bytes (12.45 MiB), `10198648` encoded bytes (9.73 MiB), `10198648` compressed data bytes (9.73 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00032.parquet`: `20238` rows, `10248513` file bytes (9.77 MiB), `13008878` physical bytes (12.41 MiB), `10218575` encoded bytes (9.75 MiB), `10218575` compressed data bytes (9.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00033.parquet`: `20106` rows, `10239682` file bytes (9.77 MiB), `12940414` physical bytes (12.34 MiB), `10209831` encoded bytes (9.74 MiB), `10209831` compressed data bytes (9.74 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00034.parquet`: `20402` rows, `10221582` file bytes (9.75 MiB), `13047848` physical bytes (12.44 MiB), `10191874` encoded bytes (9.72 MiB), `10191874` compressed data bytes (9.72 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00035.parquet`: `20316` rows, `10235476` file bytes (9.76 MiB), `13080332` physical bytes (12.47 MiB), `10205360` encoded bytes (9.73 MiB), `10205360` compressed data bytes (9.73 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00036.parquet`: `20227` rows, `10234308` file bytes (9.76 MiB), `13005804` physical bytes (12.40 MiB), `10204723` encoded bytes (9.73 MiB), `10204723` compressed data bytes (9.73 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00037.parquet`: `20635` rows, `10212332` file bytes (9.74 MiB), `13042265` physical bytes (12.44 MiB), `10182737` encoded bytes (9.71 MiB), `10182737` compressed data bytes (9.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00038.parquet`: `19868` rows, `10254357` file bytes (9.78 MiB), `13008264` physical bytes (12.41 MiB), `10224542` encoded bytes (9.75 MiB), `10224542` compressed data bytes (9.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00039.parquet`: `20179` rows, `10233769` file bytes (9.76 MiB), `13005488` physical bytes (12.40 MiB), `10203949` encoded bytes (9.73 MiB), `10203949` compressed data bytes (9.73 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00040.parquet`: `20436` rows, `10218708` file bytes (9.75 MiB), `13044881` physical bytes (12.44 MiB), `10189103` encoded bytes (9.72 MiB), `10189103` compressed data bytes (9.72 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00041.parquet`: `20711` rows, `10221935` file bytes (9.75 MiB), `13103831` physical bytes (12.50 MiB), `10192155` encoded bytes (9.72 MiB), `10192155` compressed data bytes (9.72 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00042.parquet`: `20527` rows, `10208143` file bytes (9.74 MiB), `13057903` physical bytes (12.45 MiB), `10178269` encoded bytes (9.71 MiB), `10178269` compressed data bytes (9.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00043.parquet`: `20271` rows, `10228210` file bytes (9.75 MiB), `13130428` physical bytes (12.52 MiB), `10198079` encoded bytes (9.73 MiB), `10198079` compressed data bytes (9.73 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00044.parquet`: `20877` rows, `10210663` file bytes (9.74 MiB), `13030996` physical bytes (12.43 MiB), `10180959` encoded bytes (9.71 MiB), `10180959` compressed data bytes (9.71 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00045.parquet`: `20429` rows, `10218535` file bytes (9.75 MiB), `13062547` physical bytes (12.46 MiB), `10188789` encoded bytes (9.72 MiB), `10188789` compressed data bytes (9.72 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-none-int-plain-str-rle-dict-date-plain-ts-plain/part-00046.parquet`: `17395` rows, `8871057` file bytes (8.46 MiB), `11302048` physical bytes (10.78 MiB), `8841685` encoded bytes (8.43 MiB), `8841685` compressed data bytes (8.43 MiB)
