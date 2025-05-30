package branding_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day19/branding"
	"github.com/stretchr/testify/assert"
)

func TestIsPossible(t *testing.T) {
	towels := branding.NewTowels([]string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"})
	tests := []struct {
		pattern  string
		expected bool
	}{
		{
			"brwrr",
			true,
		},
		{
			"bggr",
			true,
		},
		{
			"gbbr",
			true,
		},
		{
			"rrbgbr",
			true,
		},
		{
			"ubwu",
			false,
		},
		{
			"bwurrg",
			true,
		},
		{
			"brgr",
			true,
		},
		{
			"bbrgwb",
			false,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			result := towels.IsPossible(tt.pattern)
			assert.Equal(t, tt.expected, result)
		})
	}
}
