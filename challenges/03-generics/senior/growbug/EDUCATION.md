# Reservation Thrown Away

## Intuition

Slices are values: `Grow` and `append` both return new headers, and the caller only sees whichever header you return. Appending to one variable and returning another discards the work.

## Approach

1. Assign `slices.Grow`'s result.
2. Append the values to that same slice.
3. Return the slice you appended to.

## Solution

```go
func Collect[T any](s []T, vs ...T) []T {
	out := slices.Grow(s, len(vs))
	out = append(out, vs...)
	return out
}
```

## Walkthrough

`append` writes into the grown array and returns a header of length 3; returning `s` hands back the original header of length 1, so the new values are invisible.

## Pitfalls

- Ignoring the return value of `Grow`, `append`, `Compact`, or `Delete`.
- Assuming a helper mutates the slice header through the argument.
- Measuring only correctness and missing the lost optimisation.
