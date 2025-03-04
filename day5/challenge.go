package day5

import (
	"strconv"
	"strings"

	"github.com/brainboxweb/advent-2024/day5/safety"
)

func ChallengeOne(data []string) int {
	rules, updates := parse(data)
	manual := safety.NewManual(rules)
	selected := [][]int{}
	for _, update := range updates {
		if manual.ValidateOrder(update) {
			selected = append(selected, update)
		}
	}

	return getScore(selected)
}

func ChallengeTwo(data []string) int {
	rules, updates := parse(data)
	manual := safety.NewManual(rules)
	invalidUpdates := [][]int{}
	for _, update := range updates {
		if !manual.ValidateOrder(update) {
			invalidUpdates = append(invalidUpdates, update)
		}
	}
	fixedUpdates := [][]int{}
	for _, update := range invalidUpdates {
		fixedUpdate := manual.FixUpdate(update)
		fixedUpdates = append(fixedUpdates, fixedUpdate)
	}
	selected := [][]int{}
	for _, fixes := range fixedUpdates {
		if manual.ValidateOrder(fixes) {
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
			ordering = getOrdering(line, ordering)
			continue
		}
		updates = getUpdates(line, updates)
	}

	return ordering, updates
}

func getUpdates(line string, updates [][]int) [][]int {
	parts := strings.Split(line, ",")
	update := []int{}
	for _, part := range parts {
		val, _ := strconv.Atoi(part)
		update = append(update, val)
	}
	updates = append(updates, update)

	return updates
}

func getOrdering(line string, ordering [][]int) [][]int {
	parts := strings.Split(line, "|")
	order := []int{}
	for _, part := range parts {
		val, _ := strconv.Atoi(part)
		order = append(order, val)
	}
	ordering = append(ordering, order)

	return ordering
}
