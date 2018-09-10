package skiplist

import "unsafe"

const (
	// MaxLevelLimit is the max level for SkipList
	MaxLevelLimit = 40
)

type l1node struct {
	item    Item
	forward [1]*Node
}

type l2node struct {
	item    Item
	forward [2]*Node
}

type l3node struct {
	item    Item
	forward [3]*Node
}

type l4node struct {
	item    Item
	forward [4]*Node
}

type l5node struct {
	item    Item
	forward [5]*Node
}

type l6node struct {
	item    Item
	forward [6]*Node
}

type l7node struct {
	item    Item
	forward [7]*Node
}

type l8node struct {
	item    Item
	forward [8]*Node
}

type l9node struct {
	item    Item
	forward [9]*Node
}

type l10node struct {
	item    Item
	forward [10]*Node
}

type l11node struct {
	item    Item
	forward [11]*Node
}

type l12node struct {
	item    Item
	forward [12]*Node
}

type l13node struct {
	item    Item
	forward [13]*Node
}

type l14node struct {
	item    Item
	forward [14]*Node
}

type l15node struct {
	item    Item
	forward [15]*Node
}

type l16node struct {
	item    Item
	forward [16]*Node
}

type l17node struct {
	item    Item
	forward [17]*Node
}

type l18node struct {
	item    Item
	forward [18]*Node
}

type l19node struct {
	item    Item
	forward [19]*Node
}

type l20node struct {
	item    Item
	forward [20]*Node
}

type l21node struct {
	item    Item
	forward [21]*Node
}

type l22node struct {
	item    Item
	forward [22]*Node
}

type l23node struct {
	item    Item
	forward [23]*Node
}

type l24node struct {
	item    Item
	forward [24]*Node
}

type l25node struct {
	item    Item
	forward [25]*Node
}

type l26node struct {
	item    Item
	forward [26]*Node
}

type l27node struct {
	item    Item
	forward [27]*Node
}

type l28node struct {
	item    Item
	forward [28]*Node
}

type l29node struct {
	item    Item
	forward [29]*Node
}

type l30node struct {
	item    Item
	forward [30]*Node
}

type l31node struct {
	item    Item
	forward [31]*Node
}

type l32node struct {
	item    Item
	forward [32]*Node
}
type l33node struct {
	item    Item
	forward [33]*Node
}

type l34node struct {
	item    Item
	forward [34]*Node
}

type l35node struct {
	item    Item
	forward [35]*Node
}

type l36node struct {
	item    Item
	forward [36]*Node
}

type l37node struct {
	item    Item
	forward [37]*Node
}

type l38node struct {
	item    Item
	forward [38]*Node
}

type l39node struct {
	item    Item
	forward [39]*Node
}

type l40node struct {
	item    Item
	forward [40]*Node
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

func allocNode(item Item, level int) *Node {
	node := (*Node)(allocFuncs[level-1]())
	node.Item = item
	return node
}
