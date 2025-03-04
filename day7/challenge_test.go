package day7_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day7"
	"github.com/brainboxweb/advent-2024/helpers"
	"github.com/stretchr/testify/assert"
)

const testDataPath = "../data/"

func TestChallenge(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"day7_test.txt",
			3749,
		},
		{
			"day7.txt",
			1985268524462,
		},
	}

	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(testDataPath + tt.dataFile)
			result := day7.ChallengeOne(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestChallengeDay2(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"day7_test.txt",
			11387,
		},
		{
			"day7.txt",
			150077710195188,
		},
	}

	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(testDataPath + tt.dataFile)
			result := day7.ChallengeTwo(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}
