// Package fieldkinds — Gopher Workplace challenge.
package fieldkinds

import "reflect"

// FieldKinds returns "Name:kind" for each exported field of v, in
// declaration order.
//
// A non-struct, or a nil interface, yields nil.
//
// Examples:
//
//	FieldKinds(row{}) => []string{"ID:int", "Name:string"}
func FieldKinds(v any) []string {
	panic("not implemented")
}
