package roof_test

import (
	"strings"
	"testing"

	"github.com/brainboxweb/advent-2024/day8/roof"
	"github.com/brainboxweb/advent-2024/helpers"
	"github.com/stretchr/testify/assert"
)

func TestNewGrid(t *testing.T) {
	tests := []struct {
		data             []string
		expectedGridSize int
		expectedAntennas map[string][]roof.Point
	}{
		{
			testData(t1),
			10,
			map[string][]roof.Point{
				"a": {
					{4, 3},
					{5, 5},
				},
			},
		},
		{
			testData(t2),
			10,
			map[string][]roof.Point{
				"a": {
					{4, 3},
					{8, 4},
					{5, 5},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			dataSet := helpers.ToXY(tt.data)
			grid := roof.NewGrid(dataSet)
			assert.Equal(t, tt.expectedGridSize, grid.GridSize)
			assert.Equal(t, tt.expectedAntennas, grid.Antennas)
		})
	}
}

func TestAntinodes(t *testing.T) {
	tests := []struct {
		data          []string
		comprehensive bool
		expected      map[string]roof.Point
	}{
		{
			data:          testData(t1),
			comprehensive: false,
			expected: map[string]roof.Point{
				"3-1": {3, 1},
				"6-7": {6, 7},
			},
		},
		{
			data:          testData(t2),
			comprehensive: false,
			expected: map[string]roof.Point{
				"0-2": {0, 2},
				"2-6": {2, 6},
				"3-1": {3, 1},
				"6-7": {6, 7},
			},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			dataSet := helpers.ToXY(tt.data)
			grid := roof.NewGrid(dataSet)
			antinodes := grid.Antinodes()
			assert.Equal(t, tt.expected, antinodes)
		})
	}
}

func TestAntinodesComprehensive(t *testing.T) {
	tests := []struct {
		data          []string
		comprehensive bool
		expected      map[string]roof.Point
	}{
		{
			data:          testData(t3),
			comprehensive: true,
			expected: map[string]roof.Point{
				"0-0": {0, 0},
				"1-2": {1, 2},
				"2-4": {2, 4},
				"3-1": {3, 1},
				"3-6": {3, 6},
				"4-8": {4, 8},
				"5-0": {5, 0},
				"6-2": {6, 2},
				"9-3": {9, 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			dataSet := helpers.ToXY(tt.data)
			grid := roof.NewGrid(dataSet)
			antinodes := grid.AntinodesComprehensive()
			assert.Equal(t, tt.expected, antinodes)
		})
	}
}

func testData(name string) []string {
	return strings.Split(name, "\n")
}

var t1 = `..........
..........
..........
....a.....
..........
.....a....
..........
..........
..........
..........`

var t2 = `..........
..........
..........
....a.....
........a.
.....a....
..........
..........
..........
..........`

var t3 = `T.........
...T......
.T........
..........
..........
..........
..........
..........
..........
..........`
