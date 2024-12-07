package main

import (
	"fmt"
	"strings"

	aoc "github.com/zenion/advent-of-code"
)

type rule struct {
	before, after int
}

type update struct {
	pages   []int
	indexes map[int]int
}

func main() {
	fileLines := aoc.ReadFileLines("input.txt")

	rules := make([]rule, 0)
	updates := make([]update, 0)

	for _, line := range fileLines {
		if strings.Contains(line, "|") {
			beforeAfter := aoc.ToIntSlice(strings.Split(line, "|"))
			rules = append(rules, rule{beforeAfter[0], beforeAfter[1]})
		} else if line == "" {
			continue
		} else {
			pages := aoc.ToIntSlice(strings.Split(line, ","))
			indexes := make(map[int]int)
			for i, page := range pages {
				indexes[page] = i
			}
			updates = append(updates, update{pages, indexes})
		}
	}

	middleSum := 0

	invalidUpdates := make([]update, 0)
	for _, update := range updates {
		if isUpdateValid(update, rules) {
			middleSum += update.pages[len(update.pages)/2]
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	repairedMiddleSum := 0

	for _, update := range invalidUpdates {
		brutePermute(update, rules)
		// fmt.Println("Repaired update", repairedUpdate)
		repairedMiddleSum += update.pages[len(update.pages)/2]
	}

	fmt.Println("Middle Sum:", middleSum)
	fmt.Println("Repaired Middle Sum:", repairedMiddleSum)
}

func isUpdateValid(u update, rules []rule) bool {
	return findBrokenRule(u, rules) == nil
}

func findBrokenRule(u update, rules []rule) *rule {
	for _, rule := range rules {
		beforeIndex, beforeOk := u.indexes[rule.before]
		if !beforeOk {
			continue
		}

		afterIndex, afterOk := u.indexes[rule.after]
		if !afterOk {
			continue
		}

		if beforeIndex > afterIndex {
			return &rule
		}
	}
	return nil
}

func brutePermute(u update, rules []rule) {
	for rule := findBrokenRule(u, rules); rule != nil; rule = findBrokenRule(u, rules) {
		// swap the two elements that violate the rule
		// fmt.Println("Rule", fmt.Sprintf("%d|%d", rule.before, rule.after), "broken by", u.pages)
		beforeIndex := u.indexes[rule.before]
		afterIndex := u.indexes[rule.after]
		u.pages[beforeIndex] = rule.after
		u.pages[afterIndex] = rule.before
		u.indexes[rule.before] = afterIndex
		u.indexes[rule.after] = beforeIndex
	}
}
