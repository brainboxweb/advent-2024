package day14_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day14"
	"github.com/brainboxweb/advent-2024/helpers"
	"github.com/stretchr/testify/assert"
)

func TestChallenge(t *testing.T) {
	tests := []struct {
		dataFile  string
		maxX      int
		maxY      int
		stepCount int
		expected  int
	}{
		{
			"../testdata/day14_test.txt",
			11,
			7,
			100,
			12,
		},
		{
			"../testdata/day14.txt",
			101,
			103,
			100,
			221655456,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			data := helpers.GetDataString(tt.dataFile)
			result := day14.ChallengeOne(data, tt.maxX, tt.maxY, tt.stepCount)
			assert.Equal(t, tt.expected, result)
		})
	}
}
