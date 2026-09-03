// Package safecall — Gopher Workplace challenge.
package safecall

import (
	"errors"
	"fmt"
)

// ErrNilFunc reports a missing callback.
var ErrNilFunc = errors.New("nil function")

// SafeCall runs f, converting a panic into an error.
//
// Examples:
//
//	SafeCall(func() error { return nil }) => nil
//	SafeCall(nil)                         => ErrNilFunc
func SafeCall(f func() error) (err error) {
	// TODO(candidate): implement this.
	_ = fmt.Errorf
	panic("not implemented")
}
