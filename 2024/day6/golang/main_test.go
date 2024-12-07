package main

import "testing"

var input = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`

var incorrectDocuments [][]int
var orderingRules [][]int

func TestParsePart1(t *testing.T) {
	want := 41
	got := ParsePart1(input)

	if got != want {
		t.Errorf("Part 1: got %d want %d", got, want)
	}
}

func TestParsePart2(t *testing.T) {
	want := 6
	got := ParsePart2(input)

	if got != want {
		t.Errorf("Part 2: got %d want %d", got, want)
	}
}
