// Package stack is a stack implementation
package stack

// Stack represents the stack
type Stack[T any] struct {
	keys []T
}

// New returns a Stack
func New[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

// Push adds to the stack
func (s *Stack[T]) Push(key T) {
	s.keys = append(s.keys, key)
}

// Top returns the top of the stack
func (s *Stack[T]) Top() (T, bool) {
	var x T
	if len(s.keys) > 0 {
		x = s.keys[len(s.keys)-1]
		return x, true
	}
	return x, false
}

// Pop returns the top item and removes it from the stack
func (s *Stack[T]) Pop() (T, bool) {
	var x T
	if len(s.keys) > 0 {
		x, s.keys = s.keys[len(s.keys)-1], s.keys[:len(s.keys)-1]
		return x, true
	}
	return x, false
}

// IsEmpty is true when the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.keys) == 0
}
