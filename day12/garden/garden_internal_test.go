package garden

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMatchingNeighbours(t *testing.T) {
	tests := []struct {
		grid     [][]string
		point    point
		expected []point
	}{
		{
			grid:  getTestGrid("grid1"),
			point: point{1, 1, "B", 0},
			expected: []point{
				{0, 1, "B", 0},
				{1, 2, "B", 0},
			},
		},
		{
			grid:  getTestGrid("grid1"),
			point: point{3, 3, "C", 0},
			expected: []point{
				{3, 2, "C", 0},
			},
		},
		{
			grid:  getTestGrid("grid2"),
			point: point{3, 9, "I", 0},
			expected: []point{
				{3, 8, "I", 0},
			},
		},
		{
			grid:  getTestGrid("grid2"),
			point: point{3, 8, "I", 0},
			expected: []point{
				{3, 7, "I", 0},
				{2, 8, "I", 0},
				{3, 9, "I", 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run("tests", func(t *testing.T) {
			result := getMatchingNeighbours(tt.point, tt.grid)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMergeN(t *testing.T) {
	gg := map[string][]point{
		"0_0": {{x: 0, y: 0, name: "A"}},
		"1_0": {{x: 1, y: 0, name: "A"}},
		"2_0": {{x: 2, y: 0, name: "A"}},
		"3_0": {{x: 3, y: 0, name: "A"}},

		"0_1": {{x: 0, y: 1, name: "B"}},
		"1_1": {{x: 1, y: 1, name: "B"}},
		"2_1": {{x: 2, y: 1, name: "C"}},
		"3_1": {{x: 3, y: 1, name: "D"}},

		"0_2": {{x: 0, y: 2, name: "B"}},
		"1_2": {{x: 1, y: 2, name: "B"}},
		"2_2": {{x: 2, y: 2, name: "C"}},
		"3_2": {{x: 3, y: 2, name: "C"}},

		"0_3": {{x: 0, y: 3, name: "E"}},
		"1_3": {{x: 1, y: 3, name: "E"}},
		"2_3": {{x: 2, y: 3, name: "S"}},
		"3_3": {{x: 3, y: 3, name: "C"}},
	}

	expected := map[string][]point{
		"0_0": {
			{x: 0, y: 0, name: "A"},
			{x: 1, y: 0, name: "A"},
		},
		"2_0": {{x: 2, y: 0, name: "A"}},
		"3_0": {{x: 3, y: 0, name: "A"}},

		"0_1": {{x: 0, y: 1, name: "B"}},
		"1_1": {{x: 1, y: 1, name: "B"}},
		"2_1": {{x: 2, y: 1, name: "C"}},
		"3_1": {{x: 3, y: 1, name: "D"}},

		"0_2": {{x: 0, y: 2, name: "B"}},
		"1_2": {{x: 1, y: 2, name: "B"}},
		"2_2": {{x: 2, y: 2, name: "C"}},
		"3_2": {{x: 3, y: 2, name: "C"}},

		"0_3": {{x: 0, y: 3, name: "E"}},
		"1_3": {{x: 1, y: 3, name: "E"}},
		"2_3": {{x: 2, y: 3, name: "S"}},
		"3_3": {{x: 3, y: 3, name: "C"}},
	}

	tests := []struct {
		grouper      map[string][]point
		currentPoint point
		neighbours   []point
		expected     map[string][]point
	}{
		{
			grouper:      gg,
			currentPoint: point{0, 0, "A", 0},
			neighbours: []point{
				{1, 0, "A", 0},
			},
			expected: expected,
		},
	}

	for _, tt := range tests {
		t.Run("tests", func(t *testing.T) {
			res := mergeNeighbours(tt.grouper, tt.currentPoint, tt.neighbours)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestFindRegions(t *testing.T) {
	// AAAA
	// BBCD
	// BBCC
	// EEEC

	input := [][]string{
		{"A", "A", "A", "A"},
		{"B", "B", "C", "D"},
		{"B", "B", "C", "C"},
		{"E", "E", "E", "C"},
	}

	expected := map[string][]point{
		"2_3": {
			{x: 0, y: 3, name: "E", neighbourCount: 0},
			{x: 1, y: 3, name: "E", neighbourCount: 0},
			{x: 2, y: 3, name: "E", neighbourCount: 0},
		},
		"3_3": {
			{x: 2, y: 1, name: "C", neighbourCount: 0},
			{x: 2, y: 2, name: "C", neighbourCount: 0},
			{x: 3, y: 2, name: "C", neighbourCount: 0},
			{x: 3, y: 3, name: "C", neighbourCount: 0},
		},
		"3_1": {
			{x: 3, y: 1, name: "D", neighbourCount: 0},
		},
		"3_0": {
			{x: 0, y: 0, name: "A", neighbourCount: 0},
			{x: 1, y: 0, name: "A", neighbourCount: 0},
			{x: 2, y: 0, name: "A", neighbourCount: 0},
			{x: 3, y: 0, name: "A", neighbourCount: 0},
		},
		"1_2": {
			{x: 0, y: 1, name: "B", neighbourCount: 0},
			{x: 1, y: 1, name: "B", neighbourCount: 0},
			{x: 0, y: 2, name: "B", neighbourCount: 0},
			{x: 1, y: 2, name: "B", neighbourCount: 0},
		},
	}

	result := findRegions(input)
	assert.Equal(t, expected, result)
}

func TestFindRegions2(t *testing.T) {
	input2 := [][]string{
		{"R", "R", "R", "R", "I", "I", "C", "C", "F", "F"},
		{"R", "R", "R", "R", "I", "I", "C", "C", "C", "F"},
		{"V", "V", "R", "R", "R", "C", "C", "F", "F", "F"},
		{"V", "V", "R", "C", "C", "C", "J", "F", "F", "F"},
		{"V", "V", "V", "V", "C", "J", "J", "C", "F", "E"},
		{"V", "V", "I", "V", "C", "C", "J", "J", "E", "E"},
		{"V", "V", "I", "I", "I", "C", "J", "J", "E", "E"},
		{"M", "I", "I", "I", "I", "I", "J", "J", "E", "E"},
		{"M", "I", "I", "I", "S", "I", "J", "E", "E", "E"},
		{"M", "M", "M", "I", "S", "S", "J", "E", "E", "E"},
	}

	expected2 := map[string][]point{
		"5_1": {
			{x: 4, y: 0, name: "I", neighbourCount: 0},
			{x: 5, y: 0, name: "I", neighbourCount: 0},
			{x: 4, y: 1, name: "I", neighbourCount: 0},
			{x: 5, y: 1, name: "I", neighbourCount: 0},
		},
		"7_4": {
			{x: 7, y: 4, name: "C", neighbourCount: 0}, // This is correct
		},
		"8_4": {

			{x: 7, y: 2, name: "F", neighbourCount: 0},
			{x: 8, y: 2, name: "F", neighbourCount: 0},
			{x: 7, y: 3, name: "F", neighbourCount: 0},
			{x: 8, y: 0, name: "F", neighbourCount: 0},
			{x: 9, y: 0, name: "F", neighbourCount: 0},
			{x: 9, y: 1, name: "F", neighbourCount: 0},
			{x: 9, y: 2, name: "F", neighbourCount: 0},
			{x: 8, y: 3, name: "F", neighbourCount: 0},
			{x: 9, y: 3, name: "F", neighbourCount: 0},
			{x: 8, y: 4, name: "F", neighbourCount: 0},
		},
		"1_6": {
			{x: 0, y: 2, name: "V", neighbourCount: 0},
			{x: 1, y: 2, name: "V", neighbourCount: 0},
			{x: 0, y: 3, name: "V", neighbourCount: 0},
			{x: 1, y: 3, name: "V", neighbourCount: 0},
			{x: 0, y: 4, name: "V", neighbourCount: 0},
			{x: 1, y: 4, name: "V", neighbourCount: 0},
			{x: 0, y: 5, name: "V", neighbourCount: 0},
			{x: 2, y: 4, name: "V", neighbourCount: 0},
			{x: 1, y: 5, name: "V", neighbourCount: 0},
			{x: 3, y: 4, name: "V", neighbourCount: 0},
			{x: 3, y: 5, name: "V", neighbourCount: 0},
			{x: 0, y: 6, name: "V", neighbourCount: 0},
			{x: 1, y: 6, name: "V", neighbourCount: 0},
		},
		"9_9": {
			{x: 7, y: 8, name: "E", neighbourCount: 0},
			{x: 8, y: 5, name: "E", neighbourCount: 0},
			{x: 9, y: 4, name: "E", neighbourCount: 0},
			{x: 9, y: 5, name: "E", neighbourCount: 0},
			{x: 8, y: 6, name: "E", neighbourCount: 0},
			{x: 9, y: 6, name: "E", neighbourCount: 0},
			{x: 8, y: 7, name: "E", neighbourCount: 0},
			{x: 9, y: 7, name: "E", neighbourCount: 0},
			{x: 8, y: 8, name: "E", neighbourCount: 0},
			{x: 9, y: 8, name: "E", neighbourCount: 0},
			{x: 7, y: 9, name: "E", neighbourCount: 0},
			{x: 8, y: 9, name: "E", neighbourCount: 0},
			{x: 9, y: 9, name: "E", neighbourCount: 0},
		},
		"2_9": {
			{x: 0, y: 7, name: "M", neighbourCount: 0},
			{x: 0, y: 8, name: "M", neighbourCount: 0},
			{x: 0, y: 9, name: "M", neighbourCount: 0},
			{x: 1, y: 9, name: "M", neighbourCount: 0},
			{x: 2, y: 9, name: "M", neighbourCount: 0},
		},
		"5_9": {
			{x: 4, y: 8, name: "S", neighbourCount: 0},
			{x: 4, y: 9, name: "S", neighbourCount: 0},
			{x: 5, y: 9, name: "S", neighbourCount: 0},
		},
		"2_3": {
			{x: 0, y: 0, name: "R", neighbourCount: 0},
			{x: 1, y: 0, name: "R", neighbourCount: 0},
			{x: 0, y: 1, name: "R", neighbourCount: 0},
			{x: 2, y: 0, name: "R", neighbourCount: 0},
			{x: 1, y: 1, name: "R", neighbourCount: 0},
			{x: 3, y: 0, name: "R", neighbourCount: 0},
			{x: 2, y: 1, name: "R", neighbourCount: 0},
			{x: 3, y: 1, name: "R", neighbourCount: 0},
			{x: 2, y: 2, name: "R", neighbourCount: 0},
			{x: 3, y: 2, name: "R", neighbourCount: 0},
			{x: 2, y: 3, name: "R", neighbourCount: 0},
			{x: 4, y: 2, name: "R", neighbourCount: 0},
		},
		"5_6": {
			{x: 3, y: 3, name: "C", neighbourCount: 0},
			{x: 4, y: 3, name: "C", neighbourCount: 0},
			{x: 5, y: 2, name: "C", neighbourCount: 0},
			{x: 6, y: 0, name: "C", neighbourCount: 0},
			{x: 7, y: 0, name: "C", neighbourCount: 0},
			{x: 6, y: 1, name: "C", neighbourCount: 0},
			{x: 7, y: 1, name: "C", neighbourCount: 0},
			{x: 6, y: 2, name: "C", neighbourCount: 0},
			{x: 8, y: 1, name: "C", neighbourCount: 0},
			{x: 5, y: 3, name: "C", neighbourCount: 0},
			{x: 4, y: 4, name: "C", neighbourCount: 0},
			{x: 4, y: 5, name: "C", neighbourCount: 0},
			{x: 5, y: 5, name: "C", neighbourCount: 0},
			{x: 5, y: 6, name: "C", neighbourCount: 0},
		},
		"3_9": {
			{x: 1, y: 7, name: "I", neighbourCount: 0},
			{x: 2, y: 5, name: "I", neighbourCount: 0},
			{x: 2, y: 6, name: "I", neighbourCount: 0},
			{x: 3, y: 6, name: "I", neighbourCount: 0},
			{x: 2, y: 7, name: "I", neighbourCount: 0},
			{x: 4, y: 6, name: "I", neighbourCount: 0},
			{x: 3, y: 7, name: "I", neighbourCount: 0},
			{x: 4, y: 7, name: "I", neighbourCount: 0},
			{x: 1, y: 8, name: "I", neighbourCount: 0},
			{x: 2, y: 8, name: "I", neighbourCount: 0},
			{x: 3, y: 8, name: "I", neighbourCount: 0},
			{x: 5, y: 7, name: "I", neighbourCount: 0},
			{x: 5, y: 8, name: "I", neighbourCount: 0},
			{x: 3, y: 9, name: "I", neighbourCount: 0},
		},
		"6_9": {
			{x: 5, y: 4, name: "J", neighbourCount: 0},
			{x: 6, y: 3, name: "J", neighbourCount: 0},
			{x: 6, y: 4, name: "J", neighbourCount: 0},
			{x: 6, y: 5, name: "J", neighbourCount: 0},
			{x: 7, y: 5, name: "J", neighbourCount: 0},
			{x: 6, y: 6, name: "J", neighbourCount: 0},
			{x: 7, y: 6, name: "J", neighbourCount: 0},
			{x: 6, y: 7, name: "J", neighbourCount: 0},
			{x: 7, y: 7, name: "J", neighbourCount: 0},
			{x: 6, y: 8, name: "J", neighbourCount: 0},
			{x: 6, y: 9, name: "J", neighbourCount: 0},
		},
	}

	result2 := findRegions(input2)
	assert.Equal(t, expected2, result2)
}

func TestScan(t *testing.T) { // rename
	input := [][]string{
		{"A", "A", "A", "A"},
		{"B", "B", "C", "D"},
		{"B", "B", "C", "C"},
		{"E", "E", "E", "C"},
	}
	expected := []point{
		{
			x:              0,
			y:              0,
			name:           "A",
			neighbourCount: 0,
		},
		{
			x:              1,
			y:              0,
			name:           "A",
			neighbourCount: 0,
		},
		{
			x:              2,
			y:              0,
			name:           "A",
			neighbourCount: 0,
		},
		{
			x:              3,
			y:              0,
			name:           "A",
			neighbourCount: 0,
		},
		{
			x:              0,
			y:              1,
			name:           "B",
			neighbourCount: 0,
		},
		{
			x:              1,
			y:              1,
			name:           "B",
			neighbourCount: 0,
		},
		{
			x:              2,
			y:              1,
			name:           "C",
			neighbourCount: 0,
		},
		{
			x:              3,
			y:              1,
			name:           "D",
			neighbourCount: 0,
		},
		{
			x:              0,
			y:              2,
			name:           "B",
			neighbourCount: 0,
		},
		{
			x:              1,
			y:              2,
			name:           "B",
			neighbourCount: 0,
		},
		{
			x:              2,
			y:              2,
			name:           "C",
			neighbourCount: 0,
		},
		{
			x:              3,
			y:              2,
			name:           "C",
			neighbourCount: 0,
		},
		{
			x:              0,
			y:              3,
			name:           "E",
			neighbourCount: 0,
		},
		{
			x:              1,
			y:              3,
			name:           "E",
			neighbourCount: 0,
		},
		{
			x:              2,
			y:              3,
			name:           "E",
			neighbourCount: 0,
		},
		{
			x:              3,
			y:              3,
			name:           "C",
			neighbourCount: 0,
		},
	}

	result := scan(input)

	assert.Equal(t, expected, result)
}

func getTestGrid(name string) [][]string {
	grid1 := [][]string{
		{"A", "A", "A", "A"},
		{"B", "B", "C", "D"},
		{"B", "B", "C", "C"},
		{"E", "E", "E", "C"},
	}
	grid2 := [][]string{
		{"R", "R", "R", "R", "I", "I", "C", "C", "F", "F"},
		{"R", "R", "R", "R", "I", "I", "C", "C", "C", "F"},
		{"V", "V", "R", "R", "R", "C", "C", "F", "F", "F"},
		{"V", "V", "R", "C", "C", "C", "J", "F", "F", "F"},
		{"V", "V", "V", "V", "C", "J", "J", "C", "F", "E"},
		{"V", "V", "I", "V", "C", "C", "J", "J", "E", "E"},
		{"V", "V", "I", "I", "I", "C", "J", "J", "E", "E"},
		{"M", "I", "I", "I", "I", "I", "J", "J", "E", "E"},
		{"M", "I", "I", "I", "S", "I", "J", "E", "E", "E"},
		{"M", "M", "M", "I", "S", "S", "J", "E", "E", "E"},
	}
	if name == "grid1" {
		return grid1
	}

	return grid2
}

func TestPerimeter(t *testing.T) {
	plot := Plot{}
	plot.grid = [][]string{
		{"R", "R", "R", "R", "I", "I", "C", "C", "F", "F"},
		{"R", "R", "R", "R", "I", "I", "C", "C", "C", "F"},
		{"V", "V", "R", "R", "R", "C", "C", "F", "F", "F"},
		{"V", "V", "R", "C", "C", "C", "J", "F", "F", "F"},
		{"V", "V", "V", "V", "C", "J", "J", "C", "F", "E"},
		{"V", "V", "I", "V", "C", "C", "J", "J", "E", "E"},
		{"V", "V", "I", "I", "I", "C", "J", "J", "E", "E"},
		{"M", "I", "I", "I", "I", "I", "J", "J", "E", "E"},
		{"M", "I", "I", "I", "S", "I", "J", "E", "E", "E"},
		{"M", "M", "M", "I", "S", "S", "J", "E", "E", "E"},
	}

	plot.regions = [][]point{
		{
			{x: 4, y: 0, name: "I", neighbourCount: 0},
			{x: 5, y: 0, name: "I", neighbourCount: 0},
			{x: 4, y: 1, name: "I", neighbourCount: 0},
			{x: 5, y: 1, name: "I", neighbourCount: 0},
		},
		{
			{x: 7, y: 8, name: "E", neighbourCount: 0},
			{x: 8, y: 5, name: "E", neighbourCount: 0},
			{x: 9, y: 4, name: "E", neighbourCount: 0},
			{x: 9, y: 5, name: "E", neighbourCount: 0},
			{x: 8, y: 6, name: "E", neighbourCount: 0},
			{x: 9, y: 6, name: "E", neighbourCount: 0},
			{x: 8, y: 7, name: "E", neighbourCount: 0},
			{x: 9, y: 7, name: "E", neighbourCount: 0},
			{x: 8, y: 8, name: "E", neighbourCount: 0},
			{x: 9, y: 8, name: "E", neighbourCount: 0},
			{x: 7, y: 9, name: "E", neighbourCount: 0},
			{x: 8, y: 9, name: "E", neighbourCount: 0},
			{x: 9, y: 9, name: "E", neighbourCount: 0},
		},
	}

	tests := []struct {
		region   []point
		expected int
	}{
		{
			[]point{
				{x: 4, y: 0, name: "I", neighbourCount: 0},
				{x: 5, y: 0, name: "I", neighbourCount: 0},
				{x: 4, y: 1, name: "I", neighbourCount: 0},
				{x: 5, y: 1, name: "I", neighbourCount: 0},
			},
			8,
		},
		{
			[]point{
				{x: 7, y: 8, name: "E", neighbourCount: 0},
				{x: 8, y: 5, name: "E", neighbourCount: 0},
				{x: 9, y: 4, name: "E", neighbourCount: 0},
				{x: 9, y: 5, name: "E", neighbourCount: 0},
				{x: 8, y: 6, name: "E", neighbourCount: 0},
				{x: 9, y: 6, name: "E", neighbourCount: 0},
				{x: 8, y: 7, name: "E", neighbourCount: 0},
				{x: 9, y: 7, name: "E", neighbourCount: 0},
				{x: 8, y: 8, name: "E", neighbourCount: 0},
				{x: 9, y: 8, name: "E", neighbourCount: 0},
				{x: 7, y: 9, name: "E", neighbourCount: 0},
				{x: 8, y: 9, name: "E", neighbourCount: 0},
				{x: 9, y: 9, name: "E", neighbourCount: 0},
			},
			18,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			result := plot.perimeter(tt.region)
			assert.Equal(t, tt.expected, result)
		})
	}
}
