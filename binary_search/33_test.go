package binary_search

import (
	"testing"
)

func TestSearchLittle(t *testing.T) {
	t.Log(searchLittle([]int{4, 5, 6, 1, 2, 3}))
	t.Log(searchLittle([]int{4, 1}))
	t.Log(searchLittle([]int{4, 1, 2}))
	t.Log(searchLittle([]int{4, 1, 2, 3}))
	t.Log(searchLittle([]int{4, 5, 1, 2, 3}))
	t.Log(searchLittle([]int{2, 3, 4, 5, 1}))

	t.Log(searchLittle([]int{2, 2, 3, 4, 4, 5, 5, 1, 1}))
	t.Log(searchLittle([]int{2, 2, 3, 4, 4, 5, 5, 1, 1}))
	t.Log(searchLittle([]int{2, 2, 3, 4, 4, 5, 5, 1, 1}))
	t.Log(searchLittle([]int{2, 2, 3, 4, 4, 5, 5}))
	t.Log(searchLittle([]int{2, 2, 3, 4, 4, 5, 1}))
	t.Log(searchLittle([]int{1, 2, 2, 3, 4, 4, 5, 1}))
	t.Log(searchLittle([]int{2, 2, 2, 2, 2}))
	t.Log(searchLittle([]int{1, 2}))
	t.Log(searchLittle([]int{1, 1, 2, 1}))
	t.Log(searchLittle([]int{2, 1}))
	t.Log(searchLittle([]int{1, 2, 1}))
}

func TestSearchIndex(t *testing.T) {
	t.Log(SearchIndex([]int{4, 1, 2, 3}, 4))
	t.Log(SearchIndex([]int{4, 1, 2, 3}, 1))
	t.Log(SearchIndex([]int{4, 1, 2, 3}, 2))
	t.Log(SearchIndex([]int{4, 1, 2, 3}, 3))
	t.Log(SearchIndex([]int{4, 1, 2, 3}, 5))
	t.Log(SearchIndex([]int{4, 1}, 1))
	t.Log(SearchIndex([]int{4, 1}, 4))
	t.Log(SearchIndex([]int{1, 2, 3, 4}, 1))
	t.Log(SearchIndex([]int{1, 2, 3, 4}, 2))
	t.Log(SearchIndex([]int{1, 2, 3, 4}, 3))
	t.Log(SearchIndex([]int{1, 2, 3, 4}, 4))
	t.Log(SearchIndex([]int{1, 2, 3, 4, 1, 1, 1, 1}, 4))
	t.Log(SearchIndex([]int{1, 2, 3, 4, 1, 1, 1, 1, 1}, 4))
	t.Log(SearchIndex([]int{1, 2, 3, 4, 1, 1, 1, 1, 1}, 2))
	t.Log(SearchIndex([]int{1, 2, 3, 4, 1, 1, 1, 1, 1}, 3))
	t.Log(SearchIndex([]int{1, 2, 3, 4, 1, 1, 1, 1, 1}, 1))
	t.Log(SearchIndex([]int{2, 3, 4, 0, 1, 1, 1, 1}, 1))
	t.Log(SearchIndex([]int{1, 1, 2, 1}, 2))
}
