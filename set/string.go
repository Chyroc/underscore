package set

type String interface {
	Add(i string)
	Delete(i string)

	List() []string
}

func (s *set) Add(i string) {
	s.m.Store(i, struct{}{})
}

func (s *set) Delete(i string) {
	s.m.Delete(i)
}

func (s *set) List() []string {
	var list []string
	s.m.Range(func(key, value interface{}) bool {
		list = append(list, key.(string))
		return true
	})
	return list
}

func NewString() String {
	return &set{}
}
