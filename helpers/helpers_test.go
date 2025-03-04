package helpers_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/helpers"
	"github.com/stretchr/testify/assert"
)

func TestToXY(t *testing.T) {

	input := []string{
		"abc",
		"def",
		"ghi",
	}
	result := helpers.ToXY(input)

	tests := []struct {
		x        int
		y        int
		expected string
	}{
		{
			0,
			0,
			"a",
		},
		{
			2,
			0,
			"c",
		},
		{
			1,
			2,
			"h",
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			assert.Equal(t, tt.expected, result[tt.x][tt.y])
		})
	}
}
