package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

//go:embed puzzle1.txt
var puzzle1 string

func solve(puzzle string, part2 bool) int {
	inputLines := strings.Split(puzzle, "\n")

	var total int
	for _, line := range inputLines {
		before, after, _ := strings.Cut(line, ":")
		equalNumber, _ := strconv.Atoi(before)

		var elements []int
		for _, e := range slices.Backward(strings.Fields(after)) {
			var el, _ = strconv.Atoi(e)
			elements = append(elements, el)
		}

		if check(equalNumber, elements, part2) {
			total += equalNumber
		}
	}
	return total
}

func check(goal int, elements []int, part2 bool) bool {
	if len(elements) == 1 {
		if elements[0] == goal {
			return true
		}
		return false
	}

	head := elements[0]
	tail := elements[1:]

	if next := goal - head; next >= 0 && check(next, tail, part2) {
		return true
	}
	if next := goal / head; next*head == goal && check(next, tail, part2) {
		return true
	}
	if part2 {
		var p = 1
		for h := head; h > 0; h /= 10 {
			p *= 10
		}
		if next := (goal - head) / p; next*p+head == goal && check(next, tail, part2) {
			return true
		}
	}
	return false
}

func part1(input string) int {
	return solve(input, false)
}

func part2(input string) int {
	return solve(input, true)
}

func main() {
	start := time.Now()
	fmt.Println("part1:", part1(puzzle1))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2:", part2(puzzle1))
	fmt.Println(time.Since(start))
}
