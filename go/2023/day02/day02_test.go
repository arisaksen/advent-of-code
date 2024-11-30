package main

import (
	utils "aoc/golang"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {

	t.Run("getNumberFromString", func(t *testing.T) {
		testString := ": 3 blue"
		expected := 3
		actual := getNumberFromString(testString)
		assert.Equal(t, expected, actual)
	})

	exampleLines := utils.ReadLines("example.txt")
	t.Run("game", func(t *testing.T) {
		expected := 8
		actual := game(exampleLines)
		assert.Equal(t, expected, actual)
	})

	puzzle := utils.ReadLines("puzzle.txt")
	t.Run("game", func(t *testing.T) {
		expected := 2776
		actual := game(puzzle)
		assert.Equal(t, expected, actual)
	})

}
