// Package storage — Gopher Workplace challenge.
package storage

// Store keeps key/value pairs.
type Store interface {
	Put(key, value string)
	Get(key string) (string, bool)
}

// MemStore keeps everything in a map.
type MemStore struct {
	data map[string]string
}

// NewMemStore returns an empty, ready-to-use store.
func NewMemStore() *MemStore {
	return &MemStore{data: make(map[string]string)}
}

// Put stores a value under key.
func (m *MemStore) Put(key, value string) {
	// TODO(candidate): write into the map.
	panic("not implemented")
}

// Get reads a value; ok is false when the key is absent.
//
// Examples:
//
//	s.Put("a", "1"); s.Get("a")   => "1", true
//	s.Get("missing")              => "", false
func (m *MemStore) Get(key string) (string, bool) {
	// TODO(candidate): comma-ok map read.
	panic("not implemented")
}

// Copy moves the listed keys from src to dst and returns how many were copied.
func Copy(src, dst Store, keys []string) int {
	// TODO(candidate): copy the keys that exist.
	panic("not implemented")
}
