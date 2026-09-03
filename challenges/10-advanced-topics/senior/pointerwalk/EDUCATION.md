# Walk An Array By Pointer

## Intuition

Pointer arithmetic in Go counts bytes, not elements. Every step has to be scaled by the element size, which is exactly what indexing a slice does for you when you have one.

## Approach

1. Guard nil and non-positive `n`.
2. For each `i`, offset `p` by `i * unsafe.Sizeof(*p)` and read through the typed pointer.
3. Accumulate into an int64.

## Solution

```go
import "unsafe"

// SumInt32 totals n consecutive int32 values starting at p.
//
// This is the shape a C API hands you: a pointer and a count, with no
// slice. n <= 0 or a nil pointer totals 0.
//
// Examples:
//
// 	SumInt32(&a[0], 3) => a[0] + a[1] + a[2]
func SumInt32(p *int32, n int) int64 {
	if p == nil || n <= 0 {
		return 0
	}
	var total int64
	for i := 0; i < n; i++ {
		q := (*int32)(unsafe.Add(unsafe.Pointer(p), uintptr(i)*unsafe.Sizeof(*p)))
		total += int64(*q)
	}
	return total
}
```

## Walkthrough

For four int32s, the offsets are 0, 4, 8 and 12 bytes. Adding `i` instead of `i*4` would read overlapping, misaligned values.

## Pitfalls

- `unsafe.Add(p, i)` — a one-byte step through a four-byte array.
- `unsafe.Slice(p, n)` is the idiomatic answer in real code; here the loop is the exercise.
