package main

import (
	_ "embed"
	"github.com/lmittmann/tint"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"strconv"
	"testing"
	"time"
)

//go:embed puzzle1_test.txt
var puzzle1Test string

func TestIsRightOrder(t *testing.T) {
	slog.SetDefault(
		slog.New(
			tint.NewHandler(os.Stdout, &tint.Options{
				AddSource:  true,
				Level:      slog.LevelDebug,
				TimeFormat: time.Kitchen,
			}),
		),
	)

	assert.True(t, "70" > "69")
	/* This works because, in two-digit numbers,
	the first character corresponds to the tens place,
	and the second character corresponds to the units place,
	which matches how numbers are numerically compared.

	NB!
	Be cautious if there’s any chance that the numbers might not always be two digits,
	as lexicographical comparison could produce incorrect results in such cases.
	For instance:
	*/
	assert.True(t, "100" < "20")

	pageOrderingRules := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13`
	pageOrderingRulesMap := createPageOrderingRuleMap(pageOrderingRules)

	tests := []struct {
		name           string
		updates        string
		expected       bool
		middleReturned int
	}{
		{
			name:           "test example line 1",
			updates:        "75,47,61,53,29",
			expected:       true,
			middleReturned: 61,
		},
		{
			name:           "test example line 2",
			updates:        "97,61,53,29,13",
			expected:       true,
			middleReturned: 53,
		},
		{
			name:           "test example line 3",
			updates:        "75,29,13",
			expected:       true,
			middleReturned: 29,
		},
		{
			name:           "test example line 4",
			updates:        "75,97,47,61,53",
			expected:       false,
			middleReturned: 0,
		},
		{
			name:           "test example line 4",
			updates:        "61,13,29",
			expected:       false,
			middleReturned: 0,
		},
		{
			name:           "test example line 4",
			updates:        "97,13,75,29,47",
			expected:       false,
			middleReturned: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			rightOrder, output := isRightOrder(pageOrderingRulesMap, test.updates)
			middleSum := 0
			if rightOrder {
				lengthMiddle := len(output)
				middleString := output[(lengthMiddle/2 - 1):(lengthMiddle/2 + 1)]
				middleInt, _ := strconv.Atoi(middleString)
				middleSum += middleInt
			}

			assert.Equal(t, test.expected, rightOrder)
			assert.Equal(t, test.middleReturned, middleSum)
		})
	}

}

func TestIsCurrentPageAfter(t *testing.T) {
	// updates:  []string{"75,97,47,61,53"},  []string is the orderingRules for "97"
	assert.True(t, isCurrentPageAfter("75", []string{"13", "61", "47", "29", "53", "75"}))
}

func TestPart1(t *testing.T) {
	slog.SetDefault(
		slog.New(
			tint.NewHandler(os.Stdout, &tint.Options{
				AddSource:  true,
				Level:      slog.LevelDebug,
				TimeFormat: time.Kitchen,
			}),
		),
	)

	actual := part1(puzzle1Test)

	assert.Equal(t, 143, actual)
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = part1(puzzle1Test)
	}
	b.ReportAllocs()
}
