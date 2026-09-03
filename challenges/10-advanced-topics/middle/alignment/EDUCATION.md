# Is This Address Aligned

## Intuition

Alignment is a property of the address, so you have to look at the address as a number. The mask trick works because a power-of-two boundary means the low bits of the address are all zero.

## Approach

1. Reject an empty slice, `n == 0`, and non-powers of two.
2. Convert the data pointer to `uintptr` and test `addr & (n-1) == 0`.

## Solution

```go
import "unsafe"

// Aligned reports whether b's first byte sits at an address that is a
// multiple of n.
//
// n must be a power of two; anything else, or an empty slice, reports
// false.
//
// Examples:
//
// 	Aligned(buf, 8) => true when buf starts on an 8-byte boundary
func Aligned(b []byte, n uintptr) bool {
	if len(b) == 0 || n == 0 || n&(n-1) != 0 {
		return false
	}
	return uintptr(unsafe.Pointer(unsafe.SliceData(b)))&(n-1) == 0
}
```

## Walkthrough

A `[]uint64`'s storage is 8-byte aligned, so its address ends in three zero bits and `addr & 7` is 0. Taking `b[1:]` adds 1, so the test fails; `b[8:]` adds 8 and it passes again.

## Pitfalls

- Storing the `uintptr` in a variable and using it later — the collector may move the object, and the number goes stale.
- Using `%` with a non-power-of-two `n`, which the guard exists to reject.
