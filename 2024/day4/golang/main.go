package main

import (
	"regexp"

	aoc "github.com/zenion/advent-of-code"
)

func main() {
	fileLines := aoc.ReadFileLines("input.txt")
	vertLines := genVerticals(fileLines)
	diagLines := genDiagonals(fileLines)

	sum := 0

	allLines := make([]string, 0)
	allLines = append(allLines, fileLines...)
	allLines = append(allLines, vertLines...)
	allLines = append(allLines, diagLines...)

	xmas, _ := regexp.Compile(`XMAS`)
	samx, _ := regexp.Compile(`SAMX`)

	for _, line := range allLines {
		// println(line)
		// println(len(xmas.FindAllString(line, -1)))
		// println(len(samx.FindAllString(line, -1)))
		// println()

		sum += len(xmas.FindAllString(line, -1))
		sum += len(samx.FindAllString(line, -1))
	}

	xSum := 0

	// scan through looking for A
	// upon finding an A, get the cross values to see if they are S,M or M,S and if so then add to sum
	for i := 0; i < len(fileLines[0]); i++ {
		for j := 0; j < len(fileLines); j++ {
			if (fileLines[i][j]) == 'A' && isX(fileLines, i, j) {
				xSum++
			}

		}
	}

	println("XMAS:", sum)
	println("X-MAS:", xSum)
}

func isX(lines []string, i int, j int) bool {
	// check for out of bounds
	if i-1 < 0 || j-1 < 0 || i+1 >= len(lines[0]) || j+1 >= len(lines) {
		return false
	}

	ul := lines[i-1][j-1]
	ur := lines[i+1][j-1]
	dl := lines[i-1][j+1]
	dr := lines[i+1][j+1]

	return ((ul == 'S' && dr == 'M') || (ul == 'M' && dr == 'S')) && ((dl == 'S' && ur == 'M') || (dl == 'M' && ur == 'S'))
}

func genVerticals(lines []string) []string {
	verticalLines := make([]string, 0)

	// assume all lines are the same length
	for i := 0; i < len(lines[0]); i++ {
		vLine := make([]byte, 0, len(lines))
		for _, line := range lines {
			vLine = append(vLine, line[i])
		}
		verticalLines = append(verticalLines, string(vLine))
	}

	return verticalLines
}

func genDiagonals(lines []string) []string {
	diagonalLines := make([]string, 0)

	// bottom to top
	for k := len(lines); k >= 0; k-- {
		dLine := make([]byte, 0, len(lines))
		for i, j := k, 0; i < len(lines) && j < len(lines[i]); i, j = i+1, j+1 {
			dLine = append(dLine, lines[i][j])
		}
		diagonalLines = append(diagonalLines, string(dLine))
	}

	// left to right
	for k := 1; k < len(lines[0]); k++ {
		dLine := make([]byte, 0, len(lines))
		for i, j := 0, k; i < len(lines) && j < len(lines[i]); i, j = i+1, j+1 {
			dLine = append(dLine, lines[i][j])
		}
		diagonalLines = append(diagonalLines, string(dLine))
	}

	// right to left
	for k := len(lines[0]); k >= 0; k-- {
		dLine := make([]byte, 0, len(lines))
		for i, j := k, len(lines)-1; i < len(lines[0]) && j >= 0; i, j = i+1, j-1 {
			dLine = append(dLine, lines[i][j])
		}
		diagonalLines = append(diagonalLines, string(dLine))
	}

	// bottom to top
	for k := len(lines) - 2; k >= 0; k-- {
		dLine := make([]byte, 0, len(lines))
		for i, j := 0, k; i < len(lines[0]) && j >= 0; i, j = i+1, j-1 {
			dLine = append(dLine, lines[i][j])
		}
		diagonalLines = append(diagonalLines, string(dLine))
	}

	return diagonalLines
}
