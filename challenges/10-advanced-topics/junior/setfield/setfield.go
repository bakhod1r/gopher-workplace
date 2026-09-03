// Package setfield — Gopher Workplace challenge.
package setfield

import (
	"errors"
	"reflect"
)

// ErrNotSettable is returned when the target cannot be written.
var ErrNotSettable = errors.New("target is not a settable int field")

// SetInt sets the named int field of the struct ptr points at.
//
// Reflection can only write through a pointer: a value passed by interface
// is a copy, and the reflect package refuses to modify it.
//
// Examples:
//
//	SetInt(&counters{}, "Hits", 3) => nil, Hits is 3
func SetInt(ptr any, field string, n int) error {
	panic("not implemented")
}
