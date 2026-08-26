// Package onceinit — Gopher Workplace challenge.
package onceinit

import "sync"

// LazyData initializes thread-safely exactly once.
type LazyData struct {
	once sync.Once
	data string
	init func() string
}

// New creates a LazyData.
func New(init func() string) *LazyData {
	return &LazyData{init: init}
}

// Get returns the initialized data.
func (l *LazyData) Get() string {
	// TODO(candidate): use sync.Once to call init() exactly once,
	// store result in l.data, return l.data.
	panic("not implemented")
}
