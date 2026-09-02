# Index Of Max

## Intuition

Storing the index rather than the value means there is never a second variable to keep in sync, and the strict `>` is what makes the first maximum win.

## Approach

1. Return `-1` for an empty slice.
2. Start `best` at 0.
3. Update `best` whenever `s[i]` is strictly greater.

## Solution

```go
func IndexOfMax[T cmp.Ordered](s []T) int {
	if len(s) == 0 {
		return -1
	}
	best := 0
	for i := 1; i < len(s); i++ {
		if s[i] > s[best] {
			best = i
		}
	}
	return best
}
```

## Walkthrough

`IndexOfMax([]int{1, 9, 9})` moves `best` to 1 and leaves it there, since `s[2] > s[1]` is false.

## Pitfalls

- Using `>=`, which returns the last maximum.
- Returning the value instead of the index.
- Returning `0` for an empty slice.
