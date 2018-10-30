package set

import "sync"

type Int8Set interface {
	Exist(i int8) bool
	Add(i int8)
	Adds(list ...int8)
	Delete(i int8)
	Deletes(list ...int8)
	List() []int8
}

var _ Int8Set = (*int8Set)(nil)

type int8Set struct {
	m sync.Map
}

func (s *int8Set) Exist(i int8) bool {
	_, ok := s.m.Load(i)
	return ok
}

func (s *int8Set) Add(i int8) {
	s.m.Store(i, struct{}{})
}

func (s *int8Set) Adds(list ...int8) {
	for _, v := range list {
		s.m.Store(v, struct{}{})
	}
}

func (s *int8Set) Delete(i int8) {
	s.m.Delete(i)
}

func (s *int8Set) Deletes(list ...int8) {
	for _, v := range list {
		s.m.Delete(v)
	}
}

func (s *int8Set) List() []int8 {
	var list []int8
	s.m.Range(func(key, value interface{}) bool {
		list = append(list, key.(int8))
		return true
	})
	return list
}

func NewInt8() Int8Set {
	return &int8Set{}
}
