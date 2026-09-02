# Grow Before Appending

## Intuition

Reserving once turns a sequence of doubling reallocations into one, which is the whole point of the call in a hot path.

## Approach

1. Grow `s` by `len(vs)`.
2. Append the values.
3. Return the result.

## Solution

```go
func Collect[T any](s []T, vs ...T) []T {
	out := slices.Grow(s, len(vs))
	out = append(out, vs...)
	return out
}
```

## Walkthrough

`Collect([]int{1}, 2, 3)` reserves room for two more, then appends without a second allocation.

## Pitfalls

- Assuming `Grow` appends anything — it does not.
- Ignoring `Grow`'s return value and appending to the original.
- Growing by the total length rather than by the number of new elements.
