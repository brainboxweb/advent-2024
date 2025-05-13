// Package day14 is for Advent of Code Day 14
package day14

import (
	"strconv"
	"strings"

	"github.com/brainboxweb/advent-2024/day14/bathroom"
)

// ChallengeOne is part one
func ChallengeOne(data []string, lengthX, lengthY, stepCount int) int {
	isle := bathroom.New(lengthX, lengthY)
	for _, robotData := range data {
		point, velocity := parse(robotData)
		isle.AddRobot(point.x, point.y, velocity.x, velocity.y)
	}
	isle.Step(stepCount)
	totals := isle.CountQuadrants()
	ret := 1
	for _, total := range totals {
		ret *= total
	}
	return ret
}

func parse(input string) (point, velocity) {
	parts := strings.Split(input, " ")

	posnPart := parts[0]
	posnPart = strings.Trim(posnPart, "p=")
	xyStrings := strings.Split(posnPart, ",")

	x, err := strconv.Atoi(xyStrings[0])
	if err != nil {
		panic("not expected")
	}
	y, err := strconv.Atoi(xyStrings[1])
	if err != nil {
		panic("not expected")
	}
	p := point{x, y}

	velocityPart := parts[1]
	velocityPart = strings.Trim(velocityPart, "v=")
	velocityStrings := strings.Split(velocityPart, ",")
	x, err = strconv.Atoi(velocityStrings[0])
	if err != nil {
		panic("not expected")
	}
	y, err = strconv.Atoi(velocityStrings[1])
	if err != nil {
		panic("not expected")
	}
	v := velocity{x, y}

	return p, v
}

type point struct {
	x int
	y int
}

type velocity struct {
	x int
	y int
}
