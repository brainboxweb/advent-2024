package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input     string
		expectedP point
		expectedV velocity
	}{
		{
			"p=0,4 v=3,-3",
			point{0, 4},
			velocity{3, -3},
		},
		{
			"p=10,3 v=-1,2",
			point{10, 3},
			velocity{-1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			point, velocity := parse(tt.input)
			assert.Equal(t, tt.expectedP, point)
			assert.Equal(t, tt.expectedV, velocity)
		})
	}
}
