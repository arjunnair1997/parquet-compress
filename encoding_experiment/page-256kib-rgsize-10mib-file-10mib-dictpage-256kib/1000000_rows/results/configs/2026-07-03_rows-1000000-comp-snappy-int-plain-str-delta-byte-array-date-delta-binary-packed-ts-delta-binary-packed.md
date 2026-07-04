# ClickBench Parquet Experiment

- Started: `2026-07-03T23:30:39-04:00`
- Write elapsed: `11.512s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `501194834` (477.98 MiB)
- Compressed column data bytes after codec compression: `116950116` (111.53 MiB)
- Parquet file bytes: `117879284` (112.42 MiB)
- Physical/encoded ratio: `1.421x`
- Encoded/compressed-data ratio: `4.286x`
- Physical/compressed-data ratio: `6.091x`
- Physical/parquet-file ratio: `6.043x`
- Files: `30`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `delta-byte-array`
- Date encoding: `delta-binary-packed`
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
- Elapsed: `7.069s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004715` (7.63 MiB) | `8005128` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `261945` (255.81 KiB) | `0.999x` | `15.285x` | `15.270x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:516` | `1000000` | `138409995` (132.00 MiB) | `64473794` (61.49 MiB) | `17115419` (16.32 MiB) | `2.147x` | `3.767x` | `8.087x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `204241` (199.45 KiB) | `0.999x` | `19.603x` | `19.585x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:59` | `1000000` | `8000000` (7.63 MiB) | `3370830` (3.21 MiB) | `3039665` (2.90 MiB) | `2.373x` | `1.109x` | `2.632x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:59` | `1000000` | `4000000` (3.81 MiB) | `51937` (50.72 KiB) | `7914` (7.73 KiB) | `77.016x` | `6.563x` | `505.433x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `204263` (199.48 KiB) | `0.999x` | `19.601x` | `19.583x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `719097` (702.24 KiB) | `0.999x` | `5.568x` | `5.563x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `396306` (387.02 KiB) | `0.999x` | `10.103x` | `10.093x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004715` (7.63 MiB) | `1084714` (1.03 MiB) | `0.999x` | `7.380x` | `7.375x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `204076` (199.29 KiB) | `0.999x` | `19.619x` | `19.601x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003711` (3.82 MiB) | `317984` (310.53 KiB) | `0.999x` | `12.591x` | `12.579x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `342015` (334.00 KiB) | `0.999x` | `11.706x` | `11.695x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:339` | `1000000` | `88562192` (84.46 MiB) | `40459829` (38.59 MiB) | `18056198` (17.22 MiB) | `2.189x` | `2.241x` | `4.905x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:310` | `1000000` | `79583339` (75.90 MiB) | `38992810` (37.19 MiB) | `17321007` (16.52 MiB) | `2.041x` | `2.251x` | `4.595x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003743` (3.82 MiB) | `491817` (480.29 KiB) | `0.999x` | `8.141x` | `8.133x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003746` (3.82 MiB) | `508782` (496.86 KiB) | `0.999x` | `7.869x` | `7.862x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003737` (3.82 MiB) | `457896` (447.16 KiB) | `0.999x` | `8.744x` | `8.736x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `271848` (265.48 KiB) | `0.999x` | `14.728x` | `14.714x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `232147` (226.71 KiB) | `0.999x` | `17.246x` | `17.230x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `373199` (364.45 KiB) | `0.999x` | `10.728x` | `10.718x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `372983` (364.24 KiB) | `0.999x` | `10.734x` | `10.724x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003717` (3.82 MiB) | `282846` (276.22 KiB) | `0.999x` | `14.155x` | `14.142x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `245265` (239.52 KiB) | `0.999x` | `16.324x` | `16.309x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `328322` (320.63 KiB) | `0.999x` | `12.194x` | `12.183x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `3354477` (3.20 MiB) | `1042888` (1018.45 KiB) | `394674` (385.42 KiB) | `3.217x` | `2.642x` | `8.499x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `219059` (213.92 KiB) | `0.999x` | `18.277x` | `18.260x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `217216` (212.12 KiB) | `0.999x` | `18.432x` | `18.415x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `364866` (356.31 KiB) | `0.999x` | `10.973x` | `10.963x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `3767530` (3.59 MiB) | `859464` (839.32 KiB) | `235276` (229.76 KiB) | `4.384x` | `3.653x` | `16.013x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `204926` (200.12 KiB) | `0.999x` | `19.537x` | `19.519x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `205035` (200.23 KiB) | `0.999x` | `19.527x` | `19.509x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `218393` (213.27 KiB) | `0.999x` | `18.333x` | `18.316x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `215340` (210.29 KiB) | `0.999x` | `18.593x` | `18.575x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `81583` (79.67 KiB) | `283256` (276.62 KiB) | `57249` (55.91 KiB) | `0.288x` | `4.948x` | `1.425x` | `81583` (79.67 KiB) |
| `Params` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `0` (0 B) | `81428` (79.52 KiB) | `6912` (6.75 KiB) | `0.000x` | `11.781x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003757` (3.82 MiB) | `560336` (547.20 KiB) | `0.999x` | `7.145x` | `7.139x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003753` (3.82 MiB) | `526720` (514.38 KiB) | `0.999x` | `7.601x` | `7.594x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `293014` (286.15 KiB) | `0.999x` | `13.664x` | `13.651x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `3528017` (3.36 MiB) | `2962177` (2.82 MiB) | `994597` (971.29 KiB) | `1.191x` | `2.978x` | `3.547x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `220229` (215.07 KiB) | `0.999x` | `18.180x` | `18.163x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003747` (3.82 MiB) | `505707` (493.85 KiB) | `0.999x` | `7.917x` | `7.910x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003750` (3.82 MiB) | `516128` (504.03 KiB) | `0.999x` | `7.757x` | `7.750x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003760` (3.82 MiB) | `551321` (538.40 KiB) | `0.999x` | `7.262x` | `7.255x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `285624` (278.93 KiB) | `0.999x` | `14.017x` | `14.004x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:59` | `1000000` | `8000000` (7.63 MiB) | `3490384` (3.33 MiB) | `3054652` (2.91 MiB) | `2.292x` | `1.143x` | `2.619x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `287127` (280.40 KiB) | `0.999x` | `13.944x` | `13.931x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `269622` (263.30 KiB) | `0.999x` | `14.849x` | `14.836x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `304706` (297.56 KiB) | `0.999x` | `13.140x` | `13.127x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `204112` (199.33 KiB) | `0.999x` | `19.615x` | `19.597x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:87` | `1000000` | `13587860` (12.96 MiB) | `141854` (138.53 KiB) | `26884` (26.25 KiB) | `95.788x` | `5.277x` | `505.426x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003709` (3.82 MiB) | `205063` (200.26 KiB) | `0.999x` | `19.524x` | `19.506x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `256397` (250.39 KiB) | `0.999x` | `15.615x` | `15.601x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `206540` (201.70 KiB) | `0.999x` | `19.385x` | `19.367x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003718` (3.82 MiB) | `235273` (229.76 KiB) | `0.999x` | `17.017x` | `17.002x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004714` (7.63 MiB) | `1151399` (1.10 MiB) | `0.999x` | `6.952x` | `6.948x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:145` | `1000000` | `27797671` (26.51 MiB) | `21052485` (20.08 MiB) | `6746526` (6.43 MiB) | `1.320x` | `3.120x` | `4.120x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `3689257` (3.52 MiB) | `0.999x` | `1.085x` | `1.084x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `204076` (199.29 KiB) | `0.999x` | `19.619x` | `19.601x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `204076` (199.29 KiB) | `0.999x` | `19.619x` | `19.601x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `204076` (199.29 KiB) | `0.999x` | `19.619x` | `19.601x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003718` (3.82 MiB) | `295543` (288.62 KiB) | `0.999x` | `13.547x` | `13.534x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `204076` (199.29 KiB) | `0.999x` | `19.619x` | `19.601x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `1000000` (976.56 KiB) | `207337` (202.48 KiB) | `74334` (72.59 KiB) | `4.823x` | `2.789x` | `13.453x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:59` | `1000000` | `8000000` (7.63 MiB) | `3384213` (3.23 MiB) | `3055897` (2.91 MiB) | `2.364x` | `1.107x` | `2.618x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `333805` (325.98 KiB) | `0.999x` | `11.994x` | `11.983x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `302475` (295.39 KiB) | `0.999x` | `13.237x` | `13.224x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `318165` (310.71 KiB) | `0.999x` | `12.584x` | `12.572x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `385290` (376.26 KiB) | `0.999x` | `10.391x` | `10.382x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003719` (3.82 MiB) | `381452` (372.51 KiB) | `0.999x` | `10.496x` | `10.486x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003770` (3.82 MiB) | `706808` (690.24 KiB) | `0.999x` | `5.665x` | `5.659x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003718` (3.82 MiB) | `295106` (288.19 KiB) | `0.999x` | `13.567x` | `13.554x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `204077` (199.29 KiB) | `0.999x` | `19.619x` | `19.600x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003718` (3.82 MiB) | `258252` (252.20 KiB) | `0.999x` | `15.503x` | `15.489x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `2001192` (1.91 MiB) | `332127` (324.34 KiB) | `89669` (87.57 KiB) | `6.025x` | `3.704x` | `22.318x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `3325142` (3.17 MiB) | `964633` (942.02 KiB) | `253999` (248.05 KiB) | `3.447x` | `3.798x` | `13.091x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `0` (0 B) | `81428` (79.52 KiB) | `6912` (6.75 KiB) | `0.000x` | `11.781x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `0` (0 B) | `81428` (79.52 KiB) | `6912` (6.75 KiB) | `0.000x` | `11.781x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `204076` (199.29 KiB) | `0.999x` | `19.619x` | `19.601x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `268888` (262.59 KiB) | `0.999x` | `14.890x` | `14.876x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `326903` (319.24 KiB) | `0.999x` | `12.247x` | `12.236x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003753` (3.82 MiB) | `564134` (550.91 KiB) | `0.999x` | `7.097x` | `7.091x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003771` (3.82 MiB) | `1722741` (1.64 MiB) | `0.999x` | `2.324x` | `2.322x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003772` (3.82 MiB) | `1281263` (1.22 MiB) | `0.999x` | `3.125x` | `3.122x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003766` (3.82 MiB) | `808928` (789.97 KiB) | `0.999x` | `4.949x` | `4.945x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003708` (3.82 MiB) | `204461` (199.67 KiB) | `0.999x` | `19.582x` | `19.564x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `1024` (1.00 KiB) | `86733` (84.70 KiB) | `9592` (9.37 KiB) | `0.012x` | `9.042x` | `0.107x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004660` (7.63 MiB) | `405026` (395.53 KiB) | `0.999x` | `19.763x` | `19.752x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `0` (0 B) | `81428` (79.52 KiB) | `6912` (6.75 KiB) | `0.000x` | `11.781x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `3000000` (2.86 MiB) | `86349` (84.33 KiB) | `10314` (10.07 KiB) | `34.743x` | `8.372x` | `290.867x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `204076` (199.29 KiB) | `0.999x` | `19.619x` | `19.601x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `58030` (56.67 KiB) | `301351` (294.29 KiB) | `47779` (46.66 KiB) | `0.193x` | `6.307x` | `1.215x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `22051` (21.53 KiB) | `196999` (192.38 KiB) | `38443` (37.54 KiB) | `0.112x` | `5.124x` | `0.574x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `25445` (24.85 KiB) | `208514` (203.63 KiB) | `44424` (43.38 KiB) | `0.122x` | `4.694x` | `0.573x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `48191` (47.06 KiB) | `231124` (225.71 KiB) | `33750` (32.96 KiB) | `0.209x` | `6.848x` | `1.428x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `49433` (48.27 KiB) | `267477` (261.21 KiB) | `52027` (50.81 KiB) | `0.185x` | `5.141x` | `0.950x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `16873` (16.48 KiB) | `213222` (208.22 KiB) | `38786` (37.88 KiB) | `0.079x` | `5.497x` | `0.435x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `91870` (89.72 KiB) | `327656` (319.98 KiB) | `68261` (66.66 KiB) | `0.280x` | `4.800x` | `1.346x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `13001` (12.70 KiB) | `150177` (146.66 KiB) | `27880` (27.23 KiB) | `0.087x` | `5.387x` | `0.466x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `28101` (27.44 KiB) | `190375` (185.91 KiB) | `31113` (30.38 KiB) | `0.148x` | `6.119x` | `0.903x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:59` | `1000000` | `45607` (44.54 KiB) | `261404` (255.28 KiB) | `50283` (49.10 KiB) | `0.174x` | `5.199x` | `0.907x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `213466` (208.46 KiB) | `0.999x` | `18.756x` | `18.738x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004719` (7.63 MiB) | `3638544` (3.47 MiB) | `0.999x` | `2.200x` | `2.199x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004713` (7.63 MiB) | `4383261` (4.18 MiB) | `0.999x` | `1.826x` | `1.825x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003710` (3.82 MiB) | `204528` (199.73 KiB) | `0.999x` | `19.575x` | `19.557x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00000.parquet`: `34460` rows, `3771438` file bytes (3.60 MiB), `26764585` physical bytes (25.52 MiB), `16719667` encoded bytes (15.95 MiB), `3739473` compressed data bytes (3.57 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00001.parquet`: `34521` rows, `3700586` file bytes (3.53 MiB), `26563351` physical bytes (25.33 MiB), `16513943` encoded bytes (15.75 MiB), `3668279` compressed data bytes (3.50 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00002.parquet`: `34025` rows, `3718771` file bytes (3.55 MiB), `26618571` physical bytes (25.39 MiB), `16524112` encoded bytes (15.76 MiB), `3685671` compressed data bytes (3.51 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00003.parquet`: `34808` rows, `3719719` file bytes (3.55 MiB), `26963795` physical bytes (25.71 MiB), `16711533` encoded bytes (15.94 MiB), `3687474` compressed data bytes (3.52 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00004.parquet`: `34197` rows, `3685894` file bytes (3.52 MiB), `26609752` physical bytes (25.38 MiB), `16473775` encoded bytes (15.71 MiB), `3652952` compressed data bytes (3.48 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00005.parquet`: `34843` rows, `3691764` file bytes (3.52 MiB), `26800356` physical bytes (25.56 MiB), `16653915` encoded bytes (15.88 MiB), `3659592` compressed data bytes (3.49 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00006.parquet`: `34506` rows, `3694865` file bytes (3.52 MiB), `26574995` physical bytes (25.34 MiB), `16515938` encoded bytes (15.75 MiB), `3662375` compressed data bytes (3.49 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00007.parquet`: `34306` rows, `3670367` file bytes (3.50 MiB), `26635385` physical bytes (25.40 MiB), `16473750` encoded bytes (15.71 MiB), `3637761` compressed data bytes (3.47 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00008.parquet`: `34325` rows, `3661408` file bytes (3.49 MiB), `26626061` physical bytes (25.39 MiB), `16526668` encoded bytes (15.76 MiB), `3628800` compressed data bytes (3.46 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00009.parquet`: `34452` rows, `3641105` file bytes (3.47 MiB), `26400830` physical bytes (25.18 MiB), `16480258` encoded bytes (15.72 MiB), `3608782` compressed data bytes (3.44 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00010.parquet`: `34079` rows, `3640283` file bytes (3.47 MiB), `26415730` physical bytes (25.19 MiB), `16366931` encoded bytes (15.61 MiB), `3608281` compressed data bytes (3.44 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00011.parquet`: `34627` rows, `3622372` file bytes (3.45 MiB), `26641243` physical bytes (25.41 MiB), `16487979` encoded bytes (15.72 MiB), `3589900` compressed data bytes (3.42 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00012.parquet`: `34390` rows, `3608441` file bytes (3.44 MiB), `26426168` physical bytes (25.20 MiB), `16427095` encoded bytes (15.67 MiB), `3576457` compressed data bytes (3.41 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00013.parquet`: `34352` rows, `3677053` file bytes (3.51 MiB), `26333215` physical bytes (25.11 MiB), `16453038` encoded bytes (15.69 MiB), `3644858` compressed data bytes (3.48 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00014.parquet`: `34389` rows, `3820721` file bytes (3.64 MiB), `25991058` physical bytes (24.79 MiB), `16577914` encoded bytes (15.81 MiB), `3788513` compressed data bytes (3.61 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00015.parquet`: `34432` rows, `4151896` file bytes (3.96 MiB), `23501602` physical bytes (22.41 MiB), `16836918` encoded bytes (16.06 MiB), `4121571` compressed data bytes (3.93 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00016.parquet`: `33570` rows, `4204322` file bytes (4.01 MiB), `23809923` physical bytes (22.71 MiB), `16511072` encoded bytes (15.75 MiB), `4174105` compressed data bytes (3.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00017.parquet`: `33196` rows, `4252883` file bytes (4.06 MiB), `21244166` physical bytes (20.26 MiB), `17309938` encoded bytes (16.51 MiB), `4222579` compressed data bytes (4.03 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00018.parquet`: `33467` rows, `4388464` file bytes (4.19 MiB), `21620754` physical bytes (20.62 MiB), `17881386` encoded bytes (17.05 MiB), `4357746` compressed data bytes (4.16 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00019.parquet`: `33045` rows, `4466817` file bytes (4.26 MiB), `21559655` physical bytes (20.56 MiB), `17777053` encoded bytes (16.95 MiB), `4436470` compressed data bytes (4.23 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00020.parquet`: `33173` rows, `4291702` file bytes (4.09 MiB), `21138595` physical bytes (20.16 MiB), `17520291` encoded bytes (16.71 MiB), `4260921` compressed data bytes (4.06 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00021.parquet`: `33038` rows, `4373472` file bytes (4.17 MiB), `21235724` physical bytes (20.25 MiB), `17541035` encoded bytes (16.73 MiB), `4342609` compressed data bytes (4.14 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00022.parquet`: `33482` rows, `4349065` file bytes (4.15 MiB), `21436021` physical bytes (20.44 MiB), `17707004` encoded bytes (16.89 MiB), `4318444` compressed data bytes (4.12 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00023.parquet`: `33536` rows, `4330241` file bytes (4.13 MiB), `21423437` physical bytes (20.43 MiB), `17664532` encoded bytes (16.85 MiB), `4299856` compressed data bytes (4.10 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00024.parquet`: `32895` rows, `4416371` file bytes (4.21 MiB), `21437268` physical bytes (20.44 MiB), `17789223` encoded bytes (16.97 MiB), `4385687` compressed data bytes (4.18 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00025.parquet`: `33402` rows, `4303815` file bytes (4.10 MiB), `21256801` physical bytes (20.27 MiB), `17481042` encoded bytes (16.67 MiB), `4273666` compressed data bytes (4.08 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00026.parquet`: `33294` rows, `4334639` file bytes (4.13 MiB), `21277424` physical bytes (20.29 MiB), `17579751` encoded bytes (16.77 MiB), `4304137` compressed data bytes (4.10 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00027.parquet`: `33529` rows, `4267061` file bytes (4.07 MiB), `21307179` physical bytes (20.32 MiB), `17482025` encoded bytes (16.67 MiB), `4236486` compressed data bytes (4.04 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00028.parquet`: `33336` rows, `4273421` file bytes (4.08 MiB), `21152125` physical bytes (20.17 MiB), `17483462` encoded bytes (16.67 MiB), `4243209` compressed data bytes (4.05 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00029.parquet`: `16325` rows, `2150328` file bytes (2.05 MiB), `10632855` physical bytes (10.14 MiB), `8723586` encoded bytes (8.32 MiB), `2133462` compressed data bytes (2.03 MiB)
