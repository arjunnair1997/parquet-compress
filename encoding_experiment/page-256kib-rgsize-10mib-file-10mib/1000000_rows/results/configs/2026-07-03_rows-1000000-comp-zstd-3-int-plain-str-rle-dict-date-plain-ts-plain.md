# ClickBench Parquet Experiment

- Started: `2026-07-03T15:00:38-04:00`
- Write elapsed: `11.612s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `475404478` (453.38 MiB)
- Compressed column data bytes after codec compression: `74994244` (71.52 MiB)
- Parquet file bytes: `75881570` (72.37 MiB)
- Physical/encoded ratio: `1.499x`
- Encoded/compressed-data ratio: `6.339x`
- Physical/compressed-data ratio: `9.499x`
- Physical/parquet-file ratio: `9.388x`
- Files: `30`

## Settings

- Compression: `zstd-3`
- Int encoding: `plain`
- String encoding: `rle-dict`
- Date encoding: `plain`
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
- Files read: `30`
- Elapsed: `7.159s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004715` (7.63 MiB) | `8005549` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `64975` (63.45 KiB) | `0.999x` | `61.619x` | `61.562x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `138409995` (132.00 MiB) | `28805660` (27.47 MiB) | `7973171` (7.60 MiB) | `4.805x` | `3.613x` | `17.359x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003652` (3.82 MiB) | `5093` (4.97 KiB) | `0.999x` | `786.109x` | `785.392x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004716` (7.63 MiB) | `2517377` (2.40 MiB) | `0.999x` | `3.180x` | `3.178x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003657` (3.82 MiB) | `5098` (4.98 KiB) | `0.999x` | `785.339x` | `784.621x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003652` (3.82 MiB) | `5110` (4.99 KiB) | `0.999x` | `783.494x` | `782.779x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003719` (3.82 MiB) | `408356` (398.79 KiB) | `0.999x` | `9.804x` | `9.795x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `191014` (186.54 KiB) | `0.999x` | `20.960x` | `20.941x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004716` (7.63 MiB) | `618365` (603.87 KiB) | `0.999x` | `12.945x` | `12.937x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003651` (3.82 MiB) | `4384` (4.28 KiB) | `0.999x` | `913.242x` | `912.409x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `106695` (104.19 KiB) | `0.999x` | `37.525x` | `37.490x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003709` (3.82 MiB) | `134745` (131.59 KiB) | `0.999x` | `29.713x` | `29.686x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `88562192` (84.46 MiB) | `44076342` (42.03 MiB) | `12708183` (12.12 MiB) | `2.009x` | `3.468x` | `6.969x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `79583339` (75.90 MiB) | `34397247` (32.80 MiB) | `11732666` (11.19 MiB) | `2.314x` | `2.932x` | `6.783x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003711` (3.82 MiB) | `179019` (174.82 KiB) | `0.999x` | `22.365x` | `22.344x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `275509` (269.05 KiB) | `0.999x` | `14.532x` | `14.519x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `231684` (226.25 KiB) | `0.999x` | `17.281x` | `17.265x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `87902` (85.84 KiB) | `0.999x` | `45.547x` | `45.505x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `48450` (47.31 KiB) | `0.999x` | `82.636x` | `82.559x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `187595` (183.20 KiB) | `0.999x` | `21.342x` | `21.323x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003711` (3.82 MiB) | `186530` (182.16 KiB) | `0.999x` | `21.464x` | `21.444x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `82176` (80.25 KiB) | `0.999x` | `48.721x` | `48.676x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `53618` (52.36 KiB) | `0.999x` | `74.671x` | `74.602x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003713` (3.82 MiB) | `128549` (125.54 KiB) | `0.999x` | `31.145x` | `31.117x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `3354477` (3.20 MiB) | `267186` (260.92 KiB) | `146104` (142.68 KiB) | `12.555x` | `1.829x` | `22.960x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `26413` (25.79 KiB) | `0.999x` | `151.581x` | `151.441x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003711` (3.82 MiB) | `25275` (24.68 KiB) | `0.999x` | `158.406x` | `158.259x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003711` (3.82 MiB) | `155184` (151.55 KiB) | `0.999x` | `25.800x` | `25.776x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `3767530` (3.59 MiB) | `150388` (146.86 KiB) | `83780` (81.82 KiB) | `25.052x` | `1.795x` | `44.969x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003661` (3.82 MiB) | `6254` (6.11 KiB) | `0.999x` | `640.176x` | `639.591x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003663` (3.82 MiB) | `6703` (6.55 KiB) | `0.999x` | `597.294x` | `596.748x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `29311` (28.62 KiB) | `0.999x` | `136.594x` | `136.468x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `22907` (22.37 KiB) | `0.999x` | `174.781x` | `174.619x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `81583` (79.67 KiB) | `25127` (24.54 KiB) | `20257` (19.78 KiB) | `3.247x` | `1.240x` | `4.027x` | `81583` (79.67 KiB) |
| `Params` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `0` (0 B) | `4185` (4.09 KiB) | `5247` (5.12 KiB) | `0.000x` | `0.798x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003711` (3.82 MiB) | `324112` (316.52 KiB) | `0.999x` | `12.353x` | `12.341x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `288545` (281.78 KiB) | `0.999x` | `13.876x` | `13.863x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `102120` (99.73 KiB) | `0.999x` | `39.206x` | `39.170x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `3528017` (3.36 MiB) | `1642817` (1.57 MiB) | `636274` (621.36 KiB) | `2.148x` | `2.582x` | `5.545x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003704` (3.82 MiB) | `30854` (30.13 KiB) | `0.999x` | `129.763x` | `129.643x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003711` (3.82 MiB) | `165126` (161.26 KiB) | `0.999x` | `24.246x` | `24.224x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `306119` (298.94 KiB) | `0.999x` | `13.079x` | `13.067x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `319680` (312.19 KiB) | `0.999x` | `12.524x` | `12.513x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `100187` (97.84 KiB) | `0.999x` | `39.962x` | `39.925x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004715` (7.63 MiB) | `2474364` (2.36 MiB) | `0.999x` | `3.235x` | `3.233x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `90165` (88.05 KiB) | `0.999x` | `44.404x` | `44.363x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003715` (3.82 MiB) | `74387` (72.64 KiB) | `0.999x` | `53.823x` | `53.773x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `124223` (121.31 KiB) | `0.999x` | `32.230x` | `32.200x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003652` (3.82 MiB) | `4515` (4.41 KiB) | `0.999x` | `886.745x` | `885.936x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `13587860` (12.96 MiB) | `8862` (8.65 KiB) | `9926` (9.69 KiB) | `1533.272x` | `0.893x` | `1368.916x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003675` (3.82 MiB) | `7277` (7.11 KiB) | `0.999x` | `550.182x` | `549.677x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003680` (3.82 MiB) | `56906` (55.57 KiB) | `0.999x` | `70.356x` | `70.291x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003679` (3.82 MiB) | `8272` (8.08 KiB) | `0.999x` | `484.004x` | `483.559x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003660` (3.82 MiB) | `25328` (24.73 KiB) | `0.999x` | `158.072x` | `157.928x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004716` (7.63 MiB) | `694191` (677.92 KiB) | `0.999x` | `11.531x` | `11.524x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `27797671` (26.51 MiB) | `21274781` (20.29 MiB) | `4881011` (4.65 MiB) | `1.307x` | `4.359x` | `5.695x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003772` (3.82 MiB) | `3817370` (3.64 MiB) | `0.999x` | `1.049x` | `1.048x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003651` (3.82 MiB) | `4384` (4.28 KiB) | `0.999x` | `913.242x` | `912.409x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003651` (3.82 MiB) | `4384` (4.28 KiB) | `0.999x` | `913.242x` | `912.409x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003651` (3.82 MiB) | `4384` (4.28 KiB) | `0.999x` | `913.242x` | `912.409x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `82706` (80.77 KiB) | `0.999x` | `48.409x` | `48.364x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003651` (3.82 MiB) | `4384` (4.28 KiB) | `0.999x` | `913.242x` | `912.409x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `1000000` (976.56 KiB) | `33649` (32.86 KiB) | `24265` (23.70 KiB) | `29.719x` | `1.387x` | `41.212x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004713` (7.63 MiB) | `2517524` (2.40 MiB) | `0.999x` | `3.180x` | `3.178x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `142821` (139.47 KiB) | `0.999x` | `28.033x` | `28.007x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003710` (3.82 MiB) | `108087` (105.55 KiB) | `0.999x` | `37.042x` | `37.007x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `122810` (119.93 KiB) | `0.999x` | `32.601x` | `32.571x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003712` (3.82 MiB) | `193807` (189.26 KiB) | `0.999x` | `20.658x` | `20.639x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003716` (3.82 MiB) | `174059` (169.98 KiB) | `0.999x` | `23.002x` | `22.981x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003722` (3.82 MiB) | `426688` (416.69 KiB) | `0.999x` | `9.383x` | `9.375x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003705` (3.82 MiB) | `71038` (69.37 KiB) | `0.999x` | `56.360x` | `56.308x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003655` (3.82 MiB) | `4388` (4.29 KiB) | `0.999x` | `912.410x` | `911.577x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003698` (3.82 MiB) | `55420` (54.12 KiB) | `0.999x` | `72.243x` | `72.176x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `2001192` (1.91 MiB) | `42888` (41.88 KiB) | `27536` (26.89 KiB) | `46.661x` | `1.558x` | `72.675x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `3325142` (3.17 MiB) | `136664` (133.46 KiB) | `69665` (68.03 KiB) | `24.331x` | `1.962x` | `47.730x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `0` (0 B) | `4185` (4.09 KiB) | `5247` (5.12 KiB) | `0.000x` | `0.798x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `0` (0 B) | `4185` (4.09 KiB) | `5247` (5.12 KiB) | `0.000x` | `0.798x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003651` (3.82 MiB) | `4384` (4.28 KiB) | `0.999x` | `913.242x` | `912.409x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003667` (3.82 MiB) | `61691` (60.25 KiB) | `0.999x` | `64.899x` | `64.839x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003714` (3.82 MiB) | `135526` (132.35 KiB) | `0.999x` | `29.542x` | `29.515x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003722` (3.82 MiB) | `333919` (326.09 KiB) | `0.999x` | `11.990x` | `11.979x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003770` (3.82 MiB) | `1246200` (1.19 MiB) | `0.999x` | `3.213x` | `3.210x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003771` (3.82 MiB) | `938048` (916.06 KiB) | `0.999x` | `4.268x` | `4.264x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003752` (3.82 MiB) | `550203` (537.31 KiB) | `0.999x` | `7.277x` | `7.270x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003654` (3.82 MiB) | `5487` (5.36 KiB) | `0.999x` | `729.662x` | `728.996x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `1024` (1.00 KiB) | `5854` (5.72 KiB) | `6916` (6.75 KiB) | `0.175x` | `0.846x` | `0.148x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004602` (7.63 MiB) | `5915` (5.78 KiB) | `0.999x` | `1353.272x` | `1352.494x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `0` (0 B) | `4185` (4.09 KiB) | `5247` (5.12 KiB) | `0.000x` | `0.798x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `3000000` (2.86 MiB) | `5070` (4.95 KiB) | `6132` (5.99 KiB) | `591.716x` | `0.827x` | `489.237x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003651` (3.82 MiB) | `4384` (4.28 KiB) | `0.999x` | `913.242x` | `912.409x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `58030` (56.67 KiB) | `18496` (18.06 KiB) | `17787` (17.37 KiB) | `3.137x` | `1.040x` | `3.262x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `22051` (21.53 KiB) | `16327` (15.94 KiB) | `15738` (15.37 KiB) | `1.351x` | `1.037x` | `1.401x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `25445` (24.85 KiB) | `22713` (22.18 KiB) | `18389` (17.96 KiB) | `1.120x` | `1.235x` | `1.384x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `48191` (47.06 KiB) | `14731` (14.39 KiB) | `12981` (12.68 KiB) | `3.271x` | `1.135x` | `3.712x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `49433` (48.27 KiB) | `22696` (22.16 KiB) | `18403` (17.97 KiB) | `2.178x` | `1.233x` | `2.686x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `16873` (16.48 KiB) | `16596` (16.21 KiB) | `14813` (14.47 KiB) | `1.017x` | `1.120x` | `1.139x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `91870` (89.72 KiB) | `37222` (36.35 KiB) | `27202` (26.56 KiB) | `2.468x` | `1.368x` | `3.377x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `13001` (12.70 KiB) | `14824` (14.48 KiB) | `15030` (14.68 KiB) | `0.877x` | `0.986x` | `0.865x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `28101` (27.44 KiB) | `16498` (16.11 KiB) | `16041` (15.67 KiB) | `1.703x` | `1.028x` | `1.752x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:59, DICTIONARY_PAGE/PLAIN:59` | `1000000` | `45607` (44.54 KiB) | `41244` (40.28 KiB) | `22585` (22.06 KiB) | `1.106x` | `1.826x` | `2.019x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003683` (3.82 MiB) | `21337` (20.84 KiB) | `0.999x` | `187.640x` | `187.468x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004715` (7.63 MiB) | `2843415` (2.71 MiB) | `0.999x` | `2.815x` | `2.814x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `8000000` (7.63 MiB) | `8004717` (7.63 MiB) | `3581620` (3.42 MiB) | `0.999x` | `2.235x` | `2.234x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:59` | `1000000` | `4000000` (3.81 MiB) | `4003665` (3.82 MiB) | `5807` (5.67 KiB) | `0.999x` | `689.455x` | `688.824x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00000.parquet`: `35803` rows, `2254657` file bytes (2.15 MiB), `27806866` physical bytes (26.52 MiB), `16124845` encoded bytes (15.38 MiB), `2224271` compressed data bytes (2.12 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00001.parquet`: `35961` rows, `2254846` file bytes (2.15 MiB), `27624668` physical bytes (26.34 MiB), `16092839` encoded bytes (15.35 MiB), `2224297` compressed data bytes (2.12 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00002.parquet`: `35766` rows, `2273903` file bytes (2.17 MiB), `28014106` physical bytes (26.72 MiB), `16126834` encoded bytes (15.38 MiB), `2242711` compressed data bytes (2.14 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00003.parquet`: `35893` rows, `2255187` file bytes (2.15 MiB), `27810045` physical bytes (26.52 MiB), `16104115` encoded bytes (15.36 MiB), `2224874` compressed data bytes (2.12 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00004.parquet`: `35960` rows, `2247633` file bytes (2.14 MiB), `27944887` physical bytes (26.65 MiB), `16076223` encoded bytes (15.33 MiB), `2216464` compressed data bytes (2.11 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00005.parquet`: `36162` rows, `2243957` file bytes (2.14 MiB), `27713193` physical bytes (26.43 MiB), `16078358` encoded bytes (15.33 MiB), `2213453` compressed data bytes (2.11 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00006.parquet`: `35965` rows, `2260774` file bytes (2.16 MiB), `27852804` physical bytes (26.56 MiB), `16091564` encoded bytes (15.35 MiB), `2230198` compressed data bytes (2.13 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00007.parquet`: `35915` rows, `2266835` file bytes (2.16 MiB), `27796181` physical bytes (26.51 MiB), `16086682` encoded bytes (15.34 MiB), `2236062` compressed data bytes (2.13 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00008.parquet`: `36054` rows, `2228434` file bytes (2.13 MiB), `27945096` physical bytes (26.65 MiB), `16082345` encoded bytes (15.34 MiB), `2197777` compressed data bytes (2.10 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00009.parquet`: `35909` rows, `2269551` file bytes (2.16 MiB), `27709544` physical bytes (26.43 MiB), `16106358` encoded bytes (15.36 MiB), `2239314` compressed data bytes (2.14 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00010.parquet`: `36124` rows, `2206897` file bytes (2.10 MiB), `27768225` physical bytes (26.48 MiB), `16098082` encoded bytes (15.35 MiB), `2176505` compressed data bytes (2.08 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00011.parquet`: `36245` rows, `2209945` file bytes (2.11 MiB), `27946846` physical bytes (26.65 MiB), `16085171` encoded bytes (15.34 MiB), `2179547` compressed data bytes (2.08 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00012.parquet`: `36112` rows, `2244257` file bytes (2.14 MiB), `27721707` physical bytes (26.44 MiB), `16111031` encoded bytes (15.36 MiB), `2214046` compressed data bytes (2.11 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00013.parquet`: `35958` rows, `2239769` file bytes (2.14 MiB), `27774029` physical bytes (26.49 MiB), `16102539` encoded bytes (15.36 MiB), `2209092` compressed data bytes (2.11 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00014.parquet`: `32133` rows, `2673680` file bytes (2.55 MiB), `22403884` physical bytes (21.37 MiB), `16278128` encoded bytes (15.52 MiB), `2643990` compressed data bytes (2.52 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00015.parquet`: `31531` rows, `2697654` file bytes (2.57 MiB), `21828293` physical bytes (20.82 MiB), `16354466` encoded bytes (15.60 MiB), `2668578` compressed data bytes (2.54 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00016.parquet`: `31425` rows, `2799826` file bytes (2.67 MiB), `21485188` physical bytes (20.49 MiB), `16347039` encoded bytes (15.59 MiB), `2770729` compressed data bytes (2.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00017.parquet`: `32460` rows, `2901195` file bytes (2.77 MiB), `20683534` physical bytes (19.73 MiB), `16147935` encoded bytes (15.40 MiB), `2871778` compressed data bytes (2.74 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00018.parquet`: `31903` rows, `2919917` file bytes (2.78 MiB), `20761526` physical bytes (19.80 MiB), `16206258` encoded bytes (15.46 MiB), `2890359` compressed data bytes (2.76 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00019.parquet`: `31991` rows, `2949219` file bytes (2.81 MiB), `20543256` physical bytes (19.59 MiB), `16147437` encoded bytes (15.40 MiB), `2919415` compressed data bytes (2.78 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00020.parquet`: `32176` rows, `2930016` file bytes (2.79 MiB), `20575029` physical bytes (19.62 MiB), `16181891` encoded bytes (15.43 MiB), `2900322` compressed data bytes (2.77 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00021.parquet`: `31992` rows, `2929601` file bytes (2.79 MiB), `20633482` physical bytes (19.68 MiB), `16170022` encoded bytes (15.42 MiB), `2900027` compressed data bytes (2.77 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00022.parquet`: `32207` rows, `2906811` file bytes (2.77 MiB), `20753818` physical bytes (19.79 MiB), `16180196` encoded bytes (15.43 MiB), `2876961` compressed data bytes (2.74 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00023.parquet`: `32228` rows, `2908179` file bytes (2.77 MiB), `20665909` physical bytes (19.71 MiB), `16161540` encoded bytes (15.41 MiB), `2878992` compressed data bytes (2.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00024.parquet`: `31955` rows, `2931027` file bytes (2.80 MiB), `20631477` physical bytes (19.68 MiB), `16203103` encoded bytes (15.45 MiB), `2901205` compressed data bytes (2.77 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00025.parquet`: `32644` rows, `2891173` file bytes (2.76 MiB), `20718511` physical bytes (19.76 MiB), `16111894` encoded bytes (15.37 MiB), `2861962` compressed data bytes (2.73 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00026.parquet`: `32464` rows, `2912218` file bytes (2.78 MiB), `20703581` physical bytes (19.74 MiB), `16137626` encoded bytes (15.39 MiB), `2882348` compressed data bytes (2.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00027.parquet`: `32492` rows, `2898827` file bytes (2.76 MiB), `20773631` physical bytes (19.81 MiB), `16135471` encoded bytes (15.39 MiB), `2869056` compressed data bytes (2.74 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00028.parquet`: `32624` rows, `2895090` file bytes (2.76 MiB), `20734621` physical bytes (19.77 MiB), `16165151` encoded bytes (15.42 MiB), `2865788` compressed data bytes (2.73 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-plain/part-00029.parquet`: `13948` rows, `1280492` file bytes (1.22 MiB), `9074687` physical bytes (8.65 MiB), `7109335` encoded bytes (6.78 MiB), `1264123` compressed data bytes (1.21 MiB)
