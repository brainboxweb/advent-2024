package day3

import (
	"regexp"
	"strconv"
)

func ChallengeOne(data string) int {
	rx, _ := regexp.Compile("(mul)\\((\\d+)\\,(\\d+)\\)")
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
	rx2, _ := regexp.Compile("(mul)\\((\\d+)\\,(\\d+)\\)|do\\(\\)|don't\\(\\)")
	matches := rx2.FindAllStringSubmatch(data, -1)
	ret := 0
	do := true
	for _, mul := range matches {
		if mul[0] == "do()" {
			do = true
		}
		if mul[0] == "don't()" {
			do = false
		}
		if do == false {
			continue
		}
		x, _ := strconv.Atoi(mul[2])
		y, _ := strconv.Atoi(mul[3])
		ret += x * y
	}

	return ret
}
