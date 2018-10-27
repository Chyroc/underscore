package regutils_test

import (
	"github.com/Chyroc/underscore/utils/regutils"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestFindSubmatch(t *testing.T) {
	as := assert.New(t)
	r := regexp.MustCompile(`a = "(.*?)"`)

	{
		s, err := regutils.FindSubmatch("a = \"test\"", r)
		as.Nil(err)
		as.Equal("test", s)
	}

	{
		s, err := regutils.FindSubmatch("b = \"test\"", r)
		as.NotNil(err)
		as.Equal("reg: a = \"(.*?)\" missed", err.Error())
		as.Empty(s)
	}
}
