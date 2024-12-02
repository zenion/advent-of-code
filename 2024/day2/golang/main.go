package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file_data, _ := os.ReadFile("./input.txt")
	safeReportsCount := 0

	for _, line := range strings.Split(string(file_data), "\n") {
		if len(line) == 0 {
			continue
		}
		var reports []int
		chars := strings.Fields(string(line))
		for _, char := range chars {
			if num, err := strconv.Atoi(char); err == nil {
				reports = append(reports, num)
			}
		}

		// Check if sequence is already safe
		if isValidSequence(reports) {
			safeReportsCount++
			fmt.Printf("%v: true\n", line)
			continue
		}

		// Try removing each number to see if it makes the sequence valid
		canBeMadeSafe := false
		for i := 0; i < len(reports); i++ {
			tempReports := make([]int, 0)
			tempReports = append(tempReports, reports[:i]...)
			tempReports = append(tempReports, reports[i+1:]...)

			if isValidSequence(tempReports) {
				canBeMadeSafe = true
				break
			}
		}

		if canBeMadeSafe {
			safeReportsCount++
		}
		fmt.Printf("%v: %v\n", line, canBeMadeSafe)
	}

	fmt.Printf("Safe reports count: %v\n", safeReportsCount)
}

func isValidSequence(reports []int) bool {
	if len(reports) <= 1 {
		return false
	}

	isIncrementing := true
	isDecrementing := true
	for i := 0; i < len(reports)-1; i++ {
		diff := reports[i+1] - reports[i]
		if diff == 0 || diff > 3 || diff < -3 {
			isIncrementing = false
			isDecrementing = false
		} else if diff < 0 {
			isIncrementing = false
		} else if diff > 0 {
			isDecrementing = false
		}
	}
	return isIncrementing || isDecrementing
}
