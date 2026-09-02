// Package validator — Gopher Workplace challenge.
package validator

import "errors"

// Rule failures.
var (
	ErrEmpty   = errors.New("empty")
	ErrTooLong = errors.New("too long")
)

// Validator checks one rule.
type Validator interface {
	Validate(s string) error
}

// NotEmpty rejects the empty string.
type NotEmpty struct{}

// Validate rejects "".
func (n NotEmpty) Validate(s string) error {
	// TODO(candidate): ErrEmpty when s is "".
	panic("not implemented")
}

// MaxLen rejects strings longer than N bytes.
type MaxLen struct {
	N int
}

// Validate rejects over-long input.
//
// Examples:
//
//	MaxLen{N: 2}.Validate("ab")  => nil
//	MaxLen{N: 2}.Validate("abc") => ErrTooLong
func (m MaxLen) Validate(s string) error {
	// TODO(candidate): ErrTooLong when len(s) > N.
	panic("not implemented")
}

// ValidateAll runs every rule and returns the first failure.
func ValidateAll(vs []Validator, s string) error {
	// TODO(candidate): stop at the first error.
	panic("not implemented")
}
