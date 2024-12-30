package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/zenion/advent-of-code"
)

func main() {
	fileLines := aoc.ReadFileLines("input.txt")

	stones := aoc.MapFunc(strings.Split(fileLines[0], " "), aoc.AtoiNoError)

	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}

	fmt.Println(len(stones))
}

func blink(stones []int) []int {
	result := make([]int, 0, len(stones))

	for _, s := range stones {
		if s == 0 {
			result = append(result, 1)
		} else if len(strconv.Itoa(s))%2 == 0 {
			str := strconv.Itoa(s)
			result = append(result, aoc.AtoiNoError(str[0:len(str)/2]))
			result = append(result, aoc.AtoiNoError(str[len(str)/2:]))
		} else {
			result = append(result, s*2024)
		}
	}

	return result
}
