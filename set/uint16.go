package set

import "sync"

type Uint16Set interface {
	Exist(i uint16) bool
	Add(i uint16)
	Adds(list ...uint16)
	Delete(i uint16)
	Deletes(list ...uint16)
	List() []uint16
}

var _ Uint16Set = (*uint16Set)(nil)

type uint16Set struct {
	m sync.Map
}

func (s *uint16Set) Exist(i uint16) bool {
	_, ok := s.m.Load(i)
	return ok
}

func (s *uint16Set) Add(i uint16) {
	s.m.Store(i, struct{}{})
}

func (s *uint16Set) Adds(list ...uint16) {
	for _, v := range list {
		s.m.Store(v, struct{}{})
	}
}

func (s *uint16Set) Delete(i uint16) {
	s.m.Delete(i)
}

func (s *uint16Set) Deletes(list ...uint16) {
	for _, v := range list {
		s.m.Delete(v)
	}
}

func (s *uint16Set) List() []uint16 {
	var list []uint16
	s.m.Range(func(key, value interface{}) bool {
		list = append(list, key.(uint16))
		return true
	})
	return list
}

func NewUint16() Uint16Set {
	return &uint16Set{}
}
