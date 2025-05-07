// Package garden represents the garden plots
package garden

import (
	"strconv"
)

// NewPlot returns a new Plot
func NewPlot(grid [][]string) Plot {
	grouper := findRegions(grid)
	regions := doGouping(grouper)

	return Plot{grid: grid, regions: regions}
}

// Plot is a garden plot
type Plot struct {
	regions [][]point
	grid    [][]string
}

type point struct {
	x              int
	y              int
	name           string
	neighbourCount int
}

// TotalCost is the total cost of fencing
func (pl Plot) TotalCost() int {
	totalCost := 0
	for _, region := range pl.regions {
		perim := pl.perimeter(region)
		cost := len(region) * perim
		totalCost += cost
	}
	return totalCost
}

func doGouping(grouper map[string][]point) [][]point {
	regions := [][]point{}
	for _, points := range grouper {
		regions = append(regions, points)
	}
	return regions
}

func scan(data [][]string) []point {
	layout := []point{}
	for y, line := range data {
		for x, pnt := range line {
			layout = append(layout, point{name: pnt, x: x, y: y})
		}
	}
	return layout
}

func makeIndex(a, b int) string {
	return strconv.Itoa(a) + "_" + strconv.Itoa(b)
}

func findRegions(grid [][]string) map[string][]point {
	grouper := make(map[string][]point)
	for y, line := range grid {
		for x, name := range line {
			id := makeIndex(x, y)
			grouper[id] = append(grouper[id], point{x: x, y: y, name: name})
		}
	}
	for y, line := range grid {
		for x, name := range line {
			currentPoint := point{x: x, y: y, name: name}
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
	currentID := makeIndex(currentPoint.x, currentPoint.y)
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

func inRegion(x, y int, region []point) bool {
	for _, pnt := range region {
		if pnt.x == x && pnt.y == y {
			return true
		}
	}
	return false
}

//revive:disable:cognitive-complexity
//revive:disable:cyclomatic

func getMatchingNeighbours(pnt point, grid [][]string) []point {
	ret := []point{}
	x := pnt.x
	y := pnt.y
	label := grid[y][x]
	var compare string

	// up
	yy := y - 1
	xx := x
	if xx < len(grid) && xx > -1 && yy < len(grid) && yy > -1 {
		compare = grid[yy][xx]
		if label == compare {
			ret = append(ret, point{x: xx, y: yy, name: label})
		}
	}

	// left
	yy = y
	xx = x - 1
	if xx < len(grid) && xx > -1 && yy < len(grid) && yy > -1 {
		compare = grid[yy][xx]
		if label == compare {
			ret = append(ret, point{x: xx, y: yy, name: label})
		}
	}

	// right
	yy = y
	xx = x + 1
	if xx < len(grid) && xx > -1 && yy < len(grid) && yy > -1 {
		compare = grid[yy][xx]
		if label == compare {
			ret = append(ret, point{x: xx, y: yy, name: label})
		}
	}

	// down
	yy = y + 1
	xx = x
	if xx < len(grid) && xx > -1 && yy < len(grid) && yy > -1 {
		compare = grid[yy][xx]
		if label == compare {
			ret = append(ret, point{x: xx, y: yy, name: label})
		}
	}

	return ret
}
func (pl Plot) perimeter(region []point) int {
	data := pl.grid
	maxi := len(data)
	edgeCount := 0
	for y := range maxi {
		inBlock := false
		for x := range maxi {
			if inRegion(x, y, region) {
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
	for x := range maxi {
		inBlock := false
		for y := range maxi {
			// count edges
			if inRegion(x, y, region) {
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
//revive:enable:cyclomatic
