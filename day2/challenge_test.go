package day2_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day2"
	"github.com/brainboxweb/advent-2024/helpers"

	"github.com/stretchr/testify/assert"
)

const testDataPath = "../testdata/"

func TestDay2(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"day2_test.txt",
			2,
		},
		{
			"day2.txt",
			559,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(testDataPath + tt.dataFile)
			result := day2.ChallengeOne(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDay2Part2(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"day2_test.txt",
			4,
		},
		{
			"day2.txt",
			601,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(testDataPath + tt.dataFile)
			result := day2.ChallengeTwo(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}
