package day11_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day11"
	"github.com/stretchr/testify/assert"
)

func TestChallenge(t *testing.T) {
	tests := []struct {
		name       string
		data       string
		blinkCount int
		expected   int
	}{
		{
			"Part One",
			"17639 47 3858 0 470624 9467423 5 188",
			25,
			203228,
		},
		{
			"Part Two",
			"17639 47 3858 0 470624 9467423 5 188",
			75,
			240884656550923,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := day11.Challenge(tt.data, tt.blinkCount)
			assert.Equal(t, tt.expected, result)
		})
	}
}
