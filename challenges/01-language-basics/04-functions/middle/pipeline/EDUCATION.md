# Building pipelines from function slices

## Intuition

Threading a value through a slice of functions gives a data-flow pipeline with the order fixed by slice order.

## Approach

1. Fold the value through each function left to right.

## Solution

```go
func Pipe(x int, fns ...func(int) int) int {
	for _, f := range fns {
		x = f(x)
	}
	return x
}
```

## Walkthrough

`Pipe(3, inc, double)`: inc(3)=4, double(4)=8.

## Pitfalls

- Empty `fns` must return x unchanged (identity).
- Order is fns[0] first — opposite of mathematical Compose.
