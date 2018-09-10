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
goos: linux
goarch: amd64
pkg: github.com/petar/GoLLRB/llrb
BenchmarkInsert-32       	 1000000	      1269 ns/op
BenchmarkDelete-32       	 1000000	      1261 ns/op
BenchmarkDeleteMin-32    	 2000000	       800 ns/op
```

**SkipList:**  
```bash
goos: linux
goarch: amd64
pkg: github.com/xiaonanln/go-skiplist
BenchmarkInsert-32         	 2000000	       773 ns/op
BenchmarkDelete-32         	 3000000	       547 ns/op
BenchmarkDeleteMin-32      	50000000	        27.4 ns/op
```
