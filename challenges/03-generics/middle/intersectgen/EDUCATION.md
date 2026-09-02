# Intersect Slices

## Intuition

Order is taken from `a` deliberately: `Intersect(a, b)` and `Intersect(b, a)` hold the same values but need not list them alike.

## Approach

1. Index `b` into a membership set.
2. Walk `a`, keeping values present in `b` and not yet emitted.

## Solution

```go
func Intersect[T comparable](a, b []T) []T {
	in := make(map[T]bool, len(b))
	for _, v := range b {
		in[v] = true
	}
	seen := make(map[T]bool, len(a))
	out := make([]T, 0, len(a))
	for _, v := range a {
		if !in[v] || seen[v] {
			continue
		}
		seen[v] = true
		out = append(out, v)
	}
	return out
}
```

## Walkthrough

`Intersect([]int{1,2,2,3}, []int{2,3})` skips `1`, emits `2`, skips the duplicate, emits `3`.

## Pitfalls

- Returning duplicates when `a` repeats a value.
- Taking the order from `b`.
- Nested loops instead of a set probe.
