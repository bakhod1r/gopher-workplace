# Keep The Capacity, Drop The Contents

## Intuition

Allocating is the expensive part, not zeroing. `buf[:0]` throws away the *view* while keeping the array, so the next round of appends writes into memory you have already paid for.

## Approach

1. Return `buf[:0]`.

## Solution

```go
// Reset returns buf emptied for reuse, keeping the capacity it already
// has so the next writer does not have to allocate again.
//
// Examples:
//
// 	Reset(make([]byte, 8, 64)) => length 0, capacity 64
func Reset(buf []byte) []byte {
	return buf[:0]
}
```

## Walkthrough

`make([]byte, 8, 64)` owns a 64-byte array. `buf[:0]` still points at it with cap 64, so the next 64 bytes of appends need no allocation.

## Pitfalls

- `return nil` or `return []byte{}` — correct length, capacity thrown away.
- Reusing a buffer whose old contents someone still holds a view of.
