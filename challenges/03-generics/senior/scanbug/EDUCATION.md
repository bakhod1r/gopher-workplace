# Running Total Off By One Step

## Intuition

Appending before folding records the state *before* each element, so the series is shifted: the opening value appears and the final one is lost.

## Approach

1. Start `acc` at `init`.
2. Apply `f` first, then append the new accumulator.

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

`Scan([]int{1,2,3}, 0, add)` must end at `6`; the buggy version ends at `3` and starts at `0`.

## Pitfalls

- Including `init` in the output.
- Appending before applying the fold.
- Testing only the length of the result.
