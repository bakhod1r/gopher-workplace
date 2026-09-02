# Lexicographic Comparison

## Intuition

This is dictionary order generalised: compare positions until they differ, then fall back to length — the same rule strings follow.

## Approach

1. Return `slices.CompareFunc(a, b, ...)` with `cmp.Compare` on the name.

## Solution

```go
func CompareNames(a, b []Item) int {
	return slices.CompareFunc(a, b, func(x, y Item) int {
		return cmp.Compare(x.Name, y.Name)
	})
}
```

## Walkthrough

`CompareNames([a], [a,b])` finds every shared position equal, so the shorter slice sorts first.

## Pitfalls

- Comparing lengths first, which is not lexicographic.
- Returning a `bool`.
- Assuming equal-length slices always compare as 0.
