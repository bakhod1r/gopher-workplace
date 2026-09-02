// Package lookupkey — Gopher Workplace challenge.
package lookupkey

import "errors"

// ErrNotFound reports a missing key.
var ErrNotFound = errors.New("key not found")

// Lookup returns m[key], or ErrNotFound when the key is absent.
//
// Examples:
//
//	Lookup(map[string]int{"a": 1}, "a") => 1, nil
//	Lookup(map[string]int{"a": 1}, "z") => 0, ErrNotFound
func Lookup(m map[string]int, key string) (int, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
