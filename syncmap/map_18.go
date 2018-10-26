// +build !go1.9

// <= go1.8 version

package syncmap

type syncMap struct {
}

func (m *syncMap) Delete(key interface{}) {
	panic("implement me")
}

func (m *syncMap) Range(f func(key, value interface{}) bool) {
	panic("implement me")
}

func (m *syncMap) Store(key, value interface{}) {
	panic("implement me")
}

func (m *syncMap) Load(key interface{}) (value interface{}, ok bool) {
	panic("implement me")
}

func (m *syncMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	panic("implement me")
}

func New() Map {
	return new(syncMap)
}
