# A Struct To Bytes And Back

## Intuition

Reinterpreting is only sound when the bytes really are a `Frame` — the right size, at the right alignment. The dereference then copies the struct out, which is what makes the result independent of the buffer.

## Approach

1. Compare `len(b)` with `unsafe.Sizeof` of a zero Frame.
2. Check the data pointer against `unsafe.Alignof`.
3. Return `*(*Frame)(p), true`.

## Solution

```go
import "unsafe"

// Frame is a fixed-layout record of scalars.
type Frame struct {
	Kind  uint32
	Seq   uint32
	Stamp int64
}

// Encode returns a byte view of f, for the tests to feed back in.
func Encode(f *Frame) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(f)), unsafe.Sizeof(*f))
}

// Decode reinterprets b as a Frame, copying it out so the result does not
// alias b.
//
// The length must be exactly the frame's size and the start must be
// correctly aligned; otherwise the second result is false.
//
// Examples:
//
// 	Decode(encoded) => the frame, true
func Decode(b []byte) (Frame, bool) {
	var zero Frame
	size := unsafe.Sizeof(zero)
	if uintptr(len(b)) != size {
		return zero, false
	}
	p := unsafe.Pointer(unsafe.SliceData(b))
	if uintptr(p)%unsafe.Alignof(zero) != 0 {
		return zero, false
	}
	return *(*Frame)(p), true
}
```

## Walkthrough

A well-formed 16-byte aligned buffer decodes by one struct load. Changing the source afterwards cannot affect the result, because the dereference copied the fields out.

## Pitfalls

- Returning `(*Frame)(p)` — a pointer into the caller's buffer, which is the aliasing the spec forbids.
- Accepting `len(b) >= size`, which silently ignores a framing error.
- Assuming the padding bytes carry information; they do not.
