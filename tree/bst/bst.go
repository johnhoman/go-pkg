package bst

type Value interface {
	Less(Value) bool
	Eq(Value) bool
}

type Node struct {
	v     Value
	Left  *Node
	Right *Node
}

func (n *Node) insert(v Value) {
	if n.v == nil {
		n.v = v
	} else {
		if v.Less(n.v) {
			if n.Left == nil {
				n.Left = &Node{}
			}
			n.Left.insert(v)
		}
		if greater(v, n.v) {
			if n.Right == nil {
				n.Right = &Node{}
			}
			n.Right.insert(v)
		}
	}
}

func (n *Node) height() int {
	if n == nil {
		return 0
	}

	m := 0
	if n.Right != nil {
		m = max(m, n.Right.height())
	}
	if n.Left != nil {
		m = max(m, n.Left.height())
	}
	return 1 + m
}

// isBalanced returns true if the tree is a balanced
// tree. A balanced tree is defined as a binary tree in
// which the height of the left and right subtree differ
// by no more than 1
func (n *Node) isBalanced() bool {
	if n == nil {
		return true
	}
	return isBalanced(n.Left, n.Right)
}

type Tree interface {
	Search(v Value)
	Insert(v Value)
}

type bst struct{ node *Node }

// Insert a value into the Tree. Insertion
// happens in O(log(n)) time for a balanced tree
// and O(n) for an unbalanced tree
func (t *bst) Insert(v Value) { t.node.insert(v) }

// Height return the height of the binary search tree
func (t *bst) Height() int { return t.node.height() }

// IsBalanced returns true if the tree is balanced
func (t *bst) IsBalanced() bool { return t.node.isBalanced() }

// New returns a new binary search tree
func New() *bst { return &bst{node: &Node{}} }

func greater(v Value, o Value) bool { return !v.Eq(o) && !v.Less(o) }
func max(m int, nums ...int) int {
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	return m
}

func abs(i int) int {
	if i < 0 {
		return 0 - i
	}
	return i
}

func isBalanced(n1 *Node, n2 *Node) bool {
	if n1 != nil && n2 != nil {
		if abs(n1.height()-n2.height()) > 1 {
			return false
		}
		return n1.isBalanced() && n2.isBalanced()
	}
	if n1 != nil && n1.height() > 1 || n2 != nil && n2.height() > 1 {
		return false
	}
	return true
}
