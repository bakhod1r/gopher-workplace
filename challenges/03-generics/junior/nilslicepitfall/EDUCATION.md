# Nil Versus Empty

## Intuition

Both forms append identically, so the difference only shows when nothing is appended — which is exactly the case that reaches JSON as `null`.

## Approach

1. `Collect`: allocate with `make`, append the accepted elements, return.
2. `IsNil`: return `s == nil`.

## Solution

```go
func Collect[T any](s []T, keep func(T) bool) []T {
	out := make([]T, 0, len(s))
	for _, v := range s {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}

func IsNil[T any](s []T) bool {
	return s == nil
}
```

## Walkthrough

`Collect([]int{1}, none)` appends nothing, but the result is still the allocated empty slice, so `IsNil` reports `false`.

## Pitfalls

- Declaring `var out []T`, which returns nil when nothing matches.
- Testing `len(s) == 0` and concluding the slice is nil.
- Returning `nil` early as a shortcut for the no-match case.
