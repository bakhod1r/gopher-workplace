// Package alignassert — Gopher Workplace challenge.
package alignassert

import "unsafe"

// Counter is a struct whose Value field is updated atomically.
type Counter struct {
	Value int64
	Name  string
}

// Check reports whether Counter's Value field is aligned well enough for
// 64-bit atomic operations.
//
// The requirement is the type's own alignment, which unsafe.Alignof
// reports. Hard-coding a number is how this check passes on the machine it
// was written on and nowhere else.
//
// Examples:
//
//	Check() => true for a correctly laid out Counter
func Check() bool {
	// CHANGE CODE BELOW THIS LINE
	var c Counter
	return uintptr(unsafe.Pointer(&c.Value))%8 == 0 && unsafe.Offsetof(c.Value) == 0
	// CHANGE CODE ABOVE THIS LINE
}
