package skiplist

type ItemIterator func(i Item) bool

func (sl *SkipList) Ascend(iterator ItemIterator) {
	tail := sl.tail
	for n := sl.head.forward[0]; n != tail; n = n.forward[0] {
		if !iterator(n.Item) {
			return
		}
	}
}

func (sl *SkipList) AscendRange(greaterOrEqual, lessThan Item, iterator ItemIterator) {
	p := sl.head
	for level := sl.maxLevel - 1; level >= 0; level-- {
		for less(p.forward[level].Item, greaterOrEqual) {
			p = p.forward[level]
		}
	}
	// n >= pivot
	for n := p.forward[0]; less(n.Item, lessThan); n = n.forward[0] {
		if !iterator(n.Item) {
			return
		}
	}
}

// AscendGreaterOrEqual will call iterator once for each element greater or equal to
// pivot in ascending order. It will stop whenever the iterator returns false.
func (sl *SkipList) AscendGreaterOrEqual(pivot Item, iterator ItemIterator) {
	p := sl.head
	for level := sl.maxLevel - 1; level >= 0; level-- {
		for less(p.forward[level].Item, pivot) {
			p = p.forward[level]
		}
	}
	tail := sl.tail
	// n >= pivot
	for n := p.forward[0]; n != tail; n = n.forward[0] {
		if !iterator(n.Item) {
			return
		}
	}
}

// AscendLessThan will call iterator once for each element less than the
// pivot in ascending order. It will stop whenever the iterator returns false.
func (sl *SkipList) AscendLessThan(pivot Item, iterator ItemIterator) {
	for n := sl.head.forward[0]; less(n.Item, pivot); n = n.forward[0] {
		if !iterator(n.Item) {
			return
		}
	}
}
