package skiplist

import (
	"fmt"
	"unsafe"
)

const (
	validate = false // turn on for debug only, will run very slowly
	// DefaultMaxLevel is the default max level. 32 should be large enough for most cases.
	DefaultMaxLevel = 32
)

// Item is the interface that all SkipList items should implement
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

type node struct {
	Item
}

func get_forward(n *node, level int) *node {
	return *(**node)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + unsafe.Sizeof(node{}) + uintptr(level)*unsafe.Sizeof((*node)(nil))))
}

func set_forward(n *node, level int, fn *node) {
	*(**node)(unsafe.Pointer(uintptr(unsafe.Pointer(n)) + unsafe.Sizeof(node{}) + uintptr(level)*unsafe.Sizeof((*node)(nil)))) = fn
}

func releaseNode(n *node) (item Item) {
	item = n.Item
	return
}

// SkipList is the type of skip list
type SkipList struct {
	maxLevel int
	level    int
	len      int
	head     *node
}

// New creates a new SkipList with max level = 32, which should be enough for most cases.
func New() *SkipList {
	return NewMaxLevel(DefaultMaxLevel)
}

// NewMaxLevel creates a new SkipList with specified max level. Max level's value should be in range 1 ~ `MaxLevelLimit`
func NewMaxLevel(maxLevel int) *SkipList {
	if maxLevel > MaxLevelLimit {
		maxLevel = MaxLevelLimit
	} else if maxLevel < 1 {
		panic(fmt.Errorf("maxLevel must in range %d~%d", 1, MaxLevelLimit))
	}

	sl := &SkipList{
		maxLevel: maxLevel,
		level:    1,
		head:     allocNode(ninf, maxLevel),
	}
	sl.validate()
	return sl
}

// ReplaceOrInsert inserts item into the skiplist. If an existing
// element has the same order, it is removed from the skiplist and returned.
func (sl *SkipList) ReplaceOrInsert(item Item) Item {
	update := make([]*node, sl.level)
	p := sl.head
	var end *node
	for level := sl.level - 1; level >= 0; level-- {
		n := get_forward(p, level)
		for n != end {
			if n.Item.Less(item) {
				p = n
				n = get_forward(p, level)
			} else {
				end = n
				break
			}
		}
		update[level] = p
	}

	// insert after p
	n := get_forward(p, 0)
	if n == nil || item.Less(n.Item) {
		// insert before n
		newLevel := sl.randomLevel()
		if newLevel > sl.level {
			newLevel = sl.level + 1
			sl.level = newLevel
			//sl.head.forward = append(sl.head.forward, nil)
			update = append(update, sl.head)
		}

		newNode := allocNode(item, newLevel)
		for i := 0; i < newLevel; i++ {
			p := update[i]
			set_forward(newNode, i, get_forward(p, i))
			set_forward(p, i, newNode)
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
	update := make([]*node, sl.level)
	p := sl.head
	var end *node
	for level := sl.level - 1; level >= 0; level-- {
		n := get_forward(p, level)
		for n != end {
			if n.Item.Less(item) {
				p = n
				n = get_forward(p, level)
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
		//sl.head.forward = append(sl.head.forward, nil)
		update = append(update, sl.head)
	}
	newNode := allocNode(item, newLevel)
	for i := 0; i < newLevel; i++ {
		p := update[i]
		set_forward(newNode, i, get_forward(p, i))
		set_forward(p, i, newNode)
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

// Has returns true if the skiplist contains an element whose order is the same as that of key.
func (sl *SkipList) Has(item Item) bool {
	p := sl.head
	var end *node
	for level := sl.level - 1; level >= 0; level-- {
		n := get_forward(p, level)
		for n != end {
			if n.Item.Less(item) {
				p = n
				n = get_forward(p, level)
			} else {
				end = n
				break
			}
		}
	}

	n := get_forward(p, 0)
	return n != nil && !item.Less(n.Item)
}

// Delete deletes an item from the skiplist whose key equals key.
// The deleted item is return, otherwise nil is returned.
func (sl *SkipList) Delete(item Item) Item {
	update := make([]*node, sl.level)
	p := sl.head
	var end *node
	for level := sl.level - 1; level >= 0; level-- {
		n := get_forward(p, level)
		for n != end {
			if n.Item.Less(item) {
				p = n
				n = get_forward(p, level)
			} else {
				end = n
				break
			}
		}
		update[level] = p
	}

	// delete n if n matches
	n := get_forward(p, 0)
	if n == nil || item.Less(n.Item) {
		// item not found
		return nil
	}

	// n is the item to delete
	for i := 0; i < sl.level && get_forward(update[i], i) == n; i++ {
		set_forward(update[i], i, get_forward(n, i))
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
	n := get_forward(head, 0)
	if validate {
		if n == nil {
			panic(fmt.Errorf("bad skiplist"))
		}
	}

	for i := 0; i < sl.level && get_forward(head, i) == n; i++ {
		set_forward(head, i, get_forward(n, i))
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

	if sl.head.Item != ninf {
		panic(fmt.Errorf("head is not ninf"))
	}

	for i := 0; i < sl.level; i++ {
		pv := sl.head.Item
		levelSize := 0
		for p := get_forward(sl.head, i); p != nil; p = get_forward(p, i) {
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
