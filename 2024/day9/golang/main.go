package main

import (
	"fmt"

	aoc "github.com/zenion/advent-of-code"
)

func main() {
	fileLines := aoc.ReadFileLines("input.txt")
	diskMap := aoc.MapFunc([]byte(fileLines[0]), func(b byte) int {
		return aoc.AtoiNoError(string(b))
	})

	blocks := calcBlocks(diskMap)

	compact(blocks)
	checksum := calcChecksum(blocks)

	fmt.Println(checksum)
}

func calcChecksum(blocks []int) int {
	checksum := 0
	for i := 0; i < len(blocks); i++ {
		if blocks[i] >= 0 {
			checksum += i * blocks[i]
		}
	}
	return checksum
}

func calcBlocks(diskMap []int) []int {
	blocks := make([]int, 0, 2*len(diskMap))
	id := 0
	empty := false
	for i := 0; i < len(diskMap); i++ {
		for j := 0; j < diskMap[i]; j++ {
			if empty {
				blocks = append(blocks, -1)
			} else {
				blocks = append(blocks, id)
			}
		}
		if !empty {
			id++
		}
		empty = !empty
	}
	return blocks
}

func compact(blocks []int) {
	left := 0
	right := len(blocks) - 1
	for left < right {
		if blocks[right] < 0 {
			right--
		} else if blocks[left] < 0 {
			blocks[left] = blocks[right]
			blocks[right] = -1
			left++
		} else {
			left++
		}
	}
}
