package algorithms

type VisitFunc func([]interface{})

func subsets(set []interface{}, subset []interface{}, k int, visitFunc VisitFunc) {

    visitFunc(subset)
    for x := k; x < len(set); x++ {
        subset = append(subset, set[x])
        subsets(set, subset, x+1, visitFunc)
        subset = subset[:len(subset)-1]
    }
}

func Subsets(items []interface{}, visitFunc VisitFunc) {
    subset := make([]interface{}, 0)
    subsets(items, subset, 0, visitFunc)
}