package bst

import (
	"fmt"
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
	if !ok {
		return false
	}
	return i.v == o.v
}

func NewInteger(v int) *Integer {
	return &Integer{v: v}
}

func TestTree_Height(t *testing.T) {
	tests := []struct {
		ints     []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{1, 2, -1, -2, -3}, 4},
		{[]int{0, 2, -2, -1}, 3},
	}

	for k, subtest := range tests {
		t.Run(fmt.Sprintf("%d", k), func(t *testing.T) {
			tree := New()
			for _, i := range subtest.ints {
				tree.Insert(NewInteger(i))
			}
			require.Equal(t, subtest.expected, tree.Height())
		})
	}
}

func BenchmarkTree_IsBalanced(b *testing.B) {
}

func TestTree_IsBalanced(t *testing.T) {
	//            0
	//          /   \
	//        -2     4
	//              / \
	//             2   6
	//              \
	//               1
	tests := []struct {
		ints     []int
		expected bool
	}{
		{[]int{}, true},
		{[]int{0}, true},
		{[]int{0, -1, 1}, true},
		{[]int{0, -2, -3, -1, 4}, true},
		{[]int{0, -2, -3, -1, 4, 2, 1, 6}, true},
		{[]int{0, -2, -3, -1, 4, 2, 1}, false},
		{[]int{0, -2, 4, 2, 1, 6}, false},
	}

	for k, subtest := range tests {
		t.Run(fmt.Sprintf("%d", k), func(t *testing.T) {
			tree := New()
			for _, i := range subtest.ints {
				tree.Insert(NewInteger(i))
			}
			require.Equal(t, subtest.expected, tree.IsBalanced())
		})
	}
}

func TestTree_InOrder(t *testing.T) {
	//            0
	//          /   \
	//        -2     4
	//        / \   / \
	//      -3  -1 2   6
	//              \
	//               1
	tests := []struct {
		ints     []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{0}, []int{0}},
		{[]int{0, -1, 1}, []int{-1, 0, 1}},
		{[]int{0, -2, -3, -1, 4}, []int{-3, -2, -1, 0, 4}},
		{[]int{0, -2, -3, -1, 4, 2, 1, 6}, []int{-3, -2, -1, 0, 1, 2, 4, 6}},
	}

	for k, subtest := range tests {
		t.Run(fmt.Sprintf("%d", k), func(t *testing.T) {
			tree := New()
			for _, i := range subtest.ints {
				tree.Insert(NewInteger(i))
			}
			expected := make([]int, 0, len(subtest.expected))
			inOrder := tree.InOrder()
			for _, item := range inOrder {
				expected = append(expected, item.(*Integer).v)
			}
			require.Equal(t, subtest.expected, expected)
		})
	}
}