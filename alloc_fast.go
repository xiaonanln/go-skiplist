package skiplist

import "C"

import (
	"reflect"
	"unsafe"
)

func newNodeFast(item Item, level int) (node *Node) {
	node = (*Node)(C.malloc(int(unsafe.Sizeof(Node{}) + uintptr(level)*unsafe.Sizeof((*Node)(nil)))))
	h := (*reflect.SliceHeader)(unsafe.Pointer(&node.forward))
	h.Data = uintptr(unsafe.Pointer(node)) + unsafe.Sizeof(Node{})
	h.Cap = level
	h.Len = level
	return
}
