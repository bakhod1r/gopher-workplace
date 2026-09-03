// Package panicerror — Gopher Workplace challenge.
package panicerror

import (
	"errors"
	"fmt"
)

// ErrStop is a stand-in control-flow error used by the tests.
var ErrStop = errors.New("stop")

// Capture runs f and returns whatever it panicked with, as an error.
//
// Examples:
//
//	Capture(func() { panic(ErrStop) }) => ErrStop
//	Capture(func() {})                 => nil
func Capture(f func()) (err error) {
	// TODO(candidate): implement this.
	_ = fmt.Errorf
	panic("not implemented")
}
