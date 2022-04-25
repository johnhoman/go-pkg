package stack

import (
	"container/list"

	"github.com/johnhoman/go-pkg/collection"
)

type Stack interface {
	collection.Interface
	Push(v interface{})
	Pop() interface{}
	Top() interface{}
}

type stack struct {
	items *list.List
}

func (s *stack) Len() int { return s.items.Len() }

func (s *stack) Push(v interface{}) { s.items.PushBack(v) }

func (s *stack) Pop() interface{} {
	x := s.items.Back()
	s.items.Remove(x)
	return x.Value
}

func (s *stack) Top() interface{} {
	return s.items.Back()
}

func (s *stack) IsEmpty() bool { return s.Len() == 0 }

func New() *stack {
	s := &stack{items: list.New()}
	s.items.Init()
	return s
}

var _ Stack = &stack{}