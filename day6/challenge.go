package day6

import "github.com/brainboxweb/advent-2024/day6/location"

func ChallengeOne(data [][]string) int {
	theMap := location.NewMap(data)
	res := theMap.Move()

	return res
}

func ChallengeTwo(data [][]string) int {
	counter := 0
	for y := range len(data[0]) {
		for x := range len(data) {
			altData := makeVariation(data, x, y)
			theMap := location.NewMap(altData)
			isLoop := theMap.MoveTwo()
			if isLoop {
				counter++
			}
		}
	}

	return counter
}

func makeVariation(data [][]string, targetX, targetY int) [][]string {
	limit := len(data) // assume square
	altData := [][]string{}
	for _, line := range data {
		contents := make([]string, limit)
		copy(contents, line)
		altData = append(altData, contents)
	}
	if altData[targetX][targetY] != "^" {
		altData[targetX][targetY] = "#" // new obstacle
	}

	return altData
}
