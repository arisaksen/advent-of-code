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
	var lines = strings.Split(input, "\n")
	var maxX = len(lines[0]) - 1
	var maxY = len(lines) - 1

	var antenna [256][]position
	var antiNodes = set.NewSet[position]()
	for i, line := range lines {
		for j, c := range line {
			if c == '.' {
				continue
			}
			antenna[uint8(c)] = append(antenna[uint8(c)], position{x: j, y: i})
		}
	}

	for _, s := range antenna {
		if len(s) > 0 {
			for i, a1 := range s[:len(s)-1] {
				for _, a2 := range s[i+1:] {
					dx := a2.x - a1.x
					dy := a2.y - a1.y
					addAntiNodes(a2.x, a2.y, dx, dy, maxX, maxY, antiNodes)
					addAntiNodes(a1.x, a1.y, -dx, -dy, maxX, maxY, antiNodes)
				}
			}
		}
	}

	return antiNodes.Len()
}

func part1(input string) int {
	var addAntiNodes = func(ax, ay int, dx, dy int, maxX, maxY int, antiNodes set.Set[position]) {
		ax += dx
		ay += dy
		if ax >= 0 && ax <= maxX && ay >= 0 && ay <= maxY {
			antiNodes.Add(position{x: ax, y: ay})
		}
	}
	return getAntiNodes(input, addAntiNodes)
}

func part2(input string) int {
	var addAntiNodes = func(ax, ay int, dx, dy int, maxX, maxY int, antinodes set.Set[position]) {
		for ax >= 0 && ax <= maxX && ay >= 0 && ay <= maxY {
			antinodes.Add(position{x: ax, y: ay})
			ax += dx
			ay += dy
		}
	}
	return getAntiNodes(input, addAntiNodes)
}

func main() {
	fmt.Println("--2024 day 08 solution--")
	//var inputDay = inputTest
	start := time.Now()
	fmt.Println("part1: ", part1(puzzle1))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println("part2: ", part2(puzzle1))
	fmt.Println(time.Since(start))
}
