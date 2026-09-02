// Package orzero — Gopher Workplace challenge.
package orzero

import "errors"

// ErrHost is a stand-in failure used by the tests.
var ErrHost = errors.New("host unreachable")

// OrZero returns v when err is nil, and 0 otherwise.
//
// Examples:
//
//	OrZero(42, nil)     => 42
//	OrZero(42, ErrHost) => 0
func OrZero(v int, err error) int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
