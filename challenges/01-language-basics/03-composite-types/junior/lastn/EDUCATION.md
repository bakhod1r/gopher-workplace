# Slice expressions

## The idea

`xs[a:b]` creates a slice sharing the same backing array, covering indices
`[a, b)`. A tail of `n` is `xs[len(xs)-n:]` — but only after clamping `n`:

```go
if n > len(xs) { n = len(xs) }
if n < 0 { n = 0 }
return xs[len(xs)-n:]
```

## Why it matters

Slicing is O(1) and allocation-free, but an out-of-range index panics. Clamping
turns "give me the last N" into a safe operation for any N.

## Watch out

- `xs[len(xs)-n:]` panics if `n > len(xs)` — clamp first.
- The result **shares** memory with `xs`; appending to it may overwrite.
- `xs[i:]` is shorthand for `xs[i:len(xs)]`.
