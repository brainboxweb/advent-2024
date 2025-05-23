package day16_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day16"
	"github.com/brainboxweb/advent-2024/helpers"
	"github.com/stretchr/testify/assert"
)

func TestChallenge(t *testing.T) {
	tests := []struct {
		dataFile      string
		expectedCost  int // Part 1
		expectedCount int // Part 2
	}{
		{
			"../testdata/day16_test.txt",
			7036,
			45,
		},
		{
			"../testdata/day16_test2.txt",
			11048,
			64,
		},
		{
			"../testdata/day16.txt",
			123540,
			665,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			data := helpers.GetDataString(tt.dataFile)
			res := day16.ChallengeOne(data)
			assert.Equal(t, tt.expectedCost, res.MinCost)
			assert.Equal(t, tt.expectedCount, res.WinningPathCount)
		})
	}
}
