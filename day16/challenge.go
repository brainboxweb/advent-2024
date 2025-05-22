// Package day16 is for Advent of Code Day 16
package day16

import (
	"github.com/brainboxweb/advent-2024/day16/maze"
)

// ChallengeOne is part one of today's challenge
func ChallengeOne(data []string) int {
	mz := maze.New(data)
	return mz.Solve()
}
