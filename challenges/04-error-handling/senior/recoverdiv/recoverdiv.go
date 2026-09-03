// Package recoverdiv — Gopher Workplace challenge.
package recoverdiv

import "errors"

// ErrPanic reports a recovered runtime panic.
var ErrPanic = errors.New("recovered panic")

// SafeDivide divides a by b, converting any panic into an error.
//
// Examples:
//
//	SafeDivide(10, 2) => 5, nil
//	SafeDivide(1, 0)  => 0, ErrPanic
func SafeDivide(a, b int) (result int, err error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
