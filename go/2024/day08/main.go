package main

import (
	_ "embed"
	"fmt"
	"log/slog"
	"slices"
	"strings"
	"time"
)

//go:embed puzzle1.txt
var puzzle1 string

type position struct {
	x, y int
}

type vector struct {
	x, y int
}

func (a position) getAntiNode(b position) position {
	var vec vector
	vec = vector{
		x: a.x - b.x,
		y: a.y - b.y,
	}
	slog.Debug("getAntiNode", slog.Any("vector", vec))

	antiNode := position{
		x: a.x + vec.x,
		y: a.y + vec.y,
	}

	return antiNode
}

func createPositionMap(puzzle string) (map[rune][]position, int, int) {
	antennaMap := make(map[rune][]position)

	inputLines := strings.Split(puzzle, "\n")
	sizeX := len(inputLines[0]) - 1
	sizeY := len(inputLines) - 1

	for y, inputLine := range inputLines {
		for x, char := range inputLine {
			if char != '.' {
				antennaMap[char] = append(antennaMap[char], position{x, y})
			}
		}
	}
	slog.Debug("rune(int32) values", slog.Int("#", '#'), slog.Int("a", 'a'))
	slog.Debug("", slog.Any("antennaMap", antennaMap), slog.Int("x", sizeX), slog.Int("y", sizeY))

	return antennaMap, sizeX, sizeY
}

func getAntiNodesForSingleAntennaType(antennaLocations []position, gridSizeX int, gridSizeY int) []position {
	var antiNodes []position
	for i := 0; i < len(antennaLocations); i++ {
		currentAntenna := antennaLocations[i]
		for _, location := range antennaLocations {
			if location == currentAntenna {
				continue
			}
			antiNode := currentAntenna.getAntiNode(location)
			if antiNode.x >= 0 && antiNode.x <= gridSizeX && antiNode.y >= 0 && antiNode.y <= gridSizeY {
				antiNodes = append(antiNodes, antiNode)
			} else {
				slog.Warn("dropping anitNode", slog.Any("antiNode", antiNode))
			}
		}
	}

	return antiNodes
}

func part1(puzzle string) int {
	antennaMap, gridSizeX, gridSizeY := createPositionMap(puzzle)

	//total := 0
	var allAntiNodesCombined []position
	for k, v := range antennaMap {
		antiNodes := getAntiNodesForSingleAntennaType(v, gridSizeX, gridSizeY)
		slog.Info(fmt.Sprintf("antenna %c: %d", k, len(antiNodes)))
		for _, antiNode := range antiNodes {
			if slices.Contains(allAntiNodesCombined, antiNode) {
				slog.Warn("allAntiNodesCombined allready contains this", slog.Any("antiNode", antiNode))
				continue
			}
			allAntiNodesCombined = append(allAntiNodesCombined, antiNode)
		}
	}

	return len(allAntiNodesCombined)
}

func main() {
	start1 := time.Now()
	fmt.Println()
	fmt.Println("Part 1:", part1(puzzle1))
	fmt.Println(time.Since(start1))

	//start2 := time.Now()
	//fmt.Println()
	//fmt.Println("Part 2:", part2(puzzle2))
	//fmt.Println(time.Since(start2))
}
