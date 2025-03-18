package disk_test

import (
	"strconv"
	"testing"

	"github.com/brainboxweb/advent-2024/day9/disk"
	"github.com/stretchr/testify/assert"
)

func TestBasicDisk(t *testing.T) {
	tests := []struct {
		data     string
		expected int
	}{
		{
			"2333133121414131402",
			1928,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			code := toIntSlice(tt.data)
			disk := disk.NewBasicDisk(code)
			disk.Compress()
			result := disk.Checksum()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestAdvancedDisk(t *testing.T) {
	tests := []struct {
		data     string
		expected int
	}{
		{
			"2333133121414131402",
			2858,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			code := toIntSlice(tt.data)
			disk := disk.NewAdvancedDisk(code)
			disk.Compress()
			result := disk.Checksum()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func toIntSlice(data string) []int {
	ret := []int{}
	for _, val := range data {
		numb, _ := strconv.Atoi(string(val))
		ret = append(ret, numb)
	}

	return ret
}
