# ClickBench Parquet Experiment

- Started: `2026-07-03T19:41:50-04:00`
- Write elapsed: `11.607s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `819416584` (781.46 MiB)
- Compressed column data bytes after codec compression: `137855747` (131.47 MiB)
- Parquet file bytes: `138833336` (132.40 MiB)
- Physical/encoded ratio: `0.869x`
- Encoded/compressed-data ratio: `5.944x`
- Physical/compressed-data ratio: `5.168x`
- Physical/parquet-file ratio: `5.131x`
- Files: `31`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `plain`
- Date encoding: `rle-dict`
- Timestamp encoding: `rle-dict`
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
- Elapsed: `7.014s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004955` (7.63 MiB) | `8005388` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003897` (3.82 MiB) | `262178` (256.03 KiB) | `0.999x` | `15.272x` | `15.257x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:498` | `1000000` | `138409995` (132.00 MiB) | `142876268` (136.26 MiB) | `20906997` (19.94 MiB) | `0.969x` | `6.834x` | `6.620x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `204453` (199.66 KiB) | `0.999x` | `19.583x` | `19.564x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `7358034` (7.02 MiB) | `5536365` (5.28 MiB) | `1.087x` | `1.329x` | `1.445x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `5388` (5.26 KiB) | `5636` (5.50 KiB) | `742.390x` | `0.956x` | `709.723x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `204475` (199.68 KiB) | `0.999x` | `19.581x` | `19.562x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003960` (3.82 MiB) | `719526` (702.66 KiB) | `0.999x` | `5.565x` | `5.559x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `396994` (387.69 KiB) | `0.999x` | `10.086x` | `10.076x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004954` (7.63 MiB) | `1085033` (1.03 MiB) | `0.999x` | `7.378x` | `7.373x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `204273` (199.49 KiB) | `0.999x` | `19.601x` | `19.582x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `318288` (310.83 KiB) | `0.999x` | `12.579x` | `12.567x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `342438` (334.41 KiB) | `0.999x` | `11.692x` | `11.681x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:340` | `1000000` | `88562192` (84.46 MiB) | `92649058` (88.36 MiB) | `20460071` (19.51 MiB) | `0.956x` | `4.528x` | `4.329x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:294` | `1000000` | `79583339` (75.90 MiB) | `83648518` (79.77 MiB) | `19057579` (18.17 MiB) | `0.951x` | `4.389x` | `4.176x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003933` (3.82 MiB) | `492078` (480.54 KiB) | `0.999x` | `8.137x` | `8.129x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003931` (3.82 MiB) | `509175` (497.24 KiB) | `0.999x` | `7.864x` | `7.856x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003916` (3.82 MiB) | `458225` (447.49 KiB) | `0.999x` | `8.738x` | `8.729x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `271997` (265.62 KiB) | `0.999x` | `14.720x` | `14.706x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `232438` (226.99 KiB) | `0.999x` | `17.226x` | `17.209x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `373638` (364.88 KiB) | `0.999x` | `10.716x` | `10.706x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003906` (3.82 MiB) | `373312` (364.56 KiB) | `0.999x` | `10.725x` | `10.715x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `283130` (276.49 KiB) | `0.999x` | `14.142x` | `14.128x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `245463` (239.71 KiB) | `0.999x` | `16.312x` | `16.296x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `328526` (320.83 KiB) | `0.999x` | `12.187x` | `12.176x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `3354477` (3.20 MiB) | `7358157` (7.02 MiB) | `536776` (524.20 KiB) | `0.456x` | `13.708x` | `6.249x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003896` (3.82 MiB) | `219254` (214.12 KiB) | `0.999x` | `18.261x` | `18.244x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `217438` (212.34 KiB) | `0.999x` | `18.414x` | `18.396x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003905` (3.82 MiB) | `365219` (356.66 KiB) | `0.999x` | `10.963x` | `10.952x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `3767530` (3.59 MiB) | `7771204` (7.41 MiB) | `467339` (456.39 KiB) | `0.485x` | `16.629x` | `8.062x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `205101` (200.29 KiB) | `0.999x` | `19.522x` | `19.503x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003896` (3.82 MiB) | `205058` (200.25 KiB) | `0.999x` | `19.526x` | `19.507x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `218583` (213.46 KiB) | `0.999x` | `18.318x` | `18.300x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `215515` (210.46 KiB) | `0.999x` | `18.578x` | `18.560x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `81583` (79.67 KiB) | `4085177` (3.90 MiB) | `220066` (214.91 KiB) | `0.020x` | `18.563x` | `0.371x` | `81583` (79.67 KiB) |
| `Params` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `0` (0 B) | `4002411` (3.82 MiB) | `202785` (198.03 KiB) | `0.000x` | `19.737x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003934` (3.82 MiB) | `561206` (548.05 KiB) | `0.999x` | `7.135x` | `7.128x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003932` (3.82 MiB) | `526971` (514.62 KiB) | `0.999x` | `7.598x` | `7.591x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003896` (3.82 MiB) | `293342` (286.47 KiB) | `0.999x` | `13.649x` | `13.636x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `3528017` (3.36 MiB) | `7535671` (7.19 MiB) | `1098213` (1.05 MiB) | `0.468x` | `6.862x` | `3.213x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `220405` (215.24 KiB) | `0.999x` | `18.166x` | `18.148x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003934` (3.82 MiB) | `505887` (494.03 KiB) | `0.999x` | `7.915x` | `7.907x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003933` (3.82 MiB) | `516922` (504.81 KiB) | `0.999x` | `7.746x` | `7.738x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003944` (3.82 MiB) | `552131` (539.19 KiB) | `0.999x` | `7.252x` | `7.245x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `285966` (279.26 KiB) | `0.999x` | `14.001x` | `13.988x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `7235642` (6.90 MiB) | `5464725` (5.21 MiB) | `1.106x` | `1.324x` | `1.464x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `287348` (280.61 KiB) | `0.999x` | `13.934x` | `13.920x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003897` (3.82 MiB) | `269806` (263.48 KiB) | `0.999x` | `14.840x` | `14.825x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `304905` (297.76 KiB) | `0.999x` | `13.132x` | `13.119x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `204313` (199.52 KiB) | `0.999x` | `19.597x` | `19.578x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:92` | `1000000` | `13587860` (12.96 MiB) | `17595579` (16.78 MiB) | `928271` (906.51 KiB) | `0.772x` | `18.955x` | `14.638x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `205316` (200.50 KiB) | `0.999x` | `19.501x` | `19.482x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `256590` (250.58 KiB) | `0.999x` | `15.604x` | `15.589x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `206778` (201.93 KiB) | `0.999x` | `19.363x` | `19.344x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `235488` (229.97 KiB) | `0.999x` | `17.003x` | `16.986x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004953` (7.63 MiB) | `1151714` (1.10 MiB) | `0.999x` | `6.950x` | `6.946x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:144` | `1000000` | `27797671` (26.51 MiB) | `31860754` (30.38 MiB) | `7040512` (6.71 MiB) | `0.872x` | `4.525x` | `3.948x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003959` (3.82 MiB) | `3688293` (3.52 MiB) | `0.999x` | `1.086x` | `1.085x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `204273` (199.49 KiB) | `0.999x` | `19.601x` | `19.582x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `204273` (199.49 KiB) | `0.999x` | `19.601x` | `19.582x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `204273` (199.49 KiB) | `0.999x` | `19.601x` | `19.582x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `295772` (288.84 KiB) | `0.999x` | `13.537x` | `13.524x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `204273` (199.49 KiB) | `0.999x` | `19.601x` | `19.582x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `1000000` (976.56 KiB) | `5003162` (4.77 MiB) | `294490` (287.59 KiB) | `0.200x` | `16.989x` | `3.396x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `7350831` (7.01 MiB) | `5532484` (5.28 MiB) | `1.088x` | `1.329x` | `1.446x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `334116` (326.29 KiB) | `0.999x` | `11.984x` | `11.972x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `302680` (295.59 KiB) | `0.999x` | `13.228x` | `13.215x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `318380` (310.92 KiB) | `0.999x` | `12.576x` | `12.564x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `385632` (376.59 KiB) | `0.999x` | `10.383x` | `10.373x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `381912` (372.96 KiB) | `0.999x` | `10.484x` | `10.474x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003961` (3.82 MiB) | `707806` (691.22 KiB) | `0.999x` | `5.657x` | `5.651x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `295404` (288.48 KiB) | `0.999x` | `13.554x` | `13.541x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `204273` (199.49 KiB) | `0.999x` | `19.601x` | `19.582x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003905` (3.82 MiB) | `258505` (252.45 KiB) | `0.999x` | `15.489x` | `15.474x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `2001192` (1.91 MiB) | `6004719` (5.73 MiB) | `321343` (313.81 KiB) | `0.333x` | `18.686x` | `6.228x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `3325142` (3.17 MiB) | `7328818` (6.99 MiB) | `457245` (446.53 KiB) | `0.454x` | `16.028x` | `7.272x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `0` (0 B) | `4002411` (3.82 MiB) | `202785` (198.03 KiB) | `0.000x` | `19.737x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `0` (0 B) | `4002411` (3.82 MiB) | `202785` (198.03 KiB) | `0.000x` | `19.737x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `204273` (199.49 KiB) | `0.999x` | `19.601x` | `19.582x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003905` (3.82 MiB) | `269336` (263.02 KiB) | `0.999x` | `14.866x` | `14.851x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `327322` (319.65 KiB) | `0.999x` | `12.232x` | `12.220x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003936` (3.82 MiB) | `565212` (551.96 KiB) | `0.999x` | `7.084x` | `7.077x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003964` (3.82 MiB) | `1725011` (1.65 MiB) | `0.999x` | `2.321x` | `2.319x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003961` (3.82 MiB) | `1283282` (1.22 MiB) | `0.999x` | `3.120x` | `3.117x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003956` (3.82 MiB) | `810534` (791.54 KiB) | `0.999x` | `4.940x` | `4.935x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `204669` (199.87 KiB) | `0.999x` | `19.563x` | `19.544x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `1024` (1.00 KiB) | `4004203` (3.82 MiB) | `204355` (199.57 KiB) | `0.000x` | `19.594x` | `0.005x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004891` (7.63 MiB) | `405270` (395.77 KiB) | `0.999x` | `19.752x` | `19.740x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `0` (0 B) | `4002411` (3.82 MiB) | `202785` (198.03 KiB) | `0.000x` | `19.737x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `3000000` (2.86 MiB) | `7003650` (6.68 MiB) | `354387` (346.08 KiB) | `0.428x` | `19.763x` | `8.465x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `204273` (199.49 KiB) | `0.999x` | `19.601x` | `19.582x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `58030` (56.67 KiB) | `4062882` (3.87 MiB) | `218579` (213.46 KiB) | `0.014x` | `18.588x` | `0.265x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `22051` (21.53 KiB) | `4025887` (3.84 MiB) | `214730` (209.70 KiB) | `0.005x` | `18.749x` | `0.103x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `25445` (24.85 KiB) | `4029315` (3.84 MiB) | `219210` (214.07 KiB) | `0.006x` | `18.381x` | `0.116x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `48191` (47.06 KiB) | `4052087` (3.86 MiB) | `214450` (209.42 KiB) | `0.012x` | `18.895x` | `0.225x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `49433` (48.27 KiB) | `4053503` (3.87 MiB) | `219151` (214.01 KiB) | `0.012x` | `18.496x` | `0.226x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `16873` (16.48 KiB) | `4020198` (3.83 MiB) | `214379` (209.35 KiB) | `0.004x` | `18.753x` | `0.079x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `91870` (89.72 KiB) | `4097949` (3.91 MiB) | `230323` (224.92 KiB) | `0.022x` | `17.792x` | `0.399x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `13001` (12.70 KiB) | `4016965` (3.83 MiB) | `212439` (207.46 KiB) | `0.003x` | `18.909x` | `0.061x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `28101` (27.44 KiB) | `4033121` (3.85 MiB) | `214670` (209.64 KiB) | `0.007x` | `18.788x` | `0.131x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `45607` (44.54 KiB) | `4048573` (3.86 MiB) | `221867` (216.67 KiB) | `0.011x` | `18.248x` | `0.206x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `213669` (208.66 KiB) | `0.999x` | `18.739x` | `18.721x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004955` (7.63 MiB) | `3641069` (3.47 MiB) | `0.999x` | `2.199x` | `2.197x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004951` (7.63 MiB) | `4387178` (4.18 MiB) | `0.999x` | `1.825x` | `1.823x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003897` (3.82 MiB) | `204708` (199.91 KiB) | `0.999x` | `19.559x` | `19.540x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00000.parquet`: `32763` rows, `4355820` file bytes (4.15 MiB), `25447536` physical bytes (24.27 MiB), `28806734` encoded bytes (27.47 MiB), `4323777` compressed data bytes (4.12 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00001.parquet`: `32786` rows, `4306976` file bytes (4.11 MiB), `25225716` physical bytes (24.06 MiB), `28594272` encoded bytes (27.27 MiB), `4274584` compressed data bytes (4.08 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00002.parquet`: `32625` rows, `4330195` file bytes (4.13 MiB), `25504450` physical bytes (24.32 MiB), `28859955` encoded bytes (27.52 MiB), `4297241` compressed data bytes (4.10 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00003.parquet`: `32808` rows, `4327933` file bytes (4.13 MiB), `25469429` physical bytes (24.29 MiB), `28839485` encoded bytes (27.50 MiB), `4295379` compressed data bytes (4.10 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00004.parquet`: `33080` rows, `4356795` file bytes (4.15 MiB), `25665598` physical bytes (24.48 MiB), `29062076` encoded bytes (27.72 MiB), `4323576` compressed data bytes (4.12 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00005.parquet`: `33284` rows, `4299184` file bytes (4.10 MiB), `25692437` physical bytes (24.50 MiB), `29112555` encoded bytes (27.76 MiB), `4266556` compressed data bytes (4.07 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00006.parquet`: `33064` rows, `4325048` file bytes (4.12 MiB), `25439185` physical bytes (24.26 MiB), `28831731` encoded bytes (27.50 MiB), `4292697` compressed data bytes (4.09 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00007.parquet`: `32903` rows, `4308233` file bytes (4.11 MiB), `25488798` physical bytes (24.31 MiB), `28874359` encoded bytes (27.54 MiB), `4275563` compressed data bytes (4.08 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00008.parquet`: `32901` rows, `4304373` file bytes (4.10 MiB), `25494357` physical bytes (24.31 MiB), `28868898` encoded bytes (27.53 MiB), `4271321` compressed data bytes (4.07 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00009.parquet`: `33003` rows, `4286692` file bytes (4.09 MiB), `25511004` physical bytes (24.33 MiB), `28897791` encoded bytes (27.56 MiB), `4254247` compressed data bytes (4.06 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00010.parquet`: `33014` rows, `4311646` file bytes (4.11 MiB), `25464601` physical bytes (24.28 MiB), `28848671` encoded bytes (27.51 MiB), `4279433` compressed data bytes (4.08 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00011.parquet`: `33145` rows, `4278794` file bytes (4.08 MiB), `25475575` physical bytes (24.30 MiB), `28876569` encoded bytes (27.54 MiB), `4246253` compressed data bytes (4.05 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00012.parquet`: `33095` rows, `4275535` file bytes (4.08 MiB), `25511048` physical bytes (24.33 MiB), `28911627` encoded bytes (27.57 MiB), `4243112` compressed data bytes (4.05 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00013.parquet`: `33096` rows, `4312777` file bytes (4.11 MiB), `25484659` physical bytes (24.30 MiB), `28880583` encoded bytes (27.54 MiB), `4280656` compressed data bytes (4.08 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00014.parquet`: `32775` rows, `4235345` file bytes (4.04 MiB), `25148235` physical bytes (23.98 MiB), `28505249` encoded bytes (27.18 MiB), `4202861` compressed data bytes (4.01 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00015.parquet`: `32946` rows, `4581012` file bytes (4.37 MiB), `23860901` physical bytes (22.76 MiB), `27380885` encoded bytes (26.11 MiB), `4549677` compressed data bytes (4.34 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00016.parquet`: `32889` rows, `4729591` file bytes (4.51 MiB), `22646275` physical bytes (21.60 MiB), `26242747` encoded bytes (25.03 MiB), `4699037` compressed data bytes (4.48 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00017.parquet`: `31440` rows, `4703254` file bytes (4.49 MiB), `22002354` physical bytes (20.98 MiB), `25479903` encoded bytes (24.30 MiB), `4672891` compressed data bytes (4.46 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00018.parquet`: `32109` rows, `4620924` file bytes (4.41 MiB), `20274264` physical bytes (19.34 MiB), `23866032` encoded bytes (22.76 MiB), `4590611` compressed data bytes (4.38 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00019.parquet`: `32316` rows, `4766163` file bytes (4.55 MiB), `20913714` physical bytes (19.94 MiB), `24529929` encoded bytes (23.39 MiB), `4735304` compressed data bytes (4.52 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00020.parquet`: `31844` rows, `4872299` file bytes (4.65 MiB), `20763043` physical bytes (19.80 MiB), `24319299` encoded bytes (23.19 MiB), `4841810` compressed data bytes (4.62 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00021.parquet`: `32126` rows, `4740602` file bytes (4.52 MiB), `20566985` physical bytes (19.61 MiB), `24157483` encoded bytes (23.04 MiB), `4709739` compressed data bytes (4.49 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00022.parquet`: `32145` rows, `4794824` file bytes (4.57 MiB), `20682367` physical bytes (19.72 MiB), `24279673` encoded bytes (23.15 MiB), `4764069` compressed data bytes (4.54 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00023.parquet`: `32637` rows, `4776056` file bytes (4.55 MiB), `20858071` physical bytes (19.89 MiB), `24509253` encoded bytes (23.37 MiB), `4745285` compressed data bytes (4.53 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00024.parquet`: `32597` rows, `4783321` file bytes (4.56 MiB), `20848567` physical bytes (19.88 MiB), `24491091` encoded bytes (23.36 MiB), `4752788` compressed data bytes (4.53 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00025.parquet`: `32000` rows, `4852121` file bytes (4.63 MiB), `20888231` physical bytes (19.92 MiB), `24469600` encoded bytes (23.34 MiB), `4821283` compressed data bytes (4.60 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00026.parquet`: `32293` rows, `4722620` file bytes (4.50 MiB), `20564416` physical bytes (19.61 MiB), `24175732` encoded bytes (23.06 MiB), `4692336` compressed data bytes (4.47 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00027.parquet`: `32313` rows, `4741920` file bytes (4.52 MiB), `20593177` physical bytes (19.64 MiB), `24210077` encoded bytes (23.09 MiB), `4711221` compressed data bytes (4.49 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00028.parquet`: `32312` rows, `4678956` file bytes (4.46 MiB), `20637996` physical bytes (19.68 MiB), `24250297` encoded bytes (23.13 MiB), `4648315` compressed data bytes (4.43 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00029.parquet`: `32498` rows, `4717278` file bytes (4.50 MiB), `20583949` physical bytes (19.63 MiB), `24217353` encoded bytes (23.10 MiB), `4686928` compressed data bytes (4.47 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00030.parquet`: `21193` rows, `3137049` file bytes (2.99 MiB), `13691686` physical bytes (13.06 MiB), `16066675` encoded bytes (15.32 MiB), `3107197` compressed data bytes (2.96 MiB)
