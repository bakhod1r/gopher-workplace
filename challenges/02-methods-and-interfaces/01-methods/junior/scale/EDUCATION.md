# Pointer Receivers

## Intuition

A **pointer receiver** `(v *Vector)` gives the method access to the *original*
struct, not a copy. Any changes to `v.X` or `v.Y` persist after the method
returns. Use a pointer receiver when the method needs to **mutate** the receiver.

## Approach

1. Use `*Vector` as the receiver type.
2. Multiply `v.X` and `v.Y` by `factor`.

## Solution

```go
func (v *Vector) Scale(factor float64) {
	v.X *= factor
	v.Y *= factor
}
```

## Walkthrough

For `v := Vector{3, 4}; v.Scale(2)`:
- `v.X *= 2` → `v.X` = 6.
- `v.Y *= 2` → `v.Y` = 8.
- `v` is now `{6, 8}`.

## Pitfalls

- Using `(v Vector)` (value receiver) compiles and runs, but the caller's
  struct stays unchanged — the test will fail.
- Go auto-dereferences: `v.X` on a `*Vector` works without `(*v).X`.
