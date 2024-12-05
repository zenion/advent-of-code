package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file_data, _ := os.ReadFile("input.txt")
	part1Sum, incorrectDocuments, orderingRules := ParsePart1(string(file_data))

	fmt.Printf("part1: %d\n", part1Sum)
	fmt.Printf("part2: %v\n", ParsePart2(incorrectDocuments, orderingRules))
}

func ParsePart1(input string) (int, [][]int, [][]int) {
	parts := strings.Split(input, "\n\n")
	orderingRules := [][]int{}
	documents := [][]int{}

	// parse ordering rules from the input
	orderingRulesRaw := strings.Split(parts[0], "\n")
	for _, line := range orderingRulesRaw {
		if line == "" {
			continue
		}
		nums := strings.Split(line, "|")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		orderingRules = append(orderingRules, []int{num1, num2})
	}

	// parse documents from the input
	documentsRaw := strings.Split(parts[1], "\n")
	for _, line := range documentsRaw {
		if line == "" {
			continue
		}
		numsRaw := strings.Split(line, ",")
		nums := []int{}
		for _, numRaw := range numsRaw {
			num, _ := strconv.Atoi(numRaw)
			nums = append(nums, num)
		}
		documents = append(documents, nums)
	}

	validMiddles := []int{}
	invalidDocuments := [][]int{} // we need this for part 2 it seems

	for _, document := range documents {
		if isValidDocument(document, orderingRules) {
			validMiddles = append(validMiddles, getMiddlePage(document))
		} else {
			invalidDocuments = append(invalidDocuments, document)
		}
	}

	middleSum := 0
	for _, middle := range validMiddles {
		middleSum += middle
	}

	return middleSum, invalidDocuments, orderingRules
}

func isValidDocument(document []int, rules [][]int) bool {
	for _, rule := range rules {
		before, after := rule[0], rule[1]
		beforeIdx := -1
		afterIdx := -1

		// Find the indices of both the 'before' and 'after' pages in the document
		for i, page := range document {
			if page == before {
				beforeIdx = i
			}
			if page == after {
				afterIdx = i
			}
		}

		// if both are found then we have a valid rule
		if beforeIdx != -1 && afterIdx != -1 {
			// if its backward its no good baby
			if beforeIdx > afterIdx {
				return false
			}
		}
	}

	// otherwise we are chillin'
	return true
}

func getMiddlePage(document []int) int {
	return document[len(document)/2]
}

func ParsePart2(incorrectDocuments [][]int, orderingRules [][]int) int {
	middleSum := 0

	for _, doc := range incorrectDocuments {
		orderedDoc := []int{}

		// fucking graphs man
		graph := make(map[int][]int)   // full graph map
		edgeCount := make(map[int]int) // map of page to number of incoming edges

		// Initialize all pages in both maps
		for _, page := range doc {
			graph[page] = []int{}
			edgeCount[page] = 0
		}

		// Build the graph with the rules
		for _, rule := range orderingRules {
			before, after := rule[0], rule[1]
			// Only evaluate rules where both pages of rule exist in the document
			if containsPage(doc, before) && containsPage(doc, after) {
				graph[before] = append(graph[before], after)
				edgeCount[after]++
			}
		}

		// make a poor mans queue from a int slice because golang is too basic
		queue := []int{}

		// Find pages with no incoming edges and queue em up
		for page := range graph {
			if edgeCount[page] == 0 {
				queue = append(queue, page)
			}
		}

		for len(queue) > 0 {
			// Take the first page from the "queue" we made and do "queue" things to "pop" what we took
			current := queue[0]
			queue = queue[1:]
			orderedDoc = append(orderedDoc, current)

			// find all pages that come after the current page from the graph
			for _, next := range graph[current] {
				// reduce our edging intensity
				edgeCount[next]--
				if edgeCount[next] == 0 {
					// we're done edging... let's add it to the queue
					queue = append(queue, next)
				}
			}
		}

		middleSum += getMiddlePage(orderedDoc)
	}

	return middleSum
}

func containsPage(doc []int, page int) bool {
	for _, p := range doc {
		if p == page {
			return true
		}
	}
	return false
}
