package main

import "testing"

func BenchmarkInt(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Int()
	}
	b.StopTimer()
}

func BenchmarkBigInt(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BigInt()
	}
	b.StopTimer()
}
