# Reinterpret Bytes As Wider Values

## Intuition

A slice header is a pointer, a length and a capacity. Reinterpreting means keeping the pointer and changing the element type — so the length has to be recomputed in the new element's units.

## Approach

1. Reject an empty slice and a length that is not a multiple of four.
2. Reject a data pointer whose low two bits are set.
3. Return `unsafe.Slice((*int32)(p), len(b)/4)`.

## Solution

```go
import "unsafe"

// Int32s returns a []int32 view over b's bytes, sharing the storage.
//
// The view is only valid when b's length is a multiple of four and its
// first byte is 4-byte aligned; otherwise the second result is false.
//
// Examples:
//
// 	Int32s(eightBytes) => a 2-element view, true
func Int32s(b []byte) ([]int32, bool) {
	if len(b) == 0 || len(b)%4 != 0 {
		return nil, false
	}
	p := unsafe.Pointer(unsafe.SliceData(b))
	if uintptr(p)&3 != 0 {
		return nil, false
	}
	return unsafe.Slice((*int32)(p), len(b)/4), true
}
```

## Walkthrough

Eight aligned bytes become two int32s sharing the same memory: writing `v[1] = -1` sets bytes 4 through 7 to 0xff.

## Pitfalls

- Passing `len(b)` as the element count, which produces a view four times too long — and reads past the buffer.
- Forgetting that the result's endianness is the machine's, so it is not a wire format.
