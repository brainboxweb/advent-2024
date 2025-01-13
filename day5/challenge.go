package day5

import (
	"strconv"
	"strings"
)

func ChallengeOne(data []string) int {
	rules, updates := parse(data)
	safety := newSafetyManual(rules)
	ret := 0
	selected := [][]int{}
	for _, update := range updates { // extract
		if safety.ValidateOrder(update) {
			selected = append(selected, update)
			midIndex := (len(update) - 1) / 2
			mid := update[midIndex]
			ret += mid
		}
	}

	return getScore(selected)
}

func ChallengeTwo(data []string) int {
	rules, updates := parse(data)
	safety := newSafetyManual(rules)
	invalidUpdates := [][]int{}
	for _, update := range updates {
		if !safety.ValidateOrder(update) {
			invalidUpdates = append(invalidUpdates, update)
		}
	}
	fixedUpdates := [][]int{}
	for _, update := range invalidUpdates {
		fixedUpdate := safety.FixUpdate(update)
		fixedUpdates = append(fixedUpdates, fixedUpdate)
	}
	selected := [][]int{}
	for _, fixes := range fixedUpdates { // extrac
		if safety.ValidateOrder(fixes) {
			selected = append(selected, fixes)
		}
	}

	return getScore(selected)
}

func getScore(updates [][]int) int {
	score := 0
	for _, update := range updates {
		midIndex := (len(update) - 1) / 2
		score += update[midIndex]
	}

	return score
}

type SafetyManual struct {
	rules [][]int
}

func newSafetyManual(ordering [][]int) *SafetyManual {
	thing := SafetyManual{rules: ordering}

	return &thing
}

func (t *SafetyManual) ValidateOrder(update []int) bool {
	applicableRules := t.findApplicableRules(update)
	for _, rule := range applicableRules {
		passed := applyRule(update, rule)
		if !passed {
			return false
		}
	}

	return true
}

func applyRule(update []int, rule []int) bool { // method??
	isCandidate := false
	passed := false
	for i := 0; i < len(update); i++ {
		if update[i] == rule[0] {
			isCandidate = true
			continue // skip to look for the second match.
		}
		if update[i] == rule[1] {
			if isCandidate == true {
				passed = true
			}
		}
	}

	return passed
}

func (t *SafetyManual) FixUpdate(update []int) []int {
	applicableRules := t.findApplicableRules(update)

loop:
	for _, rule := range applicableRules {	
		passed := applyRule(update, rule) // will always fail first time
		if !passed {
			update = swapElements(update, rule)
			goto loop // start again
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

func (t *SafetyManual) findApplicableRules(update []int) [][]int { // return the indexes???
	var newRules [][]int
	for _, rule := range t.rules {
		matchCount := 0
		for _, page := range update {
			if page == rule[0] {
				matchCount++
			}
			if page == rule[1] {
				matchCount++
			}
		}
		if matchCount == 2 { // both in rule
			newRules = append(newRules, rule)
		}
	}

	return newRules
}

func parse(data []string) ([][]int, [][]int) {
	var ordering [][]int
	var updates [][]int
	splitByPipe := true
	for _, line := range data {
		if line == "" {
			splitByPipe = false
			continue
		}
		if splitByPipe {
			parts := strings.Split(line, "|")
			order := []int{}
			for _, part := range parts {
				val, _ := strconv.Atoi(part)
				order = append(order, val)
			}
			ordering = append(ordering, order)
		}
		if !splitByPipe {
			parts := strings.Split(line, ",")
			update := []int{}
			for _, part := range parts {
				val, _ := strconv.Atoi(part)
				update = append(update, val)
			}
			updates = append(updates, update)
		}
	}

	return ordering, updates
}
