// Package maze represents the maze
package maze

import (
	"slices"
	"strings"
)

const (
	stepCost = 1
	turnCost = 1000
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

// Maze represents the maze
type Maze struct {
	allowed        [][]bool
	startPoint     point
	endPoint       point
	moveVectorCost map[moveVector]int
}

// Solve solves the maze with minmium cost
func (mz *Maze) Solve() Result {
	startPoint := mz.startPoint
	startDirection := direction{x: 1, y: 0} // Start moving right
	startVectorSimple := moveVector{startPoint, startDirection}
	startVector := moveVectorLoaded{
		moveVector: moveVector{
			point:     startPoint,
			direction: startDirection,
		},
		cost:    0,
		history: []moveVector{startVectorSimple},
	}
	return mz.step(startVector)
}

func (mz *Maze) getCost(mv moveVectorLoaded) int {
	key := moveVector{mv.point, mv.direction} // NB
	val, ok := mz.moveVectorCost[key]
	if !ok {
		return 100000000
	}
	return val
}

func (mz *Maze) setCost(mv moveVectorLoaded) {
	key := moveVector{mv.point, mv.direction} // NB
	mz.moveVectorCost[key] = mv.cost
}

//revive:disable:cognitive-complexity
func (mz *Maze) step(startMove moveVectorLoaded) Result {
	startMoves := []moveVectorLoaded{}
	startMoves = append(startMoves, startMove)
	successVectors := []moveVectorLoaded{}
	for len(startMoves) > 0 { // end condition
		newStartMoves := []moveVectorLoaded{}
		for _, startMove := range startMoves {
			targets := mz.getTargets(startMove)
			for _, target := range targets {
				if target.point == mz.endPoint {
					successVectors = append(successVectors, target)
					break
				}
				newStartMoves = append(newStartMoves, target)
			}
		}
		// Prepare to loop
		startMoves = newStartMoves
	}
	minCost := minCost(successVectors)
	winningPathCount := winningPathCount(successVectors, minCost)

	return Result{minCost, winningPathCount}
}

// Result includes MinCost and the WinningPathCount
type Result struct {
	MinCost          int
	WinningPathCount int
}

func minCost(successVectors []moveVectorLoaded) int {
	minCost := 10000000
	for _, ep := range successVectors {
		if ep.cost < minCost {
			minCost = ep.cost
		}
	}
	return minCost
}

func winningPathCount(successVectors []moveVectorLoaded, minCost int) int {
	winningPointSequences := [][]moveVector{}
	for _, ep := range successVectors {
		if ep.cost == minCost {
			winningPointSequences = append(winningPointSequences, ep.history)
		}
	}

	unique := make(map[point]bool)
	for _, sequ := range winningPointSequences {
		for _, vec := range sequ {
			unique[vec.point] = true
		}
	}

	return len(unique)
}

//revive:enable:cognitive-complexity

func (mz *Maze) isAllowed(candidate point) bool {
	return mz.allowed[candidate.y][candidate.x]
}

// think I need refererences!!!!!
func (mz *Maze) getTargets(mv moveVectorLoaded) []moveVectorLoaded {
	candidates := mv.getCandidates()

	// valid? Cheapest?
	cheapMoveVectors := []moveVectorLoaded{}
	for _, candidate := range candidates {
		if !mz.isAllowed(candidate.point) {
			continue
		}
		lowestCost := mz.getCost(candidate)
		if candidate.cost > lowestCost {
			continue // drop this route (too expensive)
		}
		mz.setCost(candidate)
		cheapMoveVectors = append(cheapMoveVectors, candidate)
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
}

type moveVectorLoaded struct {
	moveVector
	cost    int
	history []moveVector
}

// get THREE candndates from the current one
// don't worry about valid / cheap. just get three
func (originMV moveVectorLoaded) getCandidates() []moveVectorLoaded {
	candidates := []moveVectorLoaded{}
	var left direction
	var right direction
	switch {
	case originMV.direction.y == 0:
		left = direction{0, 1}
		right = direction{0, -1}
	case originMV.direction.x == 0:
		left = direction{-1, 0}
		right = direction{1, 0}
	}
	targetNames := []string{"straight", "left", "right"}
	for _, targetName := range targetNames {
		var newPoint point
		var newVector moveVector
		costIncrement := 0
		switch targetName {
		case "straight":
			newPoint = point{originMV.point.x + originMV.direction.x,
				originMV.point.y + originMV.direction.y}
			newVector = moveVector{
				newPoint,
				originMV.direction,
			}
			costIncrement = stepCost
		case "left":
			newPoint = point{
				originMV.point.x + left.x,
				originMV.point.y + left.y,
			}
			newVector = moveVector{
				newPoint,
				direction{x: left.x, y: left.y},
			}
			costIncrement = stepCost + turnCost
		case "right":
			newPoint = point{
				originMV.point.x + right.x,
				originMV.point.y + right.y,
			}
			newVector = moveVector{
				newPoint,
				direction{x: right.x, y: right.y},
			}
			costIncrement = stepCost + turnCost
		}

		newCandidate := buildNewCandidate(originMV, newVector, costIncrement)
		candidates = append(candidates, newCandidate)
	}
	return candidates
}

func buildNewCandidate(
	originMV moveVectorLoaded,
	newVector moveVector,
	costIncrement int) moveVectorLoaded {
	newHistory := slices.Clone(originMV.history)
	newHistory = append(newHistory, newVector)
	newCandidate := moveVectorLoaded{
		moveVector: newVector,
		cost:       originMV.cost + costIncrement,
		history:    newHistory,
	}

	return newCandidate
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

// --------------------  DEBUGGING

// // ToString is for dedugging
// func (mv moveVectorLoaded) ToString() string {
// 	return fmt.Sprintf("vect: %d, %d (%d, %d)",
// 		mv.point.x, mv.point.y, mv.direction.x, mv.direction.y)
// }

// // ToString is for dedugging
// func (mv moveVector) ToString() string {
// 	return fmt.Sprintf("vect: %d, %d (%d, %d)",
// 		mv.point.x, mv.point.y, mv.direction.x, mv.direction.y)
// }

// // Dump is for debugging
// func (mz *Maze) Dump() {
// 	maxLen := len(mz.allowed)
// 	for y := range maxLen {
// 		line := ""
// 		for x := range maxLen {
// 			candidate := "#"
// 			if mz.allowed[y][x] {
// 				candidate = "."
// 			}
// 			switch {
// 			case mz.startPoint == point{x, y}:
// 				candidate = ("S")
// 			case mz.endPoint == point{x, y}:
// 				candidate = ("E")
// 			}
// 			line += candidate
// 		}
// 		fmt.Println(line)
// 	}
// }

// // DumpPath is for debugging
// func (mz *Maze) DumpPath(history []moveVector) {
// 	maxLen := len(mz.allowed)
// 	for y := range maxLen {
// 		line := ""
// 		for x := range maxLen {
// 			candidate := "#"
// 			if mz.allowed[y][x] {
// 				candidate = "."
// 			}
// 			for _, vect := range history {
// 				if y == vect.point.y && x == vect.point.x {
// 					switch {
// 					case vect.direction == direction{1, 0}:
// 						candidate = ">"
// 					case vect.direction == direction{-1, 0}:
// 						candidate = "<"
// 					case vect.direction == direction{0, 1}:
// 						candidate = "v"
// 					case vect.direction == direction{0, -1}:
// 						candidate = "^"
// 					}
// 				}
// 			}
// 			line += candidate
// 		}
// 		fmt.Println(line)
// 	}
// }
