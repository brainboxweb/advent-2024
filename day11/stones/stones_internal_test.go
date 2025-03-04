package stones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyRules(t *testing.T) {
	tests := []struct {
		data     int
		expected []int
	}{
		{
			0,
			[]int{1},
		},
		{
			1,
			[]int{2024},
		},
		{
			12345678,
			[]int{1234, 5678},
		},
		{
			12300078,
			[]int{1230, 78},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			result := applyRules(tt.data)
			assert.Equal(t, tt.expected, result)
		})
	}
}
