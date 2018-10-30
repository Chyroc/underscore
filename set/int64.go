package set

import "sync"

type Int64Set interface {
	Exist(i int64) bool
	Add(i int64)
	Adds(list ...int64)
	Delete(i int64)
	Deletes(list ...int64)
	List() []int64
}

var _ Int64Set = (*int64Set)(nil)

type int64Set struct {
	m sync.Map
}

func (s *int64Set) Exist(i int64) bool {
	_, ok := s.m.Load(i)
	return ok
}

func (s *int64Set) Add(i int64) {
	s.m.Store(i, struct{}{})
}

func (s *int64Set) Adds(list ...int64) {
	for _, v := range list {
		s.m.Store(v, struct{}{})
	}
}

func (s *int64Set) Delete(i int64) {
	s.m.Delete(i)
}

func (s *int64Set) Deletes(list ...int64) {
	for _, v := range list {
		s.m.Delete(v)
	}
}

func (s *int64Set) List() []int64 {
	var list []int64
	s.m.Range(func(key, value interface{}) bool {
		list = append(list, key.(int64))
		return true
	})
	return list
}

func NewInt64() Int64Set {
	return &int64Set{}
}