package main

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var testInputLines2 = strings.Split(puzzle1Test, "\n")

func TestCount(t *testing.T) {
	assert.Equal(t, 5, countHorizontal(testInputLines))
	assert.Equal(t, 3, countVertical(testInputLines))
	assert.Equal(t, 10, countDiagonals(testInputLines))

	assert.Equal(t, 18, part1(puzzle1Test))
}
