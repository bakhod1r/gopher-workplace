# The Map That Ships N Zero Values

## Intuition

`make([]U, len(s))` produces a slice that is already `len(s)` long and full of zero values. Appending to it adds the real results *after* those zeros and forces a reallocation on top.

## Approach

1. Allocate with length 0 and capacity `len(s)`.
2. Append one mapped value per input element.
3. Return the slice.

## Solution

```go
func Map[T, U any](s []T, f func(T) U) []U {
	out := make([]U, 0, len(s))
	for _, v := range s {
		out = append(out, f(v))
	}
	return out
}
```

## Walkthrough

`Map([]int{1,2}, double)` returns `[0 0 2 4]` — right values, wrong length, and a needless regrow.

## Pitfalls

- Fixing it by indexing `out[i]` while keeping `append` as well — pick one strategy.
- Dropping the capacity hint entirely, which trades one bug for repeated reallocation.
