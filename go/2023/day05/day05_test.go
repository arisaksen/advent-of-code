package main

import (
	utils "aoc/golang"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		expected := 35
		actual := part1("example.txt")

		assert.Equal(t, expected, actual)
	})

	t.Run("Part 2", func(t *testing.T) {
		expected := 46
		actual := part2("example.txt")

		assert.Equal(t, expected, actual)
	})

	t.Run("test", func(t *testing.T) {
		input := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
		result := make([]string, 0, len(input)/2)
		for i := 1; i < len(input); i += 2 {
			result = append(result, input[i-1]+input[i])
		}
		fmt.Println(result)
	})

}

func TestMaps(t *testing.T) {

	t.Run("getSeedsFromMap", func(t *testing.T) {
		exampleInput := utils.ReadText("example.txt")
		categoryMap := createMap(exampleInput)

		expected := [][]int{{79, 14, 55, 13}}
		actual := categoryMap["seeds"]
		assert.Equal(t, expected, actual)

		expectedStF := [][]int{{0, 15, 37}, {37, 52, 2}, {39, 0, 15}}
		actualStf := categoryMap["soil-to-fertilizer"]
		assert.Equal(t, expectedStF, actualStf)
	})

}

func TestMapping(t *testing.T) {
	exampleInput := utils.ReadText("example.txt")
	categoryMap := createMap(exampleInput)

	t.Run("seedToSoil", func(t *testing.T) {
		seedsMap := categoryMap["seeds"][0]
		seedToSoilMap := categoryMap["seed-to-soil"]
		expected := []int{81, 14, 57, 13}
		actual := destinationToSource(seedsMap, seedToSoilMap)
		assert.Equal(t, expected, actual)
	})

	t.Run("soilToFertilizer", func(t *testing.T) {
		soil := []int{81, 14, 57, 13}
		soilToFertMap := categoryMap["soil-to-fertilizer"]
		expected := []int{81, 53, 57, 52}
		actual := destinationToSource(soil, soilToFertMap)
		assert.Equal(t, expected, actual)
	})

	t.Run("fertToWater", func(t *testing.T) {
		fert := []int{81, 53, 57, 52}
		fertToWater := categoryMap["fertilizer-to-water"]
		expected := []int{81, 49, 53, 41}
		actual := destinationToSource(fert, fertToWater)
		assert.Equal(t, expected, actual)
	})

	t.Run("waterToLight", func(t *testing.T) {
		water := []int{81, 49, 53, 41}
		waterToLight := categoryMap["water-to-light"]
		expected := []int{74, 42, 46, 34}
		actual := destinationToSource(water, waterToLight)
		assert.Equal(t, expected, actual)
	})

	t.Run("lightToTemp", func(t *testing.T) {
		light := []int{74, 42, 46, 34}
		lightToTemp := categoryMap["light-to-temperature"]
		expected := []int{78, 42, 82, 34}
		actual := destinationToSource(light, lightToTemp)
		assert.Equal(t, expected, actual)
	})

	t.Run("tempToHumidity", func(t *testing.T) {
		temp := []int{78, 42, 82, 34}
		tempToHumidity := categoryMap["temperature-to-humidity"]
		expected := []int{78, 43, 82, 35}
		actual := destinationToSource(temp, tempToHumidity)
		assert.Equal(t, expected, actual)
	})

	t.Run("humidityToLocation", func(t *testing.T) {
		humidity := []int{78, 43, 82, 35}
		humidityToLocation := categoryMap["humidity-to-location"]
		expected := []int{82, 43, 86, 35}
		actual := destinationToSource(humidity, humidityToLocation)
		assert.Equal(t, expected, actual)
	})

}
