// Package kindof — Gopher Workplace challenge.
package kindof

import "reflect"

// KindName returns the name of v's underlying kind: "int", "slice",
// "struct" and so on.
//
// A nil interface has no type at all, so it reports "invalid".
//
// Examples:
//
//	KindName(3) => "int"
func KindName(v any) string {
	panic("not implemented")
}
