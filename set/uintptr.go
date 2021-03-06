// Code generated by underscore_generate. DO NOT EDIT.

package set

import "sync"

type UintptrSet interface {
	Exist(i uintptr) bool
	Add(i uintptr)
	Adds(list ...uintptr)
	Delete(i uintptr)
	Deletes(list ...uintptr)
	List() []uintptr
}

var _ UintptrSet = (*uintptrSet)(nil)

type uintptrSet struct {
	m sync.Map
}

func (s *uintptrSet) Exist(i uintptr) bool {
	_, ok := s.m.Load(i)
	return ok
}

func (s *uintptrSet) Add(i uintptr) {
	s.m.Store(i, struct{}{})
}

func (s *uintptrSet) Adds(list ...uintptr) {
	for _, v := range list {
		s.m.Store(v, struct{}{})
	}
}

func (s *uintptrSet) Delete(i uintptr) {
	s.m.Delete(i)
}

func (s *uintptrSet) Deletes(list ...uintptr) {
	for _, v := range list {
		s.m.Delete(v)
	}
}

func (s *uintptrSet) List() []uintptr {
	var list []uintptr
	s.m.Range(func(key, value interface{}) bool {
		list = append(list, key.(uintptr))
		return true
	})
	return list
}

func NewUintptr() UintptrSet {
	return &uintptrSet{}
}
