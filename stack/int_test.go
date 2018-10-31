package stack_test

import (
	"github.com/Chyroc/underscore/stack"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt(t *testing.T) {
	as := assert.New(t)

	s := stack.NewInt()

	as.True(s.IsEmpty())
	as.Equal(0, s.Len())

	s.Push(1)
	s.Push(2)

	as.False(s.IsEmpty())
	as.Equal(2, s.Len())

	as.Equal(2, s.Peek())
}
