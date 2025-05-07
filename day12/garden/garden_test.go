package garden_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day12/garden"
	"github.com/brainboxweb/advent-2024/helpers"
	"github.com/stretchr/testify/assert"
)

const testDataPath = "../../testdata/"

func TestTotalCost(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"day12_test.txt",
			140,
		},
		{
			"day12_test1.txt",
			772,
		},
		{
			"day12_test2.txt",
			1930,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			data := helpers.GetDataString(testDataPath + tt.dataFile)
			dataSet := helpers.ToXY(data)
			plot := garden.NewPlot(dataSet)
			result := plot.TotalCost()
			assert.Equal(t, tt.expected, result)
		})
	}
}
