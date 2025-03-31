package day9

import (
	"strconv"

	"github.com/brainboxweb/advent-2024/day9/disk"
)

func ChallengeOne(data string) int {
	code := toIntSlice(data)
	theDisk := disk.NewBasicDisk(code)
	theDisk.Compress()

	return theDisk.Checksum()
}

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
