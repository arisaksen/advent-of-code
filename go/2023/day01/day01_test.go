package main

import (
	utils "aoc/golang"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	exampleLines := utils.ReadLines("example1.txt")

	t.Run("getCalibrations", func(t *testing.T) {
		expected := []int{12, 38, 15, 77}
		actual := getCalibrations(exampleLines)
		assert.Equal(t, expected, actual)
	})

	t.Run("sumUpCalibrations", func(t *testing.T) {
		expected := 142
		actual := sumUpCalibrations(getCalibrations(exampleLines))
		assert.Equal(t, expected, actual)
	})

}

func TestPart2(t *testing.T) {
	exampleLines := "two1ninetwo"

	t.Run("getNumbersFromText", func(t *testing.T) {
		expected := "2two19nine2two"
		actual := addNumbersBeforeWords(exampleLines)
		assert.Equal(t, expected, actual)
	})

	exampleLines2 := utils.ReadLines("example2.txt")
	var calibrationsEdited []string
	for _, line := range exampleLines2 {
		calibrationsEdited = append(calibrationsEdited, addNumbersBeforeWords(line))
	}
	t.Run("getCalibrationsPart2", func(t *testing.T) {
		expected := []int{29, 83, 13, 24, 42, 14, 76}
		actual := getCalibrations(calibrationsEdited)
		assert.Equal(t, expected, actual)
	})

	t.Run("sumUpCalibrations", func(t *testing.T) {
		expected := 281
		actual := sumUpCalibrations(getCalibrations(calibrationsEdited))
		assert.Equal(t, expected, actual)
	})

}
