package jsonutils

import "fmt"

func DecodeJavascriptJSON(s string) []map[string]interface{} {
	r := &javascriptJSON{s: []rune(s)}
	return r.run()
}

type javascriptJSON struct {
	s         []rune
	startList bool
	startDict bool
	startStmt bool
	index     int
}

func (r *javascriptJSON) run() []map[string]interface{} {
	var result []map[string]interface{}
	var m = make(map[string]interface{})
	var key string
	for r.index < len(r.s) {
		c := r.s[r.index]
		switch c {
		case '[':
			if r.startList {
				panic(fmt.Sprintf("start list, but got [, cur: %d", r.index))
			}
			r.startList = true
			r.index++
		case '{':
			if r.startDict {
				panic(fmt.Sprintf("start dict, but got {, cur: %d", r.index))
			}
			r.startDict = true
			r.index++
		case ']':
			if !r.startList {
				panic(fmt.Sprintf("end list, but got ], cur: %d", r.index))
			}
			r.startList = false
			r.index++
		case '}':
			if !r.startDict {
				panic(fmt.Sprintf("end dict, but got }, cur: %d", r.index))
			}
			result = append(result, m)
			m = make(map[string]interface{})
			r.startDict = false
			r.index++
		case '\'':
			r.index++
			text := r.findRawText('\'')
			r.startStmt = true
			r.index++
			m[key] = text
		case ' ':
			// pass
			r.index++
		case ',':
			r.index++
		default:
			text := r.findRawText(':')
			r.index++
			key = text
		}
	}

	return result
}

func (r *javascriptJSON) findRawText(sep rune) string {
	var s []rune
	for r.index < len(r.s) {
		c := r.s[r.index]
		switch c {
		case sep:
			return string(s)
		case '\\':
			r.index++
			if r.index < len(r.s) {
				s = append(s, r.s[r.index])
			}
		default:
			s = append(s, c)
		}
		r.index++
	}

	return string(s)
}
