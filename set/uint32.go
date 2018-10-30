package set

import "sync"

type Uint32Set interface {
	Exist(i uint32) bool
	Add(i uint32)
	Adds(list ...uint32)
	Delete(i uint32)
	Deletes(list ...uint32)
	List() []uint32
}

var _ Uint32Set = (*uint32Set)(nil)

type uint32Set struct {
	m sync.Map
}

func (s *uint32Set) Exist(i uint32) bool {
	_, ok := s.m.Load(i)
	return ok
}

func (s *uint32Set) Add(i uint32) {
	s.m.Store(i, struct{}{})
}

func (s *uint32Set) Adds(list ...uint32) {
	for _, v := range list {
		s.m.Store(v, struct{}{})
	}
}

func (s *uint32Set) Delete(i uint32) {
	s.m.Delete(i)
}

func (s *uint32Set) Deletes(list ...uint32) {
	for _, v := range list {
		s.m.Delete(v)
	}
}

func (s *uint32Set) List() []uint32 {
	var list []uint32
	s.m.Range(func(key, value interface{}) bool {
		list = append(list, key.(uint32))
		return true
	})
	return list
}

func NewUint32() Uint32Set {
	return &uint32Set{}
}
