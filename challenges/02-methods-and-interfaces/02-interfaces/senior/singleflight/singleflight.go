// Package singleflight — Gopher Workplace challenge.
package singleflight

import "sync"

// Loader fetches a value for a key.
type Loader interface {
	Load(key string) string
}

// LoaderFunc adapts a function to Loader.
type LoaderFunc func(string) string

// Load calls the underlying function.
func (f LoaderFunc) Load(key string) string { return f(key) }

// call is one in-flight load.
type call struct {
	wg  sync.WaitGroup
	val string
}

// Group collapses concurrent loads of the same key into one.
type Group struct {
	mu    sync.Mutex
	calls map[string]*call
}

// NewGroup returns an empty group.
func NewGroup() *Group {
	return &Group{calls: make(map[string]*call)}
}

// Do runs l.Load(key), sharing the result with concurrent callers for the
// same key.
//
// Examples:
//
//	100 concurrent Do calls for one key => the loader runs once
func (g *Group) Do(key string, l Loader) string {
	// TODO(candidate): join an in-flight call, or become the leader.
	panic("not implemented")
}
