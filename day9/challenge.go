// Package day9 is AOC Day 9
package day9

import (
	"strconv"

	"github.com/brainboxweb/advent-2024/day9/disk"
)

// ChallengeOne is Part 1
func ChallengeOne(data string) int {
	code := toIntSlice(data)
	theDisk := disk.NewBasicDisk(code)
	theDisk.Compress()

	return theDisk.Checksum()
}

// ChallengeTwo is Part 2
func ChallengeTwo(data string) int {
	code := toIntSlice(data)
	theDisk := disk.NewAdvancedDisk(code)
	theDisk.Compress()

	return theDisk.Checksum()
}

func toIntSlice(data string) []int {
	ret := []int{}
	for _, val := range data {
		numb, _ := strconv.Atoi(string(val))
		ret = append(ret, numb)
	}

	return ret
}
