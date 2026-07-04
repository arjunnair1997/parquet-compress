# ClickBench Parquet Experiment

- Started: `2026-07-03T15:32:20-04:00`
- Write elapsed: `11.944s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `819345407` (781.39 MiB)
- Compressed column data bytes after codec compression: `91377768` (87.14 MiB)
- Parquet file bytes: `92280698` (88.01 MiB)
- Physical/encoded ratio: `0.869x`
- Encoded/compressed-data ratio: `8.967x`
- Physical/compressed-data ratio: `7.796x`
- Physical/parquet-file ratio: `7.720x`
- Files: `29`

## Settings

- Compression: `zstd-3`
- Int encoding: `plain`
- String encoding: `plain`
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
- Files read: `29`
- Elapsed: `7.511s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004557` (7.63 MiB) | `8005361` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `64020` (62.52 KiB) | `0.999x` | `62.537x` | `62.480x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:528` | `1000000` | `138409995` (132.00 MiB) | `142867610` (136.25 MiB) | `13953285` (13.31 MiB) | `0.969x` | `10.239x` | `9.920x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4920` (4.80 KiB) | `0.999x` | `813.725x` | `813.008x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `7357539` (7.02 MiB) | `4039124` (3.85 MiB) | `1.087x` | `1.822x` | `1.981x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4953` (4.84 KiB) | `5979` (5.84 KiB) | `807.591x` | `0.828x` | `669.008x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003533` (3.82 MiB) | `4942` (4.83 KiB) | `0.999x` | `810.104x` | `809.389x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003601` (3.82 MiB) | `408178` (398.61 KiB) | `0.999x` | `9.808x` | `9.800x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `190813` (186.34 KiB) | `0.999x` | `20.982x` | `20.963x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004551` (7.63 MiB) | `617972` (603.49 KiB) | `0.999x` | `12.953x` | `12.946x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4236` (4.14 KiB) | `0.999x` | `945.120x` | `944.287x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `106753` (104.25 KiB) | `0.999x` | `37.503x` | `37.470x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `133847` (130.71 KiB) | `0.999x` | `29.912x` | `29.885x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:356` | `1000000` | `88562192` (84.46 MiB) | `92651103` (88.36 MiB) | `15307125` (14.60 MiB) | `0.956x` | `6.053x` | `5.786x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:328` | `1000000` | `79583339` (75.90 MiB) | `83645458` (79.77 MiB) | `14221544` (13.56 MiB) | `0.951x` | `5.882x` | `5.596x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `178877` (174.68 KiB) | `0.999x` | `22.382x` | `22.362x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `275126` (268.68 KiB) | `0.999x` | `14.552x` | `14.539x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `231319` (225.90 KiB) | `0.999x` | `17.308x` | `17.292x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `87533` (85.48 KiB) | `0.999x` | `45.738x` | `45.697x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `48077` (46.95 KiB) | `0.999x` | `83.275x` | `83.200x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `187250` (182.86 KiB) | `0.999x` | `21.381x` | `21.362x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `186190` (181.83 KiB) | `0.999x` | `21.503x` | `21.483x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `82256` (80.33 KiB) | `0.999x` | `48.672x` | `48.629x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003591` (3.82 MiB) | `53287` (52.04 KiB) | `0.999x` | `75.133x` | `75.065x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `128915` (125.89 KiB) | `0.999x` | `31.056x` | `31.028x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3354477` (3.20 MiB) | `7357830` (7.02 MiB) | `246331` (240.56 KiB) | `0.456x` | `29.870x` | `13.618x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `26694` (26.07 KiB) | `0.999x` | `149.981x` | `149.846x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `25191` (24.60 KiB) | `0.999x` | `158.929x` | `158.787x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `155224` (151.59 KiB) | `0.999x` | `25.792x` | `25.769x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3767530` (3.59 MiB) | `7770889` (7.41 MiB) | `136972` (133.76 KiB) | `0.485x` | `56.733x` | `27.506x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003539` (3.82 MiB) | `6053` (5.91 KiB) | `0.999x` | `661.414x` | `660.829x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003544` (3.82 MiB) | `6561` (6.41 KiB) | `0.999x` | `610.203x` | `609.663x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `28902` (28.22 KiB) | `0.999x` | `138.523x` | `138.399x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `22977` (22.44 KiB) | `0.999x` | `174.243x` | `174.087x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `81583` (79.67 KiB) | `4084892` (3.90 MiB) | `22657` (22.13 KiB) | `0.020x` | `180.293x` | `3.601x` | `81583` (79.67 KiB) |
| `Params` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002161` (3.82 MiB) | `2868` (2.80 KiB) | `0.000x` | `1395.454x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003592` (3.82 MiB) | `323741` (316.15 KiB) | `0.999x` | `12.367x` | `12.356x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `288165` (281.41 KiB) | `0.999x` | `13.893x` | `13.881x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `101739` (99.35 KiB) | `0.999x` | `39.352x` | `39.316x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3528017` (3.36 MiB) | `7534999` (7.19 MiB) | `721542` (704.63 KiB) | `0.468x` | `10.443x` | `4.890x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `30732` (30.01 KiB) | `0.999x` | `130.274x` | `130.157x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `164931` (161.07 KiB) | `0.999x` | `24.274x` | `24.253x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `305881` (298.71 KiB) | `0.999x` | `13.089x` | `13.077x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `320010` (312.51 KiB) | `0.999x` | `12.511x` | `12.500x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `99843` (97.50 KiB) | `0.999x` | `40.099x` | `40.063x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `7208507` (6.87 MiB) | `3976226` (3.79 MiB) | `1.110x` | `1.813x` | `2.012x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `90970` (88.84 KiB) | `0.999x` | `44.010x` | `43.971x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `74605` (72.86 KiB) | `0.999x` | `53.664x` | `53.616x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `123955` (121.05 KiB) | `0.999x` | `32.299x` | `32.270x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003527` (3.82 MiB) | `4360` (4.26 KiB) | `0.999x` | `918.240x` | `917.431x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:85` | `1000000` | `13587860` (12.96 MiB) | `17594877` (16.78 MiB) | `14749` (14.40 KiB) | `0.772x` | `1192.954x` | `921.273x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003553` (3.82 MiB) | `7128` (6.96 KiB) | `0.999x` | `561.666x` | `561.167x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003557` (3.82 MiB) | `56734` (55.40 KiB) | `0.999x` | `70.567x` | `70.504x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003554` (3.82 MiB) | `8069` (7.88 KiB) | `0.999x` | `496.165x` | `495.724x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003539` (3.82 MiB) | `25202` (24.61 KiB) | `0.999x` | `158.858x` | `158.718x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004558` (7.63 MiB) | `693965` (677.70 KiB) | `0.999x` | `11.535x` | `11.528x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:137` | `1000000` | `27797671` (26.51 MiB) | `31856668` (30.38 MiB) | `5316593` (5.07 MiB) | `0.873x` | `5.992x` | `5.228x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003644` (3.82 MiB) | `3853120` (3.67 MiB) | `0.999x` | `1.039x` | `1.038x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4236` (4.14 KiB) | `0.999x` | `945.120x` | `944.287x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4236` (4.14 KiB) | `0.999x` | `945.120x` | `944.287x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4236` (4.14 KiB) | `0.999x` | `945.120x` | `944.287x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003589` (3.82 MiB) | `82547` (80.61 KiB) | `0.999x` | `48.501x` | `48.457x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4236` (4.14 KiB) | `0.999x` | `945.120x` | `944.287x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `1000000` (976.56 KiB) | `5002904` (4.77 MiB) | `29467` (28.78 KiB) | `0.200x` | `169.780x` | `33.936x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `7354048` (7.01 MiB) | `4039947` (3.85 MiB) | `1.088x` | `1.820x` | `1.980x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `142628` (139.29 KiB) | `0.999x` | `28.070x` | `28.045x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `107835` (105.31 KiB) | `0.999x` | `37.127x` | `37.094x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `122853` (119.97 KiB) | `0.999x` | `32.588x` | `32.559x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `193449` (188.92 KiB) | `0.999x` | `20.696x` | `20.677x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `173958` (169.88 KiB) | `0.999x` | `23.015x` | `22.994x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003601` (3.82 MiB) | `426500` (416.50 KiB) | `0.999x` | `9.387x` | `9.379x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003582` (3.82 MiB) | `70816` (69.16 KiB) | `0.999x` | `56.535x` | `56.484x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4236` (4.14 KiB) | `0.999x` | `945.120x` | `944.287x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003574` (3.82 MiB) | `55294` (54.00 KiB) | `0.999x` | `72.405x` | `72.341x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `2001192` (1.91 MiB) | `6004437` (5.73 MiB) | `32234` (31.48 KiB) | `0.333x` | `186.277x` | `62.083x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3325142` (3.17 MiB) | `7328504` (6.99 MiB) | `122185` (119.32 KiB) | `0.454x` | `59.979x` | `27.214x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002161` (3.82 MiB) | `2868` (2.80 KiB) | `0.000x` | `1395.454x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002161` (3.82 MiB) | `2868` (2.80 KiB) | `0.000x` | `1395.454x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4236` (4.14 KiB) | `0.999x` | `945.120x` | `944.287x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003546` (3.82 MiB) | `61456` (60.02 KiB) | `0.999x` | `65.145x` | `65.087x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `135525` (132.35 KiB) | `0.999x` | `29.541x` | `29.515x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003599` (3.82 MiB) | `333164` (325.36 KiB) | `0.999x` | `12.017x` | `12.006x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003646` (3.82 MiB) | `1245827` (1.19 MiB) | `0.999x` | `3.214x` | `3.211x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003647` (3.82 MiB) | `938027` (916.04 KiB) | `0.999x` | `4.268x` | `4.264x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003624` (3.82 MiB) | `550130` (537.24 KiB) | `0.999x` | `7.278x` | `7.271x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003534` (3.82 MiB) | `5334` (5.21 KiB) | `0.999x` | `750.569x` | `749.906x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `1024` (1.00 KiB) | `4004077` (3.82 MiB) | `4956` (4.84 KiB) | `0.000x` | `807.925x` | `0.207x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004443` (7.63 MiB) | `5704` (5.57 KiB) | `0.999x` | `1403.303x` | `1402.525x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002161` (3.82 MiB) | `2868` (2.80 KiB) | `0.000x` | `1395.454x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3000000` (2.86 MiB) | `7003298` (6.68 MiB) | `5258` (5.13 KiB) | `0.428x` | `1331.932x` | `570.559x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4236` (4.14 KiB) | `0.999x` | `945.120x` | `944.287x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `58030` (56.67 KiB) | `4062531` (3.87 MiB) | `19283` (18.83 KiB) | `0.014x` | `210.679x` | `3.009x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `22051` (21.53 KiB) | `4025591` (3.84 MiB) | `16486` (16.10 KiB) | `0.005x` | `244.182x` | `1.338x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `25445` (24.85 KiB) | `4029067` (3.84 MiB) | `18960` (18.52 KiB) | `0.006x` | `212.504x` | `1.342x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `48191` (47.06 KiB) | `4051763` (3.86 MiB) | `13345` (13.03 KiB) | `0.012x` | `303.617x` | `3.611x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `49433` (48.27 KiB) | `4053237` (3.87 MiB) | `21138` (20.64 KiB) | `0.012x` | `191.751x` | `2.339x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `16873` (16.48 KiB) | `4019952` (3.83 MiB) | `16987` (16.59 KiB) | `0.004x` | `236.649x` | `0.993x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `91870` (89.72 KiB) | `4097525` (3.91 MiB) | `29881` (29.18 KiB) | `0.022x` | `137.128x` | `3.075x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `13001` (12.70 KiB) | `4016724` (3.83 MiB) | `14181` (13.85 KiB) | `0.003x` | `283.247x` | `0.917x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `28101` (27.44 KiB) | `4032717` (3.85 MiB) | `15663` (15.30 KiB) | `0.007x` | `257.468x` | `1.794x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `45607` (44.54 KiB) | `4048297` (3.86 MiB) | `28416` (27.75 KiB) | `0.011x` | `142.465x` | `1.605x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003560` (3.82 MiB) | `21166` (20.67 KiB) | `0.999x` | `189.151x` | `188.982x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004555` (7.63 MiB) | `2842781` (2.71 MiB) | `0.999x` | `2.816x` | `2.814x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004556` (7.63 MiB) | `3580857` (3.41 MiB) | `0.999x` | `2.235x` | `2.234x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003544` (3.82 MiB) | `5655` (5.52 KiB) | `0.999x` | `707.965x` | `707.339x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00000.parquet`: `35785` rows, `3001838` file bytes (2.86 MiB), `27793026` physical bytes (26.51 MiB), `31463247` encoded bytes (30.01 MiB), `2969684` compressed data bytes (2.83 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00001.parquet`: `35363` rows, `2910014` file bytes (2.78 MiB), `27158090` physical bytes (25.90 MiB), `30793268` encoded bytes (29.37 MiB), `2877427` compressed data bytes (2.74 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00002.parquet`: `35727` rows, `3018667` file bytes (2.88 MiB), `27984129` physical bytes (26.69 MiB), `31655835` encoded bytes (30.19 MiB), `2985266` compressed data bytes (2.85 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00003.parquet`: `36136` rows, `2982722` file bytes (2.84 MiB), `27986082` physical bytes (26.69 MiB), `31696945` encoded bytes (30.23 MiB), `2950196` compressed data bytes (2.81 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00004.parquet`: `35520` rows, `2927053` file bytes (2.79 MiB), `27629210` physical bytes (26.35 MiB), `31278651` encoded bytes (29.83 MiB), `2893771` compressed data bytes (2.76 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00005.parquet`: `36197` rows, `2964394` file bytes (2.83 MiB), `27774550` physical bytes (26.49 MiB), `31489367` encoded bytes (30.03 MiB), `2931770` compressed data bytes (2.80 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00006.parquet`: `36189` rows, `2977452` file bytes (2.84 MiB), `27972902` physical bytes (26.68 MiB), `31693375` encoded bytes (30.23 MiB), `2944951` compressed data bytes (2.81 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00007.parquet`: `36145` rows, `2989725` file bytes (2.85 MiB), `27970887` physical bytes (26.68 MiB), `31685345` encoded bytes (30.22 MiB), `2956753` compressed data bytes (2.82 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00008.parquet`: `36153` rows, `2936577` file bytes (2.80 MiB), `28024743` physical bytes (26.73 MiB), `31732600` encoded bytes (30.26 MiB), `2903758` compressed data bytes (2.77 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00009.parquet`: `35750` rows, `2952388` file bytes (2.82 MiB), `27598233` physical bytes (26.32 MiB), `31264895` encoded bytes (29.82 MiB), `2920107` compressed data bytes (2.78 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00010.parquet`: `36385` rows, `2941724` file bytes (2.81 MiB), `27968296` physical bytes (26.67 MiB), `31700727` encoded bytes (30.23 MiB), `2909127` compressed data bytes (2.77 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00011.parquet`: `36317` rows, `2939869` file bytes (2.80 MiB), `28011955` physical bytes (26.71 MiB), `31744239` encoded bytes (30.27 MiB), `2907254` compressed data bytes (2.77 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00012.parquet`: `36210` rows, `2954787` file bytes (2.82 MiB), `27788987` physical bytes (26.50 MiB), `31501669` encoded bytes (30.04 MiB), `2922544` compressed data bytes (2.79 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00013.parquet`: `35765` rows, `2920998` file bytes (2.79 MiB), `27619973` physical bytes (26.34 MiB), `31285168` encoded bytes (29.84 MiB), `2888319` compressed data bytes (2.75 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00014.parquet`: `35418` rows, `3383722` file bytes (3.23 MiB), `24647464` physical bytes (23.51 MiB), `28490663` encoded bytes (27.17 MiB), `3352691` compressed data bytes (3.20 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00015.parquet`: `34980` rows, `3427976` file bytes (3.27 MiB), `24371725` physical bytes (23.24 MiB), `28202290` encoded bytes (26.90 MiB), `3397662` compressed data bytes (3.24 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00016.parquet`: `34515` rows, `3506828` file bytes (3.34 MiB), `22970681` physical bytes (21.91 MiB), `26806012` encoded bytes (25.56 MiB), `3476483` compressed data bytes (3.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00017.parquet`: `34452` rows, `3549087` file bytes (3.38 MiB), `22152833` physical bytes (21.13 MiB), `26000693` encoded bytes (24.80 MiB), `3518235` compressed data bytes (3.36 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00018.parquet`: `33904` rows, `3691880` file bytes (3.52 MiB), `22275944` physical bytes (21.24 MiB), `26065929` encoded bytes (24.86 MiB), `3661395` compressed data bytes (3.49 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00019.parquet`: `34555` rows, `3464426` file bytes (3.30 MiB), `21774627` physical bytes (20.77 MiB), `25633264` encoded bytes (24.45 MiB), `3433595` compressed data bytes (3.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00020.parquet`: `33840` rows, `3601424` file bytes (3.43 MiB), `21889410` physical bytes (20.88 MiB), `25672763` encoded bytes (24.48 MiB), `3570684` compressed data bytes (3.41 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00021.parquet`: `34486` rows, `3561294` file bytes (3.40 MiB), `22093869` physical bytes (21.07 MiB), `25948222` encoded bytes (24.75 MiB), `3530321` compressed data bytes (3.37 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00022.parquet`: `34071` rows, `3494011` file bytes (3.33 MiB), `21782259` physical bytes (20.77 MiB), `25585588` encoded bytes (24.40 MiB), `3463567` compressed data bytes (3.30 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00023.parquet`: `34114` rows, `3597147` file bytes (3.43 MiB), `22115853` physical bytes (21.09 MiB), `25927189` encoded bytes (24.73 MiB), `3566323` compressed data bytes (3.40 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00024.parquet`: `34286` rows, `3530410` file bytes (3.37 MiB), `21894060` physical bytes (20.88 MiB), `25722859` encoded bytes (24.53 MiB), `3500039` compressed data bytes (3.34 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00025.parquet`: `34339` rows, `3548386` file bytes (3.38 MiB), `21922281` physical bytes (20.91 MiB), `25761098` encoded bytes (24.57 MiB), `3517766` compressed data bytes (3.35 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00026.parquet`: `34555` rows, `3486890` file bytes (3.33 MiB), `21955492` physical bytes (20.94 MiB), `25813447` encoded bytes (24.62 MiB), `3456191` compressed data bytes (3.30 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00027.parquet`: `34305` rows, `3501858` file bytes (3.34 MiB), `21800467` physical bytes (20.79 MiB), `25633061` encoded bytes (24.45 MiB), `3471537` compressed data bytes (3.31 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-rle-dict/part-00028.parquet`: `14538` rows, `1517151` file bytes (1.45 MiB), `9470596` physical bytes (9.03 MiB), `11096998` encoded bytes (10.58 MiB), `1500352` compressed data bytes (1.43 MiB)
