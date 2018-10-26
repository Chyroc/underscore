// +build go1.9

package syncmap

import "sync"

func New() Map {
	return new(sync.Map)
}
