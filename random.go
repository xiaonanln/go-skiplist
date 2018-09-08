package skiplist

import "time"

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
