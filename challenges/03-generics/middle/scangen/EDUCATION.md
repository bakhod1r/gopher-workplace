# Scan

## Intuition

Emitting the accumulator after applying `f` (rather than before) is what makes the result align element-for-element with the input.

## Approach

1. Start `acc` at `init`.
2. For each element, apply `f` and append the new accumulator.

## Solution

```go
func Scan[T, A any](s []T, init A, f func(A, T) A) []A {
	out := make([]A, 0, len(s))
	acc := init
	for _, v := range s {
		acc = f(acc, v)
		out = append(out, acc)
	}
	return out
}
```

## Walkthrough

`Scan([]int{1,2,3}, 0, add)` records `1`, then `3`, then `6`.

## Pitfalls

- Appending `acc` before applying `f`, which shifts the output by one.
- Including `init` and returning `len(s)+1` elements.
- Returning only the final accumulator.
