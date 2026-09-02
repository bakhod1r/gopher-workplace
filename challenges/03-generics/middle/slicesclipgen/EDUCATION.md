# Clip Spare Capacity

## Intuition

Returning a sub-slice with spare capacity hands the caller a write pointer into your buffer; clipping converts that hazard into a harmless allocation.

## Approach

1. Return `slices.Clip(s)`, normalising nil to an empty slice.

## Solution

```go
func Freeze[T any](s []T) []T {
	out := slices.Clip(s)
	if out == nil {
		out = []T{}
	}
	return out
}
```

## Walkthrough

A slice of length 2 and capacity 8 becomes length 2 and capacity 2, so appending copies rather than overwriting.

## Pitfalls

- Returning the raw sub-slice and letting callers write into your buffer.
- Confusing `Clip` with `Compact` — the former changes capacity, the latter removes duplicates.
- Assuming `Clip` copies: it does not, it only narrows the capacity.
