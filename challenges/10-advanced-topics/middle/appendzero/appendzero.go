// Package appendzero — Gopher Workplace challenge.
package appendzero

import (
	"errors"
	"reflect"
)

// ErrTarget is returned when slicePtr is not a non-nil pointer to a slice.
var ErrTarget = errors.New("target must be a non-nil pointer to a slice")

// AppendZero appends n zero values to the slice that slicePtr points at.
//
// The element type comes from the slice itself, so one implementation
// serves every slice type.
//
// Examples:
//
//	s := []int{1}; AppendZero(&s, 2) => s is [1 0 0]
func AppendZero(slicePtr any, n int) error {
	panic("not implemented")
}
