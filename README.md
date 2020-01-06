# LUBM Benchmarks

- [LUBM homepage](http://swat.cse.lehigh.edu/projects/lubm/)
- [Queries in SVG format](http://swat.cse.lehigh.edu/projects/lubm/lubm.svg)
- [AllegroGraph results (archive)](https://web.archive.org/web/20090208165243/http://agraph.franz.com/allegrograph/agraph_bench_lubm50.lhtml)
- [Stardog results](https://docs.google.com/spreadsheets/d/1oHSWX_0ChZ61ofipZ1CMsW7OhyujioR28AfHzU9d56k/pubhtml#)

## Utils

Generate data

```
java edu.lehigh.swat.bench.uba.Generator -univ 1 -index 0 -seed 0 -onto http://swat.cse.lehigh.edu/onto/univ-bench.owl
```

`convert.sh`

```
rapper -i rdfxml -o nquads -q $1  > $1.nt
curl --data-binary @$1.nt -H 'Content-Type: application/n-quads' localhost:8086/
```

Ingest data

```
find *.owl -maxdepth 1 -type f -exec ./convert.sh {} \;
```

Count triples

```
find *.owl.nt | xargs wc -l
```

first: 2019/12/12 10:34:53 Message: /ipfs/bafybeidmxqxhplipsnnrwmjoqccs4zqwvolp4xvxxzyvhjeufcyqvx7f4y
last : 2019/12/12 10:58:22 Handled message in 538.707755ms

## Results

Results measured on an other-wise empty database on a 2017 Macbook Pro with 16 GB 2133 MHz LPDDR3 and a 2.9 GHz Intel Core i7, with Go 1.13 and Badger 1.6.

### LUBM1

103,104 total triples

```
curl --data-binary @query1.json -H 'Content-Type: application/ld+json' localhost:8086/
```

|        | avg            | 1        | 2        | 3        | 4        | 5        | 6        | 7        | 8        | 9        | 10       |
| ------ | -------------- | -------- | -------- | -------- | -------- | -------- | -------- | -------- | -------- | -------- | -------- |
| query1 | **0.394347ms** | 0.362919 | 0.288016 | 0.540329 | 0.415967 | 0.392153 | 0.340841 | 0.333268 | 0.558645 | 0.30682  | 0.404513 |
| query2 | **0.463441ms** | 0.558941 | 0.469040 | 0.462761 | 0.387005 | 0.643515 | 0.456333 | 0.415154 | 0.472529 | 0.381791 | 0.387346 |
| query3 | **0.158669ms** | 0.214564 | 0.166385 | 0.110453 | 0.141024 | 0.126236 | 0.209382 | 0.177337 | 0.135050 | 0.132056 | 0.174203 |

### LUBM50

6,890,640 total triples

|        | avg            | 1        | 2        | 3        | 4        | 5        | 6        | 7        | 8        | 9        | 10       |
| ------ | -------------- | -------- | -------- | -------- | -------- | -------- | -------- | -------- | -------- | -------- | -------- |
| query1 | **0.589180ms** | 0.559499 | 0.581531 | 0.534952 | 0.560218 | 0.587980 | 0.736740 | 0.536250 | 0.548940 | 0.652013 | 0.593672 |
| query2 | **2.648918ms** | 2.382153 | 3.141992 | 2.337339 | 2.658253 | 2.695001 | 2.360368 | 2.193803 | 2.70467  | 3.040395 | 2.975206 |
| query3 | **0.168975ms** | 0.128307 | 0.166007 | 0.169277 | 0.149976 | 0.15824  | 0.155821 | 0.207412 | 0.187419 | 0.173274 | 0.194018 |

### LUBM100

13,879,970 total triples

|        | avg             | 1        | 2        | 3        | 4        | 5        | 6        | 7        | 8        | 9        | 10       |
| ------ | --------------- | -------- | -------- | -------- | -------- | -------- | -------- | -------- | -------- | -------- | -------- |
| query1 | **0.5209887ms** | 0.445630 | 0.576025 | 0.408285 | 0.621839 | 0.484363 | 0.570337 | 0.477194 | 0.592727 | 0.499838 | 0.533649 |
| query2 | **1.8194817ms** | 1.645319 | 2.145748 | 2.054479 | 1.500583 | 1.436753 | 1.585419 | 1.886467 | 2.120749 | 2.381642 | 1.437658 |
| query3 | **0.1765463ms** | 0.181701 | 0.187435 | 0.221821 | 0.159671 | 0.156800 | 0.169522 | 0.196278 | 0.179919 | 0.159784 | 0.152532 |

---

lubm100
query 1:

- 10.347774ms
- 01.417844ms
- 01.682618ms
- 01.490160ms
- 01.231483ms

query 2:

- 5.889972ms
- 3.825925ms
- 4.004241ms
-

GOMAXPROCS=128
