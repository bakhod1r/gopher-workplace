# Slice bounds and clamping

## Intuition

`xs[:n]` requires `n <= len(xs)`; otherwise it panics at runtime. A "take up to n"
operation clamps the end:

```go
if n > len(xs) { n = len(xs) }
return xs[:n]
```

## Approach

1. Bug: return xs[:n] panics when n > len(xs) (slice bound out of range). 2. Fix: clamp n to len(xs) before slicing: if n > len(xs) { n = len(xs) }. 3. Then xs[:n] is always in range.

## Solution

```go
func Take(xs []int, n int) []int {
	if n <= 0 {
		return xs[:0]
	}
	if n > len(xs) {
		n = len(xs)
	}
	return xs[:n]
}
```

## Walkthrough

xs=[1,2,3], n=10: xs[:10] panics. After clamping n=3, xs[:3]=[1,2,3].

## Pitfalls

- Slicing bounds are checked at runtime; out-of-range panics.
- `min(n, len(xs))` (Go 1.21+) expresses the clamp.
- Capacity allows `xs[:n]` up to cap for the two-index form — but that exposes
  stale data; prefer length clamping.
