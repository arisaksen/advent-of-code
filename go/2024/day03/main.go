package main

import (
	_ "embed"
	"fmt"
	"log/slog"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// RegexTester plugin
// "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
// "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

var (

	//go:embed puzzle1.txt
	puzzle1 string

	re       = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	digitRe  = regexp.MustCompile("[0-9]+")
	reDoDont = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
)

func matchToDigits(match string) (int, int, error) {
	digitOnlyString := strings.Split(match, ",")
	digit0String := digitRe.FindString(digitOnlyString[0])
	digit0, err := strconv.Atoi(digit0String)
	if err != nil {
		return 0, 0, err
	}
	digit1String := digitRe.FindString(digitOnlyString[1])
	digit1, err := strconv.Atoi(digit1String)
	if err != nil {
		return 0, 0, err
	}

	return digit0, digit1, nil
}

func part1(puzzle string) int {
	matches := re.FindAllString(puzzle, -1)

	sum := 0
	for _, match := range matches {
		slog.Debug(match)
		digit0, digit1, err := matchToDigits(match)
		if err != nil {
			slog.Error("Error matchToDigits", slog.Any("error", err))
		}
		sum += digit0 * digit1
	}

	return sum
}

func part2(puzzle string) int {

	matches := reDoDont.FindAllString(puzzle, -1)

	sum := 0
	enable := true
	for _, match := range matches {

		if match == "do()" {
			enable = true
		} else if match == "don't()" {
			enable = false
		}

		if enable && match != "do()" {
			slog.Debug("sum", slog.Bool("enable", true), slog.String("match", match))
			digit0, digit1, err := matchToDigits(match)
			if err != nil {
				slog.Error("Error matchToDigits", slog.Any("error", err))
			}
			sum += digit0 * digit1
		} else {
			slog.Debug("sum", slog.Bool("enable", false), slog.String("match", match))
		}

	}

	return sum
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
