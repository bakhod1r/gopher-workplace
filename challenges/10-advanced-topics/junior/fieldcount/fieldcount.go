// Package fieldcount — Gopher Workplace challenge.
package fieldcount

import "reflect"

// FieldCount returns how many fields v's struct type has in total, and
// how many of them are exported.
//
// A non-struct, or a nil interface, reports 0, 0.
//
// Examples:
//
//	FieldCount(rec{}) => 3, 2
func FieldCount(v any) (total, exported int) {
	panic("not implemented")
}
