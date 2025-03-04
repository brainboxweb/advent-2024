package safety_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day5/safety"

	"github.com/stretchr/testify/assert"
)

var rules = [][]int{
	{47, 53},
	{97, 13},
	{97, 61},
	{97, 47},
	{75, 29},
	{61, 13},
	{75, 53},
	{29, 13},
	{97, 29},
	{53, 29},
	{61, 53},
	{97, 53},
	{61, 29},
	{47, 13},
	{75, 47},
	{97, 75},
	{47, 61},
	{75, 61},
	{47, 29},
	{75, 13},
	{53, 13},
}

func TestValidateOrder(t *testing.T) {
	tests := []struct {
		update   []int
		expected bool
	}{
		{
			[]int{75, 47, 61, 53, 29},
			true,
		}, {
			[]int{97, 61, 53, 29, 13},
			true,
		},
		{
			[]int{75, 29, 13},
			true,
		},
		{
			[]int{75, 97, 47, 61, 53},
			false,
		},
		{
			[]int{61, 13, 29},
			false,
		},
		{
			[]int{97, 13, 75, 29, 47},
			false,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			safety := safety.NewManual(rules)
			result := safety.ValidateOrder(tt.update)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFixOrder(t *testing.T) {
	tests := []struct {
		update   []int
		expected []int
	}{
		{
			[]int{75, 97, 47, 61, 53},
			[]int{97, 75, 47, 61, 53},
		},
		{
			[]int{61, 13, 29},
			[]int{61, 29, 13},
		},
		{
			[]int{97, 13, 75, 29, 47},
			[]int{97, 75, 47, 29, 13},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			safety := safety.NewManual(rules)
			result := safety.FixUpdate(tt.update)
			assert.Equal(t, tt.expected, result)
		})
	}
}
