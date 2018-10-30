package set

import "sync"

type StringSet interface {
	Exist(i string) bool
	Add(i string)
	Adds(list ...string)
	Delete(i string)
	Deletes(list ...string)
	List() []string
}

var _ StringSet = (*stringSet)(nil)

type stringSet struct {
	m sync.Map
}

func (s *stringSet) Exist(i string) bool {
	_, ok := s.m.Load(i)
	return ok
}

func (s *stringSet) Add(i string) {
	s.m.Store(i, struct{}{})
}

func (s *stringSet) Adds(list ...string) {
	for _, v := range list {
		s.m.Store(v, struct{}{})
	}
}

func (s *stringSet) Delete(i string) {
	s.m.Delete(i)
}

func (s *stringSet) Deletes(list ...string) {
	for _, v := range list {
		s.m.Delete(v)
	}
}

func (s *stringSet) List() []string {
	var list []string
	s.m.Range(func(key, value interface{}) bool {
		list = append(list, key.(string))
		return true
	})
	return list
}

func NewString() StringSet {
	return &stringSet{}
}
