// Package dedupejoin — Gopher Workplace challenge.
package dedupejoin

import "errors"

// Sample failures used by the tests.
var (
	ErrA = errors.New("a")
	ErrB = errors.New("b")
)

// Distinct joins errs, keeping one error per distinct message.
//
// Examples:
//
//	Distinct(nil) => nil
func Distinct(errs ...error) error {
	// TODO(candidate): implement this.
	_ = errors.Join
	panic("not implemented")
}
