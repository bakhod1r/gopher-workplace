// Package copyonwrite — Gopher Workplace challenge.
package copyonwrite

import (
	"sync"
	"sync/atomic"
)

// Store is a read-mostly map published by pointer swap.
type Store struct {
	mu sync.Mutex // serialises writers only
	m  atomic.Pointer[map[string]int]
}

// Get reads from the current snapshot without locking.
func (s *Store) Get(key string) (int, bool) {
	p := s.m.Load()
	if p == nil {
		return 0, false
	}
	v, ok := (*p)[key]
	return v, ok
}

// Len reports the current snapshot's size.
func (s *Store) Len() int {
	p := s.m.Load()
	if p == nil {
		return 0
	}
	return len(*p)
}

// Set publishes a new snapshot of the map with key set to val.
//
// Readers hold whatever snapshot was current when they loaded it, so a
// published map must never be modified again: build a copy, then swap.
//
// Examples:
//
//	s.Set("a", 1); s.Get("a") => 1, true
func (s *Store) Set(key string, val int) {
	panic("not implemented")
}
