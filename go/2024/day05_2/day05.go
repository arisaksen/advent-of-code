package main

import (
	"aoc/golang/utils/set"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

//go:embed puzzle1.txt
var puzzle1 string

type pair struct {
	first  string
	second string
}

func parse(input string) (set.Set[pair], [][]string) {
	var parts = strings.Split(input, "\n\n")

	var orders = set.NewSet[pair]()
	var updates [][]string
	for _, line := range strings.Split(parts[0], "\n") {
		var first = line[:2]
		var second = line[3:5]
		orders.Add(pair{first, second})
	}

	for _, line := range strings.Split(parts[1], "\n") {
		var elements = strings.Split(line, ",")
		updates = append(updates, elements)
	}
	return orders, updates
}

func checkUpdate(orders set.Set[pair], update []string) bool {
	for i := 0; i < len(update)-1; i++ {
		var p = pair{update[i], update[i+1]}
		if !orders.Contains(p) {
			return false
		}
	}
	return true
}

func part1(input string) int {
	var orders, updates = parse(input)
	var res int
	for _, update := range updates {
		if checkUpdate(orders, update) {
			var middle, _ = strconv.Atoi(update[len(update)/2])
			res += middle
		}
	}
	return res
}

func part2(input string) int {
	var orders, updates = parse(input)
	var cmp = func(a, b string) int {
		if _, ok := orders[pair{a, b}]; ok {
			return -1
		}
		if _, ok := orders[pair{b, a}]; ok {
			return 1
		}
		return 0
	}
	var res int
	for _, update := range updates {
		if !checkUpdate(orders, update) {
			slices.SortFunc(update, cmp)
			var middle, _ = strconv.Atoi(update[len(update)/2])
			res += middle
		}
	}
	return res
}

func main() {
	start := time.Now()
	fmt.Println("part1: ", part1(puzzle1))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", part2(puzzle1))
	fmt.Println(time.Since(start))
}
