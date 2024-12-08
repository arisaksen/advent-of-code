package main

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

	actual := part1(input)

	assert.Equal(t, 1, actual)
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = part1(puzzle1Test)
	}
	b.ReportAllocs()
}
