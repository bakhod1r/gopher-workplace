# Slice index bounds

## Intuition

Valid indices run `0 .. len(xs)-1`; seeding a pointer at `len(xs)` and dereferencing it panics immediately.

## Approach

1. The right index starts at `len(xs)-1`, the last valid element.
2. The bug seeds `j := len(xs)`, indexing out of range.

## Solution

```go
func Reverse(xs []int) {
	i, j := 0, len(xs)-1
	for i < j {
		xs[i], xs[j] = xs[j], xs[i]
		i++
		j--
	}
}
```

## Walkthrough

`xs[len(xs)]` is out of bounds on the first swap. Seeding `j` at `len(xs)-1` swaps the true ends inward.

## Pitfalls

- The highest valid index is `len(xs)-1`, not `len(xs)`.
- `xs[len(xs)]` always panics.
