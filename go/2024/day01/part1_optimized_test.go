package main

import (
	_ "embed"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPart1EqualPart1Optimized(t *testing.T) {
	part1Original := part1(puzzle1)
	part1Opt := part1Optimized(puzzle1)

	assert.Equal(t, part1Original, part1Opt)
}

func TestCompared(t *testing.T) {
	start1 := time.Now()
	_ = part1(puzzle1)
	fmt.Println("Part 1:", time.Since(start1))

	start2 := time.Now()
	_ = part1Optimized(puzzle1)
	fmt.Println("Optimised:", time.Since(start2))

	start3 := time.Now()
	_ = part1(puzzle1)
	fmt.Println("Part 1 again:", time.Since(start3))
}

func BenchmarkTestOptimized(b *testing.B) {

	b.Run("part1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = part1(puzzle1Test)
		}
		b.ReportAllocs()
	})

	b.Run("part1Optimised", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = part1Optimized(puzzle1Test)
		}
		b.ReportAllocs()
	})

}
