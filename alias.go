package pkg

import (
	"github.com/johnhoman/go-pkg/stack"
	"github.com/johnhoman/go-pkg/tree/bst"
)

var (
	NewBST = bst.New
	Stack  = stack.New
)

var (
	_ = NewBST
	_ = Stack
)
