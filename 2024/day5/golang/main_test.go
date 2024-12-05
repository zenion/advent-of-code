package main

import "testing"

var input = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`

var incorrectDocuments [][]int
var orderingRules [][]int

func TestParsePart1(t *testing.T) {
	want := 143
	got, iDocs, oRules := ParsePart1(input)

	incorrectDocuments = iDocs
	orderingRules = oRules

	if got != want {
		t.Errorf("Part 1: got %d want %d", got, want)
	}
}

func TestParsePart2(t *testing.T) {
	want := 123
	got := ParsePart2(incorrectDocuments, orderingRules)

	if got != want {
		t.Errorf("Part 2: got %d want %d", got, want)
	}
}
