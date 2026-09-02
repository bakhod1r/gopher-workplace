// Package connregistry - Gopher Workplace challenge.
package connregistry

import "sync"

// Registry maps upstream instance IDs to addresses.
type Registry struct {
	mu    sync.RWMutex
	conns map[string]string
}

// NewRegistry returns an empty registry.
func NewRegistry() *Registry {
	return &Registry{conns: make(map[string]string)}
}

// Register records the address of an upstream instance.
//
// Examples:
//
//	r.Register("a", "10.0.0.1"); r.Lookup("a") => "10.0.0.1", true
func (r *Registry) Register(id, addr string) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Lookup returns the address of id and whether it is registered.
//
// Examples:
//
//	NewRegistry().Lookup("ghost") => "", false
func (r *Registry) Lookup(id string) (string, bool) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// IDs returns a copy of the registered instance IDs, in any order.
//
// Examples:
//
//	r.Register("b", "x"); r.Register("a", "y"); len(r.IDs()) => 2
func (r *Registry) IDs() []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
