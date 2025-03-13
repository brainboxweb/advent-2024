package roof

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPairs(t *testing.T) {
	tests := []struct {
		start    []Point
		expected [][]Point
	}{
		{
			start: []Point{
				{4, 3},
				{5, 5},
				{9, 4},
			},
			expected: [][]Point{
				{
					{4, 3},
					{5, 5},
				},
				{
					{4, 3},
					{9, 4},
				},
				{
					{5, 5},
					{9, 4},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			result := getPairs(tt.start, nil)
			assert.Equal(t, tt.expected, result)
		})
	}
}
