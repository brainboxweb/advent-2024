// Package day6 is AOC Day 6
package day6

import (
	"github.com/brainboxweb/advent-2024/day6/location"
)

// ChallengeOne is Part 1
func ChallengeOne(data [][]string) int {
	theMap := location.NewMap(data)
	res := theMap.Move()

	return res
}

// ChallengeTwo is Part 2
func ChallengeTwo(data [][]string) int {
	counter := 0
	for y := range len(data[0]) {
		for x := range data {
			altData := makeVariation(data, x, y)
			theMap := location.NewMap(altData)
			isLoop := theMap.IsEndlessLoop()
			if isLoop {
				counter++
			}
		}
	}

	return counter
}

func makeVariation(data [][]string, targetX, targetY int) [][]string {
	limit := len(data) // assume square
	altData := [][]string{}
	for _, line := range data {
		contents := make([]string, limit)
		copy(contents, line)
		altData = append(altData, contents)
	}
	if altData[targetX][targetY] != "^" {
		altData[targetX][targetY] = "#" // new obstacle
	}

	return altData
}
