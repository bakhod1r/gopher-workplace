// Package fieldnames — Gopher Workplace challenge.
package fieldnames

import "reflect"

// FieldNames returns the names of v's exported fields, in declaration
// order.
//
// A non-struct, or a nil interface, yields nil.
//
// Examples:
//
//	FieldNames(struct{ A, b int }{}) => []string{"A"}
func FieldNames(v any) []string {
	panic("not implemented")
}
