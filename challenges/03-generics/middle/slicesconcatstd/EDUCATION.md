# Concat From Stdlib

## Intuition

Writing this by hand means summing lengths and appending; the stdlib version does exactly that, tested, in one call.

## Approach

1. Call `slices.Concat(parts...)`.
2. Replace a nil result with an empty slice.

## Solution

```go
func Join[T any](parts ...[]T) []T {
	out := slices.Concat(parts...)
	if out == nil {
		out = []T{}
	}
	return out
}
```

## Walkthrough

`Join[int]()` concatenates nothing, so the guard turns the nil result into `[]int{}`.

## Pitfalls

- Starting from `parts[0]` and appending, which can mutate a caller's slice.
- Panicking on zero parts.
- Forgetting the `...` when forwarding the variadic argument.
