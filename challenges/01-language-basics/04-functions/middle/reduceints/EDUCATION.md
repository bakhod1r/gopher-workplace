# Folding with an accumulator

## Intuition

Reduce expresses sum, product, max, and concatenation as one shape parameterised by the combining function and seed.

## Approach

1. Start `acc := init`.
2. Fold each element with `acc = f(acc, v)`.

## Solution

```go
func Reduce(xs []int, init int, f func(acc, x int) int) int {
	acc := init
	for _, v := range xs {
		acc = f(acc, v)
	}
	return acc
}
```

## Walkthrough

`Reduce([1 2 3 4], 0, add)` accumulates 0+1+2+3+4 = 10.

## Pitfalls

- With an empty slice, the result is `init`.
- Argument order in `f(acc, x)` must match the signature.
