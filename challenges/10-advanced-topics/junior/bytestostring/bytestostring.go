// Package bytestostring — Gopher Workplace challenge.
package bytestostring

import "unsafe"

// Str returns a string that shares b's bytes instead of copying them.
//
// The result is only valid while b is not written to again: a string is
// supposed to be immutable, and this one is not.
//
// Examples:
//
//	Str([]byte("hi")) => "hi"
func Str(b []byte) string {
	panic("not implemented")
}
