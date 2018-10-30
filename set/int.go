package set

import "sync"

type IntSet interface {
	Exist(i int) bool
	Add(i int)
	Adds(list ...int)
	Delete(i int)
	Deletes(list ...int)
	List() []int
}

var _ IntSet = (*intSet)(nil)

type intSet struct {
	m sync.Map
}

func (s *intSet) Exist(i int) bool {
	_, ok := s.m.Load(i)
	return ok
}

func (s *intSet) Add(i int) {
	s.m.Store(i, struct{}{})
}

func (s *intSet) Adds(list ...int) {
	for _, v := range list {
		s.m.Store(v, struct{}{})
	}
}

func (s *intSet) Delete(i int) {
	s.m.Delete(i)
}

func (s *intSet) Deletes(list ...int) {
	for _, v := range list {
		s.m.Delete(v)
	}
}

func (s *intSet) List() []int {
	var list []int
	s.m.Range(func(key, value interface{}) bool {
		list = append(list, key.(int))
		return true
	})
	return list
}

func NewInt() IntSet {
	return &intSet{}
}
