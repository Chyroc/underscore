package stack

import (
	"fmt"
	"strings"
)

type IntStack interface {
	fmt.Stringer

	Push(i int)
	Pop() int
	Peek() int
	Len() int
	IsEmpty() bool
	Clone() IntStack
}

func NewInt() IntStack {
	return &intStack{}
}

var _ IntStack = (*intStack)(nil)

type intStack struct {
	is []int
}

func (s *intStack) String() string {
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

func (s *intStack) Push(i int) {
	s.is = append(s.is, i)
}

func (s *intStack) Pop() int {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	p := s.Peek()
	s.is = s.is[:len(s.is)-1]
	return p
}

func (s *intStack) Peek() int {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s.is[len(s.is)-1]
}

func (s *intStack) Len() int {
	return len(s.is)
}

func (s *intStack) IsEmpty() bool {
	return len(s.is) == 0
}

func (s *intStack) Clone() IntStack {
	s2 := &intStack{is: make([]int, 0, len(s.is))}
	for _, v := range s.is {
		s2.is = append(s2.is, v)
	}
	return s2
}
