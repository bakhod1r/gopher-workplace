# Equal By Predicate

## Intuition

Because equality is supplied rather than assumed, this works for float tolerances, case-insensitive strings, and cross-type comparisons alike.

## Approach

1. Return `false` on differing lengths.
2. Return `false` at the first failing pair.
3. Return `true` after the loop.

## Solution

```go
func EqualBy[T, U any](a []T, b []U, eq func(T, U) bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !eq(a[i], b[i]) {
			return false
		}
	}
	return true
}
```

## Walkthrough

`EqualBy([]int{1}, []string{"1"}, matches)` compares one pair, which the predicate accepts.

## Pitfalls

- Requiring `comparable` and losing cross-type use.
- Skipping the length check and indexing out of range.
- Returning `true` when only a prefix matches.
