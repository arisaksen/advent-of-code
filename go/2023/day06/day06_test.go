package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		expected := 288
		actual := part1("example.txt")
		assert.Equal(t, expected, actual)
	})

	t.Run("Part 2", func(t *testing.T) {
		expected := 71503
		actual := part2("example.txt")
		assert.Equal(t, expected, actual)
	})

}

func TestUnits(t *testing.T) {

	t.Run("numberOfWins7", func(t *testing.T) {
		time := 7
		distance := 9

		expected := 4
		actual := numberOfWins(time, distance)
		assert.Equal(t, expected, actual)
	})

	t.Run("numberOfWins15", func(t *testing.T) {
		time := 15
		distance := 40

		expected := 8
		actual := numberOfWins(time, distance)
		assert.Equal(t, expected, actual)
	})

	t.Run("numberOfWins30", func(t *testing.T) {
		time := 30
		distance := 200

		expected := 9
		actual := numberOfWins(time, distance)
		assert.Equal(t, expected, actual)
	})

	t.Run("test", func(t *testing.T) {
		input := "Time:      7  15   30"

		expected := 71530
		actual := extractNumber(input)
		assert.Equal(t, expected, actual)
	})

}
