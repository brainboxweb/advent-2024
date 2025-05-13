package island_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day13/island"
	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		thing    island.Machine
		expected int
	}{
		{
			island.Machine{
				A: island.Point{
					94,
					34,
				},
				B: island.Point{
					22,
					67,
				},
				Prize: island.Point{
					X: 8400,
					Y: 5400,
				},
			},
			280,
		},
		{
			island.Machine{
				A: island.Point{
					26,
					66,
				},
				B: island.Point{
					67,
					21,
				},
				Prize: island.Point{
					X: 12748,
					Y: 12176,
				},
			},
			0,
		},
		{
			island.Machine{
				A: island.Point{
					17,
					86,
				},
				B: island.Point{
					84,
					37,
				},
				Prize: island.Point{
					X: 7870,
					Y: 6450,
				},
			},
			200,
		},
		{
			island.Machine{
				A: island.Point{
					69,
					23,
				},
				B: island.Point{
					27,
					71,
				},
				Prize: island.Point{
					X: 18641,
					Y: 10279,
				},
			},
			0,
		},
		{
			island.Machine{
				A: island.Point{
					36,
					99,
				},
				B: island.Point{
					49,
					23,
				},
				Prize: island.Point{
					X: 2107,
					Y: 5012,
				},
			},
			154,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			result := tt.thing.Calculate(100)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// func TestCalculateTwo(t *testing.T) { // Too slow to run!
// 	tests := []struct {
// 		thing    island.Machine
// 		expected int
// 	}{
// 		{
// 			island.Machine{
// 				A: island.Point{
// 					26,
// 					66,
// 				},
// 				B: island.Point{
// 					67,
// 					21,
// 				},
// 				Prize: island.Point{
// 					X: 10000000012748,
// 					Y: 10000000012176,
// 				},
// 			},
// 			280,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run("test", func(t *testing.T) {
// 			result := tt.thing.Calculate(-1)
// 			assert.Equal(t, tt.expected, result)
// 		})
// 	}
// }
