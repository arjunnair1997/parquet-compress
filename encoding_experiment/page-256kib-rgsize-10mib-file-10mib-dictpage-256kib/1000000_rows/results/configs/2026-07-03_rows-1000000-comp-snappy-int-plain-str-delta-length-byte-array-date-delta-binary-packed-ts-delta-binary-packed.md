# ClickBench Parquet Experiment

- Started: `2026-07-03T23:30:39-04:00`
- Write elapsed: `11.036s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `703490627` (670.90 MiB)
- Compressed column data bytes after codec compression: `127333899` (121.44 MiB)
- Parquet file bytes: `128293589` (122.35 MiB)
- Physical/encoded ratio: `1.013x`
- Encoded/compressed-data ratio: `5.525x`
- Physical/compressed-data ratio: `5.595x`
- Physical/parquet-file ratio: `5.553x`
- Files: `31`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `delta-length-byte-array`
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
- Files read: `31`
- Elapsed: `6.892s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004874` (7.63 MiB) | `8005300` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `262062` (255.92 KiB) | `0.999x` | `15.278x` | `15.264x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:510` | `1000000` | `138409995` (132.00 MiB) | `140027101` (133.54 MiB) | `21321353` (20.33 MiB) | `0.988x` | `6.567x` | `6.492x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204379` (199.59 KiB) | `0.999x` | `19.590x` | `19.571x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `3368836` (3.21 MiB) | `3037787` (2.90 MiB) | `2.375x` | `1.109x` | `2.633x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `4000000` (3.81 MiB) | `51914` (50.70 KiB) | `8125` (7.93 KiB) | `77.051x` | `6.389x` | `492.308x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204399` (199.61 KiB) | `0.999x` | `19.588x` | `19.570x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003895` (3.82 MiB) | `719328` (702.47 KiB) | `0.999x` | `5.566x` | `5.561x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `396735` (387.44 KiB) | `0.999x` | `10.092x` | `10.082x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004876` (7.63 MiB) | `1084894` (1.03 MiB) | `0.999x` | `7.378x` | `7.374x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `204207` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `318136` (310.68 KiB) | `0.999x` | `12.585x` | `12.573x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `342206` (334.19 KiB) | `0.999x` | `11.700x` | `11.689x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:341` | `1000000` | `88562192` (84.46 MiB) | `89784417` (85.63 MiB) | `20785414` (19.82 MiB) | `0.986x` | `4.320x` | `4.261x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:305` | `1000000` | `79583339` (75.90 MiB) | `80835044` (77.09 MiB) | `19457003` (18.56 MiB) | `0.985x` | `4.155x` | `4.090x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003867` (3.82 MiB) | `492011` (480.48 KiB) | `0.999x` | `8.138x` | `8.130x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003866` (3.82 MiB) | `509045` (497.11 KiB) | `0.999x` | `7.865x` | `7.858x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003857` (3.82 MiB) | `458065` (447.33 KiB) | `0.999x` | `8.741x` | `8.732x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `272009` (265.63 KiB) | `0.999x` | `14.719x` | `14.705x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `232258` (226.81 KiB) | `0.999x` | `17.239x` | `17.222x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `373508` (364.75 KiB) | `0.999x` | `10.720x` | `10.709x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `373204` (364.46 KiB) | `0.999x` | `10.728x` | `10.718x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `283118` (276.48 KiB) | `0.999x` | `14.142x` | `14.128x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `245403` (239.65 KiB) | `0.999x` | `16.315x` | `16.300x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `328519` (320.82 KiB) | `0.999x` | `12.188x` | `12.176x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3354477` (3.20 MiB) | `3706693` (3.53 MiB) | `432545` (422.41 KiB) | `0.905x` | `8.569x` | `7.755x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `219242` (214.10 KiB) | `0.999x` | `18.262x` | `18.245x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `217401` (212.31 KiB) | `0.999x` | `18.417x` | `18.399x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `364983` (356.43 KiB) | `0.999x` | `10.970x` | `10.959x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3767530` (3.59 MiB) | `4017252` (3.83 MiB) | `328057` (320.37 KiB) | `0.938x` | `12.246x` | `11.484x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `204863` (200.06 KiB) | `0.999x` | `19.544x` | `19.525x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003833` (3.82 MiB) | `205002` (200.20 KiB) | `0.999x` | `19.531x` | `19.512x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `218526` (213.40 KiB) | `0.999x` | `18.322x` | `18.304x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `215461` (210.41 KiB) | `0.999x` | `18.583x` | `18.565x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `81583` (79.67 KiB) | `231341` (225.92 KiB) | `42797` (41.79 KiB) | `0.353x` | `5.406x` | `1.906x` | `81583` (79.67 KiB) |
| `Params` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `0` (0 B) | `41869` (40.89 KiB) | `4898` (4.78 KiB) | `0.000x` | `8.548x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003874` (3.82 MiB) | `560770` (547.63 KiB) | `0.999x` | `7.140x` | `7.133x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003873` (3.82 MiB) | `526874` (514.53 KiB) | `0.999x` | `7.599x` | `7.592x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `293209` (286.34 KiB) | `0.999x` | `13.655x` | `13.642x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3528017` (3.36 MiB) | `4225581` (4.03 MiB) | `1014907` (991.12 KiB) | `0.835x` | `4.164x` | `3.476x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `220387` (215.22 KiB) | `0.999x` | `18.167x` | `18.150x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003872` (3.82 MiB) | `505818` (493.96 KiB) | `0.999x` | `7.916x` | `7.908x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003868` (3.82 MiB) | `516560` (504.45 KiB) | `0.999x` | `7.751x` | `7.744x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003879` (3.82 MiB) | `551940` (539.00 KiB) | `0.999x` | `7.254x` | `7.247x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `285706` (279.01 KiB) | `0.999x` | `14.014x` | `14.000x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `3492403` (3.33 MiB) | `3050868` (2.91 MiB) | `2.291x` | `1.145x` | `2.622x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `287305` (280.57 KiB) | `0.999x` | `13.936x` | `13.922x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `269780` (263.46 KiB) | `0.999x` | `14.841x` | `14.827x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `304922` (297.78 KiB) | `0.999x` | `13.131x` | `13.118x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `204241` (199.45 KiB) | `0.999x` | `19.603x` | `19.585x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:91` | `1000000` | `13587860` (12.96 MiB) | `13657851` (13.03 MiB) | `700057` (683.65 KiB) | `0.995x` | `19.510x` | `19.410x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `205235` (200.42 KiB) | `0.999x` | `19.509x` | `19.490x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003833` (3.82 MiB) | `256545` (250.53 KiB) | `0.999x` | `15.607x` | `15.592x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `206701` (201.86 KiB) | `0.999x` | `19.370x` | `19.352x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `235424` (229.91 KiB) | `0.999x` | `17.007x` | `16.991x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004871` (7.63 MiB) | `1151559` (1.10 MiB) | `0.999x` | `6.951x` | `6.947x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:142` | `1000000` | `27797671` (26.51 MiB) | `28793780` (27.46 MiB) | `7049442` (6.72 MiB) | `0.965x` | `4.085x` | `3.943x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003896` (3.82 MiB) | `3688231` (3.52 MiB) | `0.999x` | `1.086x` | `1.085x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `204207` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `204207` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `204207` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `295675` (288.75 KiB) | `0.999x` | `13.541x` | `13.528x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `204207` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `1000000` (976.56 KiB) | `1043314` (1018.86 KiB) | `77131` (75.32 KiB) | `0.958x` | `13.527x` | `12.965x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `3384020` (3.23 MiB) | `3048006` (2.91 MiB) | `2.364x` | `1.110x` | `2.625x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003833` (3.82 MiB) | `333966` (326.14 KiB) | `0.999x` | `11.989x` | `11.977x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `302569` (295.48 KiB) | `0.999x` | `13.233x` | `13.220x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `318304` (310.84 KiB) | `0.999x` | `12.579x` | `12.567x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `385713` (376.67 KiB) | `0.999x` | `10.380x` | `10.370x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `381910` (372.96 KiB) | `0.999x` | `10.484x` | `10.474x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003896` (3.82 MiB) | `707189` (690.61 KiB) | `0.999x` | `5.662x` | `5.656x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `295374` (288.45 KiB) | `0.999x` | `13.555x` | `13.542x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204208` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `258408` (252.35 KiB) | `0.999x` | `15.494x` | `15.479x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `2001192` (1.91 MiB) | `2051472` (1.96 MiB) | `125389` (122.45 KiB) | `0.975x` | `16.361x` | `15.960x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3325142` (3.17 MiB) | `3638033` (3.47 MiB) | `352567` (344.30 KiB) | `0.914x` | `10.319x` | `9.431x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `0` (0 B) | `41869` (40.89 KiB) | `4898` (4.78 KiB) | `0.000x` | `8.548x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `0` (0 B) | `41869` (40.89 KiB) | `4898` (4.78 KiB) | `0.000x` | `8.548x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `204207` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `269130` (262.82 KiB) | `0.999x` | `14.877x` | `14.863x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `327302` (319.63 KiB) | `0.999x` | `12.233x` | `12.221x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003872` (3.82 MiB) | `565110` (551.87 KiB) | `0.999x` | `7.085x` | `7.078x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `1726130` (1.65 MiB) | `0.999x` | `2.320x` | `2.317x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `1283326` (1.22 MiB) | `0.999x` | `3.120x` | `3.117x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003896` (3.82 MiB) | `810079` (791.09 KiB) | `0.999x` | `4.943x` | `4.938x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `204595` (199.80 KiB) | `0.999x` | `19.570x` | `19.551x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `1024` (1.00 KiB) | `46334` (45.25 KiB) | `7287` (7.12 KiB) | `0.022x` | `6.358x` | `0.141x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004812` (7.63 MiB) | `405180` (395.68 KiB) | `0.999x` | `19.756x` | `19.744x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `0` (0 B) | `41869` (40.89 KiB) | `4898` (4.78 KiB) | `0.000x` | `8.548x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3000000` (2.86 MiB) | `3044392` (2.90 MiB) | `157396` (153.71 KiB) | `0.985x` | `19.342x` | `19.060x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `204207` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `58030` (56.67 KiB) | `216705` (211.63 KiB) | `37189` (36.32 KiB) | `0.268x` | `5.827x` | `1.560x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `22051` (21.53 KiB) | `122837` (119.96 KiB) | `28579` (27.91 KiB) | `0.180x` | `4.298x` | `0.772x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `25445` (24.85 KiB) | `130598` (127.54 KiB) | `34691` (33.88 KiB) | `0.195x` | `3.765x` | `0.733x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `48191` (47.06 KiB) | `155965` (152.31 KiB) | `24689` (24.11 KiB) | `0.309x` | `6.317x` | `1.952x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `49433` (48.27 KiB) | `188423` (184.01 KiB) | `41622` (40.65 KiB) | `0.262x` | `4.527x` | `1.188x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `16873` (16.48 KiB) | `133706` (130.57 KiB) | `28600` (27.93 KiB) | `0.126x` | `4.675x` | `0.590x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `91870` (89.72 KiB) | `253785` (247.84 KiB) | `57793` (56.44 KiB) | `0.362x` | `4.391x` | `1.590x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `13001` (12.70 KiB) | `94475` (92.26 KiB) | `22288` (21.77 KiB) | `0.138x` | `4.239x` | `0.583x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `28101` (27.44 KiB) | `128892` (125.87 KiB) | `25450` (24.85 KiB) | `0.218x` | `5.065x` | `1.104x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `45607` (44.54 KiB) | `211055` (206.11 KiB) | `46288` (45.20 KiB) | `0.216x` | `4.560x` | `0.985x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `213627` (208.62 KiB) | `0.999x` | `18.742x` | `18.724x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004876` (7.63 MiB) | `3640955` (3.47 MiB) | `0.999x` | `2.199x` | `2.197x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004876` (7.63 MiB) | `4386889` (4.18 MiB) | `0.999x` | `1.825x` | `1.824x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `204635` (199.84 KiB) | `0.999x` | `19.566x` | `19.547x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00000.parquet`: `33162` rows, `4136095` file bytes (3.94 MiB), `25755694` physical bytes (24.56 MiB), `25493656` encoded bytes (24.31 MiB), `4104168` compressed data bytes (3.91 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00001.parquet`: `33223` rows, `4072595` file bytes (3.88 MiB), `25546347` physical bytes (24.36 MiB), `25263458` encoded bytes (24.09 MiB), `4040286` compressed data bytes (3.85 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00002.parquet`: `32740` rows, `4075471` file bytes (3.89 MiB), `25627388` physical bytes (24.44 MiB), `25353583` encoded bytes (24.18 MiB), `4042678` compressed data bytes (3.86 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00003.parquet`: `32998` rows, `4070465` file bytes (3.88 MiB), `25582354` physical bytes (24.40 MiB), `25307837` encoded bytes (24.14 MiB), `4038102` compressed data bytes (3.85 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00004.parquet`: `33213` rows, `4086833` file bytes (3.90 MiB), `25794675` physical bytes (24.60 MiB), `25515965` encoded bytes (24.33 MiB), `4053708` compressed data bytes (3.87 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00005.parquet`: `33442` rows, `4043314` file bytes (3.86 MiB), `25814661` physical bytes (24.62 MiB), `25527444` encoded bytes (24.34 MiB), `4010934` compressed data bytes (3.83 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00006.parquet`: `33289` rows, `4068653` file bytes (3.88 MiB), `25581912` physical bytes (24.40 MiB), `25307667` encoded bytes (24.14 MiB), `4036489` compressed data bytes (3.85 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00007.parquet`: `33270` rows, `4077714` file bytes (3.89 MiB), `25801368` physical bytes (24.61 MiB), `25517512` encoded bytes (24.34 MiB), `4045211` compressed data bytes (3.86 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00008.parquet`: `33064` rows, `4044038` file bytes (3.86 MiB), `25610948` physical bytes (24.42 MiB), `25335250` encoded bytes (24.16 MiB), `4011271` compressed data bytes (3.83 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00009.parquet`: `33205` rows, `4029088` file bytes (3.84 MiB), `25618444` physical bytes (24.43 MiB), `25340260` encoded bytes (24.17 MiB), `3996737` compressed data bytes (3.81 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00010.parquet`: `33317` rows, `4086070` file bytes (3.90 MiB), `25788242` physical bytes (24.59 MiB), `25504832` encoded bytes (24.32 MiB), `4054093` compressed data bytes (3.87 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00011.parquet`: `32937` rows, `3956130` file bytes (3.77 MiB), `25287110` physical bytes (24.12 MiB), `25005789` encoded bytes (23.85 MiB), `3923910` compressed data bytes (3.74 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00012.parquet`: `33255` rows, `4011732` file bytes (3.83 MiB), `25644148` physical bytes (24.46 MiB), `25362209` encoded bytes (24.19 MiB), `3979469` compressed data bytes (3.80 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00013.parquet`: `33371` rows, `4047775` file bytes (3.86 MiB), `25609471` physical bytes (24.42 MiB), `25329869` encoded bytes (24.16 MiB), `4015762` compressed data bytes (3.83 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00014.parquet`: `33033` rows, `3993706` file bytes (3.81 MiB), `25476896` physical bytes (24.30 MiB), `25199445` encoded bytes (24.03 MiB), `3961331` compressed data bytes (3.78 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00015.parquet`: `33352` rows, `4306558` file bytes (4.11 MiB), `23865134` physical bytes (22.76 MiB), `23558127` encoded bytes (22.47 MiB), `4275542` compressed data bytes (4.08 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00016.parquet`: `33432` rows, `4393254` file bytes (4.19 MiB), `22995692` physical bytes (21.93 MiB), `22664476` encoded bytes (21.61 MiB), `4362964` compressed data bytes (4.16 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00017.parquet`: `32838` rows, `4436021` file bytes (4.23 MiB), `22600545` physical bytes (21.55 MiB), `22283057` encoded bytes (21.25 MiB), `4405744` compressed data bytes (4.20 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00018.parquet`: `33080` rows, `4432119` file bytes (4.23 MiB), `21112472` physical bytes (20.13 MiB), `20808446` encoded bytes (19.84 MiB), `4401802` compressed data bytes (4.20 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00019.parquet`: `33037` rows, `4514918` file bytes (4.31 MiB), `21465481` physical bytes (20.47 MiB), `21148460` encoded bytes (20.17 MiB), `4484349` compressed data bytes (4.28 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00020.parquet`: `32987` rows, `4505650` file bytes (4.30 MiB), `21180148` physical bytes (20.20 MiB), `20872599` encoded bytes (19.91 MiB), `4474873` compressed data bytes (4.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00021.parquet`: `33203` rows, `4488961` file bytes (4.28 MiB), `21243314` physical bytes (20.26 MiB), `20931189` encoded bytes (19.96 MiB), `4458282` compressed data bytes (4.25 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00022.parquet`: `32965` rows, `4490091` file bytes (4.28 MiB), `21214668` physical bytes (20.23 MiB), `20907450` encoded bytes (19.94 MiB), `4459490` compressed data bytes (4.25 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00023.parquet`: `33270` rows, `4503761` file bytes (4.30 MiB), `21455868` physical bytes (20.46 MiB), `21141743` encoded bytes (20.16 MiB), `4472923` compressed data bytes (4.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00024.parquet`: `33179` rows, `4468925` file bytes (4.26 MiB), `21298945` physical bytes (20.31 MiB), `20987043` encoded bytes (20.01 MiB), `4438287` compressed data bytes (4.23 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00025.parquet`: `32824` rows, `4522814` file bytes (4.31 MiB), `21244834` physical bytes (20.26 MiB), `20933040` encoded bytes (19.96 MiB), `4492446` compressed data bytes (4.28 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00026.parquet`: `33177` rows, `4372920` file bytes (4.17 MiB), `21014238` physical bytes (20.04 MiB), `20705288` encoded bytes (19.75 MiB), `4342656` compressed data bytes (4.14 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00027.parquet`: `32995` rows, `4457085` file bytes (4.25 MiB), `21092253` physical bytes (20.12 MiB), `20783217` encoded bytes (19.82 MiB), `4426266` compressed data bytes (4.22 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00028.parquet`: `33194` rows, `4350350` file bytes (4.15 MiB), `21015663` physical bytes (20.04 MiB), `20706184` encoded bytes (19.75 MiB), `4319982` compressed data bytes (4.12 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00029.parquet`: `33186` rows, `4460658` file bytes (4.25 MiB), `21294900` physical bytes (20.31 MiB), `20985949` encoded bytes (20.01 MiB), `4430178` compressed data bytes (4.22 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed/part-00030.parquet`: `5762` rows, `789825` file bytes (771.31 KiB), `3764811` physical bytes (3.59 MiB), `3709583` encoded bytes (3.54 MiB), `773966` compressed data bytes (755.83 KiB)
