# Value-receiver methods

## The idea

A method with a **value receiver** operates on a copy, so it can't mutate the
caller — ideal for pure computations that return new values:

```go
func (r Rect) Area() int      { return r.W * r.H }
func (r Rect) Scale(f int) Rect { return Rect{r.W*f, r.H*f} }
```

## Why it matters

Small, immutable value types (points, sizes, money) are clearer with value
receivers: no aliasing, safe to copy, comparable with `==`.

## Watch out

- A value receiver copies the struct; mutations inside are lost. Use a pointer
  receiver to mutate.
- Structs of comparable fields support `==` and can be map keys.
- Returning a new struct keeps the original untouched (no surprise mutation).
