package day3_test

import (
	"strings"
	"testing"

	"github.com/brainboxweb/advent-2024/day3"
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
			"day3_test.txt",
			161,
		},
		{
			"day3.txt",
			161289189,
		},
	}

	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(testDataPath + tt.dataFile)
			data := strings.Join(dataSet, "") // Handle line breaks
			result := day3.ChallengeOne(data)
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
			"day3_test2.txt",
			48,
		},
		{
			"day3.txt",
			83595109,
		},
	}

	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(testDataPath + tt.dataFile)
			data := strings.Join(dataSet, "") // Handle line breaks
			result := day3.ChallengeTwo(data)
			assert.Equal(t, tt.expected, result)
		})
	}
}
