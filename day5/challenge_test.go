package day5_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day5"
	"github.com/brainboxweb/advent-2024/helpers"
	"github.com/stretchr/testify/assert"
)

func TestChallenge(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"../data/day5_test.txt",
			143,
		},
		{
			"../data/day5.txt",
			5275,
		},
	}

	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day5.ChallengeOne(dataSet)
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
			"../data/day5_test.txt",
			123,
		},
		{
			"../data/day5.txt",
			6191,
		},
	}

	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day5.ChallengeTwo(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}
