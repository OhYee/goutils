package stack

import (
	"container/list"
	"fmt"
)

// Stack stack struct
type Stack struct {
	stack *list.List
}

// ErrStackEmpty want using Top() or Pop() on a empty stack.
var ErrStackEmpty = fmt.Errorf("Stack error: stack is empty, no value to get")

// NewStack declare a stack
func NewStack() *Stack {
	return &Stack{list.New()}
}

// Push a value to the stack
func (s *Stack) Push(value interface{}) {
	s.stack.PushBack(value)
}

// Pop a value from the stack
func (s *Stack) Pop() (value interface{}, err error) {
	e := s.stack.Back()
	if e != nil {
		value = e.Value
		s.stack.Remove(e)
	} else {
		err = ErrStackEmpty
	}
	return
}

// Top return the top value of the stack and do not pop it
func (s *Stack) Top() (value interface{}, err error) {
	e := s.stack.Back()
	if e != nil {
		value = e.Value
	} else {
		err = ErrStackEmpty
	}
	return
}

// Len return the legnth of the stack
func (s *Stack) Len() int {
	return s.stack.Len()
}

// Empty return the stack is empty
func (s *Stack) Empty() bool {
	return s.stack.Len() == 0
}
