# go-skiplist
Skip List implementation in Golang

SkipList has almost same interface as red-black tree in [github.com/petar/GoLLRB/llrb](https://github.com/petar/GoLLRB/llrb).

**Thanks to petar, I used LLRB test files directly, and most test cases just work without modification**

SkipList outperforms LLRB significantly.

**LLRB:**
```bash
goos: windows
goarch: amd64
pkg: github.com/petar/GoLLRB/llrb
BenchmarkInsert-8      	 1000000	      1852 ns/op
BenchmarkDelete-8      	 1000000	      1918 ns/op
BenchmarkDeleteMin-8   	 1000000	      1338 ns/op
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

**SkipList of level 16:**  
```bash
goos: windows
goarch: amd64
pkg: github.com/xiaonanln/go-skiplist
BenchmarkInsert-8      	 2000000	      1140 ns/op
BenchmarkDelete-8      	 2000000	       607 ns/op
BenchmarkDeleteMin-8   	30000000	        49.1 ns/op
```
