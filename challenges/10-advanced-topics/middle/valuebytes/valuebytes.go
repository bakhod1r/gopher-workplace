// Package valuebytes — Gopher Workplace challenge.
package valuebytes

import "unsafe"

// Bytes returns an 8-byte view of the uint64 p points at, sharing its
// storage.
//
// A nil pointer yields nil. The view is the machine's layout, so it is not
// a portable encoding.
//
// Examples:
//
//	v := uint64(1); Bytes(&v) => 8 bytes sharing v
func Bytes(p *uint64) []byte {
	panic("not implemented")
}
