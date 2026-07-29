# Clipping capacity

## Intuition

The three-index slice expression `xs[low:high:max]` sets capacity to `max-low`.
`xs[:len(xs):len(xs)]` yields cap == len, so any later `append` must allocate a
new array instead of writing into shared spare capacity:

```go
return xs[:len(xs):len(xs)]
```

## Approach

1. Bug: `return xs` leaves cap unchanged (10), so a later append can mutate shared memory.
2. Fix: `return xs[:len(xs):len(xs)]` sets cap == len via the three-index slice.

## Solution

```go
func Clip(xs []int) []int {
	return xs[:len(xs):len(xs)]
}
```

## Walkthrough

xs: len 3 cap 10. xs[:3:3] yields the same elements but cap 3. Now append would reallocate rather than reuse shared backing.

## Pitfalls

- Capacity, not length, governs append reuse.
- Clipping doesn't copy — it still shares the (now cap-limited) array.
- Pair with a copy when you also need to release the backing array.
