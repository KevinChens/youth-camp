package stringjoint

import "testing"

var (
	n   = 10000
	str = "hello"
)

func BenchmarkPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Plus(n, str)
	}
}

func BenchmarkStrBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StrBuilder(n, str)
	}
}

func BenchmarkByteBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ByteBuffer(n, str)
	}
}

func BenchmarkPreStrBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PreStrBuilder(n, str)
	}
}

func BenchmarkPreByteBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PreByteBuffer(n, str)
	}
}
