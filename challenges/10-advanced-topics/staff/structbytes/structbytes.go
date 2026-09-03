// Package structbytes — Gopher Workplace challenge.
package structbytes

import (
	"reflect"
	"unsafe"
)

// Frame is a fixed-layout wire frame of scalars only.
type Frame struct {
	Kind  uint32
	Seq   uint32
	Stamp int64
}

// hasPointers reports whether t contains any pointer-shaped field.
func hasPointers(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Pointer, reflect.UnsafePointer, reflect.Slice, reflect.Map,
		reflect.Chan, reflect.Func, reflect.Interface, reflect.String:
		return true
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			if hasPointers(t.Field(i).Type) {
				return true
			}
		}
		return false
	case reflect.Array:
		return hasPointers(t.Elem())
	default:
		return false
	}
}

// Bytes returns a byte view of the frame p points at, for writing to a
// socket without an intermediate copy.
//
// This is only defined when the struct contains no pointers: a byte view of
// a pointer field would let the bytes outlive what they point at, and would
// hand the peer an address. Report false rather than producing one.
//
// Examples:
//
//	Bytes(&Frame{}) => a view of unsafe.Sizeof(Frame{}) bytes, true
func Bytes(p *Frame) ([]byte, bool) {
	panic("not implemented")
}
