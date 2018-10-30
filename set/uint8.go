package set

import "sync"

type Uint8Set interface {
	Exist(i uint8) bool
	Add(i uint8)
	Adds(list ...uint8)
	Delete(i uint8)
	Deletes(list ...uint8)
	List() []uint8
}

var _ Uint8Set = (*uint8Set)(nil)

type uint8Set struct {
	m sync.Map
}

func (s *uint8Set) Exist(i uint8) bool {
	_, ok := s.m.Load(i)
	return ok
}

func (s *uint8Set) Add(i uint8) {
	s.m.Store(i, struct{}{})
}

func (s *uint8Set) Adds(list ...uint8) {
	for _, v := range list {
		s.m.Store(v, struct{}{})
	}
}

func (s *uint8Set) Delete(i uint8) {
	s.m.Delete(i)
}

func (s *uint8Set) Deletes(list ...uint8) {
	for _, v := range list {
		s.m.Delete(v)
	}
}

func (s *uint8Set) List() []uint8 {
	var list []uint8
	s.m.Range(func(key, value interface{}) bool {
		list = append(list, key.(uint8))
		return true
	})
	return list
}

func NewUint8() Uint8Set {
	return &uint8Set{}
}
