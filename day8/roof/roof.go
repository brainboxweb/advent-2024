// Package roof describes the roof
package roof

import (
	"fmt"
)

// Point is a unique location
type Point struct {
	X int
	Y int
}

// Grid is were the antennae are located
type Grid struct {
	GridSize      int
	Antennas      map[string][]Point
	Pairs         [][]Point
	comprehensive bool
}

// NewGrid returns a new Grid object
func NewGrid(data [][]string) *Grid {
	gridSize := len(data[0])
	antennas := antennas(data)
	pairs := [][]Point{}
	for _, antenna := range antennas {
		ps := getPairs(antenna, nil)
		pairs = append(pairs, ps...)
	}

	return &Grid{GridSize: gridSize, Antennas: antennas, Pairs: pairs}
}

// AntinodesComprehensive is the advanced version of Antinodes
func (g *Grid) AntinodesComprehensive() map[string]Point {
	g.comprehensive = true
	return g.Antinodes()
}

// Antinodes are the antenna antinodes
func (g *Grid) Antinodes() map[string]Point {
	antis := []Point{}
	for _, pair := range g.Pairs {
		ans := g.antinodesForPair(pair)
		antis = append(antis, ans...)
	}

	// ---- dedupe
	antinodes := make(map[string]Point)
	for _, ant := range antis {
		key := fmt.Sprintf("%d-%d", ant.X, ant.Y)
		antinodes[key] = ant
	}
	return antinodes
}

func antennas(grid [][]string) map[string][]Point {
	points := make(map[string][]Point)
	for y := range len(grid[0]) { // collect everything that's not a dot
		for x := range grid {
			if grid[x][y] != "." {
				val := grid[x][y]
				pp := Point{x, y}
				points[val] = append(points[val], pp)
			}
		}
	}

	return points
}

func (g *Grid) antinodesForPair(pair []Point) []Point {
	if len(pair) != 2 {
		panic("not supported")
	}

	maximum := g.GridSize

	pnt1 := pair[0]
	pnt2 := pair[1]

	diffX := pnt1.X - pnt2.X
	diffY := pnt1.Y - pnt2.Y

	cands := []Point{} // candidates
	if g.comprehensive {
		for mult := range maximum {
			p1 := Point{pnt1.X + diffX*mult, pnt1.Y + diffY*mult}
			p2 := Point{pnt2.X - diffX*mult, pnt2.Y - diffY*mult}
			cands = append(cands, p1, p2)
		}
	} else {
		mult := 1
		p1 := Point{pnt1.X + diffX*mult, pnt1.Y + diffY*mult}
		p2 := Point{pnt2.X - diffX*mult, pnt2.Y - diffY*mult}
		cands = append(cands, p1, p2)
	}

	antinodes := []Point{}
	for _, c := range cands { // @todo rename
		if c.X < maximum && c.Y < maximum && c.X > -1 && c.Y > -1 {
			antinodes = append(antinodes, c)
		}
	}

	return antinodes
}

func getPairs(input []Point, pairs [][]Point) [][]Point {
	if len(input) < 2 {
		return pairs
	}
	first, others := input[0], input[1:]
	for _, item := range others {
		pair := []Point{}
		pair = append(pair, first)
		pair = append(pair, item)
		pairs = append(pairs, pair)
	}

	return getPairs(others, pairs)
}
