package tree

// Colors
const (
	RED   = true
	BLACK = false
)

// RBTreeNode represents a node in the Red-Black Tree.
type RBTreeNode[T TreeValue] struct {
	Value  T
	Left   *RBTreeNode[T]
	Right  *RBTreeNode[T]
	Parent *RBTreeNode[T]
	Color  bool
}

// RedBlackTree represents the Red-Black Tree.
type RedBlackTree[T TreeValue] struct {
	root *RBTreeNode[T]
}

func (t *RedBlackTree[T]) Root() *RBTreeNode[T] {
	return t.root
}

// LeftRotate performs a left rotation on the given subtree.
func (t *RedBlackTree[T]) LeftRotate(x *RBTreeNode[T]) {
	y := x.Right
	x.Right = y.Left
	if y.Left != nil {
		y.Left.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		t.root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

// RightRotate performs a right rotation on the given subtree.
func (t *RedBlackTree[T]) RightRotate(y *RBTreeNode[T]) {
	x := y.Left
	y.Left = x.Right
	if x.Right != nil {
		x.Right.Parent = y
	}
	x.Parent = y.Parent
	if y.Parent == nil {
		t.root = x
	} else if y == y.Parent.Left {
		y.Parent.Left = x
	} else {
		y.Parent.Right = x
	}
	x.Right = y
	y.Parent = x
}

// FixInsert fixes the Red-Black Tree properties after an insertion.
func (t *RedBlackTree[T]) FixInsert(z *RBTreeNode[T]) {
	for z.Parent != nil && z.Parent.Color == RED {
		if z.Parent == z.Parent.Parent.Left {
			y := z.Parent.Parent.Right
			if y != nil && y.Color == RED {
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Right {
					z = z.Parent
					t.LeftRotate(z)
				}
				z.Parent.Color = BLACK
				z.Parent.Parent.Color = RED
				t.RightRotate(z.Parent.Parent)
			}
		} else {
			y := z.Parent.Parent.Left
			if y != nil && y.Color == RED {
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Left {
					z = z.Parent
					t.RightRotate(z)
				}
				z.Parent.Color = BLACK
				z.Parent.Parent.Color = RED
				t.LeftRotate(z.Parent.Parent)
			}
		}
	}
	t.root.Color = BLACK
}

// Insert inserts a new node with the given value into the Red-Black Tree.
func (t *RedBlackTree[T]) Insert(value T) {
	z := &RBTreeNode[T]{Value: value, Color: RED}
	y := (*RBTreeNode[T])(nil)
	x := t.root

	for x != nil {
		y = x
		if z.Value < x.Value {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	z.Parent = y
	if y == nil {
		t.root = z
	} else if z.Value < y.Value {
		y.Left = z
	} else {
		y.Right = z
	}

	t.FixInsert(z)
}

// InOrder traverses the Red-Black Tree in order.
func (t *RedBlackTree[T]) InOrder(n *RBTreeNode[T]) {
	if n == nil {
		return
	}
	t.InOrder(n.Left)
	// fmt.Print(n.Value, " ")
	t.InOrder(n.Right)
}
