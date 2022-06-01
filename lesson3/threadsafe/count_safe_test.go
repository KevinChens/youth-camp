package threadsafe

import "testing"

func BenchmarkAtomicAddOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := &atomicCounter{i: 0}
		AtomicAddOne(a)
	}
}

func BenchmarkMutexAddOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := &mutexCounter{i: 0}
		MutexAddOne(m)
	}
}
