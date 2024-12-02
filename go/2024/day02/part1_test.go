package main

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed puzzle1_test.txt
var puzzle1Test string

// The levels are either all increasing or all decreasing.
// Any two adjacent levels differ by at least one and at most three.

func TestAdjacent(t *testing.T) {
	assert.Equal(t, true, checkAdjacent(7, 6, 3))
	assert.Equal(t, true, checkAdjacent(1, 2, 3))
	assert.Equal(t, false, checkAdjacent(2, 7, 3))
	assert.Equal(t, false, checkAdjacent(2, -7, 3))
	assert.Equal(t, false, checkAdjacent(-7, 2, 3))
}

func TestSafeReport(t *testing.T) {
	tests := []struct {
		name         string
		report       []int
		expectedSafe bool
	}{
		{
			name:         "Report 0",
			report:       []int{7, 6, 4, 2, 1},
			expectedSafe: true,
		},
		{
			name:         "Report 1",
			report:       []int{1, 2, 7, 8, 9},
			expectedSafe: false,
		},
		{
			name:         "Report 2",
			report:       []int{9, 7, 6, 2, 1},
			expectedSafe: false,
		},
		{
			name:         "Report 3",
			report:       []int{1, 3, 2, 4, 5},
			expectedSafe: false,
		},
		{
			name:         "Report 4",
			report:       []int{8, 6, 4, 4, 1},
			expectedSafe: false,
		},
		{
			name:         "Report 5",
			report:       []int{1, 3, 6, 7, 9},
			expectedSafe: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedSafe, checkIfSafe(test.report, 0))
		})
	}

}

func TestPart1(t *testing.T) {
	actual := part1(puzzle1Test)

	assert.Equal(t, 2, actual)
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = part1(puzzle1Test)
	}
	b.ReportAllocs()
}
