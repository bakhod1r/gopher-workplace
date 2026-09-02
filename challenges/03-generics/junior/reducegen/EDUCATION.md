# Reduce

## Intuition

The accumulator's type is separate from the element type, which is what makes `Reduce` general enough to express sums, joins, and counts with the same code.

## Approach

1. Copy `init` into `acc`.
2. Reassign `acc = f(acc, e)` for each element.
3. Return `acc`.

## Solution

```go
func Reduce[T, A any](s []T, init A, f func(A, T) A) A {
	acc := init
	for _, e := range s {
		acc = f(acc, e)
	}
	return acc
}
```

## Walkthrough

`Reduce([]int{1, 2, 3}, 0, add)` computes `add(0,1)=1`, `add(1,2)=3`, `add(3,3)=6`.

## Pitfalls

- Calling `f(e, acc)` — argument order is fixed by the signature.
- Starting from the zero value of `A` instead of `init`.
- Folding right-to-left, which changes the result for operations like subtraction or concatenation.
