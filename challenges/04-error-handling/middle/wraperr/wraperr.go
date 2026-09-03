// Package wraperr — Gopher Workplace challenge.
package wraperr

import (
	"errors"
	"fmt"
)

// ErrDisk is a stand-in low-level failure used by the tests.
var ErrDisk = errors.New("disk offline")

// Wrap annotates err with the operation name, preserving the error chain.
//
// Examples:
//
//	Wrap("read", ErrDisk) => "read: disk offline"
//	Wrap("read", nil)     => nil
func Wrap(op string, err error) error {
	// TODO(candidate): implement this.
	_ = fmt.Errorf
	panic("not implemented")
}
