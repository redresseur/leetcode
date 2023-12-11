package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	obj := Constructor()
	obj.AppendTail(1)
	assert.Equal(t, 1, obj.DeleteHead())

	obj.AppendTail(2)
	obj.AppendTail(3)
	assert.Equal(t, 2, obj.DeleteHead())
}
