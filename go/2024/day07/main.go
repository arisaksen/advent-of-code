package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

//go:embed puzzle1.txt
var puzzle1 string

//go:embed puzzle2.txt
var puzzle2 string

var (
	operators = []string{"+", "*"}
)

func generateCombinations(signs []string, length int, current string, combinations *[]string) *[]string {
	if len(current) == length {
		*combinations = append(*combinations, current)
		return combinations
	}
	for _, sign := range signs {
		generateCombinations(signs, length, current+sign, combinations)
	}

	return combinations
}

func checkOperator(line string, operatorCombinations []string) int {
	lineSplit := strings.Split(line, ":")
	equalNumber, err := strconv.Atoi(lineSplit[0])
	if err != nil {
		panic(err)
	}
	operatorNumbersStrings := strings.Fields(lineSplit[1])

	total := 0
	for _, operatorCombination := range operatorCombinations {
		answer := addOrMulti(operatorNumbersStrings, operatorCombination)
		if answer == equalNumber {
			total += answer
			break
		}
	}

	return total
}

func addOrMulti(stringNumbers []string, combinations string) int {
	total := 0
	for i := 0; i < len(stringNumbers); i++ {
		number, err := strconv.Atoi(stringNumbers[i])
		if err != nil {
			log.Fatal(err)
		}
		if i == 0 {
			total += number
			continue
		}
		operator := combinations[i-1]
		if operator == '+' {
			total += number
		} else if operator == '*' {
			total *= number
		}
	}

	return total
}

func part1(puzzle string) int {
	var inputLines []string
	for _, line := range strings.Split(strings.TrimSuffix(puzzle, "\n"), "\n") {
		inputLines = append(inputLines, line)
	}
	for _, equation := range inputLines {
		fmt.Println(equation)
	}

	total := 0
	for _, line := range inputLines {
		operatorsLen := len(strings.Fields(line)) - 2
		operatorCombinations := new([]string)
		generateCombinations(operators, operatorsLen, "", operatorCombinations)

		total += checkOperator(line, *operatorCombinations)
	}

	return total
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

	//start2 := time.Now()
	//fmt.Println()
	//fmt.Println("Part 2:", part2(puzzle2))
	//fmt.Println(time.Since(start2))
}
