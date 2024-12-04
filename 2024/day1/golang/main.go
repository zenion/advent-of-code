package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	aoc "github.com/zenion/advent-of-code"
)

func main() {
	fileLines := aoc.ReadFileLines("input.txt")

	var list1, list2 []int

	for _, line := range fileLines {
		tokens := strings.Split(line, "   ")
		if len(tokens) != 2 {
			println("Unexpected number of tokens in string:", line)
		}
		val1, _ := strconv.Atoi(tokens[0])
		val2, _ := strconv.Atoi(tokens[1])
		list1 = append(list1, val1)
		list2 = append(list2, val2)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	totalDist := calcTotalDist(list1, list2)
	similarity := calcSimilarity(list1, list2)

	fmt.Println("Distance:", totalDist)
	fmt.Println("Similarity:", similarity)
}

func calcSimilarity(list1, list2 []int) int {
	freqMap := make(map[int]int)
	for _, v2 := range list2 {
		freq, ok := freqMap[v2]
		if !ok {
			freq = 0
		}
		freq++
		freqMap[v2] = freq
	}

	similarity := 0
	for _, v1 := range list1 {
		freq, ok := freqMap[v1]
		if !ok {
			freq = 0
		}
		similarity += v1 * freq
	}

	return similarity
}

func calcTotalDist(list1, list2 []int) int {
	totalDist := 0
	for i, v1 := range list1 {
		v2 := list2[i]
		dist := v1 - v2
		if dist < 0 {
			dist = -dist
		}
		totalDist += dist
	}

	return totalDist
}
