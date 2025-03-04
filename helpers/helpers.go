package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	for i := range xl {
		for j := range yl {
			result[i][j] = slice[j][i]
		}
	}

	return result
}

func GetDataString(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var ret []string
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	return ret
}

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

func DumpXY(data [][]string) {
	printable := TransposeSliceOfSlices(data)
	for _, line := range printable {
		_, _ = fmt.Println(strings.Join(line, ""))
	}
}
