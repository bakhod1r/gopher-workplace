# The three-index slice expression

## The idea

`xs[:2]` has length 2 but inherits `xs`'s **capacity**, so `append` writes into
the shared tail (`xs[2]`). The three-index form `xs[low:high:max]` caps capacity,
forcing `append` to allocate a fresh array:

```go
head := xs[:2:2] // len 2, cap 2 -> append reallocates
return append(head, extra)
```

## Why it matters

"Append to a sub-slice" silently mutating the parent is a notorious Go aliasing
bug (it depends on capacity, so it's intermittent). `xs[:n:n]` is the standard
guard when handing out sub-slices that will be appended to.

## Watch out

- Capacity, not length, decides whether append reuses memory.
- The bug only manifests when spare capacity exists — easy to miss in tests.
- `slices.Clip` sets cap==len for the same effect.
