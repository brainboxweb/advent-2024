package day8

import (
	"github.com/brainboxweb/advent-2024/day8/roof"
)

func ChallengeOne(data [][]string) int {
	grid := roof.NewGrid(data)
	antinodes := grid.Antinodes()

	return len(antinodes)
}

func ChallengeTwo(data [][]string) int {
	grid := roof.NewGrid(data)
	antinodes := grid.AntinodesComprehensive()

	return len(antinodes)
}
