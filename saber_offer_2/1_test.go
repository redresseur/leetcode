package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	exp := 1<<31 - 1
	min := -(1 << 31)
	max := (1<<31 - 1)
	assert.Equal(t, 15/2, divide(15, 2))
	assert.Equal(t, -100867/90, divide(-100867, 90))
	assert.Equal(t, exp, divide(1, 0))
	assert.Equal(t, 1, divide(exp, exp))
	assert.Equal(t, max, divide(max, 1))
	assert.Equal(t, -max, divide(max, -1))
	assert.Equal(t, 1<<29, divide(1<<30, 2))
	assert.Equal(t, min, divide(min, 1))
	assert.Equal(t, 1, divide(min, min))
	assert.Equal(t, 8, divide(16, 2))
}
