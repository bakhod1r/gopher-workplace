// Package int32view — Gopher Workplace challenge.
package int32view

import "unsafe"

// Int32s returns a []int32 view over b's bytes, sharing the storage.
//
// The view is only valid when b's length is a multiple of four and its
// first byte is 4-byte aligned; otherwise the second result is false.
//
// Examples:
//
//	Int32s(eightBytes) => a 2-element view, true
func Int32s(b []byte) ([]int32, bool) {
	panic("not implemented")
}
