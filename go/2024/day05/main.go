package main

import (
	_ "embed"
	"fmt"
	"log/slog"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

//go:embed puzzle1.txt
var puzzle1 string

func isCurrentPageAfter(currentPage string, nextPageOrderingRules []string) bool {
	return slices.Contains(nextPageOrderingRules, currentPage)
}

func isRightOrder(pageOrderingRules map[string][]string, update string) (bool, string) {
	pages := strings.Split(update, ",")
	slog.Debug("pages", slog.Any("pages", pages))

	for i, page := range pages {
		currentPageOrdering := pageOrderingRules[page]
		slog.Debug("handling",
			slog.String("page", page),
			slog.Any("currentPageOrdering", currentPageOrdering),
		)

		for nextPageIndex := i + 1; nextPageIndex < len(pages); nextPageIndex++ {
			nextPage := pages[nextPageIndex]

			if !isPageBefore(page, nextPage, pageOrderingRules) {
				slog.Debug("incorrect order found",
					slog.String("currentPage", page),
					slog.String("nextPage", nextPage),
				)
				return false, update
			}
		}
	}

	return true, update
}

func isPageBefore(currentPage string, nextPage string, pageOrderingRules map[string][]string) bool {
	for _, afterPage := range pageOrderingRules[currentPage] {
		if afterPage == nextPage {
			return true
		}
	}
	return false
}

func createPageOrderingRuleMap(input string) map[string][]string {
	pageOrderingRulesMap := map[string][]string{}

	pageOrderingRules := strings.Split(input, "\n")
	for _, pageOrderingRule := range pageOrderingRules {
		pageOrderingRuleSplit := strings.Split(pageOrderingRule, "|")
		pageOrderingRulesMap[pageOrderingRuleSplit[0]] = append(
			pageOrderingRulesMap[pageOrderingRuleSplit[0]], pageOrderingRuleSplit[1],
		)
	}

	return pageOrderingRulesMap
}

func part1(puzzle string) int {
	var inputLines []string
	for _, line := range strings.Split(puzzle, "\n\n") {
		inputLines = append(inputLines, line)
	}
	pageOrderingRulesMap := createPageOrderingRuleMap(inputLines[0])
	updates := strings.Split(inputLines[1], "\n")
	slog.Debug("inputs split printed for read",
		slog.Any("pageOrderingRulesMap", pageOrderingRulesMap),
		slog.Any("updates", updates),
	)

	middleSum := 0
	for _, update := range updates {
		rightOrder, output := isRightOrder(pageOrderingRulesMap, update)

		if rightOrder {
			lengthMiddle := len(output)
			middleString := output[(lengthMiddle/2 - 1):(lengthMiddle/2 + 1)]
			middleInt, _ := strconv.Atoi(middleString)
			middleSum += middleInt
		}
	}

	return middleSum
}

func correctOrder(update string, pageOrderingRules map[string][]string) string {
	pages := strings.Split(update, ",")
	precedence := map[string]int{}

	for page, dependencies := range pageOrderingRules {
		for _, dependency := range dependencies {
			precedence[dependency]++
		}
		if _, exists := precedence[page]; !exists {
			precedence[page] = 0
		}
	}

	sort.SliceStable(pages, func(i, j int) bool {
		return precedence[pages[i]] < precedence[pages[j]]
	})

	return strings.Join(pages, ",")
}

func part2(puzzle string) int {
	inputLines := strings.Split(puzzle, "\n\n")
	pageOrderingRulesMap := createPageOrderingRuleMap(inputLines[0])
	updates := strings.Split(inputLines[1], "\n")

	middleSum := 0
	for _, update := range updates {
		rightOrder, output := isRightOrder(pageOrderingRulesMap, update)

		if !rightOrder {
			ordered := correctOrder(output, pageOrderingRulesMap)
			length := len(ordered)

			if length%2 == 0 && length > 1 {
				middleString := ordered[(length/2 - 1):(length/2 + 1)]
				middleInt, err := strconv.Atoi(middleString)
				if err != nil {
					slog.Error("Failed to parse middle string to int",
						slog.String("middleString", middleString),
					)
					continue
				}
				middleSum += middleInt
			}
		}
	}

	return middleSum
}

func main() {
	start1 := time.Now()
	fmt.Println()
	fmt.Println("Part 1:", part1(puzzle1))
	fmt.Println(time.Since(start1))

	//slog.SetDefault(
	//	slog.New(
	//		tint.NewHandler(os.Stdout, &tint.Options{
	//			AddSource:  true,
	//			Level:      slog.LevelDebug,
	//			TimeFormat: time.Kitchen,
	//		}),
	//	),
	//)

	start2 := time.Now()
	fmt.Println()
	fmt.Println("Part 2:", part2(puzzle1))
	fmt.Println(time.Since(start2))
}

// 5112 answer low
