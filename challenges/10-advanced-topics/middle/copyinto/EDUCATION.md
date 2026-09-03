# Copy Stops At The Shorter Side

## Intuition

`copy` is the one builtin that refuses to write past either side. That makes "fill what fits" a single call, with the count falling out for free.

## Approach

1. Return `copy(dst, src)`.

## Solution

```go
// CopyInto copies as many elements as fit from src into dst and returns
// how many were copied.
//
// Neither slice is resized: the copy is bounded by the shorter of the two.
//
// Examples:
//
// 	CopyInto(make([]int, 2), []int{1, 2, 3}) => 2
func CopyInto(dst, src []int) int {
	return copy(dst, src)
}
```

## Walkthrough

A destination of length 2 and a source of length 3 copies 2. A destination of length 0 and capacity 8 copies nothing, because capacity is not writable through `copy`.

## Pitfalls

- Looping with `len(src)` as the bound, which overruns a shorter `dst`.
- Expecting a `dst` with spare capacity to receive more than its length.
