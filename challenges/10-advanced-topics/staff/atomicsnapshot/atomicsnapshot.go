// Package atomicsnapshot — Gopher Workplace challenge.
package atomicsnapshot

import "sync/atomic"

// Config is one immutable settings snapshot.
type Config struct {
	Retries int
	Timeout int
}

// Store holds the current Config for concurrent readers.
type Store struct {
	v atomic.Pointer[Config]
}

// Get returns the current snapshot, or the zero Config if none is set.
func (s *Store) Get() Config {
	if p := s.v.Load(); p != nil {
		return *p
	}
	return Config{}
}

// Set publishes c as the current configuration.
//
// Readers must see either the old snapshot or the new one, never a mix.
// The snapshot escapes by construction — it outlives the call and is shared
// with every reader.
//
// Examples:
//
//	s.Set(Config{Retries: 3}); s.Get().Retries => 3
func (s *Store) Set(c Config) {
	panic("not implemented")
}
