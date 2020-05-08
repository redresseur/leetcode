package binary_tree

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestTree_Insert(t *testing.T) {
	tree := Tree{}
	elements := []int{50, 6, 15, 3, 10, 47, 23, 7, 12, 17, 7, 66, 100}
	for _, e := range elements {
		tree.Insert(e)
	}

	tree.DLR();fmt.Println()
	assert.Equal(t, tree.Find(16), -1)
	assert.Equal(t, tree.Find(7), 2)

	for _, e := range elements {
		tree.Remove(e)
		tree.DLR();fmt.Println()
	}
	//tree.Remove(66)
	//tree.DLR();fmt.Println()
}

func TestTree_High(t *testing.T) {
	tree := Tree{}
	elements := []int{50, 6, 15, 3, 10, 47, 23, 7, 12, 17, 7, 66, 100}
	for _, e := range elements {
		tree.Insert(e)
	}

	t.Log(tree.DFSHigh())
	t.Log(tree.BFSHigh())
}