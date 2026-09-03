// Package twophase — Gopher Workplace challenge.
package twophase

import "errors"

// Do applies a change, confirms it, and rolls back on a failed confirm.
//
// Examples:
//
//	Do(ok, ok, rb) => nil
func Do(apply, confirm, rollback func() error) error {
	// TODO(candidate): implement this.
	_ = errors.Join
	panic("not implemented")
}
