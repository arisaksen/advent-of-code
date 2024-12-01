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
	// mixed feelings about moving every var to the start
	var (
		inputLines       []string
		leftNumbers      []int
		rightNumbers     []int
		distanceSumTotal int
	)
	for _, line := range strings.Split(strings.TrimSuffix(puzzle, "\n"), "\n") {
		inputLines = append(inputLines, line)
	}

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
	sort.Ints(leftNumbers)
	//sort.Slice(leftNumbers, func(i, j int) bool {
	//	return leftNumbers[i] < leftNumbers[j]
	//})
	sort.Ints(rightNumbers)
	//sort.Slice(rightNumbers, func(i, j int) bool {
	//	return rightNumbers[i] < rightNumbers[j]
	//})

	for i := 0; i < len(leftNumbers); i++ {
		smallestLeft := leftNumbers[i]
		smallestRight := rightNumbers[i]

		distanceSum := smallestRight - smallestLeft

		//if distanceSum < 0 {
		//	distanceSumTotal += -1 * distanceSum
		//} else {
		//	distanceSumTotal += distanceSum
		//}
		distanceSumTotal += int(math.Abs(float64(distanceSum)))
		//log.Println("smallestLeft:", smallestLeft, "smallestRight:", smallestRight, "-", "distance:", distanceSum, "sizeLeft:", len(leftNumbers), "sum:", distanceSumTotal)

	}

	return distanceSumTotal
}
