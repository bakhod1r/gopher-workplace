# Appending twice to one base

## The idea

If `base` has spare capacity, `append(base, x)` writes into that capacity without
reallocating — and so does a second `append(base, y)`, into the **same** slot.
The two results share memory and stomp each other. Clip first:

```go
base := a[:len(a):len(a)] // cap == len -> every append reallocates
```

## Why it matters

"Fork a slice into two variants" is common (building candidates, backtracking).
The shared-capacity clobber is capacity-dependent, so it passes tests with
zero spare cap and fails in production.

## Watch out

- Only clipping (or copying) guarantees independence.
- `append` reuses spare capacity; that's the whole hazard.
- `slices.Clip` expresses `a[:len(a):len(a)]`.
