package main

import (
	"regexp"
	"strings"

	aoc "github.com/zenion/advent-of-code"
)

func main() {
	fileLines := aoc.ReadFileLines("input.txt")

	pattern, _ := regexp.Compile(`mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\)`)

	sum := 0
	enabledSum := 0

	enabled := true
	for _, line := range fileLines {
		matches := pattern.FindAllString(line, -1)
		for _, match := range matches {
			if match == "do()" {
				enabled = true
			} else if match == "don't()" {
				enabled = false
			} else {
				// mul(nnn,nnn)
				ints := aoc.ToIntSlice(strings.Split(match[4:len(match)-1], ","))
				product := ints[0] * ints[1]
				sum += product
				if enabled {
					enabledSum += product
				}
			}
			// println(match)
		}
	}

	println("Sum:", sum)
	println("Enabled Sum:", enabledSum)
}
