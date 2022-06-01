package slicedemo

import "testing"

func BenchmarkNoPreAlloc10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoPreAlloc(10)
	}
}

func BenchmarkPreAlloc10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PreAlloc(10)
	}
}
