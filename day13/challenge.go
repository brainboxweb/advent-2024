// Package day13 if for Advent of Code Day 13
package day13

import (
	"strconv"
	"strings"

	"github.com/brainboxweb/advent-2024/day13/island"
)

// ChallengeOne is part one
func ChallengeOne(data []string) int {
	ret := 0
	things := parse(data, 0)
	for _, thing := range things {
		cost := thing.Calculate(100)
		ret += cost
	}

	return ret
}

// func ChallengeTwo(data []string) int {
// 	ret := 0
// 	things := parse(data, 10000000000000)
// 	for _, thing := range things {
// 		cost := thing.Calculate(-1)
// 		ret += cost
// 	}

// 	return ret
// }

func parse(data []string, offset int) []island.Machine {
	things := []island.Machine{}
	t := island.Machine{} // new thing
	for _, line := range data {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		if parts[1] == "A:" {
			x := getXorY(parts[2], "+")
			y := getXorY(parts[3], "+")
			t.A = island.Point{X: x, Y: y}
		}
		if parts[1] == "B:" {
			x := getXorY(parts[2], "+")
			y := getXorY(parts[3], "+")
			t.B = island.Point{X: x, Y: y}
		}
		if parts[0] == "Prize:" {
			x := getXorY(parts[1], "=") + offset
			y := getXorY(parts[2], "=") + offset
			t.Prize = island.Point{X: x, Y: y}

			things = append(things, t)

			t = island.Machine{} // start again
		}
	}

	return things
}

func getXorY(in, delim string) int {
	ret := strings.Trim(in, ",")
	pp := strings.Split(ret, delim)
	val, err := strconv.Atoi(pp[1])
	if err != nil {
		panic("not expected")
	}

	return val
}
