package day4_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day4"
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
			"day4_test.txt",
			18,
		},
		{
			"day4.txt",
			2454,
		},
	}

	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(testDataPath + tt.dataFile)
			result := day4.ChallengeOne(dataSet)
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
			"day4_test.txt",
			9,
		},
		{
			"day4.txt",
			1858,
		},
	}

	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(testDataPath + tt.dataFile)
			result := day4.ChallengeTwo(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}
