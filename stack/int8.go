package stack

import (
	"fmt"
	"strings"
)

type Int8Stack interface {
	fmt.Stringer

	Push(i int8)
	Pop() int8
	Peek() int8
	Len() int
	IsEmpty() bool
	Clone() Int8Stack
}

func NewInt8() Int8Stack {
	return &int8Stack{}
}

var _ Int8Stack = (*int8Stack)(nil)

type int8Stack struct {
	is []int8
}

func (s *int8Stack) String() string {
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

func (s *int8Stack) Push(i int8) {
	s.is = append(s.is, i)
}

func (s *int8Stack) Pop() int8 {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	p := s.Peek()
	s.is = s.is[:len(s.is)-1]
	return p
}

func (s *int8Stack) Peek() int8 {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s.is[len(s.is)-1]
}

func (s *int8Stack) Len() int {
	return len(s.is)
}

func (s *int8Stack) IsEmpty() bool {
	return len(s.is) == 0
}

func (s *int8Stack) Clone() Int8Stack {
	s2 := &int8Stack{is: make([]int8, 0, len(s.is))}
	for _, v := range s.is {
		s2.is = append(s2.is, v)
	}
	return s2
}
