package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReversePrint(t *testing.T) {
	head := &ListNode{}
	for p, i := head, 0; ; {
		p.Val = i
		i++
		if i >= 4 {
			break
		}
		p.Next = &ListNode{}
		p = p.Next
	}
	assert.Equal(t, []int{3, 2, 1, 0}, reversePrint(head))
}
