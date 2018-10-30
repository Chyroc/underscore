package set

import "sync"

type Int16Set interface {
	Exist(i int16) bool
	Add(i int16)
	Adds(list ...int16)
	Delete(i int16)
	Deletes(list ...int16)
	List() []int16
}

var _ Int16Set = (*int16Set)(nil)

type int16Set struct {
	m sync.Map
}

func (s *int16Set) Exist(i int16) bool {
	_, ok := s.m.Load(i)
	return ok
}

func (s *int16Set) Add(i int16) {
	s.m.Store(i, struct{}{})
}

func (s *int16Set) Adds(list ...int16) {
	for _, v := range list {
		s.m.Store(v, struct{}{})
	}
}

func (s *int16Set) Delete(i int16) {
	s.m.Delete(i)
}

func (s *int16Set) Deletes(list ...int16) {
	for _, v := range list {
		s.m.Delete(v)
	}
}

func (s *int16Set) List() []int16 {
	var list []int16
	s.m.Range(func(key, value interface{}) bool {
		list = append(list, key.(int16))
		return true
	})
	return list
}

func NewInt16() Int16Set {
	return &int16Set{}
}
