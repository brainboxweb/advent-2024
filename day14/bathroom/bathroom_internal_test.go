package bathroom

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveRobot(t *testing.T) {
	lenX := 11
	lenY := 7
	tests := []struct {
		stepCount int
		expectedX int
		expectedY int
	}{
		{
			1,
			4,
			1,
		},
		{
			2,
			6,
			5,
		},
		{
			3,
			8,
			2,
		},
		{
			4,
			10,
			6,
		},
		{
			5,
			1,
			3,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			rob := robot{
				x:  2,
				y:  4,
				vx: 2,
				vy: -3,
			}
			for range tt.stepCount {
				rob.step(lenX, lenY)
			}
			assert.Equal(t, tt.expectedX, rob.x)
			assert.Equal(t, tt.expectedY, rob.y)
		})
	}
}

func TestIncrement(t *testing.T) {
	length := 5
	tests := []struct {
		val      int
		inc      int
		expected int
	}{
		{
			val:      2,
			inc:      2,
			expected: 4,
		},
		{
			val:      2,
			inc:      3,
			expected: 0,
		},
		{
			val:      2,
			inc:      5,
			expected: 2,
		},
		{
			val:      2,
			inc:      -4,
			expected: 3,
		},
		{
			val:      2,
			inc:      -5,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			result := increment(tt.val, tt.inc, length)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestQuadrant(t *testing.T) {
	lenX := 5
	lenY := 5
	tests := []struct {
		x        int
		y        int
		expected int
	}{
		// quadrant 1
		{0, 0, 1},
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
		// quadrant 2
		{0, 3, 2},
		{0, 4, 2},
		{1, 3, 2},
		{1, 4, 2},
		// quadrant 3
		{3, 0, 3},
		{3, 1, 3},
		{4, 0, 3},
		{4, 1, 3},
		// quadrant 4
		{3, 3, 4},
		{3, 4, 4},
		{4, 3, 4},
		{4, 4, 4},
		// excluded
		{0, 2, 0},
		{1, 2, 0},
		{2, 2, 0},
		{3, 2, 0},
		{4, 2, 0},
		{2, 0, 0},
		{2, 1, 0},
		{2, 3, 0},
		{2, 4, 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d,%d", tt.x, tt.y), func(t *testing.T) {
			b := Bathroom{lenX: lenX, lenY: lenY}
			result := b.quadrant(tt.x, tt.y)
			assert.Equal(t, tt.expected, result)
		})
	}
}
