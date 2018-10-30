package main

import (
	"github.com/Chyroc/underscore/utils"
	"os"
	"strings"
)

func GenerateSet(pkg, typ, output string) error {
	bs, err := utils.ParseTmpl(setTmpl, map[string]interface{}{
		"pkg":        pkg,
		"type_upper": strings.Title(typ),
		"type_lower": typ,
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

var setTmpl = `package {{.pkg}}

import "sync"

type {{.type_upper}}Set interface {
	Exist(i {{.type_lower}}) bool
	Add(i {{.type_lower}})
	Adds(list ...{{.type_lower}})
	Delete(i {{.type_lower}})
	Deletes(list ...{{.type_lower}})
	List() []{{.type_lower}}
}

var _ {{.type_upper}}Set = (*{{.type_lower}}Set)(nil)

type {{.type_lower}}Set struct {
	m sync.Map
}

func (s *{{.type_lower}}Set) Exist(i {{.type_lower}}) bool {
	_, ok := s.m.Load(i)
	return ok
}

func (s *{{.type_lower}}Set) Add(i {{.type_lower}}) {
	s.m.Store(i, struct{}{})
}

func (s *{{.type_lower}}Set) Adds(list ...{{.type_lower}}) {
	for _, v := range list {
		s.m.Store(v, struct{}{})
	}
}

func (s *{{.type_lower}}Set) Delete(i {{.type_lower}}) {
	s.m.Delete(i)
}

func (s *{{.type_lower}}Set) Deletes(list ...{{.type_lower}}) {
	for _, v := range list {
		s.m.Delete(v)
	}
}

func (s *{{.type_lower}}Set) List() []{{.type_lower}} {
	var list []{{.type_lower}}
	s.m.Range(func(key, value interface{}) bool {
		list = append(list, key.({{.type_lower}}))
		return true
	})
	return list
}

func New{{.type_upper}}() {{.type_upper}}Set {
	return &{{.type_lower}}Set{}
}
`
