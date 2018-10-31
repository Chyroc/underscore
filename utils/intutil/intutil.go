package intutil

import (
	"errors"
	"github.com/Chyroc/readcode/debug"
)

func IsDigital(i byte) bool {
	return i >= '0' && i <= '9'
}

func ParserIntRange(s string) ([]int, error) {
	p := &intRange{
		s:    s,
		list: make([]int, 0, len(s)),
	}
	return p.parse()
}

type intRange struct {
	s     string
	cur   int  // 0,n-1
	comma bool // ,
	bar   bool // -
	list  []int
	//pre      int
	preValid bool
}

func (r *intRange) parse() ([]int, error) {
	var pre int
	var err error
	for r.cur < len(r.s) {
		debug.Println(r.cur, string([]byte{r.s[r.cur]}), r.list)
		switch {
		case IsDigital(r.s[r.cur]):
			//debug.Println(r.cur, pre, r.preValid, err)
			pre, err = r.scanInt()
			//debug.Println(r.cur, pre, r.preValid, err)
			if err != nil {
				return nil, err
			}
		case r.s[r.cur] == '-':
			if r.cur == len(r.s)-1 {
				return nil, errors.New("1]invalid int range string: " + r.s)
			}
			if !r.preValid {
				return nil, errors.New("2]invalid int range string: " + r.s)
			}

			r.cur++
			r.preValid = false
			next, err := r.scanInt()
			if err != nil {
				return nil, err
			}

			for j := pre; j <= next; j++ {
				r.list = append(r.list, j)
			}
			r.preValid = false
			pre = 0
			r.bar = false
			r.cur++
		case r.s[r.cur] == ',':
			//debug.Println("t1", r.cur, len(r.s)-1)
			if r.cur == len(r.s)-1 {
				return r.list, nil
			}
			r.list = append(r.list, pre)
			r.preValid = false
			r.comma = false
			pre = 0
			r.cur++
		default:
			return nil, errors.New("3]invalid int range string: " + r.s)
		}
	}

	if r.preValid {
		r.list = append(r.list, pre)
	}

	return r.list, nil
}

func (r *intRange) scanInt() (int, error) {
	var pre int
	if r.preValid {
		//debug.Println(r.list, r.preValid, r.cur)
		return 0, errors.New("4]invalid int range string: " + r.s)
	}
	for r.cur < len(r.s) {
		switch {
		case IsDigital(r.s[r.cur]):
			pre = joinInt(pre, r.s[r.cur])
			if !r.preValid {
				r.preValid = true
			}
			r.cur++
		case r.s[r.cur] == '-':
			if r.comma || r.bar {
				return 0, errors.New("5]invalid int range string: " + r.s)
			}
			r.bar = true
			return pre, nil
		case r.s[r.cur] == ',':
			if r.comma || r.bar {
				return 0, errors.New("6]invalid int range string: " + r.s)
			}
			r.comma = true
			return pre, nil
		default:
			return 0, errors.New("7]invalid int range string: " + r.s)
		}
	}
	return pre, nil
}

func joinInt(pre int, current byte) int {
	return pre*10 + int(current-'0')
}
