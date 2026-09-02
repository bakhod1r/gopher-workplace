// Package validage — Gopher Workplace challenge.
package validage

import "errors"

// Validation failures.
var (
	ErrTooYoung = errors.New("age below zero")
	ErrTooOld   = errors.New("age above 130")
)

// ValidAge reports whether age is plausible.
//
// Examples:
//
//	ValidAge(30)  => nil
//	ValidAge(-1)  => ErrTooYoung
//	ValidAge(200) => ErrTooOld
func ValidAge(age int) error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
