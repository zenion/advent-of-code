package main

import "testing"

func TestSolutions(t *testing.T) {
	contents := `3   4
4   3
2   5
1   3
3   9
3   3
`
	part1, err := solvePart1(contents)
	if err != nil {
		t.Fatal(err)
	}
	if part1 != 11 {
		t.Errorf("Part 1: expected 11, got %d", part1)
	}

	part2, err := solvePart2(contents)
	if err != nil {
		t.Fatal(err)
	}
	if part2 != 31 {
		t.Errorf("Part 2: expected 31, got %d", part2)
	}
}
