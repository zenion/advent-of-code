package main

import (
	"fmt"
	"strings"

	aoc "github.com/zenion/advent-of-code"
)

func main() {
	fileLines := aoc.ReadFileLines("input.txt")
	calibration := 0
	for _, line := range fileLines {
		tokens := strings.Split(line, ": ")
		target := aoc.AtoiNoError(tokens[0])
		numbers := aoc.ToIntSlice(strings.Split(tokens[1], " "))
		// fmt.Println(target, numbers, couldBeTrue(target, numbers))
		if couldBeTrue(target, numbers) {
			calibration += target
		}
	}
	fmt.Println("Calibration:", calibration)
}

func couldBeTrue(target int, numbers []int) bool {
	// use bits in an int to represent whether operators are + or *
	operatorsMax := 1 << (len(numbers) - 1)

	for ops := 0; ops <= operatorsMax; ops++ {
		result := numbers[0]
		for i := 0; i < len(numbers)-1; i++ {
			bitMask := 1 << i
			if ops&bitMask > 0 {
				result = result * numbers[i+1]
			} else {
				result = result + numbers[i+1]
			}
		}
		if result == target {
			return true
		}
	}

	return false
}
