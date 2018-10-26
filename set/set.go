package set

import "sync"

type set struct {
	m sync.Map
}
