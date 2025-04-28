// Package day10 is AOC Day 10
package day10

import "github.com/brainboxweb/advent-2024/day10/area"

// ChallengeOne is part 1
func ChallengeOne(data [][]string) int {
	terrain := area.New(data)
	terrain.FindStartingPoints()
	score := 0
	for _, start := range terrain.StartingPoints {
		steps := terrain.Walk(start)
		score += len(steps)
	}

	return score
}

// ChallengeTwo is part 2
func ChallengeTwo(data [][]string) int {
	terrain := area.New(data)
	terrain.FindStartingPoints()
	score := 0
	for _, start := range terrain.StartingPoints {
		stepScore := terrain.Walk(start)
		for _, val := range stepScore {
			score += val
		}
	}

	return score
}
