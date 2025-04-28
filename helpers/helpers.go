// Package helpers contains useful tools
package helpers

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ReverseSlice reverses a slice
func ReverseSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// ReverseSliceOfSlices reverses a slice of slices
func ReverseSliceOfSlices(s [][]string) [][]string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// TransposeSliceOfSlices transposes a slice of slices
func TransposeSliceOfSlices(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := range xl {
		for j := range yl {
			result[i][j] = slice[j][i]
		}
	}

	return result
}

// GetDataString returns a slice of strings from a file
func GetDataString(filename string) []string {
	file, err := os.Open(filepath.Clean(filename))
	if err != nil {
		panic(err)
	}
	defer checkClose(file, &err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var ret []string
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	return ret
}

// GetDataInt returns a slice of ints from a file
func GetDataInt(filename string) ([]int, error) {
	data := GetDataString(filename)
	var ret []int
	for _, line := range data {
		val, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		ret = append(ret, val)
	}

	return ret, nil
}

// ToXY converts lines to an XY grid
func ToXY(data []string) [][]string {
	limit := len(data)
	sliceOfSlices := make([][]string, limit)
	for i := range limit {
		sliceOfSlices[i] = make([]string, limit)
		line := data[i]
		chars := strings.Split(line, "")
		sliceOfSlices[i] = chars
	}

	return TransposeSliceOfSlices(sliceOfSlices)
}

// DumpXY prints an XY grid
func DumpXY(data [][]string) {
	printable := TransposeSliceOfSlices(data)
	for _, line := range printable {
		_, _ = fmt.Println(strings.Join(line, ""))
	}
}

func checkClose(c io.Closer, err *error) {
	cerr := c.Close()
	if *err == nil {
		*err = cerr
	}
}
