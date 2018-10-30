package stack

import (
	"fmt"
	"strings"
)

type Int16Stack interface {
	fmt.Stringer

	Push(i int16)
	Pop() int16
	Peek() int16
	Len() int
	IsEmpty() bool
	Clone() Int16Stack
}

func NewInt16() Int16Stack {
	return &int16Stack{}
}

var _ Int16Stack = (*int16Stack)(nil)

type int16Stack struct {
	is []int16
}

func (s *int16Stack) String() string {
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

func (s *int16Stack) Push(i int16) {
	s.is = append(s.is, i)
}

func (s *int16Stack) Pop() int16 {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	p := s.Peek()
	s.is = s.is[:len(s.is)-1]
	return p
}

func (s *int16Stack) Peek() int16 {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s.is[len(s.is)-1]
}

func (s *int16Stack) Len() int {
	return len(s.is)
}

func (s *int16Stack) IsEmpty() bool {
	return len(s.is) == 0
}

func (s *int16Stack) Clone() Int16Stack {
	s2 := &int16Stack{is: make([]int16, 0, len(s.is))}
	for _, v := range s.is {
		s2.is = append(s2.is, v)
	}
	return s2
}
