package helpers

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func ReverseSliceOfSlices(s [][]string) [][]string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func TransposeSliceOfSlices(slice [][]string) [][]string {
    xl := len(slice[0])
    yl := len(slice)
    result := make([][]string, xl)
    for i := range result {
        result[i] = make([]string, yl)
    }
    for i := 0; i < xl; i++ {
        for j := 0; j < yl; j++ {
            result[i][j] = slice[j][i]
        }
    }
    return result
}

func GetData(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var dat []int
	for scanner.Scan() {
		if scanner.Text() == "" {
			dat = append(dat, 0)
			continue
		}
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic("not expected")
		}
		dat = append(dat, i)
	}
	return dat
}

func GetDataString(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var dat []string
	for scanner.Scan() {
		dat = append(dat, scanner.Text())
	}
	return dat
}
