package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMin(t *testing.T) {
	assert.Equal(t, 1, minArray([]int{1, 1, 1}), 1)
	assert.Equal(t, 1, minArray([]int{10, 1, 10, 10, 10}))
	assert.Equal(t, 1, minArray([]int{10, 10, 10, 1, 10}))
	// assert.Equal(t, 10, findMin([]int{10, 10, 10, 10, 10}))
}
