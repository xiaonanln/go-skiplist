package skiplist

import (
	"time"
)

var (
	seed uint32
)

func init() {
	seed = uint32(time.Now().UnixNano())
}

const (
	M = 2147483647
	A = 16807
)

func random() uint32 {
	product := seed * A

	// Compute (product % M) using the fact that ((x << 31) % M) == x.
	seed = ((product >> 31) + (product & M))
	// The first reduction may overflow by 1 bit, so we may need to
	// repeat.  mod == M is not possible; using > allows the faster
	// sign-bit-based test.
	if seed > M {
		seed -= M
	}
	return seed
}

//func randomLevel(maxLevel int) int {
//	l := 31 - int(math.Log2(float64(random()%2147483647)))
//	if l == 0 {
//		l = 1
//	} else if l > maxLevel {
//		l = maxLevel
//	}
//	return l
//}
//
//func randomLevel2(maxLevel int) int {
//	l := 1
//	for l < maxLevel && rand.Float32() < 0.5 {
//		l++
//	}
//	return l
//}

func randomLevel(maxLevel int) int {
	l := 1
	for l < maxLevel && random()%100 < 50 {
		l++
	}
	return l
}
