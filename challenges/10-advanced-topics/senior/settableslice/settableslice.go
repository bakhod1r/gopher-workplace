// Package settableslice — Gopher Workplace challenge.
package settableslice

import (
	"errors"
	"reflect"
)

// ErrShape is returned when slice is not a slice of ints.
var ErrShape = errors.New("argument must be a slice of ints")

// Double multiplies every element of the int slice in place.
//
// reflect.ValueOf gives a copy of the interface's contents, but a slice's
// elements live in the shared backing array — which is exactly why the
// elements are settable and the slice header is not.
//
// Examples:
//
//	s := []int{1, 2}; Double(s) => s is [2 4]
func Double(slice any) error {
	// CHANGE CODE BELOW THIS LINE
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice || rv.Type().Elem().Kind() != reflect.Int {
		return ErrShape
	}
	for i := 0; i < rv.Len(); i++ {
		e := rv.Index(i)
		v := e.Int()
		e = reflect.ValueOf(int(v * 2))
		_ = e
	}
	return nil
	// CHANGE CODE ABOVE THIS LINE
}
