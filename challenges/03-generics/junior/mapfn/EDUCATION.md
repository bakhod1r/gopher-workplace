# Map Over a Slice

## Intuition

One type parameter would force the result to have the same element type as the input. `U` lets `Map` change the type, which is how `[]int` becomes `[]string`.

## Approach

1. Allocate `out` as `[]U` with capacity `len(s)`.
2. Append `f(e)` for each element.
3. Return `out`.

## Solution

```go
func Map[T, U any](s []T, f func(T) U) []U {
	out := make([]U, 0, len(s))
	for _, e := range s {
		out = append(out, f(e))
	}
	return out
}
```

## Walkthrough

`Map([]int{1, 2}, itoa)` infers `T = int` from the slice and `U = string` from the function's return type, so `out` is `[]string`.

## Pitfalls

- Declaring `out` as `[]T` — the results may be a different type.
- Calling `f` on the index instead of the element.
- Skipping elements when `f` returns a zero value — every input must produce one output.
