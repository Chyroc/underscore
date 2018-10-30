package stack

import (
	"fmt"
	"strings"
)

type InterfaceStack interface {
	fmt.Stringer

	Push(i interface{})
	Pop() interface{}
	Peek() interface{}
	Len() int
	IsEmpty() bool
	Clone() InterfaceStack
}

func NewInterface() InterfaceStack {
	return &interfaceStack{}
}

var _ InterfaceStack = (*interfaceStack)(nil)

type interfaceStack struct {
	is []interface{}
}

func (s *interfaceStack) String() string {
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

func (s *interfaceStack) Push(i interface{}) {
	s.is = append(s.is, i)
}

func (s *interfaceStack) Pop() interface{} {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	p := s.Peek()
	s.is = s.is[:len(s.is)-1]
	return p
}

func (s *interfaceStack) Peek() interface{} {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s.is[len(s.is)-1]
}

func (s *interfaceStack) Len() int {
	return len(s.is)
}

func (s *interfaceStack) IsEmpty() bool {
	return len(s.is) == 0
}

func (s *interfaceStack) Clone() InterfaceStack {
	s2 := &interfaceStack{is: make([]interface{}, 0, len(s.is))}
	for _, v := range s.is {
		s2.is = append(s2.is, v)
	}
	return s2
}
