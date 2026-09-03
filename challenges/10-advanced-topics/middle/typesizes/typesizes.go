// Package typesizes — Gopher Workplace challenge.
package typesizes

import "reflect"

// Sizes returns the size and alignment of v's dynamic type.
//
// A nil interface has no type, so it reports false.
//
// Examples:
//
//	Sizes(int64(0)) => 8, 8, true
func Sizes(v any) (size, align uintptr, ok bool) {
	panic("not implemented")
}
