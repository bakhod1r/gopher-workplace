// Package derefsafe — Gopher Workplace challenge.
package derefsafe

// Value returns what p points at, or 0 when p is nil.
//
// A nil pointer is a valid value; dereferencing one is not.
//
// Examples:
//
//	Value(nil) => 0
func Value(p *int) int {
	panic("not implemented")
}
