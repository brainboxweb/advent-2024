// Package branding respresents the towel collections
package branding

import (
	"strings"
)

// NewTowels returns a new towel collection
func NewTowels(input []string) *Towels {
	return &Towels{towelOptions: input}
}

// Towels contains the range of towels
type Towels struct {
	towelOptions []string
}

//revive:disable:cognitive-complexity

// IsPossible determines if the specified pattern is possible
func (tt *Towels) IsPossible(pattern string) bool {
	patterns := []string{pattern}
	var remainder string // try reusing
	for len(patterns) > 0 {
		var newPatterns = make(map[string]struct{})
		for _, pattern := range patterns {
			for _, towel := range tt.towelOptions {
				i := strings.Index(pattern, towel)
				if i == 0 { // starts with
					remainder = pattern[len(towel):]
					if len(remainder) == 0 { // is possible
						return true
					}
					newPatterns[remainder] = struct{}{}
				}
			}
		}
		patterns = nil
		for patt := range newPatterns {
			patterns = append(patterns, patt)
		}
	}
	return false
}

//revive:enable:cognitive-complexity
