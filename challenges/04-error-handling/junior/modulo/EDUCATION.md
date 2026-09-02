# Safe Remainder

## Intuition

`%` is division's sibling and shares its failure mode: a zero right-hand operand panics at runtime. The guard is identical.

## Approach

1. Return the error when `b == 0`.
2. Return `a % b, nil` otherwise.

## Solution

```go
if b == 0 {
	return 0, ErrZeroModulus
}
return a % b, nil
```

## Walkthrough

`Mod(-7, 3)` is `-1` in Go — the result takes the sign of the dividend, which is why the test pins that case.

## Pitfalls

- Assuming `%` always returns a non-negative result.
- Computing before the guard.
- Normalising the sign, which changes the documented behaviour.
