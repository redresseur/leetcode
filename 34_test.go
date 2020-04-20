package algo

import "testing"

func TestSearchMin(t *testing.T) {
	t.Log(searchMin([]int{1, 2, 3, 4, 5, 6}, 1))
	t.Log(searchMin([]int{1, 2, 3, 4, 5, 6}, 2))
	t.Log(searchMin([]int{1, 2, 3, 4, 5, 6}, 3))
	t.Log(searchMin([]int{1, 2, 3, 4, 5, 6}, 4))
	t.Log(searchMin([]int{1, 2, 3, 4, 5, 6}, 5))
	t.Log(searchMin([]int{1, 2, 3, 4, 5, 6}, 6))

	t.Log(searchMin([]int{1, 2, 3, 3, 4, 5, 6}, 3))
	t.Log(searchMin([]int{1, 2, 3, 3, 4, 5, 6}, 3))
	t.Log(searchMin([]int{3, 3, 4, 5, 6}, 3))
	t.Log(searchMin([]int{3, 3}, 3))
	t.Log(searchMin([]int{1, 3}, 3))
	t.Log(searchMin([]int{1, 3}, 1))
	t.Log(searchMin([]int{1, 3, 3}, 3))
}

func TestSearchRange(t *testing.T) {
	t.Log(searchRange([]int{1, 2, 3, 4, 5, 6}, 1))
	t.Log(searchRange([]int{1, 2, 3, 4, 5, 6}, 2))
	t.Log(searchRange([]int{1, 2, 3, 4, 5, 6}, 3))
	t.Log(searchRange([]int{1, 2, 3, 4, 5, 6}, 4))
	t.Log(searchRange([]int{1, 2, 3, 4, 5, 6}, 5))
	t.Log(searchRange([]int{1, 2, 3, 4, 5, 6}, 6))
	t.Log(searchRange([]int{1, 2, 3, 4, 5, 6, 6}, 6))
	t.Log(searchRange([]int{1, 6, 6, 6}, 6))

	t.Log(searchRange([]int{1, 2, 3, 3, 4, 5, 6}, 3))
	t.Log(searchRange([]int{1, 2, 3, 3, 4, 5, 6}, 3))
	t.Log(searchRange([]int{3, 3, 4, 5, 6}, 3))
	t.Log(searchRange([]int{3, 3}, 3))
	t.Log(searchRange([]int{1, 3}, 3))
	t.Log(searchRange([]int{1, 3}, 1))
	t.Log(searchRange([]int{1, 3, 3}, 3))
	t.Log(searchRange([]int{3, 3, 3, 3, 3}, 3))
}
