# Partial application via closures

## The idea

Capturing a factory parameter produces specialised functions without extra state types.

## Why it matters

It underlies configurable callbacks, middleware, and dependency injection in Go.

## Watch out

- `base` is captured by reference but never mutated here, so each Adder is stable.
- The returned type must exactly match `func(int) int`.
