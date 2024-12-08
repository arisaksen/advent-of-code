package main

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCheckCombinationsPart2(t *testing.T) {
	assert.Equal(t, 12, checkCombinations([]string{"1", "1", "1"}, "|+"))
	assert.Equal(t, 41, checkCombinations([]string{"2", "2", "1"}, "+|"))
}

func TestCheckOperatorPart2(t *testing.T) {
	sings := []string{"+", "*", "|"}

	tests := []struct {
		name           string
		line           string
		expectedAnswer int
	}{
		{
			name:           "test 0",
			line:           "190: 10 19",
			expectedAnswer: 190,
		},
		{
			name:           "test 1",
			line:           "3267: 81 40 27",
			expectedAnswer: 3267,
		},
		{
			name:           "test 2",
			line:           "161011: 16 10 13",
			expectedAnswer: 0,
		},
		{
			name:           "test 3",
			line:           "21037: 9 7 18 13",
			expectedAnswer: 0,
		},
		{
			name:           "test 4",
			line:           "292: 11 6 16 20",
			expectedAnswer: 292,
		},
		{
			name:           "test 5",
			line:           "7290: 6 8 6 15",
			expectedAnswer: 7290,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			operatorsLen := len(strings.Fields(test.line)) - 2
			operatorCombinations := new([]string)
			generateCombinations(sings, operatorsLen, "", operatorCombinations)

			answer := checkOperator(test.line, *operatorCombinations)
			assert.Equal(t, test.expectedAnswer, answer)
		})
	}

}

func TestPart2(t *testing.T) {
	actual := part2(puzzle1Test)

	assert.Equal(t, 11387, actual)
}

func BenchmarkGenerateCombinations(b *testing.B) {
	operatorCombinations := new([]string)
	for i := 0; i < b.N; i++ {
		generateCombinations(operators2, 5, "", operatorCombinations)
	}
	b.ReportAllocs()
}

func BenchmarkCheckOperator(b *testing.B) {
	sings := []string{"+", "*", "|"}
	line := "7290: 6 8 6 15"

	operatorsLen := len(strings.Fields(line)) - 2
	operatorCombinations := new([]string)
	generateCombinations(sings, operatorsLen, "", operatorCombinations)

	for i := 0; i < b.N; i++ {
		_ = checkOperator(line, *operatorCombinations)
	}
	b.ReportAllocs()
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = part2(puzzle1)
	}
	b.ReportAllocs()
}
