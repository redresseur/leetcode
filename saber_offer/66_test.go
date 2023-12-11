package main

import "testing"

func TestMaxProfit(t *testing.T) {
	maxProfit([]int{1, 2})
	maxProfit([]int{2, 1, 4})
	maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4})
	maxProfit([]int{7, 1, 5, 3, 6, 4})
}
