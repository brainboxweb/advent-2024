package day3

import (
	"regexp"
	"strconv"
)

var rx = regexp.MustCompile(`(mul)\((\d+)\,(\d+)\)`)
var rx2 = regexp.MustCompile(`(mul)\((\d+)\,(\d+)\)|do\(\)|don't\(\)`)

func ChallengeOne(data string) int {
	matches := rx.FindAllStringSubmatch(data, -1)
	ret := 0
	for _, mul := range matches {
		x, _ := strconv.Atoi(mul[2])
		y, _ := strconv.Atoi(mul[3])
		ret += x * y
	}

	return ret
}

func ChallengeTwo(data string) int {
	matches := rx2.FindAllStringSubmatch(data, -1)
	ret := 0
	do := true
	for _, mul := range matches {
		if mul[0] == "do()" {
			do = true
			continue
		}
		if mul[0] == "don't()" {
			do = false
			continue
		}
		if !do {
			continue
		}
		x, _ := strconv.Atoi(mul[2])
		y, _ := strconv.Atoi(mul[3])
		ret += x * y
	}

	return ret
}
