package httputils_test

import (
	"github.com/Chyroc/underscore/utils/httputils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJoinURL(t *testing.T) {
	as := assert.New(t)

	{
		uri, err := httputils.JoinURL("http://a.com/1/2.html", "/3/")
		as.Nil(err)
		as.Equal("http://a.com/3/", uri)
	}

	{
		uri, err := httputils.JoinURL("http://a.com/1/2.html", "3.html")
		as.Nil(err)
		as.Equal("http://a.com/1/3.html", uri)
	}
}
