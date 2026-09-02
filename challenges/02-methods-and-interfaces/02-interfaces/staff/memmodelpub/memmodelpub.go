// Package memmodelpub — Gopher Workplace challenge.
package memmodelpub

import "sync/atomic"

// Config is published as an immutable object.
type Config struct {
	Name    string
	Version int
	Tags    []string
}

// Loader hands out the current config.
type Loader interface {
	Load() (*Config, bool)
}

// Publisher publishes config objects to lock-free readers.
type Publisher struct {
	cur atomic.Pointer[Config]
}

// Publish makes cfg visible to readers.
//
// Every write that initialised cfg must happen-before any reader observes it.
func (p *Publisher) Publish(cfg *Config) {
	// TODO(candidate): one atomic store of the fully built pointer.
	panic("not implemented")
}

// Load returns the current config, if any.
//
// Examples:
//
//	before any Publish => nil, false
func (p *Publisher) Load() (*Config, bool) {
	// TODO(candidate): atomic load; report whether anything is published.
	panic("not implemented")
}

// Ready reports whether a config has been published.
func (p *Publisher) Ready() bool {
	// TODO(candidate): report publication state.
	panic("not implemented")
}

// BuildAndPublish constructs a config and publishes it safely.
func BuildAndPublish(p *Publisher, name string, version int, tags []string) {
	// TODO(candidate): fully initialise first, publish last.
	panic("not implemented")
}
