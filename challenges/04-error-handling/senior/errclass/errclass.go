// Package errclass — Gopher Workplace challenge.
package errclass

import "errors"

// Failure classes used by the tests.
var (
	ErrTimeout = errors.New("timeout")
	ErrDenied  = errors.New("denied")
)

// Classify counts failures by class name.
//
// Examples:
//
//	Classify(nil) => {}
func Classify(errs []error) map[string]int {
	// TODO(candidate): implement this.
	_ = errors.Is
	panic("not implemented")
}
