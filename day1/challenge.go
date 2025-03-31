package day1

import (
	"sort"
	"strconv"
	"strings"
)

func ChallengeOne(data []string) int {
	one, two := parse(data)
	sort.Ints(one)
	sort.Ints(two)
	diff := 0
	for i := 0; i < len(one); i++ {
		d := one[i] - two[i]
		if d < 0 {
			d *= -1
		}
		diff += d
	}

	return diff
}

func parse(data []string) (one, two []int) {
	for _, line := range data {
		parts := strings.Split(line, "   ")
		p1, _ := strconv.Atoi(parts[0])
		p2, _ := strconv.Atoi(parts[1])
		one = append(one, p1)
		two = append(two, p2)
	}

	return one, two
}

func ChallengeTwo(data []string) int {
	one, two := parseTwo(data)
	ret := 0
	for _, item := range one {
		mult := item * two[item]
		ret += mult
	}

	return ret
}

func parseTwo(data []string) ([]int, map[int]int) {
	one := []int{}
	two := make(map[int]int)
	for _, line := range data {
		parts := strings.Split(line, "   ")
		p1, _ := strconv.Atoi(parts[0])
		p2, _ := strconv.Atoi(parts[1])
		one = append(one, p1)
		two[p2]++
	}

	return one, two
}
