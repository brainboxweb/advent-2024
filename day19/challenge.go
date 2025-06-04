// Package day19 is for Advent of Code Day 19
package day19

import (
	"strings"

	"github.com/brainboxweb/advent-2024/day19/branding"
)

// ChallengeOne is part one of today's challenge
func ChallengeOne(data []string) int {
	result := parse(data)
	counter := 0
	tt := branding.NewTowels(result.towels)
	for _, patt := range result.patterns {
		if tt.CanObtain(patt) > 0 {
			counter++
		}
	}
	return counter
}

// ChallengeTwo is part one of today's challenge
func ChallengeTwo(data []string) int {
	result := parse(data)

	counter := 0
	tt := branding.NewTowels(result.towels)
	for _, patt := range result.patterns {
		counter += tt.CanObtain(patt)
	}
	return counter
}

func parse(input []string) result {
	towelsPart := input[0]
	towels := strings.Split(towelsPart, ", ")
	patterns := input[2:]
	return result{towels, patterns}
}

type result struct {
	towels   []string
	patterns []string
}
