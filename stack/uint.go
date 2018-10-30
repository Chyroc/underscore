package stack

import (
	"fmt"
	"strings"
)

type UintStack interface {
	fmt.Stringer

	Push(i uint)
	Pop() uint
	Peek() uint
	Len() int
	IsEmpty() bool
	Clone() UintStack
}

func NewUint() UintStack {
	return &uintStack{}
}

var _ UintStack = (*uintStack)(nil)

type uintStack struct {
	is []uint
}

func (s *uintStack) String() string {
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

func (s *uintStack) Push(i uint) {
	s.is = append(s.is, i)
}

func (s *uintStack) Pop() uint {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	p := s.Peek()
	s.is = s.is[:len(s.is)-1]
	return p
}

func (s *uintStack) Peek() uint {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s.is[len(s.is)-1]
}

func (s *uintStack) Len() int {
	return len(s.is)
}

func (s *uintStack) IsEmpty() bool {
	return len(s.is) == 0
}

func (s *uintStack) Clone() UintStack {
	s2 := &uintStack{is: make([]uint, 0, len(s.is))}
	for _, v := range s.is {
		s2.is = append(s2.is, v)
	}
	return s2
}
