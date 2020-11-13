package engine

import (
	"sync"
)

var (
	mt sync.RWMutex
)

func readAccess(fn func()) {
	mt.RLock()
	defer mt.RUnlock()

	fn()
}

func writeAccess(fn func()) {
	mt.Lock()
	defer mt.Unlock()

	fn()
}
