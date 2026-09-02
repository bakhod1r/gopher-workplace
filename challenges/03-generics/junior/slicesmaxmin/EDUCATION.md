# Max And Min From Stdlib

## Intuition

Delegating is only safe once you know the callee's preconditions. Here the wrapper exists precisely to add the missing empty-slice contract.

## Approach

1. Return `zero, false` when the slice is empty.
2. Otherwise delegate to `slices.Max` / `slices.Min` and report `true`.

## Solution

```go
func Peak[T cmp.Ordered](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	return slices.Max(s), true
}

func Floor[T cmp.Ordered](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	return slices.Min(s), true
}
```

## Walkthrough

`Peak([]int{})` never reaches `slices.Max`, so the panic the stdlib would raise cannot happen.

## Pitfalls

- Calling `slices.Max(s)` unguarded and panicking on empty input.
- Sorting the slice to find the extremes, which mutates the caller's data.
- Assuming `slices.Max` returns a second `ok` value — it does not.
