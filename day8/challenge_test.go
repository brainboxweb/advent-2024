package day8_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day8"
	"github.com/brainboxweb/advent-2024/helpers"
	"github.com/stretchr/testify/assert"
)

const testDataPath = "../testdata/"

func TestChallenge(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"day8_test.txt",
			2,
		},
		{
			"day8_test2.txt",
			4,
		},
		{
			"day8_test3.txt",
			14,
		},
		{
			"day8.txt",
			364,
		},
	}

	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			data := helpers.GetDataString(testDataPath + tt.dataFile)
			dataSet := helpers.ToXY(data)
			result := day8.ChallengeOne(dataSet)
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
			"day8_test4.txt",
			9,
		},
		{
			"day8.txt",
			1231,
		},
	}

	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			data := helpers.GetDataString(testDataPath + tt.dataFile)
			dataSet := helpers.ToXY(data)
			result := day8.ChallengeTwo(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}
