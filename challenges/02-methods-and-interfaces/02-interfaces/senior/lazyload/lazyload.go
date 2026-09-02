// Package lazyload — Gopher Workplace challenge.
package lazyload

import (
	"sync"
	"sync/atomic"
)

// Builder produces the expensive value.
type Builder interface {
	Build() string
}

// BuilderFunc adapts a function to Builder.
type BuilderFunc func() string

// Build calls the underlying function.
func (f BuilderFunc) Build() string { return f() }

// Lazy builds a value on first use.
type Lazy struct {
	builder Builder
	once    sync.Once
	value   string
	built   atomic.Bool
}

// NewLazy returns a lazy value backed by b.
func NewLazy(b Builder) *Lazy {
	return &Lazy{builder: b}
}

// Get returns the value, building it on the first call.
//
// Examples:
//
//	two Get calls => the builder runs once
func (l *Lazy) Get() string {
	// TODO(candidate): build exactly once, then reuse.
	panic("not implemented")
}

// Built reports whether the value has been built yet.
func (l *Lazy) Built() bool {
	// TODO(candidate): report the flag.
	panic("not implemented")
}
