# Copy Words Only When The Buffer Allows

## Intuition

Reinterpretation and copying compose: the view gives you typed elements without moving anything, and `copy` then moves exactly as many as both sides can hold.

## Approach

1. Reject an empty or non-multiple-of-four `src`, and a misaligned start.
2. Build `unsafe.Slice((*uint32)(p), len(src)/4)`.
3. Return `copy(dst, view), true`.

## Solution

```go
import "unsafe"

// CopyWords copies as many whole uint32 values as fit from src into dst
// and returns how many were copied.
//
// It reports false when src's length is not a multiple of four or its start
// is not 4-byte aligned; nothing is copied in that case.
//
// Examples:
//
// 	CopyWords(make([]uint32, 2), eightBytes) => 2, true
func CopyWords(dst []uint32, src []byte) (int, bool) {
	if len(src) == 0 || len(src)%4 != 0 {
		return 0, false
	}
	p := unsafe.Pointer(unsafe.SliceData(src))
	if uintptr(p)&3 != 0 {
		return 0, false
	}
	view := unsafe.Slice((*uint32)(p), len(src)/4)
	return copy(dst, view), true
}
```

## Walkthrough

Twelve source bytes become a three-element view; a two-element `dst` receives two words and the count is 2.

## Pitfalls

- Passing `len(src)` as the element count, which builds a view four times too long.
- Returning the view itself, which would alias the caller's bytes.
