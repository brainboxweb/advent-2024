package stack_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/stack"
	"github.com/stretchr/testify/assert"
)

func TestEmptyStack(t *testing.T) {
	stk := stack.New[int]()
	result := stk.IsEmpty()
	assert.True(t, result)
}

func TestPushPop(t *testing.T) {
	stk := stack.New[int]()
	stk.Push(7)
	stk.Push(11)

	result, ok := stk.Pop()
	assert.True(t, ok)
	assert.Equal(t, 11, result)

	result, ok = stk.Pop()
	assert.True(t, ok)
	assert.Equal(t, 7, result)

	_, ok = stk.Pop()
	assert.False(t, ok)
}

func TestTop(t *testing.T) {
	stk := stack.New[int]()

	_, ok := stk.Top()
	assert.False(t, ok)

	stk.Push(7)
	stk.Push(11)

	result, ok := stk.Top()
	assert.True(t, ok)
	assert.Equal(t, 11, result)
}
