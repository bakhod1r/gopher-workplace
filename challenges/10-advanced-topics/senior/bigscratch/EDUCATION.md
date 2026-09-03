# The Scratch Array That Went To The Heap

## Intuition

A stack array costs nothing until something outlives the frame. Returning a slice of it forces the whole array onto the heap — so the "generous" size is not insurance, it is the per-call cost.

## Approach

1. Allocate a right-sized buffer instead: twenty bytes holds any int64 with its sign.
2. Append the digits into it and return the result.

## Solution

```go
import "strconv"

// Format returns v's decimal digits.
//
// The result escapes, so whatever it points into escapes with it — sizing
// the scratch buffer for the worst imaginable case then costs that much
// heap on every call.
//
// Examples:
//
// 	Format(42) => []byte("42")
func Format(v int64) []byte {
	return strconv.AppendInt(make([]byte, 0, 20), v, 10)
}
```

## Walkthrough

Two thousand calls allocate about 8 MB with the 4 KiB array and about 40 KB with a twenty-byte buffer, for identical output.

## Pitfalls

- Trusting `AllocsPerRun` alone; it counts allocations, not bytes.
- Keeping the array and copying out of it, which allocates twice.
