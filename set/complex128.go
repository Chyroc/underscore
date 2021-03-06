// Code generated by underscore_generate. DO NOT EDIT.

package set

import "sync"

type Complex128Set interface {
	Exist(i complex128) bool
	Add(i complex128)
	Adds(list ...complex128)
	Delete(i complex128)
	Deletes(list ...complex128)
	List() []complex128
}

var _ Complex128Set = (*complex128Set)(nil)

type complex128Set struct {
	m sync.Map
}

func (s *complex128Set) Exist(i complex128) bool {
	_, ok := s.m.Load(i)
	return ok
}

func (s *complex128Set) Add(i complex128) {
	s.m.Store(i, struct{}{})
}

func (s *complex128Set) Adds(list ...complex128) {
	for _, v := range list {
		s.m.Store(v, struct{}{})
	}
}

func (s *complex128Set) Delete(i complex128) {
	s.m.Delete(i)
}

func (s *complex128Set) Deletes(list ...complex128) {
	for _, v := range list {
		s.m.Delete(v)
	}
}

func (s *complex128Set) List() []complex128 {
	var list []complex128
	s.m.Range(func(key, value interface{}) bool {
		list = append(list, key.(complex128))
		return true
	})
	return list
}

func NewComplex128() Complex128Set {
	return &complex128Set{}
}
