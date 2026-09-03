// Package reflectlen — Gopher Workplace challenge.
package reflectlen

import "reflect"

// Length returns the length of v when it has one — a string, slice,
// array, map or channel — and reports false otherwise.
//
// Examples:
//
//	Length([]int{1, 2}) => 2, true
func Length(v any) (int, bool) {
	panic("not implemented")
}
