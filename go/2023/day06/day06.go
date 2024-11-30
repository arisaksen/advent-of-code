package main

import (
	utils "aoc/golang"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1: ", part1("day06/puzzle.txt"))
	fmt.Println("Part 2: ", part2("day06/puzzle.txt"))
}

func part1(filename string) int {
	puzzle1Input := utils.ReadLines(filename)

	times := extractNumbers(puzzle1Input[0])
	distances := extractNumbers(puzzle1Input[1])

	var numberOfWaysToWinPerRace []int
	for i, time := range times {
		numberOfWaysToWinPerRace = append(numberOfWaysToWinPerRace, numberOfWins(time, distances[i]))
	}

	var winWaysMultiplied = 1
	for _, number := range numberOfWaysToWinPerRace {
		winWaysMultiplied = winWaysMultiplied * number
	}

	return winWaysMultiplied
}

func part2(filename string) int {
	puzzle2Input := utils.ReadLines(filename)

	time := extractNumber(puzzle2Input[0])
	distance := extractNumber(puzzle2Input[1])

	return numberOfWins(time, distance)
}

func extractNumber(input string) int {
	removeLabel := strings.Split(input, ":")[1]
	removeSpaces := strings.ReplaceAll(removeLabel, " ", "")
	number, _ := strconv.Atoi(removeSpaces)

	return number
}

func extractNumbers(input string) []int {
	// Split the string by whitespace
	parts := strings.Fields(input)

	// The first part is the label, so we discard it
	parts = parts[1:]

	var numbers []int
	for _, part := range parts {
		number, _ := strconv.Atoi(part)
		numbers = append(numbers, number)
	}
	return numbers
}

func numberOfWins(time int, distance int) int {
	timeToBeat := time
	maxHoldTime := time

	var waysToWin []int
	for holdTime := 1; holdTime < maxHoldTime; holdTime++ {
		speed := holdTime
		raceTime := distance / speed
		totalRaceTime := raceTime + holdTime
		if totalRaceTime < timeToBeat {
			waysToWin = append(waysToWin, holdTime)
		}
		//fmt.Printf("timeToBeat: %d, totalRaceTime: %d", timeToBeat, totalRaceTime)
		//fmt.Println()
	}

	return len(waysToWin)
}
