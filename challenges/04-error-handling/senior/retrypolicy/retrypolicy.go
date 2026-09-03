// Package retrypolicy — Gopher Workplace challenge.
package retrypolicy

import "errors"

// Failure classes used by the tests.
var (
	ErrTransient = errors.New("transient")
	ErrInvalid   = errors.New("invalid")
)

// Retry calls f while it fails transiently, up to attempts times.
//
// Examples:
//
//	Retry(3, func() error { return nil }) => nil
func Retry(attempts int, f func() error) error {
	// TODO(candidate): implement this.
	_ = errors.Is
	panic("not implemented")
}
