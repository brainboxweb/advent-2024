package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ([]string) (len=10 cap=16) {
//  (string) (len=8) "########",
//  (string) (len=8) "#..O.O.#",
//  (string) (len=8) "##@.O..#",
//  (string) (len=8) "#...O..#",
//  (string) (len=8) "#.#.O..#",
//  (string) (len=8) "#...O..#",
//  (string) (len=8) "#......#",
//  (string) (len=8) "########",
//  (string) "",
//  (string) (len=15) "<^^>>>vv<v>>v<<"
// }

func TestParseMap(t *testing.T) {
	tests := []struct {
		input    []string
		expected [][]string
	}{
		{
			[]string{
				"########",
				"#..O.O.#",
				"##@.O..#",
				"#...O..#",
				"#.#.O..#",
				"#...O..#",
				"#......#",
				"########",
			},
			[][]string{
				{"#", "#", "#", "#", "#", "#", "#", "#"},
				{"#", ".", ".", "O", ".", "O", ".", "#"},
				{"#", "#", "@", ".", "O", ".", ".", "#"},
				{"#", ".", ".", ".", "O", ".", ".", "#"},
				{"#", ".", "#", ".", "O", ".", ".", "#"},
				{"#", ".", ".", ".", "O", ".", ".", "#"},
				{"#", ".", ".", ".", ".", ".", ".", "#"},
				{"#", "#", "#", "#", "#", "#", "#", "#"},
			},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			result, _ := parse(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
