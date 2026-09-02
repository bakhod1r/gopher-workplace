# Counts That Never Grow

## Intuition

A missing key reads as `0`, so `m[v]++` needs no initialisation. Assigning `1` instead resets the tally on every hit.

## Approach

1. Make the map.
2. Increment the entry for each element.
3. Return the map.

## Solution

```go
func Count[T comparable](s []T) map[T]int {
	m := make(map[T]int, len(s))
	for _, v := range s {
		m[v]++
	}
	return m
}
```

## Walkthrough

For `[a a b]` the second `a` overwrites the count with `1` rather than raising it to `2`.

## Pitfalls

- Guarding the increment with a presence check — unnecessary for `int` values.
- Assuming the zero value is always a safe default; for pointers or slices it is `nil`.
