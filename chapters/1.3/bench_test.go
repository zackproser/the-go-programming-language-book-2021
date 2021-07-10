package main

import "testing"

var testStrings = []string{"echo", "wikka", "wakka", "wooo"}

// BenchmarkBruteForceJoin allows for simple benchmarking of the bruteForceJoin function
func BenchmarkBruteForceJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BruteForceJoin(testStrings)
	}
}

func BenchmarkFasterJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FasterJoin(testStrings)
	}
}
