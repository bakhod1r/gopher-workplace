# Repeat Slice

## Intuition

`append` with `...` works the same for a type parameter as for a concrete type: the compiler knows the element size once `T` is instantiated.

## Approach

1. Allocate `out` with capacity `len(s)*max(n, 0)`.
2. Append all of `s` on each of the `n` iterations.
3. Return `out`.

## Solution

```go
func Repeat[T any](s []T, n int) []T {
	out := make([]T, 0, len(s)*max(n, 0))
	for i := 0; i < n; i++ {
		out = append(out, s...)
	}
	return out
}
```

## Walkthrough

`Repeat([]int{1, 2}, 2)` runs twice, each time splicing `[1 2]` onto `out`, producing `[1 2 1 2]`.

## Pitfalls

- Appending `s` itself (`append(out, s)`) — that does not compile for `[]T`.
- Returning the input slice when `n == 1` instead of a fresh copy.
- Forgetting the `n <= 0` case and passing a negative capacity to `make`.
