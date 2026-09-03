// Package setallints — Gopher Workplace challenge.
package setallints

import (
	"errors"
	"reflect"
)

// ErrTarget is returned when ptr is not a non-nil pointer to a struct.
var ErrTarget = errors.New("target must be a non-nil pointer to a struct")

// SetAllInts sets every settable int field of the struct ptr points at to
// v, and reports how many fields it wrote.
//
// Unexported fields and fields of other kinds are skipped.
//
// Examples:
//
//	SetAllInts(&rec{}, 7) => 2, nil for a struct with two int fields
func SetAllInts(ptr any, v int) (int, error) {
	panic("not implemented")
}
