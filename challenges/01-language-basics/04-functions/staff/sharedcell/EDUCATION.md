# Shared vs per-closure captured cells

## Intuition

Closures capture variables by identity; a variable declared outside the loop is one cell shared by all, while one declared inside gives each closure its own.

## Approach

1. A single `c` declared outside the loop is shared by all closures.
2. Declare `c := 0` **inside** the loop so each closure captures its own cell.

## Solution

```go
func Counters(n int) []func() int {
	out := make([]func() int, 0, n)
	for i := 0; i < n; i++ {
		c := 0
		out = append(out, func() int { c++; return c })
	}
	return out
}
```

## Walkthrough

With the shared `c`, incrementing one counter moves them all. A per-iteration `c` gives each closure independent state.

## Pitfalls

- Hoisting `c` out of the loop makes every closure share it.
- Declare per-iteration state inside the loop body.
