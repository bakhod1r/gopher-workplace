// Package positive — Gopher Workplace challenge.
package positive

import "errors"

// ErrNotPositive reports a quantity that is not greater than zero.
var ErrNotPositive = errors.New("value must be positive")

// Positive returns n when it is greater than zero.
//
// Examples:
//
//	Positive(5) => 5, nil
//	Positive(0) => 0, ErrNotPositive
func Positive(n int) (int, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
