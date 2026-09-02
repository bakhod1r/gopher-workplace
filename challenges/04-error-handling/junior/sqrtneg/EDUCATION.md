# Square Root Guard

## Intuition

`math.Sqrt` never fails loudly: given a negative input it returns NaN, which then poisons every later computation. An explicit error converts a silent wrong answer into a visible failure.

## Approach

1. Return `0, ErrNegative` when `x < 0`.
2. Otherwise return `math.Sqrt(x), nil`.

## Solution

```go
if x < 0 {
	return 0, ErrNegative
}
return math.Sqrt(x), nil
```

## Walkthrough

For `-1` the guard fires and NaN is never produced. For `0.25` the root is `0.5` with a nil error.

## Pitfalls

- Returning `math.NaN()` alongside the error instead of `0`.
- Testing `x <= 0` — zero has a valid root.
- Checking for NaN after the fact instead of validating the input.
