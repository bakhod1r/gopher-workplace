# Filter That Forgets To Shrink

## Intuition

The compaction is correct but the original full-length slice is returned, so the tail still shows the pre-filter elements.

## Approach

1. Walk the slice with a write cursor.
2. Copy each kept element to the cursor.
3. Return `s[:n]`.

## Solution

```go
func FilterInPlace[T any](s []T, pred func(T) bool) []T {
	n := 0
	for _, v := range s {
		if pred(v) {
			s[n] = v
			n++
		}
	}
	return s[:n]
}
```

## Walkthrough

`FilterInPlace([1,2,3,4], even)` writes `[2 4 3 4]` and returns all four elements.

## Pitfalls

- Forgetting that the caller's original slice is now scrambled — that is the price of in-place.
- Clearing the tail when `T` holds pointers is worth doing to release references.
