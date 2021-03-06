// Code generated by underscore_generate. DO NOT EDIT.

package stack

import (
	"fmt"
	"strings"
)

type Uint16Stack interface {
	fmt.Stringer

	Push(i uint16)
	Pop() uint16
	Peek() uint16
	Len() int
	IsEmpty() bool
	Clone() Uint16Stack
}

func NewUint16() Uint16Stack {
	return &uint16Stack{}
}

var _ Uint16Stack = (*uint16Stack)(nil)

type uint16Stack struct {
	is []uint16
}

func (s *uint16Stack) String() string {
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

func (s *uint16Stack) Push(i uint16) {
	s.is = append(s.is, i)
}

func (s *uint16Stack) Pop() uint16 {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	p := s.Peek()
	s.is = s.is[:len(s.is)-1]
	return p
}

func (s *uint16Stack) Peek() uint16 {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s.is[len(s.is)-1]
}

func (s *uint16Stack) Len() int {
	return len(s.is)
}

func (s *uint16Stack) IsEmpty() bool {
	return len(s.is) == 0
}

func (s *uint16Stack) Clone() Uint16Stack {
	s2 := &uint16Stack{is: make([]uint16, 0, len(s.is))}
	for _, v := range s.is {
		s2.is = append(s2.is, v)
	}
	return s2
}
