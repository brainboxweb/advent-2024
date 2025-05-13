// Package bathroom represents the bathroom
package bathroom

// New return a new bathroom
func New(lenX, lenY int) *Bathroom {
	return &Bathroom{lenX: lenX, lenY: lenY}
}

// AddRobot adds a robot to the bathroom
func (b *Bathroom) AddRobot(x, y, vx, vy int) {
	b.robots = append(b.robots, &robot{x, y, vx, vy})
}

type robot struct {
	x  int
	y  int
	vx int
	vy int
}

// Bathroom represents the bathroom
type Bathroom struct {
	lenX   int
	lenY   int
	robots []*robot
}

// Step increments the steps for ALL of the robots
func (b *Bathroom) Step(stepCount int) {
	for _, rob := range b.robots {
		for range stepCount {
			rob.step(b.lenX, b.lenY) // make naming consistent
		}
	}
}

//revive:disable:cognitive-complexity

// CountQuadrants counts the robots in each quadrant
func (b *Bathroom) CountQuadrants() map[int]int {
	totals := make(map[int]int)
	for y := range b.lenY {
		for x := range b.lenX {
			for _, rob := range b.robots {
				if x == rob.x && y == rob.y {
					quad := b.quadrant(x, y)
					if quad != 0 {
						totals[quad]++
					}
				}
			}
		}
	}
	return totals
}

//revive:enable:cognitive-complexity

func (b *Bathroom) quadrant(x, y int) int {
	midX := (b.lenX) / 2
	midY := (b.lenY) / 2
	if x == midX || y == midY {
		return 0 // Not in a quadrant
	}
	if x < midX {
		if y < midY {
			return 1
		}
		return 2
	}
	if x > midX {
		if y < midY {
			return 3
		}
		return 4
	}
	return 0 // not expected
}

func (r *robot) step(lenX, lenY int) {
	r.x = increment(r.x, r.vx, lenX)
	r.y = increment(r.y, r.vy, lenY)
}

func increment(val, increment, length int) int {
	if increment > length || -increment > length {
		panic("not expected")
	}
	val += increment
	if val < 0 {
		val = length + val
	}
	if val > length-1 {
		val -= length
	}
	return val
}
