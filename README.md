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
```
goos: linux
goarch: amd64
pkg: github.com/petar/GoLLRB/llrb
BenchmarkInsert-32       	 1000000	      1269 ns/op
BenchmarkDelete-32       	 1000000	      1261 ns/op
BenchmarkDeleteMin-32    	 2000000	       800 ns/op
```

**SkipList:**  
```
goos: linux
goarch: amd64
pkg: github.com/xiaonanln/go-skiplist
BenchmarkInsert-32         	 2000000	       773 ns/op
BenchmarkDelete-32         	 3000000	       547 ns/op
BenchmarkDeleteMin-32      	50000000	        27.4 ns/op
```

## RAM consumption

With significant optimization effort, the current RAM consumption of SkipList `(p=0.25)` is more or less same with LLRB. 

**LLRB:**
```
Showing nodes accounting for 24.62MB, 99.63% of 24.71MB total
Dropped 62 nodes (cum <= 0.12MB)
      flat  flat%   sum%        cum   cum%
   21.10MB 85.39% 85.39%    21.10MB 85.39%  github.com/petar/GoLLRB/llrb.newNode (inline)
    3.52MB 14.23% 99.63%    24.62MB 99.63%  github.com/xiaonanln/go-skiplist.TestMemoryLLRB
         0     0% 99.63%    21.10MB 85.39%  github.com/petar/GoLLRB/llrb.(*LLRB).ReplaceOrInsert
         0     0% 99.63%    21.10MB 85.39%  github.com/petar/GoLLRB/llrb.(*LLRB).replaceOrInsert
```

**SkipList with probability `p = 0.25`:**
```
Showing nodes accounting for 23.07MB, 98.96% of 23.31MB total
Dropped 71 nodes (cum <= 0.12MB)
      flat  flat%   sum%        cum   cum%
   13.56MB 58.15% 58.15%    13.56MB 58.15%  github.com/xiaonanln/go-skiplist.glob..func1
    4.52MB 19.38% 77.53%    23.21MB 99.56%  github.com/xiaonanln/go-skiplist.TestMemorySkipList
    3.39MB 14.55% 92.08%     3.39MB 14.55%  github.com/xiaonanln/go-skiplist.glob..func2
    1.27MB  5.45% 97.53%     1.27MB  5.45%  github.com/xiaonanln/go-skiplist.glob..func3
    0.32MB  1.35% 98.88%     0.32MB  1.35%  github.com/xiaonanln/go-skiplist.glob..func4
    0.02MB 0.075% 98.96%    18.69MB 80.17%  github.com/xiaonanln/go-skiplist.(*SkipList).ReplaceOrInsert
         0     0% 98.96%    18.67MB 80.10%  github.com/xiaonanln/go-skiplist.allocNode
```
