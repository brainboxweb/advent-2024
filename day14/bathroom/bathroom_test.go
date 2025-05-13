package bathroom_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day14/bathroom"
	"github.com/stretchr/testify/assert"
)

func TestSteps(t *testing.T) {
	lenX := 11
	lenY := 7
	tests := []struct {
		robX      int
		robY      int
		robVX     int
		robVY     int
		stepCount int
		expected  map[int]int
	}{
		{
			2,
			4,
			2,
			-3,
			1,
			map[int]int{
				1: 1,
			},
		},
		{
			2,
			4,
			2,
			-3,
			2,
			map[int]int{
				4: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			isle := bathroom.New(lenX, lenY)
			isle.AddRobot(tt.robX, tt.robY, tt.robVX, tt.robVY)
			isle.Step(tt.stepCount)
			result := isle.CountQuadrants()
			assert.Equal(t, tt.expected, result)
		})
	}
}
