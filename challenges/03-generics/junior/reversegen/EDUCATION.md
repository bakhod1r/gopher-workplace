# Reverse

## Intuition

Reordering in place would mutate the caller's backing array. Allocating a fresh slice keeps the function free of side effects, and the capacity hint avoids regrowth.

## Approach

1. Allocate `out` with length 0 and capacity `len(s)`.
2. Walk `s` from the last index down to 0, appending each element.
3. Return `out`.

## Solution

```go
func Reverse[T any](s []T) []T {
	out := make([]T, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		out = append(out, s[i])
	}
	return out
}
```

## Walkthrough

`Reverse([]int{1, 2, 3})` appends `s[2]`, `s[1]`, `s[0]` in turn, producing `[3 2 1]` while `s` still reads `[1 2 3]`.

## Pitfalls

- Swapping in place — the caller's slice changes too.
- Returning `nil` for an empty input when the tests expect a non-nil empty slice.
- Starting the loop at `len(s)` instead of `len(s)-1`.
