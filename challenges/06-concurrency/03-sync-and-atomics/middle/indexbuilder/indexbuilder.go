// Package indexbuilder — Gopher Workplace challenge.
package indexbuilder

import "sync"

// Registry compiles a search index per collection. Compiling is expensive, so
// when several request goroutines ask for the same collection at once they
// must share a single build rather than each starting their own.
type Registry struct {
	mu     sync.Mutex
	builds map[string]*sync.Once
	ready  map[string]string
	build  func(collection string) string
}

// NewRegistry returns a Registry that compiles indexes with build.
//
// Examples:
//
//	NewRegistry(func(c string) string { return c + "-idx" }) != nil => true
func NewRegistry(build func(collection string) string) *Registry {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Index returns the compiled index for a collection, building it on first use.
// build runs exactly once per collection no matter how many goroutines ask.
//
// Examples:
//
//	r := NewRegistry(f); r.Index("orders")               => "orders-idx"
//	r.Index("orders"); r.Index("orders")                 => build ran once
func (r *Registry) Index(collection string) string {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Built returns the collections whose index has been compiled.
func (r *Registry) Built() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return len(r.ready)
}
