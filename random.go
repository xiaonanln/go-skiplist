package skiplist

import (
	"math/rand"
)

// TODO: use faster random ?
func randomLevel(maxLevel int) int {
	l := 1
	for l < maxLevel && rand.Float32() < 0.5 {
		l++
	}
	return l
}
