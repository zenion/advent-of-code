package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file_data, _ := os.ReadFile("../input.txt")

	outputPart1 := ParsePart1(string(file_data))
	fmt.Println(outputPart1)

	outputPart2 := ParsePart2(string(file_data))
	fmt.Println(outputPart2)
}

func makeMatrix(input string) [][]string {
	scanner := bufio.NewScanner(strings.NewReader(input))

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	matrix := make([][]string, len(lines))

	for i, line := range lines {
		matrix[i] = strings.Split(line, "")
	}

	return matrix
}

func ParsePart1(input string) int {
	matrix := makeMatrix(input)

	count := 0

	// hacky directions index offsets
	directions := [][2]int{
		{0, 1},   // right
		{1, 0},   // down
		{1, 1},   // diagonal down-right
		{1, -1},  // diagonal down-left
		{0, -1},  // left
		{-1, 0},  // up
		{-1, 1},  // diagonal up-right
		{-1, -1}, // diagonal up-left
	}

	for lineIndex := range matrix {
		for charIndex := range matrix[lineIndex] {
			for _, dir := range directions {
				word := ""

				// check all 4 characters in the direction
				for i := 0; i < 4; i++ {
					newLineIndex := lineIndex + dir[0]*i
					newCharIndex := charIndex + dir[1]*i

					// check bounds so no go kaboom
					if newLineIndex < 0 || newLineIndex >= len(matrix) || newCharIndex < 0 || newCharIndex >= len(matrix[lineIndex]) {
						break
					}

					word += matrix[newLineIndex][newCharIndex]
				}

				if word == "XMAS" {
					count++
				}
			}
		}
	}

	return count
}

func ParsePart2(input string) int {
	matrix := makeMatrix(input)
	count := 0

	// i dont think i need to check all possible combinations but fuck it
	diagonalPairs := [][][2]int{
		{{1, 1}, {1, -1}},   // down-right and down-left
		{{-1, 1}, {-1, -1}}, // up-right and up-left
		{{1, 1}, {-1, 1}},   // down-right and up-right
		{{1, -1}, {-1, -1}}, // down-left and up-left
		{{1, 1}, {-1, -1}},  // down-right and up-left
		{{1, -1}, {-1, 1}},  // down-left and up-right
	}

	for lineIndex := range matrix {
		for charIndex := range matrix[lineIndex] {
			if matrix[lineIndex][charIndex] != "A" {
				continue
			}

			for _, pair := range diagonalPairs {
				dir1, dir2 := pair[0], pair[1]

				// Check first diagonal
				word1 := checkDiagonal(matrix, lineIndex, charIndex, dir1)
				// Check second diagonal
				word2 := checkDiagonal(matrix, lineIndex, charIndex, dir2)

				if word1 == "MAS" && word2 == "MAS" {
					count++
					break
				}
			}
		}
	}

	return count
}

func checkDiagonal(matrix [][]string, startLine, startChar int, dir [2]int) string {
	word := "A" // Start with center A

	// Check backwards for M (-1 step)
	newLineM := startLine - dir[0]
	newCharM := startChar - dir[1]
	if newLineM >= 0 && newLineM < len(matrix) &&
		newCharM >= 0 && newCharM < len(matrix[0]) {
		word = matrix[newLineM][newCharM] + word
	}

	// Check forward for S (+1 step)
	newLineS := startLine + dir[0]
	newCharS := startChar + dir[1]
	if newLineS >= 0 && newLineS < len(matrix) &&
		newCharS >= 0 && newCharS < len(matrix[0]) {
		word = word + matrix[newLineS][newCharS]
	}

	return word
}
