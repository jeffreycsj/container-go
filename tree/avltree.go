package tree

// AvlTreeNode represents a node in the AVL tree.
type AvlTreeNode[T TreeValue] struct {
	Value  T
	Left   *AvlTreeNode[T]
	Right  *AvlTreeNode[T]
	Height int
}

// Height returns the height of the tree node.
func (n *AvlTreeNode[T]) GetHeight() int {
	if n == nil {
		return 0
	}
	return n.Height
}

// Max returns the maximum of two integers.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// RightRotate performs a right rotation on the given subtree.
func RightRotate[T TreeValue](y *AvlTreeNode[T]) *AvlTreeNode[T] {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	y.Height = Max(y.Left.GetHeight(), y.Right.GetHeight()) + 1
	x.Height = Max(x.Left.GetHeight(), x.Right.GetHeight()) + 1

	return x
}

// LeftRotate performs a left rotation on the given subtree.
func LeftRotate[T TreeValue](x *AvlTreeNode[T]) *AvlTreeNode[T] {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	x.Height = Max(x.Left.GetHeight(), x.Right.GetHeight()) + 1
	y.Height = Max(y.Left.GetHeight(), y.Right.GetHeight()) + 1

	return y
}

// BalanceFactor returns the balance factor of the node.
func (n *AvlTreeNode[T]) BalanceFactor() int {
	if n == nil {
		return 0
	}
	return n.Left.GetHeight() - n.Right.GetHeight()
}

// Insert inserts a new node with the given value into the AVL tree.
func (n *AvlTreeNode[T]) Insert(value T) *AvlTreeNode[T] {
	if n == nil {
		return &AvlTreeNode[T]{Value: value, Height: 1}
	}

	if value < n.Value {
		n.Left = n.Left.Insert(value)
	} else if value > n.Value {
		n.Right = n.Right.Insert(value)
	} else {
		return n
	}

	n.Height = Max(n.Left.GetHeight(), n.Right.GetHeight()) + 1
	bf := n.BalanceFactor()

	if bf > 1 && value < n.Left.Value {
		return RightRotate(n)
	}

	if bf < -1 && value > n.Right.Value {
		return LeftRotate(n)
	}

	if bf > 1 && value > n.Left.Value {
		n.Left = LeftRotate(n.Left)
		return RightRotate(n)
	}

	if bf < -1 && value < n.Right.Value {
		n.Right = RightRotate(n.Right)
		return LeftRotate(n)
	}

	return n
}

// InOrder traverses the AVL tree in order.
func (n *AvlTreeNode[T]) InOrder() {
	if n == nil {
		return
	}
	n.Left.InOrder()
	// fmt.Print(n.Value, " ")
	n.Right.InOrder()
}
