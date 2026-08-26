# Value Receivers

## Intuition

A value receiver `(r Rect)` means the method gets its own *copy* of the struct.
Any changes to `r` inside the method are invisible to the caller. Use a value
receiver when the method only needs to *read* the fields.

## Approach

1. Multiply `r.Width` by `r.Height`.
2. Return the result.

## Solution

```go
func (r Rect) Area() float64 {
	return r.Width * r.Height
}
```

## Walkthrough

For `Rect{3, 4}`:
- `r.Width` = 3, `r.Height` = 4.
- `3 * 4` = 12.

## Pitfalls

- Using a pointer receiver `*Rect` works but is unnecessary overhead for a
  small, read-only struct.
- Forgetting that `float64` multiplication handles zero correctly — no special
  case needed.
