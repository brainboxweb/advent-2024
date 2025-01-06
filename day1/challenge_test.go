package day1_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day1"
	"github.com/brainboxweb/advent-2024/helpers"

	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"../data/day1_test.txt",
			11,
		},
		{
			"../data/day1.txt",
			1197984,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day1.ChallengeOne(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDay1Part2(t *testing.T) {

	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"../data/day1_test.txt",
			31,
		},
		{
			"../data/day1.txt",
			23387399,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day1.ChallengeTwo(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}
