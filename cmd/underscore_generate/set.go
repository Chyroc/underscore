package main

func GenerateSet(pkg, typ, output string) error {
	return generateStackFactory("set")(pkg, typ, output)
}

var setTmpl = `package {{.pkg}}

import "sync"

type {{.type_upper}}Set interface {
	Exist(i {{.var_type}}) bool
	Add(i {{.var_type}})
	Adds(list ...{{.var_type}})
	Delete(i {{.var_type}})
	Deletes(list ...{{.var_type}})
	List() []{{.var_type}}
}

var _ {{.type_upper}}Set = (*{{.type_lower}}Set)(nil)

type {{.type_lower}}Set struct {
	m sync.Map
}

func (s *{{.type_lower}}Set) Exist(i {{.var_type}}) bool {
	_, ok := s.m.Load(i)
	return ok
}

func (s *{{.type_lower}}Set) Add(i {{.var_type}}) {
	s.m.Store(i, struct{}{})
}

func (s *{{.type_lower}}Set) Adds(list ...{{.var_type}}) {
	for _, v := range list {
		s.m.Store(v, struct{}{})
	}
}

func (s *{{.type_lower}}Set) Delete(i {{.var_type}}) {
	s.m.Delete(i)
}

func (s *{{.type_lower}}Set) Deletes(list ...{{.var_type}}) {
	for _, v := range list {
		s.m.Delete(v)
	}
}

func (s *{{.type_lower}}Set) List() []{{.var_type}} {
	var list []{{.var_type}}
	s.m.Range(func(key, value interface{}) bool {
		list = append(list, key.({{.var_type}}))
		return true
	})
	return list
}

func New{{.type_upper}}() {{.type_upper}}Set {
	return &{{.type_lower}}Set{}
}
`
