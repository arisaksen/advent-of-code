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

//go:embed puzzle2.txt
var puzzle2 string

func findAndRemoveSmallestNumber(numbers *[]int) int {
	if len(*numbers) < 1 {
		log.Fatal("expect numbers to be greater than 0")
	}

	smallest := (*numbers)[0]
	smallestIndex := 0
	for i := 0; i < len(*numbers); i++ {
		if (*numbers)[i] < smallest {
			smallest = (*numbers)[i]
			smallestIndex = i
		}
	}
	removeIndex(numbers, smallestIndex)

	return smallest
}

func removeIndex(slice *[]int, index int) {
	if index < 0 || index >= len(*slice) {
		log.Fatal("out of bound")
	}
	copy((*slice)[index:], (*slice)[index+1:])
	*slice = (*slice)[:len(*slice)-1]
}

func part1(puzzle string) int {
	var inputLines []string
	for _, line := range strings.Split(strings.TrimSuffix(puzzle, "\n"), "\n") {
		inputLines = append(inputLines, line)
	}

	var leftNumbers []int
	var rightNumbers []int
	for _, line := range inputLines {
		numberStrings := strings.Split(line, "   ")
		leftNumber, err := strconv.Atoi(numberStrings[0])
		if err != nil {
			log.Fatal(err)
		}
		leftNumbers = append(leftNumbers, leftNumber)
		rightNumber, err := strconv.Atoi(numberStrings[1])
		if err != nil {
			log.Fatal(err)
		}
		rightNumbers = append(rightNumbers, rightNumber)
	}

	var distanceSumTotal int
	for {
		smallestLeft := findAndRemoveSmallestNumber(&leftNumbers)
		smallestRight := findAndRemoveSmallestNumber(&rightNumbers)

		distanceSum := smallestRight - smallestLeft

		distanceSumTotal += int(math.Abs(float64(distanceSum)))
		//log.Println("smallestLeft:", smallestLeft, "smallestRight:", smallestRight, "-", "distance:", distanceSum, "sizeLeft:", len(leftNumbers), "sum:", distanceSumTotal)

		if len(leftNumbers) == 0 {
			break
		}
	}

	return distanceSumTotal
}

func part2(puzzle string) int {
	var inputLines []string
	for _, line := range strings.Split(strings.TrimSuffix(puzzle, "\n"), "\n") {
		inputLines = append(inputLines, line)
	}

	var leftNumbers []int
	var rightNumbers []int
	for _, line := range inputLines {
		numberStrings := strings.Split(line, "   ")
		leftNumber, err := strconv.Atoi(numberStrings[0])
		if err != nil {
			log.Fatal(err)
		}
		leftNumbers = append(leftNumbers, leftNumber)
		rightNumber, err := strconv.Atoi(numberStrings[1])
		if err != nil {
			log.Fatal(err)
		}
		rightNumbers = append(rightNumbers, rightNumber)
	}

	var simularityScore int
	for _, numberFirst := range leftNumbers {
		var appearanceCount int
		for _, numberSecond := range rightNumbers {
			if numberSecond == numberFirst {
				appearanceCount++
			}
		}

		simularityScore += numberFirst * appearanceCount
		//log.Println("numberFirst:", numberFirst, "appearenceCount:", appearanceCount, "-", "simScore:", simularityScore)
	}

	return simularityScore
}

func main() {
	start1 := time.Now()
	fmt.Println()
	fmt.Println("Part 1: ", part1(puzzle1))
	fmt.Println(time.Since(start1))

	start2 := time.Now()
	fmt.Println()
	fmt.Println("Part 2: ", part2(puzzle2))
	fmt.Println(time.Since(start2))
}
