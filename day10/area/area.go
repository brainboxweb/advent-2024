package area

import (
	"fmt"
	"strconv"
)

type Terrain struct {
	XY             [][]int
	StartingPoints []Point
}

type Point struct {
	X      int
	Y      int
	Height int
}

func New(data [][]string) *Terrain {
	limit := len(data)
	sliceOfSlices := make([][]int, limit)
	for i := range limit {
		sliceOfSlices[i] = []int{}
		for _, item := range data[i] {
			numb, _ := strconv.Atoi(item)
			sliceOfSlices[i] = append(sliceOfSlices[i], numb)
		}
	}

	return &Terrain{XY: sliceOfSlices}
}

func (t *Terrain) FindStartingPoints() {
	limit := len(t.XY)
	for i := range limit {
		for j := range limit {
			if t.XY[i][j] == 0 {
				p := Point{X: i, Y: j, Height: 0}
				t.StartingPoints = append(t.StartingPoints, p)
			}
		}
	}
}

func (t *Terrain) Walk(start Point) map[string]int {
	endPoints := make(map[string]int) // Need distinct
	startArray := []Point{}
	startArray = append(startArray, start)
	for {
		if len(startArray) < 1 {
			break
		}
		theBigNextSteps := []Point{}
		for _, current := range startArray {
			if current.Height == 9 {
				key := fmt.Sprintf("%d_%d", current.X, current.Y)
				endPoints[key]++
				continue
			}
			nextSteps := t.Step(current)
			theBigNextSteps = append(theBigNextSteps, nextSteps...)
		}
		startArray = theBigNextSteps
	}

	return endPoints
}

func (t *Terrain) Step(start Point) []Point {
	limit := len(t.XY)
	targetHeight := start.Height + 1 // Step up
	nextSteps := []Point{}
	for i := -1; i < 2; i++ { // left and right
		x := start.X + i
		y := start.Y
		if x < 0 || x > limit-1 || x == start.X {
			continue
		}
		if t.XY[x][y] == targetHeight {
			pnt := Point{x, y, targetHeight}
			nextSteps = append(nextSteps, pnt)
		}
	}
	for j := -1; j < 2; j++ { // up and down
		x := start.X
		y := start.Y + j
		if y < 0 || y > limit-1 || y == start.Y {
			continue
		}
		if t.XY[x][y] == targetHeight {
			pnt := Point{x, y, targetHeight}
			nextSteps = append(nextSteps, pnt)
		}
	}

	return nextSteps
}
