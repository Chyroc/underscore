package stack

import (
	"fmt"
	"strings"
)

type Uint32Stack interface {
	fmt.Stringer

	Push(i uint32)
	Pop() uint32
	Peek() uint32
	Len() int
	IsEmpty() bool
	Clone() Uint32Stack
}

func NewUint32() Uint32Stack {
	return &uint32Stack{}
}

var _ Uint32Stack = (*uint32Stack)(nil)

type uint32Stack struct {
	is []uint32
}

func (s *uint32Stack) String() string {
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

func (s *uint32Stack) Push(i uint32) {
	s.is = append(s.is, i)
}

func (s *uint32Stack) Pop() uint32 {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	p := s.Peek()
	s.is = s.is[:len(s.is)-1]
	return p
}

func (s *uint32Stack) Peek() uint32 {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s.is[len(s.is)-1]
}

func (s *uint32Stack) Len() int {
	return len(s.is)
}

func (s *uint32Stack) IsEmpty() bool {
	return len(s.is) == 0
}

func (s *uint32Stack) Clone() Uint32Stack {
	s2 := &uint32Stack{is: make([]uint32, 0, len(s.is))}
	for _, v := range s.is {
		s2.is = append(s2.is, v)
	}
	return s2
}
