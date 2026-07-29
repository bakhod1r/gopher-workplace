# Read-modify-write through a pointer

## Intuition

`*p += d` loads, adds, and stores at the pointee in one expression.

## Approach

1. `p` aliases the caller's int.
2. `*p += delta` reads, adds, and stores back at the same address.

## Solution

```go
func Add(p *int, delta int) {
	*p += delta
}
```

## Walkthrough

`x := 10`, `Add(&x, 5)`: `*p += 5` computes `10 + 5` and writes `15` back to `x`.

## Pitfalls

- `*p += delta` mutates the caller's variable.
- Works for negative deltas too.
