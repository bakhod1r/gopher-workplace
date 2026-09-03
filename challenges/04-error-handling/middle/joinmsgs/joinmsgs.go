// Package joinmsgs — Gopher Workplace challenge.
package joinmsgs

import "errors"

// Validation failures used by the tests.
var (
	ErrA = errors.New("a")
	ErrB = errors.New("b")
)

// Lines returns one message per error inside err.
//
// Examples:
//
//	Lines(errors.Join(ErrA, ErrB)) => ["a" "b"]
//	Lines(ErrA)                    => ["a"]
func Lines(err error) []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
