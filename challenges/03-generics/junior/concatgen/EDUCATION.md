# Concat

## Intuition

Starting from a freshly allocated `out` — rather than from `slices[0]` — guarantees no input's backing array is written to.

## Approach

1. Sum the lengths of all inputs.
2. Allocate `out` with that capacity.
3. Append each slice with `...`.

## Solution

```go
func Concat[T any](slices ...[]T) []T {
	n := 0
	for _, s := range slices {
		n += len(s)
	}
	out := make([]T, 0, n)
	for _, s := range slices {
		out = append(out, s...)
	}
	return out
}
```

## Walkthrough

`Concat[int]()` has zero arguments, so `n` stays 0 and the function returns an empty, non-nil `[]int`.

## Pitfalls

- Starting from `slices[0]` and appending onto it, which can mutate the caller's data.
- Panicking on zero arguments by indexing `slices[0]`.
- Forgetting that `Concat()` alone cannot infer `T` — the caller must write `Concat[int]()`.
