package skiplist

import (
	"os"
	"runtime/pprof"
	"testing"

	"github.com/petar/GoLLRB/llrb"
)

const (
	cpuTestSize = 1000000
)

func TestCPUSkipList(t *testing.T) {
	fd, err := os.OpenFile("skiplist.cpu.pprof", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(fd)
	tree := New()
	for i := 0; i < cpuTestSize; i++ {
		tree.ReplaceOrInsert(Int(i))
	}
	pprof.StopCPUProfile()
}

func TestCPULLRB(t *testing.T) {
	fd, err := os.OpenFile("llrb.cpu.pprof", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	pprof.StartCPUProfile(fd)
	tree := llrb.New()
	for i := 0; i < cpuTestSize; i++ {
		tree.ReplaceOrInsert(llrb.Int(i))
	}
	pprof.StopCPUProfile()
}
