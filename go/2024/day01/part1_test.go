package main

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed puzzle1_test.txt
var puzzle1Test string

func TestRemoveIndex(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	removeIndex(&input, 2)

	assert.Equal(t, 4, len(input))
	t.Log(input)
}

func TestFindAndRemoveSmallest(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	smallest := findAndRemoveSmallestNumber(&input)

	assert.Equal(t, 4, len(input))
	assert.Equal(t, 1, smallest)
	t.Log(input)
}

func TestPart1(t *testing.T) {
	actual := part1(puzzle1Test)

	assert.Equal(t, 11, actual)
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = part1(puzzle1Test)
	}
	b.ReportAllocs()
}
