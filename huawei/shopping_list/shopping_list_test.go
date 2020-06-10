package main

import (
	"testing"
)

func TestShoppingList(t *testing.T) {
	t.Log(ShoppingList1(1500, []*Goods{
		{500, 1, nil, -1},
		{400, 4, nil, -1},
		{300, 5, nil, 0},
		{400, 5, nil, 0},
		{200, 5, nil, -1},
		{500, 4, nil, -1},
		{400, 4, nil, -1},
	}))
}

