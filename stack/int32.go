package stack

import (
	"fmt"
	"strings"
)

type Int32Stack interface {
	fmt.Stringer

	Push(i int32)
	Pop() int32
	Peek() int32
	Len() int
	IsEmpty() bool
	Clone() Int32Stack
}

func NewInt32() Int32Stack {
	return &int32Stack{}
}

var _ Int32Stack = (*int32Stack)(nil)

type int32Stack struct {
	is []int32
}

func (s *int32Stack) String() string {
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

func (s *int32Stack) Push(i int32) {
	s.is = append(s.is, i)
}

func (s *int32Stack) Pop() int32 {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	p := s.Peek()
	s.is = s.is[:len(s.is)-1]
	return p
}

func (s *int32Stack) Peek() int32 {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s.is[len(s.is)-1]
}

func (s *int32Stack) Len() int {
	return len(s.is)
}

func (s *int32Stack) IsEmpty() bool {
	return len(s.is) == 0
}

func (s *int32Stack) Clone() Int32Stack {
	s2 := &int32Stack{is: make([]int32, 0, len(s.is))}
	for _, v := range s.is {
		s2.is = append(s2.is, v)
	}
	return s2
}
