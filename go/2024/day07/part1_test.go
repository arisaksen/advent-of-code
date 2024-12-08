package main

import (
	_ "embed"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

//go:embed puzzle1_test.txt
var puzzle1Test string

func TestGenerateCombinations(t *testing.T) {
	sings := []string{"a", "b", "c"}

	combinations := new([]string)
	generateCombinations(sings, len(sings), "", combinations)

	assert.Equal(t, 27, len(*combinations))
	for _, combination := range *combinations {
		fmt.Println(combination)
	}
}

func TestAddOrMulti(t *testing.T) {
	assert.Equal(t, 3, addOrMulti([]string{"1", "1", "1"}, "++"))
	assert.Equal(t, 1, addOrMulti([]string{"1", "1", "1"}, "**"))
	assert.Equal(t, 27, addOrMulti([]string{"3", "3", "3"}, "**"))
	assert.Equal(t, 18, addOrMulti([]string{"3", "3", "3"}, "+*"))
	assert.Equal(t, 4, addOrMulti([]string{"1", "1", "1", "1"}, "+++"))
}

func TestCheckOperator(t *testing.T) {
	sings := []string{"+", "*"}

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

func TestPart1(t *testing.T) {
	actual := part1(puzzle1Test)

	assert.Equal(t, 3749, actual)
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = part1(puzzle1Test)
	}
	b.ReportAllocs()
}
