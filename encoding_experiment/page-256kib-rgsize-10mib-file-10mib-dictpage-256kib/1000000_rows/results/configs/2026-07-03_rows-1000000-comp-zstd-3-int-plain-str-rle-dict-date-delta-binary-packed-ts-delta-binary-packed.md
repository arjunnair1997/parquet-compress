# ClickBench Parquet Experiment

- Started: `2026-07-03T23:36:47-04:00`
- Write elapsed: `11.452s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `457674213` (436.47 MiB)
- Compressed column data bytes after codec compression: `76135511` (72.61 MiB)
- Parquet file bytes: `77022375` (73.45 MiB)
- Physical/encoded ratio: `1.557x`
- Encoded/compressed-data ratio: `6.011x`
- Physical/compressed-data ratio: `9.357x`
- Physical/parquet-file ratio: `9.249x`
- Files: `30`

## Settings

- Compression: `zstd-3`
- Int encoding: `plain`
- String encoding: `rle-dict`
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
- Elapsed: `7.029s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004714` (7.63 MiB) | `8005548` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `64910` (63.39 KiB) | `0.999x` | `61.681x` | `61.624x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `138409995` (132.00 MiB) | `28809750` (27.48 MiB) | `7969796` (7.60 MiB) | `4.804x` | `3.615x` | `17.367x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003653` (3.82 MiB) | `5094` (4.97 KiB) | `0.999x` | `785.955x` | `785.238x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:59` | `1000000` | `8000000` (7.63 MiB) | `3369293` (3.21 MiB) | `2885504` (2.75 MiB) | `2.374x` | `1.168x` | `2.772x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:59` | `1000000` | `4000000` (3.81 MiB) | `51777` (50.56 KiB) | `6549` (6.40 KiB) | `77.254x` | `7.906x` | `610.780x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003656` (3.82 MiB) | `5114` (4.99 KiB) | `0.999x` | `782.882x` | `782.167x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003723` (3.82 MiB) | `408427` (398.85 KiB) | `0.999x` | `9.803x` | `9.794x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `191111` (186.63 KiB) | `0.999x` | `20.950x` | `20.930x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004717` (7.63 MiB) | `618464` (603.97 KiB) | `0.999x` | `12.943x` | `12.935x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003654` (3.82 MiB) | `4387` (4.28 KiB) | `0.999x` | `912.618x` | `911.785x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `107828` (105.30 KiB) | `0.999x` | `37.131x` | `37.096x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003710` (3.82 MiB) | `135093` (131.93 KiB) | `0.999x` | `29.637x` | `29.609x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `88562192` (84.46 MiB) | `44074234` (42.03 MiB) | `12702847` (12.11 MiB) | `2.009x` | `3.470x` | `6.972x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `79583339` (75.90 MiB) | `34392102` (32.80 MiB) | `11729940` (11.19 MiB) | `2.314x` | `2.932x` | `6.785x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `179065` (174.87 KiB) | `0.999x` | `22.359x` | `22.338x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `275454` (269.00 KiB) | `0.999x` | `14.535x` | `14.521x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `231542` (226.12 KiB) | `0.999x` | `17.292x` | `17.275x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `87796` (85.74 KiB) | `0.999x` | `45.602x` | `45.560x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `48445` (47.31 KiB) | `0.999x` | `82.645x` | `82.568x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `187716` (183.32 KiB) | `0.999x` | `21.329x` | `21.309x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003707` (3.82 MiB) | `186607` (182.23 KiB) | `0.999x` | `21.455x` | `21.435x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `82064` (80.14 KiB) | `0.999x` | `48.788x` | `48.742x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `53836` (52.57 KiB) | `0.999x` | `74.369x` | `74.300x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003711` (3.82 MiB) | `130113` (127.06 KiB) | `0.999x` | `30.771x` | `30.743x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `3354477` (3.20 MiB) | `266974` (260.72 KiB) | `145507` (142.10 KiB) | `12.565x` | `1.835x` | `23.054x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `26552` (25.93 KiB) | `0.999x` | `150.788x` | `150.648x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `25253` (24.66 KiB) | `0.999x` | `158.544x` | `158.397x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `154644` (151.02 KiB) | `0.999x` | `25.890x` | `25.866x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `3767530` (3.59 MiB) | `150402` (146.88 KiB) | `84491` (82.51 KiB) | `25.050x` | `1.780x` | `44.591x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003662` (3.82 MiB) | `6257` (6.11 KiB) | `0.999x` | `639.869x` | `639.284x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003668` (3.82 MiB) | `6706` (6.55 KiB) | `0.999x` | `597.028x` | `596.481x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `29407` (28.72 KiB) | `0.999x` | `136.148x` | `136.022x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `22766` (22.23 KiB) | `0.999x` | `175.864x` | `175.701x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `81583` (79.67 KiB) | `25102` (24.51 KiB) | `20027` (19.56 KiB) | `3.250x` | `1.253x` | `4.074x` | `81583` (79.67 KiB) |
| `Params` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `0` (0 B) | `4189` (4.09 KiB) | `5251` (5.13 KiB) | `0.000x` | `0.798x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `323969` (316.38 KiB) | `0.999x` | `12.358x` | `12.347x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `288471` (281.71 KiB) | `0.999x` | `13.879x` | `13.866x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `101885` (99.50 KiB) | `0.999x` | `39.296x` | `39.260x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `3528017` (3.36 MiB) | `1642359` (1.57 MiB) | `636612` (621.69 KiB) | `2.148x` | `2.580x` | `5.542x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003703` (3.82 MiB) | `30937` (30.21 KiB) | `0.999x` | `129.415x` | `129.295x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `165200` (161.33 KiB) | `0.999x` | `24.236x` | `24.213x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `306158` (298.98 KiB) | `0.999x` | `13.077x` | `13.065x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `319612` (312.12 KiB) | `0.999x` | `12.527x` | `12.515x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `100295` (97.94 KiB) | `0.999x` | `39.919x` | `39.882x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:59` | `1000000` | `8000000` (7.63 MiB) | `3491977` (3.33 MiB) | `2897356` (2.76 MiB) | `2.291x` | `1.205x` | `2.761x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `90465` (88.34 KiB) | `0.999x` | `44.257x` | `44.216x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003717` (3.82 MiB) | `74280` (72.54 KiB) | `0.999x` | `53.900x` | `53.850x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `124238` (121.33 KiB) | `0.999x` | `32.226x` | `32.196x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003653` (3.82 MiB) | `4514` (4.41 KiB) | `0.999x` | `886.941x` | `886.132x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `13587860` (12.96 MiB) | `8842` (8.63 KiB) | `9906` (9.67 KiB) | `1536.741x` | `0.893x` | `1371.680x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003674` (3.82 MiB) | `7240` (7.07 KiB) | `0.999x` | `552.994x` | `552.486x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003686` (3.82 MiB) | `57004` (55.67 KiB) | `0.999x` | `70.235x` | `70.171x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003680` (3.82 MiB) | `8222` (8.03 KiB) | `0.999x` | `486.947x` | `486.500x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003657` (3.82 MiB) | `25379` (24.78 KiB) | `0.999x` | `157.755x` | `157.611x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004713` (7.63 MiB) | `694201` (677.93 KiB) | `0.999x` | `11.531x` | `11.524x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `27797671` (26.51 MiB) | `21270049` (20.28 MiB) | `4883639` (4.66 MiB) | `1.307x` | `4.355x` | `5.692x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003775` (3.82 MiB) | `3792156` (3.62 MiB) | `0.999x` | `1.056x` | `1.055x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003654` (3.82 MiB) | `4387` (4.28 KiB) | `0.999x` | `912.618x` | `911.785x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003654` (3.82 MiB) | `4387` (4.28 KiB) | `0.999x` | `912.618x` | `911.785x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003654` (3.82 MiB) | `4387` (4.28 KiB) | `0.999x` | `912.618x` | `911.785x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `82819` (80.88 KiB) | `0.999x` | `48.343x` | `48.298x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003654` (3.82 MiB) | `4387` (4.28 KiB) | `0.999x` | `912.618x` | `911.785x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `1000000` (976.56 KiB) | `33593` (32.81 KiB) | `26178` (25.56 KiB) | `29.768x` | `1.283x` | `38.200x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `delta-binary-packed` | `DELTA_BINARY_PACKED` | `DATA_PAGE_V2/DELTA_BINARY_PACKED:59` | `1000000` | `8000000` (7.63 MiB) | `3383514` (3.23 MiB) | `2895172` (2.76 MiB) | `2.364x` | `1.169x` | `2.763x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `142939` (139.59 KiB) | `0.999x` | `28.010x` | `27.984x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `107830` (105.30 KiB) | `0.999x` | `37.130x` | `37.095x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `122919` (120.04 KiB) | `0.999x` | `32.572x` | `32.542x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `193830` (189.29 KiB) | `0.999x` | `20.656x` | `20.637x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `174113` (170.03 KiB) | `0.999x` | `22.995x` | `22.974x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003728` (3.82 MiB) | `426837` (416.83 KiB) | `0.999x` | `9.380x` | `9.371x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003705` (3.82 MiB) | `71062` (69.40 KiB) | `0.999x` | `56.341x` | `56.289x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003655` (3.82 MiB) | `4388` (4.29 KiB) | `0.999x` | `912.410x` | `911.577x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003701` (3.82 MiB) | `55494` (54.19 KiB) | `0.999x` | `72.147x` | `72.080x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `2001192` (1.91 MiB) | `42748` (41.75 KiB) | `27536` (26.89 KiB) | `46.814x` | `1.552x` | `72.675x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `3325142` (3.17 MiB) | `136714` (133.51 KiB) | `69463` (67.83 KiB) | `24.322x` | `1.968x` | `47.869x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `0` (0 B) | `4189` (4.09 KiB) | `5251` (5.13 KiB) | `0.000x` | `0.798x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `0` (0 B) | `4189` (4.09 KiB) | `5251` (5.13 KiB) | `0.000x` | `0.798x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003654` (3.82 MiB) | `4387` (4.28 KiB) | `0.999x` | `912.618x` | `911.785x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003671` (3.82 MiB) | `61771` (60.32 KiB) | `0.999x` | `64.815x` | `64.755x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `135632` (132.45 KiB) | `0.999x` | `29.519x` | `29.492x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003720` (3.82 MiB) | `333564` (325.75 KiB) | `0.999x` | `12.003x` | `11.992x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003771` (3.82 MiB) | `1246315` (1.19 MiB) | `0.999x` | `3.212x` | `3.209x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003768` (3.82 MiB) | `937980` (916.00 KiB) | `0.999x` | `4.269x` | `4.264x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003755` (3.82 MiB) | `550492` (537.59 KiB) | `0.999x` | `7.273x` | `7.266x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003658` (3.82 MiB) | `5489` (5.36 KiB) | `0.999x` | `729.397x` | `728.730x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `1024` (1.00 KiB) | `5861` (5.72 KiB) | `6923` (6.76 KiB) | `0.175x` | `0.847x` | `0.148x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004599` (7.63 MiB) | `5912` (5.77 KiB) | `0.999x` | `1353.958x` | `1353.180x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `0` (0 B) | `4189` (4.09 KiB) | `5251` (5.13 KiB) | `0.000x` | `0.798x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `3000000` (2.86 MiB) | `5074` (4.96 KiB) | `6136` (5.99 KiB) | `591.250x` | `0.827x` | `488.918x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003654` (3.82 MiB) | `4387` (4.28 KiB) | `0.999x` | `912.618x` | `911.785x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `58030` (56.67 KiB) | `18468` (18.04 KiB) | `17634` (17.22 KiB) | `3.142x` | `1.047x` | `3.291x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `22051` (21.53 KiB) | `16269` (15.89 KiB) | `15832` (15.46 KiB) | `1.355x` | `1.028x` | `1.393x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `25445` (24.85 KiB) | `22766` (22.23 KiB) | `18430` (18.00 KiB) | `1.118x` | `1.235x` | `1.381x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `48191` (47.06 KiB) | `14663` (14.32 KiB) | `12983` (12.68 KiB) | `3.287x` | `1.129x` | `3.712x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `49433` (48.27 KiB) | `22613` (22.08 KiB) | `18362` (17.93 KiB) | `2.186x` | `1.232x` | `2.692x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `16873` (16.48 KiB) | `16570` (16.18 KiB) | `14836` (14.49 KiB) | `1.018x` | `1.117x` | `1.137x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `91870` (89.72 KiB) | `37159` (36.29 KiB) | `27192` (26.55 KiB) | `2.472x` | `1.367x` | `3.379x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `13001` (12.70 KiB) | `14719` (14.37 KiB) | `14853` (14.50 KiB) | `0.883x` | `0.991x` | `0.875x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `28101` (27.44 KiB) | `16489` (16.10 KiB) | `16106` (15.73 KiB) | `1.704x` | `1.024x` | `1.745x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `45607` (44.54 KiB) | `41272` (40.30 KiB) | `22632` (22.10 KiB) | `1.105x` | `1.824x` | `2.015x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003684` (3.82 MiB) | `21320` (20.82 KiB) | `0.999x` | `187.790x` | `187.617x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004709` (7.63 MiB) | `2843216` (2.71 MiB) | `0.999x` | `2.815x` | `2.814x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004714` (7.63 MiB) | `3581585` (3.42 MiB) | `0.999x` | `2.235x` | `2.234x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003663` (3.82 MiB) | `5811` (5.67 KiB) | `0.999x` | `688.980x` | `688.350x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00000.parquet`: `35743` rows, `2314500` file bytes (2.21 MiB), `27762714` physical bytes (26.48 MiB), `15473465` encoded bytes (14.76 MiB), `2284121` compressed data bytes (2.18 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00001.parquet`: `35886` rows, `2324351` file bytes (2.22 MiB), `27565035` physical bytes (26.29 MiB), `15418040` encoded bytes (14.70 MiB), `2293808` compressed data bytes (2.19 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00002.parquet`: `35700` rows, `2332070` file bytes (2.22 MiB), `27969621` physical bytes (26.67 MiB), `15467155` encoded bytes (14.75 MiB), `2300885` compressed data bytes (2.19 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00003.parquet`: `35832` rows, `2314236` file bytes (2.21 MiB), `27756711` physical bytes (26.47 MiB), `15438942` encoded bytes (14.72 MiB), `2283930` compressed data bytes (2.18 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00004.parquet`: `35908` rows, `2307101` file bytes (2.20 MiB), `27902171` physical bytes (26.61 MiB), `15409571` encoded bytes (14.70 MiB), `2275931` compressed data bytes (2.17 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00005.parquet`: `36116` rows, `2295293` file bytes (2.19 MiB), `27692048` physical bytes (26.41 MiB), `15423510` encoded bytes (14.71 MiB), `2264795` compressed data bytes (2.16 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00006.parquet`: `35924` rows, `2317617` file bytes (2.21 MiB), `27808620` physical bytes (26.52 MiB), `15428119` encoded bytes (14.71 MiB), `2287288` compressed data bytes (2.18 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00007.parquet`: `35846` rows, `2325178` file bytes (2.22 MiB), `27730397` physical bytes (26.45 MiB), `15423940` encoded bytes (14.71 MiB), `2294409` compressed data bytes (2.19 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00008.parquet`: `36006` rows, `2288697` file bytes (2.18 MiB), `27904204` physical bytes (26.61 MiB), `15408753` encoded bytes (14.69 MiB), `2258046` compressed data bytes (2.15 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00009.parquet`: `35864` rows, `2324049` file bytes (2.22 MiB), `27683539` physical bytes (26.40 MiB), `15440256` encoded bytes (14.72 MiB), `2293823` compressed data bytes (2.19 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00010.parquet`: `36026` rows, `2279714` file bytes (2.17 MiB), `27705452` physical bytes (26.42 MiB), `15428269` encoded bytes (14.71 MiB), `2249260` compressed data bytes (2.15 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00011.parquet`: `36168` rows, `2273576` file bytes (2.17 MiB), `27896883` physical bytes (26.60 MiB), `15417499` encoded bytes (14.70 MiB), `2243184` compressed data bytes (2.14 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00012.parquet`: `36110` rows, `2282609` file bytes (2.18 MiB), `27671668` physical bytes (26.39 MiB), `15456855` encoded bytes (14.74 MiB), `2252354` compressed data bytes (2.15 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00013.parquet`: `35917` rows, `2292435` file bytes (2.19 MiB), `27789310` physical bytes (26.50 MiB), `15441002` encoded bytes (14.73 MiB), `2261802` compressed data bytes (2.16 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00014.parquet`: `32170` rows, `2704189` file bytes (2.58 MiB), `22474327` physical bytes (21.43 MiB), `15693029` encoded bytes (14.97 MiB), `2674504` compressed data bytes (2.55 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00015.parquet`: `31553` rows, `2724375` file bytes (2.60 MiB), `21825004` physical bytes (20.81 MiB), `15782892` encoded bytes (15.05 MiB), `2695258` compressed data bytes (2.57 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00016.parquet`: `31363` rows, `2816127` file bytes (2.69 MiB), `21491123` physical bytes (20.50 MiB), `15790667` encoded bytes (15.06 MiB), `2787035` compressed data bytes (2.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00017.parquet`: `32362` rows, `2921684` file bytes (2.79 MiB), `20662724` physical bytes (19.71 MiB), `15573539` encoded bytes (14.85 MiB), `2892322` compressed data bytes (2.76 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00018.parquet`: `31905` rows, `2932162` file bytes (2.80 MiB), `20740717` physical bytes (19.78 MiB), `15633235` encoded bytes (14.91 MiB), `2902626` compressed data bytes (2.77 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00019.parquet`: `32014` rows, `2958492` file bytes (2.82 MiB), `20561622` physical bytes (19.61 MiB), `15576235` encoded bytes (14.85 MiB), `2928715` compressed data bytes (2.79 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00020.parquet`: `32234` rows, `2937341` file bytes (2.80 MiB), `20586707` physical bytes (19.63 MiB), `15612333` encoded bytes (14.89 MiB), `2907654` compressed data bytes (2.77 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00021.parquet`: `31912` rows, `2953862` file bytes (2.82 MiB), `20622327` physical bytes (19.67 MiB), `15598293` encoded bytes (14.88 MiB), `2924331` compressed data bytes (2.79 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00022.parquet`: `32234` rows, `2912106` file bytes (2.78 MiB), `20752678` physical bytes (19.79 MiB), `15609115` encoded bytes (14.89 MiB), `2882265` compressed data bytes (2.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00023.parquet`: `32308` rows, `2916363` file bytes (2.78 MiB), `20668205` physical bytes (19.71 MiB), `15572255` encoded bytes (14.85 MiB), `2887231` compressed data bytes (2.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00024.parquet`: `31812` rows, `2952509` file bytes (2.82 MiB), `20603679` physical bytes (19.65 MiB), `15637642` encoded bytes (14.91 MiB), `2922732` compressed data bytes (2.79 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00025.parquet`: `32677` rows, `2903413` file bytes (2.77 MiB), `20717195` physical bytes (19.76 MiB), `15522259` encoded bytes (14.80 MiB), `2874212` compressed data bytes (2.74 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00026.parquet`: `32418` rows, `2931003` file bytes (2.80 MiB), `20689422` physical bytes (19.73 MiB), `15566392` encoded bytes (14.85 MiB), `2901093` compressed data bytes (2.77 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00027.parquet`: `32518` rows, `2909895` file bytes (2.78 MiB), `20792591` physical bytes (19.83 MiB), `15548169` encoded bytes (14.83 MiB), `2880131` compressed data bytes (2.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00028.parquet`: `32636` rows, `2910819` file bytes (2.78 MiB), `20692655` physical bytes (19.73 MiB), `15576463` encoded bytes (14.85 MiB), `2881524` compressed data bytes (2.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed/part-00029.parquet`: `14838` rows, `1366609` file bytes (1.30 MiB), `9679275` physical bytes (9.23 MiB), `7306319` encoded bytes (6.97 MiB), `1350242` compressed data bytes (1.29 MiB)
