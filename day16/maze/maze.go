// Package maze represents the maze
package maze

import (
	"fmt"
	"strings"
)

// New return a new warehouse
func New(data []string) *Maze {
	moveCost := make(map[moveVector]int)
	mz := &Maze{moveVectorCost: moveCost}
	mz.populateAllowed(len(data))
	for y, line := range data {
		items := strings.Split(line, "")
		for x, item := range items {
			switch item {
			case "#":
				mz.allowed[y][x] = false
			case "S":
				mz.startPoint = point{x, y}
			case "E":
				mz.endPoint = point{x, y}
			}
		}
	}
	return mz
}

// Solve solves the maze with minmium cost
func (mz *Maze) Solve() int {
	startVector := moveVector{
		point:     mz.startPoint,
		direction: direction{x: 1, y: 0}, // Start moving right
	}
	return mz.step(startVector)
}

// Maze represents the maze
type Maze struct {
	allowed        [][]bool
	startPoint     point
	endPoint       point
	moveVectorCost map[moveVector]int // moveVector, lowest cost
}

func (mz *Maze) getCost(mv moveVector) int {
	key := moveVector{mv.point, mv.direction, 0} // NB
	val, ok := mz.moveVectorCost[key]
	if !ok {
		return 100000000
	}
	return val
}

func (mz *Maze) setCost(mv moveVector) {
	key := moveVector{mv.point, mv.direction, 0} // NB
	mz.moveVectorCost[key] = mv.cost
}

//revive:disable:cognitive-complexity
func (mz *Maze) step(startMove moveVector) int {
	startMoves := []moveVector{}
	startMoves = append(startMoves, startMove)
	endPointRoutes := []moveVector{}
	for len(startMoves) > 0 {
		// if len(startMoves) == 0 { // End condition
		// 	break
		// }
		newStartMoves := []moveVector{}
		for _, startMove := range startMoves {
			targets := mz.getTargets(startMove)
			for _, target := range targets {
				if target.point == mz.endPoint {
					endPointRoutes = append(endPointRoutes, target)
					break
				}
				newStartMoves = append(newStartMoves, target)
			}
		}
		// Prepare to loop
		startMoves = newStartMoves
	}
	minCost := 10000000
	for _, ep := range endPointRoutes {
		if ep.cost < minCost {
			minCost = ep.cost
		}
	}
	return minCost
}

//revive:enable:cognitive-complexity

func (mz *Maze) isAllowed(candidate point) bool {
	return mz.allowed[candidate.y][candidate.x]
}

func (mz *Maze) getTargets(mv moveVector) []moveVector {
	targetMoveVectors := []moveVector{}
	var left direction
	var right direction
	switch {
	case mv.direction.y == 0:
		left = direction{0, 1}
		right = direction{0, -1}
	case mv.direction.x == 0:
		left = direction{-1, 0}
		right = direction{1, 0}
	}
	strPoint := point{mv.point.x + mv.direction.x, mv.point.y + mv.direction.y}
	leftPoint := point{mv.point.x + left.x, mv.point.y + left.y}
	rightPoint := point{mv.point.x + right.x, mv.point.y + right.y}

	// Allowed?
	if mz.isAllowed(strPoint) {
		targetMoveVectors = append(targetMoveVectors,
			moveVector{strPoint, mv.direction, mv.cost + 1})
	}
	if mz.isAllowed(leftPoint) {
		targetMoveVectors = append(targetMoveVectors,
			moveVector{leftPoint, left, mv.cost + 1001})
	}
	if mz.isAllowed(rightPoint) {
		targetMoveVectors = append(targetMoveVectors,
			moveVector{rightPoint, right, mv.cost + 1001})
	}

	// Cheapest?
	cheapMoveVectors := []moveVector{}
	for _, mv := range targetMoveVectors {
		lowestCost := mz.getCost(mv)
		if mv.cost > lowestCost {
			continue // drop this route
		}
		mz.setCost(mv)                                  // set/update the cost
		cheapMoveVectors = append(cheapMoveVectors, mv) // add
	}

	return cheapMoveVectors
}

type point struct {
	x int
	y int
}

type direction struct {
	x int
	y int
}

type moveVector struct {
	point
	direction
	cost int
}

func (mz *Maze) populateAllowed(limit int) {
	sliceOfSlices := make([][]bool, limit)
	for i := range limit {
		sliceOfSlices[i] = make([]bool, limit)
	}
	for y := range limit {
		for x := range limit {
			sliceOfSlices[y][x] = true
		}
	}
	mz.allowed = sliceOfSlices
}

// Dump is for debugging
func (mz *Maze) Dump() {
	maxLen := len(mz.allowed)
	for y := range maxLen {
		line := ""
		for x := range maxLen {
			candidate := "#"
			if mz.allowed[y][x] {
				candidate = "."
			}
			switch {
			case mz.startPoint == point{x, y}:
				candidate = ("S")
			case mz.endPoint == point{x, y}:
				candidate = ("E")
			}
			line += candidate
		}
		fmt.Println(line)
	}
}
