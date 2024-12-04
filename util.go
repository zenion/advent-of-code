package aoc

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFileLines(filePath string) []string {
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines
}

func MapFunc[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func AtoiNoError(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func ToIntSlice(strings []string) []int {
	return MapFunc(strings, AtoiNoError)
}
