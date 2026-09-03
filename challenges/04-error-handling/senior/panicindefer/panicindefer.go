// Package panicindefer — Gopher Workplace challenge.
package panicindefer

import (
	"errors"
	"fmt"
)

// ErrPanic reports a recovered panic.
var ErrPanic = errors.New("recovered panic")

// Guard runs work then cleanup, converting panics into errors.
//
// Examples:
//
//	Guard(func() error { return nil }, func() { }) => nil
func Guard(work func() error, cleanup func()) (err error) {
	// TODO(candidate): implement this.
	_ = fmt.Errorf
	panic("not implemented")
}
