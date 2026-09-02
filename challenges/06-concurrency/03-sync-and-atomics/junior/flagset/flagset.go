// Package flagset - Gopher Workplace challenge.
package flagset

import "sync"

// FlagSet holds feature flags: read on every request, written rarely.
type FlagSet struct {
	mu    sync.RWMutex
	flags map[string]bool
}

// NewFlagSet returns an empty flag set.
func NewFlagSet() *FlagSet {
	return &FlagSet{flags: make(map[string]bool)}
}

// Set records the state of a feature flag.
//
// Examples:
//
//	f.Set("new_ui", true); f.Enabled("new_ui") => true
func (f *FlagSet) Set(name string, on bool) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Enabled reports whether the named flag is on. Unknown flags are off.
//
// Examples:
//
//	f.Enabled("unknown") => false
func (f *FlagSet) Enabled(name string) bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Len returns the number of configured flags.
//
// Examples:
//
//	f.Set("a", true); f.Set("b", false); f.Len() => 2
func (f *FlagSet) Len() int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
