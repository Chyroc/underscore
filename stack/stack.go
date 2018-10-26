package stack

import (
	"fmt"
	"strings"
)

func New() Stack {
	return &stack{}
}

type Stack interface {
	fmt.Stringer

	Push(i interface{})
	Pop() interface{}
	Peek() interface{}
	Len() int
	IsEmpty() bool
	Clone() Stack
}

type stack struct {
	is []interface{}
}

func (s *stack) String() string {
	var buf = new(strings.Builder)
	for i, v := range s.is {
		if i == 0 {
			buf.WriteString(fmt.Sprintf("%v", v))
		} else {
			buf.WriteString(fmt.Sprintf(" < %v", v))
		}
	}
	return buf.String()
}

func (s *stack) Push(i interface{}) {
	s.is = append(s.is, i)
}

func (s *stack) Pop() interface{} {
	if s.IsEmpty() {
		panic("is empty")
	}
	p := s.Peek()
	s.is = s.is[:len(s.is)-1]
	return p
}

func (s *stack) Peek() interface{} {
	if s.IsEmpty() {
		panic("is empty")
	}
	return s.is[len(s.is)-1]
}

func (s *stack) Len() int {
	return len(s.is)
}

func (s *stack) IsEmpty() bool {
	return len(s.is) == 0
}

func (s *stack) Clone() Stack {
	s2 := &stack{is: make([]interface{}, 0, len(s.is))}
	for _, v := range s.is {
		s2.is = append(s2.is, v)
	}
	return s2
}
