# ClickBench Parquet Experiment

- Started: `2026-07-05T18:26:58-04:00`
- Write elapsed: `280ms`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/10000_rows/parquet/page_encoding_distribution/rle-dict-zstd`
- Rows: `10000`
- Writer stats JSON: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/10000_rows/results/page_encoding_distribution/stats/rle-dict-zstd_writer-stats.json`
- Source TSV bytes for rows, reference only: `8365486` (7.98 MiB)
- Parquet physical bytes before page encoding: `7711890` (7.35 MiB)
- Encoded column bytes before codec compression: `1637809` (1.56 MiB)
- Compressed column data bytes after codec compression: `751801` (734.18 KiB)
- Parquet file bytes: `769307` (751.28 KiB)
- Physical/encoded ratio: `4.709x`
- Encoded/compressed-data ratio: `2.179x`
- Physical/compressed-data ratio: `10.258x`
- Physical/parquet-file ratio: `10.024x`
- Files: `1`

## Settings

- Compression: `zstd-3`
- Int encoding: `rle-dict`
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
- Rows read and compared: `10000`
- Files read: `1`
- Elapsed: `75ms`
- Source TSV bytes checked: `8365486` (7.98 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-05_rows-10000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-rle-dict_columns.tsv](../../../tsvs/page_encoding_distribution/2026-07-05_rows-10000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-rle-dict_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Dict pages | Dict page bytes | Compressed bytes without dict pages | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `97608` (95.32 KiB) | `97630` (95.34 KiB) | `1` | `80021` (78.15 KiB) | `17609` (17.20 KiB) | `0.820x` | `1.000x` | `0.819x` | `190000` (185.55 KiB) |
| `JavaEnable` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `851` (851 B) | `585` (585 B) | `1` | `26` (26 B) | `559` (559 B) | `47.004x` | `1.455x` | `68.376x` | `10000` (9.77 KiB) |
| `Title` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `2200605` (2.10 MiB) | `545453` (532.67 KiB) | `148809` (145.32 KiB) | `1` | `135715` (132.53 KiB) | `13094` (12.79 KiB) | `4.034x` | `3.665x` | `14.788x` | `2200673` (2.10 MiB) |
| `GoodEvent` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `105` (105 B) | `1` | `22` (22 B) | `83` (83 B) | `459.770x` | `0.829x` | `380.952x` | `10000` (9.77 KiB) |
| `EventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `58661` (57.29 KiB) | `33469` (32.68 KiB) | `1` | `17111` (16.71 KiB) | `16358` (15.97 KiB) | `1.364x` | `1.753x` | `2.390x` | `190000` (185.55 KiB) |
| `EventDate` | `date` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `105` (105 B) | `1` | `22` (22 B) | `83` (83 B) | `459.770x` | `0.829x` | `380.952x` | `100000` (97.66 KiB) |
| `CounterID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `101` (101 B) | `119` (119 B) | `1` | `30` (30 B) | `89` (89 B) | `396.040x` | `0.849x` | `336.134x` | `20000` (19.53 KiB) |
| `ClientIP` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `8213` (8.02 KiB) | `7233` (7.06 KiB) | `1` | `2955` (2.89 KiB) | `4278` (4.18 KiB) | `4.870x` | `1.135x` | `5.530x` | `100283` (97.93 KiB) |
| `RegionID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `3687` (3.60 KiB) | `2608` (2.55 KiB) | `1` | `532` (532 B) | `2076` (2.03 KiB) | `10.849x` | `1.414x` | `15.337x` | `24459` (23.89 KiB) |
| `UserID` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `10952` (10.70 KiB) | `10018` (9.78 KiB) | `1` | `5803` (5.67 KiB) | `4215` (4.12 KiB) | `7.305x` | `1.093x` | `7.986x` | `199987` (195.30 KiB) |
| `CounterClass` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `105` (105 B) | `1` | `22` (22 B) | `83` (83 B) | `459.770x` | `0.829x` | `380.952x` | `10000` (9.77 KiB) |
| `OS` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `2326` (2.27 KiB) | `1354` (1.32 KiB) | `1` | `102` (102 B) | `1252` (1.22 KiB) | `17.197x` | `1.718x` | `29.542x` | `16130` (15.75 KiB) |
| `UserAgent` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `2298` (2.24 KiB) | `1470` (1.44 KiB) | `1` | `82` (82 B) | `1388` (1.36 KiB) | `17.406x` | `1.563x` | `27.211x` | `11313` (11.05 KiB) |
| `URL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `760012` (742.20 KiB) | `247077` (241.29 KiB) | `73426` (71.71 KiB) | `1` | `60125` (58.72 KiB) | `13301` (12.99 KiB) | `3.076x` | `3.365x` | `10.351x` | `760022` (742.21 KiB) |
| `Referer` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `797798` (779.10 KiB) | `233930` (228.45 KiB) | `82816` (80.88 KiB) | `1` | `71372` (69.70 KiB) | `11444` (11.18 KiB) | `3.410x` | `2.825x` | `9.633x` | `797822` (779.12 KiB) |
| `IsRefresh` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `1353` (1.32 KiB) | `1372` (1.34 KiB) | `1` | `26` (26 B) | `1346` (1.31 KiB) | `29.564x` | `0.986x` | `29.155x` | `10000` (9.77 KiB) |
| `RefererCategoryID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `4330` (4.23 KiB) | `2204` (2.15 KiB) | `1` | `146` (146 B) | `2058` (2.01 KiB) | `9.238x` | `1.965x` | `18.149x` | `47318` (46.21 KiB) |
| `RefererRegionID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `3616` (3.53 KiB) | `1764` (1.72 KiB) | `1` | `74` (74 B) | `1690` (1.65 KiB) | `11.062x` | `2.050x` | `22.676x` | `28586` (27.92 KiB) |
| `URLCategoryID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `546` (546 B) | `418` (418 B) | `1` | `46` (46 B) | `372` (372 B) | `73.260x` | `1.306x` | `95.694x` | `49424` (48.27 KiB) |
| `URLRegionID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `558` (558 B) | `448` (448 B) | `1` | `46` (46 B) | `402` (402 B) | `71.685x` | `1.246x` | `89.286x` | `29545` (28.85 KiB) |
| `ResolutionWidth` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `3278` (3.20 KiB) | `2152` (2.10 KiB) | `1` | `262` (262 B) | `1890` (1.85 KiB) | `12.203x` | `1.523x` | `18.587x` | `39093` (38.18 KiB) |
| `ResolutionHeight` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `3235` (3.16 KiB) | `2087` (2.04 KiB) | `1` | `254` (254 B) | `1833` (1.79 KiB) | `12.365x` | `1.550x` | `19.166x` | `34257` (33.45 KiB) |
| `ResolutionDepth` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `1094` (1.07 KiB) | `734` (734 B) | `1` | `34` (34 B) | `700` (700 B) | `36.563x` | `1.490x` | `54.496x` | `19813` (19.35 KiB) |
| `FlashMajor` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `1184` (1.16 KiB) | `720` (720 B) | `1` | `38` (38 B) | `682` (682 B) | `33.784x` | `1.644x` | `55.556x` | `18909` (18.47 KiB) |
| `FlashMinor` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `2093` (2.04 KiB) | `1228` (1.20 KiB) | `1` | `54` (54 B) | `1174` (1.15 KiB) | `19.111x` | `1.704x` | `32.573x` | `11787` (11.51 KiB) |
| `FlashMinor2` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `37170` (36.30 KiB) | `2794` (2.73 KiB) | `1692` (1.65 KiB) | `1` | `127` (127 B) | `1565` (1.53 KiB) | `13.304x` | `1.651x` | `21.968x` | `37170` (36.30 KiB) |
| `NetMajor` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `483` (483 B) | `361` (361 B) | `1` | `34` (34 B) | `327` (327 B) | `82.816x` | `1.338x` | `110.803x` | `10000` (9.77 KiB) |
| `NetMinor` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `438` (438 B) | `334` (334 B) | `1` | `30` (30 B) | `304` (304 B) | `91.324x` | `1.311x` | `119.760x` | `10000` (9.77 KiB) |
| `UserAgentMajor` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `2804` (2.74 KiB) | `2095` (2.05 KiB) | `1` | `130` (130 B) | `1965` (1.92 KiB) | `14.265x` | `1.338x` | `19.093x` | `18125` (17.70 KiB) |
| `UserAgentMinor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `36930` (36.06 KiB) | `1986` (1.94 KiB) | `1278` (1.25 KiB) | `1` | `135` (135 B) | `1143` (1.12 KiB) | `18.595x` | `1.554x` | `28.897x` | `37108` (36.24 KiB) |
| `CookieEnable` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `232` (232 B) | `250` (250 B) | `1` | `26` (26 B) | `224` (224 B) | `172.414x` | `0.928x` | `160.000x` | `10000` (9.77 KiB) |
| `JavascriptEnable` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `231` (231 B) | `249` (249 B) | `1` | `26` (26 B) | `223` (223 B) | `173.160x` | `0.928x` | `160.643x` | `10000` (9.77 KiB) |
| `IsMobile` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `305` (305 B) | `323` (323 B) | `1` | `26` (26 B) | `297` (297 B) | `131.148x` | `0.944x` | `123.839x` | `10000` (9.77 KiB) |
| `MobilePhone` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `355` (355 B) | `290` (290 B) | `1` | `42` (42 B) | `248` (248 B) | `112.676x` | `1.224x` | `137.931x` | `10011` (9.78 KiB) |
| `MobilePhoneModel` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `1140` (1.11 KiB) | `216` (216 B) | `234` (234 B) | `1` | `30` (30 B) | `204` (204 B) | `5.278x` | `0.923x` | `4.872x` | `1140` (1.11 KiB) |
| `Params` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `0` (0 B) | `71` (71 B) | `89` (89 B) | `1` | `22` (22 B) | `67` (67 B) | `0.000x` | `0.798x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `6819` (6.66 KiB) | `6058` (5.92 KiB) | `1` | `2031` (1.98 KiB) | `4027` (3.93 KiB) | `5.866x` | `1.126x` | `6.603x` | `68529` (66.92 KiB) |
| `TraficSourceID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `2896` (2.83 KiB) | `1961` (1.92 KiB) | `1` | `50` (50 B) | `1911` (1.87 KiB) | `13.812x` | `1.477x` | `20.398x` | `17645` (17.23 KiB) |
| `SearchEngineID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `2763` (2.70 KiB) | `1214` (1.19 KiB) | `1` | `62` (62 B) | `1152` (1.12 KiB) | `14.477x` | `2.276x` | `32.949x` | `10166` (9.93 KiB) |
| `SearchPhrase` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `65246` (63.72 KiB) | `28478` (27.81 KiB) | `11366` (11.10 KiB) | `1` | `8654` (8.45 KiB) | `2712` (2.65 KiB) | `2.291x` | `2.506x` | `5.740x` | `65250` (63.72 KiB) |
| `AdvEngineID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `680` (680 B) | `461` (461 B) | `1` | `30` (30 B) | `431` (431 B) | `58.824x` | `1.475x` | `86.768x` | `10142` (9.90 KiB) |
| `IsArtifical` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `1352` (1.32 KiB) | `1371` (1.34 KiB) | `1` | `26` (26 B) | `1345` (1.31 KiB) | `29.586x` | `0.986x` | `29.176x` | `10000` (9.77 KiB) |
| `WindowClientWidth` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `6544` (6.39 KiB) | `4128` (4.03 KiB) | `1` | `1000` (1000 B) | `3128` (3.05 KiB) | `6.112x` | `1.585x` | `9.690x` | `37695` (36.81 KiB) |
| `WindowClientHeight` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `6871` (6.71 KiB) | `5983` (5.84 KiB) | `1` | `1579` (1.54 KiB) | `4404` (4.30 KiB) | `5.822x` | `1.148x` | `6.686x` | `30023` (29.32 KiB) |
| `ClientTimeZone` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `1228` (1.20 KiB) | `741` (741 B) | `1` | `70` (70 B) | `671` (671 B) | `32.573x` | `1.657x` | `53.981x` | `29805` (29.11 KiB) |
| `ClientEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `55120` (53.83 KiB) | `32305` (31.55 KiB) | `1` | `16016` (15.64 KiB) | `16289` (15.91 KiB) | `1.451x` | `1.706x` | `2.476x` | `190000` (185.55 KiB) |
| `SilverlightVersion1` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `1479` (1.44 KiB) | `840` (840 B) | `1` | `42` (42 B) | `798` (798 B) | `27.045x` | `1.761x` | `47.619x` | `10017` (9.78 KiB) |
| `SilverlightVersion2` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `833` (833 B) | `558` (558 B) | `1` | `26` (26 B) | `532` (532 B) | `48.019x` | `1.493x` | `71.685x` | `10000` (9.77 KiB) |
| `SilverlightVersion3` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `2142` (2.09 KiB) | `1238` (1.21 KiB) | `1` | `98` (98 B) | `1140` (1.11 KiB) | `18.674x` | `1.730x` | `32.310x` | `26504` (25.88 KiB) |
| `SilverlightVersion4` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `105` (105 B) | `1` | `22` (22 B) | `83` (83 B) | `459.770x` | `0.829x` | `380.952x` | `10000` (9.77 KiB) |
| `PageCharset` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `196540` (191.93 KiB) | `280` (280 B) | `298` (298 B) | `1` | `46` (46 B) | `252` (252 B) | `701.929x` | `0.940x` | `659.530x` | `196540` (191.93 KiB) |
| `CodeVersion` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `231` (231 B) | `249` (249 B) | `1` | `26` (26 B) | `223` (223 B) | `173.160x` | `0.928x` | `160.643x` | `39532` (38.61 KiB) |
| `IsLink` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `105` (105 B) | `1` | `22` (22 B) | `83` (83 B) | `459.770x` | `0.829x` | `380.952x` | `10000` (9.77 KiB) |
| `IsDownload` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `105` (105 B) | `1` | `22` (22 B) | `83` (83 B) | `459.770x` | `0.829x` | `380.952x` | `10000` (9.77 KiB) |
| `IsNotBounce` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `105` (105 B) | `1` | `22` (22 B) | `83` (83 B) | `459.770x` | `0.829x` | `380.952x` | `10000` (9.77 KiB) |
| `FUniqID` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `10320` (10.08 KiB) | `9361` (9.14 KiB) | `1` | `5171` (5.05 KiB) | `4190` (4.09 KiB) | `7.752x` | `1.102x` | `8.546x` | `176860` (172.71 KiB) |
| `OriginalURL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `74750` (73.00 KiB) | `32330` (31.57 KiB) | `8962` (8.75 KiB) | `1` | `7232` (7.06 KiB) | `1730` (1.69 KiB) | `2.312x` | `3.607x` | `8.341x` | `74750` (73.00 KiB) |
| `HID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `35968` (35.12 KiB) | `35988` (35.14 KiB) | `1` | `19715` (19.25 KiB) | `16273` (15.89 KiB) | `1.112x` | `0.999x` | `1.111x` | `88404` (86.33 KiB) |
| `IsOldCounter` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `105` (105 B) | `1` | `22` (22 B) | `83` (83 B) | `459.770x` | `0.829x` | `380.952x` | `10000` (9.77 KiB) |
| `IsEvent` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `105` (105 B) | `1` | `22` (22 B) | `83` (83 B) | `459.770x` | `0.829x` | `380.952x` | `10000` (9.77 KiB) |
| `IsParameter` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `105` (105 B) | `1` | `22` (22 B) | `83` (83 B) | `459.770x` | `0.829x` | `380.952x` | `10000` (9.77 KiB) |
| `DontCountHits` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `194` (194 B) | `212` (212 B) | `1` | `26` (26 B) | `186` (186 B) | `206.186x` | `0.915x` | `188.679x` | `10000` (9.77 KiB) |
| `WithHash` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `105` (105 B) | `1` | `22` (22 B) | `83` (83 B) | `459.770x` | `0.829x` | `380.952x` | `10000` (9.77 KiB) |
| `HitColor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `10000` (9.77 KiB) | `121` (121 B) | `139` (139 B) | `1` | `28` (28 B) | `111` (111 B) | `82.645x` | `0.871x` | `71.942x` | `10000` (9.77 KiB) |
| `LocalEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `58525` (57.15 KiB) | `33468` (32.68 KiB) | `1` | `17110` (16.71 KiB) | `16358` (15.97 KiB) | `1.367x` | `1.749x` | `2.390x` | `190000` (185.55 KiB) |
| `Age` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `1981` (1.93 KiB) | `1371` (1.34 KiB) | `1` | `42` (42 B) | `1329` (1.30 KiB) | `20.192x` | `1.445x` | `29.176x` | `17412` (17.00 KiB) |
| `Sex` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `1433` (1.40 KiB) | `935` (935 B) | `1` | `30` (30 B) | `905` (905 B) | `27.913x` | `1.533x` | `42.781x` | `10000` (9.77 KiB) |
| `Income` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `1474` (1.44 KiB) | `981` (981 B) | `1` | `34` (34 B) | `947` (947 B) | `27.137x` | `1.503x` | `40.775x` | `10000` (9.77 KiB) |
| `Interests` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `4556` (4.45 KiB) | `2885` (2.82 KiB) | `1` | `780` (780 B) | `2105` (2.06 KiB) | `8.780x` | `1.579x` | `13.865x` | `22570` (22.04 KiB) |
| `Robotness` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `3879` (3.79 KiB) | `2762` (2.70 KiB) | `1` | `352` (352 B) | `2410` (2.35 KiB) | `10.312x` | `1.404x` | `14.482x` | `13859` (13.53 KiB) |
| `RemoteIP` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `8272` (8.08 KiB) | `7289` (7.12 KiB) | `1` | `2983` (2.91 KiB) | `4306` (4.21 KiB) | `4.836x` | `1.135x` | `5.488x` | `100205` (97.86 KiB) |
| `WindowName` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `106` (106 B) | `124` (124 B) | `1` | `30` (30 B) | `94` (94 B) | `377.358x` | `0.855x` | `322.581x` | `20010` (19.54 KiB) |
| `OpenerName` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `105` (105 B) | `1` | `22` (22 B) | `83` (83 B) | `459.770x` | `0.829x` | `380.952x` | `20000` (19.53 KiB) |
| `HistoryLength` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `99` (99 B) | `117` (117 B) | `1` | `26` (26 B) | `91` (91 B) | `404.040x` | `0.846x` | `341.880x` | `19996` (19.53 KiB) |
| `BrowserLanguage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `20156` (19.68 KiB) | `432` (432 B) | `353` (353 B) | `1` | `49` (49 B) | `304` (304 B) | `46.657x` | `1.224x` | `57.099x` | `20156` (19.68 KiB) |
| `BrowserCountry` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `32430` (31.67 KiB) | `1096` (1.07 KiB) | `653` (653 B) | `1` | `44` (44 B) | `609` (609 B) | `29.589x` | `1.678x` | `49.663x` | `38645` (37.74 KiB) |
| `SocialNetwork` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `0` (0 B) | `71` (71 B) | `89` (89 B) | `1` | `22` (22 B) | `67` (67 B) | `0.000x` | `0.798x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `0` (0 B) | `71` (71 B) | `89` (89 B) | `1` | `22` (22 B) | `67` (67 B) | `0.000x` | `0.798x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `105` (105 B) | `1` | `22` (22 B) | `83` (83 B) | `459.770x` | `0.829x` | `380.952x` | `10000` (9.77 KiB) |
| `SendTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `105` (105 B) | `1` | `22` (22 B) | `83` (83 B) | `459.770x` | `0.829x` | `380.952x` | `10000` (9.77 KiB) |
| `DNSTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `2152` (2.10 KiB) | `1113` (1.09 KiB) | `1` | `304` (304 B) | `809` (809 B) | `18.587x` | `1.934x` | `35.939x` | `10125` (9.89 KiB) |
| `ConnectTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `3896` (3.80 KiB) | `2001` (1.95 KiB) | `1` | `536` (536 B) | `1465` (1.43 KiB) | `10.267x` | `1.947x` | `19.990x` | `10493` (10.25 KiB) |
| `ResponseStartTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `13242` (12.93 KiB) | `9718` (9.49 KiB) | `1` | `3379` (3.30 KiB) | `6339` (6.19 KiB) | `3.021x` | `1.363x` | `4.116x` | `16080` (15.70 KiB) |
| `ResponseEndTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `11871` (11.59 KiB) | `7911` (7.73 KiB) | `1` | `2199` (2.15 KiB) | `5712` (5.58 KiB) | `3.370x` | `1.501x` | `5.056x` | `13892` (13.57 KiB) |
| `FetchTiming` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `10559` (10.31 KiB) | `6096` (5.95 KiB) | `1` | `2099` (2.05 KiB) | `3997` (3.90 KiB) | `3.788x` | `1.732x` | `6.562x` | `11851` (11.57 KiB) |
| `SocialSourceNetworkID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `102` (102 B) | `120` (120 B) | `1` | `26` (26 B) | `94` (94 B) | `392.157x` | `0.850x` | `333.333x` | `10000` (9.77 KiB) |
| `SocialSourcePage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `42` (42 B) | `102` (102 B) | `120` (120 B) | `1` | `33` (33 B) | `87` (87 B) | `0.412x` | `0.850x` | `0.350x` | `42` (42 B) |
| `ParamPrice` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `107` (107 B) | `125` (125 B) | `1` | `26` (26 B) | `99` (99 B) | `747.664x` | `0.856x` | `640.000x` | `10000` (9.77 KiB) |
| `ParamOrderID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `0` (0 B) | `71` (71 B) | `89` (89 B) | `1` | `22` (22 B) | `67` (67 B) | `0.000x` | `0.798x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `30000` (29.30 KiB) | `86` (86 B) | `104` (104 B) | `1` | `25` (25 B) | `79` (79 B) | `348.837x` | `0.827x` | `288.462x` | `30000` (29.30 KiB) |
| `ParamCurrencyID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `105` (105 B) | `1` | `22` (22 B) | `83` (83 B) | `459.770x` | `0.829x` | `380.952x` | `10000` (9.77 KiB) |
| `OpenstatServiceName` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `560` (560 B) | `202` (202 B) | `220` (220 B) | `1` | `42` (42 B) | `178` (178 B) | `2.772x` | `0.918x` | `2.545x` | `560` (560 B) |
| `OpenstatCampaignID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `272` (272 B) | `213` (213 B) | `231` (231 B) | `1` | `47` (47 B) | `184` (184 B) | `1.277x` | `0.922x` | `1.177x` | `272` (272 B) |
| `OpenstatAdID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `358` (358 B) | `340` (340 B) | `288` (288 B) | `1` | `106` (106 B) | `182` (182 B) | `1.053x` | `1.181x` | `1.243x` | `358` (358 B) |
| `OpenstatSourceID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `692` (692 B) | `250` (250 B) | `268` (268 B) | `1` | `68` (68 B) | `200` (200 B) | `2.768x` | `0.933x` | `2.582x` | `692` (692 B) |
| `UTMSource` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `1621` (1.58 KiB) | `509` (509 B) | `379` (379 B) | `1` | `77` (77 B) | `302` (302 B) | `3.185x` | `1.343x` | `4.277x` | `1621` (1.58 KiB) |
| `UTMMedium` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `502` (502 B) | `377` (377 B) | `307` (307 B) | `1` | `47` (47 B) | `260` (260 B) | `1.332x` | `1.228x` | `1.635x` | `502` (502 B) |
| `UTMCampaign` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `3304` (3.23 KiB) | `637` (637 B) | `491` (491 B) | `1` | `142` (142 B) | `349` (349 B) | `5.187x` | `1.297x` | `6.729x` | `3304` (3.23 KiB) |
| `UTMContent` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `474` (474 B) | `203` (203 B) | `221` (221 B) | `1` | `32` (32 B) | `189` (189 B) | `2.335x` | `0.919x` | `2.145x` | `474` (474 B) |
| `UTMTerm` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `1288` (1.26 KiB) | `288` (288 B) | `306` (306 B) | `1` | `58` (58 B) | `248` (248 B) | `4.472x` | `0.941x` | `4.209x` | `1288` (1.26 KiB) |
| `FromTag` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `0` (0 B) | `71` (71 B) | `89` (89 B) | `1` | `22` (22 B) | `67` (67 B) | `0.000x` | `0.798x` | `0.000x` | `0` (0 B) |
| `HasGCLID` | `int16` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `398` (398 B) | `418` (418 B) | `1` | `26` (26 B) | `392` (392 B) | `100.503x` | `0.952x` | `95.694x` | `10000` (9.77 KiB) |
| `RefererHash` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `33011` (32.24 KiB) | `29820` (29.12 KiB) | `1` | `18163` (17.74 KiB) | `11657` (11.38 KiB) | `2.423x` | `1.107x` | `2.683x` | `193030` (188.51 KiB) |
| `URLHash` | `int64` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `36727` (35.87 KiB) | `35278` (34.45 KiB) | `1` | `21707` (21.20 KiB) | `13571` (13.25 KiB) | `2.178x` | `1.041x` | `2.268x` | `193212` (188.68 KiB) |
| `CLID` | `int32` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `87` (87 B) | `105` (105 B) | `1` | `22` (22 B) | `83` (83 B) | `459.770x` | `0.829x` | `380.952x` | `10000` (9.77 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/10000_rows/parquet/page_encoding_distribution/rle-dict-zstd/part-00000.parquet`: `10000` rows, `769307` file bytes (751.28 KiB), `7711890` physical bytes (7.35 MiB), `1637809` encoded bytes (1.56 MiB), `751801` compressed data bytes (734.18 KiB)
