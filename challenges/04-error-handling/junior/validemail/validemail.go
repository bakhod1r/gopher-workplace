// Package validemail — Gopher Workplace challenge.
package validemail

import (
	"errors"
	"strings"
)

// Validation failures.
var (
	ErrNoAt      = errors.New("missing @")
	ErrEmptyPart = errors.New("empty local or domain part")
)

// ValidEmail reports whether s has the shape local@domain.
//
// Examples:
//
//	ValidEmail("a@b.com") => nil
//	ValidEmail("abc")     => ErrNoAt
func ValidEmail(s string) error {
	// TODO(candidate): implement this.
	_ = strings.Cut
	panic("not implemented")
}
