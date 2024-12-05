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

	orderingRulesRaw := strings.Split(parts[0], "\n")

	orderingRules := [][]int{}
	for _, line := range orderingRulesRaw {
		nums := strings.Split(line, "|")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		orderingRules = append(orderingRules, []int{
			num1, num2,
		})
	}

	documentsRaw := strings.Split(parts[1], "\n")
	documentsRaw = documentsRaw[:len(documentsRaw)-1] // remove last empty line

	documents := [][]int{}
	for _, line := range documentsRaw {
		numsRaw := strings.Split(line, ",")
		nums := []int{}
		for _, numRaw := range numsRaw {
			num, _ := strconv.Atoi(numRaw)
			nums = append(nums, num)
		}
		documents = append(documents, nums)
	}

	incorrectDocuments := [][]int{}

	middles := []int{}
	for _, document := range documents {
		fmt.Printf("document: %v\n", document)
		valid := true
		for i, pageNum := range document {
			fmt.Printf("document: %v, pageNum: %d\n", document, pageNum)
			for _, rule := range orderingRules {
				if pageNum == rule[0] {
					fmt.Printf("document: %v, pageNum: %d, rule: %v\n", document, pageNum, rule)
					// check to see if rule[1] appears backward from our current position of i and if so the document is invalid
					for j := i - 1; j >= 0; j-- {
						if document[j] == rule[1] {
							fmt.Printf("document: %v, pageNum: %d, rule[1]: %d found at position %d - INVALID!", document, pageNum, rule[1], j)
							valid = false
							incorrectDocuments = append(incorrectDocuments, document)
							break
						}
					}
				} else if pageNum == rule[1] {
					fmt.Printf("document: %v, pageNum: %d, rule: %v\n", document, pageNum, rule)
					// check to see if rule[0] appears forward from our current position of i and if so the document is invalid
					for j := i + 1; j < len(document); j++ {
						if document[j] == rule[0] {
							fmt.Printf("document: %v, pageNum: %d, rule[0]: %d found at position %d - INVALID!", document, pageNum, rule[0], j)
							valid = false
							incorrectDocuments = append(incorrectDocuments, document)
							break
						}
					}
				}
			}
			if !valid {
				fmt.Printf("document %v is INVALID\n", document)
				break
			}
		}
		if valid {
			middles = append(middles, getMiddlePage(document))
			fmt.Printf("document %v is VALID WOOHOO!!\n", document)
		}
	}

	fmt.Printf("middles: %v\n", middles)

	middleSum := 0
	for _, middle := range middles {
		middleSum += middle
	}

	return middleSum, incorrectDocuments, orderingRules
}

func ParsePart2(incorrectDocuments [][]int, orderingRules [][]int) int {
	reorderedMiddlesSum := 0

	for _, document := range incorrectDocuments {
		reorderedDocument := reorderDocument(document, orderingRules)
		middlePage := getMiddlePage(reorderedDocument)
		reorderedMiddlesSum += middlePage
	}

	return reorderedMiddlesSum
}

func reorderDocument(document []int, orderingRules [][]int) []int {
	// Convert rules to a more usable format
	// where graph[a][b] means a must come before b
	graph := make(map[int][]int)
	inDegree := make(map[int]int)

	// Initialize graph and in-degree count
	for _, rule := range orderingRules {
		graph[rule[0]] = append(graph[rule[0]], rule[1])
		inDegree[rule[1]]++
	}

	// Initialize queue with nodes having zero in-degree
	queue := []int{}
	for _, page := range document {
		if inDegree[page] == 0 {
			queue = append(queue, page)
		}
	}

	// Perform topological sort
	sorted := []int{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		sorted = append(sorted, node)

		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// If sorted length is not equal to document length, there was a cycle
	if len(sorted) != len(document) {
		fmt.Println("Cycle detected, cannot reorder document")
		return document
	}

	return sorted
}

func getMiddlePage(document []int) int {
	return document[len(document)/2]
}
