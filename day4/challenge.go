package day4

import (
	"regexp"
	"strings"

	"github.com/brainboxweb/advent-2024/helpers"
)

var rx = regexp.MustCompile("XMAS")
var rx2 = regexp.MustCompile("SAMX")

func ChallengeOne(data []string) int {
	square := getSquare(data)
	count := 0
	count += findHoriz(rx, rx2, square)
	count += findVertical(rx, rx2, square)
	count += findDiagonal(rx, rx2, square)
	count += findDiagonal2(rx, rx2, square)

	return count
}

func ChallengeTwo(data []string) int {
	square := getSquare(data)
	maximum := len(square[0])
	count := 0
	for x := 1; x < maximum-1; x++ {
		for y := 1; y < maximum-1; y++ {
			if checkA(square, x, y) {
				count++
			}
		}
	}

	return count
}

func checkA(square [][]string, x, y int) bool {
	if square[x][y] == "A" {
		if isAnX(square, x, y) {
			return true
		}
	}

	return false
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

func isCandidate(a, b string) bool {
	if a == "M" && b == "S" || a == "S" && b == "M" {
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
		for x := range size {
			y := x + c
			if y > size-1 || x > size-1 || y < 0 || x < 0 {
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
		for x := range size {
			y := -x + c
			if y > size-1 || x > size-1 || y < 0 || x < 0 {
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

// For debug
// func printX(square [][]string, x, y int) {
// 	_, _ = fmt.Println("---------")
// 	_, _ = fmt.Println(square[x-1][y+1] + "." + square[x-1][y+1])
// 	_, _ = fmt.Println(".A.")
// 	_, _ = fmt.Println(square[x-1][y-1] + "." + square[x+1][y+1])
// }
