package day4

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/brainboxweb/advent-2024/helpers"
)

func ChallengeOne(data []string) int {

	square := getSquare(data)
	rx, _ := regexp.Compile("XMAS")
	rx1, _ := regexp.Compile("SAMX")
	count := 0

	count += findHoriz(rx, rx1, square)
	count += findVertical(rx, rx1, square)

	count += findDiagonal(rx, rx1, square)
	count += findDiagonal2(rx, rx1, square)

	return count
}

func ChallengeTwo(data []string) int {
	square := getSquare(data)
	max := len(square[0])
	count := 0
	for x := 1; x < max-1; x++ {
		for y := 1; y < max-1; y++ {
			if square[x][y] == "A" {
				if isAnX(square, x, y) {
					count++
				}
			}
		}
	}

	return count
}

func isAnX(square [][]string, x, y int) bool {
	ends := []string{}
	ends = append(ends, square[x-1][y+1]) // pair with next
	ends = append(ends, square[x+1][y-1])
	ends = append(ends, square[x+1][y+1]) // pair with next
	ends = append(ends, square[x-1][y-1])

	if isCandidate(ends[0], ends[1]) && isCandidate(ends[2], ends[3]) {
		return true
	}

	return false
}

func getSquare(data []string) [][]string {
	square := [][]string{}
	for _, line := range data {
		parts := strings.Split(line, "")
		square = append(square, parts)
	}
	return square
}

// For debug
func printX(square [][]string, x, y int) {
	fmt.Println("---------")
	fmt.Println(square[x-1][y+1] + "." + square[x-1][y+1])
	fmt.Println(".A.")
	fmt.Println(square[x-1][y-1] + "." + square[x+1][y+1])
}

func isCandidate(a, b string) bool {
	if a == "M" && b == "S" {
		return true
	}
	if a == "S" && b == "M" {
		return true
	}

	return false
}

func findHoriz(rx, rx1 *regexp.Regexp, data [][]string) int {
	count := 0
	for _, line := range data {
		str := strings.Join(line, "")
		matches := rx.FindAllString(str, -1) // forward
		count += len(matches)
		matches1 := rx1.FindAllString(str, -1) // and reverse
		count += len(matches1)
	}
	return count
}

func findVertical(rx, rx1 *regexp.Regexp, data [][]string) int {
	count := 0

	dataTransposed := helpers.TransposeSliceOfSlices(data)

	for _, line := range dataTransposed {
		str := strings.Join(line, "")
		matches := rx.FindAllString(str, -1)
		count += len(matches)
		matches1 := rx1.FindAllString(str, -1) // and reverse
		count += len(matches1)
	}
	return count
}

func findDiagonal(rx, rx1 *regexp.Regexp, data [][]string) int {
	count := 0
	size := len(data[0])
	for c := -1 * size; c < size; c++ {
		chunk := []string{}
		for x := 0; x < size; x++ {
			y := x + c
			if y > size-1 || x > size-1 {
				continue
			}
			if y < 0 || x < 0 { // cobine with the above
				continue
			}
			chunk = append(chunk, data[x][y])
		}
		str := strings.Join(chunk, "")
		matches := rx.FindAllString(str, -1)
		count += len(matches)
		matches1 := rx1.FindAllString(str, -1) // and reverse
		count += len(matches1)
	}

	return count
}

func findDiagonal2(rx, rx1 *regexp.Regexp, data [][]string) int {
	count := 0
	size := len(data[0])
	for c := 1; c < size*2; c++ {
		chunk := []string{}
		for x := 0; x < size; x++ {
			y := -x + c
			if y > size-1 || x > size-1 {
				continue
			}
			if y < 0 || x < 0 { // cobine with the above
				continue
			}
			chunk = append(chunk, data[x][y])
		}
		str := strings.Join(chunk, "")
		matches := rx.FindAllString(str, -1)
		count += len(matches)
		matches1 := rx1.FindAllString(str, -1) // and reverse
		count += len(matches1)
	}

	return count
}
