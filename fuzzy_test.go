package skiplist

import (
	"math/rand"
	"testing"

	"github.com/petar/GoLLRB/llrb"
)

const (
	TestSize = 100000
)

func TestFuzzyLLRB(t *testing.T) {
	rand.Seed(1)
	tree := llrb.New()
	for i := 0; i < TestSize; i++ {
		v := rand.Intn(100000)
		if tree.Delete(llrb.Int(v)) == nil {
			tree.ReplaceOrInsert(llrb.Int(v))
		}
	}
}

func TestFuzzySkipList(t *testing.T) {
	rand.Seed(1)
	tree := New()
	for i := 0; i < TestSize; i++ {
		v := rand.Intn(100000)
		if tree.Delete(Int(v)) == nil {
			tree.ReplaceOrInsert(Int(v))
		}
	}
}
