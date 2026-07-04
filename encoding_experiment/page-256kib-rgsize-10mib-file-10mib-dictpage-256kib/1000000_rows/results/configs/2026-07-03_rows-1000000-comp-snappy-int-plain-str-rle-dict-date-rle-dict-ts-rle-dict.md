# ClickBench Parquet Experiment

- Started: `2026-07-03T19:41:50-04:00`
- Write elapsed: `11.541s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `469402357` (447.66 MiB)
- Compressed column data bytes after codec compression: `111042946` (105.90 MiB)
- Parquet file bytes: `111977642` (106.79 MiB)
- Physical/encoded ratio: `1.518x`
- Encoded/compressed-data ratio: `4.227x`
- Physical/compressed-data ratio: `6.416x`
- Physical/parquet-file ratio: `6.362x`
- Files: `31`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `rle-dict`
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
- Elapsed: `6.8s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004954` (7.63 MiB) | `8005387` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `262157` (256.01 KiB) | `0.999x` | `15.273x` | `15.258x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `138409995` (132.00 MiB) | `28851140` (27.51 MiB) | `10016895` (9.55 MiB) | `4.797x` | `2.880x` | `13.818x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `204445` (199.65 KiB) | `0.999x` | `19.584x` | `19.565x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `7344624` (7.00 MiB) | `5516487` (5.26 MiB) | `1.089x` | `1.331x` | `1.450x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `5390` (5.26 KiB) | `5638` (5.51 KiB) | `742.115x` | `0.956x` | `709.471x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `204467` (199.67 KiB) | `0.999x` | `19.582x` | `19.563x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003954` (3.82 MiB) | `719522` (702.66 KiB) | `0.999x` | `5.565x` | `5.559x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `396952` (387.65 KiB) | `0.999x` | `10.087x` | `10.077x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004952` (7.63 MiB) | `1085045` (1.03 MiB) | `0.999x` | `7.378x` | `7.373x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `204274` (199.49 KiB) | `0.999x` | `19.601x` | `19.582x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `318178` (310.72 KiB) | `0.999x` | `12.584x` | `12.572x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `342417` (334.39 KiB) | `0.999x` | `11.693x` | `11.682x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `88562192` (84.46 MiB) | `44096024` (42.05 MiB) | `16097159` (15.35 MiB) | `2.008x` | `2.739x` | `5.502x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `79583339` (75.90 MiB) | `34412758` (32.82 MiB) | `14810368` (14.12 MiB) | `2.313x` | `2.324x` | `5.373x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003931` (3.82 MiB) | `492045` (480.51 KiB) | `0.999x` | `8.137x` | `8.129x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003934` (3.82 MiB) | `509148` (497.21 KiB) | `0.999x` | `7.864x` | `7.856x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003918` (3.82 MiB) | `458326` (447.58 KiB) | `0.999x` | `8.736x` | `8.727x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `272034` (265.66 KiB) | `0.999x` | `14.718x` | `14.704x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `232309` (226.86 KiB) | `0.999x` | `17.235x` | `17.218x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `373607` (364.85 KiB) | `0.999x` | `10.717x` | `10.706x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `373389` (364.64 KiB) | `0.999x` | `10.723x` | `10.713x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `283100` (276.46 KiB) | `0.999x` | `14.143x` | `14.129x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `245464` (239.71 KiB) | `0.999x` | `16.312x` | `16.296x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `328545` (320.84 KiB) | `0.999x` | `12.187x` | `12.175x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `3354477` (3.20 MiB) | `266945` (260.69 KiB) | `182893` (178.61 KiB) | `12.566x` | `1.460x` | `18.341x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `219269` (214.13 KiB) | `0.999x` | `18.260x` | `18.242x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `217430` (212.33 KiB) | `0.999x` | `18.415x` | `18.397x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `365435` (356.87 KiB) | `0.999x` | `10.957x` | `10.946x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `3767530` (3.59 MiB) | `150758` (147.22 KiB) | `104679` (102.23 KiB) | `24.991x` | `1.440x` | `35.991x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `204938` (200.13 KiB) | `0.999x` | `19.537x` | `19.518x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `205090` (200.28 KiB) | `0.999x` | `19.523x` | `19.504x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `218606` (213.48 KiB) | `0.999x` | `18.316x` | `18.298x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `215515` (210.46 KiB) | `0.999x` | `18.578x` | `18.560x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `81583` (79.67 KiB) | `25292` (24.70 KiB) | `23267` (22.72 KiB) | `3.226x` | `1.087x` | `3.506x` | `81583` (79.67 KiB) |
| `Params` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `0` (0 B) | `4398` (4.29 KiB) | `4646` (4.54 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003937` (3.82 MiB) | `561200` (548.05 KiB) | `0.999x` | `7.135x` | `7.128x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003935` (3.82 MiB) | `526982` (514.63 KiB) | `0.999x` | `7.598x` | `7.590x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `293256` (286.38 KiB) | `0.999x` | `13.653x` | `13.640x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `3528017` (3.36 MiB) | `1643743` (1.57 MiB) | `786559` (768.12 KiB) | `2.146x` | `2.090x` | `4.485x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003897` (3.82 MiB) | `220416` (215.25 KiB) | `0.999x` | `18.165x` | `18.148x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003935` (3.82 MiB) | `505905` (494.05 KiB) | `0.999x` | `7.914x` | `7.907x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003930` (3.82 MiB) | `517073` (504.95 KiB) | `0.999x` | `7.743x` | `7.736x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003941` (3.82 MiB) | `552216` (539.27 KiB) | `0.999x` | `7.251x` | `7.244x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `285811` (279.11 KiB) | `0.999x` | `14.009x` | `13.995x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `7220619` (6.89 MiB) | `5442970` (5.19 MiB) | `1.108x` | `1.327x` | `1.470x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003896` (3.82 MiB) | `287388` (280.65 KiB) | `0.999x` | `13.932x` | `13.918x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `269865` (263.54 KiB) | `0.999x` | `14.837x` | `14.822x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003896` (3.82 MiB) | `305069` (297.92 KiB) | `0.999x` | `13.125x` | `13.112x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `204311` (199.52 KiB) | `0.999x` | `19.597x` | `19.578x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `13587860` (12.96 MiB) | `9160` (8.95 KiB) | `9344` (9.12 KiB) | `1483.391x` | `0.980x` | `1454.180x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `205293` (200.48 KiB) | `0.999x` | `19.503x` | `19.484x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `256589` (250.58 KiB) | `0.999x` | `15.604x` | `15.589x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `206761` (201.92 KiB) | `0.999x` | `19.365x` | `19.346x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003904` (3.82 MiB) | `235482` (229.96 KiB) | `0.999x` | `17.003x` | `16.986x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004952` (7.63 MiB) | `1151769` (1.10 MiB) | `0.999x` | `6.950x` | `6.946x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `27797671` (26.51 MiB) | `21276498` (20.29 MiB) | `6167090` (5.88 MiB) | `1.306x` | `3.450x` | `4.507x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003961` (3.82 MiB) | `3688295` (3.52 MiB) | `0.999x` | `1.086x` | `1.085x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `204274` (199.49 KiB) | `0.999x` | `19.601x` | `19.582x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `204274` (199.49 KiB) | `0.999x` | `19.601x` | `19.582x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `204274` (199.49 KiB) | `0.999x` | `19.601x` | `19.582x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `295720` (288.79 KiB) | `0.999x` | `13.540x` | `13.526x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `204274` (199.49 KiB) | `0.999x` | `19.601x` | `19.582x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `1000000` (976.56 KiB) | `33945` (33.15 KiB) | `32155` (31.40 KiB) | `29.459x` | `1.056x` | `31.099x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `7339519` (7.00 MiB) | `5514562` (5.26 MiB) | `1.090x` | `1.331x` | `1.451x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `334116` (326.29 KiB) | `0.999x` | `11.984x` | `11.972x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `302687` (295.59 KiB) | `0.999x` | `13.228x` | `13.215x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003904` (3.82 MiB) | `318392` (310.93 KiB) | `0.999x` | `12.575x` | `12.563x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `385870` (376.83 KiB) | `0.999x` | `10.376x` | `10.366x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003906` (3.82 MiB) | `382068` (373.11 KiB) | `0.999x` | `10.480x` | `10.469x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003960` (3.82 MiB) | `707499` (690.92 KiB) | `0.999x` | `5.659x` | `5.654x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003905` (3.82 MiB) | `295428` (288.50 KiB) | `0.999x` | `13.553x` | `13.540x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `204271` (199.48 KiB) | `0.999x` | `19.601x` | `19.582x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003904` (3.82 MiB) | `258695` (252.63 KiB) | `0.999x` | `15.477x` | `15.462x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `2001192` (1.91 MiB) | `42688` (41.69 KiB) | `34792` (33.98 KiB) | `46.879x` | `1.227x` | `57.519x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `3325142` (3.17 MiB) | `138093` (134.86 KiB) | `95281` (93.05 KiB) | `24.079x` | `1.449x` | `34.898x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `0` (0 B) | `4398` (4.29 KiB) | `4646` (4.54 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `0` (0 B) | `4398` (4.29 KiB) | `4646` (4.54 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `204274` (199.49 KiB) | `0.999x` | `19.601x` | `19.582x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003909` (3.82 MiB) | `269160` (262.85 KiB) | `0.999x` | `14.876x` | `14.861x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `327504` (319.83 KiB) | `0.999x` | `12.225x` | `12.214x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003938` (3.82 MiB) | `565924` (552.66 KiB) | `0.999x` | `7.075x` | `7.068x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003964` (3.82 MiB) | `1726695` (1.65 MiB) | `0.999x` | `2.319x` | `2.317x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003958` (3.82 MiB) | `1284229` (1.22 MiB) | `0.999x` | `3.118x` | `3.115x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003958` (3.82 MiB) | `810550` (791.55 KiB) | `0.999x` | `4.940x` | `4.935x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `204769` (199.97 KiB) | `0.999x` | `19.553x` | `19.534x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `1024` (1.00 KiB) | `6015` (5.87 KiB) | `6265` (6.12 KiB) | `0.170x` | `0.960x` | `0.163x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004894` (7.63 MiB) | `405272` (395.77 KiB) | `0.999x` | `19.752x` | `19.740x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `0` (0 B) | `4398` (4.29 KiB) | `4646` (4.54 KiB) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `3000000` (2.86 MiB) | `5328` (5.20 KiB) | `5576` (5.45 KiB) | `563.063x` | `0.956x` | `538.020x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `204274` (199.49 KiB) | `0.999x` | `19.601x` | `19.582x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `58030` (56.67 KiB) | `18912` (18.47 KiB) | `18109` (17.68 KiB) | `3.068x` | `1.044x` | `3.204x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `22051` (21.53 KiB) | `16710` (16.32 KiB) | `16168` (15.79 KiB) | `1.320x` | `1.034x` | `1.364x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `25445` (24.85 KiB) | `23129` (22.59 KiB) | `21153` (20.66 KiB) | `1.100x` | `1.093x` | `1.203x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `48191` (47.06 KiB) | `14862` (14.51 KiB) | `13598` (13.28 KiB) | `3.243x` | `1.093x` | `3.544x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `49433` (48.27 KiB) | `22897` (22.36 KiB) | `20590` (20.11 KiB) | `2.159x` | `1.112x` | `2.401x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `16873` (16.48 KiB) | `16808` (16.41 KiB) | `16098` (15.72 KiB) | `1.004x` | `1.044x` | `1.048x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `91870` (89.72 KiB) | `37870` (36.98 KiB) | `32268` (31.51 KiB) | `2.426x` | `1.174x` | `2.847x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `13001` (12.70 KiB) | `14963` (14.61 KiB) | `14789` (14.44 KiB) | `0.869x` | `1.012x` | `0.879x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `28101` (27.44 KiB) | `16776` (16.38 KiB) | `15924` (15.55 KiB) | `1.675x` | `1.054x` | `1.765x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `45607` (44.54 KiB) | `41628` (40.65 KiB) | `29002` (28.32 KiB) | `1.096x` | `1.435x` | `1.573x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `213677` (208.67 KiB) | `0.999x` | `18.738x` | `18.720x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004955` (7.63 MiB) | `3638436` (3.47 MiB) | `0.999x` | `2.200x` | `2.199x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004953` (7.63 MiB) | `4382576` (4.18 MiB) | `0.999x` | `1.827x` | `1.825x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `204726` (199.93 KiB) | `0.999x` | `19.557x` | `19.538x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00000.parquet`: `34636` rows, `3353751` file bytes (3.20 MiB), `26900676` physical bytes (25.65 MiB), `15236313` encoded bytes (14.53 MiB), `3323189` compressed data bytes (3.17 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00001.parquet`: `34775` rows, `3329358` file bytes (3.18 MiB), `26775214` physical bytes (25.53 MiB), `15214194` encoded bytes (14.51 MiB), `3298796` compressed data bytes (3.15 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00002.parquet`: `34618` rows, `3349880` file bytes (3.19 MiB), `27039518` physical bytes (25.79 MiB), `15227998` encoded bytes (14.52 MiB), `3318591` compressed data bytes (3.16 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00003.parquet`: `34765` rows, `3315800` file bytes (3.16 MiB), `26923521` physical bytes (25.68 MiB), `15210349` encoded bytes (14.51 MiB), `3285417` compressed data bytes (3.13 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00004.parquet`: `34787` rows, `3323536` file bytes (3.17 MiB), `27085267` physical bytes (25.83 MiB), `15206071` encoded bytes (14.50 MiB), `3292345` compressed data bytes (3.14 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00005.parquet`: `34984` rows, `3310684` file bytes (3.16 MiB), `26911922` physical bytes (25.67 MiB), `15211935` encoded bytes (14.51 MiB), `3280070` compressed data bytes (3.13 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00006.parquet`: `34880` rows, `3317225` file bytes (3.16 MiB), `26926993` physical bytes (25.68 MiB), `15208777` encoded bytes (14.50 MiB), `3286661` compressed data bytes (3.13 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00007.parquet`: `34782` rows, `3325119` file bytes (3.17 MiB), `26980371` physical bytes (25.73 MiB), `15218018` encoded bytes (14.51 MiB), `3294217` compressed data bytes (3.14 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00008.parquet`: `34881` rows, `3307247` file bytes (3.15 MiB), `27020184` physical bytes (25.77 MiB), `15208322` encoded bytes (14.50 MiB), `3276288` compressed data bytes (3.12 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00009.parquet`: `34981` rows, `3299985` file bytes (3.15 MiB), `26843785` physical bytes (25.60 MiB), `15209621` encoded bytes (14.51 MiB), `3269504` compressed data bytes (3.12 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00010.parquet`: `34727` rows, `3307514` file bytes (3.15 MiB), `26830877` physical bytes (25.59 MiB), `15221920` encoded bytes (14.52 MiB), `3276976` compressed data bytes (3.13 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00011.parquet`: `35098` rows, `3273927` file bytes (3.12 MiB), `27092154` physical bytes (25.84 MiB), `15207885` encoded bytes (14.50 MiB), `3243288` compressed data bytes (3.09 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00012.parquet`: `35100` rows, `3279035` file bytes (3.13 MiB), `26872318` physical bytes (25.63 MiB), `15207316` encoded bytes (14.50 MiB), `3248696` compressed data bytes (3.10 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00013.parquet`: `34791` rows, `3317826` file bytes (3.16 MiB), `26795858` physical bytes (25.55 MiB), `15220460` encoded bytes (14.52 MiB), `3287126` compressed data bytes (3.13 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00014.parquet`: `32954` rows, `3545599` file bytes (3.38 MiB), `24465631` physical bytes (23.33 MiB), `15451081` encoded bytes (14.74 MiB), `3515195` compressed data bytes (3.35 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00015.parquet`: `30681` rows, `3747711` file bytes (3.57 MiB), `20905766` physical bytes (19.94 MiB), `15460236` encoded bytes (14.74 MiB), `3718196` compressed data bytes (3.55 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00016.parquet`: `29368` rows, `3799049` file bytes (3.62 MiB), `20757446` physical bytes (19.80 MiB), `15581107` encoded bytes (14.86 MiB), `3769838` compressed data bytes (3.60 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00017.parquet`: `30626` rows, `3950925` file bytes (3.77 MiB), `19955283` physical bytes (19.03 MiB), `15398770` encoded bytes (14.69 MiB), `3921537` compressed data bytes (3.74 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00018.parquet`: `30783` rows, `3962975` file bytes (3.78 MiB), `19813758` physical bytes (18.90 MiB), `15343042` encoded bytes (14.63 MiB), `3933057` compressed data bytes (3.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00019.parquet`: `29937` rows, `4053770` file bytes (3.87 MiB), `19580810` physical bytes (18.67 MiB), `15415577` encoded bytes (14.70 MiB), `4024257` compressed data bytes (3.84 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00020.parquet`: `30647` rows, `4003513` file bytes (3.82 MiB), `19553793` physical bytes (18.65 MiB), `15366736` encoded bytes (14.65 MiB), `3973643` compressed data bytes (3.79 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00021.parquet`: `30715` rows, `3998757` file bytes (3.81 MiB), `19657167` physical bytes (18.75 MiB), `15355780` encoded bytes (14.64 MiB), `3968971` compressed data bytes (3.79 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00022.parquet`: `30554` rows, `3986586` file bytes (3.80 MiB), `19629849` physical bytes (18.72 MiB), `15352298` encoded bytes (14.64 MiB), `3956815` compressed data bytes (3.77 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00023.parquet`: `30699` rows, `3971553` file bytes (3.79 MiB), `19751933` physical bytes (18.84 MiB), `15358104` encoded bytes (14.65 MiB), `3941523` compressed data bytes (3.76 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00024.parquet`: `30607` rows, `3989322` file bytes (3.80 MiB), `19676963` physical bytes (18.77 MiB), `15364836` encoded bytes (14.65 MiB), `3959496` compressed data bytes (3.78 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00025.parquet`: `30383` rows, `3999800` file bytes (3.81 MiB), `19606541` physical bytes (18.70 MiB), `15378886` encoded bytes (14.67 MiB), `3970234` compressed data bytes (3.79 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00026.parquet`: `30994` rows, `3965879` file bytes (3.78 MiB), `19675121` physical bytes (18.76 MiB), `15287763` encoded bytes (14.58 MiB), `3936486` compressed data bytes (3.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00027.parquet`: `30720` rows, `4001095` file bytes (3.82 MiB), `19690183` physical bytes (18.78 MiB), `15345864` encoded bytes (14.63 MiB), `3971105` compressed data bytes (3.79 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00028.parquet`: `31146` rows, `3946286` file bytes (3.76 MiB), `19828284` physical bytes (18.91 MiB), `15288967` encoded bytes (14.58 MiB), `3916491` compressed data bytes (3.74 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00029.parquet`: `30940` rows, `3964824` file bytes (3.78 MiB), `19621876` physical bytes (18.71 MiB), `15349434` encoded bytes (14.64 MiB), `3935366` compressed data bytes (3.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict/part-00030.parquet`: `20441` rows, `2679111` file bytes (2.55 MiB), `13229562` physical bytes (12.62 MiB), `10294697` encoded bytes (9.82 MiB), `2649572` compressed data bytes (2.53 MiB)
