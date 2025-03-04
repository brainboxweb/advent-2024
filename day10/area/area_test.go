package area_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day10/area"
	"github.com/brainboxweb/advent-2024/helpers"
	"github.com/stretchr/testify/assert"
)

const testDataFile = "../../data/"

func TestStartingPoints(t *testing.T) {
	tests := []struct {
		dataFile string
		expected []area.Point
	}{
		{
			"day10_test.txt",
			[]area.Point{
				{X: 0, Y: 0},
			},
		},
		{
			"day10_test2.txt",
			[]area.Point{
				{X: 3, Y: 0},
			},
		},
		{
			"day10_test3.txt",
			[]area.Point{
				{X: 3, Y: 0},
			},
		},
		{
			"day10_test4.txt",
			[]area.Point{
				{X: 0, Y: 6},
				{X: 1, Y: 7},
				{X: 2, Y: 0},
				{X: 2, Y: 5},
				{X: 4, Y: 0},
				{X: 4, Y: 2},
				{X: 5, Y: 5},
				{X: 6, Y: 4},
				{X: 6, Y: 6},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			data := helpers.GetDataString(testDataFile + tt.dataFile)
			dataSet := helpers.ToXY(data)

			terrain := area.New(dataSet)
			terrain.FindStartingPoints()
			result := terrain.StartingPoints
			assert.Equal(t, tt.expected, result)
		})
	}
}
