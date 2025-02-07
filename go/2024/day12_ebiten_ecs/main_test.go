package main

import (
	_ "embed"
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		expectedPartOne int
		expectedPartTwo int
	}{
		{
			name:            "sample 1",
			input:           PuzzleTest1,
			expectedPartOne: 140,
			expectedPartTwo: 80,
		}, {
			name:            "sample 2",
			input:           PuzzleTest2,
			expectedPartOne: 772,
			expectedPartTwo: 436,
		}, {
			name:            "sample 3",
			input:           PuzzleTest3,
			expectedPartOne: 1930,
			expectedPartTwo: 1206,
		}, {
			name:            "sample 4",
			input:           PuzzleTest4,
			expectedPartOne: -1,
			expectedPartTwo: 236,
		}, {
			name:            "sample 5",
			input:           PuzzleTest5,
			expectedPartOne: -1,
			expectedPartTwo: 368,
		}, {
			name:            "Puzzle",
			input:           Puzzle1,
			expectedPartOne: 0,
			expectedPartTwo: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			var grid [][]byte
			grid = parseInput(test.input)
			if test.expectedPartOne != -1 {
				assert.Equal(t, test.expectedPartOne, part1(grid))
			}

			if test.expectedPartTwo != -1 {
				assert.Equal(t, test.expectedPartTwo, part2(grid))
			}
		})

	}

}

func parseInput(input string) [][]byte {
	lines := strings.Split(input, "\n")
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}
	return grid
}

func dfs(grid [][]byte, visited [][]bool, i, j int, sides *[][3]int) (int, int) {

	visited[i][j] = true
	neighbors := getNeighbors(i, j)
	perimeter := 0
	area := 1
	for _, n := range neighbors {
		ni, nj := n[0], n[1]
		if ni < 0 || ni >= len(grid) || nj < 0 || nj >= len(grid[ni]) || grid[ni][nj] != grid[i][j] {
			if sides != nil {
				*sides = append(*sides, [3]int{ni, nj, n[2]})
			}
			perimeter++
		} else if !visited[ni][nj] {
			a, p := dfs(grid, visited, ni, nj, sides)
			area += a
			perimeter += p
		}
	}
	return area, perimeter
}

func getNeighbors(i, j int) [][3]int {
	return [][3]int{
		{i - 1, j, 0},
		{i + 1, j, 1},
		{i, j - 1, 2},
		{i, j + 1, 3},
	}
}

func part1(grid [][]byte) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !visited[i][j] {
				area, perimeter := dfs(grid, visited, i, j, nil)
				fmt.Printf("Char: %s = Area: %d, Perimeter: %d\n", string(grid[i][j]), area, perimeter)
				sum += area * perimeter
			}
		}
	}
	return sum
}

func getSideCount(sides [][3]int) int {
	sideMap := make(map[[3]int]bool)

	sort.Slice(sides, func(i, j int) bool {
		if sides[i][0] == sides[j][0] {
			return sides[i][1] < sides[j][1]
		}
		return sides[i][0] < sides[j][0]
	})

	fmt.Println(sides)

	sideCount := 0

	for _, s := range sides {
		getCombinations := getNeighbors(s[0], s[1])
		combFound := false

		for _, c := range getCombinations {
			c[2] = s[2]
			if _, found := sideMap[c]; found {
				combFound = true
			}
		}
		if !combFound {
			sideCount++
		}

		sideMap[s] = true

	}

	return sideCount
}

func part2(grid [][]byte) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !visited[i][j] {
				sides := make([][3]int, 0)
				area, _ := dfs(grid, visited, i, j, &sides)
				sideCount := getSideCount(sides)
				fmt.Printf("Char: %s = Area: %d, Sides: %d\n", string(grid[i][j]), area, sideCount)
				sum += area * sideCount
			}
		}
	}
	return sum
}
