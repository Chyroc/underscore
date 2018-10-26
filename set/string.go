package set

type StringSet interface {
	Exist(i string) bool
	Add(i string)
	Adds(list ...string)
	Delete(i string)
	Deletes(list ...string)
	List() []string
}

var _ StringSet = (*set)(nil)

func (s *set) Exist(i string) bool {
	_, ok := s.m.Load(i)
	return ok
}

func (s *set) Add(i string) {
	s.m.Store(i, struct{}{})
}

func (s *set) Adds(list ...string) {
	for _, v := range list {
		s.m.Store(v, struct{}{})
	}
}

func (s *set) Delete(i string) {
	s.m.Delete(i)
}

func (s *set) Deletes(list ...string) {
	for _, v := range list {
		s.m.Delete(v)
	}
}

func (s *set) List() []string {
	var list []string
	s.m.Range(func(key, value interface{}) bool {
		list = append(list, key.(string))
		return true
	})
	return list
}

func NewString() StringSet {
	return &set{}
}
