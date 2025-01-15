package day7_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day7"
	"github.com/brainboxweb/advent-2024/helpers"
	"github.com/stretchr/testify/assert"
)

func TestChallenge(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"../data/day7_test.txt",
			3749,
		},
		{
			"../data/day7.txt",
			1985268524462,
		},
	}

	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day7.ChallengeOne(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestChallengeDay2(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"../data/day7_test.txt",
			11387,
		},
		{
			"../data/day7.txt",
			150077710195188,
		},
	}

	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			dataSet := helpers.GetDataString(tt.dataFile)
			result := day7.ChallengeTwo(dataSet)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		equation day7.Equation
		operands []string
		expected []int
	}{
		{
			day7.Equation{Operands: []int{10, 19}},
			[]string{"*", "+"},
			[]int{190, 29},
		},
		{
			day7.Equation{Operands: []int{81, 40, 27}},
			[]string{"*", "+"},
			[]int{87480, 3267, 3267, 148},
		},
		{
			day7.Equation{Operands: []int{10, 19}},
			[]string{"*", "+", "||"},
			[]int{190, 29, 1019},
		},
		{
			day7.Equation{Operands: []int{81, 40, 27}},
			[]string{"*", "+", "||"},
			[]int{87480, 3267, 219780, 3267, 324027, 148, 12127, 8167, 814027},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			results := day7.Calculate(tt.equation, tt.operands)
			assert.Equal(t, tt.expected, results)
		})
	}
}
