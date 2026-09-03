// Package wrapnonnil — Gopher Workplace challenge.
package wrapnonnil

import (
	"errors"
	"fmt"
)

// ErrX is a stand-in failure used by the tests.
var ErrX = errors.New("boom")

// WrapNonNil annotates err with msg, passing nil through untouched.
//
// Examples:
//
//	WrapNonNil("step", nil) => nil
//	WrapNonNil("step", ErrX) => "step: boom"
func WrapNonNil(msg string, err error) error {
	// TODO(candidate): implement this.
	_ = fmt.Errorf
	panic("not implemented")
}
