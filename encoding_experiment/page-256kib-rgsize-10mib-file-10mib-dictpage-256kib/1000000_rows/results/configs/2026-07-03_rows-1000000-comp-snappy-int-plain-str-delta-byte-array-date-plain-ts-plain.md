# ClickBench Parquet Experiment

- Started: `2026-07-03T19:40:57-04:00`
- Write elapsed: `11.645s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `518922971` (494.88 MiB)
- Compressed column data bytes after codec compression: `120804059` (115.21 MiB)
- Parquet file bytes: `121746190` (116.11 MiB)
- Physical/encoded ratio: `1.373x`
- Encoded/compressed-data ratio: `4.296x`
- Physical/compressed-data ratio: `5.897x`
- Physical/parquet-file ratio: `5.852x`
- Files: `30`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `delta-byte-array`
- Date encoding: `plain`
- Timestamp encoding: `plain`
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
- Elapsed: `7.058s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `8004790` (7.63 MiB) | `8005208` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `262013` (255.87 KiB) | `0.999x` | `15.281x` | `15.266x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:510` | `1000000` | `138409995` (132.00 MiB) | `64482734` (61.50 MiB) | `17105665` (16.31 MiB) | `2.146x` | `3.770x` | `8.091x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003769` (3.82 MiB) | `204303` (199.51 KiB) | `0.999x` | `19.597x` | `19.579x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `8004794` (7.63 MiB) | `4283088` (4.08 MiB) | `0.999x` | `1.869x` | `1.868x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003773` (3.82 MiB) | `204307` (199.52 KiB) | `0.999x` | `19.597x` | `19.578x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003773` (3.82 MiB) | `204329` (199.54 KiB) | `0.999x` | `19.595x` | `19.576x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `719163` (702.31 KiB) | `0.999x` | `5.567x` | `5.562x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003778` (3.82 MiB) | `396483` (387.19 KiB) | `0.999x` | `10.098x` | `10.089x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `8004790` (7.63 MiB) | `1084820` (1.03 MiB) | `0.999x` | `7.379x` | `7.374x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003772` (3.82 MiB) | `204142` (199.36 KiB) | `0.999x` | `19.613x` | `19.594x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003776` (3.82 MiB) | `318070` (310.62 KiB) | `0.999x` | `12.588x` | `12.576x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `342050` (334.03 KiB) | `0.999x` | `11.705x` | `11.694x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:334` | `1000000` | `88562192` (84.46 MiB) | `40460653` (38.59 MiB) | `18055310` (17.22 MiB) | `2.189x` | `2.241x` | `4.905x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:304` | `1000000` | `79583339` (75.90 MiB) | `38989131` (37.18 MiB) | `17323045` (16.52 MiB) | `2.041x` | `2.251x` | `4.594x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003803` (3.82 MiB) | `491957` (480.43 KiB) | `0.999x` | `8.139x` | `8.131x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003803` (3.82 MiB) | `509037` (497.11 KiB) | `0.999x` | `7.865x` | `7.858x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003792` (3.82 MiB) | `458091` (447.35 KiB) | `0.999x` | `8.740x` | `8.732x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `271928` (265.55 KiB) | `0.999x` | `14.724x` | `14.710x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003776` (3.82 MiB) | `232256` (226.81 KiB) | `0.999x` | `17.239x` | `17.222x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003776` (3.82 MiB) | `373505` (364.75 KiB) | `0.999x` | `10.719x` | `10.709x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003776` (3.82 MiB) | `373262` (364.51 KiB) | `0.999x` | `10.726x` | `10.716x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `282917` (276.29 KiB) | `0.999x` | `14.152x` | `14.138x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003776` (3.82 MiB) | `245304` (239.55 KiB) | `0.999x` | `16.322x` | `16.306x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `328450` (320.75 KiB) | `0.999x` | `12.190x` | `12.178x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `3354477` (3.20 MiB) | `1041552` (1017.14 KiB) | `394617` (385.37 KiB) | `3.221x` | `2.639x` | `8.501x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `219151` (214.01 KiB) | `0.999x` | `18.269x` | `18.252x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003771` (3.82 MiB) | `217313` (212.22 KiB) | `0.999x` | `18.424x` | `18.407x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `364920` (356.37 KiB) | `0.999x` | `10.972x` | `10.961x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `3767530` (3.59 MiB) | `861305` (841.12 KiB) | `237150` (231.59 KiB) | `4.374x` | `3.632x` | `15.887x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `204962` (200.16 KiB) | `0.999x` | `19.534x` | `19.516x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003772` (3.82 MiB) | `205206` (200.40 KiB) | `0.999x` | `19.511x` | `19.493x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003771` (3.82 MiB) | `218416` (213.30 KiB) | `0.999x` | `18.331x` | `18.314x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003770` (3.82 MiB) | `215370` (210.32 KiB) | `0.999x` | `18.590x` | `18.573x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `81583` (79.67 KiB) | `282522` (275.90 KiB) | `57018` (55.68 KiB) | `0.289x` | `4.955x` | `1.431x` | `81583` (79.67 KiB) |
| `Params` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `0` (0 B) | `81484` (79.57 KiB) | `6953` (6.79 KiB) | `0.000x` | `11.719x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003813` (3.82 MiB) | `560292` (547.16 KiB) | `0.999x` | `7.146x` | `7.139x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003813` (3.82 MiB) | `526699` (514.35 KiB) | `0.999x` | `7.602x` | `7.594x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `293110` (286.24 KiB) | `0.999x` | `13.660x` | `13.647x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `3528017` (3.36 MiB) | `2960815` (2.82 MiB) | `993688` (970.40 KiB) | `1.192x` | `2.980x` | `3.550x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003776` (3.82 MiB) | `220295` (215.13 KiB) | `0.999x` | `18.175x` | `18.157x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003807` (3.82 MiB) | `505810` (493.96 KiB) | `0.999x` | `7.916x` | `7.908x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003810` (3.82 MiB) | `516292` (504.19 KiB) | `0.999x` | `7.755x` | `7.748x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003817` (3.82 MiB) | `551436` (538.51 KiB) | `0.999x` | `7.261x` | `7.254x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003773` (3.82 MiB) | `285650` (278.96 KiB) | `0.999x` | `14.016x` | `14.003x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `8004791` (7.63 MiB) | `4229373` (4.03 MiB) | `0.999x` | `1.893x` | `1.892x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003777` (3.82 MiB) | `287282` (280.55 KiB) | `0.999x` | `13.937x` | `13.924x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `269714` (263.39 KiB) | `0.999x` | `14.845x` | `14.831x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003773` (3.82 MiB) | `304813` (297.67 KiB) | `0.999x` | `13.135x` | `13.123x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003773` (3.82 MiB) | `204186` (199.40 KiB) | `0.999x` | `19.608x` | `19.590x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:90` | `1000000` | `13587860` (12.96 MiB) | `142671` (139.33 KiB) | `27008` (26.38 KiB) | `95.239x` | `5.283x` | `503.105x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003773` (3.82 MiB) | `205264` (200.45 KiB) | `0.999x` | `19.505x` | `19.487x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003773` (3.82 MiB) | `256469` (250.46 KiB) | `0.999x` | `15.611x` | `15.596x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003773` (3.82 MiB) | `206625` (201.78 KiB) | `0.999x` | `19.377x` | `19.359x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003777` (3.82 MiB) | `235358` (229.84 KiB) | `0.999x` | `17.011x` | `16.995x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `8004787` (7.63 MiB) | `1151557` (1.10 MiB) | `0.999x` | `6.951x` | `6.947x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:144` | `1000000` | `27797671` (26.51 MiB) | `21047560` (20.07 MiB) | `6750085` (6.44 MiB) | `1.321x` | `3.118x` | `4.118x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003832` (3.82 MiB) | `3690256` (3.52 MiB) | `0.999x` | `1.085x` | `1.084x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003772` (3.82 MiB) | `204142` (199.36 KiB) | `0.999x` | `19.613x` | `19.594x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003772` (3.82 MiB) | `204142` (199.36 KiB) | `0.999x` | `19.613x` | `19.594x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003772` (3.82 MiB) | `204142` (199.36 KiB) | `0.999x` | `19.613x` | `19.594x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003780` (3.82 MiB) | `295634` (288.71 KiB) | `0.999x` | `13.543x` | `13.530x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003772` (3.82 MiB) | `204142` (199.36 KiB) | `0.999x` | `19.613x` | `19.594x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `1000000` (976.56 KiB) | `207120` (202.27 KiB) | `74477` (72.73 KiB) | `4.828x` | `2.781x` | `13.427x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `8004794` (7.63 MiB) | `4284625` (4.09 MiB) | `0.999x` | `1.868x` | `1.867x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003773` (3.82 MiB) | `333962` (326.13 KiB) | `0.999x` | `11.989x` | `11.977x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `302563` (295.47 KiB) | `0.999x` | `13.233x` | `13.220x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `318268` (310.81 KiB) | `0.999x` | `12.580x` | `12.568x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `385543` (376.51 KiB) | `0.999x` | `10.385x` | `10.375x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003778` (3.82 MiB) | `381736` (372.79 KiB) | `0.999x` | `10.488x` | `10.478x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003834` (3.82 MiB) | `707154` (690.58 KiB) | `0.999x` | `5.662x` | `5.656x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003778` (3.82 MiB) | `295204` (288.29 KiB) | `0.999x` | `13.563x` | `13.550x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `204144` (199.36 KiB) | `0.999x` | `19.612x` | `19.594x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003776` (3.82 MiB) | `258387` (252.33 KiB) | `0.999x` | `15.495x` | `15.481x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `2001192` (1.91 MiB) | `335190` (327.33 KiB) | `90246` (88.13 KiB) | `5.970x` | `3.714x` | `22.175x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `3325142` (3.17 MiB) | `959922` (937.42 KiB) | `254070` (248.12 KiB) | `3.464x` | `3.778x` | `13.088x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `0` (0 B) | `81484` (79.57 KiB) | `6953` (6.79 KiB) | `0.000x` | `11.719x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `0` (0 B) | `81484` (79.57 KiB) | `6953` (6.79 KiB) | `0.000x` | `11.719x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003772` (3.82 MiB) | `204142` (199.36 KiB) | `0.999x` | `19.613x` | `19.594x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003779` (3.82 MiB) | `269094` (262.79 KiB) | `0.999x` | `14.879x` | `14.865x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003777` (3.82 MiB) | `327084` (319.42 KiB) | `0.999x` | `12.241x` | `12.229x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003807` (3.82 MiB) | `564424` (551.20 KiB) | `0.999x` | `7.094x` | `7.087x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003833` (3.82 MiB) | `1723633` (1.64 MiB) | `0.999x` | `2.323x` | `2.321x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003833` (3.82 MiB) | `1282061` (1.22 MiB) | `0.999x` | `3.123x` | `3.120x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003829` (3.82 MiB) | `809298` (790.33 KiB) | `0.999x` | `4.947x` | `4.943x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `204525` (199.73 KiB) | `0.999x` | `19.576x` | `19.558x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `1024` (1.00 KiB) | `86815` (84.78 KiB) | `9675` (9.45 KiB) | `0.012x` | `8.973x` | `0.106x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `8004732` (7.63 MiB) | `405100` (395.61 KiB) | `0.999x` | `19.760x` | `19.748x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `0` (0 B) | `81484` (79.57 KiB) | `6953` (6.79 KiB) | `0.000x` | `11.719x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `3000000` (2.86 MiB) | `86341` (84.32 KiB) | `10365` (10.12 KiB) | `34.746x` | `8.330x` | `289.436x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003772` (3.82 MiB) | `204142` (199.36 KiB) | `0.999x` | `19.613x` | `19.594x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `58030` (56.67 KiB) | `301670` (294.60 KiB) | `48050` (46.92 KiB) | `0.192x` | `6.278x` | `1.208x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `22051` (21.53 KiB) | `197850` (193.21 KiB) | `38908` (38.00 KiB) | `0.111x` | `5.085x` | `0.567x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `25445` (24.85 KiB) | `209607` (204.69 KiB) | `44895` (43.84 KiB) | `0.121x` | `4.669x` | `0.567x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `48191` (47.06 KiB) | `232248` (226.80 KiB) | `34195` (33.39 KiB) | `0.207x` | `6.792x` | `1.409x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `49433` (48.27 KiB) | `268624` (262.33 KiB) | `52306` (51.08 KiB) | `0.184x` | `5.136x` | `0.945x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `16873` (16.48 KiB) | `214018` (209.00 KiB) | `39024` (38.11 KiB) | `0.079x` | `5.484x` | `0.432x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `91870` (89.72 KiB) | `328350` (320.65 KiB) | `68514` (66.91 KiB) | `0.280x` | `4.792x` | `1.341x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `13001` (12.70 KiB) | `149219` (145.72 KiB) | `27629` (26.98 KiB) | `0.087x` | `5.401x` | `0.471x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `28101` (27.44 KiB) | `189334` (184.90 KiB) | `31094` (30.37 KiB) | `0.148x` | `6.089x` | `0.904x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:60` | `1000000` | `45607` (44.54 KiB) | `261434` (255.31 KiB) | `50506` (49.32 KiB) | `0.174x` | `5.176x` | `0.903x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003771` (3.82 MiB) | `213536` (208.53 KiB) | `0.999x` | `18.750x` | `18.732x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `8004787` (7.63 MiB) | `3638740` (3.47 MiB) | `0.999x` | `2.200x` | `2.199x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `8000000` (7.63 MiB) | `8004794` (7.63 MiB) | `4384068` (4.18 MiB) | `0.999x` | `1.826x` | `1.825x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:60` | `1000000` | `4000000` (3.81 MiB) | `4003774` (3.82 MiB) | `204570` (199.78 KiB) | `0.999x` | `19.572x` | `19.553x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00000.parquet`: `34156` rows, `3808749` file bytes (3.63 MiB), `26531780` physical bytes (25.30 MiB), `17164575` encoded bytes (16.37 MiB), `3776830` compressed data bytes (3.60 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00001.parquet`: `34461` rows, `3779622` file bytes (3.60 MiB), `26514512` physical bytes (25.29 MiB), `17097277` encoded bytes (16.31 MiB), `3747309` compressed data bytes (3.57 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00002.parquet`: `33968` rows, `3785405` file bytes (3.61 MiB), `26578217` physical bytes (25.35 MiB), `17098159` encoded bytes (16.31 MiB), `3752297` compressed data bytes (3.58 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00003.parquet`: `34071` rows, `3717149` file bytes (3.54 MiB), `26373783` physical bytes (25.15 MiB), `16949833` encoded bytes (16.16 MiB), `3685311` compressed data bytes (3.51 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00004.parquet`: `34141` rows, `3761496` file bytes (3.59 MiB), `26557328` physical bytes (25.33 MiB), `17066528` encoded bytes (16.28 MiB), `3728386` compressed data bytes (3.56 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00005.parquet`: `34700` rows, `3762546` file bytes (3.59 MiB), `26771932` physical bytes (25.53 MiB), `17220330` encoded bytes (16.42 MiB), `3730236` compressed data bytes (3.56 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00006.parquet`: `34507` rows, `3772365` file bytes (3.60 MiB), `26520759` physical bytes (25.29 MiB), `17114055` encoded bytes (16.32 MiB), `3739842` compressed data bytes (3.57 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00007.parquet`: `34255` rows, `3749051` file bytes (3.58 MiB), `26594135` physical bytes (25.36 MiB), `17060297` encoded bytes (16.27 MiB), `3716467` compressed data bytes (3.54 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00008.parquet`: `34251` rows, `3740696` file bytes (3.57 MiB), `26574996` physical bytes (25.34 MiB), `17117926` encoded bytes (16.32 MiB), `3708194` compressed data bytes (3.54 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00009.parquet`: `34148` rows, `3686328` file bytes (3.52 MiB), `26178365` physical bytes (24.97 MiB), `16924977` encoded bytes (16.14 MiB), `3653961` compressed data bytes (3.48 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00010.parquet`: `34464` rows, `3772209` file bytes (3.60 MiB), `26716815` physical bytes (25.48 MiB), `17180945` encoded bytes (16.39 MiB), `3740108` compressed data bytes (3.57 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00011.parquet`: `34560` rows, `3695168` file bytes (3.52 MiB), `26606845` physical bytes (25.37 MiB), `17084713` encoded bytes (16.29 MiB), `3662688` compressed data bytes (3.49 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00012.parquet`: `34147` rows, `3671433` file bytes (3.50 MiB), `26176468` physical bytes (24.96 MiB), `16900102` encoded bytes (16.12 MiB), `3639442` compressed data bytes (3.47 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00013.parquet`: `34042` rows, `3713117` file bytes (3.54 MiB), `26120688` physical bytes (24.91 MiB), `16921305` encoded bytes (16.14 MiB), `3680965` compressed data bytes (3.51 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00014.parquet`: `34752` rows, `3930652` file bytes (3.75 MiB), `26425444` physical bytes (25.20 MiB), `17350349` encoded bytes (16.55 MiB), `3898325` compressed data bytes (3.72 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00015.parquet`: `34274` rows, `4304516` file bytes (4.11 MiB), `23431745` physical bytes (22.35 MiB), `17355572` encoded bytes (16.55 MiB), `4274130` compressed data bytes (4.08 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00016.parquet`: `33635` rows, `4357346` file bytes (4.16 MiB), `23704234` physical bytes (22.61 MiB), `17088773` encoded bytes (16.30 MiB), `4327130` compressed data bytes (4.13 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00017.parquet`: `33118` rows, `4432806` file bytes (4.23 MiB), `21373789` physical bytes (20.38 MiB), `17832419` encoded bytes (17.01 MiB), `4402473` compressed data bytes (4.20 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00018.parquet`: `33397` rows, `4538245` file bytes (4.33 MiB), `21556966` physical bytes (20.56 MiB), `18379310` encoded bytes (17.53 MiB), `4507463` compressed data bytes (4.30 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00019.parquet`: `32848` rows, `4636974` file bytes (4.42 MiB), `21469410` physical bytes (20.47 MiB), `18287944` encoded bytes (17.44 MiB), `4606613` compressed data bytes (4.39 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00020.parquet`: `33361` rows, `4471940` file bytes (4.26 MiB), `21182822` physical bytes (20.20 MiB), `18169273` encoded bytes (17.33 MiB), `4441216` compressed data bytes (4.24 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00021.parquet`: `32809` rows, `4538156` file bytes (4.33 MiB), `21151065` physical bytes (20.17 MiB), `18037118` encoded bytes (17.20 MiB), `4507437` compressed data bytes (4.30 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00022.parquet`: `33354` rows, `4513154` file bytes (4.30 MiB), `21332938` physical bytes (20.34 MiB), `18230477` encoded bytes (17.39 MiB), `4482343` compressed data bytes (4.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00023.parquet`: `33345` rows, `4503856` file bytes (4.30 MiB), `21336686` physical bytes (20.35 MiB), `18160679` encoded bytes (17.32 MiB), `4473451` compressed data bytes (4.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00024.parquet`: `32834` rows, `4562956` file bytes (4.35 MiB), `21339621` physical bytes (20.35 MiB), `18287162` encoded bytes (17.44 MiB), `4532214` compressed data bytes (4.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00025.parquet`: `33086` rows, `4504056` file bytes (4.30 MiB), `21144902` physical bytes (20.17 MiB), `18039382` encoded bytes (17.20 MiB), `4473799` compressed data bytes (4.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00026.parquet`: `33306` rows, `4489997` file bytes (4.28 MiB), `21214134` physical bytes (20.23 MiB), `18095296` encoded bytes (17.26 MiB), `4459396` compressed data bytes (4.25 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00027.parquet`: `33264` rows, `4430183` file bytes (4.22 MiB), `21240215` physical bytes (20.26 MiB), `18014155` encoded bytes (17.18 MiB), `4399631` compressed data bytes (4.20 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00028.parquet`: `33270` rows, `4447927` file bytes (4.24 MiB), `21054025` physical bytes (20.08 MiB), `17985709` encoded bytes (17.15 MiB), `4417704` compressed data bytes (4.21 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-plain/part-00029.parquet`: `19476` rows, `2668092` file bytes (2.54 MiB), `12624005` physical bytes (12.04 MiB), `10708331` encoded bytes (10.21 MiB), `2638698` compressed data bytes (2.52 MiB)
