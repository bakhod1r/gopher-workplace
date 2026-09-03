// Package bytesequal — Gopher Workplace challenge.
package bytesequal

import "unsafe"

// EqualString reports whether b's bytes are exactly s.
//
// Neither side may be converted: a conversion in either direction copies
// the payload just to throw the copy away.
//
// Examples:
//
//	EqualString([]byte("hi"), "hi") => true
func EqualString(b []byte, s string) bool {
	panic("not implemented")
}
