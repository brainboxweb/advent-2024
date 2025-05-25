package maze_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day16/maze"
	"github.com/stretchr/testify/assert"
)

func TestChallenge(t *testing.T) {
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

	tests := []struct {
		data          []string
		expectedCost  int // Part 1
		expectedCount int // Part 2
	}{
		{
			data,
			7036,
			45,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			mz := maze.New(data)
			// return
			res := mz.Solve()
			assert.Equal(t, tt.expectedCost, res.MinCost)
			assert.Equal(t, tt.expectedCount, res.WinningPathCount)
		})
	}
}
