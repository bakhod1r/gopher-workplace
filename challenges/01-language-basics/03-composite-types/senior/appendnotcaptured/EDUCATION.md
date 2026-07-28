# append returns a new header

## The idea

A slice header carries length and capacity by value. `append` may write into the
existing backing array, but it always returns an updated header — you must assign
it back:

```go
out = append(out, x*2)
```

## Why it matters

Ignoring `append`'s return leaves the caller's `len` unchanged even if the
element landed in spare capacity — so the value is invisible. It's a subtle,
capacity-dependent bug.

## Watch out

- Always `s = append(s, ...)`.
- `append` may reallocate; the old header is stale afterward.
- The compiler doesn't require using the result (it's just a value), so this
  compiles.
