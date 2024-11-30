package main

// from https://github.com/jasontconnell/advent
import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"time"
)

const (
	myLocalhostFolder = "/workspace/learning/advent-of-code"
	goFolder          = "go"
)

func main() {
	date := time.Now()
	//date, err := time.Parse(time.DateOnly, "2024-12-07")
	//if err != nil {
	//	log.Fatal(err)
	//}

	year := strconv.Itoa(date.Year())
	day := fmt.Sprintf("day%02d", date.Day())
	myUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dirPath := filepath.Join(myUser.HomeDir, myLocalhostFolder, goFolder, year, day)
	log.Println("creating directory", dirPath)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	textFiles := []string{"puzzle1.txt", "puzzle1_test.txt", "puzzle2.txt", "puzzle2_test.txt"}
	for _, textFile := range textFiles {
		textFilePath := filepath.Join(dirPath, textFile)
		log.Println("creating file", textFilePath)
		if err = os.WriteFile(textFilePath, nil, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	mainGoFile := filepath.Join(dirPath, "main.go")
	log.Println("creating file", dirPath)
	if err = os.WriteFile(mainGoFile, []byte(`package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed puzzle1.txt
var puzzle1 string

//go:embed puzzle2.txt
var puzzle2 string

func part1(puzzle string) int {
	var inputLines []string
	for _, line := range strings.Split(strings.TrimSuffix(puzzle, "\n"), "\n") {
		inputLines = append(inputLines, line)
	}

	return 1
}

func part2(puzzle string) int {
	var inputLines []string
	for _, line := range strings.Split(strings.TrimSuffix(puzzle, "\n"), "\n") {
		inputLines = append(inputLines, line)
	}

	return 1
}

func main() {
	start1 := time.Now()
	fmt.Println()
	fmt.Println("Part 1: ", part1(puzzle1))
	fmt.Println(time.Since(start1))

	start2 := time.Now()
	fmt.Println()
	fmt.Println("Part 2: ", part2(puzzle2))
	fmt.Println(time.Since(start2))
}
	`), os.ModePerm); err != nil {
		fmt.Printf("unable to write file: %v", err)
	}

	testPart1GoFile := filepath.Join(dirPath, "part1_test.go")
	log.Println("creating file", dirPath)
	if err = os.WriteFile(testPart1GoFile, []byte(`package main

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed puzzle1_test.txt
var puzzle1Test string

func TestPart1(t *testing.T) {
	actual := part1(puzzle1Test)

	assert.Equal(t, 1, actual)
}
	`), os.ModePerm); err != nil {
		fmt.Printf("unable to write file: %v", err)
	}

}
