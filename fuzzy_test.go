package skiplist

import (
	"math/rand"
	"testing"

	handu_skiplist "github.com/huandu/skiplist"
	"github.com/petar/GoLLRB/llrb"
)

const (
	TestSize = 1000000
)

type slIntInt struct {
	key, val int
}

func (ii slIntInt) Less(other Item) bool {
	return ii.key < other.(slIntInt).key
}

type llrbIntInt struct {
	key, val int
}

func (ii llrbIntInt) Less(other llrb.Item) bool {
	return ii.key < other.(llrbIntInt).key
}

func TestFuzzyLLRB(t *testing.T) {
	rand.Seed(1)
	tree := llrb.New()
	for i := 0; i < TestSize; i++ {
		v := rand.Intn(100000)
		item := llrbIntInt{v, v}
		if tree.Delete(item) == nil {
			tree.ReplaceOrInsert(item)
		}
	}
}

func TestFuzzySkipList(t *testing.T) {
	rand.Seed(1)
	tree := New()
	for i := 0; i < TestSize; i++ {
		v := rand.Intn(100000)
		item := slIntInt{v, v}
		if tree.Delete(item) == nil {
			tree.ReplaceOrInsert(item)
		}
	}
}

func TestFuzzyHangduSkipList(t *testing.T) {
	rand.Seed(1)
	tree := handu_skiplist.New(handu_skiplist.Int)
	for i := 0; i < TestSize; i++ {
		v := rand.Intn(100000)
		if tree.Remove(v) == nil {
			tree.Set(v, v)
		}
	}
}
