package set

import "sync"

type InterfaceSet interface {
	Exist(i interface{}) bool
	Add(i interface{})
	Adds(list ...interface{})
	Delete(i interface{})
	Deletes(list ...interface{})
	List() []interface{}
}

var _ InterfaceSet = (*interfaceSet)(nil)

type interfaceSet struct {
	m sync.Map
}

func (s *interfaceSet) Exist(i interface{}) bool {
	_, ok := s.m.Load(i)
	return ok
}

func (s *interfaceSet) Add(i interface{}) {
	s.m.Store(i, struct{}{})
}

func (s *interfaceSet) Adds(list ...interface{}) {
	for _, v := range list {
		s.m.Store(v, struct{}{})
	}
}

func (s *interfaceSet) Delete(i interface{}) {
	s.m.Delete(i)
}

func (s *interfaceSet) Deletes(list ...interface{}) {
	for _, v := range list {
		s.m.Delete(v)
	}
}

func (s *interfaceSet) List() []interface{} {
	var list []interface{}
	s.m.Range(func(key, value interface{}) bool {
		list = append(list, key.(interface{}))
		return true
	})
	return list
}

func NewInterface() InterfaceSet {
	return &interfaceSet{}
}
