package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed puzzle1.txt
var puzzle1 string

//go:embed puzzle2.txt
var puzzle2 string

func part1(puzzle string) int {
	var inputLines []string
	for _, line := range strings.Split(strings.TrimSuffix(puzzle, "\n"), "\n") {
		inputLines = append(inputLines, line)
	}

	return 1
}

func part2(puzzle string) int {
	var inputLines []string
	for _, line := range strings.Split(strings.TrimSuffix(puzzle, "\n"), "\n") {
		inputLines = append(inputLines, line)
	}

	return 1
}

func main() {
	start1 := time.Now()
	fmt.Println()
	fmt.Println("Part 1:", part1(puzzle1))
	fmt.Println(time.Since(start1))

	start2 := time.Now()
	fmt.Println()
	fmt.Println("Part 2:", part2(puzzle2))
	fmt.Println(time.Since(start2))
}
	