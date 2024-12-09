package main

import (
	_ "embed"
	"github.com/lmittmann/tint"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"os"
	"testing"
	"time"
)

//go:embed puzzle1_test.txt
var puzzle1Test string

func function123() int {
	slog.Debug("testmessage")

	return 1
}

func TestCreatePositionMap(t *testing.T) {
	slog.SetDefault(
		slog.New(
			tint.NewHandler(os.Stdout, &tint.Options{
				AddSource:  true,
				Level:      slog.LevelDebug,
				TimeFormat: time.Kitchen,
			}),
		),
	)
	input := `..........
...#......
..........
....a.....
..........
.....a....
..........
......#...
..........
..........`
	antennaMap, gridSizeX, gridSizeY := createPositionMap(input)
	assert.Equal(t, 2, len(antennaMap['a']))
	assert.Equal(t, 10, gridSizeX)
	assert.Equal(t, 10, gridSizeY)
}

func TestGetSingleAntiNode(t *testing.T) {
	slog.SetDefault(
		slog.New(
			tint.NewHandler(os.Stdout, &tint.Options{
				AddSource:  true,
				Level:      slog.LevelDebug,
				TimeFormat: time.Kitchen,
			}),
		),
	)

	assert.Equal(t, position{1, 3}, position{3, 4}.getAntiNode(position{5, 5}))
	assert.Equal(t, position{7, 6}, position{5, 5}.getAntiNode(position{3, 4}))
}

func TestGetAntiNodes(t *testing.T) {
	slog.SetDefault(
		slog.New(
			tint.NewHandler(os.Stdout, &tint.Options{
				AddSource:  true,
				Level:      slog.LevelInfo,
				TimeFormat: time.Kitchen,
			}),
		),
	)
	input := `..........
	...#......
	#.........
	....a.....
	........a.
	.....a....
	..#.......
	......#...
	..........
	..........`
	antennaMap, gridSizeX, gridSizeY := createPositionMap(input)

	t.Log(antennaMap['#'])
	expectedAnitNodes := antennaMap['#']

	antiNodes := getAntiNodesForSingleAntennaType(antennaMap['a'], gridSizeX, gridSizeY)
	assert.Equal(t, len(expectedAnitNodes), len(antiNodes))

	t.Log(antiNodes)
	for _, antiNode := range antiNodes {
		assert.Contains(t, expectedAnitNodes, antiNode)
	}
}

func TestPart1(t *testing.T) {
	actual := part1(puzzle1Test)

	assert.Equal(t, 14, actual)
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = part1(puzzle1Test)
	}
	b.ReportAllocs()
}
