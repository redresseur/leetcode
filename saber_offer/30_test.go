package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStack(t *testing.T) {
	minStack := Constructor30()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	assert.Equal(t, -3, minStack.Min())
	minStack.Pop()
	assert.Equal(t, -2, minStack.Min())
}
