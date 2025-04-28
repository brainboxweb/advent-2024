// Package day11 is AOC Day 11
package day11

import (
	"github.com/brainboxweb/advent-2024/day11/stones"
)

// Challenge is both Part 1 and Part 2
func Challenge(data string, blinkCount int) int {
	stoneSet := stones.NewStoneSet(data)

	return stoneSet.Blinks(blinkCount)
}
