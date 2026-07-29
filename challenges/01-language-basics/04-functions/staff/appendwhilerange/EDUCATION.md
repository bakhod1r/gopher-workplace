# Bounding loops that grow their own target

## Intuition

A loop appending to `out` must not use `len(out)` as its bound; iterate the fixed source length instead to avoid skips or runaway growth.

## Approach

1. Looping over `len(out)` while appending to `out` never terminates (it grows).
2. Bound the loop by the fixed input `len(xs)`.

## Solution

```go
func DupAll(xs []int) []int {
	out := make([]int, 0, len(xs)*2)
	for i := 0; i < len(xs); i++ {
		out = append(out, xs[i], xs[i]*2)
	}
	return out
}
```

## Walkthrough

Ranging over the growing `out` keeps finding new elements forever. Iterating `len(xs)` visits each original element once.

## Pitfalls

- Bound on the stable input (`len(xs)`), never the growing output.
- `for i := range xs` snapshots the count safely.
