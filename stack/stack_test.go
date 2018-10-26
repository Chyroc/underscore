package stack_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Chyroc/underscore/stack"
)

func Test_Stack(t *testing.T) {
	as := assert.New(t)

	s := stack.New()
	as.True(s.IsEmpty())
	s.Push(2)
	s.Push(3)
	as.False(s.IsEmpty())
	as.Equal(3, s.Pop())
	as.Equal(2, s.Peek())
	as.False(s.IsEmpty())
	as.Equal(2, s.Pop())
	as.True(s.IsEmpty())
}

func Test_stack_ref_clone(t *testing.T) {
	as := assert.New(t)

	{
		var stackRef = func(stack stack.Stack) { stack.Push(100) }

		s := stack.New()
		stackRef(s)
		as.Equal(100, s.Pop())
	}

	{
		s := stack.New()
		s.Push(2)
		s.Push(3)
		s2 := s.Clone()

		s.Push(4)
		as.Equal("2 < 3 < 4", s.String())
		as.Equal("2 < 3", s2.String())
	}
}
