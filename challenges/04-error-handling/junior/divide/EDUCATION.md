# Safe Division

## Intuition

Go has no exceptions. A function that can fail returns an extra `error` value, and the caller must look at it. Failure is part of the signature.

## Approach

1. Guard the failing condition first: `if b == 0`.
2. Return the zero value of the result plus the sentinel error.
3. Otherwise return the quotient and `nil`.

## Solution

```go
if b == 0 {
	return 0, ErrDivideByZero
}
return a / b, nil
```

## Walkthrough

For `Divide(1, 0)`: `b == 0` is true, so the function returns `0, ErrDivideByZero` before the division runs.

## Pitfalls

- Dividing first and checking after — `1 / 0` panics.
- Returning a non-zero result alongside an error.
- Returning a fresh error instead of the sentinel, so `errors.Is` fails.
