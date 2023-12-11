package binary_tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTree_Insert(t *testing.T) {
	tree := Tree{}
	elements := []int{50, 6, 15, 3, 10, 47, 23, 7, 12, 17, 7, 66, 100}
	for _, e := range elements {
		tree.Insert(e)
	}

	tree.DLR()
	fmt.Println()
	assert.Equal(t, -1, tree.Find(16))
	assert.Equal(t, 2, tree.Find(7))

	for _, e := range elements {
		tree.Remove(e)
		tree.DLR()
		fmt.Println()
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

func TestTree_LayerErgodic(t *testing.T) {
	tree := Tree{}
	elements := []int{50, 6, 15, 3, 10, 47, 23, 7, 12, 17, 66, 100}
	for _, e := range elements {
		tree.Insert(e)
	}
	assert.Equal(t, len(elements), len(tree.BfsErgodic()))
	t.Log(tree.BfsErgodic())
}

func TestTree_DfsErgodic(t *testing.T) {
	tree := Tree{}
	elements := []int{50, 6, 15, 3, 10, 47, 23, 7, 12, 17, 66, 100}
	for _, e := range elements {
		tree.Insert(e)
	}
	assert.Equal(t, len(elements), len(tree.DfsErgodic()))
	t.Log(tree.DfsErgodic())
}
