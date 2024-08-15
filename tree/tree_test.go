package tree_test

import (
	"container-go/tree"
	"testing"
)

func BenchmarkMaps(b *testing.B) {
	b.Run("Tree", func(b *testing.B) {
		root := &tree.TreeNode[int]{Value: 10}
		root.Insert(5)
		root.Insert(15)
		root.Insert(3)
		root.Insert(7)
		root.Insert(12)
		root.Insert(18)
	
		root.InOrder() // Output: 3 5 7 10 12 15 18
	})

	b.Run("AvlTree", func(b *testing.B) {
		root := &tree.AvlTreeNode[int]{Value: 10}
		root.Insert(5)
		root.Insert(15)
		root.Insert(3)
		root.Insert(7)
		root.Insert(12)
		root.Insert(18)
	
		root.InOrder() // Output: 3 5 7 10 12 15 18
	})

	b.Run("RBTree", func(b *testing.B) {
		tree := &tree.RedBlackTree[int]{}
		values := []int{10, 20, 30, 15, 25, 5, 1}
		for _, value := range values {
			tree.Insert(value)
		}
		tree.InOrder(tree.Root()) // Output will depend on the tree structure
	})
}