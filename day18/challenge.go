// Package day18 is for Advent of Code Day 18
package day18

import (
	"strconv"
	"strings"

	"github.com/brainboxweb/advent-2024/day18/computer"
)

// ChallengeOne is part one of today's challenge
func ChallengeOne(data []string, length, corruptLimit int) int {
	input := parse(data)
	comp := computer.New(length)
	comp.AddCorrupted(input, corruptLimit)
	return comp.Walk()
}

// @TODO move to challenege
func parse(input []string) [][]int {
	ret := [][]int{} // make???
	for _, line := range input {
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			panic("expected")
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			panic("expected")
		}
		ln := []int{x, y}
		ret = append(ret, ln)
	}
	return ret
}
