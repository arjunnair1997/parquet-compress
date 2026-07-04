# ClickBench Parquet Experiment

- Started: `2026-07-03T15:25:52-04:00`
- Write elapsed: `112ms`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/10000_rows/parquet/rows-10000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain`
- Rows: `10000`
- Source TSV bytes for rows, reference only: `8365486` (7.98 MiB)
- Parquet physical bytes before page encoding: `7711890` (7.35 MiB)
- Encoded column bytes before codec compression: `4542751` (4.33 MiB)
- Compressed column data bytes after codec compression: `950487` (928.21 KiB)
- Parquet file bytes: `967146` (944.48 KiB)
- Physical/encoded ratio: `1.698x`
- Encoded/compressed-data ratio: `4.779x`
- Physical/compressed-data ratio: `8.114x`
- Physical/parquet-file ratio: `7.974x`
- Files: `1`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `rle-dict`
- Date encoding: `plain`
- Timestamp encoding: `plain`
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
- Rows read and compared: `10000`
- Files read: `1`
- Elapsed: `71ms`
- Source TSV bytes checked: `8365486` (7.98 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-10000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-10000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `80087` (78.21 KiB) | `0.999x` | `1.000x` | `0.999x` | `190000` (185.55 KiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2641` (2.58 KiB) | `0.998x` | `15.170x` | `15.146x` | `10000` (9.77 KiB) |
| `Title` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `2200605` (2.10 MiB) | `545453` (532.67 KiB) | `185786` (181.43 KiB) | `4.034x` | `2.936x` | `11.845x` | `2200673` (2.10 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2072` (2.02 KiB) | `0.998x` | `19.335x` | `19.305x` | `10000` (9.77 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `35183` (34.36 KiB) | `0.999x` | `2.276x` | `2.274x` | `190000` (185.55 KiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2072` (2.02 KiB) | `0.998x` | `19.335x` | `19.305x` | `100000` (97.66 KiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2079` (2.03 KiB) | `0.998x` | `19.270x` | `19.240x` | `20000` (19.53 KiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40062` (39.12 KiB) | `6760` (6.60 KiB) | `0.998x` | `5.926x` | `5.917x` | `100283` (97.93 KiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3690` (3.60 KiB) | `0.998x` | `10.857x` | `10.840x` | `24459` (23.89 KiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `10611` (10.36 KiB) | `0.999x` | `7.547x` | `7.539x` | `199987` (195.30 KiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3137` (3.06 KiB) | `0.998x` | `12.771x` | `12.751x` | `16130` (15.75 KiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3477` (3.40 KiB) | `0.998x` | `11.522x` | `11.504x` | `11313` (11.05 KiB) |
| `URL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `760012` (742.20 KiB) | `247077` (241.29 KiB) | `87979` (85.92 KiB) | `3.076x` | `2.808x` | `8.639x` | `760022` (742.21 KiB) |
| `Referer` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `797798` (779.10 KiB) | `233929` (228.45 KiB) | `100692` (98.33 KiB) | `3.410x` | `2.323x` | `7.923x` | `797822` (779.12 KiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `7387` (7.21 KiB) | `0.998x` | `5.423x` | `5.415x` | `10000` (9.77 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `4754` (4.64 KiB) | `0.998x` | `8.427x` | `8.414x` | `47318` (46.21 KiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `4762` (4.65 KiB) | `0.998x` | `8.413x` | `8.400x` | `28586` (27.92 KiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2267` (2.21 KiB) | `0.998x` | `17.672x` | `17.644x` | `49424` (48.27 KiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2300` (2.25 KiB) | `0.998x` | `17.419x` | `17.391x` | `29545` (28.85 KiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3638` (3.55 KiB) | `0.998x` | `11.012x` | `10.995x` | `39093` (38.18 KiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3619` (3.53 KiB) | `0.998x` | `11.070x` | `11.053x` | `34257` (33.45 KiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2730` (2.67 KiB) | `0.998x` | `14.675x` | `14.652x` | `19813` (19.35 KiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2622` (2.56 KiB) | `0.998x` | `15.280x` | `15.256x` | `18909` (18.47 KiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3219` (3.14 KiB) | `0.998x` | `12.446x` | `12.426x` | `11787` (11.51 KiB) |
| `FlashMinor2` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `37170` (36.30 KiB) | `2793` (2.73 KiB) | `2124` (2.07 KiB) | `13.308x` | `1.315x` | `17.500x` | `37170` (36.30 KiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2270` (2.22 KiB) | `0.998x` | `17.649x` | `17.621x` | `10000` (9.77 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2241` (2.19 KiB) | `0.998x` | `17.877x` | `17.849x` | `10000` (9.77 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3641` (3.56 KiB) | `0.998x` | `11.003x` | `10.986x` | `18125` (17.70 KiB) |
| `UserAgentMinor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `36930` (36.06 KiB) | `1985` (1.94 KiB) | `1522` (1.49 KiB) | `18.605x` | `1.304x` | `24.264x` | `37108` (36.24 KiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2157` (2.11 KiB) | `0.998x` | `18.573x` | `18.544x` | `10000` (9.77 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2132` (2.08 KiB) | `0.998x` | `18.791x` | `18.762x` | `10000` (9.77 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2201` (2.15 KiB) | `0.998x` | `18.202x` | `18.174x` | `10000` (9.77 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2188` (2.14 KiB) | `0.998x` | `18.310x` | `18.282x` | `10011` (9.78 KiB) |
| `MobilePhoneModel` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `1140` (1.11 KiB) | `216` (216 B) | `222` (222 B) | `5.278x` | `0.973x` | `5.135x` | `1140` (1.11 KiB) |
| `Params` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `0` (0 B) | `71` (71 B) | `75` (75 B) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `5433` (5.31 KiB) | `0.998x` | `7.374x` | `7.362x` | `68529` (66.92 KiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `4728` (4.62 KiB) | `0.998x` | `8.474x` | `8.460x` | `17645` (17.23 KiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3534` (3.45 KiB) | `0.998x` | `11.336x` | `11.319x` | `10166` (9.93 KiB) |
| `SearchPhrase` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `65246` (63.72 KiB) | `28478` (27.81 KiB) | `14059` (13.73 KiB) | `2.291x` | `2.026x` | `4.641x` | `65250` (63.72 KiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2434` (2.38 KiB) | `0.998x` | `16.460x` | `16.434x` | `10142` (9.90 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `7248` (7.08 KiB) | `0.998x` | `5.527x` | `5.519x` | `10000` (9.77 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `5249` (5.13 KiB) | `0.998x` | `7.633x` | `7.620x` | `37695` (36.81 KiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `4961` (4.84 KiB) | `0.998x` | `8.076x` | `8.063x` | `30023` (29.32 KiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40062` (39.12 KiB) | `2656` (2.59 KiB) | `0.998x` | `15.084x` | `15.060x` | `29805` (29.11 KiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80079` (78.20 KiB) | `33528` (32.74 KiB) | `0.999x` | `2.388x` | `2.386x` | `190000` (185.55 KiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2756` (2.69 KiB) | `0.998x` | `14.537x` | `14.514x` | `10017` (9.78 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2575` (2.51 KiB) | `0.998x` | `15.558x` | `15.534x` | `10000` (9.77 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2908` (2.84 KiB) | `0.998x` | `13.777x` | `13.755x` | `26504` (25.88 KiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `PageCharset` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `196540` (191.93 KiB) | `280` (280 B) | `286` (286 B) | `701.929x` | `0.979x` | `687.203x` | `196540` (191.93 KiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40062` (39.12 KiB) | `2132` (2.08 KiB) | `0.998x` | `18.791x` | `18.762x` | `39532` (38.61 KiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `10655` (10.41 KiB) | `0.999x` | `7.516x` | `7.508x` | `176860` (172.71 KiB) |
| `OriginalURL` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `74750` (73.00 KiB) | `32331` (31.57 KiB) | `11044` (10.79 KiB) | `2.312x` | `2.927x` | `6.768x` | `74750` (73.00 KiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `34415` (33.61 KiB) | `0.998x` | `1.164x` | `1.162x` | `88404` (86.33 KiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2109` (2.06 KiB) | `0.998x` | `18.996x` | `18.966x` | `10000` (9.77 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `HitColor` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `10000` (9.77 KiB) | `121` (121 B) | `125` (125 B) | `82.645x` | `0.968x` | `80.000x` | `10000` (9.77 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `35180` (34.36 KiB) | `0.999x` | `2.276x` | `2.274x` | `190000` (185.55 KiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3434` (3.35 KiB) | `0.998x` | `11.667x` | `11.648x` | `17412` (17.00 KiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3116` (3.04 KiB) | `0.998x` | `12.857x` | `12.837x` | `10000` (9.77 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3172` (3.10 KiB) | `0.998x` | `12.630x` | `12.610x` | `10000` (9.77 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3809` (3.72 KiB) | `0.998x` | `10.518x` | `10.501x` | `22570` (22.04 KiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `3720` (3.63 KiB) | `0.998x` | `10.770x` | `10.753x` | `13859` (13.53 KiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `6807` (6.65 KiB) | `0.998x` | `5.886x` | `5.876x` | `100205` (97.86 KiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2086` (2.04 KiB) | `0.998x` | `19.206x` | `19.175x` | `20010` (19.54 KiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `20000` (19.53 KiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2080` (2.03 KiB) | `0.998x` | `19.261x` | `19.231x` | `19996` (19.53 KiB) |
| `BrowserLanguage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `20156` (19.68 KiB) | `433` (433 B) | `405` (405 B) | `46.550x` | `1.069x` | `49.768x` | `20156` (19.68 KiB) |
| `BrowserCountry` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `32430` (31.67 KiB) | `1096` (1.07 KiB) | `917` (917 B) | `29.589x` | `1.195x` | `35.365x` | `38645` (37.74 KiB) |
| `SocialNetwork` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `0` (0 B) | `71` (71 B) | `75` (75 B) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `0` (0 B) | `71` (71 B) | `75` (75 B) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2693` (2.63 KiB) | `0.998x` | `14.877x` | `14.853x` | `10125` (9.89 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40062` (39.12 KiB) | `3882` (3.79 KiB) | `0.998x` | `10.320x` | `10.304x` | `10493` (10.25 KiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `11343` (11.08 KiB) | `0.998x` | `3.532x` | `3.526x` | `16080` (15.70 KiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40064` (39.12 KiB) | `9934` (9.70 KiB) | `0.998x` | `4.033x` | `4.027x` | `13892` (13.57 KiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `7294` (7.12 KiB) | `0.998x` | `5.493x` | `5.484x` | `11851` (11.57 KiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2079` (2.03 KiB) | `0.998x` | `19.270x` | `19.240x` | `10000` (9.77 KiB) |
| `SocialSourcePage` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `42` (42 B) | `102` (102 B) | `106` (106 B) | `0.412x` | `0.962x` | `0.396x` | `42` (42 B) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80079` (78.20 KiB) | `4086` (3.99 KiB) | `0.999x` | `19.598x` | `19.579x` | `10000` (9.77 KiB) |
| `ParamOrderID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `0` (0 B) | `71` (71 B) | `75` (75 B) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `30000` (29.30 KiB) | `86` (86 B) | `90` (90 B) | `348.837x` | `0.956x` | `333.333x` | `30000` (29.30 KiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |
| `OpenstatServiceName` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `560` (560 B) | `201` (201 B) | `206` (206 B) | `2.786x` | `0.976x` | `2.718x` | `560` (560 B) |
| `OpenstatCampaignID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `272` (272 B) | `213` (213 B) | `218` (218 B) | `1.277x` | `0.977x` | `1.248x` | `272` (272 B) |
| `OpenstatAdID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `358` (358 B) | `339` (339 B) | `318` (318 B) | `1.056x` | `1.066x` | `1.126x` | `358` (358 B) |
| `OpenstatSourceID` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `692` (692 B) | `250` (250 B) | `236` (236 B) | `2.768x` | `1.059x` | `2.932x` | `692` (692 B) |
| `UTMSource` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `1621` (1.58 KiB) | `508` (508 B) | `451` (451 B) | `3.191x` | `1.126x` | `3.594x` | `1621` (1.58 KiB) |
| `UTMMedium` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `502` (502 B) | `376` (376 B) | `361` (361 B) | `1.335x` | `1.042x` | `1.391x` | `502` (502 B) |
| `UTMCampaign` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `3304` (3.23 KiB) | `637` (637 B) | `587` (587 B) | `5.187x` | `1.085x` | `5.629x` | `3304` (3.23 KiB) |
| `UTMContent` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `474` (474 B) | `203` (203 B) | `208` (208 B) | `2.335x` | `0.976x` | `2.279x` | `474` (474 B) |
| `UTMTerm` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `1288` (1.26 KiB) | `288` (288 B) | `294` (294 B) | `4.472x` | `0.980x` | `4.381x` | `1288` (1.26 KiB) |
| `FromTag` | `string` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:1, DICTIONARY_PAGE/PLAIN:1` | `10000` | `0` (0 B) | `71` (71 B) | `75` (75 B) | `0.000x` | `0.947x` | `0.000x` | `0` (0 B) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2273` (2.22 KiB) | `0.998x` | `17.626x` | `17.598x` | `10000` (9.77 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `34180` (33.38 KiB) | `0.999x` | `2.343x` | `2.341x` | `193030` (188.51 KiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `80000` (78.12 KiB) | `80080` (78.20 KiB) | `38484` (37.58 KiB) | `0.999x` | `2.081x` | `2.079x` | `193212` (188.68 KiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:1` | `10000` | `40000` (39.06 KiB) | `40063` (39.12 KiB) | `2069` (2.02 KiB) | `0.998x` | `19.363x` | `19.333x` | `10000` (9.77 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/10000_rows/parquet/rows-10000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain/part-00000.parquet`: `10000` rows, `967146` file bytes (944.48 KiB), `7711890` physical bytes (7.35 MiB), `4542751` encoded bytes (4.33 MiB), `950487` compressed data bytes (928.21 KiB)
