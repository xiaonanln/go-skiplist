package skiplist

import (
	"fmt"
)

const (
	validate = false // turn on for debug only, will run very slowly
	// DefaultMaxLevel is the default max level. 32 should be large enough for most cases.
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

func newNode(item Item, level int) (node *Node) {
	node = &Node{
		Item:    item,
		forward: make([]*Node, level),
	}
	return
}

func releaseNode(node *Node) (item Item) {
	item = node.Item
	return
}

type SkipList struct {
	maxLevel int
	level    int
	len      int
	head     *Node
}

// New creates a new SkipList with max level = 32, which should be enough for most cases.
func New() *SkipList {
	return NewMaxLevel(DefaultMaxLevel)
}

func NewMaxLevel(maxLevel int) *SkipList {
	sl := &SkipList{
		maxLevel: maxLevel,
		level:    1,
		head:     newNode(ninf, 1),
	}
	sl.validate()
	return sl
}

// ReplaceOrInsert inserts item into the skiplist. If an existing
// element has the same order, it is removed from the skiplist and returned.
func (sl *SkipList) ReplaceOrInsert(item Item) Item {
	update := make([]*Node, sl.level)
	p := sl.head
	var end *Node
	for level := sl.level - 1; level >= 0; level-- {
		n := p.forward[level]
		for n != end {
			if n.Item.Less(item) {
				p = n
				n = p.forward[level]
			} else {
				end = n
				break
			}
		}
		update[level] = p
	}

	// insert after p
	n := p.forward[0]
	if n == nil || item.Less(n.Item) {
		// insert before n
		newLevel := sl.randomLevel()
		if newLevel > sl.level {
			newLevel = sl.level + 1
			sl.level = newLevel
			sl.head.forward = append(sl.head.forward, nil)
			update = append(update, sl.head)
		}

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
	update := make([]*Node, sl.level)
	p := sl.head
	var end *Node
	for level := sl.level - 1; level >= 0; level-- {
		n := p.forward[level]
		for n != end {
			if n.Item.Less(item) {
				p = n
				n = p.forward[level]
			} else {
				end = n
				break
			}
		}
		update[level] = p
	}

	// always insert after p without check if n equals to item
	// insert before n
	newLevel := sl.randomLevel()
	if newLevel > sl.level {
		newLevel = sl.level + 1
		sl.level = newLevel
		sl.head.forward = append(sl.head.forward, nil)
		update = append(update, sl.head)
	}
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

// MaxLevel returns the max level of the skiplist
func (sl *SkipList) MaxLevel() int {
	return sl.maxLevel
}

// Level returns the current level of the skiplist
func (sl *SkipList) Level() int {
	return sl.level
}

// Has returns true if the skiplist contains an element whose order is the same as that of key.
func (sl *SkipList) Has(item Item) bool {
	p := sl.head
	var end *Node
	for level := sl.level - 1; level >= 0; level-- {
		n := p.forward[level]
		for n != end {
			if n.Item.Less(item) {
				p = n
				n = p.forward[level]
			} else {
				end = n
				break
			}
		}
	}

	n := p.forward[0]
	return n != nil && !item.Less(n.Item)
}

// Delete deletes an item from the skiplist whose key equals key.
// The deleted item is return, otherwise nil is returned.
func (sl *SkipList) Delete(item Item) Item {
	update := make([]*Node, sl.level)
	p := sl.head
	var end *Node
	for level := sl.level - 1; level >= 0; level-- {
		n := p.forward[level]
		for n != end {
			if n.Item.Less(item) {
				p = n
				n = p.forward[level]
			} else {
				end = n
				break
			}
		}
		update[level] = p
	}

	// delete n if n matches
	n := p.forward[0]
	if n == nil || item.Less(n.Item) {
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
	return releaseNode(n)
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
		if n == nil {
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
	return releaseNode(n)
}

func (sl *SkipList) randomLevel() int {
	return randomLevel(sl.maxLevel)
}

func (sl *SkipList) validate() {
	if !validate {
		return
	}

	if len(sl.head.forward) != sl.level {
		panic(fmt.Errorf("wrong head.forward level: %d, should be %d", len(sl.head.forward), sl.level))
	}

	if sl.head.Item != ninf {
		panic(fmt.Errorf("head is not ninf"))
	}

	for i := 0; i < sl.level; i++ {
		pv := sl.head.Item
		levelSize := 0
		for p := sl.head.forward[i]; p != nil; p = p.forward[i] {
			if less(p.Item, pv) {
				panic(fmt.Errorf("wrong order: %#v > %#v", pv, p.Item))
			}
			pv = p.Item
			levelSize++
		}
		if i == 0 {
			if levelSize != sl.len {
				panic(fmt.Errorf("bad len: %d, should be %d", sl.len, levelSize))
			}
		} else {
			if levelSize > sl.len {
				panic(fmt.Errorf("bad len: %d, should be larger than or equal to %d", sl.len, levelSize))
			}
		}
	}
}
