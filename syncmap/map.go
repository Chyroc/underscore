package syncmap

type Map interface {
	Delete(key interface{})
	Range(f func(key, value interface{}) bool)
	Store(key, value interface{})
	Load(key interface{}) (value interface{}, ok bool)
	LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
}
