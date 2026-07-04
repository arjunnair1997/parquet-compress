# ClickBench Parquet Experiment

- Started: `2026-07-03T15:29:36-04:00`
- Write elapsed: `11.871s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `715151956` (682.02 MiB)
- Compressed column data bytes after codec compression: `134796005` (128.55 MiB)
- Parquet file bytes: `135771440` (129.48 MiB)
- Physical/encoded ratio: `0.996x`
- Encoded/compressed-data ratio: `5.305x`
- Physical/compressed-data ratio: `5.285x`
- Physical/parquet-file ratio: `5.247x`
- Files: `31`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `delta-length-byte-array`
- Date encoding: `rle-dict`
- Timestamp encoding: `rle-dict`
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
- Files read: `31`
- Elapsed: `6.922s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004951` (7.63 MiB) | `8005383` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `262155` (256.01 KiB) | `0.999x` | `15.273x` | `15.258x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:502` | `1000000` | `138409995` (132.00 MiB) | `140029217` (133.54 MiB) | `21333079` (20.34 MiB) | `0.988x` | `6.564x` | `6.488x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204453` (199.66 KiB) | `0.999x` | `19.583x` | `19.564x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `7358537` (7.02 MiB) | `5537299` (5.28 MiB) | `1.087x` | `1.329x` | `1.445x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `5383` (5.26 KiB) | `5631` (5.50 KiB) | `743.080x` | `0.956x` | `710.353x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003895` (3.82 MiB) | `204468` (199.68 KiB) | `0.999x` | `19.582x` | `19.563x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003961` (3.82 MiB) | `719595` (702.73 KiB) | `0.999x` | `5.564x` | `5.559x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `396757` (387.46 KiB) | `0.999x` | `10.092x` | `10.082x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004955` (7.63 MiB) | `1085037` (1.03 MiB) | `0.999x` | `7.378x` | `7.373x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204277` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `318149` (310.69 KiB) | `0.999x` | `12.585x` | `12.573x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `342330` (334.31 KiB) | `0.999x` | `11.696x` | `11.685x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:342` | `1000000` | `88562192` (84.46 MiB) | `89784917` (85.63 MiB) | `20794324` (19.83 MiB) | `0.986x` | `4.318x` | `4.259x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:294` | `1000000` | `79583339` (75.90 MiB) | `80834238` (77.09 MiB) | `19478433` (18.58 MiB) | `0.985x` | `4.150x` | `4.086x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003927` (3.82 MiB) | `492067` (480.53 KiB) | `0.999x` | `8.137x` | `8.129x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003931` (3.82 MiB) | `509269` (497.33 KiB) | `0.999x` | `7.862x` | `7.854x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003917` (3.82 MiB) | `458273` (447.53 KiB) | `0.999x` | `8.737x` | `8.728x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `272029` (265.65 KiB) | `0.999x` | `14.719x` | `14.704x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `232356` (226.91 KiB) | `0.999x` | `17.232x` | `17.215x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `373493` (364.74 KiB) | `0.999x` | `10.720x` | `10.710x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `373343` (364.59 KiB) | `0.999x` | `10.724x` | `10.714x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `283101` (276.47 KiB) | `0.999x` | `14.143x` | `14.129x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `245465` (239.71 KiB) | `0.999x` | `16.311x` | `16.296x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `328566` (320.87 KiB) | `0.999x` | `12.186x` | `12.174x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `3354477` (3.20 MiB) | `3708060` (3.54 MiB) | `433024` (422.88 KiB) | `0.905x` | `8.563x` | `7.747x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `219223` (214.08 KiB) | `0.999x` | `18.264x` | `18.246x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `217409` (212.31 KiB) | `0.999x` | `18.416x` | `18.399x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `365240` (356.68 KiB) | `0.999x` | `10.962x` | `10.952x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `3767530` (3.59 MiB) | `4017259` (3.83 MiB) | `328469` (320.77 KiB) | `0.938x` | `12.230x` | `11.470x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `204998` (200.19 KiB) | `0.999x` | `19.531x` | `19.512x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003897` (3.82 MiB) | `205055` (200.25 KiB) | `0.999x` | `19.526x` | `19.507x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003896` (3.82 MiB) | `218589` (213.47 KiB) | `0.999x` | `18.317x` | `18.299x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `215534` (210.48 KiB) | `0.999x` | `18.577x` | `18.559x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `81583` (79.67 KiB) | `232005` (226.57 KiB) | `42824` (41.82 KiB) | `0.352x` | `5.418x` | `1.905x` | `81583` (79.67 KiB) |
| `Params` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `0` (0 B) | `41915` (40.93 KiB) | `4924` (4.81 KiB) | `0.000x` | `8.512x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003937` (3.82 MiB) | `561309` (548.15 KiB) | `0.999x` | `7.133x` | `7.126x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003931` (3.82 MiB) | `526989` (514.64 KiB) | `0.999x` | `7.598x` | `7.590x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003896` (3.82 MiB) | `293339` (286.46 KiB) | `0.999x` | `13.649x` | `13.636x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `3528017` (3.36 MiB) | `4223822` (4.03 MiB) | `1015455` (991.66 KiB) | `0.835x` | `4.160x` | `3.474x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `220419` (215.25 KiB) | `0.999x` | `18.165x` | `18.147x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003935` (3.82 MiB) | `505884` (494.03 KiB) | `0.999x` | `7.915x` | `7.907x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003929` (3.82 MiB) | `516951` (504.83 KiB) | `0.999x` | `7.745x` | `7.738x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003946` (3.82 MiB) | `551997` (539.06 KiB) | `0.999x` | `7.254x` | `7.246x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `285921` (279.22 KiB) | `0.999x` | `14.004x` | `13.990x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `7237063` (6.90 MiB) | `5466243` (5.21 MiB) | `1.105x` | `1.324x` | `1.464x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `287374` (280.64 KiB) | `0.999x` | `13.933x` | `13.919x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `269839` (263.51 KiB) | `0.999x` | `14.838x` | `14.824x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `305032` (297.88 KiB) | `0.999x` | `13.126x` | `13.113x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204318` (199.53 KiB) | `0.999x` | `19.596x` | `19.577x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:92` | `1000000` | `13587860` (12.96 MiB) | `13657767` (13.03 MiB) | `700772` (684.35 KiB) | `0.995x` | `19.490x` | `19.390x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `205297` (200.49 KiB) | `0.999x` | `19.503x` | `19.484x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `256581` (250.57 KiB) | `0.999x` | `15.605x` | `15.590x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `206758` (201.91 KiB) | `0.999x` | `19.365x` | `19.346x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003906` (3.82 MiB) | `235479` (229.96 KiB) | `0.999x` | `17.003x` | `16.987x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004954` (7.63 MiB) | `1151743` (1.10 MiB) | `0.999x` | `6.950x` | `6.946x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:148` | `1000000` | `27797671` (26.51 MiB) | `28790394` (27.46 MiB) | `7055252` (6.73 MiB) | `0.966x` | `4.081x` | `3.940x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003962` (3.82 MiB) | `3688305` (3.52 MiB) | `0.999x` | `1.086x` | `1.085x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204277` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204277` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204277` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003906` (3.82 MiB) | `295753` (288.82 KiB) | `0.999x` | `13.538x` | `13.525x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204277` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `1000000` (976.56 KiB) | `1043364` (1018.91 KiB) | `77230` (75.42 KiB) | `0.958x` | `13.510x` | `12.948x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `7355810` (7.02 MiB) | `5537170` (5.28 MiB) | `1.088x` | `1.328x` | `1.445x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `334164` (326.33 KiB) | `0.999x` | `11.982x` | `11.970x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `302738` (295.64 KiB) | `0.999x` | `13.226x` | `13.213x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `318446` (310.98 KiB) | `0.999x` | `12.573x` | `12.561x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `385702` (376.66 KiB) | `0.999x` | `10.381x` | `10.371x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `382008` (373.05 KiB) | `0.999x` | `10.481x` | `10.471x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003961` (3.82 MiB) | `707584` (691.00 KiB) | `0.999x` | `5.659x` | `5.653x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `295407` (288.48 KiB) | `0.999x` | `13.554x` | `13.541x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `204278` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003906` (3.82 MiB) | `258510` (252.45 KiB) | `0.999x` | `15.488x` | `15.473x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `2001192` (1.91 MiB) | `2051474` (1.96 MiB) | `125458` (122.52 KiB) | `0.975x` | `16.352x` | `15.951x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `3325142` (3.17 MiB) | `3638960` (3.47 MiB) | `352795` (344.53 KiB) | `0.914x` | `10.315x` | `9.425x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `0` (0 B) | `41915` (40.93 KiB) | `4924` (4.81 KiB) | `0.000x` | `8.512x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `0` (0 B) | `41915` (40.93 KiB) | `4924` (4.81 KiB) | `0.000x` | `8.512x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204277` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003907` (3.82 MiB) | `268996` (262.69 KiB) | `0.999x` | `14.885x` | `14.870x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `327465` (319.79 KiB) | `0.999x` | `12.227x` | `12.215x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003935` (3.82 MiB) | `565079` (551.83 KiB) | `0.999x` | `7.086x` | `7.079x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003963` (3.82 MiB) | `1725537` (1.65 MiB) | `0.999x` | `2.320x` | `2.318x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003963` (3.82 MiB) | `1283342` (1.22 MiB) | `0.999x` | `3.120x` | `3.117x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003959` (3.82 MiB) | `810824` (791.82 KiB) | `0.999x` | `4.938x` | `4.933x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204666` (199.87 KiB) | `0.999x` | `19.563x` | `19.544x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `1024` (1.00 KiB) | `46392` (45.30 KiB) | `7348` (7.18 KiB) | `0.022x` | `6.314x` | `0.139x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004893` (7.63 MiB) | `405270` (395.77 KiB) | `0.999x` | `19.752x` | `19.740x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `0` (0 B) | `41915` (40.93 KiB) | `4924` (4.81 KiB) | `0.000x` | `8.512x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `3000000` (2.86 MiB) | `3044437` (2.90 MiB) | `157429` (153.74 KiB) | `0.985x` | `19.338x` | `19.056x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204277` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `58030` (56.67 KiB) | `215801` (210.74 KiB) | `37321` (36.45 KiB) | `0.269x` | `5.782x` | `1.555x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `22051` (21.53 KiB) | `122823` (119.94 KiB) | `28803` (28.13 KiB) | `0.180x` | `4.264x` | `0.766x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `25445` (24.85 KiB) | `130394` (127.34 KiB) | `34828` (34.01 KiB) | `0.195x` | `3.744x` | `0.731x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `48191` (47.06 KiB) | `155646` (152.00 KiB) | `24740` (24.16 KiB) | `0.310x` | `6.291x` | `1.948x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `49433` (48.27 KiB) | `187805` (183.40 KiB) | `41788` (40.81 KiB) | `0.263x` | `4.494x` | `1.183x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `16873` (16.48 KiB) | `133216` (130.09 KiB) | `28833` (28.16 KiB) | `0.127x` | `4.620x` | `0.585x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `91870` (89.72 KiB) | `252714` (246.79 KiB) | `57749` (56.40 KiB) | `0.364x` | `4.376x` | `1.591x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `13001` (12.70 KiB) | `94213` (92.00 KiB) | `22091` (21.57 KiB) | `0.138x` | `4.265x` | `0.589x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `28101` (27.44 KiB) | `128562` (125.55 KiB) | `25339` (24.75 KiB) | `0.219x` | `5.074x` | `1.109x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `45607` (44.54 KiB) | `212412` (207.43 KiB) | `46371` (45.28 KiB) | `0.215x` | `4.581x` | `0.984x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `213678` (208.67 KiB) | `0.999x` | `18.738x` | `18.720x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004951` (7.63 MiB) | `3642092` (3.47 MiB) | `0.999x` | `2.198x` | `2.197x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004953` (7.63 MiB) | `4388128` (4.18 MiB) | `0.999x` | `1.824x` | `1.823x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `204733` (199.93 KiB) | `0.999x` | `19.557x` | `19.538x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00000.parquet`: `33060` rows, `4283864` file bytes (4.09 MiB), `25671405` physical bytes (24.48 MiB), `25629198` encoded bytes (24.44 MiB), `4251814` compressed data bytes (4.05 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00001.parquet`: `33102` rows, `4233765` file bytes (4.04 MiB), `25456125` physical bytes (24.28 MiB), `25407121` encoded bytes (24.23 MiB), `4201333` compressed data bytes (4.01 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00002.parquet`: `32629` rows, `4228344` file bytes (4.03 MiB), `25549489` physical bytes (24.37 MiB), `25507778` encoded bytes (24.33 MiB), `4195418` compressed data bytes (4.00 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00003.parquet`: `32880` rows, `4222112` file bytes (4.03 MiB), `25499037` physical bytes (24.32 MiB), `25453692` encoded bytes (24.27 MiB), `4189628` compressed data bytes (4.00 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00004.parquet`: `33121` rows, `4242317` file bytes (4.05 MiB), `25706850` physical bytes (24.52 MiB), `25659508` encoded bytes (24.47 MiB), `4209053` compressed data bytes (4.01 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00005.parquet`: `33333` rows, `4200237` file bytes (4.01 MiB), `25723565` physical bytes (24.53 MiB), `25671691` encoded bytes (24.48 MiB), `4167710` compressed data bytes (3.97 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00006.parquet`: `32901` rows, `4196544` file bytes (4.00 MiB), `25311964` physical bytes (24.14 MiB), `25265257` encoded bytes (24.09 MiB), `4164289` compressed data bytes (3.97 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00007.parquet`: `32918` rows, `4202695` file bytes (4.01 MiB), `25533723` physical bytes (24.35 MiB), `25488523` encoded bytes (24.31 MiB), `4170129` compressed data bytes (3.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00008.parquet`: `32956` rows, `4204100` file bytes (4.01 MiB), `25523218` physical bytes (24.34 MiB), `25477495` encoded bytes (24.30 MiB), `4171119` compressed data bytes (3.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00009.parquet`: `33105` rows, `4171372` file bytes (3.98 MiB), `25547829` physical bytes (24.36 MiB), `25496468` encoded bytes (24.32 MiB), `4138987` compressed data bytes (3.95 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00010.parquet`: `33013` rows, `4218726` file bytes (4.02 MiB), `25492033` physical bytes (24.31 MiB), `25443919` encoded bytes (24.27 MiB), `4186525` compressed data bytes (3.99 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00011.parquet`: `33021` rows, `4141531` file bytes (3.95 MiB), `25366568` physical bytes (24.19 MiB), `25313582` encoded bytes (24.14 MiB), `4109104` compressed data bytes (3.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00012.parquet`: `33146` rows, `4167708` file bytes (3.97 MiB), `25550688` physical bytes (24.37 MiB), `25503026` encoded bytes (24.32 MiB), `4135328` compressed data bytes (3.94 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00013.parquet`: `33150` rows, `4210283` file bytes (4.02 MiB), `25529418` physical bytes (24.35 MiB), `25482473` encoded bytes (24.30 MiB), `4178196` compressed data bytes (3.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00014.parquet`: `32800` rows, `4134397` file bytes (3.94 MiB), `25188426` physical bytes (24.02 MiB), `25135902` encoded bytes (23.97 MiB), `4102068` compressed data bytes (3.91 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00015.parquet`: `33067` rows, `4475464` file bytes (4.27 MiB), `23908408` physical bytes (22.80 MiB), `23984762` encoded bytes (22.87 MiB), `4444125` compressed data bytes (4.24 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00016.parquet`: `32959` rows, `4616143` file bytes (4.40 MiB), `22674370` physical bytes (21.62 MiB), `22809400` encoded bytes (21.75 MiB), `4585807` compressed data bytes (4.37 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00017.parquet`: `32164` rows, `4672003` file bytes (4.46 MiB), `22374482` physical bytes (21.34 MiB), `22553574` encoded bytes (21.51 MiB), `4641650` compressed data bytes (4.43 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00018.parquet`: `32434` rows, `4625706` file bytes (4.41 MiB), `20613971` physical bytes (19.66 MiB), `20862801` encoded bytes (19.90 MiB), `4595373` compressed data bytes (4.38 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00019.parquet`: `31875` rows, `4688566` file bytes (4.47 MiB), `20739536` physical bytes (19.78 MiB), `20976611` encoded bytes (20.00 MiB), `4657804` compressed data bytes (4.44 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00020.parquet`: `32298` rows, `4754871` file bytes (4.53 MiB), `20785747` physical bytes (19.82 MiB), `21030758` encoded bytes (20.06 MiB), `4724429` compressed data bytes (4.51 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00021.parquet`: `32299` rows, `4703629` file bytes (4.49 MiB), `20709206` physical bytes (19.75 MiB), `20952736` encoded bytes (19.98 MiB), `4672668` compressed data bytes (4.46 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00022.parquet`: `32115` rows, `4718463` file bytes (4.50 MiB), `20690000` physical bytes (19.73 MiB), `20938526` encoded bytes (19.97 MiB), `4687922` compressed data bytes (4.47 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00023.parquet`: `32477` rows, `4665085` file bytes (4.45 MiB), `20747265` physical bytes (19.79 MiB), `20992298` encoded bytes (20.02 MiB), `4634380` compressed data bytes (4.42 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00024.parquet`: `32435` rows, `4659524` file bytes (4.44 MiB), `20729714` physical bytes (19.77 MiB), `20970993` encoded bytes (20.00 MiB), `4629157` compressed data bytes (4.41 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00025.parquet`: `31980` rows, `4774466` file bytes (4.55 MiB), `20878153` physical bytes (19.91 MiB), `21118774` encoded bytes (20.14 MiB), `4743649` compressed data bytes (4.52 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00026.parquet`: `32340` rows, `4635318` file bytes (4.42 MiB), `20564934` physical bytes (19.61 MiB), `20809718` encoded bytes (19.85 MiB), `4605180` compressed data bytes (4.39 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00027.parquet`: `32194` rows, `4667381` file bytes (4.45 MiB), `20589765` physical bytes (19.64 MiB), `20834776` encoded bytes (19.87 MiB), `4636816` compressed data bytes (4.42 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00028.parquet`: `32415` rows, `4597028` file bytes (4.38 MiB), `20630996` physical bytes (19.68 MiB), `20875019` encoded bytes (19.91 MiB), `4566399` compressed data bytes (4.35 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00029.parquet`: `32514` rows, `4636548` file bytes (4.42 MiB), `20598501` physical bytes (19.64 MiB), `20843541` encoded bytes (19.88 MiB), `4606192` compressed data bytes (4.39 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-rle-dict-ts-rle-dict/part-00030.parquet`: `19299` rows, `2823250` file bytes (2.69 MiB), `12513238` physical bytes (11.93 MiB), `12662036` encoded bytes (12.08 MiB), `2793753` compressed data bytes (2.66 MiB)
