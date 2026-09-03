// Package joinsplit — Gopher Workplace challenge.
package joinsplit

import "errors"

// Rule failures used by the tests.
var (
	ErrA = errors.New("rule a failed")
	ErrB = errors.New("rule b failed")
)

// Split returns the individual errors inside a joined error.
//
// Examples:
//
//	Split(errors.Join(ErrA, ErrB)) => [ErrA ErrB]
//	Split(ErrA)                    => [ErrA]
func Split(err error) []error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
