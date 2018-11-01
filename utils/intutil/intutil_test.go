package intutil_test

import (
	"github.com/Chyroc/underscore/utils/intutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

type IntRangeTest struct {
	s    string
	list []int
	err  string
}

func TestParserIntRange(t *testing.T) {
	as := assert.New(t)

	for _, test := range []IntRangeTest{
		//{"1", []int{1}, ""},
		//{"1,2", []int{1, 2}, ""},
		//{"1,2,3", []int{1, 2, 3}, ""},
		//{"1,2,3,4", []int{1, 2, 3, 4}, ""},
		//{"1-4", []int{1, 2, 3, 4}, ""},
		//{"1,2,3,8-10", []int{1, 2, 3, 8, 9, 10}, ""},
		//{"1,2,3,8-10,14,16", []int{1, 2, 3, 8, 9, 10, 14, 16}, ""},
		//{"1,2,3,8-10,14,16,20-30", []int{1, 2, 3, 8, 9, 10, 14, 16, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}, ""},
		//{"1,2,3,8-10,14,16,20-30,99", []int{1, 2, 3, 8, 9, 10, 14, 16, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 99}, ""},
		//
		{"1-1-1",nil,""},
	} {
		list, err := intutil.ParserIntRange(test.s)
		if test.err != "" {
			as.NotNil(err)
			as.Equal(test.err, err.Error())
		} else {
			as.Nil(err)
			as.Equal(test.list, list)
		}
	}

	list, err := intutil.ParserIntRange("1")
	as.Nil(err)
	as.Equal([]int{1}, list)
}
