# A Byte View Of A Struct, Only When It Is Safe

## Intuition

Reinterpreting a struct as bytes is sound only when the bytes mean the same thing to whoever reads them. A pointer's bytes do not: they name an address in this process, and copying them out both leaks a layout detail and hides a reference from the collector.

## Approach

1. Reject a nil pointer.
2. Reject a type that `hasPointers` reports on.
3. Return `unsafe.Slice((*byte)(unsafe.Pointer(p)), unsafe.Sizeof(*p))`.

## Solution

```go
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
// 	Bytes(&Frame{}) => a view of unsafe.Sizeof(Frame{}) bytes, true
func Bytes(p *Frame) ([]byte, bool) {
	if p == nil {
		return nil, false
	}
	t := reflect.TypeOf(*p)
	if hasPointers(t) {
		return nil, false
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(p)), unsafe.Sizeof(*p)), true
}
```

## Walkthrough

`Frame` is two uint32s and an int64 — sixteen bytes, no pointers. The view aliases the struct, so writing `f.Kind` changes the first four bytes of the slice.

## Pitfalls

- Sending the view over the wire as a portable format; the layout, padding and endianness are all the local machine's.
- Using `len(b)` from a hard-coded number instead of `Sizeof`, which breaks the moment a field is added.
- Assuming the padding bytes are zero — they are whatever was there.
