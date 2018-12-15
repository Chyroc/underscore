package encodingutls_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Chyroc/underscore/utils/encodingutls"
)

func TestGbkToUtf8(t *testing.T) {
	as := assert.New(t)

	gbk := []byte{181, 218, 210, 187, 213, 194, 32, 203, 240, 187, 181, 181, 196, 207, 181, 205, 179}
	utf8, err := encodingutls.GbkToUtf8(gbk)
	as.Nil(err)
	as.Equal("第一章 损坏的系统", string(utf8))

	gbk2, err := encodingutls.Utf8ToGbk([]byte("第一章 损坏的系统"))
	as.Nil(err)
	as.Equal(gbk, gbk2)
}
