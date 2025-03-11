package day6_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day6"
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
			"day6_test.txt",
			41,
		},
		{
			"day6.txt",
			5516,
		},
	}

	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			data := helpers.GetDataString(testDataPath + tt.dataFile)
			dataSet := helpers.ToXY(data)
			result := day6.ChallengeOne(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestChallengeTwo(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"day6_test.txt",
			6,
		},
		{
			"day6.txt",
			2008,
		},
	}

	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			data := helpers.GetDataString(testDataPath + tt.dataFile)
			dataSet := helpers.ToXY(data)
			result := day6.ChallengeTwo(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}
