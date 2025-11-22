package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/zenion/advent-of-code"
)

type seqKey struct {
	initVal int
	blinks  int
}

type seq struct {
	stones []seqKey
	key    seqKey
	length int
}

func (s *seq) recalcLength(seqLens map[seqKey]int) {
	s.length = 0
	for _, sk := range s.stones {
		len, ok := seqLens[sk]
		if !ok {
			len = 1
		} else {
		}
		s.length += len
	}
}

func (s *seq) blink(seqVals map[int]*seq, seqLens map[seqKey]int) {
	nextStones := make([]seqKey, 0, len(s.stones))

	for _, stone := range s.stones {
		// is the next iteration of the sequence already known
		nextKey := seqKey{stone.initVal, stone.blinks + 1}
		_, ok := seqLens[nextKey]
		if ok {
			// use the pre-existing result and move on to the next stone
			nextStones = append(nextStones, nextKey)
			continue
		} else if s.key.initVal >= 0 {
			sv, ok := seqVals[stone.initVal]
			if ok && stone.initVal != s.key.initVal {
				sv.blink(seqVals, seqLens)
				nextStones = append(nextStones, nextKey)
			} else {
				// apply the rules to the stone
				if stone.initVal == 0 {
					nextStones = append(nextStones, seqKey{1, 0})
					addSeqIfMissing(1, seqVals)
				} else if len(strconv.Itoa(stone.initVal))%2 == 0 {
					str := strconv.Itoa(stone.initVal)
					leftVal := aoc.AtoiNoError(str[0 : len(str)/2])
					rightVal := aoc.AtoiNoError(str[len(str)/2:])

					nextStones = append(nextStones, seqKey{leftVal, 0})
					addSeqIfMissing(leftVal, seqVals)

					nextStones = append(nextStones, seqKey{rightVal, 0})
					addSeqIfMissing(rightVal, seqVals)
				} else {
					nextStones = append(nextStones, seqKey{stone.initVal * 2024, 0})
					addSeqIfMissing(stone.initVal*2024, seqVals)
				}
			}
		} else {
			sv, ok := seqVals[stone.initVal]
			if ok {
				sv.blink(seqVals, seqLens)
				nextStones = append(nextStones, nextKey)
			} else {
				panic(fmt.Sprint("Missing seq for ", stone.initVal))
			}
		}
	}

	s.stones = nextStones
	s.key.blinks++
	s.recalcLength(seqLens)
	if s.key.initVal >= 0 {
		seqLens[s.key] = s.length
	}
}

func addSeqIfMissing(initVal int, seqVals map[int]*seq) {
	_, ok := seqVals[initVal]
	if !ok {
		seqVals[initVal] = &seq{[]seqKey{{initVal, 0}}, seqKey{initVal, 0}, 1}
	}
}

func main() {
	fileLines := aoc.ReadFileLines("input.txt")

	stones := aoc.MapFunc(aoc.MapFunc(strings.Split(fileLines[0], " "), aoc.AtoiNoError), intToSeqKey)

	seqVals := make(map[int]*seq)
	seqLens := make(map[seqKey]int)

	for _, s := range stones {
		addSeqIfMissing(s.initVal, seqVals)
	}

	inputSeq := seq{stones, seqKey{-1, 0}, len(stones)}

	for i := range 75 {
		fmt.Println(i, ":", inputSeq.length)
		inputSeq.blink(seqVals, seqLens)
	}
	fmt.Println(inputSeq.length)
}

func intToSeqKey(i int) seqKey {
	return seqKey{i, 0}
}
