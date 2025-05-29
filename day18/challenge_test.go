package day18_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day18"
	"github.com/brainboxweb/advent-2024/helpers"
	"github.com/stretchr/testify/assert"
)

const testDataFile = "../testdata/"

func TestChallengeOne(t *testing.T) {
	tests := []struct {
		dataFile     string
		length       int
		corruptLimit int
		expected     int
	}{
		{
			"day18_test.txt",
			7,
			12,
			22,
		},
		{
			"day18.txt",
			71,
			1024,
			312,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			data := helpers.GetDataString(testDataFile + tt.dataFile)
			result := day18.ChallengeOne(data, tt.length, tt.corruptLimit)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestChallengeTwo(t *testing.T) {
	tests := []struct {
		dataFile string
		length   int
		expected string
	}{
		{
			"day18_test.txt",
			7,
			"6,1",
		},
		{
			"day18.txt",
			71,
			"28,26",
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			data := helpers.GetDataString(testDataFile + tt.dataFile)
			result := day18.ChallengeTwo(data, tt.length)
			assert.Equal(t, tt.expected, result)
		})
	}
}
