package main

import "testing"

var input = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX
`

func TestParsePart1(t *testing.T) {
	want := 18
	got := ParsePart1(input)

	if got != want {
		t.Errorf("Part 1: got %d want %d", got, want)
	}
}

func TestParsePart2(t *testing.T) {
	want := 9
	got := ParsePart2(input)

	if got != want {
		t.Errorf("Part 2: got %d want %d", got, want)
	}
}
