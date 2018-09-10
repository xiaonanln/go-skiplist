package skiplist

import (
	"testing"
	"unsafe"
)

func TestAlloc(t *testing.T) {
	println("Node size", unsafe.Sizeof(Node{}))
	println("l1node size", unsafe.Sizeof(l1node{}))
}
