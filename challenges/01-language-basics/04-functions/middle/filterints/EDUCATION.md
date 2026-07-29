# Predicate-driven filtering

## Intuition

Accepting a `func(T) bool` generalises selection so callers supply the rule, not a new function per condition.

## Approach

1. Range the input.
2. Keep elements where `keep(v)` is true.

## Solution

```go
func Filter(xs []int, keep func(int) bool) []int {
	var out []int
	for _, v := range xs {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}
```

## Walkthrough

`Filter([1 2 3 4], even)` keeps 2 and 4.

## Pitfalls

- Return empty (not nil-panic) when nothing matches.
- Don't mutate the input slice.
