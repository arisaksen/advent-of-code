package main

import (
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

/* This was done after submit of main.go
I suspected that this could be improved by sorting the list of numbers first.
We can then just pick current index as smallest.
This way we don't need to loop over the whole lists of numbers each time.
*/

func part1Optimized(puzzle string) int {
	var leftNumbers, rightNumbers []int

	for _, line := range strings.Split(strings.TrimSpace(puzzle), "\n") {
		numberStrings := strings.Fields(line)
		leftNumber, err := strconv.Atoi(numberStrings[0])
		if err != nil {
			log.Fatal(err)
		}
		rightNumber, err := strconv.Atoi(numberStrings[1])
		if err != nil {
			log.Fatal(err)
		}

		leftNumbers = append(leftNumbers, leftNumber)
		rightNumbers = append(rightNumbers, rightNumber)
	}
	sort.Ints(leftNumbers)
	//sort.Slice(leftNumbers, func(i, j int) bool {
	//	return leftNumbers[i] < leftNumbers[j]
	//})
	sort.Ints(rightNumbers)
	//sort.Slice(rightNumbers, func(i, j int) bool {
	//	return rightNumbers[i] < rightNumbers[j]
	//})

	var distanceSumTotal int
	for i := 0; i < len(leftNumbers); i++ {
		distanceSumTotal += int(math.Abs(float64(rightNumbers[i] - leftNumbers[i])))
	}
	return distanceSumTotal
}
