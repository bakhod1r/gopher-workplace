// Package runall — Gopher Workplace challenge.
package runall

import "errors"

// RunAll runs every step, collecting all failures.
//
// Examples:
//
//	RunAll(func() error { return nil }) => nil
func RunAll(fs ...func() error) error {
	// TODO(candidate): implement this.
	_ = errors.Join
	panic("not implemented")
}
