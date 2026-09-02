# Is Zero

## Intuition

There is no literal that means "the zero of an unknown type", so you name one with `var zero T`. Comparing needs `comparable`, not `any`.

## Approach

1. Declare `var zero T`.
2. Return `v == zero`.

## Solution

```go
func IsZero[T comparable](v T) bool {
	var zero T
	return v == zero
}
```

## Walkthrough

`IsZero("")` instantiates `T = string`, so `zero` is `""` and the comparison is `"" == ""`, giving `true`.

## Pitfalls

- Using `[T any]`, which cannot use `==`.
- Comparing against `0` — that only compiles for numeric types.
- Reaching for `reflect` when the constraint already gives you `==`.
