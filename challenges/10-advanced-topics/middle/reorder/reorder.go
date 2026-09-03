// Package reorder — Gopher Workplace challenge.
package reorder

import "unsafe"

// Entry is one cache record. Reorder its fields to remove the padding.
//
// CHANGE STRUCT BELOW THIS LINE
type Entry struct {
	Flag byte
	Ref  int64
	Kind byte
	Seq  int32
}

// CHANGE STRUCT ABOVE THIS LINE

// Size returns the size of the Entry type.
//
// Entry's fields are declared in an order that forces the compiler to
// insert padding between them. Reordering them from widest to narrowest
// removes it without changing what the struct holds.
//
// Examples:
//
//	Size() => 16 once the fields are ordered well
func Size() uintptr {
	// CHANGE CODE BELOW THIS LINE
	return unsafe.Sizeof(Entry{})
	// CHANGE CODE ABOVE THIS LINE
}
