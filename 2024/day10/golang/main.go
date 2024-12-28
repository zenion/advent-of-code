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

	trailheads, puzzle := parsePuzzle(fileLines)

	totalScore := scoreTrailheads(trailheads, puzzle)
	fmt.Println("Total Score:", totalScore)
	totalRating := rateTrailheads(trailheads, puzzle)
	fmt.Println("Total Rating:", totalRating)
}

func parsePuzzle(fileLines []string) ([]position, [][]int) {
	positions := make([]position, 0)
	puzzle := make([][]int, len(fileLines))
	for y := 0; y < len(fileLines); y++ {
		puzzle[y] = make([]int, len(fileLines[y]))
		for x := 0; x < len(fileLines[y]); x++ {
			puzzle[y][x] = aoc.AtoiNoError(string(fileLines[y][x]))
			if string(fileLines[y][x]) == "0" {
				positions = append(positions, position{x, y})
			}
		}
	}
	return positions, puzzle
}

func scoreTrailheads(trailheads []position, puzzle [][]int) int {
	total := 0
	for _, p := range trailheads {
		visited := make(map[position]bool)
		score := scoreTrailhead(p.x, p.y, puzzle, visited, -1, 0, true)
		// fmt.Println(p, ":", score)
		total += score
	}
	return total
}

func rateTrailheads(trailheads []position, puzzle [][]int) int {
	total := 0
	for _, p := range trailheads {
		visited := make(map[position]bool)
		score := scoreTrailhead(p.x, p.y, puzzle, visited, -1, 0, false)
		// fmt.Println(p, ":", score)
		total += score
	}
	return total
}

func scoreTrailhead(x int, y int, puzzle [][]int, visited map[position]bool, currentHeight int, currentScore int, trackVisited bool) int {
	if x < 0 || y < 0 || y >= len(puzzle) || x >= len(puzzle[0]) || visited[position{x, y}] {
		// nothing
	} else if puzzle[y][x] == currentHeight+1 {
		if trackVisited {
			visited[position{x, y}] = true
		}
		// fmt.Println("{", x, y, "}", "val:", puzzle[y][x], "target:", currentHeight+1, "score:", currentScore)
		if currentHeight == 8 {
			return currentScore + 1
		} else {
			// check the four cardinal directions
			currentScore = scoreTrailhead(x-1, y, puzzle, visited, currentHeight+1, currentScore, trackVisited)
			currentScore = scoreTrailhead(x+1, y, puzzle, visited, currentHeight+1, currentScore, trackVisited)
			currentScore = scoreTrailhead(x, y-1, puzzle, visited, currentHeight+1, currentScore, trackVisited)
			currentScore = scoreTrailhead(x, y+1, puzzle, visited, currentHeight+1, currentScore, trackVisited)
		}
	}

	return currentScore
}
