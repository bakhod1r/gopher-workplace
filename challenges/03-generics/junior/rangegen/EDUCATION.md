# Range

## Intuition

Making the counter a `T` keeps `i < n` and `append(out, i)` type-correct without a single conversion in the loop body.

## Approach

1. Allocate an empty `out`.
2. Count `i` from `T(0)` while `i < n`.
3. Append each `i`.

## Solution

```go
func Range[T Integer](n T) []T {
	out := make([]T, 0)
	for i := T(0); i < n; i++ {
		out = append(out, i)
	}
	return out
}
```

## Walkthrough

`Range(int64(2))` instantiates `T = int64`, so the counter is an `int64` and the result is `[]int64{0, 1}`.

## Pitfalls

- Using an `int` counter and comparing it to `n`, which does not compile for `int64`.
- Returning `nil` for `n <= 0` when an empty slice is expected.
- Producing `1..n` instead of `0..n-1`.
