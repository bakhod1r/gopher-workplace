// Package mapfield — Gopher Workplace challenge.
package mapfield

import (
	"errors"
	"reflect"
)

// ErrTarget is returned when ptr has no settable Tags map field.
var ErrTarget = errors.New("target must be a pointer to a struct with a settable Tags map[string]string")

// PutTag sets ptr's Tags map entry, creating the map when the field is
// nil.
//
// Writing to a nil map panics, and reflection will not create one for you.
//
// Examples:
//
//	PutTag(&doc{}, "a", "1") => nil, doc.Tags["a"] == "1"
func PutTag(ptr any, key, val string) error {
	panic("not implemented")
}
