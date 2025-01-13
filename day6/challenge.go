package day6

import (
	"fmt"
)

func ChallengeOne(data [][]string) int {
	theMap := newMap(data)

	// spew.Dump(theMap)
	// panic("stiop")
	// return 7

	res := theMap.Move()

	return res
}

func ChallengeTwo(data [][]string) int {
	counter := 0

	for x := 0; x < len(data); x++ {
		for y := 0; y < len(data[0]); y++ {
			alteredData := data // one change at a time!!!
			if data[x][y] == "^" {
				continue
			}
			if data[x][y] == "#" {
				continue
			}
			alteredData[x][y] = "#" // add an obstruction
			theMap := newMap(alteredData)
			theMap.Move()
			if theMap.Loop {
				counter++
			}
		}
	}

	return counter
}

func newMap(data [][]string) *Map {
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

func (m *Map) findChar(char string) (int, int) {
	for x := 0; x < m.lenY-1; x++ {
		for y := 0; y < m.lenY-1; y++ {
			if m.grid[x][y] == char {
				return x, y
			}
		}
	}

	return 0, 0
}

type Point struct {
	X int
	Y int
}

type Map struct {
	grid      [][]string
	lenX      int
	lenY      int
	start     Point
	current   Point
	direction string
	visited   map[string]Point
	Loop      bool
}

func (m *Map) Move() int {

	for {
		stillMoving := m.moveInDirection(m.direction)
		if !stillMoving {
			break
		}
	}

	return len(m.visited)
}

func (m *Map) moveInDirection(dirn string) bool {
	fmt.Println("-------> moving ", dirn)

	stillMoving := true
	if dirn == "N" {
		// if m.current == m.start && len(m.visited) > 1 { // Think this is the only direction I have to worry about
		// 	fmt.Println("Passing go!")
		// 	// in a loop
		// 	m.Loop = true
		// 	return false // I.e., we're done here
		// }
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
	count := 0
	for {
		candidate = Point{candidate.X, candidate.Y - 1} // have incremented
		if candidate.Y < 0 {
			return false // stepping outside
		}
		if m.grid[candidate.X][candidate.Y] == "#" { // turn
			m.direction = "E" // inverted
			return true       // still moving
		}
		// take the step
		m.visited[fmt.Sprintf("%d,%d", candidate.X, candidate.Y)] = candidate
		m.current = candidate
		count++
	}
}

func (m *Map) moveSouth() bool { // S is "inverted"
	candidate := m.current
	count := 0
	for {
		candidate = Point{candidate.X, candidate.Y + 1} // have incremented
		if candidate.Y > m.lenY-1 {
			return false // stepping outside
		}
		if m.grid[candidate.X][candidate.Y] == "#" { // turn
			m.direction = "W" // inverted
			return true       // still moving
		}
		// take the step
		m.visited[fmt.Sprintf("%d,%d", candidate.X, candidate.Y)] = candidate
		m.current = candidate
		count++
	}
}

func (m *Map) moveEast() bool {
	candidate := m.current
	count := 0
	for {
		candidate = Point{candidate.X + 1, candidate.Y}
		if candidate.X > m.lenX-1 {
			return false // stepping outside
		}
		if m.grid[candidate.X][candidate.Y] == "#" { // turn
			m.direction = "S"
			return true // still moving
		}
		// take the step
		m.visited[fmt.Sprintf("%d,%d", candidate.X, candidate.Y)] = candidate
		m.current = candidate
		count++
	}
}

func (m *Map) moveWest() bool {
	candidate := m.current
	count := 0
	for {
		candidate = Point{candidate.X - 1, candidate.Y}
		if candidate.X < 0 {
			return false // stepping outside
		}
		if m.grid[candidate.X][candidate.Y] == "#" { // turn
			m.direction = "N"
			return true // still moving
		}
		// take the step
		m.visited[fmt.Sprintf("%d,%d", candidate.X, candidate.Y)] = candidate
		m.current = candidate
		count++
	}
}
