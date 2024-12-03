package main

import (
	_ "embed"
	"fmt"
	"log"
	"log/slog"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//go:embed puzzle1.txt
var puzzle1 string

var (
	re      = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	digitRe = regexp.MustCompile("[0-9]+")
	doRe    = regexp.MustCompile(`do\(\)`)
	dontRe  = regexp.MustCompile(`don't\(\)`)
)

func part1(puzzle string) int {

	matches := re.FindAllString(puzzle, -1)

	sum := 0
	for _, match := range matches {
		slog.Debug(match)
		digitOnlyString := strings.Split(match, ",")
		digit0String := digitRe.FindString(digitOnlyString[0])
		digit0, err := strconv.Atoi(digit0String)
		if err != nil {
			log.Fatal(err)
		}
		digit1String := digitRe.FindString(digitOnlyString[1])
		digit1, err := strconv.Atoi(digit1String)
		if err != nil {
			log.Fatal(err)
		}

		sum += digit0 * digit1
	}

	return sum
}

func main() {
	start1 := time.Now()
	fmt.Println()
	fmt.Println("Part 1:", part1(puzzle1))
	fmt.Println(time.Since(start1))

	//start2 := time.Now()
	//fmt.Println()
	//fmt.Println("Part 2:", part2(puzzle1))
	//fmt.Println(time.Since(start2))
}
