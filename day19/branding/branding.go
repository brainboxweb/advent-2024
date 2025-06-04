// Package branding respresents the towel collections
package branding

import (
	"slices"

	"github.com/brainboxweb/advent-2024/cache"
)

// NewTowels returns a new towel collection
func NewTowels(input []string) *Towels {
	myCache := cache.New[string, int]()
	maxTowelLength := 0
	for _, val := range input {
		if len(val) > maxTowelLength {
			maxTowelLength = len(val)
		}
	}
	return &Towels{
		towelOptions:   input,
		maxTowelLength: maxTowelLength,
		patternCache:   myCache,
	}
}

// Towels represents the collection of towels
type Towels struct {
	towelOptions   []string
	maxTowelLength int
	patternCache   *cache.Cache[string, int]
}

// CanObtain returns the number of permutations of the provided pattern
func (tt *Towels) CanObtain(pattern string) int {
	if pattern == "" { // exit condition
		return 1
	}
	val, ok := tt.patternCache.Get(pattern)
	if ok {
		return val
	}
	count := 0
	for i := 1; i < min(tt.maxTowelLength, len(pattern))+1; i++ {
		if slices.Contains(tt.towelOptions, pattern[:i]) {
			count += tt.CanObtain(pattern[i:])
		}
	}
	tt.patternCache.Set(pattern, count)
	return count
}
