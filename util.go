package skiplist

type Int int

func (x Int) Less(other Item) bool {
	return x < other.(Int)
}

type String string

func (x String) Less(other Item) bool {
	return x < other.(String)
}
