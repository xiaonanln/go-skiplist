package skiplist

import (
	"os"
	"runtime/pprof"
	"testing"

	"github.com/petar/GoLLRB/llrb"
)

const (
	memoryTestSize = 1000000
)

func TestMemorySkipList(t *testing.T) {
	fd, err := os.OpenFile("skiplist.heap.pprof", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	tree := New()
	for i := 0; i < memoryTestSize; i++ {
		tree.ReplaceOrInsert(Int(i))
	}
	pprof.WriteHeapProfile(fd)
}

func TestMemoryLLRB(t *testing.T) {
	fd, err := os.OpenFile("llrb.heap.pprof", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	tree := llrb.New()
	for i := 0; i < memoryTestSize; i++ {
		tree.ReplaceOrInsert(llrb.Int(i))
	}

	pprof.WriteHeapProfile(fd)
}
