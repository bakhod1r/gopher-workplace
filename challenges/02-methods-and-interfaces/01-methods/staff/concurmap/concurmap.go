// Package concurmap — Gopher Workplace challenge.
package concurmap

import "sync"

// ConcurrentMap is a thread-safe map.
type ConcurrentMap struct {
	mu   sync.RWMutex
	data map[string]int
}

// New creates a new ConcurrentMap.
func New() *ConcurrentMap {
	return &ConcurrentMap{data: make(map[string]int)}
}

// Get returns the value and whether it exists, using a read lock.
func (m *ConcurrentMap) Get(key string) (int, bool) {
	// TODO(candidate): thread-safe read
	panic("not implemented")
}

// Set sets the value, using a write lock.
func (m *ConcurrentMap) Set(key string, val int) {
	// TODO(candidate): thread-safe write
	panic("not implemented")
}
