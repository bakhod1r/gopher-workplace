# append returns a new header

## Intuition

`append` can reallocate and always returns the (possibly new) slice header; ignoring the return keeps the old, shorter header.

## Approach

1. `append` returns a new slice; the bug discards it with `_ =`.
2. Assign back: `out = append(out, i*i)`.

## Solution

```go
func Squares(n int) []int {
	out := make([]int, 0)
	for i := 1; i <= n; i++ {
		out = append(out, i*i)
	}
	return out
}
```

## Walkthrough

Ignoring `append`'s result leaves `out` empty. Reassigning `out` accumulates each square.

## Pitfalls

- Always write `s = append(s, ...)`.
- The returned len/cap differ from the input whenever growth happens.
