package main

func isSafe(report []int) bool {
	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(report); i++ {
		distance := report[i] - report[i-1]

		if distance < -3 || distance > 3 || distance == 0 {
			return false
		}

		if distance > 0 {
			isDecreasing = false
		}
		if distance < 0 {
			isIncreasing = false
		}
	}

	return isIncreasing || isDecreasing
}

func isSafeWithOneRemoval(report []int) bool {
	for i := 0; i < len(report); i++ {
		modifiedLevels := make([]int, 0, len(report)-1)
		modifiedLevels = append(modifiedLevels, report[:i]...)
		modifiedLevels = append(modifiedLevels, report[i+1:]...)

		if isSafe(modifiedLevels) {
			return true
		}
	}
	return false
}
