package main

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed puzzle2_test.txt
var puzzle2Test string

func TestPart2(t *testing.T) {
	actual := part2(puzzle1Test)

	assert.Equal(t, 31, actual)
}
