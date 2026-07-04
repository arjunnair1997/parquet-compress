# ClickBench Parquet Experiment

- Started: `2026-07-03T23:29:26-04:00`
- Write elapsed: `10.992s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `707451252` (674.68 MiB)
- Compressed column data bytes after codec compression: `127535033` (121.63 MiB)
- Parquet file bytes: `128494708` (122.54 MiB)
- Physical/encoded ratio: `1.007x`
- Encoded/compressed-data ratio: `5.547x`
- Physical/compressed-data ratio: `5.586x`
- Physical/parquet-file ratio: `5.544x`
- Files: `31`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `delta-length-byte-array`
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
- Files read: `31`
- Elapsed: `6.828s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004873` (7.63 MiB) | `8005299` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `262107` (255.96 KiB) | `0.999x` | `15.276x` | `15.261x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:510` | `1000000` | `138409995` (132.00 MiB) | `140027711` (133.54 MiB) | `21324409` (20.34 MiB) | `0.988x` | `6.567x` | `6.491x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `204386` (199.60 KiB) | `0.999x` | `19.590x` | `19.571x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `3369390` (3.21 MiB) | `3034102` (2.89 MiB) | `2.374x` | `1.111x` | `2.637x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `204379` (199.59 KiB) | `0.999x` | `19.590x` | `19.571x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `204399` (199.61 KiB) | `0.999x` | `19.588x` | `19.570x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `719381` (702.52 KiB) | `0.999x` | `5.566x` | `5.560x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `396725` (387.43 KiB) | `0.999x` | `10.092x` | `10.083x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004875` (7.63 MiB) | `1084921` (1.03 MiB) | `0.999x` | `7.378x` | `7.374x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204208` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `318193` (310.74 KiB) | `0.999x` | `12.583x` | `12.571x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `342206` (334.19 KiB) | `0.999x` | `11.700x` | `11.689x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:341` | `1000000` | `88562192` (84.46 MiB) | `89785659` (85.63 MiB) | `20787905` (19.82 MiB) | `0.986x` | `4.319x` | `4.260x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:305` | `1000000` | `79583339` (75.90 MiB) | `80835228` (77.09 MiB) | `19458967` (18.56 MiB) | `0.985x` | `4.154x` | `4.090x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003868` (3.82 MiB) | `491977` (480.45 KiB) | `0.999x` | `8.138x` | `8.130x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003866` (3.82 MiB) | `509066` (497.13 KiB) | `0.999x` | `7.865x` | `7.858x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003854` (3.82 MiB) | `458113` (447.38 KiB) | `0.999x` | `8.740x` | `8.731x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `272037` (265.66 KiB) | `0.999x` | `14.718x` | `14.704x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `232278` (226.83 KiB) | `0.999x` | `17.237x` | `17.221x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `373496` (364.74 KiB) | `0.999x` | `10.720x` | `10.710x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `373143` (364.40 KiB) | `0.999x` | `10.730x` | `10.720x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `283059` (276.42 KiB) | `0.999x` | `14.145x` | `14.131x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `245431` (239.68 KiB) | `0.999x` | `16.314x` | `16.298x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `328464` (320.77 KiB) | `0.999x` | `12.190x` | `12.178x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3354477` (3.20 MiB) | `3708666` (3.54 MiB) | `432876` (422.73 KiB) | `0.904x` | `8.568x` | `7.749x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `219230` (214.09 KiB) | `0.999x` | `18.263x` | `18.246x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `217394` (212.30 KiB) | `0.999x` | `18.417x` | `18.400x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `364971` (356.42 KiB) | `0.999x` | `10.970x` | `10.960x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3767530` (3.59 MiB) | `4016952` (3.83 MiB) | `328318` (320.62 KiB) | `0.938x` | `12.235x` | `11.475x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204854` (200.05 KiB) | `0.999x` | `19.545x` | `19.526x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `205008` (200.20 KiB) | `0.999x` | `19.530x` | `19.511x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `218506` (213.38 KiB) | `0.999x` | `18.324x` | `18.306x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `215449` (210.40 KiB) | `0.999x` | `18.584x` | `18.566x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `81583` (79.67 KiB) | `231424` (226.00 KiB) | `43022` (42.01 KiB) | `0.353x` | `5.379x` | `1.896x` | `81583` (79.67 KiB) |
| `Params` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `0` (0 B) | `41859` (40.88 KiB) | `4893` (4.78 KiB) | `0.000x` | `8.555x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003876` (3.82 MiB) | `560811` (547.67 KiB) | `0.999x` | `7.139x` | `7.133x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003876` (3.82 MiB) | `526939` (514.59 KiB) | `0.999x` | `7.598x` | `7.591x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003833` (3.82 MiB) | `293244` (286.37 KiB) | `0.999x` | `13.654x` | `13.641x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3528017` (3.36 MiB) | `4224516` (4.03 MiB) | `1014712` (990.93 KiB) | `0.835x` | `4.163x` | `3.477x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `220362` (215.20 KiB) | `0.999x` | `18.169x` | `18.152x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003874` (3.82 MiB) | `505890` (494.03 KiB) | `0.999x` | `7.915x` | `7.907x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003868` (3.82 MiB) | `516703` (504.59 KiB) | `0.999x` | `7.749x` | `7.741x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003885` (3.82 MiB) | `552029` (539.09 KiB) | `0.999x` | `7.253x` | `7.246x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `285706` (279.01 KiB) | `0.999x` | `14.014x` | `14.000x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `3492333` (3.33 MiB) | `3049975` (2.91 MiB) | `2.291x` | `1.145x` | `2.623x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `287333` (280.60 KiB) | `0.999x` | `13.934x` | `13.921x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `269807` (263.48 KiB) | `0.999x` | `14.840x` | `14.825x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `304955` (297.81 KiB) | `0.999x` | `13.129x` | `13.117x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204249` (199.46 KiB) | `0.999x` | `19.603x` | `19.584x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:91` | `1000000` | `13587860` (12.96 MiB) | `13657913` (13.03 MiB) | `699978` (683.57 KiB) | `0.995x` | `19.512x` | `19.412x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `205224` (200.41 KiB) | `0.999x` | `19.510x` | `19.491x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `256530` (250.52 KiB) | `0.999x` | `15.608x` | `15.593x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `206732` (201.89 KiB) | `0.999x` | `19.367x` | `19.349x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `235437` (229.92 KiB) | `0.999x` | `17.006x` | `16.990x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004868` (7.63 MiB) | `1151558` (1.10 MiB) | `0.999x` | `6.951x` | `6.947x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:142` | `1000000` | `27797671` (26.51 MiB) | `28792683` (27.46 MiB) | `7048161` (6.72 MiB) | `0.965x` | `4.085x` | `3.944x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `3688231` (3.52 MiB) | `0.999x` | `1.086x` | `1.085x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204208` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204208` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204208` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `295687` (288.76 KiB) | `0.999x` | `13.541x` | `13.528x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204208` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `1000000` (976.56 KiB) | `1043300` (1018.85 KiB) | `77109` (75.30 KiB) | `0.958x` | `13.530x` | `12.969x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:61` | `1000000` | `8000000` (7.63 MiB) | `3385194` (3.23 MiB) | `3051076` (2.91 MiB) | `2.363x` | `1.110x` | `2.622x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `333995` (326.17 KiB) | `0.999x` | `11.988x` | `11.976x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `302606` (295.51 KiB) | `0.999x` | `13.231x` | `13.219x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `318321` (310.86 KiB) | `0.999x` | `12.578x` | `12.566x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `385678` (376.64 KiB) | `0.999x` | `10.381x` | `10.371x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `381959` (373.01 KiB) | `0.999x` | `10.482x` | `10.472x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003896` (3.82 MiB) | `707218` (690.64 KiB) | `0.999x` | `5.661x` | `5.656x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `295303` (288.38 KiB) | `0.999x` | `13.558x` | `13.545x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204207` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `258403` (252.35 KiB) | `0.999x` | `15.495x` | `15.480x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `2001192` (1.91 MiB) | `2051458` (1.96 MiB) | `125341` (122.40 KiB) | `0.975x` | `16.367x` | `15.966x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3325142` (3.17 MiB) | `3638776` (3.47 MiB) | `352830` (344.56 KiB) | `0.914x` | `10.313x` | `9.424x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `0` (0 B) | `41859` (40.88 KiB) | `4893` (4.78 KiB) | `0.000x` | `8.555x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `0` (0 B) | `41859` (40.88 KiB) | `4893` (4.78 KiB) | `0.000x` | `8.555x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204208` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003845` (3.82 MiB) | `269304` (262.99 KiB) | `0.999x` | `14.867x` | `14.853x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `327297` (319.63 KiB) | `0.999x` | `12.233x` | `12.221x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003876` (3.82 MiB) | `564879` (551.64 KiB) | `0.999x` | `7.088x` | `7.081x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003896` (3.82 MiB) | `1725296` (1.65 MiB) | `0.999x` | `2.321x` | `2.318x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003896` (3.82 MiB) | `1283129` (1.22 MiB) | `0.999x` | `3.120x` | `3.117x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003891` (3.82 MiB) | `809941` (790.96 KiB) | `0.999x` | `4.943x` | `4.939x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204587` (199.79 KiB) | `0.999x` | `19.570x` | `19.552x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `1024` (1.00 KiB) | `46400` (45.31 KiB) | `7270` (7.10 KiB) | `0.022x` | `6.382x` | `0.141x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004815` (7.63 MiB) | `405184` (395.69 KiB) | `0.999x` | `19.756x` | `19.744x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `0` (0 B) | `41859` (40.88 KiB) | `4893` (4.78 KiB) | `0.000x` | `8.555x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3000000` (2.86 MiB) | `3044368` (2.90 MiB) | `157392` (153.70 KiB) | `0.985x` | `19.343x` | `19.061x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204208` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `58030` (56.67 KiB) | `216576` (211.50 KiB) | `36922` (36.06 KiB) | `0.268x` | `5.866x` | `1.572x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `22051` (21.53 KiB) | `122937` (120.06 KiB) | `28632` (27.96 KiB) | `0.179x` | `4.294x` | `0.770x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `25445` (24.85 KiB) | `130843` (127.78 KiB) | `34711` (33.90 KiB) | `0.194x` | `3.769x` | `0.733x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `48191` (47.06 KiB) | `156172` (152.51 KiB) | `24509` (23.93 KiB) | `0.309x` | `6.372x` | `1.966x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `49433` (48.27 KiB) | `189096` (184.66 KiB) | `41796` (40.82 KiB) | `0.261x` | `4.524x` | `1.183x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `16873` (16.48 KiB) | `134213` (131.07 KiB) | `28694` (28.02 KiB) | `0.126x` | `4.677x` | `0.588x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `91870` (89.72 KiB) | `254695` (248.73 KiB) | `57886` (56.53 KiB) | `0.361x` | `4.400x` | `1.587x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `13001` (12.70 KiB) | `94848` (92.62 KiB) | `22261` (21.74 KiB) | `0.137x` | `4.261x` | `0.584x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `28101` (27.44 KiB) | `129719` (126.68 KiB) | `25511` (24.91 KiB) | `0.217x` | `5.085x` | `1.102x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `45607` (44.54 KiB) | `211893` (206.93 KiB) | `46326` (45.24 KiB) | `0.215x` | `4.574x` | `0.984x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `213614` (208.61 KiB) | `0.999x` | `18.743x` | `18.725x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004874` (7.63 MiB) | `3641031` (3.47 MiB) | `0.999x` | `2.199x` | `2.197x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004875` (7.63 MiB) | `4386816` (4.18 MiB) | `0.999x` | `1.825x` | `1.824x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204647` (199.85 KiB) | `0.999x` | `19.565x` | `19.546x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00000.parquet`: `33159` rows, `4142228` file bytes (3.95 MiB), `25752972` physical bytes (24.56 MiB), `25622050` encoded bytes (24.44 MiB), `4110297` compressed data bytes (3.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00001.parquet`: `33220` rows, `4078467` file bytes (3.89 MiB), `25544801` physical bytes (24.36 MiB), `25392787` encoded bytes (24.22 MiB), `4046154` compressed data bytes (3.86 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00002.parquet`: `32736` rows, `4081010` file bytes (3.89 MiB), `25624047` physical bytes (24.44 MiB), `25480016` encoded bytes (24.30 MiB), `4048211` compressed data bytes (3.86 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00003.parquet`: `32995` rows, `4076355` file bytes (3.89 MiB), `25577832` physical bytes (24.39 MiB), `25433336` encoded bytes (24.26 MiB), `4043988` compressed data bytes (3.86 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00004.parquet`: `33207` rows, `4093545` file bytes (3.90 MiB), `25790496` physical bytes (24.60 MiB), `25645200` encoded bytes (24.46 MiB), `4060416` compressed data bytes (3.87 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00005.parquet`: `33438` rows, `4049904` file bytes (3.86 MiB), `25811014` physical bytes (24.62 MiB), `25656400` encoded bytes (24.47 MiB), `4017522` compressed data bytes (3.83 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00006.parquet`: `33285` rows, `4072878` file bytes (3.88 MiB), `25582690` physical bytes (24.40 MiB), `25439073` encoded bytes (24.26 MiB), `4040710` compressed data bytes (3.85 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00007.parquet`: `33264` rows, `4086107` file bytes (3.90 MiB), `25794659` physical bytes (24.60 MiB), `25645124` encoded bytes (24.46 MiB), `4053596` compressed data bytes (3.87 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00008.parquet`: `33061` rows, `4047977` file bytes (3.86 MiB), `25607816` physical bytes (24.42 MiB), `25462572` encoded bytes (24.28 MiB), `4015208` compressed data bytes (3.83 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00009.parquet`: `33202` rows, `4034407` file bytes (3.85 MiB), `25617003` physical bytes (24.43 MiB), `25467040` encoded bytes (24.29 MiB), `4002048` compressed data bytes (3.82 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00010.parquet`: `33315` rows, `4094982` file bytes (3.91 MiB), `25782694` physical bytes (24.59 MiB), `25630152` encoded bytes (24.44 MiB), `4062999` compressed data bytes (3.87 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00011.parquet`: `32917` rows, `3963263` file bytes (3.78 MiB), `25287264` physical bytes (24.12 MiB), `25138131` encoded bytes (23.97 MiB), `3931040` compressed data bytes (3.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00012.parquet`: `33267` rows, `4015163` file bytes (3.83 MiB), `25641045` physical bytes (24.45 MiB), `25490655` encoded bytes (24.31 MiB), `3982898` compressed data bytes (3.80 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00013.parquet`: `33362` rows, `4056836` file bytes (3.87 MiB), `25602586` physical bytes (24.42 MiB), `25457727` encoded bytes (24.28 MiB), `4024806` compressed data bytes (3.84 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00014.parquet`: `33024` rows, `4001230` file bytes (3.82 MiB), `25472550` physical bytes (24.29 MiB), `25327474` encoded bytes (24.15 MiB), `3968854` compressed data bytes (3.78 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00015.parquet`: `33341` rows, `4310720` file bytes (4.11 MiB), `23862480` physical bytes (22.76 MiB), `23686591` encoded bytes (22.59 MiB), `4279704` compressed data bytes (4.08 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00016.parquet`: `33436` rows, `4400670` file bytes (4.20 MiB), `22987550` physical bytes (21.92 MiB), `22790237` encoded bytes (21.73 MiB), `4370378` compressed data bytes (4.17 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00017.parquet`: `32837` rows, `4435291` file bytes (4.23 MiB), `22599364` physical bytes (21.55 MiB), `22412214` encoded bytes (21.37 MiB), `4405012` compressed data bytes (4.20 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00018.parquet`: `33043` rows, `4437055` file bytes (4.23 MiB), `21105065` physical bytes (20.13 MiB), `20932835` encoded bytes (19.96 MiB), `4406736` compressed data bytes (4.20 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00019.parquet`: `33030` rows, `4520252` file bytes (4.31 MiB), `21464368` physical bytes (20.47 MiB), `21276870` encoded bytes (20.29 MiB), `4489680` compressed data bytes (4.28 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00020.parquet`: `32989` rows, `4509389` file bytes (4.30 MiB), `21174140` physical bytes (20.19 MiB), `20998354` encoded bytes (20.03 MiB), `4478609` compressed data bytes (4.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00021.parquet`: `33187` rows, `4497489` file bytes (4.29 MiB), `21238901` physical bytes (20.25 MiB), `21058254` encoded bytes (20.08 MiB), `4466808` compressed data bytes (4.26 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00022.parquet`: `32961` rows, `4498750` file bytes (4.29 MiB), `21214236` physical bytes (20.23 MiB), `21036594` encoded bytes (20.06 MiB), `4468195` compressed data bytes (4.26 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00023.parquet`: `33049` rows, `4466241` file bytes (4.26 MiB), `21301600` physical bytes (20.31 MiB), `21121075` encoded bytes (20.14 MiB), `4435435` compressed data bytes (4.23 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00024.parquet`: `33154` rows, `4482153` file bytes (4.27 MiB), `21286129` physical bytes (20.30 MiB), `21104533` encoded bytes (20.13 MiB), `4451529` compressed data bytes (4.25 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00025.parquet`: `33035` rows, `4566644` file bytes (4.36 MiB), `21386337` physical bytes (20.40 MiB), `21204101` encoded bytes (20.22 MiB), `4536236` compressed data bytes (4.33 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00026.parquet`: `33157` rows, `4378746` file bytes (4.18 MiB), `21008883` physical bytes (20.04 MiB), `20830225` encoded bytes (19.87 MiB), `4348528` compressed data bytes (4.15 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00027.parquet`: `33013` rows, `4461903` file bytes (4.26 MiB), `21089743` physical bytes (20.11 MiB), `20909318` encoded bytes (19.94 MiB), `4431083` compressed data bytes (4.23 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00028.parquet`: `33179` rows, `4354863` file bytes (4.15 MiB), `21012852` physical bytes (20.04 MiB), `20835410` encoded bytes (19.87 MiB), `4324491` compressed data bytes (4.12 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00029.parquet`: `33200` rows, `4460035` file bytes (4.25 MiB), `21293723` physical bytes (20.31 MiB), `21117302` encoded bytes (20.14 MiB), `4429567` compressed data bytes (4.22 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed/part-00030.parquet`: `5937` rows, `820155` file bytes (800.93 KiB), `3883784` physical bytes (3.70 MiB), `3849602` encoded bytes (3.67 MiB), `804295` compressed data bytes (785.44 KiB)
