package skiplist

import (
	"math/rand"
	"testing"

	"github.com/petar/GoLLRB/llrb"
)

const (
	TestSize = 1000000
)

func TestLLRB(t *testing.T) {
	rand.Seed(1)
	tree := llrb.New()
	vals := []int{}
	for i := 0; i < TestSize; i++ {
		v := rand.Int()
		tree.ReplaceOrInsert(llrb.Int(v))
		vals = append(vals, v)
	}
	//for tree.Len() > 0 {
	//	tree.DeleteMin()
	//}
	for _, v := range vals {
		tree.Has(llrb.Int(v))
		tree.Delete(llrb.Int(v))
	}
}

func TestSkipList(t *testing.T) {
	rand.Seed(1)
	tree := New()
	vals := []int{}
	for i := 0; i < TestSize; i++ {
		v := rand.Int()
		tree.ReplaceOrInsert(Int(v))
		vals = append(vals, v)
	}
	//for tree.Len() > 0 {
	//	tree.DeleteMin()
	//}
	for _, v := range vals {
		tree.Has(Int(v))
		tree.Delete(Int(v))
	}
}
