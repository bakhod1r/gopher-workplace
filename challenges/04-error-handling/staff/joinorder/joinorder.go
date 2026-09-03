// Package joinorder — Gopher Workplace challenge.
package joinorder

import (
	"errors"
	"sort"
)

// Sample failures used by the tests.
var (
	ErrA = errors.New("a")
	ErrB = errors.New("b")
	ErrC = errors.New("c")
)

// Sorted joins the non-nil errors ordered by message.
//
// Examples:
//
//	Sorted() => nil
func Sorted(errs ...error) error {
	// TODO(candidate): implement this.
	_, _ = sort.Slice, errors.Join
	panic("not implemented")
}
