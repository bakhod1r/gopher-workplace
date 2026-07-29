# Pointer vs value receivers

## Intuition

A value receiver copies the struct; mutations don't reach the caller. Any state-changing method must take a pointer receiver.

## Approach

1. A **value** receiver `(c Counter)` mutates a copy, so the caller never sees `c.N` change.
2. Switch to a **pointer** receiver `(c *Counter)` so `c.N++` writes back to the caller's struct.

## Solution

```go
type Counter struct{ N int }

func (c *Counter) Inc() {
	c.N++
}
```

## Walkthrough

With the value receiver, each `Inc` bumps a throwaway copy and `c.N` stays 0. The pointer receiver shares storage, so two calls leave `c.N == 2`.

## Pitfalls

- `(c Counter)` increments a copy; `(c *Counter)` increments the caller's value.
- Keep receiver types consistent across a type's methods.
