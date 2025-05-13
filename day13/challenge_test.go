package day13_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day13"
	"github.com/brainboxweb/advent-2024/helpers"
	"github.com/stretchr/testify/assert"
)

func TestChallenge(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"../testdata/day13_test.txt",
			480,
		},
		{
			"../testdata/day13.txt",
			39748,
		},
	}
	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			data := helpers.GetDataString(tt.dataFile)
			result := day13.ChallengeOne(data)
			assert.Equal(t, tt.expected, result)
		})
	}
}
