# Greatest Common Divisor

## Intuition

The constraint choice is doing real work: `%` simply does not exist for `~float64`, so a numeric set including floats would not compile.

## Approach

1. Take the magnitude of both arguments.
2. Loop until `b` is zero, swapping and taking the remainder.
3. Return `a`.

## Solution

```go
func GCD[T Integer](a, b T) T {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```

## Walkthrough

`GCD(12, 18)` runs `(18, 12)`, `(12, 6)`, `(6, 0)` and returns `6`.

## Pitfalls

- Including float types in the constraint, which breaks `%`.
- Returning a negative result for negative inputs.
- Looping while `a != 0`, which returns the wrong operand.
