package stack

import (
	"fmt"
	"strings"
)

type ByteStack interface {
	fmt.Stringer

	Push(i byte)
	Pop() byte
	Peek() byte
	Len() int
	IsEmpty() bool
	Clone() ByteStack
}

func NewByte() ByteStack {
	return &byteStack{}
}

var _ ByteStack = (*byteStack)(nil)

type byteStack struct {
	is []byte
}

func (s *byteStack) String() string {
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

func (s *byteStack) Push(i byte) {
	s.is = append(s.is, i)
}

func (s *byteStack) Pop() byte {
	if s.IsEmpty() {
		panic("is empty")
	}
	p := s.Peek()
	s.is = s.is[:len(s.is)-1]
	return p
}

func (s *byteStack) Peek() byte {
	if s.IsEmpty() {
		panic("is empty")
	}
	return s.is[len(s.is)-1]
}

func (s *byteStack) Len() int {
	return len(s.is)
}

func (s *byteStack) IsEmpty() bool {
	return len(s.is) == 0
}

func (s *byteStack) Clone() ByteStack {
	s2 := &byteStack{is: make([]byte, 0, len(s.is))}
	for _, v := range s.is {
		s2.is = append(s2.is, v)
	}
	return s2
}
