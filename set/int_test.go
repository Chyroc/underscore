package set_test

import (
	"github.com/Chyroc/underscore/set"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt(t *testing.T) {
	as := assert.New(t)

	s := set.NewInt()
	s.Add(1)
	s.Adds(2, 3, 4)
	for i := 1; i <= 4; i++ {
		as.True(s.Exist(i))
	}
}
