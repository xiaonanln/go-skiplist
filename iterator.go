package skiplist

type ItemIterator func(i Item) bool

func (sl *SkipList) Ascend(iterator ItemIterator) {
	var end *Node
	for n := get_forward(sl.head, 0); n != end; n = get_forward(n, 0) {
		if !iterator(n.Item) {
			return
		}
	}
}

func (sl *SkipList) AscendRange(greaterOrEqual, lessThan Item, iterator ItemIterator) {
	p := sl.head
	var end *Node
	for level := sl.level - 1; level >= 0; level-- {
		n := get_forward(p, level)
		for n != end {
			if less(n.Item, greaterOrEqual) {
				p = n
				n = get_forward(p, level)
			} else {
				end = n
				break
			}
		}
	}

	// n >= pivot
	for n := get_forward(p, 0); n != nil && less(n.Item, lessThan); n = get_forward(n, 0) {
		if !iterator(n.Item) {
			return
		}
	}
}

// AscendGreaterOrEqual will call iterator once for each element greater or equal to
// pivot in ascending order. It will stop whenever the iterator returns false.
func (sl *SkipList) AscendGreaterOrEqual(pivot Item, iterator ItemIterator) {
	p := sl.head
	var end *Node
	for level := sl.level - 1; level >= 0; level-- {
		n := get_forward(p, level)
		for n != end {
			if less(n.Item, pivot) {
				p = n
				n = get_forward(p, level)
			} else {
				end = n
				break
			}
		}
	}

	// n >= pivot
	for n := get_forward(p, 0); n != nil; n = get_forward(n, 0) {
		if !iterator(n.Item) {
			return
		}
	}
}

// AscendLessThan will call iterator once for each element less than the
// pivot in ascending order. It will stop whenever the iterator returns false.
func (sl *SkipList) AscendLessThan(pivot Item, iterator ItemIterator) {
	for n := get_forward(sl.head, 0); n != nil && less(n.Item, pivot); n = get_forward(n, 0) {
		if !iterator(n.Item) {
			return
		}
	}
}
