package skiplist

import (
	"testing"
	"unsafe"
)

func TestAlloc(t *testing.T) {
	println("Node size", unsafe.Sizeof(Node{}))
	println("l1node size", unsafe.Sizeof(l1node{}))
}

func BenchmarkAllocNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		allocNode(nil, 1)
		allocNode(nil, 1)
		allocNode(nil, 1)
		allocNode(nil, 1)
		allocNode(nil, 2)
		allocNode(nil, 2)
		allocNode(nil, 3)
		allocNode(nil, 1)
		allocNode(nil, 1)
		allocNode(nil, 1)
		allocNode(nil, 1)
		allocNode(nil, 2)
		allocNode(nil, 2)
		allocNode(nil, 3)
		allocNode(nil, 4)
	}
}

type oldnode struct {
	item    Item
	forward []*Node
}

func oldAllocNode(item Item, level int) *oldnode {
	return &oldnode{
		item:    item,
		forward: make([]*Node, level),
	}
}

func BenchmarkAllocNodeOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		oldAllocNode(nil, 1)
		oldAllocNode(nil, 1)
		oldAllocNode(nil, 1)
		oldAllocNode(nil, 1)
		oldAllocNode(nil, 2)
		oldAllocNode(nil, 2)
		oldAllocNode(nil, 3)
		oldAllocNode(nil, 1)
		oldAllocNode(nil, 1)
		oldAllocNode(nil, 1)
		oldAllocNode(nil, 1)
		oldAllocNode(nil, 2)
		oldAllocNode(nil, 2)
		oldAllocNode(nil, 3)
		oldAllocNode(nil, 4)
	}
}
