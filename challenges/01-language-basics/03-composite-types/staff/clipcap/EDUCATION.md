# Clipping capacity

## The idea

The three-index slice expression `xs[low:high:max]` sets capacity to `max-low`.
`xs[:len(xs):len(xs)]` yields cap == len, so any later `append` must allocate a
new array instead of writing into shared spare capacity:

```go
return xs[:len(xs):len(xs)]
```

## Why it matters

When you return or store a sub-slice of a larger buffer, leftover capacity is a
footgun: the caller's `append` silently corrupts the rest of your buffer. Clipping
is the standard defense (`slices.Clip`).

## Watch out

- Capacity, not length, governs append reuse.
- Clipping doesn't copy — it still shares the (now cap-limited) array.
- Pair with a copy when you also need to release the backing array.
