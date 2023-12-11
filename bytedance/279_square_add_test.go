package bytedance

import (
	"sort"
	"testing"
)

func TestNumSquares(t *testing.T) {
	numSquaresByDp(1)
	numSquaresByDp(12)
}

func TestSort(t *testing.T) {
	arr := []int{2, 3, 1}
	sort.SliceStable(arr[1:], func(i, j int) bool {
		return arr[i] > arr[j]
	})

	t.Log(arr)
}
