# Clone

## Intuition

Cloning is shallow: for `[]int` that is total independence, for `[]*Job` the two slices still share the jobs. Knowing which you have is the whole skill.

## Approach

1. Clone the input.
2. Turn a nil clone into an empty slice.
3. Return it.

## Solution

```go
func Detach[T any](s []T) []T {
	out := slices.Clone(s)
	if out == nil {
		out = []T{}
	}
	return out
}
```

## Walkthrough

Writing `out[0] = 99` after `Detach` leaves the caller's `s[0]` alone, because the clone owns a different backing array.

## Pitfalls

- Returning `s[:]`, which shares the backing array entirely.
- Assuming a clone of `[]*T` deep-copies the pointed-to values.
- Returning nil for a nil input when the contract says non-nil.
