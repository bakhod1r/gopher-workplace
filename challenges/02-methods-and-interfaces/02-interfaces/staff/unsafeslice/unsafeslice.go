// Package unsafeslice — Gopher Workplace challenge.
package unsafeslice

import "unsafe"

// Converter turns bytes into a string.
type Converter interface {
	Convert(b []byte) string
}

// BytesToString reinterprets b as a string without copying.
//
// The result aliases b: mutating b afterwards changes the string. Only use
// this when b is never written again.
//
// Examples:
//
//	BytesToString([]byte("abc")) => "abc", no allocation
func BytesToString(b []byte) string {
	// TODO(candidate): unsafe.String, with the empty case handled.
	panic("not implemented")
}

// StringToBytes reinterprets s as a byte slice without copying.
//
// The result must never be written to: string data is read-only.
func StringToBytes(s string) []byte {
	// TODO(candidate): unsafe.Slice, with the empty case handled.
	panic("not implemented")
}

// SafeString copies b into a new string.
func SafeString(b []byte) string {
	// TODO(candidate): the copying conversion.
	panic("not implemented")
}

var _ = unsafe.Pointer(nil)
