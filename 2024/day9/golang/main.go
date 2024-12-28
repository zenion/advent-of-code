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
	compactBlocks(blocks)
	compactBlocksChecksum := calcChecksum(blocks)

	blocks2 := calcBlocks(diskMap)
	compactFiles(blocks2)
	compactFilesChecksum := calcChecksum(blocks2)

	fmt.Println("Compact Blocks Checksum:", compactBlocksChecksum)
	fmt.Println("Compact Files Checksum:", compactFilesChecksum)
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

func compactBlocks(blocks []int) {
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

func findHighestId(blocks []int) int {
	for i := len(blocks) - 1; i >= 0; i-- {
		if blocks[i] >= 0 {
			return blocks[i]
		}
	}
	return -1
}

func compactFiles(blocks []int) {
	currentId := findHighestId(blocks)
	right := len(blocks) - 1

	for currentId > 0 {
		currentLen := 0

		// look backwards from the right to find the length of the current block
		for blocks[right] != currentId {
			right--
		}
		for blocks[right] == currentId {
			currentLen++
			right--
		}

		left := 0
		freeLen := 0

		// look forwards from the left to find the first free space that can fit the block
		for left <= right {
			if blocks[left] >= 0 {
				freeLen = 0
				left++
			} else {
				freeLen++
				left++
			}
			if freeLen == currentLen {
				break
			}
		}

		if freeLen == currentLen {
			// move the block into the free space
			for i := 0; i < currentLen; i++ {
				blocks[left-i-1] = currentId
				blocks[right+i+1] = -1
			}
		}

		currentId--
	}
}
