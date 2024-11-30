package main

import (
	utils "aoc/golang"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func part2(filename string) int {
	startPart2 := time.Now()
	puzzle2Input := utils.ReadLines(filename)
	fullContent := strings.Join(puzzle2Input, "\n")
	sections := strings.Split(fullContent, "\n\n")

	seedLine := strings.TrimPrefix(sections[0], "seeds: ")
	seedToSoilInput := strings.TrimPrefix(sections[1], "seed-to-soil map:\n")
	soilToFertilizerInput := strings.TrimPrefix(sections[2], "soil-to-fertilizer map:\n")
	fertilizerToWaterInput := strings.TrimPrefix(sections[3], "fertilizer-to-water map:\n")
	waterToLightInput := strings.TrimPrefix(sections[4], "water-to-light map:\n")
	lightToTemperatureInput := strings.TrimPrefix(sections[5], "light-to-temperature map:\n")
	temperatureToHumidityInput := strings.TrimPrefix(sections[6], "temperature-to-humidity map:\n")
	humidityToLocationInput := strings.TrimPrefix(sections[7], "humidity-to-location map:\n")

	seedToSoil := parseRangeMapping(seedToSoilInput)
	soilToFertilizer := parseRangeMapping(soilToFertilizerInput)
	fertilizerToWater := parseRangeMapping(fertilizerToWaterInput)
	waterToLight := parseRangeMapping(waterToLightInput)
	lightToTemperature := parseRangeMapping(lightToTemperatureInput)
	temperatureToHumidity := parseRangeMapping(temperatureToHumidityInput)
	humidityToLocation := parseRangeMapping(humidityToLocationInput)

	seedRanges := parseSeedRanges(seedLine)

	lowestLocation := math.MaxInt64
	for i, seedRange := range seedRanges {
		startTimeSeedRange := time.Now()
		start := seedRange[0]
		length := seedRange[1]
		for i := 0; i < length; i++ {
			seed := start + i
			soil := convertNumber(seed, seedToSoil)
			fertilizer := convertNumber(soil, soilToFertilizer)
			water := convertNumber(fertilizer, fertilizerToWater)
			light := convertNumber(water, waterToLight)
			temperature := convertNumber(light, lightToTemperature)
			humidity := convertNumber(temperature, temperatureToHumidity)
			location := convertNumber(humidity, humidityToLocation)
			if location < lowestLocation {
				lowestLocation = location
			}
		}
		duration := time.Since(startTimeSeedRange)
		fmt.Printf("Execution time seedRange %d: ", i)
		fmt.Println(duration)
	}

	duration := time.Since(startPart2)
	fmt.Print("Execution time whole part2: ")
	fmt.Println(duration)
	return lowestLocation
}

type RangeMapping struct {
	destStart   int
	sourceStart int
	length      int
}

func parseRangeMapping(input string) []RangeMapping {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var mappings []RangeMapping
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		destStart, _ := strconv.Atoi(parts[0])
		sourceStart, _ := strconv.Atoi(parts[1])
		length, _ := strconv.Atoi(parts[2])
		mappings = append(mappings, RangeMapping{destStart, sourceStart, length})
	}
	return mappings
}

func convertNumber(number int, mappings []RangeMapping) int {
	for _, mapping := range mappings {
		if number >= mapping.sourceStart && number < mapping.sourceStart+mapping.length {
			return mapping.destStart + (number - mapping.sourceStart)
		}
	}
	return number
}

func parseSeedRanges(seedLine string) [][2]int {
	start := time.Now()
	parts := strings.Fields(seedLine)
	var ranges [][2]int
	for i := 0; i < len(parts); i += 2 {
		start, _ := strconv.Atoi(parts[i])
		length, _ := strconv.Atoi(parts[i+1])
		ranges = append(ranges, [2]int{start, length})
	}

	duration := time.Since(start)
	fmt.Print("Execution time parseRangeMappings: ")
	fmt.Println(duration)
	return ranges
}
