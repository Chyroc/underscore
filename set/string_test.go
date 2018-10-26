package set_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Chyroc/underscore/set"
)

func TestStringSet(t *testing.T) {
	as := assert.New(t)

	s := set.NewString()
	s.Add("1")
	s.Add("2")
	s.Add("1")
	as.Equal([]string{"1", "2"}, s.List())

	s.Delete("1")
	as.Equal([]string{"2"}, s.List())
}
