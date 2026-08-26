# Method Values with Pointer Receivers

## Intuition

When you take a method value from a pointer receiver (`c.Inc` where `c` is
`*Counter`), the returned function captures the **pointer**. Every call through
that function mutates the same underlying Counter.

This is exactly how closures work — but instead of closing over a variable, you
close over a receiver.

## Approach

1. Create `c := &Counter{}`.
2. Return `c.Inc`.

## Solution

```go
func NewCounter() func() int {
	c := &Counter{}
	return c.Inc
}
```

## Walkthrough

`next := NewCounter()`:
- `c` points to `Counter{N: 0}`.
- `c.Inc` is a `func() int` bound to `c`.
- `next()` → `c.N++` → returns 1.
- `next()` → `c.N++` → returns 2.

## Pitfalls

- Using `Counter{}` (value, not pointer) and taking `c.Inc` — Go auto-takes the
  address, but `c` must be addressable. Using `&Counter{}` is explicit and safe.
- Creating `Counter{}` (value) and returning an anonymous closure
  `func() int { c.N++; return c.N }` also works but misses the point of method
  values.
