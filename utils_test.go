package main

import "testing"

func BenchmarkCalculateConcurrently(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateConcurrently(i)
	}
}

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate(i)
	}
}
