package bst

import "github.com/johnhoman/go-pkg/stack"

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
		return
	}
	if v.Less(n.v) {
        insertOrSet(&n.Left, v)
	}
	if greater(v, n.v) {
        insertOrSet(&n.Right, v)
	}
}

func (n *Node) remove(v Value) {
	if n == nil {
		return
	}
	if n.Right == nil && n.Left == nil {
		//     3            _
		//   /   \    =>
		//  _     _
		// This is node technically a real nil -- *n = nil is not possible
		n.v = nil
		return
	}
	if n.v.Eq(v) {
		// (a)     3      (b)  3     (c)    3
		//       /   \          \          /
		//      1     5          5        1
		//     / \   / \        / \      / \
		//    0   2 4   7      4   7    0   2

		// Removing the root - 3
		// Result tree
		//            5
		//           / \
		//          4   7
		//         /
		//        1
		//       / \
		//      0   2
		l := n.Left
		r := n.Right

		// (a) - Neither Right nor Left are nil
		// (b) - Left is nil and Right is non nil
		// (c) - Right is nil and Left is non nil
		if l != nil && r != nil {
			*n = *r
			prev := n
			current := n.Left
			for current != nil {
				prev = current
				current = current.Left
			}
			prev.Left = l
		} else {
			if r != nil { *n = *r } else { *n = *l }
		}
	} else {
		if v.Less(n.v) {
			n.Left.remove(v)
		} else {
			n.Right.remove(v)
		}
	}
}

func (n *Node) height() int {
	if n == nil {
		return 0
	}

    return max(n.Right.height(), n.Left.height()) + 1
}

func (n *Node) traverse(fn visitor) []Value {
	items := make([]Value, 0)
	fn(n, func(current *Node) *Node {
		if current.v != nil {
			items = append(items, current.v)
		}
		return current
	})
	return items
}

func (n *Node) inOrder() []Value { return n.traverse(inOrder) }

// preOrder traversal of the tree is root, left, right
func (n *Node) preOrder() []Value { return n.traverse(preOrder) }

// postOrder traversal of the tree is left, right, root
func (n *Node) postOrder() []Value { return n.traverse(postOrder) }

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

func (n *Node) max() Value {
	if n == nil {
		return nil
	}
	if n.Right == nil {
		return n.v
	} else {
		return n.Right.max()
	}
}

func (n *Node) min() Value {
	if n == nil {
		return nil
	}
	if n.Left == nil {
		return n.v
	} else {
		return n.Left.min()
	}
}

type Tree interface {
	Search(v Value)
	Insert(v Value)
}

type bst struct{ node *Node }

// Insert a value into the Tree. Insertion
// happens in O(log(n)) time for a balanced tree
// and O(n) for an unbalanced tree
func (t *bst) Insert(v Value) { insertOrSet(&(t.node), v) }

// Remove a value from the Tree. Removal happens
// in O(log(n)) time for a balanced tree
// and O(n) for an unbalanced tree
func (t *bst) Remove(v Value) { t.node.remove(v) }

// Height return the height of the binary search tree
func (t *bst) Height() int { return t.node.height() }

// IsBalanced returns true if the tree is balanced
func (t *bst) IsBalanced() bool { return t.node.isBalanced() }

// InOrder returns the in order traversal of the tree
func (t *bst) InOrder() []Value { return t.node.inOrder() }

// PreOrder returns the pre-order traversal of the tree
func (t *bst) PreOrder() []Value { return t.node.preOrder() }

// PostOrder returns the post-order traversal of the tree
func (t *bst) PostOrder() []Value { return t.node.postOrder() }

// Max returns the max value in the Tree
func (t *bst) Max() Value { return t.node.max() }

// Min returns the min value in the Tree
func (t *bst) Min() Value { return t.node.min() }

// New returns a new binary search tree
func New() *bst { return &bst{} }

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

func insertOrSet(n **Node, v Value) {
    if *n == nil {
       *n = &Node{}
    }
    (*n).insert(v)
}

type visitor func(n *Node, visit func(*Node) *Node)

func preOrder(n *Node, visit func(*Node) *Node) {

	current := n
	s := stack.New()

	for current != nil || !s.IsEmpty() {
		if current != nil {
			current = visit(current)
			s.Push(current)
			current = current.Left
		} else {
			current = s.Pop().(*Node)
			current = current.Right
		}
	}
}

func inOrder(n *Node, visit func(*Node) *Node) {

	s := stack.New()
	current := n

	for current != nil || !s.IsEmpty() {
		if current != nil {
			s.Push(current)
			current = current.Left
		} else {
			current = visit(s.Pop().(*Node))
			current = current.Right
		}
	}
}

func postOrderRecursive(n *Node, visit func(*Node) *Node) {
	if n == nil {
		return
	}
	postOrderRecursive(n.Left, visit)
	postOrderRecursive(n.Right, visit)
	visit(n)
}

func postOrder(n *Node, visit func(*Node) *Node) { postOrderRecursive(n, visit) }