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
    items := make([][]interface{}, 0)

    set := []interface{}{1, 2}
    Subsets(set, func(subset []interface{}) {
        sub := make([]interface{}, len(subset))
        reflect.Copy(reflect.ValueOf(sub), reflect.ValueOf(subset))
        items = append(items, sub)
        fmt.Printf("%#v\n", items)
    })
    expected := [][]interface{}{
        {},
        {1},
        {1, 2},
        {2},
    }
    require.Len(t, items, power(2, len(set)))
    require.Equal(t, expected, items)
}
