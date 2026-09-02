// Package modulo — Gopher Workplace challenge.
package modulo

import "errors"

// ErrZeroModulus reports a zero modulus.
var ErrZeroModulus = errors.New("zero modulus")

// Mod returns a%b, or ErrZeroModulus when b is zero.
//
// Examples:
//
//	Mod(10, 3) => 1, nil
//	Mod(5, 0)  => 0, ErrZeroModulus
func Mod(a, b int) (int, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
