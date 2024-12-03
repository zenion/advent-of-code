package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file_data, _ := os.ReadFile("input.txt")

	mulRegex := regexp.MustCompile(`^mul\((\d+),(\d+)\)`)
	doStr := "do()"
	dontStr := "don't()"

	sum := 0

	enabled := true

	for _, line := range strings.Split(string(file_data), "\n") {
		fmt.Println("Processing line:", line)

		i := 0
		for i < len(line) {
			fmt.Printf("i: %d\n", i)
			if i+len(doStr) <= len(line) && strings.Contains(line[i:i+len(doStr)], doStr) {
				enabled = true
				i += len(doStr)
			} else if i+len(dontStr) <= len(line) && strings.Contains(line[i:i+len(dontStr)], dontStr) {
				enabled = false
				i += len(dontStr)
			} else if enabled {
				matches := mulRegex.FindStringSubmatch(line[i:])
				if len(matches) > 0 {
					num1, _ := strconv.Atoi(matches[1])
					num2, _ := strconv.Atoi(matches[2])
					sum += num1 * num2
					i += len(matches[0])
				} else {
					i++
				}
			} else {
				i++
			}
		}
	}

	fmt.Println("Final sum:", sum)
}
