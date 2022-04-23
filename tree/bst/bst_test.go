package bst

import (
	"testing"

    "github.com/stretchr/testify/require"
)

type Integer struct {
    v int
}

func (i *Integer) Less(other Value) bool {
    o, ok := other.(*Integer)
    if !ok {
        return false
    }
    return o.v > i.v
}

func (i *Integer) Eq(other Value) bool {
    o, ok := other.(*Integer)
    if !ok { return false }
    return i.v == o.v
}

func NewInteger(v int) *Integer {
    return &Integer{v: v}
}

func TestTree_Search(t *testing.T) {

    bst := New()
    bst.Insert(NewInteger(1))
    bst.Insert(NewInteger(2))
    bst.Insert(NewInteger(3))
    bst.Insert(NewInteger(4))

    // 1
    //  \
    //   2
    //    \
    //     3
    //      \
    //       4
    require.Equal(t, 4, bst.Height())
    bst.Insert(NewInteger(-1))
    require.Equal(t, 4, bst.Height())
    bst.Insert(NewInteger(-2))
    require.Equal(t, 4, bst.Height())
    bst.Insert(NewInteger(-3))
    require.Equal(t, 4, bst.Height())
    bst.Insert(NewInteger(-4))
    require.Equal(t, 5, bst.Height())
}
