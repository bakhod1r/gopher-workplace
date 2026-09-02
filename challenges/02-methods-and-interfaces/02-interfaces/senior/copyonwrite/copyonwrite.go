// Package copyonwrite — Gopher Workplace challenge.
package copyonwrite

import (
	"sync"
	"sync/atomic"
)

// Snapshot is an immutable settings map.
type Snapshot map[string]string

// Mutator produces the next snapshot from a copy of the current one.
type Mutator interface {
	Mutate(next Snapshot)
}

// MutatorFunc adapts a function to Mutator.
type MutatorFunc func(Snapshot)

// Mutate calls the underlying function.
func (f MutatorFunc) Mutate(next Snapshot) { f(next) }

// Config holds a snapshot that readers load without locking.
type Config struct {
	v  atomic.Value
	mu sync.Mutex // serialises writers only
}

// NewConfig returns a config holding an empty snapshot.
func NewConfig() *Config {
	c := &Config{}
	c.v.Store(Snapshot{})
	return c
}

// Load returns the current snapshot. Callers must not modify it.
func (c *Config) Load() Snapshot {
	// TODO(candidate): atomic load.
	panic("not implemented")
}

// Store publishes a snapshot.
func (c *Config) Store(s Snapshot) {
	// TODO(candidate): atomic store.
	panic("not implemented")
}

// Update applies m to a copy of the current snapshot and publishes the result.
//
// Examples:
//
//	Update(adding "k") => a new snapshot; the previously loaded one is unchanged
func (c *Config) Update(m Mutator) {
	// TODO(candidate): copy, mutate the copy, publish.
	panic("not implemented")
}
