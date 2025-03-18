package stones_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day11/stones"
	"github.com/stretchr/testify/assert"
)

func TestChallenge(t *testing.T) {
	tests := []struct {
		data       string
		blinkCount int
		expected   int
	}{
		{
			"0 1 10 99 999",
			1,
			7,
		},
		{
			"125 17",
			2,
			4,
		},
		{
			"125 17",
			5,
			13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.data, func(t *testing.T) {
			ss := stones.NewStoneSet(tt.data)
			result := ss.Blink(tt.blinkCount)

			assert.Equal(t, tt.expected, result)
		})
	}
}
