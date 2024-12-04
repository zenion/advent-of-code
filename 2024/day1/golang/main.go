package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readNumberPairs(filePath string) ([]int, []int, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, err
	}

	lines := strings.Split(string(content), "\n")

	// Remove any empty lines
	lines = slices.DeleteFunc(lines, func(line string) bool {
		return line == ""
	})

	var left, right []int
	for _, line := range lines {
		nums := strings.Split(line, "   ")

		numA, numB := 0, 0
		if numA, err = strconv.Atoi(nums[0]); err != nil {
			return nil, nil, err
		}
		if numB, err = strconv.Atoi(nums[1]); err != nil {
			return nil, nil, err
		}

		left = append(left, numA)
		right = append(right, numB)
	}

	return left, right, nil
}

func calculateAbsoluteDifferences(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}
	return sum
}

func calculateMatchingProducts(left, right []int) int {
	sum := 0
	for i := 0; i < len(left); i++ {
		countInRight := 0
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				countInRight++
			}
		}
		sum += left[i] * countInRight
	}
	return sum
}

func main() {
	left, right, err := readNumberPairs("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1 := calculateAbsoluteDifferences(left, right)
	part2 := calculateMatchingProducts(left, right)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
