package day10_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day10"
	"github.com/brainboxweb/advent-2024/helpers"
	"github.com/stretchr/testify/assert"
)

const testDataFile = "../data/"

func TestChallenge(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"day10_test.txt",
			1,
		},
		{
			"day10_test2.txt",
			2,
		},
		{
			"day10_test3.txt",
			4,
		},
		{
			"day10_test4.txt",
			36,
		},
		{
			"day10.txt",
			593,
		},
	}

	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			data := helpers.GetDataString(testDataFile + tt.dataFile)
			dataSet := helpers.ToXY(data)
			result := day10.ChallengeOne(dataSet)
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
			"day10_test5.txt",
			3,
		},
		{
			"day10_test4.txt",
			81,
		},
		{
			"day10.txt",
			1192,
		},
	}

	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			data := helpers.GetDataString(testDataFile + tt.dataFile)
			dataSet := helpers.ToXY(data)
			result := day10.ChallengeTwo(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}
