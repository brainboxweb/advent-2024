package calculations_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day7/calculations"
	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		subtotal  int
		operands  []int
		operators []string
		expected  []int
	}{
		{
			10,
			[]int{19},
			[]string{"*", "+"},
			[]int{190, 29},
		},
		{
			81,
			[]int{40, 27},
			[]string{"*", "+"},
			[]int{87480, 3267, 3267, 148},
		},
		{
			10,
			[]int{19},
			[]string{"*", "+", "||"},
			[]int{190, 29, 1019},
		},
		{
			81,
			[]int{40, 27},
			[]string{"*", "+", "||"},
			[]int{87480, 3267, 219780, 3267, 324027, 148, 12127, 8167, 814027},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			calc := calculations.NewCalculator(tt.operators)
			results := calc.Calculate(tt.subtotal, tt.operands)
			assert.Equal(t, tt.expected, results)
		})
	}
}
