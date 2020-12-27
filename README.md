# benchmarks

This repository contains benchmarks for a variety of algorithms.

## `list_test.go`

`list_test.go` contains benchmarks for removing a single element from
int64 Slices and linked lists. The benchmarks runs over a variety of
list sizes and removal points.

Sample results:

```
goos: linux
goarch: amd64
BenchmarkSliceRemoval/Size=100000,CutPoint=0.10-8                  43477             27543 ns/op
BenchmarkSliceRemoval/Size=100000,CutPoint=0.50-8                  80136             14942 ns/op
BenchmarkSliceRemoval/Size=100000,CutPoint=0.90-8                 493446              2400 ns/op
BenchmarkSliceRemoval/Size=1000000,CutPoint=0.10-8                  1632            735074 ns/op
BenchmarkSliceRemoval/Size=1000000,CutPoint=0.50-8                  2922            419355 ns/op
BenchmarkSliceRemoval/Size=1000000,CutPoint=0.90-8                 38446             31898 ns/op
BenchmarkSliceRemoval/Size=10000000,CutPoint=0.10-8                  127           9836437 ns/op
BenchmarkSliceRemoval/Size=10000000,CutPoint=0.50-8                  208           5226091 ns/op
BenchmarkSliceRemoval/Size=10000000,CutPoint=0.90-8                 1370            906357 ns/op
BenchmarkSliceRemoval/Size=100000000,CutPoint=0.10-8                  14          77721050 ns/op
BenchmarkSliceRemoval/Size=100000000,CutPoint=0.50-8                  27          43314248 ns/op
BenchmarkSliceRemoval/Size=100000000,CutPoint=0.90-8                 138           8593500 ns/op
BenchmarkSliceRemoval/Size=1000000000,CutPoint=0.10-8                  1        1269305388 ns/op
BenchmarkSliceRemoval/Size=1000000000,CutPoint=0.50-8                  3         416622873 ns/op
BenchmarkSliceRemoval/Size=1000000000,CutPoint=0.90-8                 12          86023684 ns/op
BenchmarkListRemoval/Size=100000,CutPoint=0.10-8                  101442             11725 ns/op
BenchmarkListRemoval/Size=100000,CutPoint=0.50-8                   20371             58877 ns/op
BenchmarkListRemoval/Size=100000,CutPoint=0.90-8                   10000            107385 ns/op
BenchmarkListRemoval/Size=1000000,CutPoint=0.10-8                   9968            116967 ns/op
BenchmarkListRemoval/Size=1000000,CutPoint=0.50-8                   1836            628306 ns/op
BenchmarkListRemoval/Size=1000000,CutPoint=0.90-8                    933           1290872 ns/op
BenchmarkListRemoval/Size=10000000,CutPoint=0.10-8                   842           1404688 ns/op
BenchmarkListRemoval/Size=10000000,CutPoint=0.50-8                   153           7616816 ns/op
BenchmarkListRemoval/Size=10000000,CutPoint=0.90-8                    85          14027545 ns/op
BenchmarkListRemoval/Size=100000000,CutPoint=0.10-8                   75          15572063 ns/op
BenchmarkListRemoval/Size=100000000,CutPoint=0.50-8                   14          77940169 ns/op
BenchmarkListRemoval/Size=100000000,CutPoint=0.90-8                    8         141617153 ns/op
BenchmarkListRemoval/Size=1000000000,CutPoint=0.10-8                   7         162538920 ns/op
BenchmarkListRemoval/Size=1000000000,CutPoint=0.50-8                   2         883383742 ns/op
BenchmarkListRemoval/Size=1000000000,CutPoint=0.90-8                   1        1487546397 ns/op
```
