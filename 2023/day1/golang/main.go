package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")

	// strip empty lines
	lines = slices.DeleteFunc(lines, func(line string) bool {
		return line == ""
	})

	// part 1
	sumPart1 := 0
	for _, line := range lines {
		var digits []string

		// Check each position in the string for digits
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				digits = append(digits, string(line[i]))
			}
		}

		if len(digits) == 0 {
			continue
		}

		// Take first and last digit
		first := digits[0]
		last := digits[len(digits)-1]

		num, err := strconv.Atoi(first + last)
		if err != nil {
			panic(err)
		}
		sumPart1 += num
	}

	// part 2
	stringToNumber := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	sumPart2 := 0
	for _, line := range lines {
		var digits []string

		// Check each position in the string for either a digit or a word
		for i := 0; i < len(line); i++ {
			// Check for numeric digit
			if line[i] >= '0' && line[i] <= '9' {
				digits = append(digits, string(line[i]))
				continue
			}

			// Check for word numbers
			for word, num := range stringToNumber {
				if strings.HasPrefix(line[i:], word) {
					digits = append(digits, num)
				}
			}
		}

		if len(digits) == 0 {
			continue
		}

		// Take first and last digit
		first := digits[0]
		last := digits[len(digits)-1]

		num, err := strconv.Atoi(first + last)
		if err != nil {
			panic(err)
		}
		sumPart2 += num
	}

	fmt.Println(sumPart1)
	fmt.Println(sumPart2)
}
