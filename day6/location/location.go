// Package location relates to the map
package location

import (
	"fmt"
)

// NewMap returns a new Map
func NewMap(data [][]string) *Map {
	visited := make(map[string]Point)
	m := Map{
		grid:      data,
		lenX:      len(data[0]),
		lenY:      len(data),
		direction: "N",
		visited:   visited,
	}
	x, y := m.findChar("^")
	start := Point{x, y}
	m.start = start
	m.current = start
	m.visited[fmt.Sprintf("%d,%d", x, y)] = start

	return &m
}

// Point is an x, y location
type Point struct {
	X int
	Y int
}

// Map is area of terrain
type Map struct {
	grid         [][]string
	lenX         int
	lenY         int
	start        Point
	current      Point
	direction    string
	visited      map[string]Point
	visitedCount int
}

// Move runs a series of moves
func (m *Map) Move() int {
	for {
		stillMoving := m.moveInDirection(m.direction)
		if !stillMoving {
			break
		}
	}

	return len(m.visited)
}

// IsEndlessLoop tests for leakage
func (m *Map) IsEndlessLoop() bool {
	for {
		if m.visitedCount > len(m.grid)*50 { // endless loop
			return true
		}
		if !m.moveInDirection(m.direction) {
			return false // leaked out of the grid
		}
	}
}

func (m *Map) findChar(char string) (x, y int) {
	for x := range m.lenY {
		for y := range m.lenY {
			if m.grid[x][y] == char {
				return x, y
			}
		}
	}

	return 0, 0
}

func (m *Map) moveInDirection(dirn string) bool {
	stillMoving := true
	if dirn == "N" {
		stillMoving = m.moveNorth()
	}
	if dirn == "S" {
		stillMoving = m.moveSouth()
	}
	if dirn == "E" {
		stillMoving = m.moveEast()
	}
	if dirn == "W" {
		stillMoving = m.moveWest()
	}

	return stillMoving
}

func (m *Map) moveNorth() bool {
	candidate := m.current
	for {
		candidate = Point{candidate.X, candidate.Y - 1} // have incremented
		if candidate.Y < 0 {
			return false // stepping outside
		}
		if m.grid[candidate.X][candidate.Y] == "#" { // turn
			m.direction = "E" // inverted
			return true       // still moving
		}
		m.takeStep(candidate)
	}
}

func (m *Map) moveSouth() bool { // S is "inverted"
	candidate := m.current
	for {
		candidate = Point{candidate.X, candidate.Y + 1} // have incremented
		if candidate.Y > m.lenY-1 {
			return false // stepping outside
		}
		if m.grid[candidate.X][candidate.Y] == "#" { // turn
			m.direction = "W" // inverted
			return true       // still moving
		}
		m.takeStep(candidate)
	}
}

func (m *Map) moveEast() bool {
	candidate := m.current
	for {
		candidate = Point{candidate.X + 1, candidate.Y}
		if candidate.X > m.lenX-1 {
			return false // stepping outside
		}
		if m.grid[candidate.X][candidate.Y] == "#" { // turn
			m.direction = "S"
			return true // still moving
		}
		m.takeStep(candidate)
	}
}

func (m *Map) moveWest() bool {
	candidate := m.current
	for {
		candidate = Point{candidate.X - 1, candidate.Y}
		if candidate.X < 0 {
			return false // stepping outside
		}
		if m.grid[candidate.X][candidate.Y] == "#" { // turn
			m.direction = "N"
			return true // still moving
		}
		m.takeStep(candidate)
	}
}

func (m *Map) takeStep(candidate Point) {
	m.visitedCount++
	m.visited[fmt.Sprintf("%d,%d", candidate.X, candidate.Y)] = candidate
	m.current = candidate
}
