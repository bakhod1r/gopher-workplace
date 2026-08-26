# Pointer Receivers for Mutation

## Intuition

Translating a point changes it — so we need a pointer receiver. Without `*`,
the `+=` operations would modify a temporary copy and the caller's point would
stay put.

## Approach

1. Add `dx` to `p.X`.
2. Add `dy` to `p.Y`.

## Solution

```go
func (p *Point) Translate(dx, dy float64) {
	p.X += dx
	p.Y += dy
}
```

## Walkthrough

For `p := Point{1, 2}; p.Translate(3, 4)`:
- `p.X += 3` → 4.
- `p.Y += 4` → 6.
- `p` is now `{4, 6}`.

## Pitfalls

- Value receiver `(p Point)` would compile but callers would see no change.
- Forgetting to add `dy` to `Y` — easy typo `p.Y += dx`.
