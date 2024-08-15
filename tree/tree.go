package tree

type TreeValue interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// TreeNode represents a node in a binary tree.
type TreeNode[T TreeValue] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

// Insert inserts a new node with the given value into the binary tree.
func (n *TreeNode[T]) Insert(value T) {
	if value < n.Value {
		if n.Left == nil {
			n.Left = &TreeNode[T]{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode[T]{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

// InOrder traverses the binary tree in order.
func (n *TreeNode[T]) InOrder() {
	if n == nil {
		return
	}
	n.Left.InOrder()
	// fmt.Print(n.Value, " ")
	n.Right.InOrder()
}
