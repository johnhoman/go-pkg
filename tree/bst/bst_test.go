package bst

import (
    "bytes"
    "fmt"
    "strings"
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

func TestTree_Height(t *testing.T) {
    tests := []struct{
        ints []int
        expected int
    } {
        {[]int{1, 2, 3, 4, 5}, 5},
        {[]int{1, 2, -1, -2, -3}, 4},
        {[]int{0, 2, -2, -1}, 3},
    }

    for _, subtest := range tests {
        name := fmt.Sprintf("Height(%s)=%d", repr(subtest.ints), subtest.expected)
        t.Run(name, func(t *testing.T) {
            tree := New()
            for _, i := range subtest.ints {
                tree.Insert(NewInteger(i))
            }
            require.Equal(t, subtest.expected, tree.Height())
        })
    }
}


func repr(items []int) string {
    buf := new(bytes.Buffer)
    values := make([]string, 0, len(items))
    for _, item := range items {
        values = append(values, fmt.Sprintf("%d", item))
    }
    buf.WriteString("[")
    buf.WriteString(strings.Join(values, ","))
    buf.WriteString("]")
    return buf.String()
}