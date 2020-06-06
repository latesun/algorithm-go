package stack

import (
	"fmt"
)

// ArrayStack is array based on stack.
type ArrayStack struct {
	data []interface{}
	top  int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

func (s *ArrayStack) IsEmpty() bool {
	return s.top < 0
}

func (s *ArrayStack) Push(v interface{}) {
	if s.top < 0 {
		s.top = 0
	} else {
		s.top += 1
	}

	if s.top > len(s.data)-1 {
		s.data = append(s.data, v)
	} else {
		s.data[s.top] = v
	}
}

func (s *ArrayStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	s.top -= 1
	return s.data[s.top+1]
}

func (s *ArrayStack) Print() {
	if s.IsEmpty() {
		fmt.Println("empty stack")
	}

	for i := s.top; i >= 0; i-- {
		fmt.Println(s.data[i])
	}
}
