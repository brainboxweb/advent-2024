package computer_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day18/computer"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		length       int
		corruptLimit int
		data         [][]int
		expected     int
	}{
		{
			7,
			12,
			[][]int{
				{5, 4},
				{4, 2},
				{4, 5},
				{3, 0},
				{2, 1},
				{6, 3},
				{2, 4},
				{1, 5},
				{0, 6},
				{3, 3},
				{2, 6},
				{5, 1},
				{1, 2},
				{5, 5},
				{2, 5},
				{6, 5},
				{1, 4},
				{0, 4},
				{6, 4},
				{1, 1},
				{6, 1},
				{1, 0},
				{0, 5},
				{1, 6},
				{2, 0},
			},
			22,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			comp := computer.New(tt.length)
			for i, item := range tt.data {
				if i > tt.corruptLimit-1 {
					break
				}
				comp.AddCorruptedByte(item)
			}
			_, result := comp.Walk()
			assert.Equal(t, tt.expected, result)
		})
	}
}
