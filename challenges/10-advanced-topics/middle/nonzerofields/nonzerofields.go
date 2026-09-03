// Package nonzerofields — Gopher Workplace challenge.
package nonzerofields

import "reflect"

// NonZero returns the names of v's exported fields that hold something
// other than their zero value, in declaration order.
//
// A non-struct, or a nil interface, yields nil.
//
// Examples:
//
//	NonZero(patch{Name: "x"}) => []string{"Name"}
func NonZero(v any) []string {
	panic("not implemented")
}
