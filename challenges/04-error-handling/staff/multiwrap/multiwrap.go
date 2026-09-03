// Package multiwrap — Gopher Workplace challenge.
package multiwrap

import (
	"errors"
	"fmt"
)

// Sample failures used by the tests.
var (
	ErrA = errors.New("a")
	ErrB = errors.New("b")
)

// Both combines two failures into one error, keeping both matchable.
//
// Examples:
//
//	Both(ErrA, ErrB) => "a; b"
//	Both(nil, nil)   => nil
func Both(primary, fallback error) error {
	// TODO(candidate): implement this.
	_ = fmt.Errorf
	panic("not implemented")
}
