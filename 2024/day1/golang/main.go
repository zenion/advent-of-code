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

type Columns struct {
	left  []int
	right []int
}

func parseColumns(contents string) (Columns, error) {
	lines := strings.Split(contents, "\n")

	lines = slices.DeleteFunc(lines, func(line string) bool {
		return line == ""
	})

	var cols Columns
	for _, line := range lines {
		nums := strings.Fields(line)

		left, err := strconv.Atoi(nums[0])
		if err != nil {
			return Columns{}, err
		}
		right, err := strconv.Atoi(nums[1])
		if err != nil {
			return Columns{}, err
		}

		cols.left = append(cols.left, left)
		cols.right = append(cols.right, right)
	}

	slices.Sort(cols.left)
	slices.Sort(cols.right)

	return cols, nil
}

func solvePart1(contents string) (int, error) {
	cols, err := parseColumns(contents)
	if err != nil {
		return 0, err
	}

	sum := 0
	for i := 0; i < len(cols.left); i++ {
		sum += int(math.Abs(float64(cols.left[i] - cols.right[i])))
	}
	return sum, nil
}

func solvePart2(contents string) (int, error) {
	cols, err := parseColumns(contents)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, leftNum := range cols.left {
		countInRight := 0
		for _, rightNum := range cols.right {
			if leftNum == rightNum {
				countInRight++
			}
		}
		sum += leftNum * countInRight
	}
	return sum, nil
}

func main() {
	contents, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1, err := solvePart1(string(contents))
	if err != nil {
		log.Fatal(err)
	}
	part2, err := solvePart2(string(contents))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
