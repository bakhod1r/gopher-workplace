# Split At

## Intuition

Returning sub-slices would be cheaper but leaves both halves sharing one array, so appending to the left half can overwrite the right half's first element.

## Approach

1. Clamp `i` to `[0, len(s)]`.
2. Copy `s[:i]` and `s[i:]` into fresh slices.

## Solution

```go
func SplitAt[T any](s []T, i int) ([]T, []T) {
	if i < 0 {
		i = 0
	}
	if i > len(s) {
		i = len(s)
	}
	left := make([]T, i)
	copy(left, s[:i])
	right := make([]T, len(s)-i)
	copy(right, s[i:])
	return left, right
}
```

## Walkthrough

`SplitAt([]int{1,2}, 9)` clamps `i` to 2, producing the whole slice and an empty second half.

## Pitfalls

- Returning `s[:i], s[i:]`, which aliases the input and each other.
- Panicking on an out-of-range index.
- Returning nil halves when one side is empty.
