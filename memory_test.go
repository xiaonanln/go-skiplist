package skiplist

import (
	"math/rand"
	"os"
	"runtime/pprof"
	"testing"

	handu_skiplist "github.com/huandu/skiplist"
	"github.com/petar/GoLLRB/llrb"
)

const (
	memoryTestSize = 10000000
)

func TestMemorySkipList(t *testing.T) {
	fd, err := os.OpenFile("sl.heap.pprof", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	tree := New()
	for i := 0; i < memoryTestSize; i++ {
		item := slIntInt{i, i}
		tree.ReplaceOrInsert(item)
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
		item := llrbIntInt{i, i}
		tree.ReplaceOrInsert(item)
	}

	pprof.WriteHeapProfile(fd)
}

func TestMemoryHanduSkipList(t *testing.T) {
	fd, err := os.OpenFile("hdsl.heap.pprof", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	tree := handu_skiplist.New(handu_skiplist.Int)
	for i := 0; i < memoryTestSize; i++ {
		tree.Set(i, i)
	}

	pprof.WriteHeapProfile(fd)
}

func TestMemoryLeak(t *testing.T) {
	t.SkipNow()
	tree := New()
	for {
		tree.Delete(Int(rand.Intn(10000)))
		tree.ReplaceOrInsert(Int(rand.Intn(10000)))
	}
}
