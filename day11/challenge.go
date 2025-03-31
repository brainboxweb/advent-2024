package day11

import (
	"github.com/brainboxweb/advent-2024/day11/stones"
)

func Challenge(data string, blinkCount int) int {
	stoneSet := stones.NewStoneSet(data)

	return stoneSet.Blinks(blinkCount)
}
