package day19_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day19"
	"github.com/brainboxweb/advent-2024/helpers"
	"github.com/stretchr/testify/assert"
)

const testDataFile = "../testdata/"

func TestChallengeOne(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"day19_test.txt",
			6,
		},
		{
			"day19.txt",
			260,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			data := helpers.GetDataString(testDataFile + tt.dataFile)
			result := day19.ChallengeOne(data)
			assert.Equal(t, tt.expected, result)
		})
	}
}
