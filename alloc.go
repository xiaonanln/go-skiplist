package skiplist

import "unsafe"

const (
	// MaxLevelLimit is the max level for SkipList
	MaxLevelLimit = 40
)

type l1node struct {
	item    Item
	forward [1]*node
}

type l2node struct {
	item    Item
	forward [2]*node
}

type l3node struct {
	item    Item
	forward [3]*node
}

type l4node struct {
	item    Item
	forward [4]*node
}

type l5node struct {
	item    Item
	forward [5]*node
}

type l6node struct {
	item    Item
	forward [6]*node
}

type l7node struct {
	item    Item
	forward [7]*node
}

type l8node struct {
	item    Item
	forward [8]*node
}

type l9node struct {
	item    Item
	forward [9]*node
}

type l10node struct {
	item    Item
	forward [10]*node
}

type l11node struct {
	item    Item
	forward [11]*node
}

type l12node struct {
	item    Item
	forward [12]*node
}

type l13node struct {
	item    Item
	forward [13]*node
}

type l14node struct {
	item    Item
	forward [14]*node
}

type l15node struct {
	item    Item
	forward [15]*node
}

type l16node struct {
	item    Item
	forward [16]*node
}

type l17node struct {
	item    Item
	forward [17]*node
}

type l18node struct {
	item    Item
	forward [18]*node
}

type l19node struct {
	item    Item
	forward [19]*node
}

type l20node struct {
	item    Item
	forward [20]*node
}

type l21node struct {
	item    Item
	forward [21]*node
}

type l22node struct {
	item    Item
	forward [22]*node
}

type l23node struct {
	item    Item
	forward [23]*node
}

type l24node struct {
	item    Item
	forward [24]*node
}

type l25node struct {
	item    Item
	forward [25]*node
}

type l26node struct {
	item    Item
	forward [26]*node
}

type l27node struct {
	item    Item
	forward [27]*node
}

type l28node struct {
	item    Item
	forward [28]*node
}

type l29node struct {
	item    Item
	forward [29]*node
}

type l30node struct {
	item    Item
	forward [30]*node
}

type l31node struct {
	item    Item
	forward [31]*node
}

type l32node struct {
	item    Item
	forward [32]*node
}
type l33node struct {
	item    Item
	forward [33]*node
}

type l34node struct {
	item    Item
	forward [34]*node
}

type l35node struct {
	item    Item
	forward [35]*node
}

type l36node struct {
	item    Item
	forward [36]*node
}

type l37node struct {
	item    Item
	forward [37]*node
}

type l38node struct {
	item    Item
	forward [38]*node
}

type l39node struct {
	item    Item
	forward [39]*node
}

type l40node struct {
	item    Item
	forward [40]*node
}

var (
	allocFuncs = []func() unsafe.Pointer{
		func() unsafe.Pointer { return unsafe.Pointer(&l1node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l2node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l3node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l4node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l5node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l6node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l7node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l8node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l9node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l10node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l11node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l12node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l13node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l14node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l15node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l16node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l17node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l18node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l19node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l20node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l21node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l22node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l23node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l24node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l25node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l26node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l27node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l28node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l29node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l30node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l31node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l32node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l33node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l34node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l35node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l36node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l37node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l38node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l39node{}) },
		func() unsafe.Pointer { return unsafe.Pointer(&l40node{}) },
	}
)

func allocNode(item Item, level int) *node {
	node := (*node)(allocFuncs[level-1]())
	node.Item = item
	return node
}
