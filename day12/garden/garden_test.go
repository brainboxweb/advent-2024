package garden_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day12/garden"
	"github.com/brainboxweb/advent-2024/helpers"
	"github.com/stretchr/testify/assert"
)

const testDataPath = "../../testdata/" // Hmm do I need this

// fix and add back
// func TestNewPlot(t *testing.T) {
// 	input := [][]string{
// 		{
// 			"A", "A", "A", "A",
// 		},
// 		{
// 			"B", "B", "C", "D",
// 		},
// 		{
// 			"B", "B", "C", "C",
// 		},
// 		{
// 			"E", "E", "E", "C",
// 		},
// 	}

// 	plot := garden.NewPlot(input)

// 	for _, region := range plot.Areas {
// 		fmt.Println(region[0].Name, len(region)) // works!
// 	}

// 	fmt.Println("Permimeter")
// 	for _, region := range plot.Areas {
// 		// fmt.Println(region[0].Name, len(region)) // works!
// 		perim := plot.Perimeter(region)
// 		fmt.Println(region[0].Name, perim)
// 	}

// }

// func TestArea(t *testing.T) { // rename

// 	// this is not right. Need to match the areas and the perims

// 	tests := []struct {
// 		dataFile string
// 		letter   string
// 		expected int
// 	}{
// 		{
// 			"day12_test.txt",
// 			"A",
// 			4,
// 		},
// 		{
// 			"day12_test.txt",
// 			"B",
// 			4,
// 		},
// 		{
// 			"day12_test.txt",
// 			"C",
// 			4,
// 		},
// 		{
// 			"day12_test.txt",
// 			"D",
// 			1,
// 		},
// 		{
// 			"day12_test.txt",
// 			"E",
// 			3,
// 		},
// 		{
// 			"day12_test1.txt",
// 			"O",
// 			21,
// 		},
// 		{
// 			"day12_test1.txt",
// 			"X",
// 			4,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.letter, func(t *testing.T) {
// 			data := helpers.GetDataString(testDataPath + tt.dataFile)
// 			dataSet := helpers.ToXY(data)
// 			plot := garden.NewPlot(dataSet)

// 			result := plot.Area(tt.letter)

// 			assert.Equal(t, tt.expected, result)
// 		})
// 	}
// }

// func TestPerimeter(t *testing.T) { // rename
// 	tests := []struct {
// 		dataFile string
// 		letter   string
// 		expected int
// 	}{
// 		{
// 			"day12_test.txt",
// 			"A",
// 			10,
// 		},
// 		{
// 			"day12_test.txt",
// 			"B",
// 			8,
// 		},
// 		{
// 			"day12_test.txt",
// 			"C",
// 			10,
// 		},
// 		{
// 			"day12_test.txt",
// 			"D",
// 			4,
// 		},
// 		{
// 			"day12_test.txt",
// 			"E",
// 			8,
// 		},
// 		{
// 			"day12_test1.txt",
// 			"O",
// 			36,
// 		},
// 		{
// 			"day12_test1.txt",
// 			"X",
// 			16,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.letter, func(t *testing.T) {
// 			data := helpers.GetDataString(testDataPath + tt.dataFile)
// 			dataSet := helpers.ToXY(data)
// 			plot := garden.NewPlot(dataSet)

// 			result := plot.Perimeter(tt.letter)

// 			assert.Equal(t, tt.expected, result)
// 		})
// 	}
// }

// func TestCost(t *testing.T) { // rename
// 	tests := []struct {
// 		dataFile string
// 		letter   string
// 		expected int
// 	}{
// 		{
// 			"day12_test.txt",
// 			"A",
// 			40,
// 		},
// 		{
// 			"day12_test.txt",
// 			"B",
// 			32,
// 		},
// 		{
// 			"day12_test.txt",
// 			"C",
// 			40,
// 		},
// 		{
// 			"day12_test.txt",
// 			"D",
// 			4,
// 		},
// 		{
// 			"day12_test.txt",
// 			"E",
// 			24,
// 		},
// 		{
// 			"day12_test1.txt",
// 			"O",
// 			756,
// 		},
// 		{
// 			"day12_test1.txt",
// 			"X",
// 			64,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.letter, func(t *testing.T) {
// 			data := helpers.GetDataString(testDataPath + tt.dataFile)
// 			dataSet := helpers.ToXY(data)
// 			plot := garden.NewPlot(dataSet)

// 			result := plot.Cost(tt.letter)

// 			assert.Equal(t, tt.expected, result)
// 		})
// 	}
// }

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
		{
			"day12.txt",
			1522850,
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
