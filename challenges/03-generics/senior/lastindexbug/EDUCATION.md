# Last Index That Finds The First

## Intuition

A forward scan returning on the first match implements `Index`, not `LastIndex`.

## Approach

1. Walk from the last index down to zero.
2. Return the first equal element found.
3. Return -1 if the loop completes.

## Solution

```go
func LastIndex[T comparable](s []T, v T) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == v {
			return i
		}
	}
	return -1
}
```

## Walkthrough

`LastIndex([1,2,1], 1)` returns 0 instead of 2.

## Pitfalls

- Tracking the best index forward without returning early — correct but scans the whole slice.
- Starting at `len(s)` and indexing out of range.
