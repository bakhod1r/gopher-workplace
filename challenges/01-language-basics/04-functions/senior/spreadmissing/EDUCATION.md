# Spreading a slice into a variadic call

## Intuition

`f(xs...)` forwards each element of `xs`; without the `...` you pass a single value (here the length), not the contents.

## Approach

1. To pass a slice to a variadic function, spread it with `...`.
2. The bug passes `len(xs)` (a single int); use `sum(xs...)`.

## Solution

```go
func sum(nums ...int) int {
	t := 0
	for _, n := range nums {
		t += n
	}
	return t
}

func Total(xs []int) int {
	return sum(xs...)
}
```

## Walkthrough

`sum(len(xs))` sums the length, not the elements. `sum(xs...)` unpacks the slice into the variadic parameter.

## Pitfalls

- `sum(xs...)` spreads; `sum(xs)` won't compile for `...int`, and `sum(len(xs))` compiles but is wrong.
- A nil slice spreads to zero arguments.
