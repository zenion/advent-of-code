package main

import (
	"fmt"
	"strings"

	aoc "github.com/zenion/advent-of-code"
)

func main() {
	fileLines := aoc.ReadFileLines("input.txt")
	calibration := 0
	concatCalib := 0
	for _, line := range fileLines {
		tokens := strings.Split(line, ": ")
		target := aoc.AtoiNoError(tokens[0])
		numbers := aoc.ToIntSlice(strings.Split(tokens[1], " "))
		// fmt.Println(target, numbers, couldBeTrue(target, numbers))
		if couldBeTrue(target, numbers) {
			calibration += target
			concatCalib += target
		} else if couldBeTrueTwo(target, numbers) {
			concatCalib += target
		}
	}
	fmt.Println("Calibration:", calibration)
	fmt.Println("Calibration Concatenation:", concatCalib)
}

func couldBeTrueTwo(target int, numbers []int) bool {
	operators := make([]int, len(numbers)-1)

	for {
		result := numbers[0]
		for i := 0; i < len(numbers)-1; i++ {
			if operators[i] == 0 {
				result = result + numbers[i+1]
			} else if operators[i] == 1 {
				result = result * numbers[i+1]
			} else {
				result = aoc.AtoiNoError(fmt.Sprint(result) + fmt.Sprint(numbers[i+1]))
			}
		}
		if result == target {
			return true
		}

		operators[0]++
		for i := 0; i < len(operators); i++ {
			if operators[i] > 2 {
				if i+1 >= len(operators) {
					return false
				}
				operators[i] = 0
				operators[i+1]++
			}
		}
		// fmt.Println(operators)
	}
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
