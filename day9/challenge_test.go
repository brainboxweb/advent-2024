package day9_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day9"
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
			"day9_test.txt",
			1928,
		},
		{
			"day9.txt",
			6201130364722,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			data := helpers.GetDataString(testDataPath + tt.dataFile)
			result := day9.ChallengeOne(data[0])
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
			"day9_test.txt",
			2858,
		},
		{
			"day9.txt",
			6221662795602,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			data := helpers.GetDataString(testDataPath + tt.dataFile)
			result := day9.ChallengeTwo(data[0])
			assert.Equal(t, tt.expected, result)
		})
	}
}
