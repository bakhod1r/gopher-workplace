// Package fastdispatch — Gopher Workplace challenge.
package fastdispatch

import (
	"reflect"
	"strconv"
)

// Render appends v's text form to dst.
//
// The common types are handled by a type switch, which costs nothing;
// everything else falls back to reflection. The fast path must not
// allocate.
//
// Examples:
//
//	Render(nil, 42) => []byte("42")
func Render(dst []byte, v any) []byte {
	panic("not implemented")
}
