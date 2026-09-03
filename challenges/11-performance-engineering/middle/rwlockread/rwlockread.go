// Package rwlockread — Gopher Workplace challenge.
package rwlockread

import "sync"

// Config is a read-mostly value: written rarely, read on every request. That
// asymmetry is what an RWMutex is for — any number of readers may hold it at
// once, and a writer excludes them all.
type Config struct {
	mu      sync.RWMutex
	values  map[string]string
	version int
}

// Get returns a value and whether it is present, under a read lock so
// concurrent readers do not serialise.
//
// Examples:
//
//	c.Get("k") => "v", true
func (c *Config) Get(key string) (string, bool) {
	panic("not implemented")
}

// Version returns the current version, also under a read lock.
//
// Examples:
//
//	c.Version() => 1
func (c *Config) Version() int {
	panic("not implemented")
}

// Replace swaps in a whole new configuration under the write lock and bumps
// the version. It copies the map it is given, so the caller cannot mutate the
// live configuration afterwards.
//
// Examples:
//
//	c.Replace(map[string]string{"k": "v"})
func (c *Config) Replace(values map[string]string) {
	panic("not implemented")
}

// Snapshot returns a copy of the values and the version they belong to, taken
// together so the pair is always consistent.
//
// Examples:
//
//	vals, ver := c.Snapshot()
func (c *Config) Snapshot() (map[string]string, int) {
	panic("not implemented")
}
