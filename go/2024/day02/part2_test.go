package main

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSafeReportWithDampener(t *testing.T) {
	tests := []struct {
		name            string
		report          []int
		problemDampener int
		expectedSafe    bool
	}{
		{
			name:            "Report 0",
			report:          []int{7, 6, 4, 2, 1},
			problemDampener: 1,
			expectedSafe:    true,
		},
		{
			name:            "Report 1",
			report:          []int{1, 2, 7, 8, 9},
			problemDampener: 1,
			expectedSafe:    false,
		},
		{
			name:            "Report 2",
			report:          []int{9, 7, 6, 2, 1},
			problemDampener: 1,
			expectedSafe:    false,
		},
		{
			name:            "Report 3",
			report:          []int{1, 3, 2, 4, 5},
			problemDampener: 1,
			expectedSafe:    true,
		},
		{
			name:            "Report 4",
			report:          []int{8, 6, 4, 4, 1},
			problemDampener: 1,
			expectedSafe:    true,
		},
		{
			name:            "Report 5",
			report:          []int{1, 3, 6, 7, 9},
			problemDampener: 1,
			expectedSafe:    true,
		},
		{
			name:            "Report 5",
			report:          []int{58, 59, 62, 63, 64, 63},
			problemDampener: 1,
			expectedSafe:    true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedSafe, checkIfSafe(test.report, test.problemDampener))
		})
	}

}
