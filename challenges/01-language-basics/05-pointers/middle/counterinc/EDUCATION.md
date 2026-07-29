# Pointer receivers mutate

## Intuition

A pointer-receiver method operates on the caller's value; a value-receiver method gets a copy and can't persist mutations.

## Approach

1. The receiver is `*Counter`, so mutations persist.
2. `c.N++` increments the caller's field.

## Solution

```go
type Counter struct{ N int }

func (c *Counter) Inc() {
	c.N++
}
```

## Walkthrough

`c.Inc()` on `&Counter{}` sets `N` to `1`; a pointer receiver is required or the increment would be lost on a copy.

## Pitfalls

- `(c Counter)` (value) would increment a copy.
- Go auto-takes the address for `c.Inc()` when `c` is addressable.
