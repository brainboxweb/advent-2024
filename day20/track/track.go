// Package track represents the track layout
package track

import (
	"slices"
)

// New returns a new warehouse
func New(mapDetails [][]string) *Track {
	w := &Track{}
	w.populateAllowed(len(mapDetails))
	for y, line := range mapDetails {
		for x, val := range line {
			switch val {
			case "#":
				w.allowed[y][x] = false
			case ".":
				w.allowed[y][x] = true
			case "S":
				w.start = point{x, y}
			case "E":
				w.end = point{x, y}
			}
		}
	}
	return w
}

// Track represents the track layout
type Track struct {
	allowed [][]bool
	start   point
	end     point
}

// Go solves the track (maze)
func (w *Track) Go(minimumSaving int) int {
	pnt := w.start
	directions := []direction{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	visited := []point{pnt}
	var newPnt point
	for pnt != w.end {
		for _, dirn := range directions {
			candidate := point{pnt.x + dirn.x, pnt.y + dirn.y}
			if slices.Contains(visited, candidate) {
				continue
			}
			if w.isAllowed(candidate) {
				newPnt = candidate
				visited = append(visited, newPnt)
				continue
			}
		}
		pnt = newPnt
	}
	return w.findCheats(visited, minimumSaving)
}

func (w *Track) findCheats(path []point, minimumSaving int) int {
	var savingCount int
	directions := []direction{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for i, pnt := range path {
		for _, dirn := range directions {
			candidateWall := point{pnt.x + dirn.x, pnt.y + dirn.y}
			if w.isAllowed(candidateWall) {
				continue
			}
			candidate := point{pnt.x + 2*dirn.x, pnt.y + 2*dirn.y}
			saving := slices.Index(path, candidate) - i - 2
			if saving >= minimumSaving {
				savingCount++
			}
		}
	}
	return savingCount
}

func (w *Track) isAllowed(candidate point) bool {
	return w.allowed[candidate.y][candidate.x]
}

type point struct {
	x int
	y int
}

type direction struct {
	x int
	y int
}

func (w *Track) populateAllowed(limit int) {
	sliceOfSlices := make([][]bool, limit)
	for i := range limit {
		sliceOfSlices[i] = make([]bool, limit)
	}
	for y := range limit {
		for x := range limit {
			sliceOfSlices[y][x] = true
		}
	}
	w.allowed = sliceOfSlices
}

// // Dump is for debugging
// func (w *Track) Dump() {
// 	maxLen := len(w.allowed)
// 	for y := range maxLen {
// 		line := ""
// 		for x := range maxLen {
// 			candidate := "#"
// 			if w.allowed[y][x] {
// 				candidate = "."
// 			}
// 			switch {
// 			case point{x, y} == w.start:
// 				candidate = ("S")
// 			case point{x, y} == w.end:
// 				candidate = ("E")
// 			}
// 			line += candidate
// 		}
// 		fmt.Println(line)
// 	}
// }

// // Dump is for debugging
// func (w *Track) DumpPath(pnts []point) {
// 	maxLen := len(w.allowed)
// 	for y := range maxLen {
// 		line := ""
// 		for x := range maxLen {
// 			candidate := "#"
// 			if w.allowed[y][x] {
// 				candidate = "."
// 			}
// 			if slices.Contains(pnts, point{x, y}) {
// 				candidate = "0"
// 			}
// 			switch {
// 			case point{x, y} == w.start:
// 				candidate = ("S")
// 			case point{x, y} == w.end:
// 				candidate = ("E")
// 			}
// 			line += candidate
// 		}
// 		fmt.Println(line)
// 	}
// }
