package main

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

//go:embed puzzle1_test.txt
var puzzle1Test string
var testInputLines = strings.Split(puzzle1Test, "\n")

func TestCountWordOccurrences(t *testing.T) {
	tests := []struct {
		name          string
		text          string
		word          string
		expectedCount int
	}{
		{
			name:          "1 xmas",
			text:          "XMASA",
			word:          xmas,
			expectedCount: 1,
		},
		{
			name:          "0 xmas",
			text:          "ABCDEFG",
			word:          xmas,
			expectedCount: 0,
		},
		{
			name:          "1 xmas in larger string",
			text:          "MSAXMASMSMXMAS",
			word:          xmas,
			expectedCount: 2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedCount, countWordOccurrences(test.text, test.word))
		})
	}
}

func TestCount(t *testing.T) {
	assert.Equal(t, 5, countHorizontal(testInputLines, xmas, samx))
	assert.Equal(t, 3, countVertical(testInputLines, xmas, samx))
	assert.Equal(t, 10, countDiagonals(testInputLines, xmas))

	assert.Equal(t, 18, part1(puzzle1Test))
}
