package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed puzzle1.txt
var puzzle1 string

const (
	word         = "XMAS"
	reverseWord  = "SAMX"
	word2        = "MAS"
	reverseWord2 = "SAM"
)

// Check for occurrences of the word in a horizontal or vertical string
func countWordOccurrences(text string, word string) int {
	var count int
	var i int
	for {
		start := i
		var end int
		end = i + len(word)
		if end > len(text) {
			break
		}
		window := text[start:end]
		if window == word {
			count++
		}
		i++
	}

	return count
}

func countHorizontal(inputLines []string) int {
	count := 0
	for _, row := range inputLines {
		count += countWordOccurrences(row, word)
		count += countWordOccurrences(row, reverseWord)
	}
	return count
}

func countVertical(inputLines []string) int {
	count := 0
	numRows := len(inputLines)
	numCols := len(inputLines[0])
	for col := 0; col < numCols; col++ {
		var vertical string
		for row := 0; row < numRows; row++ {
			vertical += string(inputLines[row][col])
		}
		count += countWordOccurrences(vertical, word)
		count += countWordOccurrences(vertical, reverseWord)
	}
	return count
}

func countDiagonals(inputLines []string) int {
	count := 0
	numRows := len(inputLines)
	numCols := len(inputLines[0])

	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if row+len(word) <= numRows && col+len(word) <= numCols {
				diag := ""
				for k := 0; k < len(word); k++ {
					diag += string(inputLines[row+k][col+k])
				}
				if diag == word {
					count++
				}
			}
		}
	}

	for row := numRows - 1; row >= 0; row-- {
		for col := 0; col < numCols; col++ {
			if row-len(word)+1 >= 0 && col+len(word) <= numCols {
				diag := ""
				for k := 0; k < len(word); k++ {
					diag += string(inputLines[row-k][col+k])
				}
				if diag == word {
					count++
				}
			}
		}
	}

	for row := 0; row < numRows; row++ {
		for col := numCols - 1; col >= 0; col-- {
			if row+len(word) <= numRows && col-len(word)+1 >= 0 {
				diag := ""
				for k := 0; k < len(word); k++ {
					diag += string(inputLines[row+k][col-k])
				}
				if diag == word {
					count++
				}
			}
		}
	}

	for row := numRows - 1; row >= 0; row-- {
		for col := numCols - 1; col >= 0; col-- {
			if row-len(word)+1 >= 0 && col-len(word)+1 >= 0 {
				diag := ""
				for k := 0; k < len(word); k++ {
					diag += string(inputLines[row-k][col-k])
				}
				if diag == word {
					count++
				}
			}
		}
	}

	return count
}

func part1(puzzle string) int {
	inputLines := strings.Split(puzzle, "\n")

	horizontalCount := countHorizontal(inputLines)
	verticalCount := countVertical(inputLines)
	diagonalCount := countDiagonals(inputLines)

	return horizontalCount + verticalCount + diagonalCount
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
