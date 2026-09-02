# Pointers To Elements

## Intuition

This puzzle is the classic loop-variable capture bug in pointer form. Modern Go fixes the scoping, but code that hoists the variable out of the loop still produces N identical pointers.

## Approach

1. Allocate `[]*T` with capacity `len(s)`.
2. For each element, take the address of a per-iteration copy.
3. Return the result.

## Solution

```go
func PtrsTo[T any](s []T) []*T {
	out := make([]*T, 0, len(s))
	for _, v := range s {
		v := v
		out = append(out, &v)
	}
	return out
}
```

## Walkthrough

`PtrsTo([]int{1, 2})` yields two pointers to two separate variables holding `1` and `2`.

## Pitfalls

- Declaring `var v T` before the loop and appending `&v` every iteration — all pointers alias one variable.
- Appending `&s[i]`, which aliases the caller's backing array.
- Returning nil for an empty input.
