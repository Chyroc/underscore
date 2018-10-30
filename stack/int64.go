package stack

import (
	"fmt"
	"strings"
)

type Int64Stack interface {
	fmt.Stringer

	Push(i int64)
	Pop() int64
	Peek() int64
	Len() int
	IsEmpty() bool
	Clone() Int64Stack
}

func NewInt64() Int64Stack {
	return &int64Stack{}
}

var _ Int64Stack = (*int64Stack)(nil)

type int64Stack struct {
	is []int64
}

func (s *int64Stack) String() string {
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

func (s *int64Stack) Push(i int64) {
	s.is = append(s.is, i)
}

func (s *int64Stack) Pop() int64 {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	p := s.Peek()
	s.is = s.is[:len(s.is)-1]
	return p
}

func (s *int64Stack) Peek() int64 {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s.is[len(s.is)-1]
}

func (s *int64Stack) Len() int {
	return len(s.is)
}

func (s *int64Stack) IsEmpty() bool {
	return len(s.is) == 0
}

func (s *int64Stack) Clone() Int64Stack {
	s2 := &int64Stack{is: make([]int64, 0, len(s.is))}
	for _, v := range s.is {
		s2.is = append(s2.is, v)
	}
	return s2
}
