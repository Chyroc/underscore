package stack

import (
	"fmt"
	"strings"
)

type Uint64Stack interface {
	fmt.Stringer

	Push(i uint64)
	Pop() uint64
	Peek() uint64
	Len() int
	IsEmpty() bool
	Clone() Uint64Stack
}

func NewUint64() Uint64Stack {
	return &uint64Stack{}
}

var _ Uint64Stack = (*uint64Stack)(nil)

type uint64Stack struct {
	is []uint64
}

func (s *uint64Stack) String() string {
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

func (s *uint64Stack) Push(i uint64) {
	s.is = append(s.is, i)
}

func (s *uint64Stack) Pop() uint64 {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	p := s.Peek()
	s.is = s.is[:len(s.is)-1]
	return p
}

func (s *uint64Stack) Peek() uint64 {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s.is[len(s.is)-1]
}

func (s *uint64Stack) Len() int {
	return len(s.is)
}

func (s *uint64Stack) IsEmpty() bool {
	return len(s.is) == 0
}

func (s *uint64Stack) Clone() Uint64Stack {
	s2 := &uint64Stack{is: make([]uint64, 0, len(s.is))}
	for _, v := range s.is {
		s2.is = append(s2.is, v)
	}
	return s2
}
