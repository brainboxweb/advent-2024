// Package computer respresents the computer
package computer

import (
	"slices"
)

// New returns a new Computer instance
func New(length int) *Computer {
	comp := &Computer{length: length}
	return comp
}

// Computer is a collection of points and starting points
type Computer struct {
	length         int
	corrupted      []point
	StartingPoints []point
}

// AddCorruptedByte adds one corrupted byte to the Computer
func (c *Computer) AddCorruptedByte(data []int) {
	p := point{data[0], data[1]}
	c.corrupted = append(c.corrupted, p)
}

type point struct {
	X int
	Y int
}

type pointLoaded struct {
	point
	cnt           int // count
	visitedPoints []point
}

type move struct {
	X int
	Y int
}

//revive:disable:cognitive-complexity
//revive:disable:cyclomatic

// Walk solves the maze and returns the move count
func (c *Computer) Walk() (bool, int) { // ok, step count
	start := point{0, 0}
	end := point{c.length - 1, c.length - 1}
	visited := []point{start}
	winningCounts := []int{}
	winningMinimumCount := 1000000
	startLoaded := pointLoaded{start, 0, []point{start}}
	loadedPoints := []pointLoaded{startLoaded}
	for len(loadedPoints) > 0 {
		var theNewPoints []pointLoaded
		for _, loaded := range loadedPoints {
			// Shorter route?
			if len(loaded.visitedPoints) > winningMinimumCount {
				continue
			}
			candidates := c.getCandidates(loaded.point)
			for _, candidate := range candidates {
				// Already visited?
				if slices.Contains(visited, candidate) {
					continue
				}
				// Success?
				if candidate == end {
					winningCount := loaded.cnt + 1
					winningCounts = append(winningCounts, winningCount)
					if winningCount <= winningMinimumCount {
						winningMinimumCount = winningCount
					}
					continue
				}
				visited = append(visited, candidate)
				newLoadedPoint := makeNewLoadedPoint(loaded, candidate)
				theNewPoints = append(theNewPoints, newLoadedPoint)
			}
		}
		loadedPoints = theNewPoints
	}
	if len(winningCounts) == 0 {
		return false, 0
	}
	minWinCount := 100000
	for _, val := range winningCounts {
		if val < minWinCount {
			minWinCount = val
		}
	}
	return true, minWinCount
}

//revive:enable:cognitive-complexity
//revive:enable:cyclomatic

func makeMove(pnt point, mv move) point {
	return point{pnt.X + mv.X, pnt.Y + mv.Y}
}

func (c *Computer) getCandidates(pnt point) []point {
	var ret []point
	moves := []move{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	for _, mv := range moves {
		newPoint := makeMove(pnt, mv)
		if !c.isPossible(newPoint) {
			continue
		}
		if slices.Contains(c.corrupted, newPoint) {
			continue
		}
		ret = append(ret, newPoint)
	}
	return ret
}

func makeNewLoadedPoint(source pointLoaded, target point) pointLoaded {
	newVisited := slices.Clone(source.visitedPoints)
	newVisited = append(newVisited, target)
	return pointLoaded{
		point:         target,
		cnt:           source.cnt + 1,
		visitedPoints: newVisited,
	}
}

func (c *Computer) isPossible(pnt point) bool {
	if pnt.X < 0 || pnt.Y < 0 || pnt.X >= c.length || pnt.Y >= c.length {
		return false
	}
	return true
}

// // Dump is for debugging
// func (c *Computer) Dump() {
// 	for y := range c.length {
// 		line := ""
// 		for x := range c.length {
// 			candidate := "."
// 			if c.isCorrupted(point{x, y}) {
// 				candidate = "#"
// 			}
// 			line += candidate
// 		}
// 		fmt.Println(line)
// 	}
// }

// // DumpRoute is for debugging a single path
// func (c *Computer) DumpRoute(hist []point) {
// 	for y := range c.length {
// 		line := ""
// 		for x := range c.length {
// 			candidate := "."
// 			if c.isCorrupted(point{x, y}) {
// 				candidate = "#"
// 			}

// 			if slices.Contains(hist, point{x, y}) {
// 				candidate = "0"
// 			}
// 			line += candidate
// 		}
// 		fmt.Println(line)
// 	}
// }
