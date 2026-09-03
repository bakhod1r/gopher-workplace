// Package fieldsview — Gopher Workplace challenge.
package fieldsview

import "strings"

// Fields splits s on sep and returns the pieces.
//
// Substrings of a string share the original bytes, so the pieces cost
// nothing but their headers. Only the header slice may be allocated.
//
// Examples:
//
//	Fields("a,b,c", ',') => []string{"a", "b", "c"}
func Fields(s string, sep byte) []string {
	panic("not implemented")
}
