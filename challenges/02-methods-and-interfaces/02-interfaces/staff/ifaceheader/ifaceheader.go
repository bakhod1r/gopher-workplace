// Package ifaceheader — Gopher Workplace challenge.
package ifaceheader

import "unsafe"

// Words returns the two machine words of an interface value:
// the type pointer and the data pointer.
//
// Examples:
//
//	Words(nil)          => 0, 0
//	Words((*int)(nil))  => non-zero, 0
func Words(v any) (typ uintptr, data uintptr) {
	// TODO(candidate): read the two-word header.
	panic("not implemented")
}

// IsTypedNil reports whether v holds a type but a nil value.
func IsTypedNil(v any) bool {
	// TODO(candidate): non-zero type word, zero data word.
	panic("not implemented")
}

// Classify returns "nil", "typed-nil", or "value".
func Classify(v any) string {
	// TODO(candidate): classify by the header words.
	panic("not implemented")
}

var _ = unsafe.Pointer(nil)
