package main

import (
	_ "embed"
	"testing"
)

//go:embed puzzle1_test.txt
var puzzleTest string

func TestPart1Input(t *testing.T) {
	result := part1(puzzleTest)
	expected := 3749
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}

func TestPart2Input(t *testing.T) {
	result := part2(puzzleTest)
	expected := 11387
	if result != expected {
		t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
	}
}
