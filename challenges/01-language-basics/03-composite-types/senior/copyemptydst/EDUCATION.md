# copy respects length, not capacity

## Intuition

`copy(dst, src)` copies `min(len(dst), len(src))` elements. A slice made with
`make([]int, 0, n)` has **length 0**, so `copy` writes nothing — capacity is
irrelevant:

```go
dst := make([]int, len(xs)) // length = len(xs)
copy(dst, xs)
```

## Approach

1. Bug: make([]int, 0, len(xs)) gives a destination of length 0; copy copies min(len(dst),len(src)) = 0 elements. 2. copy is bounded by the shorter length, not capacity. 3. Fix: make([]int, len(xs)) so dst has full length and copy fills it.

## Solution

```go
func Clone(xs []int) []int {
	dst := make([]int, len(xs))
	copy(dst, xs)
	return dst
}
```

## Walkthrough

dst len 0, copy(dst,xs) returns 0, dst stays empty. With len(xs) length, copy transfers all elements -> independent clone.

## Pitfalls

- To pre-allocate for `append`, use length 0 + capacity; to `copy`, use full
  length.
- `copy` returns the number copied — useful to assert.
- `slices.Clone` avoids the footgun entirely.
