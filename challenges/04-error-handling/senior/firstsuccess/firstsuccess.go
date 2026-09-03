// Package firstsuccess — Gopher Workplace challenge.
package firstsuccess

import "errors"

// ErrNoSources reports an empty source list.
var ErrNoSources = errors.New("no sources")

// First returns the value from the first source that succeeds.
//
// Examples:
//
//	First() => 0, ErrNoSources
func First(sources ...func() (int, error)) (int, error) {
	// TODO(candidate): implement this.
	_ = errors.Join
	panic("not implemented")
}
