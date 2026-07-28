# Slice bounds and clamping

## The idea

`xs[:n]` requires `n <= len(xs)`; otherwise it panics at runtime. A "take up to n"
operation clamps the end:

```go
if n > len(xs) { n = len(xs) }
return xs[:n]
```

## Why it matters

Pagination, previews, and "first N" helpers routinely get an N larger than the
data. Clamping turns a crash into graceful truncation — defensive slicing at
trust boundaries.

## Watch out

- Slicing bounds are checked at runtime; out-of-range panics.
- `min(n, len(xs))` (Go 1.21+) expresses the clamp.
- Capacity allows `xs[:n]` up to cap for the two-index form — but that exposes
  stale data; prefer length clamping.
