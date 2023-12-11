package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseLeftWords(t *testing.T) {
	assert.Equal(t, "cdefgab", reverseLeftWords("abcdefg", 2))
	assert.Equal(t, "efabcd", reverseLeftWords("abcdef", 4))
	assert.Equal(t, "ab", reverseLeftWords("ba", 1))
	assert.Equal(t, "a", reverseLeftWords("a", 0))
}
