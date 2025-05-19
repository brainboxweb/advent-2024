// Package day15 is for Advent of Code Day 15
package day15

import (
	"strings"

	"github.com/brainboxweb/advent-2024/day15/warehouse"
)

// ChallengeOne is part one of today's challenge
func ChallengeOne(data []string) int {
	warehouseData, comms := parse(data)
	w := warehouse.New(warehouseData)
	for _, comm := range comms {
		w.Move(comm)
	}
	return w.GPS()
}

func parse(input []string) ([][]string, []string) {
	theMap := [][]string{}
	commandString := ""
	isCommandString := false
	for _, line := range input {
		if line == "" {
			isCommandString = true
			continue
		}
		if isCommandString {
			commandString += line
			continue
		}
		lineParts := strings.Split(line, "")
		theMap = append(theMap, lineParts)
	}
	comms := strings.Split(commandString, "")
	return theMap, comms
}
