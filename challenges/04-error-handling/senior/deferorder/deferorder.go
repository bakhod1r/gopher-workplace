// Package deferorder — Gopher Workplace challenge.
package deferorder

import "errors"

// CloseAll runs every closer in reverse order, collecting failures.
//
// Examples:
//
//	CloseAll() => nil
func CloseAll(closers ...func() error) error {
	// TODO(candidate): implement this.
	_ = errors.Join
	panic("not implemented")
}
