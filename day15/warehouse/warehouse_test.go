package warehouse_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day15/warehouse"
	"github.com/stretchr/testify/assert"
)

func TestMoves(t *testing.T) {
	data := [][]string{
		{
			string("#"),
			string("#"),
			string("#"),
			string("#"),
			string("#"),
			string("#"),
			string("#"),
			string("#"),
		},
		{
			string("#"),
			string("."),
			string("."),
			string("O"),
			string("."),
			string("O"),
			string("."),
			string("#"),
		},
		{
			string("#"),
			string("#"),
			string("@"),
			string("."),
			string("O"),
			string("."),
			string("."),
			string("#"),
		},
		{
			string("#"),
			string("."),
			string("."),
			string("."),
			string("O"),
			string("."),
			string("."),
			string("#"),
		},
		{
			string("#"),
			string("."),
			string("#"),
			string("."),
			string("O"),
			string("."),
			string("."),
			string("#"),
		},
		{
			string("#"),
			string("."),
			string("."),
			string("."),
			string("O"),
			string("."),
			string("."),
			string("#"),
		},
		{
			string("#"),
			string("."),
			string("."),
			string("."),
			string("."),
			string("."),
			string("."),
			string("#"),
		},
		{
			string("#"),
			string("#"),
			string("#"),
			string("#"),
			string("#"),
			string("#"),
			string("#"),
			string("#"),
		},
	}
	w := warehouse.New(data)

	// <^^>>>vv<v>>v<<
	w.Move("<")
	w.Move("^")
	w.Move("^")
	w.Move(">")
	w.Move(">")
	w.Move(">")
	w.Move("v")
	w.Move("v")
	w.Move("<")
	w.Move("v")
	w.Move(">")
	w.Move(">")
	w.Move("v")
	w.Move("<")
	w.Move("<")

	result := w.GPS()
	assert.Equal(t, 2028, result)
}
