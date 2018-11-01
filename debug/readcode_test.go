package debug_test

import (
	"bytes"
	"github.com/Chyroc/underscore/debug"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintln(t *testing.T) {
	as := assert.New(t)

	buf := new(bytes.Buffer)
	debug.Fprintln(buf, "test", 2)
	as.Equal("1 readcode_test.go:14 TestPrintln test 2\n", buf.String())
}
