// Package mapwrap — Gopher Workplace challenge.
package mapwrap

import (
	"errors"
	"fmt"
)

// ErrNotFound is the shared lookup failure.
var ErrNotFound = errors.New("not found")

// Find returns m[key], annotating a miss with the key.
//
// Examples:
//
//	Find(map[string]int{"a": 1}, "a") => 1, nil
//	Find(nil, "a")                    => 0, "key a: not found"
func Find(m map[string]int, key string) (int, error) {
	// TODO(candidate): implement this.
	_ = fmt.Errorf
	panic("not implemented")
}
