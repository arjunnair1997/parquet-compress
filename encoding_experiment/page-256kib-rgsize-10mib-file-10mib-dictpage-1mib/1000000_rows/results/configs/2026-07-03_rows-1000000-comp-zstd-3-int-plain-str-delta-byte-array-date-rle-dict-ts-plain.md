# ClickBench Parquet Experiment

- Started: `2026-07-03T15:32:05-04:00`
- Write elapsed: `12.226s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `514915477` (491.06 MiB)
- Compressed column data bytes after codec compression: `87752370` (83.69 MiB)
- Parquet file bytes: `88651901` (84.55 MiB)
- Physical/encoded ratio: `1.384x`
- Encoded/compressed-data ratio: `5.868x`
- Physical/compressed-data ratio: `8.118x`
- Physical/parquet-file ratio: `8.036x`
- Files: `29`

## Settings

- Compression: `zstd-3`
- Int encoding: `plain`
- String encoding: `delta-byte-array`
- Date encoding: `rle-dict`
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
- Rows read and compared: `1000000`
- Files read: `29`
- Elapsed: `7.496s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004555` (7.63 MiB) | `8005363` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `63363` (61.88 KiB) | `0.999x` | `63.185x` | `63.128x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:528` | `1000000` | `138409995` (132.00 MiB) | `64476943` (61.49 MiB) | `13967585` (13.32 MiB) | `2.147x` | `4.616x` | `9.909x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4923` (4.81 KiB) | `0.999x` | `813.229x` | `812.513x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004553` (7.63 MiB) | `2518043` (2.40 MiB) | `0.999x` | `3.179x` | `3.177x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4950` (4.83 KiB) | `5976` (5.84 KiB) | `808.081x` | `0.828x` | `669.344x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4941` (4.83 KiB) | `0.999x` | `810.267x` | `809.553x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003596` (3.82 MiB) | `408067` (398.50 KiB) | `0.999x` | `9.811x` | `9.802x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `191221` (186.74 KiB) | `0.999x` | `20.937x` | `20.918x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004548` (7.63 MiB) | `618261` (603.77 KiB) | `0.999x` | `12.947x` | `12.940x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `4242` (4.14 KiB) | `0.999x` | `943.784x` | `942.951x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `106139` (103.65 KiB) | `0.999x` | `37.720x` | `37.686x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `135008` (131.84 KiB) | `0.999x` | `29.654x` | `29.628x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:354` | `1000000` | `88562192` (84.46 MiB) | `40463506` (38.59 MiB) | `15098901` (14.40 MiB) | `2.189x` | `2.680x` | `5.865x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:326` | `1000000` | `79583339` (75.90 MiB) | `38993213` (37.19 MiB) | `14563597` (13.89 MiB) | `2.041x` | `2.677x` | `5.465x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `178766` (174.58 KiB) | `0.999x` | `22.396x` | `22.376x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `275202` (268.75 KiB) | `0.999x` | `14.548x` | `14.535x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `231331` (225.91 KiB) | `0.999x` | `17.307x` | `17.291x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003582` (3.82 MiB) | `87579` (85.53 KiB) | `0.999x` | `45.714x` | `45.673x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `48142` (47.01 KiB) | `0.999x` | `83.162x` | `83.088x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `187165` (182.78 KiB) | `0.999x` | `21.391x` | `21.372x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `186041` (181.68 KiB) | `0.999x` | `21.520x` | `21.501x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `81846` (79.93 KiB) | `0.999x` | `48.916x` | `48.872x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `53719` (52.46 KiB) | `0.999x` | `74.528x` | `74.462x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `128114` (125.11 KiB) | `0.999x` | `31.250x` | `31.222x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `3354477` (3.20 MiB) | `1042299` (1017.87 KiB) | `325949` (318.31 KiB) | `3.218x` | `3.198x` | `10.291x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `26187` (25.57 KiB) | `0.999x` | `152.885x` | `152.748x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `24936` (24.35 KiB) | `0.999x` | `160.554x` | `160.411x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `156119` (152.46 KiB) | `0.999x` | `25.644x` | `25.621x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `3767530` (3.59 MiB) | `863680` (843.44 KiB) | `179568` (175.36 KiB) | `4.362x` | `4.810x` | `20.981x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003532` (3.82 MiB) | `6089` (5.95 KiB) | `0.999x` | `657.502x` | `656.922x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003544` (3.82 MiB) | `6485` (6.33 KiB) | `0.999x` | `617.355x` | `616.808x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `28629` (27.96 KiB) | `0.999x` | `139.844x` | `139.718x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `22551` (22.02 KiB) | `0.999x` | `177.535x` | `177.376x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `81583` (79.67 KiB) | `283596` (276.95 KiB) | `40719` (39.76 KiB) | `0.288x` | `6.965x` | `2.004x` | `81583` (79.67 KiB) |
| `Params` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `81289` (79.38 KiB) | `3682` (3.60 KiB) | `0.000x` | `22.077x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003591` (3.82 MiB) | `323812` (316.22 KiB) | `0.999x` | `12.364x` | `12.353x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `288206` (281.45 KiB) | `0.999x` | `13.891x` | `13.879x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `101599` (99.22 KiB) | `0.999x` | `39.406x` | `39.370x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `3528017` (3.36 MiB) | `2959595` (2.82 MiB) | `825816` (806.46 KiB) | `1.192x` | `3.584x` | `4.272x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003578` (3.82 MiB) | `30678` (29.96 KiB) | `0.999x` | `130.503x` | `130.387x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `164973` (161.11 KiB) | `0.999x` | `24.268x` | `24.246x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003590` (3.82 MiB) | `305758` (298.59 KiB) | `0.999x` | `13.094x` | `13.082x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003591` (3.82 MiB) | `319803` (312.31 KiB) | `0.999x` | `12.519x` | `12.508x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003582` (3.82 MiB) | `99975` (97.63 KiB) | `0.999x` | `40.046x` | `40.010x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004552` (7.63 MiB) | `2476454` (2.36 MiB) | `0.999x` | `3.232x` | `3.230x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `90129` (88.02 KiB) | `0.999x` | `44.421x` | `44.381x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `73932` (72.20 KiB) | `0.999x` | `54.152x` | `54.104x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `123759` (120.86 KiB) | `0.999x` | `32.350x` | `32.321x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003530` (3.82 MiB) | `4371` (4.27 KiB) | `0.999x` | `915.930x` | `915.122x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:85` | `1000000` | `13587860` (12.96 MiB) | `140943` (137.64 KiB) | `21461` (20.96 KiB) | `96.407x` | `6.567x` | `633.142x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003547` (3.82 MiB) | `7031` (6.87 KiB) | `0.999x` | `569.414x` | `568.909x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003558` (3.82 MiB) | `56675` (55.35 KiB) | `0.999x` | `70.641x` | `70.578x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003553` (3.82 MiB) | `8030` (7.84 KiB) | `0.999x` | `498.574x` | `498.132x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003535` (3.82 MiB) | `25177` (24.59 KiB) | `0.999x` | `159.016x` | `158.875x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004553` (7.63 MiB) | `694037` (677.77 KiB) | `0.999x` | `11.533x` | `11.527x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:137` | `1000000` | `27797671` (26.51 MiB) | `21050725` (20.08 MiB) | `5595515` (5.34 MiB) | `1.321x` | `3.762x` | `4.968x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003641` (3.82 MiB) | `3831650` (3.65 MiB) | `0.999x` | `1.045x` | `1.044x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `4242` (4.14 KiB) | `0.999x` | `943.784x` | `942.951x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `4242` (4.14 KiB) | `0.999x` | `943.784x` | `942.951x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `4242` (4.14 KiB) | `0.999x` | `943.784x` | `942.951x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `82509` (80.58 KiB) | `0.999x` | `48.523x` | `48.480x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `4242` (4.14 KiB) | `0.999x` | `943.784x` | `942.951x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `1000000` (976.56 KiB) | `206935` (202.08 KiB) | `51083` (49.89 KiB) | `4.832x` | `4.051x` | `19.576x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004555` (7.63 MiB) | `2520039` (2.40 MiB) | `0.999x` | `3.176x` | `3.175x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `142797` (139.45 KiB) | `0.999x` | `28.037x` | `28.012x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `108054` (105.52 KiB) | `0.999x` | `37.052x` | `37.019x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `122704` (119.83 KiB) | `0.999x` | `32.628x` | `32.599x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `193522` (188.99 KiB) | `0.999x` | `20.688x` | `20.669x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `173662` (169.59 KiB) | `0.999x` | `23.054x` | `23.033x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003601` (3.82 MiB) | `426543` (416.55 KiB) | `0.999x` | `9.386x` | `9.378x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003580` (3.82 MiB) | `70834` (69.17 KiB) | `0.999x` | `56.521x` | `56.470x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003526` (3.82 MiB) | `4237` (4.14 KiB) | `0.999x` | `944.896x` | `944.064x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003572` (3.82 MiB) | `55094` (53.80 KiB) | `0.999x` | `72.668x` | `72.603x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `2001192` (1.91 MiB) | `332270` (324.48 KiB) | `65243` (63.71 KiB) | `6.023x` | `5.093x` | `30.673x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `3325142` (3.17 MiB) | `963161` (940.59 KiB) | `195429` (190.85 KiB) | `3.452x` | `4.928x` | `17.015x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `81289` (79.38 KiB) | `3682` (3.60 KiB) | `0.000x` | `22.077x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `81289` (79.38 KiB) | `3682` (3.60 KiB) | `0.000x` | `22.077x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `4242` (4.14 KiB) | `0.999x` | `943.784x` | `942.951x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003541` (3.82 MiB) | `61466` (60.03 KiB) | `0.999x` | `65.134x` | `65.077x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003581` (3.82 MiB) | `135420` (132.25 KiB) | `0.999x` | `29.564x` | `29.538x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003598` (3.82 MiB) | `333375` (325.56 KiB) | `0.999x` | `12.009x` | `11.999x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003642` (3.82 MiB) | `1245932` (1.19 MiB) | `0.999x` | `3.213x` | `3.210x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003639` (3.82 MiB) | `937990` (916.01 KiB) | `0.999x` | `4.268x` | `4.264x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003623` (3.82 MiB) | `549886` (537.00 KiB) | `0.999x` | `7.281x` | `7.274x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `5330` (5.21 KiB) | `0.999x` | `751.132x` | `750.469x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `1024` (1.00 KiB) | `86534` (84.51 KiB) | `6144` (6.00 KiB) | `0.012x` | `14.084x` | `0.167x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004440` (7.63 MiB) | `5725` (5.59 KiB) | `0.999x` | `1398.155x` | `1397.380x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `0` (0 B) | `81289` (79.38 KiB) | `3682` (3.60 KiB) | `0.000x` | `22.077x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `3000000` (2.86 MiB) | `86024` (84.01 KiB) | `6930` (6.77 KiB) | `34.874x` | `12.413x` | `432.900x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003531` (3.82 MiB) | `4242` (4.14 KiB) | `0.999x` | `943.784x` | `942.951x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `58030` (56.67 KiB) | `303452` (296.34 KiB) | `33966` (33.17 KiB) | `0.191x` | `8.934x` | `1.708x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `22051` (21.53 KiB) | `198992` (194.33 KiB) | `28243` (27.58 KiB) | `0.111x` | `7.046x` | `0.781x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `25445` (24.85 KiB) | `210973` (206.03 KiB) | `33655` (32.87 KiB) | `0.121x` | `6.269x` | `0.756x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `48191` (47.06 KiB) | `234052` (228.57 KiB) | `23444` (22.89 KiB) | `0.206x` | `9.983x` | `2.056x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `49433` (48.27 KiB) | `267362` (261.10 KiB) | `40012` (39.07 KiB) | `0.185x` | `6.682x` | `1.235x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `16873` (16.48 KiB) | `213162` (208.17 KiB) | `27888` (27.23 KiB) | `0.079x` | `7.644x` | `0.605x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `91870` (89.72 KiB) | `327069` (319.40 KiB) | `54122` (52.85 KiB) | `0.281x` | `6.043x` | `1.697x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `13001` (12.70 KiB) | `149609` (146.10 KiB) | `20665` (20.18 KiB) | `0.087x` | `7.240x` | `0.629x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `28101` (27.44 KiB) | `189991` (185.54 KiB) | `23115` (22.57 KiB) | `0.148x` | `8.219x` | `1.216x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `delta-byte-array` | `DELTA_BYTE_ARRAY` | `DATA_PAGE_V2/DELTA_BYTE_ARRAY:57` | `1000000` | `45607` (44.54 KiB) | `261004` (254.89 KiB) | `38352` (37.45 KiB) | `0.175x` | `6.805x` | `1.189x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003559` (3.82 MiB) | `21190` (20.69 KiB) | `0.999x` | `188.936x` | `188.768x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004558` (7.63 MiB) | `2841956` (2.71 MiB) | `0.999x` | `2.817x` | `2.815x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004552` (7.63 MiB) | `3580374` (3.41 MiB) | `0.999x` | `2.236x` | `2.234x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003547` (3.82 MiB) | `5652` (5.52 KiB) | `0.999x` | `708.342x` | `707.714x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00000.parquet`: `35889` rows, `2857207` file bytes (2.72 MiB), `27875599` physical bytes (26.58 MiB), `17887730` encoded bytes (17.06 MiB), `2825147` compressed data bytes (2.69 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00001.parquet`: `35984` rows, `2806613` file bytes (2.68 MiB), `27645074` physical bytes (26.36 MiB), `17696421` encoded bytes (16.88 MiB), `2774011` compressed data bytes (2.65 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00002.parquet`: `35879` rows, `2848295` file bytes (2.72 MiB), `28094680` physical bytes (26.79 MiB), `17937953` encoded bytes (17.11 MiB), `2815013` compressed data bytes (2.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00003.parquet`: `36267` rows, `2824537` file bytes (2.69 MiB), `28070313` physical bytes (26.77 MiB), `17880336` encoded bytes (17.05 MiB), `2792117` compressed data bytes (2.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00004.parquet`: `35612` rows, `2783867` file bytes (2.65 MiB), `27705357` physical bytes (26.42 MiB), `17615674` encoded bytes (16.80 MiB), `2750668` compressed data bytes (2.62 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00005.parquet`: `36078` rows, `2800483` file bytes (2.67 MiB), `27657416` physical bytes (26.38 MiB), `17772909` encoded bytes (16.95 MiB), `2768006` compressed data bytes (2.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00006.parquet`: `36220` rows, `2847459` file bytes (2.72 MiB), `28044091` physical bytes (26.74 MiB), `17831923` encoded bytes (17.01 MiB), `2815006` compressed data bytes (2.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00007.parquet`: `36251` rows, `2850061` file bytes (2.72 MiB), `28054804` physical bytes (26.76 MiB), `17963967` encoded bytes (17.13 MiB), `2817226` compressed data bytes (2.69 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00008.parquet`: `36287` rows, `2790405` file bytes (2.66 MiB), `28103637` physical bytes (26.80 MiB), `17886472` encoded bytes (17.06 MiB), `2757660` compressed data bytes (2.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00009.parquet`: `35856` rows, `2820293` file bytes (2.69 MiB), `27656610` physical bytes (26.38 MiB), `17718657` encoded bytes (16.90 MiB), `2788120` compressed data bytes (2.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00010.parquet`: `36424` rows, `2790303` file bytes (2.66 MiB), `28071901` physical bytes (26.77 MiB), `17903223` encoded bytes (17.07 MiB), `2757690` compressed data bytes (2.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00011.parquet`: `36007` rows, `2747202` file bytes (2.62 MiB), `27725869` physical bytes (26.44 MiB), `17710709` encoded bytes (16.89 MiB), `2714806` compressed data bytes (2.59 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00012.parquet`: `36365` rows, `2804936` file bytes (2.67 MiB), `27870003` physical bytes (26.58 MiB), `17882256` encoded bytes (17.05 MiB), `2772805` compressed data bytes (2.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00013.parquet`: `36250` rows, `2826986` file bytes (2.70 MiB), `28063872` physical bytes (26.76 MiB), `17894933` encoded bytes (17.07 MiB), `2794344` compressed data bytes (2.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00014.parquet`: `35142` rows, `3233049` file bytes (3.08 MiB), `24304292` physical bytes (23.18 MiB), `17625518` encoded bytes (16.81 MiB), `3202230` compressed data bytes (3.05 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00015.parquet`: `36092` rows, `3403765` file bytes (3.25 MiB), `25220580` physical bytes (24.05 MiB), `18192726` encoded bytes (17.35 MiB), `3373317` compressed data bytes (3.22 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00016.parquet`: `35295` rows, `3460588` file bytes (3.30 MiB), `23213523` physical bytes (22.14 MiB), `18666218` encoded bytes (17.80 MiB), `3430310` compressed data bytes (3.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00017.parquet`: `34782` rows, `3502716` file bytes (3.34 MiB), `22444441` physical bytes (21.40 MiB), `18992159` encoded bytes (18.11 MiB), `3471935` compressed data bytes (3.31 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00018.parquet`: `34306` rows, `3584877` file bytes (3.42 MiB), `22377859` physical bytes (21.34 MiB), `18919547` encoded bytes (18.04 MiB), `3554535` compressed data bytes (3.39 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00019.parquet`: `34884` rows, `3479712` file bytes (3.32 MiB), `22193699` physical bytes (21.17 MiB), `18888121` encoded bytes (18.01 MiB), `3448957` compressed data bytes (3.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00020.parquet`: `34469` rows, `3537553` file bytes (3.37 MiB), `22170983` physical bytes (21.14 MiB), `18781872` encoded bytes (17.91 MiB), `3506645` compressed data bytes (3.34 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00021.parquet`: `35006` rows, `3492367` file bytes (3.33 MiB), `22384217` physical bytes (21.35 MiB), `18966732` encoded bytes (18.09 MiB), `3461720` compressed data bytes (3.30 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00022.parquet`: `34968` rows, `3488533` file bytes (3.33 MiB), `22373044` physical bytes (21.34 MiB), `18898470` encoded bytes (18.02 MiB), `3458231` compressed data bytes (3.30 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00023.parquet`: `34414` rows, `3563315` file bytes (3.40 MiB), `22375841` physical bytes (21.34 MiB), `19049534` encoded bytes (18.17 MiB), `3532531` compressed data bytes (3.37 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00024.parquet`: `35170` rows, `3488187` file bytes (3.33 MiB), `22346573` physical bytes (21.31 MiB), `18879166` encoded bytes (18.00 MiB), `3458026` compressed data bytes (3.30 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00025.parquet`: `34778` rows, `3497484` file bytes (3.34 MiB), `22178523` physical bytes (21.15 MiB), `18809651` encoded bytes (17.94 MiB), `3466783` compressed data bytes (3.31 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00026.parquet`: `35162` rows, `3483313` file bytes (3.32 MiB), `22379656` physical bytes (21.34 MiB), `18848108` encoded bytes (17.97 MiB), `3452805` compressed data bytes (3.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00027.parquet`: `34944` rows, `3510987` file bytes (3.35 MiB), `22389426` physical bytes (21.35 MiB), `18963685` encoded bytes (18.09 MiB), `3480522` compressed data bytes (3.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-1mib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain/part-00028.parquet`: `5219` rows, `526808` file bytes (514.46 KiB), `3406741` physical bytes (3.25 MiB), `2850807` encoded bytes (2.72 MiB), `511204` compressed data bytes (499.22 KiB)
