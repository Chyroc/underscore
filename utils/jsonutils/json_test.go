package jsonutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeJavascriptJSON(t *testing.T) {
	as := assert.New(t)

	var r *javascriptJSON
	{
		r = &javascriptJSON{s: []rune(`text'`)}
		as.Equal("text", r.findRawText('\''))
		as.Equal(4, r.index)

		r = &javascriptJSON{s: []rune(`text'text`)}
		as.Equal("text", r.findRawText('\''))
		as.Equal(4, r.index)

		r = &javascriptJSON{s: []rune(`'text'`)}
		as.Equal("", r.findRawText('\''))
		as.Equal("", r.findRawText('\''))
		as.Equal(0, r.index)
	}

	{

		as.Equal([]map[string]interface{}{{"key": "value"}}, DecodeJavascriptJSON(`[{key:'value'}]`))
		as.Equal([]map[string]interface{}{{"key": "value"}}, DecodeJavascriptJSON(`[{key: 'value'}]`))

		as.Equal([]map[string]interface{}{{"k1": "v1"}, {"k2": "v2"}}, DecodeJavascriptJSON(`[{k1: 'v1'}, {k2:'v2'}]`))
		as.Equal([]map[string]interface{}{{"k1": "v1", "k2": "v2"}}, DecodeJavascriptJSON(`[{k1: 'v1', k2:'v2'}]`))
	}
}
