package day12_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day12"
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
			"day12_test.txt",
			140,
		},
		{
			"day12_test1.txt",
			772,
		},
	}

	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			data := helpers.GetDataString(testDataPath + tt.dataFile)
			dataSet := helpers.ToXY(data)
			result := day12.ChallengeOne(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}
