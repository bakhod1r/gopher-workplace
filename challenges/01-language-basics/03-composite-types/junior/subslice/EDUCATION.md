# Slice expressions

## Intuition

`xs[low:high]` creates a slice over the same backing array covering `[low, high)`.
It is O(1) and shares memory with the original:

```go
mid := xs[1:4] // elements 1,2,3
```

## Approach

1. Clamp n down to len(s) if it is larger.
2. If n <= 0, return an empty slice.
3. make a fresh len-n slice and copy s[:n] into it — no shared backing array.
4. Return the copy.

## Solution

```go
func Head(s []int, n int) []int {
	if n > len(s) {
		n = len(s)
	}
	if n <= 0 {
		return []int{}
	}
	result := make([]int, n)
	copy(result, s[:n])
	return result
}
```

## Walkthrough

Head([1,2,3,4],2): n=2, make len-2 slice, copy [1,2]. Writing result[0]=99 leaves s[0]=1 because backing arrays differ.

## Pitfalls

- High index must be ≤ `len` (or `cap` for the two-index form) or it panics.
- Sub-slices share the backing array; append may or may not alias depending on
  capacity.
- The three-index form `xs[a:b:c]` caps capacity.
