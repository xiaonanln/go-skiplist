package skiplist

import (
	"fmt"
)

const (
	validate        = false // turn on for debug only
	DefaultMaxLevel = 32
)

type Item interface {
	Less(other Item) bool
}

func less(x, y Item) bool {
	if y == pinf {
		return true
	}
	if y == ninf {
		return false
	}
	return x.Less(y)
}

type Node struct {
	Item
	forward []*Node
}

func newNode(item Item, level int) *Node {
	return &Node{
		Item:    item,
		forward: make([]*Node, level),
	}
}

type SkipList struct {
	maxLevel int
	len      int
	head     *Node
	tail     *Node
}

// New creates a new SkipList with max maxLevel = 32, which should be enough for most cases.
func New() *SkipList {
	return NewMaxLevel(DefaultMaxLevel)
}

func NewMaxLevel(maxLevel int) *SkipList {
	sl := &SkipList{
		maxLevel: maxLevel,
		head:     newNode(ninf, maxLevel),
		tail:     newNode(pinf, maxLevel),
	}
	for i := 0; i < maxLevel; i++ {
		sl.head.forward[i] = sl.tail
		//sl.tail.forward[i] = nil
	}
	sl.validate()
	return sl
}

// ReplaceOrInsert inserts item into the skiplist. If an existing
// element has the same order, it is removed from the skiplist and returned.
func (sl *SkipList) ReplaceOrInsert(item Item) Item {
	update := make([]*Node, sl.maxLevel)
	p := sl.head
	for level := sl.maxLevel - 1; level >= 0; level-- {
		for less(p.forward[level].Item, item) {
			p = p.forward[level]
		}
		update[level] = p
	}

	// insert after p
	n := p.forward[0]
	if less(item, n.Item) {
		// insert before n
		newLevel := sl.randomLevel()
		newNode := newNode(item, newLevel)
		for i := 0; i < newLevel; i++ {
			p := update[i]
			newNode.forward[i] = p.forward[i]
			p.forward[i] = newNode
		}
		sl.len++
		sl.validate()
		return nil
	} else {
		// replace n
		replacedItem := n.Item
		n.Item = item
		sl.validate()
		return replacedItem
	}
}

// InsertNoReplace inserts item into the skiplist. If an existing
// element has the same order, both elements remain in the skiplist.
func (sl *SkipList) InsertNoReplace(item Item) {
	update := make([]*Node, sl.maxLevel)
	p := sl.head
	for level := sl.maxLevel - 1; level >= 0; level-- {
		for less(p.forward[level].Item, item) {
			p = p.forward[level]
		}
		update[level] = p
	}

	// always insert after p without check if n equals to item
	// insert before n
	newLevel := sl.randomLevel()
	newNode := newNode(item, newLevel)
	for i := 0; i < newLevel; i++ {
		p := update[i]
		newNode.forward[i] = p.forward[i]
		p.forward[i] = newNode
	}
	sl.len++
	sl.validate()
}

// Len returns the number of nodes in the skiplist.
func (sl *SkipList) Len() int {
	return sl.len
}

// MaxLevel returns the max maxLevel of the skiplist
func (sl *SkipList) MaxLevel() int {
	return sl.maxLevel
}

// Has returns true if the skiplist contains an element whose order is the same as that of key.
func (sl *SkipList) Has(item Item) bool {
	p := sl.head
	for level := sl.maxLevel - 1; level >= 0; level-- {
		for less(p.forward[level].Item, item) {
			p = p.forward[level]
		}
	}
	n := p.forward[0]
	return !less(item, n.Item)
}

// Delete deletes an item from the skiplist whose key equals key.
// The deleted item is return, otherwise nil is returned.
func (sl *SkipList) Delete(item Item) Item {
	update := make([]*Node, sl.maxLevel)
	p := sl.head
	for level := sl.maxLevel - 1; level >= 0; level-- {
		for less(p.forward[level].Item, item) {
			p = p.forward[level]
		}
		update[level] = p
	}

	// delete n if n matches
	n := p.forward[0]
	if less(item, n.Item) {
		// item not found
		return nil
	}

	// n is the item to delete
	level := len(n.forward)
	for i := 0; i < level; i++ {
		if validate {
			if update[i].forward[i] != n {
				panic(fmt.Errorf("wrong forward pointer"))
			}
		}

		update[i].forward[i] = n.forward[i]
	}
	sl.len--

	sl.validate()
	return n.Item
}

// DeleteMin deletes the minimum element in the skiplist and returns the
// deleted item or nil otherwise.
func (sl *SkipList) DeleteMin() Item {
	if sl.len == 0 {
		return nil
	}

	// find the min item
	head := sl.head
	n := head.forward[0]
	if validate {
		if n == sl.tail {
			panic(fmt.Errorf("bad skiplist"))
		}
	}

	level := len(n.forward)
	for i := 0; i < level; i++ {
		if validate && head.forward[i] != n {
			panic(fmt.Errorf("bad skiplist"))
		}

		head.forward[i] = n.forward[i]
	}

	sl.len--
	sl.validate()
	return n.Item
}

func (sl *SkipList) randomLevel() int {
	return randomLevel(sl.maxLevel)
}

func (sl *SkipList) validate() {
	if !validate {
		return
	}

	if len(sl.head.forward) != sl.maxLevel {
		panic(fmt.Errorf("wrong head.forward maxLevel: %d, should be %d", len(sl.head.forward), sl.maxLevel))
	}

	if len(sl.tail.forward) != sl.maxLevel {
		panic(fmt.Errorf("wrong tail.forward maxLevel"))
	}

	if sl.head.Item != ninf {
		panic(fmt.Errorf("head is not ninf"))
	}

	if sl.tail.Item != pinf {
		panic(fmt.Errorf("tail is not pinf"))
	}

	for i := 0; i < sl.maxLevel; i++ {
		pv := sl.head.Item
		levelSize := 0
		for p := sl.head.forward[i]; p != nil; p = p.forward[i] {
			if less(p.Item, pv) {
				panic(fmt.Errorf("wrong order: %#v > %#v", pv, p.Item))
			}
			pv = p.Item
			levelSize++
		}
		if pv != pinf {
			panic(fmt.Errorf("maxLevel %d not ends with tail", i))
		}
		if i == 0 {
			if levelSize != sl.len+1 {
				panic(fmt.Errorf("bad len: %d, should be %d", sl.len, levelSize-1))
			}
		} else {
			if levelSize > sl.len+1 {
				panic(fmt.Errorf("bad len: %d, should be larger than or equal to %d", sl.len, levelSize-1))
			}
		}
	}
}
