package main

import (
	"fmt"

	aoc "github.com/zenion/advent-of-code"
)

func main() {
	fileLines := aoc.ReadFileLines("input.txt")

	for _, s := range fileLines {
		fmt.Println(s)
	}
}
