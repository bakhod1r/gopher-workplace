# Prefix Skip Turned Into A Filter

## Intuition

A filter has no notion of position, so it removes every matching element. The specified behaviour stops caring about the predicate after the first failure.

## Approach

1. Advance an index while the predicate holds.
2. Copy everything from that index onward into a fresh slice.

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

For `[2 4 5 6]` the split lands at index 2, so `5` and `6` are both kept even though `6` matches.

## Pitfalls

- Reimplementing `Filter` and calling it `DropWhile`.
- Returning `s[i:]`, which aliases the caller's array.
- Running off the end when every element matches.
