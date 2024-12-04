package main

import (
	"strings"

	aoc "github.com/zenion/advent-of-code"
)

func main() {
	fileLines := aoc.ReadFileLines("input.txt")

	safeReports := 0
	safeReportsWithDampener := 0

	for _, line := range fileLines {
		tokens := strings.Split(line, " ")
		levels := aoc.ToIntSlice(tokens)

		if isSafe(levels) {
			safeReports++
		}

		if isSafeWithDampener(levels) {
			safeReportsWithDampener++
		}
	}

	println("Safe reports:", safeReports)
	println("Safe reports with dampener:", safeReportsWithDampener)
}

func isSafe(levels []int) bool {
	allIncreasing := true
	allDecreasing := true

	for i := 0; i < len(levels)-1; i++ {
		if levels[i] > levels[i+1] {
			allIncreasing = false
			if levels[i]-levels[i+1] > 3 {
				// increase of more than 3, so not safe
				return false
			}
		} else if levels[i] < levels[i+1] {
			allDecreasing = false
			if levels[i+1]-levels[i] > 3 {
				// decrease of more than 3, so not safe
				return false
			}
		} else {
			// two adjacent values are the same so definitely not safe
			return false
		}
	}

	return allIncreasing || allDecreasing
}

func isSafeWithDampener(levels []int) bool {
	if isSafe(levels) {
		// if safe without manipulation, return true
		return true
	}

	// create a new slice with a single element removed and see if it is safe
	for i := 0; i < len(levels); i++ {
		dampenedLevels := removeLevelFromReport(levels, i)

		if isSafe(dampenedLevels) {
			// if safe with a single level removed then return true
			return true
		}
	}

	// no safe slice was found, so the report is not safe
	return false
}

func removeLevelFromReport(levels []int, index int) []int {
	result := make([]int, 0, len(levels)-1)
	result = append(result, levels[:index]...)
	result = append(result, levels[index+1:]...)
	return result
}
