# ClickBench Parquet Experiment

- Started: `2026-07-03T15:28:57-04:00`
- Write elapsed: `11.789s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `823411941` (785.27 MiB)
- Compressed column data bytes after codec compression: `138062083` (131.67 MiB)
- Parquet file bytes: `139039066` (132.60 MiB)
- Physical/encoded ratio: `0.865x`
- Encoded/compressed-data ratio: `5.964x`
- Physical/compressed-data ratio: `5.160x`
- Physical/parquet-file ratio: `5.124x`
- Files: `31`

## Settings

- Compression: `snappy`
- Int encoding: `plain`
- String encoding: `plain`
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
- Elapsed: `7.003s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004955` (7.63 MiB) | `8005388` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003896` (3.82 MiB) | `262152` (256.01 KiB) | `0.999x` | `15.273x` | `15.258x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:496` | `1000000` | `138409995` (132.00 MiB) | `142872716` (136.25 MiB) | `20899403` (19.93 MiB) | `0.969x` | `6.836x` | `6.623x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204457` (199.67 KiB) | `0.999x` | `19.583x` | `19.564x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `7355268` (7.01 MiB) | `5533695` (5.28 MiB) | `1.088x` | `1.329x` | `1.446x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `204457` (199.67 KiB) | `0.999x` | `19.583x` | `19.564x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `204475` (199.68 KiB) | `0.999x` | `19.581x` | `19.562x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003960` (3.82 MiB) | `719538` (702.67 KiB) | `0.999x` | `5.565x` | `5.559x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `397009` (387.70 KiB) | `0.999x` | `10.085x` | `10.075x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004954` (7.63 MiB) | `1085045` (1.03 MiB) | `0.999x` | `7.378x` | `7.373x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `204280` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `318247` (310.79 KiB) | `0.999x` | `12.581x` | `12.569x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `342368` (334.34 KiB) | `0.999x` | `11.695x` | `11.683x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:340` | `1000000` | `88562192` (84.46 MiB) | `92651814` (88.36 MiB) | `20469052` (19.52 MiB) | `0.956x` | `4.526x` | `4.327x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:294` | `1000000` | `79583339` (75.90 MiB) | `83649248` (79.77 MiB) | `19063936` (18.18 MiB) | `0.951x` | `4.388x` | `4.175x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003930` (3.82 MiB) | `492072` (480.54 KiB) | `0.999x` | `8.137x` | `8.129x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003929` (3.82 MiB) | `509210` (497.28 KiB) | `0.999x` | `7.863x` | `7.855x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003917` (3.82 MiB) | `458244` (447.50 KiB) | `0.999x` | `8.738x` | `8.729x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `272034` (265.66 KiB) | `0.999x` | `14.718x` | `14.704x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `232383` (226.94 KiB) | `0.999x` | `17.230x` | `17.213x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `373689` (364.93 KiB) | `0.999x` | `10.715x` | `10.704x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `373290` (364.54 KiB) | `0.999x` | `10.726x` | `10.716x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `283102` (276.47 KiB) | `0.999x` | `14.143x` | `14.129x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `245500` (239.75 KiB) | `0.999x` | `16.309x` | `16.293x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003897` (3.82 MiB) | `328568` (320.87 KiB) | `0.999x` | `12.186x` | `12.174x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `3354477` (3.20 MiB) | `7358155` (7.02 MiB) | `536739` (524.16 KiB) | `0.456x` | `13.709x` | `6.250x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `219251` (214.11 KiB) | `0.999x` | `18.262x` | `18.244x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `217418` (212.32 KiB) | `0.999x` | `18.416x` | `18.398x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003899` (3.82 MiB) | `365183` (356.62 KiB) | `0.999x` | `10.964x` | `10.953x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `3767530` (3.59 MiB) | `7771208` (7.41 MiB) | `467315` (456.36 KiB) | `0.485x` | `16.629x` | `8.062x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `205038` (200.23 KiB) | `0.999x` | `19.528x` | `19.509x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `205072` (200.27 KiB) | `0.999x` | `19.524x` | `19.505x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `218571` (213.45 KiB) | `0.999x` | `18.319x` | `18.301x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `215504` (210.45 KiB) | `0.999x` | `18.579x` | `18.561x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `81583` (79.67 KiB) | `4085178` (3.90 MiB) | `220095` (214.94 KiB) | `0.020x` | `18.561x` | `0.371x` | `81583` (79.67 KiB) |
| `Params` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `0` (0 B) | `4002415` (3.82 MiB) | `202792` (198.04 KiB) | `0.000x` | `19.737x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003936` (3.82 MiB) | `561193` (548.04 KiB) | `0.999x` | `7.135x` | `7.128x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003933` (3.82 MiB) | `527011` (514.66 KiB) | `0.999x` | `7.597x` | `7.590x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `293325` (286.45 KiB) | `0.999x` | `13.650x` | `13.637x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `3528017` (3.36 MiB) | `7535671` (7.19 MiB) | `1098288` (1.05 MiB) | `0.468x` | `6.861x` | `3.212x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `220403` (215.24 KiB) | `0.999x` | `18.166x` | `18.149x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003933` (3.82 MiB) | `505910` (494.05 KiB) | `0.999x` | `7.914x` | `7.907x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003930` (3.82 MiB) | `517002` (504.88 KiB) | `0.999x` | `7.745x` | `7.737x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003945` (3.82 MiB) | `552155` (539.21 KiB) | `0.999x` | `7.251x` | `7.244x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `285932` (279.23 KiB) | `0.999x` | `14.003x` | `13.989x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `7234902` (6.90 MiB) | `5464176` (5.21 MiB) | `1.106x` | `1.324x` | `1.464x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `287405` (280.67 KiB) | `0.999x` | `13.931x` | `13.918x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `269894` (263.57 KiB) | `0.999x` | `14.835x` | `14.821x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `304963` (297.82 KiB) | `0.999x` | `13.129x` | `13.116x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `204317` (199.53 KiB) | `0.999x` | `19.597x` | `19.577x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:92` | `1000000` | `13587860` (12.96 MiB) | `17595564` (16.78 MiB) | `927116` (905.39 KiB) | `0.772x` | `18.979x` | `14.656x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003900` (3.82 MiB) | `205326` (200.51 KiB) | `0.999x` | `19.500x` | `19.481x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `256619` (250.60 KiB) | `0.999x` | `15.603x` | `15.587x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `206764` (201.92 KiB) | `0.999x` | `19.365x` | `19.346x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003906` (3.82 MiB) | `235484` (229.96 KiB) | `0.999x` | `17.003x` | `16.986x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004956` (7.63 MiB) | `1151764` (1.10 MiB) | `0.999x` | `6.950x` | `6.946x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:144` | `1000000` | `27797671` (26.51 MiB) | `31860918` (30.38 MiB) | `7043550` (6.72 MiB) | `0.872x` | `4.523x` | `3.947x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003963` (3.82 MiB) | `3688296` (3.52 MiB) | `0.999x` | `1.086x` | `1.085x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `204280` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `204280` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `204280` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003906` (3.82 MiB) | `295788` (288.86 KiB) | `0.999x` | `13.536x` | `13.523x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `204280` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `1000000` (976.56 KiB) | `5003161` (4.77 MiB) | `294512` (287.61 KiB) | `0.200x` | `16.988x` | `3.395x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:62, DICTIONARY_PAGE/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `7350971` (7.01 MiB) | `5532521` (5.28 MiB) | `1.088x` | `1.329x` | `1.446x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `334115` (326.28 KiB) | `0.999x` | `11.984x` | `11.972x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `302682` (295.59 KiB) | `0.999x` | `13.228x` | `13.215x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `318371` (310.91 KiB) | `0.999x` | `12.576x` | `12.564x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003905` (3.82 MiB) | `385612` (376.57 KiB) | `0.999x` | `10.383x` | `10.373x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `381924` (372.97 KiB) | `0.999x` | `10.484x` | `10.473x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003964` (3.82 MiB) | `707357` (690.78 KiB) | `0.999x` | `5.660x` | `5.655x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003904` (3.82 MiB) | `295435` (288.51 KiB) | `0.999x` | `13.553x` | `13.539x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `204279` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003905` (3.82 MiB) | `258509` (252.45 KiB) | `0.999x` | `15.488x` | `15.473x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `2001192` (1.91 MiB) | `6004717` (5.73 MiB) | `321359` (313.83 KiB) | `0.333x` | `18.685x` | `6.227x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `3325142` (3.17 MiB) | `7328819` (6.99 MiB) | `457222` (446.51 KiB) | `0.454x` | `16.029x` | `7.272x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `0` (0 B) | `4002415` (3.82 MiB) | `202792` (198.04 KiB) | `0.000x` | `19.737x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `0` (0 B) | `4002415` (3.82 MiB) | `202792` (198.04 KiB) | `0.000x` | `19.737x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `204280` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003909` (3.82 MiB) | `269437` (263.12 KiB) | `0.999x` | `14.860x` | `14.846x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `327255` (319.58 KiB) | `0.999x` | `12.235x` | `12.223x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003935` (3.82 MiB) | `565320` (552.07 KiB) | `0.999x` | `7.083x` | `7.076x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003962` (3.82 MiB) | `1725341` (1.65 MiB) | `0.999x` | `2.321x` | `2.318x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003962` (3.82 MiB) | `1283582` (1.22 MiB) | `0.999x` | `3.119x` | `3.116x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003957` (3.82 MiB) | `810519` (791.52 KiB) | `0.999x` | `4.940x` | `4.935x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003898` (3.82 MiB) | `204772` (199.97 KiB) | `0.999x` | `19.553x` | `19.534x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `1024` (1.00 KiB) | `4004206` (3.82 MiB) | `204476` (199.68 KiB) | `0.000x` | `19.583x` | `0.005x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004892` (7.63 MiB) | `405268` (395.77 KiB) | `0.999x` | `19.752x` | `19.740x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `0` (0 B) | `4002415` (3.82 MiB) | `202792` (198.04 KiB) | `0.000x` | `19.737x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `3000000` (2.86 MiB) | `7003654` (6.68 MiB) | `354390` (346.08 KiB) | `0.428x` | `19.763x` | `8.465x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003903` (3.82 MiB) | `204280` (199.49 KiB) | `0.999x` | `19.600x` | `19.581x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `58030` (56.67 KiB) | `4062922` (3.87 MiB) | `218655` (213.53 KiB) | `0.014x` | `18.581x` | `0.265x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `22051` (21.53 KiB) | `4025885` (3.84 MiB) | `214734` (209.70 KiB) | `0.005x` | `18.748x` | `0.103x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `25445` (24.85 KiB) | `4029356` (3.84 MiB) | `219240` (214.10 KiB) | `0.006x` | `18.379x` | `0.116x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `48191` (47.06 KiB) | `4052087` (3.86 MiB) | `214439` (209.41 KiB) | `0.012x` | `18.896x` | `0.225x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `49433` (48.27 KiB) | `4053501` (3.87 MiB) | `219149` (214.01 KiB) | `0.012x` | `18.497x` | `0.226x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `16873` (16.48 KiB) | `4020198` (3.83 MiB) | `214395` (209.37 KiB) | `0.004x` | `18.751x` | `0.079x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `91870` (89.72 KiB) | `4097902` (3.91 MiB) | `230280` (224.88 KiB) | `0.022x` | `17.795x` | `0.399x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `13001` (12.70 KiB) | `4016968` (3.83 MiB) | `212425` (207.45 KiB) | `0.003x` | `18.910x` | `0.061x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `28101` (27.44 KiB) | `4033117` (3.85 MiB) | `214669` (209.64 KiB) | `0.007x` | `18.788x` | `0.131x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `45607` (44.54 KiB) | `4048574` (3.86 MiB) | `221869` (216.67 KiB) | `0.011x` | `18.248x` | `0.206x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003902` (3.82 MiB) | `213663` (208.66 KiB) | `0.999x` | `18.739x` | `18.721x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004954` (7.63 MiB) | `3641033` (3.47 MiB) | `0.999x` | `2.199x` | `2.197x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `8000000` (7.63 MiB) | `8004955` (7.63 MiB) | `4387222` (4.18 MiB) | `0.999x` | `1.825x` | `1.823x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:62` | `1000000` | `4000000` (3.81 MiB) | `4003901` (3.82 MiB) | `204743` (199.94 KiB) | `0.999x` | `19.556x` | `19.537x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00000.parquet`: `32760` rows, `4361931` file bytes (4.16 MiB), `25444697` physical bytes (24.27 MiB), `28934599` encoded bytes (27.59 MiB), `4329911` compressed data bytes (4.13 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00001.parquet`: `32781` rows, `4312621` file bytes (4.11 MiB), `25223905` physical bytes (24.06 MiB), `28722970` encoded bytes (27.39 MiB), `4280252` compressed data bytes (4.08 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00002.parquet`: `32623` rows, `4336053` file bytes (4.14 MiB), `25501630` physical bytes (24.32 MiB), `28987375` encoded bytes (27.64 MiB), `4303122` compressed data bytes (4.10 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00003.parquet`: `32804` rows, `4333369` file bytes (4.13 MiB), `25464243` physical bytes (24.28 MiB), `28965141` encoded bytes (27.62 MiB), `4300838` compressed data bytes (4.10 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00004.parquet`: `33072` rows, `4363830` file bytes (4.16 MiB), `25661047` physical bytes (24.47 MiB), `29189147` encoded bytes (27.84 MiB), `4330634` compressed data bytes (4.13 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00005.parquet`: `33277` rows, `4305438` file bytes (4.11 MiB), `25689273` physical bytes (24.50 MiB), `29240519` encoded bytes (27.89 MiB), `4272833` compressed data bytes (4.07 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00006.parquet`: `33065` rows, `4330299` file bytes (4.13 MiB), `25436124` physical bytes (24.26 MiB), `28961086` encoded bytes (27.62 MiB), `4297969` compressed data bytes (4.10 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00007.parquet`: `32894` rows, `4313193` file bytes (4.11 MiB), `25488737` physical bytes (24.31 MiB), `29003775` encoded bytes (27.66 MiB), `4280546` compressed data bytes (4.08 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00008.parquet`: `32906` rows, `4312742` file bytes (4.11 MiB), `25488255` physical bytes (24.31 MiB), `28995076` encoded bytes (27.65 MiB), `4279711` compressed data bytes (4.08 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00009.parquet`: `32989` rows, `4290478` file bytes (4.09 MiB), `25511599` physical bytes (24.33 MiB), `29028419` encoded bytes (27.68 MiB), `4258056` compressed data bytes (4.06 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00010.parquet`: `33021` rows, `4316864` file bytes (4.12 MiB), `25461016` physical bytes (24.28 MiB), `28977788` encoded bytes (27.64 MiB), `4284674` compressed data bytes (4.09 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00011.parquet`: `33133` rows, `4284902` file bytes (4.09 MiB), `25472597` physical bytes (24.29 MiB), `29004519` encoded bytes (27.66 MiB), `4252381` compressed data bytes (4.06 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00012.parquet`: `33100` rows, `4279575` file bytes (4.08 MiB), `25513804` physical bytes (24.33 MiB), `29045964` encoded bytes (27.70 MiB), `4247174` compressed data bytes (4.05 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00013.parquet`: `33102` rows, `4316114` file bytes (4.12 MiB), `25483836` physical bytes (24.30 MiB), `29012796` encoded bytes (27.67 MiB), `4284016` compressed data bytes (4.09 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00014.parquet`: `32763` rows, `4242063` file bytes (4.05 MiB), `25143062` physical bytes (23.98 MiB), `28629648` encoded bytes (27.30 MiB), `4209604` compressed data bytes (4.01 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00015.parquet`: `32943` rows, `4583624` file bytes (4.37 MiB), `23860082` physical bytes (22.75 MiB), `27511334` encoded bytes (26.24 MiB), `4552312` compressed data bytes (4.34 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00016.parquet`: `32882` rows, `4735583` file bytes (4.52 MiB), `22642074` physical bytes (21.59 MiB), `26368926` encoded bytes (25.15 MiB), `4705053` compressed data bytes (4.49 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00017.parquet`: `31427` rows, `4708810` file bytes (4.49 MiB), `22004361` physical bytes (20.98 MiB), `25606088` encoded bytes (24.42 MiB), `4678472` compressed data bytes (4.46 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00018.parquet`: `32110` rows, `4626166` file bytes (4.41 MiB), `20272524` physical bytes (19.33 MiB), `23992875` encoded bytes (22.88 MiB), `4595879` compressed data bytes (4.38 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00019.parquet`: `31775` rows, `4696704` file bytes (4.48 MiB), `20557399` physical bytes (19.61 MiB), `24240752` encoded bytes (23.12 MiB), `4665948` compressed data bytes (4.45 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00020.parquet`: `31814` rows, `4880752` file bytes (4.65 MiB), `20771886` physical bytes (19.81 MiB), `24452635` encoded bytes (23.32 MiB), `4850268` compressed data bytes (4.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00021.parquet`: `32113` rows, `4749081` file bytes (4.53 MiB), `20564977` physical bytes (19.61 MiB), `24283123` encoded bytes (23.16 MiB), `4718204` compressed data bytes (4.50 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00022.parquet`: `32196` rows, `4809788` file bytes (4.59 MiB), `20676224` physical bytes (19.72 MiB), `24408935` encoded bytes (23.28 MiB), `4779044` compressed data bytes (4.56 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00023.parquet`: `32619` rows, `4785826` file bytes (4.56 MiB), `20870675` physical bytes (19.90 MiB), `24651212` encoded bytes (23.51 MiB), `4755077` compressed data bytes (4.53 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00024.parquet`: `32624` rows, `4786421` file bytes (4.56 MiB), `20849345` physical bytes (19.88 MiB), `24625865` encoded bytes (23.49 MiB), `4755904` compressed data bytes (4.54 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00025.parquet`: `32034` rows, `4845039` file bytes (4.62 MiB), `20880628` physical bytes (19.91 MiB), `24592752` encoded bytes (23.45 MiB), `4814240` compressed data bytes (4.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00026.parquet`: `32183` rows, `4740572` file bytes (4.52 MiB), `20543916` physical bytes (19.59 MiB), `24268813` encoded bytes (23.14 MiB), `4710310` compressed data bytes (4.49 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00027.parquet`: `32341` rows, `4743147` file bytes (4.52 MiB), `20598948` physical bytes (19.64 MiB), `24348082` encoded bytes (23.22 MiB), `4712487` compressed data bytes (4.49 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00028.parquet`: `32525` rows, `4713439` file bytes (4.50 MiB), `20756858` physical bytes (19.80 MiB), `24523639` encoded bytes (23.39 MiB), `4682771` compressed data bytes (4.47 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00029.parquet`: `32450` rows, `4726545` file bytes (4.51 MiB), `20584830` physical bytes (19.63 MiB), `24342456` encoded bytes (23.21 MiB), `4696215` compressed data bytes (4.48 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-snappy-int-plain-str-plain-date-plain-ts-rle-dict/part-00030.parquet`: `21674` rows, `3208097` file bytes (3.06 MiB), `13980072` physical bytes (13.33 MiB), `16495632` encoded bytes (15.73 MiB), `3178178` compressed data bytes (3.03 MiB)
