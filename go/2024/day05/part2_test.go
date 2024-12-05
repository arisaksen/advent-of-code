package main

import (
	_ "embed"
	"fmt"
	"github.com/lmittmann/tint"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"testing"
	"time"
)

func TestCorrectOrder(t *testing.T) {
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

	assert.Equal(t, "97,75,47,61,53", correctOrder("75,97,47,61,53", pageOrderingRulesMap))
	assert.Equal(t, "61,29,13", correctOrder("61,13,29", pageOrderingRulesMap))
	assert.Equal(t, "97,75,47,29,13", correctOrder("97,13,75,29,47", pageOrderingRulesMap))
}

func TestCorrectOrders(t *testing.T) {
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

			rightOrder, middle := isRightOrder(pageOrderingRulesMap, test.updates)
			if !rightOrder {
				fmt.Println(middle)
			}

			assert.Equal(t, test.expected, rightOrder)
			//assert.Equal(t, test.middleReturned, middle)
		})
	}

}

func TestPart2(t *testing.T) {
	slog.SetDefault(
		slog.New(
			tint.NewHandler(os.Stdout, &tint.Options{
				AddSource:  true,
				Level:      slog.LevelDebug,
				TimeFormat: time.Kitchen,
			}),
		),
	)

	actual := part2(puzzle1Test)

	assert.Equal(t, 123, actual)
}
