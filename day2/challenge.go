// Package day2 is AOC Day 2
package day2

import (
	"strconv"
	"strings"
)

// ChallengeOne is part 1 of the challenge
func ChallengeOne(data []string) int {
	reports := parse(data)
	count := 0
	for _, report := range reports {
		if isSafe(report) {
			count++
		}
	}

	return count
}

//revive:disable:cognitive-complexity

// ChallengeTwo is part 2 of the challenge
func ChallengeTwo(data []string) int {
	reports := parse(data)
	count := 0
OUTER:
	for _, report := range reports {
		if isSafe(report) {
			count++
			continue
		}
		for i := range report {
			rep := dropElement(report, i)
			if isSafe(rep) {
				count++
				continue OUTER
			}
		}
	}

	return count
}

//revive:enable:cognitive-complexity

func dropElement(report []int, element int) []int {
	ret := []int{}
	for i := range report {
		if i == element {
			continue
		}
		ret = append(ret, report[i])
	}
	return ret
}

func parse(data []string) [][]int {
	ret := [][]int{}
	for _, line := range data {
		rep := []int{}
		parts := strings.Split(line, " ")
		for _, i := range parts {
			p1, _ := strconv.Atoi(i)
			rep = append(rep, p1)
		}
		ret = append(ret, rep)
	}

	return ret
}

func isSafe(report []int) bool {
	if !areSmallSteps(report) {
		return false
	}

	return isSameDirection(report)
}

func areSmallSteps(report []int) bool {
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if diff > 3 {
			return false
		}
		if diff < -3 {
			return false
		}
	}

	return true
}

//revive:disable:cognitive-complexity
func isSameDirection(report []int) bool {
	direction := 0
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if diff == 0 {
			return false
		}
		if diff > 0 {
			if direction == -1 {
				return false
			}
			direction = 1
		}
		if diff < 0 {
			if direction == 1 {
				return false
			}
			direction = -1
		}
	}

	return true
}

//revive:enable:cognitive-complexity
