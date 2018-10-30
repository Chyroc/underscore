package main

import (
	"errors"
	"github.com/Chyroc/underscore/utils"
	"os"
	"strings"
)

func GenerateStack(pkg, typ, output string) error {
	return generateStackFactory("stack")(pkg, typ, output)
}

func generateStackFactory(schema string) func(pkg, typ, output string) error {
	var tmpl = ""
	switch schema {
	case "set":
		tmpl = setTmpl
	case "stack":
		tmpl = stackTmpl
	default:
		return func(pkg, typ, output string) error {
			return errors.New("invalid schema: " + schema)
		}
	}
	return func(pkg, typ, output string) error {
		typeLower := typ
		if typ == "interface{}" {
			typeLower = "interface"
		}
		bs, err := utils.ParseTmpl(tmpl, map[string]interface{}{
			"pkg":        pkg,
			"type_upper": strings.Title(typeLower),
			"type_lower": typeLower,
			"var_type":   typ,
		})
		if err != nil {
			return err
		}
		f, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.Write(bs)
		return err
	}
}

var stackTmpl = `package {{.pkg}}

import (
	"fmt"
	"strings"
)

type {{.type_upper}}Stack interface {
	fmt.Stringer

	Push(i {{.var_type}})
	Pop() {{.var_type}}
	Peek() {{.var_type}}
	Len() int
	IsEmpty() bool
	Clone() {{.type_upper}}Stack
}

func New{{.type_upper}}() {{.type_upper}}Stack {
	return &{{.type_lower}}Stack{}
}

var _ {{.type_upper}}Stack = (*{{.type_lower}}Stack)(nil)

type {{.type_lower}}Stack struct {
	is []{{.var_type}}
}

func (s *{{.type_lower}}Stack) String() string {
	var buf = new(strings.Builder)
	for i, v := range s.is {
		if i == 0 {
			buf.WriteString(fmt.Sprintf("%v", v))
		} else {
			buf.WriteString(fmt.Sprintf(" < %v", v))
		}
	}
	return buf.String()
}

func (s *{{.type_lower}}Stack) Push(i {{.var_type}}) {
	s.is = append(s.is, i)
}

func (s *{{.type_lower}}Stack) Pop() {{.var_type}} {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	p := s.Peek()
	s.is = s.is[:len(s.is)-1]
	return p
}

func (s *{{.type_lower}}Stack) Peek() {{.var_type}} {
	if s.IsEmpty() {
		panic("stack is empty")
	}
	return s.is[len(s.is)-1]
}

func (s *{{.type_lower}}Stack) Len() int {
	return len(s.is)
}

func (s *{{.type_lower}}Stack) IsEmpty() bool {
	return len(s.is) == 0
}

func (s *{{.type_lower}}Stack) Clone() {{.type_upper}}Stack {
	s2 := &{{.type_lower}}Stack{is: make([]{{.var_type}}, 0, len(s.is))}
	for _, v := range s.is {
		s2.is = append(s2.is, v)
	}
	return s2
}
`
