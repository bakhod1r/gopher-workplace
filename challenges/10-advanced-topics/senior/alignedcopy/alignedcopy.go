// Package alignedcopy — Gopher Workplace challenge.
package alignedcopy

import "unsafe"

// CopyWords copies as many whole uint32 values as fit from src into dst
// and returns how many were copied.
//
// It reports false when src's length is not a multiple of four or its start
// is not 4-byte aligned; nothing is copied in that case.
//
// Examples:
//
//	CopyWords(make([]uint32, 2), eightBytes) => 2, true
func CopyWords(dst []uint32, src []byte) (int, bool) {
	panic("not implemented")
}
