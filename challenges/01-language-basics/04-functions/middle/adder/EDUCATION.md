# Partial application via closures

## Intuition

Capturing a factory parameter produces specialised functions without extra state types.

## Approach

1. Capture `base` in a closure.
2. Return `func(x int) int { return base + x }`.

## Solution

```go
func Adder(base int) func(int) int {
	return func(x int) int {
		return base + x
	}
}
```

## Walkthrough

`Adder(5)` returns a function that always adds 5, so `add5(3)` is 8.

## Pitfalls

- `base` is captured by reference but never mutated here, so each Adder is stable.
- The returned type must exactly match `func(int) int`.
