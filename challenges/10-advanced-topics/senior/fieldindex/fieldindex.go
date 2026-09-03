// Package fieldindex — Gopher Workplace challenge.
package fieldindex

import (
	"errors"
	"reflect"
)

// ErrShape is returned when rows is not a slice of structs with an int
// field of the given name.
var ErrShape = errors.New("rows must be a slice of structs with that int field")

// SumColumn totals the named int field across a slice of structs.
//
// Resolving the field by name is a string search through the struct's field
// table; doing it per row makes the cost O(rows x fields).
//
// Examples:
//
//	SumColumn([]rec{{N: 1}, {N: 2}}, "N") => 3, nil
func SumColumn(rows any, field string) (int64, error) {
	panic("not implemented")
}
