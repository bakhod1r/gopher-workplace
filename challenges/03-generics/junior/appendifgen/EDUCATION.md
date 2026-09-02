# Append If

## Intuition

`append` may return a slice with a different backing array, so its result must be returned rather than discarded. That is why this helper returns `[]T` instead of mutating in place.

## Approach

1. Return `s` when `cond` is false.
2. Otherwise return `append(s, v)`.

## Solution

```go
func AppendIf[T any](s []T, v T, cond bool) []T {
	if !cond {
		return s
	}
	return append(s, v)
}
```

## Walkthrough

`AppendIf([]int{1}, 2, false)` returns the same slice header it received, so the caller's `s = AppendIf(...)` assignment is a no-op.

## Pitfalls

- Calling `append(s, v)` and ignoring the result.
- Appending first and truncating afterwards.
- Taking a `*[]T` to mutate in place, which the signature does not ask for.
