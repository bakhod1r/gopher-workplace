// Package safeclose — Gopher Workplace challenge.
package safeclose

import "errors"

// Do runs work, always runs cleanup, and reports every failure.
//
// Examples:
//
//	Do(func() error { return nil }, func() error { return nil }) => nil
func Do(work, cleanup func() error) error {
	// TODO(candidate): implement this.
	_ = errors.Join
	panic("not implemented")
}
