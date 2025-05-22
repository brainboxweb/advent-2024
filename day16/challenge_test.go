package day16_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day16"
	"github.com/brainboxweb/advent-2024/helpers"
	"github.com/stretchr/testify/assert"
)

func TestChallenge(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"../testdata/day16_test.txt",
			7036,
		},
		{
			"../testdata/day16_test2.txt",
			11048,
		},
		{
			"../testdata/day16.txt",
			123540,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			data := helpers.GetDataString(tt.dataFile)
			result := day16.ChallengeOne(data)
			assert.Equal(t, tt.expected, result)
		})
	}
}
