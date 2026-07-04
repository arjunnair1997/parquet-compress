# ClickBench Parquet Experiment

- Started: `2026-07-03T19:44:18-04:00`
- Write elapsed: `11.718s`
- Input: `data/clickbench/hits.tsv.gz`
- Output directory: `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain`
- Rows: `1000000`
- Source TSV bytes for rows, reference only: `778360762` (742.30 MiB)
- Parquet physical bytes before page encoding: `712398624` (679.40 MiB)
- Encoded column bytes before codec compression: `821443744` (783.39 MiB)
- Compressed column data bytes after codec compression: `86813070` (82.79 MiB)
- Parquet file bytes: `87713320` (83.65 MiB)
- Physical/encoded ratio: `0.867x`
- Encoded/compressed-data ratio: `9.462x`
- Physical/compressed-data ratio: `8.206x`
- Physical/parquet-file ratio: `8.122x`
- Files: `29`

## Settings

- Compression: `zstd-3`
- Int encoding: `plain`
- String encoding: `plain`
- Date encoding: `rle-dict`
- Timestamp encoding: `plain`
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
- Files read: `29`
- Elapsed: `7.498s`
- Source TSV bytes checked: `778360762` (742.30 MiB)

## Columns

Physical bytes are Parquet physical value payloads before page encoding: fixed-width physical sizes for ints, dates, and timestamps, and BYTE_ARRAY payload bytes after TSV unescaping for strings, excluding PLAIN length prefixes. Encoded bytes are Parquet column chunk total uncompressed sizes after Parquet encoding and before the snappy/zstd codec. Compressed bytes are Parquet column chunk total compressed sizes after the codec. Source field bytes are included only as a TSV reference and exclude delimiters and line endings.

Column stats TSV: [2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain_columns.tsv](../../tsvs/2026-07-03_rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain_columns.tsv)

| Column | Type | Config encoding | Metadata encodings | Page encodings | Values | Physical bytes | Encoded bytes | Compressed bytes | Physical/encoded | Encoded/compressed | Physical/compressed | Source field bytes |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `WatchID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004557` (7.63 MiB) | `8005365` (7.63 MiB) | `0.999x` | `1.000x` | `0.999x` | `19000000` (18.12 MiB) |
| `JavaEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `63619` (62.13 KiB) | `0.999x` | `62.931x` | `62.874x` | `1000000` (976.56 KiB) |
| `Title` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:528` | `1000000` | `138409995` (132.00 MiB) | `142869580` (136.25 MiB) | `13939156` (13.29 MiB) | `0.969x` | `10.250x` | `9.930x` | `138440901` (132.03 MiB) |
| `GoodEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003530` (3.82 MiB) | `4924` (4.81 KiB) | `0.999x` | `813.065x` | `812.348x` | `1000000` (976.56 KiB) |
| `EventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004555` (7.63 MiB) | `2517870` (2.40 MiB) | `0.999x` | `3.179x` | `3.177x` | `19000000` (18.12 MiB) |
| `EventDate` | `date` | `rle-dict` | `PLAIN, RLE_DICTIONARY` | `DATA_PAGE_V2/RLE_DICTIONARY:57, DICTIONARY_PAGE/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4954` (4.84 KiB) | `5980` (5.84 KiB) | `807.428x` | `0.828x` | `668.896x` | `10000000` (9.54 MiB) |
| `CounterID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003528` (3.82 MiB) | `4939` (4.82 KiB) | `0.999x` | `810.595x` | `809.881x` | `2000000` (1.91 MiB) |
| `ClientIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003599` (3.82 MiB) | `408068` (398.50 KiB) | `0.999x` | `9.811x` | `9.802x` | `10032124` (9.57 MiB) |
| `RegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `191175` (186.69 KiB) | `0.999x` | `20.942x` | `20.923x` | `2539898` (2.42 MiB) |
| `UserID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004551` (7.63 MiB) | `618297` (603.81 KiB) | `0.999x` | `12.946x` | `12.939x` | `18637316` (17.77 MiB) |
| `CounterClass` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4240` (4.14 KiB) | `0.999x` | `944.229x` | `943.396x` | `1000000` (976.56 KiB) |
| `OS` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `105900` (103.42 KiB) | `0.999x` | `37.805x` | `37.771x` | `1733675` (1.65 MiB) |
| `UserAgent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `135159` (131.99 KiB) | `0.999x` | `29.621x` | `29.595x` | `1088411` (1.04 MiB) |
| `URL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:356` | `1000000` | `88562192` (84.46 MiB) | `92652667` (88.36 MiB) | `15307029` (14.60 MiB) | `0.956x` | `6.053x` | `5.786x` | `88563396` (84.46 MiB) |
| `Referer` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:326` | `1000000` | `79583339` (75.90 MiB) | `83647730` (79.77 MiB) | `14217928` (13.56 MiB) | `0.951x` | `5.883x` | `5.597x` | `79585848` (75.90 MiB) |
| `IsRefresh` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `178937` (174.74 KiB) | `0.999x` | `22.374x` | `22.354x` | `1000000` (976.56 KiB) |
| `RefererCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `275151` (268.70 KiB) | `0.999x` | `14.550x` | `14.537x` | `4634835` (4.42 MiB) |
| `RefererRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `231384` (225.96 KiB) | `0.999x` | `17.303x` | `17.287x` | `2814059` (2.68 MiB) |
| `URLCategoryID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `87561` (85.51 KiB) | `0.999x` | `45.723x` | `45.682x` | `4525496` (4.32 MiB) |
| `URLRegionID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `48118` (46.99 KiB) | `0.999x` | `83.203x` | `83.129x` | `2981244` (2.84 MiB) |
| `ResolutionWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `187145` (182.76 KiB) | `0.999x` | `21.393x` | `21.374x` | `3967065` (3.78 MiB) |
| `ResolutionHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `186024` (181.66 KiB) | `0.999x` | `21.522x` | `21.503x` | `3407277` (3.25 MiB) |
| `ResolutionDepth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `82097` (80.17 KiB) | `0.999x` | `48.767x` | `48.723x` | `1994256` (1.90 MiB) |
| `FlashMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `53588` (52.33 KiB) | `0.999x` | `74.710x` | `74.644x` | `1923540` (1.83 MiB) |
| `FlashMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `126767` (123.80 KiB) | `0.999x` | `31.582x` | `31.554x` | `1318244` (1.26 MiB) |
| `FlashMinor2` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3354477` (3.20 MiB) | `7357827` (7.02 MiB) | `246551` (240.77 KiB) | `0.456x` | `29.843x` | `13.606x` | `3354477` (3.20 MiB) |
| `NetMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `26146` (25.53 KiB) | `0.999x` | `153.124x` | `152.987x` | `1000000` (976.56 KiB) |
| `NetMinor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `24917` (24.33 KiB) | `0.999x` | `160.677x` | `160.533x` | `1000000` (976.56 KiB) |
| `UserAgentMajor` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `154186` (150.57 KiB) | `0.999x` | `25.966x` | `25.943x` | `1885645` (1.80 MiB) |
| `UserAgentMinor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3767530` (3.59 MiB) | `7770884` (7.41 MiB) | `137048` (133.84 KiB) | `0.485x` | `56.702x` | `27.491x` | `3777059` (3.60 MiB) |
| `CookieEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003537` (3.82 MiB) | `6094` (5.95 KiB) | `0.999x` | `656.964x` | `656.383x` | `1000000` (976.56 KiB) |
| `JavascriptEnable` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003540` (3.82 MiB) | `6508` (6.36 KiB) | `0.999x` | `615.172x` | `614.628x` | `1000000` (976.56 KiB) |
| `IsMobile` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `28686` (28.01 KiB) | `0.999x` | `139.566x` | `139.441x` | `1000000` (976.56 KiB) |
| `MobilePhone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `22518` (21.99 KiB) | `0.999x` | `177.795x` | `177.636x` | `1001922` (978.44 KiB) |
| `MobilePhoneModel` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `81583` (79.67 KiB) | `4084883` (3.90 MiB) | `22608` (22.08 KiB) | `0.020x` | `180.683x` | `3.609x` | `81583` (79.67 KiB) |
| `Params` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002161` (3.82 MiB) | `2872` (2.80 KiB) | `0.000x` | `1393.510x` | `0.000x` | `0` (0 B) |
| `IPNetworkID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003591` (3.82 MiB) | `323779` (316.19 KiB) | `0.999x` | `12.365x` | `12.354x` | `6865544` (6.55 MiB) |
| `TraficSourceID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `288327` (281.57 KiB) | `0.999x` | `13.886x` | `13.873x` | `1728158` (1.65 MiB) |
| `SearchEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003585` (3.82 MiB) | `101625` (99.24 KiB) | `0.999x` | `39.396x` | `39.360x` | `1006573` (982.98 KiB) |
| `SearchPhrase` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3528017` (3.36 MiB) | `7534819` (7.19 MiB) | `719928` (703.05 KiB) | `0.468x` | `10.466x` | `4.901x` | `3528108` (3.36 MiB) |
| `AdvEngineID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003575` (3.82 MiB) | `30698` (29.98 KiB) | `0.999x` | `130.418x` | `130.302x` | `1004631` (981.08 KiB) |
| `IsArtifical` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `164982` (161.12 KiB) | `0.999x` | `24.267x` | `24.245x` | `1000000` (976.56 KiB) |
| `WindowClientWidth` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `305748` (298.58 KiB) | `0.999x` | `13.094x` | `13.083x` | `3825106` (3.65 MiB) |
| `WindowClientHeight` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003591` (3.82 MiB) | `319808` (312.31 KiB) | `0.999x` | `12.519x` | `12.508x` | `3055745` (2.91 MiB) |
| `ClientTimeZone` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `99928` (97.59 KiB) | `0.999x` | `40.065x` | `40.029x` | `2989177` (2.85 MiB) |
| `ClientEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004556` (7.63 MiB) | `2475740` (2.36 MiB) | `0.999x` | `3.233x` | `3.231x` | `19000000` (18.12 MiB) |
| `SilverlightVersion1` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `90618` (88.49 KiB) | `0.999x` | `44.181x` | `44.141x` | `1000017` (976.58 KiB) |
| `SilverlightVersion2` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `74035` (72.30 KiB) | `0.999x` | `54.077x` | `54.029x` | `1000000` (976.56 KiB) |
| `SilverlightVersion3` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `123784` (120.88 KiB) | `0.999x` | `32.343x` | `32.314x` | `2728282` (2.60 MiB) |
| `SilverlightVersion4` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4369` (4.27 KiB) | `0.999x` | `916.349x` | `915.541x` | `1000000` (976.56 KiB) |
| `PageCharset` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:85` | `1000000` | `13587860` (12.96 MiB) | `17594800` (16.78 MiB) | `14611` (14.27 KiB) | `0.772x` | `1204.216x` | `929.975x` | `13587860` (12.96 MiB) |
| `CodeVersion` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003549` (3.82 MiB) | `7063` (6.90 KiB) | `0.999x` | `566.834x` | `566.332x` | `3997297` (3.81 MiB) |
| `IsLink` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003556` (3.82 MiB) | `56801` (55.47 KiB) | `0.999x` | `70.484x` | `70.421x` | `1000000` (976.56 KiB) |
| `IsDownload` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003551` (3.82 MiB) | `8096` (7.91 KiB) | `0.999x` | `494.510x` | `494.071x` | `1000000` (976.56 KiB) |
| `IsNotBounce` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003536` (3.82 MiB) | `25199` (24.61 KiB) | `0.999x` | `158.877x` | `158.736x` | `1000000` (976.56 KiB) |
| `FUniqID` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004554` (7.63 MiB) | `694034` (677.77 KiB) | `0.999x` | `11.533x` | `11.527x` | `18077896` (17.24 MiB) |
| `OriginalURL` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:137` | `1000000` | `27797671` (26.51 MiB) | `31856344` (30.38 MiB) | `5323299` (5.08 MiB) | `0.873x` | `5.984x` | `5.222x` | `27797732` (26.51 MiB) |
| `HID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003641` (3.82 MiB) | `3849794` (3.67 MiB) | `0.999x` | `1.040x` | `1.039x` | `8956330` (8.54 MiB) |
| `IsOldCounter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4240` (4.14 KiB) | `0.999x` | `944.229x` | `943.396x` | `1000000` (976.56 KiB) |
| `IsEvent` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4240` (4.14 KiB) | `0.999x` | `944.229x` | `943.396x` | `1000000` (976.56 KiB) |
| `IsParameter` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4240` (4.14 KiB) | `0.999x` | `944.229x` | `943.396x` | `1000000` (976.56 KiB) |
| `DontCountHits` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `82561` (80.63 KiB) | `0.999x` | `48.492x` | `48.449x` | `1000000` (976.56 KiB) |
| `WithHash` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4240` (4.14 KiB) | `0.999x` | `944.229x` | `943.396x` | `1000000` (976.56 KiB) |
| `HitColor` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `1000000` (976.56 KiB) | `5002903` (4.77 MiB) | `29176` (28.49 KiB) | `0.200x` | `171.473x` | `34.275x` | `1000000` (976.56 KiB) |
| `LocalEventTime` | `timestamp_millis` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004556` (7.63 MiB) | `2520207` (2.40 MiB) | `0.999x` | `3.176x` | `3.174x` | `19000000` (18.12 MiB) |
| `Age` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003588` (3.82 MiB) | `142620` (139.28 KiB) | `0.999x` | `28.072x` | `28.047x` | `1740725` (1.66 MiB) |
| `Sex` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `107976` (105.45 KiB) | `0.999x` | `37.078x` | `37.045x` | `1000000` (976.56 KiB) |
| `Income` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003583` (3.82 MiB) | `122682` (119.81 KiB) | `0.999x` | `32.634x` | `32.605x` | `1000000` (976.56 KiB) |
| `Interests` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003584` (3.82 MiB) | `193520` (188.98 KiB) | `0.999x` | `20.688x` | `20.670x` | `2180312` (2.08 MiB) |
| `Robotness` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003587` (3.82 MiB) | `173704` (169.63 KiB) | `0.999x` | `23.048x` | `23.028x` | `1423051` (1.36 MiB) |
| `RemoteIP` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003601` (3.82 MiB) | `426597` (416.60 KiB) | `0.999x` | `9.385x` | `9.377x` | `10016734` (9.55 MiB) |
| `WindowName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003579` (3.82 MiB) | `70812` (69.15 KiB) | `0.999x` | `56.538x` | `56.488x` | `2197789` (2.10 MiB) |
| `OpenerName` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003530` (3.82 MiB) | `4241` (4.14 KiB) | `0.999x` | `944.006x` | `943.174x` | `2000000` (1.91 MiB) |
| `HistoryLength` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003572` (3.82 MiB) | `55134` (53.84 KiB) | `0.999x` | `72.615x` | `72.551x` | `1940175` (1.85 MiB) |
| `BrowserLanguage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `2001192` (1.91 MiB) | `6004437` (5.73 MiB) | `32328` (31.57 KiB) | `0.333x` | `185.735x` | `61.903x` | `2001192` (1.91 MiB) |
| `BrowserCountry` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3325142` (3.17 MiB) | `7328498` (6.99 MiB) | `122217` (119.35 KiB) | `0.454x` | `59.963x` | `27.207x` | `3987713` (3.80 MiB) |
| `SocialNetwork` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002161` (3.82 MiB) | `2872` (2.80 KiB) | `0.000x` | `1393.510x` | `0.000x` | `0` (0 B) |
| `SocialAction` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002161` (3.82 MiB) | `2872` (2.80 KiB) | `0.000x` | `1393.510x` | `0.000x` | `0` (0 B) |
| `HTTPError` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4240` (4.14 KiB) | `0.999x` | `944.229x` | `943.396x` | `1000000` (976.56 KiB) |
| `SendTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003538` (3.82 MiB) | `61455` (60.01 KiB) | `0.999x` | `65.146x` | `65.088x` | `1035866` (1011.59 KiB) |
| `DNSTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003586` (3.82 MiB) | `135564` (132.39 KiB) | `0.999x` | `29.533x` | `29.506x` | `1026953` (1002.88 KiB) |
| `ConnectTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003599` (3.82 MiB) | `333331` (325.52 KiB) | `0.999x` | `12.011x` | `12.000x` | `1145637` (1.09 MiB) |
| `ResponseStartTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003642` (3.82 MiB) | `1245895` (1.19 MiB) | `0.999x` | `3.213x` | `3.211x` | `2288673` (2.18 MiB) |
| `ResponseEndTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003640` (3.82 MiB) | `937781` (915.80 KiB) | `0.999x` | `4.269x` | `4.265x` | `1573269` (1.50 MiB) |
| `FetchTiming` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003625` (3.82 MiB) | `549831` (536.94 KiB) | `0.999x` | `7.282x` | `7.275x` | `1273464` (1.21 MiB) |
| `SocialSourceNetworkID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003530` (3.82 MiB) | `5318` (5.19 KiB) | `0.999x` | `752.826x` | `752.162x` | `1000056` (976.62 KiB) |
| `SocialSourcePage` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `1024` (1.00 KiB) | `4003990` (3.82 MiB) | `4857` (4.74 KiB) | `0.000x` | `824.375x` | `0.211x` | `1024` (1.00 KiB) |
| `ParamPrice` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004440` (7.63 MiB) | `5725` (5.59 KiB) | `0.999x` | `1398.155x` | `1397.380x` | `1000000` (976.56 KiB) |
| `ParamOrderID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `0` (0 B) | `4002161` (3.82 MiB) | `2872` (2.80 KiB) | `0.000x` | `1393.510x` | `0.000x` | `0` (0 B) |
| `ParamCurrency` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `3000000` (2.86 MiB) | `7003303` (6.68 MiB) | `5261` (5.14 KiB) | `0.428x` | `1331.173x` | `570.234x` | `3000000` (2.86 MiB) |
| `ParamCurrencyID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003529` (3.82 MiB) | `4240` (4.14 KiB) | `0.999x` | `944.229x` | `943.396x` | `1000000` (976.56 KiB) |
| `OpenstatServiceName` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `58030` (56.67 KiB) | `4062525` (3.87 MiB) | `19232` (18.78 KiB) | `0.014x` | `211.238x` | `3.017x` | `58030` (56.67 KiB) |
| `OpenstatCampaignID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `22051` (21.53 KiB) | `4025606` (3.84 MiB) | `16543` (16.16 KiB) | `0.005x` | `243.342x` | `1.333x` | `22051` (21.53 KiB) |
| `OpenstatAdID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `25445` (24.85 KiB) | `4029111` (3.84 MiB) | `19026` (18.58 KiB) | `0.006x` | `211.769x` | `1.337x` | `25445` (24.85 KiB) |
| `OpenstatSourceID` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `48191` (47.06 KiB) | `4051759` (3.86 MiB) | `13322` (13.01 KiB) | `0.012x` | `304.140x` | `3.617x` | `48191` (47.06 KiB) |
| `UTMSource` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `49433` (48.27 KiB) | `4053208` (3.87 MiB) | `21033` (20.54 KiB) | `0.012x` | `192.707x` | `2.350x` | `49433` (48.27 KiB) |
| `UTMMedium` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `16873` (16.48 KiB) | `4019933` (3.83 MiB) | `16985` (16.59 KiB) | `0.004x` | `236.675x` | `0.993x` | `16873` (16.48 KiB) |
| `UTMCampaign` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `91870` (89.72 KiB) | `4097443` (3.91 MiB) | `29859` (29.16 KiB) | `0.022x` | `137.226x` | `3.077x` | `91871` (89.72 KiB) |
| `UTMContent` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `13001` (12.70 KiB) | `4016538` (3.83 MiB) | `13965` (13.64 KiB) | `0.003x` | `287.615x` | `0.931x` | `13001` (12.70 KiB) |
| `UTMTerm` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `28101` (27.44 KiB) | `4032774` (3.85 MiB) | `15726` (15.36 KiB) | `0.007x` | `256.440x` | `1.787x` | `28101` (27.44 KiB) |
| `FromTag` | `string` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `45607` (44.54 KiB) | `4048287` (3.86 MiB) | `28371` (27.71 KiB) | `0.011x` | `142.691x` | `1.608x` | `45607` (44.54 KiB) |
| `HasGCLID` | `int16` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003558` (3.82 MiB) | `21186` (20.69 KiB) | `0.999x` | `188.972x` | `188.804x` | `1000000` (976.56 KiB) |
| `RefererHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004554` (7.63 MiB) | `2842017` (2.71 MiB) | `0.999x` | `2.817x` | `2.815x` | `19349242` (18.45 MiB) |
| `URLHash` | `int64` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `8000000` (7.63 MiB) | `8004553` (7.63 MiB) | `3580493` (3.41 MiB) | `0.999x` | `2.236x` | `2.234x` | `19343177` (18.45 MiB) |
| `CLID` | `int32` | `plain` | `PLAIN` | `DATA_PAGE_V2/PLAIN:57` | `1000000` | `4000000` (3.81 MiB) | `4003543` (3.82 MiB) | `5642` (5.51 KiB) | `0.999x` | `709.596x` | `708.968x` | `1000380` (976.93 KiB) |

## Files

- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00000.parquet`: `35889` rows, `2840401` file bytes (2.71 MiB), `27875599` physical bytes (26.58 MiB), `31797643` encoded bytes (30.32 MiB), `2808326` compressed data bytes (2.68 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00001.parquet`: `35984` rows, `2790998` file bytes (2.66 MiB), `27645074` physical bytes (26.36 MiB), `31586087` encoded bytes (30.12 MiB), `2758381` compressed data bytes (2.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00002.parquet`: `35875` rows, `2831962` file bytes (2.70 MiB), `28091748` physical bytes (26.79 MiB), `32022406` encoded bytes (30.54 MiB), `2798657` compressed data bytes (2.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00003.parquet`: `36277` rows, `2803693` file bytes (2.67 MiB), `28077471` physical bytes (26.78 MiB), `32048770` encoded bytes (30.56 MiB), `2771255` compressed data bytes (2.64 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00004.parquet`: `35612` rows, `2768305` file bytes (2.64 MiB), `27705258` physical bytes (26.42 MiB), `31606154` encoded bytes (30.14 MiB), `2735089` compressed data bytes (2.61 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00005.parquet`: `36089` rows, `2778159` file bytes (2.65 MiB), `27666502` physical bytes (26.38 MiB), `31613481` encoded bytes (30.15 MiB), `2745667` compressed data bytes (2.62 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00006.parquet`: `36248` rows, `2820164` file bytes (2.69 MiB), `28058832` physical bytes (26.76 MiB), `32030321` encoded bytes (30.55 MiB), `2787694` compressed data bytes (2.66 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00007.parquet`: `36246` rows, `2828740` file bytes (2.70 MiB), `28062130` physical bytes (26.76 MiB), `32029295` encoded bytes (30.55 MiB), `2795811` compressed data bytes (2.67 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00008.parquet`: `36295` rows, `2773444` file bytes (2.64 MiB), `28105131` physical bytes (26.80 MiB), `32073434` encoded bytes (30.59 MiB), `2740680` compressed data bytes (2.61 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00009.parquet`: `35881` rows, `2788934` file bytes (2.66 MiB), `27676562` physical bytes (26.39 MiB), `31600606` encoded bytes (30.14 MiB), `2756743` compressed data bytes (2.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00010.parquet`: `36424` rows, `2783407` file bytes (2.65 MiB), `28067097` physical bytes (26.77 MiB), `32055733` encoded bytes (30.57 MiB), `2750778` compressed data bytes (2.62 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00011.parquet`: `36019` rows, `2728665` file bytes (2.60 MiB), `27731317` physical bytes (26.45 MiB), `31675062` encoded bytes (30.21 MiB), `2696270` compressed data bytes (2.57 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00012.parquet`: `36350` rows, `2791581` file bytes (2.66 MiB), `27863603` physical bytes (26.57 MiB), `31836592` encoded bytes (30.36 MiB), `2759436` compressed data bytes (2.63 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00013.parquet`: `36240` rows, `2814035` file bytes (2.68 MiB), `28059163` physical bytes (26.76 MiB), `32024818` encoded bytes (30.54 MiB), `2781345` compressed data bytes (2.65 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00014.parquet`: `35148` rows, `3218793` file bytes (3.07 MiB), `24298781` physical bytes (23.17 MiB), `28124669` encoded bytes (26.82 MiB), `3187965` compressed data bytes (3.04 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00015.parquet`: `36070` rows, `3406155` file bytes (3.25 MiB), `25205990` physical bytes (24.04 MiB), `29128491` encoded bytes (27.78 MiB), `3375717` compressed data bytes (3.22 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00016.parquet`: `35325` rows, `3428789` file bytes (3.27 MiB), `23230806` physical bytes (22.15 MiB), `27070324` encoded bytes (25.82 MiB), `3398515` compressed data bytes (3.24 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00017.parquet`: `34862` rows, `3455104` file bytes (3.30 MiB), `22496277` physical bytes (21.45 MiB), `26282633` encoded bytes (25.07 MiB), `3424306` compressed data bytes (3.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00018.parquet`: `34371` rows, `3542240` file bytes (3.38 MiB), `22421574` physical bytes (21.38 MiB), `26154885` encoded bytes (24.94 MiB), `3511883` compressed data bytes (3.35 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00019.parquet`: `34981` rows, `3423427` file bytes (3.26 MiB), `22256396` physical bytes (21.23 MiB), `26055831` encoded bytes (24.85 MiB), `3392663` compressed data bytes (3.24 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00020.parquet`: `34541` rows, `3495273` file bytes (3.33 MiB), `22210413` physical bytes (21.18 MiB), `25961846` encoded bytes (24.76 MiB), `3464335` compressed data bytes (3.30 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00021.parquet`: `35110` rows, `3441487` file bytes (3.28 MiB), `22437635` physical bytes (21.40 MiB), `26250313` encoded bytes (25.03 MiB), `3410832` compressed data bytes (3.25 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00022.parquet`: `34848` rows, `3455884` file bytes (3.30 MiB), `22448651` physical bytes (21.41 MiB), `26232116` encoded bytes (25.02 MiB), `3425581` compressed data bytes (3.27 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00023.parquet`: `34665` rows, `3507549` file bytes (3.35 MiB), `22408521` physical bytes (21.37 MiB), `26172786` encoded bytes (24.96 MiB), `3476767` compressed data bytes (3.32 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00024.parquet`: `35257` rows, `3438269` file bytes (3.28 MiB), `22397011` physical bytes (21.36 MiB), `26224182` encoded bytes (25.01 MiB), `3408105` compressed data bytes (3.25 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00025.parquet`: `34867` rows, `3446905` file bytes (3.29 MiB), `22225491` physical bytes (21.20 MiB), `26011865` encoded bytes (24.81 MiB), `3416199` compressed data bytes (3.26 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00026.parquet`: `35306` rows, `3426529` file bytes (3.27 MiB), `22447576` physical bytes (21.41 MiB), `26281260` encoded bytes (25.06 MiB), `3395991` compressed data bytes (3.24 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00027.parquet`: `35237` rows, `3482351` file bytes (3.32 MiB), `22630942` physical bytes (21.58 MiB), `26457336` encoded bytes (25.23 MiB), `3451769` compressed data bytes (3.29 MiB)
- `encoding_experiment/page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows/parquet/rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain/part-00028.parquet`: `3983` rows, `402077` file bytes (392.65 KiB), `2597073` physical bytes (2.48 MiB), `3034805` encoded bytes (2.89 MiB), `386310` compressed data bytes (377.26 KiB)
