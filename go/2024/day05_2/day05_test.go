package main

import "testing"

func BenchmarkPart1(b *testing.B) {
	for range b.N {
		part1(puzzle1)
	}
	b.ReportAllocs()
}

func BenchmarkPart2(b *testing.B) {
	for range b.N {
		part2(puzzle1)
	}
	b.ReportAllocs()
}
