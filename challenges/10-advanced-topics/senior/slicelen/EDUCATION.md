# The Length Argument Is In Elements

## Intuition

`unsafe.Slice(p, n)` builds a slice of `n` elements of `*p`'s type. Handing it a byte count for a four-byte element type asks for four times the buffer — and nothing in the runtime will object.

## Approach

1. Keep the validation as it is.
2. Pass `len(b)/4` as the element count.

## Solution

```go
import "unsafe"

// Words returns a []uint32 view over b's bytes.
//
// unsafe.Slice takes a count of elements, not of bytes: passing the byte
// length produces a view four times too long, running off the end of the
// buffer.
//
// Examples:
//
// 	Words(eightBytes) => a 2-element view, true
func Words(b []byte) ([]uint32, bool) {
	if len(b) == 0 || len(b)%4 != 0 {
		return nil, false
	}
	p := unsafe.Pointer(unsafe.SliceData(b))
	if uintptr(p)&3 != 0 {
		return nil, false
	}
	return unsafe.Slice((*uint32)(p), len(b)/4), true
}
```

## Walkthrough

For an 8-byte buffer the correct view is 2 elements. The buggy version asks for 8, so `v[2]` through `v[7]` read 24 bytes that belong to something else.

## Pitfalls

- Testing only that the values look right — the first two do.
- Fixing it with `[:len(b)/4]` after the fact, which narrows the view but was already constructed out of bounds.
