package day17_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day17/computer"
	"github.com/stretchr/testify/assert"
)

func TestChallenge(t *testing.T) {
	tests := []struct {
		regA         int
		instructions []int
		expected     string
	}{
		{
			52042868,
			[]int{2, 4, 1, 7, 7, 5, 0, 3, 4, 4, 1, 7, 5, 5, 3, 0},
			"2,1,0,1,7,2,5,0,3", // <-- Today's solution
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			comp := computer.New()
			comp.SetRegisterA(tt.regA)
			result := comp.Run(tt.instructions)
			assert.Equal(t, tt.expected, result)
		})
	}
}
