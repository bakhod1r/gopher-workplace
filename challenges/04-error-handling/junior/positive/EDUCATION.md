# Positive Amount

## Intuition

Domain rules live in comparisons, and the boundary is where they are usually wrong. "Positive" means strictly greater than zero.

## Approach

1. Reject `n <= 0`.
2. Return `n, nil` otherwise.

## Solution

```go
if n <= 0 {
	return 0, ErrNotPositive
}
return n, nil
```

## Walkthrough

`Positive(0)` fails because `0 <= 0` holds; `Positive(1)` passes as the smallest valid input.

## Pitfalls

- Using `n < 0`, which lets zero through.
- Returning `n` alongside the error instead of 0.
- Confusing "non-negative" with "positive".
