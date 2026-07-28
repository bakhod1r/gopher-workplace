# Inserting into a slice

## The idea

Insertion stitches three pieces: the head, the new value, and the tail:

```go
out := append([]int{}, xs[:i]...)
out = append(out, v)
out = append(out, xs[i:]...)
```

## Why it matters

Ordered inserts (sorted sets, priority lists) need this. The subtle trap is doing
`append(xs[:i], v)` in place: it can overwrite `xs[i]` before you copy the tail,
because they share the backing array.

## Watch out

- Building into a fresh slice avoids the aliasing clobber.
- `slices.Insert` (Go 1.21+) handles this correctly and efficiently.
- Clamp `i`; `xs[:i]` panics for `i > len`.
