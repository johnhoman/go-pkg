package stack

import (
    "fmt"
    "github.com/stretchr/testify/require"
    "testing"
)

func TestStack_Push(t *testing.T) {
    tests := []struct{
        test int
        ints []int
        expectedLen int
    } {
        {1, []int{}, 0},
        {2, []int{1}, 1},
        {3, []int{1, 2}, 2},
    }
    for _, subtest := range tests {
        t.Run(fmt.Sprintf("%d", subtest.test), func(t *testing.T) {
            s := New()
            for _, i := range subtest.ints {
                s.Push(i)
            }
            require.Equal(t, subtest.expectedLen, s.Len())
        })
    }
}