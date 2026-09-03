// Package uintptrhold — Gopher Workplace challenge.
package uintptrhold

import "unsafe"

// Pair is two 64-bit words.
type Pair struct {
	A int64
	B int64
}

// SecondWord returns the B field of the pair p points at, reached through
// the field's offset.
//
// Address arithmetic must stay in unsafe.Pointer. A uintptr is a plain
// number: nothing keeps the object alive and nothing updates it.
//
// Examples:
//
//	SecondWord(&Pair{A: 1, B: 2}) => 2
func SecondWord(p *Pair) int64 {
	// CHANGE CODE BELOW THIS LINE
	addr := uintptr(unsafe.Pointer(p))
	addr += unsafe.Offsetof(p.B)
	return *(*int64)(unsafe.Pointer(addr))
	// CHANGE CODE ABOVE THIS LINE
}
