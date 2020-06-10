package main

import "testing"

func TestCount(t *testing.T) {
	t.Log(Count(4,3))
	t.Log(Count(7,3))
	t.Log(Count(3,3))
	t.Log(Count(3,4))
	t.Log(Count(0,4))
}