# The comma-ok return pattern

## Intuition

A trailing boolean result reports whether the other values are meaningful, mirroring map lookups and type assertions.

## Approach

1. Guard the empty slice, returning `ok == false`.
2. Otherwise return `xs[0]`, `xs[len-1]`, true.

## Solution

```go
func FirstLast(xs []int) (first, last int, ok bool) {
	if len(xs) == 0 {
		return 0, 0, false
	}
	return xs[0], xs[len(xs)-1], true
}
```

## Walkthrough

`[1 2 3]` returns first 1, last 3, ok true; an empty slice returns zeros and false.

## Pitfalls

- Reading `xs[0]` before the length guard panics on an empty slice.
- Return the zero values (`0, 0`) with `false`, not garbage.
