package main

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart2(t *testing.T) {
	assert.Equal(t, 9, part2(puzzle1Test))
}
