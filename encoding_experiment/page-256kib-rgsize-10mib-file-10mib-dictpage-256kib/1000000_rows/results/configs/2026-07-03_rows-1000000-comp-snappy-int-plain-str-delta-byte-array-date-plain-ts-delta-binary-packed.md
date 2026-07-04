# ClickBench Parquet Experiment

- Started: `2026-07-03T23:29:25-04:00`
- Write elapsed: `11.572s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `505152789` (481.75 MiB)
- Compressed column data bytes after codec compression: `117124381` (111.70 MiB)
- Parquet file bytes: `118053475` (112.58 MiB)
- Physical/encoded ratio: `1.410x`
- Encoded/compressed-data ratio: `4.313x`
- Physical/compressed-data ratio: `6.082x`
- Physical/parquet-file ratio: `6.035x`
- Files: `30`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `delta-byte-array`
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
- Files read: `30`
- Elapsed: `7.077s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004717` (7.63 MiB) | `8005130` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003711` (3.82 MiB) | `261946` (255.81 KiB) | `0.999x` | `15.284x` | `15.270x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:513` | `1000000` | `138409995` (132.00 MiB) | `64477801` (61.49 MiB) | `17105312` (16.31 MiB) | `2.147x` | `3.769x` | `8.092x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `204227` (199.44 KiB) | `0.999x` | `19.604x` | `19.586x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:59` | `1000000` | `8000000` (7.63 MiB) | `3370560` (3.21 MiB) | `3038726` (2.90 MiB) | `2.373x` | `1.109x` | `2.633x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003711` (3.82 MiB) | `204226` (199.44 KiB) | `0.999x` | `19.604x` | `19.586x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `204252` (199.46 KiB) | `0.999x` | `19.602x` | `19.584x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003772` (3.82 MiB) | `719059` (702.21 KiB) | `0.999x` | `5.568x` | `5.563x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `396421` (387.13 KiB) | `0.999x` | `10.100x` | `10.090x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004717` (7.63 MiB) | `1084730` (1.03 MiB) | `0.999x` | `7.379x` | `7.375x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `204060` (199.28 KiB) | `0.999x` | `19.620x` | `19.602x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `318013` (310.56 KiB) | `0.999x` | `12.590x` | `12.578x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `342090` (334.07 KiB) | `0.999x` | `11.704x` | `11.693x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:339` | `1000000` | `88562192` (84.46 MiB) | `40456964` (38.58 MiB) | `18047077` (17.21 MiB) | `2.189x` | `2.242x` | `4.907x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:313` | `1000000` | `79583339` (75.90 MiB) | `38988769` (37.18 MiB) | `17315959` (16.51 MiB) | `2.041x` | `2.252x` | `4.596x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003740` (3.82 MiB) | `491863` (480.33 KiB) | `0.999x` | `8.140x` | `8.132x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003746` (3.82 MiB) | `508852` (496.93 KiB) | `0.999x` | `7.868x` | `7.861x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003735` (3.82 MiB) | `458016` (447.28 KiB) | `0.999x` | `8.741x` | `8.733x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `271779` (265.41 KiB) | `0.999x` | `14.732x` | `14.718x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `232146` (226.71 KiB) | `0.999x` | `17.247x` | `17.231x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `373266` (364.52 KiB) | `0.999x` | `10.726x` | `10.716x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003717` (3.82 MiB) | `372941` (364.20 KiB) | `0.999x` | `10.736x` | `10.726x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003710` (3.82 MiB) | `282885` (276.25 KiB) | `0.999x` | `14.153x` | `14.140x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003711` (3.82 MiB) | `245188` (239.44 KiB) | `0.999x` | `16.329x` | `16.314x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003717` (3.82 MiB) | `328349` (320.65 KiB) | `0.999x` | `12.193x` | `12.182x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `3354477` (3.20 MiB) | `1042321` (1017.89 KiB) | `394668` (385.42 KiB) | `3.218x` | `2.641x` | `8.499x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003711` (3.82 MiB) | `219057` (213.92 KiB) | `0.999x` | `18.277x` | `18.260x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `217211` (212.12 KiB) | `0.999x` | `18.432x` | `18.415x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `364841` (356.29 KiB) | `0.999x` | `10.974x` | `10.964x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `3767530` (3.59 MiB) | `859642` (839.49 KiB) | `236469` (230.93 KiB) | `4.383x` | `3.635x` | `15.932x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `204824` (200.02 KiB) | `0.999x` | `19.547x` | `19.529x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `204978` (200.17 KiB) | `0.999x` | `19.532x` | `19.514x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003711` (3.82 MiB) | `218332` (213.21 KiB) | `0.999x` | `18.338x` | `18.321x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003709` (3.82 MiB) | `215318` (210.27 KiB) | `0.999x` | `18.594x` | `18.577x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `81583` (79.67 KiB) | `282717` (276.09 KiB) | `56951` (55.62 KiB) | `0.289x` | `4.964x` | `1.433x` | `81583` (79.67 KiB) |
| `Params` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `0` (0 B) | `81417` (79.51 KiB) | `6904` (6.74 KiB) | `0.000x` | `11.793x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003756` (3.82 MiB) | `560369` (547.24 KiB) | `0.999x` | `7.145x` | `7.138x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003756` (3.82 MiB) | `526787` (514.44 KiB) | `0.999x` | `7.600x` | `7.593x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `292991` (286.12 KiB) | `0.999x` | `13.665x` | `13.652x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `3528017` (3.36 MiB) | `2961049` (2.82 MiB) | `994070` (970.77 KiB) | `1.191x` | `2.979x` | `3.549x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `220229` (215.07 KiB) | `0.999x` | `18.180x` | `18.163x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003749` (3.82 MiB) | `505682` (493.83 KiB) | `0.999x` | `7.918x` | `7.910x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003747` (3.82 MiB) | `516116` (504.02 KiB) | `0.999x` | `7.757x` | `7.750x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003759` (3.82 MiB) | `551444` (538.52 KiB) | `0.999x` | `7.260x` | `7.254x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `285629` (278.93 KiB) | `0.999x` | `14.017x` | `14.004x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:59` | `1000000` | `8000000` (7.63 MiB) | `3492513` (3.33 MiB) | `3052545` (2.91 MiB) | `2.291x` | `1.144x` | `2.621x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `287130` (280.40 KiB) | `0.999x` | `13.944x` | `13.931x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `269638` (263.32 KiB) | `0.999x` | `14.848x` | `14.835x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `304718` (297.58 KiB) | `0.999x` | `13.139x` | `13.127x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003711` (3.82 MiB) | `204099` (199.32 KiB) | `0.999x` | `19.617x` | `19.598x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:89` | `1000000` | `13587860` (12.96 MiB) | `142193` (138.86 KiB) | `26906` (26.28 KiB) | `95.559x` | `5.285x` | `505.012x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `205059` (200.25 KiB) | `0.999x` | `19.525x` | `19.507x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `256364` (250.36 KiB) | `0.999x` | `15.617x` | `15.603x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003711` (3.82 MiB) | `206523` (201.68 KiB) | `0.999x` | `19.386x` | `19.368x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003717` (3.82 MiB) | `235282` (229.77 KiB) | `0.999x` | `17.017x` | `17.001x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004717` (7.63 MiB) | `1151442` (1.10 MiB) | `0.999x` | `6.952x` | `6.948x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:145` | `1000000` | `27797671` (26.51 MiB) | `21056832` (20.08 MiB) | `6747020` (6.43 MiB) | `1.320x` | `3.121x` | `4.120x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003768` (3.82 MiB) | `3689362` (3.52 MiB) | `0.999x` | `1.085x` | `1.084x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `204060` (199.28 KiB) | `0.999x` | `19.620x` | `19.602x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `204060` (199.28 KiB) | `0.999x` | `19.620x` | `19.602x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `204060` (199.28 KiB) | `0.999x` | `19.620x` | `19.602x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003719` (3.82 MiB) | `295579` (288.65 KiB) | `0.999x` | `13.545x` | `13.533x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `204060` (199.28 KiB) | `0.999x` | `19.620x` | `19.602x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `1000000` (976.56 KiB) | `207248` (202.39 KiB) | `74548` (72.80 KiB) | `4.825x` | `2.780x` | `13.414x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:59` | `1000000` | `8000000` (7.63 MiB) | `3385933` (3.23 MiB) | `3057126` (2.92 MiB) | `2.363x` | `1.108x` | `2.617x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003711` (3.82 MiB) | `333858` (326.03 KiB) | `0.999x` | `11.992x` | `11.981x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `302471` (295.38 KiB) | `0.999x` | `13.237x` | `13.224x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `318144` (310.69 KiB) | `0.999x` | `12.585x` | `12.573x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `385397` (376.36 KiB) | `0.999x` | `10.389x` | `10.379x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003719` (3.82 MiB) | `381560` (372.62 KiB) | `0.999x` | `10.493x` | `10.483x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003769` (3.82 MiB) | `706815` (690.25 KiB) | `0.999x` | `5.665x` | `5.659x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003718` (3.82 MiB) | `295142` (288.22 KiB) | `0.999x` | `13.565x` | `13.553x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `204064` (199.28 KiB) | `0.999x` | `19.620x` | `19.602x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003717` (3.82 MiB) | `258233` (252.18 KiB) | `0.999x` | `15.504x` | `15.490x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `2001192` (1.91 MiB) | `332657` (324.86 KiB) | `89633` (87.53 KiB) | `6.016x` | `3.711x` | `22.327x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `3325142` (3.17 MiB) | `961166` (938.64 KiB) | `254074` (248.12 KiB) | `3.459x` | `3.783x` | `13.087x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `0` (0 B) | `81417` (79.51 KiB) | `6904` (6.74 KiB) | `0.000x` | `11.793x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `0` (0 B) | `81417` (79.51 KiB) | `6904` (6.74 KiB) | `0.000x` | `11.793x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `204060` (199.28 KiB) | `0.999x` | `19.620x` | `19.602x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003719` (3.82 MiB) | `269082` (262.78 KiB) | `0.999x` | `14.879x` | `14.865x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `326912` (319.25 KiB) | `0.999x` | `12.247x` | `12.236x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003750` (3.82 MiB) | `564264` (551.04 KiB) | `0.999x` | `7.096x` | `7.089x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003773` (3.82 MiB) | `1722707` (1.64 MiB) | `0.999x` | `2.324x` | `2.322x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `1281142` (1.22 MiB) | `0.999x` | `3.125x` | `3.122x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003767` (3.82 MiB) | `808637` (789.68 KiB) | `0.999x` | `4.951x` | `4.947x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `204453` (199.66 KiB) | `0.999x` | `19.583x` | `19.564x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `1024` (1.00 KiB) | `86744` (84.71 KiB) | `9620` (9.39 KiB) | `0.012x` | `9.017x` | `0.106x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004657` (7.63 MiB) | `405011` (395.52 KiB) | `0.999x` | `19.764x` | `19.753x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `0` (0 B) | `81417` (79.51 KiB) | `6904` (6.74 KiB) | `0.000x` | `11.793x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `3000000` (2.86 MiB) | `86334` (84.31 KiB) | `10318` (10.08 KiB) | `34.749x` | `8.367x` | `290.754x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `204060` (199.28 KiB) | `0.999x` | `19.620x` | `19.602x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `58030` (56.67 KiB) | `301901` (294.83 KiB) | `48309` (47.18 KiB) | `0.192x` | `6.249x` | `1.201x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `22051` (21.53 KiB) | `197672` (193.04 KiB) | `38832` (37.92 KiB) | `0.112x` | `5.090x` | `0.568x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `25445` (24.85 KiB) | `209315` (204.41 KiB) | `44719` (43.67 KiB) | `0.122x` | `4.681x` | `0.569x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `48191` (47.06 KiB) | `231729` (226.30 KiB) | `34058` (33.26 KiB) | `0.208x` | `6.804x` | `1.415x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `49433` (48.27 KiB) | `268493` (262.20 KiB) | `52286` (51.06 KiB) | `0.184x` | `5.135x` | `0.945x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `16873` (16.48 KiB) | `214024` (209.01 KiB) | `38785` (37.88 KiB) | `0.079x` | `5.518x` | `0.435x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `91870` (89.72 KiB) | `328728` (321.02 KiB) | `68263` (66.66 KiB) | `0.279x` | `4.816x` | `1.346x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `13001` (12.70 KiB) | `150182` (146.66 KiB) | `27888` (27.23 KiB) | `0.087x` | `5.385x` | `0.466x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `28101` (27.44 KiB) | `190438` (185.97 KiB) | `31021` (30.29 KiB) | `0.148x` | `6.139x` | `0.906x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `45607` (44.54 KiB) | `261764` (255.63 KiB) | `50498` (49.31 KiB) | `0.174x` | `5.184x` | `0.903x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `213485` (208.48 KiB) | `0.999x` | `18.754x` | `18.737x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004718` (7.63 MiB) | `3638786` (3.47 MiB) | `0.999x` | `2.200x` | `2.199x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004715` (7.63 MiB) | `4383602` (4.18 MiB) | `0.999x` | `1.826x` | `1.825x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `204516` (199.72 KiB) | `0.999x` | `19.577x` | `19.558x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00000.parquet`: `34195` rows, `3743896` file bytes (3.57 MiB), `26565625` physical bytes (25.33 MiB), `16725528` encoded bytes (15.95 MiB), `3711983` compressed data bytes (3.54 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00001.parquet`: `34522` rows, `3705830` file bytes (3.53 MiB), `26556151` physical bytes (25.33 MiB), `16644510` encoded bytes (15.87 MiB), `3673523` compressed data bytes (3.50 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00002.parquet`: `34011` rows, `3720635` file bytes (3.55 MiB), `26612338` physical bytes (25.38 MiB), `16652179` encoded bytes (15.88 MiB), `3687534` compressed data bytes (3.52 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00003.parquet`: `34320` rows, `3665812` file bytes (3.50 MiB), `26562508` physical bytes (25.33 MiB), `16592370` encoded bytes (15.82 MiB), `3633931` compressed data bytes (3.47 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00004.parquet`: `34160` rows, `3695857` file bytes (3.52 MiB), `26589163` physical bytes (25.36 MiB), `16614283` encoded bytes (15.84 MiB), `3662741` compressed data bytes (3.49 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00005.parquet`: `34767` rows, `3699458` file bytes (3.53 MiB), `26808451` physical bytes (25.57 MiB), `16775292` encoded bytes (16.00 MiB), `3667301` compressed data bytes (3.50 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00006.parquet`: `34592` rows, `3705778` file bytes (3.53 MiB), `26564772` physical bytes (25.33 MiB), `16672182` encoded bytes (15.90 MiB), `3673279` compressed data bytes (3.50 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00007.parquet`: `34263` rows, `3673644` file bytes (3.50 MiB), `26641159` physical bytes (25.41 MiB), `16599642` encoded bytes (15.83 MiB), `3641036` compressed data bytes (3.47 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00008.parquet`: `34341` rows, `3668814` file bytes (3.50 MiB), `26614814` physical bytes (25.38 MiB), `16670637` encoded bytes (15.90 MiB), `3636202` compressed data bytes (3.47 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00009.parquet`: `34170` rows, `3621892` file bytes (3.45 MiB), `26209736` physical bytes (25.00 MiB), `16474590` encoded bytes (15.71 MiB), `3589481` compressed data bytes (3.42 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00010.parquet`: `34081` rows, `3649128` file bytes (3.48 MiB), `26415579` physical bytes (25.19 MiB), `16514217` encoded bytes (15.75 MiB), `3617153` compressed data bytes (3.45 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00011.parquet`: `34621` rows, `3624812` file bytes (3.46 MiB), `26642262` physical bytes (25.41 MiB), `16630627` encoded bytes (15.86 MiB), `3592340` compressed data bytes (3.43 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00012.parquet`: `34666` rows, `3653399` file bytes (3.48 MiB), `26609271` physical bytes (25.38 MiB), `16692097` encoded bytes (15.92 MiB), `3621426` compressed data bytes (3.45 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00013.parquet`: `34354` rows, `3674744` file bytes (3.50 MiB), `26349013` physical bytes (25.13 MiB), `16587046` encoded bytes (15.82 MiB), `3642534` compressed data bytes (3.47 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00014.parquet`: `34934` rows, `3879413` file bytes (3.70 MiB), `26423181` physical bytes (25.20 MiB), `16968678` encoded bytes (16.18 MiB), `3847139` compressed data bytes (3.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00015.parquet`: `34414` rows, `4160949` file bytes (3.97 MiB), `23505776` physical bytes (22.42 MiB), `16969396` encoded bytes (16.18 MiB), `4130638` compressed data bytes (3.94 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00016.parquet`: `33598` rows, `4207641` file bytes (4.01 MiB), `23807228` physical bytes (22.70 MiB), `16650296` encoded bytes (15.88 MiB), `4177422` compressed data bytes (3.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00017.parquet`: `33181` rows, `4259938` file bytes (4.06 MiB), `21257824` physical bytes (20.27 MiB), `17429911` encoded bytes (16.62 MiB), `4229650` compressed data bytes (4.03 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00018.parquet`: `33481` rows, `4385076` file bytes (4.18 MiB), `21637945` physical bytes (20.64 MiB), `18015778` encoded bytes (17.18 MiB), `4354353` compressed data bytes (4.15 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00019.parquet`: `33072` rows, `4465659` file bytes (4.26 MiB), `21563774` physical bytes (20.56 MiB), `17911657` encoded bytes (17.08 MiB), `4435305` compressed data bytes (4.23 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00020.parquet`: `33166` rows, `4300647` file bytes (4.10 MiB), `21132498` physical bytes (20.15 MiB), `17653266` encoded bytes (16.84 MiB), `4269865` compressed data bytes (4.07 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00021.parquet`: `33027` rows, `4382663` file bytes (4.18 MiB), `21231180` physical bytes (20.25 MiB), `17656933` encoded bytes (16.84 MiB), `4351798` compressed data bytes (4.15 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00022.parquet`: `33468` rows, `4355533` file bytes (4.15 MiB), `21442654` physical bytes (20.45 MiB), `17853418` encoded bytes (17.03 MiB), `4324906` compressed data bytes (4.12 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00023.parquet`: `33541` rows, `4340195` file bytes (4.14 MiB), `21418977` physical bytes (20.43 MiB), `17789312` encoded bytes (16.97 MiB), `4309810` compressed data bytes (4.11 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00024.parquet`: `33137` rows, `4449005` file bytes (4.24 MiB), `21588949` physical bytes (20.59 MiB), `18047273` encoded bytes (17.21 MiB), `4418302` compressed data bytes (4.21 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00025.parquet`: `33193` rows, `4281234` file bytes (4.08 MiB), `21113358` physical bytes (20.14 MiB), `17494968` encoded bytes (16.68 MiB), `4251133` compressed data bytes (4.05 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00026.parquet`: `33269` rows, `4347382` file bytes (4.15 MiB), `21274554` physical bytes (20.29 MiB), `17712744` encoded bytes (16.89 MiB), `4316878` compressed data bytes (4.12 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00027.parquet`: `33536` rows, `4269843` file bytes (4.07 MiB), `21312760` physical bytes (20.33 MiB), `17619598` encoded bytes (16.80 MiB), `4239267` compressed data bytes (4.04 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00028.parquet`: `33346` rows, `4275984` file bytes (4.08 MiB), `21150532` physical bytes (20.17 MiB), `17610672` encoded bytes (16.79 MiB), `4245770` compressed data bytes (4.05 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed/part-00029.parquet`: `16574` rows, `2188614` file bytes (2.09 MiB), `10796592` physical bytes (10.30 MiB), `8923689` encoded bytes (8.51 MiB), `2171681` compressed data bytes (2.07 MiB)
