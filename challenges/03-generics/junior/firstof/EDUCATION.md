# First Element

## Intuition

For an unknown type there is no literal you can write for "nothing". `var zero T` gives you exactly the zero value of whatever `T` turned out to be: `0`, `""`, `false`, or `nil`.

## Approach

1. Return early with `var zero T, false` when `len(s) == 0`.
2. Otherwise return `s[0], true`.

## Solution

```go
func First[T any](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	return s[0], true
}
```

## Walkthrough

`First([]int{})` takes the guard branch: `zero` is `int(0)`, and the `false` tells the caller that `0` is not real data.

## Pitfalls

- Indexing `s[0]` before the length check — that panics on an empty slice.
- Trying `return nil, false`: `nil` is not assignable to an arbitrary `T`.
- Returning only the element and using a sentinel value like `-1` to mean empty.
