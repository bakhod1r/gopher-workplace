# Is Sorted

## Intuition

Checking `s[i] < s[i-1]` (rather than `s[i] <= s[i-1]`) is what makes duplicates acceptable — the difference between sorted and strictly increasing.

## Approach

1. Loop from index 1 to the end.
2. Return `false` when an element is smaller than its predecessor.
3. Return `true` after the loop.

## Solution

```go
func IsSorted[T cmp.Ordered](s []T) bool {
	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			return false
		}
	}
	return true
}
```

## Walkthrough

`IsSorted([]int{1, 2, 2})` compares `2 < 1` (no) and `2 < 2` (no), so it returns `true`.

## Pitfalls

- Starting at index 0 and reading `s[-1]`.
- Using `<=`, which rejects legitimate duplicates.
- Returning `false` for an empty slice.
