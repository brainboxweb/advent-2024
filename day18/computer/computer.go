// Package computer respresents the computer
package computer

import (
	"fmt"
	"slices"
)

// Computer is a collection of points and starting points
type Computer struct {
	length         int
	corrupted      []point
	StartingPoints []point
}

// New returns a new Computer instance
func New(length int) *Computer {
	comp := &Computer{length: length}
	return comp
}

// AddCorrupted adds corruption to the Computer
func (c *Computer) AddCorrupted(data [][]int, limit int) {
	for i, line := range data {
		if i >= limit {
			return
		}
		p := point{line[0], line[1]}
		c.corrupted = append(c.corrupted, p)
	}
}

type point struct {
	X int
	Y int
}

type pointLoaded struct {
	point
	cnt           int     // count
	visitedPoints []point // may not be needed
}

type move struct {
	X int
	Y int
}

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
		newMove := makeMove(pnt, mv)
		if !c.isPossible(newMove) {
			continue
		}
		if c.isCorrupted(newMove) {
			continue
		}
		ret = append(ret, newMove)
	}
	return ret
}

//revive:disable:cognitive-complexity

// Walk solves the maze and returns the move count
func (c *Computer) Walk() int { // step count
	start := point{0, 0}
	end := point{c.length - 1, c.length - 1}
	globalVisited := []point{start}
	startLoaded := pointLoaded{start, 0, []point{start}}
	winningCounts := []int{}
	winningMinimumCount := 1000000
	loadedPoints := []pointLoaded{startLoaded} // seed with just one
	for len(loadedPoints) > 0 {
		var theNewPoints []pointLoaded        // for the loop
		for _, loaded := range loadedPoints { // passed in
			// Shorter route?
			if len(loaded.visitedPoints) > winningMinimumCount {
				continue
			}
			candidates := c.getCandidates(loaded.point)
			for _, candi := range candidates {
				// Already visited?
				if slices.Contains(globalVisited, candi) {
					continue
				}
				// At end?
				if candi == end {
					winningCount := loaded.cnt + 1
					winningCounts = append(winningCounts, winningCount)
					if winningCount <= winningMinimumCount {
						winningMinimumCount = winningCount
					}
					continue
				}
				globalVisited = append(globalVisited, candi)
				newLoadedPoint := makeNewLoadedPoint(loaded, candi)
				theNewPoints = append(theNewPoints, newLoadedPoint)
			}
		}
		loadedPoints = theNewPoints
	}
	minWinCount := 100000
	for _, val := range winningCounts {
		if val < minWinCount {
			minWinCount = val
		}
	}
	return minWinCount
}

//revive:enable:cognitive-complexity

func makeNewLoadedPoint(source pointLoaded, target point) pointLoaded {
	ret := pointLoaded{}
	ret.point = target
	ret.cnt = source.cnt + 1
	newHistory := slices.Clone(source.visitedPoints)
	newHistory = append(newHistory, target)
	ret.visitedPoints = newHistory
	return ret
}

func (c *Computer) isPossible(pnt point) bool {
	if pnt.X < 0 || pnt.Y < 0 || pnt.X >= c.length || pnt.Y >= c.length {
		return false
	}
	return true
}

func (c *Computer) isCorrupted(pnt point) bool {
	for _, p := range c.corrupted {
		if p == pnt {
			return true
		}
	}
	return false
}

// Dump is for debugging
func (c *Computer) Dump() {
	for y := range c.length {
		line := ""
		for x := range c.length {
			candidate := "."
			if c.isCorrupted(point{x, y}) {
				candidate = "#"
			}
			line += candidate
		}
		fmt.Println(line)
	}
}

// DumpRoute is for debugging a single path
func (c *Computer) DumpRoute(hist []point) {
	for y := range c.length {
		line := ""
		for x := range c.length {
			candidate := "."
			if c.isCorrupted(point{x, y}) {
				candidate = "#"
			}

			if slices.Contains(hist, point{x, y}) {
				candidate = "0"
			}
			line += candidate
		}
		fmt.Println(line)
	}
}
