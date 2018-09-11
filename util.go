package skiplist

// Int is a type for wrapping int to Item
type Int int

// Less compares int values
func (x Int) Less(other Item) bool {
	return x < other.(Int)
}

// String is a type for wrapping string to Item
type String string

func (x String) Less(other Item) bool {
	return x < other.(String)
}
