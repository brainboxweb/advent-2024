// Package safety contains the Safety Manual
package safety

// Manual contains the rules
type Manual struct {
	rules [][]int
}

// NewManual returns a new Manual
func NewManual(ordering [][]int) *Manual {
	manual := Manual{rules: ordering}

	return &manual
}

// ValidateOrder validates the order
func (t *Manual) ValidateOrder(update []int) bool {
	applicableRules := t.findApplicableRules(update)
	for _, rule := range applicableRules {
		passed := applyRule(update, rule)
		if !passed {
			return false
		}
	}

	return true
}

func applyRule(update []int, rule []int) bool {
	isCandidate := false
	passed := false
	for i := range update {
		if update[i] == rule[0] {
			isCandidate = true
			continue
		}
		if update[i] == rule[1] {
			if isCandidate {
				passed = true
			}
		}
	}

	return passed
}

// FixUpdate corrects the Update
func (t *Manual) FixUpdate(update []int) []int {
	applicableRules := t.findApplicableRules(update)
LOOP:
	for _, rule := range applicableRules {
		passed := applyRule(update, rule)
		if !passed {
			update = swapElements(update, rule)
			goto LOOP // restart
		}
	}

	return update
}

func swapElements(update, rule []int) []int {
	var ret []int
	for k, item := range update {
		val := update[k]
		if item == rule[0] { // swap
			val = rule[1]
		}
		if item == rule[1] { // swap
			val = rule[0]
		}
		ret = append(ret, val)
	}

	return ret
}

func (t *Manual) findApplicableRules(update []int) [][]int {
	var newRules [][]int
	for _, rule := range t.rules {
		if isApplicable(update, rule) {
			newRules = append(newRules, rule)
		}
	}

	return newRules
}

func isApplicable(update []int, rule []int) bool {
	matchCount := 0
	for _, page := range update {
		if page == rule[0] {
			matchCount++
		}
		if page == rule[1] {
			matchCount++
		}
	}

	return matchCount == 2
}
