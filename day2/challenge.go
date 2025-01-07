package day2

import (
	"strconv"
	"strings"
)

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

func ChallengeTwo(data []string) int {
	reports := parse(data)
	count := 0
OUTER:
	for _, report := range reports {
		if isSafe(report) {
			count++
			continue OUTER
		}
		for i := 0; i < len(report); i++ {
			rep := dropElement(report, i)
			if isSafe(rep) {	
				count++
				continue OUTER
			}
		}
	}

	return count
}

func dropElement(report []int, element int) []int {
	ret := []int{}
	for i := 0; i < len(report); i++ {
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
	if smallSteps(report) == false {
		return false
	}

	return sameDirection(report)
}

func smallSteps(report []int) bool {
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

func sameDirection(report []int) bool {
	direction := 0 // -1, 0, 1
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
