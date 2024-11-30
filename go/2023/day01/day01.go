package main

import (
	utils "aoc/golang"
	"fmt"
	"strings"
	"unicode"
)

func main() {

	//part1()

	part2()

}
func part1() {
	puzzle1Input := utils.ReadLines("day01/puzzle1.txt")

	calibrations := getCalibrations(puzzle1Input)
	sumCalibrations := sumUpCalibrations(calibrations)
	fmt.Println("Part 1: ", sumCalibrations)
}

func part2() {
	puzzle2Input := utils.ReadLines("day01/puzzle2.txt")
	var calibrationsEdited []string
	for _, line := range puzzle2Input {
		calibrationsEdited = append(calibrationsEdited, addNumbersBeforeWords(line))
	}
	calibrations := getCalibrations(calibrationsEdited)
	sumCalibrations := sumUpCalibrations(calibrations)
	fmt.Println("Part 2: ", sumCalibrations)
}

func getCalibrations(lines []string) []int {
	var calibrations []int

	for _, line := range lines {
		var firstDigit, lastDigit rune
		var firstDigitFound bool
		var lastDigitIndex int

		fmt.Println(line)
		for y, char := range line {
			if unicode.IsDigit(char) && !firstDigitFound {
				firstDigit = char
				firstDigitFound = true
				fmt.Printf("firstDigit found at %d: %c \n", y, char)
			}
			if unicode.IsDigit(char) {
				lastDigit = char
				lastDigitIndex = y
			}
		}
		fmt.Printf("lastDigit found at %d: %c \n", lastDigitIndex, lastDigit)
		firstDigitInt := int(firstDigit - '0')
		lastDigitInt := int(lastDigit - '0')
		combineDigits := firstDigitInt*10 + lastDigitInt

		calibrations = append(calibrations, combineDigits)
	}

	return calibrations
}

func addNumbersBeforeWords(input string) string {
	var spelledDigits = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	var result strings.Builder
	i := 0
	length := len(input)

	for i < length {
		matchFound := false
		for word, digit := range spelledDigits {
			if strings.HasPrefix(input[i:], word) {
				result.WriteString(digit + word)
				i += len(word)
				matchFound = true
				break
			}
		}
		if !matchFound {
			result.WriteByte(input[i])
			i++
		}
	}
	return result.String()
}

func sumUpCalibrations(calibrations []int) int {
	var sum = 0
	for _, number := range calibrations {
		sum += number
	}

	return sum
}
