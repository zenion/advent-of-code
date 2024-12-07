package main

import (
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func (d Direction) TurnRight() Direction {
	return Direction((int(d) + 1) % 4)
}

type Game struct {
	Board            [][]string
	CharPos          []int
	CharDir          Direction
	VisitedCount     int
	StuckInLoop      bool
	NumOfRepeatSteps int
	BaracadeLocation []int
}

func main() {
	file_data, _ := os.ReadFile("input.txt")
	fmt.Println(ParsePart1(string(file_data)))
	fmt.Println(ParsePart2(string(file_data)))
}

func ParsePart1(input string) int {
	game := NewGameFromInput(input)
	game.Play()

	return game.VisitedCount
}

func ParsePart2(input string) int {
	game := NewGameFromInput(input)

	loopCount := 0

	for i := range game.Board {
		for j := range game.Board[i] {
			if !(i == game.CharPos[0] && j == game.CharPos[1]) && game.Board[i][j] != "#" {
				game.BaracadeLocation = []int{i, j}
				game.Board[i][j] = "#"
				res := game.Play()
				if res == false {
					loopCount++
					game.Board[game.BaracadeLocation[0]][game.BaracadeLocation[1]] = "O"
					// game.PrintBoard()
				}
				game = NewGameFromInput(input)
			}
		}
	}

	game.Play()

	return loopCount
}

func (g *Game) Play() bool {
	for {
		nextPos := g.getNextPosition()

		// do we go off screen? if so we endgame
		if nextPos[0] < 0 || nextPos[0] >= len(g.Board) ||
			nextPos[1] < 0 || nextPos[1] >= len(g.Board[0]) {
			return true
		}

		// are we in a loop?
		if g.Board[nextPos[0]][nextPos[1]] == "X" {
			g.NumOfRepeatSteps++
			if g.NumOfRepeatSteps > 10000 {
				g.StuckInLoop = true
				return false
			}
		}

		// do we hit a wall? if so we turn right
		if g.Board[nextPos[0]][nextPos[1]] == "#" {
			g.CharDir = g.CharDir.TurnRight()
			g.Board[g.CharPos[0]][g.CharPos[1]] = setDirectionString(g.CharDir)
			continue
		}

		g.Board[g.CharPos[0]][g.CharPos[1]] = "X"
		g.CharPos = nextPos

		if g.Board[g.CharPos[0]][g.CharPos[1]] != "X" {
			g.VisitedCount++
		}
		g.Board[g.CharPos[0]][g.CharPos[1]] = setDirectionString(g.CharDir)

		// g.PrintBoard()
	}
}

func (g *Game) PrintBoard() {
	fmt.Println("")
	for _, row := range g.Board {
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println("")
}

func (g *Game) getNextPosition() []int {
	nextPos := make([]int, 2)
	copy(nextPos, g.CharPos)

	switch g.CharDir {
	case Up:
		nextPos[0]--
	case Right:
		nextPos[1]++
	case Down:
		nextPos[0]++
	case Left:
		nextPos[1]--
	}

	return nextPos
}

func NewGameFromInput(input string) Game {
	game := Game{
		VisitedCount: 1,
	}
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		if line == "" {
			continue
		}

		chars := strings.Split(line, "")
		for j, char := range chars {
			if char == "^" || char == "v" || char == "<" || char == ">" {
				game.CharPos = []int{i, j}
				game.CharDir = getDirection(char)
			}
		}

		game.Board = append(game.Board, chars)
	}
	return game
}

func getDirection(char string) Direction {
	switch char {
	case "^":
		return Up
	case "v":
		return Down
	case "<":
		return Left
	case ">":
		return Right
	}
	return 0
}

func setDirectionString(dir Direction) string {
	switch dir {
	case Up:
		return "^"
	case Right:
		return ">"
	case Down:
		return "v"
	case Left:
		return "<"
	}
	return ""
}
