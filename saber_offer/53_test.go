package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerchTargetRepeatedNum(t *testing.T) {
	assert.Equal(t, 1, search([]int{0, 1, 2, 3, 4}, 2))
	assert.Equal(t, 3, search([]int{0, 1, 1, 1, 3}, 1))
	assert.Equal(t, 2, search([]int{0, 1, 2, 2, 3}, 2))
	assert.Equal(t, 5, search([]int{0, 0, 0, 0, 0}, 0))
	assert.Equal(t, 0, search([]int{1, 2, 3, 4}, 0), 0)
	assert.Equal(t, 1, search([]int{0, 1, 2, 3, 4}, 4))
	assert.Equal(t, 1, search([]int{0, 1, 2, 3, 4}, 0))
	assert.Equal(t, 0, search([]int{1, 2, 3, 4}, 5))
}
