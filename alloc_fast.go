package skiplist

import "unsafe"

type l1n struct {
	Item
	forward []*Node
	array   [1]Node
}

type l2n struct {
	Item
	forward []*Node
	array   [2]Node
}

type l3n struct {
	Item
	forward []*Node
	array   [3]Node
}

type l4n struct {
	Item
	forward []*Node
	array   [4]Node
}

type l5n struct {
	Item
	forward []*Node
	array   [5]Node
}

type l6n struct {
	Item
	forward []*Node
	array   [6]Node
}

type l7n struct {
	Item
	forward []*Node
	array   [7]Node
}

type l8n struct {
	Item
	forward []*Node
	array   [8]Node
}

func newNodeFast(item Item, level int) (node *Node) {
	switch level {
	case 1:
		node = (*Node)(unsafe.Pointer(&l1n{Item: item}))
	case 2:
		node = (*Node)(unsafe.Pointer(&l2n{Item: item}))
	case 3:
		node = (*Node)(unsafe.Pointer(&l3n{Item: item}))
	case 4:
		node = (*Node)(unsafe.Pointer(&l4n{Item: item}))
	//case 5:
	//	node = (*Node)(unsafe.Pointer(&l5n{Item: item}))
	//case 6:
	//	node = (*Node)(unsafe.Pointer(&l6n{Item: item}))
	//case 7:
	//	node = (*Node)(unsafe.Pointer(&l7n{Item: item}))
	//case 8:
	//	node = (*Node)(unsafe.Pointer(&l8n{Item: item}))
	default:
		node = newNode(item, level)
	}
	return
}
