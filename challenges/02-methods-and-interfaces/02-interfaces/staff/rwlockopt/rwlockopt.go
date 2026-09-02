// Package rwlockopt — Gopher Workplace challenge.
package rwlockopt

import "sync"

// Store is a concurrent key/value map.
type Store interface {
	Get(key string) (int, bool)
	Set(key string, v int)
	Len() int
}

// RWStore uses a read-write mutex.
type RWStore struct {
	mu   sync.RWMutex
	data map[string]int

	readers  int
	maxRead  int
	readerMu sync.Mutex
}

// NewRWStore returns an empty store.
func NewRWStore() *RWStore {
	return &RWStore{data: make(map[string]int)}
}

// Get reads a key under a read lock.
//
// Examples:
//
//	Set("a", 1); Get("a") => 1, true
func (s *RWStore) Get(key string) (int, bool) {
	// TODO(candidate): read under RLock; track concurrent readers.
	panic("not implemented")
}

// Set writes a key under a write lock.
func (s *RWStore) Set(key string, v int) {
	// TODO(candidate): write under Lock.
	panic("not implemented")
}

// Len returns the number of keys.
func (s *RWStore) Len() int {
	// TODO(candidate): read under RLock.
	panic("not implemented")
}

// Snapshot returns an independent copy taken under a read lock.
func (s *RWStore) Snapshot() map[string]int {
	// TODO(candidate): copy under RLock.
	panic("not implemented")
}

// MaxConcurrentReaders reports the highest number of readers seen inside
// Get at the same time.
func (s *RWStore) MaxConcurrentReaders() int {
	s.readerMu.Lock()
	defer s.readerMu.Unlock()
	return s.maxRead
}

// enterRead and exitRead track reader overlap for the test.
func (s *RWStore) enterRead() {
	s.readerMu.Lock()
	s.readers++
	if s.readers > s.maxRead {
		s.maxRead = s.readers
	}
	s.readerMu.Unlock()
}

func (s *RWStore) exitRead() {
	s.readerMu.Lock()
	s.readers--
	s.readerMu.Unlock()
}

// MutexStore uses a plain mutex.
type MutexStore struct {
	mu   sync.Mutex
	data map[string]int
}

// NewMutexStore returns an empty store.
func NewMutexStore() *MutexStore {
	return &MutexStore{data: make(map[string]int)}
}

// Get reads a key.
func (s *MutexStore) Get(key string) (int, bool) {
	// TODO(candidate): read under the lock.
	panic("not implemented")
}

// Set writes a key.
func (s *MutexStore) Set(key string, v int) {
	// TODO(candidate): write under the lock.
	panic("not implemented")
}

// Len returns the number of keys.
func (s *MutexStore) Len() int {
	// TODO(candidate): read under the lock.
	panic("not implemented")
}
