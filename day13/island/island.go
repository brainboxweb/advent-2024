// Package island represents the tropical island
package island

import "slices"

// Point represents an xy location
type Point struct {
	X int
	Y int
}

// Machine represents the claw machine
type Machine struct {
	A     Point
	B     Point
	Prize Point
	// Limit int
}

//revive:disable:cognitive-complexity
//revive:disable:cyclomatic

// Calculate runs the machine
func (t Machine) Calculate(pressLimit int) int {
	type winner struct {
		CountA int
		countB int
	}
	winners := []winner{}
	bCount := pressLimit + 1

	// if pressLimit == -1 { // @TODO - For part 2
	// 	// figure out a resonable stating value
	// 	BCount = 1000000000000
	// }

	for {
		bCount--

		if bCount < 0 { // exit condition
			break
		}

		bTotalX := bCount * t.B.X
		if bTotalX > t.Prize.X { // overshoot X
			continue
		}
		bTotalY := bCount * t.B.Y
		if bTotalY > t.Prize.Y { // overshoot Y
			continue
		}

		// If we get here, there's a gap
		remainderX := t.Prize.X - bTotalX
		remainderY := t.Prize.Y - bTotalY

		// we need remainders to be zero
		if remainderX%t.A.X != 0 {
			continue
		}
		if remainderY%t.A.Y != 0 {
			continue
		}

		// Divisible in both dimensions?
		// (apply mod first?)
		aCount := remainderY / t.A.Y
		aCountAgain := remainderX / t.A.X
		if aCount != aCountAgain {
			continue
		}

		// Check button count
		if pressLimit != -1 {
			if aCount > pressLimit || bCount > pressLimit {
				continue
			}
		}

		// Collect the valid options
		winners = append(winners, winner{aCount, bCount})
	}

	// Calculate scores
	scores := []int{}
	for _, win := range winners {
		score := 3*win.CountA + win.countB
		scores = append(scores, score)
	}

	return getLowestScore(scores)
}

//revive:enable:cognitive-complexity
//revive:enable:cyclomatic

func getLowestScore(a []int) int {
	slices.Sort(a)
	for _, v := range a {
		if v != 0 {
			return v // return the frst non-zero
		}
	}

	return 0
}
