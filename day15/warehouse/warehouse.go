// Package warehouse represents the bathroom
package warehouse

import (
	"fmt"

	"github.com/brainboxweb/advent-2024/stack"
)

// New return a new warehouse
func New(mapDetails [][]string) *Warehouse {
	w := &Warehouse{}
	w.populateAllowed(len(mapDetails))
	w.boxes = make(map[string]point)
	for y, line := range mapDetails {
		for x, val := range line {
			switch val {
			case "@":
				fmt.Println("robot!")
				w.robotLocn = point{x, y}
			case "O":
				fmt.Println("box!")
				w.boxes[makeIndex(x, y)] = point{x, y}
			case "#":
				fmt.Println("wall!")
				w.allowed[y][x] = false // or remove?
			}
		}
	}
	return w
}

// Warehouse represents the warehouse
type Warehouse struct {
	allowed   [][]bool
	robotLocn point            // just one.
	boxes     map[string]point // many // Is this a bit silly???
}

// Move makes a move in the specified direction
func (w *Warehouse) Move(move string) {
	var dirn direction
	switch move {
	case "^":
		dirn = direction{0, -1}
	case "<":
		dirn = direction{-1, 0}
	case ">":
		dirn = direction{1, 0}
	case "v":
		dirn = direction{0, +1}
	}
	w.makeMove(dirn)

	w.Dump() // For debug
}

// GPS returns the GPS representation of the warehouse
func (w *Warehouse) GPS() int {
	total := 0
	for _, p := range w.boxes {
		total += p.x + p.y*100
	}
	return total
}

func (w *Warehouse) makeMove(dirn direction) {
	currentPoint := point{w.robotLocn.x, w.robotLocn.y}
	myStack := stack.New[moveVector]()
	for {
		myStack.Push(moveVector{currentPoint, dirn})
		candidate := getTarget(currentPoint, dirn)
		// Not allowed?
		if !w.isAllowed(candidate) {
			return // ends everything
		}
		// Not empty?
		_, ok := w.boxes[makeIndex(candidate.x, candidate.y)]
		if ok {
			// there's a box in the way. Loop.
			currentPoint = candidate
			continue
		}
		// Move is possible: run the stack
		for {
			vel, ok := myStack.Pop()
			if !ok {
				return
			}
			w.makeTheMove(vel)
		}
	}
}

func (w *Warehouse) isAllowed(candidate point) bool {
	return w.allowed[candidate.y][candidate.x]
}

func (w *Warehouse) makeTheMove(vel moveVector) {
	start := vel.point
	target := getTarget(start, vel.direction)
	// If it's a robot...
	if start == w.robotLocn {
		w.robotLocn = target
		return
	}
	// If it's a box
	delete(w.boxes, makeIndex(start.x, start.y))
	w.boxes[makeIndex(target.x, target.y)] = target
}

func getTarget(p point, d direction) point {
	return point{p.x + d.x, p.y + d.y}
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

func makeIndex(a, b int) string {
	return fmt.Sprint(a, "_", b)
}

func (w *Warehouse) populateAllowed(limit int) {
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

// Dump is for debugging
func (w *Warehouse) Dump() {
	maxLen := len(w.allowed)
	for y := range maxLen {
		line := ""
		for x := range maxLen {
			candidate := "#"
			if w.allowed[y][x] {
				candidate = "."
			}
			switch {
			case w.robotLocn == point{x, y}:
				candidate = ("@")
			case w.boxes[makeIndex(x, y)] == point{x, y}:
				candidate = ("O")
			}
			line += candidate
		}
		fmt.Println(line)
	}
}
