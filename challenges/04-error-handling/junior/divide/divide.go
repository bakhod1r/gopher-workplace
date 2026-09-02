// Package divide — Gopher Workplace challenge.
package divide

import "errors"

// ErrDivideByZero reports a zero divisor.
var ErrDivideByZero = errors.New("divide by zero")

// Divide returns a/b, or ErrDivideByZero when b is zero.
//
// Examples:
//
//	Divide(10, 2) => 5, nil
//	Divide(1, 0)  => 0, ErrDivideByZero
func Divide(a, b int) (int, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
