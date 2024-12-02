package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed puzzle1.txt
var puzzle1 string

func removeIndex(slice *[]int, index int) {
	if index < 0 || index >= len(*slice) {
		log.Fatal("out of bound")
	}
	copy((*slice)[index:], (*slice)[index+1:])
	*slice = (*slice)[:len(*slice)-1]
}

func checkIfSafe(report []int, initialProblemDampener int) bool {
	problemDampener := initialProblemDampener

	sateIncrease := report[0] < report[1]
	for i, _ := range report {
		if i+1 < len(report) {
			if !checkAdjacent(report[i], report[i+1], 3) {
				problemDampener--
				removeIndex(&report, i+1)
				if problemDampener < 0 {
					return false
				}
				continue
			}
			if (report[i] < report[i+1]) != sateIncrease {
				removeIndex(&report, i+1)
				problemDampener--
				if problemDampener < 0 {
					return false
				}
				continue
			}
			if report[i] == report[i+1] {
				removeIndex(&report, i+1)
				problemDampener--
				if problemDampener < 0 {
					return false
				}
				continue
			}
		}
		if i+2 < len(report) {
			if !checkAdjacent(report[i+1], report[i+2], 3) {
				removeIndex(&report, i+2)
				problemDampener--
				if problemDampener < 0 {
					return false
				}
				continue
			}
			if (report[i] < report[i+2]) != sateIncrease {
				removeIndex(&report, i+2)
				problemDampener--
				if problemDampener < 0 {
					return false
				}
				continue
			}
		}
	}
	return true
}

func checkAdjacent(a int, b int, maxLevelDiff int) bool {
	distance := math.Abs(float64(a) - float64(b))
	return distance <= float64(maxLevelDiff)
}

func part1(puzzle string) int {
	var inputLines []string
	for _, line := range strings.Split(strings.TrimSuffix(puzzle, "\n"), "\n") {
		inputLines = append(inputLines, line)
	}
	var reports [][]int
	for _, line := range inputLines {
		chars := strings.Split(line, " ")

		var numbers []int
		for _, char := range chars {
			number, err := strconv.Atoi(char)
			if err != nil {
				log.Fatal(err, char)
			}
			numbers = append(numbers, number)
		}
		reports = append(reports, numbers)
	}

	var safeReportCounter int
	for _, report := range reports {
		if checkIfSafe(report, 0) {
			safeReportCounter++
		}
	}

	return safeReportCounter
}

func part2(puzzle string) int {
	var inputLines []string
	for _, line := range strings.Split(strings.TrimSuffix(puzzle, "\n"), "\n") {
		inputLines = append(inputLines, line)
	}
	var reports [][]int
	for _, line := range inputLines {
		chars := strings.Split(line, " ")

		var numbers []int
		for _, char := range chars {
			number, err := strconv.Atoi(char)
			if err != nil {
				log.Fatal(err, char)
			}
			numbers = append(numbers, number)
		}
		reports = append(reports, numbers)
	}

	var safeReportCounter int
	for i, report := range reports {
		if checkIfSafe(report, 1) {
			safeReportCounter++
		} else {
			fmt.Println("unsafe report index", i, report)
		}
	}

	return safeReportCounter
}

func main() {
	start1 := time.Now()
	fmt.Println()
	fmt.Println("Part 1:", part1(puzzle1))
	fmt.Println(time.Since(start1))

	start2 := time.Now()
	fmt.Println()
	fmt.Println("Part 2:", part2(puzzle1))
	fmt.Println(time.Since(start2))
}
