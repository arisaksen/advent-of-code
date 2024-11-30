package main

import (
	_ "embed"
	"testing"
)

//go:embed example.txt
var inputTest1 string

//go:embed example2.txt
var inputTest2 string

func TestUnits(t *testing.T) {

	t.Run("splitInstructionsTwo", func(t *testing.T) {
		stringInput := "AAA = (BBB, CCC)"
		expectedDest, expectedLeft, expectedRight := "AAA", "BBB", "CCC"
		actualDest, actualLeft, actualRight := splitInstructionTwo(stringInput)

		if actualDest != expectedDest {
			t.Errorf("Result is incorrect, got: %s, want: %s.", actualDest, expectedDest)
		}
		if actualLeft != expectedLeft {
			t.Errorf("Result is incorrect, got: %s, want: %s.", actualLeft, expectedLeft)
		}
		if actualLeft != expectedLeft {
			t.Errorf("Result is incorrect, got: %s, want: %s.", actualRight, expectedRight)
		}

	})

	t.Run("splitInstructions", func(t *testing.T) {
		stringInput := "AAA = (BBB, CCC)"
		expected := map[string]choice{
			"AAA": {
				left:  "BBB",
				right: "CCC",
			},
		}
		actual := splitInstruction(stringInput)

		if actual["AAA"] != expected["AAA"] {
			t.Errorf("Result is incorrect, got: %s, want: %s.", actual["AAA"], expected["AAA"])
		}
	})

}

func TestPart1(t *testing.T) {

	t.Run("example1", func(t *testing.T) {
		result := Part1(inputTest1)
		expected := 2
		if result != expected {
			t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
		}
	})

	t.Run("example2", func(t *testing.T) {
		result := Part1(inputTest2)
		expected := 6
		if result != expected {
			t.Errorf("Result is incorrect, got: %d, want: %d.", result, expected)
		}
	})

}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1(inputDay)
	}
}
func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2(inputDay)
	}
}
