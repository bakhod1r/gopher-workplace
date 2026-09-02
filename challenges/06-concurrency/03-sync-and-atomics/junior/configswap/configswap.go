// Package configswap - Gopher Workplace challenge.
package configswap

import "sync/atomic"

// Config is a service configuration snapshot.
type Config struct {
	Version int
	Region  string
}

// Store publishes configuration snapshots to concurrent readers.
type Store struct {
	v atomic.Value
}

// Store publishes c as the current configuration.
//
// Examples:
//
//	var s Store; s.Store(Config{Version: 2, Region: "eu"}); s.Load().Region => "eu"
func (s *Store) Store(c Config) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Load returns the current configuration, or the zero Config if none.
//
// Examples:
//
//	var s Store; s.Load() => Config{}
func (s *Store) Load() Config {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Version returns the published configuration version.
//
// Examples:
//
//	s.Store(Config{Version: 3}); s.Version() => 3
func (s *Store) Version() int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
