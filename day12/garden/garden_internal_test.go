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
		"0_0": {{X: 0, Y: 0, Name: "A"}},
		"1_0": {{X: 1, Y: 0, Name: "A"}},
		"2_0": {{X: 2, Y: 0, Name: "A"}},
		"3_0": {{X: 3, Y: 0, Name: "A"}},

		"0_1": {{X: 0, Y: 1, Name: "B"}},
		"1_1": {{X: 1, Y: 1, Name: "B"}},
		"2_1": {{X: 2, Y: 1, Name: "C"}},
		"3_1": {{X: 3, Y: 1, Name: "D"}},

		"0_2": {{X: 0, Y: 2, Name: "B"}},
		"1_2": {{X: 1, Y: 2, Name: "B"}},
		"2_2": {{X: 2, Y: 2, Name: "C"}},
		"3_2": {{X: 3, Y: 2, Name: "C"}},

		"0_3": {{X: 0, Y: 3, Name: "E"}},
		"1_3": {{X: 1, Y: 3, Name: "E"}},
		"2_3": {{X: 2, Y: 3, Name: "S"}},
		"3_3": {{X: 3, Y: 3, Name: "C"}},
	}

	expected := map[string][]point{
		"0_0": {
			{X: 0, Y: 0, Name: "A"},
			{X: 1, Y: 0, Name: "A"},
		},
		"2_0": {{X: 2, Y: 0, Name: "A"}},
		"3_0": {{X: 3, Y: 0, Name: "A"}},

		"0_1": {{X: 0, Y: 1, Name: "B"}},
		"1_1": {{X: 1, Y: 1, Name: "B"}},
		"2_1": {{X: 2, Y: 1, Name: "C"}},
		"3_1": {{X: 3, Y: 1, Name: "D"}},

		"0_2": {{X: 0, Y: 2, Name: "B"}},
		"1_2": {{X: 1, Y: 2, Name: "B"}},
		"2_2": {{X: 2, Y: 2, Name: "C"}},
		"3_2": {{X: 3, Y: 2, Name: "C"}},

		"0_3": {{X: 0, Y: 3, Name: "E"}},
		"1_3": {{X: 1, Y: 3, Name: "E"}},
		"2_3": {{X: 2, Y: 3, Name: "S"}},
		"3_3": {{X: 3, Y: 3, Name: "C"}},
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

func TestFindAreas(t *testing.T) { // rename
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
			{X: 0, Y: 3, Name: "E", neighbourCount: 0},
			{X: 1, Y: 3, Name: "E", neighbourCount: 0},
			{X: 2, Y: 3, Name: "E", neighbourCount: 0},
		},
		"3_3": {
			{X: 2, Y: 1, Name: "C", neighbourCount: 0},
			{X: 2, Y: 2, Name: "C", neighbourCount: 0},
			{X: 3, Y: 2, Name: "C", neighbourCount: 0},
			{X: 3, Y: 3, Name: "C", neighbourCount: 0},
		},
		"3_1": {
			{X: 3, Y: 1, Name: "D", neighbourCount: 0},
		},
		"3_0": {
			{X: 0, Y: 0, Name: "A", neighbourCount: 0},
			{X: 1, Y: 0, Name: "A", neighbourCount: 0},
			{X: 2, Y: 0, Name: "A", neighbourCount: 0},
			{X: 3, Y: 0, Name: "A", neighbourCount: 0},
		},
		"1_2": {
			{X: 0, Y: 1, Name: "B", neighbourCount: 0},
			{X: 1, Y: 1, Name: "B", neighbourCount: 0},
			{X: 0, Y: 2, Name: "B", neighbourCount: 0},
			{X: 1, Y: 2, Name: "B", neighbourCount: 0},
		},
	}

	result := findAreas(input)
	assert.Equal(t, expected, result)
}

func TestFindAreasAgain(t *testing.T) { // merge with above
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
			{X: 4, Y: 0, Name: "I", neighbourCount: 0},
			{X: 5, Y: 0, Name: "I", neighbourCount: 0},
			{X: 4, Y: 1, Name: "I", neighbourCount: 0},
			{X: 5, Y: 1, Name: "I", neighbourCount: 0},
		},
		"7_4": {
			{X: 7, Y: 4, Name: "C", neighbourCount: 0}, // This is correct
		},
		"8_4": {

			{X: 7, Y: 2, Name: "F", neighbourCount: 0},
			{X: 8, Y: 2, Name: "F", neighbourCount: 0},
			{X: 7, Y: 3, Name: "F", neighbourCount: 0},
			{X: 8, Y: 0, Name: "F", neighbourCount: 0},
			{X: 9, Y: 0, Name: "F", neighbourCount: 0},
			{X: 9, Y: 1, Name: "F", neighbourCount: 0},
			{X: 9, Y: 2, Name: "F", neighbourCount: 0},
			{X: 8, Y: 3, Name: "F", neighbourCount: 0},
			{X: 9, Y: 3, Name: "F", neighbourCount: 0},
			{X: 8, Y: 4, Name: "F", neighbourCount: 0},
		},
		"1_6": {
			{X: 0, Y: 2, Name: "V", neighbourCount: 0},
			{X: 1, Y: 2, Name: "V", neighbourCount: 0},
			{X: 0, Y: 3, Name: "V", neighbourCount: 0},
			{X: 1, Y: 3, Name: "V", neighbourCount: 0},
			{X: 0, Y: 4, Name: "V", neighbourCount: 0},
			{X: 1, Y: 4, Name: "V", neighbourCount: 0},
			{X: 0, Y: 5, Name: "V", neighbourCount: 0},
			{X: 2, Y: 4, Name: "V", neighbourCount: 0},
			{X: 1, Y: 5, Name: "V", neighbourCount: 0},
			{X: 3, Y: 4, Name: "V", neighbourCount: 0},
			{X: 3, Y: 5, Name: "V", neighbourCount: 0},
			{X: 0, Y: 6, Name: "V", neighbourCount: 0},
			{X: 1, Y: 6, Name: "V", neighbourCount: 0},
		},
		"9_9": {
			{X: 7, Y: 8, Name: "E", neighbourCount: 0},
			{X: 8, Y: 5, Name: "E", neighbourCount: 0},
			{X: 9, Y: 4, Name: "E", neighbourCount: 0},
			{X: 9, Y: 5, Name: "E", neighbourCount: 0},
			{X: 8, Y: 6, Name: "E", neighbourCount: 0},
			{X: 9, Y: 6, Name: "E", neighbourCount: 0},
			{X: 8, Y: 7, Name: "E", neighbourCount: 0},
			{X: 9, Y: 7, Name: "E", neighbourCount: 0},
			{X: 8, Y: 8, Name: "E", neighbourCount: 0},
			{X: 9, Y: 8, Name: "E", neighbourCount: 0},
			{X: 7, Y: 9, Name: "E", neighbourCount: 0},
			{X: 8, Y: 9, Name: "E", neighbourCount: 0},
			{X: 9, Y: 9, Name: "E", neighbourCount: 0},
		},
		"2_9": {
			{X: 0, Y: 7, Name: "M", neighbourCount: 0},
			{X: 0, Y: 8, Name: "M", neighbourCount: 0},
			{X: 0, Y: 9, Name: "M", neighbourCount: 0},
			{X: 1, Y: 9, Name: "M", neighbourCount: 0},
			{X: 2, Y: 9, Name: "M", neighbourCount: 0},
		},
		"5_9": {
			{X: 4, Y: 8, Name: "S", neighbourCount: 0},
			{X: 4, Y: 9, Name: "S", neighbourCount: 0},
			{X: 5, Y: 9, Name: "S", neighbourCount: 0},
		},
		"2_3": {
			{X: 0, Y: 0, Name: "R", neighbourCount: 0},
			{X: 1, Y: 0, Name: "R", neighbourCount: 0},
			{X: 0, Y: 1, Name: "R", neighbourCount: 0},
			{X: 2, Y: 0, Name: "R", neighbourCount: 0},
			{X: 1, Y: 1, Name: "R", neighbourCount: 0},
			{X: 3, Y: 0, Name: "R", neighbourCount: 0},
			{X: 2, Y: 1, Name: "R", neighbourCount: 0},
			{X: 3, Y: 1, Name: "R", neighbourCount: 0},
			{X: 2, Y: 2, Name: "R", neighbourCount: 0},
			{X: 3, Y: 2, Name: "R", neighbourCount: 0},
			{X: 2, Y: 3, Name: "R", neighbourCount: 0},
			{X: 4, Y: 2, Name: "R", neighbourCount: 0},
		},
		"5_6": {
			{X: 3, Y: 3, Name: "C", neighbourCount: 0},
			{X: 4, Y: 3, Name: "C", neighbourCount: 0},
			{X: 5, Y: 2, Name: "C", neighbourCount: 0},
			{X: 6, Y: 0, Name: "C", neighbourCount: 0},
			{X: 7, Y: 0, Name: "C", neighbourCount: 0},
			{X: 6, Y: 1, Name: "C", neighbourCount: 0},
			{X: 7, Y: 1, Name: "C", neighbourCount: 0},
			{X: 6, Y: 2, Name: "C", neighbourCount: 0},
			{X: 8, Y: 1, Name: "C", neighbourCount: 0},
			{X: 5, Y: 3, Name: "C", neighbourCount: 0},
			{X: 4, Y: 4, Name: "C", neighbourCount: 0},
			{X: 4, Y: 5, Name: "C", neighbourCount: 0},
			{X: 5, Y: 5, Name: "C", neighbourCount: 0},
			{X: 5, Y: 6, Name: "C", neighbourCount: 0},
		},
		"3_9": {
			{X: 1, Y: 7, Name: "I", neighbourCount: 0},
			{X: 2, Y: 5, Name: "I", neighbourCount: 0},
			{X: 2, Y: 6, Name: "I", neighbourCount: 0},
			{X: 3, Y: 6, Name: "I", neighbourCount: 0},
			{X: 2, Y: 7, Name: "I", neighbourCount: 0},
			{X: 4, Y: 6, Name: "I", neighbourCount: 0},
			{X: 3, Y: 7, Name: "I", neighbourCount: 0},
			{X: 4, Y: 7, Name: "I", neighbourCount: 0},
			{X: 1, Y: 8, Name: "I", neighbourCount: 0},
			{X: 2, Y: 8, Name: "I", neighbourCount: 0},
			{X: 3, Y: 8, Name: "I", neighbourCount: 0},
			{X: 5, Y: 7, Name: "I", neighbourCount: 0},
			{X: 5, Y: 8, Name: "I", neighbourCount: 0},
			{X: 3, Y: 9, Name: "I", neighbourCount: 0},
		},
		"6_9": {
			{X: 5, Y: 4, Name: "J", neighbourCount: 0},
			{X: 6, Y: 3, Name: "J", neighbourCount: 0},
			{X: 6, Y: 4, Name: "J", neighbourCount: 0},
			{X: 6, Y: 5, Name: "J", neighbourCount: 0},
			{X: 7, Y: 5, Name: "J", neighbourCount: 0},
			{X: 6, Y: 6, Name: "J", neighbourCount: 0},
			{X: 7, Y: 6, Name: "J", neighbourCount: 0},
			{X: 6, Y: 7, Name: "J", neighbourCount: 0},
			{X: 7, Y: 7, Name: "J", neighbourCount: 0},
			{X: 6, Y: 8, Name: "J", neighbourCount: 0},
			{X: 6, Y: 9, Name: "J", neighbourCount: 0},
		},
	}

	result2 := findAreas(input2)
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
			X:              0,
			Y:              0,
			Name:           "A",
			neighbourCount: 0,
		},
		{
			X:              1,
			Y:              0,
			Name:           "A",
			neighbourCount: 0,
		},
		{
			X:              2,
			Y:              0,
			Name:           "A",
			neighbourCount: 0,
		},
		{
			X:              3,
			Y:              0,
			Name:           "A",
			neighbourCount: 0,
		},
		{
			X:              0,
			Y:              1,
			Name:           "B",
			neighbourCount: 0,
		},
		{
			X:              1,
			Y:              1,
			Name:           "B",
			neighbourCount: 0,
		},
		{
			X:              2,
			Y:              1,
			Name:           "C",
			neighbourCount: 0,
		},
		{
			X:              3,
			Y:              1,
			Name:           "D",
			neighbourCount: 0,
		},
		{
			X:              0,
			Y:              2,
			Name:           "B",
			neighbourCount: 0,
		},
		{
			X:              1,
			Y:              2,
			Name:           "B",
			neighbourCount: 0,
		},
		{
			X:              2,
			Y:              2,
			Name:           "C",
			neighbourCount: 0,
		},
		{
			X:              3,
			Y:              2,
			Name:           "C",
			neighbourCount: 0,
		},
		{
			X:              0,
			Y:              3,
			Name:           "E",
			neighbourCount: 0,
		},
		{
			X:              1,
			Y:              3,
			Name:           "E",
			neighbourCount: 0,
		},
		{
			X:              2,
			Y:              3,
			Name:           "E",
			neighbourCount: 0,
		},
		{
			X:              3,
			Y:              3,
			Name:           "C",
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
