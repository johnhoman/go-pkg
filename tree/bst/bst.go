package bst


type Value interface {
    Less(Value) bool
    Eq(Value)   bool
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

type Tree interface {
    Search(v Value)
    Insert(v Value)
}

type bst struct { node *Node }

// Insert a value into the Tree
func (t *bst) Insert(v Value) {
    if t.node == nil {
        t.node = &Node{v: v}
    }
    t.node.insert(v)
}

func (t *bst) Height() int {
    if t.node != nil {
        return t.node.height()
    }
    return 0
}

// New returns a new binary search tree
func New() *bst { return &bst{node: nil} }

func greater(v Value, o Value) bool { return !v.Eq(o) && !v.Less(o) }
func max(m int, nums ...int) int {
    for _, num := range nums {
        if num > m {
            m = num
        }
    }
    return m
}
