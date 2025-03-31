package stones

import (
	"fmt"
	"strconv"
	"strings"
)

func NewStoneSet(data string) *Set {
	stones := parse(data)
	stonesMap := make(map[int]int)
	for _, stone := range stones {
		stonesMap[stone]++
	}

	return &Set{stones, stonesMap}
}

type Set struct {
	stones    []int
	stonesMap map[int]int
}

func (ss *Set) Blinks(blinkCount int) int {
	for range blinkCount {
		ss.blink()
	}
	ret := 0
	for _, val := range ss.stonesMap {
		ret += val
	}

	return ret
}

func (ss *Set) blink() {
	type store struct {
		multiplier int
		newVals    []int
	}
	things := []store{}
	for k, val := range ss.stonesMap {
		next := applyRules(k)
		t := store{val, next}
		things = append(things, t)
	}
	// rebuild stonesMap
	for k := range ss.stonesMap { // empty it
		delete(ss.stonesMap, k)
	}
	for _, thing := range things {
		for _, val := range thing.newVals {
			ss.stonesMap[val] += thing.multiplier
		}
	}
}

var cache = make(map[int][]int)

func applyRules(input int) []int {
	if input == 0 { // --------------- If 0
		return []int{1}
	}
	val, ok := cache[input] // try to use cache
	if ok {
		return val
	}
	str := fmt.Sprint(input) // ------ If even no. of digits
	length := len(str)
	if length%2 == 0 {
		part1 := str[:(length / 2)]
		part2 := str[length/2:]
		n1, _ := strconv.Atoi(part1)
		n2, _ := strconv.Atoi(part2)
		ret := []int{n1, n2}
		cache[input] = ret
		return []int{n1, n2}
	}
	ret := []int{input * 2024}
	cache[input] = ret
	return ret // -------------------- Else
}

func parse(data string) []int {
	ret := []int{}
	parts := strings.Split(data, " ")
	for _, part := range parts {
		numb, err := strconv.Atoi(part)
		if err != nil {
			panic("Not expected")
		}
		ret = append(ret, numb)
	}

	return ret
}
