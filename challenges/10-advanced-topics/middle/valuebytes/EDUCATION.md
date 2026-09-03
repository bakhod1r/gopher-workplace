# A Byte View Of One Value

## Intuition

A value's bytes are already in memory in the machine's layout. Viewing them costs a slice header; copying them into a buffer costs an allocation and achieves the same reading.

## Approach

1. Return nil for a nil pointer.
2. `unsafe.Slice((*byte)(unsafe.Pointer(p)), unsafe.Sizeof(*p))`.

## Solution

```go
import "unsafe"

// Bytes returns an 8-byte view of the uint64 p points at, sharing its
// storage.
//
// A nil pointer yields nil. The view is the machine's layout, so it is not
// a portable encoding.
//
// Examples:
//
// 	v := uint64(1); Bytes(&v) => 8 bytes sharing v
func Bytes(p *uint64) []byte {
	if p == nil {
		return nil
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(p)), unsafe.Sizeof(*p))
}
```

## Walkthrough

Setting `v` to all ones makes every byte of the view 0xff, because the view aliases `v` rather than copying it.

## Pitfalls

- Writing 8 instead of `unsafe.Sizeof(*p)`, which is right until the type changes.
- Sending the view over a network — the byte order is the local machine's.
