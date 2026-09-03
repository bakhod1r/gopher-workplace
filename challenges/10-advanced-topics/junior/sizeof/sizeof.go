// Package sizeof — Gopher Workplace challenge.
package sizeof

import "unsafe"

// Header is the fixed part of a record.
type Header struct {
	Id   int64
	Name string
	Tags []string
}

// Sizes returns the size in bytes of the Header type and of its Id and
// Name fields.
//
// unsafe.Sizeof is a compile-time constant: it measures the type, not the
// data a pointer or slice header refers to.
//
// Examples:
//
//	Sizes() => 40, 8, 16 on a 64-bit build
func Sizes() (header, id, name uintptr) {
	panic("not implemented")
}
