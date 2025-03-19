// Package garden represents the garden plots
package garden

import (
	"strconv"
)

// NewPlot returns a new Plot
func NewPlot(grid [][]string) Plot {
	grouper := findAreas(grid)
	areas := doGouping(grouper)

	return Plot{grid: grid, Areas: areas}
}

func doGouping(grouper map[string][]point) [][]point {
	areas := [][]point{}
	for _, points := range grouper {
		if points == nil {
			for id, group := range grouper {
				if group == nil {
					// fmt.Println("OOOOPS. Nil id = ", id)
					panic("not expdted" + id)
				}
			}
		}
		areas = append(areas, points)
	}
	return areas
}

type point struct {
	X              int
	Y              int
	Name           string
	neighbourCount int
}

// Plot is a garden plot
type Plot struct {
	Areas [][]point // make internal later
	grid  [][]string
	// lenX         int
	// lenY         int
	// start        point
	// current      point
	// direction    string
	// visited      map[string]point
	// visitedCount int
}

func scan(data [][]string) []point {
	layout := []point{}
	for y, line := range data {
		for x, pnt := range line {
			layout = append(layout, point{Name: pnt, X: x, Y: y})
		}
	}
	return layout
}

func makeIndex(a, b int) string {
	return strconv.Itoa(a) + "_" + strconv.Itoa(b)
}

func findAreas(grid [][]string) map[string][]point {
	grouper := make(map[string][]point)
	for y, line := range grid {
		for x, name := range line {
			id := makeIndex(x, y)
			grouper[id] = append(grouper[id], point{X: x, Y: y, Name: name})
		}
	}
	for y, line := range grid {
		for x, name := range line {
			currentPoint := point{X: x, Y: y, Name: name}
			neighbours := getMatchingNeighbours(currentPoint, grid)
			if len(neighbours) == 0 {
				continue
			}
			grouper = mergeNeighbours(grouper, currentPoint, neighbours)
		}
	}

	return grouper
}

func mergeNeighbours(
	grouper map[string][]point,
	currentPoint point,
	neighbours []point,
) map[string][]point {
	currentID := makeIndex(currentPoint.X, currentPoint.Y)
	targetIDs := getNeighbourIDs(neighbours, grouper)
	for _, targetID := range targetIDs {
		if targetID == currentID {
			continue
		}
		grouper[currentID] = append(grouper[currentID], grouper[targetID]...)
		delete(grouper, targetID)
	}

	return grouper
}

func getNeighbourIDs(neighbours []point, grouper map[string][]point) []string {
	targetIDs := []string{}
	for _, point := range neighbours {
		for groupID, grouperPoints := range grouper {
			for _, grouperPoint := range grouperPoints {
				if point == grouperPoint { // found a matching point
					targetIDs = append(targetIDs, groupID)
				}
			}
		}
	}
	return targetIDs
}

//revive:disable:cognitive-complexity
//revive:disable:cyclomatic

func getMatchingNeighbours(pnt point, grid [][]string) []point {
	ret := []point{}
	x := pnt.X
	y := pnt.Y
	label := grid[y][x]
	var compare string

	// up
	yy := y - 1
	xx := x
	if xx < len(grid) && xx > -1 && yy < len(grid) && yy > -1 {
		compare = grid[yy][xx]
		if label == compare {
			ret = append(ret, point{X: xx, Y: yy, Name: label})
		}
	}

	// left
	yy = y
	xx = x - 1
	if xx < len(grid) && xx > -1 && yy < len(grid) && yy > -1 {
		compare = grid[yy][xx]
		if label == compare {
			ret = append(ret, point{X: xx, Y: yy, Name: label})
		}
	}

	// right
	yy = y
	xx = x + 1
	if xx < len(grid) && xx > -1 && yy < len(grid) && yy > -1 {
		compare = grid[yy][xx]
		if label == compare {
			ret = append(ret, point{X: xx, Y: yy, Name: label})
		}
	}

	// down
	yy = y + 1
	xx = x
	if xx < len(grid) && xx > -1 && yy < len(grid) && yy > -1 {
		compare = grid[yy][xx]
		if label == compare {
			ret = append(ret, point{X: xx, Y: yy, Name: label})
		}
	}

	return ret
}

//revive:enable:cognitive-complexity
//revive:enable:cyclomatic

// func isNeighbour(a, b Point) bool {
// 	if a.X == b.X+1 && a.Y == b.Y || a.X == b.X-1 && a.Y == b.Y ||
// 		a.Y == b.Y+1 && a.X == b.X || a.Y == b.Y-1 && a.X == b.X {
// 		return true
// 	}
// 	return false
// }

// func (pl *Plot) Neighbours() {
// 	for name, pnts := range pl.thing {
// 		if len(pnts) == 0 { // No neighbours
// 			continue
// 		}

// 		// Need to compare all the points in the plot
// 		for index, pnt := range pnts { // Could worl!
// 			neighbourCount := getNeighbourCount(pnt, pnts)
// 			pnts[index] = Point{pnt.X, pnt.Y, neighbourCount}
// 		}

// 		pl.thing[name] = pnts

// 	}

// }

// Try to get
// func (pl *Plot) getPerimeter()  int { // rename

// 	data := pl.grid
// 	maxi := len(data)

// 	fmt.Println("max is", maxi)

// 	// for letter := range pl.thing {
// 	// scan letter
// 	// override for now
// 	// inBlock := false

// 	for letter := range pl.thing {
// 		return pl.Perimeter( letter)

// 		// 	// AAAA
// 		// 	// BBCD
// 		// 	// BBCC
// 		// 	// EEEC
// 	}

// }

// TotalCost is the total cost of fencing
func (pl Plot) TotalCost() int {
	totalCost := 0
	for _, area := range pl.Areas { // Rename to regions
		perim := pl.Perimeter(area)
		cost := len(area) * perim
		totalCost += cost
	}

	return totalCost
}

func inRegion(x, y int, region []point) bool {
	for _, pnt := range region {
		if pnt.X == x && pnt.Y == y {
			return true
		}
	}

	return false
}

//revive:disable:cognitive-complexity

// Perimeter is the perimeter of a specific region
func (pl Plot) Perimeter(region []point) int {
	data := pl.grid
	maxi := len(data)
	edgeCount := 0
	// ---> or <---
	for y := range maxi {
		inBlock := false
		for x := range maxi { // update me!
			// count edges
			if inRegion(x, y, region) {
				// if data[x][y] == letter {
				if !inBlock {
					edgeCount++
					inBlock = true
					continue
				}
				continue
			}
			inBlock = false
		}
	}
	// up or down
	for x := range maxi {
		inBlock := false
		for y := range maxi {
			// count edges
			if inRegion(x, y, region) {
				// if data[x][y] == letter {
				if !inBlock {
					edgeCount++
					inBlock = true
					continue
				}
				continue
			}
			inBlock = false
		}
	}
	perim := edgeCount * 2

	return perim
}

//revive:enable:cognitive-complexity

// func (pl *Plot) getContiguous(x, y int) int {

// 	// grid := pl.grid

// 	// forget diagonals

// 	// AAAA
// 	// BBCD
// 	// BBCC
// 	// EEEC

// 	count := 0
// 	// for _, pnt := range pnts {
// 	// 	if pnt == myPnt {
// 	// 		continue
// 	// 	}
// 	// 	if neighbour(myPnt, pnt) {
// 	// 		count++
// 	// 	}
// 	// }

// 	return count
// }

// func (pl *Plot) Area(name string) int {
// 	area := len(pl.thing[name])
// 	return len(pl.)
// 	fmt.Println(area)

// 	return area
// }

// func (pl *Plot) Perimeter(name string) int {
// 	// spew.Dump("\n====\nthe is... ", name, pl.thing[name])

// 	perim := 0
// 	// perim := len(pl.thing)*4
// 	// fmt.Println("Staring with... ", perim)

// 	for _, pnt := range pl.thing[name] {
// 		perim += 4 - pnt.neighbourCount
// 	}

// 	return perim
// }

// func getNeighbourCount(myPnt Point, pnts []Point) int {
// 	count := 0
// 	for _, pnt := range pnts {
// 		if pnt == myPnt {
// 			continue
// 		}
// 		if neighbour(myPnt, pnt) {
// 			count++
// 		}
// 	}

// 	return count
// }

// func neighbour(a, b Point) bool {
// 	if a.X == b.X+1 && a.Y == b.Y || a.X == b.X-1 && a.Y == b.Y ||
// 		a.Y == b.Y+1 && a.X == b.X || a.Y == b.Y-1 && a.X == b.X {
// 		return true
// 	}

// 	return false
// }

// func (m *Map) Move() int {
// 	for {
// 		stillMoving := m.moveInDirection(m.direction)
// 		if !stillMoving {
// 			break
// 		}
// 	}

// 	return len(m.visited)
// }

// func (m *Map) IsEndlessLoop() bool {
// 	for {
// 		if m.visitedCount > len(m.grid)*50 { // endless loop
// 			return true
// 		}
// 		if !m.moveInDirection(m.direction) {
// 			return false // leaked out of the grid
// 		}
// 	}
// }

// func (m *Map) moveInDirection(dirn string) bool {
// 	stillMoving := true
// 	if dirn == "N" {
// 		stillMoving = m.moveNorth()
// 	}
// 	if dirn == "S" {
// 		stillMoving = m.moveSouth()
// 	}
// 	if dirn == "E" {
// 		stillMoving = m.moveEast()
// 	}
// 	if dirn == "W" {
// 		stillMoving = m.moveWest()
// 	}

// 	return stillMoving
// }

// func (m *Map) moveNorth() bool {
// 	candidate := m.current
// 	for {
// 		candidate = Point{candidate.X, candidate.Y - 1} // have incremented
// 		if candidate.Y < 0 {
// 			return false // stepping outside
// 		}
// 		if m.grid[candidate.X][candidate.Y] == "#" { // turn
// 			m.direction = "E" // inverted
// 			return true       // still moving
// 		}
// 		m.takeStep(candidate)
// 	}
// }

// func (m *Map) moveSouth() bool { // S is "inverted"
// 	candidate := m.current
// 	for {
// 		candidate = Point{candidate.X, candidate.Y + 1} // have incremented
// 		if candidate.Y > m.lenY-1 {
// 			return false // stepping outside
// 		}
// 		if m.grid[candidate.X][candidate.Y] == "#" { // turn
// 			m.direction = "W" // inverted
// 			return true       // still moving
// 		}
// 		m.takeStep(candidate)
// 	}
// }

// func (m *Map) moveEast() bool {
// 	candidate := m.current
// 	for {
// 		candidate = Point{candidate.X + 1, candidate.Y}
// 		if candidate.X > m.lenX-1 {
// 			return false // stepping outside
// 		}
// 		if m.grid[candidate.X][candidate.Y] == "#" { // turn
// 			m.direction = "S"
// 			return true // still moving
// 		}
// 		m.takeStep(candidate)
// 	}
// }

// func (m *Map) moveWest() bool {
// 	candidate := m.current
// 	for {
// 		candidate = Point{candidate.X - 1, candidate.Y}
// 		if candidate.X < 0 {
// 			return false // stepping outside
// 		}
// 		if m.grid[candidate.X][candidate.Y] == "#" { // turn
// 			m.direction = "N"
// 			return true // still moving
// 		}
// 		m.takeStep(candidate)
// 	}
// }

// func (m *Map) takeStep(candidate Point) {
// 	m.visitedCount++
// 	m.visited[fmt.Sprintf("%d,%d", candidate.X, candidate.Y)] = candidate
// 	m.current = candidate
// }
