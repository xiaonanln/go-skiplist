package skiplist

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandom(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		t.Logf("random: %d", randomLevel(32))
	}
}

func BenchmarkRandomLevel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randomLevel(32)
	}
}

//func BenchmarkRandomLevel2(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		randomLevel2(32)
//	}
//}
//
//func BenchmarkRandomLevel3(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		randomLevel3(32)
//	}
//}
