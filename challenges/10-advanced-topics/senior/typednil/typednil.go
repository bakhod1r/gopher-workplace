// Package typednil — Gopher Workplace challenge.
package typednil

import "reflect"

// IsNilValue reports whether v is nil or holds a nil pointer, map,
// slice, channel, function or interface.
//
// An interface holding a typed nil pointer is not == nil, which is the trap
// this function exists to close.
//
// Examples:
//
//	var p *int; IsNilValue(p) => true
func IsNilValue(v any) bool {
	panic("not implemented")
}
