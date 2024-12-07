package main

import (
	"fmt"

	aoc "github.com/zenion/advent-of-code"
)

func main() {
	fileLines := aoc.ReadFileLines("input.txt")
	partOne(fileLines)
	partTwo(fileLines)
}

func partTwo(fileLines []string) {
	// the guard is in a loop if they enter a space they have visited before from the same direction as before
	// they might cross a space in multiple directions, so all dirs need to be recorded

	// loop over all possible positions for an obstruction
	obstructions := 0
	for i := 0; i < len(fileLines); i++ {
		for j := 0; j < len(fileLines[i]); j++ {
			x, y, puzzle := parsePuzzle(fileLines)
			if puzzle[i][j] != '#' && puzzle[i][j] != '^' {
				puzzle[i][j] = '#'
				if calcPatrol(x, y, puzzle) {
					// fmt.Println()
					// for _, line := range puzzle {
					// 	fmt.Println(string(line))
					// }
					obstructions++
				}
			}
		}
	}

	fmt.Println("Obstructions:", obstructions)
}

func calcPatrol(x int, y int, puzzle [][]byte) bool {
	dirs := [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
	puzzle[y][x] = '.'
	dirIndex := 0

	isLoop := false
	for {
		bitMask := byte(1 << dirIndex)
		if (puzzle[y][x]-'.')&bitMask > 0 {
			isLoop = true
			break
		}
		puzzle[y][x] = ((puzzle[y][x] - '.') | bitMask) + '.'

		dir := dirs[dirIndex]
		newX := x + dir[0]
		newY := y + dir[1]

		if newX < 0 || newX >= len(puzzle[y]) || newY < 0 || newY >= len(puzzle) {
			break
		}

		if puzzle[newY][newX] == '#' {
			dirIndex = (dirIndex + 1) % 4
		} else {
			x, y = newX, newY
		}
	}

	return isLoop
}

func partOne(fileLines []string) {
	x, y, puzzle := parsePuzzle(fileLines)

	calcPatrol(x, y, puzzle)

	sum := 0
	for _, line := range puzzle {
		// fmt.Println(string(line))
		for _, b := range line {
			if b != '.' && b != '#' {
				sum++
			}
		}
	}

	fmt.Println("Positions:", sum)
}

func parsePuzzle(fileLines []string) (int, int, [][]byte) {
	guardX, guardY := -1, -1
	puzzle := make([][]byte, len(fileLines))
	for y := 0; y < len(fileLines); y++ {
		puzzle[y] = make([]byte, len(fileLines[y]))
		for x := 0; x < len(fileLines[y]); x++ {
			puzzle[y][x] = fileLines[y][x]
			if string(fileLines[y][x]) == "^" {
				guardX, guardY = x, y
			}
		}
	}
	return guardX, guardY, puzzle
}
