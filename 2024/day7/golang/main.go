package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file_data, _ := os.ReadFile("../input.txt")
	fmt.Printf("Part 1: %d\n", ParsePart1(string(file_data)))
	fmt.Printf("Part 2: %d\n", ParsePart2(string(file_data)))
}

func ParsePart1(input string) int {
	expressions := parseInput(input)
	sum := 0

	for _, expr := range expressions {
		if isValidExpression(expr, false) {
			sum += expr.Answer
		}
	}

	return sum
}

func ParsePart2(input string) int {
	expressions := parseInput(input)
	sum := 0

	for _, expr := range expressions {
		if isValidExpression(expr, true) {
			sum += expr.Answer
		}
	}

	return sum
}

func isValidExpression(expr Expression, concat bool) bool {
	return tryOperations(&expr.Numbers, 0, expr.Numbers[0], expr.Answer, concat)
}

func tryOperations(numbers *[]int, currentIndex int, currentResult int, target int, concat bool) bool {
	if currentIndex == len(*numbers)-1 {
		return currentResult == target
	}

	nextIndex := currentIndex + 1
	nextNum := (*numbers)[nextIndex]

	if tryOperations(numbers, nextIndex, currentResult+nextNum, target, concat) {
		return true
	}

	if tryOperations(numbers, nextIndex, currentResult*nextNum, target, concat) {
		return true
	}

	if concat && tryOperations(numbers, nextIndex, concatNumbers(currentResult, nextNum), target, concat) {
		return true
	}

	return false
}

func concatNumbers(numberA int, numberB int) int {
	if numberB == 0 {
		return numberA * 10
	}
	digits := int(math.Floor(math.Log10(float64(numberB)))) + 1
	return numberA*int(math.Pow10(digits)) + numberB
}

type Expression struct {
	Answer  int
	Numbers []int
}

func parseInput(input string) []Expression {
	lines := strings.Split(input, "\n")
	expressions := make([]Expression, 0)

	for _, line := range lines {
		if line == "" {
			continue
		}

		expression := Expression{}

		sides := strings.Split(line, ":")
		answer, _ := strconv.Atoi(sides[0])
		expression.Answer = answer

		numbers := strings.Split(strings.TrimSpace(sides[1]), " ")

		for _, number := range numbers {
			num, _ := strconv.Atoi(number)
			expression.Numbers = append(expression.Numbers, num)
		}

		expressions = append(expressions, expression)
	}

	return expressions
}
