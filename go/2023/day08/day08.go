package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed puzzle.txt
var inputDay string

type choice struct {
	left  string
	right string
}

func Part1(input string) int {
	var inputList []string
	for _, line := range strings.Split(strings.TrimSuffix(input, "\n"), "\n") {
		inputList = append(inputList, line)
	}
	nodeMap := createMapTwo(inputList)

	nextDest := "AAA"
	instructions := inputList[0]
	var steps int
	for {
		for _, instruction := range instructions {
			//fmt.Printf("steps: %d, current: %s, instruction: %s, dest: %s\n", steps, nextDest, string(instruction), nodeMap[nextDest])
			if instruction == rune('L') {
				nextDest = nodeMap[nextDest].left
			} else if instruction == rune('R') {
				nextDest = nodeMap[nextDest].right
			} else {
				panic("Instruction is not L - Left or R - Right")
			}

			steps += 1
		}

		if nextDest == "ZZZ" {
			//fmt.Println("you have arrived")
			//fmt.Printf("steps: %d\n", steps)
			break
		}
	}

	return steps
}

func createMapTwo(inputList []string) map[string]choice {
	nodeMap := make(map[string]choice)
	for _, line := range inputList[2:] {
		//fmt.Println(line)
		k, l, r := splitInstructionTwo(line)
		nodeMap[k] = choice{
			left:  l,
			right: r,
		}
	}
	return nodeMap
}

func splitInstructionTwo(instruction string) (string, string, string) {
	instructionSplit := strings.Split(instruction, " ")
	destination := instructionSplit[0]
	leftDest := strings.TrimSuffix(strings.TrimPrefix(instructionSplit[2], "("), ",")
	rightDest := strings.TrimSuffix(instructionSplit[3], ")")

	return destination, leftDest, rightDest
}

func splitInstruction(instruction string) map[string]choice {
	instructionSplit := strings.Split(instruction, " ")
	destination := instructionSplit[0]
	leftDest := strings.TrimSuffix(strings.TrimPrefix(instructionSplit[2], "("), ",")
	rightDest := strings.TrimSuffix(instructionSplit[3], ")")

	return map[string]choice{
		destination: {
			left:  leftDest,
			right: rightDest,
		},
	}
}

func Part2(input string) int {

	return 1
}

func main() {
	fmt.Println("--2023 day 08 solution--")
	start := time.Now()
	fmt.Println("part1: ", Part1(inputDay))
	fmt.Println(time.Since(start))

	//start = time.Now()
	//fmt.Println("part2: ", Part2(inputDay))
	//fmt.Println(time.Since(start))
}
