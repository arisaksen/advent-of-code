package main

import (
	utils "aoc/golang"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Part 1: ", part1("day05/puzzle.txt"))
	fmt.Println("Part 2: ", part2("day05/puzzle.txt"))
}

func part1(filename string) int {
	start := time.Now()

	puzzle2Input := utils.ReadText(filename)
	categoryMap := createMap(puzzle2Input)
	seeds := categoryMap["seeds"][0]

	soil := destinationToSource(seeds, categoryMap["seed-to-soil"])
	fertilizer := destinationToSource(soil, categoryMap["soil-to-fertilizer"])
	water := destinationToSource(fertilizer, categoryMap["fertilizer-to-water"])
	light := destinationToSource(water, categoryMap["water-to-light"])
	temp := destinationToSource(light, categoryMap["light-to-temperature"])
	humidity := destinationToSource(temp, categoryMap["temperature-to-humidity"])
	location := destinationToSource(humidity, categoryMap["humidity-to-location"])

	var lowestNumber int
	for _, number := range location {
		if number < lowestNumber || lowestNumber == 0 {
			lowestNumber = number
		}
	}

	duration := time.Since(start)
	fmt.Print("Execution time: ")
	fmt.Println(duration)

	return lowestNumber
}

func destinationToSource(values []int, toMaps [][]int) []int {
	var correspondingNumbers []int
	for _, value := range values {

		noValueSet := true
		var correspondingNumber int
		for _, toMap := range toMaps {
			destinationRangeStart := toMap[0]
			sourceRangeStart := toMap[1]
			rangeLength := toMap[2]

			//destinationRangeEnd := destinationRangeStart + rangeLength
			sourceRangeEnd := sourceRangeStart + rangeLength

			if value >= sourceRangeStart && value < sourceRangeEnd {
				correspondingNumber = value + destinationRangeStart - sourceRangeStart
				noValueSet = false
			}
		}
		if noValueSet {
			correspondingNumber = value
		}
		correspondingNumbers = append(correspondingNumbers, correspondingNumber)
	}

	return correspondingNumbers
}

func createMap(input string) map[string][][]int {
	lines := strings.Split(input, "\n\n")

	var categoryMap = make(map[string][][]int)
	for _, line := range lines {
		splitString := strings.Split(line, ":")
		key := strings.TrimSuffix(splitString[0], " map")
		value := cleanValue(splitString[1])

		categoryMap[key] = value
	}

	return categoryMap
}

func cleanValue(values string) [][]int {
	valuesTrimmedFirstSpace := strings.TrimPrefix(values, " ")
	valuesTrimmedFirstNewline := strings.TrimPrefix(valuesTrimmedFirstSpace, "\n")
	valuesSplitAtNewLine := strings.Split(valuesTrimmedFirstNewline, "\n")

	var intValuesList [][]int
	for _, line := range valuesSplitAtNewLine {
		valuesSplitAtSpace := strings.Split(line, " ")

		var intValues []int
		for _, stringValue := range valuesSplitAtSpace {
			valueAsInt, err := strconv.Atoi(stringValue)
			if err != nil {
				panic("Value '" + stringValue + "' can not be converted to int")
			}
			intValues = append(intValues, valueAsInt)
		}
		intValuesList = append(intValuesList, intValues)
	}

	return intValuesList
}
