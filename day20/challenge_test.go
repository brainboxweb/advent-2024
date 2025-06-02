package day20_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day20"
	"github.com/brainboxweb/advent-2024/helpers"
	"github.com/stretchr/testify/assert"
)

func TestChallenge(t *testing.T) {
	tests := []struct {
		dataFile      string
		minimumSaving int
		expected      int
	}{
		{
			"../testdata/day20_test.txt",
			20,
			5,
		},
		{
			"../testdata/day20.txt",
			100,
			1459,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			data := helpers.GetDataString(tt.dataFile)
			result := day20.ChallengeOne(data, tt.minimumSaving)
			assert.Equal(t, tt.expected, result)
		})
	}
}
