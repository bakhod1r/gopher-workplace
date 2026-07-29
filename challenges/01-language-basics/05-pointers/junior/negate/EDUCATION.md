# Unary mutation through a pointer

## Intuition

`*p = -*p` loads, negates, and stores back at the pointee.

## Approach

1. Read the pointee `*p`.
2. Store its negation: `*p = -*p`.

## Solution

```go
func Negate(p *int) {
	*p = -*p
}
```

## Walkthrough

`Negate(&x)` with `x = 5`: `-*p` is `-5`, written back to `x`.

## Pitfalls

- Negating twice is a no-op (`-(-x)==x`).
- `-*p` reads first, then assign.
