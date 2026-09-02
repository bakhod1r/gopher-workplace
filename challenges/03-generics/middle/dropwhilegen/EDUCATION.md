# Drop While

## Intuition

Keeping later matches is the point: only the *leading* run is dropped, which is why a filter would give the wrong answer.

## Approach

1. Advance `i` while `pred(s[i])` holds.
2. Copy `s[i:]` into a fresh slice.

## Solution

```go
func DropWhile[T any](s []T, pred func(T) bool) []T {
	i := 0
	for i < len(s) && pred(s[i]) {
		i++
	}
	out := make([]T, 0, len(s)-i)
	out = append(out, s[i:]...)
	return out
}
```

## Walkthrough

`DropWhile([]int{2,4,5,6}, isEven)` stops at index 2 and keeps `[5 6]` — the `6` survives.

## Pitfalls

- Filtering out every match instead of only the prefix.
- Returning `s[i:]`, which aliases the input.
- Running off the end when every element matches.
