// Package day12 is AOC Day 12
package day12

import (
	"github.com/brainboxweb/advent-2024/day12/garden"
)

// ChallengeOne is part 1 of today's challenge
func ChallengeOne(data [][]string) int {
	plot := garden.NewPlot(data)
	return plot.TotalCost()
}
