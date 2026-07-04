# ClickBench Parquet Experiment

- Started: `2026-07-03T14:58:32-04:00`
- Write elapsed: `11.548s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `717225339` (684.00 MiB)
- Compressed column data bytes after codec compression: `131021010` (124.95 MiB)
- Parquet file bytes: `131981649` (125.87 MiB)
- Physical/encoded ratio: `0.993x`
- Encoded/compressed-data ratio: `5.474x`
- Physical/compressed-data ratio: `5.437x`
- Physical/parquet-file ratio: `5.398x`
- Files: `31`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
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
- Rows read and compared: `1000000`
- Files read: `31`
- Elapsed: `6.833s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004880` (7.63 MiB) | `8005307` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `262099` (255.96 KiB) | `0.999x` | `15.276x` | `15.261x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:505` | `1000000` | `138409995` (132.00 MiB) | `140040620` (133.55 MiB) | `21338904` (20.35 MiB) | `0.988x` | `6.563x` | `6.486x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204385` (199.59 KiB) | `0.999x` | `19.590x` | `19.571x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004875` (7.63 MiB) | `4282758` (4.08 MiB) | `0.999x` | `1.869x` | `1.868x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:61, DICTIONARY_PAGE/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `5305` (5.18 KiB) | `5549` (5.42 KiB) | `754.006x` | `0.956x` | `720.851x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `204405` (199.61 KiB) | `0.999x` | `19.588x` | `19.569x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `719368` (702.51 KiB) | `0.999x` | `5.566x` | `5.560x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003843` (3.82 MiB) | `396585` (387.29 KiB) | `0.999x` | `10.096x` | `10.086x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004875` (7.63 MiB) | `1084896` (1.03 MiB) | `0.999x` | `7.378x` | `7.374x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204208` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `318082` (310.63 KiB) | `0.999x` | `12.587x` | `12.575x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `342283` (334.26 KiB) | `0.999x` | `11.697x` | `11.686x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:342` | `1000000` | `88562192` (84.46 MiB) | `89785538` (85.63 MiB) | `20794541` (19.83 MiB) | `0.986x` | `4.318x` | `4.259x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:298` | `1000000` | `79583339` (75.90 MiB) | `80833108` (77.09 MiB) | `19456070` (18.55 MiB) | `0.985x` | `4.155x` | `4.090x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003871` (3.82 MiB) | `491991` (480.46 KiB) | `0.999x` | `8.138x` | `8.130x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003869` (3.82 MiB) | `509023` (497.09 KiB) | `0.999x` | `7.866x` | `7.858x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003852` (3.82 MiB) | `458113` (447.38 KiB) | `0.999x` | `8.740x` | `8.731x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `272044` (265.67 KiB) | `0.999x` | `14.718x` | `14.704x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `232270` (226.83 KiB) | `0.999x` | `17.238x` | `17.221x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `373490` (364.74 KiB) | `0.999x` | `10.720x` | `10.710x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `373230` (364.48 KiB) | `0.999x` | `10.728x` | `10.717x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003833` (3.82 MiB) | `282953` (276.32 KiB) | `0.999x` | `14.150x` | `14.137x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `245406` (239.65 KiB) | `0.999x` | `16.315x` | `16.300x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `328465` (320.77 KiB) | `0.999x` | `12.190x` | `12.178x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3354477` (3.20 MiB) | `3707302` (3.54 MiB) | `433436` (423.28 KiB) | `0.905x` | `8.553x` | `7.739x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `219212` (214.07 KiB) | `0.999x` | `18.265x` | `18.247x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `217383` (212.29 KiB) | `0.999x` | `18.418x` | `18.401x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `365196` (356.64 KiB) | `0.999x` | `10.964x` | `10.953x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3767530` (3.59 MiB) | `4018004` (3.83 MiB) | `328036` (320.35 KiB) | `0.938x` | `12.249x` | `11.485x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `204988` (200.18 KiB) | `0.999x` | `19.532x` | `19.513x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `204982` (200.18 KiB) | `0.999x` | `19.533x` | `19.514x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `218506` (213.38 KiB) | `0.999x` | `18.324x` | `18.306x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `215461` (210.41 KiB) | `0.999x` | `18.583x` | `18.565x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `81583` (79.67 KiB) | `231662` (226.23 KiB) | `42953` (41.95 KiB) | `0.352x` | `5.393x` | `1.899x` | `81583` (79.67 KiB) |
| `Params` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `0` (0 B) | `41853` (40.87 KiB) | `4896` (4.78 KiB) | `0.000x` | `8.548x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003877` (3.82 MiB) | `560905` (547.76 KiB) | `0.999x` | `7.138x` | `7.131x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003869` (3.82 MiB) | `526843` (514.50 KiB) | `0.999x` | `7.600x` | `7.592x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `293199` (286.33 KiB) | `0.999x` | `13.656x` | `13.643x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3528017` (3.36 MiB) | `4223785` (4.03 MiB) | `1015203` (991.41 KiB) | `0.835x` | `4.161x` | `3.475x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003838` (3.82 MiB) | `220379` (215.21 KiB) | `0.999x` | `18.168x` | `18.151x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003877` (3.82 MiB) | `505891` (494.03 KiB) | `0.999x` | `7.915x` | `7.907x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003869` (3.82 MiB) | `516814` (504.70 KiB) | `0.999x` | `7.747x` | `7.740x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003883` (3.82 MiB) | `551875` (538.94 KiB) | `0.999x` | `7.255x` | `7.248x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `285747` (279.05 KiB) | `0.999x` | `14.012x` | `13.998x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004871` (7.63 MiB) | `4229161` (4.03 MiB) | `0.999x` | `1.893x` | `1.892x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `287291` (280.56 KiB) | `0.999x` | `13.937x` | `13.923x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `269743` (263.42 KiB) | `0.999x` | `14.843x` | `14.829x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `304920` (297.77 KiB) | `0.999x` | `13.131x` | `13.118x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003835` (3.82 MiB) | `204242` (199.46 KiB) | `0.999x` | `19.603x` | `19.585x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:91` | `1000000` | `13587860` (12.96 MiB) | `13658214` (13.03 MiB) | `699152` (682.77 KiB) | `0.995x` | `19.535x` | `19.435x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `205211` (200.40 KiB) | `0.999x` | `19.511x` | `19.492x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `256546` (250.53 KiB) | `0.999x` | `15.607x` | `15.592x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `206716` (201.87 KiB) | `0.999x` | `19.369x` | `19.350x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003845` (3.82 MiB) | `235451` (229.93 KiB) | `0.999x` | `17.005x` | `16.989x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004875` (7.63 MiB) | `1151631` (1.10 MiB) | `0.999x` | `6.951x` | `6.947x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:140` | `1000000` | `27797671` (26.51 MiB) | `28795793` (27.46 MiB) | `7054064` (6.73 MiB) | `0.965x` | `4.082x` | `3.941x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `3688228` (3.52 MiB) | `0.999x` | `1.086x` | `1.085x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204208` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204208` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204208` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003845` (3.82 MiB) | `295712` (288.78 KiB) | `0.999x` | `13.540x` | `13.527x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204208` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `1000000` (976.56 KiB) | `1043318` (1018.87 KiB) | `77092` (75.29 KiB) | `0.958x` | `13.533x` | `12.972x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004878` (7.63 MiB) | `4284716` (4.09 MiB) | `0.999x` | `1.868x` | `1.867x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `333977` (326.15 KiB) | `0.999x` | `11.988x` | `11.977x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `302566` (295.47 KiB) | `0.999x` | `13.233x` | `13.220x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `318320` (310.86 KiB) | `0.999x` | `12.578x` | `12.566x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `385691` (376.65 KiB) | `0.999x` | `10.381x` | `10.371x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003841` (3.82 MiB) | `381802` (372.85 KiB) | `0.999x` | `10.487x` | `10.477x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `706746` (690.18 KiB) | `0.999x` | `5.665x` | `5.660x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003847` (3.82 MiB) | `295425` (288.50 KiB) | `0.999x` | `13.553x` | `13.540x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003842` (3.82 MiB) | `204214` (199.43 KiB) | `0.999x` | `19.606x` | `19.587x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `258396` (252.34 KiB) | `0.999x` | `15.495x` | `15.480x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `2001192` (1.91 MiB) | `2051492` (1.96 MiB) | `125436` (122.50 KiB) | `0.975x` | `16.355x` | `15.954x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3325142` (3.17 MiB) | `3637763` (3.47 MiB) | `352764` (344.50 KiB) | `0.914x` | `10.312x` | `9.426x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `0` (0 B) | `41853` (40.87 KiB) | `4896` (4.78 KiB) | `0.000x` | `8.548x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `0` (0 B) | `41853` (40.87 KiB) | `4896` (4.78 KiB) | `0.000x` | `8.548x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204208` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `268978` (262.67 KiB) | `0.999x` | `14.885x` | `14.871x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003840` (3.82 MiB) | `327362` (319.69 KiB) | `0.999x` | `12.231x` | `12.219x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003875` (3.82 MiB) | `564612` (551.38 KiB) | `0.999x` | `7.091x` | `7.085x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003904` (3.82 MiB) | `1725690` (1.65 MiB) | `0.999x` | `2.320x` | `2.318x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003895` (3.82 MiB) | `1283140` (1.22 MiB) | `0.999x` | `3.120x` | `3.117x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `809971` (790.99 KiB) | `0.999x` | `4.943x` | `4.938x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003837` (3.82 MiB) | `204593` (199.80 KiB) | `0.999x` | `19.570x` | `19.551x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `1024` (1.00 KiB) | `46360` (45.27 KiB) | `7296` (7.12 KiB) | `0.022x` | `6.354x` | `0.140x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004816` (7.63 MiB) | `405189` (395.69 KiB) | `0.999x` | `19.756x` | `19.744x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `0` (0 B) | `41853` (40.87 KiB) | `4896` (4.78 KiB) | `0.000x` | `8.548x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `3000000` (2.86 MiB) | `3044413` (2.90 MiB) | `157374` (153.69 KiB) | `0.985x` | `19.345x` | `19.063x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003836` (3.82 MiB) | `204208` (199.42 KiB) | `0.999x` | `19.607x` | `19.588x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `58030` (56.67 KiB) | `216681` (211.60 KiB) | `37080` (36.21 KiB) | `0.268x` | `5.844x` | `1.565x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `22051` (21.53 KiB) | `122698` (119.82 KiB) | `28606` (27.94 KiB) | `0.180x` | `4.289x` | `0.771x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `25445` (24.85 KiB) | `130467` (127.41 KiB) | `34750` (33.94 KiB) | `0.195x` | `3.754x` | `0.732x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `48191` (47.06 KiB) | `155513` (151.87 KiB) | `24691` (24.11 KiB) | `0.310x` | `6.298x` | `1.952x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `49433` (48.27 KiB) | `187748` (183.35 KiB) | `41427` (40.46 KiB) | `0.263x` | `4.532x` | `1.193x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `16873` (16.48 KiB) | `133186` (130.06 KiB) | `28824` (28.15 KiB) | `0.127x` | `4.621x` | `0.585x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `91870` (89.72 KiB) | `252810` (246.88 KiB) | `57643` (56.29 KiB) | `0.363x` | `4.386x` | `1.594x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `13001` (12.70 KiB) | `94028` (91.82 KiB) | `22079` (21.56 KiB) | `0.138x` | `4.259x` | `0.589x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `28101` (27.44 KiB) | `128898` (125.88 KiB) | `25369` (24.77 KiB) | `0.218x` | `5.081x` | `1.108x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:61` | `1000000` | `45607` (44.54 KiB) | `211540` (206.58 KiB) | `46199` (45.12 KiB) | `0.216x` | `4.579x` | `0.987x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `213616` (208.61 KiB) | `0.999x` | `18.743x` | `18.725x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004876` (7.63 MiB) | `3641315` (3.47 MiB) | `0.999x` | `2.198x` | `2.197x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `8000000` (7.63 MiB) | `8004875` (7.63 MiB) | `4387065` (4.18 MiB) | `0.999x` | `1.825x` | `1.824x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:61` | `1000000` | `4000000` (3.81 MiB) | `4003839` (3.82 MiB) | `204662` (199.87 KiB) | `0.999x` | `19.563x` | `19.544x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00000.parquet`: `33122` rows, `4197701` file bytes (4.00 MiB), `25723561` physical bytes (24.53 MiB), `25906856` encoded bytes (24.71 MiB), `4165732` compressed data bytes (3.97 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00001.parquet`: `33171` rows, `4144656` file bytes (3.95 MiB), `25509930` physical bytes (24.33 MiB), `25685771` encoded bytes (24.50 MiB), `4112307` compressed data bytes (3.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00002.parquet`: `32706` rows, `4138419` file bytes (3.95 MiB), `25599070` physical bytes (24.41 MiB), `25778847` encoded bytes (24.58 MiB), `4105574` compressed data bytes (3.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00003.parquet`: `32955` rows, `4135717` file bytes (3.94 MiB), `25551069` physical bytes (24.37 MiB), `25729579` encoded bytes (24.54 MiB), `4103316` compressed data bytes (3.91 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00004.parquet`: `33164` rows, `4152846` file bytes (3.96 MiB), `25755586` physical bytes (24.56 MiB), `25937765` encoded bytes (24.74 MiB), `4119679` compressed data bytes (3.93 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00005.parquet`: `33400` rows, `4109964` file bytes (3.92 MiB), `25777567` physical bytes (24.58 MiB), `25952886` encoded bytes (24.75 MiB), `4077534` compressed data bytes (3.89 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00006.parquet`: `32963` rows, `4107604` file bytes (3.92 MiB), `25356937` physical bytes (24.18 MiB), `25538251` encoded bytes (24.36 MiB), `4075432` compressed data bytes (3.89 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00007.parquet`: `33216` rows, `4144039` file bytes (3.95 MiB), `25768336` physical bytes (24.57 MiB), `25946953` encoded bytes (24.74 MiB), `4111502` compressed data bytes (3.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00008.parquet`: `33047` rows, `4110665` file bytes (3.92 MiB), `25569097` physical bytes (24.38 MiB), `25750228` encoded bytes (24.56 MiB), `4077767` compressed data bytes (3.89 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00009.parquet`: `33162` rows, `4098266` file bytes (3.91 MiB), `25583982` physical bytes (24.40 MiB), `25763831` encoded bytes (24.57 MiB), `4065867` compressed data bytes (3.88 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00010.parquet`: `33052` rows, `4129072` file bytes (3.94 MiB), `25556129` physical bytes (24.37 MiB), `25733933` encoded bytes (24.54 MiB), `4097022` compressed data bytes (3.91 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00011.parquet`: `33069` rows, `4047872` file bytes (3.86 MiB), `25419958` physical bytes (24.24 MiB), `25596187` encoded bytes (24.41 MiB), `4015652` compressed data bytes (3.83 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00012.parquet`: `33227` rows, `4079459` file bytes (3.89 MiB), `25599269` physical bytes (24.41 MiB), `25780416` encoded bytes (24.59 MiB), `4047179` compressed data bytes (3.86 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00013.parquet`: `33281` rows, `4121631` file bytes (3.93 MiB), `25571261` physical bytes (24.39 MiB), `25750781` encoded bytes (24.56 MiB), `4089634` compressed data bytes (3.90 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00014.parquet`: `32751` rows, `4041440` file bytes (3.85 MiB), `25248680` physical bytes (24.08 MiB), `25424377` encoded bytes (24.25 MiB), `4009231` compressed data bytes (3.82 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00015.parquet`: `33225` rows, `4391270` file bytes (4.19 MiB), `23859995` physical bytes (22.75 MiB), `24004901` encoded bytes (22.89 MiB), `4360079` compressed data bytes (4.16 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00016.parquet`: `33306` rows, `4532779` file bytes (4.32 MiB), `22919298` physical bytes (21.86 MiB), `23041997` encoded bytes (21.97 MiB), `4502524` compressed data bytes (4.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00017.parquet`: `32423` rows, `4561410` file bytes (4.35 MiB), `22448292` physical bytes (21.41 MiB), `22571400` encoded bytes (21.53 MiB), `4531148` compressed data bytes (4.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00018.parquet`: `33023` rows, `4583156` file bytes (4.37 MiB), `21051979` physical bytes (20.08 MiB), `21199240` encoded bytes (20.22 MiB), `4552822` compressed data bytes (4.34 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00019.parquet`: `32751` rows, `4635702` file bytes (4.42 MiB), `21211811` physical bytes (20.23 MiB), `21348955` encoded bytes (20.36 MiB), `4605146` compressed data bytes (4.39 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00020.parquet`: `32246` rows, `4599451` file bytes (4.39 MiB), `20768489` physical bytes (19.81 MiB), `20910901` encoded bytes (19.94 MiB), `4568754` compressed data bytes (4.36 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00021.parquet`: `33128` rows, `4639403` file bytes (4.42 MiB), `21160601` physical bytes (20.18 MiB), `21303415` encoded bytes (20.32 MiB), `4608682` compressed data bytes (4.40 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00022.parquet`: `32739` rows, `4655559` file bytes (4.44 MiB), `21141221` physical bytes (20.16 MiB), `21284485` encoded bytes (20.30 MiB), `4625067` compressed data bytes (4.41 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00023.parquet`: `32955` rows, `4608221` file bytes (4.39 MiB), `21210480` physical bytes (20.23 MiB), `21351728` encoded bytes (20.36 MiB), `4577325` compressed data bytes (4.37 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00024.parquet`: `32766` rows, `4605850` file bytes (4.39 MiB), `21046959` physical bytes (20.07 MiB), `21186341` encoded bytes (20.20 MiB), `4575673` compressed data bytes (4.36 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00025.parquet`: `32797` rows, `4652652` file bytes (4.44 MiB), `21149039` physical bytes (20.17 MiB), `21289476` encoded bytes (20.30 MiB), `4621833` compressed data bytes (4.41 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00026.parquet`: `33021` rows, `4561273` file bytes (4.35 MiB), `20983488` physical bytes (20.01 MiB), `21125673` encoded bytes (20.15 MiB), `4531170` compressed data bytes (4.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00027.parquet`: `32932` rows, `4595316` file bytes (4.38 MiB), `20999318` physical bytes (20.03 MiB), `21138438` encoded bytes (20.16 MiB), `4564511` compressed data bytes (4.35 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00028.parquet`: `32994` rows, `4551346` file bytes (4.34 MiB), `21034486` physical bytes (20.06 MiB), `21178219` encoded bytes (20.20 MiB), `4520877` compressed data bytes (4.31 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00029.parquet`: `33211` rows, `4607340` file bytes (4.39 MiB), `21196925` physical bytes (20.21 MiB), `21343102` encoded bytes (20.35 MiB), `4576835` compressed data bytes (4.36 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-plain/part-00030.parquet`: `10197` rows, `1441570` file bytes (1.37 MiB), `6625811` physical bytes (6.32 MiB), `6670407` encoded bytes (6.36 MiB), `1425136` compressed data bytes (1.36 MiB)
