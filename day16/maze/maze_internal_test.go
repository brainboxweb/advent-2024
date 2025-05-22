package maze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	startVector := moveVector{mz.startPoint, direction{1, 0}, 0}
	targets := mz.getTargets(startVector) // seems to work!
	expected := []moveVector{
		{
			point{ // move straight (right)
				x: 2,
				y: 13,
			},
			direction{
				x: 1,
				y: 0,
			},
			1,
		},
		{
			point{ // move "left" (up)
				x: 1,
				y: 12,
			},
			direction{
				x: 0,
				y: -1,
			},
			1001,
		},
		// down not possible
	}
	assert.Equal(t, expected, targets)
}
