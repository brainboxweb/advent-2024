// Package disk is the disc drive
package disk

import (
	"errors"
)

// NewBasicDisk returns a basic disc
func NewBasicDisk(code []int) Disk {
	disk := buildBasicDisc(code)
	return &BasicDisk{disk: disk}
}

// BasicDisk represents a basic disk
type BasicDisk struct {
	disk []int
}

// Compress compresses a simple disk
func (d *BasicDisk) Compress() {
	for {
		_, err := moveLeft(d.disk)
		if err != nil {
			break
		}
	}
}

// Checksum gets the checksum of a simple disk
func (d *BasicDisk) Checksum() int {
	checksum := 0
	for k, val := range d.disk {
		if val == space {
			break
		}
		checksum += k * val
	}

	return checksum
}

func moveLeft(data []int) ([]int, error) {
	fromIndex := 0
	toIndex := 0
	blockValue := 0
	for i := len(data) - 1; i > -1; i-- {
		if data[i] != space {
			blockValue = data[i]
			fromIndex = i
			break
		}
	}
	for k, val := range data {
		if val == space {
			toIndex = k
			break
		}
	}
	// End condition
	if fromIndex-toIndex < 2 {
		return nil, errors.New("no change")
	}
	// Update the entries
	data[fromIndex] = space
	data[toIndex] = blockValue

	return data, nil
}

func buildBasicDisc(code []int) []int {
	disk := []int{}
	id := -1
	isFile := false
	for _, val := range code {
		isFile = !isFile
		var next int
		if isFile {
			id++
			next = id
		} else {
			next = space
		}
		for range val {
			disk = append(disk, next)
		}
	}
	return disk
}

// for debug
// func dump(data []string) {
// 	str := ""
// 	for _, val := range data {
// 		str += val
// 	}
// 	fmt.Println(str)
// }
