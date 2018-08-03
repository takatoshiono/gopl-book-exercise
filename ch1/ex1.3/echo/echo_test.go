package echo

import "testing"

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo1()
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo3()
	}
}
