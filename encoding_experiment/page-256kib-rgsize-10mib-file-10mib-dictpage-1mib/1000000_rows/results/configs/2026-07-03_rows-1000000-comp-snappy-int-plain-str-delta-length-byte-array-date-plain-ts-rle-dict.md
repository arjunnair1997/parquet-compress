# ClickBench Parquet Experiment

- Started: `2026-07-03T15:28:58-04:00`
- Write elapsed: `11.622s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `719157844` (685.84 MiB)
- Compressed column data bytes after codec compression: `134982207` (128.73 MiB)
- Parquet file bytes: `135957107` (129.66 MiB)
- Physical/encoded ratio: `0.991x`
- Encoded/compressed-data ratio: `5.328x`
- Physical/compressed-data ratio: `5.278x`
- Physical/parquet-file ratio: `5.240x`
- Files: `31`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `delta-length-byte-array`
- Date encoding: `plain`
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
- Elapsed: `6.862s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004952` (7.63 MiB) | `8005384` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `262184` (256.04 KiB) | `0.999x` | `15.271x` | `15.256x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:502` | `1000000` | `138409995` (132.00 MiB) | `140028969` (133.54 MiB) | `21331689` (20.34 MiB) | `0.988x` | `6.564x` | `6.488x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `204452` (199.66 KiB) | `0.999x` | `19.584x` | `19.564x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `7358904` (7.02 MiB) | `5537371` (5.28 MiB) | `1.087x` | `1.329x` | `1.445x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003895` (3.82 MiB) | `204449` (199.66 KiB) | `0.999x` | `19.584x` | `19.565x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `204474` (199.68 KiB) | `0.999x` | `19.581x` | `19.562x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003960` (3.82 MiB) | `719550` (702.69 KiB) | `0.999x` | `5.565x` | `5.559x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `396840` (387.54 KiB) | `0.999x` | `10.089x` | `10.080x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004954` (7.63 MiB) | `1085028` (1.03 MiB) | `0.999x` | `7.378x` | `7.373x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `204280` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003895` (3.82 MiB) | `318147` (310.69 KiB) | `0.999x` | `12.585x` | `12.573x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `342318` (334.29 KiB) | `0.999x` | `11.696x` | `11.685x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:340` | `1000000` | `88562192` (84.46 MiB) | `89783900` (85.62 MiB) | `20794298` (19.83 MiB) | `0.986x` | `4.318x` | `4.259x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:294` | `1000000` | `79583339` (75.90 MiB) | `80833557` (77.09 MiB) | `19469820` (18.57 MiB) | `0.985x` | `4.152x` | `4.088x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003931` (3.82 MiB) | `492079` (480.55 KiB) | `0.999x` | `8.137x` | `8.129x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003929` (3.82 MiB) | `509257` (497.32 KiB) | `0.999x` | `7.862x` | `7.855x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003916` (3.82 MiB) | `458275` (447.53 KiB) | `0.999x` | `8.737x` | `8.728x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `272038` (265.66 KiB) | `0.999x` | `14.718x` | `14.704x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `232368` (226.92 KiB) | `0.999x` | `17.231x` | `17.214x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003904` (3.82 MiB) | `373462` (364.71 KiB) | `0.999x` | `10.721x` | `10.711x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `373298` (364.55 KiB) | `0.999x` | `10.726x` | `10.715x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `283055` (276.42 KiB) | `0.999x` | `14.145x` | `14.132x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `245459` (239.71 KiB) | `0.999x` | `16.312x` | `16.296x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `328573` (320.87 KiB) | `0.999x` | `12.186x` | `12.174x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `3354477` (3.20 MiB) | `3706973` (3.54 MiB) | `433186` (423.03 KiB) | `0.905x` | `8.557x` | `7.744x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `219224` (214.09 KiB) | `0.999x` | `18.264x` | `18.246x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `217408` (212.31 KiB) | `0.999x` | `18.417x` | `18.399x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `365318` (356.76 KiB) | `0.999x` | `10.960x` | `10.949x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `3767530` (3.59 MiB) | `4017260` (3.83 MiB) | `328684` (320.98 KiB) | `0.938x` | `12.222x` | `11.462x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `204952` (200.15 KiB) | `0.999x` | `19.536x` | `19.517x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `205061` (200.25 KiB) | `0.999x` | `19.525x` | `19.506x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `218596` (213.47 KiB) | `0.999x` | `18.316x` | `18.299x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `215513` (210.46 KiB) | `0.999x` | `18.578x` | `18.560x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `81583` (79.67 KiB) | `231385` (225.96 KiB) | `42795` (41.79 KiB) | `0.353x` | `5.407x` | `1.906x` | `81583` (79.67 KiB) |
| `Params` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `0` (0 B) | `41920` (40.94 KiB) | `4928` (4.81 KiB) | `0.000x` | `8.506x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003934` (3.82 MiB) | `561210` (548.06 KiB) | `0.999x` | `7.134x` | `7.127x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003931` (3.82 MiB) | `526993` (514.64 KiB) | `0.999x` | `7.598x` | `7.590x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003896` (3.82 MiB) | `293267` (286.39 KiB) | `0.999x` | `13.653x` | `13.639x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `3528017` (3.36 MiB) | `4223971` (4.03 MiB) | `1014938` (991.15 KiB) | `0.835x` | `4.162x` | `3.476x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `220431` (215.26 KiB) | `0.999x` | `18.164x` | `18.146x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003933` (3.82 MiB) | `505940` (494.08 KiB) | `0.999x` | `7.914x` | `7.906x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003931` (3.82 MiB) | `516888` (504.77 KiB) | `0.999x` | `7.746x` | `7.739x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003940` (3.82 MiB) | `551924` (538.99 KiB) | `0.999x` | `7.255x` | `7.247x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `285910` (279.21 KiB) | `0.999x` | `14.004x` | `13.990x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `7237357` (6.90 MiB) | `5466410` (5.21 MiB) | `1.105x` | `1.324x` | `1.463x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `287412` (280.68 KiB) | `0.999x` | `13.931x` | `13.917x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `269878` (263.55 KiB) | `0.999x` | `14.836x` | `14.822x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `305072` (297.92 KiB) | `0.999x` | `13.124x` | `13.112x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `204317` (199.53 KiB) | `0.999x` | `19.597x` | `19.577x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:92` | `1000000` | `13587860` (12.96 MiB) | `13657641` (13.02 MiB) | `700639` (684.22 KiB) | `0.995x` | `19.493x` | `19.394x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `205287` (200.48 KiB) | `0.999x` | `19.504x` | `19.485x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `256598` (250.58 KiB) | `0.999x` | `15.604x` | `15.589x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `206776` (201.93 KiB) | `0.999x` | `19.363x` | `19.345x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003905` (3.82 MiB) | `235490` (229.97 KiB) | `0.999x` | `17.002x` | `16.986x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004953` (7.63 MiB) | `1151729` (1.10 MiB) | `0.999x` | `6.950x` | `6.946x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:146` | `1000000` | `27797671` (26.51 MiB) | `28794521` (27.46 MiB) | `7055129` (6.73 MiB) | `0.965x` | `4.081x` | `3.940x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003962` (3.82 MiB) | `3688307` (3.52 MiB) | `0.999x` | `1.086x` | `1.085x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `204280` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `204280` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `204280` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `295773` (288.84 KiB) | `0.999x` | `13.537x` | `13.524x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `204280` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `1000000` (976.56 KiB) | `1043360` (1018.91 KiB) | `77220` (75.41 KiB) | `0.958x` | `13.512x` | `12.950x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `7355536` (7.01 MiB) | `5537019` (5.28 MiB) | `1.088x` | `1.328x` | `1.445x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `334192` (326.36 KiB) | `0.999x` | `11.981x` | `11.969x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `302734` (295.64 KiB) | `0.999x` | `13.226x` | `13.213x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003896` (3.82 MiB) | `318400` (310.94 KiB) | `0.999x` | `12.575x` | `12.563x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `385796` (376.75 KiB) | `0.999x` | `10.378x` | `10.368x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `381982` (373.03 KiB) | `0.999x` | `10.482x` | `10.472x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003961` (3.82 MiB) | `707420` (690.84 KiB) | `0.999x` | `5.660x` | `5.654x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003906` (3.82 MiB) | `295428` (288.50 KiB) | `0.999x` | `13.553x` | `13.540x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `204279` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003904` (3.82 MiB) | `258562` (252.50 KiB) | `0.999x` | `15.485x` | `15.470x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `2001192` (1.91 MiB) | `2051409` (1.96 MiB) | `125388` (122.45 KiB) | `0.976x` | `16.360x` | `15.960x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `3325142` (3.17 MiB) | `3639390` (3.47 MiB) | `353048` (344.77 KiB) | `0.914x` | `10.308x` | `9.418x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `0` (0 B) | `41920` (40.94 KiB) | `4928` (4.81 KiB) | `0.000x` | `8.506x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `0` (0 B) | `41920` (40.94 KiB) | `4928` (4.81 KiB) | `0.000x` | `8.506x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `204280` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003908` (3.82 MiB) | `269430` (263.12 KiB) | `0.999x` | `14.861x` | `14.846x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `327371` (319.70 KiB) | `0.999x` | `12.230x` | `12.219x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003938` (3.82 MiB) | `565072` (551.83 KiB) | `0.999x` | `7.086x` | `7.079x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003961` (3.82 MiB) | `1724442` (1.64 MiB) | `0.999x` | `2.322x` | `2.320x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003960` (3.82 MiB) | `1282848` (1.22 MiB) | `0.999x` | `3.121x` | `3.118x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003959` (3.82 MiB) | `810184` (791.20 KiB) | `0.999x` | `4.942x` | `4.937x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `204670` (199.87 KiB) | `0.999x` | `19.563x` | `19.544x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `1024` (1.00 KiB) | `46467` (45.38 KiB) | `7349` (7.18 KiB) | `0.022x` | `6.323x` | `0.139x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004894` (7.63 MiB) | `405268` (395.77 KiB) | `0.999x` | `19.752x` | `19.740x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `0` (0 B) | `41920` (40.94 KiB) | `4928` (4.81 KiB) | `0.000x` | `8.506x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `3000000` (2.86 MiB) | `3044430` (2.90 MiB) | `157428` (153.74 KiB) | `0.985x` | `19.339x` | `19.056x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `204280` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `58030` (56.67 KiB) | `216969` (211.88 KiB) | `37190` (36.32 KiB) | `0.267x` | `5.834x` | `1.560x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `22051` (21.53 KiB) | `123175` (120.29 KiB) | `28880` (28.20 KiB) | `0.179x` | `4.265x` | `0.764x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `25445` (24.85 KiB) | `130853` (127.79 KiB) | `34714` (33.90 KiB) | `0.194x` | `3.769x` | `0.733x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `48191` (47.06 KiB) | `156140` (152.48 KiB) | `24637` (24.06 KiB) | `0.309x` | `6.338x` | `1.956x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `49433` (48.27 KiB) | `189064` (184.63 KiB) | `41779` (40.80 KiB) | `0.261x` | `4.525x` | `1.183x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `16873` (16.48 KiB) | `134295` (131.15 KiB) | `28750` (28.08 KiB) | `0.126x` | `4.671x` | `0.587x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `91870` (89.72 KiB) | `254335` (248.37 KiB) | `57676` (56.32 KiB) | `0.361x` | `4.410x` | `1.593x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `13001` (12.70 KiB) | `94737` (92.52 KiB) | `22198` (21.68 KiB) | `0.137x` | `4.268x` | `0.586x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `28101` (27.44 KiB) | `129434` (126.40 KiB) | `25224` (24.63 KiB) | `0.217x` | `5.131x` | `1.114x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `delta-length-byte-array` | `DELTA_LENGTH_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_LENGTH_BYTE_ARRAY:62` | `1000000` | `45607` (44.54 KiB) | `210615` (205.68 KiB) | `46123` (45.04 KiB) | `0.217x` | `4.566x` | `0.989x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `213684` (208.68 KiB) | `0.999x` | `18.737x` | `18.719x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004950` (7.63 MiB) | `3642160` (3.47 MiB) | `0.999x` | `2.198x` | `2.196x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004954` (7.63 MiB) | `4388336` (4.19 MiB) | `0.999x` | `1.824x` | `1.823x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204731` (199.93 KiB) | `0.999x` | `19.557x` | `19.538x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00000.parquet`: `33055` rows, `4289812` file bytes (4.09 MiB), `25668101` physical bytes (24.48 MiB), `25758103` encoded bytes (24.56 MiB), `4257783` compressed data bytes (4.06 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00001.parquet`: `33101` rows, `4239356` file bytes (4.04 MiB), `25453482` physical bytes (24.27 MiB), `25536685` encoded bytes (24.35 MiB), `4206947` compressed data bytes (4.01 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00002.parquet`: `32620` rows, `4234392` file bytes (4.04 MiB), `25543984` physical bytes (24.36 MiB), `25632662` encoded bytes (24.45 MiB), `4201489` compressed data bytes (4.01 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00003.parquet`: `32872` rows, `4230781` file bytes (4.03 MiB), `25493099` physical bytes (24.31 MiB), `25579417` encoded bytes (24.39 MiB), `4198320` compressed data bytes (4.00 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00004.parquet`: `33119` rows, `4246684` file bytes (4.05 MiB), `25705063` physical bytes (24.51 MiB), `25788457` encoded bytes (24.59 MiB), `4213441` compressed data bytes (4.02 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00005.parquet`: `33331` rows, `4206659` file bytes (4.01 MiB), `25722284` physical bytes (24.53 MiB), `25804003` encoded bytes (24.61 MiB), `4174155` compressed data bytes (3.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00006.parquet`: `32899` rows, `4200646` file bytes (4.01 MiB), `25308693` physical bytes (24.14 MiB), `25395339` encoded bytes (24.22 MiB), `4168398` compressed data bytes (3.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00007.parquet`: `32910` rows, `4206825` file bytes (4.01 MiB), `25530870` physical bytes (24.35 MiB), `25617651` encoded bytes (24.43 MiB), `4174282` compressed data bytes (3.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00008.parquet`: `32958` rows, `4210099` file bytes (4.02 MiB), `25519062` physical bytes (24.34 MiB), `25606876` encoded bytes (24.42 MiB), `4177142` compressed data bytes (3.98 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00009.parquet`: `33097` rows, `4177770` file bytes (3.98 MiB), `25545003` physical bytes (24.36 MiB), `25624912` encoded bytes (24.44 MiB), `4145406` compressed data bytes (3.95 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00010.parquet`: `33018` rows, `4223209` file bytes (4.03 MiB), `25489422` physical bytes (24.31 MiB), `25575430` encoded bytes (24.39 MiB), `4191031` compressed data bytes (4.00 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00011.parquet`: `33004` rows, `4148037` file bytes (3.96 MiB), `25362314` physical bytes (24.19 MiB), `25443074` encoded bytes (24.26 MiB), `4115615` compressed data bytes (3.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00012.parquet`: `33149` rows, `4171210` file bytes (3.98 MiB), `25550053` physical bytes (24.37 MiB), `25637094` encoded bytes (24.45 MiB), `4138853` compressed data bytes (3.95 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00013.parquet`: `33146` rows, `4217571` file bytes (4.02 MiB), `25527626` physical bytes (24.35 MiB), `25613484` encoded bytes (24.43 MiB), `4185507` compressed data bytes (3.99 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00014.parquet`: `32803` rows, `4139677` file bytes (3.95 MiB), `25185628` physical bytes (24.02 MiB), `25265551` encoded bytes (24.10 MiB), `4107371` compressed data bytes (3.92 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00015.parquet`: `33292` rows, `4511521` file bytes (4.30 MiB), `24062050` physical bytes (22.95 MiB), `24272345` encoded bytes (23.15 MiB), `4480168` compressed data bytes (4.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00016.parquet`: `32948` rows, `4622280` file bytes (4.41 MiB), `22669871` physical bytes (21.62 MiB), `22937272` encoded bytes (21.87 MiB), `4591967` compressed data bytes (4.38 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00017.parquet`: `32161` rows, `4678353` file bytes (4.46 MiB), `22368658` physical bytes (21.33 MiB), `22676472` encoded bytes (21.63 MiB), `4648022` compressed data bytes (4.43 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00018.parquet`: `32435` rows, `4630498` file bytes (4.42 MiB), `20607649` physical bytes (19.65 MiB), `20988274` encoded bytes (20.02 MiB), `4600104` compressed data bytes (4.39 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00019.parquet`: `31838` rows, `4698217` file bytes (4.48 MiB), `20733218` physical bytes (19.77 MiB), `21097440` encoded bytes (20.12 MiB), `4667478` compressed data bytes (4.45 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00020.parquet`: `31811` rows, `4680882` file bytes (4.46 MiB), `20460999` physical bytes (19.51 MiB), `20828943` encoded bytes (19.86 MiB), `4650551` compressed data bytes (4.44 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00021.parquet`: `32273` rows, `4709930` file bytes (4.49 MiB), `20700861` physical bytes (19.74 MiB), `21073145` encoded bytes (20.10 MiB), `4678992` compressed data bytes (4.46 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00022.parquet`: `32100` rows, `4727488` file bytes (4.51 MiB), `20683254` physical bytes (19.73 MiB), `21059369` encoded bytes (20.08 MiB), `4696963` compressed data bytes (4.48 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00023.parquet`: `32495` rows, `4668556` file bytes (4.45 MiB), `20743866` physical bytes (19.78 MiB), `21119553` encoded bytes (20.14 MiB), `4637846` compressed data bytes (4.42 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00024.parquet`: `32618` rows, `4692926` file bytes (4.48 MiB), `20855513` physical bytes (19.89 MiB), `21227681` encoded bytes (20.24 MiB), `4662550` compressed data bytes (4.45 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00025.parquet`: `31968` rows, `4783911` file bytes (4.56 MiB), `20876102` physical bytes (19.91 MiB), `21245300` encoded bytes (20.26 MiB), `4753118` compressed data bytes (4.53 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00026.parquet`: `32357` rows, `4634393` file bytes (4.42 MiB), `20560823` physical bytes (19.61 MiB), `20933922` encoded bytes (19.96 MiB), `4604283` compressed data bytes (4.39 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00027.parquet`: `32194` rows, `4670484` file bytes (4.45 MiB), `20580575` physical bytes (19.63 MiB), `20953420` encoded bytes (19.98 MiB), `4639881` compressed data bytes (4.42 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00028.parquet`: `32378` rows, `4603795` file bytes (4.39 MiB), `20636248` physical bytes (19.68 MiB), `21007985` encoded bytes (20.03 MiB), `4573204` compressed data bytes (4.36 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00029.parquet`: `32515` rows, `4640982` file bytes (4.43 MiB), `20592948` physical bytes (19.64 MiB), `20967491` encoded bytes (20.00 MiB), `4610651` compressed data bytes (4.40 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-rle-dict/part-00030.parquet`: `19535` rows, `2860163` file bytes (2.73 MiB), `12661305` physical bytes (12.07 MiB), `12890494` encoded bytes (12.29 MiB), `2830689` compressed data bytes (2.70 MiB)
