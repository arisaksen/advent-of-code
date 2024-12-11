package main

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed puzzle1_test.txt
var puzzle1Test string

func TestPart1(t *testing.T) {
	actual := part1(puzzle1Test)

	assert.Equal(t, 1, actual)
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = part1(puzzle1Test)
	}
	b.ReportAllocs()
}
	