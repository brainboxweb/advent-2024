// Package day8 is AOC Day 8
package day8

import (
	"github.com/brainboxweb/advent-2024/day8/roof"
)

// ChallengeOne is Part 1
func ChallengeOne(data [][]string) int {
	grid := roof.NewGrid(data)
	antinodes := grid.Antinodes()

	return len(antinodes)
}

// ChallengeTwo is Part 2
func ChallengeTwo(data [][]string) int {
	grid := roof.NewGrid(data)
	antinodes := grid.AntinodesComprehensive()

	return len(antinodes)
}
