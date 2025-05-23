package maze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCandidates(t *testing.T) {
	tests := []struct {
		name     string
		input    moveVectorLoaded
		expected []moveVectorLoaded
	}{
		{
			"right",
			moveVectorLoaded{
				moveVector{
					point{5, 5},
					direction{1, 0},
				},
				7,
				[]moveVector{
					{
						point{5, 5},
						direction{1, 0},
					},
				},
			},
			[]moveVectorLoaded{
				{
					moveVector{
						point{6, 5}, // straight (right)
						direction{1, 0},
					},
					8,
					[]moveVector{
						{point{5, 5}, direction{1, 0}},
						{point{6, 5}, direction{1, 0}},
					},
				},
				{
					moveVector{
						point{5, 6}, // "left" (up)
						direction{0, 1},
					},
					1008,
					[]moveVector{
						{point{5, 5}, direction{1, 0}},
						{point{5, 6}, direction{0, 1}},
					},
				},
				{
					moveVector{
						point{5, 4}, // "right" (down)
						direction{0, -1},
					},
					1008,
					[]moveVector{
						{point{5, 5}, direction{1, 0}},
						{point{5, 4}, direction{0, -1}},
					},
				},
			},
		},
		{
			"up",
			moveVectorLoaded{
				moveVector{
					point{5, 5},
					direction{0, 1},
				},
				7,
				[]moveVector{
					{
						point{5, 5},
						direction{0, 1},
					},
				},
			},
			[]moveVectorLoaded{
				{
					moveVector{
						point{5, 6}, // straight (up)
						direction{0, 1},
					},
					8,
					[]moveVector{
						{point{5, 5}, direction{0, 1}},
						{point{5, 6}, direction{0, 1}},
					},
				},
				{
					moveVector{
						point{4, 5}, // "left" (left)
						direction{-1, 0},
					},
					1008,
					[]moveVector{
						{point{5, 5}, direction{0, 1}},
						{point{4, 5}, direction{-1, 0}},
					},
				},
				{
					moveVector{
						point{6, 5}, // "right" (right)
						direction{1, 0},
					},
					1008,
					[]moveVector{
						{point{5, 5}, direction{0, 1}},
						{point{6, 5}, direction{1, 0}},
					},
				},
			},
		},
		{
			"left",
			moveVectorLoaded{
				moveVector{
					point{5, 5},
					direction{-1, 0}, // left
				},
				7,
				[]moveVector{
					{
						point{5, 5},
						direction{-1, 0},
					},
				},
			},
			[]moveVectorLoaded{
				{
					moveVector{
						point{4, 5}, // straight (left)
						direction{-1, 0},
					},
					8,
					[]moveVector{
						{point{5, 5}, direction{-1, 0}},
						{point{4, 5}, direction{-1, 0}},
					},
				},
				{
					moveVector{
						point{5, 6}, // "right" (up)
						direction{0, 1},
					},
					1008,
					[]moveVector{
						{point{5, 5}, direction{-1, 0}},
						{point{5, 6}, direction{0, 1}},
					},
				},
				{
					moveVector{
						point{5, 4}, // "left" (down)
						direction{0, -1},
					},
					1008,
					[]moveVector{
						{point{5, 5}, direction{-1, 0}},
						{point{5, 4}, direction{0, -1}},
					},
				},
			},
		},
		{
			"down",
			moveVectorLoaded{
				moveVector{
					point{5, 5},
					direction{0, -1}, // down
				},
				7,
				[]moveVector{
					{
						point{5, 5},
						direction{0, -1},
					},
				},
			},
			[]moveVectorLoaded{
				{
					moveVector{
						point{5, 4}, // straight (down)
						direction{0, -1},
					},
					8,
					[]moveVector{
						{point{5, 5}, direction{0, -1}},
						{point{5, 4}, direction{0, -1}},
					},
				},
				{
					moveVector{
						point{4, 5}, // "right" (left)
						direction{-1, 0},
					},
					1008,
					[]moveVector{
						{point{5, 5}, direction{0, -1}},
						{point{4, 5}, direction{-1, 0}},
					},
				},
				{
					moveVector{
						point{6, 5}, // "left" (right)
						direction{1, 0},
					},
					1008,
					[]moveVector{
						{point{5, 5}, direction{0, -1}},
						{point{6, 5}, direction{1, 0}},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveVec := tt.input
			result := moveVec.getCandidates()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestBuildCandidate(t *testing.T) {
	tests := []struct {
		name      string
		input     moveVectorLoaded
		newVector moveVector
		addCost   int
		expected  moveVectorLoaded
	}{
		{
			name: "Not sure",
			input: moveVectorLoaded{
				moveVector{point{5, 5}, direction{1, 0}},
				7,
				[]moveVector{
					{point{5, 5}, direction{1, 0}},
				},
			},
			newVector: moveVector{point{6, 5}, direction{1, 0}},
			addCost:   1,
			expected: moveVectorLoaded{
				moveVector{
					point{6, 5},
					direction{1, 0},
				},
				8,
				[]moveVector{
					{point{5, 5}, direction{1, 0}},
					{point{6, 5}, direction{1, 0}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := buildNewCandidate(tt.input, tt.newVector, tt.addCost)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetTargets(t *testing.T) {
	data := []string{
		"###############",
		"#.......#....E#",
		"#.#.###.#.###.#",
		"#.....#.#...#.#",
		"#.###.#####.#.#",
		"#.#.#.......#.#",
		"#.#.#####.###.#",
		"#...........#.#",
		"###.#.#####.#.#",
		"#...#.....#.#.#",
		"#.#.#.###.#.#.#",
		"#.....#...#.#.#",
		"#.###.#.#.#.#.#",
		"#S..#.....#...#",
		"###############",
	}

	mz := New(data)
	startVector := moveVectorLoaded{
		moveVector: moveVector{
			point: mz.startPoint, direction: direction{1, 0},
		},
		cost: 0,
		history: []moveVector{
			{point: mz.startPoint, direction: direction{1, 0}},
		},
	}
	targets := mz.getTargets(startVector)
	expected := []moveVectorLoaded{
		{
			moveVector{
				point{ // move straight (right)
					x: 2,
					y: 13,
				},
				direction{
					x: 1,
					y: 0,
				},
			},
			1,
			[]moveVector{
				{point: mz.startPoint, direction: direction{1, 0}},
				{point: point{2, 13}, direction: direction{1, 0}},
			},
		},
		{
			moveVector{
				point{ // move "left" (up)
					x: 1,
					y: 12,
				},
				direction{
					x: 0,
					y: -1,
				},
			},
			1001,
			[]moveVector{
				{point: mz.startPoint, direction: direction{1, 0}},
				{point: point{1, 12}, direction: direction{0, -1}},
			},
		},
		// down not possible
	}
	assert.Equal(t, expected, targets)
}
