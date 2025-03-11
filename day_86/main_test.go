package main

import "testing"

func benchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Soma(10, 20)
	}
}
