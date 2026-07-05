# Column Top 5 Encoding Rankings

- Experiment: `page-256kib-rgsize-10mib-file-10mib-dictpage-256kib/1000000_rows`
- Source data: [2026-07-05_rows-1000000_encoding-matrix_column-results.tsv](../tsvs/2026-07-05_rows-1000000_encoding-matrix_column-results.tsv)
- Rows: `1,000,000`
- Ranking metric: per-column `compressed_bytes`, after Parquet page encoding and Snappy/ZSTD compression.
- Each numbered item starts with the achieved compressed size for that encoding/compression choice.
- Duplicate matrix rows with the same effective column encoding are collapsed to the best observed row before ranking.
- Encodings in this matrix: `plain`, `rle-dict`, `delta-binary-packed`, `delta-byte-array`, `delta-length-byte-array`.
- Column shape stats: [column_shape_stats.json](column_shape_stats/column_shape_stats.json)

## Winner Distribution

Counts are based on each column's first `Compressed overall` ranking below: one winner per column, grouped by compression algorithm and configured column encoding.

![Column winner distribution](images/column_winner_distribution.svg)

| Compression | Encoding | Column wins |
| --- | --- | ---: |
| `zstd-3` | `plain` | 54 |
| `zstd-3` | `rle-dict` | 46 |
| `snappy` | `rle-dict` | 2 |
| `zstd-3` | `delta-binary-packed` | 2 |
| `snappy` | `plain` | 1 |

## Encoding Rank Distribution

For each column and compression codec, duplicate matrix rows with the same effective column encoding are collapsed to the smallest compressed byte count. The remaining encodings are sorted by compressed bytes; counts below show how often each compression + encoding landed at rank 1, rank 2, and so on. Encodings that are not valid for a column type are not counted for that column.

![Encoding rank distribution by compression](images/encoding_rank_distribution.svg)

| Compression | Encoding | Ranked columns | Rank 1 | Rank 2 | Rank 3 | Rank 4 |
| --- | --- | ---: | ---: | ---: | ---: | ---: |
| `zstd-3` | `plain` | 105 | 55 | 41 | 9 | 0 |
| `zstd-3` | `rle-dict` | 105 | 48 | 30 | 22 | 5 |
| `zstd-3` | `delta-binary-packed` | 77 | 2 | 26 | 49 | 0 |
| `zstd-3` | `delta-byte-array` | 28 | 0 | 1 | 8 | 19 |
| `zstd-3` | `delta-length-byte-array` | 28 | 0 | 7 | 17 | 4 |
| `snappy` | `plain` | 105 | 7 | 18 | 56 | 24 |
| `snappy` | `rle-dict` | 105 | 91 | 8 | 6 | 0 |
| `snappy` | `delta-binary-packed` | 77 | 7 | 51 | 19 | 0 |
| `snappy` | `delta-byte-array` | 28 | 0 | 12 | 16 | 0 |
| `snappy` | `delta-length-byte-array` | 28 | 0 | 16 | 8 | 4 |

## ZSTD Plain Winner Second-Place Distribution

For columns where `zstd + plain` is rank 1 in the ZSTD-only compressed-byte ranking, this counts which encoding landed at rank 2 after collapsing duplicate matrix rows to each encoding's smallest compressed byte count.

![ZSTD plain winner second-place distribution](images/zstd_plain_winner_second_place_distribution.svg)

- Columns where `zstd + plain` ranked first: `55`
- Missing second-place rows: `0`

| Second-place encoding | Columns |
| --- | ---: |
| `zstd + rle-dict` | 29 |
| `zstd + delta-binary-packed` | 20 |
| `zstd + delta-length-byte-array` | 6 |

## ZSTD Plain vs RLE Dict Improvement Distribution

For each column, this compares the best observed `zstd + plain` compressed byte count with the best observed `zstd + rle-dict` compressed byte count. Improvement is `(larger compressed bytes - smaller compressed bytes) / larger compressed bytes * 100`.

![ZSTD plain versus RLE dictionary improvement distribution](images/zstd_plain_vs_rle_dict_improvement.svg)

- Compared columns: `105`
- `zstd + plain` smaller: `56`; `zstd + rle-dict` smaller: `49`; ties: `0`; missing comparisons: `0`

| Improvement bucket | `zstd + plain` better | `zstd + rle-dict` better |
| --- | ---: | ---: |
| `0-10%` | 6 | 13 |
| `10-20%` | 19 | 15 |
| `20-30%` | 15 | 10 |
| `30-40%` | 7 | 5 |
| `40-50%` | 8 | 4 |
| `50-60%` | 1 | 2 |

### Page-Level Winner Distribution

This is the page-level version of the same `plain + zstd` vs `rle-dict + zstd` comparison. Page ranges differ between the two runs, so the distribution is computed over overlap windows from the union of page row ranges. RLE dictionary bytes include `compressed_page_bytes_with_amortized_dictionary`, where each column chunk's compressed dictionary page bytes are spread evenly across that chunk's data pages before comparing windows.

- Source TSV: [2026-07-05_rows-1000000_plain-zstd_vs_rle-dict-zstd_page-distribution.tsv](../tsvs/page_encoding_distribution/2026-07-05_rows-1000000_plain-zstd_vs_rle-dict-zstd_page-distribution.tsv)
- Compared columns: `105`; mixed page winners: `84`; plain-only columns: `16`; rle-dict-only columns: `5`; tie-only columns: `0`
- Overlap windows: `13,449`; `plain + zstd` wins: `7,557`; `rle-dict + zstd` wins: `5,892`; ties: `0`
- Row-weighted wins: `plain + zstd` `64,862,049` (`61.77%`); `rle-dict + zstd` `40,137,951` (`38.23%`); ties `0` (`0.00%`)
- Allocated page bytes: `plain + zstd` `86,812,277 B (82.79 MiB)`; `rle-dict + zstd` `86,384,156 B (82.38 MiB)`; rle-dict vs plain `-0.493157%`
- Exact matched page ranges: `0`; unmatched plain pages: `7,254`; unmatched rle-dict pages: `6,300`

![Page-window winner distribution by column](page_encoding_distribution/images/2026-07-05_rows-1000000_plain-zstd_vs_rle-dict-zstd_page-distribution.svg)

| Column | Type | Windows | Plain wins | RLE dict wins | Ties | Row-weighted plain | Row-weighted RLE dict | Plain allocated bytes | RLE dict allocated bytes | RLE dict vs plain | Exact matches | Unmatched pages |
| --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| `BrowserCountry` | `STRING` | `116` | `0` (`0.00%`) | `116` (`100.00%`) | `0` (`0.00%`) | `0` (`0.00%`) | `1,000,000` (`100.00%`) | `122,222 B (119.36 KiB)` | `69,669 B (68.04 KiB)` | `-42.997987%` | `0` | `57 / 60` |
| `FlashMinor2` | `STRING` | `116` | `0` (`0.00%`) | `116` (`100.00%`) | `0` (`0.00%`) | `0` (`0.00%`) | `1,000,000` (`100.00%`) | `246,539 B (240.76 KiB)` | `146,647 B (143.21 KiB)` | `-40.517727%` | `0` | `57 / 60` |
| `RefererRegionID` | `INT(32,true)` | `116` | `0` (`0.00%`) | `116` (`100.00%`) | `0` (`0.00%`) | `0` (`0.00%`) | `1,000,000` (`100.00%`) | `231,356 B (225.93 KiB)` | `165,218 B (161.35 KiB)` | `-28.587113%` | `0` | `57 / 60` |
| `Sex` | `INT(16,true)` | `116` | `0` (`0.00%`) | `116` (`100.00%`) | `0` (`0.00%`) | `0` (`0.00%`) | `1,000,000` (`100.00%`) | `107,976 B (105.45 KiB)` | `80,827 B (78.93 KiB)` | `-25.143550%` | `0` | `57 / 60` |
| `UserAgentMinor` | `STRING` | `116` | `0` (`0.00%`) | `116` (`100.00%`) | `0` (`0.00%`) | `0` (`0.00%`) | `1,000,000` (`100.00%`) | `137,050 B (133.84 KiB)` | `83,983 B (82.01 KiB)` | `-38.720905%` | `0` | `57 / 60` |
| `Title` | `STRING` | `635` | `3` (`0.47%`) | `632` (`99.53%`) | `0` (`0.00%`) | `17,791` (`1.78%`) | `982,209` (`98.22%`) | `13,939,135 B (13.29 MiB)` | `7,975,230 B (7.61 MiB)` | `-42.785331%` | `0` | `576 / 60` |
| `RefererCategoryID` | `INT(16,true)` | `116` | `1` (`0.86%`) | `115` (`99.14%`) | `0` (`0.00%`) | `639` (`0.06%`) | `999,361` (`99.94%`) | `275,135 B (268.69 KiB)` | `216,084 B (211.02 KiB)` | `-21.462555%` | `0` | `57 / 60` |
| `TraficSourceID` | `INT(16,true)` | `116` | `1` (`0.86%`) | `115` (`99.14%`) | `0` (`0.00%`) | `639` (`0.06%`) | `999,361` (`99.94%`) | `288,326 B (281.57 KiB)` | `178,371 B (174.19 KiB)` | `-38.135652%` | `0` | `57 / 60` |
| `ClientTimeZone` | `INT(16,true)` | `116` | `2` (`1.72%`) | `114` (`98.28%`) | `0` (`0.00%`) | `1,315` (`0.13%`) | `998,685` (`99.87%`) | `99,924 B (97.58 KiB)` | `85,542 B (83.54 KiB)` | `-14.392939%` | `0` | `57 / 60` |
| `Income` | `INT(16,true)` | `116` | `2` (`1.72%`) | `114` (`98.28%`) | `0` (`0.00%`) | `807` (`0.08%`) | `999,193` (`99.92%`) | `122,679 B (119.80 KiB)` | `89,413 B (87.32 KiB)` | `-27.116295%` | `0` | `57 / 60` |
| `Age` | `INT(16,true)` | `116` | `3` (`2.59%`) | `113` (`97.41%`) | `0` (`0.00%`) | `17,288` (`1.73%`) | `982,712` (`98.27%`) | `142,624 B (139.28 KiB)` | `118,882 B (116.10 KiB)` | `-16.646567%` | `0` | `57 / 60` |
| `IsRefresh` | `INT(16,true)` | `116` | `4` (`3.45%`) | `112` (`96.55%`) | `0` (`0.00%`) | `29,594` (`2.96%`) | `970,406` (`97.04%`) | `178,941 B (174.75 KiB)` | `84,103 B (82.13 KiB)` | `-52.999592%` | `0` | `57 / 60` |
| `Referer` | `STRING` | `411` | `18` (`4.38%`) | `393` (`95.62%`) | `0` (`0.00%`) | `50,233` (`5.02%`) | `949,767` (`94.98%`) | `14,218,039 B (13.56 MiB)` | `11,752,116 B (11.21 MiB)` | `-17.343622%` | `0` | `352 / 60` |
| `URL` | `STRING` | `449` | `22` (`4.90%`) | `427` (`95.10%`) | `0` (`0.00%`) | `44,586` (`4.46%`) | `955,414` (`95.54%`) | `15,306,955 B (14.60 MiB)` | `12,707,286 B (12.12 MiB)` | `-16.983580%` | `0` | `390 / 60` |
| `URLCategoryID` | `INT(16,true)` | `116` | `7` (`6.03%`) | `109` (`93.97%`) | `0` (`0.00%`) | `16,054` (`1.61%`) | `983,946` (`98.39%`) | `87,570 B (85.52 KiB)` | `72,205 B (70.51 KiB)` | `-17.545963%` | `0` | `57 / 60` |
| `SilverlightVersion2` | `INT(16,true)` | `116` | `8` (`6.90%`) | `108` (`93.10%`) | `0` (`0.00%`) | `59,183` (`5.92%`) | `940,817` (`94.08%`) | `74,028 B (72.29 KiB)` | `56,776 B (55.45 KiB)` | `-23.304696%` | `0` | `57 / 60` |
| `URLRegionID` | `INT(32,true)` | `116` | `8` (`6.90%`) | `108` (`93.10%`) | `0` (`0.00%`) | `44,477` (`4.45%`) | `955,523` (`95.55%`) | `48,124 B (47.00 KiB)` | `40,697 B (39.74 KiB)` | `-15.433048%` | `0` | `57 / 60` |
| `ResolutionDepth` | `INT(16,true)` | `116` | `9` (`7.76%`) | `107` (`92.24%`) | `0` (`0.00%`) | `89,540` (`8.95%`) | `910,460` (`91.05%`) | `82,090 B (80.17 KiB)` | `62,736 B (61.27 KiB)` | `-23.576562%` | `0` | `57 / 60` |
| `JavaEnable` | `INT(16,true)` | `116` | `13` (`11.21%`) | `103` (`88.79%`) | `0` (`0.00%`) | `99,698` (`9.97%`) | `900,302` (`90.03%`) | `63,644 B (62.15 KiB)` | `51,484 B (50.28 KiB)` | `-19.106279%` | `0` | `57 / 60` |
| `SearchPhrase` | `STRING` | `116` | `18` (`15.52%`) | `98` (`84.48%`) | `0` (`0.00%`) | `145,597` (`14.56%`) | `854,403` (`85.44%`) | `719,930 B (703.06 KiB)` | `636,259 B (621.35 KiB)` | `-11.622102%` | `0` | `57 / 60` |
| `BrowserLanguage` | `STRING` | `116` | `19` (`16.38%`) | `97` (`83.62%`) | `0` (`0.00%`) | `108,497` (`10.85%`) | `891,503` (`89.15%`) | `32,361 B (31.60 KiB)` | `27,773 B (27.12 KiB)` | `-14.177559%` | `0` | `57 / 60` |
| `OriginalURL` | `STRING` | `210` | `38` (`18.10%`) | `172` (`81.90%`) | `0` (`0.00%`) | `107,992` (`10.80%`) | `892,008` (`89.20%`) | `5,323,588 B (5.08 MiB)` | `4,887,325 B (4.66 MiB)` | `-8.194905%` | `0` | `151 / 60` |
| `SearchEngineID` | `INT(16,true)` | `116` | `27` (`23.28%`) | `89` (`76.72%`) | `0` (`0.00%`) | `203,249` (`20.32%`) | `796,751` (`79.68%`) | `101,613 B (99.23 KiB)` | `76,937 B (75.13 KiB)` | `-24.284294%` | `0` | `57 / 60` |
| `FlashMinor` | `INT(16,true)` | `116` | `28` (`24.14%`) | `88` (`75.86%`) | `0` (`0.00%`) | `229,133` (`22.91%`) | `770,867` (`77.09%`) | `126,766 B (123.79 KiB)` | `114,356 B (111.68 KiB)` | `-9.789691%` | `0` | `57 / 60` |
| `SilverlightVersion1` | `INT(16,true)` | `116` | `31` (`26.72%`) | `85` (`73.28%`) | `0` (`0.00%`) | `269,583` (`26.96%`) | `730,417` (`73.04%`) | `90,600 B (88.48 KiB)` | `84,773 B (82.79 KiB)` | `-6.431567%` | `0` | `57 / 60` |
| `PageCharset` | `STRING` | `144` | `40` (`27.78%`) | `104` (`72.22%`) | `0` (`0.00%`) | `313,297` (`31.33%`) | `686,703` (`68.67%`) | `14,612 B (14.27 KiB)` | `10,012 B (9.78 KiB)` | `-31.480975%` | `0` | `85 / 60` |
| `IsMobile` | `INT(16,true)` | `116` | `33` (`28.45%`) | `83` (`71.55%`) | `0` (`0.00%`) | `297,984` (`29.80%`) | `702,016` (`70.20%`) | `28,686 B (28.01 KiB)` | `24,770 B (24.19 KiB)` | `-13.651258%` | `0` | `57 / 60` |
| `FlashMajor` | `INT(16,true)` | `116` | `35` (`30.17%`) | `81` (`69.83%`) | `0` (`0.00%`) | `321,814` (`32.18%`) | `678,186` (`67.82%`) | `53,604 B (52.35 KiB)` | `49,164 B (48.01 KiB)` | `-8.282964%` | `0` | `57 / 60` |
| `OpenstatServiceName` | `STRING` | `116` | `35` (`30.17%`) | `81` (`69.83%`) | `0` (`0.00%`) | `240,681` (`24.07%`) | `759,319` (`75.93%`) | `19,232 B (18.78 KiB)` | `17,819 B (17.40 KiB)` | `-7.347130%` | `0` | `57 / 60` |
| `UserAgent` | `INT(16,true)` | `116` | `36` (`31.03%`) | `80` (`68.97%`) | `0` (`0.00%`) | `276,895` (`27.69%`) | `723,105` (`72.31%`) | `135,169 B (132.00 KiB)` | `125,928 B (122.98 KiB)` | `-6.836627%` | `0` | `57 / 60` |
| `IsArtifical` | `INT(16,true)` | `116` | `39` (`33.62%`) | `77` (`66.38%`) | `0` (`0.00%`) | `326,111` (`32.61%`) | `673,889` (`67.39%`) | `164,989 B (161.12 KiB)` | `82,436 B (80.50 KiB)` | `-50.035457%` | `0` | `57 / 60` |
| `MobilePhoneModel` | `STRING` | `116` | `39` (`33.62%`) | `77` (`66.38%`) | `0` (`0.00%`) | `334,137` (`33.41%`) | `665,863` (`66.59%`) | `22,607 B (22.08 KiB)` | `20,619 B (20.14 KiB)` | `-8.793736%` | `0` | `57 / 60` |
| `UTMCampaign` | `STRING` | `116` | `39` (`33.62%`) | `77` (`66.38%`) | `0` (`0.00%`) | `363,098` (`36.31%`) | `636,902` (`63.69%`) | `29,858 B (29.16 KiB)` | `27,397 B (26.75 KiB)` | `-8.242347%` | `0` | `57 / 60` |
| `UTMMedium` | `STRING` | `116` | `39` (`33.62%`) | `77` (`66.38%`) | `0` (`0.00%`) | `328,680` (`32.87%`) | `671,320` (`67.13%`) | `16,986 B (16.59 KiB)` | `14,496 B (14.16 KiB)` | `-14.659131%` | `0` | `57 / 60` |
| `DontCountHits` | `INT(16,true)` | `116` | `40` (`34.48%`) | `76` (`65.52%`) | `0` (`0.00%`) | `379,813` (`37.98%`) | `620,187` (`62.02%`) | `82,558 B (80.62 KiB)` | `50,501 B (49.32 KiB)` | `-38.829671%` | `0` | `57 / 60` |
| `UTMSource` | `STRING` | `116` | `41` (`35.34%`) | `75` (`64.66%`) | `0` (`0.00%`) | `367,562` (`36.76%`) | `632,438` (`63.24%`) | `21,033 B (20.54 KiB)` | `18,473 B (18.04 KiB)` | `-12.171350%` | `0` | `57 / 60` |
| `OpenstatAdID` | `STRING` | `116` | `46` (`39.66%`) | `70` (`60.34%`) | `0` (`0.00%`) | `329,773` (`32.98%`) | `670,227` (`67.02%`) | `19,026 B (18.58 KiB)` | `18,400 B (17.97 KiB)` | `-3.290234%` | `0` | `57 / 60` |
| `HitColor` | `STRING` | `116` | `47` (`40.52%`) | `69` (`59.48%`) | `0` (`0.00%`) | `344,752` (`34.48%`) | `655,248` (`65.52%`) | `29,173 B (28.49 KiB)` | `25,054 B (24.47 KiB)` | `-14.119220%` | `0` | `57 / 60` |
| `OpenstatCampaignID` | `STRING` | `116` | `50` (`43.10%`) | `66` (`56.90%`) | `0` (`0.00%`) | `359,749` (`35.97%`) | `640,251` (`64.03%`) | `16,542 B (16.15 KiB)` | `15,897 B (15.52 KiB)` | `-3.899166%` | `0` | `57 / 60` |
| `AdvEngineID` | `INT(16,true)` | `116` | `51` (`43.97%`) | `65` (`56.03%`) | `0` (`0.00%`) | `431,822` (`43.18%`) | `568,178` (`56.82%`) | `30,696 B (29.98 KiB)` | `23,249 B (22.70 KiB)` | `-24.260490%` | `0` | `57 / 60` |
| `UTMTerm` | `STRING` | `116` | `52` (`44.83%`) | `64` (`55.17%`) | `0` (`0.00%`) | `402,277` (`40.23%`) | `597,723` (`59.77%`) | `15,728 B (15.36 KiB)` | `16,158 B (15.78 KiB)` | `2.733978%` | `0` | `57 / 60` |
| `IsLink` | `INT(16,true)` | `116` | `56` (`48.28%`) | `60` (`51.72%`) | `0` (`0.00%`) | `505,429` (`50.54%`) | `494,571` (`49.46%`) | `56,781 B (55.45 KiB)` | `37,020 B (36.15 KiB)` | `-34.802135%` | `0` | `57 / 60` |
| `NetMinor` | `INT(16,true)` | `116` | `58` (`50.00%`) | `58` (`50.00%`) | `0` (`0.00%`) | `453,837` (`45.38%`) | `546,163` (`54.62%`) | `24,917 B (24.33 KiB)` | `24,001 B (23.44 KiB)` | `-3.676205%` | `0` | `57 / 60` |
| `HasGCLID` | `INT(16,true)` | `116` | `61` (`52.59%`) | `55` (`47.41%`) | `0` (`0.00%`) | `504,977` (`50.50%`) | `495,023` (`49.50%`) | `21,186 B (20.69 KiB)` | `17,019 B (16.62 KiB)` | `-19.668649%` | `0` | `57 / 60` |
| `FromTag` | `STRING` | `116` | `63` (`54.31%`) | `53` (`45.69%`) | `0` (`0.00%`) | `572,451` (`57.25%`) | `427,549` (`42.75%`) | `28,372 B (27.71 KiB)` | `22,821 B (22.29 KiB)` | `-19.565064%` | `0` | `57 / 60` |
| `OpenstatSourceID` | `STRING` | `116` | `63` (`54.31%`) | `53` (`45.69%`) | `0` (`0.00%`) | `502,301` (`50.23%`) | `497,699` (`49.77%`) | `13,322 B (13.01 KiB)` | `12,930 B (12.63 KiB)` | `-2.942501%` | `0` | `57 / 60` |
| `NetMajor` | `INT(16,true)` | `116` | `64` (`55.17%`) | `52` (`44.83%`) | `0` (`0.00%`) | `500,986` (`50.10%`) | `499,014` (`49.90%`) | `26,145 B (25.53 KiB)` | `26,304 B (25.69 KiB)` | `0.608147%` | `0` | `57 / 60` |
| `UTMContent` | `STRING` | `116` | `65` (`56.03%`) | `51` (`43.97%`) | `0` (`0.00%`) | `577,387` (`57.74%`) | `422,613` (`42.26%`) | `13,966 B (13.64 KiB)` | `15,105 B (14.75 KiB)` | `8.155521%` | `0` | `57 / 60` |
| `HID` | `INT(32,true)` | `116` | `67` (`57.76%`) | `49` (`42.24%`) | `0` (`0.00%`) | `572,460` (`57.25%`) | `427,540` (`42.75%`) | `3,849,795 B (3.67 MiB)` | `4,492,493 B (4.28 MiB)` | `16.694343%` | `0` | `57 / 60` |
| `SilverlightVersion3` | `INT(32,true)` | `116` | `70` (`60.34%`) | `46` (`39.66%`) | `0` (`0.00%`) | `640,580` (`64.06%`) | `359,420` (`35.94%`) | `123,785 B (120.88 KiB)` | `124,901 B (121.97 KiB)` | `0.901563%` | `0` | `57 / 60` |
| `MobilePhone` | `INT(16,true)` | `116` | `79` (`68.10%`) | `37` (`31.90%`) | `0` (`0.00%`) | `714,212` (`71.42%`) | `285,788` (`28.58%`) | `22,518 B (21.99 KiB)` | `24,212 B (23.64 KiB)` | `7.522871%` | `0` | `57 / 60` |
| `OS` | `INT(16,true)` | `116` | `85` (`73.28%`) | `31` (`26.72%`) | `0` (`0.00%`) | `739,875` (`73.99%`) | `260,125` (`26.01%`) | `105,905 B (103.42 KiB)` | `121,811 B (118.96 KiB)` | `15.019121%` | `0` | `57 / 60` |
| `HistoryLength` | `INT(16,true)` | `116` | `89` (`76.72%`) | `27` (`23.28%`) | `0` (`0.00%`) | `865,995` (`86.60%`) | `134,005` (`13.40%`) | `55,137 B (53.84 KiB)` | `51,071 B (49.87 KiB)` | `-7.374358%` | `0` | `57 / 60` |
| `IsDownload` | `INT(16,true)` | `116` | `96` (`82.76%`) | `20` (`17.24%`) | `0` (`0.00%`) | `852,847` (`85.28%`) | `147,153` (`14.72%`) | `8,096 B (7.91 KiB)` | `9,596 B (9.37 KiB)` | `18.527668%` | `0` | `57 / 60` |
| `UserAgentMajor` | `INT(16,true)` | `116` | `96` (`82.76%`) | `20` (`17.24%`) | `0` (`0.00%`) | `813,364` (`81.34%`) | `186,636` (`18.66%`) | `154,189 B (150.58 KiB)` | `180,166 B (175.94 KiB)` | `16.847505%` | `0` | `57 / 60` |
| `WindowName` | `INT(32,true)` | `116` | `98` (`84.48%`) | `18` (`15.52%`) | `0` (`0.00%`) | `949,167` (`94.92%`) | `50,833` (`5.08%`) | `70,817 B (69.16 KiB)` | `148,984 B (145.49 KiB)` | `110.378864%` | `0` | `57 / 60` |
| `JavascriptEnable` | `INT(16,true)` | `116` | `101` (`87.07%`) | `15` (`12.93%`) | `0` (`0.00%`) | `934,689` (`93.47%`) | `65,311` (`6.53%`) | `6,510 B (6.36 KiB)` | `7,740 B (7.56 KiB)` | `18.894009%` | `0` | `57 / 60` |
| `CookieEnable` | `INT(16,true)` | `116` | `102` (`87.93%`) | `14` (`12.07%`) | `0` (`0.00%`) | `958,225` (`95.82%`) | `41,775` (`4.18%`) | `6,094 B (5.95 KiB)` | `7,425 B (7.25 KiB)` | `21.841155%` | `0` | `57 / 60` |
| `DNSTiming` | `INT(32,true)` | `116` | `102` (`87.93%`) | `14` (`12.07%`) | `0` (`0.00%`) | `891,026` (`89.10%`) | `108,974` (`10.90%`) | `135,564 B (132.39 KiB)` | `158,368 B (154.66 KiB)` | `16.821575%` | `0` | `57 / 60` |
| `IsNotBounce` | `INT(16,true)` | `116` | `102` (`87.93%`) | `14` (`12.07%`) | `0` (`0.00%`) | `874,119` (`87.41%`) | `125,881` (`12.59%`) | `25,193 B (24.60 KiB)` | `17,680 B (17.27 KiB)` | `-29.821776%` | `0` | `57 / 60` |
| `ResolutionWidth` | `INT(16,true)` | `116` | `102` (`87.93%`) | `14` (`12.07%`) | `0` (`0.00%`) | `868,251` (`86.83%`) | `131,749` (`13.17%`) | `187,141 B (182.75 KiB)` | `205,670 B (200.85 KiB)` | `9.901091%` | `0` | `57 / 60` |
| `SocialSourcePage` | `STRING` | `116` | `102` (`87.93%`) | `14` (`12.07%`) | `0` (`0.00%`) | `962,871` (`96.29%`) | `37,129` (`3.71%`) | `4,856 B (4.74 KiB)` | `6,940 B (6.78 KiB)` | `42.915980%` | `0` | `57 / 60` |
| `CodeVersion` | `INT(32,true)` | `116` | `103` (`88.79%`) | `13` (`11.21%`) | `0` (`0.00%`) | `953,155` (`95.32%`) | `46,845` (`4.68%`) | `7,064 B (6.90 KiB)` | `8,314 B (8.12 KiB)` | `17.695357%` | `0` | `57 / 60` |
| `ConnectTiming` | `INT(32,true)` | `116` | `107` (`92.24%`) | `9` (`7.76%`) | `0` (`0.00%`) | `979,098` (`97.91%`) | `20,902` (`2.09%`) | `333,314 B (325.50 KiB)` | `380,387 B (371.47 KiB)` | `14.122719%` | `0` | `57 / 60` |
| `RefererHash` | `INT(64,true)` | `116` | `107` (`92.24%`) | `9` (`7.76%`) | `0` (`0.00%`) | `942,696` (`94.27%`) | `57,304` (`5.73%`) | `2,842,003 B (2.71 MiB)` | `3,530,099 B (3.37 MiB)` | `24.211656%` | `0` | `57 / 60` |
| `CLID` | `INT(32,true)` | `116` | `109` (`93.97%`) | `7` (`6.03%`) | `0` (`0.00%`) | `951,412` (`95.14%`) | `48,588` (`4.86%`) | `5,641 B (5.51 KiB)` | `7,234 B (7.06 KiB)` | `28.239674%` | `0` | `57 / 60` |
| `ResolutionHeight` | `INT(16,true)` | `116` | `111` (`95.69%`) | `5` (`4.31%`) | `0` (`0.00%`) | `960,461` (`96.05%`) | `39,539` (`3.95%`) | `186,024 B (181.66 KiB)` | `208,351 B (203.47 KiB)` | `12.002215%` | `0` | `57 / 60` |
| `ResponseEndTiming` | `INT(32,true)` | `116` | `112` (`96.55%`) | `4` (`3.45%`) | `0` (`0.00%`) | `992,596` (`99.26%`) | `7,404` (`0.74%`) | `937,782 B (915.80 KiB)` | `1,056,836 B (1.01 MiB)` | `12.695275%` | `0` | `57 / 60` |
| `SendTiming` | `INT(32,true)` | `116` | `112` (`96.55%`) | `4` (`3.45%`) | `0` (`0.00%`) | `962,463` (`96.25%`) | `37,537` (`3.75%`) | `61,444 B (60.00 KiB)` | `80,125 B (78.25 KiB)` | `30.403294%` | `0` | `57 / 60` |
| `SocialSourceNetworkID` | `INT(16,true)` | `116` | `112` (`96.55%`) | `4` (`3.45%`) | `0` (`0.00%`) | `995,467` (`99.55%`) | `4,533` (`0.45%`) | `5,318 B (5.19 KiB)` | `7,223 B (7.05 KiB)` | `35.821737%` | `0` | `57 / 60` |
| `FetchTiming` | `INT(32,true)` | `116` | `114` (`98.28%`) | `2` (`1.72%`) | `0` (`0.00%`) | `992,406` (`99.24%`) | `7,594` (`0.76%`) | `549,819 B (536.93 KiB)` | `695,763 B (679.46 KiB)` | `26.544008%` | `0` | `57 / 60` |
| `RemoteIP` | `INT(32,true)` | `116` | `114` (`98.28%`) | `2` (`1.72%`) | `0` (`0.00%`) | `988,642` (`98.86%`) | `11,358` (`1.14%`) | `426,607 B (416.61 KiB)` | `701,005 B (684.58 KiB)` | `64.321026%` | `0` | `57 / 60` |
| `CounterClass` | `INT(16,true)` | `116` | `115` (`99.14%`) | `1` (`0.86%`) | `0` (`0.00%`) | `996,017` (`99.60%`) | `3,983` (`0.40%`) | `4,240 B (4.14 KiB)` | `6,292 B (6.14 KiB)` | `48.396226%` | `0` | `57 / 60` |
| `CounterID` | `INT(32,true)` | `116` | `115` (`99.14%`) | `1` (`0.86%`) | `0` (`0.00%`) | `996,017` (`99.60%`) | `3,983` (`0.40%`) | `4,941 B (4.83 KiB)` | `6,343 B (6.19 KiB)` | `28.374823%` | `0` | `57 / 60` |
| `EventDate` | `DATE` | `116` | `115` (`99.14%`) | `1` (`0.86%`) | `0` (`0.00%`) | `996,017` (`99.60%`) | `3,983` (`0.40%`) | `4,923 B (4.81 KiB)` | `6,292 B (6.14 KiB)` | `27.808247%` | `0` | `57 / 60` |
| `GoodEvent` | `INT(16,true)` | `116` | `115` (`99.14%`) | `1` (`0.86%`) | `0` (`0.00%`) | `996,017` (`99.60%`) | `3,983` (`0.40%`) | `4,925 B (4.81 KiB)` | `6,292 B (6.14 KiB)` | `27.756345%` | `0` | `57 / 60` |
| `HTTPError` | `INT(16,true)` | `116` | `115` (`99.14%`) | `1` (`0.86%`) | `0` (`0.00%`) | `996,017` (`99.60%`) | `3,983` (`0.40%`) | `4,240 B (4.14 KiB)` | `6,292 B (6.14 KiB)` | `48.396226%` | `0` | `57 / 60` |
| `IsEvent` | `INT(16,true)` | `116` | `115` (`99.14%`) | `1` (`0.86%`) | `0` (`0.00%`) | `996,017` (`99.60%`) | `3,983` (`0.40%`) | `4,240 B (4.14 KiB)` | `6,292 B (6.14 KiB)` | `48.396226%` | `0` | `57 / 60` |
| `IsOldCounter` | `INT(16,true)` | `116` | `115` (`99.14%`) | `1` (`0.86%`) | `0` (`0.00%`) | `996,017` (`99.60%`) | `3,983` (`0.40%`) | `4,240 B (4.14 KiB)` | `6,292 B (6.14 KiB)` | `48.396226%` | `0` | `57 / 60` |
| `IsParameter` | `INT(16,true)` | `116` | `115` (`99.14%`) | `1` (`0.86%`) | `0` (`0.00%`) | `996,017` (`99.60%`) | `3,983` (`0.40%`) | `4,240 B (4.14 KiB)` | `6,292 B (6.14 KiB)` | `48.396226%` | `0` | `57 / 60` |
| `OpenerName` | `INT(32,true)` | `116` | `115` (`99.14%`) | `1` (`0.86%`) | `0` (`0.00%`) | `996,017` (`99.60%`) | `3,983` (`0.40%`) | `4,239 B (4.14 KiB)` | `6,292 B (6.14 KiB)` | `48.431234%` | `0` | `57 / 60` |
| `ParamCurrency` | `STRING` | `116` | `115` (`99.14%`) | `1` (`0.86%`) | `0` (`0.00%`) | `996,017` (`99.60%`) | `3,983` (`0.40%`) | `5,261 B (5.14 KiB)` | `6,232 B (6.09 KiB)` | `18.456567%` | `0` | `57 / 60` |
| `ParamCurrencyID` | `INT(16,true)` | `116` | `115` (`99.14%`) | `1` (`0.86%`) | `0` (`0.00%`) | `996,017` (`99.60%`) | `3,983` (`0.40%`) | `4,240 B (4.14 KiB)` | `6,292 B (6.14 KiB)` | `48.396226%` | `0` | `57 / 60` |
| `ParamPrice` | `INT(64,true)` | `116` | `115` (`99.14%`) | `1` (`0.86%`) | `0` (`0.00%`) | `996,017` (`99.60%`) | `3,983` (`0.40%`) | `5,725 B (5.59 KiB)` | `7,492 B (7.32 KiB)` | `30.864629%` | `0` | `57 / 60` |
| `RegionID` | `INT(32,true)` | `116` | `115` (`99.14%`) | `1` (`0.86%`) | `0` (`0.00%`) | `991,605` (`99.16%`) | `8,395` (`0.84%`) | `191,174 B (186.69 KiB)` | `248,192 B (242.38 KiB)` | `29.825185%` | `0` | `57 / 60` |
| `ResponseStartTiming` | `INT(32,true)` | `116` | `115` (`99.14%`) | `1` (`0.86%`) | `0` (`0.00%`) | `999,832` (`99.98%`) | `168` (`0.02%`) | `1,245,891 B (1.19 MiB)` | `1,587,891 B (1.51 MiB)` | `27.450234%` | `0` | `57 / 60` |
| `Robotness` | `INT(16,true)` | `116` | `115` (`99.14%`) | `1` (`0.86%`) | `0` (`0.00%`) | `998,589` (`99.86%`) | `1,411` (`0.14%`) | `173,697 B (169.63 KiB)` | `243,015 B (237.32 KiB)` | `39.907425%` | `0` | `57 / 60` |
| `SilverlightVersion4` | `INT(16,true)` | `116` | `115` (`99.14%`) | `1` (`0.86%`) | `0` (`0.00%`) | `996,017` (`99.60%`) | `3,983` (`0.40%`) | `4,368 B (4.27 KiB)` | `6,401 B (6.25 KiB)` | `46.543040%` | `0` | `57 / 60` |
| `WithHash` | `INT(16,true)` | `116` | `115` (`99.14%`) | `1` (`0.86%`) | `0` (`0.00%`) | `996,017` (`99.60%`) | `3,983` (`0.40%`) | `4,240 B (4.14 KiB)` | `6,292 B (6.14 KiB)` | `48.396226%` | `0` | `57 / 60` |
| `ClientEventTime` | `TIMESTAMP(isAdjustedToUTC=true,unit=MILLIS)` | `116` | `116` (`100.00%`) | `0` (`0.00%`) | `0` (`0.00%`) | `1,000,000` (`100.00%`) | `0` (`0.00%`) | `2,475,752 B (2.36 MiB)` | `3,958,950 B (3.78 MiB)` | `59.908989%` | `0` | `57 / 60` |
| `ClientIP` | `INT(32,true)` | `116` | `116` (`100.00%`) | `0` (`0.00%`) | `0` (`0.00%`) | `1,000,000` (`100.00%`) | `0` (`0.00%`) | `408,069 B (398.50 KiB)` | `813,095 B (794.04 KiB)` | `99.254293%` | `0` | `57 / 60` |
| `EventTime` | `TIMESTAMP(isAdjustedToUTC=true,unit=MILLIS)` | `116` | `116` (`100.00%`) | `0` (`0.00%`) | `0` (`0.00%`) | `1,000,000` (`100.00%`) | `0` (`0.00%`) | `2,517,896 B (2.40 MiB)` | `4,021,665 B (3.84 MiB)` | `59.723237%` | `0` | `57 / 60` |
| `FUniqID` | `INT(64,true)` | `116` | `116` (`100.00%`) | `0` (`0.00%`) | `0` (`0.00%`) | `1,000,000` (`100.00%`) | `0` (`0.00%`) | `694,033 B (677.77 KiB)` | `1,065,272 B (1.02 MiB)` | `53.490108%` | `0` | `57 / 60` |
| `IPNetworkID` | `INT(32,true)` | `116` | `116` (`100.00%`) | `0` (`0.00%`) | `0` (`0.00%`) | `1,000,000` (`100.00%`) | `0` (`0.00%`) | `323,783 B (316.19 KiB)` | `623,466 B (608.85 KiB)` | `92.556743%` | `0` | `57 / 60` |
| `Interests` | `INT(16,true)` | `116` | `116` (`100.00%`) | `0` (`0.00%`) | `0` (`0.00%`) | `1,000,000` (`100.00%`) | `0` (`0.00%`) | `193,520 B (188.98 KiB)` | `310,229 B (302.96 KiB)` | `60.308495%` | `0` | `57 / 60` |
| `LocalEventTime` | `TIMESTAMP(isAdjustedToUTC=true,unit=MILLIS)` | `116` | `116` (`100.00%`) | `0` (`0.00%`) | `0` (`0.00%`) | `1,000,000` (`100.00%`) | `0` (`0.00%`) | `2,520,192 B (2.40 MiB)` | `4,023,641 B (3.84 MiB)` | `59.656129%` | `0` | `57 / 60` |
| `ParamOrderID` | `STRING` | `116` | `116` (`100.00%`) | `0` (`0.00%`) | `0` (`0.00%`) | `1,000,000` (`100.00%`) | `0` (`0.00%`) | `2,872 B (2.80 KiB)` | `5,332 B (5.21 KiB)` | `85.654596%` | `0` | `57 / 60` |
| `Params` | `STRING` | `116` | `116` (`100.00%`) | `0` (`0.00%`) | `0` (`0.00%`) | `1,000,000` (`100.00%`) | `0` (`0.00%`) | `2,872 B (2.80 KiB)` | `5,332 B (5.21 KiB)` | `85.654596%` | `0` | `57 / 60` |
| `SocialAction` | `STRING` | `116` | `116` (`100.00%`) | `0` (`0.00%`) | `0` (`0.00%`) | `1,000,000` (`100.00%`) | `0` (`0.00%`) | `2,872 B (2.80 KiB)` | `5,332 B (5.21 KiB)` | `85.654596%` | `0` | `57 / 60` |
| `SocialNetwork` | `STRING` | `116` | `116` (`100.00%`) | `0` (`0.00%`) | `0` (`0.00%`) | `1,000,000` (`100.00%`) | `0` (`0.00%`) | `2,872 B (2.80 KiB)` | `5,332 B (5.21 KiB)` | `85.654596%` | `0` | `57 / 60` |
| `URLHash` | `INT(64,true)` | `116` | `116` (`100.00%`) | `0` (`0.00%`) | `0` (`0.00%`) | `1,000,000` (`100.00%`) | `0` (`0.00%`) | `3,580,480 B (3.41 MiB)` | `4,558,127 B (4.35 MiB)` | `27.304914%` | `0` | `57 / 60` |
| `UserID` | `INT(64,true)` | `116` | `116` (`100.00%`) | `0` (`0.00%`) | `0` (`0.00%`) | `1,000,000` (`100.00%`) | `0` (`0.00%`) | `618,298 B (603.81 KiB)` | `1,106,815 B (1.06 MiB)` | `79.009960%` | `0` | `57 / 60` |
| `WatchID` | `INT(64,true)` | `116` | `116` (`100.00%`) | `0` (`0.00%`) | `0` (`0.00%`) | `1,000,000` (`100.00%`) | `0` (`0.00%`) | `8,005,365 B (7.63 MiB)` | `9,852,495 B (9.40 MiB)` | `23.073651%` | `0` | `57 / 60` |
| `WindowClientHeight` | `INT(16,true)` | `116` | `116` (`100.00%`) | `0` (`0.00%`) | `0` (`0.00%`) | `1,000,000` (`100.00%`) | `0` (`0.00%`) | `319,807 B (312.31 KiB)` | `591,986 B (578.11 KiB)` | `85.107268%` | `0` | `57 / 60` |
| `WindowClientWidth` | `INT(16,true)` | `116` | `116` (`100.00%`) | `0` (`0.00%`) | `0` (`0.00%`) | `1,000,000` (`100.00%`) | `0` (`0.00%`) | `305,751 B (298.58 KiB)` | `475,942 B (464.79 KiB)` | `55.663268%` | `0` | `57 / 60` |

## ZSTD RLE Dict Worse Distribution By Category

For columns where the best observed `zstd + plain` compressed byte count is smaller than the best observed `zstd + rle-dict` compressed byte count, each category image plots `plain + zstd` compressed bytes on the x-axis and `rle-dict + zstd` compressed bytes on the y-axis using the same log byte scale. Points above the diagonal are larger with RLE dictionary encoding. Point color is bucketed by `plain/no-compression encoded bytes / rle-dict + zstd compressed bytes`, so high-ratio colors identify columns where RLE dictionary lost the head-to-head but still compressed the baseline dramatically.

The bucket tables below each image still show how much worse RLE dictionary encoding was. Worse-by percentage is `(rle_dict_compressed_bytes / plain_compressed_bytes - 1) * 100`, so values can exceed 100%.

The compressed bytes are Parquet column-chunk bytes, including dictionary pages and page headers.

`Plain encoded bytes before compression` is the same column's byte count from the all-plain/no-compression baseline run. The `/ plain encoded` percentage columns compare compressed column bytes against that baseline denominator.

Categorization uses only measured byte sizes, row-group cardinality, and column type: `True dictionary bloat` means RLE dictionary encoded bytes exceeded plain encoded bytes before ZSTD; `Tiny/constant plain stream` means median row-group cardinality is at most 2 or median cardinality/rows is at most 0.0006; `Structured medium/high-cardinality numeric streams` means a numeric or temporal column has median cardinality/rows at least 0.09; the remaining losing columns fall into `Small-domain fixed-width literals`. Sortedness, page min/max, and value-length distributions are shown elsewhere in this report but are not currently used for this category assignment.

- Compared columns: `105`
- `zstd + rle-dict` worse than `zstd + plain`: `56`; better: `49`; ties: `0`; missing comparisons: `0`
- Missing shape stats while categorizing: `0`

| Category | Columns | Worse by min/median/max |
| --- | ---: | ---: |
| True dictionary bloat | 2 | 18.468449% / 20.666456% / 22.864463% |
| Tiny/constant plain stream | 30 | 1.738241% / 34.805520% / 110.023587% |
| Small-domain fixed-width literals | 18 | 8.762892% / 33.500720% / 97.747379% |
| Structured medium/high-cardinality numeric streams | 6 | 23.259941% / 43.172756% / 59.989014% |

### True dictionary bloat

RLE dictionary encoding was already larger than plain before ZSTD, usually because the dictionary itself was too large for the column.

![RLE dictionary worse: True dictionary bloat](images/zstd_rle_dict_worse_true_dictionary_bloat.svg)

| Improvement bucket | `zstd + rle-dict` worse by |
| --- | ---: |
| `0-10%` | 0 |
| `10-20%` | 1 |
| `20-30%` | 1 |
| `30-40%` | 0 |
| `40-50%` | 0 |
| `50-60%` | 0 |
| `60-70%` | 0 |
| `70-80%` | 0 |
| `80-90%` | 0 |
| `90-100%` | 0 |
| `100-200%` | 0 |
| `200-500%` | 0 |
| `500%+` | 0 |

| Column | Type | Category | Measured feature | Measured reason | Row-group cardinality min/median/max | Median cardinality / rows | Min value length (B) | Median value length (B) | Max value length (B) | Physical bytes before encoding/compression | Plain encoded bytes before compression | Plain + ZSTD compressed bytes | Plain + ZSTD / physical | Plain + ZSTD / plain encoded | RLE dict + ZSTD compressed bytes | RLE dict + ZSTD / physical | RLE dict + ZSTD / plain encoded | RLE dict + ZSTD vs plain + ZSTD | RLE dict + ZSTD without dict pages | RLE dict without dict pages / physical | RLE dict without dict pages / plain encoded | RLE dict without dict pages vs plain + ZSTD | RLE + dict is better without including dict page | Dictionary pages |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | --- | ---: |
| `WatchID` | `int64` | True dictionary bloat | plain encoded 8,004,555 B (7.63 MiB); rle encoded 9,834,361 B (9.38 MiB) | RLE dictionary was larger than plain before ZSTD; the compressed result stayed larger. | 9,315 / 11,938 / 14,202 | 100.000000% | 8 | 8 | 8 | 8,000,000 B (7.63 MiB) | 8,006,312 B (7.64 MiB) | 8,005,353 B (7.63 MiB) | 100.066913% | 99.988022% | 9,835,734 B (9.38 MiB) | 122.946675% | 122.849747% | 22.864463% | 1,834,419 B (1.75 MiB) | 22.930238% | 22.912160% | -77.085095% | yes | 58 |
| `HID` | `int32` | True dictionary bloat | plain encoded 4,003,775 B (3.82 MiB); rle encoded 4,491,293 B (4.28 MiB) | RLE dictionary was larger than plain before ZSTD; the compressed result stayed larger. | 5,818 / 5,965 / 13,281 | 49.966494% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,051 B (3.82 MiB) | 3,792,143 B (3.62 MiB) | 94.803575% | 94.684013% | 4,492,493 B (4.28 MiB) | 112.312325% | 112.170681% | 18.468449% | 1,752,569 B (1.67 MiB) | 43.814225% | 43.758968% | -53.784206% | yes | 60 |

### Tiny/constant plain stream

The column is tiny or nearly constant per row group; plain pages give ZSTD an extremely repetitive stream, while dictionary pages add overhead.

![RLE dictionary worse: Tiny/constant plain stream](images/zstd_rle_dict_worse_tiny_constant_plain_stream.svg)

| Improvement bucket | `zstd + rle-dict` worse by |
| --- | ---: |
| `0-10%` | 3 |
| `10-20%` | 5 |
| `20-30%` | 7 |
| `30-40%` | 1 |
| `40-50%` | 9 |
| `50-60%` | 0 |
| `60-70%` | 0 |
| `70-80%` | 0 |
| `80-90%` | 4 |
| `90-100%` | 0 |
| `100-200%` | 1 |
| `200-500%` | 0 |
| `500%+` | 0 |

| Column | Type | Category | Measured feature | Measured reason | Row-group cardinality min/median/max | Median cardinality / rows | Min value length (B) | Median value length (B) | Max value length (B) | Physical bytes before encoding/compression | Plain encoded bytes before compression | Plain + ZSTD compressed bytes | Plain + ZSTD / physical | Plain + ZSTD / plain encoded | RLE dict + ZSTD compressed bytes | RLE dict + ZSTD / physical | RLE dict + ZSTD / plain encoded | RLE dict + ZSTD vs plain + ZSTD | RLE dict + ZSTD without dict pages | RLE dict without dict pages / physical | RLE dict without dict pages / plain encoded | RLE dict without dict pages vs plain + ZSTD | RLE + dict is better without including dict page | Dictionary pages |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | --- | ---: |
| `WindowName` | `int32` | Tiny/constant plain stream | median row-group cardinality 6; median cardinality/rows 0.050260% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 6 / 3,150 | 0.050260% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,052 B (3.82 MiB) | 70,803 B (69.14 KiB) | 1.770075% | 1.767842% | 148,703 B (145.22 KiB) | 3.717575% | 3.712886% | 110.023587% | 73,928 B (72.20 KiB) | 1.848200% | 1.845869% | 4.413655% | no | 59 |
| `ParamOrderID` | `string` | Tiny/constant plain stream | median row-group cardinality 1; median cardinality/rows 0.008377% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 1 / 1 | 0.008377% | 0 | 0 | 0 | 0 B (0 B) | 4,003,158 B (3.82 MiB) | 2,848 B (2.78 KiB) | n/a | 0.071144% | 5,247 B (5.12 KiB) | n/a | 0.131072% | 84.234551% | 3,949 B (3.86 KiB) | n/a | 0.098647% | 38.658708% | no | 59 |
| `Params` | `string` | Tiny/constant plain stream | median row-group cardinality 1; median cardinality/rows 0.008377% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 1 / 1 | 0.008377% | 0 | 0 | 0 | 0 B (0 B) | 4,003,158 B (3.82 MiB) | 2,848 B (2.78 KiB) | n/a | 0.071144% | 5,247 B (5.12 KiB) | n/a | 0.131072% | 84.234551% | 3,949 B (3.86 KiB) | n/a | 0.098647% | 38.658708% | no | 59 |
| `SocialAction` | `string` | Tiny/constant plain stream | median row-group cardinality 1; median cardinality/rows 0.008377% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 1 / 1 | 0.008377% | 0 | 0 | 0 | 0 B (0 B) | 4,003,158 B (3.82 MiB) | 2,848 B (2.78 KiB) | n/a | 0.071144% | 5,247 B (5.12 KiB) | n/a | 0.131072% | 84.234551% | 3,949 B (3.86 KiB) | n/a | 0.098647% | 38.658708% | no | 59 |
| `SocialNetwork` | `string` | Tiny/constant plain stream | median row-group cardinality 1; median cardinality/rows 0.008377% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 1 / 1 | 0.008377% | 0 | 0 | 0 | 0 B (0 B) | 4,003,158 B (3.82 MiB) | 2,848 B (2.78 KiB) | n/a | 0.071144% | 5,247 B (5.12 KiB) | n/a | 0.131072% | 84.234551% | 3,949 B (3.86 KiB) | n/a | 0.098647% | 38.658708% | no | 59 |
| `CounterClass` | `int16` | Tiny/constant plain stream | median row-group cardinality 1; median cardinality/rows 0.008377% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 1 / 1 | 0.008377% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,054 B (3.82 MiB) | 4,228 B (4.13 KiB) | 0.105700% | 0.105567% | 5,975 B (5.83 KiB) | 0.149375% | 0.149187% | 41.319773% | 4,721 B (4.61 KiB) | 0.118025% | 0.117876% | 11.660360% | no | 57 |
| `HTTPError` | `int16` | Tiny/constant plain stream | median row-group cardinality 1; median cardinality/rows 0.008377% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 1 / 1 | 0.008377% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,054 B (3.82 MiB) | 4,228 B (4.13 KiB) | 0.105700% | 0.105567% | 5,975 B (5.83 KiB) | 0.149375% | 0.149187% | 41.319773% | 4,721 B (4.61 KiB) | 0.118025% | 0.117876% | 11.660360% | no | 57 |
| `IsEvent` | `int16` | Tiny/constant plain stream | median row-group cardinality 1; median cardinality/rows 0.008377% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 1 / 1 | 0.008377% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,054 B (3.82 MiB) | 4,228 B (4.13 KiB) | 0.105700% | 0.105567% | 5,975 B (5.83 KiB) | 0.149375% | 0.149187% | 41.319773% | 4,721 B (4.61 KiB) | 0.118025% | 0.117876% | 11.660360% | no | 57 |
| `IsOldCounter` | `int16` | Tiny/constant plain stream | median row-group cardinality 1; median cardinality/rows 0.008377% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 1 / 1 | 0.008377% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,054 B (3.82 MiB) | 4,228 B (4.13 KiB) | 0.105700% | 0.105567% | 5,975 B (5.83 KiB) | 0.149375% | 0.149187% | 41.319773% | 4,721 B (4.61 KiB) | 0.118025% | 0.117876% | 11.660360% | no | 57 |
| `IsParameter` | `int16` | Tiny/constant plain stream | median row-group cardinality 1; median cardinality/rows 0.008377% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 1 / 1 | 0.008377% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,054 B (3.82 MiB) | 4,228 B (4.13 KiB) | 0.105700% | 0.105567% | 5,975 B (5.83 KiB) | 0.149375% | 0.149187% | 41.319773% | 4,721 B (4.61 KiB) | 0.118025% | 0.117876% | 11.660360% | no | 57 |
| `ParamCurrencyID` | `int16` | Tiny/constant plain stream | median row-group cardinality 1; median cardinality/rows 0.008377% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 1 / 1 | 0.008377% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,054 B (3.82 MiB) | 4,228 B (4.13 KiB) | 0.105700% | 0.105567% | 5,975 B (5.83 KiB) | 0.149375% | 0.149187% | 41.319773% | 4,721 B (4.61 KiB) | 0.118025% | 0.117876% | 11.660360% | no | 57 |
| `WithHash` | `int16` | Tiny/constant plain stream | median row-group cardinality 1; median cardinality/rows 0.008377% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 1 / 1 | 0.008377% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,054 B (3.82 MiB) | 4,228 B (4.13 KiB) | 0.105700% | 0.105567% | 5,975 B (5.83 KiB) | 0.149375% | 0.149187% | 41.319773% | 4,721 B (4.61 KiB) | 0.118025% | 0.117876% | 11.660360% | no | 57 |
| `OpenerName` | `int32` | Tiny/constant plain stream | median row-group cardinality 1; median cardinality/rows 0.008377% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 1 / 1 | 0.008377% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,046 B (3.82 MiB) | 4,232 B (4.13 KiB) | 0.105800% | 0.105667% | 5,975 B (5.83 KiB) | 0.149375% | 0.149187% | 41.186200% | 4,721 B (4.61 KiB) | 0.118025% | 0.117876% | 11.554820% | no | 57 |
| `SocialSourcePage` | `string` | Tiny/constant plain stream | median row-group cardinality 1; median cardinality/rows 0.008377% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 1 / 5 | 0.008377% | 0 | 0 | 28 | 1,024 B (1.00 KiB) | 4,005,153 B (3.82 MiB) | 4,853 B (4.74 KiB) | 473.925781% | 0.121169% | 6,826 B (6.67 KiB) | 666.601562% | 0.170430% | 40.655265% | 4,890 B (4.78 KiB) | 477.539062% | 0.122093% | 0.762415% | no | 59 |
| `SilverlightVersion4` | `int16` | Tiny/constant plain stream | median row-group cardinality 1; median cardinality/rows 0.008377% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 1 / 3 | 0.008377% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,054 B (3.82 MiB) | 4,357 B (4.25 KiB) | 0.108925% | 0.108788% | 6,084 B (5.94 KiB) | 0.152100% | 0.151908% | 39.637365% | 4,794 B (4.68 KiB) | 0.119850% | 0.119699% | 10.029837% | no | 57 |
| `SocialSourceNetworkID` | `int16` | Tiny/constant plain stream | median row-group cardinality 2; median cardinality/rows 0.016753% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 2 / 4 | 0.016753% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,053 B (3.82 MiB) | 5,318 B (5.19 KiB) | 0.132950% | 0.132782% | 6,912 B (6.75 KiB) | 0.172800% | 0.172582% | 29.973674% | 5,390 B (5.26 KiB) | 0.134750% | 0.134580% | 1.353892% | no | 57 |
| `ParamPrice` | `int64` | Tiny/constant plain stream | median row-group cardinality 1; median cardinality/rows 0.008377% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 1 / 1 | 0.008377% | 8 | 8 | 8 | 8,000,000 B (7.63 MiB) | 8,006,318 B (7.64 MiB) | 5,675 B (5.54 KiB) | 0.070938% | 0.070882% | 7,115 B (6.95 KiB) | 0.088938% | 0.088867% | 25.374449% | 5,633 B (5.50 KiB) | 0.070413% | 0.070357% | -0.740088% | yes | 57 |
| `SendTiming` | `int32` | Tiny/constant plain stream | median row-group cardinality 1; median cardinality/rows 0.008377% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 1 / 989 | 0.008377% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,054 B (3.82 MiB) | 61,415 B (59.98 KiB) | 1.535375% | 1.533438% | 76,243 B (74.46 KiB) | 1.906075% | 1.903670% | 24.143939% | 53,896 B (52.63 KiB) | 1.347400% | 1.345700% | -12.242937% | yes | 59 |
| `CLID` | `int32` | Tiny/constant plain stream | median row-group cardinality 1; median cardinality/rows 0.008377% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 1 / 2 | 0.008377% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,054 B (3.82 MiB) | 5,627 B (5.50 KiB) | 0.140675% | 0.140497% | 6,908 B (6.75 KiB) | 0.172700% | 0.172482% | 22.765239% | 5,554 B (5.42 KiB) | 0.138850% | 0.138675% | -1.297317% | yes | 57 |
| `CounterID` | `int32` | Tiny/constant plain stream | median row-group cardinality 1; median cardinality/rows 0.008377% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 1 / 4 | 0.008377% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,052 B (3.82 MiB) | 4,931 B (4.82 KiB) | 0.123275% | 0.123120% | 6,020 B (5.88 KiB) | 0.150500% | 0.150310% | 22.084770% | 4,742 B (4.63 KiB) | 0.118550% | 0.118400% | -3.832894% | yes | 57 |
| `EventDate` | `date` | Tiny/constant plain stream | median row-group cardinality 1; median cardinality/rows 0.008377% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 1 / 1 | 0.008377% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,045 B (3.82 MiB) | 4,905 B (4.79 KiB) | 0.122625% | 0.122471% | 5,976 B (5.84 KiB) | 0.149400% | 0.149212% | 21.834862% | 4,722 B (4.61 KiB) | 0.118050% | 0.117901% | -3.730887% | yes | 57 |
| `GoodEvent` | `int16` | Tiny/constant plain stream | median row-group cardinality 1; median cardinality/rows 0.008377% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 1 / 1 | 0.008377% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,053 B (3.82 MiB) | 4,915 B (4.80 KiB) | 0.122875% | 0.122720% | 5,975 B (5.83 KiB) | 0.149375% | 0.149187% | 21.566633% | 4,721 B (4.61 KiB) | 0.118025% | 0.117876% | -3.947101% | yes | 57 |
| `CookieEnable` | `int16` | Tiny/constant plain stream | median row-group cardinality 2; median cardinality/rows 0.016753% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 2 / 2 | 0.016753% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,052 B (3.82 MiB) | 6,053 B (5.91 KiB) | 0.151325% | 0.151134% | 7,077 B (6.91 KiB) | 0.176925% | 0.176702% | 16.917231% | 5,663 B (5.53 KiB) | 0.141575% | 0.141396% | -6.443086% | yes | 57 |
| `ParamCurrency` | `string` | Tiny/constant plain stream | median row-group cardinality 1; median cardinality/rows 0.008377% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 1 / 1 | 0.008377% | 3 | 3 | 3 | 3,000,000 B (2.86 MiB) | 7,004,733 B (6.68 MiB) | 5,256 B (5.13 KiB) | 0.175200% | 0.075035% | 6,132 B (5.99 KiB) | 0.204400% | 0.087541% | 16.666667% | 4,657 B (4.55 KiB) | 0.155233% | 0.066484% | -11.396499% | yes | 59 |
| `IsDownload` | `int16` | Tiny/constant plain stream | median row-group cardinality 1; median cardinality/rows 0.008377% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 1 / 2 | 0.008377% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,055 B (3.82 MiB) | 8,024 B (7.84 KiB) | 0.200600% | 0.200347% | 9,254 B (9.04 KiB) | 0.231350% | 0.231058% | 15.329013% | 7,900 B (7.71 KiB) | 0.197500% | 0.197251% | -1.545364% | yes | 57 |
| `JavascriptEnable` | `int16` | Tiny/constant plain stream | median row-group cardinality 2; median cardinality/rows 0.016753% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 2 / 2 | 0.016753% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,051 B (3.82 MiB) | 6,485 B (6.33 KiB) | 0.162125% | 0.161921% | 7,387 B (7.21 KiB) | 0.184675% | 0.184442% | 13.909021% | 5,941 B (5.80 KiB) | 0.148525% | 0.148338% | -8.388589% | yes | 57 |
| `CodeVersion` | `int32` | Tiny/constant plain stream | median row-group cardinality 2; median cardinality/rows 0.016753% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 2 / 3 | 0.016753% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,053 B (3.82 MiB) | 7,031 B (6.87 KiB) | 0.175775% | 0.175553% | 7,970 B (7.78 KiB) | 0.199250% | 0.198999% | 13.355142% | 6,424 B (6.27 KiB) | 0.160600% | 0.160397% | -8.633196% | yes | 57 |
| `UTMContent` | `string` | Tiny/constant plain stream | median row-group cardinality 3; median cardinality/rows 0.025130% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 3 / 25 | 0.025130% | 0 | 0 | 62 | 13,001 B (12.70 KiB) | 4,018,131 B (3.83 MiB) | 13,959 B (13.63 KiB) | 107.368664% | 0.347400% | 14,839 B (14.49 KiB) | 114.137374% | 0.369301% | 6.304177% | 9,904 B (9.67 KiB) | 76.178755% | 0.246483% | -29.049359% | yes | 59 |
| `MobilePhone` | `int16` | Tiny/constant plain stream | median row-group cardinality 7; median cardinality/rows 0.058636% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 3 / 7 / 11 | 0.058636% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,050 B (3.82 MiB) | 22,463 B (21.94 KiB) | 0.561575% | 0.560867% | 23,572 B (23.02 KiB) | 0.589300% | 0.588557% | 4.937008% | 20,726 B (20.24 KiB) | 0.518150% | 0.517497% | -7.732716% | yes | 57 |
| `UTMTerm` | `string` | Tiny/constant plain stream | median row-group cardinality 2; median cardinality/rows 0.016753% | Plain+ZSTD collapsed a constant or near-constant stream more than RLE-dict's dictionary/page/ID overhead. | 1 / 2 / 75 | 0.016753% | 0 | 0 | 72 | 28,101 B (27.44 KiB) | 4,034,484 B (3.85 MiB) | 15,648 B (15.28 KiB) | 55.684851% | 0.387856% | 15,920 B (15.55 KiB) | 56.652788% | 0.394598% | 1.738241% | 10,518 B (10.27 KiB) | 37.429273% | 0.260702% | -32.783742% | yes | 59 |

### Small-domain fixed-width literals

RLE dictionary shrank the encoded stream, but ZSTD compressed the repeated fixed-width plain literals better than dictionary IDs plus a dictionary page.

![RLE dictionary worse: Small-domain fixed-width literals](images/zstd_rle_dict_worse_small_domain_fixed_width_literals.svg)

| Improvement bucket | `zstd + rle-dict` worse by |
| --- | ---: |
| `0-10%` | 1 |
| `10-20%` | 6 |
| `20-30%` | 2 |
| `30-40%` | 1 |
| `40-50%` | 0 |
| `50-60%` | 2 |
| `60-70%` | 2 |
| `70-80%` | 1 |
| `80-90%` | 1 |
| `90-100%` | 2 |
| `100-200%` | 0 |
| `200-500%` | 0 |
| `500%+` | 0 |

| Column | Type | Category | Measured feature | Measured reason | Row-group cardinality min/median/max | Median cardinality / rows | Min value length (B) | Median value length (B) | Max value length (B) | Physical bytes before encoding/compression | Plain encoded bytes before compression | Plain + ZSTD compressed bytes | Plain + ZSTD / physical | Plain + ZSTD / plain encoded | RLE dict + ZSTD compressed bytes | RLE dict + ZSTD / physical | RLE dict + ZSTD / plain encoded | RLE dict + ZSTD vs plain + ZSTD | RLE dict + ZSTD without dict pages | RLE dict without dict pages / physical | RLE dict without dict pages / plain encoded | RLE dict without dict pages vs plain + ZSTD | RLE + dict is better without including dict page | Dictionary pages |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | --- | ---: |
| `ClientIP` | `int32` | Small-domain fixed-width literals | median row-group cardinality 924; median cardinality/rows 7.739990%; plain encoded 4,003,592 B (3.82 MiB); rle encoded 941,748 B (919.68 KiB) | RLE dictionary reduced pre-codec bytes, but ZSTD compressed the plain fixed-width values to fewer bytes. | 744 / 924 / 1,957 | 7.739990% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,050 B (3.82 MiB) | 408,058 B (398.49 KiB) | 10.201450% | 10.188587% | 806,924 B (788.01 KiB) | 20.173100% | 20.147664% | 97.747379% | 469,904 B (458.89 KiB) | 11.747600% | 11.732787% | 15.156179% | no | 57 |
| `IPNetworkID` | `int32` | Small-domain fixed-width literals | median row-group cardinality 600; median cardinality/rows 5.025967%; plain encoded 4,003,590 B (3.82 MiB); rle encoded 748,476 B (730.93 KiB) | RLE dictionary reduced pre-codec bytes, but ZSTD compressed the plain fixed-width values to fewer bytes. | 498 / 600 / 1,095 | 5.025967% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,052 B (3.82 MiB) | 323,708 B (316.12 KiB) | 8.092700% | 8.082492% | 619,756 B (605.23 KiB) | 15.493900% | 15.474356% | 91.455262% | 420,504 B (410.65 KiB) | 10.512600% | 10.499339% | 29.902258% | no | 57 |
| `WindowClientHeight` | `int16` | Small-domain fixed-width literals | median row-group cardinality 435; median cardinality/rows 3.643826%; plain encoded 4,003,650 B (3.82 MiB); rle encoded 750,345 B (732.76 KiB) | RLE dictionary reduced pre-codec bytes, but ZSTD compressed the plain fixed-width values to fewer bytes. | 318 / 435 / 575 | 3.643826% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,051 B (3.82 MiB) | 319,342 B (311.86 KiB) | 7.983550% | 7.973481% | 583,325 B (569.65 KiB) | 14.583125% | 14.564733% | 82.664667% | 459,637 B (448.86 KiB) | 11.490925% | 11.476433% | 43.932524% | no | 57 |
| `UserID` | `int64` | Small-domain fixed-width literals | median row-group cardinality 898; median cardinality/rows 7.522198%; plain encoded 8,004,550 B (7.63 MiB); rle encoded 1,230,153 B (1.17 MiB) | RLE dictionary reduced pre-codec bytes, but ZSTD compressed the plain fixed-width values to fewer bytes. | 716 / 898 / 1,805 | 7.522198% | 8 | 8 | 8 | 8,000,000 B (7.63 MiB) | 8,006,315 B (7.64 MiB) | 617,840 B (603.36 KiB) | 7.723000% | 7.716908% | 1,101,572 B (1.05 MiB) | 13.769650% | 13.758789% | 78.294057% | 455,193 B (444.52 KiB) | 5.689913% | 5.685425% | -26.325100% | yes | 57 |
| `RemoteIP` | `int32` | Small-domain fixed-width literals | median row-group cardinality 851; median cardinality/rows 7.128497%; plain encoded 4,003,606 B (3.82 MiB); rle encoded 927,182 B (905.45 KiB) | RLE dictionary reduced pre-codec bytes, but ZSTD compressed the plain fixed-width values to fewer bytes. | 508 / 851 / 1,951 | 7.128497% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,047 B (3.82 MiB) | 426,425 B (416.43 KiB) | 10.660625% | 10.647191% | 698,454 B (682.08 KiB) | 17.461350% | 17.439346% | 63.792930% | 407,905 B (398.34 KiB) | 10.197625% | 10.184774% | -4.343085% | yes | 59 |
| `Interests` | `int16` | Small-domain fixed-width literals | median row-group cardinality 217; median cardinality/rows 1.817725%; plain encoded 4,003,589 B (3.82 MiB); rle encoded 487,582 B (476.15 KiB) | RLE dictionary reduced pre-codec bytes, but ZSTD compressed the plain fixed-width values to fewer bytes. | 151 / 217 / 395 | 1.817725% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,050 B (3.82 MiB) | 193,369 B (188.84 KiB) | 4.834225% | 4.828129% | 310,021 B (302.75 KiB) | 7.750525% | 7.740752% | 60.326112% | 243,865 B (238.15 KiB) | 6.096625% | 6.088938% | 26.113803% | no | 60 |
| `WindowClientWidth` | `int16` | Small-domain fixed-width literals | median row-group cardinality 306; median cardinality/rows 2.563243%; plain encoded 4,003,590 B (3.82 MiB); rle encoded 695,366 B (679.07 KiB) | RLE dictionary reduced pre-codec bytes, but ZSTD compressed the plain fixed-width values to fewer bytes. | 241 / 306 / 374 | 2.563243% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,052 B (3.82 MiB) | 305,724 B (298.56 KiB) | 7.643100% | 7.633459% | 471,075 B (460.03 KiB) | 11.776875% | 11.762020% | 54.085057% | 384,571 B (375.56 KiB) | 9.614275% | 9.602147% | 25.790255% | no | 57 |
| `FUniqID` | `int64` | Small-domain fixed-width literals | median row-group cardinality 812; median cardinality/rows 6.801809%; plain encoded 8,004,554 B (7.63 MiB); rle encoded 1,186,134 B (1.13 MiB) | RLE dictionary reduced pre-codec bytes, but ZSTD compressed the plain fixed-width values to fewer bytes. | 628 / 812 / 1,703 | 6.801809% | 8 | 8 | 8 | 8,000,000 B (7.63 MiB) | 8,006,317 B (7.64 MiB) | 693,929 B (677.67 KiB) | 8.674112% | 8.667269% | 1,058,653 B (1.01 MiB) | 13.233163% | 13.222722% | 52.559268% | 451,570 B (440.99 KiB) | 5.644625% | 5.640171% | -34.925619% | yes | 57 |
| `Robotness` | `int16` | Small-domain fixed-width literals | median row-group cardinality 114; median cardinality/rows 0.954934%; plain encoded 4,003,775 B (3.82 MiB); rle encoded 415,545 B (405.81 KiB) | RLE dictionary reduced pre-codec bytes, but ZSTD compressed the plain fixed-width values to fewer bytes. | 85 / 114 / 200 | 0.954934% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,054 B (3.82 MiB) | 173,390 B (169.33 KiB) | 4.334750% | 4.329280% | 240,793 B (235.15 KiB) | 6.019825% | 6.012229% | 38.873637% | 207,281 B (202.42 KiB) | 5.182025% | 5.175486% | 19.546110% | no | 57 |
| `RegionID` | `int32` | Small-domain fixed-width literals | median row-group cardinality 149; median cardinality/rows 1.248115%; plain encoded 4,003,584 B (3.82 MiB); rle encoded 435,033 B (424.84 KiB) | RLE dictionary reduced pre-codec bytes, but ZSTD compressed the plain fixed-width values to fewer bytes. | 112 / 149 / 275 | 1.248115% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,048 B (3.82 MiB) | 190,701 B (186.23 KiB) | 4.767525% | 4.761516% | 244,341 B (238.61 KiB) | 6.108525% | 6.100826% | 28.127802% | 198,133 B (193.49 KiB) | 4.953325% | 4.947082% | 3.897200% | no | 57 |
| `FetchTiming` | `int32` | Small-domain fixed-width literals | median row-group cardinality 664; median cardinality/rows 5.562071%; plain encoded 4,003,625 B (3.82 MiB); rle encoded 1,131,789 B (1.08 MiB) | RLE dictionary reduced pre-codec bytes, but ZSTD compressed the plain fixed-width values to fewer bytes. | 329 / 664 / 1,264 | 5.562071% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,053 B (3.82 MiB) | 549,819 B (536.93 KiB) | 13.745475% | 13.728133% | 685,133 B (669.08 KiB) | 17.128325% | 17.106715% | 24.610645% | 475,578 B (464.43 KiB) | 11.889450% | 11.874450% | -13.502807% | yes | 57 |
| `DNSTiming` | `int32` | Small-domain fixed-width literals | median row-group cardinality 98; median cardinality/rows 0.820908%; plain encoded 4,003,588 B (3.82 MiB); rle encoded 343,417 B (335.37 KiB) | RLE dictionary reduced pre-codec bytes, but ZSTD compressed the plain fixed-width values to fewer bytes. | 68 / 98 / 352 | 0.820908% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,051 B (3.82 MiB) | 135,085 B (131.92 KiB) | 3.377125% | 3.372866% | 156,185 B (152.52 KiB) | 3.904625% | 3.899701% | 15.619795% | 121,048 B (118.21 KiB) | 3.026200% | 3.022383% | -10.391235% | yes | 57 |
| `UserAgentMajor` | `int16` | Small-domain fixed-width literals | median row-group cardinality 29; median cardinality/rows 0.242922%; plain encoded 4,003,585 B (3.82 MiB); rle encoded 275,570 B (269.11 KiB) | RLE dictionary reduced pre-codec bytes, but ZSTD compressed the plain fixed-width values to fewer bytes. | 26 / 29 / 31 | 0.242922% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,050 B (3.82 MiB) | 154,186 B (150.57 KiB) | 3.854650% | 3.849790% | 177,456 B (173.30 KiB) | 4.436400% | 4.430806% | 15.092161% | 169,722 B (165.74 KiB) | 4.243050% | 4.237700% | 10.076142% | no | 57 |
| `ConnectTiming` | `int32` | Small-domain fixed-width literals | median row-group cardinality 222; median cardinality/rows 1.859608%; plain encoded 4,003,660 B (3.82 MiB); rle encoded 703,627 B (687.14 KiB) | RLE dictionary reduced pre-codec bytes, but ZSTD compressed the plain fixed-width values to fewer bytes. | 117 / 222 / 628 | 1.859608% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,051 B (3.82 MiB) | 333,011 B (325.21 KiB) | 8.325275% | 8.314776% | 377,469 B (368.62 KiB) | 9.436725% | 9.424824% | 13.350310% | 291,018 B (284.20 KiB) | 7.275450% | 7.266275% | -12.610094% | yes | 57 |
| `OS` | `int16` | Small-domain fixed-width literals | median row-group cardinality 23; median cardinality/rows 0.192662%; plain encoded 4,003,587 B (3.82 MiB); rle encoded 229,346 B (223.97 KiB) | RLE dictionary reduced pre-codec bytes, but ZSTD compressed the plain fixed-width values to fewer bytes. | 16 / 23 / 31 | 0.192662% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,049 B (3.82 MiB) | 105,900 B (103.42 KiB) | 2.647500% | 2.644162% | 119,893 B (117.08 KiB) | 2.997325% | 2.993546% | 13.213409% | 113,271 B (110.62 KiB) | 2.831775% | 2.828205% | 6.960340% | no | 57 |
| `ResponseEndTiming` | `int32` | Small-domain fixed-width literals | median row-group cardinality 673; median cardinality/rows 5.637460%; plain encoded 4,003,640 B (3.82 MiB); rle encoded 1,358,021 B (1.30 MiB) | RLE dictionary reduced pre-codec bytes, but ZSTD compressed the plain fixed-width values to fewer bytes. | 528 / 673 / 1,577 | 5.637460% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,051 B (3.82 MiB) | 937,781 B (915.80 KiB) | 23.444525% | 23.414958% | 1,037,700 B (1013.38 KiB) | 25.942500% | 25.909782% | 10.654833% | 796,993 B (778.31 KiB) | 19.924825% | 19.899697% | -15.012887% | yes | 57 |
| `ResolutionHeight` | `int16` | Small-domain fixed-width literals | median row-group cardinality 70; median cardinality/rows 0.586363%; plain encoded 4,003,586 B (3.82 MiB); rle encoded 371,693 B (362.98 KiB) | RLE dictionary reduced pre-codec bytes, but ZSTD compressed the plain fixed-width values to fewer bytes. | 49 / 70 / 103 | 0.586363% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,050 B (3.82 MiB) | 186,022 B (181.66 KiB) | 4.650550% | 4.644686% | 205,409 B (200.59 KiB) | 5.135225% | 5.128750% | 10.421886% | 184,227 B (179.91 KiB) | 4.605675% | 4.599868% | -0.964940% | yes | 57 |
| `ResolutionWidth` | `int16` | Small-domain fixed-width literals | median row-group cardinality 64; median cardinality/rows 0.536103%; plain encoded 4,003,581 B (3.82 MiB); rle encoded 368,799 B (360.16 KiB) | RLE dictionary reduced pre-codec bytes, but ZSTD compressed the plain fixed-width values to fewer bytes. | 48 / 64 / 84 | 0.536103% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,053 B (3.82 MiB) | 187,130 B (182.74 KiB) | 4.678250% | 4.672348% | 203,528 B (198.76 KiB) | 5.088200% | 5.081780% | 8.762892% | 185,086 B (180.75 KiB) | 4.627150% | 4.621312% | -1.092289% | yes | 57 |

### Structured medium/high-cardinality numeric streams

The column has enough distinct numeric/timestamp values that the plain stream preserves structure ZSTD can exploit better than dictionary IDs.

![RLE dictionary worse: Structured medium/high-cardinality numeric streams](images/zstd_rle_dict_worse_structured_medium_high_cardinality_numeric_streams.svg)

| Improvement bucket | `zstd + rle-dict` worse by |
| --- | ---: |
| `0-10%` | 0 |
| `10-20%` | 0 |
| `20-30%` | 3 |
| `30-40%` | 0 |
| `40-50%` | 0 |
| `50-60%` | 3 |
| `60-70%` | 0 |
| `70-80%` | 0 |
| `80-90%` | 0 |
| `90-100%` | 0 |
| `100-200%` | 0 |
| `200-500%` | 0 |
| `500%+` | 0 |

| Column | Type | Category | Measured feature | Measured reason | Row-group cardinality min/median/max | Median cardinality / rows | Min value length (B) | Median value length (B) | Max value length (B) | Physical bytes before encoding/compression | Plain encoded bytes before compression | Plain + ZSTD compressed bytes | Plain + ZSTD / physical | Plain + ZSTD / plain encoded | RLE dict + ZSTD compressed bytes | RLE dict + ZSTD / physical | RLE dict + ZSTD / plain encoded | RLE dict + ZSTD vs plain + ZSTD | RLE dict + ZSTD without dict pages | RLE dict without dict pages / physical | RLE dict without dict pages / plain encoded | RLE dict without dict pages vs plain + ZSTD | RLE + dict is better without including dict page | Dictionary pages |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | --- | ---: |
| `ClientEventTime` | `timestamp_millis` | Structured medium/high-cardinality numeric streams | median row-group cardinality 5882; median cardinality/rows 49.271235%; plain+zstd 2,474,093 B (2.36 MiB); rle+zstd 3,958,277 B (3.77 MiB) | The high-cardinality numeric/timestamp column produced a larger RLE-dict+ZSTD stream than plain+ZSTD. | 5,666 / 5,882 / 13,078 | 49.271235% | 8 | 8 | 8 | 8,000,000 B (7.63 MiB) | 8,006,314 B (7.64 MiB) | 2,474,093 B (2.36 MiB) | 30.926162% | 30.901773% | 3,958,277 B (3.77 MiB) | 49.478462% | 49.439442% | 59.989014% | 1,725,958 B (1.65 MiB) | 21.574475% | 21.557461% | -30.238758% | yes | 60 |
| `EventTime` | `timestamp_millis` | Structured medium/high-cardinality numeric streams | median row-group cardinality 6237; median cardinality/rows 52.244932%; plain+zstd 2,514,539 B (2.40 MiB); rle+zstd 4,021,616 B (3.84 MiB) | The high-cardinality numeric/timestamp column produced a larger RLE-dict+ZSTD stream than plain+ZSTD. | 5,997 / 6,237 / 13,042 | 52.244932% | 8 | 8 | 8 | 8,000,000 B (7.63 MiB) | 8,006,315 B (7.64 MiB) | 2,514,539 B (2.40 MiB) | 31.431738% | 31.406946% | 4,021,616 B (3.84 MiB) | 50.270200% | 50.230549% | 59.934525% | 1,755,738 B (1.67 MiB) | 21.946725% | 21.929414% | -30.176545% | yes | 60 |
| `LocalEventTime` | `timestamp_millis` | Structured medium/high-cardinality numeric streams | median row-group cardinality 6254; median cardinality/rows 52.387335%; plain+zstd 2,517,265 B (2.40 MiB); rle+zstd 4,023,316 B (3.84 MiB) | The high-cardinality numeric/timestamp column produced a larger RLE-dict+ZSTD stream than plain+ZSTD. | 5,968 / 6,254 / 13,047 | 52.387335% | 8 | 8 | 8 | 8,000,000 B (7.63 MiB) | 8,006,317 B (7.64 MiB) | 2,517,265 B (2.40 MiB) | 31.465813% | 31.440986% | 4,023,316 B (3.84 MiB) | 50.291450% | 50.251770% | 59.828862% | 1,755,643 B (1.67 MiB) | 21.945538% | 21.928222% | -30.255933% | yes | 60 |
| `URLHash` | `int64` | Structured medium/high-cardinality numeric streams | median row-group cardinality 3292; median cardinality/rows 27.575808%; plain+zstd 3,580,060 B (3.41 MiB); rle+zstd 4,529,372 B (4.32 MiB) | The high-cardinality numeric/timestamp column produced a larger RLE-dict+ZSTD stream than plain+ZSTD. | 3,001 / 3,292 / 7,420 | 27.575808% | 8 | 8 | 8 | 8,000,000 B (7.63 MiB) | 8,006,310 B (7.64 MiB) | 3,580,060 B (3.41 MiB) | 44.750750% | 44.715481% | 4,529,372 B (4.32 MiB) | 56.617150% | 56.572528% | 26.516651% | 1,565,061 B (1.49 MiB) | 19.563263% | 19.547844% | -56.283945% | yes | 57 |
| `ResponseStartTiming` | `int32` | Structured medium/high-cardinality numeric streams | median row-group cardinality 1112; median cardinality/rows 9.314793%; plain+zstd 1,245,745 B (1.19 MiB); rle+zstd 1,556,751 B (1.48 MiB) | The high-cardinality numeric/timestamp column produced a larger RLE-dict+ZSTD stream than plain+ZSTD. | 800 / 1,112 / 3,761 | 9.314793% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,050 B (3.82 MiB) | 1,245,745 B (1.19 MiB) | 31.143625% | 31.104356% | 1,556,751 B (1.48 MiB) | 38.918775% | 38.869702% | 24.965462% | 1,001,805 B (978.33 KiB) | 25.045125% | 25.013545% | -19.581857% | yes | 58 |
| `RefererHash` | `int64` | Structured medium/high-cardinality numeric streams | median row-group cardinality 2729; median cardinality/rows 22.859776%; plain+zstd 2,841,886 B (2.71 MiB); rle+zstd 3,502,907 B (3.34 MiB) | The high-cardinality numeric/timestamp column produced a larger RLE-dict+ZSTD stream than plain+ZSTD. | 378 / 2,729 / 6,051 | 22.859776% | 8 | 8 | 8 | 8,000,000 B (7.63 MiB) | 8,006,316 B (7.64 MiB) | 2,841,886 B (2.71 MiB) | 35.523575% | 35.495551% | 3,502,907 B (3.34 MiB) | 43.786338% | 43.751795% | 23.259941% | 1,272,178 B (1.21 MiB) | 15.902225% | 15.889680% | -55.234728% | yes | 59 |

## Snappy RLE Dict Worse Distribution By Category

For columns where the best observed `snappy + plain` compressed byte count is smaller than the best observed `snappy + rle-dict` compressed byte count, each category image plots `plain + snappy` compressed bytes on the x-axis and `rle-dict + snappy` compressed bytes on the y-axis using the same log byte scale. Points above the diagonal are larger with RLE dictionary encoding. Point color is bucketed by `plain/no-compression encoded bytes / rle-dict + snappy compressed bytes`, so high-ratio colors identify columns where RLE dictionary lost the head-to-head but still compressed the baseline dramatically.

The bucket tables below each image show how much worse RLE dictionary encoding was. Worse-by percentage is `(rle_dict_snappy_compressed_bytes / plain_snappy_compressed_bytes - 1) * 100`, so values can exceed 100%.

The compressed bytes are Parquet column-chunk bytes, including dictionary pages and page headers. Dictionary-page byte breakdown columns are left blank when the cached Snappy result TSV does not contain those byte counts.

`Plain encoded bytes before compression` is the same column's byte count from the all-plain/no-compression baseline run. The `/ plain encoded` percentage columns compare compressed column bytes against that baseline denominator.

Categorization uses measured row-group cardinality and column type: `Medium-cardinality fixed-width numeric streams` means a non-timestamp numeric column has median cardinality/rows below 9%; `High-cardinality fixed-width IDs / hashes` means a non-timestamp numeric column has median cardinality/rows at least 9%; `High-cardinality timestamp streams` covers timestamp columns. Value-length distributions are included in the table for context, but these Snappy categories are driven by fixed-width type plus cardinality.

- Compared columns: `105`
- `snappy + rle-dict` worse than `snappy + plain`: `12`; better: `93`; ties: `0`; missing comparisons: `0`
- Missing shape stats while categorizing: `0`

| Category | Columns | Worse by min/median/max |
| --- | ---: | ---: |
| Medium-cardinality fixed-width numeric streams | 6 | 0.377184% / 7.527025% / 15.194024% |
| High-cardinality fixed-width IDs / hashes | 3 | 2.639084% / 21.157028% / 22.749867% |
| High-cardinality timestamp streams | 3 | 28.674248% / 28.732596% / 28.817198% |

### Medium-cardinality fixed-width numeric streams

Non-timestamp numeric columns with medium row-group cardinality; RLE dictionary reduced the pre-compression stream, but Snappy compressed the plain fixed-width stream to fewer bytes.

![Snappy RLE dictionary worse: Medium-cardinality fixed-width numeric streams](images/snappy_rle_dict_worse_medium_cardinality_fixed_width_numeric_streams.svg)

| Improvement bucket | `snappy + rle-dict` worse by |
| --- | ---: |
| `0-10%` | 4 |
| `10-20%` | 2 |
| `20-30%` | 0 |
| `30-40%` | 0 |
| `40-50%` | 0 |
| `50-60%` | 0 |
| `60-70%` | 0 |
| `70-80%` | 0 |
| `80-90%` | 0 |
| `90-100%` | 0 |
| `100-200%` | 0 |
| `200-500%` | 0 |
| `500%+` | 0 |

| Column | Type | Category | Measured feature | Measured reason | Row-group cardinality min/median/max | Median cardinality / rows | Min value length (B) | Median value length (B) | Max value length (B) | Physical bytes before encoding/compression | Plain encoded bytes before compression | RLE dict encoded bytes before compression | RLE dict encoded / plain encoded | Plain + Snappy compressed bytes | Plain + Snappy / physical | Plain + Snappy / plain encoded | RLE dict + Snappy compressed bytes | RLE dict + Snappy / physical | RLE dict + Snappy / plain encoded | RLE dict + Snappy vs plain + Snappy | RLE dict + Snappy without dict pages | RLE dict without dict pages / physical | RLE dict without dict pages / plain encoded | RLE dict without dict pages vs plain + Snappy | RLE + dict is better without including dict page | Dictionary pages |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | --- | ---: |
| `ClientIP` | `int32` | Medium-cardinality fixed-width numeric streams | median row-group cardinality 924; median cardinality/rows 7.739990%; rle encoded 23.427048% of plain encoded | RLE dictionary reduced pre-compression bytes, but Snappy compressed the plain fixed-width stream better for this column. | 744 / 924 / 1,957 | 7.739990% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,050 B (3.82 MiB) | 938,265 B (916.27 KiB) | 23.427048% | 719,059 B (702.21 KiB) | 17.976475% | 17.953808% | 828,313 B (808.90 KiB) | 20.707825% | 20.681714% | 15.194024% | 491,147 B (479.64 KiB) | 12.278675% | 12.263193% | -31.695869% | yes | 61 |
| `IPNetworkID` | `int32` | Medium-cardinality fixed-width numeric streams | median row-group cardinality 600; median cardinality/rows 5.025967%; rle encoded 18.662230% of plain encoded | RLE dictionary reduced pre-compression bytes, but Snappy compressed the plain fixed-width stream better for this column. | 498 / 600 / 1,095 | 5.025967% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,052 B (3.82 MiB) | 747,432 B (729.91 KiB) | 18.662230% | 560,148 B (547.02 KiB) | 14.003700% | 13.986036% | 636,803 B (621.88 KiB) | 15.920075% | 15.899993% | 13.684776% | 436,391 B (426.16 KiB) | 10.909775% | 10.896013% | -22.093625% | yes | 58 |
| `WindowClientHeight` | `int16` | Medium-cardinality fixed-width numeric streams | median row-group cardinality 435; median cardinality/rows 3.643826%; rle encoded 18.722233% of plain encoded | RLE dictionary reduced pre-compression bytes, but Snappy compressed the plain fixed-width stream better for this column. | 318 / 435 / 575 | 3.643826% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,051 B (3.82 MiB) | 749,835 B (732.26 KiB) | 18.722233% | 551,321 B (538.40 KiB) | 13.783025% | 13.765642% | 601,354 B (587.26 KiB) | 15.033850% | 15.014890% | 9.075112% | 481,703 B (470.41 KiB) | 12.042575% | 12.027387% | -12.627489% | yes | 58 |
| `RemoteIP` | `int32` | Medium-cardinality fixed-width numeric streams | median row-group cardinality 851; median cardinality/rows 7.128497%; rle encoded 23.165746% of plain encoded | RLE dictionary reduced pre-compression bytes, but Snappy compressed the plain fixed-width stream better for this column. | 508 / 851 / 1,951 | 7.128497% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,047 B (3.82 MiB) | 927,799 B (906.05 KiB) | 23.165746% | 706,731 B (690.17 KiB) | 17.668275% | 17.646010% | 748,986 B (731.43 KiB) | 18.724650% | 18.701054% | 5.978937% | 458,600 B (447.85 KiB) | 11.465000% | 11.450552% | -35.109681% | yes | 61 |
| `UserID` | `int64` | Medium-cardinality fixed-width numeric streams | median row-group cardinality 898; median cardinality/rows 7.522198%; rle encoded 15.311738% of plain encoded | RLE dictionary reduced pre-compression bytes, but Snappy compressed the plain fixed-width stream better for this column. | 716 / 898 / 1,805 | 7.522198% | 8 | 8 | 8 | 8,000,000 B (7.63 MiB) | 8,006,315 B (7.64 MiB) | 1,225,906 B (1.17 MiB) | 15.311738% | 1,084,714 B (1.03 MiB) | 13.558925% | 13.548230% | 1,120,481 B (1.07 MiB) | 14.006013% | 13.994965% | 3.297367% | 474,344 B (463.23 KiB) | 5.929300% | 5.924623% | -56.270132% | yes | 61 |
| `FetchTiming` | `int32` | Medium-cardinality fixed-width numeric streams | median row-group cardinality 664; median cardinality/rows 5.562071%; rle encoded 28.140701% of plain encoded | RLE dictionary reduced pre-compression bytes, but Snappy compressed the plain fixed-width stream better for this column. | 329 / 664 / 1,264 | 5.562071% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,053 B (3.82 MiB) | 1,127,050 B (1.07 MiB) | 28.140701% | 808,624 B (789.67 KiB) | 20.215600% | 20.190095% | 811,674 B (792.65 KiB) | 20.291850% | 20.266249% | 0.377184% | 606,675 B (592.46 KiB) | 15.166875% | 15.147740% | -24.974401% | yes | 59 |

### High-cardinality fixed-width IDs / hashes

Non-timestamp numeric ID/hash-like columns with high row-group cardinality; dictionary IDs had too little repetition to beat Snappy over plain fixed-width values.

![Snappy RLE dictionary worse: High-cardinality fixed-width IDs / hashes](images/snappy_rle_dict_worse_high_cardinality_fixed_width_ids_hashes.svg)

| Improvement bucket | `snappy + rle-dict` worse by |
| --- | ---: |
| `0-10%` | 1 |
| `10-20%` | 0 |
| `20-30%` | 2 |
| `30-40%` | 0 |
| `40-50%` | 0 |
| `50-60%` | 0 |
| `60-70%` | 0 |
| `70-80%` | 0 |
| `80-90%` | 0 |
| `90-100%` | 0 |
| `100-200%` | 0 |
| `200-500%` | 0 |
| `500%+` | 0 |

| Column | Type | Category | Measured feature | Measured reason | Row-group cardinality min/median/max | Median cardinality / rows | Min value length (B) | Median value length (B) | Max value length (B) | Physical bytes before encoding/compression | Plain encoded bytes before compression | RLE dict encoded bytes before compression | RLE dict encoded / plain encoded | Plain + Snappy compressed bytes | Plain + Snappy / physical | Plain + Snappy / plain encoded | RLE dict + Snappy compressed bytes | RLE dict + Snappy / physical | RLE dict + Snappy / plain encoded | RLE dict + Snappy vs plain + Snappy | RLE dict + Snappy without dict pages | RLE dict without dict pages / physical | RLE dict without dict pages / plain encoded | RLE dict without dict pages vs plain + Snappy | RLE + dict is better without including dict page | Dictionary pages |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | --- | ---: |
| `WatchID` | `int64` | High-cardinality fixed-width IDs / hashes | median row-group cardinality 11938; median cardinality/rows 100.000000%; rle encoded 122.722135% of plain encoded | RLE dictionary was already larger than plain before Snappy; the compressed result stayed larger. | 9,315 / 11,938 / 14,202 | 100.000000% | 8 | 8 | 8 | 8,000,000 B (7.63 MiB) | 8,006,312 B (7.64 MiB) | 9,825,517 B (9.37 MiB) | 122.722135% | 8,005,128 B (7.63 MiB) | 100.064100% | 99.985212% | 9,826,284 B (9.37 MiB) | 122.828550% | 122.731715% | 22.749867% | 1,825,340 B (1.74 MiB) | 22.816750% | 22.798762% | -77.197866% | yes | 59 |
| `HID` | `int32` | High-cardinality fixed-width IDs / hashes | median row-group cardinality 5965; median cardinality/rows 49.966494%; rle encoded 111.552287% of plain encoded | RLE dictionary was already larger than plain before Snappy; the compressed result stayed larger. | 5,818 / 5,965 / 13,281 | 49.966494% | 4 | 4 | 4 | 4,000,000 B (3.81 MiB) | 4,005,051 B (3.82 MiB) | 4,467,726 B (4.26 MiB) | 111.552287% | 3,688,155 B (3.52 MiB) | 92.203875% | 92.087591% | 4,468,459 B (4.26 MiB) | 111.711475% | 111.570589% | 21.157028% | 1,728,746 B (1.65 MiB) | 43.218650% | 43.164144% | -53.127078% | yes | 61 |
| `URLHash` | `int64` | High-cardinality fixed-width IDs / hashes | median row-group cardinality 3292; median cardinality/rows 27.575808%; rle encoded 57.494464% of plain encoded | High cardinality limited dictionary benefit; Snappy over plain fixed-width values stayed smaller. | 3,001 / 3,292 / 7,420 | 27.575808% | 8 | 8 | 8 | 8,000,000 B (7.63 MiB) | 8,006,310 B (7.64 MiB) | 4,603,185 B (4.39 MiB) | 57.494464% | 4,382,430 B (4.18 MiB) | 54.780375% | 54.737201% | 4,498,086 B (4.29 MiB) | 56.226075% | 56.181762% | 2.639084% | 1,523,644 B (1.45 MiB) | 19.045550% | 19.030540% | -65.232896% | yes | 60 |

### High-cardinality timestamp streams

Timestamp columns with high row-group cardinality; RLE dictionary barely reduced the encoded stream, and Snappy did better on the plain timestamp bytes.

![Snappy RLE dictionary worse: High-cardinality timestamp streams](images/snappy_rle_dict_worse_high_cardinality_timestamp_streams.svg)

| Improvement bucket | `snappy + rle-dict` worse by |
| --- | ---: |
| `0-10%` | 0 |
| `10-20%` | 0 |
| `20-30%` | 3 |
| `30-40%` | 0 |
| `40-50%` | 0 |
| `50-60%` | 0 |
| `60-70%` | 0 |
| `70-80%` | 0 |
| `80-90%` | 0 |
| `90-100%` | 0 |
| `100-200%` | 0 |
| `200-500%` | 0 |
| `500%+` | 0 |

| Column | Type | Category | Measured feature | Measured reason | Row-group cardinality min/median/max | Median cardinality / rows | Min value length (B) | Median value length (B) | Max value length (B) | Physical bytes before encoding/compression | Plain encoded bytes before compression | RLE dict encoded bytes before compression | RLE dict encoded / plain encoded | Plain + Snappy compressed bytes | Plain + Snappy / physical | Plain + Snappy / plain encoded | RLE dict + Snappy compressed bytes | RLE dict + Snappy / physical | RLE dict + Snappy / plain encoded | RLE dict + Snappy vs plain + Snappy | RLE dict + Snappy without dict pages | RLE dict without dict pages / physical | RLE dict without dict pages / plain encoded | RLE dict without dict pages vs plain + Snappy | RLE + dict is better without including dict page | Dictionary pages |
| --- | --- | --- | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | --- | ---: |
| `EventTime` | `timestamp_millis` | High-cardinality timestamp streams | median row-group cardinality 6237; median cardinality/rows 52.244932%; rle encoded 91.735386% of plain encoded | High-cardinality timestamp values left little dictionary repetition; Snappy compressed the plain timestamp stream to fewer bytes. | 5,997 / 6,237 / 13,042 | 52.244932% | 8 | 8 | 8 | 8,000,000 B (7.63 MiB) | 8,006,315 B (7.64 MiB) | 7,344,624 B (7.00 MiB) | 91.735386% | 4,282,415 B (4.08 MiB) | 53.530188% | 53.487965% | 5,516,487 B (5.26 MiB) | 68.956087% | 68.901698% | 28.817198% | 1,730,189 B (1.65 MiB) | 21.627363% | 21.610304% | -59.597820% | yes | 62 |
| `LocalEventTime` | `timestamp_millis` | High-cardinality timestamp streams | median row-group cardinality 6254; median cardinality/rows 52.387335%; rle encoded 91.671601% of plain encoded | High-cardinality timestamp values left little dictionary repetition; Snappy compressed the plain timestamp stream to fewer bytes. | 5,968 / 6,254 / 13,047 | 52.387335% | 8 | 8 | 8 | 8,000,000 B (7.63 MiB) | 8,006,317 B (7.64 MiB) | 7,339,519 B (7.00 MiB) | 91.671601% | 4,283,734 B (4.09 MiB) | 53.546675% | 53.504427% | 5,514,562 B (5.26 MiB) | 68.932025% | 68.877637% | 28.732596% | 1,728,196 B (1.65 MiB) | 21.602450% | 21.585406% | -59.656785% | yes | 62 |
| `ClientEventTime` | `timestamp_millis` | High-cardinality timestamp streams | median row-group cardinality 5882; median cardinality/rows 49.271235%; rle encoded 90.150524% of plain encoded | High-cardinality timestamp values left little dictionary repetition; Snappy compressed the plain timestamp stream to fewer bytes. | 5,666 / 5,882 / 13,078 | 49.271235% | 8 | 8 | 8 | 8,000,000 B (7.63 MiB) | 8,006,314 B (7.64 MiB) | 7,217,734 B (6.88 MiB) | 90.150524% | 4,228,484 B (4.03 MiB) | 52.856050% | 52.814366% | 5,440,970 B (5.19 MiB) | 68.012125% | 67.958489% | 28.674248% | 1,727,979 B (1.65 MiB) | 21.599737% | 21.582703% | -59.134787% | yes | 62 |

## Delta Binary Packed Winner vs Second Best Improvement Distribution

For each column, this looks at the `Compressed overall` ranking below. Only columns where `delta-binary-packed` is the best observed compressed result are bucketed. Improvement is `(second-best compressed bytes - delta-binary-packed compressed bytes) / second-best compressed bytes * 100`.

![Delta binary packed winner improvement over second best](images/delta_binary_packed_winner_vs_second_best_improvement.svg)

- Delta-binary-packed winner columns: `2`
- Missing second-best rows: `0`

| Improvement bucket | `delta-binary-packed` better than second best |
| --- | ---: |
| `0-10%` | 2 |
| `10-20%` | 0 |

## Snappy Plain vs RLE Dict Improvement Distribution

For each column, this compares the best observed `snappy + plain` compressed byte count with the best observed `snappy + rle-dict` compressed byte count. Improvement is `(larger compressed bytes - smaller compressed bytes) / larger compressed bytes * 100`.

- Compared columns: `105`
- `snappy + rle-dict` smaller: `93`; `snappy + plain` smaller: `12`; ties: `0`; missing comparisons: `0`

![Snappy RLE dictionary improvement over plain](images/snappy_rle_dict_better_than_plain_improvement.svg)

`snappy + rle-dict` better buckets:

| Improvement bucket | `snappy + rle-dict` better |
| --- | ---: |
| `0-10%` | 5 |
| `10-20%` | 3 |
| `20-30%` | 5 |
| `30-40%` | 2 |
| `40-50%` | 5 |
| `50-60%` | 7 |
| `60-70%` | 9 |
| `70-80%` | 8 |
| `80-90%` | 13 |
| `90-100%` | 36 |

![Snappy plain improvement over RLE dictionary](images/snappy_plain_better_than_rle_dict_improvement.svg)

`snappy + plain` better buckets:

| Improvement bucket | `snappy + plain` better |
| --- | ---: |
| `0-10%` | 5 |
| `10-20%` | 4 |
| `20-30%` | 3 |

## AdvEngineID (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 3 / 4`
- Page cardinality per row group min/median/max of mins: `1 / 3 / 4`; of maxes: `1 / 3 / 4`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/advengineid_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/advengineid_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/advengineid_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/advengineid_value_length.svg)


Compressed overall:
1. 22,787 B (22.25 KiB) compressed - `zstd-3` + `rle-dict`; 34,609 B (33.80 KiB) encoded; 175.760434x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 29,802 B (29.10 KiB) compressed - `snappy` + `rle-dict`; 34,781 B (33.97 KiB) encoded; 134.388732x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 30,642 B (29.92 KiB) compressed - `zstd-3` + `plain`; 4,003,579 B (3.82 MiB) encoded; 130.704686x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
4. 44,831 B (43.78 KiB) compressed - `zstd-3` + `delta-binary-packed`; 200,545 B (195.84 KiB) encoded; 89.336687x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
5. 53,198 B (51.95 KiB) compressed - `snappy` + `delta-binary-packed`; 202,399 B (197.66 KiB) encoded; 75.285781x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 22,787 B (22.25 KiB) compressed - `zstd-3` + `rle-dict`; 34,609 B (33.80 KiB) encoded; 175.760434x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 30,642 B (29.92 KiB) compressed - `zstd-3` + `plain`; 4,003,579 B (3.82 MiB) encoded; 130.704686x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
3. 44,831 B (43.78 KiB) compressed - `zstd-3` + `delta-binary-packed`; 200,545 B (195.84 KiB) encoded; 89.336687x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

Snappy:
1. 29,802 B (29.10 KiB) compressed - `snappy` + `rle-dict`; 34,781 B (33.97 KiB) encoded; 134.388732x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 53,198 B (51.95 KiB) compressed - `snappy` + `delta-binary-packed`; 202,399 B (197.66 KiB) encoded; 75.285781x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-plain`
3. 220,197 B (215.04 KiB) compressed - `snappy` + `plain`; 4,003,714 B (3.82 MiB) encoded; 18.188499x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`

## Age (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `6 / 6 / 6`
- Page cardinality per row group min/median/max of mins: `6 / 6 / 6`; of maxes: `6 / 6 / 6`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/age_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/age_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/age_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/age_value_length.svg)


Compressed overall:
1. 117,935 B (115.17 KiB) compressed - `zstd-3` + `rle-dict`; 186,008 B (181.65 KiB) encoded; 33.959825x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-rle-dict`
2. 142,526 B (139.19 KiB) compressed - `zstd-3` + `plain`; 4,003,588 B (3.82 MiB) encoded; 28.100501x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`
3. 155,617 B (151.97 KiB) compressed - `snappy` + `rle-dict`; 186,110 B (181.75 KiB) encoded; 25.736597x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
4. 245,551 B (239.80 KiB) compressed - `zstd-3` + `delta-binary-packed`; 743,503 B (726.08 KiB) encoded; 16.310469x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
5. 286,272 B (279.56 KiB) compressed - `snappy` + `delta-binary-packed`; 742,356 B (724.96 KiB) encoded; 13.990373x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-rle-dict-date-rle-dict-ts-delta-binary-packed`

ZSTD:
1. 117,935 B (115.17 KiB) compressed - `zstd-3` + `rle-dict`; 186,008 B (181.65 KiB) encoded; 33.959825x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-rle-dict`
2. 142,526 B (139.19 KiB) compressed - `zstd-3` + `plain`; 4,003,588 B (3.82 MiB) encoded; 28.100501x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`
3. 245,551 B (239.80 KiB) compressed - `zstd-3` + `delta-binary-packed`; 743,503 B (726.08 KiB) encoded; 16.310469x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`

Snappy:
1. 155,617 B (151.97 KiB) compressed - `snappy` + `rle-dict`; 186,110 B (181.75 KiB) encoded; 25.736597x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
2. 286,272 B (279.56 KiB) compressed - `snappy` + `delta-binary-packed`; 742,356 B (724.96 KiB) encoded; 13.990373x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
3. 333,804 B (325.98 KiB) compressed - `snappy` + `plain`; 4,003,715 B (3.82 MiB) encoded; 11.998215x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`

## BrowserCountry (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `3 / 5 / 11`
- Page cardinality per row group min/median/max of mins: `3 / 5 / 11`; of maxes: `3 / 5 / 11`
- Value length min/median/max: `2 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `2 / 2 / 2`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/browsercountry_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/browsercountry_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/browsercountry_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/browsercountry_value_length.svg)


Compressed overall:
1. 69,192 B (67.57 KiB) compressed - `zstd-3` + `rle-dict`; 136,881 B (133.67 KiB) encoded; 105.935339x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
2. 94,626 B (92.41 KiB) compressed - `snappy` + `rle-dict`; 136,126 B (132.94 KiB) encoded; 77.461564x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-delta-binary-packed`
3. 122,185 B (119.32 KiB) compressed - `zstd-3` + `plain`; 7,328,504 B (6.99 MiB) encoded; 59.989999x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-rle-dict`
4. 185,035 B (180.70 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 3,639,209 B (3.47 MiB) encoded; 39.613468x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
5. 194,675 B (190.11 KiB) compressed - `zstd-3` + `delta-byte-array`; 960,047 B (937.55 KiB) encoded; 37.651871x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`

ZSTD:
1. 69,192 B (67.57 KiB) compressed - `zstd-3` + `rle-dict`; 136,881 B (133.67 KiB) encoded; 105.935339x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
2. 122,185 B (119.32 KiB) compressed - `zstd-3` + `plain`; 7,328,504 B (6.99 MiB) encoded; 59.989999x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-rle-dict`
3. 185,035 B (180.70 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 3,639,209 B (3.47 MiB) encoded; 39.613468x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 194,675 B (190.11 KiB) compressed - `zstd-3` + `delta-byte-array`; 960,047 B (937.55 KiB) encoded; 37.651871x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`

Snappy:
1. 94,626 B (92.41 KiB) compressed - `snappy` + `rle-dict`; 136,126 B (132.94 KiB) encoded; 77.461564x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-delta-binary-packed`
2. 253,396 B (247.46 KiB) compressed - `snappy` + `delta-byte-array`; 961,111 B (938.58 KiB) encoded; 28.926573x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 352,235 B (343.98 KiB) compressed - `snappy` + `delta-length-byte-array`; 3,638,826 B (3.47 MiB) encoded; 20.809624x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
4. 457,031 B (446.32 KiB) compressed - `snappy` + `plain`; 7,328,708 B (6.99 MiB) encoded; 16.038032x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`

## BrowserLanguage (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `3 / 7 / 14`
- Page cardinality per row group min/median/max of mins: `3 / 7 / 14`; of maxes: `3 / 7 / 14`
- Value length min/median/max: `2 / 2 / 3` bytes
- Value length per row group min/median/max of mins: `2 / 2 / 2`; of maxes: `2 / 3 / 3`

![Row-group cardinality](column_shape_stats/images/browserlanguage_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/browserlanguage_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/browserlanguage_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/browserlanguage_value_length.svg)


Compressed overall:
1. 27,489 B (26.84 KiB) compressed - `zstd-3` + `rle-dict`; 42,266 B (41.28 KiB) encoded; 218.478701x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 31,938 B (31.19 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 2,051,163 B (1.96 MiB) encoded; 188.044367x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
3. 32,231 B (31.48 KiB) compressed - `zstd-3` + `plain`; 6,004,431 B (5.73 MiB) encoded; 186.334926x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
4. 34,402 B (33.60 KiB) compressed - `snappy` + `rle-dict`; 42,486 B (41.49 KiB) encoded; 174.575926x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`
5. 64,686 B (63.17 KiB) compressed - `zstd-3` + `delta-byte-array`; 333,815 B (325.99 KiB) encoded; 92.844835x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

ZSTD:
1. 27,489 B (26.84 KiB) compressed - `zstd-3` + `rle-dict`; 42,266 B (41.28 KiB) encoded; 218.478701x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 31,938 B (31.19 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 2,051,163 B (1.96 MiB) encoded; 188.044367x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
3. 32,231 B (31.48 KiB) compressed - `zstd-3` + `plain`; 6,004,431 B (5.73 MiB) encoded; 186.334926x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
4. 64,686 B (63.17 KiB) compressed - `zstd-3` + `delta-byte-array`; 333,815 B (325.99 KiB) encoded; 92.844835x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 34,402 B (33.60 KiB) compressed - `snappy` + `rle-dict`; 42,486 B (41.49 KiB) encoded; 174.575926x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`
2. 89,360 B (87.27 KiB) compressed - `snappy` + `delta-byte-array`; 333,048 B (325.24 KiB) encoded; 67.208606x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 125,242 B (122.31 KiB) compressed - `snappy` + `delta-length-byte-array`; 2,051,313 B (1.96 MiB) encoded; 47.953251x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
4. 321,259 B (313.73 KiB) compressed - `snappy` + `plain`; 6,004,612 B (5.73 MiB) encoded; 18.694452x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`

## CLID (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 2`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 2`; of maxes: `1 / 1 / 2`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/clid_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/clid_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/clid_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/clid_value_length.svg)


Compressed overall:
1. 5,627 B (5.50 KiB) compressed - `zstd-3` + `plain`; 4,003,542 B (3.82 MiB) encoded; 711.756531x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed`
2. 6,231 B (6.08 KiB) compressed - `snappy` + `rle-dict`; 5,999 B (5.86 KiB) encoded; 642.762638x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 6,908 B (6.75 KiB) compressed - `zstd-3` + `rle-dict`; 5,882 B (5.74 KiB) encoded; 579.770411x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
4. 8,385 B (8.19 KiB) compressed - `zstd-3` + `delta-binary-packed`; 68,797 B (67.18 KiB) encoded; 477.645081x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
5. 10,366 B (10.12 KiB) compressed - `snappy` + `delta-binary-packed`; 68,651 B (67.04 KiB) encoded; 386.364461x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 5,627 B (5.50 KiB) compressed - `zstd-3` + `plain`; 4,003,542 B (3.82 MiB) encoded; 711.756531x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed`
2. 6,908 B (6.75 KiB) compressed - `zstd-3` + `rle-dict`; 5,882 B (5.74 KiB) encoded; 579.770411x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 8,385 B (8.19 KiB) compressed - `zstd-3` + `delta-binary-packed`; 68,797 B (67.18 KiB) encoded; 477.645081x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 6,231 B (6.08 KiB) compressed - `snappy` + `rle-dict`; 5,999 B (5.86 KiB) encoded; 642.762638x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 10,366 B (10.12 KiB) compressed - `snappy` + `delta-binary-packed`; 68,651 B (67.04 KiB) encoded; 386.364461x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`
3. 204,500 B (199.71 KiB) compressed - `snappy` + `plain`; 4,003,715 B (3.82 MiB) encoded; 19.584616x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`

## ClientEventTime (timestamp_millis)

Column shape stats:
- Parquet type: `TIMESTAMP(isAdjustedToUTC=true,unit=MILLIS)`; physical type: `INT64`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `5,666 / 5,882 / 13,078`
- Page cardinality per row group min/median/max of mins: `5,666 / 5,882 / 13,078`; of maxes: `5,666 / 5,882 / 13,078`
- Value length min/median/max: `8 / 8 / 8` bytes
- Value length per row group min/median/max of mins: `8 / 8 / 8`; of maxes: `8 / 8 / 8`

![Row-group cardinality](column_shape_stats/images/clienteventtime_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/clienteventtime_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/clienteventtime_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/clienteventtime_value_length.svg)


Compressed overall:
1. 2,474,093 B (2.36 MiB) compressed - `zstd-3` + `plain`; 8,004,796 B (7.63 MiB) encoded; 3.236060x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 2,895,535 B (2.76 MiB) compressed - `zstd-3` + `delta-binary-packed`; 3,488,803 B (3.33 MiB) encoded; 2.765055x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
3. 3,045,886 B (2.90 MiB) compressed - `snappy` + `delta-binary-packed`; 3,492,148 B (3.33 MiB) encoded; 2.628567x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`
4. 3,958,277 B (3.77 MiB) compressed - `zstd-3` + `rle-dict`; 7,210,522 B (6.88 MiB) encoded; 2.022677x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-rle-dict`
5. 4,228,484 B (4.03 MiB) compressed - `snappy` + `plain`; 8,004,794 B (7.63 MiB) encoded; 1.893424x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-rle-dict-ts-plain`

ZSTD:
1. 2,474,093 B (2.36 MiB) compressed - `zstd-3` + `plain`; 8,004,796 B (7.63 MiB) encoded; 3.236060x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 2,895,535 B (2.76 MiB) compressed - `zstd-3` + `delta-binary-packed`; 3,488,803 B (3.33 MiB) encoded; 2.765055x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
3. 3,958,277 B (3.77 MiB) compressed - `zstd-3` + `rle-dict`; 7,210,522 B (6.88 MiB) encoded; 2.022677x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-rle-dict`

Snappy:
1. 3,045,886 B (2.90 MiB) compressed - `snappy` + `delta-binary-packed`; 3,492,148 B (3.33 MiB) encoded; 2.628567x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`
2. 4,228,484 B (4.03 MiB) compressed - `snappy` + `plain`; 8,004,794 B (7.63 MiB) encoded; 1.893424x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-rle-dict-ts-plain`
3. 5,440,970 B (5.19 MiB) compressed - `snappy` + `rle-dict`; 7,217,734 B (6.88 MiB) encoded; 1.471487x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-rle-dict-date-rle-dict-ts-rle-dict`

## ClientIP (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `744 / 924 / 1,957`
- Page cardinality per row group min/median/max of mins: `744 / 924 / 1,957`; of maxes: `744 / 924 / 1,957`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/clientip_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/clientip_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/clientip_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/clientip_value_length.svg)


Compressed overall:
1. 408,058 B (398.49 KiB) compressed - `zstd-3` + `plain`; 4,003,592 B (3.82 MiB) encoded; 9.814904x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
2. 719,059 B (702.21 KiB) compressed - `snappy` + `plain`; 4,003,772 B (3.82 MiB) encoded; 5.569849x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 806,924 B (788.01 KiB) compressed - `zstd-3` + `rle-dict`; 941,748 B (919.68 KiB) encoded; 4.963355x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
4. 828,313 B (808.90 KiB) compressed - `snappy` + `rle-dict`; 938,265 B (916.27 KiB) encoded; 4.835189x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-delta-binary-packed-ts-rle-dict`
5. 859,503 B (839.36 KiB) compressed - `zstd-3` + `delta-binary-packed`; 3,746,908 B (3.57 MiB) encoded; 4.659728x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-rle-dict`

ZSTD:
1. 408,058 B (398.49 KiB) compressed - `zstd-3` + `plain`; 4,003,592 B (3.82 MiB) encoded; 9.814904x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
2. 806,924 B (788.01 KiB) compressed - `zstd-3` + `rle-dict`; 941,748 B (919.68 KiB) encoded; 4.963355x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 859,503 B (839.36 KiB) compressed - `zstd-3` + `delta-binary-packed`; 3,746,908 B (3.57 MiB) encoded; 4.659728x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-rle-dict`

Snappy:
1. 719,059 B (702.21 KiB) compressed - `snappy` + `plain`; 4,003,772 B (3.82 MiB) encoded; 5.569849x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 828,313 B (808.90 KiB) compressed - `snappy` + `rle-dict`; 938,265 B (916.27 KiB) encoded; 4.835189x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-delta-binary-packed-ts-rle-dict`
3. 1,054,378 B (1.01 MiB) compressed - `snappy` + `delta-binary-packed`; 3,744,635 B (3.57 MiB) encoded; 3.798495x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`

## ClientTimeZone (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `10 / 15 / 20`
- Page cardinality per row group min/median/max of mins: `10 / 15 / 20`; of maxes: `10 / 15 / 20`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/clienttimezone_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/clienttimezone_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/clienttimezone_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/clienttimezone_value_length.svg)


Compressed overall:
1. 84,413 B (82.43 KiB) compressed - `zstd-3` + `rle-dict`; 164,301 B (160.45 KiB) encoded; 47.445903x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
2. 99,843 B (97.50 KiB) compressed - `zstd-3` + `plain`; 4,003,590 B (3.82 MiB) encoded; 40.113488x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-rle-dict`
3. 110,114 B (107.53 KiB) compressed - `snappy` + `rle-dict`; 162,472 B (158.66 KiB) encoded; 36.371860x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
4. 211,859 B (206.89 KiB) compressed - `zstd-3` + `delta-binary-packed`; 976,871 B (953.98 KiB) encoded; 18.904323x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
5. 251,618 B (245.72 KiB) compressed - `snappy` + `delta-binary-packed`; 975,773 B (952.90 KiB) encoded; 15.917188x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 84,413 B (82.43 KiB) compressed - `zstd-3` + `rle-dict`; 164,301 B (160.45 KiB) encoded; 47.445903x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
2. 99,843 B (97.50 KiB) compressed - `zstd-3` + `plain`; 4,003,590 B (3.82 MiB) encoded; 40.113488x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-rle-dict`
3. 211,859 B (206.89 KiB) compressed - `zstd-3` + `delta-binary-packed`; 976,871 B (953.98 KiB) encoded; 18.904323x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 110,114 B (107.53 KiB) compressed - `snappy` + `rle-dict`; 162,472 B (158.66 KiB) encoded; 36.371860x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 251,618 B (245.72 KiB) compressed - `snappy` + `delta-binary-packed`; 975,773 B (952.90 KiB) encoded; 15.917188x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 285,543 B (278.85 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 14.026087x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## CodeVersion (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 2 / 3`
- Page cardinality per row group min/median/max of mins: `1 / 2 / 3`; of maxes: `1 / 2 / 3`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/codeversion_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/codeversion_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/codeversion_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/codeversion_value_length.svg)


Compressed overall:
1. 7,031 B (6.87 KiB) compressed - `zstd-3` + `plain`; 4,003,547 B (3.82 MiB) encoded; 569.627791x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`
2. 7,284 B (7.11 KiB) compressed - `snappy` + `rle-dict`; 7,063 B (6.90 KiB) encoded; 549.842532x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 7,970 B (7.78 KiB) compressed - `zstd-3` + `rle-dict`; 6,944 B (6.78 KiB) encoded; 502.516060x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict`
4. 12,547 B (12.25 KiB) compressed - `zstd-3` + `delta-binary-packed`; 95,513 B (93.27 KiB) encoded; 319.204033x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain`
5. 15,076 B (14.72 KiB) compressed - `snappy` + `delta-binary-packed`; 96,584 B (94.32 KiB) encoded; 265.657535x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-delta-binary-packed`

ZSTD:
1. 7,031 B (6.87 KiB) compressed - `zstd-3` + `plain`; 4,003,547 B (3.82 MiB) encoded; 569.627791x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`
2. 7,970 B (7.78 KiB) compressed - `zstd-3` + `rle-dict`; 6,944 B (6.78 KiB) encoded; 502.516060x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict`
3. 12,547 B (12.25 KiB) compressed - `zstd-3` + `delta-binary-packed`; 95,513 B (93.27 KiB) encoded; 319.204033x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain`

Snappy:
1. 7,284 B (7.11 KiB) compressed - `snappy` + `rle-dict`; 7,063 B (6.90 KiB) encoded; 549.842532x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 15,076 B (14.72 KiB) compressed - `snappy` + `delta-binary-packed`; 96,584 B (94.32 KiB) encoded; 265.657535x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 205,059 B (200.25 KiB) compressed - `snappy` + `plain`; 4,003,714 B (3.82 MiB) encoded; 19.531223x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## ConnectTiming (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `117 / 222 / 628`
- Page cardinality per row group min/median/max of mins: `117 / 222 / 628`; of maxes: `117 / 222 / 628`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/connecttiming_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/connecttiming_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/connecttiming_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/connecttiming_value_length.svg)


Compressed overall:
1. 333,011 B (325.21 KiB) compressed - `zstd-3` + `plain`; 4,003,660 B (3.82 MiB) encoded; 12.026783x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict`
2. 377,469 B (368.62 KiB) compressed - `zstd-3` + `rle-dict`; 703,627 B (687.14 KiB) encoded; 10.610278x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 461,361 B (450.55 KiB) compressed - `snappy` + `rle-dict`; 699,339 B (682.95 KiB) encoded; 8.680948x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 564,134 B (550.91 KiB) compressed - `snappy` + `plain`; 4,003,753 B (3.82 MiB) encoded; 7.099468x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
5. 580,418 B (566.81 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,090,679 B (1.04 MiB) encoded; 6.900287x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 333,011 B (325.21 KiB) compressed - `zstd-3` + `plain`; 4,003,660 B (3.82 MiB) encoded; 12.026783x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict`
2. 377,469 B (368.62 KiB) compressed - `zstd-3` + `rle-dict`; 703,627 B (687.14 KiB) encoded; 10.610278x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 580,418 B (566.81 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,090,679 B (1.04 MiB) encoded; 6.900287x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

Snappy:
1. 461,361 B (450.55 KiB) compressed - `snappy` + `rle-dict`; 699,339 B (682.95 KiB) encoded; 8.680948x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 564,134 B (550.91 KiB) compressed - `snappy` + `plain`; 4,003,753 B (3.82 MiB) encoded; 7.099468x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 625,501 B (610.84 KiB) compressed - `snappy` + `delta-binary-packed`; 1,090,357 B (1.04 MiB) encoded; 6.402949x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`

## CookieEnable (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 2 / 2`
- Page cardinality per row group min/median/max of mins: `1 / 2 / 2`; of maxes: `1 / 2 / 2`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/cookieenable_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/cookieenable_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/cookieenable_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/cookieenable_value_length.svg)


Compressed overall:
1. 6,053 B (5.91 KiB) compressed - `zstd-3` + `plain`; 4,003,539 B (3.82 MiB) encoded; 661.663968x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-rle-dict`
2. 6,402 B (6.25 KiB) compressed - `snappy` + `rle-dict`; 6,166 B (6.02 KiB) encoded; 625.593877x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 7,077 B (6.91 KiB) compressed - `zstd-3` + `rle-dict`; 6,049 B (5.91 KiB) encoded; 565.925110x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
4. 7,474 B (7.30 KiB) compressed - `zstd-3` + `delta-binary-packed`; 46,548 B (45.46 KiB) encoded; 535.864597x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
5. 9,014 B (8.80 KiB) compressed - `snappy` + `delta-binary-packed`; 46,660 B (45.57 KiB) encoded; 444.314622x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

ZSTD:
1. 6,053 B (5.91 KiB) compressed - `zstd-3` + `plain`; 4,003,539 B (3.82 MiB) encoded; 661.663968x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-rle-dict`
2. 7,077 B (6.91 KiB) compressed - `zstd-3` + `rle-dict`; 6,049 B (5.91 KiB) encoded; 565.925110x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 7,474 B (7.30 KiB) compressed - `zstd-3` + `delta-binary-packed`; 46,548 B (45.46 KiB) encoded; 535.864597x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 6,402 B (6.25 KiB) compressed - `snappy` + `rle-dict`; 6,166 B (6.02 KiB) encoded; 625.593877x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 9,014 B (8.80 KiB) compressed - `snappy` + `delta-binary-packed`; 46,660 B (45.57 KiB) encoded; 444.314622x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 204,770 B (199.97 KiB) compressed - `snappy` + `plain`; 4,003,777 B (3.82 MiB) encoded; 19.558783x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`

## CounterClass (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/counterclass_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/counterclass_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/counterclass_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/counterclass_value_length.svg)


Compressed overall:
1. 4,228 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
2. 4,932 B (4.82 KiB) compressed - `zstd-3` + `delta-binary-packed`; 43,044 B (42.04 KiB) encoded; 812.054745x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 5,273 B (5.15 KiB) compressed - `snappy` + `rle-dict`; 5,041 B (4.92 KiB) encoded; 759.539920x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
4. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
5. 6,203 B (6.06 KiB) compressed - `snappy` + `delta-binary-packed`; 43,179 B (42.17 KiB) encoded; 645.664034x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 4,228 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
2. 4,932 B (4.82 KiB) compressed - `zstd-3` + `delta-binary-packed`; 43,044 B (42.04 KiB) encoded; 812.054745x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 5,273 B (5.15 KiB) compressed - `snappy` + `rle-dict`; 5,041 B (4.92 KiB) encoded; 759.539920x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 6,203 B (6.06 KiB) compressed - `snappy` + `delta-binary-packed`; 43,179 B (42.17 KiB) encoded; 645.664034x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 204,060 B (199.28 KiB) compressed - `snappy` + `plain`; 4,003,712 B (3.82 MiB) encoded; 19.626845x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## CounterID (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `true`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 4`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 4`; of maxes: `1 / 1 / 4`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/counterid_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/counterid_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/counterid_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/counterid_value_length.svg)


Compressed overall:
1. 4,931 B (4.82 KiB) compressed - `zstd-3` + `plain`; 4,003,528 B (3.82 MiB) encoded; 812.219023x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 5,299 B (5.17 KiB) compressed - `snappy` + `rle-dict`; 5,067 B (4.95 KiB) encoded; 755.812795x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 5,835 B (5.70 KiB) compressed - `zstd-3` + `delta-binary-packed`; 47,098 B (45.99 KiB) encoded; 686.384233x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 6,020 B (5.88 KiB) compressed - `zstd-3` + `rle-dict`; 4,994 B (4.88 KiB) encoded; 665.291030x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
5. 7,269 B (7.10 KiB) compressed - `snappy` + `delta-binary-packed`; 46,534 B (45.44 KiB) encoded; 550.977026x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-delta-binary-packed`

ZSTD:
1. 4,931 B (4.82 KiB) compressed - `zstd-3` + `plain`; 4,003,528 B (3.82 MiB) encoded; 812.219023x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 5,835 B (5.70 KiB) compressed - `zstd-3` + `delta-binary-packed`; 47,098 B (45.99 KiB) encoded; 686.384233x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 6,020 B (5.88 KiB) compressed - `zstd-3` + `rle-dict`; 4,994 B (4.88 KiB) encoded; 665.291030x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 5,299 B (5.17 KiB) compressed - `snappy` + `rle-dict`; 5,067 B (4.95 KiB) encoded; 755.812795x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 7,269 B (7.10 KiB) compressed - `snappy` + `delta-binary-packed`; 46,534 B (45.44 KiB) encoded; 550.977026x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 204,252 B (199.46 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 19.608386x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## DNSTiming (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `68 / 98 / 352`
- Page cardinality per row group min/median/max of mins: `68 / 98 / 352`; of maxes: `68 / 98 / 352`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/dnstiming_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/dnstiming_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/dnstiming_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/dnstiming_value_length.svg)


Compressed overall:
1. 135,085 B (131.92 KiB) compressed - `zstd-3` + `plain`; 4,003,588 B (3.82 MiB) encoded; 29.648377x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 156,185 B (152.52 KiB) compressed - `zstd-3` + `rle-dict`; 343,417 B (335.37 KiB) encoded; 25.642994x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 193,951 B (189.41 KiB) compressed - `snappy` + `rle-dict`; 344,174 B (336.11 KiB) encoded; 20.649808x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 301,429 B (294.36 KiB) compressed - `zstd-3` + `delta-binary-packed`; 806,355 B (787.46 KiB) encoded; 13.286880x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
5. 326,903 B (319.24 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 12.251497x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

ZSTD:
1. 135,085 B (131.92 KiB) compressed - `zstd-3` + `plain`; 4,003,588 B (3.82 MiB) encoded; 29.648377x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 156,185 B (152.52 KiB) compressed - `zstd-3` + `rle-dict`; 343,417 B (335.37 KiB) encoded; 25.642994x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 301,429 B (294.36 KiB) compressed - `zstd-3` + `delta-binary-packed`; 806,355 B (787.46 KiB) encoded; 13.286880x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 193,951 B (189.41 KiB) compressed - `snappy` + `rle-dict`; 344,174 B (336.11 KiB) encoded; 20.649808x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
2. 326,903 B (319.24 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 12.251497x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 341,399 B (333.40 KiB) compressed - `snappy` + `delta-binary-packed`; 805,791 B (786.91 KiB) encoded; 11.731291x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## DontCountHits (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `2 / 2 / 2`
- Page cardinality per row group min/median/max of mins: `2 / 2 / 2`; of maxes: `2 / 2 / 2`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/dontcounthits_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/dontcounthits_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/dontcounthits_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/dontcounthits_value_length.svg)


Compressed overall:
1. 49,438 B (48.28 KiB) compressed - `zstd-3` + `rle-dict`; 65,952 B (64.41 KiB) encoded; 81.011550x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
2. 59,494 B (58.10 KiB) compressed - `snappy` + `rle-dict`; 66,028 B (64.48 KiB) encoded; 67.318536x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 60,370 B (58.96 KiB) compressed - `zstd-3` + `delta-binary-packed`; 168,379 B (164.43 KiB) encoded; 66.341709x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 82,484 B (80.55 KiB) compressed - `zstd-3` + `plain`; 4,003,585 B (3.82 MiB) encoded; 48.555465x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
5. 82,543 B (80.61 KiB) compressed - `snappy` + `delta-binary-packed`; 168,534 B (164.58 KiB) encoded; 48.520759x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`

ZSTD:
1. 49,438 B (48.28 KiB) compressed - `zstd-3` + `rle-dict`; 65,952 B (64.41 KiB) encoded; 81.011550x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
2. 60,370 B (58.96 KiB) compressed - `zstd-3` + `delta-binary-packed`; 168,379 B (164.43 KiB) encoded; 66.341709x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 82,484 B (80.55 KiB) compressed - `zstd-3` + `plain`; 4,003,585 B (3.82 MiB) encoded; 48.555465x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`

Snappy:
1. 59,494 B (58.10 KiB) compressed - `snappy` + `rle-dict`; 66,028 B (64.48 KiB) encoded; 67.318536x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 82,543 B (80.61 KiB) compressed - `snappy` + `delta-binary-packed`; 168,534 B (164.58 KiB) encoded; 48.520759x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
3. 295,526 B (288.60 KiB) compressed - `snappy` + `plain`; 4,003,719 B (3.82 MiB) encoded; 13.552273x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`

## EventDate (date)

Column shape stats:
- Parquet type: `DATE`; physical type: `INT32`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/eventdate_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/eventdate_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/eventdate_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/eventdate_value_length.svg)


Compressed overall:
1. 4,905 B (4.79 KiB) compressed - `zstd-3` + `plain`; 4,003,532 B (3.82 MiB) encoded; 816.522936x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-delta-binary-packed`
2. 5,275 B (5.15 KiB) compressed - `snappy` + `rle-dict`; 5,043 B (4.92 KiB) encoded; 759.250237x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
3. 5,976 B (5.84 KiB) compressed - `zstd-3` + `rle-dict`; 4,950 B (4.83 KiB) encoded; 670.188253x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
4. 6,282 B (6.13 KiB) compressed - `zstd-3` + `delta-binary-packed`; 50,562 B (49.38 KiB) encoded; 637.542980x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
5. 7,838 B (7.65 KiB) compressed - `snappy` + `delta-binary-packed`; 51,405 B (50.20 KiB) encoded; 510.977928x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

ZSTD:
1. 4,905 B (4.79 KiB) compressed - `zstd-3` + `plain`; 4,003,532 B (3.82 MiB) encoded; 816.522936x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-delta-binary-packed`
2. 5,976 B (5.84 KiB) compressed - `zstd-3` + `rle-dict`; 4,950 B (4.83 KiB) encoded; 670.188253x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
3. 6,282 B (6.13 KiB) compressed - `zstd-3` + `delta-binary-packed`; 50,562 B (49.38 KiB) encoded; 637.542980x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 5,275 B (5.15 KiB) compressed - `snappy` + `rle-dict`; 5,043 B (4.92 KiB) encoded; 759.250237x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
2. 7,838 B (7.65 KiB) compressed - `snappy` + `delta-binary-packed`; 51,405 B (50.20 KiB) encoded; 510.977928x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 204,170 B (199.38 KiB) compressed - `snappy` + `plain`; 4,003,652 B (3.82 MiB) encoded; 19.616227x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## EventTime (timestamp_millis)

Column shape stats:
- Parquet type: `TIMESTAMP(isAdjustedToUTC=true,unit=MILLIS)`; physical type: `INT64`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `5,997 / 6,237 / 13,042`
- Page cardinality per row group min/median/max of mins: `5,997 / 6,237 / 13,042`; of maxes: `5,997 / 6,237 / 13,042`
- Value length min/median/max: `8 / 8 / 8` bytes
- Value length per row group min/median/max of mins: `8 / 8 / 8`; of maxes: `8 / 8 / 8`

![Row-group cardinality](column_shape_stats/images/eventtime_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/eventtime_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/eventtime_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/eventtime_value_length.svg)


Compressed overall:
1. 2,514,539 B (2.40 MiB) compressed - `zstd-3` + `plain`; 8,004,631 B (7.63 MiB) encoded; 3.184009x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-plain`
2. 2,885,504 B (2.75 MiB) compressed - `zstd-3` + `delta-binary-packed`; 3,369,293 B (3.21 MiB) encoded; 2.774668x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`
3. 3,032,179 B (2.89 MiB) compressed - `snappy` + `delta-binary-packed`; 3,368,385 B (3.21 MiB) encoded; 2.640449x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`
4. 4,021,616 B (3.84 MiB) compressed - `zstd-3` + `rle-dict`; 7,360,398 B (7.02 MiB) encoded; 1.990820x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-rle-dict`
5. 4,282,415 B (4.08 MiB) compressed - `snappy` + `plain`; 8,004,714 B (7.63 MiB) encoded; 1.869579x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 2,514,539 B (2.40 MiB) compressed - `zstd-3` + `plain`; 8,004,631 B (7.63 MiB) encoded; 3.184009x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-plain`
2. 2,885,504 B (2.75 MiB) compressed - `zstd-3` + `delta-binary-packed`; 3,369,293 B (3.21 MiB) encoded; 2.774668x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`
3. 4,021,616 B (3.84 MiB) compressed - `zstd-3` + `rle-dict`; 7,360,398 B (7.02 MiB) encoded; 1.990820x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-rle-dict`

Snappy:
1. 3,032,179 B (2.89 MiB) compressed - `snappy` + `delta-binary-packed`; 3,368,385 B (3.21 MiB) encoded; 2.640449x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`
2. 4,282,415 B (4.08 MiB) compressed - `snappy` + `plain`; 8,004,714 B (7.63 MiB) encoded; 1.869579x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
3. 5,516,487 B (5.26 MiB) compressed - `snappy` + `rle-dict`; 7,344,624 B (7.00 MiB) encoded; 1.451343x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict`

## FUniqID (int64)

Column shape stats:
- Parquet type: `INT(64,true)`; physical type: `INT64`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `628 / 812 / 1,703`
- Page cardinality per row group min/median/max of mins: `628 / 812 / 1,703`; of maxes: `628 / 812 / 1,703`
- Value length min/median/max: `8 / 8 / 8` bytes
- Value length per row group min/median/max of mins: `8 / 8 / 8`; of maxes: `8 / 8 / 8`

![Row-group cardinality](column_shape_stats/images/funiqid_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/funiqid_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/funiqid_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/funiqid_value_length.svg)


Compressed overall:
1. 693,929 B (677.67 KiB) compressed - `zstd-3` + `plain`; 8,004,554 B (7.63 MiB) encoded; 11.537660x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
2. 1,058,653 B (1.01 MiB) compressed - `zstd-3` + `rle-dict`; 1,186,134 B (1.13 MiB) encoded; 7.562740x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 1,071,313 B (1.02 MiB) compressed - `snappy` + `rle-dict`; 1,179,670 B (1.13 MiB) encoded; 7.473369x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`
4. 1,151,399 B (1.10 MiB) compressed - `snappy` + `plain`; 8,004,714 B (7.63 MiB) encoded; 6.953556x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
5. 1,414,138 B (1.35 MiB) compressed - `zstd-3` + `delta-binary-packed`; 7,453,534 B (7.11 MiB) encoded; 5.661624x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain`

ZSTD:
1. 693,929 B (677.67 KiB) compressed - `zstd-3` + `plain`; 8,004,554 B (7.63 MiB) encoded; 11.537660x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
2. 1,058,653 B (1.01 MiB) compressed - `zstd-3` + `rle-dict`; 1,186,134 B (1.13 MiB) encoded; 7.562740x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 1,414,138 B (1.35 MiB) compressed - `zstd-3` + `delta-binary-packed`; 7,453,534 B (7.11 MiB) encoded; 5.661624x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain`

Snappy:
1. 1,071,313 B (1.02 MiB) compressed - `snappy` + `rle-dict`; 1,179,670 B (1.13 MiB) encoded; 7.473369x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`
2. 1,151,399 B (1.10 MiB) compressed - `snappy` + `plain`; 8,004,714 B (7.63 MiB) encoded; 6.953556x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 1,778,224 B (1.70 MiB) compressed - `snappy` + `delta-binary-packed`; 7,453,850 B (7.11 MiB) encoded; 4.502423x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## FetchTiming (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `329 / 664 / 1,264`
- Page cardinality per row group min/median/max of mins: `329 / 664 / 1,264`; of maxes: `329 / 664 / 1,264`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/fetchtiming_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/fetchtiming_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/fetchtiming_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/fetchtiming_value_length.svg)


Compressed overall:
1. 549,819 B (536.93 KiB) compressed - `zstd-3` + `plain`; 4,003,625 B (3.82 MiB) encoded; 7.284312x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`
2. 685,133 B (669.08 KiB) compressed - `zstd-3` + `rle-dict`; 1,131,789 B (1.08 MiB) encoded; 5.845658x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 808,624 B (789.67 KiB) compressed - `snappy` + `plain`; 4,003,769 B (3.82 MiB) encoded; 4.952924x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
4. 811,674 B (792.65 KiB) compressed - `snappy` + `rle-dict`; 1,127,050 B (1.07 MiB) encoded; 4.934312x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`
5. 963,559 B (940.98 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,547,166 B (1.48 MiB) encoded; 4.156521x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 549,819 B (536.93 KiB) compressed - `zstd-3` + `plain`; 4,003,625 B (3.82 MiB) encoded; 7.284312x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-plain`
2. 685,133 B (669.08 KiB) compressed - `zstd-3` + `rle-dict`; 1,131,789 B (1.08 MiB) encoded; 5.845658x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 963,559 B (940.98 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,547,166 B (1.48 MiB) encoded; 4.156521x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

Snappy:
1. 808,624 B (789.67 KiB) compressed - `snappy` + `plain`; 4,003,769 B (3.82 MiB) encoded; 4.952924x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
2. 811,674 B (792.65 KiB) compressed - `snappy` + `rle-dict`; 1,127,050 B (1.07 MiB) encoded; 4.934312x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`
3. 1,012,893 B (989.15 KiB) compressed - `snappy` + `delta-binary-packed`; 1,546,128 B (1.47 MiB) encoded; 3.954073x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## FlashMajor (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `4 / 5 / 7`
- Page cardinality per row group min/median/max of mins: `4 / 5 / 7`; of maxes: `4 / 5 / 7`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/flashmajor_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/flashmajor_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/flashmajor_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/flashmajor_value_length.svg)


Compressed overall:
1. 49,021 B (47.87 KiB) compressed - `zstd-3` + `rle-dict`; 88,382 B (86.31 KiB) encoded; 81.700761x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 53,248 B (52.00 KiB) compressed - `zstd-3` + `plain`; 4,003,589 B (3.82 MiB) encoded; 75.215088x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 67,838 B (66.25 KiB) compressed - `snappy` + `rle-dict`; 88,003 B (85.94 KiB) encoded; 59.038489x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
4. 92,739 B (90.57 KiB) compressed - `zstd-3` + `delta-binary-packed`; 399,413 B (390.05 KiB) encoded; 43.186286x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain`
5. 113,209 B (110.56 KiB) compressed - `snappy` + `delta-binary-packed`; 398,480 B (389.14 KiB) encoded; 35.377514x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 49,021 B (47.87 KiB) compressed - `zstd-3` + `rle-dict`; 88,382 B (86.31 KiB) encoded; 81.700761x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 53,248 B (52.00 KiB) compressed - `zstd-3` + `plain`; 4,003,589 B (3.82 MiB) encoded; 75.215088x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 92,739 B (90.57 KiB) compressed - `zstd-3` + `delta-binary-packed`; 399,413 B (390.05 KiB) encoded; 43.186286x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain`

Snappy:
1. 67,838 B (66.25 KiB) compressed - `snappy` + `rle-dict`; 88,003 B (85.94 KiB) encoded; 59.038489x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 113,209 B (110.56 KiB) compressed - `snappy` + `delta-binary-packed`; 398,480 B (389.14 KiB) encoded; 35.377514x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`
3. 245,188 B (239.44 KiB) compressed - `snappy` + `plain`; 4,003,711 B (3.82 MiB) encoded; 16.334621x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## FlashMinor (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `9 / 9 / 9`
- Page cardinality per row group min/median/max of mins: `9 / 9 / 9`; of maxes: `9 / 9 / 9`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/flashminor_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/flashminor_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/flashminor_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/flashminor_value_length.svg)


Compressed overall:
1. 113,166 B (110.51 KiB) compressed - `zstd-3` + `rle-dict`; 214,760 B (209.73 KiB) encoded; 35.390912x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
2. 126,765 B (123.79 KiB) compressed - `zstd-3` + `plain`; 4,003,586 B (3.82 MiB) encoded; 31.594273x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-plain`
3. 152,997 B (149.41 KiB) compressed - `snappy` + `rle-dict`; 214,831 B (209.80 KiB) encoded; 26.177298x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 198,971 B (194.31 KiB) compressed - `zstd-3` + `delta-binary-packed`; 523,936 B (511.66 KiB) encoded; 20.128803x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
5. 234,096 B (228.61 KiB) compressed - `snappy` + `delta-binary-packed`; 523,422 B (511.15 KiB) encoded; 17.108571x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-delta-binary-packed`

ZSTD:
1. 113,166 B (110.51 KiB) compressed - `zstd-3` + `rle-dict`; 214,760 B (209.73 KiB) encoded; 35.390912x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
2. 126,765 B (123.79 KiB) compressed - `zstd-3` + `plain`; 4,003,586 B (3.82 MiB) encoded; 31.594273x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-plain`
3. 198,971 B (194.31 KiB) compressed - `zstd-3` + `delta-binary-packed`; 523,936 B (511.66 KiB) encoded; 20.128803x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`

Snappy:
1. 152,997 B (149.41 KiB) compressed - `snappy` + `rle-dict`; 214,831 B (209.80 KiB) encoded; 26.177298x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 234,096 B (228.61 KiB) compressed - `snappy` + `delta-binary-packed`; 523,422 B (511.15 KiB) encoded; 17.108571x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 328,315 B (320.62 KiB) compressed - `snappy` + `plain`; 4,003,715 B (3.82 MiB) encoded; 12.198797x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`

## FlashMinor2 (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `18 / 24 / 29`
- Page cardinality per row group min/median/max of mins: `18 / 24 / 29`; of maxes: `18 / 24 / 29`
- Value length min/median/max: `0 / 3 / 8` bytes
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `8 / 8 / 8`

![Row-group cardinality](column_shape_stats/images/flashminor2_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/flashminor2_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/flashminor2_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/flashminor2_value_length.svg)


Compressed overall:
1. 145,400 B (141.99 KiB) compressed - `zstd-3` + `rle-dict`; 266,705 B (260.45 KiB) encoded; 50.613508x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 182,282 B (178.01 KiB) compressed - `snappy` + `rle-dict`; 266,872 B (260.62 KiB) encoded; 40.372631x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
3. 246,055 B (240.29 KiB) compressed - `zstd-3` + `plain`; 7,357,830 B (7.02 MiB) encoded; 29.908776x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict`
4. 266,508 B (260.26 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 3,707,674 B (3.54 MiB) encoded; 27.613445x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
5. 325,949 B (318.31 KiB) compressed - `zstd-3` + `delta-byte-array`; 1,042,299 B (1017.87 KiB) encoded; 22.577778x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 145,400 B (141.99 KiB) compressed - `zstd-3` + `rle-dict`; 266,705 B (260.45 KiB) encoded; 50.613508x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 246,055 B (240.29 KiB) compressed - `zstd-3` + `plain`; 7,357,830 B (7.02 MiB) encoded; 29.908776x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-rle-dict`
3. 266,508 B (260.26 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 3,707,674 B (3.54 MiB) encoded; 27.613445x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
4. 325,949 B (318.31 KiB) compressed - `zstd-3` + `delta-byte-array`; 1,042,299 B (1017.87 KiB) encoded; 22.577778x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`

Snappy:
1. 182,282 B (178.01 KiB) compressed - `snappy` + `rle-dict`; 266,872 B (260.62 KiB) encoded; 40.372631x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
2. 393,931 B (384.70 KiB) compressed - `snappy` + `delta-byte-array`; 1,042,680 B (1018.24 KiB) encoded; 18.681454x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
3. 432,351 B (422.22 KiB) compressed - `snappy` + `delta-length-byte-array`; 3,707,874 B (3.54 MiB) encoded; 17.021365x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
4. 536,601 B (524.02 KiB) compressed - `snappy` + `plain`; 7,358,039 B (7.02 MiB) encoded; 13.714481x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`

## FromTag (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 15`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 15`; of maxes: `1 / 1 / 15`
- Value length min/median/max: `0 / 0 / 12` bytes
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 0 / 12`

![Row-group cardinality](column_shape_stats/images/fromtag_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/fromtag_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/fromtag_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/fromtag_value_length.svg)


Compressed overall:
1. 22,483 B (21.96 KiB) compressed - `zstd-3` + `rle-dict`; 41,201 B (40.24 KiB) encoded; 180.110395x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
2. 28,366 B (27.70 KiB) compressed - `zstd-3` + `plain`; 4,048,288 B (3.86 MiB) encoded; 142.756187x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed`
3. 28,495 B (27.83 KiB) compressed - `snappy` + `rle-dict`; 41,170 B (40.21 KiB) encoded; 142.109914x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
4. 36,140 B (35.29 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 210,629 B (205.69 KiB) encoded; 112.048201x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
5. 38,213 B (37.32 KiB) compressed - `zstd-3` + `delta-byte-array`; 260,375 B (254.27 KiB) encoded; 105.969749x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`

ZSTD:
1. 22,483 B (21.96 KiB) compressed - `zstd-3` + `rle-dict`; 41,201 B (40.24 KiB) encoded; 180.110395x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
2. 28,366 B (27.70 KiB) compressed - `zstd-3` + `plain`; 4,048,288 B (3.86 MiB) encoded; 142.756187x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed`
3. 36,140 B (35.29 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 210,629 B (205.69 KiB) encoded; 112.048201x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
4. 38,213 B (37.32 KiB) compressed - `zstd-3` + `delta-byte-array`; 260,375 B (254.27 KiB) encoded; 105.969749x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`

Snappy:
1. 28,495 B (27.83 KiB) compressed - `snappy` + `rle-dict`; 41,170 B (40.21 KiB) encoded; 142.109914x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 45,963 B (44.89 KiB) compressed - `snappy` + `delta-length-byte-array`; 211,803 B (206.84 KiB) encoded; 88.101778x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-delta-binary-packed`
3. 49,905 B (48.74 KiB) compressed - `snappy` + `delta-byte-array`; 260,742 B (254.63 KiB) encoded; 81.142611x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
4. 221,650 B (216.46 KiB) compressed - `snappy` + `plain`; 4,048,459 B (3.86 MiB) encoded; 18.269443x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-delta-binary-packed-ts-delta-binary-packed`

## GoodEvent (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/goodevent_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/goodevent_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/goodevent_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/goodevent_value_length.svg)


Compressed overall:
1. 4,915 B (4.80 KiB) compressed - `zstd-3` + `plain`; 4,003,530 B (3.82 MiB) encoded; 814.863276x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 5,273 B (5.15 KiB) compressed - `snappy` + `rle-dict`; 5,041 B (4.92 KiB) encoded; 759.539731x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 5,585 B (5.45 KiB) compressed - `zstd-3` + `delta-binary-packed`; 43,693 B (42.67 KiB) encoded; 717.108863x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301757x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
5. 6,932 B (6.77 KiB) compressed - `snappy` + `delta-binary-packed`; 43,746 B (42.72 KiB) encoded; 577.762983x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 4,915 B (4.80 KiB) compressed - `zstd-3` + `plain`; 4,003,530 B (3.82 MiB) encoded; 814.863276x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 5,585 B (5.45 KiB) compressed - `zstd-3` + `delta-binary-packed`; 43,693 B (42.67 KiB) encoded; 717.108863x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301757x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 5,273 B (5.15 KiB) compressed - `snappy` + `rle-dict`; 5,041 B (4.92 KiB) encoded; 759.539731x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 6,932 B (6.77 KiB) compressed - `snappy` + `delta-binary-packed`; 43,746 B (42.72 KiB) encoded; 577.762983x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`
3. 204,227 B (199.44 KiB) compressed - `snappy` + `plain`; 4,003,712 B (3.82 MiB) encoded; 19.610791x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## HID (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `5,818 / 5,965 / 13,281`
- Page cardinality per row group min/median/max of mins: `5,818 / 5,965 / 13,281`; of maxes: `5,818 / 5,965 / 13,281`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/hid_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/hid_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/hid_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/hid_value_length.svg)


Compressed overall:
1. 3,688,155 B (3.52 MiB) compressed - `snappy` + `plain`; 4,003,834 B (3.82 MiB) encoded; 1.085923x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`
2. 3,759,716 B (3.59 MiB) compressed - `snappy` + `delta-binary-packed`; 3,953,327 B (3.77 MiB) encoded; 1.065254x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
3. 3,792,143 B (3.62 MiB) compressed - `zstd-3` + `plain`; 4,003,775 B (3.82 MiB) encoded; 1.056145x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed`
4. 3,949,027 B (3.77 MiB) compressed - `zstd-3` + `delta-binary-packed`; 3,952,659 B (3.77 MiB) encoded; 1.014187x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
5. 4,468,459 B (4.26 MiB) compressed - `snappy` + `rle-dict`; 4,467,726 B (4.26 MiB) encoded; 0.896294x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-rle-dict`

ZSTD:
1. 3,792,143 B (3.62 MiB) compressed - `zstd-3` + `plain`; 4,003,775 B (3.82 MiB) encoded; 1.056145x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed`
2. 3,949,027 B (3.77 MiB) compressed - `zstd-3` + `delta-binary-packed`; 3,952,659 B (3.77 MiB) encoded; 1.014187x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 4,492,493 B (4.28 MiB) compressed - `zstd-3` + `rle-dict`; 4,491,293 B (4.28 MiB) encoded; 0.891499x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-rle-dict`

Snappy:
1. 3,688,155 B (3.52 MiB) compressed - `snappy` + `plain`; 4,003,834 B (3.82 MiB) encoded; 1.085923x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`
2. 3,759,716 B (3.59 MiB) compressed - `snappy` + `delta-binary-packed`; 3,953,327 B (3.77 MiB) encoded; 1.065254x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
3. 4,468,459 B (4.26 MiB) compressed - `snappy` + `rle-dict`; 4,467,726 B (4.26 MiB) encoded; 0.896294x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-rle-dict`

## HTTPError (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/httperror_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/httperror_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/httperror_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/httperror_value_length.svg)


Compressed overall:
1. 4,228 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
2. 4,932 B (4.82 KiB) compressed - `zstd-3` + `delta-binary-packed`; 43,044 B (42.04 KiB) encoded; 812.054745x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 5,273 B (5.15 KiB) compressed - `snappy` + `rle-dict`; 5,041 B (4.92 KiB) encoded; 759.539920x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
4. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
5. 6,203 B (6.06 KiB) compressed - `snappy` + `delta-binary-packed`; 43,179 B (42.17 KiB) encoded; 645.664034x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 4,228 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
2. 4,932 B (4.82 KiB) compressed - `zstd-3` + `delta-binary-packed`; 43,044 B (42.04 KiB) encoded; 812.054745x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 5,273 B (5.15 KiB) compressed - `snappy` + `rle-dict`; 5,041 B (4.92 KiB) encoded; 759.539920x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 6,203 B (6.06 KiB) compressed - `snappy` + `delta-binary-packed`; 43,179 B (42.17 KiB) encoded; 645.664034x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 204,060 B (199.28 KiB) compressed - `snappy` + `plain`; 4,003,712 B (3.82 MiB) encoded; 19.626845x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## HasGCLID (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 2 / 2`
- Page cardinality per row group min/median/max of mins: `1 / 2 / 2`; of maxes: `1 / 2 / 2`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/hasgclid_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/hasgclid_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/hasgclid_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/hasgclid_value_length.svg)


Compressed overall:
1. 16,205 B (15.83 KiB) compressed - `zstd-3` + `rle-dict`; 21,326 B (20.83 KiB) encoded; 247.149151x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
2. 19,736 B (19.27 KiB) compressed - `snappy` + `rle-dict`; 21,395 B (20.89 KiB) encoded; 202.931293x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 21,114 B (20.62 KiB) compressed - `zstd-3` + `plain`; 4,003,558 B (3.82 MiB) encoded; 189.687032x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
4. 22,286 B (21.76 KiB) compressed - `zstd-3` + `delta-binary-packed`; 91,903 B (89.75 KiB) encoded; 179.711568x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
5. 28,367 B (27.70 KiB) compressed - `snappy` + `delta-binary-packed`; 92,303 B (90.14 KiB) encoded; 141.187013x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

ZSTD:
1. 16,205 B (15.83 KiB) compressed - `zstd-3` + `rle-dict`; 21,326 B (20.83 KiB) encoded; 247.149151x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
2. 21,114 B (20.62 KiB) compressed - `zstd-3` + `plain`; 4,003,558 B (3.82 MiB) encoded; 189.687032x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
3. 22,286 B (21.76 KiB) compressed - `zstd-3` + `delta-binary-packed`; 91,903 B (89.75 KiB) encoded; 179.711568x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

Snappy:
1. 19,736 B (19.27 KiB) compressed - `snappy` + `rle-dict`; 21,395 B (20.89 KiB) encoded; 202.931293x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 28,367 B (27.70 KiB) compressed - `snappy` + `delta-binary-packed`; 92,303 B (90.14 KiB) encoded; 141.187013x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 213,439 B (208.44 KiB) compressed - `snappy` + `plain`; 4,003,712 B (3.82 MiB) encoded; 18.764387x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`

## HistoryLength (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 4 / 87`
- Page cardinality per row group min/median/max of mins: `1 / 4 / 87`; of maxes: `1 / 4 / 87`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/historylength_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/historylength_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/historylength_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/historylength_value_length.svg)


Compressed overall:
1. 50,298 B (49.12 KiB) compressed - `zstd-3` + `rle-dict`; 69,568 B (67.94 KiB) encoded; 79.626506x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
2. 55,094 B (53.80 KiB) compressed - `zstd-3` + `plain`; 4,003,572 B (3.82 MiB) encoded; 72.694921x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 58,079 B (56.72 KiB) compressed - `snappy` + `rle-dict`; 68,952 B (67.34 KiB) encoded; 68.958729x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-delta-binary-packed`
4. 67,060 B (65.49 KiB) compressed - `zstd-3` + `delta-binary-packed`; 116,017 B (113.30 KiB) encoded; 59.723442x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain`
5. 68,715 B (67.10 KiB) compressed - `snappy` + `delta-binary-packed`; 116,370 B (113.64 KiB) encoded; 58.285003x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 50,298 B (49.12 KiB) compressed - `zstd-3` + `rle-dict`; 69,568 B (67.94 KiB) encoded; 79.626506x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
2. 55,094 B (53.80 KiB) compressed - `zstd-3` + `plain`; 4,003,572 B (3.82 MiB) encoded; 72.694921x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 67,060 B (65.49 KiB) compressed - `zstd-3` + `delta-binary-packed`; 116,017 B (113.30 KiB) encoded; 59.723442x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain`

Snappy:
1. 58,079 B (56.72 KiB) compressed - `snappy` + `rle-dict`; 68,952 B (67.34 KiB) encoded; 68.958729x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-delta-binary-packed`
2. 68,715 B (67.10 KiB) compressed - `snappy` + `delta-binary-packed`; 116,370 B (113.64 KiB) encoded; 58.285003x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`
3. 258,233 B (252.18 KiB) compressed - `snappy` + `plain`; 4,003,717 B (3.82 MiB) encoded; 15.509459x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## HitColor (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `2 / 3 / 3`
- Page cardinality per row group min/median/max of mins: `2 / 3 / 3`; of maxes: `2 / 3 / 3`
- Value length min/median/max: `1 / 1 / 1` bytes
- Value length per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`

![Row-group cardinality](column_shape_stats/images/hitcolor_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/hitcolor_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/hitcolor_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/hitcolor_value_length.svg)


Compressed overall:
1. 24,265 B (23.70 KiB) compressed - `zstd-3` + `rle-dict`; 33,649 B (32.86 KiB) encoded; 206.227323x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 28,126 B (27.47 KiB) compressed - `zstd-3` + `plain`; 5,002,901 B (4.77 MiB) encoded; 177.917443x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 28,185 B (27.52 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 1,043,109 B (1018.66 KiB) encoded; 177.545006x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
4. 31,870 B (31.12 KiB) compressed - `snappy` + `rle-dict`; 33,619 B (32.83 KiB) encoded; 157.016191x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
5. 50,721 B (49.53 KiB) compressed - `zstd-3` + `delta-byte-array`; 206,530 B (201.69 KiB) encoded; 98.659451x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

ZSTD:
1. 24,265 B (23.70 KiB) compressed - `zstd-3` + `rle-dict`; 33,649 B (32.86 KiB) encoded; 206.227323x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 28,126 B (27.47 KiB) compressed - `zstd-3` + `plain`; 5,002,901 B (4.77 MiB) encoded; 177.917443x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 28,185 B (27.52 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 1,043,109 B (1018.66 KiB) encoded; 177.545006x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
4. 50,721 B (49.53 KiB) compressed - `zstd-3` + `delta-byte-array`; 206,530 B (201.69 KiB) encoded; 98.659451x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 31,870 B (31.12 KiB) compressed - `snappy` + `rle-dict`; 33,619 B (32.83 KiB) encoded; 157.016191x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 74,207 B (72.47 KiB) compressed - `snappy` + `delta-byte-array`; 206,696 B (201.85 KiB) encoded; 67.434420x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 76,943 B (75.14 KiB) compressed - `snappy` + `delta-length-byte-array`; 1,043,282 B (1018.83 KiB) encoded; 65.036534x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
4. 294,288 B (287.39 KiB) compressed - `snappy` + `plain`; 5,003,058 B (4.77 MiB) encoded; 17.004112x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`

## IPNetworkID (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `498 / 600 / 1,095`
- Page cardinality per row group min/median/max of mins: `498 / 600 / 1,095`; of maxes: `498 / 600 / 1,095`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/ipnetworkid_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/ipnetworkid_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/ipnetworkid_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/ipnetworkid_value_length.svg)


Compressed overall:
1. 323,708 B (316.12 KiB) compressed - `zstd-3` + `plain`; 4,003,590 B (3.82 MiB) encoded; 12.372422x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`
2. 560,148 B (547.02 KiB) compressed - `snappy` + `plain`; 4,003,755 B (3.82 MiB) encoded; 7.149989x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 619,756 B (605.23 KiB) compressed - `zstd-3` + `rle-dict`; 748,476 B (730.93 KiB) encoded; 6.462305x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
4. 636,803 B (621.88 KiB) compressed - `snappy` + `rle-dict`; 747,432 B (729.91 KiB) encoded; 6.289311x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
5. 781,293 B (762.98 KiB) compressed - `zstd-3` + `delta-binary-packed`; 2,654,893 B (2.53 MiB) encoded; 5.126184x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 323,708 B (316.12 KiB) compressed - `zstd-3` + `plain`; 4,003,590 B (3.82 MiB) encoded; 12.372422x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`
2. 619,756 B (605.23 KiB) compressed - `zstd-3` + `rle-dict`; 748,476 B (730.93 KiB) encoded; 6.462305x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 781,293 B (762.98 KiB) compressed - `zstd-3` + `delta-binary-packed`; 2,654,893 B (2.53 MiB) encoded; 5.126184x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`

Snappy:
1. 560,148 B (547.02 KiB) compressed - `snappy` + `plain`; 4,003,755 B (3.82 MiB) encoded; 7.149989x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`
2. 636,803 B (621.88 KiB) compressed - `snappy` + `rle-dict`; 747,432 B (729.91 KiB) encoded; 6.289311x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 928,063 B (906.31 KiB) compressed - `snappy` + `delta-binary-packed`; 2,654,814 B (2.53 MiB) encoded; 4.315496x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`

## Income (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `3 / 4 / 4`
- Page cardinality per row group min/median/max of mins: `3 / 4 / 4`; of maxes: `3 / 4 / 4`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/income_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/income_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/income_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/income_value_length.svg)


Compressed overall:
1. 88,577 B (86.50 KiB) compressed - `zstd-3` + `rle-dict`; 142,532 B (139.19 KiB) encoded; 45.215507x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
2. 122,535 B (119.66 KiB) compressed - `zstd-3` + `plain`; 4,003,589 B (3.82 MiB) encoded; 32.684980x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
3. 122,841 B (119.96 KiB) compressed - `snappy` + `rle-dict`; 142,510 B (139.17 KiB) encoded; 32.603561x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
4. 140,019 B (136.74 KiB) compressed - `zstd-3` + `delta-binary-packed`; 328,738 B (321.03 KiB) encoded; 28.603647x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
5. 167,792 B (163.86 KiB) compressed - `snappy` + `delta-binary-packed`; 328,482 B (320.78 KiB) encoded; 23.869159x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-delta-binary-packed`

ZSTD:
1. 88,577 B (86.50 KiB) compressed - `zstd-3` + `rle-dict`; 142,532 B (139.19 KiB) encoded; 45.215507x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
2. 122,535 B (119.66 KiB) compressed - `zstd-3` + `plain`; 4,003,589 B (3.82 MiB) encoded; 32.684980x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
3. 140,019 B (136.74 KiB) compressed - `zstd-3` + `delta-binary-packed`; 328,738 B (321.03 KiB) encoded; 28.603647x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 122,841 B (119.96 KiB) compressed - `snappy` + `rle-dict`; 142,510 B (139.17 KiB) encoded; 32.603561x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 167,792 B (163.86 KiB) compressed - `snappy` + `delta-binary-packed`; 328,482 B (320.78 KiB) encoded; 23.869159x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 318,144 B (310.69 KiB) compressed - `snappy` + `plain`; 4,003,714 B (3.82 MiB) encoded; 12.588809x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## Interests (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `151 / 217 / 395`
- Page cardinality per row group min/median/max of mins: `151 / 217 / 395`; of maxes: `151 / 217 / 395`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/interests_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/interests_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/interests_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/interests_value_length.svg)


Compressed overall:
1. 193,369 B (188.84 KiB) compressed - `zstd-3` + `plain`; 4,003,589 B (3.82 MiB) encoded; 20.711955x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`
2. 310,021 B (302.75 KiB) compressed - `zstd-3` + `rle-dict`; 487,582 B (476.15 KiB) encoded; 12.918641x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-rle-dict`
3. 343,969 B (335.91 KiB) compressed - `snappy` + `rle-dict`; 483,872 B (472.53 KiB) encoded; 11.643636x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-rle-dict`
4. 385,284 B (376.25 KiB) compressed - `snappy` + `plain`; 4,003,716 B (3.82 MiB) encoded; 10.395059x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
5. 491,082 B (479.57 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,591,399 B (1.52 MiB) encoded; 8.155563x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`

ZSTD:
1. 193,369 B (188.84 KiB) compressed - `zstd-3` + `plain`; 4,003,589 B (3.82 MiB) encoded; 20.711955x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`
2. 310,021 B (302.75 KiB) compressed - `zstd-3` + `rle-dict`; 487,582 B (476.15 KiB) encoded; 12.918641x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-rle-dict`
3. 491,082 B (479.57 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,591,399 B (1.52 MiB) encoded; 8.155563x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`

Snappy:
1. 343,969 B (335.91 KiB) compressed - `snappy` + `rle-dict`; 483,872 B (472.53 KiB) encoded; 11.643636x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-rle-dict`
2. 385,284 B (376.25 KiB) compressed - `snappy` + `plain`; 4,003,716 B (3.82 MiB) encoded; 10.395059x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
3. 582,331 B (568.68 KiB) compressed - `snappy` + `delta-binary-packed`; 1,585,376 B (1.51 MiB) encoded; 6.877618x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`

## IsArtifical (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `2 / 2 / 2`
- Page cardinality per row group min/median/max of mins: `2 / 2 / 2`; of maxes: `2 / 2 / 2`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/isartifical_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/isartifical_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/isartifical_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/isartifical_value_length.svg)


Compressed overall:
1. 81,424 B (79.52 KiB) compressed - `snappy` + `rle-dict`; 81,070 B (79.17 KiB) encoded; 49.187561x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 82,011 B (80.09 KiB) compressed - `zstd-3` + `rle-dict`; 80,952 B (79.05 KiB) encoded; 48.835498x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`
3. 84,329 B (82.35 KiB) compressed - `zstd-3` + `delta-binary-packed`; 196,606 B (192.00 KiB) encoded; 47.493128x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain`
4. 118,508 B (115.73 KiB) compressed - `snappy` + `delta-binary-packed`; 196,750 B (192.14 KiB) encoded; 33.795592x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-plain`
5. 164,839 B (160.98 KiB) compressed - `zstd-3` + `plain`; 4,003,586 B (3.82 MiB) encoded; 24.296726x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`

ZSTD:
1. 82,011 B (80.09 KiB) compressed - `zstd-3` + `rle-dict`; 80,952 B (79.05 KiB) encoded; 48.835498x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`
2. 84,329 B (82.35 KiB) compressed - `zstd-3` + `delta-binary-packed`; 196,606 B (192.00 KiB) encoded; 47.493128x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain`
3. 164,839 B (160.98 KiB) compressed - `zstd-3` + `plain`; 4,003,586 B (3.82 MiB) encoded; 24.296726x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`

Snappy:
1. 81,424 B (79.52 KiB) compressed - `snappy` + `rle-dict`; 81,070 B (79.17 KiB) encoded; 49.187561x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 118,508 B (115.73 KiB) compressed - `snappy` + `delta-binary-packed`; 196,750 B (192.14 KiB) encoded; 33.795592x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-plain`
3. 505,682 B (493.83 KiB) compressed - `snappy` + `plain`; 4,003,749 B (3.82 MiB) encoded; 7.920092x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## IsDownload (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 2`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 2`; of maxes: `1 / 1 / 2`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/isdownload_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/isdownload_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/isdownload_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/isdownload_value_length.svg)


Compressed overall:
1. 8,024 B (7.84 KiB) compressed - `zstd-3` + `plain`; 4,003,554 B (3.82 MiB) encoded; 499.134472x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 8,618 B (8.42 KiB) compressed - `snappy` + `rle-dict`; 8,385 B (8.19 KiB) encoded; 464.731376x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 9,254 B (9.04 KiB) compressed - `zstd-3` + `rle-dict`; 8,328 B (8.13 KiB) encoded; 432.791766x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
4. 10,021 B (9.79 KiB) compressed - `zstd-3` + `delta-binary-packed`; 53,356 B (52.11 KiB) encoded; 399.666201x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain`
5. 12,157 B (11.87 KiB) compressed - `snappy` + `delta-binary-packed`; 53,537 B (52.28 KiB) encoded; 329.444353x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-delta-binary-packed`

ZSTD:
1. 8,024 B (7.84 KiB) compressed - `zstd-3` + `plain`; 4,003,554 B (3.82 MiB) encoded; 499.134472x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 9,254 B (9.04 KiB) compressed - `zstd-3` + `rle-dict`; 8,328 B (8.13 KiB) encoded; 432.791766x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
3. 10,021 B (9.79 KiB) compressed - `zstd-3` + `delta-binary-packed`; 53,356 B (52.11 KiB) encoded; 399.666201x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain`

Snappy:
1. 8,618 B (8.42 KiB) compressed - `snappy` + `rle-dict`; 8,385 B (8.19 KiB) encoded; 464.731376x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 12,157 B (11.87 KiB) compressed - `snappy` + `delta-binary-packed`; 53,537 B (52.28 KiB) encoded; 329.444353x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 206,523 B (201.68 KiB) compressed - `snappy` + `plain`; 4,003,711 B (3.82 MiB) encoded; 19.392779x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## IsEvent (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/isevent_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/isevent_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/isevent_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/isevent_value_length.svg)


Compressed overall:
1. 4,228 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
2. 4,932 B (4.82 KiB) compressed - `zstd-3` + `delta-binary-packed`; 43,044 B (42.04 KiB) encoded; 812.054745x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 5,273 B (5.15 KiB) compressed - `snappy` + `rle-dict`; 5,041 B (4.92 KiB) encoded; 759.539920x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
4. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
5. 6,203 B (6.06 KiB) compressed - `snappy` + `delta-binary-packed`; 43,179 B (42.17 KiB) encoded; 645.664034x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 4,228 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
2. 4,932 B (4.82 KiB) compressed - `zstd-3` + `delta-binary-packed`; 43,044 B (42.04 KiB) encoded; 812.054745x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 5,273 B (5.15 KiB) compressed - `snappy` + `rle-dict`; 5,041 B (4.92 KiB) encoded; 759.539920x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 6,203 B (6.06 KiB) compressed - `snappy` + `delta-binary-packed`; 43,179 B (42.17 KiB) encoded; 645.664034x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 204,060 B (199.28 KiB) compressed - `snappy` + `plain`; 4,003,712 B (3.82 MiB) encoded; 19.626845x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## IsLink (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 2`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 2`; of maxes: `1 / 1 / 2`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/islink_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/islink_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/islink_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/islink_value_length.svg)


Compressed overall:
1. 36,415 B (35.56 KiB) compressed - `zstd-3` + `rle-dict`; 54,152 B (52.88 KiB) encoded; 109.983606x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
2. 44,344 B (43.30 KiB) compressed - `zstd-3` + `delta-binary-packed`; 140,719 B (137.42 KiB) encoded; 90.317811x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 47,108 B (46.00 KiB) compressed - `snappy` + `rle-dict`; 54,197 B (52.93 KiB) encoded; 85.018532x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
4. 56,671 B (55.34 KiB) compressed - `zstd-3` + `plain`; 4,003,556 B (3.82 MiB) encoded; 70.672002x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
5. 59,460 B (58.07 KiB) compressed - `snappy` + `delta-binary-packed`; 140,796 B (137.50 KiB) encoded; 67.357097x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`

ZSTD:
1. 36,415 B (35.56 KiB) compressed - `zstd-3` + `rle-dict`; 54,152 B (52.88 KiB) encoded; 109.983606x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
2. 44,344 B (43.30 KiB) compressed - `zstd-3` + `delta-binary-packed`; 140,719 B (137.42 KiB) encoded; 90.317811x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 56,671 B (55.34 KiB) compressed - `zstd-3` + `plain`; 4,003,556 B (3.82 MiB) encoded; 70.672002x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`

Snappy:
1. 47,108 B (46.00 KiB) compressed - `snappy` + `rle-dict`; 54,197 B (52.93 KiB) encoded; 85.018532x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
2. 59,460 B (58.07 KiB) compressed - `snappy` + `delta-binary-packed`; 140,796 B (137.50 KiB) encoded; 67.357097x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
3. 256,364 B (250.36 KiB) compressed - `snappy` + `plain`; 4,003,714 B (3.82 MiB) encoded; 15.622525x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## IsMobile (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `2 / 2 / 2`
- Page cardinality per row group min/median/max of mins: `2 / 2 / 2`; of maxes: `2 / 2 / 2`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/ismobile_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/ismobile_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/ismobile_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/ismobile_value_length.svg)


Compressed overall:
1. 23,878 B (23.32 KiB) compressed - `zstd-3` + `rle-dict`; 29,965 B (29.26 KiB) encoded; 167.729584x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 28,143 B (27.48 KiB) compressed - `snappy` + `rle-dict`; 29,904 B (29.20 KiB) encoded; 142.310592x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
3. 28,623 B (27.95 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 139.924082x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
4. 32,828 B (32.06 KiB) compressed - `zstd-3` + `delta-binary-packed`; 114,545 B (111.86 KiB) encoded; 122.000944x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-delta-binary-packed`
5. 42,223 B (41.23 KiB) compressed - `snappy` + `delta-binary-packed`; 114,480 B (111.80 KiB) encoded; 94.854629x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 23,878 B (23.32 KiB) compressed - `zstd-3` + `rle-dict`; 29,965 B (29.26 KiB) encoded; 167.729584x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 28,623 B (27.95 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 139.924082x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
3. 32,828 B (32.06 KiB) compressed - `zstd-3` + `delta-binary-packed`; 114,545 B (111.86 KiB) encoded; 122.000944x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-delta-binary-packed`

Snappy:
1. 28,143 B (27.48 KiB) compressed - `snappy` + `rle-dict`; 29,904 B (29.20 KiB) encoded; 142.310592x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
2. 42,223 B (41.23 KiB) compressed - `snappy` + `delta-binary-packed`; 114,480 B (111.80 KiB) encoded; 94.854629x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`
3. 218,332 B (213.21 KiB) compressed - `snappy` + `plain`; 4,003,711 B (3.82 MiB) encoded; 18.343839x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## IsNotBounce (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 2`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 2`; of maxes: `1 / 1 / 2`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/isnotbounce_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/isnotbounce_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/isnotbounce_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/isnotbounce_value_length.svg)


Compressed overall:
1. 16,386 B (16.00 KiB) compressed - `zstd-3` + `delta-binary-packed`; 63,947 B (62.45 KiB) encoded; 244.419260x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain`
2. 16,585 B (16.20 KiB) compressed - `snappy` + `rle-dict`; 16,335 B (15.95 KiB) encoded; 241.486524x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 17,292 B (16.89 KiB) compressed - `zstd-3` + `rle-dict`; 16,261 B (15.88 KiB) encoded; 231.613116x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
4. 23,116 B (22.57 KiB) compressed - `snappy` + `delta-binary-packed`; 64,017 B (62.52 KiB) encoded; 173.258955x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
5. 25,161 B (24.57 KiB) compressed - `zstd-3` + `plain`; 4,003,536 B (3.82 MiB) encoded; 159.177060x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 16,386 B (16.00 KiB) compressed - `zstd-3` + `delta-binary-packed`; 63,947 B (62.45 KiB) encoded; 244.419260x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain`
2. 17,292 B (16.89 KiB) compressed - `zstd-3` + `rle-dict`; 16,261 B (15.88 KiB) encoded; 231.613116x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 25,161 B (24.57 KiB) compressed - `zstd-3` + `plain`; 4,003,536 B (3.82 MiB) encoded; 159.177060x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`

Snappy:
1. 16,585 B (16.20 KiB) compressed - `snappy` + `rle-dict`; 16,335 B (15.95 KiB) encoded; 241.486524x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 23,116 B (22.57 KiB) compressed - `snappy` + `delta-binary-packed`; 64,017 B (62.52 KiB) encoded; 173.258955x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 235,269 B (229.75 KiB) compressed - `snappy` + `plain`; 4,003,720 B (3.82 MiB) encoded; 17.023297x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`

## IsOldCounter (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/isoldcounter_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/isoldcounter_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/isoldcounter_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/isoldcounter_value_length.svg)


Compressed overall:
1. 4,228 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
2. 4,932 B (4.82 KiB) compressed - `zstd-3` + `delta-binary-packed`; 43,044 B (42.04 KiB) encoded; 812.054745x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 5,273 B (5.15 KiB) compressed - `snappy` + `rle-dict`; 5,041 B (4.92 KiB) encoded; 759.539920x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
4. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
5. 6,203 B (6.06 KiB) compressed - `snappy` + `delta-binary-packed`; 43,179 B (42.17 KiB) encoded; 645.664034x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 4,228 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
2. 4,932 B (4.82 KiB) compressed - `zstd-3` + `delta-binary-packed`; 43,044 B (42.04 KiB) encoded; 812.054745x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 5,273 B (5.15 KiB) compressed - `snappy` + `rle-dict`; 5,041 B (4.92 KiB) encoded; 759.539920x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 6,203 B (6.06 KiB) compressed - `snappy` + `delta-binary-packed`; 43,179 B (42.17 KiB) encoded; 645.664034x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 204,060 B (199.28 KiB) compressed - `snappy` + `plain`; 4,003,712 B (3.82 MiB) encoded; 19.626845x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## IsParameter (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/isparameter_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/isparameter_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/isparameter_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/isparameter_value_length.svg)


Compressed overall:
1. 4,228 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
2. 4,932 B (4.82 KiB) compressed - `zstd-3` + `delta-binary-packed`; 43,044 B (42.04 KiB) encoded; 812.054745x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 5,273 B (5.15 KiB) compressed - `snappy` + `rle-dict`; 5,041 B (4.92 KiB) encoded; 759.539920x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
4. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
5. 6,203 B (6.06 KiB) compressed - `snappy` + `delta-binary-packed`; 43,179 B (42.17 KiB) encoded; 645.664034x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 4,228 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
2. 4,932 B (4.82 KiB) compressed - `zstd-3` + `delta-binary-packed`; 43,044 B (42.04 KiB) encoded; 812.054745x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 5,273 B (5.15 KiB) compressed - `snappy` + `rle-dict`; 5,041 B (4.92 KiB) encoded; 759.539920x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 6,203 B (6.06 KiB) compressed - `snappy` + `delta-binary-packed`; 43,179 B (42.17 KiB) encoded; 645.664034x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 204,060 B (199.28 KiB) compressed - `snappy` + `plain`; 4,003,712 B (3.82 MiB) encoded; 19.626845x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## IsRefresh (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `2 / 2 / 2`
- Page cardinality per row group min/median/max of mins: `2 / 2 / 2`; of maxes: `2 / 2 / 2`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/isrefresh_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/isrefresh_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/isrefresh_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/isrefresh_value_length.svg)


Compressed overall:
1. 83,476 B (81.52 KiB) compressed - `zstd-3` + `rle-dict`; 92,576 B (90.41 KiB) encoded; 47.978473x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
2. 88,590 B (86.51 KiB) compressed - `snappy` + `rle-dict`; 92,617 B (90.45 KiB) encoded; 45.208838x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
3. 98,679 B (96.37 KiB) compressed - `zstd-3` + `delta-binary-packed`; 234,085 B (228.60 KiB) encoded; 40.586660x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
4. 137,630 B (134.40 KiB) compressed - `snappy` + `delta-binary-packed`; 233,923 B (228.44 KiB) encoded; 29.100131x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
5. 178,675 B (174.49 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 22.415285x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

ZSTD:
1. 83,476 B (81.52 KiB) compressed - `zstd-3` + `rle-dict`; 92,576 B (90.41 KiB) encoded; 47.978473x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
2. 98,679 B (96.37 KiB) compressed - `zstd-3` + `delta-binary-packed`; 234,085 B (228.60 KiB) encoded; 40.586660x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 178,675 B (174.49 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 22.415285x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 88,590 B (86.51 KiB) compressed - `snappy` + `rle-dict`; 92,617 B (90.45 KiB) encoded; 45.208838x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
2. 137,630 B (134.40 KiB) compressed - `snappy` + `delta-binary-packed`; 233,923 B (228.44 KiB) encoded; 29.100131x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 491,817 B (480.29 KiB) compressed - `snappy` + `plain`; 4,003,743 B (3.82 MiB) encoded; 8.143376x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

## JavaEnable (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `2 / 2 / 2`
- Page cardinality per row group min/median/max of mins: `2 / 2 / 2`; of maxes: `2 / 2 / 2`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/javaenable_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/javaenable_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/javaenable_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/javaenable_value_length.svg)


Compressed overall:
1. 50,727 B (49.54 KiB) compressed - `zstd-3` + `rle-dict`; 79,757 B (77.89 KiB) encoded; 78.953023x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 63,363 B (61.88 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 63.208024x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 67,923 B (66.33 KiB) compressed - `zstd-3` + `delta-binary-packed`; 197,664 B (193.03 KiB) encoded; 58.964563x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 69,196 B (67.57 KiB) compressed - `snappy` + `rle-dict`; 79,868 B (78.00 KiB) encoded; 57.879791x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
5. 90,773 B (88.65 KiB) compressed - `snappy` + `delta-binary-packed`; 198,065 B (193.42 KiB) encoded; 44.121600x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 50,727 B (49.54 KiB) compressed - `zstd-3` + `rle-dict`; 79,757 B (77.89 KiB) encoded; 78.953023x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 63,363 B (61.88 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 63.208024x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 67,923 B (66.33 KiB) compressed - `zstd-3` + `delta-binary-packed`; 197,664 B (193.03 KiB) encoded; 58.964563x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 69,196 B (67.57 KiB) compressed - `snappy` + `rle-dict`; 79,868 B (78.00 KiB) encoded; 57.879791x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 90,773 B (88.65 KiB) compressed - `snappy` + `delta-binary-packed`; 198,065 B (193.42 KiB) encoded; 44.121600x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`
3. 261,945 B (255.81 KiB) compressed - `snappy` + `plain`; 4,003,714 B (3.82 MiB) encoded; 15.289660x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

## JavascriptEnable (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 2 / 2`
- Page cardinality per row group min/median/max of mins: `1 / 2 / 2`; of maxes: `1 / 2 / 2`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/javascriptenable_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/javascriptenable_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/javascriptenable_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/javascriptenable_value_length.svg)


Compressed overall:
1. 6,485 B (6.33 KiB) compressed - `zstd-3` + `plain`; 4,003,543 B (3.82 MiB) encoded; 617.586893x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
2. 6,690 B (6.53 KiB) compressed - `snappy` + `rle-dict`; 6,472 B (6.32 KiB) encoded; 598.662332x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 7,387 B (7.21 KiB) compressed - `zstd-3` + `rle-dict`; 6,361 B (6.21 KiB) encoded; 542.175579x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
4. 8,167 B (7.98 KiB) compressed - `zstd-3` + `delta-binary-packed`; 47,574 B (46.46 KiB) encoded; 490.394392x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-delta-binary-packed`
5. 9,810 B (9.58 KiB) compressed - `snappy` + `delta-binary-packed`; 47,724 B (46.61 KiB) encoded; 408.262080x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-delta-binary-packed`

ZSTD:
1. 6,485 B (6.33 KiB) compressed - `zstd-3` + `plain`; 4,003,543 B (3.82 MiB) encoded; 617.586893x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
2. 7,387 B (7.21 KiB) compressed - `zstd-3` + `rle-dict`; 6,361 B (6.21 KiB) encoded; 542.175579x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
3. 8,167 B (7.98 KiB) compressed - `zstd-3` + `delta-binary-packed`; 47,574 B (46.46 KiB) encoded; 490.394392x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-delta-binary-packed`

Snappy:
1. 6,690 B (6.53 KiB) compressed - `snappy` + `rle-dict`; 6,472 B (6.32 KiB) encoded; 598.662332x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 9,810 B (9.58 KiB) compressed - `snappy` + `delta-binary-packed`; 47,724 B (46.61 KiB) encoded; 408.262080x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 204,912 B (200.11 KiB) compressed - `snappy` + `plain`; 4,003,774 B (3.82 MiB) encoded; 19.545224x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`

## LocalEventTime (timestamp_millis)

Column shape stats:
- Parquet type: `TIMESTAMP(isAdjustedToUTC=true,unit=MILLIS)`; physical type: `INT64`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `5,968 / 6,254 / 13,047`
- Page cardinality per row group min/median/max of mins: `5,968 / 6,254 / 13,047`; of maxes: `5,968 / 6,254 / 13,047`
- Value length min/median/max: `8 / 8 / 8` bytes
- Value length per row group min/median/max of mins: `8 / 8 / 8`; of maxes: `8 / 8 / 8`

![Row-group cardinality](column_shape_stats/images/localeventtime_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/localeventtime_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/localeventtime_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/localeventtime_value_length.svg)


Compressed overall:
1. 2,517,265 B (2.40 MiB) compressed - `zstd-3` + `plain`; 8,004,711 B (7.63 MiB) encoded; 3.180562x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 2,893,866 B (2.76 MiB) compressed - `zstd-3` + `delta-binary-packed`; 3,381,185 B (3.22 MiB) encoded; 2.766651x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed`
3. 3,042,877 B (2.90 MiB) compressed - `snappy` + `delta-binary-packed`; 3,384,399 B (3.23 MiB) encoded; 2.631167x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-rle-dict-ts-delta-binary-packed`
4. 4,023,316 B (3.84 MiB) compressed - `zstd-3` + `rle-dict`; 7,355,383 B (7.01 MiB) encoded; 1.989980x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict`
5. 4,283,734 B (4.09 MiB) compressed - `snappy` + `plain`; 8,004,716 B (7.63 MiB) encoded; 1.869004x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 2,517,265 B (2.40 MiB) compressed - `zstd-3` + `plain`; 8,004,711 B (7.63 MiB) encoded; 3.180562x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 2,893,866 B (2.76 MiB) compressed - `zstd-3` + `delta-binary-packed`; 3,381,185 B (3.22 MiB) encoded; 2.766651x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed`
3. 4,023,316 B (3.84 MiB) compressed - `zstd-3` + `rle-dict`; 7,355,383 B (7.01 MiB) encoded; 1.989980x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-rle-dict`

Snappy:
1. 3,042,877 B (2.90 MiB) compressed - `snappy` + `delta-binary-packed`; 3,384,399 B (3.23 MiB) encoded; 2.631167x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-rle-dict-ts-delta-binary-packed`
2. 4,283,734 B (4.09 MiB) compressed - `snappy` + `plain`; 8,004,716 B (7.63 MiB) encoded; 1.869004x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 5,514,562 B (5.26 MiB) compressed - `snappy` + `rle-dict`; 7,339,519 B (7.00 MiB) encoded; 1.451850x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-rle-dict-date-rle-dict-ts-rle-dict`

## MobilePhone (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `3 / 7 / 11`
- Page cardinality per row group min/median/max of mins: `3 / 7 / 11`; of maxes: `3 / 7 / 11`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/mobilephone_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/mobilephone_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/mobilephone_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/mobilephone_value_length.svg)


Compressed overall:
1. 22,463 B (21.94 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 178.295419x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed`
2. 23,572 B (23.02 KiB) compressed - `zstd-3` + `rle-dict`; 34,399 B (33.59 KiB) encoded; 169.907093x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 29,688 B (28.99 KiB) compressed - `snappy` + `rle-dict`; 34,623 B (33.81 KiB) encoded; 134.904675x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 39,491 B (38.57 KiB) compressed - `zstd-3` + `delta-binary-packed`; 144,598 B (141.21 KiB) encoded; 101.416779x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
5. 47,928 B (46.80 KiB) compressed - `snappy` + `delta-binary-packed`; 144,874 B (141.48 KiB) encoded; 83.563887x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 22,463 B (21.94 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 178.295419x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed`
2. 23,572 B (23.02 KiB) compressed - `zstd-3` + `rle-dict`; 34,399 B (33.59 KiB) encoded; 169.907093x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 39,491 B (38.57 KiB) compressed - `zstd-3` + `delta-binary-packed`; 144,598 B (141.21 KiB) encoded; 101.416779x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

Snappy:
1. 29,688 B (28.99 KiB) compressed - `snappy` + `rle-dict`; 34,623 B (33.81 KiB) encoded; 134.904675x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 47,928 B (46.80 KiB) compressed - `snappy` + `delta-binary-packed`; 144,874 B (141.48 KiB) encoded; 83.563887x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-plain`
3. 215,287 B (210.24 KiB) compressed - `snappy` + `plain`; 4,003,712 B (3.82 MiB) encoded; 18.603306x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`

## MobilePhoneModel (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `2 / 4 / 6`
- Page cardinality per row group min/median/max of mins: `2 / 4 / 6`; of maxes: `2 / 4 / 6`
- Value length min/median/max: `0 / 0 / 17` bytes
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `4 / 6 / 17`

![Row-group cardinality](column_shape_stats/images/mobilephonemodel_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/mobilephonemodel_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/mobilephonemodel_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/mobilephonemodel_value_length.svg)


Compressed overall:
1. 20,027 B (19.56 KiB) compressed - `zstd-3` + `rle-dict`; 25,102 B (24.51 KiB) encoded; 204.035702x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`
2. 22,555 B (22.03 KiB) compressed - `zstd-3` + `plain`; 4,084,876 B (3.90 MiB) encoded; 181.167058x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed`
3. 22,980 B (22.44 KiB) compressed - `snappy` + `rle-dict`; 24,988 B (24.40 KiB) encoded; 177.816493x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-delta-binary-packed`
4. 30,616 B (29.90 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 231,724 B (226.29 KiB) encoded; 133.466913x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
5. 40,717 B (39.76 KiB) compressed - `zstd-3` + `delta-byte-array`; 282,592 B (275.97 KiB) encoded; 100.356681x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

ZSTD:
1. 20,027 B (19.56 KiB) compressed - `zstd-3` + `rle-dict`; 25,102 B (24.51 KiB) encoded; 204.035702x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`
2. 22,555 B (22.03 KiB) compressed - `zstd-3` + `plain`; 4,084,876 B (3.90 MiB) encoded; 181.167058x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed`
3. 30,616 B (29.90 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 231,724 B (226.29 KiB) encoded; 133.466913x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
4. 40,717 B (39.76 KiB) compressed - `zstd-3` + `delta-byte-array`; 282,592 B (275.97 KiB) encoded; 100.356681x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 22,980 B (22.44 KiB) compressed - `snappy` + `rle-dict`; 24,988 B (24.40 KiB) encoded; 177.816493x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-delta-binary-packed`
2. 42,340 B (41.35 KiB) compressed - `snappy` + `delta-length-byte-array`; 231,343 B (225.92 KiB) encoded; 96.509754x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 56,742 B (55.41 KiB) compressed - `snappy` + `delta-byte-array`; 281,770 B (275.17 KiB) encoded; 72.014081x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
4. 219,863 B (214.71 KiB) compressed - `snappy` + `plain`; 4,085,062 B (3.90 MiB) encoded; 18.585314x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-delta-binary-packed`

## NetMajor (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `3 / 4 / 4`
- Page cardinality per row group min/median/max of mins: `3 / 4 / 4`; of maxes: `3 / 4 / 4`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/netmajor_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/netmajor_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/netmajor_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/netmajor_value_length.svg)


Compressed overall:
1. 25,268 B (24.68 KiB) compressed - `zstd-3` + `rle-dict`; 36,499 B (35.64 KiB) encoded; 158.502929x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 26,145 B (25.53 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 153.186154x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-plain`
3. 31,965 B (31.22 KiB) compressed - `snappy` + `rle-dict`; 36,717 B (35.86 KiB) encoded; 125.294916x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
4. 42,172 B (41.18 KiB) compressed - `zstd-3` + `delta-binary-packed`; 159,938 B (156.19 KiB) encoded; 94.969458x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain`
5. 51,618 B (50.41 KiB) compressed - `snappy` + `delta-binary-packed`; 160,393 B (156.63 KiB) encoded; 77.590220x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 25,268 B (24.68 KiB) compressed - `zstd-3` + `rle-dict`; 36,499 B (35.64 KiB) encoded; 158.502929x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 26,145 B (25.53 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 153.186154x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-plain`
3. 42,172 B (41.18 KiB) compressed - `zstd-3` + `delta-binary-packed`; 159,938 B (156.19 KiB) encoded; 94.969458x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain`

Snappy:
1. 31,965 B (31.22 KiB) compressed - `snappy` + `rle-dict`; 36,717 B (35.86 KiB) encoded; 125.294916x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 51,618 B (50.41 KiB) compressed - `snappy` + `delta-binary-packed`; 160,393 B (156.63 KiB) encoded; 77.590220x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 219,057 B (213.92 KiB) compressed - `snappy` + `plain`; 4,003,711 B (3.82 MiB) encoded; 18.283150x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## NetMinor (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `2 / 3 / 3`
- Page cardinality per row group min/median/max of mins: `2 / 3 / 3`; of maxes: `2 / 3 / 3`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/netminor_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/netminor_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/netminor_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/netminor_value_length.svg)


Compressed overall:
1. 23,065 B (22.52 KiB) compressed - `zstd-3` + `rle-dict`; 33,114 B (32.34 KiB) encoded; 173.641708x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 24,688 B (24.11 KiB) compressed - `zstd-3` + `plain`; 4,003,585 B (3.82 MiB) encoded; 162.226426x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 28,744 B (28.07 KiB) compressed - `snappy` + `rle-dict`; 33,059 B (32.28 KiB) encoded; 139.335026x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
4. 36,419 B (35.57 KiB) compressed - `zstd-3` + `delta-binary-packed`; 197,975 B (193.33 KiB) encoded; 109.971334x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
5. 47,620 B (46.50 KiB) compressed - `snappy` + `delta-binary-packed`; 197,804 B (193.17 KiB) encoded; 84.104284x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 23,065 B (22.52 KiB) compressed - `zstd-3` + `rle-dict`; 33,114 B (32.34 KiB) encoded; 173.641708x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 24,688 B (24.11 KiB) compressed - `zstd-3` + `plain`; 4,003,585 B (3.82 MiB) encoded; 162.226426x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 36,419 B (35.57 KiB) compressed - `zstd-3` + `delta-binary-packed`; 197,975 B (193.33 KiB) encoded; 109.971334x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 28,744 B (28.07 KiB) compressed - `snappy` + `rle-dict`; 33,059 B (32.28 KiB) encoded; 139.335026x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
2. 47,620 B (46.50 KiB) compressed - `snappy` + `delta-binary-packed`; 197,804 B (193.17 KiB) encoded; 84.104284x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`
3. 217,211 B (212.12 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 18.438504x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## OS (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `16 / 23 / 31`
- Page cardinality per row group min/median/max of mins: `16 / 23 / 31`; of maxes: `16 / 23 / 31`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/os_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/os_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/os_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/os_value_length.svg)


Compressed overall:
1. 105,900 B (103.42 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 37.819160x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain`
2. 119,893 B (117.08 KiB) compressed - `zstd-3` + `rle-dict`; 229,346 B (223.97 KiB) encoded; 33.405195x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 148,507 B (145.03 KiB) compressed - `snappy` + `rle-dict`; 230,224 B (224.83 KiB) encoded; 26.968756x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 215,676 B (210.62 KiB) compressed - `zstd-3` + `delta-binary-packed`; 846,563 B (826.72 KiB) encoded; 18.569748x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
5. 263,942 B (257.76 KiB) compressed - `snappy` + `delta-binary-packed`; 845,066 B (825.26 KiB) encoded; 15.173974x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 105,900 B (103.42 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 37.819160x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain`
2. 119,893 B (117.08 KiB) compressed - `zstd-3` + `rle-dict`; 229,346 B (223.97 KiB) encoded; 33.405195x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 215,676 B (210.62 KiB) compressed - `zstd-3` + `delta-binary-packed`; 846,563 B (826.72 KiB) encoded; 18.569748x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 148,507 B (145.03 KiB) compressed - `snappy` + `rle-dict`; 230,224 B (224.83 KiB) encoded; 26.968756x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 263,942 B (257.76 KiB) compressed - `snappy` + `delta-binary-packed`; 845,066 B (825.26 KiB) encoded; 15.173974x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 317,972 B (310.52 KiB) compressed - `snappy` + `plain`; 4,003,710 B (3.82 MiB) encoded; 12.595603x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## OpenerName (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/openername_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/openername_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/openername_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/openername_value_length.svg)


Compressed overall:
1. 4,232 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,531 B (3.82 MiB) encoded; 946.371928x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 4,987 B (4.87 KiB) compressed - `zstd-3` + `delta-binary-packed`; 43,042 B (42.03 KiB) encoded; 803.097253x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-delta-binary-packed`
3. 5,273 B (5.15 KiB) compressed - `snappy` + `rle-dict`; 5,041 B (4.92 KiB) encoded; 759.538403x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
4. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.300586x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
5. 6,253 B (6.11 KiB) compressed - `snappy` + `delta-binary-packed`; 43,182 B (42.17 KiB) encoded; 640.499920x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 4,232 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,531 B (3.82 MiB) encoded; 946.371928x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 4,987 B (4.87 KiB) compressed - `zstd-3` + `delta-binary-packed`; 43,042 B (42.03 KiB) encoded; 803.097253x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-delta-binary-packed`
3. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.300586x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 5,273 B (5.15 KiB) compressed - `snappy` + `rle-dict`; 5,041 B (4.92 KiB) encoded; 759.538403x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 6,253 B (6.11 KiB) compressed - `snappy` + `delta-binary-packed`; 43,182 B (42.17 KiB) encoded; 640.499920x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-plain`
3. 204,061 B (199.28 KiB) compressed - `snappy` + `plain`; 4,003,712 B (3.82 MiB) encoded; 19.626710x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`

## OpenstatAdID (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 9 / 18`
- Page cardinality per row group min/median/max of mins: `1 / 9 / 18`; of maxes: `1 / 9 / 18`
- Value length min/median/max: `0 / 0 / 22` bytes
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 13 / 22`

![Row-group cardinality](column_shape_stats/images/openstatadid_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/openstatadid_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/openstatadid_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/openstatadid_value_length.svg)


Compressed overall:
1. 18,284 B (17.86 KiB) compressed - `zstd-3` + `rle-dict`; 22,605 B (22.08 KiB) encoded; 220.437322x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 18,939 B (18.50 KiB) compressed - `zstd-3` + `plain`; 4,029,016 B (3.84 MiB) encoded; 212.813559x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 20,798 B (20.31 KiB) compressed - `snappy` + `rle-dict`; 22,683 B (22.15 KiB) encoded; 193.791518x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-delta-binary-packed`
4. 28,422 B (27.76 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 130,090 B (127.04 KiB) encoded; 141.808318x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
5. 33,463 B (32.68 KiB) compressed - `zstd-3` + `delta-byte-array`; 210,199 B (205.27 KiB) encoded; 120.445746x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`

ZSTD:
1. 18,284 B (17.86 KiB) compressed - `zstd-3` + `rle-dict`; 22,605 B (22.08 KiB) encoded; 220.437322x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 18,939 B (18.50 KiB) compressed - `zstd-3` + `plain`; 4,029,016 B (3.84 MiB) encoded; 212.813559x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 28,422 B (27.76 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 130,090 B (127.04 KiB) encoded; 141.808318x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
4. 33,463 B (32.68 KiB) compressed - `zstd-3` + `delta-byte-array`; 210,199 B (205.27 KiB) encoded; 120.445746x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`

Snappy:
1. 20,798 B (20.31 KiB) compressed - `snappy` + `rle-dict`; 22,683 B (22.15 KiB) encoded; 193.791518x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-delta-binary-packed`
2. 34,325 B (33.52 KiB) compressed - `snappy` + `delta-length-byte-array`; 130,064 B (127.02 KiB) encoded; 117.421005x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 44,359 B (43.32 KiB) compressed - `snappy` + `delta-byte-array`; 210,052 B (205.13 KiB) encoded; 90.860389x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`
4. 218,877 B (213.75 KiB) compressed - `snappy` + `plain`; 4,029,158 B (3.84 MiB) encoded; 18.414342x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-rle-dict-ts-delta-binary-packed`

## OpenstatCampaignID (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 4 / 6`
- Page cardinality per row group min/median/max of mins: `1 / 4 / 6`; of maxes: `1 / 4 / 6`
- Value length min/median/max: `0 / 0 / 12` bytes
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 10 / 12`

![Row-group cardinality](column_shape_stats/images/openstatcampaignid_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/openstatcampaignid_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/openstatcampaignid_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/openstatcampaignid_value_length.svg)


Compressed overall:
1. 15,591 B (15.23 KiB) compressed - `zstd-3` + `rle-dict`; 16,315 B (15.93 KiB) encoded; 258.295042x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 15,794 B (15.42 KiB) compressed - `snappy` + `rle-dict`; 16,230 B (15.85 KiB) encoded; 254.975180x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-delta-binary-packed`
3. 16,486 B (16.10 KiB) compressed - `zstd-3` + `plain`; 4,025,591 B (3.84 MiB) encoded; 244.272595x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-rle-dict`
4. 23,166 B (22.62 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 122,625 B (119.75 KiB) encoded; 173.835708x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
5. 28,149 B (27.49 KiB) compressed - `zstd-3` + `delta-byte-array`; 197,351 B (192.73 KiB) encoded; 143.062915x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`

ZSTD:
1. 15,591 B (15.23 KiB) compressed - `zstd-3` + `rle-dict`; 16,315 B (15.93 KiB) encoded; 258.295042x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 16,486 B (16.10 KiB) compressed - `zstd-3` + `plain`; 4,025,591 B (3.84 MiB) encoded; 244.272595x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-rle-dict`
3. 23,166 B (22.62 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 122,625 B (119.75 KiB) encoded; 173.835708x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 28,149 B (27.49 KiB) compressed - `zstd-3` + `delta-byte-array`; 197,351 B (192.73 KiB) encoded; 143.062915x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`

Snappy:
1. 15,794 B (15.42 KiB) compressed - `snappy` + `rle-dict`; 16,230 B (15.85 KiB) encoded; 254.975180x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-delta-binary-packed`
2. 28,470 B (27.80 KiB) compressed - `snappy` + `delta-length-byte-array`; 122,381 B (119.51 KiB) encoded; 141.449877x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 38,441 B (37.54 KiB) compressed - `snappy` + `delta-byte-array`; 198,099 B (193.46 KiB) encoded; 104.759970x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`
4. 214,606 B (209.58 KiB) compressed - `snappy` + `plain`; 4,025,810 B (3.84 MiB) encoded; 18.764983x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-rle-dict-ts-delta-binary-packed`

## OpenstatServiceName (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 2 / 6`
- Page cardinality per row group min/median/max of mins: `1 / 2 / 6`; of maxes: `1 / 2 / 6`
- Value length min/median/max: `0 / 0 / 16` bytes
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 16 / 16`

![Row-group cardinality](column_shape_stats/images/openstatservicename_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/openstatservicename_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/openstatservicename_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/openstatservicename_value_length.svg)


Compressed overall:
1. 17,512 B (17.10 KiB) compressed - `zstd-3` + `rle-dict`; 18,472 B (18.04 KiB) encoded; 232.086969x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed`
2. 17,876 B (17.46 KiB) compressed - `snappy` + `rle-dict`; 18,558 B (18.12 KiB) encoded; 227.361099x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
3. 19,169 B (18.72 KiB) compressed - `zstd-3` + `plain`; 4,062,525 B (3.87 MiB) encoded; 212.024988x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed`
4. 27,297 B (26.66 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 215,907 B (210.85 KiB) encoded; 148.892076x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
5. 33,746 B (32.96 KiB) compressed - `zstd-3` + `delta-byte-array`; 301,258 B (294.20 KiB) encoded; 120.438185x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`

ZSTD:
1. 17,512 B (17.10 KiB) compressed - `zstd-3` + `rle-dict`; 18,472 B (18.04 KiB) encoded; 232.086969x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed`
2. 19,169 B (18.72 KiB) compressed - `zstd-3` + `plain`; 4,062,525 B (3.87 MiB) encoded; 212.024988x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed`
3. 27,297 B (26.66 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 215,907 B (210.85 KiB) encoded; 148.892076x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
4. 33,746 B (32.96 KiB) compressed - `zstd-3` + `delta-byte-array`; 301,258 B (294.20 KiB) encoded; 120.438185x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`

Snappy:
1. 17,876 B (17.46 KiB) compressed - `snappy` + `rle-dict`; 18,558 B (18.12 KiB) encoded; 227.361099x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 36,802 B (35.94 KiB) compressed - `snappy` + `delta-length-byte-array`; 215,782 B (210.72 KiB) encoded; 110.437123x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
3. 47,779 B (46.66 KiB) compressed - `snappy` + `delta-byte-array`; 301,351 B (294.29 KiB) encoded; 85.064715x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 218,413 B (213.29 KiB) compressed - `snappy` + `plain`; 4,062,767 B (3.87 MiB) encoded; 18.608357x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`

## OpenstatSourceID (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 3 / 7`
- Page cardinality per row group min/median/max of mins: `1 / 3 / 7`; of maxes: `1 / 3 / 7`
- Value length min/median/max: `0 / 0 / 31` bytes
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 20 / 31`

![Row-group cardinality](column_shape_stats/images/openstatsourceid_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/openstatsourceid_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/openstatsourceid_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/openstatsourceid_value_length.svg)


Compressed overall:
1. 12,812 B (12.51 KiB) compressed - `zstd-3` + `rle-dict`; 14,560 B (14.22 KiB) encoded; 316.374337x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 13,231 B (12.92 KiB) compressed - `zstd-3` + `plain`; 4,051,710 B (3.86 MiB) encoded; 306.355378x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
3. 13,422 B (13.11 KiB) compressed - `snappy` + `rle-dict`; 14,793 B (14.45 KiB) encoded; 301.995828x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
4. 18,296 B (17.87 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 155,189 B (151.55 KiB) encoded; 221.545037x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict`
5. 23,129 B (22.59 KiB) compressed - `zstd-3` + `delta-byte-array`; 231,933 B (226.50 KiB) encoded; 175.251329x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

ZSTD:
1. 12,812 B (12.51 KiB) compressed - `zstd-3` + `rle-dict`; 14,560 B (14.22 KiB) encoded; 316.374337x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 13,231 B (12.92 KiB) compressed - `zstd-3` + `plain`; 4,051,710 B (3.86 MiB) encoded; 306.355378x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
3. 18,296 B (17.87 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 155,189 B (151.55 KiB) encoded; 221.545037x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict`
4. 23,129 B (22.59 KiB) compressed - `zstd-3` + `delta-byte-array`; 231,933 B (226.50 KiB) encoded; 175.251329x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 13,422 B (13.11 KiB) compressed - `snappy` + `rle-dict`; 14,793 B (14.45 KiB) encoded; 301.995828x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
2. 24,350 B (23.78 KiB) compressed - `snappy` + `delta-length-byte-array`; 155,292 B (151.65 KiB) encoded; 166.463573x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-delta-binary-packed`
3. 33,748 B (32.96 KiB) compressed - `snappy` + `delta-byte-array`; 232,759 B (227.30 KiB) encoded; 120.107503x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`
4. 214,290 B (209.27 KiB) compressed - `snappy` + `plain`; 4,051,956 B (3.86 MiB) encoded; 18.915432x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-delta-binary-packed-ts-delta-binary-packed`

## OriginalURL (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `173`
- Row-group cardinality min/median/max: `95 / 227 / 6,169`
- Page cardinality per row group min/median/max of mins: `12 / 183 / 850`; of maxes: `95 / 227 / 1,065`
- Value length min/median/max: `0 / 138.50 / 3,723` bytes
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `357 / 517 / 3,723`

![Row-group cardinality](column_shape_stats/images/originalurl_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/originalurl_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/originalurl_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/originalurl_value_length.svg)


Compressed overall:
1. 4,878,331 B (4.65 MiB) compressed - `zstd-3` + `rle-dict`; 21,272,072 B (20.29 MiB) encoded; 6.532149x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 5,313,092 B (5.07 MiB) compressed - `zstd-3` + `plain`; 31,855,165 B (30.38 MiB) encoded; 5.997634x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain`
3. 5,585,217 B (5.33 MiB) compressed - `zstd-3` + `delta-byte-array`; 21,049,703 B (20.07 MiB) encoded; 5.705415x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
4. 5,605,309 B (5.35 MiB) compressed - `zstd-3` + `delta-length-byte-array`; 28,790,277 B (27.46 MiB) encoded; 5.684965x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
5. 6,156,851 B (5.87 MiB) compressed - `snappy` + `rle-dict`; 21,269,469 B (20.28 MiB) encoded; 5.175695x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`

ZSTD:
1. 4,878,331 B (4.65 MiB) compressed - `zstd-3` + `rle-dict`; 21,272,072 B (20.29 MiB) encoded; 6.532149x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 5,313,092 B (5.07 MiB) compressed - `zstd-3` + `plain`; 31,855,165 B (30.38 MiB) encoded; 5.997634x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain`
3. 5,585,217 B (5.33 MiB) compressed - `zstd-3` + `delta-byte-array`; 21,049,703 B (20.07 MiB) encoded; 5.705415x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
4. 5,605,309 B (5.35 MiB) compressed - `zstd-3` + `delta-length-byte-array`; 28,790,277 B (27.46 MiB) encoded; 5.684965x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 6,156,851 B (5.87 MiB) compressed - `snappy` + `rle-dict`; 21,269,469 B (20.28 MiB) encoded; 5.175695x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`
2. 6,746,526 B (6.43 MiB) compressed - `snappy` + `delta-byte-array`; 21,052,485 B (20.08 MiB) encoded; 4.723317x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 7,036,093 B (6.71 MiB) compressed - `snappy` + `plain`; 31,860,219 B (30.38 MiB) encoded; 4.528931x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-delta-binary-packed-ts-rle-dict`
4. 7,041,489 B (6.72 MiB) compressed - `snappy` + `delta-length-byte-array`; 28,790,269 B (27.46 MiB) encoded; 4.525461x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`

## PageCharset (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `122`
- Row-group cardinality min/median/max: `1 / 2 / 3`
- Page cardinality per row group min/median/max of mins: `1 / 2 / 3`; of maxes: `1 / 2 / 3`
- Value length min/median/max: `0 / 20 / 20` bytes
- Value length per row group min/median/max of mins: `0 / 0 / 20`; of maxes: `7 / 20 / 20`

![Row-group cardinality](column_shape_stats/images/pagecharset_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/pagecharset_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/pagecharset_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/pagecharset_value_length.svg)


Compressed overall:
1. 9,045 B (8.83 KiB) compressed - `snappy` + `rle-dict`; 8,865 B (8.66 KiB) encoded; 1945.711443x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-delta-binary-packed`
2. 9,902 B (9.67 KiB) compressed - `zstd-3` + `rle-dict`; 8,838 B (8.63 KiB) encoded; 1777.313674x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed`
3. 14,575 B (14.23 KiB) compressed - `zstd-3` + `plain`; 17,594,795 B (16.78 MiB) encoded; 1207.475815x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed`
4. 18,662 B (18.22 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 13,656,839 B (13.02 MiB) encoded; 943.037188x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
5. 21,415 B (20.91 KiB) compressed - `zstd-3` + `delta-byte-array`; 140,380 B (137.09 KiB) encoded; 821.805277x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`

ZSTD:
1. 9,902 B (9.67 KiB) compressed - `zstd-3` + `rle-dict`; 8,838 B (8.63 KiB) encoded; 1777.313674x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed`
2. 14,575 B (14.23 KiB) compressed - `zstd-3` + `plain`; 17,594,795 B (16.78 MiB) encoded; 1207.475815x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed`
3. 18,662 B (18.22 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 13,656,839 B (13.02 MiB) encoded; 943.037188x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
4. 21,415 B (20.91 KiB) compressed - `zstd-3` + `delta-byte-array`; 140,380 B (137.09 KiB) encoded; 821.805277x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`

Snappy:
1. 9,045 B (8.83 KiB) compressed - `snappy` + `rle-dict`; 8,865 B (8.66 KiB) encoded; 1945.711443x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-delta-binary-packed`
2. 26,461 B (25.84 KiB) compressed - `snappy` + `delta-byte-array`; 141,618 B (138.30 KiB) encoded; 665.090511x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 699,096 B (682.71 KiB) compressed - `snappy` + `delta-length-byte-array`; 13,657,540 B (13.02 MiB) encoded; 25.173882x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 913,453 B (892.04 KiB) compressed - `snappy` + `plain`; 17,595,554 B (16.78 MiB) encoded; 19.266410x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-rle-dict-ts-delta-binary-packed`

## ParamCurrency (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length min/median/max: `3 / 3 / 3` bytes
- Value length per row group min/median/max of mins: `3 / 3 / 3`; of maxes: `3 / 3 / 3`

![Row-group cardinality](column_shape_stats/images/paramcurrency_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/paramcurrency_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/paramcurrency_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/paramcurrency_value_length.svg)


Compressed overall:
1. 5,256 B (5.13 KiB) compressed - `zstd-3` + `plain`; 7,003,298 B (6.68 MiB) encoded; 1332.711758x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 5,393 B (5.27 KiB) compressed - `snappy` + `rle-dict`; 5,153 B (5.03 KiB) encoded; 1298.856481x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
3. 5,771 B (5.64 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 3,043,923 B (2.90 MiB) encoded; 1213.781494x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 6,132 B (5.99 KiB) compressed - `zstd-3` + `rle-dict`; 5,070 B (4.95 KiB) encoded; 1142.324364x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-plain`
5. 6,869 B (6.71 KiB) compressed - `zstd-3` + `delta-byte-array`; 86,017 B (84.00 KiB) encoded; 1019.760227x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`

ZSTD:
1. 5,256 B (5.13 KiB) compressed - `zstd-3` + `plain`; 7,003,298 B (6.68 MiB) encoded; 1332.711758x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 5,771 B (5.64 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 3,043,923 B (2.90 MiB) encoded; 1213.781494x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 6,132 B (5.99 KiB) compressed - `zstd-3` + `rle-dict`; 5,070 B (4.95 KiB) encoded; 1142.324364x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-plain`
4. 6,869 B (6.71 KiB) compressed - `zstd-3` + `delta-byte-array`; 86,017 B (84.00 KiB) encoded; 1019.760227x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`

Snappy:
1. 5,393 B (5.27 KiB) compressed - `snappy` + `rle-dict`; 5,153 B (5.03 KiB) encoded; 1298.856481x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
2. 10,240 B (10.00 KiB) compressed - `snappy` + `delta-byte-array`; 86,106 B (84.09 KiB) encoded; 684.055957x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 157,287 B (153.60 KiB) compressed - `snappy` + `delta-length-byte-array`; 3,044,410 B (2.90 MiB) encoded; 44.534723x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
4. 354,236 B (345.93 KiB) compressed - `snappy` + `plain`; 7,003,535 B (6.68 MiB) encoded; 19.774199x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-rle-dict-ts-plain`

## ParamCurrencyID (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/paramcurrencyid_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/paramcurrencyid_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/paramcurrencyid_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/paramcurrencyid_value_length.svg)


Compressed overall:
1. 4,228 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
2. 4,932 B (4.82 KiB) compressed - `zstd-3` + `delta-binary-packed`; 43,044 B (42.04 KiB) encoded; 812.054745x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 5,273 B (5.15 KiB) compressed - `snappy` + `rle-dict`; 5,041 B (4.92 KiB) encoded; 759.539920x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
4. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
5. 6,203 B (6.06 KiB) compressed - `snappy` + `delta-binary-packed`; 43,179 B (42.17 KiB) encoded; 645.664034x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 4,228 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
2. 4,932 B (4.82 KiB) compressed - `zstd-3` + `delta-binary-packed`; 43,044 B (42.04 KiB) encoded; 812.054745x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 5,273 B (5.15 KiB) compressed - `snappy` + `rle-dict`; 5,041 B (4.92 KiB) encoded; 759.539920x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 6,203 B (6.06 KiB) compressed - `snappy` + `delta-binary-packed`; 43,179 B (42.17 KiB) encoded; 645.664034x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 204,060 B (199.28 KiB) compressed - `snappy` + `plain`; 4,003,712 B (3.82 MiB) encoded; 19.626845x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## ParamOrderID (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length min/median/max: `0 / 0 / 0` bytes
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 0 / 0`

![Row-group cardinality](column_shape_stats/images/paramorderid_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/paramorderid_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/paramorderid_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/paramorderid_value_length.svg)


Compressed overall:
1. 2,848 B (2.78 KiB) compressed - `zstd-3` + `plain`; 4,002,159 B (3.82 MiB) encoded; 1405.603230x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
2. 3,567 B (3.48 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 41,679 B (40.70 KiB) encoded; 1122.275862x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 3,680 B (3.59 KiB) compressed - `zstd-3` + `delta-byte-array`; 81,200 B (79.30 KiB) encoded; 1087.814674x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 4,493 B (4.39 KiB) compressed - `snappy` + `rle-dict`; 4,253 B (4.15 KiB) encoded; 890.976630x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
5. 4,832 B (4.72 KiB) compressed - `snappy` + `delta-length-byte-array`; 41,830 B (40.85 KiB) encoded; 828.468129x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`

ZSTD:
1. 2,848 B (2.78 KiB) compressed - `zstd-3` + `plain`; 4,002,159 B (3.82 MiB) encoded; 1405.603230x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
2. 3,567 B (3.48 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 41,679 B (40.70 KiB) encoded; 1122.275862x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 3,680 B (3.59 KiB) compressed - `zstd-3` + `delta-byte-array`; 81,200 B (79.30 KiB) encoded; 1087.814674x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 5,247 B (5.12 KiB) compressed - `zstd-3` + `rle-dict`; 4,185 B (4.09 KiB) encoded; 762.942253x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-plain`

Snappy:
1. 4,493 B (4.39 KiB) compressed - `snappy` + `rle-dict`; 4,253 B (4.15 KiB) encoded; 890.976630x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
2. 4,832 B (4.72 KiB) compressed - `snappy` + `delta-length-byte-array`; 41,830 B (40.85 KiB) encoded; 828.468129x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
3. 6,836 B (6.68 KiB) compressed - `snappy` + `delta-byte-array`; 81,353 B (79.45 KiB) encoded; 585.599473x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 202,701 B (197.95 KiB) compressed - `snappy` + `plain`; 4,002,334 B (3.82 MiB) encoded; 19.749079x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-delta-binary-packed`

## ParamPrice (int64)

Column shape stats:
- Parquet type: `INT(64,true)`; physical type: `INT64`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length min/median/max: `8 / 8 / 8` bytes
- Value length per row group min/median/max of mins: `8 / 8 / 8`; of maxes: `8 / 8 / 8`

![Row-group cardinality](column_shape_stats/images/paramprice_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/paramprice_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/paramprice_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/paramprice_value_length.svg)


Compressed overall:
1. 5,675 B (5.54 KiB) compressed - `zstd-3` + `plain`; 8,004,440 B (7.63 MiB) encoded; 1410.804934x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
2. 5,844 B (5.71 KiB) compressed - `zstd-3` + `delta-binary-packed`; 43,956 B (42.93 KiB) encoded; 1370.006502x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 6,433 B (6.28 KiB) compressed - `snappy` + `rle-dict`; 6,201 B (6.06 KiB) encoded; 1244.569874x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
4. 7,115 B (6.95 KiB) compressed - `zstd-3` + `rle-dict`; 6,089 B (5.95 KiB) encoded; 1125.273085x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
5. 7,147 B (6.98 KiB) compressed - `snappy` + `delta-binary-packed`; 44,123 B (43.09 KiB) encoded; 1120.234784x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 5,675 B (5.54 KiB) compressed - `zstd-3` + `plain`; 8,004,440 B (7.63 MiB) encoded; 1410.804934x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
2. 5,844 B (5.71 KiB) compressed - `zstd-3` + `delta-binary-packed`; 43,956 B (42.93 KiB) encoded; 1370.006502x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 7,115 B (6.95 KiB) compressed - `zstd-3` + `rle-dict`; 6,089 B (5.95 KiB) encoded; 1125.273085x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 6,433 B (6.28 KiB) compressed - `snappy` + `rle-dict`; 6,201 B (6.06 KiB) encoded; 1244.569874x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 7,147 B (6.98 KiB) compressed - `snappy` + `delta-binary-packed`; 44,123 B (43.09 KiB) encoded; 1120.234784x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 405,010 B (395.52 KiB) compressed - `snappy` + `plain`; 8,004,654 B (7.63 MiB) encoded; 19.768198x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## Params (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length min/median/max: `0 / 0 / 0` bytes
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 0 / 0`

![Row-group cardinality](column_shape_stats/images/params_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/params_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/params_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/params_value_length.svg)


Compressed overall:
1. 2,848 B (2.78 KiB) compressed - `zstd-3` + `plain`; 4,002,159 B (3.82 MiB) encoded; 1405.603230x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
2. 3,567 B (3.48 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 41,679 B (40.70 KiB) encoded; 1122.275862x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 3,680 B (3.59 KiB) compressed - `zstd-3` + `delta-byte-array`; 81,200 B (79.30 KiB) encoded; 1087.814674x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 4,493 B (4.39 KiB) compressed - `snappy` + `rle-dict`; 4,253 B (4.15 KiB) encoded; 890.976630x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
5. 4,832 B (4.72 KiB) compressed - `snappy` + `delta-length-byte-array`; 41,830 B (40.85 KiB) encoded; 828.468129x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`

ZSTD:
1. 2,848 B (2.78 KiB) compressed - `zstd-3` + `plain`; 4,002,159 B (3.82 MiB) encoded; 1405.603230x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
2. 3,567 B (3.48 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 41,679 B (40.70 KiB) encoded; 1122.275862x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 3,680 B (3.59 KiB) compressed - `zstd-3` + `delta-byte-array`; 81,200 B (79.30 KiB) encoded; 1087.814674x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 5,247 B (5.12 KiB) compressed - `zstd-3` + `rle-dict`; 4,185 B (4.09 KiB) encoded; 762.942253x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-plain`

Snappy:
1. 4,493 B (4.39 KiB) compressed - `snappy` + `rle-dict`; 4,253 B (4.15 KiB) encoded; 890.976630x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
2. 4,832 B (4.72 KiB) compressed - `snappy` + `delta-length-byte-array`; 41,830 B (40.85 KiB) encoded; 828.468129x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
3. 6,836 B (6.68 KiB) compressed - `snappy` + `delta-byte-array`; 81,353 B (79.45 KiB) encoded; 585.599473x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 202,701 B (197.95 KiB) compressed - `snappy` + `plain`; 4,002,334 B (3.82 MiB) encoded; 19.749079x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-delta-binary-packed`

## Referer (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `349`
- Row-group cardinality min/median/max: `361 / 2,628 / 5,960`
- Page cardinality per row group min/median/max of mins: `1 / 670 / 1,130`; of maxes: `223 / 802 / 1,644`
- Value length min/median/max: `0 / 64 / 2,007` bytes
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `514 / 1,014 / 2,007`

![Row-group cardinality](column_shape_stats/images/referer_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/referer_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/referer_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/referer_value_length.svg)


Compressed overall:
1. 11,720,762 B (11.18 MiB) compressed - `zstd-3` + `rle-dict`; 34,380,236 B (32.79 MiB) encoded; 7.136665x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 14,212,015 B (13.55 MiB) compressed - `zstd-3` + `plain`; 83,646,116 B (79.77 MiB) encoded; 5.885665x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
3. 14,535,820 B (13.86 MiB) compressed - `zstd-3` + `delta-byte-array`; 38,986,977 B (37.18 MiB) encoded; 5.754554x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 14,794,287 B (14.11 MiB) compressed - `snappy` + `rle-dict`; 34,400,577 B (32.81 MiB) encoded; 5.654017x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`
5. 14,991,338 B (14.30 MiB) compressed - `zstd-3` + `delta-length-byte-array`; 80,834,647 B (77.09 MiB) encoded; 5.579699x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 11,720,762 B (11.18 MiB) compressed - `zstd-3` + `rle-dict`; 34,380,236 B (32.79 MiB) encoded; 7.136665x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 14,212,015 B (13.55 MiB) compressed - `zstd-3` + `plain`; 83,646,116 B (79.77 MiB) encoded; 5.885665x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
3. 14,535,820 B (13.86 MiB) compressed - `zstd-3` + `delta-byte-array`; 38,986,977 B (37.18 MiB) encoded; 5.754554x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 14,991,338 B (14.30 MiB) compressed - `zstd-3` + `delta-length-byte-array`; 80,834,647 B (77.09 MiB) encoded; 5.579699x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`

Snappy:
1. 14,794,287 B (14.11 MiB) compressed - `snappy` + `rle-dict`; 34,400,577 B (32.81 MiB) encoded; 5.654017x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`
2. 17,315,959 B (16.51 MiB) compressed - `snappy` + `delta-byte-array`; 38,988,769 B (37.18 MiB) encoded; 4.830640x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 19,041,603 B (18.16 MiB) compressed - `snappy` + `plain`; 83,646,065 B (79.77 MiB) encoded; 4.392863x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`
4. 19,455,042 B (18.55 MiB) compressed - `snappy` + `delta-length-byte-array`; 80,832,992 B (77.09 MiB) encoded; 4.299510x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`

## RefererCategoryID (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `12 / 35 / 79`
- Page cardinality per row group min/median/max of mins: `12 / 35 / 79`; of maxes: `12 / 35 / 79`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/referercategoryid_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/referercategoryid_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/referercategoryid_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/referercategoryid_value_length.svg)


Compressed overall:
1. 215,941 B (210.88 KiB) compressed - `zstd-3` + `rle-dict`; 516,015 B (503.92 KiB) encoded; 18.546978x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
2. 275,092 B (268.64 KiB) compressed - `zstd-3` + `plain`; 4,003,582 B (3.82 MiB) encoded; 14.558958x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed`
3. 280,720 B (274.14 KiB) compressed - `snappy` + `rle-dict`; 512,413 B (500.40 KiB) encoded; 14.267074x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 508,782 B (496.86 KiB) compressed - `snappy` + `plain`; 4,003,746 B (3.82 MiB) encoded; 7.871845x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
5. 520,118 B (507.93 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,862,796 B (1.78 MiB) encoded; 7.700278x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

ZSTD:
1. 215,941 B (210.88 KiB) compressed - `zstd-3` + `rle-dict`; 516,015 B (503.92 KiB) encoded; 18.546978x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
2. 275,092 B (268.64 KiB) compressed - `zstd-3` + `plain`; 4,003,582 B (3.82 MiB) encoded; 14.558958x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed`
3. 520,118 B (507.93 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,862,796 B (1.78 MiB) encoded; 7.700278x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 280,720 B (274.14 KiB) compressed - `snappy` + `rle-dict`; 512,413 B (500.40 KiB) encoded; 14.267074x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 508,782 B (496.86 KiB) compressed - `snappy` + `plain`; 4,003,746 B (3.82 MiB) encoded; 7.871845x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 589,994 B (576.17 KiB) compressed - `snappy` + `delta-binary-packed`; 1,862,388 B (1.78 MiB) encoded; 6.788294x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-plain`

## RefererHash (int64)

Column shape stats:
- Parquet type: `INT(64,true)`; physical type: `INT64`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `378 / 2,729 / 6,051`
- Page cardinality per row group min/median/max of mins: `378 / 2,729 / 6,051`; of maxes: `378 / 2,729 / 6,051`
- Value length min/median/max: `8 / 8 / 8` bytes
- Value length per row group min/median/max of mins: `8 / 8 / 8`; of maxes: `8 / 8 / 8`

![Row-group cardinality](column_shape_stats/images/refererhash_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/refererhash_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/refererhash_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/refererhash_value_length.svg)


Compressed overall:
1. 2,841,886 B (2.71 MiB) compressed - `zstd-3` + `plain`; 8,004,557 B (7.63 MiB) encoded; 2.817254x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
2. 3,502,907 B (3.34 MiB) compressed - `zstd-3` + `rle-dict`; 3,770,600 B (3.60 MiB) encoded; 2.285620x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
3. 3,517,169 B (3.35 MiB) compressed - `snappy` + `rle-dict`; 3,759,258 B (3.59 MiB) encoded; 2.276352x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
4. 3,637,773 B (3.47 MiB) compressed - `snappy` + `plain`; 8,004,877 B (7.63 MiB) encoded; 2.200884x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-plain`
5. 4,464,597 B (4.26 MiB) compressed - `zstd-3` + `delta-binary-packed`; 8,119,083 B (7.74 MiB) encoded; 1.793290x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

ZSTD:
1. 2,841,886 B (2.71 MiB) compressed - `zstd-3` + `plain`; 8,004,557 B (7.63 MiB) encoded; 2.817254x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
2. 3,502,907 B (3.34 MiB) compressed - `zstd-3` + `rle-dict`; 3,770,600 B (3.60 MiB) encoded; 2.285620x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
3. 4,464,597 B (4.26 MiB) compressed - `zstd-3` + `delta-binary-packed`; 8,119,083 B (7.74 MiB) encoded; 1.793290x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 3,517,169 B (3.35 MiB) compressed - `snappy` + `rle-dict`; 3,759,258 B (3.59 MiB) encoded; 2.276352x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
2. 3,637,773 B (3.47 MiB) compressed - `snappy` + `plain`; 8,004,877 B (7.63 MiB) encoded; 2.200884x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-rle-dict-date-delta-binary-packed-ts-plain`
3. 5,094,999 B (4.86 MiB) compressed - `snappy` + `delta-binary-packed`; 8,118,138 B (7.74 MiB) encoded; 1.571407x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

## RefererRegionID (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `7 / 17 / 33`
- Page cardinality per row group min/median/max of mins: `7 / 17 / 33`; of maxes: `7 / 17 / 33`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/refererregionid_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/refererregionid_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/refererregionid_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/refererregionid_value_length.svg)


Compressed overall:
1. 162,823 B (159.01 KiB) compressed - `zstd-3` + `rle-dict`; 378,845 B (369.97 KiB) encoded; 24.597569x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict`
2. 218,717 B (213.59 KiB) compressed - `snappy` + `rle-dict`; 377,417 B (368.57 KiB) encoded; 18.311562x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 231,215 B (225.80 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 17.321757x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
4. 335,977 B (328.10 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,368,810 B (1.31 MiB) encoded; 11.920608x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain`
5. 392,279 B (383.08 KiB) compressed - `snappy` + `delta-binary-packed`; 1,366,147 B (1.30 MiB) encoded; 10.209698x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 162,823 B (159.01 KiB) compressed - `zstd-3` + `rle-dict`; 378,845 B (369.97 KiB) encoded; 24.597569x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict`
2. 231,215 B (225.80 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 17.321757x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
3. 335,977 B (328.10 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,368,810 B (1.31 MiB) encoded; 11.920608x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain`

Snappy:
1. 218,717 B (213.59 KiB) compressed - `snappy` + `rle-dict`; 377,417 B (368.57 KiB) encoded; 18.311562x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
2. 392,279 B (383.08 KiB) compressed - `snappy` + `delta-binary-packed`; 1,366,147 B (1.30 MiB) encoded; 10.209698x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 457,896 B (447.16 KiB) compressed - `snappy` + `plain`; 4,003,737 B (3.82 MiB) encoded; 8.746637x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

## RegionID (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `112 / 149 / 275`
- Page cardinality per row group min/median/max of mins: `112 / 149 / 275`; of maxes: `112 / 149 / 275`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/regionid_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/regionid_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/regionid_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/regionid_value_length.svg)


Compressed overall:
1. 190,701 B (186.23 KiB) compressed - `zstd-3` + `plain`; 4,003,584 B (3.82 MiB) encoded; 21.001715x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`
2. 244,341 B (238.61 KiB) compressed - `zstd-3` + `rle-dict`; 435,033 B (424.84 KiB) encoded; 16.391224x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 286,974 B (280.25 KiB) compressed - `snappy` + `rle-dict`; 435,032 B (424.84 KiB) encoded; 13.956135x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
4. 396,306 B (387.02 KiB) compressed - `snappy` + `plain`; 4,003,714 B (3.82 MiB) encoded; 10.105948x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
5. 463,496 B (452.63 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,489,061 B (1.42 MiB) encoded; 8.640955x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`

ZSTD:
1. 190,701 B (186.23 KiB) compressed - `zstd-3` + `plain`; 4,003,584 B (3.82 MiB) encoded; 21.001715x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`
2. 244,341 B (238.61 KiB) compressed - `zstd-3` + `rle-dict`; 435,033 B (424.84 KiB) encoded; 16.391224x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 463,496 B (452.63 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,489,061 B (1.42 MiB) encoded; 8.640955x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 286,974 B (280.25 KiB) compressed - `snappy` + `rle-dict`; 435,032 B (424.84 KiB) encoded; 13.956135x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
2. 396,306 B (387.02 KiB) compressed - `snappy` + `plain`; 4,003,714 B (3.82 MiB) encoded; 10.105948x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 545,447 B (532.66 KiB) compressed - `snappy` + `delta-binary-packed`; 1,484,217 B (1.42 MiB) encoded; 7.342690x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## RemoteIP (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `508 / 851 / 1,951`
- Page cardinality per row group min/median/max of mins: `508 / 851 / 1,951`; of maxes: `508 / 851 / 1,951`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/remoteip_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/remoteip_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/remoteip_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/remoteip_value_length.svg)


Compressed overall:
1. 426,425 B (416.43 KiB) compressed - `zstd-3` + `plain`; 4,003,606 B (3.82 MiB) encoded; 9.392149x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`
2. 698,454 B (682.08 KiB) compressed - `zstd-3` + `rle-dict`; 927,182 B (905.45 KiB) encoded; 5.734160x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
3. 706,731 B (690.17 KiB) compressed - `snappy` + `plain`; 4,003,834 B (3.82 MiB) encoded; 5.667003x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`
4. 748,986 B (731.43 KiB) compressed - `snappy` + `rle-dict`; 927,799 B (906.05 KiB) encoded; 5.347292x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-delta-binary-packed-ts-rle-dict`
5. 926,924 B (905.20 KiB) compressed - `zstd-3` + `delta-binary-packed`; 3,793,493 B (3.62 MiB) encoded; 4.320793x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`

ZSTD:
1. 426,425 B (416.43 KiB) compressed - `zstd-3` + `plain`; 4,003,606 B (3.82 MiB) encoded; 9.392149x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`
2. 698,454 B (682.08 KiB) compressed - `zstd-3` + `rle-dict`; 927,182 B (905.45 KiB) encoded; 5.734160x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
3. 926,924 B (905.20 KiB) compressed - `zstd-3` + `delta-binary-packed`; 3,793,493 B (3.62 MiB) encoded; 4.320793x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`

Snappy:
1. 706,731 B (690.17 KiB) compressed - `snappy` + `plain`; 4,003,834 B (3.82 MiB) encoded; 5.667003x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-rle-dict`
2. 748,986 B (731.43 KiB) compressed - `snappy` + `rle-dict`; 927,799 B (906.05 KiB) encoded; 5.347292x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-delta-binary-packed-ts-rle-dict`
3. 1,109,977 B (1.06 MiB) compressed - `snappy` + `delta-binary-packed`; 3,800,823 B (3.62 MiB) encoded; 3.608225x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`

## ResolutionDepth (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `4 / 4 / 5`
- Page cardinality per row group min/median/max of mins: `4 / 4 / 5`; of maxes: `4 / 4 / 5`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/resolutiondepth_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/resolutiondepth_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/resolutiondepth_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/resolutiondepth_value_length.svg)


Compressed overall:
1. 61,927 B (60.48 KiB) compressed - `zstd-3` + `rle-dict`; 111,133 B (108.53 KiB) encoded; 64.673745x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 81,692 B (79.78 KiB) compressed - `zstd-3` + `plain`; 4,003,589 B (3.82 MiB) encoded; 49.026233x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
3. 87,058 B (85.02 KiB) compressed - `snappy` + `rle-dict`; 111,224 B (108.62 KiB) encoded; 46.004399x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 135,708 B (132.53 KiB) compressed - `zstd-3` + `delta-binary-packed`; 552,646 B (539.69 KiB) encoded; 29.512269x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
5. 162,742 B (158.93 KiB) compressed - `snappy` + `delta-binary-packed`; 552,414 B (539.47 KiB) encoded; 24.609818x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-delta-binary-packed`

ZSTD:
1. 61,927 B (60.48 KiB) compressed - `zstd-3` + `rle-dict`; 111,133 B (108.53 KiB) encoded; 64.673745x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 81,692 B (79.78 KiB) compressed - `zstd-3` + `plain`; 4,003,589 B (3.82 MiB) encoded; 49.026233x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
3. 135,708 B (132.53 KiB) compressed - `zstd-3` + `delta-binary-packed`; 552,646 B (539.69 KiB) encoded; 29.512269x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

Snappy:
1. 87,058 B (85.02 KiB) compressed - `snappy` + `rle-dict`; 111,224 B (108.62 KiB) encoded; 46.004399x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 162,742 B (158.93 KiB) compressed - `snappy` + `delta-binary-packed`; 552,414 B (539.47 KiB) encoded; 24.609818x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 282,833 B (276.20 KiB) compressed - `snappy` + `plain`; 4,003,714 B (3.82 MiB) encoded; 14.160480x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`

## ResolutionHeight (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `49 / 70 / 103`
- Page cardinality per row group min/median/max of mins: `49 / 70 / 103`; of maxes: `49 / 70 / 103`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/resolutionheight_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/resolutionheight_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/resolutionheight_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/resolutionheight_value_length.svg)


Compressed overall:
1. 186,022 B (181.66 KiB) compressed - `zstd-3` + `plain`; 4,003,586 B (3.82 MiB) encoded; 21.529980x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-plain`
2. 205,409 B (200.59 KiB) compressed - `zstd-3` + `rle-dict`; 371,693 B (362.98 KiB) encoded; 19.497929x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 242,446 B (236.76 KiB) compressed - `snappy` + `rle-dict`; 373,970 B (365.21 KiB) encoded; 16.519349x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
4. 358,361 B (349.96 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,340,323 B (1.28 MiB) encoded; 11.176021x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain`
5. 372,941 B (364.20 KiB) compressed - `snappy` + `plain`; 4,003,717 B (3.82 MiB) encoded; 10.739098x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

ZSTD:
1. 186,022 B (181.66 KiB) compressed - `zstd-3` + `plain`; 4,003,586 B (3.82 MiB) encoded; 21.529980x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-plain`
2. 205,409 B (200.59 KiB) compressed - `zstd-3` + `rle-dict`; 371,693 B (362.98 KiB) encoded; 19.497929x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 358,361 B (349.96 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,340,323 B (1.28 MiB) encoded; 11.176021x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain`

Snappy:
1. 242,446 B (236.76 KiB) compressed - `snappy` + `rle-dict`; 373,970 B (365.21 KiB) encoded; 16.519349x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
2. 372,941 B (364.20 KiB) compressed - `snappy` + `plain`; 4,003,717 B (3.82 MiB) encoded; 10.739098x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 422,347 B (412.45 KiB) compressed - `snappy` + `delta-binary-packed`; 1,341,881 B (1.28 MiB) encoded; 9.482842x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-plain`

## ResolutionWidth (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `48 / 64 / 84`
- Page cardinality per row group min/median/max of mins: `48 / 64 / 84`; of maxes: `48 / 64 / 84`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/resolutionwidth_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/resolutionwidth_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/resolutionwidth_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/resolutionwidth_value_length.svg)


Compressed overall:
1. 187,130 B (182.74 KiB) compressed - `zstd-3` + `plain`; 4,003,581 B (3.82 MiB) encoded; 21.402517x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
2. 203,528 B (198.76 KiB) compressed - `zstd-3` + `rle-dict`; 368,799 B (360.16 KiB) encoded; 19.678143x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 241,573 B (235.91 KiB) compressed - `snappy` + `rle-dict`; 365,920 B (357.34 KiB) encoded; 16.579059x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 373,198 B (364.45 KiB) compressed - `snappy` + `plain`; 4,003,715 B (3.82 MiB) encoded; 10.731711x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
5. 409,455 B (399.86 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,281,717 B (1.22 MiB) encoded; 9.781424x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

ZSTD:
1. 187,130 B (182.74 KiB) compressed - `zstd-3` + `plain`; 4,003,581 B (3.82 MiB) encoded; 21.402517x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
2. 203,528 B (198.76 KiB) compressed - `zstd-3` + `rle-dict`; 368,799 B (360.16 KiB) encoded; 19.678143x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 409,455 B (399.86 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,281,717 B (1.22 MiB) encoded; 9.781424x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 241,573 B (235.91 KiB) compressed - `snappy` + `rle-dict`; 365,920 B (357.34 KiB) encoded; 16.579059x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 373,198 B (364.45 KiB) compressed - `snappy` + `plain`; 4,003,715 B (3.82 MiB) encoded; 10.731711x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
3. 477,429 B (466.24 KiB) compressed - `snappy` + `delta-binary-packed`; 1,284,851 B (1.23 MiB) encoded; 8.388793x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-rle-dict`

## ResponseEndTiming (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `528 / 673 / 1,577`
- Page cardinality per row group min/median/max of mins: `528 / 673 / 1,577`; of maxes: `528 / 673 / 1,577`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/responseendtiming_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/responseendtiming_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/responseendtiming_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/responseendtiming_value_length.svg)


Compressed overall:
1. 937,781 B (915.80 KiB) compressed - `zstd-3` + `plain`; 4,003,640 B (3.82 MiB) encoded; 4.270774x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain`
2. 1,037,700 B (1013.38 KiB) compressed - `zstd-3` + `rle-dict`; 1,358,021 B (1.30 MiB) encoded; 3.859546x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
3. 1,211,062 B (1.15 MiB) compressed - `snappy` + `rle-dict`; 1,362,045 B (1.30 MiB) encoded; 3.307057x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 1,280,646 B (1.22 MiB) compressed - `snappy` + `plain`; 4,003,770 B (3.82 MiB) encoded; 3.127368x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
5. 1,303,985 B (1.24 MiB) compressed - `zstd-3` + `delta-binary-packed`; 1,502,343 B (1.43 MiB) encoded; 3.071393x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict`

ZSTD:
1. 937,781 B (915.80 KiB) compressed - `zstd-3` + `plain`; 4,003,640 B (3.82 MiB) encoded; 4.270774x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-plain`
2. 1,037,700 B (1013.38 KiB) compressed - `zstd-3` + `rle-dict`; 1,358,021 B (1.30 MiB) encoded; 3.859546x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
3. 1,303,985 B (1.24 MiB) compressed - `zstd-3` + `delta-binary-packed`; 1,502,343 B (1.43 MiB) encoded; 3.071393x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-rle-dict`

Snappy:
1. 1,211,062 B (1.15 MiB) compressed - `snappy` + `rle-dict`; 1,362,045 B (1.30 MiB) encoded; 3.307057x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 1,280,646 B (1.22 MiB) compressed - `snappy` + `plain`; 4,003,770 B (3.82 MiB) encoded; 3.127368x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
3. 1,305,051 B (1.24 MiB) compressed - `snappy` + `delta-binary-packed`; 1,503,555 B (1.43 MiB) encoded; 3.068885x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-plain`

## ResponseStartTiming (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `800 / 1,112 / 3,761`
- Page cardinality per row group min/median/max of mins: `800 / 1,112 / 3,761`; of maxes: `800 / 1,112 / 3,761`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/responsestarttiming_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/responsestarttiming_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/responsestarttiming_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/responsestarttiming_value_length.svg)


Compressed overall:
1. 1,245,745 B (1.19 MiB) compressed - `zstd-3` + `plain`; 4,003,643 B (3.82 MiB) encoded; 3.214984x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`
2. 1,510,735 B (1.44 MiB) compressed - `zstd-3` + `delta-binary-packed`; 1,704,137 B (1.63 MiB) encoded; 2.651061x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-rle-dict`
3. 1,516,154 B (1.45 MiB) compressed - `snappy` + `delta-binary-packed`; 1,703,414 B (1.62 MiB) encoded; 2.641585x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 1,556,751 B (1.48 MiB) compressed - `zstd-3` + `rle-dict`; 1,799,520 B (1.72 MiB) encoded; 2.572698x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict`
5. 1,714,995 B (1.64 MiB) compressed - `snappy` + `rle-dict`; 1,810,426 B (1.73 MiB) encoded; 2.335313x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`

ZSTD:
1. 1,245,745 B (1.19 MiB) compressed - `zstd-3` + `plain`; 4,003,643 B (3.82 MiB) encoded; 3.214984x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`
2. 1,510,735 B (1.44 MiB) compressed - `zstd-3` + `delta-binary-packed`; 1,704,137 B (1.63 MiB) encoded; 2.651061x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-rle-dict`
3. 1,556,751 B (1.48 MiB) compressed - `zstd-3` + `rle-dict`; 1,799,520 B (1.72 MiB) encoded; 2.572698x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict`

Snappy:
1. 1,516,154 B (1.45 MiB) compressed - `snappy` + `delta-binary-packed`; 1,703,414 B (1.62 MiB) encoded; 2.641585x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 1,714,995 B (1.64 MiB) compressed - `snappy` + `rle-dict`; 1,810,426 B (1.73 MiB) encoded; 2.335313x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
3. 1,721,098 B (1.64 MiB) compressed - `snappy` + `plain`; 4,003,772 B (3.82 MiB) encoded; 2.327032x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## Robotness (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `85 / 114 / 200`
- Page cardinality per row group min/median/max of mins: `85 / 114 / 200`; of maxes: `85 / 114 / 200`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/robotness_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/robotness_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/robotness_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/robotness_value_length.svg)


Compressed overall:
1. 173,390 B (169.33 KiB) compressed - `zstd-3` + `plain`; 4,003,775 B (3.82 MiB) encoded; 23.098529x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-rle-dict`
2. 240,793 B (235.15 KiB) compressed - `zstd-3` + `rle-dict`; 415,545 B (405.81 KiB) encoded; 16.632768x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 281,529 B (274.93 KiB) compressed - `snappy` + `rle-dict`; 414,918 B (405.19 KiB) encoded; 14.226080x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
4. 336,057 B (328.18 KiB) compressed - `zstd-3` + `delta-binary-packed`; 841,286 B (821.57 KiB) encoded; 11.917782x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict`
5. 381,443 B (372.50 KiB) compressed - `snappy` + `plain`; 4,003,719 B (3.82 MiB) encoded; 10.499744x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`

ZSTD:
1. 173,390 B (169.33 KiB) compressed - `zstd-3` + `plain`; 4,003,775 B (3.82 MiB) encoded; 23.098529x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-rle-dict`
2. 240,793 B (235.15 KiB) compressed - `zstd-3` + `rle-dict`; 415,545 B (405.81 KiB) encoded; 16.632768x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 336,057 B (328.18 KiB) compressed - `zstd-3` + `delta-binary-packed`; 841,286 B (821.57 KiB) encoded; 11.917782x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict`

Snappy:
1. 281,529 B (274.93 KiB) compressed - `snappy` + `rle-dict`; 414,918 B (405.19 KiB) encoded; 14.226080x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 381,443 B (372.50 KiB) compressed - `snappy` + `plain`; 4,003,719 B (3.82 MiB) encoded; 10.499744x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
3. 384,989 B (375.97 KiB) compressed - `snappy` + `delta-binary-packed`; 838,567 B (818.91 KiB) encoded; 10.403035x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## SearchEngineID (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `3 / 11 / 15`
- Page cardinality per row group min/median/max of mins: `3 / 11 / 15`; of maxes: `3 / 11 / 15`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/searchengineid_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/searchengineid_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/searchengineid_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/searchengineid_value_length.svg)


Compressed overall:
1. 75,997 B (74.22 KiB) compressed - `zstd-3` + `rle-dict`; 166,419 B (162.52 KiB) encoded; 52.700107x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 101,565 B (99.18 KiB) compressed - `zstd-3` + `plain`; 4,003,588 B (3.82 MiB) encoded; 39.433368x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed`
3. 104,165 B (101.72 KiB) compressed - `snappy` + `rle-dict`; 166,930 B (163.02 KiB) encoded; 38.449095x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-delta-binary-packed`
4. 150,070 B (146.55 KiB) compressed - `zstd-3` + `delta-binary-packed`; 378,395 B (369.53 KiB) encoded; 26.687879x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
5. 168,281 B (164.34 KiB) compressed - `snappy` + `delta-binary-packed`; 377,704 B (368.85 KiB) encoded; 23.799775x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 75,997 B (74.22 KiB) compressed - `zstd-3` + `rle-dict`; 166,419 B (162.52 KiB) encoded; 52.700107x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 101,565 B (99.18 KiB) compressed - `zstd-3` + `plain`; 4,003,588 B (3.82 MiB) encoded; 39.433368x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-rle-dict-ts-delta-binary-packed`
3. 150,070 B (146.55 KiB) compressed - `zstd-3` + `delta-binary-packed`; 378,395 B (369.53 KiB) encoded; 26.687879x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`

Snappy:
1. 104,165 B (101.72 KiB) compressed - `snappy` + `rle-dict`; 166,930 B (163.02 KiB) encoded; 38.449095x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-plain-ts-delta-binary-packed`
2. 168,281 B (164.34 KiB) compressed - `snappy` + `delta-binary-packed`; 377,704 B (368.85 KiB) encoded; 23.799775x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-plain`
3. 292,991 B (286.12 KiB) compressed - `snappy` + `plain`; 4,003,715 B (3.82 MiB) encoded; 13.669533x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## SearchPhrase (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `7 / 421 / 558`
- Page cardinality per row group min/median/max of mins: `7 / 421 / 558`; of maxes: `7 / 421 / 558`
- Value length min/median/max: `0 / 0 / 1,939` bytes
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `37 / 148 / 1,939`

![Row-group cardinality](column_shape_stats/images/searchphrase_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/searchphrase_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/searchphrase_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/searchphrase_value_length.svg)


Compressed overall:
1. 635,270 B (620.38 KiB) compressed - `zstd-3` + `rle-dict`; 1,641,728 B (1.57 MiB) encoded; 11.865681x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`
2. 719,279 B (702.42 KiB) compressed - `zstd-3` + `plain`; 7,535,016 B (7.19 MiB) encoded; 10.479815x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 784,080 B (765.70 KiB) compressed - `snappy` + `rle-dict`; 1,642,854 B (1.57 MiB) encoded; 9.613701x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
4. 811,698 B (792.67 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 4,222,053 B (4.03 MiB) encoded; 9.286596x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
5. 825,797 B (806.44 KiB) compressed - `zstd-3` + `delta-byte-array`; 2,957,835 B (2.82 MiB) encoded; 9.128044x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 635,270 B (620.38 KiB) compressed - `zstd-3` + `rle-dict`; 1,641,728 B (1.57 MiB) encoded; 11.865681x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`
2. 719,279 B (702.42 KiB) compressed - `zstd-3` + `plain`; 7,535,016 B (7.19 MiB) encoded; 10.479815x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 811,698 B (792.67 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 4,222,053 B (4.03 MiB) encoded; 9.286596x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
4. 825,797 B (806.44 KiB) compressed - `zstd-3` + `delta-byte-array`; 2,957,835 B (2.82 MiB) encoded; 9.128044x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`

Snappy:
1. 784,080 B (765.70 KiB) compressed - `snappy` + `rle-dict`; 1,642,854 B (1.57 MiB) encoded; 9.613701x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 992,328 B (969.07 KiB) compressed - `snappy` + `delta-byte-array`; 2,964,725 B (2.83 MiB) encoded; 7.596189x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
3. 1,012,183 B (988.46 KiB) compressed - `snappy` + `delta-length-byte-array`; 4,223,215 B (4.03 MiB) encoded; 7.447182x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
4. 1,096,608 B (1.05 MiB) compressed - `snappy` + `plain`; 7,535,343 B (7.19 MiB) encoded; 6.873843x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-delta-binary-packed-ts-delta-binary-packed`

## SendTiming (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 989`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 989`; of maxes: `1 / 1 / 989`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/sendtiming_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/sendtiming_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/sendtiming_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/sendtiming_value_length.svg)


Compressed overall:
1. 61,415 B (59.98 KiB) compressed - `zstd-3` + `plain`; 4,003,537 B (3.82 MiB) encoded; 65.212961x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
2. 76,243 B (74.46 KiB) compressed - `zstd-3` + `rle-dict`; 120,257 B (117.44 KiB) encoded; 52.530121x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
3. 92,894 B (90.72 KiB) compressed - `snappy` + `rle-dict`; 119,257 B (116.46 KiB) encoded; 43.114238x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-delta-binary-packed-ts-rle-dict`
4. 108,282 B (105.74 KiB) compressed - `zstd-3` + `delta-binary-packed`; 173,667 B (169.60 KiB) encoded; 36.987255x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-rle-dict`
5. 111,631 B (109.01 KiB) compressed - `snappy` + `delta-binary-packed`; 173,175 B (169.12 KiB) encoded; 35.877615x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`

ZSTD:
1. 61,415 B (59.98 KiB) compressed - `zstd-3` + `plain`; 4,003,537 B (3.82 MiB) encoded; 65.212961x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
2. 76,243 B (74.46 KiB) compressed - `zstd-3` + `rle-dict`; 120,257 B (117.44 KiB) encoded; 52.530121x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
3. 108,282 B (105.74 KiB) compressed - `zstd-3` + `delta-binary-packed`; 173,667 B (169.60 KiB) encoded; 36.987255x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-rle-dict`

Snappy:
1. 92,894 B (90.72 KiB) compressed - `snappy` + `rle-dict`; 119,257 B (116.46 KiB) encoded; 43.114238x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-delta-binary-packed-ts-rle-dict`
2. 111,631 B (109.01 KiB) compressed - `snappy` + `delta-binary-packed`; 173,175 B (169.12 KiB) encoded; 35.877615x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
3. 268,732 B (262.43 KiB) compressed - `snappy` + `plain`; 4,003,775 B (3.82 MiB) encoded; 14.903525x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-rle-dict`

## Sex (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `3 / 3 / 3`
- Page cardinality per row group min/median/max of mins: `3 / 3 / 3`; of maxes: `3 / 3 / 3`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/sex_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/sex_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/sex_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/sex_value_length.svg)


Compressed overall:
1. 80,090 B (78.21 KiB) compressed - `zstd-3` + `rle-dict`; 133,318 B (130.19 KiB) encoded; 50.006892x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 107,246 B (104.73 KiB) compressed - `zstd-3` + `plain`; 4,003,584 B (3.82 MiB) encoded; 37.344535x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 108,767 B (106.22 KiB) compressed - `zstd-3` + `delta-binary-packed`; 253,027 B (247.10 KiB) encoded; 36.822308x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain`
4. 112,308 B (109.68 KiB) compressed - `snappy` + `rle-dict`; 133,539 B (130.41 KiB) encoded; 35.661324x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
5. 135,239 B (132.07 KiB) compressed - `snappy` + `delta-binary-packed`; 252,677 B (246.75 KiB) encoded; 29.614623x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 80,090 B (78.21 KiB) compressed - `zstd-3` + `rle-dict`; 133,318 B (130.19 KiB) encoded; 50.006892x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 107,246 B (104.73 KiB) compressed - `zstd-3` + `plain`; 4,003,584 B (3.82 MiB) encoded; 37.344535x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 108,767 B (106.22 KiB) compressed - `zstd-3` + `delta-binary-packed`; 253,027 B (247.10 KiB) encoded; 36.822308x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain`

Snappy:
1. 112,308 B (109.68 KiB) compressed - `snappy` + `rle-dict`; 133,539 B (130.41 KiB) encoded; 35.661324x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
2. 135,239 B (132.07 KiB) compressed - `snappy` + `delta-binary-packed`; 252,677 B (246.75 KiB) encoded; 29.614623x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`
3. 302,471 B (295.38 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 13.241111x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## SilverlightVersion1 (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `4 / 5 / 6`
- Page cardinality per row group min/median/max of mins: `4 / 5 / 6`; of maxes: `4 / 5 / 6`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/silverlightversion1_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/silverlightversion1_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/silverlightversion1_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/silverlightversion1_value_length.svg)


Compressed overall:
1. 82,970 B (81.03 KiB) compressed - `zstd-3` + `rle-dict`; 155,404 B (151.76 KiB) encoded; 48.271062x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
2. 88,747 B (86.67 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 45.128849x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 112,565 B (109.93 KiB) compressed - `snappy` + `rle-dict`; 156,160 B (152.50 KiB) encoded; 35.579887x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`
4. 118,729 B (115.95 KiB) compressed - `zstd-3` + `delta-binary-packed`; 437,185 B (426.94 KiB) encoded; 33.732702x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain`
5. 151,472 B (147.92 KiB) compressed - `snappy` + `delta-binary-packed`; 437,094 B (426.85 KiB) encoded; 26.440860x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 82,970 B (81.03 KiB) compressed - `zstd-3` + `rle-dict`; 155,404 B (151.76 KiB) encoded; 48.271062x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
2. 88,747 B (86.67 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 45.128849x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 118,729 B (115.95 KiB) compressed - `zstd-3` + `delta-binary-packed`; 437,185 B (426.94 KiB) encoded; 33.732702x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-plain`

Snappy:
1. 112,565 B (109.93 KiB) compressed - `snappy` + `rle-dict`; 156,160 B (152.50 KiB) encoded; 35.579887x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`
2. 151,472 B (147.92 KiB) compressed - `snappy` + `delta-binary-packed`; 437,094 B (426.85 KiB) encoded; 26.440860x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`
3. 287,124 B (280.39 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 13.948851x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`

## SilverlightVersion2 (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `2 / 2 / 2`
- Page cardinality per row group min/median/max of mins: `2 / 2 / 2`; of maxes: `2 / 2 / 2`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/silverlightversion2_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/silverlightversion2_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/silverlightversion2_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/silverlightversion2_value_length.svg)


Compressed overall:
1. 56,055 B (54.74 KiB) compressed - `zstd-3` + `rle-dict`; 88,952 B (86.87 KiB) encoded; 71.448666x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`
2. 73,199 B (71.48 KiB) compressed - `zstd-3` + `delta-binary-packed`; 207,538 B (202.67 KiB) encoded; 54.714614x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain`
3. 73,703 B (71.98 KiB) compressed - `zstd-3` + `plain`; 4,003,589 B (3.82 MiB) encoded; 54.340461x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
4. 76,117 B (74.33 KiB) compressed - `snappy` + `rle-dict`; 88,970 B (86.88 KiB) encoded; 52.617089x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
5. 98,397 B (96.09 KiB) compressed - `snappy` + `delta-binary-packed`; 208,063 B (203.19 KiB) encoded; 40.703019x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-plain`

ZSTD:
1. 56,055 B (54.74 KiB) compressed - `zstd-3` + `rle-dict`; 88,952 B (86.87 KiB) encoded; 71.448666x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`
2. 73,199 B (71.48 KiB) compressed - `zstd-3` + `delta-binary-packed`; 207,538 B (202.67 KiB) encoded; 54.714614x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain`
3. 73,703 B (71.98 KiB) compressed - `zstd-3` + `plain`; 4,003,589 B (3.82 MiB) encoded; 54.340461x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`

Snappy:
1. 76,117 B (74.33 KiB) compressed - `snappy` + `rle-dict`; 88,970 B (86.88 KiB) encoded; 52.617089x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 98,397 B (96.09 KiB) compressed - `snappy` + `delta-binary-packed`; 208,063 B (203.19 KiB) encoded; 40.703019x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-plain`
3. 269,618 B (263.30 KiB) compressed - `snappy` + `plain`; 4,003,713 B (3.82 MiB) encoded; 14.854553x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`

## SilverlightVersion3 (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `17 / 22 / 28`
- Page cardinality per row group min/median/max of mins: `17 / 22 / 28`; of maxes: `17 / 22 / 28`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/silverlightversion3_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/silverlightversion3_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/silverlightversion3_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/silverlightversion3_value_length.svg)


Compressed overall:
1. 123,341 B (120.45 KiB) compressed - `zstd-3` + `rle-dict`; 231,114 B (225.70 KiB) encoded; 32.471368x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 123,727 B (120.83 KiB) compressed - `zstd-3` + `plain`; 4,003,583 B (3.82 MiB) encoded; 32.370065x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
3. 153,901 B (150.29 KiB) compressed - `snappy` + `rle-dict`; 231,262 B (225.84 KiB) encoded; 26.023554x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
4. 304,661 B (297.52 KiB) compressed - `snappy` + `plain`; 4,003,715 B (3.82 MiB) encoded; 13.145926x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
5. 315,969 B (308.56 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,861,362 B (1.78 MiB) encoded; 12.675456x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`

ZSTD:
1. 123,341 B (120.45 KiB) compressed - `zstd-3` + `rle-dict`; 231,114 B (225.70 KiB) encoded; 32.471368x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 123,727 B (120.83 KiB) compressed - `zstd-3` + `plain`; 4,003,583 B (3.82 MiB) encoded; 32.370065x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`
3. 315,969 B (308.56 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,861,362 B (1.78 MiB) encoded; 12.675456x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 153,901 B (150.29 KiB) compressed - `snappy` + `rle-dict`; 231,262 B (225.84 KiB) encoded; 26.023554x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 304,661 B (297.52 KiB) compressed - `snappy` + `plain`; 4,003,715 B (3.82 MiB) encoded; 13.145926x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`
3. 390,759 B (381.60 KiB) compressed - `snappy` + `delta-binary-packed`; 1,859,939 B (1.77 MiB) encoded; 10.249415x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`

## SilverlightVersion4 (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 3`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 3`; of maxes: `1 / 1 / 3`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/silverlightversion4_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/silverlightversion4_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/silverlightversion4_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/silverlightversion4_value_length.svg)


Compressed overall:
1. 4,357 B (4.25 KiB) compressed - `zstd-3` + `plain`; 4,003,528 B (3.82 MiB) encoded; 919.222860x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
2. 5,164 B (5.04 KiB) compressed - `zstd-3` + `delta-binary-packed`; 43,549 B (42.53 KiB) encoded; 775.572037x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain`
3. 5,382 B (5.26 KiB) compressed - `snappy` + `rle-dict`; 5,150 B (5.03 KiB) encoded; 744.157191x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
4. 6,084 B (5.94 KiB) compressed - `zstd-3` + `rle-dict`; 5,058 B (4.94 KiB) encoded; 658.292899x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-rle-dict`
5. 6,432 B (6.28 KiB) compressed - `snappy` + `delta-binary-packed`; 43,669 B (42.65 KiB) encoded; 622.676306x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 4,357 B (4.25 KiB) compressed - `zstd-3` + `plain`; 4,003,528 B (3.82 MiB) encoded; 919.222860x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
2. 5,164 B (5.04 KiB) compressed - `zstd-3` + `delta-binary-packed`; 43,549 B (42.53 KiB) encoded; 775.572037x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain`
3. 6,084 B (5.94 KiB) compressed - `zstd-3` + `rle-dict`; 5,058 B (4.94 KiB) encoded; 658.292899x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-rle-dict`

Snappy:
1. 5,382 B (5.26 KiB) compressed - `snappy` + `rle-dict`; 5,150 B (5.03 KiB) encoded; 744.157191x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 6,432 B (6.28 KiB) compressed - `snappy` + `delta-binary-packed`; 43,669 B (42.65 KiB) encoded; 622.676306x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 204,099 B (199.32 KiB) compressed - `snappy` + `plain`; 4,003,711 B (3.82 MiB) encoded; 19.623095x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## SocialAction (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length min/median/max: `0 / 0 / 0` bytes
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 0 / 0`

![Row-group cardinality](column_shape_stats/images/socialaction_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/socialaction_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/socialaction_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/socialaction_value_length.svg)


Compressed overall:
1. 2,848 B (2.78 KiB) compressed - `zstd-3` + `plain`; 4,002,159 B (3.82 MiB) encoded; 1405.603230x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
2. 3,567 B (3.48 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 41,679 B (40.70 KiB) encoded; 1122.275862x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 3,680 B (3.59 KiB) compressed - `zstd-3` + `delta-byte-array`; 81,200 B (79.30 KiB) encoded; 1087.814674x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 4,493 B (4.39 KiB) compressed - `snappy` + `rle-dict`; 4,253 B (4.15 KiB) encoded; 890.976630x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
5. 4,832 B (4.72 KiB) compressed - `snappy` + `delta-length-byte-array`; 41,830 B (40.85 KiB) encoded; 828.468129x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`

ZSTD:
1. 2,848 B (2.78 KiB) compressed - `zstd-3` + `plain`; 4,002,159 B (3.82 MiB) encoded; 1405.603230x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
2. 3,567 B (3.48 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 41,679 B (40.70 KiB) encoded; 1122.275862x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 3,680 B (3.59 KiB) compressed - `zstd-3` + `delta-byte-array`; 81,200 B (79.30 KiB) encoded; 1087.814674x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 5,247 B (5.12 KiB) compressed - `zstd-3` + `rle-dict`; 4,185 B (4.09 KiB) encoded; 762.942253x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-plain`

Snappy:
1. 4,493 B (4.39 KiB) compressed - `snappy` + `rle-dict`; 4,253 B (4.15 KiB) encoded; 890.976630x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
2. 4,832 B (4.72 KiB) compressed - `snappy` + `delta-length-byte-array`; 41,830 B (40.85 KiB) encoded; 828.468129x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
3. 6,836 B (6.68 KiB) compressed - `snappy` + `delta-byte-array`; 81,353 B (79.45 KiB) encoded; 585.599473x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 202,701 B (197.95 KiB) compressed - `snappy` + `plain`; 4,002,334 B (3.82 MiB) encoded; 19.749079x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-delta-binary-packed`

## SocialNetwork (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length min/median/max: `0 / 0 / 0` bytes
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 0 / 0`

![Row-group cardinality](column_shape_stats/images/socialnetwork_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/socialnetwork_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/socialnetwork_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/socialnetwork_value_length.svg)


Compressed overall:
1. 2,848 B (2.78 KiB) compressed - `zstd-3` + `plain`; 4,002,159 B (3.82 MiB) encoded; 1405.603230x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
2. 3,567 B (3.48 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 41,679 B (40.70 KiB) encoded; 1122.275862x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 3,680 B (3.59 KiB) compressed - `zstd-3` + `delta-byte-array`; 81,200 B (79.30 KiB) encoded; 1087.814674x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 4,493 B (4.39 KiB) compressed - `snappy` + `rle-dict`; 4,253 B (4.15 KiB) encoded; 890.976630x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
5. 4,832 B (4.72 KiB) compressed - `snappy` + `delta-length-byte-array`; 41,830 B (40.85 KiB) encoded; 828.468129x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`

ZSTD:
1. 2,848 B (2.78 KiB) compressed - `zstd-3` + `plain`; 4,002,159 B (3.82 MiB) encoded; 1405.603230x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
2. 3,567 B (3.48 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 41,679 B (40.70 KiB) encoded; 1122.275862x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 3,680 B (3.59 KiB) compressed - `zstd-3` + `delta-byte-array`; 81,200 B (79.30 KiB) encoded; 1087.814674x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 5,247 B (5.12 KiB) compressed - `zstd-3` + `rle-dict`; 4,185 B (4.09 KiB) encoded; 762.942253x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-delta-binary-packed-ts-plain`

Snappy:
1. 4,493 B (4.39 KiB) compressed - `snappy` + `rle-dict`; 4,253 B (4.15 KiB) encoded; 890.976630x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
2. 4,832 B (4.72 KiB) compressed - `snappy` + `delta-length-byte-array`; 41,830 B (40.85 KiB) encoded; 828.468129x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
3. 6,836 B (6.68 KiB) compressed - `snappy` + `delta-byte-array`; 81,353 B (79.45 KiB) encoded; 585.599473x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 202,701 B (197.95 KiB) compressed - `snappy` + `plain`; 4,002,334 B (3.82 MiB) encoded; 19.749079x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-delta-binary-packed`

## SocialSourceNetworkID (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 2 / 4`
- Page cardinality per row group min/median/max of mins: `1 / 2 / 4`; of maxes: `1 / 2 / 4`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/socialsourcenetworkid_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/socialsourcenetworkid_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/socialsourcenetworkid_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/socialsourcenetworkid_value_length.svg)


Compressed overall:
1. 5,318 B (5.19 KiB) compressed - `zstd-3` + `plain`; 4,003,530 B (3.82 MiB) encoded; 753.112636x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-plain`
2. 6,192 B (6.05 KiB) compressed - `snappy` + `rle-dict`; 5,960 B (5.82 KiB) encoded; 646.810885x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 6,912 B (6.75 KiB) compressed - `zstd-3` + `rle-dict`; 5,886 B (5.75 KiB) encoded; 579.434751x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
4. 7,014 B (6.85 KiB) compressed - `zstd-3` + `delta-binary-packed`; 47,327 B (46.22 KiB) encoded; 571.008412x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
5. 8,282 B (8.09 KiB) compressed - `snappy` + `delta-binary-packed`; 47,482 B (46.37 KiB) encoded; 483.585245x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 5,318 B (5.19 KiB) compressed - `zstd-3` + `plain`; 4,003,530 B (3.82 MiB) encoded; 753.112636x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-plain`
2. 6,912 B (6.75 KiB) compressed - `zstd-3` + `rle-dict`; 5,886 B (5.75 KiB) encoded; 579.434751x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 7,014 B (6.85 KiB) compressed - `zstd-3` + `delta-binary-packed`; 47,327 B (46.22 KiB) encoded; 571.008412x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

Snappy:
1. 6,192 B (6.05 KiB) compressed - `snappy` + `rle-dict`; 5,960 B (5.82 KiB) encoded; 646.810885x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 8,282 B (8.09 KiB) compressed - `snappy` + `delta-binary-packed`; 47,482 B (46.37 KiB) encoded; 483.585245x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`
3. 204,453 B (199.66 KiB) compressed - `snappy` + `plain`; 4,003,714 B (3.82 MiB) encoded; 19.589113x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## SocialSourcePage (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 5`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 5`; of maxes: `1 / 1 / 5`
- Value length min/median/max: `0 / 0 / 28` bytes
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 0 / 28`

![Row-group cardinality](column_shape_stats/images/socialsourcepage_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/socialsourcepage_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/socialsourcepage_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/socialsourcepage_value_length.svg)


Compressed overall:
1. 4,853 B (4.74 KiB) compressed - `zstd-3` + `plain`; 4,003,988 B (3.82 MiB) encoded; 825.294251x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict`
2. 5,977 B (5.84 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 46,239 B (45.16 KiB) encoded; 670.094194x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-rle-dict`
3. 6,050 B (5.91 KiB) compressed - `snappy` + `rle-dict`; 5,817 B (5.68 KiB) encoded; 662.008760x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
4. 6,144 B (6.00 KiB) compressed - `zstd-3` + `delta-byte-array`; 86,534 B (84.51 KiB) encoded; 651.880371x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`
5. 6,826 B (6.67 KiB) compressed - `zstd-3` + `rle-dict`; 5,764 B (5.63 KiB) encoded; 586.749634x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`

ZSTD:
1. 4,853 B (4.74 KiB) compressed - `zstd-3` + `plain`; 4,003,988 B (3.82 MiB) encoded; 825.294251x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict`
2. 5,977 B (5.84 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 46,239 B (45.16 KiB) encoded; 670.094194x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-rle-dict`
3. 6,144 B (6.00 KiB) compressed - `zstd-3` + `delta-byte-array`; 86,534 B (84.51 KiB) encoded; 651.880371x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`
4. 6,826 B (6.67 KiB) compressed - `zstd-3` + `rle-dict`; 5,764 B (5.63 KiB) encoded; 586.749634x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`

Snappy:
1. 6,050 B (5.91 KiB) compressed - `snappy` + `rle-dict`; 5,817 B (5.68 KiB) encoded; 662.008760x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 7,270 B (7.10 KiB) compressed - `snappy` + `delta-length-byte-array`; 46,400 B (45.31 KiB) encoded; 550.915131x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
3. 9,337 B (9.12 KiB) compressed - `snappy` + `delta-byte-array`; 86,546 B (84.52 KiB) encoded; 428.955018x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
4. 204,255 B (199.47 KiB) compressed - `snappy` + `plain`; 4,004,109 B (3.82 MiB) encoded; 19.608592x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-plain-date-delta-binary-packed-ts-delta-binary-packed`

## Title (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `603`
- Row-group cardinality min/median/max: `108 / 2,277 / 2,527`
- Page cardinality per row group min/median/max of mins: `15 / 135 / 258`; of maxes: `84 / 334 / 758`
- Value length min/median/max: `0 / 121 / 1,026` bytes
- Value length per row group min/median/max of mins: `0 / 0 / 33`; of maxes: `143 / 523 / 1,026`

![Row-group cardinality](column_shape_stats/images/title_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/title_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/title_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/title_value_length.svg)


Compressed overall:
1. 7,955,232 B (7.59 MiB) compressed - `zstd-3` + `rle-dict`; 28,774,771 B (27.44 MiB) encoded; 17.961193x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 9,992,672 B (9.53 MiB) compressed - `snappy` + `rle-dict`; 28,820,866 B (27.49 MiB) encoded; 14.299024x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`
3. 13,925,809 B (13.28 MiB) compressed - `zstd-3` + `plain`; 142,862,125 B (136.24 MiB) encoded; 10.260478x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
4. 13,958,100 B (13.31 MiB) compressed - `zstd-3` + `delta-byte-array`; 64,469,305 B (61.48 MiB) encoded; 10.236741x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
5. 14,454,863 B (13.79 MiB) compressed - `zstd-3` + `delta-length-byte-array`; 140,026,422 B (133.54 MiB) encoded; 9.884940x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`

ZSTD:
1. 7,955,232 B (7.59 MiB) compressed - `zstd-3` + `rle-dict`; 28,774,771 B (27.44 MiB) encoded; 17.961193x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 13,925,809 B (13.28 MiB) compressed - `zstd-3` + `plain`; 142,862,125 B (136.24 MiB) encoded; 10.260478x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
3. 13,958,100 B (13.31 MiB) compressed - `zstd-3` + `delta-byte-array`; 64,469,305 B (61.48 MiB) encoded; 10.236741x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 14,454,863 B (13.79 MiB) compressed - `zstd-3` + `delta-length-byte-array`; 140,026,422 B (133.54 MiB) encoded; 9.884940x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`

Snappy:
1. 9,992,672 B (9.53 MiB) compressed - `snappy` + `rle-dict`; 28,820,866 B (27.49 MiB) encoded; 14.299024x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`
2. 17,100,392 B (16.31 MiB) compressed - `snappy` + `delta-byte-array`; 64,470,659 B (61.48 MiB) encoded; 8.355683x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
3. 20,886,114 B (19.92 MiB) compressed - `snappy` + `plain`; 142,865,534 B (136.25 MiB) encoded; 6.841170x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-plain-ts-delta-binary-packed`
4. 21,320,746 B (20.33 MiB) compressed - `snappy` + `delta-length-byte-array`; 140,024,907 B (133.54 MiB) encoded; 6.701710x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-plain`

## TraficSourceID (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `5 / 7 / 8`
- Page cardinality per row group min/median/max of mins: `5 / 7 / 8`; of maxes: `5 / 7 / 8`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/traficsourceid_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/traficsourceid_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/traficsourceid_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/traficsourceid_value_length.svg)


Compressed overall:
1. 178,343 B (174.16 KiB) compressed - `zstd-3` + `rle-dict`; 289,829 B (283.04 KiB) encoded; 22.457013x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-rle-dict`
2. 237,812 B (232.24 KiB) compressed - `zstd-3` + `delta-binary-packed`; 496,558 B (484.92 KiB) encoded; 16.841249x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 243,838 B (238.12 KiB) compressed - `snappy` + `rle-dict`; 289,838 B (283.04 KiB) encoded; 16.425049x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
4. 288,165 B (281.41 KiB) compressed - `zstd-3` + `plain`; 4,003,588 B (3.82 MiB) encoded; 13.898464x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-rle-dict`
5. 292,167 B (285.32 KiB) compressed - `snappy` + `delta-binary-packed`; 496,760 B (485.12 KiB) encoded; 13.708088x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`

ZSTD:
1. 178,343 B (174.16 KiB) compressed - `zstd-3` + `rle-dict`; 289,829 B (283.04 KiB) encoded; 22.457013x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-rle-dict`
2. 237,812 B (232.24 KiB) compressed - `zstd-3` + `delta-binary-packed`; 496,558 B (484.92 KiB) encoded; 16.841249x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 288,165 B (281.41 KiB) compressed - `zstd-3` + `plain`; 4,003,588 B (3.82 MiB) encoded; 13.898464x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-rle-dict`

Snappy:
1. 243,838 B (238.12 KiB) compressed - `snappy` + `rle-dict`; 289,838 B (283.04 KiB) encoded; 16.425049x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`
2. 292,167 B (285.32 KiB) compressed - `snappy` + `delta-binary-packed`; 496,760 B (485.12 KiB) encoded; 13.708088x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
3. 526,610 B (514.27 KiB) compressed - `snappy` + `plain`; 4,003,754 B (3.82 MiB) encoded; 7.605346x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## URL (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `396`
- Row-group cardinality min/median/max: `2,860 / 3,100 / 6,974`
- Page cardinality per row group min/median/max of mins: `116 / 646 / 935`; of maxes: `847 / 1,003 / 1,731`
- Value length min/median/max: `0 / 74 / 1,991` bytes
- Value length per row group min/median/max of mins: `0 / 17 / 19`; of maxes: `252 / 483 / 1,991`

![Row-group cardinality](column_shape_stats/images/url_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/url_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/url_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/url_value_length.svg)


Compressed overall:
1. 12,678,493 B (12.09 MiB) compressed - `zstd-3` + `rle-dict`; 44,047,032 B (42.01 MiB) encoded; 7.307902x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 15,058,105 B (14.36 MiB) compressed - `zstd-3` + `delta-byte-array`; 40,457,535 B (38.58 MiB) encoded; 6.153044x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-rle-dict`
3. 15,298,603 B (14.59 MiB) compressed - `zstd-3` + `plain`; 92,652,326 B (88.36 MiB) encoded; 6.056317x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-delta-binary-packed`
4. 16,024,686 B (15.28 MiB) compressed - `zstd-3` + `delta-length-byte-array`; 89,782,859 B (85.62 MiB) encoded; 5.781903x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
5. 16,085,485 B (15.34 MiB) compressed - `snappy` + `rle-dict`; 44,090,971 B (42.05 MiB) encoded; 5.760049x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`

ZSTD:
1. 12,678,493 B (12.09 MiB) compressed - `zstd-3` + `rle-dict`; 44,047,032 B (42.01 MiB) encoded; 7.307902x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
2. 15,058,105 B (14.36 MiB) compressed - `zstd-3` + `delta-byte-array`; 40,457,535 B (38.58 MiB) encoded; 6.153044x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-rle-dict`
3. 15,298,603 B (14.59 MiB) compressed - `zstd-3` + `plain`; 92,652,326 B (88.36 MiB) encoded; 6.056317x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-delta-binary-packed`
4. 16,024,686 B (15.28 MiB) compressed - `zstd-3` + `delta-length-byte-array`; 89,782,859 B (85.62 MiB) encoded; 5.781903x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 16,085,485 B (15.34 MiB) compressed - `snappy` + `rle-dict`; 44,090,971 B (42.05 MiB) encoded; 5.760049x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`
2. 18,047,077 B (17.21 MiB) compressed - `snappy` + `delta-byte-array`; 40,456,964 B (38.58 MiB) encoded; 5.133972x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 20,458,616 B (19.51 MiB) compressed - `snappy` + `plain`; 92,651,231 B (88.36 MiB) encoded; 4.528810x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-delta-binary-packed`
4. 20,776,321 B (19.81 MiB) compressed - `snappy` + `delta-length-byte-array`; 89,784,720 B (85.63 MiB) encoded; 4.459557x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-delta-binary-packed`

## URLCategoryID (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `2 / 4 / 78`
- Page cardinality per row group min/median/max of mins: `2 / 4 / 78`; of maxes: `2 / 4 / 78`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/urlcategoryid_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/urlcategoryid_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/urlcategoryid_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/urlcategoryid_value_length.svg)


Compressed overall:
1. 71,056 B (69.39 KiB) compressed - `zstd-3` + `rle-dict`; 161,772 B (157.98 KiB) encoded; 56.364670x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
2. 87,467 B (85.42 KiB) compressed - `zstd-3` + `plain`; 4,003,589 B (3.82 MiB) encoded; 45.789246x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
3. 91,524 B (89.38 KiB) compressed - `snappy` + `rle-dict`; 161,494 B (157.71 KiB) encoded; 43.759538x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
4. 202,577 B (197.83 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,084,378 B (1.03 MiB) encoded; 19.770497x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain`
5. 241,518 B (235.86 KiB) compressed - `snappy` + `delta-binary-packed`; 1,082,023 B (1.03 MiB) encoded; 16.582814x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 71,056 B (69.39 KiB) compressed - `zstd-3` + `rle-dict`; 161,772 B (157.98 KiB) encoded; 56.364670x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
2. 87,467 B (85.42 KiB) compressed - `zstd-3` + `plain`; 4,003,589 B (3.82 MiB) encoded; 45.789246x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
3. 202,577 B (197.83 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,084,378 B (1.03 MiB) encoded; 19.770497x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain`

Snappy:
1. 91,524 B (89.38 KiB) compressed - `snappy` + `rle-dict`; 161,494 B (157.71 KiB) encoded; 43.759538x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 241,518 B (235.86 KiB) compressed - `snappy` + `delta-binary-packed`; 1,082,023 B (1.03 MiB) encoded; 16.582814x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 271,779 B (265.41 KiB) compressed - `snappy` + `plain`; 4,003,714 B (3.82 MiB) encoded; 14.736415x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

## URLHash (int64)

Column shape stats:
- Parquet type: `INT(64,true)`; physical type: `INT64`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `3,001 / 3,292 / 7,420`
- Page cardinality per row group min/median/max of mins: `3,001 / 3,292 / 7,420`; of maxes: `3,001 / 3,292 / 7,420`
- Value length min/median/max: `8 / 8 / 8` bytes
- Value length per row group min/median/max of mins: `8 / 8 / 8`; of maxes: `8 / 8 / 8`

![Row-group cardinality](column_shape_stats/images/urlhash_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/urlhash_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/urlhash_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/urlhash_value_length.svg)


Compressed overall:
1. 3,580,060 B (3.41 MiB) compressed - `zstd-3` + `plain`; 8,004,555 B (7.63 MiB) encoded; 2.236362x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 4,382,430 B (4.18 MiB) compressed - `snappy` + `plain`; 8,004,876 B (7.63 MiB) encoded; 1.826911x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain`
3. 4,498,086 B (4.29 MiB) compressed - `snappy` + `rle-dict`; 4,603,185 B (4.39 MiB) encoded; 1.779937x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-delta-binary-packed`
4. 4,529,372 B (4.32 MiB) compressed - `zstd-3` + `rle-dict`; 4,619,317 B (4.41 MiB) encoded; 1.767642x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
5. 5,217,669 B (4.98 MiB) compressed - `zstd-3` + `delta-binary-packed`; 8,121,071 B (7.74 MiB) encoded; 1.534461x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict`

ZSTD:
1. 3,580,060 B (3.41 MiB) compressed - `zstd-3` + `plain`; 8,004,555 B (7.63 MiB) encoded; 2.236362x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 4,529,372 B (4.32 MiB) compressed - `zstd-3` + `rle-dict`; 4,619,317 B (4.41 MiB) encoded; 1.767642x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
3. 5,217,669 B (4.98 MiB) compressed - `zstd-3` + `delta-binary-packed`; 8,121,071 B (7.74 MiB) encoded; 1.534461x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict`

Snappy:
1. 4,382,430 B (4.18 MiB) compressed - `snappy` + `plain`; 8,004,876 B (7.63 MiB) encoded; 1.826911x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-rle-dict-date-plain-ts-plain`
2. 4,498,086 B (4.29 MiB) compressed - `snappy` + `rle-dict`; 4,603,185 B (4.39 MiB) encoded; 1.779937x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-delta-binary-packed`
3. 5,945,541 B (5.67 MiB) compressed - `snappy` + `delta-binary-packed`; 8,119,600 B (7.74 MiB) encoded; 1.346607x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

## URLRegionID (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `3 / 5 / 35`
- Page cardinality per row group min/median/max of mins: `3 / 5 / 35`; of maxes: `3 / 5 / 35`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/urlregionid_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/urlregionid_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/urlregionid_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/urlregionid_value_length.svg)


Compressed overall:
1. 40,097 B (39.16 KiB) compressed - `zstd-3` + `rle-dict`; 74,481 B (72.74 KiB) encoded; 99.884006x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict`
2. 48,046 B (46.92 KiB) compressed - `zstd-3` + `plain`; 4,003,586 B (3.82 MiB) encoded; 83.358635x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`
3. 52,198 B (50.97 KiB) compressed - `snappy` + `rle-dict`; 73,898 B (72.17 KiB) encoded; 76.728016x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
4. 99,369 B (97.04 KiB) compressed - `zstd-3` + `delta-binary-packed`; 722,823 B (705.88 KiB) encoded; 40.304813x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict`
5. 126,255 B (123.30 KiB) compressed - `snappy` + `delta-binary-packed`; 716,541 B (699.75 KiB) encoded; 31.721904x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 40,097 B (39.16 KiB) compressed - `zstd-3` + `rle-dict`; 74,481 B (72.74 KiB) encoded; 99.884006x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict`
2. 48,046 B (46.92 KiB) compressed - `zstd-3` + `plain`; 4,003,586 B (3.82 MiB) encoded; 83.358635x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`
3. 99,369 B (97.04 KiB) compressed - `zstd-3` + `delta-binary-packed`; 722,823 B (705.88 KiB) encoded; 40.304813x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict`

Snappy:
1. 52,198 B (50.97 KiB) compressed - `snappy` + `rle-dict`; 73,898 B (72.17 KiB) encoded; 76.728016x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-plain`
2. 126,255 B (123.30 KiB) compressed - `snappy` + `delta-binary-packed`; 716,541 B (699.75 KiB) encoded; 31.721904x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`
3. 232,132 B (226.69 KiB) compressed - `snappy` + `plain`; 4,003,712 B (3.82 MiB) encoded; 17.253326x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## UTMCampaign (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 10 / 18`
- Page cardinality per row group min/median/max of mins: `1 / 10 / 18`; of maxes: `1 / 10 / 18`
- Value length min/median/max: `0 / 0 / 66` bytes
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 43 / 66`

![Row-group cardinality](column_shape_stats/images/utmcampaign_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/utmcampaign_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/utmcampaign_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/utmcampaign_value_length.svg)


Compressed overall:
1. 27,122 B (26.49 KiB) compressed - `zstd-3` + `rle-dict`; 36,919 B (36.05 KiB) encoded; 151.161419x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 29,663 B (28.97 KiB) compressed - `zstd-3` + `plain`; 4,097,338 B (3.91 MiB) encoded; 138.212588x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 31,688 B (30.95 KiB) compressed - `snappy` + `rle-dict`; 37,205 B (36.33 KiB) encoded; 129.380207x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
4. 47,487 B (46.37 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 254,182 B (248.22 KiB) encoded; 86.335208x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
5. 53,867 B (52.60 KiB) compressed - `zstd-3` + `delta-byte-array`; 325,366 B (317.74 KiB) encoded; 76.109678x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`

ZSTD:
1. 27,122 B (26.49 KiB) compressed - `zstd-3` + `rle-dict`; 36,919 B (36.05 KiB) encoded; 151.161419x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 29,663 B (28.97 KiB) compressed - `zstd-3` + `plain`; 4,097,338 B (3.91 MiB) encoded; 138.212588x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 47,487 B (46.37 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 254,182 B (248.22 KiB) encoded; 86.335208x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
4. 53,867 B (52.60 KiB) compressed - `zstd-3` + `delta-byte-array`; 325,366 B (317.74 KiB) encoded; 76.109678x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`

Snappy:
1. 31,688 B (30.95 KiB) compressed - `snappy` + `rle-dict`; 37,205 B (36.33 KiB) encoded; 129.380207x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 57,171 B (55.83 KiB) compressed - `snappy` + `delta-length-byte-array`; 252,420 B (246.50 KiB) encoded; 71.711182x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 67,966 B (66.37 KiB) compressed - `snappy` + `delta-byte-array`; 328,459 B (320.76 KiB) encoded; 60.321337x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 230,205 B (224.81 KiB) compressed - `snappy` + `plain`; 4,097,894 B (3.91 MiB) encoded; 17.809344x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-rle-dict-ts-delta-binary-packed`

## UTMContent (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 3 / 25`
- Page cardinality per row group min/median/max of mins: `1 / 3 / 25`; of maxes: `1 / 3 / 25`
- Value length min/median/max: `0 / 0 / 62` bytes
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 7 / 62`

![Row-group cardinality](column_shape_stats/images/utmcontent_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/utmcontent_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/utmcontent_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/utmcontent_value_length.svg)


Compressed overall:
1. 13,959 B (13.63 KiB) compressed - `zstd-3` + `plain`; 4,016,537 B (3.83 MiB) encoded; 287.852353x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed`
2. 14,567 B (14.23 KiB) compressed - `snappy` + `rle-dict`; 14,840 B (14.49 KiB) encoded; 275.837921x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`
3. 14,839 B (14.49 KiB) compressed - `zstd-3` + `rle-dict`; 14,752 B (14.41 KiB) encoded; 270.781791x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed`
4. 18,006 B (17.58 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 94,767 B (92.55 KiB) encoded; 223.155115x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
5. 20,606 B (20.12 KiB) compressed - `zstd-3` + `delta-byte-array`; 149,975 B (146.46 KiB) encoded; 194.998107x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

ZSTD:
1. 13,959 B (13.63 KiB) compressed - `zstd-3` + `plain`; 4,016,537 B (3.83 MiB) encoded; 287.852353x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-plain-ts-delta-binary-packed`
2. 14,839 B (14.49 KiB) compressed - `zstd-3` + `rle-dict`; 14,752 B (14.41 KiB) encoded; 270.781791x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-rle-dict-date-plain-ts-delta-binary-packed`
3. 18,006 B (17.58 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 94,767 B (92.55 KiB) encoded; 223.155115x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-delta-binary-packed`
4. 20,606 B (20.12 KiB) compressed - `zstd-3` + `delta-byte-array`; 149,975 B (146.46 KiB) encoded; 194.998107x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 14,567 B (14.23 KiB) compressed - `snappy` + `rle-dict`; 14,840 B (14.49 KiB) encoded; 275.837921x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-delta-binary-packed`
2. 21,973 B (21.46 KiB) compressed - `snappy` + `delta-length-byte-array`; 94,225 B (92.02 KiB) encoded; 182.866746x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 27,353 B (26.71 KiB) compressed - `snappy` + `delta-byte-array`; 149,463 B (145.96 KiB) encoded; 146.899097x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 212,173 B (207.20 KiB) compressed - `snappy` + `plain`; 4,016,717 B (3.83 MiB) encoded; 18.937994x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-delta-binary-packed`

## UTMMedium (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 3 / 5`
- Page cardinality per row group min/median/max of mins: `1 / 3 / 5`; of maxes: `1 / 3 / 5`
- Value length min/median/max: `0 / 0 / 16` bytes
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 4 / 16`

![Row-group cardinality](column_shape_stats/images/utmmedium_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/utmmedium_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/utmmedium_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/utmmedium_value_length.svg)


Compressed overall:
1. 14,496 B (14.16 KiB) compressed - `zstd-3` + `rle-dict`; 16,556 B (16.17 KiB) encoded; 277.401214x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-rle-dict`
2. 15,832 B (15.46 KiB) compressed - `snappy` + `rle-dict`; 16,565 B (16.18 KiB) encoded; 253.992420x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
3. 16,944 B (16.55 KiB) compressed - `zstd-3` + `plain`; 4,019,915 B (3.83 MiB) encoded; 237.323418x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
4. 21,863 B (21.35 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 133,463 B (130.33 KiB) encoded; 183.927549x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-plain`
5. 27,882 B (27.23 KiB) compressed - `zstd-3` + `delta-byte-array`; 213,000 B (208.01 KiB) encoded; 144.222366x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 14,496 B (14.16 KiB) compressed - `zstd-3` + `rle-dict`; 16,556 B (16.17 KiB) encoded; 277.401214x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-rle-dict`
2. 16,944 B (16.55 KiB) compressed - `zstd-3` + `plain`; 4,019,915 B (3.83 MiB) encoded; 237.323418x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 21,863 B (21.35 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 133,463 B (130.33 KiB) encoded; 183.927549x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-plain`
4. 27,882 B (27.23 KiB) compressed - `zstd-3` + `delta-byte-array`; 213,000 B (208.01 KiB) encoded; 144.222366x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-plain-ts-plain`

Snappy:
1. 15,832 B (15.46 KiB) compressed - `snappy` + `rle-dict`; 16,565 B (16.18 KiB) encoded; 253.992420x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 28,347 B (27.68 KiB) compressed - `snappy` + `delta-length-byte-array`; 132,761 B (129.65 KiB) encoded; 141.856563x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-delta-binary-packed`
3. 38,746 B (37.84 KiB) compressed - `snappy` + `delta-byte-array`; 213,611 B (208.60 KiB) encoded; 103.783823x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`
4. 214,350 B (209.33 KiB) compressed - `snappy` + `plain`; 4,020,123 B (3.83 MiB) encoded; 18.760009x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`

## UTMSource (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 6 / 8`
- Page cardinality per row group min/median/max of mins: `1 / 6 / 8`; of maxes: `1 / 6 / 8`
- Value length min/median/max: `0 / 0 / 19` bytes
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 16 / 19`

![Row-group cardinality](column_shape_stats/images/utmsource_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/utmsource_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/utmsource_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/utmsource_value_length.svg)


Compressed overall:
1. 18,339 B (17.91 KiB) compressed - `zstd-3` + `rle-dict`; 22,575 B (22.05 KiB) encoded; 221.098261x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 20,260 B (19.79 KiB) compressed - `snappy` + `rle-dict`; 22,681 B (22.15 KiB) encoded; 200.134304x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
3. 21,004 B (20.51 KiB) compressed - `zstd-3` + `plain`; 4,053,175 B (3.87 MiB) encoded; 193.045182x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
4. 33,547 B (32.76 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 187,509 B (183.11 KiB) encoded; 120.866873x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
5. 40,012 B (39.07 KiB) compressed - `zstd-3` + `delta-byte-array`; 267,362 B (261.10 KiB) encoded; 101.337624x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 18,339 B (17.91 KiB) compressed - `zstd-3` + `rle-dict`; 22,575 B (22.05 KiB) encoded; 221.098261x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 21,004 B (20.51 KiB) compressed - `zstd-3` + `plain`; 4,053,175 B (3.87 MiB) encoded; 193.045182x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 33,547 B (32.76 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 187,509 B (183.11 KiB) encoded; 120.866873x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
4. 40,012 B (39.07 KiB) compressed - `zstd-3` + `delta-byte-array`; 267,362 B (261.10 KiB) encoded; 101.337624x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-plain`

Snappy:
1. 20,260 B (19.79 KiB) compressed - `snappy` + `rle-dict`; 22,681 B (22.15 KiB) encoded; 200.134304x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-rle-dict-ts-plain`
2. 41,112 B (40.15 KiB) compressed - `snappy` + `delta-length-byte-array`; 187,202 B (182.81 KiB) encoded; 98.626216x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-delta-binary-packed`
3. 51,686 B (50.47 KiB) compressed - `snappy` + `delta-byte-array`; 267,959 B (261.68 KiB) encoded; 78.449116x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 219,066 B (213.93 KiB) compressed - `snappy` + `plain`; 4,053,439 B (3.87 MiB) encoded; 18.509130x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-rle-dict-ts-delta-binary-packed`

## UTMTerm (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 2 / 75`
- Page cardinality per row group min/median/max of mins: `1 / 2 / 75`; of maxes: `1 / 2 / 75`
- Value length min/median/max: `0 / 0 / 72` bytes
- Value length per row group min/median/max of mins: `0 / 0 / 0`; of maxes: `0 / 16 / 72`

![Row-group cardinality](column_shape_stats/images/utmterm_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/utmterm_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/utmterm_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/utmterm_value_length.svg)


Compressed overall:
1. 15,648 B (15.28 KiB) compressed - `zstd-3` + `plain`; 4,032,689 B (3.85 MiB) encoded; 257.827454x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
2. 15,716 B (15.35 KiB) compressed - `snappy` + `rle-dict`; 16,622 B (16.23 KiB) encoded; 256.711886x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-rle-dict-date-plain-ts-delta-binary-packed`
3. 15,920 B (15.55 KiB) compressed - `zstd-3` + `rle-dict`; 16,405 B (16.02 KiB) encoded; 253.422362x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
4. 19,861 B (19.40 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 128,536 B (125.52 KiB) encoded; 203.135995x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
5. 22,914 B (22.38 KiB) compressed - `zstd-3` + `delta-byte-array`; 189,668 B (185.22 KiB) encoded; 176.070699x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 15,648 B (15.28 KiB) compressed - `zstd-3` + `plain`; 4,032,689 B (3.85 MiB) encoded; 257.827454x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-delta-binary-packed-ts-plain`
2. 15,920 B (15.55 KiB) compressed - `zstd-3` + `rle-dict`; 16,405 B (16.02 KiB) encoded; 253.422362x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-plain-ts-plain`
3. 19,861 B (19.40 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 128,536 B (125.52 KiB) encoded; 203.135995x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
4. 22,914 B (22.38 KiB) compressed - `zstd-3` + `delta-byte-array`; 189,668 B (185.22 KiB) encoded; 176.070699x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`

Snappy:
1. 15,716 B (15.35 KiB) compressed - `snappy` + `rle-dict`; 16,622 B (16.23 KiB) encoded; 256.711886x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-rle-dict-date-plain-ts-delta-binary-packed`
2. 25,137 B (24.55 KiB) compressed - `snappy` + `delta-length-byte-array`; 129,144 B (126.12 KiB) encoded; 160.499821x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-length-byte-array-date-rle-dict-ts-plain`
3. 30,896 B (30.17 KiB) compressed - `snappy` + `delta-byte-array`; 189,246 B (184.81 KiB) encoded; 130.582729x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-plain`
4. 214,528 B (209.50 KiB) compressed - `snappy` + `plain`; 4,032,992 B (3.85 MiB) encoded; 18.806328x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-rle-dict-ts-delta-binary-packed`

## UserAgent (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `11 / 14 / 18`
- Page cardinality per row group min/median/max of mins: `11 / 14 / 18`; of maxes: `11 / 14 / 18`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/useragent_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/useragent_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/useragent_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/useragent_value_length.svg)


Compressed overall:
1. 125,153 B (122.22 KiB) compressed - `zstd-3` + `rle-dict`; 226,591 B (221.28 KiB) encoded; 32.001254x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 133,847 B (130.71 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 29.922621x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-rle-dict`
3. 164,732 B (160.87 KiB) compressed - `snappy` + `rle-dict`; 223,595 B (218.35 KiB) encoded; 24.312538x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
4. 235,012 B (229.50 KiB) compressed - `zstd-3` + `delta-binary-packed`; 686,507 B (670.42 KiB) encoded; 17.041908x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
5. 281,367 B (274.77 KiB) compressed - `snappy` + `delta-binary-packed`; 685,977 B (669.90 KiB) encoded; 14.234267x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 125,153 B (122.22 KiB) compressed - `zstd-3` + `rle-dict`; 226,591 B (221.28 KiB) encoded; 32.001254x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
2. 133,847 B (130.71 KiB) compressed - `zstd-3` + `plain`; 4,003,587 B (3.82 MiB) encoded; 29.922621x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-rle-dict`
3. 235,012 B (229.50 KiB) compressed - `zstd-3` + `delta-binary-packed`; 686,507 B (670.42 KiB) encoded; 17.041908x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`

Snappy:
1. 164,732 B (160.87 KiB) compressed - `snappy` + `rle-dict`; 223,595 B (218.35 KiB) encoded; 24.312538x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
2. 281,367 B (274.77 KiB) compressed - `snappy` + `delta-binary-packed`; 685,977 B (669.90 KiB) encoded; 14.234267x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 342,015 B (334.00 KiB) compressed - `snappy` + `plain`; 4,003,715 B (3.82 MiB) encoded; 11.710168x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

## UserAgentMajor (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `26 / 29 / 31`
- Page cardinality per row group min/median/max of mins: `26 / 29 / 31`; of maxes: `26 / 29 / 31`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/useragentmajor_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/useragentmajor_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/useragentmajor_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/useragentmajor_value_length.svg)


Compressed overall:
1. 154,186 B (150.57 KiB) compressed - `zstd-3` + `plain`; 4,003,585 B (3.82 MiB) encoded; 25.975445x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-plain`
2. 177,456 B (173.30 KiB) compressed - `zstd-3` + `rle-dict`; 275,570 B (269.11 KiB) encoded; 22.569257x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 213,965 B (208.95 KiB) compressed - `snappy` + `rle-dict`; 275,877 B (269.41 KiB) encoded; 18.718248x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
4. 259,009 B (252.94 KiB) compressed - `zstd-3` + `delta-binary-packed`; 661,075 B (645.58 KiB) encoded; 15.462976x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain`
5. 301,197 B (294.14 KiB) compressed - `snappy` + `delta-binary-packed`; 661,130 B (645.63 KiB) encoded; 13.297111x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`

ZSTD:
1. 154,186 B (150.57 KiB) compressed - `zstd-3` + `plain`; 4,003,585 B (3.82 MiB) encoded; 25.975445x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-plain-date-delta-binary-packed-ts-plain`
2. 177,456 B (173.30 KiB) compressed - `zstd-3` + `rle-dict`; 275,570 B (269.11 KiB) encoded; 22.569257x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
3. 259,009 B (252.94 KiB) compressed - `zstd-3` + `delta-binary-packed`; 661,075 B (645.58 KiB) encoded; 15.462976x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain`

Snappy:
1. 213,965 B (208.95 KiB) compressed - `snappy` + `rle-dict`; 275,877 B (269.41 KiB) encoded; 18.718248x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
2. 301,197 B (294.14 KiB) compressed - `snappy` + `delta-binary-packed`; 661,130 B (645.63 KiB) encoded; 13.297111x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-plain-ts-plain`
3. 364,837 B (356.29 KiB) compressed - `snappy` + `plain`; 4,003,717 B (3.82 MiB) encoded; 10.977642x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-rle-dict-ts-plain`

## UserAgentMinor (string)

Column shape stats:
- Parquet type: `STRING`; physical type: `BYTE_ARRAY`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `19 / 22 / 28`
- Page cardinality per row group min/median/max of mins: `19 / 22 / 28`; of maxes: `19 / 22 / 28`
- Value length min/median/max: `2 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `2 / 2 / 2`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/useragentminor_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/useragentminor_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/useragentminor_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/useragentminor_value_length.svg)


Compressed overall:
1. 83,594 B (81.63 KiB) compressed - `zstd-3` + `rle-dict`; 150,305 B (146.78 KiB) encoded; 92.976183x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 103,993 B (101.56 KiB) compressed - `snappy` + `rle-dict`; 150,251 B (146.73 KiB) encoded; 74.738213x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
3. 136,655 B (133.45 KiB) compressed - `zstd-3` + `plain`; 7,770,878 B (7.41 MiB) encoded; 56.874984x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
4. 150,595 B (147.07 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 4,016,573 B (3.83 MiB) encoded; 51.610286x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
5. 178,932 B (174.74 KiB) compressed - `zstd-3` + `delta-byte-array`; 858,898 B (838.77 KiB) encoded; 43.436898x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`

ZSTD:
1. 83,594 B (81.63 KiB) compressed - `zstd-3` + `rle-dict`; 150,305 B (146.78 KiB) encoded; 92.976183x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
2. 136,655 B (133.45 KiB) compressed - `zstd-3` + `plain`; 7,770,878 B (7.41 MiB) encoded; 56.874984x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 150,595 B (147.07 KiB) compressed - `zstd-3` + `delta-length-byte-array`; 4,016,573 B (3.83 MiB) encoded; 51.610286x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
4. 178,932 B (174.74 KiB) compressed - `zstd-3` + `delta-byte-array`; 858,898 B (838.77 KiB) encoded; 43.436898x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`

Snappy:
1. 103,993 B (101.56 KiB) compressed - `snappy` + `rle-dict`; 150,251 B (146.73 KiB) encoded; 74.738213x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
2. 235,276 B (229.76 KiB) compressed - `snappy` + `delta-byte-array`; 859,464 B (839.32 KiB) encoded; 33.034610x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 327,938 B (320.25 KiB) compressed - `snappy` + `delta-length-byte-array`; 4,017,088 B (3.83 MiB) encoded; 23.700367x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-length-byte-array-date-plain-ts-plain`
4. 467,072 B (456.12 KiB) compressed - `snappy` + `plain`; 7,771,080 B (7.41 MiB) encoded; 16.640370x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-plain`

## UserID (int64)

Column shape stats:
- Parquet type: `INT(64,true)`; physical type: `INT64`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `716 / 898 / 1,805`
- Page cardinality per row group min/median/max of mins: `716 / 898 / 1,805`; of maxes: `716 / 898 / 1,805`
- Value length min/median/max: `8 / 8 / 8` bytes
- Value length per row group min/median/max of mins: `8 / 8 / 8`; of maxes: `8 / 8 / 8`

![Row-group cardinality](column_shape_stats/images/userid_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/userid_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/userid_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/userid_value_length.svg)


Compressed overall:
1. 617,840 B (603.36 KiB) compressed - `zstd-3` + `plain`; 8,004,550 B (7.63 MiB) encoded; 12.958557x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 760,398 B (742.58 KiB) compressed - `zstd-3` + `delta-binary-packed`; 4,392,093 B (4.19 MiB) encoded; 10.529111x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict`
3. 944,456 B (922.32 KiB) compressed - `snappy` + `delta-binary-packed`; 4,383,346 B (4.18 MiB) encoded; 8.477171x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-plain`
4. 1,084,714 B (1.03 MiB) compressed - `snappy` + `plain`; 8,004,715 B (7.63 MiB) encoded; 7.381038x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
5. 1,101,572 B (1.05 MiB) compressed - `zstd-3` + `rle-dict`; 1,230,153 B (1.17 MiB) encoded; 7.268081x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

ZSTD:
1. 617,840 B (603.36 KiB) compressed - `zstd-3` + `plain`; 8,004,550 B (7.63 MiB) encoded; 12.958557x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 760,398 B (742.58 KiB) compressed - `zstd-3` + `delta-binary-packed`; 4,392,093 B (4.19 MiB) encoded; 10.529111x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-rle-dict`
3. 1,101,572 B (1.05 MiB) compressed - `zstd-3` + `rle-dict`; 1,230,153 B (1.17 MiB) encoded; 7.268081x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

Snappy:
1. 944,456 B (922.32 KiB) compressed - `snappy` + `delta-binary-packed`; 4,383,346 B (4.18 MiB) encoded; 8.477171x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-plain`
2. 1,084,714 B (1.03 MiB) compressed - `snappy` + `plain`; 8,004,715 B (7.63 MiB) encoded; 7.381038x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 1,120,481 B (1.07 MiB) compressed - `snappy` + `rle-dict`; 1,225,906 B (1.17 MiB) encoded; 7.145427x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-plain-date-plain-ts-rle-dict`

## WatchID (int64)

Column shape stats:
- Parquet type: `INT(64,true)`; physical type: `INT64`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `9,315 / 11,938 / 14,202`
- Page cardinality per row group min/median/max of mins: `9,315 / 11,938 / 14,202`; of maxes: `9,315 / 11,938 / 14,202`
- Value length min/median/max: `8 / 8 / 8` bytes
- Value length per row group min/median/max of mins: `8 / 8 / 8`; of maxes: `8 / 8 / 8`

![Row-group cardinality](column_shape_stats/images/watchid_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/watchid_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/watchid_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/watchid_value_length.svg)


Compressed overall:
1. 7,987,755 B (7.62 MiB) compressed - `zstd-3` + `delta-binary-packed`; 7,989,449 B (7.62 MiB) encoded; 1.002323x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-rle-dict-date-plain-ts-plain`
2. 7,989,242 B (7.62 MiB) compressed - `snappy` + `delta-binary-packed`; 7,988,829 B (7.62 MiB) encoded; 1.002137x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
3. 8,005,128 B (7.63 MiB) compressed - `snappy` + `plain`; 8,004,715 B (7.63 MiB) encoded; 1.000148x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
4. 8,005,353 B (7.63 MiB) compressed - `zstd-3` + `plain`; 8,004,555 B (7.63 MiB) encoded; 1.000120x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
5. 9,826,284 B (9.37 MiB) compressed - `snappy` + `rle-dict`; 9,825,517 B (9.37 MiB) encoded; 0.814785x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`

ZSTD:
1. 7,987,755 B (7.62 MiB) compressed - `zstd-3` + `delta-binary-packed`; 7,989,449 B (7.62 MiB) encoded; 1.002323x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-rle-dict-date-plain-ts-plain`
2. 8,005,353 B (7.63 MiB) compressed - `zstd-3` + `plain`; 8,004,555 B (7.63 MiB) encoded; 1.000120x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 9,835,734 B (9.38 MiB) compressed - `zstd-3` + `rle-dict`; 9,834,361 B (9.38 MiB) encoded; 0.814002x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict`

Snappy:
1. 7,989,242 B (7.62 MiB) compressed - `snappy` + `delta-binary-packed`; 7,988,829 B (7.62 MiB) encoded; 1.002137x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
2. 8,005,128 B (7.63 MiB) compressed - `snappy` + `plain`; 8,004,715 B (7.63 MiB) encoded; 1.000148x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 9,826,284 B (9.37 MiB) compressed - `snappy` + `rle-dict`; 9,825,517 B (9.37 MiB) encoded; 0.814785x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-rle-dict`

## WindowClientHeight (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `318 / 435 / 575`
- Page cardinality per row group min/median/max of mins: `318 / 435 / 575`; of maxes: `318 / 435 / 575`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/windowclientheight_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/windowclientheight_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/windowclientheight_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/windowclientheight_value_length.svg)


Compressed overall:
1. 319,342 B (311.86 KiB) compressed - `zstd-3` + `plain`; 4,003,650 B (3.82 MiB) encoded; 12.541573x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict`
2. 551,321 B (538.40 KiB) compressed - `snappy` + `plain`; 4,003,760 B (3.82 MiB) encoded; 7.264463x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 551,652 B (538.72 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,249,475 B (1.19 MiB) encoded; 7.260104x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain`
4. 583,325 B (569.65 KiB) compressed - `zstd-3` + `rle-dict`; 750,345 B (732.76 KiB) encoded; 6.865900x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
5. 601,354 B (587.26 KiB) compressed - `snappy` + `rle-dict`; 749,835 B (732.26 KiB) encoded; 6.660055x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

ZSTD:
1. 319,342 B (311.86 KiB) compressed - `zstd-3` + `plain`; 4,003,650 B (3.82 MiB) encoded; 12.541573x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-rle-dict`
2. 551,652 B (538.72 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,249,475 B (1.19 MiB) encoded; 7.260104x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-rle-dict-ts-plain`
3. 583,325 B (569.65 KiB) compressed - `zstd-3` + `rle-dict`; 750,345 B (732.76 KiB) encoded; 6.865900x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 551,321 B (538.40 KiB) compressed - `snappy` + `plain`; 4,003,760 B (3.82 MiB) encoded; 7.264463x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 601,354 B (587.26 KiB) compressed - `snappy` + `rle-dict`; 749,835 B (732.26 KiB) encoded; 6.660055x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
3. 616,032 B (601.59 KiB) compressed - `snappy` + `delta-binary-packed`; 1,249,759 B (1.19 MiB) encoded; 6.501368x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-plain`

## WindowClientWidth (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `241 / 306 / 374`
- Page cardinality per row group min/median/max of mins: `241 / 306 / 374`; of maxes: `241 / 306 / 374`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/windowclientwidth_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/windowclientwidth_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/windowclientwidth_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/windowclientwidth_value_length.svg)


Compressed overall:
1. 305,724 B (298.56 KiB) compressed - `zstd-3` + `plain`; 4,003,590 B (3.82 MiB) encoded; 13.100221x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 471,075 B (460.03 KiB) compressed - `zstd-3` + `rle-dict`; 695,366 B (679.07 KiB) encoded; 8.501941x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 512,791 B (500.77 KiB) compressed - `snappy` + `rle-dict`; 695,742 B (679.44 KiB) encoded; 7.810301x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
4. 516,116 B (504.02 KiB) compressed - `snappy` + `plain`; 4,003,747 B (3.82 MiB) encoded; 7.759984x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`
5. 573,659 B (560.21 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,355,674 B (1.29 MiB) encoded; 6.981590x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-plain`

ZSTD:
1. 305,724 B (298.56 KiB) compressed - `zstd-3` + `plain`; 4,003,590 B (3.82 MiB) encoded; 13.100221x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-delta-binary-packed`
2. 471,075 B (460.03 KiB) compressed - `zstd-3` + `rle-dict`; 695,366 B (679.07 KiB) encoded; 8.501941x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-delta-length-byte-array-date-plain-ts-plain`
3. 573,659 B (560.21 KiB) compressed - `zstd-3` + `delta-binary-packed`; 1,355,674 B (1.29 MiB) encoded; 6.981590x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-delta-length-byte-array-date-plain-ts-plain`

Snappy:
1. 512,791 B (500.77 KiB) compressed - `snappy` + `rle-dict`; 695,742 B (679.44 KiB) encoded; 7.810301x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-rle-dict-ts-delta-binary-packed`
2. 516,116 B (504.02 KiB) compressed - `snappy` + `plain`; 4,003,747 B (3.82 MiB) encoded; 7.759984x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`
3. 646,095 B (630.95 KiB) compressed - `snappy` + `delta-binary-packed`; 1,354,763 B (1.29 MiB) encoded; 6.198859x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`

## WindowName (int32)

Column shape stats:
- Parquet type: `INT(32,true)`; physical type: `INT32`
- Sorted ascending across written rows: `false`; sorted descending: `false`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 6 / 3,150`
- Page cardinality per row group min/median/max of mins: `1 / 6 / 3,150`; of maxes: `1 / 6 / 3,150`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/windowname_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/windowname_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/windowname_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/windowname_value_length.svg)


Compressed overall:
1. 70,803 B (69.14 KiB) compressed - `zstd-3` + `plain`; 4,003,577 B (3.82 MiB) encoded; 56.566134x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`
2. 121,181 B (118.34 KiB) compressed - `zstd-3` + `delta-binary-packed`; 274,063 B (267.64 KiB) encoded; 33.050165x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain`
3. 130,335 B (127.28 KiB) compressed - `snappy` + `delta-binary-packed`; 274,877 B (268.43 KiB) encoded; 30.728906x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-rle-dict`
4. 148,703 B (145.22 KiB) compressed - `zstd-3` + `rle-dict`; 171,358 B (167.34 KiB) encoded; 26.933229x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-delta-binary-packed`
5. 150,771 B (147.24 KiB) compressed - `snappy` + `rle-dict`; 171,344 B (167.33 KiB) encoded; 26.563809x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`

ZSTD:
1. 70,803 B (69.14 KiB) compressed - `zstd-3` + `plain`; 4,003,577 B (3.82 MiB) encoded; 56.566134x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-rle-dict`
2. 121,181 B (118.34 KiB) compressed - `zstd-3` + `delta-binary-packed`; 274,063 B (267.64 KiB) encoded; 33.050165x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-plain-ts-plain`
3. 148,703 B (145.22 KiB) compressed - `zstd-3` + `rle-dict`; 171,358 B (167.34 KiB) encoded; 26.933229x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-rle-dict-date-rle-dict-ts-delta-binary-packed`

Snappy:
1. 130,335 B (127.28 KiB) compressed - `snappy` + `delta-binary-packed`; 274,877 B (268.43 KiB) encoded; 30.728906x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-rle-dict-ts-rle-dict`
2. 150,771 B (147.24 KiB) compressed - `snappy` + `rle-dict`; 171,344 B (167.33 KiB) encoded; 26.563809x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-rle-dict-date-delta-binary-packed-ts-plain`
3. 295,106 B (288.19 KiB) compressed - `snappy` + `plain`; 4,003,718 B (3.82 MiB) encoded; 13.571571x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-delta-binary-packed-ts-delta-binary-packed`

## WithHash (int16)

Column shape stats:
- Parquet type: `INT(16,true)`; physical type: `INT32`
- Sorted ascending across written rows: `true`; sorted descending: `true`
- Row groups: `79`; pages: `79`
- Row-group cardinality min/median/max: `1 / 1 / 1`
- Page cardinality per row group min/median/max of mins: `1 / 1 / 1`; of maxes: `1 / 1 / 1`
- Value length min/median/max: `4 / 4 / 4` bytes
- Value length per row group min/median/max of mins: `4 / 4 / 4`; of maxes: `4 / 4 / 4`

![Row-group cardinality](column_shape_stats/images/withhash_row_group_cardinality.svg)

![Page cardinality min/max per row group](column_shape_stats/images/withhash_page_cardinality.svg)

![Page min/max distribution](column_shape_stats/images/withhash_page_bounds.svg)

![Value length min/max per row group](column_shape_stats/images/withhash_value_length.svg)


Compressed overall:
1. 4,228 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
2. 4,932 B (4.82 KiB) compressed - `zstd-3` + `delta-binary-packed`; 43,044 B (42.04 KiB) encoded; 812.054745x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 5,273 B (5.15 KiB) compressed - `snappy` + `rle-dict`; 5,041 B (4.92 KiB) encoded; 759.539920x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
4. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`
5. 6,203 B (6.06 KiB) compressed - `snappy` + `delta-binary-packed`; 43,179 B (42.17 KiB) encoded; 645.664034x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`

ZSTD:
1. 4,228 B (4.13 KiB) compressed - `zstd-3` + `plain`; 4,003,527 B (3.82 MiB) encoded; 947.269158x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-plain-str-delta-length-byte-array-date-delta-binary-packed-ts-plain`
2. 4,932 B (4.82 KiB) compressed - `zstd-3` + `delta-binary-packed`; 43,044 B (42.04 KiB) encoded; 812.054745x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-delta-binary-packed-str-plain-date-delta-binary-packed-ts-delta-binary-packed`
3. 5,975 B (5.83 KiB) compressed - `zstd-3` + `rle-dict`; 4,949 B (4.83 KiB) encoded; 670.301925x post-compression ratio; experiment `rows-1000000-comp-zstd-3-int-rle-dict-str-plain-date-plain-ts-plain`

Snappy:
1. 5,273 B (5.15 KiB) compressed - `snappy` + `rle-dict`; 5,041 B (4.92 KiB) encoded; 759.539920x post-compression ratio; experiment `rows-1000000-comp-snappy-int-rle-dict-str-delta-byte-array-date-plain-ts-delta-binary-packed`
2. 6,203 B (6.06 KiB) compressed - `snappy` + `delta-binary-packed`; 43,179 B (42.17 KiB) encoded; 645.664034x post-compression ratio; experiment `rows-1000000-comp-snappy-int-delta-binary-packed-str-delta-byte-array-date-delta-binary-packed-ts-plain`
3. 204,060 B (199.28 KiB) compressed - `snappy` + `plain`; 4,003,712 B (3.82 MiB) encoded; 19.626845x post-compression ratio; experiment `rows-1000000-comp-snappy-int-plain-str-delta-byte-array-date-plain-ts-delta-binary-packed`

