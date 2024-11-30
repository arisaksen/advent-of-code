package main

import (
	utils "aoc/golang"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	part1()

	//part2()

}

func part1() {
	puzzleInput := utils.ReadLines("day02/puzzle.txt")
	possibleGamesSum := game(puzzleInput)

	fmt.Println("Part 1: ", possibleGamesSum)
}

func part2() {
	puzzleInput := utils.ReadLines("day02/puzzle.txt")
	fmt.Println("Part 1: ", puzzleInput)
}

func game(games []string) int {
	var possibleGames = 0

game:
	for index, game := range games {
		fmt.Println("NEW GAME - " + strconv.Itoa(index+1))
		gamePrefix := "Game " + strconv.Itoa(index+1) + ": "
		withoutGamePrefix := strings.TrimPrefix(game, gamePrefix)
		sets := strings.Split(withoutGamePrefix, ";")
		validGame := false
		for _, set := range sets {
			fmt.Println("NEW SET")
			gameMoves := strings.Split(set, ",")
			for _, move := range gameMoves {
				gameBag := bag{12, 13, 14}
				fmt.Printf("refill gamebag - red: %d, green: %d, blue: %d \n", gameBag.red, gameBag.green, gameBag.blue)
				fmt.Println("next move: " + move)
				switch {
				case strings.Contains(move, "blue"):
					number := getNumberFromString(move)
					gameBag.blue -= number
				case strings.Contains(move, "red"):
					number := getNumberFromString(move)
					gameBag.red -= number
				case strings.Contains(move, "green"):
					number := getNumberFromString(move)
					gameBag.green -= number
				default:
					panic("new cube color?" + move)
				}
				fmt.Printf("gamebag - red: %d, green: %d, blue: %d \n", gameBag.red, gameBag.green, gameBag.blue)
				if gameBag.green > 0 && gameBag.red > 0 && gameBag.blue > 0 {
					validGame = true
				} else {
					validGame = false
					fmt.Println("GAME UNVALID")
					fmt.Println("BREAK")
					fmt.Println()
					continue game
				}
			}
		}
		if validGame {
			fmt.Println("GAME VALID")
			possibleGames += index + 1
		} else {
		}
		println()

	}

	return possibleGames
}

func getNumberFromString(line string) int {
	re := regexp.MustCompile("[0-9]+")
	listOfAllNumbers := re.FindAllString(line, -1)
	firstNumber := listOfAllNumbers[0]
	numberAsInt, err := strconv.Atoi(firstNumber)
	if err != nil {
		panic("error")
	}
	return numberAsInt
}

type bag struct {
	red, green, blue int
}
