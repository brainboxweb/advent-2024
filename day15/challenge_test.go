package day15_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day15"
	"github.com/brainboxweb/advent-2024/helpers"
	"github.com/stretchr/testify/assert"
)

func TestChallenge(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"../testdata/day15_test.txt",
			2028,
		},
		{
			"../testdata/day15_test2.txt",
			10092,
		},
		{
			"../testdata/day15.txt",
			1479679,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			data := helpers.GetDataString(tt.dataFile)
			result := day15.ChallengeOne(data)
			assert.Equal(t, tt.expected, result)
		})
	}
}
