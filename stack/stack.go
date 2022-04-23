package stack

import "github.com/johnhoman/go-pkg/collection"

type Stack interface {
	collection.Interface
	Push(v interface{})
	Pop() interface{}
	Top() interface{}
}

type stack struct {
	items []interface{}
}

func (s *stack) Len() int { return len(s.items) }

func (s *stack) Push(v interface{}) { s.items = append(s.items, v) }

func (s *stack) Pop() interface{} {
	if len(s.items) > 0 {
		x := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		return x
	}
	return nil
}

func (s *stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.items[len(s.items)-1]
}

func (s *stack) IsEmpty() bool { return len(s.items) == 0 }

func New() *stack {
	return &stack{items: make([]interface{}, 0)}
}

var _ Stack = &stack{}