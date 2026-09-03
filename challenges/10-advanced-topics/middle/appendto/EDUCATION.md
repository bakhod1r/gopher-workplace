# Append Into The Caller's Slice

## Intuition

Who allocates is an API decision. Taking a `dst` and returning the extended slice lets a caller in a loop allocate once and reuse, while a caller that does not care can still pass nil.

## Approach

1. Loop `i` from 0 to `n-1`.
2. `dst = append(dst, i*i)`.
3. Return `dst`.

## Solution

```go
// AppendSquares appends the squares 0..n-1 to dst and returns the
// extended slice.
//
// When dst already has the capacity the call must allocate nothing: the
// result is the caller's memory, not the function's.
//
// Examples:
//
// 	AppendSquares(nil, 3) => []int{0, 1, 4}
func AppendSquares(dst []int, n int) []int {
	for i := 0; i < n; i++ {
		dst = append(dst, i*i)
	}
	return dst
}
```

## Walkthrough

With `dst[:0]` over a 64-element array, all 64 appends fit in the existing capacity and the allocator is never called. With `dst == nil`, `append` grows from nothing as usual.

## Pitfalls

- Ignoring `append`'s return value.
- Allocating a local result and copying into `dst`, which defeats the whole point.
