package main

import (
	"aoc/golang/utils/set"
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed puzzle1.txt
var puzzle1 string

type position struct {
	x, y int
}

func getAntiNodes(input string, addAntiNodes func(ax, ay int, dx, dy int, maxX, maxY int, antiNodes set.Set[position])) int {
	antiNodes := set.NewSet[position]()

	inputLines := strings.Split(input, "\n")
	maxX := len(inputLines[0]) - 1
	maxY := len(inputLines) - 1

	antenna := make(map[rune][]position)
	for y, line := range inputLines {
		for x, char := range line {
			if char != '.' {
				antenna[char] = append(antenna[char], position{x, y})
			}
		}
	}

	for _, s := range antenna {
		for i, a1 := range s[:len(s)-1] {
			for _, a2 := range s[i+1:] {
				dx := a2.x - a1.x
				dy := a2.y - a1.y
				addAntiNodes(a2.x, a2.y, dx, dy, maxX, maxY, antiNodes)
				addAntiNodes(a1.x, a1.y, -dx, -dy, maxX, maxY, antiNodes)
			}
		}
	}

	return antiNodes.Len()
}

func part1(input string) int {
	addAntiNodes := func(ax, ay int, dx, dy int, maxX, maxY int, antiNodes set.Set[position]) {
		ax += dx
		ay += dy
		if ax >= 0 && ax <= maxX && ay >= 0 && ay <= maxY {
			antiNodes.Add(position{x: ax, y: ay})
		}
	}
	return getAntiNodes(input, addAntiNodes)
}

func part2(input string) int {
	addAntiNodes := func(ax, ay int, dx, dy int, maxX, maxY int, antiNodes set.Set[position]) {
		for ax >= 0 && ax <= maxX && ay >= 0 && ay <= maxY {
			antiNodes.Add(position{x: ax, y: ay})
			ax += dx
			ay += dy
		}
	}
	return getAntiNodes(input, addAntiNodes)
}

func main() {
	start := time.Now()
	fmt.Println("part1: ", part1(puzzle1))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", part2(puzzle1))
	fmt.Println(time.Since(start))
}
