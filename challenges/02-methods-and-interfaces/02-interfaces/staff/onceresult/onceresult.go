// Package onceresult — Gopher Workplace challenge.
package onceresult

import (
	"fmt"
	"sync"
)

// Initer produces a value or an error.
type Initer interface {
	Init() (string, error)
}

// IniterFunc adapts a function to Initer.
type IniterFunc func() (string, error)

// Init calls the underlying function.
func (f IniterFunc) Init() (string, error) { return f() }

// OnceValue runs an initialiser at most once and caches its result.
type OnceValue struct {
	init  Initer
	once  sync.Once
	value string
	err   error
	runs  int
}

// NewOnceValue returns a lazily initialised value.
func NewOnceValue(i Initer) *OnceValue {
	return &OnceValue{init: i}
}

// Get returns the cached value, running the initialiser on the first call.
//
// A panic inside the initialiser is captured and cached as an error.
//
// Examples:
//
//	two Get calls => the initialiser runs once
func (o *OnceValue) Get() (string, error) {
	// TODO(candidate): run once, recover panics, cache both outcomes.
	panic("not implemented")
}

// Runs reports how many times the initialiser actually ran.
func (o *OnceValue) Runs() int {
	// TODO(candidate): report the run count.
	panic("not implemented")
}

var _ = fmt.Errorf
