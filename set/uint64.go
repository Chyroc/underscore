package set

import "sync"

type Uint64Set interface {
	Exist(i uint64) bool
	Add(i uint64)
	Adds(list ...uint64)
	Delete(i uint64)
	Deletes(list ...uint64)
	List() []uint64
}

var _ Uint64Set = (*uint64Set)(nil)

type uint64Set struct {
	m sync.Map
}

func (s *uint64Set) Exist(i uint64) bool {
	_, ok := s.m.Load(i)
	return ok
}

func (s *uint64Set) Add(i uint64) {
	s.m.Store(i, struct{}{})
}

func (s *uint64Set) Adds(list ...uint64) {
	for _, v := range list {
		s.m.Store(v, struct{}{})
	}
}

func (s *uint64Set) Delete(i uint64) {
	s.m.Delete(i)
}

func (s *uint64Set) Deletes(list ...uint64) {
	for _, v := range list {
		s.m.Delete(v)
	}
}

func (s *uint64Set) List() []uint64 {
	var list []uint64
	s.m.Range(func(key, value interface{}) bool {
		list = append(list, key.(uint64))
		return true
	})
	return list
}

func NewUint64() Uint64Set {
	return &uint64Set{}
}