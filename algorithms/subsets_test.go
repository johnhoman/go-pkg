package algorithms

import (
    "fmt"
    "reflect"
    "testing"

    "github.com/stretchr/testify/require"
)

func power(base, exp int) int {
    k := 1
    for x := 0; x < exp; x++ {
        k = k * base
    }
    return k
}

func TestSubsets(t *testing.T) {

    tests := []struct{
        values []interface{}
        expected [][]interface{}
    } {
        {
            []interface{}{1, 2},
            [][]interface{}{
                {},
                {1},
                {1, 2},
                {2},
            },
        },
    }

    for k, subtest := range tests {
        t.Run(fmt.Sprintf("%d", k), func(t *testing.T) {
            items := make([][]interface{}, 0)

            Subsets(subtest.values, func(subset []interface{}) {
                sub := make([]interface{}, len(subset))
                reflect.Copy(reflect.ValueOf(sub), reflect.ValueOf(subset))
                items = append(items, sub)
                fmt.Printf("%#v\n", items)
            })
            require.Len(t, items, power(2, len(subtest.values)))
            require.Equal(t, subtest.expected, items)
        })
    }
}
