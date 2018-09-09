# go-skiplist
Skip List implementation in Golang.

## Installation
`go get github.com/xiaonanln/go-skiplist`

## Overview
SkipList has almost same interface as red-black tree in [github.com/petar/GoLLRB/llrb](https://github.com/petar/GoLLRB).

**Thanks to petar, I used LLRB test files directly, and most test cases just work without modification**

## Performance
Current implementation of SkipList has comparable performance with LLRB, except for `DeleteMin` 
which SkipList outperforms LLRB dramatically.

**LLRB:**
```bash
goos: windows
goarch: amd64
pkg: github.com/petar/GoLLRB/llrb
BenchmarkInsert-8      	 2000000	       915 ns/op
BenchmarkDelete-8      	 2000000	       946 ns/op
BenchmarkDeleteMin-8   	 3000000	       585 ns/op
```

**SkipList of level 16:**  
```bash
goos: windows
goarch: amd64
pkg: github.com/xiaonanln/go-skiplist
BenchmarkInsert-8      	 2000000	       728 ns/op
BenchmarkDelete-8      	 3000000	       502 ns/op
BenchmarkDeleteMin-8   	50000000	        40.3 ns/op
```

**SkipList of level 24:**  
```bash
goos: windows
goarch: amd64
pkg: github.com/xiaonanln/go-skiplist
BenchmarkInsert-8      	 2000000	      1140 ns/op
BenchmarkDelete-8      	 2000000	       607 ns/op
BenchmarkDeleteMin-8   	30000000	        49.1 ns/op
```
