package main

import (
	"fmt"

	aoc "github.com/zenion/advent-of-code"
)

type position struct {
	x, y int
}

func main() {
	fileLines := aoc.ReadFileLines("input.txt")

	antennas, puzzle := parsePuzzle(fileLines)

	partOneAntinodes := findAntinodes(antennas, puzzle, false)
	partTwoAntinodes := findAntinodes(antennas, puzzle, true)

	for k := range partTwoAntinodes {
		puzzle[k.y][k.x] = '#'
	}
	for _, line := range puzzle {
		fmt.Println(string(line))
	}
	fmt.Println("Antinodes:", len(partOneAntinodes))
	fmt.Println("More Antinodes:", len(partTwoAntinodes))
}

func findAntinodes(antennas map[byte][]position, puzzle [][]byte, harmonics bool) map[position]bool {
	// make a set of positions of antinodes
	antinodes := make(map[position]bool)
	for k := range antennas {
		// fmt.Println(string(k), antennas[k])
		positions := antennas[k]
		// test every pair of positions
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				p1 := positions[i]
				p2 := positions[j]
				// the difference represents the direction from p1 to p2
				xDiff := p2.x - p1.x
				yDiff := p2.y - p1.y

				if harmonics {
					// keep adding antinodes until out of bounds
					for i := 0; inBounds(position{p1.x + i*xDiff, p1.y + i*yDiff}, puzzle); i++ {
						antinodes[position{p1.x + i*xDiff, p1.y + i*yDiff}] = true
					}
					for i := 1; inBounds(position{p1.x - i*xDiff, p1.y - i*yDiff}, puzzle); i++ {
						antinodes[position{p1.x - i*xDiff, p1.y - i*yDiff}] = true
					}
				} else {
					// calculate antinodes
					an1 := position{p1.x - xDiff, p1.y - yDiff}
					an2 := position{p2.x + xDiff, p2.y + yDiff}
					if inBounds(an1, puzzle) {
						antinodes[an1] = true
					}
					if inBounds(an2, puzzle) {
						antinodes[an2] = true
					}
				}
			}
		}
	}
	return antinodes
}

func inBounds(p position, puzzle [][]byte) bool {
	return p.x >= 0 && p.x < len(puzzle[0]) && p.y >= 0 && p.y < len(puzzle)
}

func parsePuzzle(fileLines []string) (map[byte][]position, [][]byte) {
	antennas := make(map[byte][]position)
	puzzle := make([][]byte, len(fileLines))
	for y := 0; y < len(fileLines); y++ {
		puzzle[y] = make([]byte, len(fileLines[y]))
		for x := 0; x < len(fileLines[y]); x++ {
			puzzle[y][x] = fileLines[y][x]
			if string(fileLines[y][x]) != "." {
				positions, ok := antennas[fileLines[y][x]]
				if !ok {
					positions = make([]position, 0)
				}
				positions = append(positions, position{x, y})
				antennas[fileLines[y][x]] = positions
			}
		}
	}
	return antennas, puzzle
}
