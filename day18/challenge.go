// Package day18 is for Advent of Code Day 18
package day18

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brainboxweb/advent-2024/day18/computer"
)

// ChallengeOne is part one of today's challenge
func ChallengeOne(data []string, length, corruptLimit int) int {
	input := parse(data)
	comp := computer.New(length)
	for i, line := range input {
		if i >= corruptLimit {
			break
		}
		comp.AddCorruptedByte(line)
	}
	_, res := comp.Walk()
	return res
}

// ChallengeTwo is part two of today's challenge
func ChallengeTwo(data []string, length int) string {
	input := parse(data)
	bin := NewBinarySearch(len(data))
	goBigger := true
	var index int
	for {
		if goBigger {
			index = bin.Bigger()
		} else {
			index = bin.Smaller()
		}
		bigger, result := tryTwo(length, index, input)
		if result != "" {
			return result
		}
		goBigger = bigger
	}
}

func tryTwo(length, index int, input [][]int) (bool, string) {
	comp := computer.New(length)
	for i, line := range input {
		if i > index {
			break
		}
		comp.AddCorruptedByte(line)
	}
	ok, _ := comp.Walk()
	if !ok { // too big
		return false, ""
	}
	// try the next one
	comp.AddCorruptedByte(input[index+1])
	ok, _ = comp.Walk() // run again
	if !ok {            // From pass to fail: winner!
		return true, fmt.Sprintf("%d,%d", input[index+1][0], input[index+1][1])
	}
	return true, "" // too small
}

func parse(input []string) [][]int {
	ret := [][]int{} // make???
	for _, line := range input {
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			panic("not expected")
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			panic("not expected")
		}
		ln := []int{x, y}
		ret = append(ret, ln)
	}
	return ret
}

// NewBinarySearch returns a new binary search helper
func NewBinarySearch(count int) *Binary {
	data := make([]int, count)
	for i := range count {
		data[i] = i
	}
	return &Binary{
		data:         data,
		hi:           len(data) - 1,
		lo:           0,
		currentIndex: 0,
	}
}

// Binary is a binary search helper
type Binary struct {
	data         []int
	hi           int
	lo           int
	currentIndex int
}

// Bigger returns the next (bigger) search candidate
func (b *Binary) Bigger() int {
	if b.currentIndex == 0 { // getting started
		b.currentIndex = (b.hi - b.lo) / 2
		return b.currentIndex
	}
	b.lo = b.currentIndex
	b.currentIndex = b.lo + ((b.hi - b.lo) / 2)
	return b.currentIndex
}

// Smaller returns the next (smaller) search candidate
func (b *Binary) Smaller() int {
	b.hi = b.currentIndex
	b.currentIndex = b.lo + ((b.hi - b.lo) / 2)
	return b.currentIndex
}
