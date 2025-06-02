// Package day20 is for Advent of Code Day 20
package day20

import (
	"strings"

	"github.com/brainboxweb/advent-2024/day20/track"
)

// ChallengeOne is part one of today's challenge
func ChallengeOne(data []string, minCheats int) int {
	trackData := parse(data)
	w := track.New(trackData)
	return w.Go(minCheats)
}

func parse(input []string) [][]string {
	theMap := [][]string{}
	for _, line := range input {
		lineParts := strings.Split(line, "")
		theMap = append(theMap, lineParts)
	}
	return theMap
}
